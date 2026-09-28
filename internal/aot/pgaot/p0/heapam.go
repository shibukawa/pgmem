package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_relation_copy_for_cluster(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v353 int32
	_ = v353
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v426 int32
	_ = v426
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int64
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v691 int32
	_ = v691
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int64
	_ = v735
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v978 int32
	_ = v978
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1138 int32
	_ = v1138
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1170 float64
	_ = v1170
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 float64
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1234 float64
	_ = v1234
	var v1238 float64
	_ = v1238
	var v1242 float64
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1249 float64
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 float64
	_ = v1259
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int64
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1329 float64
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1389 float64
	_ = v1389
	var v1390 int64
	_ = v1390
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1540 int64
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1594 int32
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1631 int32
	_ = v1631
	var v1636 int32
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1678 int32
	_ = v1678
	var v1707 float64
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1730 float64
	_ = v1730
	var v1734 int32
	_ = v1734
	var v1738 int32
	_ = v1738
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1771 int32
	_ = v1771
	var v1776 int32
	_ = v1776
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1954 int32
	_ = v1954
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1962 int32
	_ = v1962
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	v12 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v35) < base.Ui32(int32(_a_F_heapam_relation_copy_for_cluster_0)) {
		v44 = v34
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v47 = F_palloc_mul(m, int32(8), v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L1
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v39 == int32(99) {
		v44 = v34
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v42 = F_isTempToastNamespace(m, v39)
	mBase = m.M
	v44 = v42
	goto L2
L5:
	;
	return
L6:
	;
	v50 = F_palloc_mul(m, int32(1), v46)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if l5 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if l3 != 0 {
		goto L53
	} else {
		goto L54
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v56 = m.G0
	v58 = v56 - int32(112)
	m.G0 = v58
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	v66 = F_AllocSetContextCreateInternal(m, v61, int32(_a_F_heapam_relation_copy_for_cluster_1), int32(0), int32(_a_F_heapam_relation_copy_for_cluster_2), int32(_a_F_heapam_relation_copy_for_cluster_3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v205 = F_GetBulkInsertState(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L5
	} else {
		goto L48
	}
L12:
	;
	v68 = int32(_a_F_heapam_relation_copy_for_cluster_4)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v66
	v73 = F_palloc0(m, int32(72))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v75 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = l0
	v80 = F_RelationGetNumberOfBlocksInFork(m, l1, v75)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+40)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v73)+36)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v73)+28)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v80
	v88 = F_smgr_bulk_start_rel(m, l1, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v58)+16)) = int64(103079215116)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v73)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+44)) = v93
	v98 = v58 + int32(8)
	v100 = F_hash_create(m, int32(_a_F_heapam_relation_copy_for_cluster_5), int64(128), v98, int32(1064))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+56)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = int32(20)
	v108 = F_hash_create(m, int32(_a_F_heapam_relation_copy_for_cluster_6), int64(128), v98, int32(1064))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+60)) = v108
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v69
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[1]))
	if v114 <= int32(1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	m.G0 = v58 + int32(112)
	v209 = v73
	v214 = v12
	goto L8
L19:
	;
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+20)) = uint8(v197)
	goto L18
L20:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[2])))
	if v118&int32(1) == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+118)))
	if v125 != int32(112) {
		goto L19
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	if v114 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v123)+32))
	if v130 != 0 {
		goto L19
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v123)+56))
	goto L31
L28:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	if v131 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v154 = v58 + int32(60)
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[3]))
	v160 = F_LWLockAcquire(m, v156+int32(512), int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L38
	}
L31:
	;
	if base.Ui32(v132) < base.Ui32(int32(_a_F_heapam_relation_copy_for_cluster_0)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v135 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+20)) = uint8(v135)
	goto L30
L33:
	;
	goto L34
L34:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+180))
	if v138 == int32(0) {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+48))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+119)))
	switch v142 - int32(109) {
	case 0, 5:
		goto L36
	default:
		goto L19
	}
L36:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+20)) = uint8(v145)
	if v145 == int32(0) {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	goto L30
L38:
	;
	if v154 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[4]))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v164
	goto L41
L40:
	;
	goto L41
L41:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[3]))
	F_LWLockRelease(m, v167+int32(512))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v58)+60))
	if v172 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v175 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73+int32(20)))) = uint8(v175)
	goto L18
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+32)) = v172
	v178 = F_GetXLogInsertRecPtr(m)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v73)+48)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v58)+72)) = int64(4535485464580)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v73)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+100)) = v185
	v192 = F_hash_create(m, int32(_a_F_heapam_relation_copy_for_cluster_7), int64(128), v58-int32(-64), int32(1064))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+64)) = v192
	goto L18
L48:
	;
	v209 = v12
	v214 = v205
	goto L8
L49:
	;
	v670 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L5
	} else {
		goto L114
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L5
	} else {
		goto L111
	}
