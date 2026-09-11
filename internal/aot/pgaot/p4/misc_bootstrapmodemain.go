package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BootstrapModeMain(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v387 int32
	_ = v387
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v429 int32
	_ = v429
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v471 int32
	_ = v471
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v501 int32
	_ = v501
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v552 int64
	_ = v552
	var v553 int64
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int64
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
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
	var v605 int32
	_ = v605
	var v606 int64
	_ = v606
	var v617 int64
	_ = v617
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v626 int64
	_ = v626
	var v638 int32
	_ = v638
	var v640 int64
	_ = v640
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v805 int32
	_ = v805
	var v807 int64
	_ = v807
	var v809 int64
	_ = v809
	var v811 int64
	_ = v811
	var v813 int64
	_ = v813
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v851 int64
	_ = v851
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1161 int64
	_ = v1161
	var v1176 int32
	_ = v1176
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int64
	_ = v1200
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1258 int32
	_ = v1258
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
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
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1512 int32
	_ = v1512
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1544 int32
	_ = v1544
	var v1560 int32
	_ = v1560
	var v1566 int32
	_ = v1566
	var v1570 int32
	_ = v1570
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1633 int32
	_ = v1633
	var v1638 int32
	_ = v1638
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1730 int32
	_ = v1730
	var v1734 int32
	_ = v1734
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1772 int32
	_ = v1772
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1961 int32
	_ = v1961
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1980 int32
	_ = v1980
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1995 int32
	_ = v1995
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2018 int32
	_ = v2018
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2033 int32
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2052 int32
	_ = v2052
	var v2056 int32
	_ = v2056
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2090 int32
	_ = v2090
	var v2094 int32
	_ = v2094
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2132 int32
	_ = v2132
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2147 int32
	_ = v2147
	var v2151 int32
	_ = v2151
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2166 int32
	_ = v2166
	var v2170 int32
	_ = v2170
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
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2202 int32
	_ = v2202
	var v2206 int32
	_ = v2206
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2219 int32
	_ = v2219
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2235 int32
	_ = v2235
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2255 int32
	_ = v2255
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2272 int32
	_ = v2272
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2318 int32
	_ = v2318
	var v2322 int32
	_ = v2322
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2358 int32
	_ = v2358
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2367 int32
	_ = v2367
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2413 int32
	_ = v2413
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2432 int32
	_ = v2432
	var v2445 int32
	_ = v2445
	var v2461 int32
	_ = v2461
	var v2465 int32
	_ = v2465
	var v2467 int32
	_ = v2467
	var v2473 int32
	_ = v2473
	var v2495 int32
	_ = v2495
	var v2498 int32
	_ = v2498
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2510 int32
	_ = v2510
	var v2516 int32
	_ = v2516
	var v2537 int32
	_ = v2537
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2554 int32
	_ = v2554
	var v2558 int32
	_ = v2558
	var v2585 int32
	_ = v2585
	var v2593 int32
	_ = v2593
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2620 int32
	_ = v2620
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2655 int32
	_ = v2655
	var v2657 int32
	_ = v2657
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2662 int32
	_ = v2662
	var v2664 int32
	_ = v2664
	var v2668 int32
	_ = v2668
	var v2670 int32
	_ = v2670
	var v2694 int32
	_ = v2694
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2717 int32
	_ = v2717
	var v2719 int32
	_ = v2719
	var v2724 int32
	_ = v2724
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2756 int32
	_ = v2756
	var v2759 int32
	_ = v2759
	var v2762 int32
	_ = v2762
	var v2766 int32
	_ = v2766
	var v2782 int32
	_ = v2782
	var v2785 int32
	_ = v2785
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2801 int32
	_ = v2801
	var v2806 int32
	_ = v2806
	var v2822 int32
	_ = v2822
	var v2827 int32
	_ = v2827
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2846 int32
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2860 int32
	_ = v2860
	var v2863 int32
	_ = v2863
	var v2883 int32
	_ = v2883
	var v2886 int32
	_ = v2886
	var v2888 int32
	_ = v2888
	var v2893 int32
	_ = v2893
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2930 int32
	_ = v2930
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2937 int32
	_ = v2937
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2954 int32
	_ = v2954
	var v2955 int32
	_ = v2955
	var v2958 int32
	_ = v2958
	var v2962 int32
	_ = v2962
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2983 int32
	_ = v2983
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3009 int32
	_ = v3009
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3020 int32
	_ = v3020
	var v3024 int32
	_ = v3024
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3030 int32
	_ = v3030
	var v3034 int32
	_ = v3034
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3043 int32
	_ = v3043
	var v3044 int32
	_ = v3044
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3052 int32
	_ = v3052
	var v3055 int32
	_ = v3055
	var v3058 int32
	_ = v3058
	var v3063 int32
	_ = v3063
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3093 int32
	_ = v3093
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3118 int32
	_ = v3118
	var v3121 int32
	_ = v3121
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3127 int32
	_ = v3127
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3134 int32
	_ = v3134
	var v3139 int32
	_ = v3139
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3145 int32
	_ = v3145
	var v3146 int32
	_ = v3146
	var v3150 int32
	_ = v3150
	var v3154 int32
	_ = v3154
	var v3156 int32
	_ = v3156
	var v3159 int32
	_ = v3159
	var v3166 int32
	_ = v3166
	var v3168 int32
	_ = v3168
	var v3171 int32
	_ = v3171
	var v3174 int32
	_ = v3174
	var v3175 int32
	_ = v3175
	var v3177 int32
	_ = v3177
	var v3179 int32
	_ = v3179
	var v3180 int32
	_ = v3180
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3186 int32
	_ = v3186
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3200 int32
	_ = v3200
	var v3208 int32
	_ = v3208
	var v3209 int32
	_ = v3209
	var v3212 int32
	_ = v3212
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3223 int32
	_ = v3223
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3231 int32
	_ = v3231
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3261 int32
	_ = v3261
	var v3262 int32
	_ = v3262
	var v3265 int32
	_ = v3265
	var v3267 int32
	_ = v3267
	var v3268 int32
	_ = v3268
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3279 int32
	_ = v3279
	var v3280 int32
	_ = v3280
	var v3282 int32
	_ = v3282
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3287 int32
	_ = v3287
	var v3290 int32
	_ = v3290
	var v3291 int32
	_ = v3291
	var v3293 int32
	_ = v3293
	var v3295 int32
	_ = v3295
	var v3296 int32
	_ = v3296
	var v3300 int32
	_ = v3300
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3321 int32
	_ = v3321
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3326 int32
	_ = v3326
	var v3332 int32
	_ = v3332
	var v3336 int32
	_ = v3336
	var v3354 int32
	_ = v3354
	var v3357 int32
	_ = v3357
	var v3358 int32
	_ = v3358
	var v3363 int32
	_ = v3363
	var v3367 int32
	_ = v3367
	var v3368 int32
	_ = v3368
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3378 int32
	_ = v3378
	var v3382 int32
	_ = v3382
	var v3384 int32
	_ = v3384
	var v3385 int32
	_ = v3385
	var v3409 int32
	_ = v3409
	var v3414 int32
	_ = v3414
	var v3415 int32
	_ = v3415
	var v3417 int32
	_ = v3417
	var v3418 int32
	_ = v3418
	var v3422 int32
	_ = v3422
	var v3423 int32
	_ = v3423
	var v3428 int32
	_ = v3428
	var v3441 int32
	_ = v3441
	var v3457 int32
	_ = v3457
	var v3461 int32
	_ = v3461
	var v3463 int32
	_ = v3463
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3468 int32
	_ = v3468
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3476 int32
	_ = v3476
	var v3477 int32
	_ = v3477
	var v3493 int32
	_ = v3493
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3500 int32
	_ = v3500
	var v3501 int32
	_ = v3501
	var v3507 int32
	_ = v3507
	var v3511 int32
	_ = v3511
	var v3529 int32
	_ = v3529
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3538 int32
	_ = v3538
	var v3542 int32
	_ = v3542
	var v3543 int32
	_ = v3543
	var v3547 int32
	_ = v3547
	var v3548 int32
	_ = v3548
	var v3553 int32
	_ = v3553
	var v3557 int32
	_ = v3557
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3584 int32
	_ = v3584
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3592 int32
	_ = v3592
	var v3593 int32
	_ = v3593
	var v3597 int32
	_ = v3597
	var v3598 int32
	_ = v3598
	var v3603 int32
	_ = v3603
	var v3616 int32
	_ = v3616
	var v3632 int32
	_ = v3632
	var v3636 int32
	_ = v3636
	var v3638 int32
	_ = v3638
	var v3642 int32
	_ = v3642
	var v3647 int32
	_ = v3647
	var v3648 int32
	_ = v3648
	var v3676 int32
	_ = v3676
	var v3680 int32
	_ = v3680
	var v3686 int32
	_ = v3686
	var v3690 int32
	_ = v3690
	var v3692 int32
	_ = v3692
	var v3693 int32
	_ = v3693
	var v3695 int32
	_ = v3695
	var v3696 int32
	_ = v3696
	var v3697 int32
	_ = v3697
	var v3700 int32
	_ = v3700
	var v3702 int32
	_ = v3702
	var v3703 int32
	_ = v3703
	var v3710 int32
	_ = v3710
	var v3717 int32
	_ = v3717
	var v3718 int32
	_ = v3718
	var v3719 int32
	_ = v3719
	var v3720 int32
	_ = v3720
	var v3725 int32
	_ = v3725
	var v3729 int32
	_ = v3729
	var v3735 int32
	_ = v3735
	var v3736 int32
	_ = v3736
	var v3740 int32
	_ = v3740
	var v3751 int32
	_ = v3751
	var v3753 int32
	_ = v3753
	var v3770 int32
	_ = v3770
	var v3777 int32
	_ = v3777
	var v3781 int32
	_ = v3781
	var v3783 int32
	_ = v3783
	var v3784 int32
	_ = v3784
	var v3786 int32
	_ = v3786
	var v3787 int32
	_ = v3787
	var v3788 int32
	_ = v3788
	var v3791 int32
	_ = v3791
	var v3793 int32
	_ = v3793
	var v3794 int32
	_ = v3794
	var v3801 int32
	_ = v3801
	var v3807 int32
	_ = v3807
	var v3812 int32
	_ = v3812
	var v3816 int32
	_ = v3816
	var v3821 int32
	_ = v3821
	var v3826 int32
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3829 int32
	_ = v3829
	var v3832 int32
	_ = v3832
	var v3833 int32
	_ = v3833
	var v3834 int32
	_ = v3834
	var v3836 int32
	_ = v3836
	var v3845 int32
	_ = v3845
	var v3850 int32
	_ = v3850
	var v3854 int32
	_ = v3854
	var v3859 int32
	_ = v3859
	var v3861 int32
	_ = v3861
	var v3865 int32
	_ = v3865
	var v3871 int32
	_ = v3871
	var v3874 int32
	_ = v3874
	var v3880 int32
	_ = v3880
	var v3884 int32
	_ = v3884
	var v3886 int32
	_ = v3886
	var v3894 int32
	_ = v3894
	var v3897 int32
	_ = v3897
	var v3900 int32
	_ = v3900
	var v3904 int32
	_ = v3904
	var v3906 int32
	_ = v3906
	var v3909 int32
	_ = v3909
	var v3912 int32
	_ = v3912
	var v3913 int32
	_ = v3913
	var v3921 int32
	_ = v3921
	var v3926 int32
	_ = v3926
	var v3930 int32
	_ = v3930
	var v3931 int32
	_ = v3931
	var v3933 int32
	_ = v3933
	var v3934 int32
	_ = v3934
	var v3937 int32
	_ = v3937
	var v3938 int32
	_ = v3938
	var v3950 int32
	_ = v3950
	var v3970 int32
	_ = v3970
	var v3973 int32
	_ = v3973
	var v3977 int32
	_ = v3977
	var v3979 int32
	_ = v3979
	var v3980 int32
	_ = v3980
	var v3982 int32
	_ = v3982
	var v3984 int32
	_ = v3984
	var v3985 int32
	_ = v3985
	var v3986 int32
	_ = v3986
	var v3990 int32
	_ = v3990
	var v3994 int32
	_ = v3994
	var v3999 int32
	_ = v3999
	var v4006 int32
	_ = v4006
	var v4009 int32
	_ = v4009
	var v4016 int32
	_ = v4016
	var v4017 int32
	_ = v4017
	var v4018 int32
	_ = v4018
	var v4022 int32
	_ = v4022
	var v4024 int32
	_ = v4024
	var v4025 int32
	_ = v4025
	var v4027 int32
	_ = v4027
	var v4029 int32
	_ = v4029
	var v4036 int32
	_ = v4036
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4045 int32
	_ = v4045
	var v4051 int32
	_ = v4051
	var v4058 int32
	_ = v4058
	var v4062 int32
	_ = v4062
	var v4065 int32
	_ = v4065
	var v4071 int32
	_ = v4071
	var v4078 int32
	_ = v4078
	var v4082 int32
	_ = v4082
	var v4085 int32
	_ = v4085
	var v4088 int32
	_ = v4088
	var v4089 int32
	_ = v4089
	var v4090 int32
	_ = v4090
	var v4095 int32
	_ = v4095
	var v4096 int32
	_ = v4096
	var v4097 int32
	_ = v4097
	var v4099 int32
	_ = v4099
	var v4101 int32
	_ = v4101
	var v4102 int32
	_ = v4102
	var v4104 int32
	_ = v4104
	var v4106 int32
	_ = v4106
	var v4110 int32
	_ = v4110
	var v4111 int32
	_ = v4111
	var v4112 int32
	_ = v4112
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4119 int32
	_ = v4119
	var v4121 int32
	_ = v4121
	var v4123 int32
	_ = v4123
	var v4128 int32
	_ = v4128
	var v4140 int32
	_ = v4140
	var v4143 int32
	_ = v4143
	var v4144 int32
	_ = v4144
	var v4145 int32
	_ = v4145
	var v4146 int32
	_ = v4146
	var v4147 int32
	_ = v4147
	var v4157 int32
	_ = v4157
	var v4162 int32
	_ = v4162
	var v4166 int32
	_ = v4166
	var v4168 int32
	_ = v4168
	var v4200 int32
	_ = v4200
	var v4203 int32
	_ = v4203
	var v4205 int32
	_ = v4205
	var v4207 int32
	_ = v4207
	var v4209 int32
	_ = v4209
	var v4210 int32
	_ = v4210
	var v4211 int32
	_ = v4211
	var v4217 int32
	_ = v4217
	var v4218 int32
	_ = v4218
	var v4219 int32
	_ = v4219
	var v4221 int32
	_ = v4221
	var v4226 int32
	_ = v4226
	var v4231 int32
	_ = v4231
	var v4232 int32
	_ = v4232
	var v4234 int32
	_ = v4234
	var v4237 int32
	_ = v4237
	var v4239 int32
	_ = v4239
	var v4242 int32
	_ = v4242
	var v4245 int32
	_ = v4245
	var v4247 int32
	_ = v4247
	var v4249 int32
	_ = v4249
	var v4251 int32
	_ = v4251
	var v4252 int32
	_ = v4252
	var v4253 int32
	_ = v4253
	var v4259 int32
	_ = v4259
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4263 int32
	_ = v4263
	var v4268 int32
	_ = v4268
	var v4273 int32
	_ = v4273
	var v4274 int32
	_ = v4274
	var v4276 int32
	_ = v4276
	var v4280 int32
	_ = v4280
	var v4284 int32
	_ = v4284
	var v4285 int32
	_ = v4285
	var v4290 int32
	_ = v4290
	var v4293 int32
	_ = v4293
	var v4296 int64
	_ = v4296
	var v4300 int32
	_ = v4300
	var v4304 int32
	_ = v4304
	var v4308 int32
	_ = v4308
	var v4313 int32
	_ = v4313
	var v4316 int32
	_ = v4316
	var v4319 int32
	_ = v4319
	var v4321 int32
	_ = v4321
	var v4323 int32
	_ = v4323
	var v4325 int32
	_ = v4325
	var v4326 int32
	_ = v4326
	var v4327 int32
	_ = v4327
	var v4333 int32
	_ = v4333
	var v4334 int32
	_ = v4334
	var v4335 int32
	_ = v4335
	var v4337 int32
	_ = v4337
	var v4342 int32
	_ = v4342
	var v4347 int32
	_ = v4347
	var v4348 int32
	_ = v4348
	var v4350 int32
	_ = v4350
	var v4354 int32
	_ = v4354
	var v4356 int32
	_ = v4356
	var v4357 int32
	_ = v4357
	var v4360 int32
	_ = v4360
	var v4363 int32
	_ = v4363
	var v4365 int32
	_ = v4365
	var v4368 int32
	_ = v4368
	var v4369 int32
	_ = v4369
	var v4373 int32
	_ = v4373
	var v4378 int32
	_ = v4378
	var v4381 int32
	_ = v4381
	var v4385 int32
	_ = v4385
	var v4389 int32
	_ = v4389
	var v4392 int32
	_ = v4392
	var v4393 int32
	_ = v4393
	var v4396 int32
	_ = v4396
	var v4399 int32
	_ = v4399
	var v4406 int32
	_ = v4406
	var v4407 int32
	_ = v4407
	var v4411 int32
	_ = v4411
	var v4412 int32
	_ = v4412
	var v4418 int32
	_ = v4418
	var v4423 int32
	_ = v4423
	var v4426 int32
	_ = v4426
	var v4430 int32
	_ = v4430
	var v4433 int32
	_ = v4433
	var v4436 int32
	_ = v4436
	var v4437 int32
	_ = v4437
	var v4444 int32
	_ = v4444
	var v4452 int32
	_ = v4452
	var v4453 int32
	_ = v4453
	var v4456 int32
	_ = v4456
	var v4457 int32
	_ = v4457
	var v4465 int32
	_ = v4465
	var v4470 int32
	_ = v4470
	var v4474 int32
	_ = v4474
	var v4477 int32
	_ = v4477
	var v4479 int32
	_ = v4479
	var v4481 int32
	_ = v4481
	var v4483 int32
	_ = v4483
	var v4484 int32
	_ = v4484
	var v4485 int32
	_ = v4485
	var v4491 int32
	_ = v4491
	var v4492 int32
	_ = v4492
	var v4493 int32
	_ = v4493
	var v4495 int32
	_ = v4495
	var v4500 int32
	_ = v4500
	var v4505 int32
	_ = v4505
	var v4506 int32
	_ = v4506
	var v4508 int32
	_ = v4508
	var v4513 int32
	_ = v4513
	var v4514 int32
	_ = v4514
	var v4518 int32
	_ = v4518
	var v4523 int32
	_ = v4523
	var v4528 int32
	_ = v4528
	var v4530 int32
	_ = v4530
	var v4533 int32
	_ = v4533
	var v4536 int32
	_ = v4536
	var v4538 int32
	_ = v4538
	var v4542 int32
	_ = v4542
	var v4543 int32
	_ = v4543
	var v4545 int32
	_ = v4545
	var v4549 int32
	_ = v4549
	var v4554 int32
	_ = v4554
	var v4556 int32
	_ = v4556
	var v4558 int32
	_ = v4558
	var v4559 int32
	_ = v4559
	var v4562 int32
	_ = v4562
	var v4563 int32
	_ = v4563
	var v4565 int32
	_ = v4565
	var v4567 int32
	_ = v4567
	var v4569 int32
	_ = v4569
	var v4571 int32
	_ = v4571
	var v4574 int32
	_ = v4574
	var v4575 int32
	_ = v4575
	var v4579 int32
	_ = v4579
	var v4584 int32
	_ = v4584
	var v4586 int32
	_ = v4586
	var v4592 int32
	_ = v4592
	var v4598 int32
	_ = v4598
	var v4601 int32
	_ = v4601
	var v4603 int32
	_ = v4603
	var v4605 int32
	_ = v4605
	var v4607 int32
	_ = v4607
	var v4608 int32
	_ = v4608
	var v4609 int32
	_ = v4609
	var v4615 int32
	_ = v4615
	var v4616 int32
	_ = v4616
	var v4617 int32
	_ = v4617
	var v4619 int32
	_ = v4619
	var v4620 int32
	_ = v4620
	var v4625 int32
	_ = v4625
	var v4626 int32
	_ = v4626
	var v4629 int32
	_ = v4629
	var v4635 int32
	_ = v4635
	var v4640 int32
	_ = v4640
	var v4642 int32
	_ = v4642
	var v4647 int32
	_ = v4647
	var v4652 int32
	_ = v4652
	var v4653 int32
	_ = v4653
	var v4655 int32
	_ = v4655
	var v4660 int32
	_ = v4660
	var v4665 int32
	_ = v4665
	var v4667 int32
	_ = v4667
	var v4668 int32
	_ = v4668
	var v4672 int32
	_ = v4672
	var v4673 int32
	_ = v4673
	var v4678 int32
	_ = v4678
	var v4682 int64
	_ = v4682
	var v4702 int32
	_ = v4702
	var v4703 int32
	_ = v4703
	var v4706 int32
	_ = v4706
	var v4707 int32
	_ = v4707
	var v4716 int32
	_ = v4716
	var v4719 int32
	_ = v4719
	var v4722 int32
	_ = v4722
	var v4724 int32
	_ = v4724
	var v4726 int32
	_ = v4726
	var v4728 int32
	_ = v4728
	var v4729 int32
	_ = v4729
	var v4730 int32
	_ = v4730
	var v4736 int32
	_ = v4736
	var v4737 int32
	_ = v4737
	var v4738 int32
	_ = v4738
	var v4740 int32
	_ = v4740
	var v4741 int32
	_ = v4741
	var v4746 int32
	_ = v4746
	var v4747 int32
	_ = v4747
	var v4750 int32
	_ = v4750
	var v4756 int32
	_ = v4756
	var v4761 int32
	_ = v4761
	var v4763 int32
	_ = v4763
	var v4768 int32
	_ = v4768
	var v4773 int32
	_ = v4773
	var v4774 int32
	_ = v4774
	var v4776 int32
	_ = v4776
	var v4781 int32
	_ = v4781
	var v4786 int32
	_ = v4786
	var v4788 int32
	_ = v4788
	var v4789 int32
	_ = v4789
	var v4793 int32
	_ = v4793
	var v4794 int32
	_ = v4794
	var v4799 int32
	_ = v4799
	var v4800 int64
	_ = v4800
	var v4815 int32
	_ = v4815
	var v4825 int32
	_ = v4825
	var v4826 int32
	_ = v4826
	var v4829 int32
	_ = v4829
	var v4830 int32
	_ = v4830
	var v4839 int32
	_ = v4839
	var v4842 int32
	_ = v4842
	var v4845 int32
	_ = v4845
	var v4847 int32
	_ = v4847
	var v4849 int32
	_ = v4849
	var v4851 int32
	_ = v4851
	var v4852 int32
	_ = v4852
	var v4853 int32
	_ = v4853
	var v4859 int32
	_ = v4859
	var v4860 int32
	_ = v4860
	var v4861 int32
	_ = v4861
	var v4864 int32
	_ = v4864
	var v4865 int32
	_ = v4865
	var v4866 int32
	_ = v4866
	var v4872 int32
	_ = v4872
	var v4877 int32
	_ = v4877
	var v4879 int32
	_ = v4879
	var v4884 int32
	_ = v4884
	var v4889 int32
	_ = v4889
	var v4890 int32
	_ = v4890
	var v4892 int32
	_ = v4892
	var v4895 int32
	_ = v4895
	var v4898 int32
	_ = v4898
	var v4901 int32
	_ = v4901
	var v4902 int32
	_ = v4902
	var v4904 int32
	_ = v4904
	var v4908 int32
	_ = v4908
	var v4909 int32
	_ = v4909
	var v4911 int32
	_ = v4911
	var v4912 int32
	_ = v4912
	var v4913 int32
	_ = v4913
	var v4914 int32
	_ = v4914
	var v4920 int32
	_ = v4920
	var v4924 int32
	_ = v4924
	var v4929 int32
	_ = v4929
	var v4930 int32
	_ = v4930
	var v4934 int32
	_ = v4934
	var v4935 int32
	_ = v4935
	var v4941 int32
	_ = v4941
	var v4947 int32
	_ = v4947
	var v4952 int32
	_ = v4952
	var v4955 int32
	_ = v4955
	var v4961 int32
	_ = v4961
	var v4964 int32
	_ = v4964
	var v4966 int32
	_ = v4966
	var v4968 int32
	_ = v4968
	var v4970 int32
	_ = v4970
	var v4971 int32
	_ = v4971
	var v4972 int32
	_ = v4972
	var v4978 int32
	_ = v4978
	var v4979 int32
	_ = v4979
	var v4980 int32
	_ = v4980
	var v4982 int32
	_ = v4982
	var v4987 int32
	_ = v4987
	var v4992 int32
	_ = v4992
	var v4993 int32
	_ = v4993
	var v4995 int32
	_ = v4995
	var v4999 int32
	_ = v4999
	var v5000 int32
	_ = v5000
	var v5025 int32
	_ = v5025
	var v5027 int32
	_ = v5027
	var v5028 int32
	_ = v5028
	var v5030 int32
	_ = v5030
	var v5031 int32
	_ = v5031
	var v5033 int32
	_ = v5033
	var v5034 int32
	_ = v5034
	var v5036 int32
	_ = v5036
	var v5037 int32
	_ = v5037
	var v5038 int32
	_ = v5038
	var v5041 int32
	_ = v5041
	var v5044 int32
	_ = v5044
	var v5047 int32
	_ = v5047
	var v5048 int32
	_ = v5048
	var v5050 int32
	_ = v5050
	var v5051 int32
	_ = v5051
	var v5080 int32
	_ = v5080
	var v5083 int32
	_ = v5083
	var v5085 int32
	_ = v5085
	var v5087 int32
	_ = v5087
	var v5089 int32
	_ = v5089
	var v5090 int32
	_ = v5090
	var v5091 int32
	_ = v5091
	var v5097 int32
	_ = v5097
	var v5098 int32
	_ = v5098
	var v5099 int32
	_ = v5099
	var v5102 int32
	_ = v5102
	var v5103 int32
	_ = v5103
	var v5104 int32
	_ = v5104
	var v5105 int32
	_ = v5105
	var v5106 int32
	_ = v5106
	var v5112 int32
	_ = v5112
	var v5113 int32
	_ = v5113
	var v5115 int32
	_ = v5115
	var v5116 int32
	_ = v5116
	var v5121 int32
	_ = v5121
	var v5127 int32
	_ = v5127
	var v5128 int32
	_ = v5128
	var v5129 int32
	_ = v5129
	var v5135 int32
	_ = v5135
	var v5136 int32
	_ = v5136
	var v5142 int32
	_ = v5142
	var v5143 int32
	_ = v5143
	var v5145 int32
	_ = v5145
	var v5147 int32
	_ = v5147
	var v5153 int32
	_ = v5153
	var v5156 int32
	_ = v5156
	var v5157 int32
	_ = v5157
	var v5158 int32
	_ = v5158
	var v5160 int32
	_ = v5160
	var v5163 int32
	_ = v5163
	var v5166 int32
	_ = v5166
	var v5167 int32
	_ = v5167
	var v5171 int32
	_ = v5171
	var v5176 int32
	_ = v5176
	var v5179 int32
	_ = v5179
	var v5181 int32
	_ = v5181
	var v5184 int32
	_ = v5184
	var v5188 int32
	_ = v5188
	var v5190 int32
	_ = v5190
	var v5191 int32
	_ = v5191
	var v5193 int32
	_ = v5193
	var v5199 int32
	_ = v5199
	var v5205 int32
	_ = v5205
	var v5207 int32
	_ = v5207
	var v5214 int32
	_ = v5214
	var v5218 int32
	_ = v5218
	var v5221 int32
	_ = v5221
	var v5225 int32
	_ = v5225
	var v5226 int32
	_ = v5226
	var v5228 int32
	_ = v5228
	var v5231 int32
	_ = v5231
	var v5232 int32
	_ = v5232
	var v5233 int32
	_ = v5233
	var v5242 int32
	_ = v5242
	var v5247 int32
	_ = v5247
	var v5249 int32
	_ = v5249
	var v5251 int32
	_ = v5251
	var v5254 int32
	_ = v5254
	var v5262 int32
	_ = v5262
	var v5285 int32
	_ = v5285
	var v5292 int32
	_ = v5292
	var v5293 int32
	_ = v5293
	var v5294 int32
	_ = v5294
	var v5295 int32
	_ = v5295
	var v5296 int32
	_ = v5296
	var v5298 int32
	_ = v5298
	var v5303 int32
	_ = v5303
	var v5306 int32
	_ = v5306
	var v5307 int32
	_ = v5307
	var v5308 int32
	_ = v5308
	var v5313 int32
	_ = v5313
	var v5315 int32
	_ = v5315
	var v5318 int32
	_ = v5318
	var v5322 int32
	_ = v5322
	var v5323 int32
	_ = v5323
	var v5334 int32
	_ = v5334
	var v5339 int32
	_ = v5339
	var v5340 int32
	_ = v5340
	var v5346 int32
	_ = v5346
	var v5351 int32
	_ = v5351
	var v5353 int32
	_ = v5353
	var v5354 int32
	_ = v5354
	var v5356 int32
	_ = v5356
	var v5364 int32
	_ = v5364
	var v5377 int32
	_ = v5377
	var v5384 int32
	_ = v5384
	var v5387 int32
	_ = v5387
	var v5393 int32
	_ = v5393
	var v5416 int32
	_ = v5416
	var v5418 int32
	_ = v5418
	var v5425 int32
	_ = v5425
	var v5426 int32
	_ = v5426
	var v5427 int32
	_ = v5427
	var v5428 int32
	_ = v5428
	var v5429 int32
	_ = v5429
	var v5431 int32
	_ = v5431
	var v5436 int32
	_ = v5436
	var v5439 int32
	_ = v5439
	var v5440 int32
	_ = v5440
	var v5441 int32
	_ = v5441
	var v5446 int32
	_ = v5446
	var v5448 int32
	_ = v5448
	var v5451 int32
	_ = v5451
	var v5455 int32
	_ = v5455
	var v5456 int32
	_ = v5456
	var v5467 int32
	_ = v5467
	var v5495 int32
	_ = v5495
	var v5500 int32
	_ = v5500
	var v5502 int32
	_ = v5502
	var v5505 int32
	_ = v5505
	var v5508 int32
	_ = v5508
	var v5515 int32
	_ = v5515
	var v5538 int32
	_ = v5538
	var v5540 int32
	_ = v5540
	var v5547 int32
	_ = v5547
	var v5548 int32
	_ = v5548
	var v5549 int32
	_ = v5549
	var v5550 int32
	_ = v5550
	var v5551 int32
	_ = v5551
	var v5553 int32
	_ = v5553
	var v5558 int32
	_ = v5558
	var v5561 int32
	_ = v5561
	var v5562 int32
	_ = v5562
	var v5563 int32
	_ = v5563
	var v5568 int32
	_ = v5568
	var v5570 int32
	_ = v5570
	var v5573 int32
	_ = v5573
	var v5577 int32
	_ = v5577
	var v5578 int32
	_ = v5578
	var v5589 int32
	_ = v5589
	var v5619 int32
	_ = v5619
	var v5623 int32
	_ = v5623
	var v5628 int32
	_ = v5628
	var v5629 int32
	_ = v5629
	var v5631 int32
	_ = v5631
	var v5634 int32
	_ = v5634
	var v5636 int32
	_ = v5636
	var v5639 int32
	_ = v5639
	var v5641 int32
	_ = v5641
	var v5644 int32
	_ = v5644
	var v5646 int32
	_ = v5646
	var v5649 int32
	_ = v5649
	var v5651 int32
	_ = v5651
	var v5654 int32
	_ = v5654
	var v5656 int32
	_ = v5656
	var v5657 int32
	_ = v5657
	var v5660 int32
	_ = v5660
	var v5663 int32
	_ = v5663
	var v5665 int32
	_ = v5665
	var v5668 int32
	_ = v5668
	var v5671 int32
	_ = v5671
	var v5682 int32
	_ = v5682
	var v5699 int32
	_ = v5699
	var v5702 int32
	_ = v5702
	var v5703 int32
	_ = v5703
	var v5705 int32
	_ = v5705
	var v5706 int32
	_ = v5706
	var v5708 int32
	_ = v5708
	var v5709 int32
	_ = v5709
	var v5711 int32
	_ = v5711
	var v5712 int32
	_ = v5712
	var v5714 int32
	_ = v5714
	var v5715 int32
	_ = v5715
	var v5717 int32
	_ = v5717
	var v5720 int32
	_ = v5720
	var v5721 int32
	_ = v5721
	var v5723 int32
	_ = v5723
	var v5726 int32
	_ = v5726
	var v5729 int32
	_ = v5729
	var v5730 int32
	_ = v5730
	var v5735 int32
	_ = v5735
	var v5761 int32
	_ = v5761
	var v5762 int32
	_ = v5762
	var v5783 int32
	_ = v5783
	var v5784 int32
	_ = v5784
	var v5787 int32
	_ = v5787
	var v5788 int32
	_ = v5788
	var v5791 int32
	_ = v5791
	var v5792 int32
	_ = v5792
	var v5795 int32
	_ = v5795
	var v5799 int32
	_ = v5799
	var v5802 int32
	_ = v5802
	var v5810 int32
	_ = v5810
	var v5834 int32
	_ = v5834
	var v5835 int32
	_ = v5835
	var v5838 int32
	_ = v5838
	var v5841 int32
	_ = v5841
	var v5843 int32
	_ = v5843
	var v5850 int32
	_ = v5850
	var v5876 int32
	_ = v5876
	var v5928 int32
	_ = v5928
	var v5932 int64
	_ = v5932
	var v5934 int32
	_ = v5934
	var v5935 int32
	_ = v5935
	var v5937 int32
	_ = v5937
	var v5941 int32
	_ = v5941
	var v5943 int32
	_ = v5943
	var v5947 int32
	_ = v5947
	var v5948 int32
	_ = v5948
	var v5955 int32
	_ = v5955
	var v5960 int32
	_ = v5960
	var v5962 int32
	_ = v5962
	var v5963 int32
	_ = v5963
	var v5964 int32
	_ = v5964
	var v5971 int32
	_ = v5971
	var v5987 int32
	_ = v5987
	var v5989 int32
	_ = v5989
	var v5992 int32
	_ = v5992
	var v5993 int32
	_ = v5993
	var v5995 int32
	_ = v5995
	var v5996 int32
	_ = v5996
	var v6000 int32
	_ = v6000
	var v6001 int32
	_ = v6001
	var v6002 int32
	_ = v6002
	var v6003 int32
	_ = v6003
	var v6004 int32
	_ = v6004
	var v6005 int32
	_ = v6005
	var v6009 int32
	_ = v6009
	var v6014 int32
	_ = v6014
	var v6018 int32
	_ = v6018
	var v6020 int32
	_ = v6020
	var v6024 int32
	_ = v6024
	var v6026 int32
	_ = v6026
	var v6030 int32
	_ = v6030
	var v6031 int32
	_ = v6031
	var v6037 int32
	_ = v6037
	var v6042 int32
	_ = v6042
	var v6044 int32
	_ = v6044
	var v6045 int32
	_ = v6045
	var v6046 int32
	_ = v6046
	var v6051 int32
	_ = v6051
	var v6053 int32
	_ = v6053
	var v6059 int32
	_ = v6059
	var v6061 int32
	_ = v6061
	var v6062 int32
	_ = v6062
	var v6063 int32
	_ = v6063
	var v6064 int32
	_ = v6064
	var v6065 int32
	_ = v6065
	var v6077 int32
	_ = v6077
	var v6082 int32
	_ = v6082
	var v6091 int32
	_ = v6091
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
	var v6101 int32
	_ = v6101
	var v6102 int32
	_ = v6102
	var v6103 int32
	_ = v6103
	var v6104 int32
	_ = v6104
	var v6105 int32
	_ = v6105
	var v6106 int32
	_ = v6106
	var v6107 int32
	_ = v6107
	var v6108 int32
	_ = v6108
	var v6109 int32
	_ = v6109
	var v6110 int32
	_ = v6110
	var v6111 int32
	_ = v6111
	var v6112 int32
	_ = v6112
	var v6113 int32
	_ = v6113
	var v6114 int32
	_ = v6114
	var v6115 int32
	_ = v6115
	var v6116 int32
	_ = v6116
	var v6117 int32
	_ = v6117
	var v6118 int32
	_ = v6118
	var v6119 int32
	_ = v6119
	var v6120 int32
	_ = v6120
	var v6121 int32
	_ = v6121
	var v6122 int32
	_ = v6122
	var v6123 int32
	_ = v6123
	var v6124 int32
	_ = v6124
	var v6125 int32
	_ = v6125
	var v6126 int32
	_ = v6126
	var v6127 int32
	_ = v6127
	var v6128 int32
	_ = v6128
	var v6129 int32
	_ = v6129
	var v6130 int32
	_ = v6130
	var v6131 int32
	_ = v6131
	var v6132 int32
	_ = v6132
	var v6133 int32
	_ = v6133
	var v6134 int32
	_ = v6134
	var v6135 int32
	_ = v6135
	var v6136 int32
	_ = v6136
	var v6137 int32
	_ = v6137
	var v6138 int32
	_ = v6138
	var v6139 int32
	_ = v6139
	var v6140 int32
	_ = v6140
	var v6141 int32
	_ = v6141
	var v6142 int32
	_ = v6142
	var v6143 int32
	_ = v6143
	var v6144 int32
	_ = v6144
	var v6145 int32
	_ = v6145
	var v6146 int32
	_ = v6146
	var v6147 int32
	_ = v6147
	var v6148 int32
	_ = v6148
	var v6149 int32
	_ = v6149
	var v6150 int32
	_ = v6150
	var v6151 int32
	_ = v6151
	var v6152 int32
	_ = v6152
	var v6153 int32
	_ = v6153
	var v6156 int32
	_ = v6156
	var v6183 int32
	_ = v6183
	var v6187 int32
	_ = v6187
	var v6188 int32
	_ = v6188
	var v6191 int32
	_ = v6191
	var v6194 int32
	_ = v6194
	var v6195 int32
	_ = v6195
	var v6200 int32
	_ = v6200
	var v6204 int32
	_ = v6204
	var v6207 int32
	_ = v6207
	var v6211 int32
	_ = v6211
	var v6215 int32
	_ = v6215
	var v6217 int32
	_ = v6217
	var v6218 int32
	_ = v6218
	var v6220 int32
	_ = v6220
	var v6221 int32
	_ = v6221
	var v6222 int32
	_ = v6222
	var v6225 int32
	_ = v6225
	var v6227 int32
	_ = v6227
	var v6228 int32
	_ = v6228
	var v6238 int32
	_ = v6238
	var v6262 int32
	_ = v6262
	var v6267 int32
	_ = v6267
	var v6271 int32
	_ = v6271
	var v6273 int32
	_ = v6273
	var v6276 int32
	_ = v6276
	var v6282 int32
	_ = v6282
	var v6287 int32
	_ = v6287
	var v6291 int32
	_ = v6291
	var v6295 int32
	_ = v6295
	var v6300 int32
	_ = v6300
	var v6304 int32
	_ = v6304
	var v6308 int32
	_ = v6308
	var v6313 int32
	_ = v6313
	var v6315 int32
	_ = v6315
	var v6317 int32
	_ = v6317
	var v6321 int32
	_ = v6321
	var v6322 int32
	_ = v6322
	var v6324 int32
	_ = v6324
	var v6331 int32
	_ = v6331
	var v6333 int32
	_ = v6333
	var v6337 int32
	_ = v6337
	var v6339 int32
	_ = v6339
	var v6341 int32
	_ = v6341
	var v6343 int32
	_ = v6343
	var v6345 int32
	_ = v6345
	var v6349 int32
	_ = v6349
	var v6351 int32
	_ = v6351
	var v6354 int32
	_ = v6354
	var v6357 int32
	_ = v6357
	var v6361 int32
	_ = v6361
	var v6364 int32
	_ = v6364
	var v6366 int32
	_ = v6366
	var v6372 int32
	_ = v6372
	var v6377 int32
	_ = v6377
	var v6383 int32
	_ = v6383
	var v6388 int32
	_ = v6388
	var v6392 int32
	_ = v6392
	var v6395 int32
	_ = v6395
	var v6401 int32
	_ = v6401
	var v6404 int32
	_ = v6404
	var v6407 int32
	_ = v6407
	var v6413 int32
	_ = v6413
	var v6417 int32
	_ = v6417
	var v6421 int32
	_ = v6421
	var v6426 int32
	_ = v6426
	v4 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(96)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_InitStandaloneProcess(m, v30)
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
	v34 = l0 - int32(1)
	F_InitializeGUCOptions(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = int32(0)
	v59 = v4
	goto L11
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6417 = m.ExcPending
	if v6417 != 0 {
		goto L1
	} else {
		goto L1324
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[0])) = int32(2)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6413 = m.ExcPending
	if v6413 != 0 {
		goto L1
	} else {
		goto L1323
	}
L6:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6407 = m.ExcPending
	if v6407 != 0 {
		goto L1
	} else {
		goto L1322
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v30
	F_write_stderr(m, int32(_a_F_BootstrapModeMain_0), v28+int32(16))
	mBase = m.M
	v6401 = m.ExcPending
	if v6401 != 0 {
		goto L1
	} else {
		goto L1320
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
	F_write_stderr(m, int32(_a_F_BootstrapModeMain_1), v28)
	mBase = m.M
	v6392 = m.ExcPending
	if v6392 != 0 {
		goto L1
	} else {
		goto L1318
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v122
	F_errmsg(m, int32(_a_F_BootstrapModeMain_2), v28+int32(32))
	mBase = m.M
	v6383 = m.ExcPending
	if v6383 != 0 {
		goto L1
	} else {
		goto L1316
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6361 = m.ExcPending
	if v6361 != 0 {
		goto L1
	} else {
		goto L1312
	}
L11:
	;
	v67 = F_getopt(m, v34, l1+int32(4), int32(_a_F_BootstrapModeMain_3))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L12:
	;
	if v67 != int32(-1) {
		goto L8
	} else {
		goto L89
	}
L13:
	;
	goto L12
L14:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_4), v297, int32(0), int32(1))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L88
	}
L15:
	;
	v177 = int32(_a_F_BootstrapModeMain_5)
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	goto L59
L16:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_6), int32(_a_F_BootstrapModeMain_7), int32(1), int32(4))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L55
	}
