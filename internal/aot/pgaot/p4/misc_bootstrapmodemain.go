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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int64
	_ = v403
	var v405 int64
	_ = v405
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v427 int32
	_ = v427
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int64
	_ = v449
	var v451 int64
	_ = v451
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v473 int32
	_ = v473
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int64
	_ = v495
	var v497 int64
	_ = v497
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v519 int32
	_ = v519
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int64
	_ = v541
	var v543 int64
	_ = v543
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v574 int64
	_ = v574
	var v575 int64
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int64
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
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
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int64
	_ = v627
	var v638 int64
	_ = v638
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v647 int64
	_ = v647
	var v659 int32
	_ = v659
	var v661 int64
	_ = v661
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v827 int64
	_ = v827
	var v829 int64
	_ = v829
	var v831 int64
	_ = v831
	var v833 int64
	_ = v833
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v871 int64
	_ = v871
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1168 int32
	_ = v1168
	var v1174 int64
	_ = v1174
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1274 int32
	_ = v1274
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1522 int32
	_ = v1522
	var v1539 int32
	_ = v1539
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1614 int32
	_ = v1614
	var v1620 int32
	_ = v1620
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1661 int32
	_ = v1661
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1826 int32
	_ = v1826
	var v1830 int32
	_ = v1830
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
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1880 int32
	_ = v1880
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
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
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1978 int32
	_ = v1978
	var v1982 int32
	_ = v1982
	var v1986 int32
	_ = v1986
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2020 int32
	_ = v2020
	var v2024 int32
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2041 int32
	_ = v2041
	var v2045 int32
	_ = v2045
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2087 int32
	_ = v2087
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2112 int32
	_ = v2112
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2125 int32
	_ = v2125
	var v2129 int32
	_ = v2129
	var v2133 int32
	_ = v2133
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2175 int32
	_ = v2175
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2188 int32
	_ = v2188
	var v2192 int32
	_ = v2192
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
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
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2240 int32
	_ = v2240
	var v2245 int32
	_ = v2245
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2262 int32
	_ = v2262
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2282 int32
	_ = v2282
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2301 int32
	_ = v2301
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2393 int32
	_ = v2393
	var v2396 int32
	_ = v2396
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2409 int32
	_ = v2409
	var v2413 int32
	_ = v2413
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2441 int32
	_ = v2441
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2458 int32
	_ = v2458
	var v2471 int32
	_ = v2471
	var v2488 int32
	_ = v2488
	var v2492 int32
	_ = v2492
	var v2494 int32
	_ = v2494
	var v2501 int32
	_ = v2501
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2533 int32
	_ = v2533
	var v2535 int32
	_ = v2535
	var v2540 int32
	_ = v2540
	var v2547 int32
	_ = v2547
	var v2568 int32
	_ = v2568
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2585 int32
	_ = v2585
	var v2589 int32
	_ = v2589
	var v2619 int32
	_ = v2619
	var v2623 int32
	_ = v2623
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2637 int32
	_ = v2637
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2652 int32
	_ = v2652
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2686 int32
	_ = v2686
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2692 int32
	_ = v2692
	var v2693 int32
	_ = v2693
	var v2695 int32
	_ = v2695
	var v2697 int32
	_ = v2697
	var v2703 int32
	_ = v2703
	var v2706 int32
	_ = v2706
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2754 int32
	_ = v2754
	var v2756 int32
	_ = v2756
	var v2761 int32
	_ = v2761
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2794 int32
	_ = v2794
	var v2797 int32
	_ = v2797
	var v2801 int32
	_ = v2801
	var v2804 int32
	_ = v2804
	var v2821 int32
	_ = v2821
	var v2824 int32
	_ = v2824
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2839 int32
	_ = v2839
	var v2846 int32
	_ = v2846
	var v2863 int32
	_ = v2863
	var v2868 int32
	_ = v2868
	var v2870 int32
	_ = v2870
	var v2874 int32
	_ = v2874
	var v2876 int32
	_ = v2876
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2889 int32
	_ = v2889
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2898 int32
	_ = v2898
	var v2903 int32
	_ = v2903
	var v2906 int32
	_ = v2906
	var v2927 int32
	_ = v2927
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2936 int32
	_ = v2936
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2960 int32
	_ = v2960
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2974 int32
	_ = v2974
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v2997 int32
	_ = v2997
	var v3001 int32
	_ = v3001
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3015 int32
	_ = v3015
	var v3016 int32
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3022 int32
	_ = v3022
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3055 int32
	_ = v3055
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3080 int32
	_ = v3080
	var v3083 int32
	_ = v3083
	var v3086 int32
	_ = v3086
	var v3091 int32
	_ = v3091
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3122 int32
	_ = v3122
	var v3144 int32
	_ = v3144
	var v3145 int32
	_ = v3145
	var v3148 int32
	_ = v3148
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3157 int32
	_ = v3157
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3164 int32
	_ = v3164
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3175 int32
	_ = v3175
	var v3176 int32
	_ = v3176
	var v3180 int32
	_ = v3180
	var v3184 int32
	_ = v3184
	var v3186 int32
	_ = v3186
	var v3189 int32
	_ = v3189
	var v3196 int32
	_ = v3196
	var v3198 int32
	_ = v3198
	var v3201 int32
	_ = v3201
	var v3204 int32
	_ = v3204
	var v3205 int32
	_ = v3205
	var v3207 int32
	_ = v3207
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3216 int32
	_ = v3216
	var v3218 int32
	_ = v3218
	var v3221 int32
	_ = v3221
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3230 int32
	_ = v3230
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3250 int32
	_ = v3250
	var v3251 int32
	_ = v3251
	var v3253 int32
	_ = v3253
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3261 int32
	_ = v3261
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3278 int32
	_ = v3278
	var v3279 int32
	_ = v3279
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3288 int32
	_ = v3288
	var v3289 int32
	_ = v3289
	var v3290 int32
	_ = v3290
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3295 int32
	_ = v3295
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3312 int32
	_ = v3312
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3316 int32
	_ = v3316
	var v3318 int32
	_ = v3318
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3331 int32
	_ = v3331
	var v3332 int32
	_ = v3332
	var v3333 int32
	_ = v3333
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3357 int32
	_ = v3357
	var v3365 int32
	_ = v3365
	var v3368 int32
	_ = v3368
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3391 int32
	_ = v3391
	var v3394 int32
	_ = v3394
	var v3398 int32
	_ = v3398
	var v3401 int32
	_ = v3401
	var v3402 int32
	_ = v3402
	var v3407 int32
	_ = v3407
	var v3411 int32
	_ = v3411
	var v3414 int32
	_ = v3414
	var v3415 int32
	_ = v3415
	var v3439 int32
	_ = v3439
	var v3442 int32
	_ = v3442
	var v3443 int32
	_ = v3443
	var v3445 int32
	_ = v3445
	var v3446 int32
	_ = v3446
	var v3450 int32
	_ = v3450
	var v3451 int32
	_ = v3451
	var v3456 int32
	_ = v3456
	var v3469 int32
	_ = v3469
	var v3486 int32
	_ = v3486
	var v3490 int32
	_ = v3490
	var v3492 int32
	_ = v3492
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3496 int32
	_ = v3496
	var v3497 int32
	_ = v3497
	var v3501 int32
	_ = v3501
	var v3502 int32
	_ = v3502
	var v3503 int32
	_ = v3503
	var v3506 int32
	_ = v3506
	var v3523 int32
	_ = v3523
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3537 int32
	_ = v3537
	var v3538 int32
	_ = v3538
	var v3560 int32
	_ = v3560
	var v3561 int32
	_ = v3561
	var v3562 int32
	_ = v3562
	var v3567 int32
	_ = v3567
	var v3571 int32
	_ = v3571
	var v3572 int32
	_ = v3572
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3582 int32
	_ = v3582
	var v3586 int32
	_ = v3586
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3614 int32
	_ = v3614
	var v3617 int32
	_ = v3617
	var v3618 int32
	_ = v3618
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3625 int32
	_ = v3625
	var v3626 int32
	_ = v3626
	var v3631 int32
	_ = v3631
	var v3644 int32
	_ = v3644
	var v3661 int32
	_ = v3661
	var v3665 int32
	_ = v3665
	var v3667 int32
	_ = v3667
	var v3671 int32
	_ = v3671
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3707 int32
	_ = v3707
	var v3711 int32
	_ = v3711
	var v3717 int32
	_ = v3717
	var v3719 int32
	_ = v3719
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
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3736 int32
	_ = v3736
	var v3742 int32
	_ = v3742
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3751 int32
	_ = v3751
	var v3752 int32
	_ = v3752
	var v3755 int32
	_ = v3755
	var v3757 int32
	_ = v3757
	var v3760 int32
	_ = v3760
	var v3773 int32
	_ = v3773
	var v3775 int32
	_ = v3775
	var v3777 int32
	_ = v3777
	var v3780 int32
	_ = v3780
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
	var v3792 int32
	_ = v3792
	var v3798 int32
	_ = v3798
	var v3804 int32
	_ = v3804
	var v3806 int32
	_ = v3806
	var v3808 int32
	_ = v3808
	var v3811 int32
	_ = v3811
	var v3812 int32
	_ = v3812
	var v3814 int32
	_ = v3814
	var v3815 int32
	_ = v3815
	var v3817 int32
	_ = v3817
	var v3818 int32
	_ = v3818
	var v3819 int32
	_ = v3819
	var v3823 int32
	_ = v3823
	var v3830 int32
	_ = v3830
	var v3835 int32
	_ = v3835
	var v3839 int32
	_ = v3839
	var v3844 int32
	_ = v3844
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3852 int32
	_ = v3852
	var v3855 int32
	_ = v3855
	var v3857 int32
	_ = v3857
	var v3859 int32
	_ = v3859
	var v3861 int32
	_ = v3861
	var v3864 int32
	_ = v3864
	var v3867 int32
	_ = v3867
	var v3871 int32
	_ = v3871
	var v3873 int32
	_ = v3873
	var v3876 int32
	_ = v3876
	var v3879 int32
	_ = v3879
	var v3880 int32
	_ = v3880
	var v3888 int32
	_ = v3888
	var v3893 int32
	_ = v3893
	var v3897 int32
	_ = v3897
	var v3898 int32
	_ = v3898
	var v3900 int32
	_ = v3900
	var v3901 int32
	_ = v3901
	var v3904 int32
	_ = v3904
	var v3905 int32
	_ = v3905
	var v3917 int32
	_ = v3917
	var v3938 int32
	_ = v3938
	var v3941 int32
	_ = v3941
	var v3945 int32
	_ = v3945
	var v3947 int32
	_ = v3947
	var v3948 int32
	_ = v3948
	var v3950 int32
	_ = v3950
	var v3952 int32
	_ = v3952
	var v3953 int32
	_ = v3953
	var v3954 int32
	_ = v3954
	var v3958 int32
	_ = v3958
	var v3965 int32
	_ = v3965
	var v3968 int32
	_ = v3968
	var v3969 int32
	_ = v3969
	var v3970 int32
	_ = v3970
	var v3971 int32
	_ = v3971
	var v3972 int32
	_ = v3972
	var v3982 int32
	_ = v3982
	var v3987 int32
	_ = v3987
	var v3991 int32
	_ = v3991
	var v3993 int32
	_ = v3993
	var v4026 int32
	_ = v4026
	var v4029 int32
	_ = v4029
	var v4031 int32
	_ = v4031
	var v4033 int32
	_ = v4033
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4037 int32
	_ = v4037
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4045 int32
	_ = v4045
	var v4047 int32
	_ = v4047
	var v4052 int32
	_ = v4052
	var v4057 int32
	_ = v4057
	var v4058 int32
	_ = v4058
	var v4060 int32
	_ = v4060
	var v4063 int32
	_ = v4063
	var v4065 int32
	_ = v4065
	var v4068 int32
	_ = v4068
	var v4071 int32
	_ = v4071
	var v4073 int32
	_ = v4073
	var v4075 int32
	_ = v4075
	var v4077 int32
	_ = v4077
	var v4078 int32
	_ = v4078
	var v4079 int32
	_ = v4079
	var v4085 int32
	_ = v4085
	var v4086 int32
	_ = v4086
	var v4087 int32
	_ = v4087
	var v4089 int32
	_ = v4089
	var v4094 int32
	_ = v4094
	var v4099 int32
	_ = v4099
	var v4100 int32
	_ = v4100
	var v4102 int32
	_ = v4102
	var v4106 int32
	_ = v4106
	var v4110 int32
	_ = v4110
	var v4111 int32
	_ = v4111
	var v4116 int32
	_ = v4116
	var v4119 int32
	_ = v4119
	var v4122 int64
	_ = v4122
	var v4126 int32
	_ = v4126
	var v4130 int32
	_ = v4130
	var v4134 int32
	_ = v4134
	var v4139 int32
	_ = v4139
	var v4142 int32
	_ = v4142
	var v4145 int32
	_ = v4145
	var v4147 int32
	_ = v4147
	var v4149 int32
	_ = v4149
	var v4151 int32
	_ = v4151
	var v4152 int32
	_ = v4152
	var v4153 int32
	_ = v4153
	var v4159 int32
	_ = v4159
	var v4160 int32
	_ = v4160
	var v4161 int32
	_ = v4161
	var v4163 int32
	_ = v4163
	var v4168 int32
	_ = v4168
	var v4173 int32
	_ = v4173
	var v4174 int32
	_ = v4174
	var v4176 int32
	_ = v4176
	var v4180 int32
	_ = v4180
	var v4182 int32
	_ = v4182
	var v4183 int32
	_ = v4183
	var v4186 int32
	_ = v4186
	var v4189 int32
	_ = v4189
	var v4191 int32
	_ = v4191
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4199 int32
	_ = v4199
	var v4204 int32
	_ = v4204
	var v4207 int32
	_ = v4207
	var v4211 int32
	_ = v4211
	var v4215 int32
	_ = v4215
	var v4218 int32
	_ = v4218
	var v4219 int32
	_ = v4219
	var v4222 int32
	_ = v4222
	var v4225 int32
	_ = v4225
	var v4232 int32
	_ = v4232
	var v4233 int32
	_ = v4233
	var v4237 int32
	_ = v4237
	var v4238 int32
	_ = v4238
	var v4244 int32
	_ = v4244
	var v4249 int32
	_ = v4249
	var v4252 int32
	_ = v4252
	var v4256 int32
	_ = v4256
	var v4259 int32
	_ = v4259
	var v4262 int32
	_ = v4262
	var v4263 int32
	_ = v4263
	var v4270 int32
	_ = v4270
	var v4278 int32
	_ = v4278
	var v4279 int32
	_ = v4279
	var v4282 int32
	_ = v4282
	var v4283 int32
	_ = v4283
	var v4291 int32
	_ = v4291
	var v4296 int32
	_ = v4296
	var v4300 int32
	_ = v4300
	var v4303 int32
	_ = v4303
	var v4305 int32
	_ = v4305
	var v4307 int32
	_ = v4307
	var v4309 int32
	_ = v4309
	var v4310 int32
	_ = v4310
	var v4311 int32
	_ = v4311
	var v4317 int32
	_ = v4317
	var v4318 int32
	_ = v4318
	var v4319 int32
	_ = v4319
	var v4321 int32
	_ = v4321
	var v4326 int32
	_ = v4326
	var v4331 int32
	_ = v4331
	var v4332 int32
	_ = v4332
	var v4334 int32
	_ = v4334
	var v4339 int32
	_ = v4339
	var v4340 int32
	_ = v4340
	var v4344 int32
	_ = v4344
	var v4349 int32
	_ = v4349
	var v4354 int32
	_ = v4354
	var v4356 int32
	_ = v4356
	var v4359 int32
	_ = v4359
	var v4362 int32
	_ = v4362
	var v4364 int32
	_ = v4364
	var v4368 int32
	_ = v4368
	var v4369 int32
	_ = v4369
	var v4371 int32
	_ = v4371
	var v4375 int32
	_ = v4375
	var v4380 int32
	_ = v4380
	var v4382 int32
	_ = v4382
	var v4384 int32
	_ = v4384
	var v4385 int32
	_ = v4385
	var v4388 int32
	_ = v4388
	var v4389 int32
	_ = v4389
	var v4391 int32
	_ = v4391
	var v4393 int32
	_ = v4393
	var v4395 int32
	_ = v4395
	var v4397 int32
	_ = v4397
	var v4400 int32
	_ = v4400
	var v4401 int32
	_ = v4401
	var v4405 int32
	_ = v4405
	var v4410 int32
	_ = v4410
	var v4412 int32
	_ = v4412
	var v4413 int32
	_ = v4413
	var v4428 int32
	_ = v4428
	var v4431 int32
	_ = v4431
	var v4433 int32
	_ = v4433
	var v4435 int32
	_ = v4435
	var v4437 int32
	_ = v4437
	var v4438 int32
	_ = v4438
	var v4439 int32
	_ = v4439
	var v4445 int32
	_ = v4445
	var v4446 int32
	_ = v4446
	var v4447 int32
	_ = v4447
	var v4449 int32
	_ = v4449
	var v4450 int32
	_ = v4450
	var v4455 int32
	_ = v4455
	var v4456 int32
	_ = v4456
	var v4459 int32
	_ = v4459
	var v4465 int32
	_ = v4465
	var v4470 int32
	_ = v4470
	var v4472 int32
	_ = v4472
	var v4477 int32
	_ = v4477
	var v4482 int32
	_ = v4482
	var v4483 int32
	_ = v4483
	var v4485 int32
	_ = v4485
	var v4490 int32
	_ = v4490
	var v4495 int32
	_ = v4495
	var v4497 int32
	_ = v4497
	var v4498 int32
	_ = v4498
	var v4502 int32
	_ = v4502
	var v4503 int32
	_ = v4503
	var v4508 int32
	_ = v4508
	var v4512 int64
	_ = v4512
	var v4532 int32
	_ = v4532
	var v4533 int32
	_ = v4533
	var v4536 int32
	_ = v4536
	var v4537 int32
	_ = v4537
	var v4546 int32
	_ = v4546
	var v4549 int32
	_ = v4549
	var v4552 int32
	_ = v4552
	var v4554 int32
	_ = v4554
	var v4556 int32
	_ = v4556
	var v4558 int32
	_ = v4558
	var v4559 int32
	_ = v4559
	var v4560 int32
	_ = v4560
	var v4566 int32
	_ = v4566
	var v4567 int32
	_ = v4567
	var v4568 int32
	_ = v4568
	var v4570 int32
	_ = v4570
	var v4571 int32
	_ = v4571
	var v4576 int32
	_ = v4576
	var v4577 int32
	_ = v4577
	var v4580 int32
	_ = v4580
	var v4586 int32
	_ = v4586
	var v4591 int32
	_ = v4591
	var v4593 int32
	_ = v4593
	var v4598 int32
	_ = v4598
	var v4603 int32
	_ = v4603
	var v4604 int32
	_ = v4604
	var v4606 int32
	_ = v4606
	var v4611 int32
	_ = v4611
	var v4616 int32
	_ = v4616
	var v4618 int32
	_ = v4618
	var v4619 int32
	_ = v4619
	var v4623 int32
	_ = v4623
	var v4624 int32
	_ = v4624
	var v4629 int32
	_ = v4629
	var v4630 int64
	_ = v4630
	var v4645 int32
	_ = v4645
	var v4655 int32
	_ = v4655
	var v4656 int32
	_ = v4656
	var v4659 int32
	_ = v4659
	var v4660 int32
	_ = v4660
	var v4669 int32
	_ = v4669
	var v4672 int32
	_ = v4672
	var v4675 int32
	_ = v4675
	var v4677 int32
	_ = v4677
	var v4679 int32
	_ = v4679
	var v4681 int32
	_ = v4681
	var v4682 int32
	_ = v4682
	var v4683 int32
	_ = v4683
	var v4689 int32
	_ = v4689
	var v4690 int32
	_ = v4690
	var v4691 int32
	_ = v4691
	var v4694 int32
	_ = v4694
	var v4695 int32
	_ = v4695
	var v4696 int32
	_ = v4696
	var v4702 int32
	_ = v4702
	var v4707 int32
	_ = v4707
	var v4709 int32
	_ = v4709
	var v4714 int32
	_ = v4714
	var v4719 int32
	_ = v4719
	var v4720 int32
	_ = v4720
	var v4722 int32
	_ = v4722
	var v4725 int32
	_ = v4725
	var v4728 int32
	_ = v4728
	var v4731 int32
	_ = v4731
	var v4732 int32
	_ = v4732
	var v4734 int32
	_ = v4734
	var v4738 int32
	_ = v4738
	var v4739 int32
	_ = v4739
	var v4741 int32
	_ = v4741
	var v4742 int32
	_ = v4742
	var v4743 int32
	_ = v4743
	var v4744 int32
	_ = v4744
	var v4750 int32
	_ = v4750
	var v4754 int32
	_ = v4754
	var v4759 int32
	_ = v4759
	var v4760 int32
	_ = v4760
	var v4764 int32
	_ = v4764
	var v4765 int32
	_ = v4765
	var v4771 int32
	_ = v4771
	var v4777 int32
	_ = v4777
	var v4782 int32
	_ = v4782
	var v4785 int32
	_ = v4785
	var v4791 int32
	_ = v4791
	var v4794 int32
	_ = v4794
	var v4796 int32
	_ = v4796
	var v4798 int32
	_ = v4798
	var v4800 int32
	_ = v4800
	var v4801 int32
	_ = v4801
	var v4802 int32
	_ = v4802
	var v4808 int32
	_ = v4808
	var v4809 int32
	_ = v4809
	var v4810 int32
	_ = v4810
	var v4812 int32
	_ = v4812
	var v4817 int32
	_ = v4817
	var v4822 int32
	_ = v4822
	var v4823 int32
	_ = v4823
	var v4825 int32
	_ = v4825
	var v4829 int32
	_ = v4829
	var v4834 int32
	_ = v4834
	var v4856 int32
	_ = v4856
	var v4858 int32
	_ = v4858
	var v4859 int32
	_ = v4859
	var v4861 int32
	_ = v4861
	var v4862 int32
	_ = v4862
	var v4864 int32
	_ = v4864
	var v4865 int32
	_ = v4865
	var v4867 int32
	_ = v4867
	var v4868 int32
	_ = v4868
	var v4869 int32
	_ = v4869
	var v4872 int32
	_ = v4872
	var v4875 int32
	_ = v4875
	var v4878 int32
	_ = v4878
	var v4879 int32
	_ = v4879
	var v4881 int32
	_ = v4881
	var v4882 int32
	_ = v4882
	var v4912 int32
	_ = v4912
	var v4915 int32
	_ = v4915
	var v4917 int32
	_ = v4917
	var v4919 int32
	_ = v4919
	var v4921 int32
	_ = v4921
	var v4922 int32
	_ = v4922
	var v4923 int32
	_ = v4923
	var v4929 int32
	_ = v4929
	var v4930 int32
	_ = v4930
	var v4931 int32
	_ = v4931
	var v4934 int32
	_ = v4934
	var v4935 int32
	_ = v4935
	var v4936 int32
	_ = v4936
	var v4937 int32
	_ = v4937
	var v4938 int32
	_ = v4938
	var v4944 int32
	_ = v4944
	var v4945 int32
	_ = v4945
	var v4947 int32
	_ = v4947
	var v4948 int32
	_ = v4948
	var v4953 int32
	_ = v4953
	var v4959 int32
	_ = v4959
	var v4960 int32
	_ = v4960
	var v4961 int32
	_ = v4961
	var v4967 int32
	_ = v4967
	var v4968 int32
	_ = v4968
	var v4974 int32
	_ = v4974
	var v4975 int32
	_ = v4975
	var v4977 int32
	_ = v4977
	var v4979 int32
	_ = v4979
	var v4985 int32
	_ = v4985
	var v4988 int32
	_ = v4988
	var v4989 int32
	_ = v4989
	var v4990 int32
	_ = v4990
	var v4991 int32
	_ = v4991
	var v4993 int32
	_ = v4993
	var v4996 int32
	_ = v4996
	var v4999 int32
	_ = v4999
	var v5000 int32
	_ = v5000
	var v5004 int32
	_ = v5004
	var v5009 int32
	_ = v5009
	var v5012 int32
	_ = v5012
	var v5014 int32
	_ = v5014
	var v5017 int32
	_ = v5017
	var v5021 int32
	_ = v5021
	var v5023 int32
	_ = v5023
	var v5024 int32
	_ = v5024
	var v5026 int32
	_ = v5026
	var v5027 int32
	_ = v5027
	var v5030 int32
	_ = v5030
	var v5034 int32
	_ = v5034
	var v5039 int32
	_ = v5039
	var v5040 int32
	_ = v5040
	var v5041 int32
	_ = v5041
	var v5050 int32
	_ = v5050
	var v5055 int32
	_ = v5055
	var v5057 int32
	_ = v5057
	var v5059 int32
	_ = v5059
	var v5062 int32
	_ = v5062
	var v5074 int32
	_ = v5074
	var v5094 int32
	_ = v5094
	var v5101 int32
	_ = v5101
	var v5102 int32
	_ = v5102
	var v5103 int32
	_ = v5103
	var v5104 int32
	_ = v5104
	var v5105 int32
	_ = v5105
	var v5107 int32
	_ = v5107
	var v5113 int32
	_ = v5113
	var v5116 int32
	_ = v5116
	var v5117 int32
	_ = v5117
	var v5118 int32
	_ = v5118
	var v5123 int32
	_ = v5123
	var v5125 int32
	_ = v5125
	var v5128 int32
	_ = v5128
	var v5132 int32
	_ = v5132
	var v5133 int32
	_ = v5133
	var v5144 int32
	_ = v5144
	var v5149 int32
	_ = v5149
	var v5150 int32
	_ = v5150
	var v5156 int32
	_ = v5156
	var v5161 int32
	_ = v5161
	var v5163 int32
	_ = v5163
	var v5164 int32
	_ = v5164
	var v5166 int32
	_ = v5166
	var v5177 int32
	_ = v5177
	var v5178 int32
	_ = v5178
	var v5195 int32
	_ = v5195
	var v5198 int32
	_ = v5198
	var v5208 int32
	_ = v5208
	var v5228 int32
	_ = v5228
	var v5230 int32
	_ = v5230
	var v5237 int32
	_ = v5237
	var v5238 int32
	_ = v5238
	var v5239 int32
	_ = v5239
	var v5240 int32
	_ = v5240
	var v5241 int32
	_ = v5241
	var v5243 int32
	_ = v5243
	var v5249 int32
	_ = v5249
	var v5252 int32
	_ = v5252
	var v5253 int32
	_ = v5253
	var v5254 int32
	_ = v5254
	var v5259 int32
	_ = v5259
	var v5261 int32
	_ = v5261
	var v5264 int32
	_ = v5264
	var v5268 int32
	_ = v5268
	var v5269 int32
	_ = v5269
	var v5280 int32
	_ = v5280
	var v5309 int32
	_ = v5309
	var v5314 int32
	_ = v5314
	var v5316 int32
	_ = v5316
	var v5319 int32
	_ = v5319
	var v5322 int32
	_ = v5322
	var v5333 int32
	_ = v5333
	var v5353 int32
	_ = v5353
	var v5355 int32
	_ = v5355
	var v5362 int32
	_ = v5362
	var v5363 int32
	_ = v5363
	var v5364 int32
	_ = v5364
	var v5365 int32
	_ = v5365
	var v5366 int32
	_ = v5366
	var v5368 int32
	_ = v5368
	var v5374 int32
	_ = v5374
	var v5377 int32
	_ = v5377
	var v5378 int32
	_ = v5378
	var v5379 int32
	_ = v5379
	var v5384 int32
	_ = v5384
	var v5386 int32
	_ = v5386
	var v5389 int32
	_ = v5389
	var v5393 int32
	_ = v5393
	var v5394 int32
	_ = v5394
	var v5405 int32
	_ = v5405
	var v5436 int32
	_ = v5436
	var v5440 int32
	_ = v5440
	var v5445 int32
	_ = v5445
	var v5446 int32
	_ = v5446
	var v5448 int32
	_ = v5448
	var v5449 int32
	_ = v5449
	var v5451 int32
	_ = v5451
	var v5452 int32
	_ = v5452
	var v5454 int32
	_ = v5454
	var v5455 int32
	_ = v5455
	var v5457 int32
	_ = v5457
	var v5458 int32
	_ = v5458
	var v5460 int32
	_ = v5460
	var v5461 int32
	_ = v5461
	var v5463 int32
	_ = v5463
	var v5464 int32
	_ = v5464
	var v5467 int32
	_ = v5467
	var v5468 int32
	_ = v5468
	var v5470 int32
	_ = v5470
	var v5475 int32
	_ = v5475
	var v5479 int32
	_ = v5479
	var v5504 int32
	_ = v5504
	var v5507 int32
	_ = v5507
	var v5508 int32
	_ = v5508
	var v5510 int32
	_ = v5510
	var v5511 int32
	_ = v5511
	var v5513 int32
	_ = v5513
	var v5514 int32
	_ = v5514
	var v5516 int32
	_ = v5516
	var v5517 int32
	_ = v5517
	var v5519 int32
	_ = v5519
	var v5520 int32
	_ = v5520
	var v5522 int32
	_ = v5522
	var v5525 int32
	_ = v5525
	var v5526 int32
	_ = v5526
	var v5528 int32
	_ = v5528
	var v5531 int32
	_ = v5531
	var v5534 int32
	_ = v5534
	var v5535 int32
	_ = v5535
	var v5537 int32
	_ = v5537
	var v5564 int32
	_ = v5564
	var v5567 int32
	_ = v5567
	var v5590 int32
	_ = v5590
	var v5591 int32
	_ = v5591
	var v5594 int32
	_ = v5594
	var v5595 int32
	_ = v5595
	var v5598 int32
	_ = v5598
	var v5599 int32
	_ = v5599
	var v5602 int32
	_ = v5602
	var v5606 int32
	_ = v5606
	var v5609 int32
	_ = v5609
	var v5621 int32
	_ = v5621
	var v5640 int32
	_ = v5640
	var v5641 int32
	_ = v5641
	var v5644 int32
	_ = v5644
	var v5648 int32
	_ = v5648
	var v5659 int32
	_ = v5659
	var v5681 int32
	_ = v5681
	var v5735 int32
	_ = v5735
	var v5739 int64
	_ = v5739
	var v5741 int32
	_ = v5741
	var v5742 int32
	_ = v5742
	var v5744 int32
	_ = v5744
	var v5748 int32
	_ = v5748
	var v5750 int32
	_ = v5750
	var v5754 int32
	_ = v5754
	var v5755 int32
	_ = v5755
	var v5762 int32
	_ = v5762
	var v5767 int32
	_ = v5767
	var v5769 int32
	_ = v5769
	var v5770 int32
	_ = v5770
	var v5771 int32
	_ = v5771
	var v5778 int32
	_ = v5778
	var v5794 int32
	_ = v5794
	var v5796 int32
	_ = v5796
	var v5799 int32
	_ = v5799
	var v5800 int32
	_ = v5800
	var v5802 int32
	_ = v5802
	var v5803 int32
	_ = v5803
	var v5807 int32
	_ = v5807
	var v5808 int32
	_ = v5808
	var v5809 int32
	_ = v5809
	var v5810 int32
	_ = v5810
	var v5811 int32
	_ = v5811
	var v5812 int32
	_ = v5812
	var v5816 int32
	_ = v5816
	var v5821 int32
	_ = v5821
	var v5825 int32
	_ = v5825
	var v5827 int32
	_ = v5827
	var v5831 int32
	_ = v5831
	var v5833 int32
	_ = v5833
	var v5837 int32
	_ = v5837
	var v5838 int32
	_ = v5838
	var v5844 int32
	_ = v5844
	var v5849 int32
	_ = v5849
	var v5851 int32
	_ = v5851
	var v5853 int32
	_ = v5853
	var v5854 int32
	_ = v5854
	var v5855 int32
	_ = v5855
	var v5860 int32
	_ = v5860
	var v5866 int32
	_ = v5866
	var v5868 int32
	_ = v5868
	var v5869 int32
	_ = v5869
	var v5870 int32
	_ = v5870
	var v5871 int32
	_ = v5871
	var v5872 int32
	_ = v5872
	var v5884 int32
	_ = v5884
	var v5889 int32
	_ = v5889
	var v5898 int32
	_ = v5898
	var v5903 int32
	_ = v5903
	var v5904 int32
	_ = v5904
	var v5905 int32
	_ = v5905
	var v5906 int32
	_ = v5906
	var v5907 int32
	_ = v5907
	var v5908 int32
	_ = v5908
	var v5909 int32
	_ = v5909
	var v5910 int32
	_ = v5910
	var v5911 int32
	_ = v5911
	var v5912 int32
	_ = v5912
	var v5913 int32
	_ = v5913
	var v5914 int32
	_ = v5914
	var v5915 int32
	_ = v5915
	var v5916 int32
	_ = v5916
	var v5917 int32
	_ = v5917
	var v5918 int32
	_ = v5918
	var v5919 int32
	_ = v5919
	var v5920 int32
	_ = v5920
	var v5921 int32
	_ = v5921
	var v5922 int32
	_ = v5922
	var v5923 int32
	_ = v5923
	var v5924 int32
	_ = v5924
	var v5925 int32
	_ = v5925
	var v5926 int32
	_ = v5926
	var v5927 int32
	_ = v5927
	var v5928 int32
	_ = v5928
	var v5929 int32
	_ = v5929
	var v5930 int32
	_ = v5930
	var v5931 int32
	_ = v5931
	var v5932 int32
	_ = v5932
	var v5933 int32
	_ = v5933
	var v5934 int32
	_ = v5934
	var v5935 int32
	_ = v5935
	var v5936 int32
	_ = v5936
	var v5937 int32
	_ = v5937
	var v5938 int32
	_ = v5938
	var v5939 int32
	_ = v5939
	var v5940 int32
	_ = v5940
	var v5941 int32
	_ = v5941
	var v5942 int32
	_ = v5942
	var v5943 int32
	_ = v5943
	var v5944 int32
	_ = v5944
	var v5945 int32
	_ = v5945
	var v5946 int32
	_ = v5946
	var v5947 int32
	_ = v5947
	var v5948 int32
	_ = v5948
	var v5949 int32
	_ = v5949
	var v5950 int32
	_ = v5950
	var v5951 int32
	_ = v5951
	var v5952 int32
	_ = v5952
	var v5953 int32
	_ = v5953
	var v5954 int32
	_ = v5954
	var v5955 int32
	_ = v5955
	var v5956 int32
	_ = v5956
	var v5957 int32
	_ = v5957
	var v5958 int32
	_ = v5958
	var v5959 int32
	_ = v5959
	var v5960 int32
	_ = v5960
	var v5963 int32
	_ = v5963
	var v5989 int32
	_ = v5989
	var v5992 int32
	_ = v5992
	var v5993 int32
	_ = v5993
	var v5994 int32
	_ = v5994
	var v5997 int32
	_ = v5997
	var v6000 int32
	_ = v6000
	var v6001 int32
	_ = v6001
	var v6004 int32
	_ = v6004
	var v6010 int32
	_ = v6010
	var v6013 int32
	_ = v6013
	var v6017 int32
	_ = v6017
	var v6019 int32
	_ = v6019
	var v6021 int32
	_ = v6021
	var v6024 int32
	_ = v6024
	var v6025 int32
	_ = v6025
	var v6027 int32
	_ = v6027
	var v6028 int32
	_ = v6028
	var v6030 int32
	_ = v6030
	var v6031 int32
	_ = v6031
	var v6032 int32
	_ = v6032
	var v6036 int32
	_ = v6036
	var v6044 int32
	_ = v6044
	var v6057 int32
	_ = v6057
	var v6058 int32
	_ = v6058
	var v6061 int32
	_ = v6061
	var v6087 int32
	_ = v6087
	var v6106 int32
	_ = v6106
	var v6110 int32
	_ = v6110
	var v6112 int32
	_ = v6112
	var v6115 int32
	_ = v6115
	var v6121 int32
	_ = v6121
	var v6126 int32
	_ = v6126
	var v6130 int32
	_ = v6130
	var v6134 int32
	_ = v6134
	var v6139 int32
	_ = v6139
	var v6143 int32
	_ = v6143
	var v6147 int32
	_ = v6147
	var v6152 int32
	_ = v6152
	var v6163 int32
	_ = v6163
	var v6164 int32
	_ = v6164
	var v6179 int32
	_ = v6179
	var v6189 int32
	_ = v6189
	var v6209 int32
	_ = v6209
	var v6214 int32
	_ = v6214
	var v6216 int32
	_ = v6216
	var v6220 int32
	_ = v6220
	var v6221 int32
	_ = v6221
	var v6223 int32
	_ = v6223
	var v6230 int32
	_ = v6230
	var v6232 int32
	_ = v6232
	var v6236 int32
	_ = v6236
	var v6238 int32
	_ = v6238
	var v6240 int32
	_ = v6240
	var v6242 int32
	_ = v6242
	var v6244 int32
	_ = v6244
	var v6248 int32
	_ = v6248
	var v6250 int32
	_ = v6250
	var v6253 int32
	_ = v6253
	var v6256 int32
	_ = v6256
	var v6260 int32
	_ = v6260
	var v6263 int32
	_ = v6263
	var v6265 int32
	_ = v6265
	var v6271 int32
	_ = v6271
	var v6276 int32
	_ = v6276
	var v6282 int32
	_ = v6282
	var v6287 int32
	_ = v6287
	var v6291 int32
	_ = v6291
	var v6294 int32
	_ = v6294
	var v6300 int32
	_ = v6300
	var v6303 int32
	_ = v6303
	var v6306 int32
	_ = v6306
	var v6312 int32
	_ = v6312
	var v6316 int32
	_ = v6316
	var v6320 int32
	_ = v6320
	var v6325 int32
	_ = v6325
	v4 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(96)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_InitStandaloneProcess(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v35 = l0 - int32(1)
	F_InitializeGUCOptions(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v42 = int32(0)
	v44 = v4
	goto L11
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6316 = m.ExcPending
	if v6316 != 0 {
		goto L1
	} else {
		goto L1222
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[0])) = int32(2)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6312 = m.ExcPending
	if v6312 != 0 {
		goto L1
	} else {
		goto L1221
	}
L6:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6306 = m.ExcPending
	if v6306 != 0 {
		goto L1
	} else {
		goto L1220
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v31
	F_write_stderr(m, int32(_a_F_BootstrapModeMain_0), v29+int32(16))
	mBase = m.M
	v6300 = m.ExcPending
	if v6300 != 0 {
		goto L1
	} else {
		goto L1218
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
	F_write_stderr(m, int32(_a_F_BootstrapModeMain_1), v29)
	mBase = m.M
	v6291 = m.ExcPending
	if v6291 != 0 {
		goto L1
	} else {
		goto L1216
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v124
	F_errmsg(m, int32(_a_F_BootstrapModeMain_2), v29+int32(32))
	mBase = m.M
	v6282 = m.ExcPending
	if v6282 != 0 {
		goto L1
	} else {
		goto L1214
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6260 = m.ExcPending
	if v6260 != 0 {
		goto L1
	} else {
		goto L1210
	}
L11:
	;
	v69 = F_getopt(m, v35, l1+int32(4), int32(_a_F_BootstrapModeMain_3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L12:
	;
	if v69 != int32(-1) {
		goto L8
	} else {
		goto L88
	}
L13:
	;
	goto L12
L14:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_4), v303, int32(0), int32(1))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L87
	}
L15:
	;
	v179 = int32(_a_F_BootstrapModeMain_5)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	goto L59
L16:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_6), int32(_a_F_BootstrapModeMain_7), int32(1), int32(4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L55
	}
L17:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v154
	v160 = F_psprintf(m, int32(_a_F_BootstrapModeMain_8), v29+int32(80))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L51
	}
L18:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	v151 = F_pstrdup(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L50
	}
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	F_ParseLongOption(m, v106, v29+int32(92), v29+int32(88))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	v83 = F_strcmp(m, int32(_a_F_BootstrapModeMain_9), v81)
	mBase = m.M
	if v83 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_10), v75, int32(1), int32(4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	switch v69 - int32(45) {
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
		v42 = int32(1)
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
	if v102 != int32(5) {
		goto L10
	} else {
		goto L37
	}
L25:
	;
	v102 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v88 = F_strcmp(m, int32(_a_F_BootstrapModeMain_11), v81)
	mBase = m.M
	if v88 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v102 = int32(1)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v93 = F_strcmp(m, int32(_a_F_BootstrapModeMain_12), v81)
	mBase = m.M
	if v93 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v102 = int32(3)
	goto L24
L32:
	;
	goto L33
L33:
	;
	v100 = F_strcmp(m, int32(_a_F_BootstrapModeMain_13), v81)
	mBase = m.M
	if v100 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v101 = int32(5)
	goto L36
L35:
	;
	v101 = int32(4)
	goto L36
L36:
	;
	v102 = v101
	goto L24
L37:
	;
	goto L19
L38:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	if v113 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	F_SetConfigOption(m, v138, v113, int32(1), int32(4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L47
	}
L42:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	if v69 == int32(45) {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v124
	F_errmsg(m, int32(_a_F_BootstrapModeMain_14), v29+int32(48))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(259), int32(_a_F_BootstrapModeMain_16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
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
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	F_pfree(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	F_pfree(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L11
L50:
	;
	v44 = v151
	goto L11
L51:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_17), v160, int32(1), int32(4))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_18), v160, int32(1), int32(4))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_pfree(m, v160)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
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
	v298 = F_strlen(m, v287)
	mBase = m.M
	goto L56
L59:
	;
	goto L60
L60:
	;
	v188 = int32(1023)
	if (v179^v181)&int32(3) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v291 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v288))) = uint8(v291)
	goto L57
L62:
	;
	v272 = v267
	v273 = v268
	v274 = v269
	goto L83
L63:
	;
	if v262 == int32(0) {
		v287 = v260
		v288 = v261
		goto L61
	} else {
		goto L82
	}
L64:
	;
	v260 = v181
	v261 = v179
	v262 = v188
	goto L63
L65:
	;
	goto L66
L66:
	;
	v192 = int32(0)
	if base.B2i32(v181&int32(3) == v192)|int32(0) == v192 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v228 == int32(0) {
		v287 = v225
		v288 = v226
		goto L61
	} else {
		goto L76
	}
L68:
	;
	v204 = v181
	v205 = v179
	v206 = v188
	goto L71
L69:
	;
	goto L70
L70:
	;
	v225 = v181
	v226 = v179
	v227 = v188
	v228 = int32(1)
	goto L67
L71:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v208)
	if v208 == int32(0) {
		v267 = v204
		v268 = v205
		v269 = v206
		goto L62
	} else {
		goto L73
	}