L51:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v573 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L52:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v464 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L53:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[6]))
	v218 = m.G0
	v220 = v218 - int32(16)
	m.G0 = v220
	v222 = int32(0)
	v224 = F_tuplesort_begin_common(m, v217, v222, v222)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if l2 != 0 {
		goto L51
	} else {
		goto L85
	}
L56:
	;
	v226 = int32(_a_F_heapam_relation_copy_for_cluster_4)
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v229
	v232 = F_palloc0(m, int32(12))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[7])))
	if v235 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l2)+192))
	v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v259)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+8)) = int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+40)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v224)+60)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(2049)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+16)) = int32(2050)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+12)) = int32(2051)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+4)) = int32(2052)
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = int32(2053)
	v275 = F_BuildIndexInfo(m, l2)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L64
	}
L59:
	;
	v240 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	if v240 == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v244)+120)))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+8)) = int32(102)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v245
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_8), v220)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_9), int32(276), int32(_a_F_heapam_relation_copy_for_cluster_10))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	goto L58
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v275
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+12)))
	v279 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(base.B2i32(v278 != v279))
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v31
	v284 = F__bt_mkscankey(m, l2, v279)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+76))
	if v287 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v288 = F_CreateExecutorState(m)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v224)+40))
	v306 = F_palloc0(m, v303*int32(36))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L75
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = v288
	v292 = F_MakeSingleTupleTableSlot(m, v31, int32(_a_F_heapam_relation_copy_for_cluster_11))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+152))
	if v295 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v298 = v295
	goto L73
L72:
	;
	v296 = F_MakePerTupleExprContext(m, v294)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L5
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v298)+4)) = v292
	goto L68
L74:
	;
	v298 = v296
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224)+44)) = v306
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v224)+40))
	if v310 <= int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_pfree(m, v284)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L5
	} else {
		goto L84
	}
L77:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v284)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = v316
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	v322 = int32(base.Ui32(v318)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+9)) = uint8(v322)
	v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+10)) = uint16(v324)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+20)) = uint8(v326)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	F_PrepareSortSupportFromIndexRel(m, l2, int32(base.Ui32(v328&int32(16777216))>>(uint(int32(24))%32)), v306)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v224)+40))
	if v335 < int32(2) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v353 = int32(1)
	goto L80
L80:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v224)+44))
	v369 = v366 + v353*int32(36)
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v371
	v375 = v284 + int32(16) + v353*int32(56)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v376
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v382 = int32(base.Ui32(v378)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v369)+9)) = uint8(v382)
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v375)+4)))
	v385 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v369)+20)) = uint8(v385)
	*(*uint16)(unsafe.Add(mBase, uint32(v369)+10)) = uint16(v384)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	F_PrepareSortSupportFromIndexRel(m, l2, int32(base.Ui32(v388&int32(16777216))>>(uint(int32(24))%32)), v369)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L5
	} else {
		goto L82
	}
L81:
	;
	goto L76
L82:
	;
	v396 = v353 + int32(1)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v224)+40))
	if v396 < v397 {
		v353 = v396
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v227
	m.G0 = v220 + int32(16)
	v459 = v224
	goto L52
L85:
	;
	v459 = int32(0)
	goto L52
L86:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[8]))
	if v507 != 0 {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	goto L86
L88:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v468&int32(1) == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v473 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v476 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v475 + v476
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	*(*int32)(unsafe.Add(mBase, uint32(v464))) = v479 + v476
	v483 = int32(0)
	v485 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v486 = base.AtomicRmwOr32(m, v483, v485, v483)
	*(*int64)(unsafe.Add(mBase, uint32(v464+int32(8))+232)) = int64(1)
	v494 = base.AtomicRmwOr32(m, v483, v485, v483)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	*(*int32)(unsafe.Add(mBase, uint32(v464))) = v495 + v476
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v501 - v476
	goto L87
L90:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[11])))
	if v509&int32(1) == int32(0) {
		goto L50
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if l5 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	v516 = l5
	goto L96
L95:
	;
	v516 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	goto L96
L96:
	;
	v517 = int32(0)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+8))
	v523 = m.T0[v522].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v516, v517, v517, v517, int32(449))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v525 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v523)+40)))
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v528 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v649 = v523
	v654 = int32(0)
	v655 = v459
	goto L49
L99:
	;
	goto L98
L100:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v532&int32(1) == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v537 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v540 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v539 + v540
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = v543 + v540
	v547 = int32(0)
	v549 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v550 = base.AtomicRmwOr32(m, v547, v549, v547)
	*(*int64)(unsafe.Add(mBase, uint32(v528+int32(56))+232)) = v525
	v558 = base.AtomicRmwOr32(m, v547, v549, v547)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = v559 + v540
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v565 - v540
	goto L99
L102:
	;
	if l5 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	goto L102
L104:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v577&int32(1) == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v582 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v585 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v584 + v585
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v588 + v585
	v592 = int32(0)
	v594 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v595 = base.AtomicRmwOr32(m, v592, v594, v592)
	*(*int64)(unsafe.Add(mBase, uint32(v573+int32(8))+232)) = int64(2)
	v603 = base.AtomicRmwOr32(m, v592, v594, v592)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = v604 + v585
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v610 - v585
	goto L103