L17:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v152
	v158 = F_psprintf(m, int32(_a_F_BootstrapModeMain_8), v28+int32(80))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L51
	}
L18:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	v149 = F_pstrdup(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L50
	}
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	F_ParseLongOption(m, v104, v28+int32(92), v28+int32(88))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L38
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	v81 = F_strcmp(m, int32(_a_F_BootstrapModeMain_9), v79)
	mBase = m.M
	if v81 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_10), v73, int32(1), int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	switch v67 - int32(45) {
	case 0:
		goto L20
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 24, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 56, 57, 58, 59, 60, 61, 63, 64, 65, 66, 67, 68:
		goto L8
	case 21:
		goto L21
	case 23:
		goto L18
	case 25:
		goto L16
	case 43:
		goto L14
	case 54:
		goto L19
	case 55:
		goto L17
	case 62:
		v41 = int32(1)
		goto L11
	case 69:
		goto L15
	default:
		goto L13
	}
L23:
	;
	goto L11
L24:
	;
	if v100 != int32(5) {
		goto L10
	} else {
		goto L37
	}
L25:
	;
	v100 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v86 = F_strcmp(m, int32(_a_F_BootstrapModeMain_11), v79)
	mBase = m.M
	if v86 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v100 = int32(1)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v91 = F_strcmp(m, int32(_a_F_BootstrapModeMain_12), v79)
	mBase = m.M
	if v91 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v100 = int32(3)
	goto L24
L32:
	;
	goto L33
L33:
	;
	v98 = F_strcmp(m, int32(_a_F_BootstrapModeMain_13), v79)
	mBase = m.M
	if v98 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v99 = int32(5)
	goto L36
L35:
	;
	v99 = int32(4)
	goto L36
L36:
	;
	v100 = v99
	goto L24
L37:
	;
	goto L19
L38:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v28)+88))
	if v111 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v28)+92))
	F_SetConfigOption(m, v136, v111, int32(1), int32(4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L47
	}
L42:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	if v67 == int32(45) {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v122
	F_errmsg(m, int32(_a_F_BootstrapModeMain_14), v28+int32(48))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(259), int32(_a_F_BootstrapModeMain_16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v28)+92))
	F_pfree(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v28)+88))
	F_pfree(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L11
L50:
	;
	v59 = v149
	goto L11
L51:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_17), v158, int32(1), int32(4))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_18), v158, int32(1), int32(4))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_pfree(m, v158)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L11
L55:
	;
	goto L11
L56:
	;
	goto L11
L57:
	;
	v292 = F_strlen(m, v281)
	mBase = m.M
	goto L56
L59:
	;
	goto L60
L60:
	;
	v186 = int32(1023)
	if (v177^v179)&int32(3) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v285)
	goto L57
L62:
	;
	v266 = v261
	v267 = v262
	v268 = v263
	goto L84
L63:
	;
	if v256 == int32(0) {
		v281 = v254
		v282 = v255
		goto L61
	} else {
		goto L83
	}
L64:
	;
	v254 = v179
	v255 = v177
	v256 = v186
	goto L63
L65:
	;
	goto L66
L66:
	;
	if v179&int32(3) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v223 == int32(0) {
		v281 = v220
		v282 = v221
		goto L61
	} else {
		goto L76
	}
L68:
	;
	v220 = v179
	v221 = v177
	v222 = v186
	v223 = int32(1)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v199 = v179
	v200 = v177
	v201 = v186
	goto L71
L71:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v203)
	if v203 == int32(0) {
		v261 = v199
		v262 = v200
		v263 = v201
		goto L62
	} else {
		goto L73
	}
L72:
	;
	v220 = v214
	v221 = v208
	v222 = v210
	v223 = v212
	goto L67
L73:
	;
	v207 = int32(1)
	v208 = v200 + v207
	v210 = v201 - v207
	v211 = int32(0)
	v212 = base.B2i32(v210 != v211)
	v214 = v199 + v207
	if v214&int32(3) == v211 {
		v220 = v214
		v221 = v208
		v222 = v210
		v223 = v212
		goto L67
	} else {
		goto L74
	}
L74:
	;
	if v210 != 0 {
		v199 = v214
		v200 = v208
		v201 = v210
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	if v226 == int32(0) {
		v254 = v220
		v255 = v221
		v256 = v222
		goto L63
	} else {
		goto L77
	}
L77:
	;
	if base.Ui32(v222) < base.Ui32(int32(4)) {
		v254 = v220
		v255 = v221
		v256 = v222
		goto L63
	} else {
		goto L78
	}
L78:
	;
	v232 = v220
	v233 = v221
	v234 = v222
	goto L79
L79:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v240 = int32(-2139062144)
	if (int32(16843008)-v237|v237)&v240 != v240 {
		v261 = v232
		v262 = v233
		v263 = v234
		goto L62
	} else {
		goto L81
	}
L80:
	;
	v254 = v248
	v255 = v246
	v256 = v250
	goto L63
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v237
	v245 = int32(4)
	v246 = v233 + v245
	v248 = v232 + v245
	v250 = v234 - v245
	if base.Ui32(int32(3)) < base.Ui32(v250) {
		v232 = v248
		v233 = v246
		v234 = v250
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v261 = v254
	v262 = v255
	v263 = v256
	goto L62
L84:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	*(*uint8)(unsafe.Add(mBase, uint32(v267))) = uint8(v270)
	if v270 == int32(0) {
		v281 = v266
		v282 = v267
		goto L61
	} else {
		goto L86
	}
L85:
	;
	v281 = v277
	v282 = v275
	goto L61
L86:
	;
	v274 = int32(1)
	v275 = v267 + v274
	v277 = v266 + v274
	v279 = v268 - v274
	if v279 != 0 {
		v266 = v277
		v267 = v275
		v268 = v279
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	goto L11
L89:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[2]))
	if v34 != v305 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v307 = F_SelectConfigFiles(m, v59, v30)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v307 == int32(0) {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	F_checkDataDir(m)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_ChangeToDataDir(m)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_CreateDataDirLockFile(m, int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v319 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[3])) = uint8(v319)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[0])) = int32(0)
	F_InitializeMaxBackends(m)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_InitPostmasterChildSlots(m)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v332 = int32(1)
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[4]))
	if v335&(v335-v332) != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	F_CreateSharedMemoryAndSemaphores(m)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L108
	}
L99:
	;
	v342 = v332 << (uint(int32(32)-base.I32_clz(v335)) % 32)
	goto L101
L100:
	;
	v342 = v335
	goto L101
L101:
	;
	if base.Ui32(v342) <= base.Ui32(int32(31)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v345 = int32(31)
	goto L104
L103:
	;
	v345 = v342
	goto L104
L104:
	;
	if base.Ui32(int32(_a_F_BootstrapModeMain_19)) <= base.Ui32(v342) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v350 = int32(1024)
	goto L107
L106:
	;
	v350 = int32(base.Ui32(v345) >> (uint(int32(4)) % 32))
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[5])) = v350
	goto L98
L108:
	;
	F_set_max_safe_fds(m)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	F_InitProcess(m)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_BaseInit(m)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v361 = int32(0)
	v363 = m.G0
	v365 = v363 - int32(144)
	m.G0 = v365
	switch int32(2) {
	case 0, 2:
		v375 = v361
		goto L114
	default:
		goto L115
	}
L113:
	;
	v403 = int32(0)
	v405 = m.G0
	v407 = v405 - int32(144)
	m.G0 = v407
	switch int32(2) {
	case 0, 2:
		v417 = v403
		goto L127
	default:
		goto L128
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v365)+4)) = v375
	F_sigemptyset(m, v365+int32(8))
	mBase = m.M
	goto L117
L115:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[6])) = v361
	v375 = int32(_a_F_BootstrapModeMain_20)
	goto L114
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v365)+136)) = int32(268435456)
	v387 = v365 + int32(4)
	goto L121
L119:
	;
	m.G0 = v365 + int32(144)
	goto L113
L121:
	;
	goto L122
L122:
	;
	if v387 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v398 = F___memcpy(m, int32(_a_F_BootstrapModeMain_21), v387, int32(140))
	mBase = m.M
	goto L125
L124:
	;
	goto L125
L125:
	;
	goto L119
L126:
	;
	v445 = int32(0)
	v447 = m.G0
	v449 = v447 - int32(144)
	m.G0 = v449
	switch int32(2) {
	case 0, 2:
		v459 = v445
		goto L140
	default:
		goto L141
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407)+4)) = v417
	F_sigemptyset(m, v407+int32(8))
	mBase = m.M
	goto L130
L128:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[7])) = v403
	v417 = int32(_a_F_BootstrapModeMain_20)
	goto L127
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407)+136)) = int32(268435456)
	v429 = v407 + int32(4)
	goto L134
L132:
	;
	m.G0 = v407 + int32(144)
	goto L126
L134:
	;
	goto L135
L135:
	;
	if v429 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v440 = F___memcpy(m, int32(_a_F_BootstrapModeMain_22), v429, int32(140))
	mBase = m.M
	goto L138
L137:
	;
	goto L138
L138:
	;
	goto L132
L139:
	;
	v487 = int32(0)
	v489 = m.G0
	v491 = v489 - int32(144)
	m.G0 = v491
	switch int32(2) {
	case 0, 2:
		v501 = v487
		goto L153
	default:
		goto L154
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v449)+4)) = v459
	F_sigemptyset(m, v449+int32(8))
	mBase = m.M
	goto L143
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[8])) = v445
	v459 = int32(_a_F_BootstrapModeMain_20)
	goto L140
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v449)+136)) = int32(268435456)
	v471 = v449 + int32(4)
	goto L147
L145:
	;
	m.G0 = v449 + int32(144)
	goto L139
L147:
	;
	goto L148
L148:
	;
	if v471 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v482 = F___memcpy(m, int32(_a_F_BootstrapModeMain_23), v471, int32(140))
	mBase = m.M
	goto L151
L150:
	;
	goto L151
L151:
	;
	goto L145
L152:
	;
	v528 = m.G0
	v530 = v528 - int32(_a_F_BootstrapModeMain_24)
	m.G0 = v530
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9]))
	v537 = F_LWLockAcquire(m, v533+int32(1152), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L165
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491)+4)) = v501
	F_sigemptyset(m, v491+int32(8))
	mBase = m.M
	goto L156
L154:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[10])) = v487
	v501 = int32(_a_F_BootstrapModeMain_20)
	goto L153
L156:
	;
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v491)+136)) = int32(268435456)
	v513 = v491 + int32(4)
	goto L160
L158:
	;
	m.G0 = v491 + int32(144)
	goto L152
L160:
	;
	goto L161
L161:
	;
	if v513 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v524 = F___memcpy(m, int32(_a_F_BootstrapModeMain_25), v513, int32(140))
	mBase = m.M
	goto L164
