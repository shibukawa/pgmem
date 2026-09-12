package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommitTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v380 int32
	_ = v380
	var v381 int64
	_ = v381
	var v382 int32
	_ = v382
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v415 int64
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int64
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int64
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v470 int64
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v628 int32
	_ = v628
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int64
	_ = v655
	var v656 int64
	_ = v656
	var v661 int32
	_ = v661
	var v663 float64
	_ = v663
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int64
	_ = v674
	var v675 int64
	_ = v675
	var v683 int64
	_ = v683
	var v685 int32
	_ = v685
	var v686 int64
	_ = v686
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v707 int64
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v722 int32
	_ = v722
	var v731 int64
	_ = v731
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int64
	_ = v738
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int64
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int64
	_ = v758
	var v760 int32
	_ = v760
	var v776 int32
	_ = v776
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v846 int64
	_ = v846
	var v847 int64
	_ = v847
	var v848 int64
	_ = v848
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int64
	_ = v859
	var v860 int64
	_ = v860
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v894 int32
	_ = v894
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v974 int64
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int64
	_ = v982
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int64
	_ = v993
	var v994 int64
	_ = v994
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int64
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1180 int64
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1207 int64
	_ = v1207
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int64
	_ = v1219
	var v1220 int64
	_ = v1220
	var v1228 int64
	_ = v1228
	var v1230 int64
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int64
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1246 int64
	_ = v1246
	var v1248 int64
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int64
	_ = v1252
	var v1257 int64
	_ = v1257
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int64
	_ = v1269
	var v1270 int64
	_ = v1270
	var v1278 int64
	_ = v1278
	var v1280 int64
	_ = v1280
	var v1283 int64
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1297 int64
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1302 int64
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	var v1315 int64
	_ = v1315
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1339 int64
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1348 int32
	_ = v1348
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1363 int64
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1465 int64
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1473 int32
	_ = v1473
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1535 int32
	_ = v1535
	var v1540 int32
	_ = v1540
	var v1546 int32
	_ = v1546
	var v1554 int32
	_ = v1554
	var v1555 int64
	_ = v1555
	var v1560 int32
	_ = v1560
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1690 int32
	_ = v1690
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1741 int32
	_ = v1741
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1758 int32
	_ = v1758
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1916 int32
	_ = v1916
	var v1922 int32
	_ = v1922
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1968 int32
	_ = v1968
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1984 int32
	_ = v1984
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2016 int32
	_ = v2016
	var v2022 int32
	_ = v2022
	var v2027 int32
	_ = v2027
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2095 int32
	_ = v2095
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2129 int32
	_ = v2129
	var v2133 int32
	_ = v2133
	var v2139 int32
	_ = v2139
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2149 int32
	_ = v2149
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2170 int64
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int64
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2183 int32
	_ = v2183
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2239 int32
	_ = v2239
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2260 int32
	_ = v2260
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2296 int32
	_ = v2296
	var v2298 int32
	_ = v2298
	var v2326 int32
	_ = v2326
	var v2328 int32
	_ = v2328
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2373 int32
	_ = v2373
	var v2375 int32
	_ = v2375
	var v2382 int32
	_ = v2382
	var v2385 int32
	_ = v2385
	var v2387 int32
	_ = v2387
	var v2390 int32
	_ = v2390
	var v2392 int32
	_ = v2392
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2429 int32
	_ = v2429
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2435 int32
	_ = v2435
	var v2446 int32
	_ = v2446
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2479 int64
	_ = v2479
	var v2495 int32
	_ = v2495
	var v2497 int32
	_ = v2497
	v1 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(48)
	m.G0 = v29
	v32 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v33 == int32(5) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v36 + int32(1)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v40 = int32(10)
	goto L6
L4:
	;
	if v74 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	goto L4
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	goto L9
L7:
	;
	v59 = int32(0)
	goto L15
L9:
	;
	goto L10
L10:
	;
	goto L12
L12:
	;
	if v47 == int32(15) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	if v47 <= v40 {
		v74 = int32(1)
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L7
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	if v63 != int32(2) {
		v74 = v59
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _consts[147])))
	if v67 != 0 {
		v74 = v59
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	v74 = int32(0) | base.B2i32(v71 <= v40)
	goto L5
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	F_ShowTransactionStateRec(m, int32(268638), v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v81 == int32(2) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	return
L22:
	;
	goto L20
L23:
	;
	goto L32
L24:
	;
	v86 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	if v86 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if base.Ui32(v90) <= base.Ui32(int32(5)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32))+uint32(_consts[149])))
	v99 = v97
	goto L29
L28:
	;
	v99 = int32(565099)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v99
	F_errmsg_internal(m, int32(367039), v29+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L21
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(512898), int32(2247), int32(268638))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	F_AfterTriggerFireDeferred(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L21
	} else {
		goto L34
	}
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	if v144 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v141 = F_PreCommit_Portals(m, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	if v141 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v146 = int32(5)
	if v33 == v146 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	F_AtEOXact_Parallel(m, int32(1))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L21
	} else {
		goto L47
	}
L40:
	;
	v149 = int32(6)
	goto L42
L41:
	;
	v149 = v146
	goto L42
L42:
	;
	v152 = v144
	goto L43
L43:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	m.T0[v178].(func(*base.Module, int32, int32))(m, v149, v177)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L21
	} else {
		goto L45
	}
L44:
	;
	goto L39
L45:
	;
	if v176 != 0 {
		v152 = v176
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v33 == int32(5) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	F_AfterTriggerEndXact(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L21
	} else {
		goto L61
	}
L49:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v235
	F_errmsg_internal(m, v234, v29)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L21
	} else {
		goto L59
	}
L50:
	;
	if v210 == int32(1) {
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v210 == int32(0) {
		goto L48
	} else {
		goto L56
	}
L53:
	;
	v217 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L21
	} else {
		goto L54
	}
L54:
	;
	if v217 == int32(0) {
		goto L48
	} else {
		goto L55
	}
L55:
	;
	v233 = int32(2295)
	v234 = int32(267411)
	goto L49
L56:
	;
	v227 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	if v227 == int32(0) {
		goto L48
	} else {
		goto L58
	}
L58:
	;
	v233 = int32(2301)
	v234 = int32(267785)
	goto L49
L59:
	;
	F_errfinish(m, int32(512898), v233, int32(268638))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L21
	} else {
		goto L60
	}
L60:
	;
	goto L48
L61:
	;
	F_PreCommit_on_commit_actions(m)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L21
	} else {
		goto L62
	}
L62:
	;
	v251 = base.B2i32(v33 == int32(5))
	F_smgrDoPendingSyncs(m, int32(1), v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L21
	} else {
		goto L63
	}
L63:
	;
	F_AtEOXact_LargeObject(m, int32(1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L21
	} else {
		goto L64
	}
L64:
	;
	v257 = m.G0
	v259 = v257 - int32(8144)
	m.G0 = v259
	v262 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	v264 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	if v262|v264 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v251 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L21
	} else {
		goto L210
	}
L67:
	;
	m.G0 = v259 + int32(8144)
	goto L65
L68:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, _consts[153])))
	if v269 != int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	if v288 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v274 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L21
	} else {
		goto L71
	}
L71:
	;
	if v274 == int32(0) {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	F_errmsg_internal(m, int32(21031), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L21
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(520690), int32(868), int32(21031))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	goto L69
L75:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	if v600 == int32(0) {
		goto L67
	} else {
		goto L119
	}
L76:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v291 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v294 <= int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v316 = v1
	goto L79
L79:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323+v316<<(uint(int32(2))%32))))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	if v328 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L75
L81:
	;
	v570 = v316 + int32(1)
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v570 < v571 {
		v316 = v570
		goto L79
	} else {
		goto L118
	}
L82:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, _consts[154])))
	if v330 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, _consts[153])))
	if v332 != int32(1) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, _consts[155])))
	if v355 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v337 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L21
	} else {
		goto L86
	}
L86:
	;
	if v337 == int32(0) {
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	*(*int32)(unsafe.Add(mBase, uint32(v259)+32)) = v342
	F_errmsg_internal(m, int32(704146), v259+int32(32))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L21
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(520690), int32(1054), int32(106926))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L21
	} else {
		goto L89
	}
L89:
	;
	goto L84
L90:
	;
	F_before_shmem_exit(m, int32(514), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L21
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v370 = F_LWLockAcquire(m, v366+int32(3456), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L21
	} else {
		goto L94
	}
L93:
	;
	v363 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[155])) = uint8(v363)
	goto L92
L94:
	;
	v372 = int32(-1)
	v374 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v376 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+28))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v376)+24))
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v376)+16))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v376)+8))
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v376)))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v376)+40))
	if v382 != v372 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v393 = v382
	v401 = v372
	v402 = v378
	v410 = v377
	v415 = v379
	goto L98
L96:
	;
	v456 = v372
	v457 = v378
	v465 = v377
	v470 = v379
	goto L97
L97:
	;
	v474 = int32(5)
	v476 = v376 + v374<<(uint(v474)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v476)+84)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v476)+80)) = v457
	*(*int64)(unsafe.Add(mBase, uint32(v476)+72)) = v470
	v481 = v376 + int32(56)
	v482 = int32(4155456)
	v483 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v488 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	*(*int32)(unsafe.Add(mBase, uint32(v481+v483<<(uint(v474)%32)))) = v488
	v491 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v496 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v491<<(uint(v474)%32))+60)) = v496
	if v456 != int32(-1) {
		goto L112
	} else {
		goto L113
	}
L98:
	;
	v420 = v393 << (uint(int32(5)) % 32)
	v421 = v376 + int32(56) + v420
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	if v422 != v392 {
		v435 = v420
		v437 = v402
		v439 = v410
		v440 = v415
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v456 = v443
	v457 = v437
	v465 = v439
	v470 = v440
	goto L97
L100:
	;
	if v393 < v374 {
		goto L107
	} else {
		goto L108
	}
L101:
	;
	v425 = v393 << (uint(int32(5)) % 32)
	v426 = v376 + int32(72) + v425
	v427 = *(*int64)(unsafe.Add(mBase, uint32(v421)+16))
	if v415 < v427 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v426)+12))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
	v435 = v425
	v437 = v430
	v439 = v429
	v440 = v427
	goto L100
L103:
	;
	goto L104
L104:
	;
	if v415 != v427 {
		v435 = v425
		v437 = v402
		v439 = v410
		v440 = v415
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
	if v432 < v402 {
		v435 = v425
		v437 = v402
		v439 = v410
		v440 = v415
		goto L100
	} else {
		goto L106
	}
L106:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v426)+12))
	v435 = v425
	v437 = v432
	v439 = v434
	v440 = v415
	goto L100
L107:
	;
	v443 = v393
	goto L109
L108:
	;
	v443 = v401
	goto L109
L109:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v435+(v376-int32(-64)))))
	if v445 != int32(-1) {
		v393 = v445
		v401 = v443
		v402 = v437
		v410 = v439
		v415 = v440
		goto L98
	} else {
		goto L110
	}
L110:
	;
	goto L99
L111:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v530+int32(3456))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L21
	} else {
		goto L115
	}
L112:
	;
	v500 = int32(4155456)
	v501 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v502 = int32(5)
	v506 = v456 << (uint(v502) % 32)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v481+v506)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v481+v501<<(uint(v502)%32))+8)) = v508
	v514 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	*(*int32)(unsafe.Add(mBase, uint32(v506+v376-int32(-64)))) = v514
	goto L111
L113:
	;
	goto L114
L114:
	;
	v516 = int32(4155456)
	v517 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v376)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v517<<(uint(int32(5))%32)-int32(-64)))) = v523
	v526 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+40)) = v526
	goto L111
L115:
	;
	v536 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[154])) = uint8(v536)
	if base.B2i32(v457 == v380)&base.B2i32(v470 == v381) != 0 {
		goto L81
	} else {
		goto L116
	}
L116:
	;
	F_asyncQueueReadAllNotifications(m)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L21
	} else {
		goto L117
	}
L117:
	;
	goto L81
L118:
	;
	goto L80
L119:
	;
	v603 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L21
	} else {
		goto L120
	}
L120:
	;
	F_LockSharedObject(m, int32(1262), int32(0), int32(8))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L21
	} else {
		goto L121
	}
L121:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	if v612 == int32(0) {
		goto L67
	} else {
		goto L122
	}
L122:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v612)+12))
	if v615 == int32(0) {
		goto L67
	} else {
		goto L123
	}
L123:
	;
	v628 = v615
	goto L124
L124:
	;
	v647 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v651 = F_LWLockAcquire(m, v647+int32(3456), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L21
	} else {
		goto L126
	}
L125:
	;
	goto L67
L126:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	v655 = *(*int64)(unsafe.Add(mBase, uint32(v654)))
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v654)+16))
	if v655 == v656 {
		v820 = v654
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v846 = int64(*(*int32)(unsafe.Add(mBase, _consts[159])))
	v847 = *(*int64)(unsafe.Add(mBase, uint32(v820)))
	v848 = *(*int64)(unsafe.Add(mBase, uint32(v820)+16))
	if v846 <= v847-v848 {
		goto L66
	} else {
		goto L163
	}
L128:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	v663 = base.F64_div(base.F64_convert_i64_s(v655-v656), base.F64_convert_i32_s(v661))
	if base.F64_lt(v663, float64(0.5)) != 0 {
		v820 = v654
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v669 = m.G0
	v670 = int32(16)
	v671 = v669 - v670
	m.G0 = v671
	F___gettimeofday(m, v671)
	mBase = m.M
	v674 = *(*int64)(unsafe.Add(mBase, uint32(v671)))
	v675 = int64(*(*int32)(unsafe.Add(mBase, uint32(v671)+8)))
	m.G0 = v671 + v670
	v683 = v675 + v674*int64(1000000) - int64(946684800000000)
	goto L130
L130:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	v686 = *(*int64)(unsafe.Add(mBase, uint32(v685)+48))
	goto L131
L131:
	;
	v694 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	if base.B2i32(base.I64_extend_i32_s(int32(5000))*int64(1000) <= v683-v686) == int32(0) {
		v820 = v694
		goto L127
	} else {
		goto L132
	}
L132:
	;
	v697 = int32(-1)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v694)+40))
	if v698 != v697 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v704 = v694 + int32(56)
	v707 = *(*int64)(unsafe.Add(mBase, uint32(v694)))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v694)+8))
	v709 = v698
	v714 = v708
	v722 = v697
	v731 = v707
	goto L136
L134:
	;
	v776 = v697
	goto L135
L135:
	;
	v791 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L21
	} else {
		goto L152
	}
L136:
	;
	v736 = v709 << (uint(int32(5)) % 32)
	v737 = v694 + int32(72) + v736
	v738 = *(*int64)(unsafe.Add(mBase, uint32(v737)))
	if v738 <= v731 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v776 = v757
	goto L135
L138:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v736+(v694-int32(-64)))))
	if v760 != int32(-1) {
		v709 = v760
		v714 = v756
		v722 = v757
		v731 = v758
		goto L136
	} else {
		goto L151
	}
L139:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v736+v704)))
	v756 = v750
	v757 = v754
	v758 = v752
	goto L138
L140:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v737)+8))
	if v731 != v738 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v745 = v714
	goto L142
L142:
	;
	if v731 != v738 {
		v756 = v745
		v757 = v722
		v758 = v731
		goto L138
	} else {
		goto L149
	}
L143:
	;
	v750 = v740
	v752 = v738
	goto L139
L144:
	;
	goto L145
L145:
	;
	if v714 < v740 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v743 = v714
	goto L148
L147:
	;
	v743 = v740
	goto L148
L148:
	;
	v745 = v743
	goto L142
L149:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v736+v704)+24))
	if v745 != v748 {
		v756 = v745
		v757 = v722
		v758 = v731
		goto L138
	} else {
		goto L150
	}
L150:
	;
	v750 = v748
	v752 = v731
	goto L139
L151:
	;
	goto L137
L152:
	;
	if v791 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v259)+16)) = base.F64_mul(v663, float64(100))
	F_errmsg(m, int32(315842), v259+int32(16))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L21
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v817 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int64)(unsafe.Add(mBase, uint32(v817)+48)) = v683
	v820 = v817
	goto L127
L156:
	;
	if v776 != int32(-1) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v776
	F_errdetail(m, int32(613723), v259)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L21
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	F_errfinish(m, int32(520690), int32(1569), int32(348836))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L21
	} else {
		goto L162
	}
L160:
	;
	F_errhint(m, int32(636851), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L21
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	goto L155
L163:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v259)+56)) = v847
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v820)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v259)+48)) = v852
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v820)+12))
	v856 = *(*int32)(unsafe.Add(mBase, _consts[160]))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v856)+28))
	v859 = int64(*(*uint16)(unsafe.Add(mBase, _consts[161])))
	v860 = base.I64_rem_s(v847, v859)
	v864 = v857 + base.I32_wrap_i64(v860)<<(uint(int32(7))%32)
	v866 = F_LWLockAcquire(m, v864, int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L21
	} else {
		goto L164
	}
L164:
	;
	if v847 != int64(0) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v881 = *(*int32)(unsafe.Add(mBase, _consts[160]))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v881)+12))
	v884 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v879+v882))) = uint8(v884)
	v886 = v852
	v894 = v628
	goto L172
L166:
	;
	v877 = F_SimpleLruReadPage(m, int32(4445896), v847, int32(1), int32(0))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L21
	} else {
		goto L170
	}
L167:
	;
	if v852 != 0 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v872 = F_SimpleLruZeroPage(m, int32(4445896), int64(0))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L21
	} else {
		goto L169
	}
L169:
	;
	v879 = v872
	goto L165
L170:
	;
	v879 = v877
	goto L165
L171:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int64)(unsafe.Add(mBase, uint32(v1021))) = v982
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v259)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1021)+12)) = v854
	*(*int32)(unsafe.Add(mBase, uint32(v1021)+8)) = v1023
	F_LWLockRelease(m, v1018)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L21
	} else {
		goto L207
	}
L172:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v894)))
	v913 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v912)+2)))
	v914 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v912))))
	v916 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	*(*int32)(unsafe.Add(mBase, uint32(v259)+68)) = v916
	v918 = v913 + v914
	v922 = (v918 + int32(21)) & int32(262140)
	*(*int32)(unsafe.Add(mBase, uint32(v259)+64)) = v922
	v924 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L21
	} else {
		goto L174
	}
L173:
	;
	v1018 = v864
	v1019 = int32(0)
	goto L171
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+72)) = v924
	v928 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	*(*int32)(unsafe.Add(mBase, uint32(v259)+76)) = v928
	v933 = v918 + int32(2)
	if v933 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	if v886+v922 <= int32(8192) {
		goto L180
	} else {
		goto L181
	}
L176:
	;
	v934 = F__emscripten_memcpy_bulkmem(m, v259+int32(80), v912+int32(4), v933)
	mBase = m.M
	goto L178
L177:
	;
	goto L178
L178:
	;
	goto L175
L179:
	;
	v963 = *(*int32)(unsafe.Add(mBase, _consts[160]))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v963)+4))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v964+v879<<(uint(int32(2))%32))))
	if v959 != 0 {
		goto L187
	} else {
		goto L188
	}
L180:
	;
	v940 = v894 + int32(4)
	v943 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v944)+12))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v944)+4))
	if base.Ui32(v940) < base.Ui32(v945+v946<<(uint(int32(2))%32)) {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	goto L182
L182:
	;
	v952 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v259)+80)) = uint16(v952)
	*(*int64)(unsafe.Add(mBase, uint32(v259)+68)) = int64(0)
	v957 = int32(8192) - v886
	*(*int32)(unsafe.Add(mBase, uint32(v259)+64)) = v957
	v959 = v957
	v961 = v894
	goto L179
L183:
	;
	v951 = v940
	goto L185
L184:
	;
	v951 = int32(0)
	goto L185
L185:
	;
	v959 = v922
	v961 = v951
	goto L179
L186:
	;
	v974 = *(*int64)(unsafe.Add(mBase, uint32(v259)+56))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v259)+48))
	v976 = v975 + v959
	v978 = v976 - int32(8173)
	v980 = base.B2i32(base.Ui32(v978) < base.Ui32(int32(-8193)))
	v982 = v974 + base.I64_extend_i32_u(v980)
	*(*int64)(unsafe.Add(mBase, uint32(v259)+56)) = v982
	if base.Ui32(v978) < base.Ui32(int32(-8193)) {
		goto L190
	} else {
		goto L191
	}
L187:
	;
	v972 = F__emscripten_memcpy_bulkmem(m, v968+v886, v259-int32(-64), v959)
	mBase = m.M
	goto L189
L188:
	;
	goto L189
L189:
	;
	goto L186
L190:
	;
	v985 = int32(0)
	goto L192
L191:
	;
	v985 = v976
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+48)) = v985
	if base.Ui32(v978) <= base.Ui32(int32(-8194)) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v990 = *(*int32)(unsafe.Add(mBase, _consts[160]))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v990)+28))
	v993 = int64(*(*uint16)(unsafe.Add(mBase, _consts[161])))
	v994 = base.I64_rem_s(v982, v993)
	v998 = v991 + base.I32_wrap_i64(v994)<<(uint(int32(7))%32)
	if v864 == v998 {
		goto L197
	} else {
		goto L198
	}
L194:
	;
	goto L195
L195:
	;
	if v961 != 0 {
		v886 = v976
		v894 = v961
		goto L172
	} else {
		goto L206
	}
L196:
	;
	v1007 = F_SimpleLruZeroPage(m, int32(4445896), v982)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L21
	} else {
		goto L202
	}
L197:
	;
	v1005 = v864
	goto L196
L198:
	;
	goto L199
L199:
	;
	F_LWLockRelease(m, v864)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L21
	} else {
		goto L200
	}
L200:
	;
	v1003 = F_LWLockAcquire(m, v998, int32(0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L21
	} else {
		goto L201
	}
L201:
	;
	v1005 = v998
	goto L196
L202:
	;
	if v982&int64(3) == int64(0) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1014 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[162])) = uint8(v1014)
	goto L205
L204:
	;
	goto L205
L205:
	;
	v1018 = v1005
	v1019 = v961
	goto L171
L206:
	;
	goto L173
L207:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v1029+int32(3456))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L21
	} else {
		goto L208
	}
L208:
	;
	if v1019 != 0 {
		v628 = v1019
		goto L124
	} else {
		goto L209
	}
L209:
	;
	goto L125
L210:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L21
	} else {
		goto L211
	}
L211:
	;
	F_errmsg(m, int32(362997), int32(0))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L21
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(520690), int32(945), int32(21031))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L21
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_PreCommit_CheckForSerializationFailure(m)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L21
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1083 = int32(4543676)
	v1085 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	v1086 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v1085 + v1086
	F_AtEOXact_RelationMap(m, v1086, v251)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L21
	} else {
		goto L218
	}
L217:
	;
	goto L216
L218:
	;
	v1092 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+76)) = uint8(v1092)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v1092
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(3)
	v1099 = *(*int32)(unsafe.Add(mBase, _consts[164]))
	if v1092 < v1099 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	F_disable_timeout(m, int32(8))
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L21
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	if v33 != int32(5) {
		goto L232
	} else {
		goto L233
	}
L222:
	;
	goto L221
L223:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	F_ProcArrayEndTransaction(m, v1587, v1575)
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L21
	} else {
		goto L367
	}
L224:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	if v1567 != 0 {
		goto L361
	} else {
		goto L362
	}
L225:
	;
	v1554 = int32(4444576)
	v1555 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	*(*int64)(unsafe.Add(mBase, _consts[166])) = v1555
	*(*int64)(unsafe.Add(mBase, _consts[165])) = int64(0)
	v1560 = v1546
	goto L224