L106:
	;
	v615 = l5
	goto L108
L107:
	;
	v615 = int32(_a_F_heapam_relation_copy_for_cluster_14)
	goto L108
L108:
	;
	v616 = int32(0)
	v620 = F_index_beginscan(m, l0, l2, v615, v616, v616, v616, v616)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v622 = int32(0)
	F_index_rescan(m, v620, v622, v622, v622, v622)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v649 = int32(0)
	v654 = v620
	v655 = v12
	goto L49
L111:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_15), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_16), int32(931), int32(_a_F_heapam_relation_copy_for_cluster_17))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	v691 = int32(-1)
	goto L117
L115:
	;
	if v670 != 0 {
		goto L324
	} else {
		goto L325
	}
L116:
	;
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+188))
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1584)+12))
	m.T0[v1585].(func(*base.Module, int32))(m, v649)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L5
	} else {
		goto L323
	}
L117:
	;
	v699 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[12]))
	if v699 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	F_index_endscan(m, v654)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L5
	} else {
		goto L321
	}
L119:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L5
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	if v654 != 0 {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	goto L121
L123:
	;
	goto L118
L124:
	;
	v836 = int32(0)
	v838 = F_ExecFetchSlotHeapTuple(m, v670, v836, v836)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L5
	} else {
		goto L147
	}
L125:
	;
	v703 = F_index_getnext_slot(m, v654, int32(1), v670)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L5
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+40)) = v724
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v727)+188))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)+20))
	v730 = m.T0[v729].(func(*base.Module, int32, int32, int32) int32)(m, v649, int32(1), v670)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L5
	} else {
		goto L134
	}
L128:
	;
	if v703 == int32(0) {
		goto L123
	} else {
		goto L129
	}
L129:
	;
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654)+72)))
	if v707 != int32(1) {
		v835 = v691
		goto L124
	} else {
		goto L130
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_18), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_19), int32(722), int32(_a_F_heapam_relation_copy_for_cluster_20))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	if v730 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v735 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v649)+40)))
	v738 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v738 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	goto L137
L137:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v649)+56))
	if v691 == v779 {
		v835 = v691
		goto L124
	} else {
		goto L142
	}
L138:
	;
	goto L116
L139:
	;
	goto L138
L140:
	;
	v742 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v742&int32(1) == int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v747 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v750 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v749 + v750
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v738)))
	*(*int32)(unsafe.Add(mBase, uint32(v738))) = v753 + v750
	v757 = int32(0)
	v759 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v760 = base.AtomicRmwOr32(m, v757, v759, v757)
	*(*int64)(unsafe.Add(mBase, uint32(v738+int32(64))+232)) = v735
	v768 = base.AtomicRmwOr32(m, v757, v759, v757)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v738)))
	*(*int32)(unsafe.Add(mBase, uint32(v738))) = v769 + v750
	v775 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v775 - v750
	goto L139
L142:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v649)+40))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v649)+44))
	v786 = base.I32_rem_u_s(v779+v782-v784, v782)
	v792 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v792 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v649)+56))
	v835 = v833
	goto L124
L144:
	;
	goto L143
L145:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v796&int32(1) == int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v801 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v803 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v804 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v803 + v804
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v792)))
	*(*int32)(unsafe.Add(mBase, uint32(v792))) = v807 + v804
	v811 = int32(0)
	v813 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v814 = base.AtomicRmwOr32(m, v811, v813, v811)
	*(*int64)(unsafe.Add(mBase, uint32(v792+int32(64))+232)) = base.I64_extend_i32_u(v786 + int32(1))
	v822 = base.AtomicRmwOr32(m, v811, v813, v811)
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v792)))
	*(*int32)(unsafe.Add(mBase, uint32(v792))) = v823 + v804
	v829 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v829 - v804
	goto L144
L147:
	;
	if l5 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	F_pfree(m, v1386)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L5
	} else {
		goto L303
	}
L149:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+56)) = int64(17179869187)
	v1376 = F_reform_tuple(m, v838, l0, l1, v47, v50)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L5
	} else {
		goto L300
	}
L150:
	;
	v1267 = m.G0
	v1269 = v1267 - int32(32)
	m.G0 = v1269
	v1271 = int32(_a_F_heapam_relation_copy_for_cluster_4)
	v1272 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0]))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v655)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v1274
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v655)+60))
	v1277 = F_heap_copytuple(m, v838)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L5
	} else {
		goto L282
	}
L151:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v670)+72))
	F_LockBufferInternal(m, v842, int32(3))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L5
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v1259 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v1259, float64(1))
	if v655 == int32(0) {
		goto L149
	} else {
		goto L281
	}
L154:
	;
	v846 = F_HeapTupleSatisfiesVacuum(m, v838, l4, v842)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L5
	} else {
		goto L161
	}
L155:
	;
	F_UnlockBuffer(m, v842)
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L5
	} else {
		goto L277
	}