L163:
	;
	goto L164
L164:
	;
	goto L158
L165:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[11]))
	v541 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v540)+320)) = uint8(v541)
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9]))
	F_LWLockRelease(m, v544+int32(1152))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F___gettimeofday(m, v530-int32(-64))
	mBase = m.M
	v552 = int64(*(*int32)(unsafe.Add(mBase, uint32(v530)+72)))
	v553 = *(*int64)(unsafe.Add(mBase, uint32(v530)+64))
	v555 = F_palloc(m, int32(_a_F_BootstrapModeMain_19))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v564 = F__emscripten_memset_bulkmem(m, (v555+int32(_a_F_BootstrapModeMain_26))&int32(-8192), base.I32_extend8_s(int32(0)), int32(_a_F_BootstrapModeMain_27))
	mBase = m.M
	goto L168
L168:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[12]))
	v568 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[13])))
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[14]))
	v571 = F___time(m)
	mBase = m.M
	v572 = int32(_a_F_BootstrapModeMain_28)
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = int32(_a_F_BootstrapModeMain_29)
	*(*int64)(unsafe.Add(mBase, uint32(v573)+8)) = int64(3)
	v579 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[15]))
	v580 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v579)+4)) = v580
	F_MultiXactSetNextMXact(m, int32(1), v580)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_AdvanceOldestClogXid(m, int32(3))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	F_SetTransactionIdLimit(m, int32(3), int32(1))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v593 = int32(1)
	F_SetMultiXactIdLimit(m, v593, v593, v593)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v598 = int32(0)
	F_SetCommitTsLimit(m, v598, v598)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v564))) = int64(4295151896)
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[12]))
	v606 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+48)) = v606
	*(*int32)(unsafe.Add(mBase, uint32(v564)+36)) = int32(_a_F_BootstrapModeMain_27)
	*(*int32)(unsafe.Add(mBase, uint32(v564)+32)) = v605
	v617 = v552<<(uint(int64(12))%64) | v553<<(uint(int64(32))%64) | int64(42)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+24)) = v617
	v619 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v564)+146)) = v619
	*(*int64)(unsafe.Add(mBase, uint32(v564)+138)) = v606
	*(*int64)(unsafe.Add(mBase, uint32(v564)+130)) = v571
	v624 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v564)+122)) = v624
	v626 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+114)) = v626
	*(*int64)(unsafe.Add(mBase, uint32(v564)+106)) = int64(12884901888)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+98)) = int64(4294977296)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+90)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v564)+86)) = v570
	*(*uint8)(unsafe.Add(mBase, uint32(v564)+82)) = uint8(v568)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+74)) = v626
	v638 = int32(40)
	v640 = base.I64_extend_i32_u(v566 + v638)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+66)) = v640
	v642 = int32(_a_F_BootstrapModeMain_30)
	*(*uint16)(unsafe.Add(mBase, uint32(v564)+64)) = uint16(v642)
	*(*uint16)(unsafe.Add(mBase, uint32(v564)+56)) = uint16(v619)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+40)) = int64(114)
	*(*int64)(unsafe.Add(mBase, uint32(v564)+8)) = base.I64_extend_i32_s(v605)
	v650 = int32(-1)
	v654 = m.Env.Pgmem_crc32c(m, v650, v564|int32(64), int32(90))
	mBase = m.M
	v658 = m.Env.Pgmem_crc32c(m, v654, v564|v638, int32(20))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v564)+60)) = v658 ^ v650
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[16])) = v624
	v668 = F_XLogFileInit(m, int64(1), v624)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[17])) = v668
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(0)
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v675))) = int32(167772229)
	v678 = int32(_a_F_BootstrapModeMain_27)
	v679 = F_write(m, v668, v564, v678)
	mBase = m.M
	if v679 != v678 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v683 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	goto L177
L177:
	;
	v704 = int32(_a_F_BootstrapModeMain_31)
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	v706 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v705))) = v706
	v709 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v709))) = int32(167772228)
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[17]))
	v716 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[20])))
	if v716 != int32(1) {
		v730 = v706
		goto L192
	} else {
		goto L193
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(51)
	goto L180
L179:
	;
	goto L180
L180:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_32), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_33), int32(_a_F_BootstrapModeMain_34), int32(_a_F_BootstrapModeMain_35))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	v1147 = int32(0)
	F_InitPostgres(m, v1147, v1147, v1147, v1147, v1147, v1147)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L285
	}
L186:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L1
	} else {
		goto L281
	}
L187:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L277
	}
L188:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L273
	}
L189:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L1
	} else {
		goto L269
	}
L190:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L265
	}
L191:
	;
	if v730 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L192:
	;
	goto L191
L193:
	;
	goto L194
L194:
	;
	v721 = F_fsync(m, v713)
	mBase = m.M
	if v721 != int32(-1) {
		v730 = v721
		goto L192
	} else {
		goto L196
	}
L195:
	;
	v730 = int32(-1)
	goto L192
L196:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v725 == int32(27) {
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v734))) = int32(0)
	v738 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[17]))
	v739 = F_close(m, v738)
	mBase = m.M
	if v739 != 0 {
		goto L190
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L261
	}
L201:
	;
	v741 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[17])) = v741
	v746 = int32(0)
	v750 = m.G0
	v752 = v750 - int32(16)
	m.G0 = v752
	*(*int32)(unsafe.Add(mBase, uint32(v752))) = v746
	v758 = F_open(m, int32(_a_F_BootstrapModeMain_36), v746, v752)
	mBase = m.M
	if v758 != v741 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v791 == int32(0) {
		goto L189
	} else {
		goto L215
	}
L203:
	;
	goto L207
L204:
	;
	v791 = v746
	goto L205
L205:
	;
	m.G0 = v752 + int32(16)
	goto L202
L206:
	;
	v786 = F_close(m, v758)
	mBase = m.M
	v791 = v784
	goto L205
L207:
	;
	v764 = v530 + int32(80)
	v765 = int32(32)
	goto L208
L208:
	;
	v770 = F_read(m, v758, v764, v765)
	mBase = m.M
	if v770 <= int32(0) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v784 = int32(1)
	goto L206
L210:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v774 == int32(27) {
		goto L208
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v779 = v765 - v770
	if v779 != 0 {
		v764 = v764 + v770
		v765 = v779
		goto L208
	} else {
		goto L214
	}
L213:
	;
	v784 = int32(0)
	goto L206
L214:
	;
	goto L209
L215:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[21]))
	v805 = F__emscripten_memset_bulkmem(m, v799+int32(8), base.I32_extend8_s(int32(0)), int32(288))
	mBase = m.M
	goto L216
L216:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v799))) = v617
	v807 = *(*int64)(unsafe.Add(mBase, uint32(v530)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v799)+257)) = v807
	v809 = *(*int64)(unsafe.Add(mBase, uint32(v530)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v799)+265)) = v809
	v811 = *(*int64)(unsafe.Add(mBase, uint32(v530)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v799)+273)) = v811
	v813 = *(*int64)(unsafe.Add(mBase, uint32(v530)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v799)+281)) = v813
	*(*int64)(unsafe.Add(mBase, uint32(v799)+128)) = int64(1000)
	v817 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v799)+16)) = v817
	v820 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v799)+180)) = v820
	v823 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v799)+184)) = v823
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[24]))
	*(*int32)(unsafe.Add(mBase, uint32(v799)+188)) = v826
	v829 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[25]))
	*(*int32)(unsafe.Add(mBase, uint32(v799)+192)) = v829
	v832 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v799)+196)) = v832
	v835 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v799)+172)) = v835
	v838 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[26])))
	*(*uint8)(unsafe.Add(mBase, uint32(v799)+176)) = uint8(v838)
	v841 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[27])))
	*(*int32)(unsafe.Add(mBase, uint32(v799)+252)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v799)+200)) = uint8(v841)
	v844 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v799)+120)) = v844
	*(*int64)(unsafe.Add(mBase, uint32(v799)+112)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+104)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v799)+96)) = v817
	v851 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+88)) = v851
	*(*int64)(unsafe.Add(mBase, uint32(v799)+80)) = int64(12884901888)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+72)) = int64(4294977296)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+64)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v799)+60)) = v570
	*(*uint8)(unsafe.Add(mBase, uint32(v799)+56)) = uint8(v568)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+48)) = v851
	*(*int64)(unsafe.Add(mBase, uint32(v799)+40)) = v640
	*(*int64)(unsafe.Add(mBase, uint32(v799)+32)) = v640
	*(*int64)(unsafe.Add(mBase, uint32(v799)+24)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v799)+224)) = int32(_a_F_BootstrapModeMain_27)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+216)) = int64(562949953429504)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+208)) = int64(4698053236609777664)
	*(*int32)(unsafe.Add(mBase, uint32(v799)+204)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+8)) = int64(869757897079260936)
	v877 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[12]))
	v878 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v799)+292)) = v878
	*(*uint8)(unsafe.Add(mBase, uint32(v799)+256)) = uint8(v817)
	*(*uint8)(unsafe.Add(mBase, uint32(v799)+248)) = uint8(v844)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+240)) = int64(8796093024204)
	*(*int64)(unsafe.Add(mBase, uint32(v799)+232)) = int64(137438953536)
	*(*int32)(unsafe.Add(mBase, uint32(v799)+228)) = v877
	v891 = m.Env.Pgmem_crc32c(m, v878, v799, int32(292))
	mBase = m.M
	v893 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v893)+292)) = v891 ^ v878
	v902 = F__emscripten_memset_bulkmem(m, v530+int32(376), base.I32_extend8_s(v844), int32(_a_F_BootstrapModeMain_37))
	mBase = m.M
	goto L217
L217:
	;
	goto L219
L218:
	;
	v910 = F_BasicOpenFile(m, int32(_a_F_BootstrapModeMain_38), int32(194))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L222
	}
L219:
	;
	v906 = F__emscripten_memcpy_bulkmem(m, v530+int32(80), v893, int32(296))
	mBase = m.M
	goto L221
L221:
	;
	goto L218
L222:
	;
	if v910 < int32(0) {
		goto L188
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(0)
	v918 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v918))) = int32(167772172)
	v923 = int32(_a_F_BootstrapModeMain_27)
	v924 = F_write(m, v910, v530+int32(80), v923)
	mBase = m.M
	if v924 != v923 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v928 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v928 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	goto L226
L226:
	;
	v952 = int32(_a_F_BootstrapModeMain_31)
	v953 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	v954 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v953))) = v954
	v957 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v957))) = int32(167772170)
	v962 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[20])))
	if v962 != int32(1) {
		v976 = v954
		goto L235
	} else {
		goto L236
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(51)
	goto L229
L228:
	;
	goto L229
L229:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v530)+48)) = int32(_a_F_BootstrapModeMain_38)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_39), v530+int32(48))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_33), int32(_a_F_BootstrapModeMain_40), int32(_a_F_BootstrapModeMain_41))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	if v976 != 0 {
		goto L187
	} else {
		goto L241
	}
L235:
	;
	goto L234
L236:
	;
	goto L237
L237:
	;
	v967 = F_fsync(m, v910)
	mBase = m.M
	if v967 != int32(-1) {
		v976 = v967
		goto L235
	} else {
		goto L239
	}
L238:
	;
	v976 = int32(-1)
	goto L235
L239:
	;
	v971 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v971 == int32(27) {
		goto L237
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	v978 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v978))) = int32(0)
	v981 = F_close(m, v910)
	mBase = m.M
	if v981 != 0 {
		goto L186
	} else {
		goto L242
	}
L242:
	;
	v983 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[28]))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v983)+28))
	v986 = F_LWLockAcquire(m, v984, int32(0))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v988 = int32(_a_F_BootstrapModeMain_42)
	v991 = F_SimpleLruZeroPage(m, v988, int64(0))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	F_SimpleLruWritePage(m, v988, v991)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_LWLockRelease(m, v984)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	v998 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[29]))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v998)+28))
	v1001 = F_LWLockAcquire(m, v999, int32(0))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v1003 = int32(_a_F_BootstrapModeMain_43)
	v1006 = F_SimpleLruZeroPage(m, v1003, int64(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	F_SimpleLruWritePage(m, v1003, v1006)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_LWLockRelease(m, v999)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+28))
	v1016 = F_LWLockAcquire(m, v1014, int32(0))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1018 = int32(_a_F_BootstrapModeMain_44)
	v1021 = F_SimpleLruZeroPage(m, v1018, int64(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	F_SimpleLruWritePage(m, v1018, v1021)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	F_LWLockRelease(m, v1014)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+28))
	v1031 = F_LWLockAcquire(m, v1029, int32(0))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v1033 = int32(_a_F_BootstrapModeMain_45)
	v1036 = F_SimpleLruZeroPage(m, v1033, int64(0))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	F_SimpleLruWritePage(m, v1033, v1036)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	F_LWLockRelease(m, v1029)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	F_pfree(m, v555)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	F_ReadControlFile(m)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	m.G0 = v530 + int32(_a_F_BootstrapModeMain_24)
	goto L185
L261:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_46), int32(0))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_33), int32(_a_F_BootstrapModeMain_47), int32(_a_F_BootstrapModeMain_35))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_48), int32(0))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_33), int32(_a_F_BootstrapModeMain_49), int32(_a_F_BootstrapModeMain_35))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_50), int32(0))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_33), int32(_a_F_BootstrapModeMain_51), int32(_a_F_BootstrapModeMain_52))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v530))) = int32(_a_F_BootstrapModeMain_38)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_53), v530)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_33), int32(_a_F_BootstrapModeMain_54), int32(_a_F_BootstrapModeMain_41))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v530)+32)) = int32(_a_F_BootstrapModeMain_38)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_55), v530+int32(32))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_33), int32(_a_F_BootstrapModeMain_56), int32(_a_F_BootstrapModeMain_41))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v530)+16)) = int32(_a_F_BootstrapModeMain_38)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_57), v530+int32(16))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_33), int32(_a_F_BootstrapModeMain_58), int32(_a_F_BootstrapModeMain_41))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L285:
	;
	v1159 = F__emscripten_memset_bulkmem(m, int32(_a_F_BootstrapModeMain_59), base.I32_extend8_s(int32(0)), int32(160))
	mBase = m.M
	goto L286
L286:
	;
	v1161 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[32])) = v1161
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[33])) = v1161
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[34])) = v1161
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[35])) = v1161
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[36])) = v1161
	v1176 = v28 + int32(92)
	if v1176 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	if v1215 != 0 {
		goto L4
	} else {
		goto L296
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(28)
	v1215 = int32(1)
	goto L287
L289:
	;
	goto L290
L290:
	;
	v1184 = F_palloc(m, int32(96))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1176))) = v1184
	if v1184 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(48)
	v1215 = int32(1)
	goto L287
L293:
	;
	goto L294
L294:
	;
	v1196 = F__emscripten_memset_bulkmem(m, v1184, base.I32_extend8_s(int32(0)), int32(96))
	mBase = m.M
	goto L295
L295:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1176)))
	v1198 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1197)+60)) = v1198
	v1200 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1197)+52)) = v1200
	*(*int32)(unsafe.Add(mBase, uint32(v1197)+44)) = v1198
	*(*int64)(unsafe.Add(mBase, uint32(v1197)+36)) = v1200
	*(*int64)(unsafe.Add(mBase, uint32(v1197)+4)) = v1200
	*(*int64)(unsafe.Add(mBase, uint32(v1197)+12)) = v1200
	*(*int32)(unsafe.Add(mBase, uint32(v1197)+20)) = v1198
	v1215 = v1198
	goto L287
L296:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v28)+92))
	v1219 = m.G0
	v1221 = v1219 - int32(1344)
	m.G0 = v1221
	v1224 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[37]))
	v1227 = v1221 + int32(128)
	v1229 = v1221 + int32(928)
	v1234 = v1218
	v1238 = v1227
	v1240 = int32(-2)
	v1241 = v1221
	v1243 = v1229
	v1244 = v1229
	v1245 = v4
	v1248 = v1224
	v1250 = v1227
	v1251 = int32(200)
	goto L303
L298:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6315 = m.ExcPending
	if v6315 != 0 {
		goto L1
	} else {
		goto L1302
	}
L299:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6304 = m.ExcPending
	if v6304 != 0 {
		goto L1
	} else {
		goto L1299
	}
L300:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6291 = m.ExcPending
	if v6291 != 0 {
		goto L1
	} else {
		goto L1296
	}
L301:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6271 = m.ExcPending
	if v6271 != 0 {
		goto L1
	} else {
		goto L1293
	}
L302:
	;
	F_boot_yyerror(m, v1234, int32(_a_F_BootstrapModeMain_60))
	mBase = m.M
	v6267 = m.ExcPending
	if v6267 != 0 {
		goto L1
	} else {
		goto L1292
	}
L303:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1244))) = uint16(v1245)
	v1258 = v1251 << (uint(int32(1)) % 32)
	if base.Ui32(v1243+v1258-int32(2)) <= base.Ui32(v1244) {
		goto L309
	} else {
		goto L310
	}
L304:
	;
	F_boot_yyerror(m, v6238, int32(_a_F_BootstrapModeMain_61))
	mBase = m.M
	v6262 = m.ExcPending
	if v6262 != 0 {
		goto L1
	} else {
		goto L1291
	}
L305:
	;
	goto L304
L306:
	;
	v1234 = v6211
	v1238 = v6215
	v1240 = v6217
	v1241 = v6218
	v1243 = v6220
	v1244 = v6221 + int32(2)
	v1245 = v6222
	v1248 = v6225
	v1250 = v6227
	v1251 = v6228
	goto L303
L307:
	;
	v3801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3788)+uint32(_c_F_BootstrapModeMain[38]))))
	if v3801 == int32(0) {
		v6238 = v3777
		goto L305
	} else {
		goto L729
	}
L308:
	;
	if v3751+int32(928) != v3753 {
		goto L725
	} else {
		goto L726
	}
L309:
	;
	if base.Ui32(int32(_a_F_BootstrapModeMain_62)) < base.Ui32(v1251) {
		goto L302
	} else {
		goto L312
	}
L310:
	;
	v1312 = v1238
	v1313 = v1243
	v1314 = v1244
	v1315 = v1250
	v1316 = v1251
	goto L311
L311:
	;
	v1321 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1245<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[39]))))
	if v1321 == int32(-53) {
		v3777 = v1234
		v3781 = v1312
		v3783 = v1240
		v3784 = v1241
		v3786 = v1313
		v3787 = v1314
		v3788 = v1245
		v3791 = v1248
		v3793 = v1315
		v3794 = v1316
		goto L307
	} else {
		goto L333
	}
L312:
	;
	v1265 = int32(_a_F_BootstrapModeMain_29)
	if base.Ui32(v1265) <= base.Ui32(v1258) {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v1268 = v1265
	goto L315
L314:
	;
	v1268 = v1258
	goto L315
L315:
	;
	v1273 = F_palloc(m, v1268*int32(6)|int32(3))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	if v1273 == int32(0) {
		goto L302
	} else {
		goto L317
	}
L317:
	;
	v1278 = int32(1)
	v1281 = (v1244-v1243)>>(uint(v1278)%32) + v1278
	v1283 = v1281 << (uint(v1278) % 32)
	if v1283 != 0 {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	v1288 = v1285 + v1268<<(uint(int32(1))%32)
	v1290 = v1281 << (uint(int32(2)) % 32)
	if v1290 != 0 {
		goto L323
	} else {
		goto L324
	}
L319:
	;
	v1284 = F__emscripten_memcpy_bulkmem(m, v1273, v1243, v1283)
	mBase = m.M
	v1285 = v1284
	goto L321
L320:
	;
	v1285 = v1273
	goto L321
L321:
	;
	goto L318
L322:
	;
	if v1241+int32(928) != v1243 {
		goto L326
	} else {
		goto L327
	}
L323:
	;
	v1291 = F__emscripten_memcpy_bulkmem(m, v1288, v1250, v1290)
	mBase = m.M
	v1292 = v1291
	goto L325
L324:
	;
	v1292 = v1288
	goto L325
L325:
	;
	goto L322
L326:
	;
	F_pfree(m, v1243)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L1
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	v1298 = int32(1)
	v1300 = v1285 + v1281<<(uint(v1298)%32)
	if base.Ui32(v1285+v1268<<(uint(v1298)%32)) <= base.Ui32(v1300) {
		goto L330
	} else {
		goto L331
	}
L329:
	;
	goto L328
L330:
	;
	v3751 = v1241
	v3753 = v1285
	goto L308
L331:
	;
	goto L332
L332:
	;
	v1312 = v1292 + v1290 - int32(4)
	v1313 = v1285
	v1314 = v1300 - int32(2)
	v1315 = v1292
	v1316 = v1268
	goto L311
L333:
	;
	if v1240 == int32(-2) {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	v3720 = v1321 + v3719
	if base.Ui32(int32(169)) < base.Ui32(v3720) {
		v3777 = v3686
		v3781 = v3690
		v3783 = v3718
		v3784 = v3693
		v3786 = v3695
		v3787 = v3696
		v3788 = v3697
		v3791 = v3700
		v3793 = v3702
		v3794 = v3703
		goto L307
	} else {
		goto L716
	}
L335:
	;
	v1326 = m.G0
	v1328 = v1326 - int32(16)
	m.G0 = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+92)) = v1241 + int32(1340)
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+40))
	if v1333 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L336:
	;
	v3686 = v1234
	v3690 = v1312
	v3692 = v1240
	v3693 = v1241
	v3695 = v1313
	v3696 = v1314
	v3697 = v1245
	v3700 = v1248
	v3702 = v1315
	v3703 = v1316
	goto L337
L337:
	;
	if v3692 <= int32(0) {
		goto L712
	} else {
		goto L713
	}
L338:
	;
	v3686 = v1403
	v3690 = v1407
	v3692 = v2219
	v3693 = v1410
	v3695 = v1412
	v3696 = v1413
	v3697 = v1414
	v3700 = v1417
	v3702 = v1419
	v3703 = v1420
	goto L337
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+40)) = int32(1)
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+44))
	if v1338 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	goto L341
L341:
	;
	v1403 = v1234
	v1407 = v1312
	v1410 = v1241
	v1412 = v1313
	v1413 = v1314
	v1414 = v1245
	v1416 = v1328
	v1417 = v1248
	v1419 = v1315
	v1420 = v1316
	goto L358
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+44)) = int32(1)
	goto L344
L343:
	;
	goto L344
L344:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+4))
	if v1343 == int32(0) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+4)) = v1347
	goto L347
L346:
	;
	goto L347
L347:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+8))
	if v1349 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1353 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+8)) = v1353
	goto L350
L349:
	;
	goto L350
L350:
	;
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+20))
	if v1355 != 0 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1382)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+28)) = v1383
	v1387 = v1381 + v1380<<(uint(int32(2))%32)
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1388)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+80)) = v1389
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+36)) = v1389
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1392)))
	*(*int32)(unsafe.Add(mBase, uint32(v1234)+4)) = v1393
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1389))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1234)+24)) = uint8(v1395)
	goto L341
L352:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+12))
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1355+v1356<<(uint(int32(2))%32))))
	if v1360 != 0 {
		v1380 = v1356
		v1381 = v1355
		v1382 = v1360
		goto L351
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	F_boot_yyensure_buffer_stack(m, v1234)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L1
	} else {
		goto L356
	}
L355:
	;
	goto L354
L356:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+4))
	v1366 = F_boot_yy_create_buffer(m, v1365, v1234)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+20))
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+12))
	v1370 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1368+v1369<<(uint(v1370)%32)))) = v1366
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+20))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+12))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1374+v1375<<(uint(v1370)%32))))
	v1380 = v1375
	v1381 = v1374
	v1382 = v1379
	goto L351
L358:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+36))
	v1426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1425))) = uint8(v1426)
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1428+v1429<<(uint(int32(2))%32))))
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+28))
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+44))
	v1437 = v1425
	v1441 = v1434 + v1435
	v1443 = v1425
	goto L360
L360:
	;
	v1462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1443))))
	v1465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1462)+uint32(_c_F_BootstrapModeMain[41]))))
	v1467 = v1441 << (uint(int32(1)) % 32)
	v1470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1467)+uint32(_c_F_BootstrapModeMain[42]))))
	if v1470 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+68)) = v1443
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+64)) = v1441
	goto L364
L363:
	;
	goto L364
L364:
	;
	v1475 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1467)+uint32(_c_F_BootstrapModeMain[43]))))
	v1476 = v1465 + v1475
	v1481 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1476<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v1481 != v1441 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1485 = v1465
	v1487 = v1441
	v1488 = v1465
	goto L368
L366:
	;
	v1544 = v1476
	goto L367
L367:
	;
	v1560 = int32(1)
	v1566 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1544<<(uint(v1560)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v1566 != int32(127) {
		v1441 = v1566
		v1443 = v1443 + v1560
		goto L360
	} else {
		goto L374
	}
L368:
	;
	v1512 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1487<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[46]))))
	if int32(128) <= v1512 {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	v1544 = v1526
	goto L367
L370:
	;
	v1517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1485)+uint32(_c_F_BootstrapModeMain[47]))))
	v1518 = v1517
	goto L372
L371:
	;
	v1518 = v1488
	goto L372
L372:
	;
	v1520 = v1518 & int32(255)
	v1521 = int32(1)
	v1525 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1512<<(uint(v1521)%32))+uint32(_c_F_BootstrapModeMain[43]))))
	v1526 = v1520 + v1525
	v1531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1526<<(uint(v1521)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v1531 != v1512&int32(_a_F_BootstrapModeMain_63) {
		v1485 = v1520
		v1487 = v1512
		v1488 = v1518
		goto L368
	} else {
		goto L373
	}
L373:
	;
	goto L369
L374:
	;
	v1570 = v1437
	goto L375
L375:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+64))
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+68))
	v1597 = v1570
	v1600 = v1594
	v1602 = v1595
	goto L377
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+80)) = v1597
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+32)) = v1602 - v1597
	v1624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1602))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1403)+24)) = uint8(v1624)
	v1626 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1602))) = uint8(v1626)
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+36)) = v1602
	v1633 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1600<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[42]))))
	v1638 = v1633
	goto L379