L226:
	;
	v1473 = v1129 - int32(1)
	if v1473 < int32(0) {
		v1540 = v1109
		goto L331
	} else {
		goto L332
	}
L227:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1372)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1372)+120)) = v1373 & int32(-2)
	v1377 = int32(4543684)
	v1379 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v1380 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1379 - v1380
	v1387 = v1129 - v1380
	if v1387 < int32(0) {
		v1454 = v1109
		goto L297
	} else {
		goto L298
	}
L228:
	;
	F_XLogSetAsyncXactLSN(m, v1339)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L21
	} else {
		goto L293
	}
L229:
	;
	F_TransactionIdCommitTree(m, v1109, v1129, v1131)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L21
	} else {
		goto L292
	}
L230:
	;
	v1339 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	v1341 = int32(*(*uint8)(unsafe.Add(mBase, _consts[167])))
	v1342 = int32(0)
	if base.B2i32(v1341 == v1342)&base.B2i32(v1125 <= v1342) != 0 {
		goto L228
	} else {
		goto L289
	}
L231:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L21
	} else {
		goto L286
	}
L232:
	;
	v1107 = int32(0)
	v1109 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1107
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+35)) = uint8(v1107)
	v1117 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if int32(2) <= v1117 {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	goto L234
L234:
	;
	v1302 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	v1304 = *(*int32)(unsafe.Add(mBase, _consts[168]))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+68)) = int32(1)
	v1309 = v1304 + int32(68)
	if v1305 != 0 {
		goto L279
	} else {
		goto L280
	}
L235:
	;
	F_LogLogicalInvalidations(m)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L21
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v1125 = F_smgrGetPendingDeletes(m, int32(1), v29+int32(44))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L21
	} else {
		goto L239
	}
L238:
	;
	goto L237
L239:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+52))
	if v1129 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+48))
	v1131 = v1130
	goto L242
L241:
	;
	v1131 = v1107
	goto L242
L242:
	;
	v1135 = F_pgstat_get_transactional_drops(m, int32(1), v29+int32(40))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L21
	} else {
		goto L243
	}
L243:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	if int32(0) < v1138 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1145 = F_xactGetCommittedInvalidationMessages(m, v29+int32(36), v29+int32(35))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L21
	} else {
		goto L247
	}
L245:
	;
	v1147 = v1
	goto L246
L246:
	;
	v1149 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	if v1109 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v1147 = v1145
	goto L246
L248:
	;
	if v1125|v1135 != 0 {
		goto L231
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1190 = int32(4543684)
	v1192 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v1193 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1192 + v1193
	v1197 = int32(*(*uint16)(unsafe.Add(mBase, _consts[169])))
	v1199 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1199)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1199)+120)) = v1200 | v1193
	v1207 = *(*int64)(unsafe.Add(mBase, _consts[170]))
	if v1207 == int64(0) {
		goto L260
	} else {
		goto L261
	}
L251:
	;
	if v1147 != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+35)))
	v1155 = m.G0
	v1157 = v1155 - int32(16)
	m.G0 = v1157
	*(*int32)(unsafe.Add(mBase, uint32(v1157)+8)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1157)+8)) = uint8(v1154)
	v1163 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	*(*int32)(unsafe.Add(mBase, uint32(v1157))) = v1163
	v1166 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, uint32(v1157)+4)) = v1166
	*(*int32)(unsafe.Add(mBase, uint32(v1157)+12)) = v1147
	F_XLogBeginInsert(m)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L21
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	if v1149 != int64(0) {
		v1336 = int32(1)
		goto L230
	} else {
		goto L259
	}
L255:
	;
	F_XLogRegisterData(m, v1157, int32(16))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L21
	} else {
		goto L256
	}
L256:
	;
	F_XLogRegisterData(m, v1153, v1147<<(uint(int32(4))%32))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L21
	} else {
		goto L257
	}
L257:
	;
	v1180 = F_XLogInsert(m, int32(8), int32(32))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L21
	} else {
		goto L258
	}
L258:
	;
	m.G0 = v1157 + int32(16)
	v1336 = int32(1)
	goto L230
L259:
	;
	v1560 = int32(0)
	goto L224
L260:
	;
	v1214 = m.G0
	v1215 = int32(16)
	v1216 = v1214 - v1215
	m.G0 = v1216
	F___gettimeofday(m, v1216)
	mBase = m.M
	v1219 = *(*int64)(unsafe.Add(mBase, uint32(v1216)))
	v1220 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1216)+8)))
	m.G0 = v1216 + v1215
	v1228 = v1220 + v1219*int64(1000000) - int64(946684800000000)
	goto L263
L261:
	;
	v1230 = v1207
	goto L262
L262:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+35)))
	v1236 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	v1237 = int32(0)
	v1239 = F_XactLogCommitRecord(m, v1230, v1129, v1131, v1125, v1231, v1135, v1232, v1147, v1233, v1234, v1236, v1237, v1237)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L21
	} else {
		goto L264
	}
L263:
	;
	*(*int64)(unsafe.Add(mBase, _consts[170])) = v1228
	v1230 = v1228
	goto L262
L264:
	;
	if base.Ui32(int32(2)) <= base.Ui32((v1197+v1193)&int32(65535)) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v1286 = int32(*(*uint16)(unsafe.Add(mBase, _consts[169])))
	F_TransactionTreeSetCommitTsData(m, v1109, v1129, v1131, v1283, v1286)
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L21
	} else {
		goto L275
	}
L266:
	;
	v1246 = *(*int64)(unsafe.Add(mBase, _consts[172]))
	v1248 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	F_replorigin_session_advance(m, v1246, v1248)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L21
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v1257 = *(*int64)(unsafe.Add(mBase, _consts[170]))
	if v1257 == int64(0) {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	v1252 = *(*int64)(unsafe.Add(mBase, _consts[173]))
	if v1252 != int64(0) {
		v1283 = v1252
		goto L265
	} else {
		goto L270
	}
L270:
	;
	goto L268
L271:
	;
	v1264 = m.G0
	v1265 = int32(16)
	v1266 = v1264 - v1265
	m.G0 = v1266
	F___gettimeofday(m, v1266)
	mBase = m.M
	v1269 = *(*int64)(unsafe.Add(mBase, uint32(v1266)))
	v1270 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1266)+8)))
	m.G0 = v1266 + v1265
	v1278 = v1270 + v1269*int64(1000000) - int64(946684800000000)
	goto L274
L272:
	;
	v1280 = v1257
	goto L273
L273:
	;
	*(*int64)(unsafe.Add(mBase, _consts[173])) = v1280
	v1283 = v1280
	goto L265
L274:
	;
	*(*int64)(unsafe.Add(mBase, _consts[170])) = v1278
	v1280 = v1278
	goto L273
L275:
	;
	if v1149 == int64(0) {
		v1336 = int32(0)
		goto L230
	} else {
		goto L276
	}
L276:
	;
	v1291 = int32(1)
	v1293 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	if v1293 <= int32(0) {
		v1336 = v1291
		goto L230
	} else {
		goto L277
	}
L277:
	;
	v1297 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	F_XLogFlush(m, v1297)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L21
	} else {
		goto L278
	}
L278:
	;
	v1354 = v1291
	goto L229
L279:
	;
	F_s_lock(m, v1309, int32(517893), int32(1598), int32(446375))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L21
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(v1304)+72))
	if base.Ui64(v1315) < base.Ui64(v1302) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	goto L281
L283:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1304)+72)) = v1302
	goto L285
L284:
	;
	goto L285
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1309))) = int32(0)
	v1575 = int32(0)
	goto L223
L286:
	;
	F_errmsg_internal(m, int32(449843), int32(0))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L21
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(512898), int32(1365), int32(106859))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L21
	} else {
		goto L288
	}
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L289:
	;
	F_XLogFlush(m, v1339)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L21
	} else {
		goto L290
	}
L290:
	;
	if v1109 == int32(0) {
		goto L226
	} else {
		goto L291
	}
L291:
	;
	v1354 = v1336
	goto L229
L292:
	;
	v1369 = v1354
	goto L227
L293:
	;
	if v1109 == int32(0) {
		goto L226
	} else {
		goto L294
	}
L294:
	;
	v1363 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	F_TransactionIdAsyncCommitTree(m, v1109, v1129, v1131, v1363)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L21
	} else {
		goto L295
	}
L295:
	;
	v1369 = v1336
	goto L227
L296:
	;
	if v1369 == int32(0) {
		v1546 = v1454
		goto L225
	} else {
		goto L327
	}
L297:
	;
	goto L296
L298:
	;
	if v1129&int32(1) != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1392 = int32(2)
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1131+v1387<<(uint(v1392)%32))))
	if base.B2i32(base.Ui32(v1392) < base.Ui32(v1395))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1109)) == int32(0) {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	v1410 = v1109
	v1412 = v1387
	goto L301
L301:
	;
	if v1387 == int32(0) {
		v1454 = v1410
		goto L297
	} else {
		goto L309
	}
L302:
	;
	v1410 = v1407
	v1412 = v1129 - int32(2)
	goto L301
L303:
	;
	v1407 = v1395
	goto L302
L304:
	;
	if base.Ui32(v1109) < base.Ui32(v1395) {
		goto L303
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	if int32(0) <= v1109-v1395 {
		v1407 = v1109
		goto L302
	} else {
		goto L308
	}
L307:
	;
	v1407 = v1109
	goto L302
L308:
	;
	goto L303
L309:
	;
	v1417 = v1410
	v1418 = v1412
	goto L310
L310:
	;
	v1424 = v1418 << (uint(int32(2)) % 32)
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1131+v1424)))
	if base.Ui32(v1417) < base.Ui32(int32(3)) {
		goto L314
	} else {
		goto L315
	}
L311:
	;
	v1454 = v1449
	goto L297
L312:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1424+(v1131-int32(4)))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1437))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1435)) == int32(0) {
		goto L321
	} else {
		goto L322
	}
L313:
	;
	v1435 = v1426
	goto L312
L314:
	;
	if base.Ui32(v1426) <= base.Ui32(v1417) {
		v1435 = v1417
		goto L312
	} else {
		goto L318
	}
L315:
	;
	if base.Ui32(v1426) < base.Ui32(int32(3)) {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	if v1417-v1426 < int32(0) {
		goto L313
	} else {
		goto L317
	}
L317:
	;
	v1435 = v1417
	goto L312
L318:
	;
	goto L313
L319:
	;
	if int32(1) < v1418 {
		v1417 = v1449
		v1418 = v1418 - int32(2)
		goto L310
	} else {
		goto L326
	}
L320:
	;
	v1449 = v1437
	goto L319
L321:
	;
	if base.Ui32(v1435) < base.Ui32(v1437) {
		goto L320
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	if int32(0) <= v1435-v1437 {
		v1449 = v1435
		goto L319
	} else {
		goto L325
	}
L324:
	;
	v1449 = v1435
	goto L319
L325:
	;
	goto L320
L326:
	;
	goto L311
L327:
	;
	if v1109 == int32(0) {
		v1546 = v1454
		goto L225
	} else {
		goto L328
	}
L328:
	;
	v1465 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	F_SyncRepWaitForLSN(m, v1465, int32(1))
	mBase = m.M
	v1468 = m.ExcPending
	if v1468 != 0 {
		goto L21
	} else {
		goto L329
	}
L329:
	;
	v1546 = v1454
	goto L225
L330:
	;
	v1546 = v1540
	goto L225
L331:
	;
	goto L330
L332:
	;
	if v1129&int32(1) != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v1478 = int32(2)
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1131+v1473<<(uint(v1478)%32))))
	if base.B2i32(base.Ui32(v1478) < base.Ui32(v1481))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1109)) == int32(0) {
		goto L338
	} else {
		goto L339
	}
L334:
	;
	v1496 = v1109
	v1498 = v1473
	goto L335
L335:
	;
	if v1473 == int32(0) {
		v1540 = v1496
		goto L331
	} else {
		goto L343
	}
L336:
	;
	v1496 = v1493
	v1498 = v1129 - int32(2)
	goto L335
L337:
	;
	v1493 = v1481
	goto L336
L338:
	;
	if base.Ui32(v1109) < base.Ui32(v1481) {
		goto L337
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	if int32(0) <= v1109-v1481 {
		v1493 = v1109
		goto L336
	} else {
		goto L342
	}
L341:
	;
	v1493 = v1109
	goto L336
L342:
	;
	goto L337
L343:
	;
	v1503 = v1496
	v1504 = v1498
	goto L344
L344:
	;
	v1510 = v1504 << (uint(int32(2)) % 32)
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1131+v1510)))
	if base.Ui32(v1503) < base.Ui32(int32(3)) {
		goto L348
	} else {
		goto L349
	}
L345:
	;
	v1540 = v1535
	goto L331
L346:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1510+(v1131-int32(4)))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1523))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1521)) == int32(0) {
		goto L355
	} else {
		goto L356
	}
L347:
	;
	v1521 = v1512
	goto L346
L348:
	;
	if base.Ui32(v1512) <= base.Ui32(v1503) {
		v1521 = v1503
		goto L346
	} else {
		goto L352
	}
L349:
	;
	if base.Ui32(v1512) < base.Ui32(int32(3)) {
		goto L348
	} else {
		goto L350
	}
L350:
	;
	if v1503-v1512 < int32(0) {
		goto L347
	} else {
		goto L351
	}
L351:
	;
	v1521 = v1503
	goto L346
L352:
	;
	goto L347
L353:
	;
	if int32(1) < v1504 {
		v1503 = v1535
		v1504 = v1504 - int32(2)
		goto L344
	} else {
		goto L360
	}
L354:
	;
	v1535 = v1523
	goto L353
L355:
	;
	if base.Ui32(v1521) < base.Ui32(v1523) {
		goto L354
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	if int32(0) <= v1521-v1523 {
		v1535 = v1521
		goto L353
	} else {
		goto L359
	}
L358:
	;
	v1535 = v1521
	goto L353
L359:
	;
	goto L354
L360:
	;
	goto L345
L361:
	;
	F_pfree(m, v1567)
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L21
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	if v1135 == int32(0) {
		v1575 = v1560
		goto L223
	} else {
		goto L365
	}
L364:
	;
	goto L363
L365:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	F_pfree(m, v1572)
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L21
	} else {
		goto L366
	}
L366:
	;
	v1575 = v1560
	goto L223
L367:
	;
	v1591 = base.B2i32(v33 == int32(5))
	v1593 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	if v1593 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1596 = v1593
	goto L371
L369:
	;
	goto L370
L370:
	;
	*(*int32)(unsafe.Add(mBase, _consts[175])) = int32(0)
	v1655 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v1656 = int32(1)
	F_ResourceOwnerRelease(m, v1655, v1656, v1656, v1656)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L21
	} else {
		goto L375
	}
L371:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1596)))
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1596)+8))
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1596)+4))
	m.T0[v1622].(func(*base.Module, int32, int32))(m, v1591, v1621)
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L21
	} else {
		goto L373
	}
L372:
	;
	goto L370
L373:
	;
	if v1620 != 0 {
		v1596 = v1620
		goto L371
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	F_AtEOXact_Aio(m)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L21
	} else {
		goto L376
	}
L376:
	;
	F_AtEOXact_RelationCache(m, int32(1))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L21
	} else {
		goto L377
	}
L377:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L21
	} else {
		goto L378
	}
L378:
	;
	F_AtEOXact_Inval(m, int32(1))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L21
	} else {
		goto L379
	}
L379:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v1673 = int32(4155456)
	v1674 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v1675 = int32(2)
	v1678 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1672+v1674<<(uint(v1675)%32)))) = v1678
	v1681 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	v1683 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	*(*int32)(unsafe.Add(mBase, uint32(v1681+v1683<<(uint(v1675)%32)))) = v1678
	v1690 = int32(4146096)
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v1690
	*(*int32)(unsafe.Add(mBase, _consts[178])) = v1690
	*(*int32)(unsafe.Add(mBase, _consts[179])) = v1678
	*(*int32)(unsafe.Add(mBase, _consts[180])) = v1678
	goto L380
L380:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v1704 = int32(1)
	F_ResourceOwnerRelease(m, v1702, int32(2), v1704, v1704)
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L21
	} else {
		goto L381
	}
L381:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v1711 = int32(1)
	F_ResourceOwnerRelease(m, v1709, int32(3), v1711, v1711)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L21
	} else {
		goto L382
	}
L382:
	;
	F_smgrDoPendingDeletes(m, int32(1))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L21
	} else {
		goto L383
	}
L383:
	;
	v1719 = m.G0
	v1721 = v1719 - int32(48)
	m.G0 = v1721
	v1724 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	v1726 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	if v1724|v1726 != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1729 = int32(*(*uint8)(unsafe.Add(mBase, _consts[153])))
	if v1729 != int32(1) {
		goto L387
	} else {
		goto L388
	}
L385:
	;
	goto L386
L386:
	;
	m.G0 = v1721 + int32(48)
	v2367 = int32(1)
	F_AtEOXact_GUC(m, v2367, v2367)
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L21
	} else {
		goto L499
	}
L387:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	if v1748 == int32(0) {
		goto L393
	} else {
		goto L394
	}
L388:
	;
	v1734 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L21
	} else {
		goto L389
	}
L389:
	;
	if v1734 == int32(0) {
		goto L387
	} else {
		goto L390
	}
L390:
	;
	F_errmsg_internal(m, int32(21015), int32(0))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L21
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(520690), int32(979), int32(21015))
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L21
	} else {
		goto L392
	}
L392:
	;
	goto L387
L393:
	;
	v2095 = int32(*(*uint8)(unsafe.Add(mBase, _consts[154])))
	if v2095 == int32(0) {
		goto L452
	} else {
		goto L453
	}
L394:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+4))
	if v1751 == int32(0) {
		goto L393
	} else {
		goto L395
	}
L395:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+4))
	if v1754 <= int32(0) {
		goto L393
	} else {
		goto L396
	}
L396:
	;
	v1758 = int32(0)
	goto L397
L397:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+12))
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1783+v1758<<(uint(int32(2))%32))))
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1787)))
	switch v1788 {
	case 0:
		goto L402
	case 1:
		goto L401
	case 2:
		goto L400
	default:
		goto L399
	}
L398:
	;
	goto L393
L399:
	;
	v2065 = v1758 + int32(1)
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+4))
	if v2065 < v2066 {
		v1758 = v2065
		goto L397
	} else {
		goto L451
	}
L400:
	;
	v2006 = int32(*(*uint8)(unsafe.Add(mBase, _consts[153])))
	if v2006 != int32(1) {
		goto L444
	} else {
		goto L445
	}
L401:
	;
	v1903 = v1787 + int32(4)
	v1905 = int32(*(*uint8)(unsafe.Add(mBase, _consts[153])))
	if v1905 != int32(1) {
		goto L420
	} else {
		goto L421
	}
L402:
	;
	v1790 = v1787 + int32(4)
	m.Env.Pgmem_listen(m, v1790, int32(1))
	mBase = m.M
	v1794 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v1794 == int32(0) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1888 = int32(4549024)
	v1889 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1892 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1892
	v1894 = F_pstrdup(m, v1790)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L21
	} else {
		goto L418
	}
L404:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+4))
	if v1797 <= int32(0) {
		goto L403
	} else {
		goto L405
	}
L405:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1794)+12))
	v1802 = int32(0)
	goto L406
L406:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1800+v1802<<(uint(int32(2))%32))))
	v1834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1790))))
	v1835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1831))))
	if v1835 == int32(0) {
		v1854 = v1834
		v1855 = v1835
		goto L409
	} else {
		goto L410
	}
L407:
	;
	goto L403
L408:
	;
	if v1855-v1854 == int32(0) {
		goto L399
	} else {
		goto L416
	}
L409:
	;
	goto L408
L410:
	;
	if v1834 != v1835 {
		v1854 = v1834
		v1855 = v1835
		goto L409
	} else {
		goto L411
	}
L411:
	;
	v1839 = v1831
	v1840 = v1790
	goto L412
L412:
	;
	v1843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1840)+1)))
	v1844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1839)+1)))
	if v1844 == int32(0) {
		v1854 = v1843
		v1855 = v1844
		goto L409
	} else {
		goto L414
	}
L413:
	;
	v1854 = v1843
	v1855 = v1844
	goto L409
L414:
	;
	v1847 = int32(1)
	if v1843 == v1844 {
		v1839 = v1839 + v1847
		v1840 = v1840 + v1847
		goto L412
	} else {
		goto L415
	}
L415:
	;
	goto L413
L416:
	;
	v1860 = v1802 + int32(1)
	if v1797 != v1860 {
		v1802 = v1860
		goto L406
	} else {
		goto L417
	}
L417:
	;
	goto L407
L418:
	;
	v1896 = F_lappend(m, v1794, v1894)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L21
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1889
	*(*int32)(unsafe.Add(mBase, _consts[181])) = v1896
	goto L399
L420:
	;
	v1928 = int32(0)
	m.Env.Pgmem_listen(m, v1903, v1928)
	mBase = m.M
	v1931 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v1931 == v1928 {
		goto L399
	} else {
		goto L426
	}
L421:
	;
	v1910 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L21
	} else {
		goto L422
	}
L422:
	;
	if v1910 == int32(0) {
		goto L420
	} else {
		goto L423
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+16)) = v1903
	v1916 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+20)) = v1916
	F_errmsg_internal(m, int32(704033), v1721+int32(16))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L21
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(520690), int32(1175), int32(106883))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L21
	} else {
		goto L425
	}
L425:
	;
	goto L420
L426:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1931)+4))
	if v1934 <= int32(0) {
		goto L399
	} else {
		goto L427
	}
L427:
	;
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1931)+12))
	v1939 = int32(0)
	goto L428
L428:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v1937+v1939<<(uint(int32(2))%32))))
	v1971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1903))))
	v1972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1968))))
	if v1972 == int32(0) {
		v1991 = v1971
		v1992 = v1972
		goto L431
	} else {
		goto L432
	}
L429:
	;
	goto L399
L430:
	;
	if v1992-v1991 == int32(0) {
		goto L438
	} else {
		goto L439
	}
L431:
	;
	goto L430
L432:
	;
	if v1971 != v1972 {
		v1991 = v1971
		v1992 = v1972
		goto L431
	} else {
		goto L433
	}
L433:
	;
	v1976 = v1968
	v1977 = v1903
	goto L434
L434:
	;
	v1980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1977)+1)))
	v1981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1976)+1)))
	if v1981 == int32(0) {
		v1991 = v1980
		v1992 = v1981
		goto L431
	} else {
		goto L436
	}
L435:
	;
	v1991 = v1980
	v1992 = v1981
	goto L431
L436:
	;
	v1984 = int32(1)
	if v1980 == v1981 {
		v1976 = v1976 + v1984
		v1977 = v1977 + v1984
		goto L434
	} else {
		goto L437
	}
L437:
	;
	goto L435
L438:
	;
	v1997 = F_list_delete_nth_cell(m, v1931, v1939)
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L21
	} else {
		goto L441
	}
L439:
	;
	goto L440
L440:
	;
	v2003 = v1939 + int32(1)
	if v1934 != v2003 {
		v1939 = v2003
		goto L428
	} else {
		goto L443
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, _consts[181])) = v1997
	F_pfree(m, v1968)
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L21
	} else {
		goto L442
	}
L442:
	;
	goto L399
L443:
	;
	goto L429
L444:
	;
	m.Env.Pgmem_listen(m, int32(785690), int32(2))
	mBase = m.M
	v2032 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	F_list_free_deep(m, v2032)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L21
	} else {
		goto L450
	}