L72:
	;
	v225 = v219
	v226 = v213
	v227 = v215
	v228 = v217
	goto L67
L73:
	;
	v212 = int32(1)
	v213 = v205 + v212
	v215 = v206 - v212
	v216 = int32(0)
	v217 = base.B2i32(v215 != v216)
	v219 = v204 + v212
	if v219&int32(3) == v216 {
		v225 = v219
		v226 = v213
		v227 = v215
		v228 = v217
		goto L67
	} else {
		goto L74
	}
L74:
	;
	if v215 != 0 {
		v204 = v219
		v205 = v213
		v206 = v215
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if base.B2i32(v231 == int32(0))|base.B2i32(base.Ui32(v227) < base.Ui32(int32(4))) != 0 {
		v260 = v225
		v261 = v226
		v262 = v227
		goto L63
	} else {
		goto L77
	}
L77:
	;
	v238 = v225
	v239 = v226
	v240 = v227
	goto L78
L78:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v246 = int32(-2139062144)
	if (int32(16843008)-v243|v243)&v246 != v246 {
		v267 = v238
		v268 = v239
		v269 = v240
		goto L62
	} else {
		goto L80
	}
L79:
	;
	v260 = v254
	v261 = v252
	v262 = v256
	goto L63
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v243
	v251 = int32(4)
	v252 = v239 + v251
	v254 = v238 + v251
	v256 = v240 - v251
	if base.Ui32(int32(3)) < base.Ui32(v256) {
		v238 = v254
		v239 = v252
		v240 = v256
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v267 = v260
	v268 = v261
	v269 = v262
	goto L62
L83:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v276)
	if v276 == int32(0) {
		v287 = v272
		v288 = v273
		goto L61
	} else {
		goto L85
	}
L84:
	;
	v287 = v283
	v288 = v281
	goto L61
L85:
	;
	v280 = int32(1)
	v281 = v273 + v280
	v283 = v272 + v280
	v285 = v274 - v280
	if v285 != 0 {
		v272 = v283
		v273 = v281
		v274 = v285
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	goto L11
L88:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[2]))
	if v35 != v311 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	v313 = F_SelectConfigFiles(m, v44, v31)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if v313 == int32(0) {
		goto L6
	} else {
		goto L91
	}
L91:
	;
	F_checkDataDir(m)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_ChangeToDataDir(m)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_CreateDataDirLockFile(m, int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v325 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[3])) = uint8(v325)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[0])) = int32(0)
	F_InitializeMaxBackends(m)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_InitPostmasterChildSlots(m)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v338 = int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[4]))
	if v341&(v341-v338) != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	F_CreateSharedMemoryAndSemaphores(m)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L107
	}
L98:
	;
	v348 = v338 << (uint(int32(32)-base.I32_clz(v341)) % 32)
	goto L100
L99:
	;
	v348 = v341
	goto L100
L100:
	;
	if base.Ui32(v348) <= base.Ui32(int32(31)) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v351 = int32(31)
	goto L103
L102:
	;
	v351 = v348
	goto L103
L103:
	;
	if base.Ui32(int32(_a_F_BootstrapModeMain_19)) <= base.Ui32(v348) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v356 = int32(1024)
	goto L106
L105:
	;
	v356 = int32(base.Ui32(v351) >> (uint(int32(4)) % 32))
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[5])) = v356
	goto L97
L107:
	;
	F_set_max_safe_fds(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	F_InitProcess(m)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_BaseInit(m)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v367 = int32(0)
	v369 = m.G0
	v371 = v369 - int32(32)
	m.G0 = v371
	switch int32(2) {
	case 0, 2:
		v381 = v367
		goto L113
	default:
		goto L114
	}
L112:
	;
	v413 = int32(0)
	v415 = m.G0
	v417 = v415 - int32(32)
	m.G0 = v417
	switch int32(2) {
	case 0, 2:
		v427 = v413
		goto L126
	default:
		goto L127
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+12)) = v381
	F_sigemptyset(m, v371+int32(16))
	mBase = m.M
	goto L116
L114:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[6])) = v367
	v381 = int32(_a_F_BootstrapModeMain_20)
	goto L113
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+24)) = int32(268435456)
	v393 = v371 + int32(12)
	goto L120
L118:
	;
	m.G0 = v371 + int32(32)
	goto L112
L120:
	;
	goto L121
L121:
	;
	if v393 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v399 = int32(20)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v393)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[7])) = v401
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v393)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[8])) = v403
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v393)))
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9])) = v405
	goto L124
L123:
	;
	goto L124
L124:
	;
	goto L118
L125:
	;
	v459 = int32(0)
	v461 = m.G0
	v463 = v461 - int32(32)
	m.G0 = v463
	switch int32(2) {
	case 0, 2:
		v473 = v459
		goto L139
	default:
		goto L140
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417)+12)) = v427
	F_sigemptyset(m, v417+int32(16))
	mBase = m.M
	goto L129
L127:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[10])) = v413
	v427 = int32(_a_F_BootstrapModeMain_20)
	goto L126
L129:
	;
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417)+24)) = int32(268435456)
	v439 = v417 + int32(12)
	goto L133
L131:
	;
	m.G0 = v417 + int32(32)
	goto L125
L133:
	;
	goto L134
L134:
	;
	if v439 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v446 = int32(40)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v439)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[11])) = v447
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v439)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[12])) = v449
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v439)))
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[13])) = v451
	goto L137
L136:
	;
	goto L137
L137:
	;
	goto L131
L138:
	;
	v505 = int32(0)
	v507 = m.G0
	v509 = v507 - int32(32)
	m.G0 = v509
	switch int32(2) {
	case 0, 2:
		v519 = v505
		goto L152
	default:
		goto L153
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v463)+12)) = v473
	F_sigemptyset(m, v463+int32(16))
	mBase = m.M
	goto L142
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[14])) = v459
	v473 = int32(_a_F_BootstrapModeMain_20)
	goto L139
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v463)+24)) = int32(268435456)
	v485 = v463 + int32(12)
	goto L146
L144:
	;
	m.G0 = v463 + int32(32)
	goto L138
L146:
	;
	goto L147
L147:
	;
	if v485 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v492 = int32(300)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v485)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[15])) = v493
	v495 = *(*int64)(unsafe.Add(mBase, uint32(v485)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[16])) = v495
	v497 = *(*int64)(unsafe.Add(mBase, uint32(v485)))
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[17])) = v497
	goto L150
L149:
	;
	goto L150
L150:
	;
	goto L144
L151:
	;
	v550 = m.G0
	v552 = v550 - int32(_a_F_BootstrapModeMain_21)
	m.G0 = v552
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	v559 = F_LWLockAcquire(m, v555+int32(1152), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L164
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v509)+12)) = v519
	F_sigemptyset(m, v509+int32(16))
	mBase = m.M
	goto L155
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[19])) = v505
	v519 = int32(_a_F_BootstrapModeMain_20)
	goto L152
L155:
	;
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v509)+24)) = int32(268435456)
	v531 = v509 + int32(12)
	goto L159
L157:
	;
	m.G0 = v509 + int32(32)
	goto L151
L159:
	;
	goto L160
L160:
	;
	if v531 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v538 = int32(60)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v531)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[20])) = v539
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v531)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[21])) = v541
	v543 = *(*int64)(unsafe.Add(mBase, uint32(v531)))
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[22])) = v543
	goto L163
L162:
	;
	goto L163
L163:
	;
	goto L157