L379:
	;
	switch v1638 {
	case 0:
		goto L425
	case 1:
		goto L424
	case 2:
		goto L423
	case 3:
		goto L422
	case 4:
		goto L421
	case 5:
		goto L420
	case 6:
		goto L419
	case 7:
		goto L418
	case 8:
		goto L417
	case 9:
		goto L416
	case 10:
		goto L415
	case 11:
		goto L414
	case 12:
		goto L413
	case 13:
		goto L412
	case 14:
		goto L411
	case 15:
		goto L410
	case 16:
		goto L409
	case 17:
		goto L408
	case 18:
		goto L407
	case 19:
		goto L406
	case 20:
		goto L405
	case 21:
		goto L404
	case 22:
		goto L403
	case 23:
		goto L402
	case 24:
		goto L401
	case 25:
		goto L400
	case 26:
		goto L399
	case 27:
		goto L398
	case 28:
		goto L397
	case 29:
		goto L396
	case 30:
		goto L393
	case 31:
		goto L392
	case 32:
		goto L391
	case 33:
		v2219 = int32(0)
		goto L394
	default:
		goto L390
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+36)) = v3648
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+48)) = int32(0)
	v3676 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+44))
	v3680 = base.I32_div_s(v3676-int32(1), int32(2))
	v1638 = v3680 + int32(33)
	goto L379
L382:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_64))
	mBase = m.M
	v3647 = m.ExcPending
	if v3647 != 0 {
		goto L1
	} else {
		goto L711
	}
L383:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_65))
	mBase = m.M
	v3642 = m.ExcPending
	if v3642 != 0 {
		goto L1
	} else {
		goto L710
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	v3493 = v3472 + v3476
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+36)) = v3493
	v3498 = *(*int32)(unsafe.Add(mBase, uint32(v3473+v3477<<(uint(int32(2))%32))))
	v3499 = *(*int32)(unsafe.Add(mBase, uint32(v3498)+28))
	v3500 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+44))
	v3501 = v3499 + v3500
	if base.Ui32(v3493) <= base.Ui32(v3468) {
		v1597 = v3468
		v1600 = v3501
		v1602 = v3493
		goto L377
	} else {
		goto L691
	}
L386:
	;
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v3114)))
	*(*int32)(unsafe.Add(mBase, uint32(v3115)+16)) = v3093
	v3118 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+28))
	if v3118 != 0 {
		v3240 = int32(0)
		goto L635
	} else {
		goto L636
	}
L387:
	;
	v3084 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3093 = v3063
	v3114 = v3084 + v3085<<(uint(int32(2))%32)
	goto L386
L388:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_66))
	mBase = m.M
	v3058 = m.ExcPending
	if v3058 != 0 {
		goto L1
	} else {
		goto L634
	}
L389:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_67))
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L1
	} else {
		goto L633
	}
L390:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_68))
	mBase = m.M
	v3052 = m.ExcPending
	if v3052 != 0 {
		goto L1
	} else {
		goto L632
	}
L391:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1602))) = uint8(v2280)
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2286 = v2282 + v2283<<(uint(int32(2))%32)
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v2286)))
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v2287)+44))
	if v2288 == int32(0) {
		goto L511
	} else {
		goto L512
	}
L392:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2261 != 0 {
		goto L507
	} else {
		goto L508
	}
L393:
	;
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2224 != 0 {
		goto L501
	} else {
		goto L502
	}
L394:
	;
	m.G0 = v1416 + int32(16)
	goto L338
L395:
	;
	v2219 = int32(258)
	goto L394
L396:
	;
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2196 != 0 {
		goto L497
	} else {
		goto L498
	}
L397:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2175 != 0 {
		goto L493
	} else {
		goto L494
	}
L398:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2155 != 0 {
		goto L490
	} else {
		goto L491
	}
L399:
	;
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2136 != 0 {
		goto L487
	} else {
		goto L488
	}
L400:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2117 != 0 {
		goto L484
	} else {
		goto L485
	}
L401:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2098 != 0 {
		goto L481
	} else {
		goto L482
	}
L402:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2079 != 0 {
		goto L478
	} else {
		goto L479
	}
L403:
	;
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2060 != 0 {
		goto L475
	} else {
		goto L476
	}
L404:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2041 != 0 {
		goto L472
	} else {
		goto L473
	}
L405:
	;
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2022 != 0 {
		goto L469
	} else {
		goto L470
	}
L406:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v2003 != 0 {
		goto L466
	} else {
		goto L467
	}
L407:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1984 != 0 {
		goto L463
	} else {
		goto L464
	}
L408:
	;
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1965 != 0 {
		goto L460
	} else {
		goto L461
	}
L409:
	;
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1948 == int32(0) {
		goto L358
	} else {
		goto L459
	}
L410:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1931 == int32(0) {
		goto L358
	} else {
		goto L458
	}
L411:
	;
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1906 != 0 {
		goto L455
	} else {
		goto L456
	}
L412:
	;
	v1886 = int32(262)
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1887 == int32(0) {
		v2219 = v1886
		goto L394
	} else {
		goto L454
	}
L413:
	;
	v1868 = int32(261)
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1869 == int32(0) {
		v2219 = v1868
		goto L394
	} else {
		goto L453
	}
L414:
	;
	v1850 = int32(260)
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1851 == int32(0) {
		v2219 = v1850
		goto L394
	} else {
		goto L452
	}
L415:
	;
	v1832 = int32(259)
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1833 == int32(0) {
		v2219 = v1832
		goto L394
	} else {
		goto L451
	}
L416:
	;
	v1814 = int32(263)
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1815 == int32(0) {
		v2219 = v1814
		goto L394
	} else {
		goto L450
	}
L417:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1795 != 0 {
		goto L447
	} else {
		goto L448
	}
L418:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1776 != 0 {
		goto L444
	} else {
		goto L445
	}
L419:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1757 != 0 {
		goto L441
	} else {
		goto L442
	}
L420:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1738 != 0 {
		goto L438
	} else {
		goto L439
	}
L421:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1719 != 0 {
		goto L435
	} else {
		goto L436
	}
L422:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1700 != 0 {
		goto L432
	} else {
		goto L433
	}
L423:
	;
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1681 != 0 {
		goto L429
	} else {
		goto L430
	}
L424:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+32))
	if v1662 != 0 {
		goto L426
	} else {
		goto L427
	}
L425:
	;
	v1660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1403)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1602))) = uint8(v1660)
	v1570 = v1597
	goto L375
L426:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1663+v1664<<(uint(int32(2))%32))))
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669+v1662-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1668)+28)) = base.B2i32(v1673 == int32(10))
	goto L428
L427:
	;
	goto L428
L428:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1677))) = int32(_a_F_BootstrapModeMain_69)
	v2219 = int32(264)
	goto L394
L429:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1682+v1683<<(uint(int32(2))%32))))
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1688+v1681-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1687)+28)) = base.B2i32(v1692 == int32(10))
	goto L431
L430:
	;
	goto L431
L431:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1696))) = int32(_a_F_BootstrapModeMain_70)
	v2219 = int32(265)
	goto L394
L432:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1701+v1702<<(uint(int32(2))%32))))
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1707+v1700-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1706)+28)) = base.B2i32(v1711 == int32(10))
	goto L434
L433:
	;
	goto L434
L434:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1715))) = int32(_a_F_BootstrapModeMain_71)
	v2219 = int32(266)
	goto L394
L435:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1720+v1721<<(uint(int32(2))%32))))
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726+v1719-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1725)+28)) = base.B2i32(v1730 == int32(10))
	goto L437
L436:
	;
	goto L437
L437:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1734))) = int32(_a_F_BootstrapModeMain_72)
	v2219 = int32(276)
	goto L394
L438:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v1739+v1740<<(uint(int32(2))%32))))
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1745+v1738-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1744)+28)) = base.B2i32(v1749 == int32(10))
	goto L440
L439:
	;
	goto L440
L440:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1753))) = int32(_a_F_BootstrapModeMain_73)
	v2219 = int32(277)
	goto L394
L441:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1758+v1759<<(uint(int32(2))%32))))
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1764+v1757-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1763)+28)) = base.B2i32(v1768 == int32(10))
	goto L443
L442:
	;
	goto L443
L443:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1772))) = int32(_a_F_BootstrapModeMain_74)
	v2219 = int32(278)
	goto L394
L444:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1777+v1778<<(uint(int32(2))%32))))
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1783+v1776-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+28)) = base.B2i32(v1787 == int32(10))
	goto L446
L445:
	;
	goto L446
L446:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1791))) = int32(_a_F_BootstrapModeMain_75)
	v2219 = int32(279)
	goto L394
L447:
	;
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1796+v1797<<(uint(int32(2))%32))))
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1802+v1795-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1801)+28)) = base.B2i32(v1806 == int32(10))
	goto L449
L448:
	;
	goto L449
L449:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1810))) = int32(_a_F_BootstrapModeMain_76)
	v2219 = int32(267)
	goto L394
L450:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1818+v1819<<(uint(int32(2))%32))))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1824+v1815-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1823)+28)) = base.B2i32(v1828 == int32(10))
	v2219 = v1814
	goto L394
L451:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1836+v1837<<(uint(int32(2))%32))))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1842+v1833-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1841)+28)) = base.B2i32(v1846 == int32(10))
	v2219 = v1832
	goto L394
L452:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1854+v1855<<(uint(int32(2))%32))))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1860+v1851-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1859)+28)) = base.B2i32(v1864 == int32(10))
	v2219 = v1850
	goto L394
L453:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1872+v1873<<(uint(int32(2))%32))))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1878+v1869-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1877)+28)) = base.B2i32(v1882 == int32(10))
	v2219 = v1868
	goto L394
L454:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1890+v1891<<(uint(int32(2))%32))))
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+v1887-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1895)+28)) = base.B2i32(v1900 == int32(10))
	v2219 = v1886
	goto L394
L455:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1905+v1904<<(uint(int32(2))%32))))
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1911+v1906-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1910)+28)) = base.B2i32(v1915 == int32(10))
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1921 = v1919
	v1922 = v1920
	goto L457
L456:
	;
	v1921 = v1905
	v1922 = v1904
	goto L457
L457:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1922<<(uint(int32(2))%32)+v1921)))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1926)+32)) = v1927 + int32(1)
	goto L358
L458:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1934+v1935<<(uint(int32(2))%32))))
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1940+v1931-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1939)+28)) = base.B2i32(v1944 == int32(10))
	goto L358
L459:
	;
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1951+v1952<<(uint(int32(2))%32))))
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1957+v1948-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1956)+28)) = base.B2i32(v1961 == int32(10))
	goto L358
L460:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(v1966+v1967<<(uint(int32(2))%32))))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1972+v1965-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1971)+28)) = base.B2i32(v1976 == int32(10))
	goto L462
L461:
	;
	goto L462
L462:
	;
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1980))) = int32(_a_F_BootstrapModeMain_77)
	v2219 = int32(268)
	goto L394
L463:
	;
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v1985+v1986<<(uint(int32(2))%32))))
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v1995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1991+v1984-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1990)+28)) = base.B2i32(v1995 == int32(10))
	goto L465
L464:
	;
	goto L465
L465:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1999))) = int32(_a_F_BootstrapModeMain_78)
	v2219 = int32(272)
	goto L394
L466:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v2004+v2005<<(uint(int32(2))%32))))
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2010+v2003-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+28)) = base.B2i32(v2014 == int32(10))
	goto L468
L467:
	;
	goto L468
L468:
	;
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2018))) = int32(_a_F_BootstrapModeMain_79)
	v2219 = int32(273)
	goto L394
L469:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2023+v2024<<(uint(int32(2))%32))))
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2029+v2022-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2028)+28)) = base.B2i32(v2033 == int32(10))
	goto L471
L470:
	;
	goto L471
L471:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2037))) = int32(_a_F_BootstrapModeMain_80)
	v2219 = int32(274)
	goto L394
L472:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v2042+v2043<<(uint(int32(2))%32))))
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2048+v2041-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2047)+28)) = base.B2i32(v2052 == int32(10))
	goto L474
L473:
	;
	goto L474
L474:
	;
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2056))) = int32(_a_F_BootstrapModeMain_81)
	v2219 = int32(269)
	goto L394
L475:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v2061+v2062<<(uint(int32(2))%32))))
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2067+v2060-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2066)+28)) = base.B2i32(v2071 == int32(10))
	goto L477
L476:
	;
	goto L477
L477:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2075))) = int32(_a_F_BootstrapModeMain_82)
	v2219 = int32(270)
	goto L394
L478:
	;
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2080+v2081<<(uint(int32(2))%32))))
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2086+v2079-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2085)+28)) = base.B2i32(v2090 == int32(10))
	goto L480
L479:
	;
	goto L480
L480:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2094))) = int32(_a_F_BootstrapModeMain_83)
	v2219 = int32(271)
	goto L394
L481:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v2099+v2100<<(uint(int32(2))%32))))
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2105+v2098-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2104)+28)) = base.B2i32(v2109 == int32(10))
	goto L483
L482:
	;
	goto L483
L483:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2113))) = int32(_a_F_BootstrapModeMain_84)
	v2219 = int32(275)
	goto L394
L484:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2118+v2119<<(uint(int32(2))%32))))
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2124+v2117-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2123)+28)) = base.B2i32(v2128 == int32(10))
	goto L486
L485:
	;
	goto L486
L486:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2132))) = int32(_a_F_BootstrapModeMain_85)
	v2219 = int32(280)
	goto L394
L487:
	;
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(v2137+v2138<<(uint(int32(2))%32))))
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2143+v2136-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2142)+28)) = base.B2i32(v2147 == int32(10))
	goto L489
L488:
	;
	goto L489
L489:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2151))) = int32(_a_F_BootstrapModeMain_86)
	v2219 = int32(281)
	goto L394
L490:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v2156+v2157<<(uint(int32(2))%32))))
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2162+v2155-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2161)+28)) = base.B2i32(v2166 == int32(10))
	goto L492
L491:
	;
	goto L492
L492:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2170))) = int32(_a_F_BootstrapModeMain_87)
	v2219 = int32(282)
	goto L394
L493:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v2176+v2177<<(uint(int32(2))%32))))
	v2185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2175+v2174-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2181)+28)) = base.B2i32(v2185 == int32(10))
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2190 = v2189
	goto L495
L494:
	;
	v2190 = v2174
	goto L495
L495:
	;
	v2191 = F_pstrdup(m, v2190)
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2193))) = v2191
	goto L395
L497:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v2197+v2198<<(uint(int32(2))%32))))
	v2206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196+v2195-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2202)+28)) = base.B2i32(v2206 == int32(10))
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2211 = v2210
	goto L499
L498:
	;
	v2211 = v2195
	goto L499
L499:
	;
	v2212 = F_DeescapeQuotedString(m, v2211)
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2214))) = v2212
	goto L395
L501:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2225+v2226<<(uint(int32(2))%32))))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2231+v2224-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2230)+28)) = base.B2i32(v2235 == int32(10))
	goto L503
L502:
	;
	goto L503
L503:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v2243+v2244<<(uint(int32(2))%32))))
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v2248)+32))
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1416)+4)) = v2250
	*(*int32)(unsafe.Add(mBase, uint32(v1416))) = v2249
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_88), v1416)
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_89), int32(124), int32(_a_F_BootstrapModeMain_90))
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L507:
	;
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2263 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v2262+v2263<<(uint(int32(2))%32))))
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2268+v2261-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+28)) = base.B2i32(v2272 == int32(10))
	goto L509
L508:
	;
	goto L509
L509:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v2278 = m.ExcPending
	if v2278 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L511:
	;
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v2287)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+28)) = v2291
	v2293 = *(*int32)(unsafe.Add(mBase, uint32(v2286)))
	v2294 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2293))) = v2294
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2298 = int32(2)
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v2296+v2297<<(uint(v2298)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2301)+44)) = int32(1)
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(v2304+v2305<<(uint(v2298)%32))))
	v2310 = v2309
	v2311 = v2304
	v2312 = v2305
	goto L513
L512:
	;
	v2310 = v2287
	v2311 = v2282
	v2312 = v2283
	goto L513
L513:
	;
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+36))
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v2310)+4))
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+28))
	v2316 = v2314 + v2315
	if base.Ui32(v2313) <= base.Ui32(v2316) {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2322 = v2318 + (v2279 ^ int32(-1)) + v1602
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+36)) = v2322
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2311+v2312<<(uint(int32(2))%32))))
	v2328 = *(*int32)(unsafe.Add(mBase, uint32(v2327)+28))
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+44))
	v2330 = v2328 + v2329
	if base.Ui32(v2318) < base.Ui32(v2322) {
		goto L517
	} else {
		goto L518
	}
L515:
	;
	goto L516
L516:
	;
	if base.Ui32(v2316+int32(1)) < base.Ui32(v2313) {
		goto L389
	} else {
		goto L549
	}
L517:
	;
	v2336 = v2330
	v2338 = v2318
	goto L520
L518:
	;
	v2473 = v2330
	goto L519
L519:
	;
	v2495 = v2473 << (uint(int32(1)) % 32)
	v2498 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2495)+uint32(_c_F_BootstrapModeMain[42]))))
	if v2498 != 0 {
		goto L538
	} else {
		goto L539
	}
L520:
	;
	v2358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2338))))
	if v2358 != 0 {
		goto L522
	} else {
		goto L523
	}
L521:
	;
	v2473 = v2465
	goto L519
L522:
	;
	v2361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2358)+uint32(_c_F_BootstrapModeMain[41]))))
	v2362 = v2361
	goto L524
L523:
	;
	v2362 = int32(1)
	goto L524
L524:
	;
	v2367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2336<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[42]))))
	if v2367 != 0 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+68)) = v2338
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+64)) = v2336
	goto L527
L526:
	;
	goto L527
L527:
	;
	v2371 = v2362 & int32(255)
	v2372 = int32(1)
	v2376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2336<<(uint(v2372)%32))+uint32(_c_F_BootstrapModeMain[43]))))
	v2377 = v2371 + v2376
	v2382 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2377<<(uint(v2372)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v2382 != v2336 {
		goto L528
	} else {
		goto L529
	}
L528:
	;
	v2386 = v2362
	v2388 = v2336
	v2389 = v2371
	goto L531
L529:
	;
	v2445 = v2377
	goto L530
L530:
	;
	v2461 = int32(1)
	v2465 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2445<<(uint(v2461)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	v2467 = v2338 + v2461
	if v2467 != v2322 {
		v2336 = v2465
		v2338 = v2467
		goto L520
	} else {
		goto L537
	}
L531:
	;
	v2413 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2388<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[46]))))
	if int32(128) <= v2413 {
		goto L533
	} else {
		goto L534
	}
L532:
	;
	v2445 = v2427
	goto L530
L533:
	;
	v2418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2389)+uint32(_c_F_BootstrapModeMain[47]))))
	v2419 = v2418
	goto L535
L534:
	;
	v2419 = v2386
	goto L535
L535:
	;
	v2421 = v2419 & int32(255)
	v2422 = int32(1)
	v2426 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2413<<(uint(v2422)%32))+uint32(_c_F_BootstrapModeMain[43]))))
	v2427 = v2421 + v2426
	v2432 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2427<<(uint(v2422)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v2432 != v2413&int32(_a_F_BootstrapModeMain_63) {
		v2386 = v2419
		v2388 = v2413
		v2389 = v2421
		goto L531
	} else {
		goto L536
	}
L536:
	;
	goto L532
L537:
	;
	goto L521
L538:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+68)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+64)) = v2473
	goto L540
L539:
	;
	goto L540
L540:
	;
	v2503 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2495)+uint32(_c_F_BootstrapModeMain[43]))))
	v2504 = int32(1)
	v2505 = v2503 + v2504
	v2510 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2505<<(uint(v2504)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v2510 != v2473 {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v2516 = v2473
	goto L544
L542:
	;
	v2558 = v2505
	goto L543
L543:
	;
	v2585 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2558<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[45]))))
	if v2585 == int32(127) {
		v1570 = v2318
		goto L375
	} else {
		goto L547
	}
L544:
	;
	v2537 = int32(1)
	v2541 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2516<<(uint(v2537)%32))+uint32(_c_F_BootstrapModeMain[46]))))
	v2542 = base.I32_extend16_s(v2541)
	v2547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2542<<(uint(v2537)%32))+uint32(_c_F_BootstrapModeMain[43]))))
	v2549 = v2547 + v2537
	v2554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2549<<(uint(v2537)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v2541 != v2554 {
		v2516 = v2542
		goto L544
	} else {
		goto L546
	}
L545:
	;
	v2558 = v2549
	goto L543
L546:
	;
	goto L545
L547:
	;
	if v2558&int32(2147483647) == int32(0) {
		v1570 = v2318
		goto L375
	} else {
		goto L548
	}
L548:
	;
	v2593 = v2322 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+36)) = v2593
	v1437 = v2318
	v1441 = v2585
	v1443 = v2593
	goto L360
L549:
	;
	v2598 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+80))
	v2599 = *(*int32)(unsafe.Add(mBase, uint32(v2310)+40))
	if v2599 == int32(0) {
		goto L550
	} else {
		goto L551
	}
L550:
	;
	if v2313-v2598 != int32(1) {
		v3468 = v2598
		v3472 = v2314
		v3473 = v2311
		v3476 = v2315
		v3477 = v2312
		goto L385
	} else {
		goto L553
	}
L551:
	;
	goto L552
L552:
	;
	v2607 = v2598 ^ int32(-1) + v2313
	if v2607 != 0 {
		goto L554
	} else {
		goto L555
	}
L553:
	;
	v3648 = v2598
	goto L381
L554:
	;
	v2608 = int32(7)
	v2609 = v2607 & v2608
	if base.Ui32(v2313-v2598-int32(2)) < base.Ui32(v2608) {
		goto L558
	} else {
		goto L559
	}
L555:
	;
	v2759 = v2310
	v2762 = v2311
	v2766 = v2312
	goto L556
L556:
	;
	v2782 = *(*int32)(unsafe.Add(mBase, uint32(v2759)+44))
	if v2782 == int32(2) {
		goto L570
	} else {
		goto L571
	}
L557:
	;
	if v2609 != 0 {
		goto L564
	} else {
		goto L565
	}
L558:
	;
	v2668 = v2598
	v2670 = v2314
	goto L557
L559:
	;
	goto L560
L560:
	;
	v2620 = v2598
	v2622 = v2314
	v2623 = int32(0)
	goto L561
L561:
	;
	v2643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2622))) = uint8(v2643)
	v2645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2622)+1)) = uint8(v2645)
	v2647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2622)+2)) = uint8(v2647)
	v2649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2622)+3)) = uint8(v2649)
	v2651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2622)+4)) = uint8(v2651)
	v2653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2622)+5)) = uint8(v2653)
	v2655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2622)+6)) = uint8(v2655)
	v2657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2622)+7)) = uint8(v2657)
	v2659 = int32(8)
	v2660 = v2622 + v2659
	v2662 = v2620 + v2659
	v2664 = v2623 + v2659
	if v2664 != v2607&int32(-8) {
		v2620 = v2662
		v2622 = v2660
		v2623 = v2664
		goto L561
	} else {
		goto L563
	}
L562:
	;
	v2668 = v2662
	v2670 = v2660
	goto L557
L563:
	;
	goto L562
L564:
	;
	v2694 = v2668
	v2696 = v2670
	v2697 = int32(0)
	goto L567
L565:
	;
	goto L566
L566:
	;
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v2751+v2752<<(uint(int32(2))%32))))
	v2759 = v2756
	v2762 = v2751
	v2766 = v2752
	goto L556
L567:
	;
	v2717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2694))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2696))) = uint8(v2717)
	v2719 = int32(1)
	v2724 = v2697 + v2719
	if v2724 != v2609 {
		v2694 = v2694 + v2719
		v2696 = v2696 + v2719
		v2697 = v2724
		goto L567
	} else {
		goto L569
	}
L568:
	;
	goto L566
L569:
	;
	goto L568
L570:
	;
	v2785 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+28)) = v2785
	v3093 = v2785
	v3114 = v2762 + v2766<<(uint(int32(2))%32)
	goto L386
L571:
	;
	goto L572
L572:
	;
	v2791 = *(*int32)(unsafe.Add(mBase, uint32(v2759)+12))
	v2792 = v2598 - v2313
	v2793 = v2791 + v2792
	if v2793 == int32(0) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v2796 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+36))
	v2799 = v2759
	v2801 = v2791
	v2806 = v2796
	goto L576
L574:
	;
	v2860 = v2759
	v2863 = v2793
	goto L575
L575:
	;
	v2883 = int32(_a_F_BootstrapModeMain_27)
	if base.Ui32(v2883) <= base.Ui32(v2863) {
		goto L592
	} else {
		goto L593
	}
L576:
	;
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v2799)+20))
	if v2822 == int32(0) {
		goto L578
	} else {
		goto L579
	}
L577:
	;
	v2860 = v2853
	v2863 = v2855
	goto L575
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2799)+4)) = int32(0)
	goto L382
L579:
	;
	goto L580
L580:
	;
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(v2799)+4))
	if base.Ui32(int32(2147483646)) <= base.Ui32(v2801) {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v2833 = int32(-3)
	goto L583
L582:
	;
	v2833 = v2801 << (uint(int32(1)) % 32)
	goto L583
L583:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2799)+12)) = v2833
	v2836 = v2833 + int32(2)
	if v2827 != 0 {
		goto L585
	} else {
		goto L586
	}
L584:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2799)+4)) = v2841
	if v2841 == int32(0) {
		goto L382
	} else {
		goto L590
	}
L585:
	;
	v2837 = F_repalloc(m, v2827, v2836)
	mBase = m.M
	v2838 = m.ExcPending
	if v2838 != 0 {
		goto L1
	} else {
		goto L588
	}
L586:
	;
	goto L587
L587:
	;
	v2839 = F_palloc(m, v2836)
	mBase = m.M
	v2840 = m.ExcPending
	if v2840 != 0 {
		goto L1
	} else {
		goto L589
	}
L588:
	;
	v2841 = v2837
	goto L584
L589:
	;
	v2841 = v2839
	goto L584
L590:
	;
	v2846 = v2841 + (v2806 - v2827)
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+36)) = v2846
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2849 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2853 = *(*int32)(unsafe.Add(mBase, uint32(v2848+v2849<<(uint(int32(2))%32))))
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+12))
	v2855 = v2854 + v2792
	if v2855 == int32(0) {
		v2799 = v2853
		v2801 = v2854
		v2806 = v2846
		goto L576
	} else {
		goto L591
	}
L591:
	;
	goto L577
L592:
	;
	v2886 = v2883
	goto L594
L593:
	;
	v2886 = v2863
	goto L594