L445:
	;
	v2011 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L21
	} else {
		goto L446
	}
L446:
	;
	if v2011 == int32(0) {
		goto L444
	} else {
		goto L447
	}
L447:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	*(*int32)(unsafe.Add(mBase, uint32(v1721)+32)) = v2016
	F_errmsg_internal(m, int32(704119), v1721+int32(32))
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L21
	} else {
		goto L448
	}
L448:
	;
	F_errfinish(m, int32(520690), int32(1205), int32(106903))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L21
	} else {
		goto L449
	}
L449:
	;
	goto L444
L450:
	;
	*(*int32)(unsafe.Add(mBase, _consts[181])) = int32(0)
	goto L399
L451:
	;
	goto L398
L452:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	if v2103 != 0 {
		goto L456
	} else {
		goto L457
	}
L453:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v2099 != 0 {
		goto L452
	} else {
		goto L454
	}
L454:
	;
	F_asyncQueueUnregister(m)
	mBase = m.M
	v2101 = m.ExcPending
	if v2101 != 0 {
		goto L21
	} else {
		goto L455
	}
L455:
	;
	goto L452
L456:
	;
	v2105 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v2108 = F_palloc(m, v2105<<(uint(int32(2))%32))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L21
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	v2326 = int32(*(*uint8)(unsafe.Add(mBase, _consts[162])))
	if v2326 != 0 {
		goto L495
	} else {
		goto L496
	}
L459:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v2114 = F_palloc(m, v2111<<(uint(int32(2))%32))
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		goto L21
	} else {
		goto L460
	}
L460:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v2121 = F_LWLockAcquire(m, v2117+int32(3456), int32(0))
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L21
	} else {
		goto L461
	}
L461:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v2124)+40))
	if v2125 == int32(-1) {
		goto L463
	} else {
		goto L464
	}
L462:
	;
	F_pfree(m, v2108)
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L21
	} else {
		goto L493
	}
L463:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v2129+int32(3456))
	mBase = m.M
	v2133 = m.ExcPending
	if v2133 != 0 {
		goto L21
	} else {
		goto L466
	}
L464:
	;
	goto L465
L465:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v2142 = v2125
	v2143 = int32(0)
	v2149 = v2139
	goto L467
L466:
	;
	goto L462
L467:
	;
	v2168 = v2142 << (uint(int32(5)) % 32)
	v2169 = v2124 + int32(56) + v2168
	v2170 = *(*int64)(unsafe.Add(mBase, uint32(v2169)+16))
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2169)))
	v2172 = *(*int64)(unsafe.Add(mBase, uint32(v2124)))
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+4))
	if v2149 == v2173 {
		goto L471
	} else {
		goto L472
	}
L468:
	;
	v2201 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v2201+int32(3456))
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L21
	} else {
		goto L478
	}
L469:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v2168+(v2124-int32(-64)))))
	if v2196 != int32(-1) {
		v2142 = v2196
		v2143 = v2193
		v2149 = v2194
		goto L467
	} else {
		goto L477
	}
L470:
	;
	v2183 = v2143 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v2108+v2183))) = v2171
	*(*int32)(unsafe.Add(mBase, uint32(v2183+v2114))) = v2142
	v2191 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v2193 = v2143 + int32(1)
	v2194 = v2191
	goto L469
L471:
	;
	if v2170 != v2172 {
		goto L470
	} else {
		goto L474
	}
L472:
	;
	goto L473
L473:
	;
	if v2172-v2170 < int64(4) {
		v2193 = v2143
		v2194 = v2149
		goto L469
	} else {
		goto L476
	}
L474:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v2169)+24))
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2124)+8))
	if v2176 != v2177 {
		goto L470
	} else {
		goto L475
	}
L475:
	;
	v2193 = v2143
	v2194 = v2149
	goto L469
L476:
	;
	goto L470
L477:
	;
	goto L468
L478:
	;
	if v2193 <= int32(0) {
		goto L462
	} else {
		goto L479
	}
L479:
	;
	v2208 = int32(0)
	goto L480
L480:
	;
	v2235 = v2208 << (uint(int32(2)) % 32)
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2108+v2235)))
	v2239 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	if v2237 == v2239 {
		goto L483
	} else {
		goto L484
	}
L481:
	;
	goto L462
L482:
	;
	v2267 = v2208 + int32(1)
	if v2267 != v2193 {
		v2208 = v2267
		goto L480
	} else {
		goto L492
	}
L483:
	;
	*(*int32)(unsafe.Add(mBase, _consts[183])) = int32(1)
	goto L482
L484:
	;
	goto L485
L485:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2235+v2114)))
	v2247 = F_SendProcSignal(m, v2237, int32(1), v2246)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L21
	} else {
		goto L486
	}
L486:
	;
	if int32(0) <= v2247 {
		goto L482
	} else {
		goto L487
	}
L487:
	;
	v2253 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L21
	} else {
		goto L488
	}
L488:
	;
	if v2253 == int32(0) {
		goto L482
	} else {
		goto L489
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1721))) = v2237
	F_errmsg_internal(m, int32(307750), v1721)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L21
	} else {
		goto L490
	}
L490:
	;
	F_errfinish(m, int32(520690), int32(1665), int32(181798))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L21
	} else {
		goto L491
	}
L491:
	;
	goto L482
L492:
	;
	goto L481
L493:
	;
	F_pfree(m, v2114)
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L21
	} else {
		goto L494
	}
L494:
	;
	goto L458
L495:
	;
	v2328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[162])) = uint8(v2328)
	F_asyncQueueAdvanceTail(m)
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L21
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	v2333 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[152])) = v2333
	*(*int32)(unsafe.Add(mBase, _consts[151])) = v2333
	goto L386
L498:
	;
	goto L497
L499:
	;
	F_AtEOXact_SPI(m, int32(1))
	mBase = m.M
	v2373 = m.ExcPending
	if v2373 != 0 {
		goto L21
	} else {
		goto L500
	}
L500:
	;
	v2375 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[184])) = v2375
	*(*int32)(unsafe.Add(mBase, _consts[185])) = v2375
	goto L501
L501:
	;
	F_AtEOXact_on_commit_actions(m, int32(1))
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L21
	} else {
		goto L502
	}
L502:
	;
	F_AtEOXact_Namespace(m, int32(1), v1591)
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L21
	} else {
		goto L503
	}
L503:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		goto L21
	} else {
		goto L504
	}
L504:
	;
	F_AtEOXact_Files(m, int32(1))
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L21
	} else {
		goto L505
	}
L505:
	;
	v2392 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[68])) = v2392
	*(*int32)(unsafe.Add(mBase, _consts[186])) = v2392
	*(*int32)(unsafe.Add(mBase, _consts[187])) = v2392
	*(*int32)(unsafe.Add(mBase, _consts[188])) = v2392
	goto L506
L506:
	;
	F_AtEOXact_HashTables(m, int32(1))
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L21
	} else {
		goto L507
	}
L507:
	;
	F_AtEOXact_PgStat(m, int32(1), v1591)
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L21
	} else {
		goto L508
	}
L508:
	;
	F_AtEOXact_Snapshot(m, int32(1), int32(0))
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L21
	} else {
		goto L509
	}
L509:
	;
	F_AtEOXact_ApplyLauncher(m, int32(1))
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L21
	} else {
		goto L510
	}
L510:
	;
	F_AtEOXact_LogicalRepWorkers(m, int32(1))
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L21
	} else {
		goto L511
	}
L511:
	;
	v2422 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v2422 != int32(1) {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	v2452 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	F_ResourceOwnerDelete(m, v2452)
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L21
	} else {
		goto L516
	}
L513:
	;
	goto L512
L514:
	;
	v2426 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v2426 == int32(0) {
		goto L513
	} else {
		goto L515
	}
L515:
	;
	v2429 = int32(4543684)
	v2431 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2432 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2431 + v2432
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2426)))
	*(*int32)(unsafe.Add(mBase, uint32(v2426))) = v2435 + v2432
	*(*int64)(unsafe.Add(mBase, uint32(v2426)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2426))) = v2435 + int32(2)
	v2446 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2446 - v2432
	goto L513
L516:
	;
	v2455 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+40)) = v2455
	*(*int32)(unsafe.Add(mBase, _consts[176])) = v2455
	*(*int32)(unsafe.Add(mBase, _consts[189])) = v2455
	v2465 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v2465)+44))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v2466
	v2469 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	F_MemoryContextReset(m, v2469)
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L21
	} else {
		goto L517
	}
L517:
	;
	v2473 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v2473
	*(*int32)(unsafe.Add(mBase, uint32(v2465)+36)) = v2473
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v2473
	v2479 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+48)) = v2479
	*(*int64)(unsafe.Add(mBase, uint32(v32)+28)) = v2479
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v2473
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v2479
	*(*int64)(unsafe.Add(mBase, _consts[69])) = v2479
	*(*int32)(unsafe.Add(mBase, _consts[70])) = v2473
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v2473
	v2495 = int32(4543676)
	v2497 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v2497 - int32(1)
	m.G0 = v29 + int32(48)
	return
}
func F_IsTransactionBlock(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	return base.B2i32(base.Ui32(int32(1)) < base.Ui32(v3))
}
func F_PopTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v11 == int32(0) {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
		if v41 != 0 {
			*(*int32)(unsafe.Add(mBase, _consts[4])) = v41
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v45
			*(*int32)(unsafe.Add(mBase, _consts[141])) = v45
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
			*(*int32)(unsafe.Add(mBase, _consts[175])) = v50
			*(*int32)(unsafe.Add(mBase, _consts[189])) = v50
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			if v54 != 0 {
				F_pfree(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_pfree(m, v10)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(99685), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					F_errfinish(m, int32(512898), int32(5487), int32(268715))
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
	} else {
		v16 = F_errstart(m, int32(19), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 == int32(0) {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
				if v41 != 0 {
					*(*int32)(unsafe.Add(mBase, _consts[4])) = v41
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v45
					*(*int32)(unsafe.Add(mBase, _consts[141])) = v45
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
					*(*int32)(unsafe.Add(mBase, _consts[175])) = v50
					*(*int32)(unsafe.Add(mBase, _consts[189])) = v50
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v54 != 0 {
						F_pfree(m, v54)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_pfree(m, v10)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					} else {
						F_pfree(m, v10)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				} else {
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(99685), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							F_errfinish(m, int32(512898), int32(5487), int32(268715))
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
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
				if base.Ui32(v21) <= base.Ui32(int32(5)) {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_consts[149])))
					v29 = v28
				} else {
					v29 = int32(565099)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v29
				F_errmsg_internal(m, int32(367075), v7)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_errfinish(m, int32(512898), int32(5484), int32(268715))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
						if v41 != 0 {
							*(*int32)(unsafe.Add(mBase, _consts[4])) = v41
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+36))
							*(*int32)(unsafe.Add(mBase, _consts[3])) = v45
							*(*int32)(unsafe.Add(mBase, _consts[141])) = v45
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
							*(*int32)(unsafe.Add(mBase, _consts[175])) = v50
							*(*int32)(unsafe.Add(mBase, _consts[189])) = v50
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							if v54 != 0 {
								F_pfree(m, v54)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_pfree(m, v10)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							} else {
								F_pfree(m, v10)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(99685), int32(0))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									F_errfinish(m, int32(512898), int32(5487), int32(268715))
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
func F_PrepareTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int64
	_ = v258
	var v259 int64
	_ = v259
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v554 int32
	_ = v554
	var v556 int64
	_ = v556
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int64
	_ = v603
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
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
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1240 int32
	_ = v1240
	var v1243 int64
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1279 int32
	_ = v1279
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1466 int32
	_ = v1466
	var v1469 int64
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1509 int32
	_ = v1509
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1595 int32
	_ = v1595
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1701 int32
	_ = v1701
	var v1729 int64
	_ = v1729
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1742 int64
	_ = v1742
	var v1744 int64
	_ = v1744
	var v1753 int64
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1764 int64
	_ = v1764
	var v1766 int64
	_ = v1766
	var v1775 int64
	_ = v1775
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1788 int32
	_ = v1788
	var v1793 int32
	_ = v1793
	var v1794 int64
	_ = v1794
	var v1795 int64
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1803 int64
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1845 int32
	_ = v1845
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1854 int64
	_ = v1854
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1916 int32
	_ = v1916
	var v1947 int32
	_ = v1947
	var v1949 int64
	_ = v1949
	var v1951 int64
	_ = v1951
	var v1953 int32
	_ = v1953
	var v1960 int32
	_ = v1960
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2033 int32
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2040 int32
	_ = v2040
	var v2044 int32
	_ = v2044
	var v2049 int32
	_ = v2049
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2062 int32
	_ = v2062
	var v2067 int32
	_ = v2067
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2080 int32
	_ = v2080
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2091 int32
	_ = v2091
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2163 int32
	_ = v2163
	var v2164 int64
	_ = v2164
	var v2166 int64
	_ = v2166
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2279 int32
	_ = v2279
	var v2280 int64
	_ = v2280
	var v2282 int64
	_ = v2282
	var v2284 int64
	_ = v2284
	var v2286 int64
	_ = v2286
	var v2288 int64
	_ = v2288
	var v2290 int64
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2377 int32
	_ = v2377
	var v2384 int32
	_ = v2384
	var v2389 int32
	_ = v2389
	var v2391 int32
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2401 int32
	_ = v2401
	var v2404 int32
	_ = v2404
	var v2408 int32
	_ = v2408
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2444 int32
	_ = v2444
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2467 int32
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2477 int32
	_ = v2477
	var v2480 int32
	_ = v2480
	var v2482 int32
	_ = v2482
	var v2488 int64
	_ = v2488
	var v2491 int64
	_ = v2491
	var v2497 int32
	_ = v2497
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2507 int32
	_ = v2507
	var v2508 int32
	_ = v2508
	var v2513 int32
	_ = v2513
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2580 int32
	_ = v2580
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2587 int64
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2595 int64
	_ = v2595
	var v2597 int32
	_ = v2597
	var v2598 int64
	_ = v2598
	var v2599 int64
	_ = v2599
	var v2601 int32
	_ = v2601
	var v2603 int64
	_ = v2603
	var v2606 int32
	_ = v2606
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2615 int32
	_ = v2615
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2628 int32
	_ = v2628
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2637 int32
	_ = v2637
	var v2639 int32
	_ = v2639
	var v2643 int64
	_ = v2643
	var v2646 int32
	_ = v2646
	var v2648 int32
	_ = v2648
	var v2659 int32
	_ = v2659
	var v2662 int32
	_ = v2662
	var v2666 int32
	_ = v2666
	var v2671 int32
	_ = v2671
	var v2675 int32
	_ = v2675
	var v2677 int32
	_ = v2677
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2691 int32
	_ = v2691
	var v2693 int32
	_ = v2693
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2702 int32
	_ = v2702
	var v2727 int32
	_ = v2727
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2747 int32
	_ = v2747
	var v2757 int32
	_ = v2757
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2767 int32
	_ = v2767
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2796 int32
	_ = v2796
	var v2801 int32
	_ = v2801
	var v2806 int32
	_ = v2806
	var v2810 int32
	_ = v2810
	var v2816 int32
	_ = v2816
	var v2818 int32
	_ = v2818
	var v2820 int32
	_ = v2820
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2825 int32
	_ = v2825
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2857 int32
	_ = v2857
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2886 int32
	_ = v2886
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2893 int32
	_ = v2893
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2930 int64
	_ = v2930
	var v2933 int32
	_ = v2933
	var v2935 int32
	_ = v2935
	var v2969 int32
	_ = v2969
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3035 int32
	_ = v3035
	var v3037 int32
	_ = v3037
	var v3040 int32
	_ = v3040
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3075 int32
	_ = v3075
	var v3081 int32
	_ = v3081
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3119 int32
	_ = v3119
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3128 int32
	_ = v3128
	var v3133 int32
	_ = v3133
	var v3135 int32
	_ = v3135
	var v3137 int32
	_ = v3137
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3148 int32
	_ = v3148
	var v3150 int32
	_ = v3150
	var v3152 int32
	_ = v3152
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3163 int32
	_ = v3163
	var v3165 int32
	_ = v3165
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3172 int32
	_ = v3172
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3193 int32
	_ = v3193
	var v3213 int32
	_ = v3213
	var v3216 int32
	_ = v3216
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3228 int32
	_ = v3228
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3237 int32
	_ = v3237
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3267 int32
	_ = v3267
	var v3279 int32
	_ = v3279
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3296 int32
	_ = v3296
	var v3317 int32
	_ = v3317
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3356 int32
	_ = v3356
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3368 int32
	_ = v3368
	var v3373 int32
	_ = v3373
	var v3404 int32
	_ = v3404
	var v3408 int32
	_ = v3408
	var v3409 int32
	_ = v3409
	var v3415 int32
	_ = v3415
	var v3420 int32
	_ = v3420
	var v3423 int32
	_ = v3423
	var v3429 int32
	_ = v3429
	var v3493 int32
	_ = v3493
	var v3495 int32
	_ = v3495
	var v3497 int32
	_ = v3497
	var v3499 int32
	_ = v3499
	var v3504 int32
	_ = v3504
	var v3528 int32
	_ = v3528
	var v3531 int32
	_ = v3531
	var v3533 int32
	_ = v3533
	var v3543 int32
	_ = v3543
	var v3546 int32
	_ = v3546
	var v3550 int32
	_ = v3550
	var v3555 int32
	_ = v3555
	var v3559 int32
	_ = v3559
	var v3563 int32
	_ = v3563
	var v3568 int32
	_ = v3568
	var v3572 int32
	_ = v3572
	var v3576 int32
	_ = v3576
	var v3581 int32
	_ = v3581
	var v3583 int32
	_ = v3583
	var v3585 int32
	_ = v3585
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3592 int32
	_ = v3592
	var v3593 int32
	_ = v3593
	var v3594 int32
	_ = v3594
	var v3598 int32
	_ = v3598
	var v3607 int32
	_ = v3607
	var v3608 int64
	_ = v3608
	var v3612 int32
	_ = v3612
	var v3615 int32
	_ = v3615
	var v3619 int32
	_ = v3619
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3624 int32
	_ = v3624
	var v3627 int32
	_ = v3627
	var v3628 int32
	_ = v3628
	var v3636 int32
	_ = v3636
	var v3640 int32
	_ = v3640
	var v3642 int32
	_ = v3642
	var v3646 int32
	_ = v3646
	var v3672 int32
	_ = v3672
	var v3674 int32
	_ = v3674
	var v3675 int32
	_ = v3675
	var v3677 int32
	_ = v3677
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3713 int32
	_ = v3713
	var v3715 int32
	_ = v3715
	var v3718 int32
	_ = v3718
	var v3720 int32
	_ = v3720
	var v3722 int32
	_ = v3722
	var v3723 int32
	_ = v3723
	var v3727 int32
	_ = v3727
	var v3753 int32
	_ = v3753
	var v3756 int32
	_ = v3756
	var v3819 int32
	_ = v3819
	var v3824 int32
	_ = v3824
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3831 int32
	_ = v3831
	var v3858 int32
	_ = v3858
	var v3863 int32
	_ = v3863
	var v3865 int32
	_ = v3865
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3902 int32
	_ = v3902
	var v3929 int32
	_ = v3929
	var v3934 int32
	_ = v3934
	var v3936 int32
	_ = v3936
	var v4000 int32
	_ = v4000
	var v4004 int32
	_ = v4004
	var v4031 int32
	_ = v4031
	var v4034 int32
	_ = v4034
	var v4065 int32
	_ = v4065
	var v4067 int32
	_ = v4067
	var v4071 int32
	_ = v4071
	var v4073 int32
	_ = v4073
	var v4074 int32
	_ = v4074
	var v4076 int32
	_ = v4076
	var v4080 int32
	_ = v4080
	var v4081 int32
	_ = v4081
	var v4083 int32
	_ = v4083
	var v4084 int32
	_ = v4084
	var v4089 int32
	_ = v4089
	var v4096 int32
	_ = v4096
	var v4100 int32
	_ = v4100
	var v4102 int32
	_ = v4102
	var v4103 int32
	_ = v4103
	var v4106 int32
	_ = v4106
	var v4110 int32
	_ = v4110
	var v4113 int32
	_ = v4113
	var v4125 int32
	_ = v4125
	var v4129 int32
	_ = v4129
	var v4131 int32
	_ = v4131
	var v4133 int32
	_ = v4133
	var v4142 int32
	_ = v4142
	var v4144 int32
	_ = v4144
	var v4147 int32
	_ = v4147
	var v4149 int32
	_ = v4149
	var v4151 int32
	_ = v4151
	var v4154 int32
	_ = v4154
	var v4156 int32
	_ = v4156
	var v4160 int32
	_ = v4160
	var v4161 int32
	_ = v4161
	var v4163 int32
	_ = v4163
	var v4167 int32
	_ = v4167
	var v4171 int32
	_ = v4171
	var v4175 int32
	_ = v4175
	var v4178 int32
	_ = v4178
	var v4181 int32
	_ = v4181
	var v4183 int32
	_ = v4183
	var v4190 int32
	_ = v4190
	var v4194 int32
	_ = v4194
	var v4196 int32
	_ = v4196
	var v4199 int32
	_ = v4199
	var v4201 int32
	_ = v4201
	var v4214 int32
	_ = v4214
	var v4215 int32
	_ = v4215
	var v4218 int32
	_ = v4218
	var v4221 int32
	_ = v4221
	var v4224 int32
	_ = v4224
	var v4228 int32
	_ = v4228
	var v4232 int32
	_ = v4232
	var v4235 int32
	_ = v4235
	var v4237 int32
	_ = v4237
	var v4238 int32
	_ = v4238
	var v4241 int32
	_ = v4241
	var v4252 int32
	_ = v4252
	var v4261 int32
	_ = v4261
	var v4263 int32
	_ = v4263
	var v4264 int32
	_ = v4264
	var v4274 int32
	_ = v4274
	var v4275 int32
	_ = v4275
	var v4278 int32
	_ = v4278
	var v4280 int32
	_ = v4280
	var v4282 int32
	_ = v4282
	var v4288 int64
	_ = v4288
	var v4304 int32
	_ = v4304
	var v4306 int32
	_ = v4306
	var v4316 int32
	_ = v4316
	var v4319 int32
	_ = v4319
	var v4323 int32
	_ = v4323
	var v4328 int32
	_ = v4328
	var v4332 int32
	_ = v4332
	var v4335 int32
	_ = v4335
	var v4339 int32
	_ = v4339
	var v4344 int32
	_ = v4344
	v30 = m.G0
	v32 = v30 - int32(16)
	m.G0 = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v36 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35))))
	if v36 == int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_AssignTransactionId(m, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v42 = v36
	goto L3
L3:
	;
	v43 = int32(10)
	goto L8
L4:
	;
	return
L5:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v42 = v41
	goto L3
L6:
	;
	if v77 != 0 {
		goto L20
	} else {
		goto L21
	}
L7:
	;
	goto L6
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	goto L11
L9:
	;
	v62 = int32(0)
	goto L17
L11:
	;
	goto L12
L12:
	;
	goto L14
L14:
	;
	if v50 == int32(15) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	if v50 <= v43 {
		v77 = int32(1)
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	if v66 != int32(2) {
		v77 = v62
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _consts[147])))
	if v70 != 0 {
		v77 = v62
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	v77 = int32(0) | base.B2i32(v74 <= v43)
	goto L7
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	F_ShowTransactionStateRec(m, int32(268746), v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if v84 == int32(2) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	v113 = base.I32_wrap_i64(v42)
	goto L33
L25:
	;
	v89 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	if v89 == int32(0) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if base.Ui32(v93) <= base.Ui32(int32(5)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93<<(uint(int32(2))%32))+uint32(_consts[149])))
	v102 = v100
	goto L30
L29:
	;
	v102 = int32(565099)
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v102
	F_errmsg_internal(m, int32(367108), v32)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(512898), int32(2531), int32(268746))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L24