L164:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[23]))
	v563 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v562)+320)) = uint8(v563)
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	F_LWLockRelease(m, v566+int32(1152))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_gettimeofday(m, v552-int32(-64))
	mBase = m.M
	v574 = int64(*(*int32)(unsafe.Add(mBase, uint32(v552)+72)))
	v575 = *(*int64)(unsafe.Add(mBase, uint32(v552)+64))
	v577 = F_palloc(m, int32(_a_F_BootstrapModeMain_19))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v582 = (v577 + int32(_a_F_BootstrapModeMain_22)) & int32(-8192)
	v583 = int32(0)
	base.MemoryFill(m, v582, v583, int32(_a_F_BootstrapModeMain_23))
	v587 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[24]))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[25])))
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[26]))
	v592 = F_time(m)
	mBase = m.M
	v593 = int32(_a_F_BootstrapModeMain_24)
	v594 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v594))) = int32(_a_F_BootstrapModeMain_25)
	*(*int64)(unsafe.Add(mBase, uint32(v594)+8)) = int64(3)
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v600)+4)) = v583
	F_MultiXactSetNextMXact(m, int32(1), v583)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_AdvanceOldestClogXid(m, int32(3))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_SetTransactionIdLimit(m, int32(3), int32(1))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v614 = int32(1)
	F_SetMultiXactIdLimit(m, v614, v614, v614)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v619 = int32(0)
	F_SetCommitTsLimit(m, v619, v619)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v582))) = int64(4295151896)
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[24]))
	v627 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v582)+48)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v582)+36)) = int32(_a_F_BootstrapModeMain_23)
	*(*int32)(unsafe.Add(mBase, uint32(v582)+32)) = v626
	v638 = v574<<(uint(int64(12))%64) | v575<<(uint(int64(32))%64) | int64(42)
	*(*int64)(unsafe.Add(mBase, uint32(v582)+24)) = v638
	v640 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v582)+146)) = v640
	*(*int64)(unsafe.Add(mBase, uint32(v582)+138)) = v627
	*(*int64)(unsafe.Add(mBase, uint32(v582)+130)) = v592
	v645 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v582)+122)) = v645
	v647 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v582)+114)) = v647
	*(*int64)(unsafe.Add(mBase, uint32(v582)+106)) = int64(12884901888)
	*(*int64)(unsafe.Add(mBase, uint32(v582)+98)) = int64(4294977296)
	*(*int64)(unsafe.Add(mBase, uint32(v582)+90)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v582)+86)) = v591
	*(*uint8)(unsafe.Add(mBase, uint32(v582)+82)) = uint8(v589)
	*(*int64)(unsafe.Add(mBase, uint32(v582)+74)) = v647
	v659 = int32(40)
	v661 = base.I64_extend_i32_u(v587 + v659)
	*(*int64)(unsafe.Add(mBase, uint32(v582)+66)) = v661
	v663 = int32(_a_F_BootstrapModeMain_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v582)+64)) = uint16(v663)
	*(*uint16)(unsafe.Add(mBase, uint32(v582)+56)) = uint16(v640)
	*(*int64)(unsafe.Add(mBase, uint32(v582)+40)) = int64(114)
	*(*int64)(unsafe.Add(mBase, uint32(v582)+8)) = base.I64_extend_i32_s(v626)
	v671 = int32(-1)
	v675 = m.Env.Pgmem_crc32c(m, v671, v582|int32(64), int32(90))
	mBase = m.M
	v679 = m.Env.Pgmem_crc32c(m, v675, v582|v659, int32(20))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v582)+60)) = v679 ^ v671
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[28])) = v645
	v689 = F_XLogFileInit(m, int64(1), v645)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[29])) = v689
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30])) = int32(0)
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v696))) = int32(167772229)
	v699 = int32(_a_F_BootstrapModeMain_23)
	v700 = F_write(m, v689, v582, v699)
	mBase = m.M
	if v700 != v699 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v704 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	if v704 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L175
L175:
	;
	v725 = int32(_a_F_BootstrapModeMain_27)
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	v727 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v726))) = v727
	v730 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = int32(167772228)
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[29]))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[32])))
	if v737 != int32(1) {
		v751 = v727
		goto L190
	} else {
		goto L191
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30])) = int32(51)
	goto L178
L177:
	;
	goto L178
L178:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_28), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_29), int32(_a_F_BootstrapModeMain_30), int32(_a_F_BootstrapModeMain_31))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	v1161 = int32(0)
	F_InitPostgres(m, v1161, v1161, v1161, v1161, v1161, v1161)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L277
	}
L184:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L273
	}
L185:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L269
	}
L186:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L265
	}
L187:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L261
	}
L188:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L1
	} else {
		goto L257
	}
L189:
	;
	if v751 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L190:
	;
	goto L189
L191:
	;
	goto L192
L192:
	;
	v742 = F_fsync(m, v734)
	mBase = m.M
	if v742 != int32(-1) {
		v751 = v742
		goto L190
	} else {
		goto L194
	}
L193:
	;
	v751 = int32(-1)
	goto L190
L194:
	;
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	if v746 == int32(27) {
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = int32(0)
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[29]))
	v760 = F_close(m, v759)
	mBase = m.M
	if v760 != 0 {
		goto L188
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L253
	}
L199:
	;
	v762 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[29])) = v762
	v765 = v552 + int32(80)
	v767 = int32(0)
	v771 = m.G0
	v773 = v771 - int32(16)
	m.G0 = v773
	*(*int32)(unsafe.Add(mBase, uint32(v773))) = v767
	v779 = F_open(m, int32(_a_F_BootstrapModeMain_32), v767, v773)
	mBase = m.M
	if v779 != v762 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	if v812 == int32(0) {
		goto L187
	} else {
		goto L213
	}
L201:
	;
	goto L205
L202:
	;
	v812 = v767
	goto L203
L203:
	;
	m.G0 = v773 + int32(16)
	goto L200
L204:
	;
	v807 = F_close(m, v779)
	mBase = m.M
	v812 = v805
	goto L203
L205:
	;
	v785 = v765
	v786 = int32(32)
	goto L206
L206:
	;
	v791 = F_read(m, v779, v785, v786)
	mBase = m.M
	if v791 <= int32(0) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v805 = int32(1)
	goto L204
L208:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	if v795 == int32(27) {
		goto L206
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v800 = v786 - v791
	if v800 != 0 {
		v785 = v785 + v791
		v786 = v800
		goto L206
	} else {
		goto L212
	}
L211:
	;
	v805 = int32(0)
	goto L204
L212:
	;
	goto L207
L213:
	;
	v819 = int32(_a_F_BootstrapModeMain_33)
	v820 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[33]))
	v821 = int32(8)
	v823 = int32(0)
	base.MemoryFill(m, v820+v821, v823, int32(288))
	*(*int64)(unsafe.Add(mBase, uint32(v820))) = v638
	v827 = *(*int64)(unsafe.Add(mBase, uint32(v552)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v820)+257)) = v827
	v829 = *(*int64)(unsafe.Add(mBase, uint32(v552)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v820)+265)) = v829
	v831 = *(*int64)(unsafe.Add(mBase, uint32(v552)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v820)+273)) = v831
	v833 = *(*int64)(unsafe.Add(mBase, uint32(v552)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v820)+281)) = v833
	*(*int64)(unsafe.Add(mBase, uint32(v820)+128)) = int64(1000)
	v837 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+16)) = v837
	v840 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[34]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+180)) = v840
	v843 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+184)) = v843
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+188)) = v846
	v849 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[37]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+192)) = v849
	v852 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+196)) = v852
	v855 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[26]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+172)) = v855
	v858 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[38])))
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+176)) = uint8(v858)
	v861 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39])))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+252)) = v42
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+200)) = uint8(v861)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+120)) = v823
	*(*int64)(unsafe.Add(mBase, uint32(v820)+112)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+104)) = v592
	*(*int32)(unsafe.Add(mBase, uint32(v820)+96)) = v837
	v871 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+88)) = v871
	*(*int64)(unsafe.Add(mBase, uint32(v820)+80)) = int64(12884901888)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+72)) = int64(4294977296)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+64)) = int64(3)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+60)) = v591
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+56)) = uint8(v589)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+48)) = v871
	*(*int64)(unsafe.Add(mBase, uint32(v820)+40)) = v661
	*(*int64)(unsafe.Add(mBase, uint32(v820)+32)) = v661
	*(*int64)(unsafe.Add(mBase, uint32(v820)+24)) = v592
	*(*int32)(unsafe.Add(mBase, uint32(v820)+224)) = int32(_a_F_BootstrapModeMain_23)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+216)) = int64(562949953429504)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+208)) = int64(4698053236609777664)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+204)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v820)+8)) = int64(869757897079260936)
	v897 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[24]))
	v898 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+292)) = v898
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+256)) = uint8(v837)
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+248)) = uint8(v823)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+240)) = int64(8796093024204)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+232)) = int64(137438953536)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+228)) = v897
	v911 = m.Env.Pgmem_crc32c(m, v898, v820, int32(292))
	mBase = m.M
	v913 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[33]))
	*(*int32)(unsafe.Add(mBase, uint32(v913)+292)) = v911 ^ v898
	base.MemoryFill(m, v552+int32(376), v823, int32(_a_F_BootstrapModeMain_34))
	base.MemoryCopy(m, v765, v913, int32(296))
	v926 = F_BasicOpenFile(m, int32(_a_F_BootstrapModeMain_35), int32(194))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	if v926 < int32(0) {
		goto L186
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30])) = int32(0)
	v934 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v934))) = int32(167772172)
	v937 = int32(_a_F_BootstrapModeMain_23)
	v938 = F_write(m, v926, v765, v937)
	mBase = m.M
	if v938 != v937 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v942 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	if v942 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	goto L218
L218:
	;
	v966 = int32(_a_F_BootstrapModeMain_27)
	v967 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	v968 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v967))) = v968
	v971 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v971))) = int32(167772170)
	v976 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[32])))
	if v976 != int32(1) {
		v990 = v968
		goto L227
	} else {
		goto L228
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30])) = int32(51)
	goto L221
L220:
	;
	goto L221
L221:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+48)) = int32(_a_F_BootstrapModeMain_35)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_36), v552+int32(48))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_29), int32(_a_F_BootstrapModeMain_37), int32(_a_F_BootstrapModeMain_38))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	if v990 != 0 {
		goto L185
	} else {
		goto L233
	}
L227:
	;
	goto L226
L228:
	;
	goto L229
L229:
	;
	v981 = F_fsync(m, v926)
	mBase = m.M
	if v981 != int32(-1) {
		v990 = v981
		goto L227
	} else {
		goto L231
	}
L230:
	;
	v990 = int32(-1)
	goto L227
L231:
	;
	v985 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	if v985 == int32(27) {
		goto L229
	} else {
		goto L232
	}
L232:
	;
	goto L230
L233:
	;
	v992 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v992))) = int32(0)
	v995 = F_close(m, v926)
	mBase = m.M
	if v995 != 0 {
		goto L184
	} else {
		goto L234
	}
L234:
	;
	v997 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v997)+28))
	v1000 = F_LWLockAcquire(m, v998, int32(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	v1002 = int32(_a_F_BootstrapModeMain_39)
	v1005 = F_SimpleLruZeroPage(m, v1002, int64(0))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	F_SimpleLruWritePage(m, v1002, v1005)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_LWLockRelease(m, v998)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[41]))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1012)+28))
	v1015 = F_LWLockAcquire(m, v1013, int32(0))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v1017 = int32(_a_F_BootstrapModeMain_40)
	v1020 = F_SimpleLruZeroPage(m, v1017, int64(0))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_SimpleLruWritePage(m, v1017, v1020)
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	F_LWLockRelease(m, v1013)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[42]))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+28))
	v1030 = F_LWLockAcquire(m, v1028, int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v1032 = int32(_a_F_BootstrapModeMain_41)
	v1035 = F_SimpleLruZeroPage(m, v1032, int64(0))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	F_SimpleLruWritePage(m, v1032, v1035)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_LWLockRelease(m, v1028)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[43]))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+28))
	v1045 = F_LWLockAcquire(m, v1043, int32(0))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v1047 = int32(_a_F_BootstrapModeMain_42)
	v1050 = F_SimpleLruZeroPage(m, v1047, int64(0))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	F_SimpleLruWritePage(m, v1047, v1050)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	F_LWLockRelease(m, v1043)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	F_pfree(m, v577)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	F_ReadControlFile(m)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	m.G0 = v552 + int32(_a_F_BootstrapModeMain_21)
	goto L183
L253:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_43), int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_29), int32(_a_F_BootstrapModeMain_44), int32(_a_F_BootstrapModeMain_31))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_45), int32(0))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_29), int32(_a_F_BootstrapModeMain_46), int32(_a_F_BootstrapModeMain_31))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_47), int32(0))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_29), int32(_a_F_BootstrapModeMain_48), int32(_a_F_BootstrapModeMain_49))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
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
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552))) = int32(_a_F_BootstrapModeMain_35)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_50), v552)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_29), int32(_a_F_BootstrapModeMain_51), int32(_a_F_BootstrapModeMain_38))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+32)) = int32(_a_F_BootstrapModeMain_35)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_52), v552+int32(32))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_29), int32(_a_F_BootstrapModeMain_53), int32(_a_F_BootstrapModeMain_38))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
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
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+16)) = int32(_a_F_BootstrapModeMain_35)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_54), v552+int32(16))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_29), int32(_a_F_BootstrapModeMain_55), int32(_a_F_BootstrapModeMain_38))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
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
	base.MemoryFill(m, int32(_a_F_BootstrapModeMain_56), int32(0), int32(160))
	v1174 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[44])) = v1174
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[45])) = v1174
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[46])) = v1174
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[47])) = v1174
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[48])) = v1174
	v1190 = F_boot_yylex_init(m, v29+int32(92))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	if v1190 != 0 {
		goto L4
	} else {
		goto L279
	}
L279:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	v1195 = m.G0
	v1197 = v1195 - int32(1136)
	m.G0 = v1197
	*(*int32)(unsafe.Add(mBase, uint32(v1197)+1132)) = int32(0)
	v1202 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[49]))
	v1205 = v1197 + int32(128)
	v1207 = v1197 + int32(928)
	v1212 = v1194
	v1214 = int32(-2)
	v1216 = v1205
	v1219 = v1197
	v1220 = v1207
	v1222 = v4
	v1223 = v1207
	v1225 = v1202
	v1226 = v1205
	v1227 = int32(200)
	v1231 = v4
	goto L288
L281:
	;
	if v6189+int32(928) != v6179 {
		goto L1196
	} else {
		goto L1197
	}
L282:
	;
	v6179 = v6164
	v6189 = v6163
	goto L281
L283:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6143 = m.ExcPending
	if v6143 != 0 {
		goto L1
	} else {
		goto L1193
	}
L284:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6130 = m.ExcPending
	if v6130 != 0 {
		goto L1
	} else {
		goto L1190
	}
L285:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6110 = m.ExcPending
	if v6110 != 0 {
		goto L1
	} else {
		goto L1187
	}
L286:
	;
	F_boot_yyerror(m, v1212, int32(_a_F_BootstrapModeMain_57))
	mBase = m.M
	v6106 = m.ExcPending
	if v6106 != 0 {
		goto L1
	} else {
		goto L1186
	}
L287:
	;
	v6087 = v6061
	goto L1183
L288:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1223))) = uint8(v1222)
	if base.Ui32(v1220+v1227-int32(1)) <= base.Ui32(v1223) {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	switch v3792 {
	case 0:
		goto L1180
	default:
		v6057 = v3780
		v6058 = v3781
		v6061 = v3784
		goto L287
	case 3:
		goto L1179
	}
L290:
	;
	if int32(_a_F_BootstrapModeMain_58) < v1227 {
		goto L286
	} else {
		goto L293
	}
L291:
	;
	v1286 = v1216
	v1287 = v1220
	v1288 = v1223
	v1289 = v1226
	v1290 = v1227
	goto L292
L292:
	;
	if v1222 == int32(46) {
		goto L310
	} else {
		goto L311
	}
L293:
	;
	v1242 = int32(_a_F_BootstrapModeMain_25)
	v1244 = v1227 << (uint(int32(1)) % 32)
	if v1242 <= v1244 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1247 = v1242
	goto L296
L295:
	;
	v1247 = v1244
	goto L296
L296:
	;
	v1252 = F_palloc(m, v1247*int32(5)+int32(3))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	if v1252 == int32(0) {
		goto L286
	} else {
		goto L298
	}
L298:
	;
	v1256 = v1223 - v1220
	v1258 = v1256 + int32(1)
	if v1258 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	base.MemoryCopy(m, v1252, v1220, v1258)
	goto L301
L300:
	;
	goto L301
L301:
	;
	v1263 = base.I32_div_s(v1247+int32(3), int32(4))
	v1264 = int32(2)
	v1266 = v1252 + v1263<<(uint(v1264)%32)
	v1268 = v1258 << (uint(v1264) % 32)
	if v1268 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	base.MemoryCopy(m, v1266, v1226, v1268)
	goto L304
L303:
	;
	goto L304
L304:
	;
	if v1219+int32(928) != v1220 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	F_pfree(m, v1220)
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L1
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	if v1247-int32(1) <= v1256 {
		v6179 = v1252
		v6189 = v1219
		goto L281
	} else {
		goto L309
	}
L308:
	;
	goto L307
L309:
	;
	v1286 = v1266 + v1268 - int32(4)
	v1287 = v1252
	v1288 = v1252 + v1256
	v1289 = v1266
	v1290 = v1247
	goto L292
L310:
	;
	v6179 = v1287
	v6189 = v1219
	goto L281
L311:
	;
	goto L312
L312:
	;
	v1297 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1222<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[50]))))
	if v1297 == int32(-53) {
		v3773 = v1212
		v3775 = v1214
		v3777 = v1286
		v3780 = v1219
		v3781 = v1287
		v3783 = v1222
		v3784 = v1288
		v3786 = v1225
		v3787 = v1289
		v3788 = v1290
		v3792 = v1231
		goto L316
	} else {
		goto L317
	}
L313:
	;
	goto L289
L314:
	;
	v1212 = v6017
	v1214 = v6019
	v1216 = v6021
	v1219 = v6024
	v1220 = v6025
	v1222 = v6027
	v1223 = v6028 + int32(1)
	v1225 = v6030
	v1226 = v6031
	v1227 = v6032
	v1231 = v6036
	goto L288
L315:
	;
	v3830 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3814)+uint32(_c_F_BootstrapModeMain[51]))))
	v3835 = *(*int32)(unsafe.Add(mBase, uint32(v3808+(int32(1)-v3830)<<(uint(int32(2))%32))))
	switch v3814 - int32(14) {
	case 0:
		goto L738
	case 1:
		goto L737
	case 2:
		goto L736
	case 3:
		goto L735
	case 4:
		goto L734
	case 5:
		goto L733
	case 6:
		goto L732
	case 7:
		goto L731
	case 8:
		goto L730
	case 9:
		goto L729
	case 10:
		goto L728
	case 11:
		goto L727
	case 12:
		goto L726
	case 13:
		goto L725
	case 14, 16, 25:
		goto L724
	case 15, 17, 19:
		goto L723
	case 18:
		goto L722
	default:
		v5963 = v3835
		goto L695
	case 22:
		goto L721
	case 23:
		goto L720
	case 24:
		goto L719
	case 26:
		goto L718
	case 30:
		goto L717
	case 31:
		goto L716
	case 32:
		goto L715
	case 33:
		goto L714
	case 34:
		goto L713
	case 35:
		goto L712
	case 36:
		goto L711
	case 37:
		goto L710
	case 38:
		goto L709
	case 39:
		goto L708
	case 40:
		goto L707
	case 41:
		goto L706
	case 42:
		goto L705
	case 43:
		goto L704
	case 44:
		goto L703
	case 45:
		goto L702
	case 46:
		goto L701
	case 47:
		goto L700
	case 48:
		goto L699
	case 49:
		goto L698
	case 50:
		goto L697
	case 51:
		goto L696
	}
L316:
	;
	v3798 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3783)+uint32(_c_F_BootstrapModeMain[52]))))
	if v3798 == int32(0) {
		goto L313
	} else {
		goto L694
	}
L317:
	;
	if v1214 == int32(-2) {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	v3752 = v1297 + v3751
	if base.Ui32(int32(169)) < base.Ui32(v3752) {
		v3773 = v3717
		v3775 = v3750
		v3777 = v3721
		v3780 = v3724
		v3781 = v3725
		v3783 = v3727
		v3784 = v3728
		v3786 = v3730
		v3787 = v3731
		v3788 = v3732
		v3792 = v3736
		goto L316
	} else {
		goto L689
	}
L319:
	;
	v1302 = m.G0
	v1304 = v1302 - int32(16)
	m.G0 = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+92)) = v1219 + int32(1132)
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+40))
	if v1309 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	v3717 = v1212
	v3719 = v1214
	v3721 = v1286
	v3724 = v1219
	v3725 = v1287
	v3727 = v1222
	v3728 = v1288
	v3730 = v1225
	v3731 = v1289
	v3732 = v1290
	v3736 = v1231
	goto L321
L321:
	;
	if v3719 <= int32(0) {
		goto L684
	} else {
		goto L685
	}
L322:
	;
	v3717 = v1420
	v3719 = v2245
	v3721 = v1424
	v3724 = v1427
	v3725 = v1428
	v3727 = v1430
	v3728 = v1431
	v3730 = v1433
	v3731 = v1434
	v3732 = v1435
	v3736 = v1439
	goto L321
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+40)) = int32(1)
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+44))
	if v1314 == int32(0) {
		goto L326
	} else {
		goto L327
	}
L324:
	;
	goto L325
L325:
	;
	v1382 = v1212
	v1386 = v1286
	v1389 = v1219
	v1390 = v1287
	v1392 = v1222
	v1393 = v1288
	v1395 = v1225
	v1396 = v1289
	v1397 = v1290
	v1398 = v1304
	v1401 = v1231
	goto L342
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+44)) = int32(1)
	goto L328
L327:
	;
	goto L328
L328:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+4))
	if v1319 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[53]))
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+4)) = v1323
	goto L331
L330:
	;
	goto L331
L331:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+8))
	if v1325 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+8)) = v1329
	goto L334
L333:
	;
	goto L334
L334:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+20))
	if v1331 != 0 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1359)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+28)) = v1360
	v1364 = v1358 + v1357<<(uint(int32(2))%32)
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1364)))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+80)) = v1366
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+36)) = v1366
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1364)))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1369)))
	*(*int32)(unsafe.Add(mBase, uint32(v1212)+4)) = v1370
	v1372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1366))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1212)+24)) = uint8(v1372)
	goto L325
L336:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+12))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1331+v1332<<(uint(int32(2))%32))))
	if v1336 != 0 {
		v1357 = v1332
		v1358 = v1331
		v1359 = v1336
		goto L335
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	F_boot_yyensure_buffer_stack(m, v1212)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L1
	} else {
		goto L340
	}
L339:
	;
	goto L338
L340:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+4))
	v1342 = F_boot_yy_create_buffer(m, v1341, v1212)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+20))
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+12))
	v1346 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1344+v1345<<(uint(v1346)%32)))) = v1342
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+20))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1212)+12))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1350+v1351<<(uint(v1346)%32))))
	v1357 = v1351
	v1358 = v1350
	v1359 = v1355
	goto L335
L342:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1382)+36))
	v1406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1405))) = uint8(v1406)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1382)+20))
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1382)+12))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1408+v1409<<(uint(int32(2))%32))))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1413)+28))
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1382)+44))
	v1417 = v1405
	v1420 = v1382
	v1421 = v1405
	v1422 = v1414 + v1415
	v1424 = v1386
	v1427 = v1389
	v1428 = v1390
	v1430 = v1392
	v1431 = v1393
	v1433 = v1395
	v1434 = v1396
	v1435 = v1397
	v1436 = v1398
	v1439 = v1401
	goto L344
L344:
	;
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1421))))
	v1444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1443)+uint32(_c_F_BootstrapModeMain[54]))))
	v1446 = v1422 << (uint(int32(1)) % 32)
	v1449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1446)+uint32(_c_F_BootstrapModeMain[55]))))
	if v1449 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+68)) = v1421
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+64)) = v1422
	goto L348
L347:
	;
	goto L348
L348:
	;
	v1454 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1446)+uint32(_c_F_BootstrapModeMain[56]))))
	v1455 = v1454 + v1444
	v1460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1455<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v1460 != v1422 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1464 = v1444
	v1467 = v1422
	v1468 = v1444
	goto L352
L350:
	;
	v1522 = v1455
	goto L351
L351:
	;
	v1539 = int32(1)
	v1545 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1522<<(uint(v1539)%32))+uint32(_c_F_BootstrapModeMain[58]))))
	if v1545 != int32(127) {
		v1421 = v1421 + v1539
		v1422 = v1545
		goto L344
	} else {
		goto L358
	}
L352:
	;
	v1492 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1467<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[59]))))
	if int32(128) <= v1492 {
		goto L354
	} else {
		goto L355
	}
L353:
	;
	v1522 = v1504
	goto L351
L354:
	;
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1468)+uint32(_c_F_BootstrapModeMain[60]))))
	v1496 = v1495
	goto L356
L355:
	;
	v1496 = v1464
	goto L356
L356:
	;
	v1498 = v1496 & int32(255)
	v1499 = int32(1)
	v1503 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1492<<(uint(v1499)%32))+uint32(_c_F_BootstrapModeMain[56]))))
	v1504 = v1498 + v1503
	v1509 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1504<<(uint(v1499)%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v1509 != v1492&int32(_a_F_BootstrapModeMain_59) {
		v1464 = v1496
		v1467 = v1492
		v1468 = v1498
		goto L352
	} else {
		goto L357
	}
L357:
	;
	goto L353
L358:
	;
	v1549 = v1417
	goto L359
L359:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+64))
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+68))
	v1577 = v1549
	v1581 = v1574
	v1584 = v1575
	goto L361
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+80)) = v1577
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+32)) = v1584 - v1577
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1584))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1420)+24)) = uint8(v1605)
	v1607 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1584))) = uint8(v1607)
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+36)) = v1584
	v1614 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1581<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[55]))))
	v1620 = v1614
	goto L363