L156:
	;
	v1242 = *(*float64)(unsafe.Add(mBase, uint32(l10)))
	*(*float64)(unsafe.Add(mBase, uint32(l10))) = base.F64_add(v1242, float64(1))
	goto L155
L157:
	;
	F_UnlockBuffer(m, v842)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L5
	} else {
		goto L266
	}
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L5
	} else {
		goto L263
	}
L159:
	;
	if v44 != 0 {
		goto L211
	} else {
		goto L212
	}
L160:
	;
	if v44 != 0 {
		goto L155
	} else {
		goto L162
	}
L161:
	;
	switch v846 {
	case 0:
		goto L157
	case 1:
		goto L155
	case 2:
		goto L156
	case 3:
		goto L160
	case 4:
		goto L159
	default:
		goto L158
	}
L162:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v838)+16))
	v849 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v848)+20)))
	v850 = int32(768)
	if v849&v850 != v850 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v848)))
	v856 = v854
	goto L165
L164:
	;
	v856 = int32(2)
	goto L165
L165:
	;
	if base.Ui32(v856) < base.Ui32(int32(3)) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	if v988 != 0 {
		goto L155
	} else {
		goto L206
	}
L167:
	;
	v988 = int32(0)
	goto L166
L168:
	;
	goto L169
L169:
	;
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[13]))
	if v868 == v856 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v988 = int32(1)
	goto L166
L171:
	;
	goto L172
L172:
	;
	v872 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[14]))
	if v872 <= int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v988 = v978
	goto L166
L174:
	;
	v876 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[15]))
	if v876 == int32(0) {
		v978 = int32(0)
		goto L173
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v946 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[16]))
	v948 = int32(0)
	v951 = v872 - int32(1)
	goto L196
L177:
	;
	v881 = v876
	goto L178
L178:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v881)+20))
	if v887 == int32(4) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v978 = int32(0)
	goto L173
L180:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v881)+80))
	if v941 != 0 {
		v881 = v941
		goto L178
	} else {
		goto L195
	}
L181:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	if v890 == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v893 = int32(1)
	if v856 == v890 {
		v978 = v893
		goto L173
	} else {
		goto L183
	}
L183:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v881)+52))
	v897 = v895 - int32(1)
	if v897 < int32(0) {
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v881)+48))
	v903 = int32(0)
	v906 = v897
	goto L185
L185:
	;
	v911 = int32(2)
	v912 = base.I32_div_s(v906-v903, v911)
	v913 = v912 + v903
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v900+v913<<(uint(v911)%32))))
	if v917 == v856 {
		v978 = v893
		goto L173
	} else {
		goto L187
	}
L186:
	;
	goto L180
L187:
	;
	v926 = base.B2i32(v917-v856 < int32(0)) | base.B2i32(base.Ui32(v917) < base.Ui32(int32(3)))
	if v926 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v927 = v913 + int32(1)
	goto L190
L189:
	;
	v927 = v903
	goto L190
L190:
	;
	if v926 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v930 = v906
	goto L193
L192:
	;
	v930 = v913 - int32(1)
	goto L193
L193:
	;
	if v927 <= v930 {
		v903 = v927
		v906 = v930
		goto L185
	} else {
		goto L194
	}
L194:
	;
	goto L186
L195:
	;
	goto L179
L196:
	;
	v956 = int32(2)
	v957 = base.I32_div_s(v951-v948, v956)
	v958 = v957 + v948
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v946+v958<<(uint(v956)%32))))
	v963 = base.B2i32(v962 == v856)
	if v962 == v856 {
		v978 = v963
		goto L173
	} else {
		goto L198
	}
L197:
	;
	v978 = v963
	goto L173
L198:
	;
	v966 = base.B2i32(base.Ui32(v962) < base.Ui32(v856))
	if base.Ui32(v962) < base.Ui32(v856) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v967 = v958 + int32(1)
	goto L201
L200:
	;
	v967 = v948
	goto L201
L201:
	;
	if base.Ui32(v962) < base.Ui32(v856) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v970 = v951
	goto L204
L203:
	;
	v970 = v958 - int32(1)
	goto L204
L204:
	;
	if v967 <= v970 {
		v948 = v967
		v951 = v970
		goto L196
	} else {
		goto L205
	}
L205:
	;
	goto L197
L206:
	;
	v991 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	if v991 == int32(0) {
		goto L155
	} else {
		goto L208
	}
L208:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v995 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_21), v29)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_19), int32(815), int32(_a_F_heapam_relation_copy_for_cluster_20))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L5
	} else {
		goto L210
	}
L210:
	;
	goto L155
L211:
	;
	v1170 = *(*float64)(unsafe.Add(mBase, uint32(l10)))
	*(*float64)(unsafe.Add(mBase, uint32(l10))) = base.F64_add(v1170, float64(1))
	goto L155