L594:
	;
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(v2860)+24))
	if v2888 != 0 {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v2893 = int32(0)
	goto L599
L596:
	;
	goto L597
L597:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = int32(0)
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2968 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2972 = *(*int32)(unsafe.Add(mBase, uint32(v2967+v2968<<(uint(int32(2))%32))))
	v2973 = *(*int32)(unsafe.Add(mBase, uint32(v2972)+4))
	v2976 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	v2977 = F_fread(m, v2973+v2607, int32(1), v2886, v2976)
	mBase = m.M
	v2978 = m.ExcPending
	if v2978 != 0 {
		goto L1
	} else {
		goto L614
	}
L598:
	;
	switch v2918 {
	case 0:
		goto L606
	default:
		v2962 = v2932
		goto L604
	case 11:
		goto L605
	}
L599:
	;
	v2914 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	v2915 = F_do_getc(m, v2914)
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L1
	} else {
		goto L602
	}
L600:
	;
	v2932 = v2886
	goto L598
L601:
	;
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2919+v2920<<(uint(int32(2))%32))))
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2924)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2925+v2607+v2893))) = uint8(v2915)
	v2930 = v2893 + int32(1)
	if v2930 != v2886 {
		v2893 = v2930
		goto L599
	} else {
		goto L603
	}
L602:
	;
	v2918 = v2915 + int32(1)
	switch v2918 {
	case 0, 11:
		v2932 = v2893
		goto L598
	default:
		goto L601
	}
L603:
	;
	goto L600
L604:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+28)) = v2962
	v3063 = v2962
	goto L387
L605:
	;
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v2954 = *(*int32)(unsafe.Add(mBase, uint32(v2949+v2950<<(uint(int32(2))%32))))
	v2955 = *(*int32)(unsafe.Add(mBase, uint32(v2954)+4))
	v2958 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v2955+v2607+v2932))) = uint8(v2958)
	v2962 = v2932 + int32(1)
	goto L604
L606:
	;
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	v2934 = *(*int32)(unsafe.Add(mBase, uint32(v2933)+76))
	if v2934 < int32(0) {
		goto L609
	} else {
		goto L610
	}
L607:
	;
	if int32(base.Ui32(v2939)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v2962 = v2932
		goto L604
	} else {
		goto L612
	}
L608:
	;
	goto L607
L609:
	;
	v2937 = *(*int32)(unsafe.Add(mBase, uint32(v2933)))
	v2939 = v2937
	goto L608
L610:
	;
	goto L611
L611:
	;
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v2933)))
	v2939 = v2938
	goto L608
L612:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_66))
	mBase = m.M
	v2948 = m.ExcPending
	if v2948 != 0 {
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
	v2983 = v2977
	goto L615
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+28)) = v2983
	if v2983 != 0 {
		v3063 = v2983
		goto L387
	} else {
		goto L617
	}
L617:
	;
	v3005 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	v3006 = *(*int32)(unsafe.Add(mBase, uint32(v3005)+76))
	if v3006 < int32(0) {
		goto L620
	} else {
		goto L621
	}
L618:
	;
	if int32(base.Ui32(v3011)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L623
	} else {
		goto L624
	}
L619:
	;
	goto L618
L620:
	;
	v3009 = *(*int32)(unsafe.Add(mBase, uint32(v3005)))
	v3011 = v3009
	goto L619
L621:
	;
	goto L622
L622:
	;
	v3010 = *(*int32)(unsafe.Add(mBase, uint32(v3005)))
	v3011 = v3010
	goto L619
L623:
	;
	v3063 = int32(0)
	goto L387
L624:
	;
	goto L625
L625:
	;
	v3020 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	if v3020 != int32(27) {
		goto L388
	} else {
		goto L626
	}
L626:
	;
	v3024 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = v3024
	v3026 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	v3027 = *(*int32)(unsafe.Add(mBase, uint32(v3026)+76))
	if v3024 <= v3027 {
		goto L628
	} else {
		goto L629
	}
L627:
	;
	v3038 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3039 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3043 = *(*int32)(unsafe.Add(mBase, uint32(v3038+v3039<<(uint(int32(2))%32))))
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(v3043)+4))
	v3047 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	v3048 = F_fread(m, v3044+v2607, int32(1), v2886, v3047)
	mBase = m.M
	v3049 = m.ExcPending
	if v3049 != 0 {
		goto L1
	} else {
		goto L631
	}
L628:
	;
	v3030 = *(*int32)(unsafe.Add(mBase, uint32(v3026)))
	*(*int32)(unsafe.Add(mBase, uint32(v3026))) = v3030 & int32(-49)
	goto L627
L629:
	;
	goto L630
L630:
	;
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(v3026)))
	*(*int32)(unsafe.Add(mBase, uint32(v3026))) = v3034 & int32(-49)
	goto L627
L631:
	;
	v2983 = v3048
	goto L615
L632:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L633:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L634:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L635:
	;
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+28))
	v3242 = v3241 + v2607
	v3243 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3244 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3248 = *(*int32)(unsafe.Add(mBase, uint32(v3243+v3244<<(uint(int32(2))%32))))
	v3249 = *(*int32)(unsafe.Add(mBase, uint32(v3248)+12))
	if base.Ui32(v3249) < base.Ui32(v3242) {
		goto L659
	} else {
		goto L660
	}
L636:
	;
	if v2607 == int32(0) {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v3121 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	if v3122 != 0 {
		goto L642
	} else {
		goto L643
	}
L638:
	;
	goto L639
L639:
	;
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3227 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3228 = int32(2)
	v3231 = *(*int32)(unsafe.Add(mBase, uint32(v3226+v3227<<(uint(v3228)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3231)+44)) = v3228
	v3240 = v3228
	goto L635
L640:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3188)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3188))) = v3121
	v3195 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	if v3195 != 0 {
		goto L655
	} else {
		goto L656
	}
L641:
	;
	v3145 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	v3146 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v3143+v3146<<(uint(int32(2))%32))))
	if v3150 == int32(0) {
		goto L649
	} else {
		goto L650
	}
L642:
	;
	v3123 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3127 = *(*int32)(unsafe.Add(mBase, uint32(v3122+v3123<<(uint(int32(2))%32))))
	if v3127 != 0 {
		v3143 = v3122
		goto L641
	} else {
		goto L645
	}
L643:
	;
	goto L644
L644:
	;
	F_boot_yyensure_buffer_stack(m, v1403)
	mBase = m.M
	v3129 = m.ExcPending
	if v3129 != 0 {
		goto L1
	} else {
		goto L646
	}
L645:
	;
	goto L644
L646:
	;
	v3130 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+4))
	v3131 = F_boot_yy_create_buffer(m, v3130, v1403)
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	v3133 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3133+v3134<<(uint(int32(2))%32)))) = v3131
	v3139 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	if v3139 != 0 {
		v3143 = v3139
		goto L641
	} else {
		goto L648
	}
L648:
	;
	v3141 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	v3188 = int32(0)
	v3189 = v3141
	goto L640
L649:
	;
	v3188 = int32(0)
	v3189 = v3145
	goto L640
L650:
	;
	goto L651
L651:
	;
	v3154 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3150)+16)) = v3154
	v3156 = *(*int32)(unsafe.Add(mBase, uint32(v3150)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v3156))) = uint8(v3154)
	v3159 = *(*int32)(unsafe.Add(mBase, uint32(v3150)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v3159)+1)) = uint8(v3154)
	*(*int32)(unsafe.Add(mBase, uint32(v3150)+44)) = v3154
	*(*int32)(unsafe.Add(mBase, uint32(v3150)+28)) = int32(1)
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(v3150)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3150)+8)) = v3166
	v3168 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	if v3168 == v3154 {
		v3188 = v3150
		v3189 = v3145
		goto L640
	} else {
		goto L652
	}
L652:
	;
	v3171 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3174 = v3168 + v3171<<(uint(int32(2))%32)
	v3175 = *(*int32)(unsafe.Add(mBase, uint32(v3174)))
	if v3150 != v3175 {
		v3188 = v3150
		v3189 = v3145
		goto L640
	} else {
		goto L653
	}
L653:
	;
	v3177 = *(*int32)(unsafe.Add(mBase, uint32(v3175)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+28)) = v3177
	v3179 = *(*int32)(unsafe.Add(mBase, uint32(v3174)))
	v3180 = *(*int32)(unsafe.Add(mBase, uint32(v3179)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+80)) = v3180
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+36)) = v3180
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v3174)))
	v3184 = *(*int32)(unsafe.Add(mBase, uint32(v3183)))
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+4)) = v3184
	v3186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3180))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1403)+24)) = uint8(v3186)
	v3188 = v3150
	v3189 = v3145
	goto L640
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3188)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18])) = v3189
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3212 = v3208 + v3209<<(uint(int32(2))%32)
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v3212)))
	v3214 = *(*int32)(unsafe.Add(mBase, uint32(v3213)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+28)) = v3214
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(v3212)))
	v3217 = *(*int32)(unsafe.Add(mBase, uint32(v3216)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+36)) = v3217
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+80)) = v3217
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v3212)))
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(v3220)))
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+4)) = v3221
	v3223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3217))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1403)+24)) = uint8(v3223)
	v3240 = int32(1)
	goto L635
L655:
	;
	v3196 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3200 = *(*int32)(unsafe.Add(mBase, uint32(v3195+v3196<<(uint(int32(2))%32))))
	if v3188 == v3200 {
		goto L654
	} else {
		goto L658
	}
L656:
	;
	goto L657
L657:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3188)+32)) = int64(1)
	goto L654
L658:
	;
	goto L657
L659:
	;
	v3253 = v3242 + int32(base.Ui32(v3241)>>(uint(int32(1))%32))
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v3248)+4))
	if v3254 != 0 {
		goto L663
	} else {
		goto L664
	}
L660:
	;
	v3283 = v3243
	v3284 = v3242
	v3285 = v3244
	goto L661
L661:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+28)) = v3284
	v3287 = int32(2)
	v3290 = *(*int32)(unsafe.Add(mBase, uint32(v3283+v3285<<(uint(v3287)%32))))
	v3291 = *(*int32)(unsafe.Add(mBase, uint32(v3290)+4))
	v3293 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3291+v3284))) = uint8(v3293)
	v3295 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3300 = *(*int32)(unsafe.Add(mBase, uint32(v3295+v3296<<(uint(v3287)%32))))
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v3300)+4))
	v3302 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v3301+v3302)+1)) = uint8(v3293)
	v3306 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3307 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3310 = v3306 + v3307<<(uint(v3287)%32)
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(v3310)))
	v3312 = *(*int32)(unsafe.Add(mBase, uint32(v3311)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+80)) = v3312
	if v3240 == int32(1) {
		v3648 = v3312
		goto L381
	} else {
		goto L669
	}
L662:
	;
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3261 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3262 = int32(2)
	v3265 = *(*int32)(unsafe.Add(mBase, uint32(v3260+v3261<<(uint(v3262)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3265)+4)) = v3259
	v3267 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3268 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3272 = *(*int32)(unsafe.Add(mBase, uint32(v3267+v3268<<(uint(v3262)%32))))
	v3273 = *(*int32)(unsafe.Add(mBase, uint32(v3272)+4))
	if v3273 == int32(0) {
		goto L383
	} else {
		goto L668
	}
L663:
	;
	v3255 = F_repalloc(m, v3254, v3253)
	mBase = m.M
	v3256 = m.ExcPending
	if v3256 != 0 {
		goto L1
	} else {
		goto L666
	}
L664:
	;
	goto L665
L665:
	;
	v3257 = F_palloc(m, v3253)
	mBase = m.M
	v3258 = m.ExcPending
	if v3258 != 0 {
		goto L1
	} else {
		goto L667
	}
L666:
	;
	v3259 = v3255
	goto L662
L667:
	;
	v3259 = v3257
	goto L662
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3272)+12)) = v3253 - int32(2)
	v3279 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+12))
	v3280 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+28))
	v3282 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+20))
	v3283 = v3282
	v3284 = v3280 + v2607
	v3285 = v3279
	goto L661
L669:
	;
	switch v3240 - int32(1) {
	case 0:
		goto L384
	case 1:
		goto L670
	default:
		goto L671
	}
L670:
	;
	v3465 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+28))
	v3466 = *(*int32)(unsafe.Add(mBase, uint32(v3310)))
	v3467 = *(*int32)(unsafe.Add(mBase, uint32(v3466)+4))
	v3468 = v3312
	v3472 = v3467
	v3473 = v3306
	v3476 = v3465
	v3477 = v3307
	goto L385
L671:
	;
	v3321 = v3312 + (v2279 ^ int32(-1)) + v1602
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+36)) = v3321
	v3323 = *(*int32)(unsafe.Add(mBase, uint32(v3310)))
	v3324 = *(*int32)(unsafe.Add(mBase, uint32(v3323)+28))
	v3325 = *(*int32)(unsafe.Add(mBase, uint32(v1403)+44))
	v3326 = v3324 + v3325
	if base.Ui32(v3321) <= base.Ui32(v3312) {
		v1437 = v3312
		v1441 = v3326
		v1443 = v3321
		goto L360
	} else {
		goto L672
	}
L672:
	;
	v3332 = v3326
	v3336 = v3312
	goto L673
L673:
	;
	v3354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3336))))
	if v3354 != 0 {
		goto L675
	} else {
		goto L676
	}
L674:
	;
	v1437 = v3312
	v1441 = v3461
	v1443 = v3321
	goto L360
L675:
	;
	v3357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3354)+uint32(_c_F_BootstrapModeMain[41]))))
	v3358 = v3357
	goto L677
L676:
	;
	v3358 = int32(1)
	goto L677
L677:
	;
	v3363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3332<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[42]))))
	if v3363 != 0 {
		goto L678
	} else {
		goto L679
	}
L678:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+68)) = v3336
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+64)) = v3332
	goto L680
L679:
	;
	goto L680
L680:
	;
	v3367 = v3358 & int32(255)
	v3368 = int32(1)
	v3372 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3332<<(uint(v3368)%32))+uint32(_c_F_BootstrapModeMain[43]))))
	v3373 = v3367 + v3372
	v3378 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3373<<(uint(v3368)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v3378 != v3332 {
		goto L681
	} else {
		goto L682
	}
L681:
	;
	v3382 = v3358
	v3384 = v3332
	v3385 = v3367
	goto L684
L682:
	;
	v3441 = v3373
	goto L683
L683:
	;
	v3457 = int32(1)
	v3461 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3441<<(uint(v3457)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	v3463 = v3336 + v3457
	if v3321 != v3463 {
		v3332 = v3461
		v3336 = v3463
		goto L673
	} else {
		goto L690
	}
L684:
	;
	v3409 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3384<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[46]))))
	if int32(128) <= v3409 {
		goto L686
	} else {
		goto L687
	}
L685:
	;
	v3441 = v3423
	goto L683
L686:
	;
	v3414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3385)+uint32(_c_F_BootstrapModeMain[47]))))
	v3415 = v3414
	goto L688
L687:
	;
	v3415 = v3382
	goto L688
L688:
	;
	v3417 = v3415 & int32(255)
	v3418 = int32(1)
	v3422 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3409<<(uint(v3418)%32))+uint32(_c_F_BootstrapModeMain[43]))))
	v3423 = v3417 + v3422
	v3428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3423<<(uint(v3418)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v3428 != v3409&int32(_a_F_BootstrapModeMain_63) {
		v3382 = v3415
		v3384 = v3409
		v3385 = v3417
		goto L684
	} else {
		goto L689
	}
L689:
	;
	goto L685
L690:
	;
	goto L674
L691:
	;
	v3507 = v3501
	v3511 = v3468
	goto L692
L692:
	;
	v3529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3511))))
	if v3529 != 0 {
		goto L694
	} else {
		goto L695
	}
L693:
	;
	v1597 = v3468
	v1600 = v3636
	v1602 = v3493
	goto L377
L694:
	;
	v3532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3529)+uint32(_c_F_BootstrapModeMain[41]))))
	v3533 = v3532
	goto L696
L695:
	;
	v3533 = int32(1)
	goto L696
L696:
	;
	v3538 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3507<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[42]))))
	if v3538 != 0 {
		goto L697
	} else {
		goto L698
	}
L697:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+68)) = v3511
	*(*int32)(unsafe.Add(mBase, uint32(v1403)+64)) = v3507
	goto L699
L698:
	;
	goto L699
L699:
	;
	v3542 = v3533 & int32(255)
	v3543 = int32(1)
	v3547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3507<<(uint(v3543)%32))+uint32(_c_F_BootstrapModeMain[43]))))
	v3548 = v3542 + v3547
	v3553 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3548<<(uint(v3543)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v3553 != v3507 {
		goto L700
	} else {
		goto L701
	}
L700:
	;
	v3557 = v3533
	v3559 = v3507
	v3560 = v3542
	goto L703
L701:
	;
	v3616 = v3548
	goto L702
L702:
	;
	v3632 = int32(1)
	v3636 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3616<<(uint(v3632)%32))+uint32(_c_F_BootstrapModeMain[45]))))
	v3638 = v3511 + v3632
	if v3638 != v3493 {
		v3507 = v3636
		v3511 = v3638
		goto L692
	} else {
		goto L709
	}
L703:
	;
	v3584 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3559<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[46]))))
	if int32(128) <= v3584 {
		goto L705
	} else {
		goto L706
	}
L704:
	;
	v3616 = v3598
	goto L702
L705:
	;
	v3589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3560)+uint32(_c_F_BootstrapModeMain[47]))))
	v3590 = v3589
	goto L707
L706:
	;
	v3590 = v3557
	goto L707
L707:
	;
	v3592 = v3590 & int32(255)
	v3593 = int32(1)
	v3597 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3584<<(uint(v3593)%32))+uint32(_c_F_BootstrapModeMain[43]))))
	v3598 = v3592 + v3597
	v3603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3598<<(uint(v3593)%32))+uint32(_c_F_BootstrapModeMain[44]))))
	if v3603 != v3584&int32(_a_F_BootstrapModeMain_63) {
		v3557 = v3590
		v3559 = v3584
		v3560 = v3592
		goto L703
	} else {
		goto L708
	}
L708:
	;
	goto L704
L709:
	;
	goto L693
L710:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L711:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L712:
	;
	v3710 = int32(0)
	v3718 = v3710
	v3719 = v3710
	goto L334
L713:
	;
	goto L714
L714:
	;
	if base.Ui32(int32(282)) < base.Ui32(v3692) {
		v3718 = v3692
		v3719 = int32(2)
		goto L334
	} else {
		goto L715
	}
L715:
	;
	v3717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3692)+uint32(_c_F_BootstrapModeMain[48]))))
	v3718 = v3692
	v3719 = v3717
	goto L334
L716:
	;
	v3725 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3720)+uint32(_c_F_BootstrapModeMain[49]))))
	if v3719 != v3725 {
		v3777 = v3686
		v3781 = v3690
		v3783 = v3718
		v3784 = v3693
		v3786 = v3695
		v3787 = v3696
		v3788 = v3697
		v3791 = v3700
		v3793 = v3702
		v3794 = v3703
		goto L307
	} else {
		goto L717
	}
L717:
	;
	v3729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3720)+uint32(_c_F_BootstrapModeMain[50]))))
	if v3729 == int32(0) {
		v6238 = v3686
		goto L305
	} else {
		goto L718
	}
L718:
	;
	if v3720 != int32(25) {
		goto L719
	} else {
		goto L720
	}
L719:
	;
	v3735 = v3690 + int32(4)
	v3736 = *(*int32)(unsafe.Add(mBase, uint32(v3693)+1340))
	*(*int32)(unsafe.Add(mBase, uint32(v3735))) = v3736
	if v3718 != 0 {
		goto L722
	} else {
		goto L723
	}
L720:
	;
	goto L721
L721:
	;
	v3751 = v3693
	v3753 = v3695
	goto L308
L722:
	;
	v3740 = int32(-2)
	goto L724
L723:
	;
	v3740 = int32(0)
	goto L724
L724:
	;
	v6211 = v3686
	v6215 = v3735
	v6217 = v3740
	v6218 = v3693
	v6220 = v3695
	v6221 = v3696
	v6222 = v3729
	v6225 = v3700
	v6227 = v3702
	v6228 = v3703
	goto L306
L725:
	;
	F_pfree(m, v3753)
	mBase = m.M
	v3770 = m.ExcPending
	if v3770 != 0 {
		goto L1
	} else {
		goto L728
	}
L726:
	;
	goto L727
L727:
	;
	m.G0 = v3751 + int32(1344)
	goto L298
L728:
	;
	goto L727
L729:
	;
	v3807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3801)+uint32(_c_F_BootstrapModeMain[51]))))
	v3812 = *(*int32)(unsafe.Add(mBase, uint32(v3781+(int32(1)-v3807)<<(uint(int32(2))%32))))
	switch v3801 - int32(14) {
	case 0:
		goto L773
	case 1:
		goto L772
	case 2:
		goto L771
	case 3:
		goto L770
	case 4:
		goto L769
	case 5:
		goto L768
	case 6:
		goto L767
	case 7:
		goto L766
	case 8:
		goto L765
	case 9:
		goto L764
	case 10:
		goto L763
	case 11:
		goto L762
	case 12:
		goto L761
	case 13:
		goto L760
	case 14, 16, 25:
		goto L759
	case 15, 17, 19:
		goto L758
	case 18:
		goto L757
	default:
		v6156 = v3812
		goto L730
	case 22:
		goto L756
	case 23:
		goto L755
	case 24:
		goto L754
	case 26:
		goto L753
	case 30:
		goto L752
	case 31:
		goto L751
	case 32:
		goto L750
	case 33:
		goto L749
	case 34:
		goto L748
	case 35:
		goto L747
	case 36:
		goto L746
	case 37:
		goto L745
	case 38:
		goto L744
	case 39:
		goto L743
	case 40:
		goto L742
	case 41:
		goto L741
	case 42:
		goto L740
	case 43:
		goto L739
	case 44:
		goto L738
	case 45:
		goto L737
	case 46:
		goto L736
	case 47:
		goto L735
	case 48:
		goto L734
	case 49:
		goto L733
	case 50:
		goto L732
	case 51:
		goto L731
	}
L730:
	;
	v6183 = v3781 - v3807<<(uint(int32(2))%32) + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6183))) = v6156
	v6187 = v3787 - v3807<<(uint(int32(1))%32)
	v6188 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6187))))
	v6191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3801)+uint32(_c_F_BootstrapModeMain[52]))))
	v6194 = int32(*(*int8)(unsafe.Add(mBase, uint32(v6191)+uint32(_c_F_BootstrapModeMain[53]))))
	v6195 = v6188 + v6194
	if base.Ui32(int32(169)) < base.Ui32(v6195) {
		goto L1288
	} else {
		goto L1289
	}
L731:
	;
	v6151 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6152 = F_pstrdup(m, v6151)
	mBase = m.M
	v6153 = m.ExcPending
	if v6153 != 0 {
		goto L1
	} else {
		goto L1287
	}
L732:
	;
	v6148 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6149 = F_pstrdup(m, v6148)
	mBase = m.M
	v6150 = m.ExcPending
	if v6150 != 0 {
		goto L1
	} else {
		goto L1286
	}
L733:
	;
	v6145 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6146 = F_pstrdup(m, v6145)
	mBase = m.M
	v6147 = m.ExcPending
	if v6147 != 0 {
		goto L1
	} else {
		goto L1285
	}
L734:
	;
	v6142 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6143 = F_pstrdup(m, v6142)
	mBase = m.M
	v6144 = m.ExcPending
	if v6144 != 0 {
		goto L1
	} else {
		goto L1284
	}
L735:
	;
	v6139 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6140 = F_pstrdup(m, v6139)
	mBase = m.M
	v6141 = m.ExcPending
	if v6141 != 0 {
		goto L1
	} else {
		goto L1283
	}
L736:
	;
	v6136 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6137 = F_pstrdup(m, v6136)
	mBase = m.M
	v6138 = m.ExcPending
	if v6138 != 0 {
		goto L1
	} else {
		goto L1282
	}
L737:
	;
	v6133 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6134 = F_pstrdup(m, v6133)
	mBase = m.M
	v6135 = m.ExcPending
	if v6135 != 0 {
		goto L1
	} else {
		goto L1281
	}
L738:
	;
	v6130 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6131 = F_pstrdup(m, v6130)
	mBase = m.M
	v6132 = m.ExcPending
	if v6132 != 0 {
		goto L1
	} else {
		goto L1280
	}
L739:
	;
	v6127 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6128 = F_pstrdup(m, v6127)
	mBase = m.M
	v6129 = m.ExcPending
	if v6129 != 0 {
		goto L1
	} else {
		goto L1279
	}
L740:
	;
	v6124 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6125 = F_pstrdup(m, v6124)
	mBase = m.M
	v6126 = m.ExcPending
	if v6126 != 0 {
		goto L1
	} else {
		goto L1278
	}
L741:
	;
	v6121 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6122 = F_pstrdup(m, v6121)
	mBase = m.M
	v6123 = m.ExcPending
	if v6123 != 0 {
		goto L1
	} else {
		goto L1277
	}
L742:
	;
	v6118 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6119 = F_pstrdup(m, v6118)
	mBase = m.M
	v6120 = m.ExcPending
	if v6120 != 0 {
		goto L1
	} else {
		goto L1276
	}
L743:
	;
	v6115 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6116 = F_pstrdup(m, v6115)
	mBase = m.M
	v6117 = m.ExcPending
	if v6117 != 0 {
		goto L1
	} else {
		goto L1275
	}
L744:
	;
	v6112 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6113 = F_pstrdup(m, v6112)
	mBase = m.M
	v6114 = m.ExcPending
	if v6114 != 0 {
		goto L1
	} else {
		goto L1274
	}
L745:
	;
	v6109 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6110 = F_pstrdup(m, v6109)
	mBase = m.M
	v6111 = m.ExcPending
	if v6111 != 0 {
		goto L1
	} else {
		goto L1273
	}
L746:
	;
	v6106 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6107 = F_pstrdup(m, v6106)
	mBase = m.M
	v6108 = m.ExcPending
	if v6108 != 0 {
		goto L1
	} else {
		goto L1272
	}
L747:
	;
	v6103 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6104 = F_pstrdup(m, v6103)
	mBase = m.M
	v6105 = m.ExcPending
	if v6105 != 0 {
		goto L1
	} else {
		goto L1271
	}
L748:
	;
	v6100 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6101 = F_pstrdup(m, v6100)
	mBase = m.M
	v6102 = m.ExcPending
	if v6102 != 0 {
		goto L1
	} else {
		goto L1270
	}
L749:
	;
	v6097 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6098 = F_pstrdup(m, v6097)
	mBase = m.M
	v6099 = m.ExcPending
	if v6099 != 0 {
		goto L1
	} else {
		goto L1269
	}
L750:
	;
	v6096 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6156 = v6096
	goto L730