L33:
	;
	F_AfterTriggerFireDeferred(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	if v149 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v146 = F_PreCommit_Portals(m, int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	if v146 != 0 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	v153 = v149
	goto L41
L39:
	;
	goto L40
L40:
	;
	F_AfterTriggerEndXact(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L45
	}
L41:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	m.T0[v182].(func(*base.Module, int32, int32))(m, int32(7), v181)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	if v179 != 0 {
		v153 = v179
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	F_PreCommit_on_commit_actions(m)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_smgrDoPendingSyncs(m, int32(1), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_AtEOXact_LargeObject(m, int32(1))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_PreCommit_CheckForSerializationFailure(m)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, _consts[22])))
	if v228&int32(1) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4332 = m.ExcPending
	if v4332 != 0 {
		goto L4
	} else {
		goto L664
	}
L51:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[191]))
	if v234 != 0 {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4316 = m.ExcPending
	if v4316 != 0 {
		goto L4
	} else {
		goto L660
	}
L54:
	;
	v235 = int32(4543676)
	v237 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v237 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(5)
	v244 = *(*int32)(unsafe.Add(mBase, _consts[164]))
	if int32(0) < v244 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_disable_timeout(m, int32(8))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v253 = m.G0
	v254 = int32(16)
	v255 = v253 - v254
	m.G0 = v255
	F___gettimeofday(m, v255)
	mBase = m.M
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v255)))
	v259 = int64(*(*int32)(unsafe.Add(mBase, uint32(v255)+8)))
	m.G0 = v255 + v254
	goto L59
L58:
	;
	goto L57
L59:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	v271 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v273 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	v275 = m.G0
	v277 = v275 - int32(48)
	m.G0 = v277
	v279 = F_strlen(m, v269)
	mBase = m.M
	if base.Ui32(v279) < base.Ui32(int32(200)) {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v509 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[192])) = v509
	v511 = m.G0
	v513 = v511 - int32(96)
	m.G0 = v513
	v516 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v421)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v513)+8)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v513)+4)) = v509
	v526 = F_palloc0(m, int32(12))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L4
	} else {
		goto L108
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L103
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L98
	}
L63:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	if v283 == int32(0) {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L94
	}
L66:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, _consts[193])))
	if v287 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_before_shmem_exit(m, int32(393), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v302 = F_LWLockAcquire(m, v298+int32(2304), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L71
	}
L70:
	;
	v295 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[193])) = uint8(v295)
	goto L69
L71:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	if v306 <= int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v421 == int32(0) {
		goto L61
	} else {
		goto L92
	}
L73:
	;
	v314 = int32(0)
	goto L74
L74:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v305+int32(8)+v314<<(uint(int32(2))%32))))
	v345 = v343 + int32(47)
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	if v349 == int32(0) {
		v368 = v348
		v369 = v349
		goto L77
	} else {
		goto L78
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L88
	}
L76:
	;
	if v369-v368 != 0 {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	goto L76
L78:
	;
	if v348 != v349 {
		v368 = v348
		v369 = v349
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v353 = v345
	v354 = v269
	goto L80
L80:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+1)))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+1)))
	if v358 == int32(0) {
		v368 = v357
		v369 = v358
		goto L77
	} else {
		goto L82
	}
L81:
	;
	v368 = v357
	v369 = v358
	goto L77
L82:
	;
	v361 = int32(1)
	if v357 == v358 {
		v353 = v353 + v361
		v354 = v354 + v361
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v372 = v314 + int32(1)
	if v306 != v372 {
		v314 = v372
		goto L74
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	goto L75
L87:
	;
	goto L72
L88:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+16)) = v269
	F_errmsg(m, int32(375128), v277+int32(16))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(519151), int32(396), int32(345183))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v424
	F_MarkAsPreparingGuts(m, v421, v113, v269, v259+v258*int64(1000000)-int64(946684800000000), v271, v273)
	mBase = m.M
	v427 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v421)+45)) = uint8(v427)
	v430 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v430)+4)) = v431 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v430+v431<<(uint(int32(2))%32))+8)) = v421
	v440 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v440+int32(2304))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	m.G0 = v277 + int32(48)
	goto L60
L94:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v269
	F_errmsg(m, int32(341882), v277)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(519151), int32(369), int32(345183))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(472823), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	F_errhint(m, int32(652145), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(519151), int32(376), int32(345183))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(476932), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+32)) = v496
	F_errhint(m, int32(686683), v277+int32(32))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(519151), int32(406), int32(345183))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, _consts[194])) = v526
	*(*int64)(unsafe.Add(mBase, uint32(v526)+4)) = int64(0)
	v532 = int32(512)
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v532
	v535 = F_palloc(m, v532)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _consts[194]))
	*(*int32)(unsafe.Add(mBase, uint32(v538))) = v535
	*(*int32)(unsafe.Add(mBase, _consts[196])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v538
	v546 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v513)+32)) = v519
	*(*int64)(unsafe.Add(mBase, uint32(v513)+24)) = int64(1475953972)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v517+v518*int32(640))+60))
	*(*int32)(unsafe.Add(mBase, uint32(v513)+36)) = v554
	v556 = *(*int64)(unsafe.Add(mBase, uint32(v421)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v513)+40)) = v556
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v421)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v513)+48)) = v558
	v564 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+52))
	if v565 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+52)) = v569
	v574 = F_smgrGetPendingDeletes(m, int32(1), v513+int32(16))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L114
	}
L111:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v564)+48))
	v567 = v566
	goto L113
L112:
	;
	v567 = v546
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513+int32(20)))) = v567
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v564)+52))
	goto L110
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+56)) = v574
	v580 = F_smgrGetPendingDeletes(m, int32(0), v513+int32(12))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+60)) = v580
	v586 = F_pgstat_get_transactional_drops(m, int32(1), v513+int32(4))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+64)) = v586
	v592 = F_pgstat_get_transactional_drops(m, int32(0), v513+int32(8))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+68)) = v592
	v597 = F_xactGetCommittedInvalidationMessages(m, v513, v513+int32(76))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+72)) = v597
	v601 = v421 + int32(47)
	v602 = F_strlen(m, v601)
	mBase = m.M
	v603 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v513)+88)) = v603
	*(*int64)(unsafe.Add(mBase, uint32(v513)+80)) = v603
	v608 = v602 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v513)+78)) = uint16(v608)
	v611 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	if base.Ui32(int32(72)) <= base.Ui32(v611) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v644)+4))
	goto L126
L120:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	v644 = v615
	v645 = v616
	v646 = v611
	goto L119
L121:
	;
	goto L122
L122:
	;
	v618 = F_palloc0(m, int32(12))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v620 = int32(4444048)
	v621 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v621)+8)) = v618
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v618
	*(*int64)(unsafe.Add(mBase, uint32(v618)+4)) = int64(0)
	v628 = int32(512)
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v628
	v630 = int32(4444052)
	v632 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v632 + int32(1)
	v637 = F_palloc(m, v628)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v640))) = v637
	v643 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v644 = v640
	v645 = v637
	v646 = v643
	goto L119
L125:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v644)+4))
	v655 = int32(72)
	v656 = v654 + v655
	*(*int32)(unsafe.Add(mBase, uint32(v644)+4)) = v656
	v660 = v646 - v655
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v660
	v662 = int32(4444060)
	v664 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	v666 = v664 + v655
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v666
	v668 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v513)+78)))
	v672 = (v668 + int32(7)) & int32(131064)
	if base.Ui32(v672) <= base.Ui32(v660) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	v652 = F__emscripten_memcpy_bulkmem(m, v645+v647, v513+int32(24), int32(72))
	mBase = m.M
	goto L128
L128:
	;
	goto L125
L129:
	;
	if v668 != 0 {
		goto L139
	} else {
		goto L140
	}
L130:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	v707 = v644
	v708 = v660
	v709 = v656
	v710 = v666
	v711 = v674
	goto L129
L131:
	;
	goto L132
L132:
	;
	v676 = F_palloc0(m, int32(12))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	v678 = int32(4444048)
	v679 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v679)+8)) = v676
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v676
	*(*int64)(unsafe.Add(mBase, uint32(v676)+4)) = int64(0)
	v686 = int32(512)
	if base.Ui32(v672) <= base.Ui32(v686) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v689 = v686
	goto L136
L135:
	;
	v689 = v672
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v689
	v691 = int32(4444052)
	v693 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v693 + int32(1)
	v697 = F_palloc(m, v689)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v700))) = v697
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v700)+4))
	v704 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	v706 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v707 = v700
	v708 = v706
	v709 = v702
	v710 = v704
	v711 = v697
	goto L129
L138:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	v716 = v715 + v672
	*(*int32)(unsafe.Add(mBase, uint32(v707)+4)) = v716
	v719 = v672 + v710
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v719
	v722 = v708 - v672
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v722
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v513)+52))
	if v724 <= int32(0) {
		v811 = v722
		v812 = v719
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v713 = F__emscripten_memcpy_bulkmem(m, v709+v711, v601, v668)
	mBase = m.M
	goto L141
L140:
	;
	goto L141
L141:
	;
	goto L138
L142:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v513)+56))
	if int32(0) < v816 {
		goto L166
	} else {
		goto L167
	}
L143:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v513)+20))
	v729 = v724 << (uint(int32(2)) % 32)
	v733 = (v729 + int32(7)) & int32(-8)
	if base.Ui32(v733) <= base.Ui32(v722) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	if v729 != 0 {
		goto L154
	} else {
		goto L155
	}
L145:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v766 = v707
	v767 = v722
	v768 = v735
	v769 = v716
	goto L144
L146:
	;
	goto L147
L147:
	;
	v737 = F_palloc0(m, int32(12))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v739 = int32(4444048)
	v740 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v740)+8)) = v737
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v737
	*(*int64)(unsafe.Add(mBase, uint32(v737)+4)) = int64(0)
	v747 = int32(512)
	if base.Ui32(v733) <= base.Ui32(v747) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v750 = v747
	goto L151
L150:
	;
	v750 = v733
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v750
	v752 = int32(4444052)
	v754 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v754 + int32(1)
	v758 = F_palloc(m, v750)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v761 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = v758
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v761)+4))
	v765 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v766 = v761
	v767 = v765
	v768 = v758
	v769 = v763
	goto L144
L153:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v766)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v766)+4)) = v773 + v733
	v777 = v767 - v733
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v777
	v779 = int32(4444060)
	v781 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	v782 = v781 + v733
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v782
	v785 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	v790 = v786 + v787*int32(640)
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v513)+20))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v513)+52))
	if int32(65) <= v792 {
		goto L158
	} else {
		goto L159
	}
L154:
	;
	v771 = F__emscripten_memcpy_bulkmem(m, v768+v769, v727, v729)
	mBase = m.M
	goto L156
L155:
	;
	goto L156
L156:
	;
	goto L153
L157:
	;
	v804 = v800 << (uint(int32(2)) % 32)
	if v804 != 0 {
		goto L163
	} else {
		goto L164
	}
L158:
	;
	v795 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v790)+277)) = uint8(v795)
	v800 = int32(64)
	goto L157
L159:
	;
	goto L160
L160:
	;
	if v792 <= int32(0) {
		v811 = v777
		v812 = v782
		goto L142
	} else {
		goto L161
	}
L161:
	;
	v800 = v792
	goto L157
L162:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v790)+276)) = uint8(v800)
	v811 = v777
	v812 = v782
	goto L142
L163:
	;
	v805 = F__emscripten_memcpy_bulkmem(m, v790+int32(280), v791, v804)
	mBase = m.M
	goto L165
L164:
	;
	goto L165
L165:
	;
	goto L162
L166:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v513)+16))
	v821 = v816 * int32(12)
	v825 = (v821 + int32(7)) & int32(-8)
	if base.Ui32(v825) <= base.Ui32(v811) {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	goto L168
L168:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v513)+60))
	if int32(0) < v888 {
		goto L183
	} else {
		goto L184
	}
L169:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	if v821 != 0 {
		goto L179
	} else {
		goto L180
	}
L170:
	;
	v828 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)))
	v861 = v828
	v862 = v829
	v863 = v811
	v864 = v812
	goto L169
L171:
	;
	goto L172
L172:
	;
	v831 = F_palloc0(m, int32(12))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v833 = int32(4444048)
	v834 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v834)+8)) = v831
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v831
	*(*int64)(unsafe.Add(mBase, uint32(v831)+4)) = int64(0)
	v841 = int32(512)
	if base.Ui32(v825) <= base.Ui32(v841) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v844 = v841
	goto L176
L175:
	;
	v844 = v825
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v844
	v846 = int32(4444052)
	v848 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v848 + int32(1)
	v852 = F_palloc(m, v844)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	v855 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v855))) = v852
	v858 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	v860 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v861 = v855
	v862 = v852
	v863 = v860
	v864 = v858
	goto L169
L178:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v861)+4)) = v869 + v825
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v825 + v864
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v863 - v825
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v513)+16))
	F_pfree(m, v878)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L182
	}
L179:
	;
	v867 = F__emscripten_memcpy_bulkmem(m, v862+v865, v819, v821)
	mBase = m.M
	goto L181
L180:
	;
	goto L181
L181:
	;
	goto L178
L182:
	;
	goto L168
L183:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v513)+12))
	v893 = v888 * int32(12)
	v897 = (v893 + int32(7)) & int32(-8)
	v899 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	if base.Ui32(v897) <= base.Ui32(v899) {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	goto L185
L185:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v513)+64))
	if int32(0) < v960 {
		goto L200
	} else {
		goto L201
	}
L186:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	if v893 != 0 {
		goto L196
	} else {
		goto L197
	}
L187:
	;
	v902 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	v933 = v902
	v934 = v903
	v935 = v899
	goto L186
L188:
	;
	goto L189
L189:
	;
	v905 = F_palloc0(m, int32(12))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	v907 = int32(4444048)
	v908 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v908)+8)) = v905
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v905
	*(*int64)(unsafe.Add(mBase, uint32(v905)+4)) = int64(0)
	v915 = int32(512)
	if base.Ui32(v897) <= base.Ui32(v915) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v918 = v915
	goto L193
L192:
	;
	v918 = v897
	goto L193
L193:
	;
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v918
	v920 = int32(4444052)
	v922 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v922 + int32(1)
	v926 = F_palloc(m, v918)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	v929 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v929))) = v926
	v932 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v933 = v929
	v934 = v926
	v935 = v932
	goto L186
L195:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v933)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v933)+4)) = v940 + v897
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v935 - v897
	v946 = int32(4444060)
	v948 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v948 + v897
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v513)+12))
	F_pfree(m, v951)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L4
	} else {
		goto L199
	}
L196:
	;
	v938 = F__emscripten_memcpy_bulkmem(m, v934+v936, v891, v893)
	mBase = m.M
	goto L198
L197:
	;
	goto L198
L198:
	;
	goto L195
L199:
	;
	goto L185
L200:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	v965 = v960 << (uint(int32(4)) % 32)
	v967 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	if base.Ui32(v965) <= base.Ui32(v967) {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	goto L202
L202:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v513)+68))
	if int32(0) < v1027 {
		goto L217
	} else {
		goto L218
	}
L203:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	if v965 != 0 {
		goto L213
	} else {
		goto L214
	}
L204:
	;
	v970 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v970)))
	v1001 = v970
	v1002 = v971
	v1003 = v967
	goto L203
L205:
	;
	goto L206
L206:
	;
	v973 = F_palloc0(m, int32(12))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	v975 = int32(4444048)
	v976 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v976)+8)) = v973
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v973
	*(*int64)(unsafe.Add(mBase, uint32(v973)+4)) = int64(0)
	v983 = int32(512)
	if base.Ui32(v965) <= base.Ui32(v983) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v986 = v983
	goto L210
L209:
	;
	v986 = v965
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v986
	v988 = int32(4444052)
	v990 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v990 + int32(1)
	v994 = F_palloc(m, v986)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L4
	} else {
		goto L211
	}
L211:
	;
	v997 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v997))) = v994
	v1000 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v1001 = v997
	v1002 = v994
	v1003 = v1000
	goto L203
L212:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1001)+4)) = v1008 + v965
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v1003 - v965
	v1014 = int32(4444060)
	v1016 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v1016 + v965
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	F_pfree(m, v1019)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L4
	} else {
		goto L216
	}
L213:
	;
	v1006 = F__emscripten_memcpy_bulkmem(m, v1002+v1004, v963, v965)
	mBase = m.M
	goto L215
L214:
	;
	goto L215
L215:
	;
	goto L212
L216:
	;
	goto L202
L217:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v513)+8))
	v1032 = v1027 << (uint(int32(4)) % 32)
	v1034 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	if base.Ui32(v1032) <= base.Ui32(v1034) {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	goto L219
L219:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v513)+72))
	if int32(0) < v1094 {
		goto L234
	} else {
		goto L235
	}
L220:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+4))
	if v1032 != 0 {
		goto L230
	} else {
		goto L231
	}
L221:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1037)))
	v1068 = v1037
	v1069 = v1038
	v1070 = v1034
	goto L220
L222:
	;
	goto L223
L223:
	;
	v1040 = F_palloc0(m, int32(12))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L4
	} else {
		goto L224
	}
L224:
	;
	v1042 = int32(4444048)
	v1043 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v1043)+8)) = v1040
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v1040
	*(*int64)(unsafe.Add(mBase, uint32(v1040)+4)) = int64(0)
	v1050 = int32(512)
	if base.Ui32(v1032) <= base.Ui32(v1050) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1053 = v1050
	goto L227
L226:
	;
	v1053 = v1032
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v1053
	v1055 = int32(4444052)
	v1057 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v1057 + int32(1)
	v1061 = F_palloc(m, v1053)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v1064))) = v1061
	v1067 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v1068 = v1064
	v1069 = v1061
	v1070 = v1067
	goto L220
L229:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1068)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1068)+4)) = v1075 + v1032
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v1070 - v1032
	v1081 = int32(4444060)
	v1083 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v1083 + v1032
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v513)+8))
	F_pfree(m, v1086)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L4
	} else {
		goto L233
	}
L230:
	;
	v1073 = F__emscripten_memcpy_bulkmem(m, v1069+v1071, v1030, v1032)
	mBase = m.M
	goto L232
L231:
	;
	goto L232
L232:
	;
	goto L229
L233:
	;
	goto L219
L234:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	v1099 = v1094 << (uint(int32(4)) % 32)
	v1101 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	if base.Ui32(v1099) <= base.Ui32(v1101) {
		goto L238
	} else {
		goto L239
	}
L235:
	;
	goto L236
L236:
	;
	m.G0 = v513 + int32(96)
	v1165 = *(*int32)(unsafe.Add(mBase, _consts[151]))
	v1167 = *(*int32)(unsafe.Add(mBase, _consts[152]))
	if v1165|v1167 != 0 {
		goto L251
	} else {
		goto L252
	}
L237:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+4))
	if v1099 != 0 {
		goto L247
	} else {
		goto L248
	}
L238:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1104)))
	v1135 = v1104
	v1136 = v1105
	v1137 = v1101
	goto L237
L239:
	;
	goto L240
L240:
	;
	v1107 = F_palloc0(m, int32(12))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	v1109 = int32(4444048)
	v1110 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+8)) = v1107
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v1107
	*(*int64)(unsafe.Add(mBase, uint32(v1107)+4)) = int64(0)
	v1117 = int32(512)
	if base.Ui32(v1099) <= base.Ui32(v1117) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1120 = v1117
	goto L244
L243:
	;
	v1120 = v1099
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v1120
	v1122 = int32(4444052)
	v1124 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v1124 + int32(1)
	v1128 = F_palloc(m, v1120)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L4
	} else {
		goto L245
	}
L245:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v1131))) = v1128
	v1134 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v1135 = v1131
	v1136 = v1128
	v1137 = v1134
	goto L237
L246:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1135)+4)) = v1142 + v1099
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v1137 - v1099
	v1148 = int32(4444060)
	v1150 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v1150 + v1099
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	F_pfree(m, v1153)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L4
	} else {
		goto L250
	}
L247:
	;
	v1140 = F__emscripten_memcpy_bulkmem(m, v1136+v1138, v1097, v1099)
	mBase = m.M
	goto L249
L248:
	;
	goto L249
L249:
	;
	goto L246
L250:
	;
	goto L236
L251:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L4
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v1185 = m.G0
	v1187 = v1185 - int32(80)
	m.G0 = v1187
	*(*int64)(unsafe.Add(mBase, uint32(v1187)+48)) = int64(85899345936)
	v1192 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+72)) = v1192
	v1199 = F_hash_create(m, int32(408651), int32(256), v1187+int32(32), int32(1064))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L4
	} else {
		goto L258
	}
L254:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	F_errmsg(m, int32(530735), int32(0))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L4
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(520690), int32(841), int32(21048))
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	F_hash_seq_init(m, v1187+int32(8), v1204)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L4
	} else {
		goto L259
	}
L259:
	;
	v1209 = F_hash_seq_search(m, v1187+int32(8))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L4
	} else {
		goto L262
	}
L260:
	;
	v2097 = m.G0
	v2099 = v2097 - int32(32)
	m.G0 = v2099
	v2102 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if v2102 != 0 {
		goto L389
	} else {
		goto L390
	}
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2084 = m.ExcPending
	if v2084 != 0 {
		goto L4
	} else {
		goto L385
	}
L262:
	;
	if v1209 != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1211 = v1209
	goto L266
L264:
	;
	goto L265
L265:
	;
	F_hash_destroy(m, v1199)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L4
	} else {
		goto L300
	}
L266:
	;
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211)+14)))
	if v1240 == int32(6) {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	goto L265
L268:
	;
	v1394 = F_hash_seq_search(m, v1187+int32(8))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L4
	} else {
		goto L298
	}
L269:
	;
	v1243 = *(*int64)(unsafe.Add(mBase, uint32(v1211)+32))
	if v1243 <= int64(0) {
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+48))
	v1250 = F_hash_search(m, v1199, v1211, int32(1), v1187+int32(7))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187)+7)))
	if v1252 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1255 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1250)+16)) = uint16(v1255)
	goto L274
L273:
	;
	goto L274
L274:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1211)+40))
	v1259 = v1257 - int32(1)
	if v1259 < int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1250)+16)))
	if v1357 != int32(1) {
		goto L268
	} else {
		goto L296
	}
L276:
	;
	if v1257&int32(1) != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1246+v1259<<(uint(int32(4))%32))))
	if v1267 != 0 {
		goto L281
	} else {
		goto L282
	}
L278:
	;
	v1274 = v1259
	goto L279
L279:
	;
	if v1259 == int32(0) {
		goto L275
	} else {
		goto L284
	}
L280:
	;
	v1274 = v1257 - int32(2)
	goto L279
L281:
	;
	v1268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1250)+17)) = uint8(v1268)
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1250)+16)) = uint8(v1270)
	goto L280
L284:
	;
	v1279 = v1274
	goto L285
L285:
	;
	v1309 = v1279 << (uint(int32(4)) % 32)
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1246+v1309)))
	if v1311 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	goto L275
L287:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1309+(v1246-int32(16)))))
	if v1319 != 0 {
		goto L292
	} else {
		goto L293
	}
L288:
	;
	v1314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1250)+16)) = uint8(v1314)
	goto L287