L212:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v838)+16))
	v1008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1007)+20)))
	if v1008&int32(_a_F_heapam_relation_copy_for_cluster_22) == int32(_a_F_heapam_relation_copy_for_cluster_23) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	if base.Ui32(v1016) < base.Ui32(int32(3)) {
		goto L219
	} else {
		goto L220
	}
L214:
	;
	v1013 = F_HeapTupleGetUpdateXid(m, v1007)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L5
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+4))
	v1016 = v1015
	goto L213
L217:
	;
	v1016 = v1013
	goto L213
L218:
	;
	if v1148 != 0 {
		goto L211
	} else {
		goto L258
	}
L219:
	;
	v1148 = int32(0)
	goto L218
L220:
	;
	goto L221
L221:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[13]))
	if v1028 == v1016 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1148 = int32(1)
	goto L218
L223:
	;
	goto L224
L224:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[14]))
	if v1032 <= int32(0) {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v1148 = v1138
	goto L218
L226:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[15]))
	if v1036 == int32(0) {
		v1138 = int32(0)
		goto L225
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[16]))
	v1108 = int32(0)
	v1111 = v1032 - int32(1)
	goto L248
L229:
	;
	v1041 = v1036
	goto L230
L230:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+20))
	if v1047 == int32(4) {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v1138 = int32(0)
	goto L225
L232:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+80))
	if v1101 != 0 {
		v1041 = v1101
		goto L230
	} else {
		goto L247
	}
L233:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1041)))
	if v1050 == int32(0) {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v1053 = int32(1)
	if v1016 == v1050 {
		v1138 = v1053
		goto L225
	} else {
		goto L235
	}
L235:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+52))
	v1057 = v1055 - int32(1)
	if v1057 < int32(0) {
		goto L232
	} else {
		goto L236
	}
L236:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+48))
	v1063 = int32(0)
	v1066 = v1057
	goto L237
L237:
	;
	v1071 = int32(2)
	v1072 = base.I32_div_s(v1066-v1063, v1071)
	v1073 = v1072 + v1063
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1060+v1073<<(uint(v1071)%32))))
	if v1077 == v1016 {
		v1138 = v1053
		goto L225
	} else {
		goto L239
	}
L238:
	;
	goto L232
L239:
	;
	v1086 = base.B2i32(v1077-v1016 < int32(0)) | base.B2i32(base.Ui32(v1077) < base.Ui32(int32(3)))
	if v1086 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1087 = v1073 + int32(1)
	goto L242
L241:
	;
	v1087 = v1063
	goto L242
L242:
	;
	if v1086 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1090 = v1066
	goto L245
L244:
	;
	v1090 = v1073 - int32(1)
	goto L245
L245:
	;
	if v1087 <= v1090 {
		v1063 = v1087
		v1066 = v1090
		goto L237
	} else {
		goto L246
	}
L246:
	;
	goto L238
L247:
	;
	goto L231
L248:
	;
	v1116 = int32(2)
	v1117 = base.I32_div_s(v1111-v1108, v1116)
	v1118 = v1117 + v1108
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1106+v1118<<(uint(v1116)%32))))
	v1123 = base.B2i32(v1122 == v1016)
	if v1122 == v1016 {
		v1138 = v1123
		goto L225
	} else {
		goto L250
	}
L249:
	;
	v1138 = v1123
	goto L225
L250:
	;
	v1126 = base.B2i32(base.Ui32(v1122) < base.Ui32(v1016))
	if base.Ui32(v1122) < base.Ui32(v1016) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1127 = v1118 + int32(1)
	goto L253
L252:
	;
	v1127 = v1108
	goto L253
L253:
	;
	if base.Ui32(v1122) < base.Ui32(v1016) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1130 = v1111
	goto L256
L255:
	;
	v1130 = v1118 - int32(1)
	goto L256
L256:
	;
	if v1127 <= v1130 {
		v1108 = v1127
		v1111 = v1130
		goto L248
	} else {
		goto L257
	}
L257:
	;
	goto L249
L258:
	;
	v1151 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L5
	} else {
		goto L259
	}
L259:
	;
	if v1151 == int32(0) {
		goto L211
	} else {
		goto L260
	}
L260:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1155 + int32(4)
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_24), v27+int32(-48))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_19), int32(827), int32(_a_F_heapam_relation_copy_for_cluster_20))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L5
	} else {
		goto L262
	}
L262:
	;
	goto L211
L263:
	;
	F_errmsg_internal(m, int32(_a_F_heapam_relation_copy_for_cluster_25), int32(0))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L5
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_19), int32(833), int32(_a_F_heapam_relation_copy_for_cluster_20))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L5
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	v1189 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1189, float64(1))
	v1193 = m.G0
	v1195 = v1193 - int32(16)
	m.G0 = v1195
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+12)) = int32(0)
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v838)+16))
	v1200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1199)+20)))
	v1201 = int32(768)
	if v1200&v1201 != v1201 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1199)))
	v1207 = v1205
	goto L269