L363:
	;
	switch v1620 {
	case 0:
		goto L409
	case 1:
		goto L408
	case 2:
		goto L407
	case 3:
		goto L406
	case 4:
		goto L405
	case 5:
		goto L404
	case 6:
		goto L403
	case 7:
		goto L402
	case 8:
		goto L401
	case 9:
		goto L400
	case 10:
		goto L399
	case 11:
		goto L398
	case 12:
		goto L397
	case 13:
		goto L396
	case 14:
		goto L395
	case 15:
		goto L394
	case 16:
		goto L393
	case 17:
		goto L392
	case 18:
		goto L391
	case 19:
		goto L390
	case 20:
		goto L389
	case 21:
		goto L388
	case 22:
		goto L387
	case 23:
		goto L386
	case 24:
		goto L385
	case 25:
		goto L384
	case 26:
		goto L383
	case 27:
		goto L382
	case 28:
		goto L381
	case 29:
		goto L380
	case 30:
		goto L377
	case 31:
		goto L376
	case 32:
		goto L375
	case 33:
		v2245 = int32(0)
		goto L378
	default:
		goto L374
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+36)) = v3678
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+48)) = int32(0)
	v3707 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+44))
	v3711 = base.I32_div_s(v3707-int32(1), int32(2))
	v1620 = v3711 + int32(33)
	goto L363
L366:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_60))
	mBase = m.M
	v3677 = m.ExcPending
	if v3677 != 0 {
		goto L1
	} else {
		goto L683
	}
L367:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_61))
	mBase = m.M
	v3671 = m.ExcPending
	if v3671 != 0 {
		goto L1
	} else {
		goto L682
	}
L368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L369:
	;
	v3523 = v3501 + v3502
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+36)) = v3523
	v3528 = *(*int32)(unsafe.Add(mBase, uint32(v3503+v3506<<(uint(int32(2))%32))))
	v3529 = *(*int32)(unsafe.Add(mBase, uint32(v3528)+28))
	v3530 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+44))
	v3531 = v3529 + v3530
	if base.Ui32(v3523) <= base.Ui32(v3497) {
		v1577 = v3497
		v1581 = v3531
		v1584 = v3523
		goto L361
	} else {
		goto L663
	}
L370:
	;
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(v3144)))
	*(*int32)(unsafe.Add(mBase, uint32(v3145)+16)) = v3122
	v3148 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+28))
	if v3148 != 0 {
		v3270 = int32(0)
		goto L607
	} else {
		goto L608
	}
L371:
	;
	v3113 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3122 = v3091
	v3144 = v3113 + v3114<<(uint(int32(2))%32)
	goto L370
L372:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_62))
	mBase = m.M
	v3086 = m.ExcPending
	if v3086 != 0 {
		goto L1
	} else {
		goto L606
	}
L373:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_63))
	mBase = m.M
	v3083 = m.ExcPending
	if v3083 != 0 {
		goto L1
	} else {
		goto L605
	}
L374:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_64))
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L1
	} else {
		goto L604
	}
L375:
	;
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1584))) = uint8(v2309)
	v2311 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2315 = v2311 + v2312<<(uint(int32(2))%32)
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(v2315)))
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v2316)+44))
	if v2317 == int32(0) {
		goto L495
	} else {
		goto L496
	}
L376:
	;
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2288 {
		goto L491
	} else {
		goto L492
	}
L377:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2249 {
		goto L485
	} else {
		goto L486
	}
L378:
	;
	m.G0 = v1436 + int32(16)
	goto L322
L379:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2240))) = v2238
	v2245 = int32(258)
	goto L378
L380:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2218 {
		goto L481
	} else {
		goto L482
	}
L381:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2197 {
		goto L477
	} else {
		goto L478
	}
L382:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2175 {
		goto L474
	} else {
		goto L475
	}
L383:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2154 {
		goto L471
	} else {
		goto L472
	}
L384:
	;
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2133 {
		goto L468
	} else {
		goto L469
	}
L385:
	;
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2112 {
		goto L465
	} else {
		goto L466
	}
L386:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2091 {
		goto L462
	} else {
		goto L463
	}
L387:
	;
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2070 {
		goto L459
	} else {
		goto L460
	}
L388:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2049 {
		goto L456
	} else {
		goto L457
	}
L389:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2028 {
		goto L453
	} else {
		goto L454
	}
L390:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v2007 {
		goto L450
	} else {
		goto L451
	}
L391:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1986 {
		goto L447
	} else {
		goto L448
	}
L392:
	;
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1965 {
		goto L444
	} else {
		goto L445
	}
L393:
	;
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if v1948 <= int32(0) {
		v1382 = v1420
		v1386 = v1424
		v1389 = v1427
		v1390 = v1428
		v1392 = v1430
		v1393 = v1431
		v1395 = v1433
		v1396 = v1434
		v1397 = v1435
		v1398 = v1436
		v1401 = v1439
		goto L342
	} else {
		goto L443
	}
L394:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if v1931 <= int32(0) {
		v1382 = v1420
		v1386 = v1424
		v1389 = v1427
		v1390 = v1428
		v1392 = v1430
		v1393 = v1431
		v1395 = v1433
		v1396 = v1434
		v1397 = v1435
		v1398 = v1436
		v1401 = v1439
		goto L342
	} else {
		goto L442
	}
L395:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1904 {
		goto L439
	} else {
		goto L440
	}
L396:
	;
	v1884 = int32(262)
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if v1885 <= int32(0) {
		v2245 = v1884
		goto L378
	} else {
		goto L438
	}
L397:
	;
	v1866 = int32(261)
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if v1867 <= int32(0) {
		v2245 = v1866
		goto L378
	} else {
		goto L437
	}
L398:
	;
	v1848 = int32(260)
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if v1849 <= int32(0) {
		v2245 = v1848
		goto L378
	} else {
		goto L436
	}
L399:
	;
	v1830 = int32(259)
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if v1831 <= int32(0) {
		v2245 = v1830
		goto L378
	} else {
		goto L435
	}
L400:
	;
	v1812 = int32(263)
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if v1813 <= int32(0) {
		v2245 = v1812
		goto L378
	} else {
		goto L434
	}
L401:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1791 {
		goto L431
	} else {
		goto L432
	}
L402:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1770 {
		goto L428
	} else {
		goto L429
	}
L403:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1749 {
		goto L425
	} else {
		goto L426
	}
L404:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1728 {
		goto L422
	} else {
		goto L423
	}
L405:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1707 {
		goto L419
	} else {
		goto L420
	}
L406:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1686 {
		goto L416
	} else {
		goto L417
	}
L407:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1665 {
		goto L413
	} else {
		goto L414
	}
L408:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+32))
	if int32(0) < v1644 {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	v1642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1584))) = uint8(v1642)
	v1549 = v1577
	goto L359
L410:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1647+v1648<<(uint(int32(2))%32))))
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1653+v1644-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1652)+28)) = base.B2i32(v1657 == int32(10))
	goto L412
L411:
	;
	goto L412
L412:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1661))) = int32(_a_F_BootstrapModeMain_65)
	v2245 = int32(264)
	goto L378
L413:
	;
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1668+v1669<<(uint(int32(2))%32))))
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1674+v1665-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1673)+28)) = base.B2i32(v1678 == int32(10))
	goto L415
L414:
	;
	goto L415
L415:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1682))) = int32(_a_F_BootstrapModeMain_66)
	v2245 = int32(265)
	goto L378
L416:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1689+v1690<<(uint(int32(2))%32))))
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1695+v1686-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1694)+28)) = base.B2i32(v1699 == int32(10))
	goto L418
L417:
	;
	goto L418
L418:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1703))) = int32(_a_F_BootstrapModeMain_67)
	v2245 = int32(266)
	goto L378
L419:
	;
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1710+v1711<<(uint(int32(2))%32))))
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1716+v1707-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1715)+28)) = base.B2i32(v1720 == int32(10))
	goto L421
L420:
	;
	goto L421
L421:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1724))) = int32(_a_F_BootstrapModeMain_68)
	v2245 = int32(276)
	goto L378
L422:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1731+v1732<<(uint(int32(2))%32))))
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1737+v1728-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1736)+28)) = base.B2i32(v1741 == int32(10))
	goto L424
L423:
	;
	goto L424
L424:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1745))) = int32(_a_F_BootstrapModeMain_69)
	v2245 = int32(277)
	goto L378
L425:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1752+v1753<<(uint(int32(2))%32))))
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1758+v1749-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1757)+28)) = base.B2i32(v1762 == int32(10))
	goto L427
L426:
	;
	goto L427
L427:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1766))) = int32(_a_F_BootstrapModeMain_70)
	v2245 = int32(278)
	goto L378
L428:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1773+v1774<<(uint(int32(2))%32))))
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1779+v1770-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1778)+28)) = base.B2i32(v1783 == int32(10))
	goto L430
L429:
	;
	goto L430
L430:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1787))) = int32(_a_F_BootstrapModeMain_71)
	v2245 = int32(279)
	goto L378
L431:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1794+v1795<<(uint(int32(2))%32))))
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1800+v1791-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1799)+28)) = base.B2i32(v1804 == int32(10))
	goto L433
L432:
	;
	goto L433
L433:
	;
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1808))) = int32(_a_F_BootstrapModeMain_72)
	v2245 = int32(267)
	goto L378
L434:
	;
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1816+v1817<<(uint(int32(2))%32))))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1822+v1813-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+28)) = base.B2i32(v1826 == int32(10))
	v2245 = v1812
	goto L378
L435:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1834+v1835<<(uint(int32(2))%32))))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1840+v1831-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1839)+28)) = base.B2i32(v1844 == int32(10))
	v2245 = v1830
	goto L378
L436:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1852+v1853<<(uint(int32(2))%32))))
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1858+v1849-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1857)+28)) = base.B2i32(v1862 == int32(10))
	v2245 = v1848
	goto L378
L437:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1870+v1871<<(uint(int32(2))%32))))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1876+v1867-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1875)+28)) = base.B2i32(v1880 == int32(10))
	v2245 = v1866
	goto L378
L438:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1888+v1889<<(uint(int32(2))%32))))
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1894+v1885-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1893)+28)) = base.B2i32(v1898 == int32(10))
	v2245 = v1884
	goto L378
L439:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1903+v1902<<(uint(int32(2))%32))))
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1911+v1904-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1910)+28)) = base.B2i32(v1915 == int32(10))
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1921 = v1919
	v1922 = v1920
	goto L441
L440:
	;
	v1921 = v1903
	v1922 = v1902
	goto L441
L441:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1922<<(uint(int32(2))%32)+v1921)))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1926)+32)) = v1927 + int32(1)
	v1382 = v1420
	v1386 = v1424
	v1389 = v1427
	v1390 = v1428
	v1392 = v1430
	v1393 = v1431
	v1395 = v1433
	v1396 = v1434
	v1397 = v1435
	v1398 = v1436
	v1401 = v1439
	goto L342
L442:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1934+v1935<<(uint(int32(2))%32))))
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1940+v1931-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1939)+28)) = base.B2i32(v1944 == int32(10))
	v1382 = v1420
	v1386 = v1424
	v1389 = v1427
	v1390 = v1428
	v1392 = v1430
	v1393 = v1431
	v1395 = v1433
	v1396 = v1434
	v1397 = v1435
	v1398 = v1436
	v1401 = v1439
	goto L342
L443:
	;
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1951+v1952<<(uint(int32(2))%32))))
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1957+v1948-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1956)+28)) = base.B2i32(v1961 == int32(10))
	v1382 = v1420
	v1386 = v1424
	v1389 = v1427
	v1390 = v1428
	v1392 = v1430
	v1393 = v1431
	v1395 = v1433
	v1396 = v1434
	v1397 = v1435
	v1398 = v1436
	v1401 = v1439
	goto L342
L444:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(v1968+v1969<<(uint(int32(2))%32))))
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1974+v1965-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1973)+28)) = base.B2i32(v1978 == int32(10))
	goto L446
L445:
	;
	goto L446
L446:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1982))) = int32(_a_F_BootstrapModeMain_73)
	v2245 = int32(268)
	goto L378
L447:
	;
	v1989 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1989+v1990<<(uint(int32(2))%32))))
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v1999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1995+v1986-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1994)+28)) = base.B2i32(v1999 == int32(10))
	goto L449
L448:
	;
	goto L449
L449:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2003))) = int32(_a_F_BootstrapModeMain_74)
	v2245 = int32(272)
	goto L378
L450:
	;
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v2010+v2011<<(uint(int32(2))%32))))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2020 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2016+v2007-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2015)+28)) = base.B2i32(v2020 == int32(10))
	goto L452
L451:
	;
	goto L452
L452:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2024))) = int32(_a_F_BootstrapModeMain_75)
	v2245 = int32(273)
	goto L378
L453:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v2031+v2032<<(uint(int32(2))%32))))
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2037+v2028-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2036)+28)) = base.B2i32(v2041 == int32(10))
	goto L455
L454:
	;
	goto L455
L455:
	;
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2045))) = int32(_a_F_BootstrapModeMain_76)
	v2245 = int32(274)
	goto L378
L456:
	;
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v2052+v2053<<(uint(int32(2))%32))))
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2058+v2049-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2057)+28)) = base.B2i32(v2062 == int32(10))
	goto L458
L457:
	;
	goto L458
L458:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2066))) = int32(_a_F_BootstrapModeMain_77)
	v2245 = int32(269)
	goto L378
L459:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v2073+v2074<<(uint(int32(2))%32))))
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2079+v2070-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2078)+28)) = base.B2i32(v2083 == int32(10))
	goto L461
L460:
	;
	goto L461
L461:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2087))) = int32(_a_F_BootstrapModeMain_78)
	v2245 = int32(270)
	goto L378
L462:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v2094+v2095<<(uint(int32(2))%32))))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2100+v2091-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2099)+28)) = base.B2i32(v2104 == int32(10))
	goto L464
L463:
	;
	goto L464
L464:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2108))) = int32(_a_F_BootstrapModeMain_79)
	v2245 = int32(271)
	goto L378
L465:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v2115+v2116<<(uint(int32(2))%32))))
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2121+v2112-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2120)+28)) = base.B2i32(v2125 == int32(10))
	goto L467
L466:
	;
	goto L467
L467:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2129))) = int32(_a_F_BootstrapModeMain_80)
	v2245 = int32(275)
	goto L378
L468:
	;
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v2136+v2137<<(uint(int32(2))%32))))
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2142+v2133-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2141)+28)) = base.B2i32(v2146 == int32(10))
	goto L470
L469:
	;
	goto L470
L470:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2150))) = int32(_a_F_BootstrapModeMain_81)
	v2245 = int32(280)
	goto L378
L471:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v2157+v2158<<(uint(int32(2))%32))))
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2163+v2154-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2162)+28)) = base.B2i32(v2167 == int32(10))
	goto L473
L472:
	;
	goto L473
L473:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2171))) = int32(_a_F_BootstrapModeMain_82)
	v2245 = int32(281)
	goto L378
L474:
	;
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v2178+v2179<<(uint(int32(2))%32))))
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2184+v2175-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2183)+28)) = base.B2i32(v2188 == int32(10))
	goto L476
L475:
	;
	goto L476
L476:
	;
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2192))) = int32(_a_F_BootstrapModeMain_83)
	v2245 = int32(282)
	goto L378
L477:
	;
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(v2200+v2201<<(uint(int32(2))%32))))
	v2209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196+v2197-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2205)+28)) = base.B2i32(v2209 == int32(10))
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2214 = v2213
	goto L479
L478:
	;
	v2214 = v2196
	goto L479
L479:
	;
	v2215 = F_pstrdup(m, v2214)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	v2238 = v2215
	goto L379
L481:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2221+v2222<<(uint(int32(2))%32))))
	v2230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2217+v2218-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2226)+28)) = base.B2i32(v2230 == int32(10))
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2235 = v2234
	goto L483
L482:
	;
	v2235 = v2217
	goto L483
L483:
	;
	v2236 = F_DeescapeQuotedString(m, v2235)
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L1
	} else {
		goto L484
	}
L484:
	;
	v2238 = v2236
	goto L379
L485:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v2252+v2253<<(uint(int32(2))%32))))
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2258+v2249-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2257)+28)) = base.B2i32(v2262 == int32(10))
	goto L487
L486:
	;
	goto L487
L487:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v2270+v2271<<(uint(int32(2))%32))))
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v2275)+32))
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1436)+4)) = v2277
	*(*int32)(unsafe.Add(mBase, uint32(v1436))) = v2276
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_84), v1436)
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_85), int32(124), int32(_a_F_BootstrapModeMain_86))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L1
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
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2291+v2292<<(uint(int32(2))%32))))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2297+v2288-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2296)+28)) = base.B2i32(v2301 == int32(10))
	goto L493
L492:
	;
	goto L493
L493:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_87))
	mBase = m.M
	v2307 = m.ExcPending
	if v2307 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L495:
	;
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v2316)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+28)) = v2320
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v2315)))
	v2323 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2322))) = v2323
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2326 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2327 = int32(2)
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v2325+v2326<<(uint(v2327)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2330)+44)) = int32(1)
	v2333 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2334 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2333+v2334<<(uint(v2327)%32))))
	v2339 = v2338
	v2340 = v2333
	v2341 = v2334
	goto L497
L496:
	;
	v2339 = v2316
	v2340 = v2311
	v2341 = v2312
	goto L497
L497:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+36))
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v2339)+4))
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+28))
	v2345 = v2343 + v2344
	if base.Ui32(v2342) <= base.Ui32(v2345) {
		goto L498
	} else {
		goto L499
	}
L498:
	;
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2350 = v2308 ^ int32(-1) + v1584
	v2351 = v2347 + v2350
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+36)) = v2351
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v2340+v2341<<(uint(int32(2))%32))))
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(v2356)+28))
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+44))
	v2359 = v2357 + v2358
	if int32(0) < v2350 {
		goto L501
	} else {
		goto L502
	}
L499:
	;
	goto L500
L500:
	;
	if base.Ui32(v2345+int32(1)) < base.Ui32(v2342) {
		goto L373
	} else {
		goto L533
	}
L501:
	;
	v2366 = v2347
	v2367 = v2359
	goto L504
L502:
	;
	v2501 = v2359
	goto L503
L503:
	;
	v2526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2501<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[55]))))
	if v2526 != 0 {
		goto L522
	} else {
		goto L523
	}
L504:
	;
	v2389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2366))))
	if v2389 != 0 {
		goto L506
	} else {
		goto L507
	}
L505:
	;
	v2501 = v2492
	goto L503
L506:
	;
	v2390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2389)+uint32(_c_F_BootstrapModeMain[54]))))
	v2391 = v2390
	goto L508
L507:
	;
	v2391 = int32(1)
	goto L508
L508:
	;
	v2393 = v2367 << (uint(int32(1)) % 32)
	v2396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+uint32(_c_F_BootstrapModeMain[55]))))
	if v2396 != 0 {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+68)) = v2366
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+64)) = v2367
	goto L511
L510:
	;
	goto L511
L511:
	;
	v2400 = v2391 & int32(255)
	v2403 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2393)+uint32(_c_F_BootstrapModeMain[56]))))
	v2404 = v2400 + v2403
	v2409 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2404<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v2409 != v2367 {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v2413 = v2391
	v2416 = v2367
	v2417 = v2400
	goto L515
L513:
	;
	v2471 = v2404
	goto L514
L514:
	;
	v2488 = int32(1)
	v2492 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2471<<(uint(v2488)%32))+uint32(_c_F_BootstrapModeMain[58]))))
	v2494 = v2366 + v2488
	if v2494 != v2351 {
		v2366 = v2494
		v2367 = v2492
		goto L504
	} else {
		goto L521
	}
L515:
	;
	v2441 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2416<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[59]))))
	if int32(128) <= v2441 {
		goto L517
	} else {
		goto L518
	}
L516:
	;
	v2471 = v2453
	goto L514
L517:
	;
	v2444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2417)+uint32(_c_F_BootstrapModeMain[60]))))
	v2445 = v2444
	goto L519
L518:
	;
	v2445 = v2413
	goto L519
L519:
	;
	v2447 = v2445 & int32(255)
	v2448 = int32(1)
	v2452 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2441<<(uint(v2448)%32))+uint32(_c_F_BootstrapModeMain[56]))))
	v2453 = v2447 + v2452
	v2458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2453<<(uint(v2448)%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v2458 != v2441&int32(_a_F_BootstrapModeMain_59) {
		v2413 = v2445
		v2416 = v2441
		v2417 = v2447
		goto L515
	} else {
		goto L520
	}
L520:
	;
	goto L516
L521:
	;
	goto L505
L522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+68)) = v2351
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+64)) = v2501
	goto L524
L523:
	;
	goto L524
L524:
	;
	v2529 = int32(1)
	v2533 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2501<<(uint(v2529)%32))+uint32(_c_F_BootstrapModeMain[56]))))
	v2535 = v2533 + v2529
	v2540 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2535<<(uint(v2529)%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v2540 != v2501 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v2547 = v2501
	goto L528
L526:
	;
	v2589 = v2535
	goto L527
L527:
	;
	if v2589 == int32(0) {
		v1549 = v2347
		goto L359
	} else {
		goto L531
	}
L528:
	;
	v2568 = int32(1)
	v2572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2547<<(uint(v2568)%32))+uint32(_c_F_BootstrapModeMain[59]))))
	v2573 = base.I32_extend16_s(v2572)
	v2578 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2573<<(uint(v2568)%32))+uint32(_c_F_BootstrapModeMain[56]))))
	v2580 = v2578 + v2568
	v2585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2580<<(uint(v2568)%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v2572 != v2585 {
		v2547 = v2573
		goto L528
	} else {
		goto L530
	}
L529:
	;
	v2589 = v2580
	goto L527
L530:
	;
	goto L529
L531:
	;
	v2619 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2589<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[58]))))
	if v2619 == int32(127) {
		v1549 = v2347
		goto L359
	} else {
		goto L532
	}
L532:
	;
	v2623 = v2351 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+36)) = v2623
	v1417 = v2347
	v1421 = v2623
	v1422 = v2619
	goto L344
L533:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+80))
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v2339)+40))
	if v2629 == int32(0) {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	if v2342-v2628 != int32(1) {
		v3497 = v2628
		v3501 = v2344
		v3502 = v2343
		v3503 = v2340
		v3506 = v2341
		goto L369
	} else {
		goto L537
	}
L535:
	;
	goto L536
L536:
	;
	v2637 = v2628 ^ int32(-1) + v2342
	if int32(0) < v2637 {
		goto L538
	} else {
		goto L539
	}
L537:
	;
	v3678 = v2628
	goto L365
L538:
	;
	v2640 = int32(7)
	v2641 = v2637 & v2640
	if base.Ui32(v2342-v2628-int32(2)) < base.Ui32(v2640) {
		goto L543
	} else {
		goto L544
	}
L539:
	;
	v2797 = v2339
	v2801 = v2340
	v2804 = v2341
	goto L540
L540:
	;
	v2821 = *(*int32)(unsafe.Add(mBase, uint32(v2797)+44))
	if v2821 == int32(2) {
		goto L553
	} else {
		goto L554
	}
L541:
	;
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(v2789+v2790<<(uint(int32(2))%32))))
	v2797 = v2794
	v2801 = v2789
	v2804 = v2790
	goto L540
L542:
	;
	v2730 = v2703
	v2733 = v2706
	v2734 = int32(0)
	goto L550
L543:
	;
	v2703 = v2628
	v2706 = v2343
	goto L542
L544:
	;
	goto L545
L545:
	;
	v2652 = v2628
	v2655 = v2343
	v2656 = int32(0)
	goto L546
L546:
	;
	v2676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2652))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2655))) = uint8(v2676)
	v2678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2652)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2655)+1)) = uint8(v2678)
	v2680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2652)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2655)+2)) = uint8(v2680)
	v2682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2652)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2655)+3)) = uint8(v2682)
	v2684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2652)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2655)+4)) = uint8(v2684)
	v2686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2652)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2655)+5)) = uint8(v2686)
	v2688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2652)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2655)+6)) = uint8(v2688)
	v2690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2652)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2655)+7)) = uint8(v2690)
	v2692 = int32(8)
	v2693 = v2655 + v2692
	v2695 = v2652 + v2692
	v2697 = v2656 + v2692
	if v2697 != v2637&int32(2147483640) {
		v2652 = v2695
		v2655 = v2693
		v2656 = v2697
		goto L546
	} else {
		goto L548
	}
L547:
	;
	if v2641 == int32(0) {
		goto L541
	} else {
		goto L549
	}
L548:
	;
	goto L547
L549:
	;
	v2703 = v2695
	v2706 = v2693
	goto L542
L550:
	;
	v2754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2730))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2733))) = uint8(v2754)
	v2756 = int32(1)
	v2761 = v2734 + v2756
	if v2761 != v2641 {
		v2730 = v2730 + v2756
		v2733 = v2733 + v2756
		v2734 = v2761
		goto L550
	} else {
		goto L552
	}
L551:
	;
	goto L541
L552:
	;
	goto L551
L553:
	;
	v2824 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+28)) = v2824
	v3122 = v2824
	v3144 = v2801 + v2804<<(uint(int32(2))%32)
	goto L370