L751:
	;
	v6018 = int32(_a_F_BootstrapModeMain_92)
	v6020 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54])) = v6020 + int32(1)
	v6024 = m.G0
	v6026 = v6024 - int32(32)
	m.G0 = v6026
	v6030 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v6031 = m.ExcPending
	if v6031 != 0 {
		goto L1
	} else {
		goto L1257
	}
L752:
	;
	v5934 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v5935 = int32(_a_F_BootstrapModeMain_92)
	v5937 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54])) = v5937 + int32(1)
	v5941 = m.G0
	v5943 = v5941 - int32(48)
	m.G0 = v5943
	v5947 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5948 = m.ExcPending
	if v5948 != 0 {
		goto L1
	} else {
		goto L1242
	}
L753:
	;
	v5928 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v5932 = F_strtox_2(m, v5928, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L1241
L754:
	;
	v6156 = int32(2)
	goto L730
L755:
	;
	v6156 = int32(3)
	goto L730
L756:
	;
	v5143 = int32(_a_F_BootstrapModeMain_93)
	v5145 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	v5147 = v5145 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55])) = v5147
	if int32(41) <= v5147 {
		goto L299
	} else {
		goto L1108
	}
L757:
	;
	v5142 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v6156 = v5142
	goto L730
L758:
	;
	v6156 = int32(0)
	goto L730
L759:
	;
	v6156 = int32(1)
	goto L730
L760:
	;
	v5115 = F_palloc0(m, int32(36))
	mBase = m.M
	v5116 = m.ExcPending
	if v5116 != 0 {
		goto L1
	} else {
		goto L1105
	}
L761:
	;
	v5106 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+96)) = v5106
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+108)) = v5106
	v5112 = F_list_make1_impl(m, int32(1), v3784+int32(96))
	mBase = m.M
	v5113 = m.ExcPending
	if v5113 != 0 {
		goto L1
	} else {
		goto L1104
	}
L762:
	;
	v5102 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(8))))
	v5103 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v5104 = F_lappend(m, v5102, v5103)
	mBase = m.M
	v5105 = m.ExcPending
	if v5105 != 0 {
		goto L1
	} else {
		goto L1103
	}
L763:
	;
	v4982 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4982 == int32(0) {
		goto L1080
	} else {
		goto L1081
	}
L764:
	;
	v4864 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4865 = m.ExcPending
	if v4865 != 0 {
		goto L1
	} else {
		goto L1047
	}
L765:
	;
	v4740 = F_palloc0(m, int32(72))
	mBase = m.M
	v4741 = m.ExcPending
	if v4741 != 0 {
		goto L1
	} else {
		goto L1025
	}
L766:
	;
	v4619 = F_palloc0(m, int32(72))
	mBase = m.M
	v4620 = m.ExcPending
	if v4620 != 0 {
		goto L1
	} else {
		goto L1003
	}
L767:
	;
	v4528 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54]))
	v4530 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	if v4528 != v4530 {
		goto L301
	} else {
		goto L972
	}
L768:
	;
	v4495 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4495 == int32(0) {
		goto L962
	} else {
		goto L963
	}
L769:
	;
	v4337 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4337 == int32(0) {
		goto L919
	} else {
		goto L920
	}
L770:
	;
	v4316 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4316
	v4319 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4319)
	mBase = m.M
	v4321 = m.ExcPending
	if v4321 != 0 {
		goto L1
	} else {
		goto L911
	}
L771:
	;
	v4263 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4263 == int32(0) {
		goto L897
	} else {
		goto L898
	}
L772:
	;
	v4221 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4221 == int32(0) {
		goto L884
	} else {
		goto L885
	}
L773:
	;
	v3816 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v3816 == int32(0) {
		goto L774
	} else {
		goto L775
	}
L774:
	;
	v3821 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v3826 = F_AllocSetContextCreateInternal(m, v3821, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_27), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v3827 = m.ExcPending
	if v3827 != 0 {
		goto L1
	} else {
		goto L777
	}
L775:
	;
	v3829 = v3816
	goto L776
L776:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v3829
	v3832 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v3833 = int32(0)
	v3834 = m.G0
	v3836 = v3834 - int32(48)
	m.G0 = v3836
	if v3832&int32(3) == v3833 {
		v3861 = v3832
		goto L780
	} else {
		goto L781
	}
L777:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v3826
	v3829 = v3826
	goto L776
L778:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v3894) {
		goto L795
	} else {
		goto L796
	}
L779:
	;
	v3894 = v3886 - v3832
	goto L778
L780:
	;
	v3865 = v3861
	goto L789
L781:
	;
	v3845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3832))))
	if v3845 == int32(0) {
		goto L782
	} else {
		goto L783
	}
L782:
	;
	v3894 = int32(0)
	goto L778
L783:
	;
	goto L784
L784:
	;
	v3850 = v3832
	goto L785
L785:
	;
	v3854 = v3850 + int32(1)
	if v3854&int32(3) == int32(0) {
		v3861 = v3854
		goto L780
	} else {
		goto L787
	}
L786:
	;
	v3886 = v3854
	goto L779
L787:
	;
	v3859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3854))))
	if v3859 != 0 {
		v3850 = v3854
		goto L785
	} else {
		goto L788
	}
L788:
	;
	goto L786
L789:
	;
	v3871 = *(*int32)(unsafe.Add(mBase, uint32(v3865)))
	v3874 = int32(-2139062144)
	if (int32(16843008)-v3871|v3871)&v3874 == v3874 {
		v3865 = v3865 + int32(4)
		goto L789
	} else {
		goto L791
	}
L790:
	;
	v3880 = v3865
	goto L792
L791:
	;
	goto L790
L792:
	;
	v3884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3880))))
	if v3884 != 0 {
		v3880 = v3880 + int32(1)
		goto L792
	} else {
		goto L794
	}
L793:
	;
	v3886 = v3880
	goto L779
L794:
	;
	goto L793
L795:
	;
	v3897 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3832)+63)) = uint8(v3897)
	goto L797
L796:
	;
	goto L797
L797:
	;
	v3900 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59]))
	if v3900 == int32(0) {
		goto L798
	} else {
		goto L799
	}
L798:
	;
	F_populate_typ_list(m)
	mBase = m.M
	v3904 = m.ExcPending
	if v3904 != 0 {
		goto L1
	} else {
		goto L801
	}
L799:
	;
	goto L800
L800:
	;
	v3906 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v3906 != 0 {
		goto L802
	} else {
		goto L803
	}
L801:
	;
	goto L800
L802:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v3909 = m.ExcPending
	if v3909 != 0 {
		goto L1
	} else {
		goto L805
	}
L803:
	;
	goto L804
L804:
	;
	v3912 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v3913 = m.ExcPending
	if v3913 != 0 {
		goto L1
	} else {
		goto L806
	}
L805:
	;
	goto L804
L806:
	;
	if v3912 != 0 {
		goto L807
	} else {
		goto L808
	}
L807:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3836)+36)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v3836)+32)) = v3832
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_96), v3836+int32(32))
	mBase = m.M
	v3921 = m.ExcPending
	if v3921 != 0 {
		goto L1
	} else {
		goto L810
	}
L808:
	;
	goto L809
L809:
	;
	v3930 = F_makeRangeVar(m, int32(0), v3832, int32(-1))
	mBase = m.M
	v3931 = m.ExcPending
	if v3931 != 0 {
		goto L1
	} else {
		goto L812
	}
L810:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(458), int32(_a_F_BootstrapModeMain_97))
	mBase = m.M
	v3926 = m.ExcPending
	if v3926 != 0 {
		goto L1
	} else {
		goto L811
	}
L811:
	;
	goto L809
L812:
	;
	v3933 = F_table_openrv(m, v3930, int32(0))
	mBase = m.M
	v3934 = m.ExcPending
	if v3934 != 0 {
		goto L1
	} else {
		goto L813
	}
L813:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60])) = v3933
	v3937 = *(*int32)(unsafe.Add(mBase, uint32(v3933)+48))
	v3938 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3937)+120)))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55])) = v3938
	if int32(0) < v3938 {
		goto L814
	} else {
		goto L815
	}
L814:
	;
	v3950 = v3833
	goto L817
L815:
	;
	goto L816
L816:
	;
	m.G0 = v3836 + int32(48)
	v4200 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4200
	v4203 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4203)
	mBase = m.M
	v4205 = m.ExcPending
	if v4205 != 0 {
		goto L1
	} else {
		goto L876
	}
L817:
	;
	v3970 = v3950 << (uint(int32(2)) % 32)
	v3973 = *(*int32)(unsafe.Add(mBase, uint32(v3970)+uint32(_c_F_BootstrapModeMain[61])))
	if v3973 == int32(0) {
		goto L819
	} else {
		goto L820
	}
L818:
	;
	goto L816
L819:
	;
	v3977 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[62]))
	v3979 = F_MemoryContextAllocZero(m, v3977, int32(100))
	mBase = m.M
	v3980 = m.ExcPending
	if v3980 != 0 {
		goto L1
	} else {
		goto L822
	}
L820:
	;
	v3982 = v3973
	goto L821
L821:
	;
	v3984 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	v3985 = *(*int32)(unsafe.Add(mBase, uint32(v3984)+52))
	v3986 = *(*int32)(unsafe.Add(mBase, uint32(v3985)))
	v3990 = int32(100)
	v3994 = v3985 + v3986<<(uint(int32(4))%32) + v3950*v3990 + int32(20)
	if v3982 == v3994 {
		goto L824
	} else {
		goto L825
	}
L822:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3970)+uint32(_c_F_BootstrapModeMain[61]))) = v3979
	v3982 = v3979
	goto L821
L823:
	;
	v4140 = *(*int32)(unsafe.Add(mBase, uint32(v3970)+uint32(_c_F_BootstrapModeMain[61])))
	v4143 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4144 = m.ExcPending
	if v4144 != 0 {
		goto L1
	} else {
		goto L869
	}
L824:
	;
	goto L823
L825:
	;
	v3999 = v3982 + v3990
	if base.Ui32(v3994-v3999) <= base.Ui32(int32(-200)) {
		goto L826
	} else {
		goto L827
	}
L826:
	;
	v4006 = F___memcpy(m, v3982, v3994, v3990)
	mBase = m.M
	goto L823
L827:
	;
	goto L828
L828:
	;
	v4009 = (v3982 ^ v3994) & int32(3)
	if base.Ui32(v3982) < base.Ui32(v3994) {
		goto L831
	} else {
		goto L832
	}
L829:
	;
	if v4111 == int32(0) {
		goto L824
	} else {
		goto L865
	}
L830:
	;
	if base.Ui32(v4089) <= base.Ui32(int32(3)) {
		v4110 = v4088
		v4111 = v4089
		v4112 = v4090
		goto L829
	} else {
		goto L861
	}
L831:
	;
	if v4009 != 0 {
		goto L834
	} else {
		goto L835
	}
L832:
	;
	goto L833
L833:
	;
	if v4009 != 0 {
		v4071 = v3990
		goto L844
	} else {
		goto L845
	}
L834:
	;
	v4110 = v3994
	v4111 = v3990
	v4112 = v3982
	goto L829
L835:
	;
	goto L836
L836:
	;
	if v3982&int32(3) == int32(0) {
		goto L837
	} else {
		goto L838
	}
L837:
	;
	v4088 = v3994
	v4089 = v3990
	v4090 = v3982
	goto L830
L838:
	;
	goto L839
L839:
	;
	v4016 = v3994
	v4017 = v3990
	v4018 = v3982
	goto L840
L840:
	;
	if v4017 == int32(0) {
		goto L824
	} else {
		goto L842
	}
L841:
	;
	v4088 = v4025
	v4089 = v4027
	v4090 = v4029
	goto L830
L842:
	;
	v4022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4016))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4018))) = uint8(v4022)
	v4024 = int32(1)
	v4025 = v4016 + v4024
	v4027 = v4017 - v4024
	v4029 = v4018 + v4024
	if v4029&int32(3) != 0 {
		v4016 = v4025
		v4017 = v4027
		v4018 = v4029
		goto L840
	} else {
		goto L843
	}
L843:
	;
	goto L841
L844:
	;
	if v4071 == int32(0) {
		goto L824
	} else {
		goto L857
	}
L845:
	;
	if v3999&int32(3) != 0 {
		goto L846
	} else {
		goto L847
	}
L846:
	;
	v4036 = v3990
	goto L849
L847:
	;
	v4051 = v3990
	goto L848
L848:
	;
	if base.Ui32(v4051) <= base.Ui32(int32(3)) {
		v4071 = v4051
		goto L844
	} else {
		goto L853
	}
L849:
	;
	if v4036 == int32(0) {
		goto L824
	} else {
		goto L851
	}
L850:
	;
	v4051 = v4042
	goto L848
L851:
	;
	v4042 = v4036 - int32(1)
	v4043 = v3982 + v4042
	v4045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3994+v4042))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4043))) = uint8(v4045)
	if v4043&int32(3) != 0 {
		v4036 = v4042
		goto L849
	} else {
		goto L852
	}
L852:
	;
	goto L850
L853:
	;
	v4058 = v4051
	goto L854
L854:
	;
	v4062 = v4058 - int32(4)
	v4065 = *(*int32)(unsafe.Add(mBase, uint32(v3994+v4062)))
	*(*int32)(unsafe.Add(mBase, uint32(v3982+v4062))) = v4065
	if base.Ui32(int32(3)) < base.Ui32(v4062) {
		v4058 = v4062
		goto L854
	} else {
		goto L856
	}
L855:
	;
	v4071 = v4062
	goto L844
L856:
	;
	goto L855
L857:
	;
	v4078 = v4071
	goto L858
L858:
	;
	v4082 = v4078 - int32(1)
	v4085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3994+v4082))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3982+v4082))) = uint8(v4085)
	if v4082 != 0 {
		v4078 = v4082
		goto L858
	} else {
		goto L860
	}
L859:
	;
	goto L824
L860:
	;
	goto L859
L861:
	;
	v4095 = v4088
	v4096 = v4089
	v4097 = v4090
	goto L862
L862:
	;
	v4099 = *(*int32)(unsafe.Add(mBase, uint32(v4095)))
	*(*int32)(unsafe.Add(mBase, uint32(v4097))) = v4099
	v4101 = int32(4)
	v4102 = v4095 + v4101
	v4104 = v4097 + v4101
	v4106 = v4096 - v4101
	if base.Ui32(int32(3)) < base.Ui32(v4106) {
		v4095 = v4102
		v4096 = v4106
		v4097 = v4104
		goto L862
	} else {
		goto L864
	}
L863:
	;
	v4110 = v4102
	v4111 = v4106
	v4112 = v4104
	goto L829
L864:
	;
	goto L863
L865:
	;
	v4117 = v4110
	v4118 = v4111
	v4119 = v4112
	goto L866
L866:
	;
	v4121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4117))))
	*(*uint8)(unsafe.Add(mBase, uint32(v4119))) = uint8(v4121)
	v4123 = int32(1)
	v4128 = v4118 - v4123
	if v4128 != 0 {
		v4117 = v4117 + v4123
		v4118 = v4128
		v4119 = v4119 + v4123
		goto L866
	} else {
		goto L868
	}
L867:
	;
	goto L824
L868:
	;
	goto L867
L869:
	;
	if v4143 != 0 {
		goto L870
	} else {
		goto L871
	}
L870:
	;
	v4145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4140)+72)))
	v4146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v4140)+74)))
	v4147 = *(*int32)(unsafe.Add(mBase, uint32(v4140)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v3836+int32(16)))) = v4147
	*(*int32)(unsafe.Add(mBase, uint32(v3836)+12)) = v4146
	*(*int32)(unsafe.Add(mBase, uint32(v3836)+8)) = v4145
	*(*int32)(unsafe.Add(mBase, uint32(v3836)+4)) = v4140 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v3836))) = v3950
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_98), v3836)
	mBase = m.M
	v4157 = m.ExcPending
	if v4157 != 0 {
		goto L1
	} else {
		goto L873
	}
L871:
	;
	goto L872
L872:
	;
	v4166 = v3950 + int32(1)
	v4168 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	if v4166 < v4168 {
		v3950 = v4166
		goto L817
	} else {
		goto L875
	}
L873:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(475), int32(_a_F_BootstrapModeMain_97))
	mBase = m.M
	v4162 = m.ExcPending
	if v4162 != 0 {
		goto L1
	} else {
		goto L874
	}
L874:
	;
	goto L872
L875:
	;
	goto L818
L876:
	;
	v4207 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4207 != 0 {
		goto L877
	} else {
		goto L878
	}
L877:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4209 = m.ExcPending
	if v4209 != 0 {
		goto L1
	} else {
		goto L880
	}
L878:
	;
	goto L879
L879:
	;
	v4210 = int32(0)
	v4211 = F_isatty(m, v4210)
	mBase = m.M
	if v4211 == v4210 {
		v6156 = v3812
		goto L730
	} else {
		goto L881
	}
L880:
	;
	goto L879
L881:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4217 = m.ExcPending
	if v4217 != 0 {
		goto L1
	} else {
		goto L882
	}
L882:
	;
	v4218 = F_fflush(m, v3791)
	mBase = m.M
	v4219 = m.ExcPending
	if v4219 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	v6156 = v3812
	goto L730
L884:
	;
	v4226 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4231 = F_AllocSetContextCreateInternal(m, v4226, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_27), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4232 = m.ExcPending
	if v4232 != 0 {
		goto L1
	} else {
		goto L887
	}
L885:
	;
	v4234 = v4221
	goto L886
L886:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4234
	v4237 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	F_closerel(m, v4237)
	mBase = m.M
	v4239 = m.ExcPending
	if v4239 != 0 {
		goto L1
	} else {
		goto L888
	}
L887:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4231
	v4234 = v4231
	goto L886
L888:
	;
	v4242 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4242
	v4245 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4245)
	mBase = m.M
	v4247 = m.ExcPending
	if v4247 != 0 {
		goto L1
	} else {
		goto L889
	}
L889:
	;
	v4249 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4249 != 0 {
		goto L890
	} else {
		goto L891
	}
L890:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4251 = m.ExcPending
	if v4251 != 0 {
		goto L1
	} else {
		goto L893
	}
L891:
	;
	goto L892
L892:
	;
	v4252 = int32(0)
	v4253 = F_isatty(m, v4252)
	mBase = m.M
	if v4253 == v4252 {
		v6156 = v3812
		goto L730
	} else {
		goto L894
	}
L893:
	;
	goto L892
L894:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4259 = m.ExcPending
	if v4259 != 0 {
		goto L1
	} else {
		goto L895
	}
L895:
	;
	v4260 = F_fflush(m, v3791)
	mBase = m.M
	v4261 = m.ExcPending
	if v4261 != 0 {
		goto L1
	} else {
		goto L896
	}
L896:
	;
	v6156 = v3812
	goto L730
L897:
	;
	v4268 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4273 = F_AllocSetContextCreateInternal(m, v4268, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_27), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4274 = m.ExcPending
	if v4274 != 0 {
		goto L1
	} else {
		goto L900
	}
L898:
	;
	v4276 = v4263
	goto L899
L899:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4276
	v4280 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55])) = v4280
	v4284 = F_errstart(m, int32(11), v4280)
	mBase = m.M
	v4285 = m.ExcPending
	if v4285 != 0 {
		goto L1
	} else {
		goto L901
	}
L900:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4273
	v4276 = v4273
	goto L899
L901:
	;
	if v4284 == int32(0) {
		v6156 = v3812
		goto L730
	} else {
		goto L902
	}
L902:
	;
	v4290 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(12))))
	v4293 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(8))))
	v4296 = *(*int64)(unsafe.Add(mBase, uint32(v3781-int32(20))))
	*(*int64)(unsafe.Add(mBase, uint32(v3784)+8)) = v4296
	if v4293 != 0 {
		goto L903
	} else {
		goto L904
	}
L903:
	;
	v4300 = int32(_a_F_BootstrapModeMain_100)
	goto L905
L904:
	;
	v4300 = int32(_a_F_BootstrapModeMain_101)
	goto L905
L905:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+4)) = v4300
	if v4290 != 0 {
		goto L906
	} else {
		goto L907
	}
L906:
	;
	v4304 = int32(_a_F_BootstrapModeMain_102)
	goto L908
L907:
	;
	v4304 = int32(_a_F_BootstrapModeMain_101)
	goto L908
L908:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3784))) = v4304
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_103), v3784)
	mBase = m.M
	v4308 = m.ExcPending
	if v4308 != 0 {
		goto L1
	} else {
		goto L909
	}
L909:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(166), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v4313 = m.ExcPending
	if v4313 != 0 {
		goto L1
	} else {
		goto L910
	}
L910:
	;
	v6156 = v3812
	goto L730
L911:
	;
	v4323 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4323 != 0 {
		goto L912
	} else {
		goto L913
	}
L912:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4325 = m.ExcPending
	if v4325 != 0 {
		goto L1
	} else {
		goto L915
	}
L913:
	;
	goto L914
L914:
	;
	v4326 = int32(0)
	v4327 = F_isatty(m, v4326)
	mBase = m.M
	if v4327 == v4326 {
		v6156 = v3812
		goto L730
	} else {
		goto L916
	}
L915:
	;
	goto L914
L916:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4333 = m.ExcPending
	if v4333 != 0 {
		goto L1
	} else {
		goto L917
	}
L917:
	;
	v4334 = F_fflush(m, v3791)
	mBase = m.M
	v4335 = m.ExcPending
	if v4335 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	v6156 = v3812
	goto L730
L919:
	;
	v4342 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4347 = F_AllocSetContextCreateInternal(m, v4342, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_27), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L1
	} else {
		goto L922
	}
L920:
	;
	v4350 = v4337
	goto L921
L921:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4350
	v4354 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	v4356 = F_CreateTupleDesc(m, v4354, int32(_a_F_BootstrapModeMain_59))
	mBase = m.M
	v4357 = m.ExcPending
	if v4357 != 0 {
		goto L1
	} else {
		goto L923
	}
L922:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4347
	v4350 = v4347
	goto L921
L923:
	;
	v4360 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(24))))
	v4363 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(28))))
	if v4363 != 0 {
		goto L925
	} else {
		goto L926
	}
L924:
	;
	v4474 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4474
	v4477 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4477)
	mBase = m.M
	v4479 = m.ExcPending
	if v4479 != 0 {
		goto L1
	} else {
		goto L954
	}
L925:
	;
	v4365 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v4365 != 0 {
		goto L928
	} else {
		goto L929
	}
L926:
	;
	goto L927
L927:
	;
	v4426 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(36))))
	if v4360 != 0 {
		goto L946
	} else {
		goto L947
	}
L928:
	;
	v4368 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4369 = m.ExcPending
	if v4369 != 0 {
		goto L1
	} else {
		goto L931
	}
L929:
	;
	goto L930
L930:
	;
	v4385 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(36))))
	if v4360 != 0 {
		goto L938
	} else {
		goto L939
	}
L931:
	;
	if v4368 != 0 {
		goto L932
	} else {
		goto L933
	}
L932:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_106), int32(0))
	mBase = m.M
	v4373 = m.ExcPending
	if v4373 != 0 {
		goto L1
	} else {
		goto L935
	}
L933:
	;
	goto L934
L934:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v4381 = m.ExcPending
	if v4381 != 0 {
		goto L1
	} else {
		goto L937
	}
L935:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(202), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v4378 = m.ExcPending
	if v4378 != 0 {
		goto L1
	} else {
		goto L936
	}
L936:
	;
	goto L934
L937:
	;
	goto L930
L938:
	;
	v4389 = int32(1664)
	goto L940
L939:
	;
	v4389 = int32(0)
	goto L940
L940:
	;
	v4392 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(32))))
	v4393 = int32(0)
	v4396 = int32(112)
	v4399 = int32(1)
	v4406 = F_heap_create(m, v4385, int32(11), v4389, v4392, v4393, int32(2), v4356, int32(114), v4396, base.B2i32(v4360 != v4393), v4399, v4399, v3784+v4396, v3784+int32(124), v4399)
	mBase = m.M
	v4407 = m.ExcPending
	if v4407 != 0 {
		goto L1
	} else {
		goto L941
	}
L941:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60])) = v4406
	v4411 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4412 = m.ExcPending
	if v4412 != 0 {
		goto L1
	} else {
		goto L942
	}
L942:
	;
	if v4411 == int32(0) {
		goto L924
	} else {
		goto L943
	}
L943:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_107), int32(0))
	mBase = m.M
	v4418 = m.ExcPending
	if v4418 != 0 {
		goto L1
	} else {
		goto L944
	}
L944:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(221), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v4423 = m.ExcPending
	if v4423 != 0 {
		goto L1
	} else {
		goto L945
	}
L945:
	;
	goto L924
L946:
	;
	v4430 = int32(1664)
	goto L948
L947:
	;
	v4430 = int32(0)
	goto L948
L948:
	;
	v4433 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(32))))
	v4436 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(20))))
	v4437 = int32(0)
	v4444 = base.B2i32(v4360 != v4437)
	v4452 = F_heap_create_with_catalog(m, v4426, int32(11), v4430, v4433, v4436, v4437, int32(10), int32(2), v4356, v4437, int32(114), int32(112), v4444, v4444, v4437, v4437, v4437, int32(1), v4437, v4437, v4437)
	mBase = m.M
	v4453 = m.ExcPending
	if v4453 != 0 {
		goto L1
	} else {
		goto L949
	}
L949:
	;
	v4456 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4457 = m.ExcPending
	if v4457 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	if v4456 == int32(0) {
		goto L924
	} else {
		goto L951
	}
L951:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+16)) = v4452
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_108), v3784+int32(16))
	mBase = m.M
	v4465 = m.ExcPending
	if v4465 != 0 {
		goto L1
	} else {
		goto L952
	}
L952:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(248), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v4470 = m.ExcPending
	if v4470 != 0 {
		goto L1
	} else {
		goto L953
	}
L953:
	;
	goto L924
L954:
	;
	v4481 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4481 != 0 {
		goto L955
	} else {
		goto L956
	}
L955:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4483 = m.ExcPending
	if v4483 != 0 {
		goto L1
	} else {
		goto L958
	}
L956:
	;
	goto L957
L957:
	;
	v4484 = int32(0)
	v4485 = F_isatty(m, v4484)
	mBase = m.M
	if v4485 == v4484 {
		v6156 = v3812
		goto L730
	} else {
		goto L959
	}
L958:
	;
	goto L957
L959:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4491 = m.ExcPending
	if v4491 != 0 {
		goto L1
	} else {
		goto L960
	}
L960:
	;
	v4492 = F_fflush(m, v3791)
	mBase = m.M
	v4493 = m.ExcPending
	if v4493 != 0 {
		goto L1
	} else {
		goto L961
	}