L268:
	;
	v1207 = int32(2)
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+4)) = v1207
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v838)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1195)+8)) = v1209
	v1211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v838)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1195)+12)) = uint16(v1211)
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v209)+56))
	v1215 = v1195 + int32(4)
	v1216 = int32(0)
	v1218 = F_hash_search(m, v1213, v1215, v1216, v1216)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L5
	} else {
		goto L270
	}
L270:
	;
	if v1218 != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+20))
	F_pfree(m, v1220)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L5
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	m.G0 = v1195 + int32(16)
	if v1218 == int32(0) {
		v691 = v835
		goto L117
	} else {
		goto L276
	}
L274:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v209)+56))
	v1227 = F_hash_search(m, v1223, v1215, int32(2), v1195+int32(3))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L5
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	v1234 = *(*float64)(unsafe.Add(mBase, uint32(l9)))
	*(*float64)(unsafe.Add(mBase, uint32(l9))) = base.F64_add(v1234, float64(1))
	v1238 = *(*float64)(unsafe.Add(mBase, uint32(l10)))
	*(*float64)(unsafe.Add(mBase, uint32(l10))) = base.F64_add(v1238, float64(-1))
	v691 = v835
	goto L117
L277:
	;
	v1249 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	*(*float64)(unsafe.Add(mBase, uint32(l8))) = base.F64_add(v1249, float64(1))
	if v655 != 0 {
		goto L150
	} else {
		goto L278
	}
L278:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+56)) = int64(17179869187)
	v1255 = F_reform_tuple(m, v838, l0, l1, v47, v50)
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L5
	} else {
		goto L279
	}
L279:
	;
	F_rewrite_heap_tuple(m, v209, v838, v1255)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L5
	} else {
		goto L280
	}
L280:
	;
	v1386 = v1255
	goto L148
L281:
	;
	goto L150
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+8)) = v1277
	v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655)+36)))
	if v1280 == int32(1) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+4))
	v1284 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1283)+12)))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v1276)))
	v1288 = F_heap_getattr_1(m, v1277, v1284, v1285, v1269+int32(24))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L5
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655)+52)))
	if v1291&int32(2) == int32(0) {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1269)+16)) = v1288
	goto L285
L287:
	;
	v1307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655)+36)))
	if v1307 != int32(1) {
		v1318 = int32(0)
		goto L292
	} else {
		goto L293
	}
L288:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	v1303 = (v1296 + int32(31)) & int32(-8)
	goto L287
L289:
	;
	goto L290
L290:
	;
	v1301 = F_GetMemoryChunkSpace(m, v1277)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L5
	} else {
		goto L291
	}
L291:
	;
	v1303 = v1301
	goto L287
L292:
	;
	F_tuplesort_puttuple_common(m, v655, v1269+int32(8), v1318&int32(1), v1303)
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L5
	} else {
		goto L295
	}
L293:
	;
	v1310 = int32(0)
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v655)+44))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+24))
	if v1312 == v1310 {
		v1318 = v1310
		goto L292
	} else {
		goto L294
	}
L294:
	;
	v1315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1269)+24)))
	v1318 = v1315 ^ int32(1)
	goto L292
L295:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[0])) = v1272
	m.G0 = v1269 + int32(32)
	v1329 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	v1333 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v1333 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	v691 = v835
	goto L117
L297:
	;
	goto L296
L298:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v1337&int32(1) == int32(0) {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1342 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v1344 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v1345 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1344 + v1345
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1333)))
	*(*int32)(unsafe.Add(mBase, uint32(v1333))) = v1348 + v1345
	v1352 = int32(0)
	v1354 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1355 = base.AtomicRmwOr32(m, v1352, v1354, v1352)
	*(*int64)(unsafe.Add(mBase, uint32(v1333+int32(24))+232)) = base.I64_trunc_sat_f64_s(v1329)
	v1363 = base.AtomicRmwOr32(m, v1352, v1354, v1352)
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1333)))
	*(*int32)(unsafe.Add(mBase, uint32(v1333))) = v1364 + v1345
	v1370 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1370 - v1345
	goto L297
L300:
	;
	v1379 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	F_heap_insert(m, l1, v1376, v1379, int32(8), v214)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L5
	} else {
		goto L302
	}
L302:
	;
	v1386 = v1376
	goto L148
L303:
	;
	v1389 = *(*float64)(unsafe.Add(mBase, uint32(l8)))
	v1390 = base.I64_trunc_sat_f64_s(v1389)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+40)) = v1390
	*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v1390
	goto L306
L304:
	;
	v691 = v835
	goto L117
L305:
	;
	goto L304
L306:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v1407 == int32(0) {
		goto L305
	} else {
		goto L307
	}
L307:
	;
	v1411 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v1411&int32(1) == int32(0) {
		goto L305
	} else {
		goto L308
	}
L308:
	;
	v1416 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v1418 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v1419 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1418 + v1419
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1407)))
	*(*int32)(unsafe.Add(mBase, uint32(v1407))) = v1422 + v1419
	v1426 = int32(0)
	v1429 = base.AtomicRmwOr32(m, v1426, int32(_a_F_heapam_relation_copy_for_cluster_13), v1426)
	goto L310