L289:
	;
	goto L290
L290:
	;
	v1316 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1250)+17)) = uint8(v1316)
	goto L287
L291:
	;
	if int32(1) < v1279 {
		v1279 = v1279 - int32(2)
		goto L285
	} else {
		goto L295
	}
L292:
	;
	v1320 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1250)+17)) = uint8(v1320)
	goto L291
L293:
	;
	goto L294
L294:
	;
	v1322 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1250)+16)) = uint8(v1322)
	goto L291
L295:
	;
	goto L286
L296:
	;
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1250)+17)))
	if v1360 == int32(1) {
		goto L261
	} else {
		goto L297
	}
L297:
	;
	goto L268
L298:
	;
	if v1394 != 0 {
		v1211 = v1394
		goto L266
	} else {
		goto L299
	}
L299:
	;
	goto L267
L300:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	F_hash_seq_init(m, v1187+int32(32), v1430)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L4
	} else {
		goto L301
	}
L301:
	;
	v1435 = F_hash_seq_search(m, v1187+int32(32))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L4
	} else {
		goto L305
	}
L302:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L4
	} else {
		goto L382
	}
L303:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L4
	} else {
		goto L379
	}
L304:
	;
	F_LWLockRelease(m, v1701)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L4
	} else {
		goto L372
	}
L305:
	;
	if v1435 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1441 = v1435
	goto L309
L307:
	;
	goto L308
L308:
	;
	m.G0 = v1187 + int32(80)
	goto L260
L309:
	;
	v1466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1441)+14)))
	if v1466 == int32(6) {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	goto L308
L311:
	;
	v1992 = F_hash_seq_search(m, v1187+int32(32))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L4
	} else {
		goto L370
	}
L312:
	;
	v1469 = *(*int64)(unsafe.Add(mBase, uint32(v1441)+32))
	if v1469 <= int64(0) {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+40))
	v1474 = v1472 - int32(1)
	if v1474 < int32(0) {
		goto L311
	} else {
		goto L314
	}
L314:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+48))
	v1478 = int32(3)
	v1479 = v1472 & v1478
	if base.Ui32(v1474) < base.Ui32(v1478) {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	if v1479 != 0 {
		goto L322
	} else {
		goto L323
	}
L316:
	;
	v1483 = int32(0)
	v1558 = v1474
	v1559 = v1483
	v1560 = v1483
	goto L315
L317:
	;
	goto L318
L318:
	;
	v1493 = int32(0)
	v1496 = v1474
	v1497 = v1493
	v1498 = v1493
	v1509 = v1493
	goto L319
L319:
	;
	v1525 = int32(4)
	v1526 = v1496 << (uint(v1525) % 32)
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1477-int32(48)+v1526)))
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1526+(v1477-int32(32)))))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1526+(v1477-int32(16)))))
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1477+v1526)))
	v1538 = int32(0)
	v1540 = base.B2i32(v1528|v1530|v1533|v1536 != v1538) | v1498
	v1552 = base.B2i32(v1536 == v1538) | (base.B2i32(v1533 == v1538) | (base.B2i32(v1528 == v1538) | base.B2i32(v1530 == v1538))) | v1497
	v1554 = v1496 - v1525
	v1556 = v1509 + v1525
	if v1556 != v1472&int32(-4) {
		v1496 = v1554
		v1497 = v1552
		v1498 = v1540
		v1509 = v1556
		goto L319
	} else {
		goto L321
	}
L320:
	;
	v1558 = v1554
	v1559 = v1552
	v1560 = v1540
	goto L315
L321:
	;
	goto L320
L322:
	;
	v1587 = v1558
	v1588 = v1559
	v1589 = v1560
	v1595 = int32(0)
	goto L325
L323:
	;
	v1632 = v1559
	v1633 = v1560
	goto L324
L324:
	;
	if v1633&int32(1) == int32(0) {
		goto L311
	} else {
		goto L328
	}
L325:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1477+v1587<<(uint(int32(4))%32))))
	v1620 = int32(0)
	v1622 = base.B2i32(v1619 != v1620) | v1589
	v1625 = base.B2i32(v1619 == v1620) | v1588
	v1626 = int32(1)
	v1629 = v1595 + v1626
	if v1629 != v1479 {
		v1587 = v1587 - v1626
		v1588 = v1625
		v1589 = v1622
		v1595 = v1629
		goto L325
	} else {
		goto L327
	}
L326:
	;
	v1632 = v1625
	v1633 = v1622
	goto L324
L327:
	;
	goto L326
L328:
	;
	if v1632&int32(1) == int32(0) {
		goto L333
	} else {
		goto L334
	}
L329:
	;
	v1947 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1441)+52)) = uint8(v1947)
	v1949 = *(*int64)(unsafe.Add(mBase, uint32(v1441)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1187)+16)) = v1949
	v1951 = *(*int64)(unsafe.Add(mBase, uint32(v1441)))
	*(*int64)(unsafe.Add(mBase, uint32(v1187)+8)) = v1951
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+24)) = v1953
	F_RegisterTwoPhaseRecord(m, int32(1), v1187+int32(8), int32(20))
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L4
	} else {
		goto L369
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1441)+28)) = v1907
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v1907)))
	*(*int32)(unsafe.Add(mBase, uint32(v1441)+24)) = v1916
	goto L329
L331:
	;
	F_LWLockRelease(m, v1690+int32(584))
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L4
	} else {
		goto L362
	}
L332:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+16))
	v1803 = int64(1) << (uint(base.I64_extend_i32_u(v1797+base.I32_wrap_i64(v1794)-int32(1))) % 64)
	if v1803&v1795 == int64(0) {
		goto L331
	} else {
		goto L353
	}
L333:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+28))
	if v1668 != 0 {
		goto L329
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L4
	} else {
		goto L349
	}
L336:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+20))
	v1673 = *(*int32)(unsafe.Add(mBase, _consts[201]))
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+4))
	v1676 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v1680 = F_LWLockAcquire(m, v1676+int32(584), int32(0))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L4
	} else {
		goto L337
	}
L337:
	;
	v1682 = int32(268435455)
	v1686 = (v1673 + v1682) & (v1674 * int32(49157))
	v1688 = v1686 & v1682
	v1690 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+604))
	v1694 = v1691 + v1686<<(uint(int32(6))%32)
	v1701 = v1670 + v1671&int32(15)<<(uint(int32(7))%32) + int32(23296)
	v1729 = int64(0)
	goto L338
L338:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1694+base.I32_wrap_i64(v1729)<<(uint(int32(2))%32))))
	if v1736 == v1674 {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	goto L331
L340:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+600))
	v1742 = *(*int64)(unsafe.Add(mBase, uint32(v1738+v1688<<(uint(int32(3))%32))))
	v1744 = v1729 * int64(3)
	if int64(base.Ui64(v1742)>>(uint(v1744)%64))&int64(7) != int64(0) {
		v1794 = v1744
		v1795 = v1742
		goto L332
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	v1753 = v1729 | int64(1)
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1694+base.I32_wrap_i64(v1753)<<(uint(int32(2))%32))))
	if v1758 == v1674 {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	goto L342
L344:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+600))
	v1764 = *(*int64)(unsafe.Add(mBase, uint32(v1760+v1688<<(uint(int32(3))%32))))
	v1766 = v1753 * int64(3)
	if int64(base.Ui64(v1764)>>(uint(v1766)%64))&int64(7) != int64(0) {
		v1794 = v1766
		v1795 = v1764
		goto L332
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	v1775 = v1729 + int64(2)
	if v1775 != int64(16) {
		v1729 = v1775
		goto L338
	} else {
		goto L348
	}
L347:
	;
	goto L346
L348:
	;
	goto L339
L349:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L4
	} else {
		goto L350
	}
L350:
	;
	F_errmsg(m, int32(118056), int32(0))
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L4
	} else {
		goto L351
	}
L351:
	;
	F_errfinish(m, int32(518098), int32(3494), int32(162806))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L4
	} else {
		goto L352
	}
L352:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L353:
	;
	v1808 = F_LWLockAcquire(m, v1701, int32(0))
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L4
	} else {
		goto L354
	}
L354:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+20))
	v1814 = F_SetupLockInTable(m, int32(1656888), v1812, v1441, v1813, v1797)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L4
	} else {
		goto L355
	}
L355:
	;
	if v1814 == int32(0) {
		goto L304
	} else {
		goto L356
	}
L356:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1814)))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+128))
	v1820 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+128)) = v1819 + v1820
	v1825 = v1818 + v1797<<(uint(int32(2))%32)
	v1827 = v1825 + int32(88)
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1827)))
	*(*int32)(unsafe.Add(mBase, uint32(v1827))) = v1828 + v1820
	v1833 = v1820 << (uint(v1797) % 32)
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+16)) = v1833 | v1834
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1827)))
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1825)+44))
	if v1837 == v1838 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+20)) = v1840 & (v1833 ^ int32(-1))
	goto L359
L358:
	;
	goto L359
L359:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1814)+12)) = v1845 | v1833
	v1849 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1849)+600))
	v1853 = v1850 + v1688<<(uint(int32(3))%32)
	v1854 = *(*int64)(unsafe.Add(mBase, uint32(v1853)))
	*(*int64)(unsafe.Add(mBase, uint32(v1853))) = v1854 & (v1803 ^ int64(-1))
	F_LWLockRelease(m, v1701)
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L4
	} else {
		goto L360
	}
L360:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	F_LWLockRelease(m, v1862+int32(584))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L4
	} else {
		goto L361
	}
L361:
	;
	v1907 = v1814
	goto L330
L362:
	;
	v1876 = F_LWLockAcquire(m, v1701, int32(1))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L4
	} else {
		goto L363
	}
L363:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, _consts[202]))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+20))
	v1881 = int32(0)
	v1883 = F_hash_search_with_hash_value(m, v1879, v1441, v1880, v1881, v1881)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L4
	} else {
		goto L364
	}
L364:
	;
	if v1883 == int32(0) {
		goto L303
	} else {
		goto L365
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+8)) = v1883
	v1889 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+12)) = v1889
	v1892 = *(*int32)(unsafe.Add(mBase, _consts[203]))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+20))
	v1899 = int32(0)
	v1901 = F_hash_search_with_hash_value(m, v1892, v1187+int32(8), v1895^v1889<<(uint(int32(4))%32), v1899, v1899)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L4
	} else {
		goto L366
	}
L366:
	;
	if v1901 == int32(0) {
		goto L302
	} else {
		goto L367
	}
L367:
	;
	F_LWLockRelease(m, v1701)
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L4
	} else {
		goto L368
	}
L368:
	;
	v1907 = v1901
	goto L330
L369:
	;
	goto L311
L370:
	;
	if v1992 != 0 {
		v1441 = v1992
		goto L309
	} else {
		goto L371
	}
L371:
	;
	goto L310
L372:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	F_LWLockRelease(m, v2029+int32(584))
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L4
	} else {
		goto L373
	}
L373:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L4
	} else {
		goto L374
	}
L374:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L4
	} else {
		goto L375
	}
L375:
	;
	F_errmsg(m, int32(14198), int32(0))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L4
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1187))) = int32(266632)
	F_errhint(m, int32(691511), v1187)
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L4
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(518098), int32(2969), int32(12733))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(117901), int32(0))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L4
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(518098), int32(2997), int32(12733))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L4
	} else {
		goto L381
	}
L381:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L382:
	;
	F_errmsg_internal(m, int32(117860), int32(0))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L4
	} else {
		goto L383
	}
L383:
	;
	F_errfinish(m, int32(518098), int32(3010), int32(12733))
	mBase = m.M
	v2080 = m.ExcPending
	if v2080 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	F_errmsg(m, int32(118056), int32(0))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	F_errfinish(m, int32(518098), int32(3426), int32(162706))
	mBase = m.M
	v2096 = m.ExcPending
	if v2096 != 0 {
		goto L4
	} else {
		goto L388
	}
L388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2099)+8)) = int32(0)
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v2102)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v2099)+12)) = v2105
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v2102)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v2099)+16)) = v2107
	F_RegisterTwoPhaseRecord(m, int32(4), v2099+int32(8), int32(24))
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L4
	} else {
		goto L392
	}
L390:
	;
	goto L391
L391:
	;
	m.G0 = v2099 + int32(32)
	v2244 = *(*int32)(unsafe.Add(mBase, _consts[204]))
	if v2244 != 0 {
		goto L402
	} else {
		goto L403
	}
L392:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v2120 = F_LWLockAcquire(m, v2116+int32(3840), int32(1))
	mBase = m.M
	v2121 = m.ExcPending
	if v2121 != 0 {
		goto L4
	} else {
		goto L393
	}
L393:
	;
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v2102)+52))
	if v2122 == int32(0) {
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v2206+int32(3840))
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L4
	} else {
		goto L401
	}
L395:
	;
	v2126 = v2102 + int32(48)
	if v2122 == v2126 {
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v2129 = v2099 + int32(12)
	v2132 = v2122
	goto L397
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2099)+8)) = int32(1)
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2132-int32(16))))
	v2164 = *(*int64)(unsafe.Add(mBase, uint32(v2163)))
	*(*int64)(unsafe.Add(mBase, uint32(v2129))) = v2164
	v2166 = *(*int64)(unsafe.Add(mBase, uint32(v2163)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2129)+8)) = v2166
	F_RegisterTwoPhaseRecord(m, int32(4), v2099+int32(8), int32(24))
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L4
	} else {
		goto L399
	}
L398:
	;
	goto L394
L399:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+4))
	if v2174 != v2126 {
		v2132 = v2174
		goto L397
	} else {
		goto L400
	}
L400:
	;
	goto L398
L401:
	;
	goto L391
L402:
	;
	v2245 = m.G0
	v2247 = v2245 + int32(-64)
	m.G0 = v2247
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v2244)+20))
	if v2249 != 0 {
		goto L405
	} else {
		goto L406
	}
L403:
	;
	goto L404
L404:
	;
	v2366 = m.G0
	v2368 = v2366 - int32(16)
	m.G0 = v2368
	v2371 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v2373 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(v2371+v2373<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2368)+12)) = v2377
	if v2377 != 0 {
		goto L412
	} else {
		goto L413
	}
L405:
	;
	v2250 = v2249
	goto L408
L406:
	;
	goto L407
L407:
	;
	m.G0 = v2247 - int32(-64)
	goto L404
L408:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+64))
	v2280 = *(*int64)(unsafe.Add(mBase, uint32(v2250)))
	*(*int64)(unsafe.Add(mBase, uint32(v2247)+8)) = v2280
	v2282 = *(*int64)(unsafe.Add(mBase, uint32(v2250)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2247)+16)) = v2282
	v2284 = *(*int64)(unsafe.Add(mBase, uint32(v2250)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2247)+24)) = v2284
	v2286 = *(*int64)(unsafe.Add(mBase, uint32(v2250)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2247)+32)) = v2286
	v2288 = *(*int64)(unsafe.Add(mBase, uint32(v2250)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2247)+40)) = v2288
	v2290 = *(*int64)(unsafe.Add(mBase, uint32(v2250)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v2247)+48)) = v2290
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v2279)))
	*(*int32)(unsafe.Add(mBase, uint32(v2247)+56)) = v2292
	v2294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2279)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2247)+60)) = uint8(v2294)
	v2296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2250)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2247)+61)) = uint8(v2296)
	F_RegisterTwoPhaseRecord(m, int32(2), v2245+int32(-56), int32(56))
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L4
	} else {
		goto L410
	}
L409:
	;
	goto L407
L410:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2250)+68))
	if v2304 != 0 {
		v2250 = v2304
		goto L408
	} else {
		goto L411
	}
L411:
	;
	goto L409
L412:
	;
	F_RegisterTwoPhaseRecord(m, int32(3), v2368+int32(12), int32(4))
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L4
	} else {
		goto L415
	}
L413:
	;
	goto L414
L414:
	;
	m.G0 = v2368 + int32(16)
	v2389 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v2389 != 0 {
		goto L417
	} else {
		goto L418
	}
L415:
	;
	goto L414
L416:
	;
	v2415 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	if base.Ui32(int32(8)) <= base.Ui32(v2415) {
		goto L427
	} else {
		goto L428
	}
L417:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2401 = m.ExcPending
	if v2401 != 0 {
		goto L4
	} else {
		goto L422
	}
L418:
	;
	v2391 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v2391 != 0 {
		goto L417
	} else {
		goto L419
	}
L419:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	if v2393 != 0 {
		goto L417
	} else {
		goto L420
	}
L420:
	;
	v2395 = *(*int32)(unsafe.Add(mBase, _consts[208]))
	if v2395 == int32(0) {
		goto L416
	} else {
		goto L421
	}
L421:
	;
	goto L417
L422:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L4
	} else {
		goto L423
	}
L423:
	;
	F_errmsg(m, int32(348325), int32(0))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L4
	} else {
		goto L424
	}
L424:
	;
	F_errfinish(m, int32(515261), int32(596), int32(249087))
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L4
	} else {
		goto L425
	}
L425:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L426:
	;
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v2448)+4))
	v2452 = v2449 + v2451
	v2453 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2452)+6)) = uint16(v2453)
	*(*uint8)(unsafe.Add(mBase, uint32(v2452)+4)) = uint8(v2453)
	*(*int32)(unsafe.Add(mBase, uint32(v2452))) = v2453
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2448)+4))
	v2460 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v2448)+4)) = v2459 + v2460
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v2450 - v2460
	v2467 = int32(4444060)
	v2469 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	*(*int32)(unsafe.Add(mBase, _consts[198])) = v2469 + v2460
	v2474 = *(*int32)(unsafe.Add(mBase, _consts[194]))
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v2474)))
	v2477 = v2469 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v2475)+4)) = v2477
	v2480 = int32(*(*uint16)(unsafe.Add(mBase, _consts[169])))
	v2482 = v2480 - int32(1)
	if base.Ui32(v2482&int32(65535)) <= base.Ui32(int32(65533)) {
		goto L432
	} else {
		goto L433
	}
L427:
	;
	v2419 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(v2419)))
	v2448 = v2419
	v2449 = v2420
	v2450 = v2415
	goto L426
L428:
	;
	goto L429
L429:
	;
	v2422 = F_palloc0(m, int32(12))
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L4
	} else {
		goto L430
	}
L430:
	;
	v2424 = int32(4444048)
	v2425 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v2425)+8)) = v2422
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v2422
	*(*int64)(unsafe.Add(mBase, uint32(v2422)+4)) = int64(0)
	v2432 = int32(512)
	*(*int32)(unsafe.Add(mBase, _consts[195])) = v2432
	v2434 = int32(4444052)
	v2436 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v2436 + int32(1)
	v2441 = F_palloc(m, v2432)
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L4
	} else {
		goto L431
	}
L431:
	;
	v2444 = *(*int32)(unsafe.Add(mBase, _consts[197]))
	*(*int32)(unsafe.Add(mBase, uint32(v2444))) = v2441
	v2447 = *(*int32)(unsafe.Add(mBase, _consts[195]))
	v2448 = v2444
	v2449 = v2441
	v2450 = v2447
	goto L426
L432:
	;
	v2488 = *(*int64)(unsafe.Add(mBase, _consts[172]))
	*(*int64)(unsafe.Add(mBase, uint32(v2475)+56)) = v2488
	v2491 = *(*int64)(unsafe.Add(mBase, _consts[173]))
	*(*int64)(unsafe.Add(mBase, uint32(v2475)+64)) = v2491
	goto L434
L433:
	;
	goto L434
L434:
	;
	if base.Ui32(v2477) < base.Ui32(int32(1073741824)) {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	*(*int64)(unsafe.Add(mBase, _consts[165])) = int64(0)
	v2675 = m.G0
	v2677 = v2675 - int32(32)
	m.G0 = v2677
	v2680 = F_TwoPhaseGetDummyProc(m, v113, int32(0))
	mBase = m.M
	v2681 = m.ExcPending
	if v2681 != 0 {
		goto L4
	} else {
		goto L463
	}
L436:
	;
	v2497 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	F_XLogEnsureRecordSpace(m, int32(0), v2497)
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L4
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2659 = m.ExcPending
	if v2659 != 0 {
		goto L4
	} else {
		goto L459
	}
L439:
	;
	v2500 = int32(4543684)
	v2502 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2503 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2502 + v2503
	v2507 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v2508 = *(*int32)(unsafe.Add(mBase, uint32(v2507)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v2507)+120)) = v2508 | v2503
	F_XLogBeginInsert(m)
	mBase = m.M
	v2513 = m.ExcPending
	if v2513 != 0 {
		goto L4
	} else {
		goto L440
	}
L440:
	;
	v2515 = *(*int32)(unsafe.Add(mBase, _consts[194]))
	if v2515 != 0 {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v2516 = v2515
	goto L444
L442:
	;
	goto L443
L443:
	;
	v2580 = int32(4444868)
	v2582 = int32(*(*uint8)(unsafe.Add(mBase, _consts[65])))
	v2583 = v2582 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[65])) = uint8(v2583)
	goto L448
L444:
	;
	v2545 = *(*int32)(unsafe.Add(mBase, uint32(v2516)))
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(v2516)+4))
	F_XLogRegisterData(m, v2545, v2546)
	mBase = m.M
	v2548 = m.ExcPending
	if v2548 != 0 {
		goto L4
	} else {
		goto L446
	}
L445:
	;
	goto L443
L446:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(v2516)+8))
	if v2549 != 0 {
		v2516 = v2549
		goto L444
	} else {
		goto L447
	}
L447:
	;
	goto L445
L448:
	;
	v2587 = F_XLogInsert(m, int32(1), int32(16))
	mBase = m.M
	v2588 = m.ExcPending
	if v2588 != 0 {
		goto L4
	} else {
		goto L449
	}
L449:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v421)+24)) = v2587
	if base.Ui32(v2482&int32(65535)) <= base.Ui32(int32(65533)) {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v2595 = *(*int64)(unsafe.Add(mBase, _consts[172]))
	F_replorigin_session_advance(m, v2595, v2587)
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		goto L4
	} else {
		goto L453
	}
L451:
	;
	v2599 = v2587
	goto L452
L452:
	;
	F_XLogFlush(m, v2599)
	mBase = m.M
	v2601 = m.ExcPending
	if v2601 != 0 {
		goto L4
	} else {
		goto L454
	}
L453:
	;
	v2598 = *(*int64)(unsafe.Add(mBase, uint32(v421)+24))
	v2599 = v2598
	goto L452
L454:
	;
	v2603 = *(*int64)(unsafe.Add(mBase, _consts[209]))
	*(*int64)(unsafe.Add(mBase, uint32(v421)+16)) = v2603
	v2606 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v2610 = F_LWLockAcquire(m, v2606+int32(2304), int32(0))
	mBase = m.M
	v2611 = m.ExcPending
	if v2611 != 0 {
		goto L4
	} else {
		goto L455
	}
L455:
	;
	v2612 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v421)+44)) = uint8(v2612)
	v2615 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v2615+int32(2304))
	mBase = m.M
	v2619 = m.ExcPending
	if v2619 != 0 {
		goto L4
	} else {
		goto L456
	}
L456:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v2621)))
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	F_ProcArrayAdd(m, v2622+v2623*int32(640))
	mBase = m.M
	v2628 = m.ExcPending
	if v2628 != 0 {
		goto L4
	} else {
		goto L457
	}