L554:
	;
	goto L555
L555:
	;
	v2830 = int32(0)
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(v2797)+12))
	v2832 = v2628 - v2342
	v2833 = v2831 + v2832
	if v2833 <= v2830 {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	v2836 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+36))
	v2837 = v2836
	v2839 = v2797
	v2846 = v2831
	goto L559
L557:
	;
	v2903 = v2797
	v2906 = v2833
	goto L558
L558:
	;
	v2927 = int32(_a_F_BootstrapModeMain_23)
	if base.Ui32(v2927) <= base.Ui32(v2906) {
		goto L575
	} else {
		goto L576
	}
L559:
	;
	v2863 = *(*int32)(unsafe.Add(mBase, uint32(v2839)+20))
	if v2863 == int32(0) {
		goto L561
	} else {
		goto L562
	}
L560:
	;
	v2903 = v2896
	v2906 = v2898
	goto L558
L561:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2839)+4)) = int32(0)
	goto L366
L562:
	;
	goto L563
L563:
	;
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v2839)+4))
	v2870 = v2846 << (uint(int32(1)) % 32)
	if v2870 <= int32(0) {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v2874 = base.I32_div_s(v2846, int32(8))
	v2876 = v2874 + v2846
	goto L566
L565:
	;
	v2876 = v2870
	goto L566
L566:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2839)+12)) = v2876
	v2879 = v2876 + int32(2)
	if v2868 != 0 {
		goto L568
	} else {
		goto L569
	}
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2839)+4)) = v2884
	if v2884 == int32(0) {
		goto L366
	} else {
		goto L573
	}
L568:
	;
	v2880 = F_repalloc(m, v2868, v2879)
	mBase = m.M
	v2881 = m.ExcPending
	if v2881 != 0 {
		goto L1
	} else {
		goto L571
	}
L569:
	;
	goto L570
L570:
	;
	v2882 = F_palloc(m, v2879)
	mBase = m.M
	v2883 = m.ExcPending
	if v2883 != 0 {
		goto L1
	} else {
		goto L572
	}
L571:
	;
	v2884 = v2880
	goto L567
L572:
	;
	v2884 = v2882
	goto L567
L573:
	;
	v2889 = v2884 + (v2837 - v2868)
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+36)) = v2889
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2892 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v2891+v2892<<(uint(int32(2))%32))))
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v2896)+12))
	v2898 = v2897 + v2832
	if v2898 <= int32(0) {
		v2837 = v2889
		v2839 = v2896
		v2846 = v2897
		goto L559
	} else {
		goto L574
	}
L574:
	;
	goto L560
L575:
	;
	v2930 = v2927
	goto L577
L576:
	;
	v2930 = v2906
	goto L577
L577:
	;
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(v2903)+24))
	if v2931 != 0 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v2936 = v2830
	goto L582
L579:
	;
	goto L580
L580:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30])) = int32(0)
	v3006 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3007 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(v3006+v3007<<(uint(int32(2))%32))))
	v3012 = *(*int32)(unsafe.Add(mBase, uint32(v3011)+4))
	v3015 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v3016 = F_fread(m, v3012+v2637, int32(1), v2930, v3015)
	mBase = m.M
	v3017 = m.ExcPending
	if v3017 != 0 {
		goto L1
	} else {
		goto L593
	}
L581:
	;
	switch v2962 {
	case 0:
		goto L589
	default:
		v3001 = v2976
		goto L587
	case 11:
		goto L588
	}
L582:
	;
	v2958 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v2959 = F_do_getc(m, v2958)
	mBase = m.M
	v2960 = m.ExcPending
	if v2960 != 0 {
		goto L1
	} else {
		goto L585
	}
L583:
	;
	v2976 = v2930
	goto L581
L584:
	;
	v2963 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2964 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2968 = *(*int32)(unsafe.Add(mBase, uint32(v2963+v2964<<(uint(int32(2))%32))))
	v2969 = *(*int32)(unsafe.Add(mBase, uint32(v2968)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2969+v2637+v2936))) = uint8(v2959)
	v2974 = v2936 + int32(1)
	if v2974 != v2930 {
		v2936 = v2974
		goto L582
	} else {
		goto L586
	}
L585:
	;
	v2962 = v2959 + int32(1)
	switch v2962 {
	case 0, 11:
		v2976 = v2936
		goto L581
	default:
		goto L584
	}
L586:
	;
	goto L583
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+28)) = v3001
	v3091 = v3001
	goto L371
L588:
	;
	v2988 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v2989 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v2993 = *(*int32)(unsafe.Add(mBase, uint32(v2988+v2989<<(uint(int32(2))%32))))
	v2994 = *(*int32)(unsafe.Add(mBase, uint32(v2993)+4))
	v2997 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v2994+v2637+v2976))) = uint8(v2997)
	v3001 = v2976 + int32(1)
	goto L587
L589:
	;
	v2977 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v2978 = *(*int32)(unsafe.Add(mBase, uint32(v2977)))
	goto L590
L590:
	;
	if int32(base.Ui32(v2978)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v3001 = v2976
		goto L587
	} else {
		goto L591
	}
L591:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_62))
	mBase = m.M
	v2987 = m.ExcPending
	if v2987 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L593:
	;
	v3022 = v3016
	goto L594
L594:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+28)) = v3022
	if v3022 != 0 {
		v3091 = v3022
		goto L371
	} else {
		goto L596
	}
L596:
	;
	v3045 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v3045)))
	goto L597
L597:
	;
	if int32(base.Ui32(v3046)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v3091 = int32(0)
	goto L371
L599:
	;
	goto L600
L600:
	;
	v3055 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	if v3055 != int32(27) {
		goto L372
	} else {
		goto L601
	}
L601:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30])) = int32(0)
	v3061 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(v3061)))
	*(*int32)(unsafe.Add(mBase, uint32(v3061))) = v3062 & int32(-49)
	goto L602
L602:
	;
	v3066 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3067 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3071 = *(*int32)(unsafe.Add(mBase, uint32(v3066+v3067<<(uint(int32(2))%32))))
	v3072 = *(*int32)(unsafe.Add(mBase, uint32(v3071)+4))
	v3075 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v3076 = F_fread(m, v3072+v2637, int32(1), v2930, v3075)
	mBase = m.M
	v3077 = m.ExcPending
	if v3077 != 0 {
		goto L1
	} else {
		goto L603
	}
L603:
	;
	v3022 = v3076
	goto L594
L604:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L605:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L606:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L607:
	;
	v3271 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+28))
	v3272 = v3271 + v2637
	v3273 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3274 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3278 = *(*int32)(unsafe.Add(mBase, uint32(v3273+v3274<<(uint(int32(2))%32))))
	v3279 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+12))
	if v3279 < v3272 {
		goto L631
	} else {
		goto L632
	}
L608:
	;
	if v2637 == int32(0) {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	if v3152 != 0 {
		goto L614
	} else {
		goto L615
	}
L610:
	;
	goto L611
L611:
	;
	v3256 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3257 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3258 = int32(2)
	v3261 = *(*int32)(unsafe.Add(mBase, uint32(v3256+v3257<<(uint(v3258)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3261)+44)) = v3258
	v3270 = v3258
	goto L607
L612:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3218)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3218))) = v3151
	v3225 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	if v3225 != 0 {
		goto L627
	} else {
		goto L628
	}
L613:
	;
	v3175 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	v3176 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3180 = *(*int32)(unsafe.Add(mBase, uint32(v3173+v3176<<(uint(int32(2))%32))))
	if v3180 == int32(0) {
		goto L621
	} else {
		goto L622
	}
L614:
	;
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3157 = *(*int32)(unsafe.Add(mBase, uint32(v3152+v3153<<(uint(int32(2))%32))))
	if v3157 != 0 {
		v3173 = v3152
		goto L613
	} else {
		goto L617
	}
L615:
	;
	goto L616
L616:
	;
	F_boot_yyensure_buffer_stack(m, v1420)
	mBase = m.M
	v3159 = m.ExcPending
	if v3159 != 0 {
		goto L1
	} else {
		goto L618
	}
L617:
	;
	goto L616
L618:
	;
	v3160 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	v3161 = F_boot_yy_create_buffer(m, v3160, v1420)
	mBase = m.M
	v3162 = m.ExcPending
	if v3162 != 0 {
		goto L1
	} else {
		goto L619
	}
L619:
	;
	v3163 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3164 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3163+v3164<<(uint(int32(2))%32)))) = v3161
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	if v3169 != 0 {
		v3173 = v3169
		goto L613
	} else {
		goto L620
	}
L620:
	;
	v3171 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30]))
	v3218 = int32(0)
	v3221 = v3171
	goto L612
L621:
	;
	v3218 = int32(0)
	v3221 = v3175
	goto L612
L622:
	;
	goto L623
L623:
	;
	v3184 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+16)) = v3184
	v3186 = *(*int32)(unsafe.Add(mBase, uint32(v3180)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v3186))) = uint8(v3184)
	v3189 = *(*int32)(unsafe.Add(mBase, uint32(v3180)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v3189)+1)) = uint8(v3184)
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+44)) = v3184
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+28)) = int32(1)
	v3196 = *(*int32)(unsafe.Add(mBase, uint32(v3180)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+8)) = v3196
	v3198 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	if v3198 == v3184 {
		v3218 = v3180
		v3221 = v3175
		goto L612
	} else {
		goto L624
	}
L624:
	;
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3204 = v3198 + v3201<<(uint(int32(2))%32)
	v3205 = *(*int32)(unsafe.Add(mBase, uint32(v3204)))
	if v3180 != v3205 {
		v3218 = v3180
		v3221 = v3175
		goto L612
	} else {
		goto L625
	}
L625:
	;
	v3207 = *(*int32)(unsafe.Add(mBase, uint32(v3205)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+28)) = v3207
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v3204)))
	v3210 = *(*int32)(unsafe.Add(mBase, uint32(v3209)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+80)) = v3210
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+36)) = v3210
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v3204)))
	v3214 = *(*int32)(unsafe.Add(mBase, uint32(v3213)))
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+4)) = v3214
	v3216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3210))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1420)+24)) = uint8(v3216)
	v3218 = v3180
	v3221 = v3175
	goto L612
L626:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3218)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[30])) = v3221
	v3238 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3239 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3242 = v3238 + v3239<<(uint(int32(2))%32)
	v3243 = *(*int32)(unsafe.Add(mBase, uint32(v3242)))
	v3244 = *(*int32)(unsafe.Add(mBase, uint32(v3243)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+28)) = v3244
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(v3242)))
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(v3246)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+36)) = v3247
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+80)) = v3247
	v3250 = *(*int32)(unsafe.Add(mBase, uint32(v3242)))
	v3251 = *(*int32)(unsafe.Add(mBase, uint32(v3250)))
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+4)) = v3251
	v3253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3247))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1420)+24)) = uint8(v3253)
	v3270 = int32(1)
	goto L607
L627:
	;
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3230 = *(*int32)(unsafe.Add(mBase, uint32(v3225+v3226<<(uint(int32(2))%32))))
	if v3218 == v3230 {
		goto L626
	} else {
		goto L630
	}
L628:
	;
	goto L629
L629:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3218)+32)) = int64(1)
	goto L626
L630:
	;
	goto L629
L631:
	;
	v3283 = v3272 + v3271>>(uint(int32(1))%32)
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+4))
	if v3284 != 0 {
		goto L635
	} else {
		goto L636
	}
L632:
	;
	v3314 = v3273
	v3315 = v3272
	v3316 = v3274
	goto L633
L633:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+28)) = v3315
	v3318 = int32(2)
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(v3314+v3316<<(uint(v3318)%32))))
	v3322 = *(*int32)(unsafe.Add(mBase, uint32(v3321)+4))
	v3324 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3322+v3315))) = uint8(v3324)
	v3326 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3327 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3331 = *(*int32)(unsafe.Add(mBase, uint32(v3326+v3327<<(uint(v3318)%32))))
	v3332 = *(*int32)(unsafe.Add(mBase, uint32(v3331)+4))
	v3333 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v3332+v3333)+1)) = uint8(v3324)
	v3337 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3338 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3341 = v3337 + v3338<<(uint(v3318)%32)
	v3342 = *(*int32)(unsafe.Add(mBase, uint32(v3341)))
	v3343 = *(*int32)(unsafe.Add(mBase, uint32(v3342)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+80)) = v3343
	if v3270 == int32(1) {
		v3678 = v3343
		goto L365
	} else {
		goto L641
	}
L634:
	;
	v3290 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3291 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3292 = int32(2)
	v3295 = *(*int32)(unsafe.Add(mBase, uint32(v3290+v3291<<(uint(v3292)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3295)+4)) = v3289
	v3297 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3298 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3302 = *(*int32)(unsafe.Add(mBase, uint32(v3297+v3298<<(uint(v3292)%32))))
	v3303 = *(*int32)(unsafe.Add(mBase, uint32(v3302)+4))
	if v3303 == int32(0) {
		goto L367
	} else {
		goto L640
	}
L635:
	;
	v3285 = F_repalloc(m, v3284, v3283)
	mBase = m.M
	v3286 = m.ExcPending
	if v3286 != 0 {
		goto L1
	} else {
		goto L638
	}
L636:
	;
	goto L637
L637:
	;
	v3287 = F_palloc(m, v3283)
	mBase = m.M
	v3288 = m.ExcPending
	if v3288 != 0 {
		goto L1
	} else {
		goto L639
	}
L638:
	;
	v3289 = v3285
	goto L634
L639:
	;
	v3289 = v3287
	goto L634
L640:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3302)+12)) = v3283 - int32(2)
	v3309 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+12))
	v3310 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+28))
	v3312 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v3314 = v3312
	v3315 = v3310 + v2637
	v3316 = v3309
	goto L633
L641:
	;
	switch v3270 - int32(1) {
	case 0:
		goto L368
	case 1:
		goto L642
	default:
		goto L643
	}
L642:
	;
	v3494 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+28))
	v3495 = *(*int32)(unsafe.Add(mBase, uint32(v3341)))
	v3496 = *(*int32)(unsafe.Add(mBase, uint32(v3495)+4))
	v3497 = v3343
	v3501 = v3494
	v3502 = v3496
	v3503 = v3337
	v3506 = v3338
	goto L369
L643:
	;
	v3351 = v2308 ^ int32(-1) + v1584
	v3352 = v3343 + v3351
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+36)) = v3352
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(v3341)))
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v3354)+28))
	v3356 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+44))
	v3357 = v3355 + v3356
	if v3351 <= int32(0) {
		v1417 = v3343
		v1421 = v3352
		v1422 = v3357
		goto L344
	} else {
		goto L644
	}
L644:
	;
	v3365 = v3357
	v3368 = v3343
	goto L645
L645:
	;
	v3387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3368))))
	if v3387 != 0 {
		goto L647
	} else {
		goto L648
	}
L646:
	;
	v1417 = v3343
	v1421 = v3352
	v1422 = v3490
	goto L344
L647:
	;
	v3388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3387)+uint32(_c_F_BootstrapModeMain[54]))))
	v3389 = v3388
	goto L649
L648:
	;
	v3389 = int32(1)
	goto L649
L649:
	;
	v3391 = v3365 << (uint(int32(1)) % 32)
	v3394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3391)+uint32(_c_F_BootstrapModeMain[55]))))
	if v3394 != 0 {
		goto L650
	} else {
		goto L651
	}
L650:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+68)) = v3368
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+64)) = v3365
	goto L652
L651:
	;
	goto L652
L652:
	;
	v3398 = v3389 & int32(255)
	v3401 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3391)+uint32(_c_F_BootstrapModeMain[56]))))
	v3402 = v3398 + v3401
	v3407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3402<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v3407 != v3365 {
		goto L653
	} else {
		goto L654
	}
L653:
	;
	v3411 = v3389
	v3414 = v3365
	v3415 = v3398
	goto L656
L654:
	;
	v3469 = v3402
	goto L655
L655:
	;
	v3486 = int32(1)
	v3490 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3469<<(uint(v3486)%32))+uint32(_c_F_BootstrapModeMain[58]))))
	v3492 = v3368 + v3486
	if v3352 != v3492 {
		v3365 = v3490
		v3368 = v3492
		goto L645
	} else {
		goto L662
	}
L656:
	;
	v3439 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3414<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[59]))))
	if int32(128) <= v3439 {
		goto L658
	} else {
		goto L659
	}
L657:
	;
	v3469 = v3451
	goto L655
L658:
	;
	v3442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3415)+uint32(_c_F_BootstrapModeMain[60]))))
	v3443 = v3442
	goto L660
L659:
	;
	v3443 = v3411
	goto L660
L660:
	;
	v3445 = v3443 & int32(255)
	v3446 = int32(1)
	v3450 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3439<<(uint(v3446)%32))+uint32(_c_F_BootstrapModeMain[56]))))
	v3451 = v3445 + v3450
	v3456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3451<<(uint(v3446)%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v3456 != v3439&int32(_a_F_BootstrapModeMain_59) {
		v3411 = v3443
		v3414 = v3439
		v3415 = v3445
		goto L656
	} else {
		goto L661
	}
L661:
	;
	goto L657
L662:
	;
	goto L646
L663:
	;
	v3537 = v3497
	v3538 = v3531
	goto L664
L664:
	;
	v3560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3537))))
	if v3560 != 0 {
		goto L666
	} else {
		goto L667
	}
L665:
	;
	v1577 = v3497
	v1581 = v3665
	v1584 = v3523
	goto L361
L666:
	;
	v3561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3560)+uint32(_c_F_BootstrapModeMain[54]))))
	v3562 = v3561
	goto L668
L667:
	;
	v3562 = int32(1)
	goto L668
L668:
	;
	v3567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3538<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[55]))))
	if v3567 != 0 {
		goto L669
	} else {
		goto L670
	}
L669:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+68)) = v3537
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+64)) = v3538
	goto L671
L670:
	;
	goto L671
L671:
	;
	v3571 = v3562 & int32(255)
	v3572 = int32(1)
	v3576 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3538<<(uint(v3572)%32))+uint32(_c_F_BootstrapModeMain[56]))))
	v3577 = v3571 + v3576
	v3582 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3577<<(uint(v3572)%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v3582 != v3538 {
		goto L672
	} else {
		goto L673
	}
L672:
	;
	v3586 = v3562
	v3589 = v3538
	v3590 = v3571
	goto L675
L673:
	;
	v3644 = v3577
	goto L674
L674:
	;
	v3661 = int32(1)
	v3665 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3644<<(uint(v3661)%32))+uint32(_c_F_BootstrapModeMain[58]))))
	v3667 = v3537 + v3661
	if v3667 != v3523 {
		v3537 = v3667
		v3538 = v3665
		goto L664
	} else {
		goto L681
	}
L675:
	;
	v3614 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3589<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[59]))))
	if int32(128) <= v3614 {
		goto L677
	} else {
		goto L678
	}
L676:
	;
	v3644 = v3626
	goto L674
L677:
	;
	v3617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3590)+uint32(_c_F_BootstrapModeMain[60]))))
	v3618 = v3617
	goto L679
L678:
	;
	v3618 = v3586
	goto L679
L679:
	;
	v3620 = v3618 & int32(255)
	v3621 = int32(1)
	v3625 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3614<<(uint(v3621)%32))+uint32(_c_F_BootstrapModeMain[56]))))
	v3626 = v3620 + v3625
	v3631 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3626<<(uint(v3621)%32))+uint32(_c_F_BootstrapModeMain[57]))))
	if v3631 != v3614&int32(_a_F_BootstrapModeMain_59) {
		v3586 = v3618
		v3589 = v3614
		v3590 = v3620
		goto L675
	} else {
		goto L680
	}
L680:
	;
	goto L676
L681:
	;
	goto L665
L682:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L683:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L684:
	;
	v3742 = int32(0)
	v3750 = v3742
	v3751 = v3742
	goto L318
L685:
	;
	goto L686
L686:
	;
	if v3719 == int32(256) {
		v6057 = v3724
		v6058 = v3725
		v6061 = v3728
		goto L287
	} else {
		goto L687
	}
L687:
	;
	if base.Ui32(int32(282)) < base.Ui32(v3719) {
		v3750 = v3719
		v3751 = int32(2)
		goto L318
	} else {
		goto L688
	}
L688:
	;
	v3749 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3719)+uint32(_c_F_BootstrapModeMain[61]))))
	v3750 = v3719
	v3751 = v3749
	goto L318
L689:
	;
	v3755 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3752)+uint32(_c_F_BootstrapModeMain[62]))))
	if v3751 != v3755 {
		v3773 = v3717
		v3775 = v3750
		v3777 = v3721
		v3780 = v3724
		v3781 = v3725
		v3783 = v3727
		v3784 = v3728
		v3786 = v3730
		v3787 = v3731
		v3788 = v3732
		v3792 = v3736
		goto L316
	} else {
		goto L690
	}
L690:
	;
	v3757 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3752)+uint32(_c_F_BootstrapModeMain[63]))))
	if int32(0) < v3757 {
		goto L691
	} else {
		goto L692
	}
L691:
	;
	v3760 = *(*int32)(unsafe.Add(mBase, uint32(v3724)+1132))
	*(*int32)(unsafe.Add(mBase, uint32(v3721)+4)) = v3760
	v6017 = v3717
	v6019 = int32(-2)
	v6021 = v3721 + int32(4)
	v6024 = v3724
	v6025 = v3725
	v6027 = v3757
	v6028 = v3728
	v6030 = v3730
	v6031 = v3731
	v6032 = v3732
	v6036 = v3736 - base.B2i32(v3736 != int32(0))
	goto L314
L692:
	;
	goto L693
L693:
	;
	v3804 = v3717
	v3806 = v3750
	v3808 = v3721
	v3811 = v3724
	v3812 = v3725
	v3814 = int32(0) - v3757
	v3815 = v3728
	v3817 = v3730
	v3818 = v3731
	v3819 = v3732
	v3823 = v3736
	goto L315
L694:
	;
	v3804 = v3773
	v3806 = v3775
	v3808 = v3777
	v3811 = v3780
	v3812 = v3781
	v3814 = v3798
	v3815 = v3784
	v3817 = v3786
	v3818 = v3787
	v3819 = v3788
	v3823 = v3792
	goto L315
L695:
	;
	v5989 = v3808 - v3830<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v5989)+4)) = v5963
	v5992 = v5989 + int32(4)
	v5993 = v3815 - v3830
	v5994 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5993))))
	v5997 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3814)+uint32(_c_F_BootstrapModeMain[64]))))
	v6000 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5997)+uint32(_c_F_BootstrapModeMain[65]))))
	v6001 = v5994 + v6000
	if base.Ui32(int32(169)) < base.Ui32(v6001) {
		goto L1176
	} else {
		goto L1177
	}
L696:
	;
	v5958 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5959 = F_pstrdup(m, v5958)
	mBase = m.M
	v5960 = m.ExcPending
	if v5960 != 0 {
		goto L1
	} else {
		goto L1175
	}
L697:
	;
	v5955 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5956 = F_pstrdup(m, v5955)
	mBase = m.M
	v5957 = m.ExcPending
	if v5957 != 0 {
		goto L1
	} else {
		goto L1174
	}
L698:
	;
	v5952 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5953 = F_pstrdup(m, v5952)
	mBase = m.M
	v5954 = m.ExcPending
	if v5954 != 0 {
		goto L1
	} else {
		goto L1173
	}
L699:
	;
	v5949 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5950 = F_pstrdup(m, v5949)
	mBase = m.M
	v5951 = m.ExcPending
	if v5951 != 0 {
		goto L1
	} else {
		goto L1172
	}
L700:
	;
	v5946 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5947 = F_pstrdup(m, v5946)
	mBase = m.M
	v5948 = m.ExcPending
	if v5948 != 0 {
		goto L1
	} else {
		goto L1171
	}
L701:
	;
	v5943 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5944 = F_pstrdup(m, v5943)
	mBase = m.M
	v5945 = m.ExcPending
	if v5945 != 0 {
		goto L1
	} else {
		goto L1170
	}
L702:
	;
	v5940 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5941 = F_pstrdup(m, v5940)
	mBase = m.M
	v5942 = m.ExcPending
	if v5942 != 0 {
		goto L1
	} else {
		goto L1169
	}
L703:
	;
	v5937 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5938 = F_pstrdup(m, v5937)
	mBase = m.M
	v5939 = m.ExcPending
	if v5939 != 0 {
		goto L1
	} else {
		goto L1168
	}
L704:
	;
	v5934 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5935 = F_pstrdup(m, v5934)
	mBase = m.M
	v5936 = m.ExcPending
	if v5936 != 0 {
		goto L1
	} else {
		goto L1167
	}
L705:
	;
	v5931 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5932 = F_pstrdup(m, v5931)
	mBase = m.M
	v5933 = m.ExcPending
	if v5933 != 0 {
		goto L1
	} else {
		goto L1166
	}
L706:
	;
	v5928 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5929 = F_pstrdup(m, v5928)
	mBase = m.M
	v5930 = m.ExcPending
	if v5930 != 0 {
		goto L1
	} else {
		goto L1165
	}
L707:
	;
	v5925 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5926 = F_pstrdup(m, v5925)
	mBase = m.M
	v5927 = m.ExcPending
	if v5927 != 0 {
		goto L1
	} else {
		goto L1164
	}
L708:
	;
	v5922 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5923 = F_pstrdup(m, v5922)
	mBase = m.M
	v5924 = m.ExcPending
	if v5924 != 0 {
		goto L1
	} else {
		goto L1163
	}