L309:
	;
	v1556 = int32(0)
	v1559 = base.AtomicRmwOr32(m, v1556, int32(_a_F_heapam_relation_copy_for_cluster_13), v1556)
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1407)))
	v1561 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1407))) = v1560 + v1561
	v1564 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v1566 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1566 - v1561
	goto L305
L310:
	;
	goto L312
L312:
	;
	goto L313
L313:
	;
	v1521 = int32(0)
	v1524 = int32(0)
	goto L318
L318:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(-8)+v1524<<(uint(int32(2))%32))))
	v1534 = int32(3)
	v1540 = *(*int64)(unsafe.Add(mBase, uint32(v27+int32(-32)+v1524<<(uint(v1534)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1407+int32(232)+v1533<<(uint(v1534)%32)))) = v1540
	v1542 = int32(1)
	v1545 = v1521 + v1542
	if v1545 != int32(2) {
		v1521 = v1545
		v1524 = v1524 + v1542
		goto L318
	} else {
		goto L320
	}
L319:
	;
	goto L309
L320:
	;
	goto L319
L321:
	;
	if v649 == int32(0) {
		goto L115
	} else {
		goto L322
	}
L322:
	;
	goto L116
L323:
	;
	goto L115
L324:
	;
	F_ExecDropSingleTupleTableSlot(m, v670)
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L5
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	if v655 != 0 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	goto L326
L328:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v1594 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L329:
	;
	goto L330
L330:
	;
	if v209 != 0 {
		goto L364
	} else {
		goto L365
	}
L331:
	;
	F_tuplesort_performsort(m, v655)
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L5
	} else {
		goto L335
	}
L332:
	;
	goto L331
L333:
	;
	v1598 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v1598&int32(1) == int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1603 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v1605 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v1606 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1605 + v1606
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1594)))
	*(*int32)(unsafe.Add(mBase, uint32(v1594))) = v1609 + v1606
	v1613 = int32(0)
	v1615 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1616 = base.AtomicRmwOr32(m, v1613, v1615, v1613)
	*(*int64)(unsafe.Add(mBase, uint32(v1594+int32(8))+232)) = int64(3)
	v1624 = base.AtomicRmwOr32(m, v1613, v1615, v1613)
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1594)))
	*(*int32)(unsafe.Add(mBase, uint32(v1594))) = v1625 + v1606
	v1631 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1631 - v1606
	goto L332
L335:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v1641 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	v1707 = float64(0)
	goto L340
L337:
	;
	goto L336
L338:
	;
	v1645 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v1645&int32(1) == int32(0) {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1650 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v1652 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v1653 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1652 + v1653
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1641)))
	*(*int32)(unsafe.Add(mBase, uint32(v1641))) = v1656 + v1653
	v1660 = int32(0)
	v1662 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1663 = base.AtomicRmwOr32(m, v1660, v1662, v1660)
	*(*int64)(unsafe.Add(mBase, uint32(v1641+int32(8))+232)) = int64(4)
	v1671 = base.AtomicRmwOr32(m, v1660, v1662, v1660)
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1641)))
	*(*int32)(unsafe.Add(mBase, uint32(v1641))) = v1672 + v1653
	v1678 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1678 - v1653
	goto L337
L340:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[12]))
	if v1709 != 0 {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	F_tuplesort_end(m, v655)
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L5
	} else {
		goto L363
	}
L342:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L5
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	v1712 = F_tuplesort_getheaptuple(m, v655)
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L5
	} else {
		goto L346
	}
L345:
	;
	goto L344
L346:
	;
	if v1712 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1714 = F_reform_tuple(m, v1712, l0, l1, v47, v50)
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L5
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	goto L341
L350:
	;
	if l5 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	F_pfree(m, v1714)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L5
	} else {
		goto L358
	}
L352:
	;
	F_rewrite_heap_tuple(m, v209, v1712, v1714)
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L5
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1721 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L5
	} else {
		goto L356
	}
L355:
	;
	goto L351
L356:
	;
	F_heap_insert(m, l1, v1714, v1721, int32(8), v214)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L5
	} else {
		goto L357
	}
L357:
	;
	goto L351
L358:
	;
	v1730 = base.F64_add(v1707, float64(1))
	v1734 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[5]))
	if v1734 == int32(0) {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	v1707 = v1730
	goto L340
L360:
	;
	goto L359
L361:
	;
	v1738 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[9])))
	if v1738&int32(1) == int32(0) {
		goto L360
	} else {
		goto L362
	}