L457:
	;
	v2630 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v2631 = *(*int32)(unsafe.Add(mBase, uint32(v2630)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v2630)+120)) = v2631 & int32(-2)
	*(*int32)(unsafe.Add(mBase, _consts[210])) = v421
	v2637 = int32(4543684)
	v2639 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2639 - int32(1)
	v2643 = *(*int64)(unsafe.Add(mBase, uint32(v421)+24))
	F_SyncRepWaitForLSN(m, v2643, int32(0))
	mBase = m.M
	v2646 = m.ExcPending
	if v2646 != 0 {
		goto L4
	} else {
		goto L458
	}
L458:
	;
	v2648 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[197])) = v2648
	*(*int32)(unsafe.Add(mBase, _consts[194])) = v2648
	*(*int32)(unsafe.Add(mBase, _consts[196])) = v2648
	goto L435
L459:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		goto L4
	} else {
		goto L460
	}
L460:
	;
	F_errmsg(m, int32(480060), int32(0))
	mBase = m.M
	v2666 = m.ExcPending
	if v2666 != 0 {
		goto L4
	} else {
		goto L461
	}
L461:
	;
	F_errfinish(m, int32(519151), int32(1174), int32(380732))
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L4
	} else {
		goto L462
	}
L462:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L463:
	;
	v2682 = int32(4543684)
	v2684 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2684 + int32(1)
	v2691 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	F_hash_seq_init(m, v2677+int32(12), v2691)
	mBase = m.M
	v2693 = m.ExcPending
	if v2693 != 0 {
		goto L4
	} else {
		goto L464
	}
L464:
	;
	v2696 = F_hash_seq_search(m, v2677+int32(12))
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L4
	} else {
		goto L469
	}
L465:
	;
	v3583 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v3585 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v3589 = F_LWLockAcquire(m, v3585+int32(512), int32(0))
	mBase = m.M
	v3590 = m.ExcPending
	if v3590 != 0 {
		goto L4
	} else {
		goto L576
	}
L466:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3572 = m.ExcPending
	if v3572 != 0 {
		goto L4
	} else {
		goto L573
	}
L467:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3559 = m.ExcPending
	if v3559 != 0 {
		goto L4
	} else {
		goto L570
	}
L468:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3543 = m.ExcPending
	if v3543 != 0 {
		goto L4
	} else {
		goto L566
	}
L469:
	;
	if v2696 != 0 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v2702 = v2696
	goto L473
L471:
	;
	goto L472
L472:
	;
	v3035 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v3037 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v3040 = v3037
	v3045 = v3035
	v3046 = int32(0)
	goto L500
L473:
	;
	v2727 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+28))
	if v2727 == int32(0) {
		goto L476
	} else {
		goto L477
	}
L474:
	;
	goto L472
L475:
	;
	v3001 = F_hash_seq_search(m, v2677+int32(12))
	mBase = m.M
	v3002 = m.ExcPending
	if v3002 != 0 {
		goto L4
	} else {
		goto L498
	}
L476:
	;
	F_RemoveLocalLock(m, v2702)
	mBase = m.M
	v2969 = m.ExcPending
	if v2969 != 0 {
		goto L4
	} else {
		goto L497
	}
L477:
	;
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+24))
	if v2730 == int32(0) {
		goto L476
	} else {
		goto L478
	}
L478:
	;
	v2733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2702)+14)))
	if v2733 == int32(6) {
		goto L475
	} else {
		goto L479
	}
L479:
	;
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+40))
	v2738 = v2736 - int32(1)
	if v2738 < int32(0) {
		goto L475
	} else {
		goto L480
	}
L480:
	;
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+48))
	v2742 = int32(3)
	v2743 = v2736 & v2742
	if base.Ui32(v2738) < base.Ui32(v2742) {
		goto L482
	} else {
		goto L483
	}
L481:
	;
	if v2743 != 0 {
		goto L488
	} else {
		goto L489
	}
L482:
	;
	v2747 = int32(0)
	v2822 = v2747
	v2823 = v2747
	v2825 = v2738
	goto L481
L483:
	;
	goto L484
L484:
	;
	v2757 = int32(0)
	v2760 = v2757
	v2761 = v2757
	v2763 = v2738
	v2767 = v2757
	goto L485
L485:
	;
	v2789 = int32(4)
	v2790 = v2763 << (uint(v2789) % 32)
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v2741-int32(48)+v2790)))
	v2793 = int32(0)
	v2796 = *(*int32)(unsafe.Add(mBase, uint32(v2790+(v2741-int32(32)))))
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(v2790+(v2741-int32(16)))))
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v2790+v2741)))
	v2810 = base.B2i32(v2792 == v2793) | base.B2i32(v2796 == v2793) | base.B2i32(v2801 == v2793) | base.B2i32(v2806 == v2793) | v2761
	v2816 = base.B2i32(v2792|v2796|v2801|v2806 != v2793) | v2760
	v2818 = v2763 - v2789
	v2820 = v2767 + v2789
	if v2820 != v2736&int32(-4) {
		v2760 = v2816
		v2761 = v2810
		v2763 = v2818
		v2767 = v2820
		goto L485
	} else {
		goto L487
	}
L486:
	;
	v2822 = v2816
	v2823 = v2810
	v2825 = v2818
	goto L481
L487:
	;
	goto L486
L488:
	;
	v2851 = v2822
	v2852 = v2823
	v2854 = v2825
	v2857 = int32(0)
	goto L491
L489:
	;
	v2895 = v2822
	v2896 = v2823
	goto L490
L490:
	;
	if v2895&int32(1) == int32(0) {
		goto L475
	} else {
		goto L494
	}
L491:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2741+v2854<<(uint(int32(4))%32))))
	v2884 = int32(0)
	v2886 = base.B2i32(v2883 == v2884) | v2852
	v2889 = base.B2i32(v2883 != v2884) | v2851
	v2890 = int32(1)
	v2893 = v2857 + v2890
	if v2893 != v2743 {
		v2851 = v2889
		v2852 = v2886
		v2854 = v2854 - v2890
		v2857 = v2893
		goto L491
	} else {
		goto L493
	}
L492:
	;
	v2895 = v2889
	v2896 = v2886
	goto L490
L493:
	;
	goto L492
L494:
	;
	if v2896&int32(1) != 0 {
		goto L468
	} else {
		goto L495
	}
L495:
	;
	v2930 = *(*int64)(unsafe.Add(mBase, uint32(v2702)+32))
	if v2930 <= int64(0) {
		goto L476
	} else {
		goto L496
	}
L496:
	;
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v2727)+16))
	v2935 = *(*int32)(unsafe.Add(mBase, uint32(v2702)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2727)+16)) = v2933 | int32(1)<<(uint(v2935)%32)
	goto L476
L497:
	;
	goto L475
L498:
	;
	if v3001 != 0 {
		v2702 = v3001
		goto L473
	} else {
		goto L499
	}
L499:
	;
	goto L474
L500:
	;
	v3069 = v3046 << (uint(int32(3)) % 32)
	v3070 = v3040 + v3069
	v3071 = *(*int32)(unsafe.Add(mBase, uint32(v3070)+152))
	if v3071 == int32(0) {
		v3499 = v3040
		v3504 = v3045
		goto L502
	} else {
		goto L503
	}
L501:
	;
	v3531 = int32(4543684)
	v3533 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v3533 - int32(1)
	m.G0 = v2677 + int32(32)
	goto L465
L502:
	;
	v3528 = v3046 + int32(1)
	if v3528 != int32(16) {
		v3040 = v3499
		v3045 = v3504
		v3046 = v3528
		goto L500
	} else {
		goto L565
	}
L503:
	;
	v3075 = v3070 + int32(148)
	if v3071 == v3075 {
		v3499 = v3040
		v3504 = v3045
		goto L502
	} else {
		goto L504
	}
L504:
	;
	v3081 = v3045 + v3046<<(uint(int32(7))%32) + int32(23296)
	v3083 = F_LWLockAcquire(m, v3081, int32(0))
	mBase = m.M
	v3084 = m.ExcPending
	if v3084 != 0 {
		goto L4
	} else {
		goto L505
	}
L505:
	;
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v3075)+4))
	if v3085 == int32(0) {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	F_LWLockRelease(m, v3081)
	mBase = m.M
	v3493 = m.ExcPending
	if v3493 != 0 {
		goto L4
	} else {
		goto L564
	}
L507:
	;
	if v3085 == v3075 {
		goto L506
	} else {
		goto L508
	}
L508:
	;
	v3089 = v3069 + (v2680 + int32(148))
	v3090 = v3085
	goto L509
L509:
	;
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(v3090)+4))
	v3121 = v3090 - int32(28)
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(v3121)))
	v3123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3122)+14)))
	if v3123 == int32(6) {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	goto L506
L511:
	;
	if v3119 != v3075 {
		v3090 = v3119
		goto L509
	} else {
		goto L563
	}
L512:
	;
	v3128 = *(*int32)(unsafe.Add(mBase, uint32(v3090-int32(12))))
	if v3128 == int32(0) {
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v3133 = *(*int32)(unsafe.Add(mBase, uint32(v3090-int32(16))))
	if v3128 != v3133 {
		goto L467
	} else {
		goto L514
	}
L514:
	;
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(v3090)))
	*(*int32)(unsafe.Add(mBase, uint32(v3135)+4)) = v3119
	v3137 = *(*int32)(unsafe.Add(mBase, uint32(v3090)))
	*(*int32)(unsafe.Add(mBase, uint32(v3119))) = v3137
	*(*int32)(unsafe.Add(mBase, uint32(v2677)+8)) = v2680
	*(*int32)(unsafe.Add(mBase, uint32(v2677)+4)) = v3122
	*(*int32)(unsafe.Add(mBase, uint32(v3090-int32(20)))) = v2680
	v3145 = *(*int32)(unsafe.Add(mBase, _consts[203]))
	v3147 = v2677 + int32(4)
	v3148 = m.G0
	v3150 = v3148 - int32(32)
	m.G0 = v3150
	v3152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3145)+34)))
	if v3152 != int32(1) {
		goto L518
	} else {
		goto L519
	}
L515:
	;
	if v3356 == int32(0) {
		goto L466
	} else {
		goto L559
	}
L516:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L4
	} else {
		goto L556
	}
L517:
	;
	F_hash_corrupted(m, v3145)
	mBase = m.M
	v3404 = m.ExcPending
	if v3404 != 0 {
		goto L4
	} else {
		goto L555
	}
L518:
	;
	v3155 = *(*int32)(unsafe.Add(mBase, uint32(v3145)))
	v3156 = *(*int32)(unsafe.Add(mBase, uint32(v3155)+396))
	v3158 = v3121 - int32(4)
	v3159 = *(*int32)(unsafe.Add(mBase, uint32(v3158)))
	v3160 = v3156 & v3159
	v3161 = *(*int32)(unsafe.Add(mBase, uint32(v3155)+392))
	if base.Ui32(v3161) < base.Ui32(v3160) {
		goto L521
	} else {
		goto L522
	}
L519:
	;
	goto L520
L520:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3363 = m.ExcPending
	if v3363 != 0 {
		goto L4
	} else {
		goto L552
	}
L521:
	;
	v3163 = *(*int32)(unsafe.Add(mBase, uint32(v3155)+400))
	v3165 = v3163 & v3160
	goto L523
L522:
	;
	v3165 = v3160
	goto L523
L523:
	;
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+4))
	v3167 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+44))
	v3172 = *(*int32)(unsafe.Add(mBase, uint32(v3166+int32(base.Ui32(v3165)>>(uint(v3167)%32))<<(uint(int32(2))%32))))
	if v3172 == int32(0) {
		goto L517
	} else {
		goto L524
	}
L524:
	;
	v3176 = v3121 - int32(8)
	v3177 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+40))
	v3193 = v3172 + (v3177-int32(1))&v3165<<(uint(int32(2))%32)
	goto L525
L525:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v3193)))
	if v3213 != v3176 {
		goto L527
	} else {
		goto L528
	}
L526:
	;
	if v3213 == int32(0) {
		goto L516
	} else {
		goto L531
	}
L527:
	;
	v3216 = v3213
	goto L529
L528:
	;
	v3216 = int32(0)
	goto L529
L529:
	;
	if v3216 != 0 {
		v3193 = v3213
		goto L525
	} else {
		goto L530
	}
L530:
	;
	goto L526
L531:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+36))
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+8))
	v3221 = m.T0[v3220].(func(*base.Module, int32, int32) int32)(m, v3147, v3219)
	mBase = m.M
	v3222 = m.ExcPending
	if v3222 != 0 {
		goto L4
	} else {
		goto L532
	}
L532:
	;
	v3223 = *(*int32)(unsafe.Add(mBase, uint32(v3145)))
	v3224 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+396))
	v3225 = v3221 & v3224
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+392))
	if base.Ui32(v3226) < base.Ui32(v3225) {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v3228 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+400))
	v3230 = v3228 & v3225
	goto L535
L534:
	;
	v3230 = v3225
	goto L535
L535:
	;
	v3231 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+4))
	v3232 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+44))
	v3237 = *(*int32)(unsafe.Add(mBase, uint32(v3231+int32(base.Ui32(v3230)>>(uint(v3232)%32))<<(uint(int32(2))%32))))
	if v3237 == int32(0) {
		goto L517
	} else {
		goto L536
	}
L536:
	;
	v3240 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+36))
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+40))
	v3247 = v3237 + (v3241-int32(1))&v3230<<(uint(int32(2))%32)
	v3248 = *(*int32)(unsafe.Add(mBase, uint32(v3247)))
	if v3248 != 0 {
		goto L538
	} else {
		goto L539
	}
L537:
	;
	m.G0 = v3150 + int32(32)
	goto L515
L538:
	;
	v3249 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+12))
	v3267 = v3248
	goto L541
L539:
	;
	v3296 = v3247
	goto L540
L540:
	;
	if v3230 != v3165 {
		goto L548
	} else {
		goto L549
	}
L541:
	;
	v3279 = *(*int32)(unsafe.Add(mBase, uint32(v3267)+4))
	if v3279 != v3221 {
		goto L543
	} else {
		goto L544
	}
L542:
	;
	v3296 = v3267
	goto L540
L543:
	;
	v3286 = *(*int32)(unsafe.Add(mBase, uint32(v3267)))
	if v3286 != 0 {
		v3267 = v3286
		goto L541
	} else {
		goto L547
	}
L544:
	;
	v3283 = m.T0[v3249].(func(*base.Module, int32, int32, int32) int32)(m, v3267+int32(8), v3147, v3240)
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L4
	} else {
		goto L545
	}
L545:
	;
	if v3283 != 0 {
		goto L543
	} else {
		goto L546
	}
L546:
	;
	v3356 = int32(0)
	goto L537
L547:
	;
	goto L542
L548:
	;
	v3317 = *(*int32)(unsafe.Add(mBase, uint32(v3176)))
	*(*int32)(unsafe.Add(mBase, uint32(v3193))) = v3317
	*(*int32)(unsafe.Add(mBase, uint32(v3296))) = v3176
	*(*int32)(unsafe.Add(mBase, uint32(v3176))) = int32(0)
	goto L550
L549:
	;
	goto L550
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3158))) = v3221
	v3323 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+16))
	v3324 = m.T0[v3323].(func(*base.Module, int32, int32, int32) int32)(m, v3121, v3147, v3240)
	mBase = m.M
	v3325 = m.ExcPending
	if v3325 != 0 {
		goto L4
	} else {
		goto L551
	}
L551:
	;
	v3356 = int32(1)
	goto L537
L552:
	;
	v3364 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v3150))) = v3364
	F_errmsg_internal(m, int32(745509), v3150)
	mBase = m.M
	v3368 = m.ExcPending
	if v3368 != 0 {
		goto L4
	} else {
		goto L553
	}
L553:
	;
	F_errfinish(m, int32(518333), int32(1170), int32(21569))
	mBase = m.M
	v3373 = m.ExcPending
	if v3373 != 0 {
		goto L4
	} else {
		goto L554
	}
L554:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L555:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L556:
	;
	v3409 = *(*int32)(unsafe.Add(mBase, uint32(v3145)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v3150)+16)) = v3409
	F_errmsg_internal(m, int32(745413), v3150+int32(16))
	mBase = m.M
	v3415 = m.ExcPending
	if v3415 != 0 {
		goto L4
	} else {
		goto L557
	}
L557:
	;
	F_errfinish(m, int32(518333), int32(1191), int32(21569))
	mBase = m.M
	v3420 = m.ExcPending
	if v3420 != 0 {
		goto L4
	} else {
		goto L558
	}
L558:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L559:
	;
	v3423 = *(*int32)(unsafe.Add(mBase, uint32(v3089)+4))
	if v3423 == int32(0) {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3089)+4)) = v3089
	*(*int32)(unsafe.Add(mBase, uint32(v3089))) = v3089
	goto L562
L561:
	;
	goto L562
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3090)+4)) = v3089
	v3429 = *(*int32)(unsafe.Add(mBase, uint32(v3089)))
	*(*int32)(unsafe.Add(mBase, uint32(v3090))) = v3429
	*(*int32)(unsafe.Add(mBase, uint32(v3429)+4)) = v3090
	*(*int32)(unsafe.Add(mBase, uint32(v3089))) = v3090
	goto L511
L563:
	;
	goto L510
L564:
	;
	v3495 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v3497 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v3499 = v3497
	v3504 = v3495
	goto L502
L565:
	;
	goto L501
L566:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L4
	} else {
		goto L567
	}
L567:
	;
	F_errmsg(m, int32(118056), int32(0))
	mBase = m.M
	v3550 = m.ExcPending
	if v3550 != 0 {
		goto L4
	} else {
		goto L568
	}
L568:
	;
	F_errfinish(m, int32(518098), int32(3610), int32(162788))
	mBase = m.M
	v3555 = m.ExcPending
	if v3555 != 0 {
		goto L4
	} else {
		goto L569
	}
L569:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L570:
	;
	F_errmsg_internal(m, int32(380016), int32(0))
	mBase = m.M
	v3563 = m.ExcPending
	if v3563 != 0 {
		goto L4
	} else {
		goto L571
	}
L571:
	;
	F_errfinish(m, int32(518098), int32(3669), int32(162788))
	mBase = m.M
	v3568 = m.ExcPending
	if v3568 != 0 {
		goto L4
	} else {
		goto L572
	}
L572:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L573:
	;
	F_errmsg_internal(m, int32(162543), int32(0))
	mBase = m.M
	v3576 = m.ExcPending
	if v3576 != 0 {
		goto L4
	} else {
		goto L574
	}
L574:
	;
	F_errfinish(m, int32(518098), int32(3707), int32(162788))
	mBase = m.M
	v3581 = m.ExcPending
	if v3581 != 0 {
		goto L4
	} else {
		goto L575
	}
L575:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L576:
	;
	v3592 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v3593 = *(*int32)(unsafe.Add(mBase, uint32(v3592)+4))
	v3594 = *(*int32)(unsafe.Add(mBase, uint32(v3583)+48))
	v3598 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3593+v3594<<(uint(int32(2))%32)))) = v3598
	*(*int32)(unsafe.Add(mBase, uint32(v3583)+56)) = v3598
	*(*int64)(unsafe.Add(mBase, uint32(v3583)+36)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3583)+73)) = uint8(v3598)
	v3607 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v3608 = *(*int64)(unsafe.Add(mBase, uint32(v3607)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v3607)+56)) = v3608 + int64(1)
	v3612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3583)+276)))
	if v3612 == v3598 {
		goto L578
	} else {
		goto L579
	}
L577:
	;
	v3636 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v3636+int32(512))
	mBase = m.M
	v3640 = m.ExcPending
	if v3640 != 0 {
		goto L4
	} else {
		goto L582
	}
L578:
	;
	v3615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3583)+277)))
	if v3615 != int32(1) {
		goto L577
	} else {
		goto L581
	}
L579:
	;
	goto L580
L580:
	;
	v3619 = v3594 << (uint(int32(1)) % 32)
	v3620 = int32(4472464)
	v3621 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v3621)+8))
	v3624 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3619+v3622))) = uint8(v3624)
	v3627 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v3628 = *(*int32)(unsafe.Add(mBase, uint32(v3627)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v3628+v3619)+1)) = uint8(v3624)
	*(*uint16)(unsafe.Add(mBase, uint32(v3583)+276)) = uint16(v3624)
	goto L577
L581:
	;
	goto L580
L582:
	;
	v3642 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	if v3642 != 0 {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v3646 = v3642
	goto L586
L584:
	;
	goto L585
L585:
	;
	v3708 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v3709 = int32(1)
	F_ResourceOwnerRelease(m, v3708, v3709, v3709, v3709)
	mBase = m.M
	v3713 = m.ExcPending
	if v3713 != 0 {
		goto L4
	} else {
		goto L590
	}
L586:
	;
	v3672 = *(*int32)(unsafe.Add(mBase, uint32(v3646)))
	v3674 = *(*int32)(unsafe.Add(mBase, uint32(v3646)+8))
	v3675 = *(*int32)(unsafe.Add(mBase, uint32(v3646)+4))
	m.T0[v3675].(func(*base.Module, int32, int32))(m, int32(4), v3674)
	mBase = m.M
	v3677 = m.ExcPending
	if v3677 != 0 {
		goto L4
	} else {
		goto L588
	}
L587:
	;
	goto L585
L588:
	;
	if v3672 != 0 {
		v3646 = v3672
		goto L586
	} else {
		goto L589
	}
L589:
	;
	goto L587
L590:
	;
	F_AtEOXact_Aio(m)
	mBase = m.M
	v3715 = m.ExcPending
	if v3715 != 0 {
		goto L4
	} else {
		goto L591
	}
L591:
	;
	F_AtEOXact_RelationCache(m, int32(1))
	mBase = m.M
	v3718 = m.ExcPending
	if v3718 != 0 {
		goto L4
	} else {
		goto L592
	}
L592:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v3720 = m.ExcPending
	if v3720 != 0 {
		goto L4
	} else {
		goto L593
	}
L593:
	;
	v3722 = *(*int32)(unsafe.Add(mBase, _consts[204]))
	if v3722 != 0 {
		goto L594
	} else {
		goto L595
	}
L594:
	;
	v3723 = *(*int32)(unsafe.Add(mBase, uint32(v3722)+20))
	if v3723 != 0 {
		goto L597
	} else {
		goto L598
	}
L595:
	;
	goto L596
L596:
	;
	*(*int32)(unsafe.Add(mBase, _consts[204])) = int32(0)
	F_pgstat_clear_snapshot(m)
	mBase = m.M
	v3819 = m.ExcPending
	if v3819 != 0 {
		goto L4
	} else {
		goto L603
	}
L597:
	;
	v3727 = v3723
	goto L600
L598:
	;
	goto L599
L599:
	;
	goto L596
L600:
	;
	v3753 = *(*int32)(unsafe.Add(mBase, uint32(v3727)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v3753)+8)) = int32(0)
	v3756 = *(*int32)(unsafe.Add(mBase, uint32(v3727)+68))
	if v3756 != 0 {
		v3727 = v3756
		goto L600
	} else {
		goto L602
	}
L601:
	;
	goto L599
L602:
	;
	goto L601
L603:
	;
	*(*int32)(unsafe.Add(mBase, _consts[116])) = int32(0)
	v3824 = *(*int32)(unsafe.Add(mBase, _consts[211]))
	if v3824 != 0 {
		goto L604
	} else {
		goto L605
	}
L604:
	;
	v3825 = *(*int32)(unsafe.Add(mBase, uint32(v3824)+20))
	v3826 = *(*int32)(unsafe.Add(mBase, uint32(v3824)+28))
	if v3825 < v3826 {
		goto L607
	} else {
		goto L608
	}
L605:
	;
	goto L606
L606:
	;
	v4000 = *(*int32)(unsafe.Add(mBase, _consts[212]))
	if v4000 != 0 {
		goto L621
	} else {
		goto L622
	}