L709:
	;
	v5919 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5920 = F_pstrdup(m, v5919)
	mBase = m.M
	v5921 = m.ExcPending
	if v5921 != 0 {
		goto L1
	} else {
		goto L1162
	}
L710:
	;
	v5916 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5917 = F_pstrdup(m, v5916)
	mBase = m.M
	v5918 = m.ExcPending
	if v5918 != 0 {
		goto L1
	} else {
		goto L1161
	}
L711:
	;
	v5913 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5914 = F_pstrdup(m, v5913)
	mBase = m.M
	v5915 = m.ExcPending
	if v5915 != 0 {
		goto L1
	} else {
		goto L1160
	}
L712:
	;
	v5910 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5911 = F_pstrdup(m, v5910)
	mBase = m.M
	v5912 = m.ExcPending
	if v5912 != 0 {
		goto L1
	} else {
		goto L1159
	}
L713:
	;
	v5907 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5908 = F_pstrdup(m, v5907)
	mBase = m.M
	v5909 = m.ExcPending
	if v5909 != 0 {
		goto L1
	} else {
		goto L1158
	}
L714:
	;
	v5904 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5905 = F_pstrdup(m, v5904)
	mBase = m.M
	v5906 = m.ExcPending
	if v5906 != 0 {
		goto L1
	} else {
		goto L1157
	}
L715:
	;
	v5903 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5963 = v5903
	goto L695
L716:
	;
	v5825 = int32(_a_F_BootstrapModeMain_88)
	v5827 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[66]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[66])) = v5827 + int32(1)
	v5831 = m.G0
	v5833 = v5831 - int32(32)
	m.G0 = v5833
	v5837 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5838 = m.ExcPending
	if v5838 != 0 {
		goto L1
	} else {
		goto L1145
	}
L717:
	;
	v5741 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5742 = int32(_a_F_BootstrapModeMain_88)
	v5744 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[66]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[66])) = v5744 + int32(1)
	v5748 = m.G0
	v5750 = v5748 - int32(48)
	m.G0 = v5750
	v5754 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5755 = m.ExcPending
	if v5755 != 0 {
		goto L1
	} else {
		goto L1130
	}
L718:
	;
	v5735 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5739 = F_strtox_2(m, v5735, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L1129
L719:
	;
	v5963 = int32(2)
	goto L695
L720:
	;
	v5963 = int32(3)
	goto L695
L721:
	;
	v4975 = int32(_a_F_BootstrapModeMain_89)
	v4977 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67]))
	v4979 = v4977 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67])) = v4979
	if int32(41) <= v4979 {
		goto L283
	} else {
		goto L1009
	}
L722:
	;
	v4974 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v5963 = v4974
	goto L695
L723:
	;
	v5963 = int32(0)
	goto L695
L724:
	;
	v5963 = int32(1)
	goto L695
L725:
	;
	v4947 = F_palloc0(m, int32(36))
	mBase = m.M
	v4948 = m.ExcPending
	if v4948 != 0 {
		goto L1
	} else {
		goto L1006
	}
L726:
	;
	v4938 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+96)) = v4938
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+108)) = v4938
	v4944 = F_list_make1_impl(m, int32(1), v3811+int32(96))
	mBase = m.M
	v4945 = m.ExcPending
	if v4945 != 0 {
		goto L1
	} else {
		goto L1005
	}
L727:
	;
	v4934 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(8))))
	v4935 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v4936 = F_lappend(m, v4934, v4935)
	mBase = m.M
	v4937 = m.ExcPending
	if v4937 != 0 {
		goto L1
	} else {
		goto L1004
	}
L728:
	;
	v4812 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	if v4812 == int32(0) {
		goto L981
	} else {
		goto L982
	}
L729:
	;
	v4694 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4695 = m.ExcPending
	if v4695 != 0 {
		goto L1
	} else {
		goto L948
	}
L730:
	;
	v4570 = F_palloc0(m, int32(72))
	mBase = m.M
	v4571 = m.ExcPending
	if v4571 != 0 {
		goto L1
	} else {
		goto L926
	}
L731:
	;
	v4449 = F_palloc0(m, int32(72))
	mBase = m.M
	v4450 = m.ExcPending
	if v4450 != 0 {
		goto L1
	} else {
		goto L904
	}
L732:
	;
	v4354 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[66]))
	v4356 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67]))
	if v4354 != v4356 {
		goto L285
	} else {
		goto L874
	}
L733:
	;
	v4321 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	if v4321 == int32(0) {
		goto L864
	} else {
		goto L865
	}
L734:
	;
	v4163 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	if v4163 == int32(0) {
		goto L821
	} else {
		goto L822
	}
L735:
	;
	v4142 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4142
	v4145 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	F_MemoryContextReset(m, v4145)
	mBase = m.M
	v4147 = m.ExcPending
	if v4147 != 0 {
		goto L1
	} else {
		goto L813
	}
L736:
	;
	v4089 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	if v4089 == int32(0) {
		goto L799
	} else {
		goto L800
	}
L737:
	;
	v4047 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	if v4047 == int32(0) {
		goto L786
	} else {
		goto L787
	}
L738:
	;
	v3839 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	if v3839 == int32(0) {
		goto L739
	} else {
		goto L740
	}
L739:
	;
	v3844 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	v3849 = F_AllocSetContextCreateInternal(m, v3844, int32(_a_F_BootstrapModeMain_90), int32(0), int32(_a_F_BootstrapModeMain_23), int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v3850 = m.ExcPending
	if v3850 != 0 {
		goto L1
	} else {
		goto L742
	}
L740:
	;
	v3852 = v3839
	goto L741
L741:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v3852
	v3855 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v3857 = m.G0
	v3859 = v3857 - int32(48)
	m.G0 = v3859
	v3861 = F_strlen(m, v3855)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v3861) {
		goto L743
	} else {
		goto L744
	}
L742:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68])) = v3849
	v3852 = v3849
	goto L741
L743:
	;
	v3864 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3855)+63)) = uint8(v3864)
	goto L745
L744:
	;
	goto L745
L745:
	;
	v3867 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[71]))
	if v3867 == int32(0) {
		goto L746
	} else {
		goto L747
	}
L746:
	;
	F_populate_typ_list(m)
	mBase = m.M
	v3871 = m.ExcPending
	if v3871 != 0 {
		goto L1
	} else {
		goto L749
	}
L747:
	;
	goto L748
L748:
	;
	v3873 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	if v3873 != 0 {
		goto L750
	} else {
		goto L751
	}
L749:
	;
	goto L748
L750:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		goto L1
	} else {
		goto L753
	}
L751:
	;
	goto L752
L752:
	;
	v3879 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v3880 = m.ExcPending
	if v3880 != 0 {
		goto L1
	} else {
		goto L754
	}
L753:
	;
	goto L752
L754:
	;
	if v3879 != 0 {
		goto L755
	} else {
		goto L756
	}
L755:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3859)+36)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v3859)+32)) = v3855
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_92), v3859+int32(32))
	mBase = m.M
	v3888 = m.ExcPending
	if v3888 != 0 {
		goto L1
	} else {
		goto L758
	}
L756:
	;
	goto L757
L757:
	;
	v3897 = F_makeRangeVar(m, int32(0), v3855, int32(-1))
	mBase = m.M
	v3898 = m.ExcPending
	if v3898 != 0 {
		goto L1
	} else {
		goto L760
	}
L758:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(458), int32(_a_F_BootstrapModeMain_93))
	mBase = m.M
	v3893 = m.ExcPending
	if v3893 != 0 {
		goto L1
	} else {
		goto L759
	}
L759:
	;
	goto L757
L760:
	;
	v3900 = F_table_openrv(m, v3897, int32(0))
	mBase = m.M
	v3901 = m.ExcPending
	if v3901 != 0 {
		goto L1
	} else {
		goto L761
	}
L761:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72])) = v3900
	v3904 = *(*int32)(unsafe.Add(mBase, uint32(v3900)+48))
	v3905 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3904)+120)))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67])) = v3905
	if int32(0) < v3905 {
		goto L762
	} else {
		goto L763
	}
L762:
	;
	v3917 = int32(0)
	goto L765
L763:
	;
	goto L764
L764:
	;
	m.G0 = v3859 + int32(48)
	v4026 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4026
	v4029 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	F_MemoryContextReset(m, v4029)
	mBase = m.M
	v4031 = m.ExcPending
	if v4031 != 0 {
		goto L1
	} else {
		goto L778
	}
L765:
	;
	v3938 = v3917 << (uint(int32(2)) % 32)
	v3941 = *(*int32)(unsafe.Add(mBase, uint32(v3938)+uint32(_c_F_BootstrapModeMain[73])))
	if v3941 == int32(0) {
		goto L767
	} else {
		goto L768
	}
L766:
	;
	goto L764
L767:
	;
	v3945 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[74]))
	v3947 = F_MemoryContextAllocZero(m, v3945, int32(100))
	mBase = m.M
	v3948 = m.ExcPending
	if v3948 != 0 {
		goto L1
	} else {
		goto L770
	}
L768:
	;
	v3950 = v3941
	goto L769
L769:
	;
	v3952 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	v3953 = *(*int32)(unsafe.Add(mBase, uint32(v3952)+52))
	v3954 = *(*int32)(unsafe.Add(mBase, uint32(v3953)))
	v3958 = int32(100)
	base.MemoryCopy(m, v3950, v3953+v3954<<(uint(int32(4))%32)+v3917*v3958+int32(20), v3958)
	v3965 = *(*int32)(unsafe.Add(mBase, uint32(v3938)+uint32(_c_F_BootstrapModeMain[73])))
	v3968 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v3969 = m.ExcPending
	if v3969 != 0 {
		goto L1
	} else {
		goto L771
	}
L770:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3938)+uint32(_c_F_BootstrapModeMain[73]))) = v3947
	v3950 = v3947
	goto L769
L771:
	;
	if v3968 != 0 {
		goto L772
	} else {
		goto L773
	}
L772:
	;
	v3970 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3965)+72)))
	v3971 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3965)+74)))
	v3972 = *(*int32)(unsafe.Add(mBase, uint32(v3965)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v3859+int32(16)))) = v3972
	*(*int32)(unsafe.Add(mBase, uint32(v3859)+12)) = v3971
	*(*int32)(unsafe.Add(mBase, uint32(v3859)+8)) = v3970
	*(*int32)(unsafe.Add(mBase, uint32(v3859)+4)) = v3965 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v3859))) = v3917
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_94), v3859)
	mBase = m.M
	v3982 = m.ExcPending
	if v3982 != 0 {
		goto L1
	} else {
		goto L775
	}
L773:
	;
	goto L774
L774:
	;
	v3991 = v3917 + int32(1)
	v3993 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67]))
	if v3991 < v3993 {
		v3917 = v3991
		goto L765
	} else {
		goto L777
	}
L775:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(475), int32(_a_F_BootstrapModeMain_93))
	mBase = m.M
	v3987 = m.ExcPending
	if v3987 != 0 {
		goto L1
	} else {
		goto L776
	}
L776:
	;
	goto L774
L777:
	;
	goto L766
L778:
	;
	v4033 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	if v4033 != 0 {
		goto L779
	} else {
		goto L780
	}
L779:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4035 = m.ExcPending
	if v4035 != 0 {
		goto L1
	} else {
		goto L782
	}
L780:
	;
	goto L781
L781:
	;
	v4036 = int32(0)
	v4037 = F_isatty(m, v4036)
	mBase = m.M
	if v4037 == v4036 {
		v5963 = v3835
		goto L695
	} else {
		goto L783
	}
L782:
	;
	goto L781
L783:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_95), int32(0))
	mBase = m.M
	v4043 = m.ExcPending
	if v4043 != 0 {
		goto L1
	} else {
		goto L784
	}
L784:
	;
	v4044 = F_fflush(m, v3817)
	mBase = m.M
	v4045 = m.ExcPending
	if v4045 != 0 {
		goto L1
	} else {
		goto L785
	}
L785:
	;
	v5963 = v3835
	goto L695
L786:
	;
	v4052 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	v4057 = F_AllocSetContextCreateInternal(m, v4052, int32(_a_F_BootstrapModeMain_90), int32(0), int32(_a_F_BootstrapModeMain_23), int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v4058 = m.ExcPending
	if v4058 != 0 {
		goto L1
	} else {
		goto L789
	}
L787:
	;
	v4060 = v4047
	goto L788
L788:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4060
	v4063 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	F_closerel(m, v4063)
	mBase = m.M
	v4065 = m.ExcPending
	if v4065 != 0 {
		goto L1
	} else {
		goto L790
	}
L789:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68])) = v4057
	v4060 = v4057
	goto L788
L790:
	;
	v4068 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4068
	v4071 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	F_MemoryContextReset(m, v4071)
	mBase = m.M
	v4073 = m.ExcPending
	if v4073 != 0 {
		goto L1
	} else {
		goto L791
	}
L791:
	;
	v4075 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	if v4075 != 0 {
		goto L792
	} else {
		goto L793
	}
L792:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4077 = m.ExcPending
	if v4077 != 0 {
		goto L1
	} else {
		goto L795
	}
L793:
	;
	goto L794
L794:
	;
	v4078 = int32(0)
	v4079 = F_isatty(m, v4078)
	mBase = m.M
	if v4079 == v4078 {
		v5963 = v3835
		goto L695
	} else {
		goto L796
	}
L795:
	;
	goto L794
L796:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_95), int32(0))
	mBase = m.M
	v4085 = m.ExcPending
	if v4085 != 0 {
		goto L1
	} else {
		goto L797
	}
L797:
	;
	v4086 = F_fflush(m, v3817)
	mBase = m.M
	v4087 = m.ExcPending
	if v4087 != 0 {
		goto L1
	} else {
		goto L798
	}
L798:
	;
	v5963 = v3835
	goto L695
L799:
	;
	v4094 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	v4099 = F_AllocSetContextCreateInternal(m, v4094, int32(_a_F_BootstrapModeMain_90), int32(0), int32(_a_F_BootstrapModeMain_23), int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v4100 = m.ExcPending
	if v4100 != 0 {
		goto L1
	} else {
		goto L802
	}
L800:
	;
	v4102 = v4089
	goto L801
L801:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4102
	v4106 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67])) = v4106
	v4110 = F_errstart(m, int32(11), v4106)
	mBase = m.M
	v4111 = m.ExcPending
	if v4111 != 0 {
		goto L1
	} else {
		goto L803
	}
L802:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68])) = v4099
	v4102 = v4099
	goto L801
L803:
	;
	if v4110 == int32(0) {
		v5963 = v3835
		goto L695
	} else {
		goto L804
	}
L804:
	;
	v4116 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(12))))
	v4119 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(8))))
	v4122 = *(*int64)(unsafe.Add(mBase, uint32(v3808-int32(20))))
	*(*int64)(unsafe.Add(mBase, uint32(v3811)+8)) = v4122
	if v4119 != 0 {
		goto L805
	} else {
		goto L806
	}
L805:
	;
	v4126 = int32(_a_F_BootstrapModeMain_96)
	goto L807
L806:
	;
	v4126 = int32(_a_F_BootstrapModeMain_97)
	goto L807
L807:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+4)) = v4126
	if v4116 != 0 {
		goto L808
	} else {
		goto L809
	}
L808:
	;
	v4130 = int32(_a_F_BootstrapModeMain_98)
	goto L810
L809:
	;
	v4130 = int32(_a_F_BootstrapModeMain_97)
	goto L810
L810:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3811))) = v4130
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_99), v3811)
	mBase = m.M
	v4134 = m.ExcPending
	if v4134 != 0 {
		goto L1
	} else {
		goto L811
	}
L811:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(166), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v4139 = m.ExcPending
	if v4139 != 0 {
		goto L1
	} else {
		goto L812
	}
L812:
	;
	v5963 = v3835
	goto L695
L813:
	;
	v4149 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	if v4149 != 0 {
		goto L814
	} else {
		goto L815
	}
L814:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4151 = m.ExcPending
	if v4151 != 0 {
		goto L1
	} else {
		goto L817
	}
L815:
	;
	goto L816
L816:
	;
	v4152 = int32(0)
	v4153 = F_isatty(m, v4152)
	mBase = m.M
	if v4153 == v4152 {
		v5963 = v3835
		goto L695
	} else {
		goto L818
	}
L817:
	;
	goto L816
L818:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_95), int32(0))
	mBase = m.M
	v4159 = m.ExcPending
	if v4159 != 0 {
		goto L1
	} else {
		goto L819
	}
L819:
	;
	v4160 = F_fflush(m, v3817)
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		goto L1
	} else {
		goto L820
	}
L820:
	;
	v5963 = v3835
	goto L695
L821:
	;
	v4168 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	v4173 = F_AllocSetContextCreateInternal(m, v4168, int32(_a_F_BootstrapModeMain_90), int32(0), int32(_a_F_BootstrapModeMain_23), int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v4174 = m.ExcPending
	if v4174 != 0 {
		goto L1
	} else {
		goto L824
	}
L822:
	;
	v4176 = v4163
	goto L823
L823:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4176
	v4180 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67]))
	v4182 = F_CreateTupleDesc(m, v4180, int32(_a_F_BootstrapModeMain_56))
	mBase = m.M
	v4183 = m.ExcPending
	if v4183 != 0 {
		goto L1
	} else {
		goto L825
	}
L824:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68])) = v4173
	v4176 = v4173
	goto L823
L825:
	;
	v4186 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(24))))
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(28))))
	if v4189 != 0 {
		goto L827
	} else {
		goto L828
	}
L826:
	;
	v4300 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4300
	v4303 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	F_MemoryContextReset(m, v4303)
	mBase = m.M
	v4305 = m.ExcPending
	if v4305 != 0 {
		goto L1
	} else {
		goto L856
	}
L827:
	;
	v4191 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	if v4191 != 0 {
		goto L830
	} else {
		goto L831
	}
L828:
	;
	goto L829
L829:
	;
	v4252 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(36))))
	if v4186 != 0 {
		goto L848
	} else {
		goto L849
	}
L830:
	;
	v4194 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4195 = m.ExcPending
	if v4195 != 0 {
		goto L1
	} else {
		goto L833
	}
L831:
	;
	goto L832
L832:
	;
	v4211 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(36))))
	if v4186 != 0 {
		goto L840
	} else {
		goto L841
	}
L833:
	;
	if v4194 != 0 {
		goto L834
	} else {
		goto L835
	}
L834:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_102), int32(0))
	mBase = m.M
	v4199 = m.ExcPending
	if v4199 != 0 {
		goto L1
	} else {
		goto L837
	}
L835:
	;
	goto L836
L836:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v4207 = m.ExcPending
	if v4207 != 0 {
		goto L1
	} else {
		goto L839
	}
L837:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(202), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v4204 = m.ExcPending
	if v4204 != 0 {
		goto L1
	} else {
		goto L838
	}
L838:
	;
	goto L836
L839:
	;
	goto L832
L840:
	;
	v4215 = int32(1664)
	goto L842
L841:
	;
	v4215 = int32(0)
	goto L842
L842:
	;
	v4218 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(32))))
	v4219 = int32(0)
	v4222 = int32(112)
	v4225 = int32(1)
	v4232 = F_heap_create(m, v4211, int32(11), v4215, v4218, v4219, int32(2), v4182, int32(114), v4222, base.B2i32(v4186 != v4219), v4225, v4225, v3811+v4222, v3811+int32(124), v4225)
	mBase = m.M
	v4233 = m.ExcPending
	if v4233 != 0 {
		goto L1
	} else {
		goto L843
	}
L843:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72])) = v4232
	v4237 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4238 = m.ExcPending
	if v4238 != 0 {
		goto L1
	} else {
		goto L844
	}
L844:
	;
	if v4237 == int32(0) {
		goto L826
	} else {
		goto L845
	}
L845:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_103), int32(0))
	mBase = m.M
	v4244 = m.ExcPending
	if v4244 != 0 {
		goto L1
	} else {
		goto L846
	}
L846:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(221), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v4249 = m.ExcPending
	if v4249 != 0 {
		goto L1
	} else {
		goto L847
	}
L847:
	;
	goto L826
L848:
	;
	v4256 = int32(1664)
	goto L850
L849:
	;
	v4256 = int32(0)
	goto L850
L850:
	;
	v4259 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(32))))
	v4262 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(20))))
	v4263 = int32(0)
	v4270 = base.B2i32(v4186 != v4263)
	v4278 = F_heap_create_with_catalog(m, v4252, int32(11), v4256, v4259, v4262, v4263, int32(10), int32(2), v4182, v4263, int32(114), int32(112), v4270, v4270, v4263, v4263, v4263, int32(1), v4263, v4263, v4263)
	mBase = m.M
	v4279 = m.ExcPending
	if v4279 != 0 {
		goto L1
	} else {
		goto L851
	}
L851:
	;
	v4282 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4283 = m.ExcPending
	if v4283 != 0 {
		goto L1
	} else {
		goto L852
	}
L852:
	;
	if v4282 == int32(0) {
		goto L826
	} else {
		goto L853
	}
L853:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+16)) = v4278
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_104), v3811+int32(16))
	mBase = m.M
	v4291 = m.ExcPending
	if v4291 != 0 {
		goto L1
	} else {
		goto L854
	}
L854:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(248), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v4296 = m.ExcPending
	if v4296 != 0 {
		goto L1
	} else {
		goto L855
	}
L855:
	;
	goto L826
L856:
	;
	v4307 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	if v4307 != 0 {
		goto L857
	} else {
		goto L858
	}
L857:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4309 = m.ExcPending
	if v4309 != 0 {
		goto L1
	} else {
		goto L860
	}
L858:
	;
	goto L859
L859:
	;
	v4310 = int32(0)
	v4311 = F_isatty(m, v4310)
	mBase = m.M
	if v4311 == v4310 {
		v5963 = v3835
		goto L695
	} else {
		goto L861
	}
L860:
	;
	goto L859
L861:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_95), int32(0))
	mBase = m.M
	v4317 = m.ExcPending
	if v4317 != 0 {
		goto L1
	} else {
		goto L862
	}
L862:
	;
	v4318 = F_fflush(m, v3817)
	mBase = m.M
	v4319 = m.ExcPending
	if v4319 != 0 {
		goto L1
	} else {
		goto L863
	}
L863:
	;
	v5963 = v3835
	goto L695
L864:
	;
	v4326 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	v4331 = F_AllocSetContextCreateInternal(m, v4326, int32(_a_F_BootstrapModeMain_90), int32(0), int32(_a_F_BootstrapModeMain_23), int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v4332 = m.ExcPending
	if v4332 != 0 {
		goto L1
	} else {
		goto L867
	}
L865:
	;
	v4334 = v4321
	goto L866
L866:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4334
	v4339 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4340 = m.ExcPending
	if v4340 != 0 {
		goto L1
	} else {
		goto L868
	}
L867:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68])) = v4331
	v4334 = v4331
	goto L866
L868:
	;
	if v4339 != 0 {
		goto L869
	} else {
		goto L870
	}
L869:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_105), int32(0))
	mBase = m.M
	v4344 = m.ExcPending
	if v4344 != 0 {
		goto L1
	} else {
		goto L872
	}
L870:
	;
	goto L871
L871:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[66])) = int32(0)
	v5963 = v3835
	goto L695
L872:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(258), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v4349 = m.ExcPending
	if v4349 != 0 {
		goto L1
	} else {
		goto L873
	}
L873:
	;
	goto L871
L874:
	;
	v4359 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	if v4359 == int32(0) {
		goto L284
	} else {
		goto L875
	}
L875:
	;
	v4362 = m.G0
	v4364 = v4362 - int32(16)
	m.G0 = v4364
	v4368 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4369 = m.ExcPending
	if v4369 != 0 {
		goto L1
	} else {
		goto L876
	}
L876:
	;
	if v4368 != 0 {
		goto L877
	} else {
		goto L878
	}
L877:
	;
	v4371 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67]))
	*(*int32)(unsafe.Add(mBase, uint32(v4364))) = v4371
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_106), v4364)
	mBase = m.M
	v4375 = m.ExcPending
	if v4375 != 0 {
		goto L1
	} else {
		goto L880
	}
L878:
	;
	goto L879
L879:
	;
	v4382 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67]))
	v4384 = F_CreateTupleDesc(m, v4382, int32(_a_F_BootstrapModeMain_56))
	mBase = m.M
	v4385 = m.ExcPending
	if v4385 != 0 {
		goto L1
	} else {
		goto L882
	}
L880:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(635), int32(_a_F_BootstrapModeMain_107))
	mBase = m.M
	v4380 = m.ExcPending
	if v4380 != 0 {
		goto L1
	} else {
		goto L881
	}
L881:
	;
	goto L879
L882:
	;
	v4388 = F_heap_form_tuple(m, v4384, int32(_a_F_BootstrapModeMain_108), int32(_a_F_BootstrapModeMain_109))
	mBase = m.M
	v4389 = m.ExcPending
	if v4389 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	F_pfree(m, v4384)
	mBase = m.M
	v4391 = m.ExcPending
	if v4391 != 0 {
		goto L1
	} else {
		goto L884
	}
L884:
	;
	v4393 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	F_simple_heap_insert(m, v4393, v4388)
	mBase = m.M
	v4395 = m.ExcPending
	if v4395 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	F_pfree(m, v4388)
	mBase = m.M
	v4397 = m.ExcPending
	if v4397 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	v4400 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4401 = m.ExcPending
	if v4401 != 0 {
		goto L1
	} else {
		goto L887
	}