L362:
	;
	v1743 = int32(_a_F_heapam_relation_copy_for_cluster_12)
	v1745 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	v1746 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1745 + v1746
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1734)))
	*(*int32)(unsafe.Add(mBase, uint32(v1734))) = v1749 + v1746
	v1753 = int32(0)
	v1755 = int32(_a_F_heapam_relation_copy_for_cluster_13)
	v1756 = base.AtomicRmwOr32(m, v1753, v1755, v1753)
	*(*int64)(unsafe.Add(mBase, uint32(v1734+int32(32))+232)) = base.I64_trunc_sat_f64_s(v1730)
	v1764 = base.AtomicRmwOr32(m, v1753, v1755, v1753)
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1734)))
	*(*int32)(unsafe.Add(mBase, uint32(v1734))) = v1765 + v1746
	v1771 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[10])) = v1771 - v1746
	goto L360
L363:
	;
	goto L330
L364:
	;
	v1803 = m.G0
	v1805 = v1803 - int32(48)
	m.G0 = v1805
	v1808 = v1805 + int32(8)
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v209)+56))
	F_hash_seq_init(m, v1808, v1809)
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L5
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	if v214 != 0 {
		goto L409
	} else {
		goto L410
	}
L367:
	;
	v1812 = F_hash_seq_search(m, v1808)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L5
	} else {
		goto L368
	}
L368:
	;
	if v1812 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1814 = v1812
	goto L372
L370:
	;
	goto L371
L371:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	if v1879 != 0 {
		goto L377
	} else {
		goto L378
	}
L372:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+20))
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+16))
	v1842 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1841)+16)) = uint16(v1842)
	*(*int32)(unsafe.Add(mBase, uint32(v1841)+12)) = int32(-1)
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+20))
	F_raw_heap_insert(m, v209, v1846)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L5
	} else {
		goto L374
	}
L373:
	;
	goto L371
L374:
	;
	v1851 = F_hash_seq_search(m, v1805+int32(8))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L5
	} else {
		goto L375
	}
L375:
	;
	if v1851 != 0 {
		v1814 = v1851
		goto L372
	} else {
		goto L376
	}
L376:
	;
	goto L373
L377:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v209)+8))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v209)+16))
	F_smgr_bulk_write(m, v1880, v1881, v1879, int32(1))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L5
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v209)+8))
	F_smgr_bulk_finish(m, v1887)
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L5
	} else {
		goto L381
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+12)) = int32(0)
	goto L379
L381:
	;
	v1890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+20)))
	if v1890 != int32(1) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v209)+40))
	F_MemoryContextDelete(m, v1993)
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L5
	} else {
		goto L408
	}
L383:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v209)+68))
	if v1893 != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	F_logical_heap_rewrite_flush_mappings(m, v209)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L5
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	v1897 = v1805 + int32(28)
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v209)+64))
	F_hash_seq_init(m, v1897, v1898)
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L5
	} else {
		goto L388
	}
L387:
	;
	goto L386
L388:
	;
	v1901 = F_hash_seq_search(m, v1897)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L5
	} else {
		goto L389
	}
L389:
	;
	if v1901 == int32(0) {
		goto L382
	} else {
		goto L390
	}
L390:
	;
	v1905 = v1901
	goto L391
L391:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1905)+4))
	v1933 = F_FileSync(m, v1931, int32(167772199))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L5
	} else {
		goto L394
	}
L392:
	;
	goto L382
L393:
	;
	v1960 = *(*int32)(unsafe.Add(mBase, uint32(v1905)+4))
	F_FileClose(m, v1960)
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L5
	} else {
		goto L405
	}
L394:
	;
	if v1933 == int32(0) {
		goto L393
	} else {
		goto L395
	}
L395:
	;
	v1940 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_relation_copy_for_cluster[17])))
	if v1940 != 0 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	v1943 = F_errstart(m, v1941, int32(0))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L5
	} else {
		goto L400
	}
L397:
	;
	v1941 = int32(21)
	goto L399
L398:
	;
	v1941 = int32(24)
	goto L399
L399:
	;
	goto L396
L400:
	;
	if v1943 == int32(0) {
		goto L393
	} else {
		goto L401
	}
L401:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L5
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1805))) = v1905 + int32(28)
	F_errmsg(m, int32(_a_F_heapam_relation_copy_for_cluster_26), v1805)
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L5
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(_a_F_heapam_relation_copy_for_cluster_27), int32(928), int32(_a_F_heapam_relation_copy_for_cluster_28))
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L5
	} else {
		goto L404
	}
L404:
	;
	goto L393
L405:
	;
	v1965 = F_hash_seq_search(m, v1805+int32(28))
	mBase = m.M
	v1966 = m.ExcPending
	if v1966 != 0 {
		goto L5
	} else {
		goto L406
	}
L406:
	;
	if v1965 != 0 {
		v1905 = v1965
		goto L391
	} else {
		goto L407
	}
L407:
	;
	goto L392
L408:
	;
	m.G0 = v1805 + int32(48)
	goto L366
L409:
	;
	F_FreeBulkInsertState(m, v214)
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L5
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	F_pfree(m, v47)
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L5
	} else {
		goto L413
	}
L412:
	;
	goto L411
L413:
	;
	F_pfree(m, v50)
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L5
	} else {
		goto L414
	}
L414:
	;
	m.G0 = v29 - int32(-64)
	return
}