L607:
	;
	v3831 = v3825
	goto L610
L608:
	;
	goto L609
L609:
	;
	v3896 = *(*int32)(unsafe.Add(mBase, uint32(v3824)+24))
	v3897 = *(*int32)(unsafe.Add(mBase, uint32(v3824)+32))
	if v3896 < v3897 {
		goto L614
	} else {
		goto L615
	}
L610:
	;
	v3858 = *(*int32)(unsafe.Add(mBase, _consts[213]))
	F_LocalExecuteInvalidationMessage(m, v3858+v3831<<(uint(int32(4))%32))
	mBase = m.M
	v3863 = m.ExcPending
	if v3863 != 0 {
		goto L4
	} else {
		goto L612
	}
L611:
	;
	goto L609
L612:
	;
	v3865 = v3831 + int32(1)
	if v3865 != v3826 {
		v3831 = v3865
		goto L610
	} else {
		goto L613
	}
L613:
	;
	goto L611
L614:
	;
	v3902 = v3896
	goto L617
L615:
	;
	goto L616
L616:
	;
	*(*int32)(unsafe.Add(mBase, _consts[211])) = int32(0)
	goto L606
L617:
	;
	v3929 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	F_LocalExecuteInvalidationMessage(m, v3929+v3902<<(uint(int32(4))%32))
	mBase = m.M
	v3934 = m.ExcPending
	if v3934 != 0 {
		goto L4
	} else {
		goto L619
	}
L618:
	;
	goto L616
L619:
	;
	v3936 = v3902 + int32(1)
	if v3936 != v3897 {
		v3902 = v3936
		goto L617
	} else {
		goto L620
	}
L620:
	;
	goto L618
L621:
	;
	v4004 = v4000
	goto L624
L622:
	;
	goto L623
L623:
	;
	v4065 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v4067 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v4071 = *(*int32)(unsafe.Add(mBase, uint32(v4065+v4067<<(uint(int32(2))%32))))
	if v4071 != 0 {
		goto L628
	} else {
		goto L629
	}
L624:
	;
	v4031 = *(*int32)(unsafe.Add(mBase, uint32(v4004)+24))
	*(*int32)(unsafe.Add(mBase, _consts[212])) = v4031
	F_pfree(m, v4004)
	mBase = m.M
	v4034 = m.ExcPending
	if v4034 != 0 {
		goto L4
	} else {
		goto L626
	}
L625:
	;
	goto L623
L626:
	;
	if v4031 != 0 {
		v4004 = v4031
		goto L624
	} else {
		goto L627
	}
L627:
	;
	goto L625
L628:
	;
	v4073 = F_TwoPhaseGetDummyProcNumber(m, v113, int32(0))
	mBase = m.M
	v4074 = m.ExcPending
	if v4074 != 0 {
		goto L4
	} else {
		goto L631
	}
L629:
	;
	v4103 = v4067
	goto L630
L630:
	;
	v4106 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	v4110 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4106+v4103<<(uint(int32(2))%32)))) = v4110
	v4113 = int32(4146096)
	*(*int32)(unsafe.Add(mBase, _consts[177])) = v4113
	*(*int32)(unsafe.Add(mBase, _consts[178])) = v4113
	*(*int32)(unsafe.Add(mBase, _consts[179])) = v4110
	*(*int32)(unsafe.Add(mBase, _consts[180])) = v4110
	v4125 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	if v4125 != 0 {
		goto L634
	} else {
		goto L635
	}
L631:
	;
	v4076 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v4080 = F_LWLockAcquire(m, v4076+int32(1664), int32(0))
	mBase = m.M
	v4081 = m.ExcPending
	if v4081 != 0 {
		goto L4
	} else {
		goto L632
	}
L632:
	;
	v4083 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v4084 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v4083+v4073<<(uint(v4084)%32)))) = v4071
	v4089 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	*(*int32)(unsafe.Add(mBase, uint32(v4083+v4089<<(uint(v4084)%32)))) = int32(0)
	v4096 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v4096+int32(1664))
	mBase = m.M
	v4100 = m.ExcPending
	if v4100 != 0 {
		goto L4
	} else {
		goto L633
	}
L633:
	;
	v4102 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v4103 = v4102
	goto L630
L634:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4125)+112)) = int64(-4294967296)
	v4129 = *(*int32)(unsafe.Add(mBase, _consts[215]))
	F_hash_destroy(m, v4129)
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L4
	} else {
		goto L637
	}
L635:
	;
	goto L636
L636:
	;
	v4142 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v4144 = int32(1)
	F_ResourceOwnerRelease(m, v4142, int32(2), v4144, v4144)
	mBase = m.M
	v4147 = m.ExcPending
	if v4147 != 0 {
		goto L4
	} else {
		goto L638
	}
L637:
	;
	v4133 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[200])) = v4133
	*(*int32)(unsafe.Add(mBase, _consts[215])) = v4133
	*(*uint8)(unsafe.Add(mBase, _consts[216])) = uint8(v4133)
	goto L636
L638:
	;
	v4149 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	v4151 = int32(1)
	F_ResourceOwnerRelease(m, v4149, int32(3), v4151, v4151)
	mBase = m.M
	v4154 = m.ExcPending
	if v4154 != 0 {
		goto L4
	} else {
		goto L639
	}
L639:
	;
	v4156 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v4160 = F_LWLockAcquire(m, v4156+int32(2304), int32(0))
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		goto L4
	} else {
		goto L640
	}
L640:
	;
	v4163 = *(*int32)(unsafe.Add(mBase, _consts[210]))
	*(*int32)(unsafe.Add(mBase, uint32(v4163)+40)) = int32(-1)
	v4167 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v4167+int32(2304))
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		goto L4
	} else {
		goto L641
	}
L641:
	;
	*(*int32)(unsafe.Add(mBase, _consts[210])) = int32(0)
	v4175 = int32(1)
	F_AtEOXact_GUC(m, v4175, v4175)
	mBase = m.M
	v4178 = m.ExcPending
	if v4178 != 0 {
		goto L4
	} else {
		goto L642
	}
L642:
	;
	F_AtEOXact_SPI(m, int32(1))
	mBase = m.M
	v4181 = m.ExcPending
	if v4181 != 0 {
		goto L4
	} else {
		goto L643
	}
L643:
	;
	v4183 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[184])) = v4183
	*(*int32)(unsafe.Add(mBase, _consts[185])) = v4183
	goto L644
L644:
	;
	F_AtEOXact_on_commit_actions(m, int32(1))
	mBase = m.M
	v4190 = m.ExcPending
	if v4190 != 0 {
		goto L4
	} else {
		goto L645
	}
L645:
	;
	F_AtEOXact_Namespace(m, int32(1), int32(0))
	mBase = m.M
	v4194 = m.ExcPending
	if v4194 != 0 {
		goto L4
	} else {
		goto L646
	}
L646:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v4196 = m.ExcPending
	if v4196 != 0 {
		goto L4
	} else {
		goto L647
	}
L647:
	;
	F_AtEOXact_Files(m, int32(1))
	mBase = m.M
	v4199 = m.ExcPending
	if v4199 != 0 {
		goto L4
	} else {
		goto L648
	}
L648:
	;
	v4201 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[68])) = v4201
	*(*int32)(unsafe.Add(mBase, _consts[186])) = v4201
	*(*int32)(unsafe.Add(mBase, _consts[187])) = v4201
	*(*int32)(unsafe.Add(mBase, _consts[188])) = v4201
	goto L649
L649:
	;
	F_AtEOXact_HashTables(m, int32(1))
	mBase = m.M
	v4214 = m.ExcPending
	if v4214 != 0 {
		goto L4
	} else {
		goto L650
	}
L650:
	;
	v4215 = int32(1)
	F_AtEOXact_Snapshot(m, v4215, v4215)
	mBase = m.M
	v4218 = m.ExcPending
	if v4218 != 0 {
		goto L4
	} else {
		goto L651
	}
L651:
	;
	F_AtEOXact_ApplyLauncher(m, int32(0))
	mBase = m.M
	v4221 = m.ExcPending
	if v4221 != 0 {
		goto L4
	} else {
		goto L652
	}
L652:
	;
	F_AtEOXact_LogicalRepWorkers(m, int32(0))
	mBase = m.M
	v4224 = m.ExcPending
	if v4224 != 0 {
		goto L4
	} else {
		goto L653
	}
L653:
	;
	v4228 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v4228 != int32(1) {
		goto L655
	} else {
		goto L656
	}
L654:
	;
	*(*int32)(unsafe.Add(mBase, _consts[175])) = int32(0)
	v4261 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	F_ResourceOwnerDelete(m, v4261)
	mBase = m.M
	v4263 = m.ExcPending
	if v4263 != 0 {
		goto L4
	} else {
		goto L658
	}
L655:
	;
	goto L654
L656:
	;
	v4232 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v4232 == int32(0) {
		goto L655
	} else {
		goto L657
	}
L657:
	;
	v4235 = int32(4543684)
	v4237 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v4238 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v4237 + v4238
	v4241 = *(*int32)(unsafe.Add(mBase, uint32(v4232)))
	*(*int32)(unsafe.Add(mBase, uint32(v4232))) = v4241 + v4238
	*(*int64)(unsafe.Add(mBase, uint32(v4232)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4232))) = v4241 + int32(2)
	v4252 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v4252 - v4238
	goto L655
L658:
	;
	v4264 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+40)) = v4264
	*(*int32)(unsafe.Add(mBase, _consts[176])) = v4264
	*(*int32)(unsafe.Add(mBase, _consts[189])) = v4264
	v4274 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v4275 = *(*int32)(unsafe.Add(mBase, uint32(v4274)+44))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v4275
	v4278 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	F_MemoryContextReset(m, v4278)
	mBase = m.M
	v4280 = m.ExcPending
	if v4280 != 0 {
		goto L4
	} else {
		goto L659
	}
L659:
	;
	v4282 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v4282
	*(*int32)(unsafe.Add(mBase, uint32(v4274)+36)) = v4282
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v4282
	v4288 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v4288
	*(*int64)(unsafe.Add(mBase, uint32(v35)+28)) = v4288
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v4282
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v4288
	*(*int64)(unsafe.Add(mBase, _consts[69])) = v4288
	*(*int32)(unsafe.Add(mBase, _consts[70])) = v4282
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v4282
	v4304 = int32(4543676)
	v4306 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v4306 - int32(1)
	m.G0 = v32 + int32(16)
	return
L660:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4319 = m.ExcPending
	if v4319 != 0 {
		goto L4
	} else {
		goto L661
	}
L661:
	;
	F_errmsg(m, int32(132286), int32(0))
	mBase = m.M
	v4323 = m.ExcPending
	if v4323 != 0 {
		goto L4
	} else {
		goto L662
	}
L662:
	;
	F_errfinish(m, int32(512898), int32(2616), int32(268746))
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		goto L4
	} else {
		goto L663
	}
L663:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L664:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4335 = m.ExcPending
	if v4335 != 0 {
		goto L4
	} else {
		goto L665
	}
L665:
	;
	F_errmsg(m, int32(126099), int32(0))
	mBase = m.M
	v4339 = m.ExcPending
	if v4339 != 0 {
		goto L4
	} else {
		goto L666
	}
L666:
	;
	F_errfinish(m, int32(512898), int32(2626), int32(268746))
	mBase = m.M
	v4344 = m.ExcPending
	if v4344 != 0 {
		goto L4
	} else {
		goto L667
	}
L667:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RecordTransactionAbort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v79 int64
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int64
	_ = v410
	var v411 int32
	_ = v411
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v430 int64
	_ = v430
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v2
	if v23 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L7
	} else {
		goto L110
	}
L2:
	;
	m.G0 = v19 + int32(16)
	return v475
L3:
	;
	if l0 != 0 {
		v475 = v2
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v31 = F_TransactionIdDidCommit(m, v23)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, _consts[165])) = int64(0)
	v475 = v2
	goto L2
L7:
	;
	return int32(0)
L8:
	;
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, _consts[169])))
	v40 = F_smgrGetPendingDeletes(m, int32(0), v19+int32(12))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	if v44 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v46 = v45
	goto L13
L12:
	;
	v46 = v2
	goto L13
L13:
	;
	v52 = F_pgstat_get_transactional_drops(m, int32(0), v19+int32(8))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v54 = int32(4543684)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v56 + int32(1)
	if l0 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v106 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	v107 = int32(0)
	v109 = F_XactLogAbortRecord(m, v102, v44, v46, v40, v103, v52, v104, v106, v107, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L22
	}
L16:
	;
	v63 = m.G0
	v64 = int32(16)
	v65 = v63 - v64
	m.G0 = v65
	F___gettimeofday(m, v65)
	mBase = m.M
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	v69 = int64(*(*int32)(unsafe.Add(mBase, uint32(v65)+8)))
	m.G0 = v65 + v64
	goto L19
L17:
	;
	goto L18
L18:
	;
	v79 = *(*int64)(unsafe.Add(mBase, _consts[170]))
	if v79 != int64(0) {
		v102 = v79
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v102 = v69 + v68*int64(1000000) - int64(946684800000000)
	goto L15
L20:
	;
	v86 = m.G0
	v87 = int32(16)
	v88 = v86 - v87
	m.G0 = v88
	F___gettimeofday(m, v88)
	mBase = m.M
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	v92 = int64(*(*int32)(unsafe.Add(mBase, uint32(v88)+8)))
	m.G0 = v88 + v87
	v100 = v92 + v91*int64(1000000) - int64(946684800000000)
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, _consts[170])) = v100
	v102 = v100
	goto L15
L22:
	;
	if base.Ui32((v36-int32(1))&int32(65535)) <= base.Ui32(int32(65533)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v116 = *(*int64)(unsafe.Add(mBase, _consts[172]))
	v118 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	F_replorigin_session_advance(m, v116, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if l0 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v124 = *(*int64)(unsafe.Add(mBase, _consts[165]))
	F_XLogSetAsyncXactLSN(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_TransactionIdAbortTree(m, v23, v44, v46)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v129 = int32(4543684)
	v131 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v132 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v131 - v132
	v139 = v44 - v132
	if v139 < int32(0) {
		v206 = v23
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if l0 != 0 {
		goto L64
	} else {
		goto L65
	}
L33:
	;
	goto L32
L34:
	;
	if v44&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v144 = int32(2)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v46+v139<<(uint(v144)%32))))
	if base.B2i32(base.Ui32(v144) < base.Ui32(v147))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v23)) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v162 = v23
	v164 = v139
	goto L37
L37:
	;
	if v139 == int32(0) {
		v206 = v162
		goto L33
	} else {
		goto L45
	}
L38:
	;
	v162 = v159
	v164 = v44 - int32(2)
	goto L37
L39:
	;
	v159 = v147
	goto L38
L40:
	;
	if base.Ui32(v23) < base.Ui32(v147) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if int32(0) <= v23-v147 {
		v159 = v23
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v159 = v23
	goto L38
L44:
	;
	goto L39
L45:
	;
	v169 = v162
	v170 = v164
	goto L46
L46:
	;
	v176 = v170 << (uint(int32(2)) % 32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v46+v176)))
	if base.Ui32(v169) < base.Ui32(int32(3)) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v206 = v201
	goto L33
L48:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v176+(v46-int32(4)))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v189))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v187)) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L49:
	;
	v187 = v178
	goto L48
L50:
	;
	if base.Ui32(v178) <= base.Ui32(v169) {
		v187 = v169
		goto L48
	} else {
		goto L54
	}
L51:
	;
	if base.Ui32(v178) < base.Ui32(int32(3)) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	if v169-v178 < int32(0) {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v187 = v169
	goto L48
L54:
	;
	goto L49
L55:
	;
	if int32(1) < v170 {
		v169 = v201
		v170 = v170 - int32(2)
		goto L46
	} else {
		goto L62
	}
L56:
	;
	v201 = v189
	goto L55
L57:
	;
	if base.Ui32(v187) < base.Ui32(v189) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if int32(0) <= v187-v189 {
		v201 = v187
		goto L55
	} else {
		goto L61
	}
L60:
	;
	v201 = v187
	goto L55
L61:
	;
	goto L56
L62:
	;
	goto L47
L63:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v462 != 0 {
		goto L104
	} else {
		goto L105
	}
L64:
	;
	v212 = m.G0
	v214 = v212 - int32(32)
	m.G0 = v214
	v217 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v221 = F_LWLockAcquire(m, v217+int32(512), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int64)(unsafe.Add(mBase, _consts[165])) = int64(0)
	goto L63
L67:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	v227 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+48))
	v229 = int32(1)
	v231 = v225 + v228<<(uint(v229)%32)
	v233 = v44 - v229
	if int32(0) <= v233 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v236 = v233
	goto L71
L69:
	;
	v336 = v227
	goto L70
L70:
	;
	v346 = v336 + int32(280)
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+276)))
	v348 = v347
	goto L87
L71:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v255 = v253 + int32(280)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v46+v236<<(uint(int32(2))%32))))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+276)))
	v261 = v260
	goto L75
L72:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v336 = v328
	goto L70
L73:
	;
	if int32(0) < v236 {
		v236 = v236 - int32(1)
		goto L71
	} else {
		goto L84
	}
L74:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+277)))
	if v303 != 0 {
		goto L73
	} else {
		goto L79
	}
L75:
	;
	if v261 == int32(0) {
		goto L74
	} else {
		goto L77
	}
L76:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v260<<(uint(int32(2))%32)+v255-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v291
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	v294 = int32(1)
	v295 = v293 - v294
	*(*uint8)(unsafe.Add(mBase, uint32(v231))) = uint8(v295)
	v298 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+276)))
	v301 = v299 - v294
	*(*uint8)(unsafe.Add(mBase, uint32(v298)+276)) = uint8(v301)
	goto L73
L77:
	;
	v280 = v261 - int32(1)
	v283 = v255 + v280<<(uint(int32(2))%32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v284 != v259 {
		v261 = v280
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v306 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	if v306 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+16)) = v259
	F_errmsg_internal(m, int32(507434), v214+int32(16))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(511574), int32(4046), int32(182362))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	goto L73
L84:
	;
	goto L72
L85:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v409)+48))
	v411 = base.I32_wrap_i64(v410)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v206))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v411)) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+277)))
	if v390 != 0 {
		goto L85
	} else {
		goto L91
	}
L87:
	;
	if v348 == int32(0) {
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v347<<(uint(int32(2))%32)+v346-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v370))) = v378
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	v381 = int32(1)
	v382 = v380 - v381
	*(*uint8)(unsafe.Add(mBase, uint32(v231))) = uint8(v382)
	v385 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+276)))
	v388 = v386 - v381
	*(*uint8)(unsafe.Add(mBase, uint32(v385)+276)) = uint8(v388)
	goto L85
L89:
	;
	v367 = v348 - int32(1)
	v370 = v346 + v367<<(uint(int32(2))%32)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	if v371 != v23 {
		v348 = v367
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v393 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	if v393 == int32(0) {
		goto L85
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v23
	F_errmsg_internal(m, int32(507434), v214)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(511574), int32(4062), int32(182362))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	goto L85
L96:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v423 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v423 = base.B2i32(base.Ui32(v411) < base.Ui32(v206))
	goto L96
L98:
	;
	goto L99
L99:
	;
	v423 = int32(base.Ui32(v411-v206) >> (uint(int32(31)) % 32))
	goto L96
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v425)+48)) = v410 + base.I64_extend_i32_s(v206-v411)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v425)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v425)+56)) = v430 + int64(1)
	v435 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v435+int32(512))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	m.G0 = v214 + int32(32)
	goto L63
L104:
	;
	F_pfree(m, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L7
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v52 == int32(0) {
		v475 = v206
		goto L2
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	F_pfree(m, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v475 = v206
	goto L2
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v23
	F_errmsg_internal(m, int32(458171), v19)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(512898), int32(1794), int32(86871))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SaveTransactionCharacteristics(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v3
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[143])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v6)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[144])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v9)
	return
}
func F_StartTransactionCommand(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(int32(19)) < base.Ui32(v10) {
		v49 = *(*int32)(unsafe.Add(mBase, _consts[141]))
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v49
		m.G0 = v6 + int32(16)
		return
	} else {
		if v10 != 0 {
			if int32(1)<<(uint(v10)%32)&int32(1011558) == int32(0) {
				v49 = *(*int32)(unsafe.Add(mBase, _consts[141]))
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v49
				m.G0 = v6 + int32(16)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					if base.Ui32(v23) <= base.Ui32(int32(19)) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(int32(2))%32))+uint32(_consts[142])))
						v33 = v32
					} else {
						v33 = int32(565099)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v33
					F_errmsg_internal(m, int32(196363), v6)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_errfinish(m, int32(512898), int32(3114), int32(446140))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			F_StartTransaction(m)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(1)
				v49 = *(*int32)(unsafe.Add(mBase, _consts[141]))
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v49
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
func F_TransactionIdFollows(m *base.Module, l0 int32, l1 int32) int32 {
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		return base.B2i32(base.Ui32(l1) < base.Ui32(l0))
	} else {
		return base.B2i32(int32(0) < l0-l1)
	}
}
func F_TransactionIdIsInProgress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[669]))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[385]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v15))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	v27 = base.B2i32(base.Ui32(l0) < base.Ui32(v15))
	goto L2
L4:
	;
	goto L5
L5:
	;
	v27 = int32(base.Ui32(l0-v15) >> (uint(int32(31)) % 32))
	goto L2
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[782]))
	if v29 == l0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if base.Ui32(l0) < base.Ui32(int32(3)) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v150 != 0 {
		goto L48
	} else {
		goto L49
	}
L9:
	;
	v150 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v41 == l0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v150 = int32(1)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v45 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v150 = v142
	goto L8
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v49 == int32(0) {
		v142 = int32(0)
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v113 = int32(0)
	v115 = v45 - int32(1)
	goto L38
L19:
	;
	v54 = v49
	goto L20
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	if v59 == int32(4) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v142 = int32(0)
	goto L15
L22:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v54)+80))
	if v106 != 0 {
		v54 = v106
		goto L20
	} else {
		goto L37
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v62 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v65 = int32(1)
	if l0 == v62 {
		v142 = v65
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v69 = v67 - int32(1)
	if v69 < int32(0) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v74 = int32(0)
	v76 = v69
	goto L27
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	v82 = int32(2)
	v83 = base.I32_div_s(v76-v74, v82)
	v84 = v83 + v74
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v80+v84<<(uint(v82)%32))))
	if v88 == l0 {
		v142 = v65
		goto L15
	} else {
		goto L29
	}
L28:
	;
	goto L22
L29:
	;
	v92 = F_TransactionIdPrecedes(m, v88, l0)
	mBase = m.M
	if v92 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v93 = v84 + int32(1)
	goto L32
L31:
	;
	v93 = v74
	goto L32
L32:
	;
	if v92 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v76
	goto L35
L34:
	;
	v96 = v84 - int32(1)
	goto L35
L35:
	;
	if v93 <= v96 {
		v74 = v93
		v76 = v96
		goto L27
	} else {
		goto L36
	}
L36:
	;
	goto L28
L37:
	;
	goto L21
L38:
	;
	v120 = int32(2)
	v121 = base.I32_div_s(v115-v113, v120)
	v122 = v121 + v113
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v111+v122<<(uint(v120)%32))))
	v127 = base.B2i32(v126 == l0)
	if v126 == l0 {
		v142 = v127
		goto L15
	} else {
		goto L40
	}
L39:
	;
	v142 = v127
	goto L15