L887:
	;
	if v4400 != 0 {
		goto L888
	} else {
		goto L889
	}
L888:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_110), int32(0))
	mBase = m.M
	v4405 = m.ExcPending
	if v4405 != 0 {
		goto L1
	} else {
		goto L891
	}
L889:
	;
	goto L890
L890:
	;
	v4412 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67]))
	v4413 = int32(0)
	if base.B2i32(v4412 <= v4413)|base.B2i32(v4412 == v4413) == v4413 {
		goto L893
	} else {
		goto L894
	}
L891:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(643), int32(_a_F_BootstrapModeMain_107))
	mBase = m.M
	v4410 = m.ExcPending
	if v4410 != 0 {
		goto L1
	} else {
		goto L892
	}
L892:
	;
	goto L890
L893:
	;
	base.MemoryFill(m, int32(_a_F_BootstrapModeMain_109), int32(0), v4412)
	goto L895
L894:
	;
	goto L895
L895:
	;
	m.G0 = v4364 + int32(16)
	v4428 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4428
	v4431 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	F_MemoryContextReset(m, v4431)
	mBase = m.M
	v4433 = m.ExcPending
	if v4433 != 0 {
		goto L1
	} else {
		goto L896
	}
L896:
	;
	v4435 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	if v4435 != 0 {
		goto L897
	} else {
		goto L898
	}
L897:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4437 = m.ExcPending
	if v4437 != 0 {
		goto L1
	} else {
		goto L900
	}
L898:
	;
	goto L899
L899:
	;
	v4438 = int32(0)
	v4439 = F_isatty(m, v4438)
	mBase = m.M
	if v4439 == v4438 {
		v5963 = v3835
		goto L695
	} else {
		goto L901
	}
L900:
	;
	goto L899
L901:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_95), int32(0))
	mBase = m.M
	v4445 = m.ExcPending
	if v4445 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	v4446 = F_fflush(m, v3817)
	mBase = m.M
	v4447 = m.ExcPending
	if v4447 != 0 {
		goto L1
	} else {
		goto L903
	}
L903:
	;
	v5963 = v3835
	goto L695
L904:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4449))) = int32(204)
	v4455 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4456 = m.ExcPending
	if v4456 != 0 {
		goto L1
	} else {
		goto L905
	}
L905:
	;
	if v4455 != 0 {
		goto L906
	} else {
		goto L907
	}
L906:
	;
	v4459 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+48)) = v4459
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_111), v3811+int32(48))
	mBase = m.M
	v4465 = m.ExcPending
	if v4465 != 0 {
		goto L1
	} else {
		goto L909
	}
L907:
	;
	goto L908
L908:
	;
	v4472 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	if v4472 == int32(0) {
		goto L911
	} else {
		goto L912
	}
L909:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(279), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v4470 = m.ExcPending
	if v4470 != 0 {
		goto L1
	} else {
		goto L910
	}
L910:
	;
	goto L908
L911:
	;
	v4477 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	v4482 = F_AllocSetContextCreateInternal(m, v4477, int32(_a_F_BootstrapModeMain_90), int32(0), int32(_a_F_BootstrapModeMain_23), int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v4483 = m.ExcPending
	if v4483 != 0 {
		goto L1
	} else {
		goto L914
	}
L912:
	;
	v4485 = v4472
	goto L913
L913:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4485
	v4490 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4449)+4)) = v4490
	v4495 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(20))))
	v4497 = F_makeRangeVar(m, int32(0), v4495, int32(-1))
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
		goto L1
	} else {
		goto L915
	}
L914:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68])) = v4482
	v4485 = v4482
	goto L913
L915:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4449)+8)) = v4497
	v4502 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(12))))
	v4503 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4449)+16)) = v4503
	*(*int32)(unsafe.Add(mBase, uint32(v4449)+12)) = v4502
	v4508 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(4))))
	*(*uint16)(unsafe.Add(mBase, uint32(v4449)+62)) = uint16(v4503)
	*(*int32)(unsafe.Add(mBase, uint32(v4449)+20)) = v4508
	v4512 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4449)+24)) = v4512
	*(*int64)(unsafe.Add(mBase, uint32(v4449)+32)) = v4512
	*(*int64)(unsafe.Add(mBase, uint32(v4449)+40)) = v4512
	*(*int64)(unsafe.Add(mBase, uint32(v4449)+48)) = v4512
	*(*int64)(unsafe.Add(mBase, uint32(v4449)+53)) = v4512
	*(*int32)(unsafe.Add(mBase, uint32(v4449)+65)) = v4503
	*(*uint16)(unsafe.Add(mBase, uint32(v4449)+69)) = uint16(v4503)
	v4532 = F_RangeVarGetRelidExtended(m, v4497, v4503, v4503, v4503, v4503)
	mBase = m.M
	v4533 = m.ExcPending
	if v4533 != 0 {
		goto L1
	} else {
		goto L916
	}
L916:
	;
	v4536 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(28))))
	v4537 = int32(0)
	F_DefineIndex(m, v3811+int32(112), v4532, v4449, v4536, v4537, v4537, int32(-1), v4537, v4537, v4537, int32(1), v4537)
	mBase = m.M
	v4546 = m.ExcPending
	if v4546 != 0 {
		goto L1
	} else {
		goto L917
	}
L917:
	;
	v4549 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4549
	v4552 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	F_MemoryContextReset(m, v4552)
	mBase = m.M
	v4554 = m.ExcPending
	if v4554 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	v4556 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	if v4556 != 0 {
		goto L919
	} else {
		goto L920
	}
L919:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4558 = m.ExcPending
	if v4558 != 0 {
		goto L1
	} else {
		goto L922
	}
L920:
	;
	goto L921
L921:
	;
	v4559 = int32(0)
	v4560 = F_isatty(m, v4559)
	mBase = m.M
	if v4560 == v4559 {
		v5963 = v3835
		goto L695
	} else {
		goto L923
	}
L922:
	;
	goto L921
L923:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_95), int32(0))
	mBase = m.M
	v4566 = m.ExcPending
	if v4566 != 0 {
		goto L1
	} else {
		goto L924
	}
L924:
	;
	v4567 = F_fflush(m, v3817)
	mBase = m.M
	v4568 = m.ExcPending
	if v4568 != 0 {
		goto L1
	} else {
		goto L925
	}
L925:
	;
	v5963 = v3835
	goto L695
L926:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4570))) = int32(204)
	v4576 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4577 = m.ExcPending
	if v4577 != 0 {
		goto L1
	} else {
		goto L927
	}
L927:
	;
	if v4576 != 0 {
		goto L928
	} else {
		goto L929
	}
L928:
	;
	v4580 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+64)) = v4580
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_112), v3811-int32(-64))
	mBase = m.M
	v4586 = m.ExcPending
	if v4586 != 0 {
		goto L1
	} else {
		goto L931
	}
L929:
	;
	goto L930
L930:
	;
	v4593 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	if v4593 == int32(0) {
		goto L933
	} else {
		goto L934
	}
L931:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(332), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v4591 = m.ExcPending
	if v4591 != 0 {
		goto L1
	} else {
		goto L932
	}
L932:
	;
	goto L930
L933:
	;
	v4598 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	v4603 = F_AllocSetContextCreateInternal(m, v4598, int32(_a_F_BootstrapModeMain_90), int32(0), int32(_a_F_BootstrapModeMain_23), int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v4604 = m.ExcPending
	if v4604 != 0 {
		goto L1
	} else {
		goto L936
	}
L934:
	;
	v4606 = v4593
	goto L935
L935:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4606
	v4611 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4570)+4)) = v4611
	v4616 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(20))))
	v4618 = F_makeRangeVar(m, int32(0), v4616, int32(-1))
	mBase = m.M
	v4619 = m.ExcPending
	if v4619 != 0 {
		goto L1
	} else {
		goto L937
	}
L936:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68])) = v4603
	v4606 = v4603
	goto L935
L937:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4570)+8)) = v4618
	v4623 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(12))))
	v4624 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4570)+16)) = v4624
	*(*int32)(unsafe.Add(mBase, uint32(v4570)+12)) = v4623
	v4629 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(4))))
	v4630 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4570)+24)) = v4630
	*(*int32)(unsafe.Add(mBase, uint32(v4570)+20)) = v4629
	*(*int64)(unsafe.Add(mBase, uint32(v4570)+32)) = v4630
	*(*int64)(unsafe.Add(mBase, uint32(v4570)+40)) = v4630
	*(*int64)(unsafe.Add(mBase, uint32(v4570)+48)) = v4630
	*(*int32)(unsafe.Add(mBase, uint32(v4570)+56)) = v4624
	*(*int32)(unsafe.Add(mBase, uint32(v4570)+65)) = v4624
	*(*uint16)(unsafe.Add(mBase, uint32(v4570)+62)) = uint16(v4624)
	v4645 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4570)+60)) = uint8(v4645)
	*(*uint16)(unsafe.Add(mBase, uint32(v4570)+69)) = uint16(v4624)
	v4655 = F_RangeVarGetRelidExtended(m, v4618, v4624, v4624, v4624, v4624)
	mBase = m.M
	v4656 = m.ExcPending
	if v4656 != 0 {
		goto L1
	} else {
		goto L938
	}
L938:
	;
	v4659 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(28))))
	v4660 = int32(0)
	F_DefineIndex(m, v3811+int32(112), v4655, v4570, v4659, v4660, v4660, int32(-1), v4660, v4660, v4660, int32(1), v4660)
	mBase = m.M
	v4669 = m.ExcPending
	if v4669 != 0 {
		goto L1
	} else {
		goto L939
	}
L939:
	;
	v4672 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4672
	v4675 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	F_MemoryContextReset(m, v4675)
	mBase = m.M
	v4677 = m.ExcPending
	if v4677 != 0 {
		goto L1
	} else {
		goto L940
	}
L940:
	;
	v4679 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	if v4679 != 0 {
		goto L941
	} else {
		goto L942
	}
L941:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4681 = m.ExcPending
	if v4681 != 0 {
		goto L1
	} else {
		goto L944
	}
L942:
	;
	goto L943
L943:
	;
	v4682 = int32(0)
	v4683 = F_isatty(m, v4682)
	mBase = m.M
	if v4683 == v4682 {
		v5963 = v3835
		goto L695
	} else {
		goto L945
	}
L944:
	;
	goto L943
L945:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_95), int32(0))
	mBase = m.M
	v4689 = m.ExcPending
	if v4689 != 0 {
		goto L1
	} else {
		goto L946
	}
L946:
	;
	v4690 = F_fflush(m, v3817)
	mBase = m.M
	v4691 = m.ExcPending
	if v4691 != 0 {
		goto L1
	} else {
		goto L947
	}
L947:
	;
	v5963 = v3835
	goto L695
L948:
	;
	if v4694 != 0 {
		goto L949
	} else {
		goto L950
	}
L949:
	;
	v4696 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+80)) = v4696
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_113), v3811+int32(80))
	mBase = m.M
	v4702 = m.ExcPending
	if v4702 != 0 {
		goto L1
	} else {
		goto L952
	}
L950:
	;
	goto L951
L951:
	;
	v4709 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	if v4709 == int32(0) {
		goto L954
	} else {
		goto L955
	}
L952:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(382), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v4707 = m.ExcPending
	if v4707 != 0 {
		goto L1
	} else {
		goto L953
	}
L953:
	;
	goto L951
L954:
	;
	v4714 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	v4719 = F_AllocSetContextCreateInternal(m, v4714, int32(_a_F_BootstrapModeMain_90), int32(0), int32(_a_F_BootstrapModeMain_23), int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v4720 = m.ExcPending
	if v4720 != 0 {
		goto L1
	} else {
		goto L957
	}
L955:
	;
	v4722 = v4709
	goto L956
L956:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4722
	v4725 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v4728 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(12))))
	v4731 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(8))))
	v4732 = m.G0
	v4734 = v4732 - int32(32)
	m.G0 = v4734
	v4738 = F_makeRangeVar(m, int32(0), v4725, int32(-1))
	mBase = m.M
	v4739 = m.ExcPending
	if v4739 != 0 {
		goto L1
	} else {
		goto L960
	}
L957:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68])) = v4719
	v4722 = v4719
	goto L956
L958:
	;
	v4760 = int32(0)
	v4764 = F_create_toast_table(m, v4741, v4728, v4731, v4760, int32(8), v4760, v4760)
	mBase = m.M
	v4765 = m.ExcPending
	if v4765 != 0 {
		goto L1
	} else {
		goto L965
	}
L959:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4750 = m.ExcPending
	if v4750 != 0 {
		goto L1
	} else {
		goto L962
	}
L960:
	;
	v4741 = F_table_openrv(m, v4738, int32(8))
	mBase = m.M
	v4742 = m.ExcPending
	if v4742 != 0 {
		goto L1
	} else {
		goto L961
	}
L961:
	;
	v4743 = *(*int32)(unsafe.Add(mBase, uint32(v4741)+48))
	v4744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4743)+119)))
	switch v4744 - int32(109) {
	case 0, 5:
		goto L958
	default:
		goto L959
	}
L962:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4734))) = v4725
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_114), v4734)
	mBase = m.M
	v4754 = m.ExcPending
	if v4754 != 0 {
		goto L1
	} else {
		goto L963
	}
L963:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_115), int32(107), int32(_a_F_BootstrapModeMain_116))
	mBase = m.M
	v4759 = m.ExcPending
	if v4759 != 0 {
		goto L1
	} else {
		goto L964
	}
L964:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L965:
	;
	if v4764 == int32(0) {
		goto L966
	} else {
		goto L967
	}
L966:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4771 = m.ExcPending
	if v4771 != 0 {
		goto L1
	} else {
		goto L969
	}
L967:
	;
	goto L968
L968:
	;
	F_relation_close(m, v4741, int32(0))
	mBase = m.M
	v4785 = m.ExcPending
	if v4785 != 0 {
		goto L1
	} else {
		goto L972
	}
L969:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4734)+16)) = v4725
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_117), v4734+int32(16))
	mBase = m.M
	v4777 = m.ExcPending
	if v4777 != 0 {
		goto L1
	} else {
		goto L970
	}
L970:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_115), int32(113), int32(_a_F_BootstrapModeMain_116))
	mBase = m.M
	v4782 = m.ExcPending
	if v4782 != 0 {
		goto L1
	} else {
		goto L971
	}
L971:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L972:
	;
	m.G0 = v4734 + int32(32)
	v4791 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4791
	v4794 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	F_MemoryContextReset(m, v4794)
	mBase = m.M
	v4796 = m.ExcPending
	if v4796 != 0 {
		goto L1
	} else {
		goto L973
	}
L973:
	;
	v4798 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	if v4798 != 0 {
		goto L974
	} else {
		goto L975
	}
L974:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4800 = m.ExcPending
	if v4800 != 0 {
		goto L1
	} else {
		goto L977
	}
L975:
	;
	goto L976
L976:
	;
	v4801 = int32(0)
	v4802 = F_isatty(m, v4801)
	mBase = m.M
	if v4802 == v4801 {
		v5963 = v3835
		goto L695
	} else {
		goto L978
	}
L977:
	;
	goto L976
L978:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_95), int32(0))
	mBase = m.M
	v4808 = m.ExcPending
	if v4808 != 0 {
		goto L1
	} else {
		goto L979
	}
L979:
	;
	v4809 = F_fflush(m, v3817)
	mBase = m.M
	v4810 = m.ExcPending
	if v4810 != 0 {
		goto L1
	} else {
		goto L980
	}
L980:
	;
	v5963 = v3835
	goto L695
L981:
	;
	v4817 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	v4822 = F_AllocSetContextCreateInternal(m, v4817, int32(_a_F_BootstrapModeMain_90), int32(0), int32(_a_F_BootstrapModeMain_23), int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v4823 = m.ExcPending
	if v4823 != 0 {
		goto L1
	} else {
		goto L984
	}
L982:
	;
	v4825 = v4812
	goto L983
L983:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4825
	v4829 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	if v4829 != 0 {
		goto L985
	} else {
		goto L986
	}
L984:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68])) = v4822
	v4825 = v4822
	goto L983
L985:
	;
	v4834 = v4829
	goto L988
L986:
	;
	goto L987
L987:
	;
	v4912 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[69]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[70])) = v4912
	v4915 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[68]))
	F_MemoryContextReset(m, v4915)
	mBase = m.M
	v4917 = m.ExcPending
	if v4917 != 0 {
		goto L1
	} else {
		goto L996
	}
L988:
	;
	v4856 = *(*int32)(unsafe.Add(mBase, uint32(v4834)))
	v4858 = F_table_open(m, v4856, int32(0))
	mBase = m.M
	v4859 = m.ExcPending
	if v4859 != 0 {
		goto L1
	} else {
		goto L990
	}
L989:
	;
	goto L987
L990:
	;
	v4861 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	v4862 = *(*int32)(unsafe.Add(mBase, uint32(v4861)+4))
	v4864 = F_index_open(m, v4862, int32(0))
	mBase = m.M
	v4865 = m.ExcPending
	if v4865 != 0 {
		goto L1
	} else {
		goto L991
	}
L991:
	;
	v4867 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	v4868 = *(*int32)(unsafe.Add(mBase, uint32(v4867)+8))
	v4869 = int32(0)
	F_index_build(m, v4858, v4864, v4868, v4869, v4869)
	mBase = m.M
	v4872 = m.ExcPending
	if v4872 != 0 {
		goto L1
	} else {
		goto L992
	}
L992:
	;
	F_relation_close(m, v4864, int32(0))
	mBase = m.M
	v4875 = m.ExcPending
	if v4875 != 0 {
		goto L1
	} else {
		goto L993
	}
L993:
	;
	F_relation_close(m, v4858, int32(0))
	mBase = m.M
	v4878 = m.ExcPending
	if v4878 != 0 {
		goto L1
	} else {
		goto L994
	}
L994:
	;
	v4879 = int32(_a_F_BootstrapModeMain_118)
	v4881 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	v4882 = *(*int32)(unsafe.Add(mBase, uint32(v4881)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76])) = v4882
	if v4882 != 0 {
		v4834 = v4882
		goto L988
	} else {
		goto L995
	}
L995:
	;
	goto L989
L996:
	;
	v4919 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	if v4919 != 0 {
		goto L997
	} else {
		goto L998
	}
L997:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4921 = m.ExcPending
	if v4921 != 0 {
		goto L1
	} else {
		goto L1000
	}
L998:
	;
	goto L999
L999:
	;
	v4922 = int32(0)
	v4923 = F_isatty(m, v4922)
	mBase = m.M
	if v4923 == v4922 {
		v5963 = v3835
		goto L695
	} else {
		goto L1001
	}
L1000:
	;
	goto L999
L1001:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_95), int32(0))
	mBase = m.M
	v4929 = m.ExcPending
	if v4929 != 0 {
		goto L1
	} else {
		goto L1002
	}
L1002:
	;
	v4930 = F_fflush(m, v3817)
	mBase = m.M
	v4931 = m.ExcPending
	if v4931 != 0 {
		goto L1
	} else {
		goto L1003
	}
L1003:
	;
	v5963 = v3835
	goto L695
L1004:
	;
	v5963 = v4936
	goto L695
L1005:
	;
	v5963 = v4944
	goto L695
L1006:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4947))) = int32(92)
	v4953 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v4947)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4947)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4947)+4)) = v4953
	v4959 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v4960 = F_makeString(m, v4959)
	mBase = m.M
	v4961 = m.ExcPending
	if v4961 != 0 {
		goto L1
	} else {
		goto L1007
	}
L1007:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+100)) = v4960
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+104)) = v4960
	v4967 = F_list_make1_impl(m, int32(1), v3811+int32(100))
	mBase = m.M
	v4968 = m.ExcPending
	if v4968 != 0 {
		goto L1
	} else {
		goto L1008
	}
L1008:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4947)+28)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4947)+20)) = v4967
	v5963 = v4947
	goto L695
L1009:
	;
	v4985 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(12))))
	v4988 = *(*int32)(unsafe.Add(mBase, uint32(v3808-int32(4))))
	v4989 = *(*int32)(unsafe.Add(mBase, uint32(v3808)))
	v4990 = int32(0)
	v4991 = m.G0
	v4993 = v4991 - int32(48)
	m.G0 = v4993
	v4996 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	if v4996 != 0 {
		goto L1010
	} else {
		goto L1011
	}
L1010:
	;
	v4999 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v5000 = m.ExcPending
	if v5000 != 0 {
		goto L1
	} else {
		goto L1013
	}
L1011:
	;
	goto L1012
L1012:
	;
	v5014 = v4977 << (uint(int32(2)) % 32)
	v5017 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	if v5017 == int32(0) {
		goto L1020
	} else {
		goto L1021
	}
L1013:
	;
	if v4999 != 0 {
		goto L1014
	} else {
		goto L1015
	}
L1014:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_119), int32(0))
	mBase = m.M
	v5004 = m.ExcPending
	if v5004 != 0 {
		goto L1
	} else {
		goto L1017
	}
L1015:
	;
	goto L1016
L1016:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v5012 = m.ExcPending
	if v5012 != 0 {
		goto L1
	} else {
		goto L1019
	}
L1017:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(528), int32(_a_F_BootstrapModeMain_120))
	mBase = m.M
	v5009 = m.ExcPending
	if v5009 != 0 {
		goto L1
	} else {
		goto L1018
	}
L1018:
	;
	goto L1016
L1019:
	;
	goto L1012
L1020:
	;
	v5021 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[74]))
	v5023 = F_MemoryContextAllocZero(m, v5021, int32(100))
	mBase = m.M
	v5024 = m.ExcPending
	if v5024 != 0 {
		goto L1
	} else {
		goto L1023
	}
L1021:
	;
	v5026 = v5017
	goto L1022
L1022:
	;
	v5027 = int32(0)
	base.MemoryFill(m, v5026, v5027, int32(100))
	v5030 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5034 = F_strncpy(m, v5030+int32(4), v4985, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v5034)+63)) = uint8(v5027)
	goto L1024
L1023:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73]))) = v5023
	v5026 = v5023
	goto L1022
L1024:
	;
	v5039 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5040 = m.ExcPending
	if v5040 != 0 {
		goto L1
	} else {
		goto L1025
	}
L1025:
	;
	if v5039 != 0 {
		goto L1026
	} else {
		goto L1027
	}
L1026:
	;
	v5041 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	*(*int32)(unsafe.Add(mBase, uint32(v4993)+36)) = v4988
	*(*int32)(unsafe.Add(mBase, uint32(v4993)+32)) = v5041 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_121), v4993+int32(32))
	mBase = m.M
	v5050 = m.ExcPending
	if v5050 != 0 {
		goto L1
	} else {
		goto L1029
	}
L1027:
	;
	goto L1028
L1028:
	;
	v5057 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5059 = v4977 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5057)+74)) = uint16(v5059)
	v5062 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[71]))
	if v5062 == int32(0) {
		goto L1035
	} else {
		goto L1036
	}
L1029:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(537), int32(_a_F_BootstrapModeMain_120))
	mBase = m.M
	v5055 = m.ExcPending
	if v5055 != 0 {
		goto L1
	} else {
		goto L1030
	}
L1030:
	;
	goto L1028
L1031:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5564)+80)) = uint16(v5567)
	v5590 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5591 = *(*int32)(unsafe.Add(mBase, uint32(v5590)+96))
	if v5591 != 0 {
		goto L1113
	} else {
		goto L1114
	}
L1032:
	;
	v5564 = v5537
	v5567 = int32(1)
	goto L1031
L1033:
	;
	v5504 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v5479
	v5507 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5508 = *(*int32)(unsafe.Add(mBase, uint32(v5479)))
	*(*int32)(unsafe.Add(mBase, uint32(v5507)+68)) = v5508
	v5510 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5479)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v5510)+72)) = uint16(v5511)
	v5513 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5479)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5513)+82)) = uint8(v5514)
	v5516 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5479)+132)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5516)+83)) = uint8(v5517)
	v5519 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5479)+133)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5519)+84)) = uint8(v5520)
	v5522 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5522)+85)) = uint8(v5504)
	v5525 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5526 = *(*int32)(unsafe.Add(mBase, uint32(v5479)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v5525)+96)) = v5526
	v5528 = *(*int32)(unsafe.Add(mBase, uint32(v5479)+96))
	if v5528 == v5504 {
		goto L1110
	} else {
		goto L1111
	}
L1034:
	;
	v5446 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5448 = v5074 * int32(92)
	v5449 = *(*int32)(unsafe.Add(mBase, uint32(v5448)+uint32(_c_F_BootstrapModeMain[78])))
	*(*int32)(unsafe.Add(mBase, uint32(v5446)+68)) = v5449
	v5451 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5452 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5448)+uint32(_c_F_BootstrapModeMain[79]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5451)+72)) = uint16(v5452)
	v5454 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5448)+uint32(_c_F_BootstrapModeMain[80]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5454)+82)) = uint8(v5455)
	v5457 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5448)+uint32(_c_F_BootstrapModeMain[81]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5457)+83)) = uint8(v5458)
	v5460 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5448)+uint32(_c_F_BootstrapModeMain[82]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5460)+84)) = uint8(v5461)
	v5463 = int32(0)
	v5464 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5464)+85)) = uint8(v5463)
	v5467 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5468 = *(*int32)(unsafe.Add(mBase, uint32(v5448)+uint32(_c_F_BootstrapModeMain[83])))
	*(*int32)(unsafe.Add(mBase, uint32(v5467)+96)) = v5468
	v5470 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	if int32(1)<<(uint(v5074)%32)&int32(_a_F_BootstrapModeMain_122) != 0 {
		v5564 = v5470
		v5567 = v5463
		goto L1031
	} else {
		goto L1108
	}