L961:
	;
	v6156 = v3812
	goto L730
L962:
	;
	v4500 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4505 = F_AllocSetContextCreateInternal(m, v4500, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_27), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4506 = m.ExcPending
	if v4506 != 0 {
		goto L1
	} else {
		goto L965
	}
L963:
	;
	v4508 = v4495
	goto L964
L964:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4508
	v4513 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4514 = m.ExcPending
	if v4514 != 0 {
		goto L1
	} else {
		goto L966
	}
L965:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4505
	v4508 = v4505
	goto L964
L966:
	;
	if v4513 != 0 {
		goto L967
	} else {
		goto L968
	}
L967:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_109), int32(0))
	mBase = m.M
	v4518 = m.ExcPending
	if v4518 != 0 {
		goto L1
	} else {
		goto L970
	}
L968:
	;
	goto L969
L969:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54])) = int32(0)
	v6156 = v3812
	goto L730
L970:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(258), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v4523 = m.ExcPending
	if v4523 != 0 {
		goto L1
	} else {
		goto L971
	}
L971:
	;
	goto L969
L972:
	;
	v4533 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v4533 == int32(0) {
		goto L300
	} else {
		goto L973
	}
L973:
	;
	v4536 = m.G0
	v4538 = v4536 - int32(16)
	m.G0 = v4538
	v4542 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4543 = m.ExcPending
	if v4543 != 0 {
		goto L1
	} else {
		goto L974
	}
L974:
	;
	if v4542 != 0 {
		goto L975
	} else {
		goto L976
	}
L975:
	;
	v4545 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v4538))) = v4545
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_110), v4538)
	mBase = m.M
	v4549 = m.ExcPending
	if v4549 != 0 {
		goto L1
	} else {
		goto L978
	}
L976:
	;
	goto L977
L977:
	;
	v4556 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	v4558 = F_CreateTupleDesc(m, v4556, int32(_a_F_BootstrapModeMain_59))
	mBase = m.M
	v4559 = m.ExcPending
	if v4559 != 0 {
		goto L1
	} else {
		goto L980
	}
L978:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(635), int32(_a_F_BootstrapModeMain_111))
	mBase = m.M
	v4554 = m.ExcPending
	if v4554 != 0 {
		goto L1
	} else {
		goto L979
	}
L979:
	;
	goto L977
L980:
	;
	v4562 = F_heap_form_tuple(m, v4558, int32(_a_F_BootstrapModeMain_112), int32(_a_F_BootstrapModeMain_113))
	mBase = m.M
	v4563 = m.ExcPending
	if v4563 != 0 {
		goto L1
	} else {
		goto L981
	}
L981:
	;
	F_pfree(m, v4558)
	mBase = m.M
	v4565 = m.ExcPending
	if v4565 != 0 {
		goto L1
	} else {
		goto L982
	}
L982:
	;
	v4567 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	F_simple_heap_insert(m, v4567, v4562)
	mBase = m.M
	v4569 = m.ExcPending
	if v4569 != 0 {
		goto L1
	} else {
		goto L983
	}
L983:
	;
	F_pfree(m, v4562)
	mBase = m.M
	v4571 = m.ExcPending
	if v4571 != 0 {
		goto L1
	} else {
		goto L984
	}
L984:
	;
	v4574 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4575 = m.ExcPending
	if v4575 != 0 {
		goto L1
	} else {
		goto L985
	}
L985:
	;
	if v4574 != 0 {
		goto L986
	} else {
		goto L987
	}
L986:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_114), int32(0))
	mBase = m.M
	v4579 = m.ExcPending
	if v4579 != 0 {
		goto L1
	} else {
		goto L989
	}
L987:
	;
	goto L988
L988:
	;
	v4586 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	if int32(0) < v4586 {
		goto L991
	} else {
		goto L992
	}
L989:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(643), int32(_a_F_BootstrapModeMain_111))
	mBase = m.M
	v4584 = m.ExcPending
	if v4584 != 0 {
		goto L1
	} else {
		goto L990
	}
L990:
	;
	goto L988
L991:
	;
	v4592 = F__emscripten_memset_bulkmem(m, int32(_a_F_BootstrapModeMain_113), base.I32_extend8_s(int32(0)), v4586)
	mBase = m.M
	goto L994
L992:
	;
	goto L993
L993:
	;
	m.G0 = v4538 + int32(16)
	v4598 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4598
	v4601 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4601)
	mBase = m.M
	v4603 = m.ExcPending
	if v4603 != 0 {
		goto L1
	} else {
		goto L995
	}
L994:
	;
	goto L993
L995:
	;
	v4605 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4605 != 0 {
		goto L996
	} else {
		goto L997
	}
L996:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4607 = m.ExcPending
	if v4607 != 0 {
		goto L1
	} else {
		goto L999
	}
L997:
	;
	goto L998
L998:
	;
	v4608 = int32(0)
	v4609 = F_isatty(m, v4608)
	mBase = m.M
	if v4609 == v4608 {
		v6156 = v3812
		goto L730
	} else {
		goto L1000
	}
L999:
	;
	goto L998
L1000:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4615 = m.ExcPending
	if v4615 != 0 {
		goto L1
	} else {
		goto L1001
	}
L1001:
	;
	v4616 = F_fflush(m, v3791)
	mBase = m.M
	v4617 = m.ExcPending
	if v4617 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1002:
	;
	v6156 = v3812
	goto L730
L1003:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4619))) = int32(204)
	v4625 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4626 = m.ExcPending
	if v4626 != 0 {
		goto L1
	} else {
		goto L1004
	}
L1004:
	;
	if v4625 != 0 {
		goto L1005
	} else {
		goto L1006
	}
L1005:
	;
	v4629 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+48)) = v4629
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_115), v3784+int32(48))
	mBase = m.M
	v4635 = m.ExcPending
	if v4635 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1006:
	;
	goto L1007
L1007:
	;
	v4642 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4642 == int32(0) {
		goto L1010
	} else {
		goto L1011
	}
L1008:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(279), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v4640 = m.ExcPending
	if v4640 != 0 {
		goto L1
	} else {
		goto L1009
	}
L1009:
	;
	goto L1007
L1010:
	;
	v4647 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4652 = F_AllocSetContextCreateInternal(m, v4647, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_27), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4653 = m.ExcPending
	if v4653 != 0 {
		goto L1
	} else {
		goto L1013
	}
L1011:
	;
	v4655 = v4642
	goto L1012
L1012:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4655
	v4660 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4619)+4)) = v4660
	v4665 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(20))))
	v4667 = F_makeRangeVar(m, int32(0), v4665, int32(-1))
	mBase = m.M
	v4668 = m.ExcPending
	if v4668 != 0 {
		goto L1
	} else {
		goto L1014
	}
L1013:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4652
	v4655 = v4652
	goto L1012
L1014:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4619)+8)) = v4667
	v4672 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(12))))
	v4673 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4619)+16)) = v4673
	*(*int32)(unsafe.Add(mBase, uint32(v4619)+12)) = v4672
	v4678 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(4))))
	*(*uint16)(unsafe.Add(mBase, uint32(v4619)+62)) = uint16(v4673)
	*(*int32)(unsafe.Add(mBase, uint32(v4619)+20)) = v4678
	v4682 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4619)+24)) = v4682
	*(*int64)(unsafe.Add(mBase, uint32(v4619)+32)) = v4682
	*(*int64)(unsafe.Add(mBase, uint32(v4619)+40)) = v4682
	*(*int64)(unsafe.Add(mBase, uint32(v4619)+48)) = v4682
	*(*int64)(unsafe.Add(mBase, uint32(v4619)+53)) = v4682
	*(*int32)(unsafe.Add(mBase, uint32(v4619)+65)) = v4673
	*(*uint16)(unsafe.Add(mBase, uint32(v4619)+69)) = uint16(v4673)
	v4702 = F_RangeVarGetRelidExtended(m, v4667, v4673, v4673, v4673, v4673)
	mBase = m.M
	v4703 = m.ExcPending
	if v4703 != 0 {
		goto L1
	} else {
		goto L1015
	}
L1015:
	;
	v4706 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(28))))
	v4707 = int32(0)
	F_DefineIndex(m, v3784+int32(112), v4702, v4619, v4706, v4707, v4707, int32(-1), v4707, v4707, v4707, int32(1), v4707)
	mBase = m.M
	v4716 = m.ExcPending
	if v4716 != 0 {
		goto L1
	} else {
		goto L1016
	}
L1016:
	;
	v4719 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4719
	v4722 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4722)
	mBase = m.M
	v4724 = m.ExcPending
	if v4724 != 0 {
		goto L1
	} else {
		goto L1017
	}
L1017:
	;
	v4726 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4726 != 0 {
		goto L1018
	} else {
		goto L1019
	}
L1018:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4728 = m.ExcPending
	if v4728 != 0 {
		goto L1
	} else {
		goto L1021
	}
L1019:
	;
	goto L1020
L1020:
	;
	v4729 = int32(0)
	v4730 = F_isatty(m, v4729)
	mBase = m.M
	if v4730 == v4729 {
		v6156 = v3812
		goto L730
	} else {
		goto L1022
	}
L1021:
	;
	goto L1020
L1022:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4736 = m.ExcPending
	if v4736 != 0 {
		goto L1
	} else {
		goto L1023
	}
L1023:
	;
	v4737 = F_fflush(m, v3791)
	mBase = m.M
	v4738 = m.ExcPending
	if v4738 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1024:
	;
	v6156 = v3812
	goto L730
L1025:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4740))) = int32(204)
	v4746 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4747 = m.ExcPending
	if v4747 != 0 {
		goto L1
	} else {
		goto L1026
	}
L1026:
	;
	if v4746 != 0 {
		goto L1027
	} else {
		goto L1028
	}
L1027:
	;
	v4750 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+64)) = v4750
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_116), v3784-int32(-64))
	mBase = m.M
	v4756 = m.ExcPending
	if v4756 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1028:
	;
	goto L1029
L1029:
	;
	v4763 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4763 == int32(0) {
		goto L1032
	} else {
		goto L1033
	}
L1030:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(332), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v4761 = m.ExcPending
	if v4761 != 0 {
		goto L1
	} else {
		goto L1031
	}
L1031:
	;
	goto L1029
L1032:
	;
	v4768 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4773 = F_AllocSetContextCreateInternal(m, v4768, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_27), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4774 = m.ExcPending
	if v4774 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1033:
	;
	v4776 = v4763
	goto L1034
L1034:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4776
	v4781 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4740)+4)) = v4781
	v4786 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(20))))
	v4788 = F_makeRangeVar(m, int32(0), v4786, int32(-1))
	mBase = m.M
	v4789 = m.ExcPending
	if v4789 != 0 {
		goto L1
	} else {
		goto L1036
	}
L1035:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4773
	v4776 = v4773
	goto L1034
L1036:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4740)+8)) = v4788
	v4793 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(12))))
	v4794 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4740)+16)) = v4794
	*(*int32)(unsafe.Add(mBase, uint32(v4740)+12)) = v4793
	v4799 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(4))))
	v4800 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4740)+24)) = v4800
	*(*int32)(unsafe.Add(mBase, uint32(v4740)+20)) = v4799
	*(*int64)(unsafe.Add(mBase, uint32(v4740)+32)) = v4800
	*(*int64)(unsafe.Add(mBase, uint32(v4740)+40)) = v4800
	*(*int64)(unsafe.Add(mBase, uint32(v4740)+48)) = v4800
	*(*int32)(unsafe.Add(mBase, uint32(v4740)+56)) = v4794
	*(*int32)(unsafe.Add(mBase, uint32(v4740)+65)) = v4794
	*(*uint16)(unsafe.Add(mBase, uint32(v4740)+62)) = uint16(v4794)
	v4815 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4740)+60)) = uint8(v4815)
	*(*uint16)(unsafe.Add(mBase, uint32(v4740)+69)) = uint16(v4794)
	v4825 = F_RangeVarGetRelidExtended(m, v4788, v4794, v4794, v4794, v4794)
	mBase = m.M
	v4826 = m.ExcPending
	if v4826 != 0 {
		goto L1
	} else {
		goto L1037
	}
L1037:
	;
	v4829 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(28))))
	v4830 = int32(0)
	F_DefineIndex(m, v3784+int32(112), v4825, v4740, v4829, v4830, v4830, int32(-1), v4830, v4830, v4830, int32(1), v4830)
	mBase = m.M
	v4839 = m.ExcPending
	if v4839 != 0 {
		goto L1
	} else {
		goto L1038
	}
L1038:
	;
	v4842 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4842
	v4845 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4845)
	mBase = m.M
	v4847 = m.ExcPending
	if v4847 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1039:
	;
	v4849 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4849 != 0 {
		goto L1040
	} else {
		goto L1041
	}
L1040:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4851 = m.ExcPending
	if v4851 != 0 {
		goto L1
	} else {
		goto L1043
	}
L1041:
	;
	goto L1042
L1042:
	;
	v4852 = int32(0)
	v4853 = F_isatty(m, v4852)
	mBase = m.M
	if v4853 == v4852 {
		v6156 = v3812
		goto L730
	} else {
		goto L1044
	}
L1043:
	;
	goto L1042
L1044:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4859 = m.ExcPending
	if v4859 != 0 {
		goto L1
	} else {
		goto L1045
	}
L1045:
	;
	v4860 = F_fflush(m, v3791)
	mBase = m.M
	v4861 = m.ExcPending
	if v4861 != 0 {
		goto L1
	} else {
		goto L1046
	}
L1046:
	;
	v6156 = v3812
	goto L730
L1047:
	;
	if v4864 != 0 {
		goto L1048
	} else {
		goto L1049
	}
L1048:
	;
	v4866 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+80)) = v4866
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_117), v3784+int32(80))
	mBase = m.M
	v4872 = m.ExcPending
	if v4872 != 0 {
		goto L1
	} else {
		goto L1051
	}
L1049:
	;
	goto L1050
L1050:
	;
	v4879 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	if v4879 == int32(0) {
		goto L1053
	} else {
		goto L1054
	}
L1051:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(382), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v4877 = m.ExcPending
	if v4877 != 0 {
		goto L1
	} else {
		goto L1052
	}
L1052:
	;
	goto L1050
L1053:
	;
	v4884 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4889 = F_AllocSetContextCreateInternal(m, v4884, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_27), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4890 = m.ExcPending
	if v4890 != 0 {
		goto L1
	} else {
		goto L1056
	}
L1054:
	;
	v4892 = v4879
	goto L1055
L1055:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4892
	v4895 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v4898 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(12))))
	v4901 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(8))))
	v4902 = m.G0
	v4904 = v4902 - int32(32)
	m.G0 = v4904
	v4908 = F_makeRangeVar(m, int32(0), v4895, int32(-1))
	mBase = m.M
	v4909 = m.ExcPending
	if v4909 != 0 {
		goto L1
	} else {
		goto L1059
	}
L1056:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4889
	v4892 = v4889
	goto L1055
L1057:
	;
	v4930 = int32(0)
	v4934 = F_create_toast_table(m, v4911, v4898, v4901, v4930, int32(8), v4930, v4930)
	mBase = m.M
	v4935 = m.ExcPending
	if v4935 != 0 {
		goto L1
	} else {
		goto L1064
	}
L1058:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4920 = m.ExcPending
	if v4920 != 0 {
		goto L1
	} else {
		goto L1061
	}
L1059:
	;
	v4911 = F_table_openrv(m, v4908, int32(8))
	mBase = m.M
	v4912 = m.ExcPending
	if v4912 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1060:
	;
	v4913 = *(*int32)(unsafe.Add(mBase, uint32(v4911)+48))
	v4914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4913)+119)))
	switch v4914 - int32(109) {
	case 0, 5:
		goto L1057
	default:
		goto L1058
	}
L1061:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4904))) = v4895
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_118), v4904)
	mBase = m.M
	v4924 = m.ExcPending
	if v4924 != 0 {
		goto L1
	} else {
		goto L1062
	}
L1062:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_119), int32(107), int32(_a_F_BootstrapModeMain_120))
	mBase = m.M
	v4929 = m.ExcPending
	if v4929 != 0 {
		goto L1
	} else {
		goto L1063
	}
L1063:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1064:
	;
	if v4934 == int32(0) {
		goto L1065
	} else {
		goto L1066
	}
L1065:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4941 = m.ExcPending
	if v4941 != 0 {
		goto L1
	} else {
		goto L1068
	}
L1066:
	;
	goto L1067
L1067:
	;
	F_sequence_close(m, v4911, int32(0))
	mBase = m.M
	v4955 = m.ExcPending
	if v4955 != 0 {
		goto L1
	} else {
		goto L1071
	}
L1068:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4904)+16)) = v4895
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_121), v4904+int32(16))
	mBase = m.M
	v4947 = m.ExcPending
	if v4947 != 0 {
		goto L1
	} else {
		goto L1069
	}
L1069:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_119), int32(113), int32(_a_F_BootstrapModeMain_120))
	mBase = m.M
	v4952 = m.ExcPending
	if v4952 != 0 {
		goto L1
	} else {
		goto L1070
	}
L1070:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1071:
	;
	m.G0 = v4904 + int32(32)
	v4961 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4961
	v4964 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v4964)
	mBase = m.M
	v4966 = m.ExcPending
	if v4966 != 0 {
		goto L1
	} else {
		goto L1072
	}
L1072:
	;
	v4968 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v4968 != 0 {
		goto L1073
	} else {
		goto L1074
	}
L1073:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4970 = m.ExcPending
	if v4970 != 0 {
		goto L1
	} else {
		goto L1076
	}
L1074:
	;
	goto L1075
L1075:
	;
	v4971 = int32(0)
	v4972 = F_isatty(m, v4971)
	mBase = m.M
	if v4972 == v4971 {
		v6156 = v3812
		goto L730
	} else {
		goto L1077
	}
L1076:
	;
	goto L1075
L1077:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4978 = m.ExcPending
	if v4978 != 0 {
		goto L1
	} else {
		goto L1078
	}
L1078:
	;
	v4979 = F_fflush(m, v3791)
	mBase = m.M
	v4980 = m.ExcPending
	if v4980 != 0 {
		goto L1
	} else {
		goto L1079
	}
L1079:
	;
	v6156 = v3812
	goto L730
L1080:
	;
	v4987 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	v4992 = F_AllocSetContextCreateInternal(m, v4987, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_27), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4993 = m.ExcPending
	if v4993 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1081:
	;
	v4995 = v4982
	goto L1082
L1082:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v4995
	v4999 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64]))
	if v4999 != 0 {
		goto L1084
	} else {
		goto L1085
	}
L1083:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v4992
	v4995 = v4992
	goto L1082
L1084:
	;
	v5000 = v4999
	goto L1087
L1085:
	;
	goto L1086
L1086:
	;
	v5080 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58])) = v5080
	v5083 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56]))
	F_MemoryContextReset(m, v5083)
	mBase = m.M
	v5085 = m.ExcPending
	if v5085 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1087:
	;
	v5025 = *(*int32)(unsafe.Add(mBase, uint32(v5000)))
	v5027 = F_table_open(m, v5025, int32(0))
	mBase = m.M
	v5028 = m.ExcPending
	if v5028 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1088:
	;
	goto L1086
L1089:
	;
	v5030 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64]))
	v5031 = *(*int32)(unsafe.Add(mBase, uint32(v5030)+4))
	v5033 = F_index_open(m, v5031, int32(0))
	mBase = m.M
	v5034 = m.ExcPending
	if v5034 != 0 {
		goto L1
	} else {
		goto L1090
	}
L1090:
	;
	v5036 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64]))
	v5037 = *(*int32)(unsafe.Add(mBase, uint32(v5036)+8))
	v5038 = int32(0)
	F_index_build(m, v5027, v5033, v5037, v5038, v5038)
	mBase = m.M
	v5041 = m.ExcPending
	if v5041 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1091:
	;
	F_relation_close(m, v5033, int32(0))
	mBase = m.M
	v5044 = m.ExcPending
	if v5044 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1092:
	;
	F_sequence_close(m, v5027, int32(0))
	mBase = m.M
	v5047 = m.ExcPending
	if v5047 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1093:
	;
	v5048 = int32(_a_F_BootstrapModeMain_122)
	v5050 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64]))
	v5051 = *(*int32)(unsafe.Add(mBase, uint32(v5050)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[64])) = v5051
	if v5051 != 0 {
		v5000 = v5051
		goto L1087
	} else {
		goto L1094
	}
L1094:
	;
	goto L1088
L1095:
	;
	v5087 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[63]))
	if v5087 != 0 {
		goto L1096
	} else {
		goto L1097
	}
L1096:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v5089 = m.ExcPending
	if v5089 != 0 {
		goto L1
	} else {
		goto L1099
	}
L1097:
	;
	goto L1098
L1098:
	;
	v5090 = int32(0)
	v5091 = F_isatty(m, v5090)
	mBase = m.M
	if v5091 == v5090 {
		v6156 = v3812
		goto L730
	} else {
		goto L1100
	}
L1099:
	;
	goto L1098
L1100:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v5097 = m.ExcPending
	if v5097 != 0 {
		goto L1
	} else {
		goto L1101
	}
L1101:
	;
	v5098 = F_fflush(m, v3791)
	mBase = m.M
	v5099 = m.ExcPending
	if v5099 != 0 {
		goto L1
	} else {
		goto L1102
	}
L1102:
	;
	v6156 = v3812
	goto L730
L1103:
	;
	v6156 = v5104
	goto L730
L1104:
	;
	v6156 = v5112
	goto L730
L1105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5115))) = int32(92)
	v5121 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v5115)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5115)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5115)+4)) = v5121
	v5127 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v5128 = F_makeString(m, v5127)
	mBase = m.M
	v5129 = m.ExcPending
	if v5129 != 0 {
		goto L1
	} else {
		goto L1106
	}
L1106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+100)) = v5128
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+104)) = v5128
	v5135 = F_list_make1_impl(m, int32(1), v3784+int32(100))
	mBase = m.M
	v5136 = m.ExcPending
	if v5136 != 0 {
		goto L1
	} else {
		goto L1107
	}
L1107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5115)+28)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5115)+20)) = v5135
	v6156 = v5115
	goto L730
L1108:
	;
	v5153 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(12))))
	v5156 = *(*int32)(unsafe.Add(mBase, uint32(v3781-int32(4))))
	v5157 = *(*int32)(unsafe.Add(mBase, uint32(v3781)))
	v5158 = m.G0
	v5160 = v5158 - int32(48)
	m.G0 = v5160
	v5163 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v5163 != 0 {
		goto L1109
	} else {
		goto L1110
	}
L1109:
	;
	v5166 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5167 = m.ExcPending
	if v5167 != 0 {
		goto L1
	} else {
		goto L1112
	}
L1110:
	;
	goto L1111
L1111:
	;
	v5181 = v5145 << (uint(int32(2)) % 32)
	v5184 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	if v5184 == int32(0) {
		goto L1119
	} else {
		goto L1120
	}
L1112:
	;
	if v5166 != 0 {
		goto L1113
	} else {
		goto L1114
	}
L1113:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_123), int32(0))
	mBase = m.M
	v5171 = m.ExcPending
	if v5171 != 0 {
		goto L1
	} else {
		goto L1116
	}
L1114:
	;
	goto L1115
L1115:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v5179 = m.ExcPending
	if v5179 != 0 {
		goto L1
	} else {
		goto L1118
	}
L1116:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(528), int32(_a_F_BootstrapModeMain_124))
	mBase = m.M
	v5176 = m.ExcPending
	if v5176 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1117:
	;
	goto L1115
L1118:
	;
	goto L1111
L1119:
	;
	v5188 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[62]))
	v5190 = F_MemoryContextAllocZero(m, v5188, int32(100))
	mBase = m.M
	v5191 = m.ExcPending
	if v5191 != 0 {
		goto L1
	} else {
		goto L1122
	}
L1120:
	;
	v5193 = v5184
	goto L1121
L1121:
	;
	if v5193&int32(3) == int32(0) {
		goto L1124
	} else {
		goto L1125
	}
L1122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61]))) = v5190
	v5193 = v5190
	goto L1121
L1123:
	;
	v5221 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5225 = F_strncpy(m, v5221+int32(4), v5153, int32(64))
	mBase = m.M
	v5226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5225)+63)) = uint8(v5226)
	goto L1133
L1124:
	;
	v5199 = v5193 + int32(100)
	if base.Ui32(v5199) <= base.Ui32(v5193) {
		goto L1123
	} else {
		goto L1127
	}
L1125:
	;
	goto L1126
L1126:
	;
	v5218 = F__emscripten_memset_bulkmem(m, v5193, base.I32_extend8_s(int32(0)), int32(100))
	mBase = m.M
	goto L1132
L1127:
	;
	v5205 = v5193 + int32(4)
	if base.Ui32(v5205) < base.Ui32(v5199) {
		goto L1128
	} else {
		goto L1129
	}
L1128:
	;
	v5207 = v5199
	goto L1130
L1129:
	;
	v5207 = v5205
	goto L1130
L1130:
	;
	v5214 = F__emscripten_memset_bulkmem(m, v5193, base.I32_extend8_s(int32(0)), (v5193^int32(-1)+v5207)&int32(-4)+int32(4))
	mBase = m.M
	goto L1131
L1131:
	;
	goto L1123
L1132:
	;
	goto L1123
L1133:
	;
	v5228 = int32(0)
	v5231 = F_errstart(m, int32(11), v5228)
	mBase = m.M
	v5232 = m.ExcPending
	if v5232 != 0 {
		goto L1
	} else {
		goto L1134
	}
L1134:
	;
	if v5231 != 0 {
		goto L1135
	} else {
		goto L1136
	}
L1135:
	;
	v5233 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	*(*int32)(unsafe.Add(mBase, uint32(v5160)+36)) = v5156
	*(*int32)(unsafe.Add(mBase, uint32(v5160)+32)) = v5233 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_125), v5160+int32(32))
	mBase = m.M
	v5242 = m.ExcPending
	if v5242 != 0 {
		goto L1
	} else {
		goto L1138
	}
L1136:
	;
	goto L1137
L1137:
	;
	v5249 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5251 = v5145 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5249)+74)) = uint16(v5251)
	v5254 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59]))
	if v5254 == int32(0) {
		goto L1144
	} else {
		goto L1145
	}
L1138:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(537), int32(_a_F_BootstrapModeMain_124))
	mBase = m.M
	v5247 = m.ExcPending
	if v5247 != 0 {
		goto L1
	} else {
		goto L1139
	}
L1139:
	;
	goto L1137