L40:
	;
	v130 = base.B2i32(base.Ui32(v126) < base.Ui32(l0))
	if base.Ui32(v126) < base.Ui32(l0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v131 = v122 + int32(1)
	goto L43
L42:
	;
	v131 = v113
	goto L43
L43:
	;
	if base.Ui32(v126) < base.Ui32(l0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v134 = v115
	goto L46
L45:
	;
	v134 = v122 - int32(1)
	goto L46
L46:
	;
	if v131 <= v134 {
		v113 = v131
		v115 = v134
		goto L38
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	return int32(1)
L49:
	;
	goto L50
L50:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[783]))
	if v154 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v570+int32(512))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L66
	} else {
		goto L161
	}
L52:
	;
	v458 = *(*int32)(unsafe.Add(mBase, _consts[669]))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v458)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v459))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L131
	} else {
		goto L132
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L66
	} else {
		goto L126
	}
L54:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
	if v160 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	goto L56
L56:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	*(*int32)(unsafe.Add(mBase, _consts[784])) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v194 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v198 = F_LWLockAcquire(m, v194+int32(512), int32(1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	v182 = F_emscripten_builtin_malloc(m, v179<<(uint(int32(2))%32))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[783])) = v182
	if v182 == int32(0) {
		goto L53
	} else {
		goto L65
	}
L58:
	;
	if v170 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+316))
	v168 = base.B2i32(v166 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(v168)
	v170 = v168
	goto L61
L60:
	;
	v170 = int32(0)
	goto L61
L61:
	;
	goto L58
L62:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	v174 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	v179 = (v172 + v174) * int32(65)
	goto L57
L63:
	;
	goto L64
L64:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v179 = v178
	goto L57
L65:
	;
	goto L56
L66:
	;
	return int32(0)
L67:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+48))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v204)) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v216 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v216 = base.B2i32(base.Ui32(v204) < base.Ui32(l0))
	goto L68
L70:
	;
	goto L71
L71:
	;
	v216 = int32(base.Ui32(v204-l0) >> (uint(int32(31)) % 32))
	goto L68
L72:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v218+int32(512))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L66
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if int32(0) < v225 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	return int32(1)
L76:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+48))
	v236 = int32(0)
	v239 = v2
	goto L79
L77:
	;
	v346 = v2
	goto L78
L78:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, _consts[29])))
	if v354 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L79:
	;
	if v236 == v230 {
		v332 = v239
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v346 = v332
	goto L78
L81:
	;
	v339 = v236 + int32(1)
	if v339 != v225 {
		v236 = v339
		v239 = v332
		goto L79
	} else {
		goto L100
	}
L82:
	;
	v247 = v236 << (uint(int32(2)) % 32)
	v249 = *(*int32)(unsafe.Add(mBase, _consts[784]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247+v249)))
	if v251 == int32(0) {
		v332 = v239
		goto L81
	} else {
		goto L83
	}
L83:
	;
	if l0 == v251 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v256+int32(512))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L66
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v251))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	return int32(1)
L88:
	;
	if v274 != 0 {
		v332 = v239
		goto L81
	} else {
		goto L92
	}
L89:
	;
	v274 = base.B2i32(base.Ui32(l0) < base.Ui32(v251))
	goto L88
L90:
	;
	goto L91
L91:
	;
	v274 = int32(base.Ui32(l0-v251) >> (uint(int32(31)) % 32))
	goto L88
L92:
	;
	v277 = v192 + v236<<(uint(int32(1))%32)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v280 = *(*int32)(unsafe.Add(mBase, _consts[785]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v247+(v13+int32(36)))))
	v289 = v278
	goto L94
L93:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	if v316 != int32(1) {
		v332 = v239
		goto L81
	} else {
		goto L99
	}
L94:
	;
	if v289 <= int32(0) {
		goto L93
	} else {
		goto L96
	}
L95:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v309+int32(512))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L66
	} else {
		goto L98
	}
L96:
	;
	v302 = v289 - int32(1)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v280+v282*int32(640)+int32(280)+v302<<(uint(int32(2))%32))))
	if v306 != l0 {
		v289 = v302
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	return int32(1)
L99:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _consts[783]))
	*(*int32)(unsafe.Add(mBase, uint32(v320+v239<<(uint(int32(2))%32)))) = v251
	v332 = v239 + int32(1)
	goto L81
L100:
	;
	goto L80
L101:
	;
	if v364 == int32(0) {
		v562 = v346
		goto L51
	} else {
		goto L105
	}
L102:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[30]))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+316))
	v362 = base.B2i32(v360 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[29])) = uint8(v362)
	v364 = v362
	goto L104
L103:
	;
	v364 = int32(0)
	goto L104
L104:
	;
	goto L101
L105:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[669]))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+20))
	v371 = v369 - int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v368)+16))
	if v371 < v372 {
		goto L52
	} else {
		goto L106
	}
L106:
	;
	v375 = v372
	v376 = v371
	goto L107
L107:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	v387 = v375 + v376
	v388 = int32(2)
	v389 = base.I32_div_s(v387, v388)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v386+v389<<(uint(v388)%32))))
	if v393 != l0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	if v387 < int32(-1) {
		goto L52
	} else {
		goto L123
	}
L109:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v393))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	goto L111
L111:
	;
	goto L108
L112:
	;
	if v408 != 0 {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v408 = base.B2i32(base.Ui32(l0) < base.Ui32(v393))
	goto L112
L114:
	;
	goto L115
L115:
	;
	v408 = int32(base.Ui32(l0-v393) >> (uint(int32(31)) % 32))
	goto L112
L116:
	;
	v409 = v375
	goto L118
L117:
	;
	v409 = v389 + int32(1)
	goto L118
L118:
	;
	if v408 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v412 = v389 - int32(1)
	goto L121
L120:
	;
	v412 = v376
	goto L121
L121:
	;
	if v409 <= v412 {
		v375 = v409
		v376 = v412
		goto L107
	} else {
		goto L122
	}
L122:
	;
	goto L52
L123:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _consts[781]))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v389))))
	if v419 != int32(1) {
		goto L52
	} else {
		goto L124
	}
L124:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v423+int32(512))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L66
	} else {
		goto L125
	}
L125:
	;
	return int32(1)
L126:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L66
	} else {
		goto L127
	}
L127:
	;
	F_errmsg(m, int32(14012), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L66
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(511574), int32(1465), int32(135624))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L66
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	if v471 == int32(0) {
		v562 = v346
		goto L51
	} else {
		goto L134
	}
L131:
	;
	v471 = base.B2i32(base.Ui32(l0) <= base.Ui32(v459))
	goto L130
L132:
	;
	goto L133
L133:
	;
	v471 = base.B2i32(l0-v459 <= int32(0))
	goto L130
L134:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _consts[783]))
	v476 = int32(0)
	v479 = *(*int32)(unsafe.Add(mBase, _consts[669]))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+20))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v479)+16))
	if v480 <= v481 {
		v548 = v476
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v562 = v548
	goto L51
L136:
	;
	v484 = v481
	v485 = v476
	v487 = v476
	goto L137
L137:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _consts[781]))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495+v484))))
	if v497 == int32(1) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v548 = v540
	goto L135
L139:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v501+v484<<(uint(int32(2))%32))))
	if v485 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v540 = v485
	v542 = v487
	goto L141
L141:
	;
	v544 = v484 + int32(1)
	if v544 != v480 {
		v484 = v544
		v485 = v540
		v487 = v542
		goto L137
	} else {
		goto L160
	}
L142:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v487))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v505)) == int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v521 = v487
	goto L144
L144:
	;
	if l0 != 0 {
		goto L152
	} else {
		goto L153
	}
L145:
	;
	if v519 != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v519 = base.B2i32(base.Ui32(v505) < base.Ui32(v487))
	goto L145
L147:
	;
	goto L148
L148:
	;
	v519 = int32(base.Ui32(v505-v487) >> (uint(int32(31)) % 32))
	goto L145
L149:
	;
	v520 = v505
	goto L151
L150:
	;
	v520 = v487
	goto L151
L151:
	;
	v521 = v520
	goto L144
L152:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v505)) == int32(0) {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v475+v485<<(uint(int32(2))%32)))) = v505
	v540 = v485 + int32(1)
	v542 = v521
	goto L141
L155:
	;
	if v533 != 0 {
		v548 = v485
		goto L135
	} else {
		goto L159
	}
L156:
	;
	v533 = base.B2i32(base.Ui32(l0) <= base.Ui32(v505))
	goto L155
L157:
	;
	goto L158
L158:
	;
	v533 = base.B2i32(int32(0) <= v505-l0)
	goto L155
L159:
	;
	goto L154
L160:
	;
	goto L138
L161:
	;
	if v562 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v575 = F_TransactionIdDidAbort(m, l0)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L66
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, _consts[782])) = l0
	goto L1
L165:
	;
	goto L164
L166:
	;
	if v575 != 0 {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v577 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L66
	} else {
		goto L168
	}
L168:
	;
	if v577 == l0 {
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _consts[783]))
	v583 = int32(0)
	goto L170
L170:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v581+v583<<(uint(int32(2))%32))))
	v597 = base.B2i32(v577 == v596)
	if v597 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v597 == int32(0) {
		goto L165
	} else {
		goto L176
	}
L172:
	;
	v601 = v583 + int32(1)
	if v601 != v562 {
		v583 = v601
		goto L170
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	goto L171
L175:
	;
	goto L174
L176:
	;
	return int32(1)
}
func F_TransactionIdSetTreeStatus(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v39 int32
	_ = v39
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v351 int32
	_ = v351
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int64
	_ = v393
	var v394 int64
	_ = v394
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v398 int64
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int64
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v576 int32
	_ = v576
	var v590 int64
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v624 int32
	_ = v624
	var v627 int64
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int64
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v667 int32
	_ = v667
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v732 int64
	_ = v732
	var v737 int32
	_ = v737
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v828 int32
	_ = v828
	var v841 int64
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v875 int32
	_ = v875
	var v878 int64
	_ = v878
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int64
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	v6 = int32(0)
	v25 = int32(base.Ui32(l0) >> (uint(int32(15)) % 32))
	v26 = base.I64_extend_i32_u(v25)
	if l1 <= v6 {
		v72 = v6
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v557 = l1 - v72
	if l3 != int32(1) {
		goto L96
	} else {
		goto L97
	}
L2:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+28))
	v114 = int32(*(*uint16)(unsafe.Add(mBase, _consts[123])))
	v115 = base.I32_rem_u_s(base.I32_wrap_i64(v26), v114)
	v118 = v111 + v115<<(uint(int32(7))%32)
	if int32(5) < l1 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	if l1 != v72 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v39 = v6
	goto L5
L5:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2+v39<<(uint(int32(2))%32))))
	if int32(base.Ui32(v55)>>(uint(int32(15))%32)) != v25 {
		v72 = v39
		goto L3
	} else {
		goto L7
	}
L6:
	;
	goto L2
L7:
	;
	v60 = v39 + int32(1)
	if v60 != l1 {
		v39 = v60
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	goto L2
L10:
	;
	return
L11:
	;
	F_TransactionIdSetPageStatusInternal(m, l0, l1, l2, l3, l4, v26)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L38
	} else {
		goto L94
	}
L12:
	;
	v505 = F_LWLockAcquire(m, v118, int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L38
	} else {
		goto L93
	}
L13:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+36))
	if l0 != v123 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+276)))
	if l1 != v125 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if l1 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v128 = v122 + int32(280)
	v130 = l1 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v130) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	goto L18
L18:
	;
	v194 = F_LWLockConditionalAcquire(m, v118, int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	if v192 != 0 {
		goto L12
	} else {
		goto L37
	}
L20:
	;
	v192 = int32(0)
	goto L19
L21:
	;
	v166 = v161
	v167 = v162
	v168 = v163
	goto L31
L22:
	;
	if (l2|v128)&int32(3) != 0 {
		v161 = l2
		v162 = v128
		v163 = v130
		goto L21
	} else {
		goto L25
	}
L23:
	;
	v154 = l2
	v155 = v128
	v156 = v130
	goto L24
L24:
	;
	if v156 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L25:
	;
	v138 = l2
	v139 = v128
	v140 = v130
	goto L26
L26:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v143 != v144 {
		v161 = v138
		v162 = v139
		v163 = v140
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v154 = v149
	v155 = v147
	v156 = v151
	goto L24
L28:
	;
	v146 = int32(4)
	v147 = v139 + v146
	v149 = v138 + v146
	v151 = v140 - v146
	if base.Ui32(int32(3)) < base.Ui32(v151) {
		v138 = v149
		v139 = v147
		v140 = v151
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v161 = v154
	v162 = v155
	v163 = v156
	goto L21
L31:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v171 == v172 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v192 = v171 - v172
	goto L19
L33:
	;
	v174 = int32(1)
	v179 = v168 - v174
	if v179 != 0 {
		v166 = v166 + v174
		v167 = v167 + v174
		v168 = v179
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L20
L37:
	;
	goto L18
L38:
	;
	return
L39:
	;
	if v194 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v199 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	*(*int64)(unsafe.Add(mBase, uint32(v199)+576)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v199)+568)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v199)+564)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v199)+560)) = l0
	v204 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+552)) = uint8(v204)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v210 = v206
	goto L42
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+556)) = int32(-1)
	v479 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+552)) = uint8(v479)
	goto L12
L42:
	;
	if v210 != int32(-1) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+28))
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v199)+568))
	v332 = int64(*(*uint16)(unsafe.Add(mBase, _consts[123])))
	v333 = base.I64_rem_s(v330, v332)
	v337 = v329 + base.I32_wrap_i64(v333)<<(uint(int32(7))%32)
	v339 = F_LWLockAcquire(m, v337, int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L38
	} else {
		goto L65
	}
L44:
	;
	goto L43
L45:
	;
	v210 = v325
	goto L42
L46:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v234+v210*int32(640))+568))
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v199)+568))
	if v238 != v239 {
		goto L41
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v316 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+556)) = v316
	v319 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v322 = base.B2i32(v320 == v316)
	if v320 == v316 {
		goto L61
	} else {
		goto L62
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+556)) = v210
	v243 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v245 = base.B2i32(v244 == v210)
	if v244 == v210 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v246 = v243
	goto L52
L51:
	;
	v246 = v244
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+56)) = v246
	if v245 == int32(0) {
		v325 = v244
		goto L45
	} else {
		goto L53
	}
L53:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = int32(134217784)
	v256 = int32(0)
	goto L54
L54:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+552)))
	if v281 != 0 {
		v256 = v256 + int32(1)
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	v284 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v284
	if v256 <= v284 {
		goto L10
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v293 = v256
	goto L58
L58:
	;
	v312 = int32(1)
	if base.Ui32(v312) < base.Ui32(v293) {
		v293 = v293 - v312
		goto L58
	} else {
		goto L60
	}
L59:
	;
	goto L10
L60:
	;
	goto L59
L61:
	;
	v323 = v319
	goto L63
L62:
	;
	v323 = v320
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+56)) = v323
	if v320 == v316 {
		goto L44
	} else {
		goto L64
	}
L64:
	;
	v325 = v320
	goto L45
L65:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v342 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+56)) = v342
	if v341 != v342 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v349 = v337
	v350 = v330
	v351 = v341
	goto L69
L67:
	;
	v413 = v337
	goto L68
L68:
	;
	if v413 != 0 {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v374 = v371 + v351*int32(640)
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v374)+568))
	if v350 == v375 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v413 = v396
	goto L68
L71:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v374)+560))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+276)))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v374)+564))
	v404 = *(*int64)(unsafe.Add(mBase, uint32(v374)+576))
	F_TransactionIdSetPageStatusInternal(m, v399, v400, v374+int32(280), v403, v404, v398)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L38
	} else {
		goto L80
	}
L72:
	;
	v396 = v349
	v397 = v350
	v398 = v350
	goto L71
L73:
	;
	goto L74
L74:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)+28))
	v381 = int64(*(*uint16)(unsafe.Add(mBase, _consts[123])))
	v382 = base.I64_rem_s(v375, v381)
	v386 = v379 + base.I32_wrap_i64(v382)<<(uint(int32(7))%32)
	if v386 == v349 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v394 = v375
	goto L77
L76:
	;
	F_LWLockRelease(m, v349)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L38
	} else {
		goto L78
	}
L77:
	;
	v396 = v386
	v397 = v375
	v398 = v394
	goto L71
L78:
	;
	v391 = F_LWLockAcquire(m, v386, int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L38
	} else {
		goto L79
	}
L79:
	;
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v374)+568))
	v394 = v393
	goto L77
L80:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v374)+556))
	if v407 != int32(-1) {
		v349 = v396
		v350 = v397
		v351 = v407
		goto L69
	} else {
		goto L81
	}
L81:
	;
	goto L70
L82:
	;
	F_LWLockRelease(m, v413)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L38
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v341 == int32(-1) {
		goto L10
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v438 = v341
	goto L87
L87:
	;
	v461 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v465 = v462 + v438*int32(640)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+556))
	*(*int32)(unsafe.Add(mBase, uint32(v465)+556)) = int32(-1)
	v469 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v465)+552)) = uint8(v469)
	v472 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	if v472 != v465 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L10
L89:
	;
	goto L91
L90:
	;
	goto L91
L91:
	;
	if v466 != int32(-1) {
		v438 = v466
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	goto L11
L94:
	;
	F_LWLockRelease(m, v118)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L38
	} else {
		goto L95
	}
L95:
	;
	goto L10
L96:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+28))
	v798 = int32(*(*uint16)(unsafe.Add(mBase, _consts[123])))
	v799 = base.I32_rem_u_s(v25, v798)
	v802 = v796 + v799<<(uint(int32(7))%32)
	v804 = F_LWLockAcquire(m, v802, int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L38
	} else {
		goto L128
	}
L97:
	;
	if v557 <= int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v564 = l2 + v72<<(uint(int32(2))%32)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v576 = v6
	v590 = base.I64_extend_i32_u(int32(base.Ui32(v565) >> (uint(int32(15)) % 32)))
	goto L99
L99:
	;
	v593 = v576 + int32(1)
	if v593 < v557 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L96
L101:
	;
	v595 = v557
	goto L103
L102:
	;
	v595 = v593
	goto L103
L103:
	;
	v596 = v595 - v576
	v599 = v576
	v604 = int32(0)
	goto L105
L104:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+28))
	v644 = int32(*(*uint16)(unsafe.Add(mBase, _consts[123])))
	v645 = base.I32_rem_u_s(base.I32_wrap_i64(v590), v644)
	v648 = v641 + v645<<(uint(int32(7))%32)
	v650 = F_LWLockAcquire(m, v648, int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L38
	} else {
		goto L111
	}
L105:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v564+v599<<(uint(int32(2))%32))))
	v627 = base.I64_extend_i32_u(int32(base.Ui32(v624) >> (uint(int32(15)) % 32)))
	if v627 != v590 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v636 = v595
	v637 = v596
	v638 = v590
	goto L104
L107:
	;
	v636 = v599
	v637 = v604
	v638 = v627
	goto L104
L108:
	;
	goto L109
L109:
	;
	v629 = int32(1)
	v632 = v604 + v629
	if v632 != v596 {
		v599 = v599 + v629
		v604 = v632
		goto L105
	} else {
		goto L110
	}
L110:
	;
	goto L106
L111:
	;
	v654 = base.B2i32(l4 == int64(0))
	v656 = F_SimpleLruReadPage(m, int32(4443552), v590, v654, int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L38
	} else {
		goto L112
	}
L112:
	;
	if int32(0) < v637 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v667 = int32(0)
	goto L116
L114:
	;
	goto L115
L115:
	;
	v763 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v763)+12))
	v766 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v764+v656))) = uint8(v766)
	F_LWLockRelease(m, v648)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L38
	} else {
		goto L126
	}
L116:
	;
	v690 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+4))
	v692 = int32(2)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v691+v656<<(uint(v692)%32))))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v564+v576<<(uint(int32(2))%32)+v667<<(uint(v692)%32))))
	v701 = v699 & int32(32767)
	v704 = v695 + int32(base.Ui32(v701)>>(uint(v692)%32))
	v705 = int32(*(*int8)(unsafe.Add(mBase, uint32(v704))))
	v706 = int32(1)
	v709 = v699 << (uint(v706) % 32) & int32(6)
	v711 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
	if v711 == v706 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L115
L118:
	;
	v737 = v667 + int32(1)
	if v737 != v637 {
		v667 = v737
		goto L116
	} else {
		goto L125
	}
L119:
	;
	if v705>>(uint(v709)%32)&int32(3) == int32(1) {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v721 = v705 | int32(3)<<(uint(v709)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v704))) = uint8(v721)
	if l4 == int64(0) {
		goto L118
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v724)+36))
	v731 = v725 + (int32(base.Ui32(v701)>>(uint(int32(5))%32))|v656<<(uint(int32(10))%32))<<(uint(int32(3))%32)
	v732 = *(*int64)(unsafe.Add(mBase, uint32(v731)))
	if base.Ui64(l4) <= base.Ui64(v732) {
		goto L118
	} else {
		goto L124
	}
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v731))) = l4
	goto L118
L125:
	;
	goto L117
L126:
	;
	if v636 < v557 {
		v576 = v636
		v590 = v638
		goto L99
	} else {
		goto L127
	}
L127:
	;
	goto L100
L128:
	;
	F_TransactionIdSetPageStatusInternal(m, l0, v72, l2, l3, l4, v26)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L38
	} else {
		goto L129
	}
L129:
	;
	F_LWLockRelease(m, v802)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L38
	} else {
		goto L130
	}
L130:
	;
	if int32(0) < v557 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v814 = l2 + v72<<(uint(int32(2))%32)
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	v828 = int32(0)
	v841 = base.I64_extend_i32_u(int32(base.Ui32(v815) >> (uint(int32(15)) % 32)))
	goto L134
L132:
	;
	goto L133
L133:
	;
	return
L134:
	;
	v844 = v828 + int32(1)
	if v844 < v557 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L133
L136:
	;
	v846 = v557
	goto L138
L137:
	;
	v846 = v844
	goto L138
L138:
	;
	v847 = v846 - v828
	v850 = v828
	v855 = int32(0)
	goto L140
L139:
	;
	v891 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v891)+28))
	v895 = int32(*(*uint16)(unsafe.Add(mBase, _consts[123])))
	v896 = base.I32_rem_u_s(base.I32_wrap_i64(v841), v895)
	v899 = v892 + v896<<(uint(int32(7))%32)
	v901 = F_LWLockAcquire(m, v899, int32(0))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L38
	} else {
		goto L146
	}
L140:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v814+v850<<(uint(int32(2))%32))))
	v878 = base.I64_extend_i32_u(int32(base.Ui32(v875) >> (uint(int32(15)) % 32)))
	if v878 != v841 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v886 = v847
	v888 = v846
	v889 = v841
	goto L139
L142:
	;
	v886 = v855
	v888 = v850
	v889 = v878
	goto L139
L143:
	;
	goto L144
L144:
	;
	v880 = int32(1)
	v883 = v855 + v880
	if v883 != v847 {
		v850 = v850 + v880
		v855 = v883
		goto L140
	} else {
		goto L145
	}
L145:
	;
	goto L141
L146:
	;
	F_TransactionIdSetPageStatusInternal(m, int32(0), v886, v814+v828<<(uint(int32(2))%32), l3, l4, v841)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L38
	} else {
		goto L147
	}
L147:
	;
	F_LWLockRelease(m, v899)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L38
	} else {
		goto L148
	}
L148:
	;
	if v888 < v557 {
		v828 = v888
		v841 = v889
		goto L134
	} else {
		goto L149
	}
L149:
	;
	goto L135
}
func F_check_transaction_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_check_slru_buffers(m, int32(142628), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