L1035:
	;
	v5074 = v4990
	goto L1038
L1036:
	;
	v5177 = v5062
	v5178 = v4990
	goto L1037
L1037:
	;
	v5195 = *(*int32)(unsafe.Add(mBase, uint32(v5177)+4))
	if int32(0) < v5195 {
		goto L1063
	} else {
		goto L1064
	}
L1038:
	;
	v5094 = v5074*int32(92) + int32(_a_F_BootstrapModeMain_123)
	goto L1042
L1039:
	;
	v5177 = v5166
	v5178 = v5164
	goto L1037
L1040:
	;
	if v5132-v5133 == int32(0) {
		goto L1034
	} else {
		goto L1053
	}
L1042:
	;
	goto L1043
L1043:
	;
	v5101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4988))))
	if v5101 != 0 {
		goto L1044
	} else {
		goto L1045
	}
L1044:
	;
	v5102 = v4988
	v5103 = v5094
	v5104 = int32(64)
	v5105 = v5101
	goto L1048
L1045:
	;
	v5128 = v5094
	v5132 = int32(0)
	goto L1046
L1046:
	;
	v5133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5128))))
	goto L1040
L1047:
	;
	v5128 = v5123
	v5132 = v5125
	goto L1046
L1048:
	;
	v5107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5103))))
	if base.B2i32(v5105 != v5107)|base.B2i32(v5107 == int32(0)) != 0 {
		v5123 = v5103
		v5125 = v5105
		goto L1047
	} else {
		goto L1050
	}
L1049:
	;
	v5123 = v5117
	v5125 = int32(0)
	goto L1047
L1050:
	;
	v5113 = v5104 - int32(1)
	if v5113 == int32(0) {
		v5123 = v5103
		v5125 = v5105
		goto L1047
	} else {
		goto L1051
	}
L1051:
	;
	v5116 = int32(1)
	v5117 = v5103 + v5116
	v5118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5102)+1)))
	if v5118 != 0 {
		v5102 = v5102 + v5116
		v5103 = v5117
		v5104 = v5113
		v5105 = v5118
		goto L1048
	} else {
		goto L1052
	}
L1052:
	;
	goto L1049
L1053:
	;
	v5144 = v5074 + int32(1)
	if v5144 != int32(25) {
		v5074 = v5144
		goto L1038
	} else {
		goto L1054
	}
L1054:
	;
	v5149 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5150 = m.ExcPending
	if v5150 != 0 {
		goto L1
	} else {
		goto L1055
	}
L1055:
	;
	if v5149 != 0 {
		goto L1056
	} else {
		goto L1057
	}
L1056:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4993)+16)) = v4988
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_124), v4993+int32(16))
	mBase = m.M
	v5156 = m.ExcPending
	if v5156 != 0 {
		goto L1
	} else {
		goto L1059
	}
L1057:
	;
	goto L1058
L1058:
	;
	F_populate_typ_list(m)
	mBase = m.M
	v5163 = m.ExcPending
	if v5163 != 0 {
		goto L1
	} else {
		goto L1061
	}
L1059:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(817), int32(_a_F_BootstrapModeMain_125))
	mBase = m.M
	v5161 = m.ExcPending
	if v5161 != 0 {
		goto L1
	} else {
		goto L1060
	}
L1060:
	;
	goto L1058
L1061:
	;
	v5164 = int32(0)
	v5166 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[71]))
	if v5166 == v5164 {
		v5074 = v5164
		goto L1038
	} else {
		goto L1062
	}
L1062:
	;
	goto L1039
L1063:
	;
	v5198 = *(*int32)(unsafe.Add(mBase, uint32(v5177)+12))
	v5208 = v5178
	goto L1066
L1064:
	;
	goto L1065
L1065:
	;
	F_list_free_deep(m, v5177)
	mBase = m.M
	v5309 = m.ExcPending
	if v5309 != 0 {
		goto L1
	} else {
		goto L1083
	}
L1066:
	;
	v5228 = *(*int32)(unsafe.Add(mBase, uint32(v5198+v5208<<(uint(int32(2))%32))))
	v5230 = v5228 + int32(8)
	goto L1070
L1067:
	;
	goto L1065
L1068:
	;
	if v5268-v5269 == int32(0) {
		v5479 = v5228
		goto L1033
	} else {
		goto L1081
	}
L1070:
	;
	goto L1071
L1071:
	;
	v5237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5230))))
	if v5237 != 0 {
		goto L1072
	} else {
		goto L1073
	}
L1072:
	;
	v5238 = v5230
	v5239 = v4988
	v5240 = int32(64)
	v5241 = v5237
	goto L1076
L1073:
	;
	v5264 = v4988
	v5268 = int32(0)
	goto L1074
L1074:
	;
	v5269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5264))))
	goto L1068
L1075:
	;
	v5264 = v5259
	v5268 = v5261
	goto L1074
L1076:
	;
	v5243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5239))))
	if base.B2i32(v5241 != v5243)|base.B2i32(v5243 == int32(0)) != 0 {
		v5259 = v5239
		v5261 = v5241
		goto L1075
	} else {
		goto L1078
	}
L1077:
	;
	v5259 = v5253
	v5261 = int32(0)
	goto L1075
L1078:
	;
	v5249 = v5240 - int32(1)
	if v5249 == int32(0) {
		v5259 = v5239
		v5261 = v5241
		goto L1075
	} else {
		goto L1079
	}
L1079:
	;
	v5252 = int32(1)
	v5253 = v5239 + v5252
	v5254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5238)+1)))
	if v5254 != 0 {
		v5238 = v5238 + v5252
		v5239 = v5253
		v5240 = v5249
		v5241 = v5254
		goto L1076
	} else {
		goto L1080
	}
L1080:
	;
	goto L1077
L1081:
	;
	v5280 = v5208 + int32(1)
	if v5280 != v5195 {
		v5208 = v5280
		goto L1066
	} else {
		goto L1082
	}
L1082:
	;
	goto L1067
L1083:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[71])) = int32(0)
	F_populate_typ_list(m)
	mBase = m.M
	v5314 = m.ExcPending
	if v5314 != 0 {
		goto L1
	} else {
		goto L1084
	}
L1084:
	;
	v5316 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[71]))
	if v5316 == int32(0) {
		goto L1085
	} else {
		goto L1086
	}
L1085:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5436 = m.ExcPending
	if v5436 != 0 {
		goto L1
	} else {
		goto L1105
	}
L1086:
	;
	v5319 = *(*int32)(unsafe.Add(mBase, uint32(v5316)+4))
	if v5319 <= int32(0) {
		goto L1085
	} else {
		goto L1087
	}
L1087:
	;
	v5322 = *(*int32)(unsafe.Add(mBase, uint32(v5316)+12))
	v5333 = int32(0)
	goto L1088
L1088:
	;
	v5353 = *(*int32)(unsafe.Add(mBase, uint32(v5322+v5333<<(uint(int32(2))%32))))
	v5355 = v5353 + int32(8)
	goto L1092
L1089:
	;
	goto L1085
L1090:
	;
	if v5393-v5394 == int32(0) {
		v5479 = v5353
		goto L1033
	} else {
		goto L1103
	}
L1092:
	;
	goto L1093
L1093:
	;
	v5362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5355))))
	if v5362 != 0 {
		goto L1094
	} else {
		goto L1095
	}
L1094:
	;
	v5363 = v5355
	v5364 = v4988
	v5365 = int32(64)
	v5366 = v5362
	goto L1098
L1095:
	;
	v5389 = v4988
	v5393 = int32(0)
	goto L1096
L1096:
	;
	v5394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5389))))
	goto L1090
L1097:
	;
	v5389 = v5384
	v5393 = v5386
	goto L1096
L1098:
	;
	v5368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5364))))
	if base.B2i32(v5366 != v5368)|base.B2i32(v5368 == int32(0)) != 0 {
		v5384 = v5364
		v5386 = v5366
		goto L1097
	} else {
		goto L1100
	}
L1099:
	;
	v5384 = v5378
	v5386 = int32(0)
	goto L1097
L1100:
	;
	v5374 = v5365 - int32(1)
	if v5374 == int32(0) {
		v5384 = v5364
		v5386 = v5366
		goto L1097
	} else {
		goto L1101
	}
L1101:
	;
	v5377 = int32(1)
	v5378 = v5364 + v5377
	v5379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5363)+1)))
	if v5379 != 0 {
		v5363 = v5363 + v5377
		v5364 = v5378
		v5365 = v5374
		v5366 = v5379
		goto L1098
	} else {
		goto L1102
	}
L1102:
	;
	goto L1099
L1103:
	;
	v5405 = v5333 + int32(1)
	if v5319 != v5405 {
		v5333 = v5405
		goto L1088
	} else {
		goto L1104
	}
L1104:
	;
	goto L1089
L1105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4993))) = v4988
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_126), v4993)
	mBase = m.M
	v5440 = m.ExcPending
	if v5440 != 0 {
		goto L1
	} else {
		goto L1106
	}
L1106:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(821), int32(_a_F_BootstrapModeMain_125))
	mBase = m.M
	v5445 = m.ExcPending
	if v5445 != 0 {
		goto L1
	} else {
		goto L1107
	}
L1107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1108:
	;
	v5475 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5470)+72)))
	if int32(0) <= v5475 {
		v5564 = v5470
		v5567 = v5463
		goto L1031
	} else {
		goto L1109
	}
L1109:
	;
	v5537 = v5470
	goto L1032
L1110:
	;
	v5535 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5564 = v5535
	v5567 = v5504
	goto L1031
L1111:
	;
	v5531 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5479)+80)))
	if int32(0) <= v5531 {
		goto L1110
	} else {
		goto L1112
	}
L1112:
	;
	v5534 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5537 = v5534
	goto L1032
L1113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5590)+96)) = int32(950)
	v5594 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	v5595 = v5594
	goto L1115
L1114:
	;
	v5595 = v5590
	goto L1115
L1115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5595)+76)) = int32(-1)
	v5598 = int32(1)
	v5599 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5599)+92)) = uint8(v5598)
	v5602 = *(*int32)(unsafe.Add(mBase, uint32(v5014)+uint32(_c_F_BootstrapModeMain[73])))
	switch v4989 - int32(2) {
	case 0:
		goto L1119
	case 1:
		v5681 = v5598
		goto L1117
	default:
		goto L1118
	}
L1116:
	;
	m.G0 = v4993 + int32(48)
	v5963 = v3835
	goto L695
L1117:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v5602)+86)) = uint8(v5681)
	goto L1116
L1118:
	;
	v5606 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5602)+72)))
	if v5606 <= int32(0) {
		goto L1116
	} else {
		goto L1120
	}
L1119:
	;
	v5681 = int32(0)
	goto L1117
L1120:
	;
	v5609 = int32(0)
	if v4977 <= v5609 {
		v5659 = v5609
		goto L1121
	} else {
		goto L1122
	}
L1121:
	;
	if v5659 != v4977 {
		goto L1116
	} else {
		goto L1128
	}
L1122:
	;
	v5621 = v5609
	goto L1123
L1123:
	;
	v5640 = *(*int32)(unsafe.Add(mBase, uint32(v5621<<(uint(int32(2))%32))+uint32(_c_F_BootstrapModeMain[73])))
	v5641 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5640)+72)))
	if v5641 <= int32(0) {
		v5659 = v5621
		goto L1121
	} else {
		goto L1125
	}
L1124:
	;
	v5681 = v5598
	goto L1117
L1125:
	;
	v5644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5640)+86)))
	if v5644 != int32(1) {
		v5659 = v5621
		goto L1121
	} else {
		goto L1126
	}
L1126:
	;
	v5648 = v5621 + int32(1)
	if v5648 != v4977 {
		v5621 = v5648
		goto L1123
	} else {
		goto L1127
	}
L1127:
	;
	goto L1124
L1128:
	;
	v5681 = v5598
	goto L1117
L1129:
	;
	v5963 = base.I32_wrap_i64(v5739)
	goto L695
L1130:
	;
	if v5754 != 0 {
		goto L1131
	} else {
		goto L1132
	}
L1131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5750)+20)) = v5741
	*(*int32)(unsafe.Add(mBase, uint32(v5750)+16)) = v5744
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_127), v5750+int32(16))
	mBase = m.M
	v5762 = m.ExcPending
	if v5762 != 0 {
		goto L1
	} else {
		goto L1134
	}
L1132:
	;
	goto L1133
L1133:
	;
	v5769 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	v5770 = *(*int32)(unsafe.Add(mBase, uint32(v5769)+52))
	v5771 = *(*int32)(unsafe.Add(mBase, uint32(v5770)))
	v5778 = *(*int32)(unsafe.Add(mBase, uint32(v5770+v5771<<(uint(int32(4))%32)+v5744*int32(100))+88))
	F_boot_get_type_io_data(m, v5778, v5750+int32(46), v5750+int32(45), v5750+int32(44), v5750+int32(43), v5750+int32(36), v5750+int32(32), v5750+int32(28))
	mBase = m.M
	v5794 = m.ExcPending
	if v5794 != 0 {
		goto L1
	} else {
		goto L1136
	}
L1134:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(670), int32(_a_F_BootstrapModeMain_128))
	mBase = m.M
	v5767 = m.ExcPending
	if v5767 != 0 {
		goto L1
	} else {
		goto L1135
	}
L1135:
	;
	goto L1133
L1136:
	;
	v5796 = v5744 << (uint(int32(2)) % 32)
	v5799 = *(*int32)(unsafe.Add(mBase, uint32(v5750)+32))
	v5800 = *(*int32)(unsafe.Add(mBase, uint32(v5750)+36))
	v5802 = F_OidInputFunctionCall(m, v5799, v5741, v5800, int32(-1))
	mBase = m.M
	v5803 = m.ExcPending
	if v5803 != 0 {
		goto L1
	} else {
		goto L1137
	}
L1137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5796)+uint32(_c_F_BootstrapModeMain[84]))) = v5802
	v5807 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5808 = m.ExcPending
	if v5808 != 0 {
		goto L1
	} else {
		goto L1138
	}
L1138:
	;
	if v5807 != 0 {
		goto L1139
	} else {
		goto L1140
	}
L1139:
	;
	v5809 = *(*int32)(unsafe.Add(mBase, uint32(v5750)+28))
	v5810 = *(*int32)(unsafe.Add(mBase, uint32(v5796)+uint32(_c_F_BootstrapModeMain[84])))
	v5811 = F_OidOutputFunctionCall(m, v5809, v5810)
	mBase = m.M
	v5812 = m.ExcPending
	if v5812 != 0 {
		goto L1
	} else {
		goto L1142
	}
L1140:
	;
	goto L1141
L1141:
	;
	m.G0 = v5750 + int32(48)
	v5963 = v3835
	goto L695
L1142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5750))) = v5811
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_129), v5750)
	mBase = m.M
	v5816 = m.ExcPending
	if v5816 != 0 {
		goto L1
	} else {
		goto L1143
	}
L1143:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(687), int32(_a_F_BootstrapModeMain_128))
	mBase = m.M
	v5821 = m.ExcPending
	if v5821 != 0 {
		goto L1
	} else {
		goto L1144
	}
L1144:
	;
	goto L1141
L1145:
	;
	if v5837 != 0 {
		goto L1146
	} else {
		goto L1147
	}
L1146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5833)+16)) = v5827
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_130), v5833+int32(16))
	mBase = m.M
	v5844 = m.ExcPending
	if v5844 != 0 {
		goto L1
	} else {
		goto L1149
	}
L1147:
	;
	goto L1148
L1148:
	;
	v5851 = v5827 * int32(100)
	v5853 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	v5854 = *(*int32)(unsafe.Add(mBase, uint32(v5853)+52))
	v5855 = *(*int32)(unsafe.Add(mBase, uint32(v5854)))
	v5860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5851+(v5854+v5855<<(uint(int32(4))%32)))+106)))
	if v5860 == int32(1) {
		goto L1151
	} else {
		goto L1152
	}
L1149:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(697), int32(_a_F_BootstrapModeMain_131))
	mBase = m.M
	v5849 = m.ExcPending
	if v5849 != 0 {
		goto L1
	} else {
		goto L1150
	}
L1150:
	;
	goto L1148
L1151:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5866 = m.ExcPending
	if v5866 != 0 {
		goto L1
	} else {
		goto L1154
	}
L1152:
	;
	goto L1153
L1153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5827<<(uint(int32(2))%32))+uint32(_c_F_BootstrapModeMain[84]))) = int32(0)
	v5898 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5827)+uint32(_c_F_BootstrapModeMain[48]))) = uint8(v5898)
	m.G0 = v5833 + int32(32)
	v5963 = v3835
	goto L695
L1154:
	;
	v5868 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	v5869 = *(*int32)(unsafe.Add(mBase, uint32(v5868)+52))
	v5870 = *(*int32)(unsafe.Add(mBase, uint32(v5869)))
	v5871 = *(*int32)(unsafe.Add(mBase, uint32(v5868)+48))
	v5872 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v5833)+4)) = v5871 + v5872
	*(*int32)(unsafe.Add(mBase, uint32(v5833))) = v5869 + v5870<<(uint(v5872)%32) + v5851 + int32(24)
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_132), v5833)
	mBase = m.M
	v5884 = m.ExcPending
	if v5884 != 0 {
		goto L1
	} else {
		goto L1155
	}
L1155:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(703), int32(_a_F_BootstrapModeMain_131))
	mBase = m.M
	v5889 = m.ExcPending
	if v5889 != 0 {
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
	v5963 = v5905
	goto L695
L1158:
	;
	v5963 = v5908
	goto L695
L1159:
	;
	v5963 = v5911
	goto L695
L1160:
	;
	v5963 = v5914
	goto L695
L1161:
	;
	v5963 = v5917
	goto L695
L1162:
	;
	v5963 = v5920
	goto L695
L1163:
	;
	v5963 = v5923
	goto L695
L1164:
	;
	v5963 = v5926
	goto L695
L1165:
	;
	v5963 = v5929
	goto L695
L1166:
	;
	v5963 = v5932
	goto L695
L1167:
	;
	v5963 = v5935
	goto L695
L1168:
	;
	v5963 = v5938
	goto L695
L1169:
	;
	v5963 = v5941
	goto L695
L1170:
	;
	v5963 = v5944
	goto L695
L1171:
	;
	v5963 = v5947
	goto L695
L1172:
	;
	v5963 = v5950
	goto L695
L1173:
	;
	v5963 = v5953
	goto L695
L1174:
	;
	v5963 = v5956
	goto L695
L1175:
	;
	v5963 = v5959
	goto L695
L1176:
	;
	v6013 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5997)+uint32(_c_F_BootstrapModeMain[85]))))
	v6017 = v3804
	v6019 = v3806
	v6021 = v5992
	v6024 = v3811
	v6025 = v3812
	v6027 = v6013
	v6028 = v5993
	v6030 = v3817
	v6031 = v3818
	v6032 = v3819
	v6036 = v3823
	goto L314
L1177:
	;
	v6004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6001)+uint32(_c_F_BootstrapModeMain[62]))))
	if v6004 != v5994&int32(255) {
		goto L1176
	} else {
		goto L1178
	}
L1178:
	;
	v6010 = int32(*(*int8)(unsafe.Add(mBase, uint32(v6001)+uint32(_c_F_BootstrapModeMain[63]))))
	v6017 = v3804
	v6019 = v3806
	v6021 = v5992
	v6024 = v3811
	v6025 = v3812
	v6027 = v6010
	v6028 = v5993
	v6030 = v3817
	v6031 = v3818
	v6032 = v3819
	v6036 = v3823
	goto L314
L1179:
	;
	if v3775 == int32(0) {
		v6163 = v3780
		v6164 = v3781
		goto L282
	} else {
		goto L1182
	}
L1180:
	;
	F_boot_yyerror(m, v3773, int32(_a_F_BootstrapModeMain_133))
	mBase = m.M
	v6044 = m.ExcPending
	if v6044 != 0 {
		goto L1
	} else {
		goto L1181
	}
L1181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1182:
	;
	v6057 = v3780
	v6058 = v3781
	v6061 = v3784
	goto L287
L1183:
	;
	if v6058 == v6087 {
		v6163 = v6057
		v6164 = v6058
		goto L282
	} else {
		goto L1185
	}
L1185:
	;
	v6087 = v6087 - int32(1)
	goto L1183
L1186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1187:
	;
	v6112 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[67]))
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+32)) = v6112
	v6115 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[66]))
	*(*int32)(unsafe.Add(mBase, uint32(v3811)+36)) = v6115
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_134), v3811+int32(32))
	mBase = m.M
	v6121 = m.ExcPending
	if v6121 != 0 {
		goto L1
	} else {
		goto L1188
	}
L1188:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(265), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v6126 = m.ExcPending
	if v6126 != 0 {
		goto L1
	} else {
		goto L1189
	}
L1189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1190:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_135), int32(0))
	mBase = m.M
	v6134 = m.ExcPending
	if v6134 != 0 {
		goto L1
	} else {
		goto L1191
	}
L1191:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(267), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v6139 = m.ExcPending
	if v6139 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1193:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_136), int32(0))
	mBase = m.M
	v6147 = m.ExcPending
	if v6147 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1194:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_100), int32(446), int32(_a_F_BootstrapModeMain_101))
	mBase = m.M
	v6152 = m.ExcPending
	if v6152 != 0 {
		goto L1
	} else {
		goto L1195
	}
L1195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1196:
	;
	F_pfree(m, v6179)
	mBase = m.M
	v6209 = m.ExcPending
	if v6209 != 0 {
		goto L1
	} else {
		goto L1199
	}
L1197:
	;
	goto L1198
L1198:
	;
	m.G0 = v6189 + int32(1136)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6214 = m.ExcPending
	if v6214 != 0 {
		goto L1
	} else {
		goto L1200
	}
L1199:
	;
	goto L1198
L1200:
	;
	v6216 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	v6220 = F_LWLockAcquire(m, v6216+int32(3200), int32(0))
	mBase = m.M
	v6221 = m.ExcPending
	if v6221 != 0 {
		goto L1
	} else {
		goto L1201
	}
L1201:
	;
	v6223 = int32(0)
	F_write_relmap_file(m, int32(_a_F_BootstrapModeMain_137), v6223, v6223, v6223, v6223, int32(1664), int32(_a_F_BootstrapModeMain_138))
	mBase = m.M
	v6230 = m.ExcPending
	if v6230 != 0 {
		goto L1
	} else {
		goto L1202
	}
L1202:
	;
	v6232 = int32(0)
	v6236 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[86]))
	v6238 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[87]))
	v6240 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[88]))
	F_write_relmap_file(m, int32(_a_F_BootstrapModeMain_139), v6232, v6232, v6232, v6236, v6238, v6240)
	mBase = m.M
	v6242 = m.ExcPending
	if v6242 != 0 {
		goto L1
	} else {
		goto L1203
	}
L1203:
	;
	v6244 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[18]))
	F_LWLockRelease(m, v6244+int32(3200))
	mBase = m.M
	v6248 = m.ExcPending
	if v6248 != 0 {
		goto L1
	} else {
		goto L1204
	}
L1204:
	;
	v6250 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[72]))
	if v6250 != 0 {
		goto L1205
	} else {
		goto L1206
	}
L1205:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v6253 = m.ExcPending
	if v6253 != 0 {
		goto L1
	} else {
		goto L1208
	}
L1206:
	;
	goto L1207
L1207:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6256 = m.ExcPending
	if v6256 != 0 {
		goto L1
	} else {
		goto L1209
	}
L1208:
	;
	goto L1207
L1209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1210:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6263 = m.ExcPending
	if v6263 != 0 {
		goto L1
	} else {
		goto L1211
	}
L1211:
	;
	v6265 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v6265
	F_errmsg(m, int32(_a_F_BootstrapModeMain_140), v29-int32(-64))
	mBase = m.M
	v6271 = m.ExcPending
	if v6271 != 0 {
		goto L1
	} else {
		goto L1212
	}
L1212:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(239), int32(_a_F_BootstrapModeMain_16))
	mBase = m.M
	v6276 = m.ExcPending
	if v6276 != 0 {
		goto L1
	} else {
		goto L1213
	}
L1213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1214:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(254), int32(_a_F_BootstrapModeMain_16))
	mBase = m.M
	v6287 = m.ExcPending
	if v6287 != 0 {
		goto L1
	} else {
		goto L1215
	}
L1215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1216:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6294 = m.ExcPending
	if v6294 != 0 {
		goto L1
	} else {
		goto L1217
	}
L1217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1218:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6303 = m.ExcPending
	if v6303 != 0 {
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
	base.Wasm_trap_unreachable()
	for {
	}
L1221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1222:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_141), int32(0))
	mBase = m.M
	v6320 = m.ExcPending
	if v6320 != 0 {
		goto L1
	} else {
		goto L1223
	}
L1223:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_15), int32(383), int32(_a_F_BootstrapModeMain_16))
	mBase = m.M
	v6325 = m.ExcPending
	if v6325 != 0 {
		goto L1
	} else {
		goto L1224
	}
L1224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