L1140:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5761)+80)) = uint16(v5762)
	v5783 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5784 = *(*int32)(unsafe.Add(mBase, uint32(v5783)+96))
	if v5784 != 0 {
		goto L1225
	} else {
		goto L1226
	}
L1141:
	;
	v5761 = v5735
	v5762 = int32(1)
	goto L1140
L1142:
	;
	v5699 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[65])) = v5682
	v5702 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5703 = *(*int32)(unsafe.Add(mBase, uint32(v5682)))
	*(*int32)(unsafe.Add(mBase, uint32(v5702)+68)) = v5703
	v5705 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5706 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5682)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v5705)+72)) = uint16(v5706)
	v5708 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5682)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5708)+82)) = uint8(v5709)
	v5711 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5682)+132)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5711)+83)) = uint8(v5712)
	v5714 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5682)+133)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5714)+84)) = uint8(v5715)
	v5717 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5717)+85)) = uint8(v5699)
	v5720 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5721 = *(*int32)(unsafe.Add(mBase, uint32(v5682)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v5720)+96)) = v5721
	v5723 = *(*int32)(unsafe.Add(mBase, uint32(v5682)+96))
	if v5723 == v5699 {
		goto L1222
	} else {
		goto L1223
	}
L1143:
	;
	v5629 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5631 = v5262 * int32(92)
	v5634 = *(*int32)(unsafe.Add(mBase, uint32(v5631)+uint32(_c_F_BootstrapModeMain[66])))
	*(*int32)(unsafe.Add(mBase, uint32(v5629)+68)) = v5634
	v5636 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5631)+uint32(_c_F_BootstrapModeMain[67]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5636)+72)) = uint16(v5639)
	v5641 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5631)+uint32(_c_F_BootstrapModeMain[68]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5641)+82)) = uint8(v5644)
	v5646 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5631)+uint32(_c_F_BootstrapModeMain[69]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5646)+83)) = uint8(v5649)
	v5651 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5631)+uint32(_c_F_BootstrapModeMain[70]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5651)+84)) = uint8(v5654)
	v5656 = int32(0)
	v5657 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5657)+85)) = uint8(v5656)
	v5660 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5663 = *(*int32)(unsafe.Add(mBase, uint32(v5631)+uint32(_c_F_BootstrapModeMain[71])))
	*(*int32)(unsafe.Add(mBase, uint32(v5660)+96)) = v5663
	v5665 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5668 = *(*int32)(unsafe.Add(mBase, uint32(v5631)+uint32(_c_F_BootstrapModeMain[72])))
	if v5668 == v5656 {
		v5761 = v5665
		v5762 = v5656
		goto L1140
	} else {
		goto L1220
	}
L1144:
	;
	v5262 = v5228
	goto L1147
L1145:
	;
	v5364 = v5228
	v5377 = v5254
	goto L1146
L1146:
	;
	v5384 = *(*int32)(unsafe.Add(mBase, uint32(v5377)+4))
	if int32(0) < v5384 {
		goto L1173
	} else {
		goto L1174
	}
L1147:
	;
	v5285 = v5262*int32(92) + int32(_a_F_BootstrapModeMain_126)
	goto L1151
L1148:
	;
	v5364 = v5354
	v5377 = v5356
	goto L1146
L1149:
	;
	if v5322-v5323 == int32(0) {
		goto L1143
	} else {
		goto L1163
	}
L1151:
	;
	goto L1152
L1152:
	;
	v5292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5156))))
	if v5292 != 0 {
		goto L1153
	} else {
		goto L1154
	}
L1153:
	;
	v5293 = v5156
	v5294 = v5285
	v5295 = int32(64)
	v5296 = v5292
	goto L1157
L1154:
	;
	v5318 = v5285
	v5322 = int32(0)
	goto L1155
L1155:
	;
	v5323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5318))))
	goto L1149
L1156:
	;
	v5318 = v5313
	v5322 = v5315
	goto L1155
L1157:
	;
	v5298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5294))))
	if v5296 != v5298 {
		v5313 = v5294
		v5315 = v5296
		goto L1156
	} else {
		goto L1159
	}
L1158:
	;
	v5313 = v5307
	v5315 = int32(0)
	goto L1156
L1159:
	;
	if v5298 == int32(0) {
		v5313 = v5294
		v5315 = v5296
		goto L1156
	} else {
		goto L1160
	}
L1160:
	;
	v5303 = v5295 - int32(1)
	if v5303 == int32(0) {
		v5313 = v5294
		v5315 = v5296
		goto L1156
	} else {
		goto L1161
	}
L1161:
	;
	v5306 = int32(1)
	v5307 = v5294 + v5306
	v5308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5293)+1)))
	if v5308 != 0 {
		v5293 = v5293 + v5306
		v5294 = v5307
		v5295 = v5303
		v5296 = v5308
		goto L1157
	} else {
		goto L1162
	}
L1162:
	;
	goto L1158
L1163:
	;
	v5334 = v5262 + int32(1)
	if v5334 != int32(25) {
		v5262 = v5334
		goto L1147
	} else {
		goto L1164
	}
L1164:
	;
	v5339 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5340 = m.ExcPending
	if v5340 != 0 {
		goto L1
	} else {
		goto L1165
	}
L1165:
	;
	if v5339 != 0 {
		goto L1166
	} else {
		goto L1167
	}
L1166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5160)+16)) = v5156
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_127), v5160+int32(16))
	mBase = m.M
	v5346 = m.ExcPending
	if v5346 != 0 {
		goto L1
	} else {
		goto L1169
	}
L1167:
	;
	goto L1168
L1168:
	;
	F_populate_typ_list(m)
	mBase = m.M
	v5353 = m.ExcPending
	if v5353 != 0 {
		goto L1
	} else {
		goto L1171
	}
L1169:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(817), int32(_a_F_BootstrapModeMain_128))
	mBase = m.M
	v5351 = m.ExcPending
	if v5351 != 0 {
		goto L1
	} else {
		goto L1170
	}
L1170:
	;
	goto L1168
L1171:
	;
	v5354 = int32(0)
	v5356 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59]))
	if v5356 == v5354 {
		v5262 = v5354
		goto L1147
	} else {
		goto L1172
	}
L1172:
	;
	goto L1148
L1173:
	;
	v5387 = *(*int32)(unsafe.Add(mBase, uint32(v5377)+12))
	v5393 = v5364
	goto L1176
L1174:
	;
	goto L1175
L1175:
	;
	F_list_free_deep(m, v5377)
	mBase = m.M
	v5495 = m.ExcPending
	if v5495 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1176:
	;
	v5416 = *(*int32)(unsafe.Add(mBase, uint32(v5387+v5393<<(uint(int32(2))%32))))
	v5418 = v5416 + int32(8)
	goto L1180
L1177:
	;
	goto L1175
L1178:
	;
	if v5455-v5456 == int32(0) {
		v5682 = v5416
		goto L1142
	} else {
		goto L1192
	}
L1180:
	;
	goto L1181
L1181:
	;
	v5425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5418))))
	if v5425 != 0 {
		goto L1182
	} else {
		goto L1183
	}
L1182:
	;
	v5426 = v5418
	v5427 = v5156
	v5428 = int32(64)
	v5429 = v5425
	goto L1186
L1183:
	;
	v5451 = v5156
	v5455 = int32(0)
	goto L1184
L1184:
	;
	v5456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5451))))
	goto L1178
L1185:
	;
	v5451 = v5446
	v5455 = v5448
	goto L1184
L1186:
	;
	v5431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5427))))
	if v5429 != v5431 {
		v5446 = v5427
		v5448 = v5429
		goto L1185
	} else {
		goto L1188
	}
L1187:
	;
	v5446 = v5440
	v5448 = int32(0)
	goto L1185
L1188:
	;
	if v5431 == int32(0) {
		v5446 = v5427
		v5448 = v5429
		goto L1185
	} else {
		goto L1189
	}
L1189:
	;
	v5436 = v5428 - int32(1)
	if v5436 == int32(0) {
		v5446 = v5427
		v5448 = v5429
		goto L1185
	} else {
		goto L1190
	}
L1190:
	;
	v5439 = int32(1)
	v5440 = v5427 + v5439
	v5441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5426)+1)))
	if v5441 != 0 {
		v5426 = v5426 + v5439
		v5427 = v5440
		v5428 = v5436
		v5429 = v5441
		goto L1186
	} else {
		goto L1191
	}
L1191:
	;
	goto L1187
L1192:
	;
	v5467 = v5393 + int32(1)
	if v5467 != v5384 {
		v5393 = v5467
		goto L1176
	} else {
		goto L1193
	}
L1193:
	;
	goto L1177
L1194:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59])) = int32(0)
	F_populate_typ_list(m)
	mBase = m.M
	v5500 = m.ExcPending
	if v5500 != 0 {
		goto L1
	} else {
		goto L1195
	}
L1195:
	;
	v5502 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[59]))
	if v5502 == int32(0) {
		goto L1196
	} else {
		goto L1197
	}
L1196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5619 = m.ExcPending
	if v5619 != 0 {
		goto L1
	} else {
		goto L1217
	}
L1197:
	;
	v5505 = *(*int32)(unsafe.Add(mBase, uint32(v5502)+4))
	if v5505 <= int32(0) {
		goto L1196
	} else {
		goto L1198
	}
L1198:
	;
	v5508 = *(*int32)(unsafe.Add(mBase, uint32(v5502)+12))
	v5515 = int32(0)
	goto L1199
L1199:
	;
	v5538 = *(*int32)(unsafe.Add(mBase, uint32(v5508+v5515<<(uint(int32(2))%32))))
	v5540 = v5538 + int32(8)
	goto L1203
L1200:
	;
	goto L1196
L1201:
	;
	if v5577-v5578 == int32(0) {
		v5682 = v5538
		goto L1142
	} else {
		goto L1215
	}
L1203:
	;
	goto L1204
L1204:
	;
	v5547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5540))))
	if v5547 != 0 {
		goto L1205
	} else {
		goto L1206
	}
L1205:
	;
	v5548 = v5540
	v5549 = v5156
	v5550 = int32(64)
	v5551 = v5547
	goto L1209
L1206:
	;
	v5573 = v5156
	v5577 = int32(0)
	goto L1207
L1207:
	;
	v5578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5573))))
	goto L1201
L1208:
	;
	v5573 = v5568
	v5577 = v5570
	goto L1207
L1209:
	;
	v5553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5549))))
	if v5551 != v5553 {
		v5568 = v5549
		v5570 = v5551
		goto L1208
	} else {
		goto L1211
	}
L1210:
	;
	v5568 = v5562
	v5570 = int32(0)
	goto L1208
L1211:
	;
	if v5553 == int32(0) {
		v5568 = v5549
		v5570 = v5551
		goto L1208
	} else {
		goto L1212
	}
L1212:
	;
	v5558 = v5550 - int32(1)
	if v5558 == int32(0) {
		v5568 = v5549
		v5570 = v5551
		goto L1208
	} else {
		goto L1213
	}
L1213:
	;
	v5561 = int32(1)
	v5562 = v5549 + v5561
	v5563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5548)+1)))
	if v5563 != 0 {
		v5548 = v5548 + v5561
		v5549 = v5562
		v5550 = v5558
		v5551 = v5563
		goto L1209
	} else {
		goto L1214
	}
L1214:
	;
	goto L1210
L1215:
	;
	v5589 = v5515 + int32(1)
	if v5505 != v5589 {
		v5515 = v5589
		goto L1199
	} else {
		goto L1216
	}
L1216:
	;
	goto L1200
L1217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5160))) = v5156
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_129), v5160)
	mBase = m.M
	v5623 = m.ExcPending
	if v5623 != 0 {
		goto L1
	} else {
		goto L1218
	}
L1218:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(821), int32(_a_F_BootstrapModeMain_128))
	mBase = m.M
	v5628 = m.ExcPending
	if v5628 != 0 {
		goto L1
	} else {
		goto L1219
	}
L1219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1220:
	;
	v5671 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5665)+72)))
	if int32(0) <= v5671 {
		v5761 = v5665
		v5762 = v5656
		goto L1140
	} else {
		goto L1221
	}
L1221:
	;
	v5735 = v5665
	goto L1141
L1222:
	;
	v5730 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5761 = v5730
	v5762 = v5699
	goto L1140
L1223:
	;
	v5726 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5682)+80)))
	if int32(0) <= v5726 {
		goto L1222
	} else {
		goto L1224
	}
L1224:
	;
	v5729 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5735 = v5729
	goto L1141
L1225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5783)+96)) = int32(950)
	v5787 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	v5788 = v5787
	goto L1227
L1226:
	;
	v5788 = v5783
	goto L1227
L1227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5788)+76)) = int32(-1)
	v5791 = int32(1)
	v5792 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5792)+92)) = uint8(v5791)
	v5795 = *(*int32)(unsafe.Add(mBase, uint32(v5181)+uint32(_c_F_BootstrapModeMain[61])))
	switch v5157 - int32(2) {
	case 0:
		goto L1231
	case 1:
		v5876 = v5791
		goto L1229
	default:
		goto L1230
	}
L1228:
	;
	m.G0 = v5160 + int32(48)
	v6156 = v3812
	goto L730
L1229:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v5795)+86)) = uint8(v5876)
	goto L1228
L1230:
	;
	v5799 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5795)+72)))
	if v5799 <= int32(0) {
		goto L1228
	} else {
		goto L1232
	}
L1231:
	;
	v5876 = int32(0)
	goto L1229
L1232:
	;
	v5802 = int32(0)
	if v5145 <= v5802 {
		v5850 = v5802
		goto L1233
	} else {
		goto L1234
	}
L1233:
	;
	if v5145 != v5850 {
		goto L1228
	} else {
		goto L1240
	}
L1234:
	;
	v5810 = v5802
	goto L1235
L1235:
	;
	v5834 = *(*int32)(unsafe.Add(mBase, uint32(v5810<<(uint(int32(2))%32))+uint32(_c_F_BootstrapModeMain[61])))
	v5835 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5834)+72)))
	if v5835 <= int32(0) {
		v5850 = v5810
		goto L1233
	} else {
		goto L1237
	}
L1236:
	;
	v5876 = v5841
	goto L1229
L1237:
	;
	v5838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5834)+86)))
	if v5838 != int32(1) {
		v5850 = v5810
		goto L1233
	} else {
		goto L1238
	}
L1238:
	;
	v5841 = int32(1)
	v5843 = v5810 + v5841
	if v5843 != v5145 {
		v5810 = v5843
		goto L1235
	} else {
		goto L1239
	}
L1239:
	;
	goto L1236
L1240:
	;
	v5876 = int32(1)
	goto L1229
L1241:
	;
	v6156 = base.I32_wrap_i64(v5932)
	goto L730
L1242:
	;
	if v5947 != 0 {
		goto L1243
	} else {
		goto L1244
	}
L1243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5943)+20)) = v5934
	*(*int32)(unsafe.Add(mBase, uint32(v5943)+16)) = v5937
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_130), v5943+int32(16))
	mBase = m.M
	v5955 = m.ExcPending
	if v5955 != 0 {
		goto L1
	} else {
		goto L1246
	}
L1244:
	;
	goto L1245
L1245:
	;
	v5962 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	v5963 = *(*int32)(unsafe.Add(mBase, uint32(v5962)+52))
	v5964 = *(*int32)(unsafe.Add(mBase, uint32(v5963)))
	v5971 = *(*int32)(unsafe.Add(mBase, uint32(v5963+v5964<<(uint(int32(4))%32)+v5937*int32(100))+88))
	F_boot_get_type_io_data(m, v5971, v5943+int32(46), v5943+int32(45), v5943+int32(44), v5943+int32(43), v5943+int32(36), v5943+int32(32), v5943+int32(28))
	mBase = m.M
	v5987 = m.ExcPending
	if v5987 != 0 {
		goto L1
	} else {
		goto L1248
	}
L1246:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(670), int32(_a_F_BootstrapModeMain_131))
	mBase = m.M
	v5960 = m.ExcPending
	if v5960 != 0 {
		goto L1
	} else {
		goto L1247
	}
L1247:
	;
	goto L1245
L1248:
	;
	v5989 = v5937 << (uint(int32(2)) % 32)
	v5992 = *(*int32)(unsafe.Add(mBase, uint32(v5943)+32))
	v5993 = *(*int32)(unsafe.Add(mBase, uint32(v5943)+36))
	v5995 = F_OidInputFunctionCall(m, v5992, v5934, v5993, int32(-1))
	mBase = m.M
	v5996 = m.ExcPending
	if v5996 != 0 {
		goto L1
	} else {
		goto L1249
	}
L1249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5989)+uint32(_c_F_BootstrapModeMain[73]))) = v5995
	v6000 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v6001 = m.ExcPending
	if v6001 != 0 {
		goto L1
	} else {
		goto L1250
	}
L1250:
	;
	if v6000 != 0 {
		goto L1251
	} else {
		goto L1252
	}
L1251:
	;
	v6002 = *(*int32)(unsafe.Add(mBase, uint32(v5943)+28))
	v6003 = *(*int32)(unsafe.Add(mBase, uint32(v5989)+uint32(_c_F_BootstrapModeMain[73])))
	v6004 = F_OidOutputFunctionCall(m, v6002, v6003)
	mBase = m.M
	v6005 = m.ExcPending
	if v6005 != 0 {
		goto L1
	} else {
		goto L1254
	}
L1252:
	;
	goto L1253
L1253:
	;
	m.G0 = v5943 + int32(48)
	v6156 = v3812
	goto L730
L1254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5943))) = v6004
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_132), v5943)
	mBase = m.M
	v6009 = m.ExcPending
	if v6009 != 0 {
		goto L1
	} else {
		goto L1255
	}
L1255:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(687), int32(_a_F_BootstrapModeMain_131))
	mBase = m.M
	v6014 = m.ExcPending
	if v6014 != 0 {
		goto L1
	} else {
		goto L1256
	}
L1256:
	;
	goto L1253
L1257:
	;
	if v6030 != 0 {
		goto L1258
	} else {
		goto L1259
	}
L1258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6026)+16)) = v6020
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_133), v6026+int32(16))
	mBase = m.M
	v6037 = m.ExcPending
	if v6037 != 0 {
		goto L1
	} else {
		goto L1261
	}
L1259:
	;
	goto L1260
L1260:
	;
	v6044 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	v6045 = *(*int32)(unsafe.Add(mBase, uint32(v6044)+52))
	v6046 = *(*int32)(unsafe.Add(mBase, uint32(v6045)))
	v6051 = v6020 * int32(100)
	v6053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6045+v6046<<(uint(int32(4))%32)+v6051)+106)))
	if v6053 == int32(1) {
		goto L1263
	} else {
		goto L1264
	}
L1261:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(697), int32(_a_F_BootstrapModeMain_134))
	mBase = m.M
	v6042 = m.ExcPending
	if v6042 != 0 {
		goto L1
	} else {
		goto L1262
	}
L1262:
	;
	goto L1260
L1263:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6059 = m.ExcPending
	if v6059 != 0 {
		goto L1
	} else {
		goto L1266
	}
L1264:
	;
	goto L1265
L1265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6020<<(uint(int32(2))%32))+uint32(_c_F_BootstrapModeMain[73]))) = int32(0)
	v6091 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6020)+uint32(_c_F_BootstrapModeMain[36]))) = uint8(v6091)
	m.G0 = v6026 + int32(32)
	v6156 = v3812
	goto L730
L1266:
	;
	v6061 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	v6062 = *(*int32)(unsafe.Add(mBase, uint32(v6061)+52))
	v6063 = *(*int32)(unsafe.Add(mBase, uint32(v6062)))
	v6064 = *(*int32)(unsafe.Add(mBase, uint32(v6061)+48))
	v6065 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6026)+4)) = v6064 + v6065
	*(*int32)(unsafe.Add(mBase, uint32(v6026))) = v6062 + v6063<<(uint(v6065)%32) + v6051 + int32(24)
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_135), v6026)
	mBase = m.M
	v6077 = m.ExcPending
	if v6077 != 0 {
		goto L1
	} else {
		goto L1267
	}
L1267:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(703), int32(_a_F_BootstrapModeMain_134))
	mBase = m.M
	v6082 = m.ExcPending
	if v6082 != 0 {
		goto L1
	} else {
		goto L1268
	}
L1268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1269:
	;
	v6156 = v6098
	goto L730
L1270:
	;
	v6156 = v6101
	goto L730
L1271:
	;
	v6156 = v6104
	goto L730
L1272:
	;
	v6156 = v6107
	goto L730
L1273:
	;
	v6156 = v6110
	goto L730
L1274:
	;
	v6156 = v6113
	goto L730
L1275:
	;
	v6156 = v6116
	goto L730
L1276:
	;
	v6156 = v6119
	goto L730
L1277:
	;
	v6156 = v6122
	goto L730
L1278:
	;
	v6156 = v6125
	goto L730
L1279:
	;
	v6156 = v6128
	goto L730
L1280:
	;
	v6156 = v6131
	goto L730
L1281:
	;
	v6156 = v6134
	goto L730
L1282:
	;
	v6156 = v6137
	goto L730
L1283:
	;
	v6156 = v6140
	goto L730
L1284:
	;
	v6156 = v6143
	goto L730
L1285:
	;
	v6156 = v6146
	goto L730
L1286:
	;
	v6156 = v6149
	goto L730
L1287:
	;
	v6156 = v6152
	goto L730
L1288:
	;
	v6207 = int32(*(*int8)(unsafe.Add(mBase, uint32(v6191)+uint32(_c_F_BootstrapModeMain[74]))))
	v6211 = v3777
	v6215 = v6183
	v6217 = v3783
	v6218 = v3784
	v6220 = v3786
	v6221 = v6187
	v6222 = v6207
	v6225 = v3791
	v6227 = v3793
	v6228 = v3794
	goto L306
L1289:
	;
	v6200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v6195)+uint32(_c_F_BootstrapModeMain[49]))))
	if v6188 != v6200 {
		goto L1288
	} else {
		goto L1290
	}
L1290:
	;
	v6204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6195)+uint32(_c_F_BootstrapModeMain[50]))))
	v6211 = v3777
	v6215 = v6183
	v6217 = v3783
	v6218 = v3784
	v6220 = v3786
	v6221 = v6187
	v6222 = v6204
	v6225 = v3791
	v6227 = v3793
	v6228 = v3794
	goto L306
L1291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1293:
	;
	v6273 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+32)) = v6273
	v6276 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54]))
	*(*int32)(unsafe.Add(mBase, uint32(v3784)+36)) = v6276
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_136), v3784+int32(32))
	mBase = m.M
	v6282 = m.ExcPending
	if v6282 != 0 {
		goto L1
	} else {
		goto L1294
	}
L1294:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(265), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v6287 = m.ExcPending
	if v6287 != 0 {
		goto L1
	} else {
		goto L1295
	}
L1295:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1296:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_137), int32(0))
	mBase = m.M
	v6295 = m.ExcPending
	if v6295 != 0 {
		goto L1
	} else {
		goto L1297
	}
L1297:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(267), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v6300 = m.ExcPending
	if v6300 != 0 {
		goto L1
	} else {
		goto L1298
	}
L1298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1299:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_138), int32(0))
	mBase = m.M
	v6308 = m.ExcPending
	if v6308 != 0 {
		goto L1
	} else {
		goto L1300
	}
L1300:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_104), int32(446), int32(_a_F_BootstrapModeMain_105))
	mBase = m.M
	v6313 = m.ExcPending
	if v6313 != 0 {
		goto L1
	} else {
		goto L1301
	}
L1301:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1302:
	;
	v6317 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9]))
	v6321 = F_LWLockAcquire(m, v6317+int32(3200), int32(0))
	mBase = m.M
	v6322 = m.ExcPending
	if v6322 != 0 {
		goto L1
	} else {
		goto L1303
	}
L1303:
	;
	v6324 = int32(0)
	F_write_relmap_file(m, int32(_a_F_BootstrapModeMain_139), v6324, v6324, v6324, v6324, int32(1664), int32(_a_F_BootstrapModeMain_140))
	mBase = m.M
	v6331 = m.ExcPending
	if v6331 != 0 {
		goto L1
	} else {
		goto L1304
	}
L1304:
	;
	v6333 = int32(0)
	v6337 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	v6339 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	v6341 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_write_relmap_file(m, int32(_a_F_BootstrapModeMain_141), v6333, v6333, v6333, v6337, v6339, v6341)
	mBase = m.M
	v6343 = m.ExcPending
	if v6343 != 0 {
		goto L1
	} else {
		goto L1305
	}
L1305:
	;
	v6345 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9]))
	F_LWLockRelease(m, v6345+int32(3200))
	mBase = m.M
	v6349 = m.ExcPending
	if v6349 != 0 {
		goto L1
	} else {
		goto L1306
	}
L1306:
	;
	v6351 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[60]))
	if v6351 != 0 {
		goto L1307
	} else {
		goto L1308
	}
L1307:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v6354 = m.ExcPending
	if v6354 != 0 {
		goto L1
	} else {
		goto L1310
	}
L1308:
	;
	goto L1309
L1309:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6357 = m.ExcPending
	if v6357 != 0 {
		goto L1
	} else {
		goto L1311
	}
L1310:
	;
	goto L1309
L1311:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1312:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6364 = m.ExcPending
	if v6364 != 0 {
		goto L1
	} else {
		goto L1313
	}
L1313:
	;
	v6366 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v6366
	F_errmsg(m, int32(_a_F_BootstrapModeMain_142), v28-int32(-64))
	mBase = m.M
	v6372 = m.ExcPending
	if v6372 != 0 {
		goto L1
	} else {
		goto L1314
	}
L1314:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(239), int32(_a_F_BootstrapModeMain_16))
	mBase = m.M
	v6377 = m.ExcPending
	if v6377 != 0 {
		goto L1
	} else {
		goto L1315
	}
L1315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1316:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(254), int32(_a_F_BootstrapModeMain_16))
	mBase = m.M
	v6388 = m.ExcPending
	if v6388 != 0 {
		goto L1
	} else {
		goto L1317
	}
L1317:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1318:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6395 = m.ExcPending
	if v6395 != 0 {
		goto L1
	} else {
		goto L1319
	}
L1319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1320:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6404 = m.ExcPending
	if v6404 != 0 {
		goto L1
	} else {
		goto L1321
	}
L1321:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1322:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1324:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_143), int32(0))
	mBase = m.M
	v6421 = m.ExcPending
	if v6421 != 0 {
		goto L1
	} else {
		goto L1325
	}
L1325:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(383), int32(_a_F_BootstrapModeMain_16))
	mBase = m.M
	v6426 = m.ExcPending
	if v6426 != 0 {
		goto L1
	} else {
		goto L1326
	}
L1326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
