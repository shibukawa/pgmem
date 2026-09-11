package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleSatisfiesVisibility(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
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
	var v444 int32
	_ = v444
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v705 int32
	_ = v705
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v896 int32
	_ = v896
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1041 int32
	_ = v1041
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1206 int32
	_ = v1206
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1344 int32
	_ = v1344
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
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1484 int32
	_ = v1484
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1530 int32
	_ = v1530
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1582 int32
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1618 int32
	_ = v1618
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1696 int32
	_ = v1696
	var v1706 int32
	_ = v1706
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1742 int32
	_ = v1742
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1826 int32
	_ = v1826
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1911 int32
	_ = v1911
	var v1919 int32
	_ = v1919
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1954 int32
	_ = v1954
	var v1959 int32
	_ = v1959
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2011 int32
	_ = v2011
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2047 int32
	_ = v2047
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2080 int32
	_ = v2080
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2108 int32
	_ = v2108
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	var v2132 int32
	_ = v2132
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2166 int32
	_ = v2166
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2174 int32
	_ = v2174
	var v2184 int32
	_ = v2184
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2212 int32
	_ = v2212
	var v2220 int32
	_ = v2220
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2238 int32
	_ = v2238
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2257 int32
	_ = v2257
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2270 int32
	_ = v2270
	var v2275 int32
	_ = v2275
	var v2278 int32
	_ = v2278
	var v2281 int32
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2296 int32
	_ = v2296
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2304 int32
	_ = v2304
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2322 int32
	_ = v2322
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2358 int32
	_ = v2358
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2374 int32
	_ = v2374
	var v2378 int32
	_ = v2378
	var v2381 int32
	_ = v2381
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2389 int32
	_ = v2389
	var v2393 int32
	_ = v2393
	var v2398 int32
	_ = v2398
	var v2401 int32
	_ = v2401
	var v2406 int32
	_ = v2406
	var v2417 int32
	_ = v2417
	var v2421 int32
	_ = v2421
	var v2425 int32
	_ = v2425
	var v2430 int32
	_ = v2430
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2443 int32
	_ = v2443
	var v2445 int32
	_ = v2445
	var v2450 int32
	_ = v2450
	var v2452 int32
	_ = v2452
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2464 int32
	_ = v2464
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2472 int32
	_ = v2472
	var v2482 int32
	_ = v2482
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2491 int32
	_ = v2491
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2510 int32
	_ = v2510
	var v2518 int32
	_ = v2518
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2535 int32
	_ = v2535
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2553 int32
	_ = v2553
	var v2557 int32
	_ = v2557
	var v2561 int32
	_ = v2561
	var v2566 int32
	_ = v2566
	var v2571 int32
	_ = v2571
	var v2574 int32
	_ = v2574
	var v2577 int32
	_ = v2577
	var v2579 int32
	_ = v2579
	var v2581 int32
	_ = v2581
	var v2586 int32
	_ = v2586
	var v2588 int32
	_ = v2588
	var v2592 int32
	_ = v2592
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2600 int32
	_ = v2600
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2608 int32
	_ = v2608
	var v2618 int32
	_ = v2618
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2627 int32
	_ = v2627
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2646 int32
	_ = v2646
	var v2654 int32
	_ = v2654
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2691 int32
	_ = v2691
	var v2695 int32
	_ = v2695
	var v2699 int32
	_ = v2699
	var v2704 int32
	_ = v2704
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2715 int32
	_ = v2715
	var v2717 int32
	_ = v2717
	var v2719 int32
	_ = v2719
	var v2724 int32
	_ = v2724
	var v2726 int32
	_ = v2726
	var v2730 int32
	_ = v2730
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2738 int32
	_ = v2738
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2746 int32
	_ = v2746
	var v2756 int32
	_ = v2756
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2765 int32
	_ = v2765
	var v2770 int32
	_ = v2770
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2780 int32
	_ = v2780
	var v2781 int32
	_ = v2781
	var v2784 int32
	_ = v2784
	var v2792 int32
	_ = v2792
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2825 int32
	_ = v2825
	var v2829 int32
	_ = v2829
	var v2833 int32
	_ = v2833
	var v2838 int32
	_ = v2838
	var v2843 int32
	_ = v2843
	var v2846 int32
	_ = v2846
	var v2849 int32
	_ = v2849
	var v2851 int32
	_ = v2851
	var v2853 int32
	_ = v2853
	var v2858 int32
	_ = v2858
	var v2860 int32
	_ = v2860
	var v2864 int32
	_ = v2864
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2872 int32
	_ = v2872
	var v2876 int32
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2880 int32
	_ = v2880
	var v2890 int32
	_ = v2890
	var v2895 int32
	_ = v2895
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2918 int32
	_ = v2918
	var v2926 int32
	_ = v2926
	var v2934 int32
	_ = v2934
	var v2937 int32
	_ = v2937
	var v2938 int32
	_ = v2938
	var v2949 int32
	_ = v2949
	var v2953 int32
	_ = v2953
	var v2957 int32
	_ = v2957
	var v2962 int32
	_ = v2962
	var v2967 int32
	_ = v2967
	var v2970 int32
	_ = v2970
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v2977 int32
	_ = v2977
	var v2982 int32
	_ = v2982
	var v2984 int32
	_ = v2984
	var v2988 int32
	_ = v2988
	var v2990 int32
	_ = v2990
	var v2991 int32
	_ = v2991
	var v2992 int32
	_ = v2992
	var v2996 int32
	_ = v2996
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3004 int32
	_ = v3004
	var v3014 int32
	_ = v3014
	var v3019 int32
	_ = v3019
	var v3021 int32
	_ = v3021
	var v3023 int32
	_ = v3023
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3042 int32
	_ = v3042
	var v3050 int32
	_ = v3050
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3066 int32
	_ = v3066
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3075 int32
	_ = v3075
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3082 int32
	_ = v3082
	var v3084 int32
	_ = v3084
	var v3085 int32
	_ = v3085
	var v3087 int32
	_ = v3087
	var v3091 int32
	_ = v3091
	var v3093 int32
	_ = v3093
	var v3097 int32
	_ = v3097
	var v3100 int32
	_ = v3100
	var v3101 int32
	_ = v3101
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3136 int32
	_ = v3136
	var v3140 int32
	_ = v3140
	var v3144 int32
	_ = v3144
	var v3149 int32
	_ = v3149
	var v3154 int32
	_ = v3154
	var v3157 int32
	_ = v3157
	var v3160 int32
	_ = v3160
	var v3162 int32
	_ = v3162
	var v3164 int32
	_ = v3164
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3175 int32
	_ = v3175
	var v3177 int32
	_ = v3177
	var v3178 int32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3183 int32
	_ = v3183
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3191 int32
	_ = v3191
	var v3201 int32
	_ = v3201
	var v3206 int32
	_ = v3206
	var v3208 int32
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3229 int32
	_ = v3229
	var v3237 int32
	_ = v3237
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3250 int32
	_ = v3250
	var v3251 int32
	_ = v3251
	var v3254 int32
	_ = v3254
	var v3265 int32
	_ = v3265
	var v3269 int32
	_ = v3269
	var v3273 int32
	_ = v3273
	var v3278 int32
	_ = v3278
	var v3283 int32
	_ = v3283
	var v3286 int32
	_ = v3286
	var v3289 int32
	_ = v3289
	var v3291 int32
	_ = v3291
	var v3293 int32
	_ = v3293
	var v3298 int32
	_ = v3298
	var v3300 int32
	_ = v3300
	var v3304 int32
	_ = v3304
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3312 int32
	_ = v3312
	var v3316 int32
	_ = v3316
	var v3317 int32
	_ = v3317
	var v3320 int32
	_ = v3320
	var v3330 int32
	_ = v3330
	var v3335 int32
	_ = v3335
	var v3337 int32
	_ = v3337
	var v3339 int32
	_ = v3339
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3346 int32
	_ = v3346
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3358 int32
	_ = v3358
	var v3366 int32
	_ = v3366
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3385 int32
	_ = v3385
	var v3386 int32
	_ = v3386
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3395 int32
	_ = v3395
	var v3397 int32
	_ = v3397
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3404 int32
	_ = v3404
	var v3408 int32
	_ = v3408
	var v3418 int32
	_ = v3418
	var v3421 int32
	_ = v3421
	var v3423 int32
	_ = v3423
	var v3426 int32
	_ = v3426
	var v3427 int32
	_ = v3427
	var v3428 int32
	_ = v3428
	var v3429 int32
	_ = v3429
	var v3435 int32
	_ = v3435
	var v3436 int32
	_ = v3436
	var v3437 int32
	_ = v3437
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3447 int32
	_ = v3447
	var v3448 int32
	_ = v3448
	var v3451 int32
	_ = v3451
	var v3456 int32
	_ = v3456
	var v3461 int32
	_ = v3461
	var v3462 int32
	_ = v3462
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3469 int32
	_ = v3469
	var v3481 int32
	_ = v3481
	var v3482 int32
	_ = v3482
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3488 int32
	_ = v3488
	var v3500 int32
	_ = v3500
	var v3502 int32
	_ = v3502
	var v3503 int32
	_ = v3503
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3531 int32
	_ = v3531
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3541 int32
	_ = v3541
	var v3542 int32
	_ = v3542
	var v3545 int32
	_ = v3545
	var v3548 int32
	_ = v3548
	var v3553 int32
	_ = v3553
	var v3554 int32
	_ = v3554
	var v3557 int32
	_ = v3557
	var v3560 int32
	_ = v3560
	var v3562 int32
	_ = v3562
	var v3574 int32
	_ = v3574
	var v3576 int32
	_ = v3576
	var v3579 int32
	_ = v3579
	var v3580 int32
	_ = v3580
	var v3583 int32
	_ = v3583
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3597 int32
	_ = v3597
	var v3605 int32
	_ = v3605
	var v3606 int32
	_ = v3606
	var v3614 int32
	_ = v3614
	var v3615 int32
	_ = v3615
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3624 int32
	_ = v3624
	var v3625 int32
	_ = v3625
	var v3633 int32
	_ = v3633
	var v3636 int32
	_ = v3636
	var v3640 int32
	_ = v3640
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v14 {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	case 4:
		goto L5
	case 5:
		goto L4
	case 6:
		goto L3
	default:
		v3640 = v4
		goto L1
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v3640
L2:
	;
	v3633 = int32(1)
	F_MarkBufferDirtyHint(m, l2, v3633)
	mBase = m.M
	v3636 = m.ExcPending
	if v3636 != 0 {
		goto L64
	} else {
		goto L1309
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v3614 = F_HeapTupleSatisfiesVacuumHorizon(m, l0, l2, v12+int32(12))
	mBase = m.M
	v3615 = m.ExcPending
	if v3615 != 0 {
		goto L64
	} else {
		goto L1301
	}
L4:
	;
	v3426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v3427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3426)+20)))
	v3428 = int32(768)
	v3429 = v3427 & v3428
	if v3429 != v3428 {
		goto L1234
	} else {
		goto L1235
	}
L5:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(0)
	v2398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	if v2398&int32(256) != 0 {
		goto L871
	} else {
		goto L872
	}
L6:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v2099 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2098)+20)))
	if v2099&int32(256) != 0 {
		goto L761
	} else {
		goto L762
	}
L7:
	;
	v3640 = int32(1)
	goto L1
L8:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1085 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)))
	if v1085&int32(256) != 0 {
		goto L401
	} else {
		goto L402
	}
L9:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)))
	v17 = base.I32_extend16_s(v16)
	if v16&int32(256) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v771 = int32(1)
	v772 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)))
	if v772&int32(2048) != 0 {
		v3640 = v771
		goto L1
	} else {
		goto L285
	}
L11:
	;
	v764 = v155 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)) = uint16(v764)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L64
	} else {
		goto L284
	}
L12:
	;
	if v17&int32(512) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v753 = int32(768)
	if v17&v753 == v753 {
		goto L10
	} else {
		goto L281
	}
L15:
	;
	v3640 = int32(0)
	goto L1
L16:
	;
	goto L17
L17:
	;
	if v17&int32(16384) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if base.Ui32(v27) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	if v17 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L21:
	;
	if v147 != 0 {
		goto L61
	} else {
		goto L62
	}
L22:
	;
	v147 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v38 == v27 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v147 = int32(1)
	goto L21
L26:
	;
	goto L27
L27:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v42 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v147 = v139
	goto L21
L29:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v46 == int32(0) {
		v139 = int32(0)
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v110 = int32(0)
	v112 = v42 - int32(1)
	goto L51
L32:
	;
	v51 = v46
	goto L33
L33:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	if v56 == int32(4) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v139 = int32(0)
	goto L28
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
	if v103 != 0 {
		v51 = v103
		goto L33
	} else {
		goto L50
	}
L36:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v59 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v62 = int32(1)
	if v27 == v59 {
		v139 = v62
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
	v66 = v64 - int32(1)
	if v66 < int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v71 = int32(0)
	v73 = v66
	goto L40
L40:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v51)+48))
	v79 = int32(2)
	v80 = base.I32_div_s(v73-v71, v79)
	v81 = v80 + v71
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v77+v81<<(uint(v79)%32))))
	if v85 == v27 {
		v139 = v62
		goto L28
	} else {
		goto L42
	}
L41:
	;
	goto L35
L42:
	;
	v89 = F_TransactionIdPrecedes(m, v85, v27)
	mBase = m.M
	if v89 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v90 = v81 + int32(1)
	goto L45
L44:
	;
	v90 = v71
	goto L45
L45:
	;
	if v89 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v93 = v73
	goto L48
L47:
	;
	v93 = v81 - int32(1)
	goto L48
L48:
	;
	if v90 <= v93 {
		v71 = v90
		v73 = v93
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	goto L34
L51:
	;
	v117 = int32(2)
	v118 = base.I32_div_s(v112-v110, v117)
	v119 = v118 + v110
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v108+v119<<(uint(v117)%32))))
	v124 = base.B2i32(v123 == v27)
	if v123 == v27 {
		v139 = v124
		goto L28
	} else {
		goto L53
	}
L52:
	;
	v139 = v124
	goto L28
L53:
	;
	v127 = base.B2i32(base.Ui32(v123) < base.Ui32(v27))
	if base.Ui32(v123) < base.Ui32(v27) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v128 = v119 + int32(1)
	goto L56
L55:
	;
	v128 = v110
	goto L56
L56:
	;
	if base.Ui32(v123) < base.Ui32(v27) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v131 = v112
	goto L59
L58:
	;
	v131 = v119 - int32(1)
	goto L59
L59:
	;
	if v128 <= v131 {
		v110 = v128
		v112 = v131
		goto L51
	} else {
		goto L60
	}
L60:
	;
	goto L52
L61:
	;
	v3640 = int32(0)
	goto L1
L62:
	;
	goto L63
L63:
	;
	v149 = F_XidInMVCCSnapshot(m, v27, l1)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	return int32(0)
L65:
	;
	if v149 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v153 = F_TransactionIdDidCommit(m, v27)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)))
	if v153 == int32(0) {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v159 = v155 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)) = uint16(v159)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v3640 = int32(0)
	goto L1
L70:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if base.Ui32(v167) < base.Ui32(int32(3)) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if base.Ui32(v307) < base.Ui32(int32(3)) {
		goto L125
	} else {
		goto L126
	}
L73:
	;
	if v287 != 0 {
		goto L10
	} else {
		goto L113
	}
L74:
	;
	v287 = int32(0)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v178 == v167 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v287 = int32(1)
	goto L73
L78:
	;
	goto L79
L79:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v182 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v287 = v279
	goto L73
L81:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v186 == int32(0) {
		v279 = int32(0)
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v250 = int32(0)
	v252 = v182 - int32(1)
	goto L103
L84:
	;
	v191 = v186
	goto L85
L85:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	if v196 == int32(4) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v279 = int32(0)
	goto L80
L87:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v191)+80))
	if v243 != 0 {
		v191 = v243
		goto L85
	} else {
		goto L102
	}
L88:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v199 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v202 = int32(1)
	if v167 == v199 {
		v279 = v202
		goto L80
	} else {
		goto L90
	}
L90:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v191)+52))
	v206 = v204 - int32(1)
	if v206 < int32(0) {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	v211 = int32(0)
	v213 = v206
	goto L92
L92:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v191)+48))
	v219 = int32(2)
	v220 = base.I32_div_s(v213-v211, v219)
	v221 = v220 + v211
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v217+v221<<(uint(v219)%32))))
	if v225 == v167 {
		v279 = v202
		goto L80
	} else {
		goto L94
	}
L93:
	;
	goto L87
L94:
	;
	v229 = F_TransactionIdPrecedes(m, v225, v167)
	mBase = m.M
	if v229 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v230 = v221 + int32(1)
	goto L97
L96:
	;
	v230 = v211
	goto L97
L97:
	;
	if v229 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v233 = v213
	goto L100
L99:
	;
	v233 = v221 - int32(1)
	goto L100
L100:
	;
	if v230 <= v233 {
		v211 = v230
		v213 = v233
		goto L92
	} else {
		goto L101
	}
L101:
	;
	goto L93
L102:
	;
	goto L86
L103:
	;
	v257 = int32(2)
	v258 = base.I32_div_s(v252-v250, v257)
	v259 = v258 + v250
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v248+v259<<(uint(v257)%32))))
	v264 = base.B2i32(v263 == v167)
	if v263 == v167 {
		v279 = v264
		goto L80
	} else {
		goto L105
	}
L104:
	;
	v279 = v264
	goto L80
L105:
	;
	v267 = base.B2i32(base.Ui32(v263) < base.Ui32(v167))
	if base.Ui32(v263) < base.Ui32(v167) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v268 = v259 + int32(1)
	goto L108
L107:
	;
	v268 = v250
	goto L108
L108:
	;
	if base.Ui32(v263) < base.Ui32(v167) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v271 = v252
	goto L111
L110:
	;
	v271 = v259 - int32(1)
	goto L111
L111:
	;
	if v268 <= v271 {
		v250 = v268
		v252 = v271
		goto L103
	} else {
		goto L112
	}
L112:
	;
	goto L104
L113:
	;
	v288 = F_XidInMVCCSnapshot(m, v167, l1)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L64
	} else {
		goto L114
	}
L114:
	;
	if v288 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v3640 = int32(0)
	goto L1
L116:
	;
	goto L117
L117:
	;
	v291 = F_TransactionIdDidCommit(m, v167)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L64
	} else {
		goto L118
	}
L118:
	;
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)))
	if v291 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v295 = v293 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)) = uint16(v295)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L64
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v301 = v293 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)) = uint16(v301)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L64
	} else {
		goto L123
	}
L122:
	;
	goto L10
L123:
	;
	v3640 = int32(0)
	goto L1
L124:
	;
	if v427 != 0 {
		goto L164
	} else {
		goto L165
	}
L125:
	;
	v427 = int32(0)
	goto L124
L126:
	;
	goto L127
L127:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v318 == v307 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v427 = int32(1)
	goto L124
L129:
	;
	goto L130
L130:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v322 <= int32(0) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v427 = v419
	goto L124
L132:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v326 == int32(0) {
		v419 = int32(0)
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v390 = int32(0)
	v392 = v322 - int32(1)
	goto L154
L135:
	;
	v331 = v326
	goto L136
L136:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v331)+20))
	if v336 == int32(4) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v419 = int32(0)
	goto L131
L138:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v331)+80))
	if v383 != 0 {
		v331 = v383
		goto L136
	} else {
		goto L153
	}
L139:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	if v339 == int32(0) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v342 = int32(1)
	if v307 == v339 {
		v419 = v342
		goto L131
	} else {
		goto L141
	}
L141:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v331)+52))
	v346 = v344 - int32(1)
	if v346 < int32(0) {
		goto L138
	} else {
		goto L142
	}
L142:
	;
	v351 = int32(0)
	v353 = v346
	goto L143
L143:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v331)+48))
	v359 = int32(2)
	v360 = base.I32_div_s(v353-v351, v359)
	v361 = v360 + v351
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v357+v361<<(uint(v359)%32))))
	if v365 == v307 {
		v419 = v342
		goto L131
	} else {
		goto L145
	}
L144:
	;
	goto L138
L145:
	;
	v369 = F_TransactionIdPrecedes(m, v365, v307)
	mBase = m.M
	if v369 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v370 = v361 + int32(1)
	goto L148
L147:
	;
	v370 = v351
	goto L148
L148:
	;
	if v369 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v373 = v353
	goto L151
L150:
	;
	v373 = v361 - int32(1)
	goto L151
L151:
	;
	if v370 <= v373 {
		v351 = v370
		v353 = v373
		goto L143
	} else {
		goto L152
	}
L152:
	;
	goto L144
L153:
	;
	goto L137
L154:
	;
	v397 = int32(2)
	v398 = base.I32_div_s(v392-v390, v397)
	v399 = v398 + v390
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v388+v399<<(uint(v397)%32))))
	v404 = base.B2i32(v403 == v307)
	if v403 == v307 {
		v419 = v404
		goto L131
	} else {
		goto L156
	}
L155:
	;
	v419 = v404
	goto L131
L156:
	;
	v407 = base.B2i32(base.Ui32(v403) < base.Ui32(v307))
	if base.Ui32(v403) < base.Ui32(v307) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v408 = v399 + int32(1)
	goto L159
L158:
	;
	v408 = v390
	goto L159
L159:
	;
	if base.Ui32(v403) < base.Ui32(v307) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v411 = v392
	goto L162
L161:
	;
	v411 = v399 - int32(1)
	goto L162
L162:
	;
	if v408 <= v411 {
		v390 = v408
		v392 = v411
		goto L154
	} else {
		goto L163
	}
L163:
	;
	goto L155
L164:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v431&int32(32) != 0 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	goto L166
L166:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v735 = F_XidInMVCCSnapshot(m, v734, l1)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L64
	} else {
		goto L271
	}
L167:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if base.Ui32(v441) <= base.Ui32(v440) {
		v3640 = int32(0)
		goto L1
	} else {
		goto L171
	}
L168:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v435+v430<<(uint(int32(3))%32))))
	v440 = v439
	goto L170
L169:
	;
	v440 = v430
	goto L170
L170:
	;
	goto L167
L171:
	;
	v443 = int32(1)
	v444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)))
	if v444&int32(2048) != 0 {
		v3640 = v443
		goto L1
	} else {
		goto L172
	}
L172:
	;
	if v444&int32(128) != 0 {
		v3640 = v443
		goto L1
	} else {
		goto L173
	}
L173:
	;
	if v444&int32(4176) == int32(64) {
		v3640 = v443
		goto L1
	} else {
		goto L174
	}
L174:
	;
	if v444&int32(4096) != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v455 = F_HeapTupleGetUpdateXid(m, v15)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L64
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if base.Ui32(v593) < base.Ui32(int32(3)) {
		goto L225
	} else {
		goto L226
	}
L178:
	;
	if base.Ui32(v455) < base.Ui32(int32(3)) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v576 == int32(0) {
		v3640 = v443
		goto L1
	} else {
		goto L219
	}
L180:
	;
	v576 = int32(0)
	goto L179
L181:
	;
	goto L182
L182:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v467 == v455 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v576 = int32(1)
	goto L179
L184:
	;
	goto L185
L185:
	;
	v471 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v471 <= int32(0) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v576 = v568
	goto L179
L187:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v475 == int32(0) {
		v568 = int32(0)
		goto L186
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v539 = int32(0)
	v541 = v471 - int32(1)
	goto L209
L190:
	;
	v480 = v475
	goto L191
L191:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v480)+20))
	if v485 == int32(4) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v568 = int32(0)
	goto L186
L193:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v480)+80))
	if v532 != 0 {
		v480 = v532
		goto L191
	} else {
		goto L208
	}
L194:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	if v488 == int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v491 = int32(1)
	if v455 == v488 {
		v568 = v491
		goto L186
	} else {
		goto L196
	}
L196:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v480)+52))
	v495 = v493 - int32(1)
	if v495 < int32(0) {
		goto L193
	} else {
		goto L197
	}
L197:
	;
	v500 = int32(0)
	v502 = v495
	goto L198
L198:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v480)+48))
	v508 = int32(2)
	v509 = base.I32_div_s(v502-v500, v508)
	v510 = v509 + v500
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v506+v510<<(uint(v508)%32))))
	if v514 == v455 {
		v568 = v491
		goto L186
	} else {
		goto L200
	}
L199:
	;
	goto L193
L200:
	;
	v518 = F_TransactionIdPrecedes(m, v514, v455)
	mBase = m.M
	if v518 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v519 = v510 + int32(1)
	goto L203
L202:
	;
	v519 = v500
	goto L203
L203:
	;
	if v518 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v522 = v502
	goto L206
L205:
	;
	v522 = v510 - int32(1)
	goto L206
L206:
	;
	if v519 <= v522 {
		v500 = v519
		v502 = v522
		goto L198
	} else {
		goto L207
	}
L207:
	;
	goto L199
L208:
	;
	goto L192
L209:
	;
	v546 = int32(2)
	v547 = base.I32_div_s(v541-v539, v546)
	v548 = v547 + v539
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v537+v548<<(uint(v546)%32))))
	v553 = base.B2i32(v552 == v455)
	if v552 == v455 {
		v568 = v553
		goto L186
	} else {
		goto L211
	}
L210:
	;
	v568 = v553
	goto L186
L211:
	;
	v556 = base.B2i32(base.Ui32(v552) < base.Ui32(v455))
	if base.Ui32(v552) < base.Ui32(v455) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v557 = v548 + int32(1)
	goto L214
L213:
	;
	v557 = v539
	goto L214
L214:
	;
	if base.Ui32(v552) < base.Ui32(v455) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v560 = v541
	goto L217
L216:
	;
	v560 = v548 - int32(1)
	goto L217
L217:
	;
	if v557 <= v560 {
		v539 = v557
		v541 = v560
		goto L209
	} else {
		goto L218
	}
L218:
	;
	goto L210
L219:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v581&int32(32) != 0 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3640 = base.B2i32(base.Ui32(v591) <= base.Ui32(v590))
	goto L1
L221:
	;
	v585 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v585+v580<<(uint(int32(3))%32))+4))
	v590 = v589
	goto L223
L222:
	;
	v590 = v580
	goto L223
L223:
	;
	goto L220
L224:
	;
	if v713 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L225:
	;
	v713 = int32(0)
	goto L224
L226:
	;
	goto L227
L227:
	;
	v604 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v604 == v593 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v713 = int32(1)
	goto L224
L229:
	;
	goto L230
L230:
	;
	v608 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v608 <= int32(0) {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v713 = v705
	goto L224
L232:
	;
	v612 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v612 == int32(0) {
		v705 = int32(0)
		goto L231
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v676 = int32(0)
	v678 = v608 - int32(1)
	goto L254
L235:
	;
	v617 = v612
	goto L236
L236:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v617)+20))
	if v622 == int32(4) {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v705 = int32(0)
	goto L231
L238:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v617)+80))
	if v669 != 0 {
		v617 = v669
		goto L236
	} else {
		goto L253
	}
L239:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	if v625 == int32(0) {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v628 = int32(1)
	if v593 == v625 {
		v705 = v628
		goto L231
	} else {
		goto L241
	}
L241:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v617)+52))
	v632 = v630 - int32(1)
	if v632 < int32(0) {
		goto L238
	} else {
		goto L242
	}
L242:
	;
	v637 = int32(0)
	v639 = v632
	goto L243
L243:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v617)+48))
	v645 = int32(2)
	v646 = base.I32_div_s(v639-v637, v645)
	v647 = v646 + v637
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v643+v647<<(uint(v645)%32))))
	if v651 == v593 {
		v705 = v628
		goto L231
	} else {
		goto L245
	}
L244:
	;
	goto L238
L245:
	;
	v655 = F_TransactionIdPrecedes(m, v651, v593)
	mBase = m.M
	if v655 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v656 = v647 + int32(1)
	goto L248
L247:
	;
	v656 = v637
	goto L248
L248:
	;
	if v655 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v659 = v639
	goto L251
L250:
	;
	v659 = v647 - int32(1)
	goto L251
L251:
	;
	if v656 <= v659 {
		v637 = v656
		v639 = v659
		goto L243
	} else {
		goto L252
	}
L252:
	;
	goto L244
L253:
	;
	goto L237
L254:
	;
	v683 = int32(2)
	v684 = base.I32_div_s(v678-v676, v683)
	v685 = v684 + v676
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v674+v685<<(uint(v683)%32))))
	v690 = base.B2i32(v689 == v593)
	if v689 == v593 {
		v705 = v690
		goto L231
	} else {
		goto L256
	}
L255:
	;
	v705 = v690
	goto L231
L256:
	;
	v693 = base.B2i32(base.Ui32(v689) < base.Ui32(v593))
	if base.Ui32(v689) < base.Ui32(v593) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v694 = v685 + int32(1)
	goto L259
L258:
	;
	v694 = v676
	goto L259
L259:
	;
	if base.Ui32(v689) < base.Ui32(v593) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v697 = v678
	goto L262
L261:
	;
	v697 = v685 - int32(1)
	goto L262
L262:
	;
	if v694 <= v697 {
		v676 = v694
		v678 = v697
		goto L254
	} else {
		goto L263
	}
L263:
	;
	goto L255
L264:
	;
	v716 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)))
	v718 = v716 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)) = uint16(v718)
	goto L2
L265:
	;
	goto L266
L266:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v722&int32(32) != 0 {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3640 = base.B2i32(base.Ui32(v732) <= base.Ui32(v731))
	goto L1
L268:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v726+v721<<(uint(int32(3))%32))+4))
	v731 = v730
	goto L270
L269:
	;
	v731 = v721
	goto L270
L270:
	;
	goto L267
L271:
	;
	if v735 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v3640 = int32(0)
	goto L1
L273:
	;
	goto L274
L274:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v739 = F_TransactionIdDidCommit(m, v738)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L64
	} else {
		goto L275
	}
L275:
	;
	if v739 != 0 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	F_HeapTupleSetHintBits(m, v15, l2, int32(256), v742)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L64
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v745 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)))
	v747 = v745 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)) = uint16(v747)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L64
	} else {
		goto L280
	}
L279:
	;
	goto L10
L280:
	;
	v3640 = int32(0)
	goto L1
L281:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v758 = F_XidInMVCCSnapshot(m, v757, l1)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L64
	} else {
		goto L282
	}
L282:
	;
	if v758 == int32(0) {
		goto L10
	} else {
		goto L283
	}
L283:
	;
	v3640 = int32(0)
	goto L1
L284:
	;
	goto L10
L285:
	;
	if v772&int32(128) != 0 {
		v3640 = v771
		goto L1
	} else {
		goto L286
	}
L286:
	;
	if v772&int32(4176) == int32(64) {
		v3640 = v771
		goto L1
	} else {
		goto L287
	}
L287:
	;
	if v772&int32(4096) != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v783 = F_HeapTupleGetUpdateXid(m, v15)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L64
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v772&int32(1024) == int32(0) {
		goto L342
	} else {
		goto L343
	}
L291:
	;
	if base.Ui32(v783) < base.Ui32(int32(3)) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	if v904 != 0 {
		goto L332
	} else {
		goto L333
	}
L293:
	;
	v904 = int32(0)
	goto L292
L294:
	;
	goto L295
L295:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v795 == v783 {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v904 = int32(1)
	goto L292
L297:
	;
	goto L298
L298:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v799 <= int32(0) {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v904 = v896
	goto L292
L300:
	;
	v803 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v803 == int32(0) {
		v896 = int32(0)
		goto L299
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v865 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v867 = int32(0)
	v869 = v799 - int32(1)
	goto L322
L303:
	;
	v808 = v803
	goto L304
L304:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v808)+20))
	if v813 == int32(4) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v896 = int32(0)
	goto L299
L306:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v808)+80))
	if v860 != 0 {
		v808 = v860
		goto L304
	} else {
		goto L321
	}
L307:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	if v816 == int32(0) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v819 = int32(1)
	if v783 == v816 {
		v896 = v819
		goto L299
	} else {
		goto L309
	}
L309:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v808)+52))
	v823 = v821 - int32(1)
	if v823 < int32(0) {
		goto L306
	} else {
		goto L310
	}
L310:
	;
	v828 = int32(0)
	v830 = v823
	goto L311
L311:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v808)+48))
	v836 = int32(2)
	v837 = base.I32_div_s(v830-v828, v836)
	v838 = v837 + v828
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v834+v838<<(uint(v836)%32))))
	if v842 == v783 {
		v896 = v819
		goto L299
	} else {
		goto L313
	}
L312:
	;
	goto L306
L313:
	;
	v846 = F_TransactionIdPrecedes(m, v842, v783)
	mBase = m.M
	if v846 != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v847 = v838 + int32(1)
	goto L316
L315:
	;
	v847 = v828
	goto L316
L316:
	;
	if v846 != 0 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v850 = v830
	goto L319
L318:
	;
	v850 = v838 - int32(1)
	goto L319
L319:
	;
	if v847 <= v850 {
		v828 = v847
		v830 = v850
		goto L311
	} else {
		goto L320
	}
L320:
	;
	goto L312
L321:
	;
	goto L305
L322:
	;
	v874 = int32(2)
	v875 = base.I32_div_s(v869-v867, v874)
	v876 = v875 + v867
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v865+v876<<(uint(v874)%32))))
	v881 = base.B2i32(v880 == v783)
	if v880 == v783 {
		v896 = v881
		goto L299
	} else {
		goto L324
	}
L323:
	;
	v896 = v881
	goto L299
L324:
	;
	v884 = base.B2i32(base.Ui32(v880) < base.Ui32(v783))
	if base.Ui32(v880) < base.Ui32(v783) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v885 = v876 + int32(1)
	goto L327
L326:
	;
	v885 = v867
	goto L327
L327:
	;
	if base.Ui32(v880) < base.Ui32(v783) {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v888 = v869
	goto L330
L329:
	;
	v888 = v876 - int32(1)
	goto L330
L330:
	;
	if v885 <= v888 {
		v867 = v885
		v869 = v888
		goto L322
	} else {
		goto L331
	}
L331:
	;
	goto L323
L332:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v907&int32(32) != 0 {
		goto L336
	} else {
		goto L337
	}
L333:
	;
	goto L334
L334:
	;
	v919 = F_XidInMVCCSnapshot(m, v783, l1)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L64
	} else {
		goto L339
	}
L335:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3640 = base.B2i32(base.Ui32(v917) <= base.Ui32(v916))
	goto L1
L336:
	;
	v911 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v911+v906<<(uint(int32(3))%32))+4))
	v916 = v915
	goto L338
L337:
	;
	v916 = v906
	goto L338
L338:
	;
	goto L335
L339:
	;
	if v919 != 0 {
		v3640 = v771
		goto L1
	} else {
		goto L340
	}
L340:
	;
	v921 = F_TransactionIdDidCommit(m, v783)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L64
	} else {
		goto L341
	}
L341:
	;
	v3640 = v921 ^ int32(1)
	goto L1
L342:
	;
	if base.Ui32(v925) < base.Ui32(int32(3)) {
		goto L346
	} else {
		goto L347
	}
L343:
	;
	goto L344
L344:
	;
	v1081 = F_XidInMVCCSnapshot(m, v925, l1)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L64
	} else {
		goto L399
	}
L345:
	;
	if v1049 != 0 {
		goto L385
	} else {
		goto L386
	}
L346:
	;
	v1049 = int32(0)
	goto L345
L347:
	;
	goto L348
L348:
	;
	v940 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v940 == v925 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1049 = int32(1)
	goto L345
L350:
	;
	goto L351
L351:
	;
	v944 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v944 <= int32(0) {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	v1049 = v1041
	goto L345
L353:
	;
	v948 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v948 == int32(0) {
		v1041 = int32(0)
		goto L352
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v1012 = int32(0)
	v1014 = v944 - int32(1)
	goto L375
L356:
	;
	v953 = v948
	goto L357
L357:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v953)+20))
	if v958 == int32(4) {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	v1041 = int32(0)
	goto L352
L359:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v953)+80))
	if v1005 != 0 {
		v953 = v1005
		goto L357
	} else {
		goto L374
	}
L360:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v953)))
	if v961 == int32(0) {
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v964 = int32(1)
	if v925 == v961 {
		v1041 = v964
		goto L352
	} else {
		goto L362
	}
L362:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v953)+52))
	v968 = v966 - int32(1)
	if v968 < int32(0) {
		goto L359
	} else {
		goto L363
	}
L363:
	;
	v973 = int32(0)
	v975 = v968
	goto L364
L364:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v953)+48))
	v981 = int32(2)
	v982 = base.I32_div_s(v975-v973, v981)
	v983 = v982 + v973
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v979+v983<<(uint(v981)%32))))
	if v987 == v925 {
		v1041 = v964
		goto L352
	} else {
		goto L366
	}
L365:
	;
	goto L359
L366:
	;
	v991 = F_TransactionIdPrecedes(m, v987, v925)
	mBase = m.M
	if v991 != 0 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v992 = v983 + int32(1)
	goto L369
L368:
	;
	v992 = v973
	goto L369
L369:
	;
	if v991 != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v995 = v975
	goto L372
L371:
	;
	v995 = v983 - int32(1)
	goto L372
L372:
	;
	if v992 <= v995 {
		v973 = v992
		v975 = v995
		goto L364
	} else {
		goto L373
	}
L373:
	;
	goto L365
L374:
	;
	goto L358
L375:
	;
	v1019 = int32(2)
	v1020 = base.I32_div_s(v1014-v1012, v1019)
	v1021 = v1020 + v1012
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1010+v1021<<(uint(v1019)%32))))
	v1026 = base.B2i32(v1025 == v925)
	if v1025 == v925 {
		v1041 = v1026
		goto L352
	} else {
		goto L377
	}
L376:
	;
	v1041 = v1026
	goto L352
L377:
	;
	v1029 = base.B2i32(base.Ui32(v1025) < base.Ui32(v925))
	if base.Ui32(v1025) < base.Ui32(v925) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1030 = v1021 + int32(1)
	goto L380
L379:
	;
	v1030 = v1012
	goto L380
L380:
	;
	if base.Ui32(v1025) < base.Ui32(v925) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v1033 = v1014
	goto L383
L382:
	;
	v1033 = v1021 - int32(1)
	goto L383
L383:
	;
	if v1030 <= v1033 {
		v1012 = v1030
		v1014 = v1033
		goto L375
	} else {
		goto L384
	}
L384:
	;
	goto L376
L385:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v1052&int32(32) != 0 {
		goto L389
	} else {
		goto L390
	}
L386:
	;
	goto L387
L387:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v1065 = F_XidInMVCCSnapshot(m, v1064, l1)
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L64
	} else {
		goto L392
	}
L388:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3640 = base.B2i32(base.Ui32(v1062) <= base.Ui32(v1061))
	goto L1
L389:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1056+v1051<<(uint(int32(3))%32))+4))
	v1061 = v1060
	goto L391
L390:
	;
	v1061 = v1051
	goto L391
L391:
	;
	goto L388
L392:
	;
	if v1065 != 0 {
		v3640 = v771
		goto L1
	} else {
		goto L393
	}
L393:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v1068 = F_TransactionIdDidCommit(m, v1067)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L64
	} else {
		goto L394
	}
L394:
	;
	if v1068 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)))
	v1074 = v1072 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+20)) = uint16(v1074)
	goto L2
L396:
	;
	goto L397
L397:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	F_HeapTupleSetHintBits(m, v15, l2, int32(1024), v1077)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L64
	} else {
		goto L398
	}
L398:
	;
	v3640 = int32(0)
	goto L1
L399:
	;
	if v1081 != 0 {
		v3640 = v771
		goto L1
	} else {
		goto L400
	}
L400:
	;
	v3640 = int32(0)
	goto L1
L401:
	;
	v1782 = int32(1)
	v1783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)))
	if v1783&int32(2048) != 0 {
		v3640 = v1782
		goto L1
	} else {
		goto L654
	}
L402:
	;
	v1088 = base.I32_extend16_s(v1085)
	if v1088&int32(512) != 0 {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v3640 = int32(0)
	goto L1
L404:
	;
	goto L405
L405:
	;
	if v1088&int32(16384) != 0 {
		goto L407
	} else {
		goto L408
	}
L406:
	;
	v1775 = v1220 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)) = uint16(v1775)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L64
	} else {
		goto L653
	}
L407:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+8))
	if base.Ui32(v1094) < base.Ui32(int32(3)) {
		goto L411
	} else {
		goto L412
	}
L408:
	;
	goto L409
L409:
	;
	if v1088 < int32(0) {
		goto L458
	} else {
		goto L459
	}
L410:
	;
	if v1214 != 0 {
		goto L450
	} else {
		goto L451
	}
L411:
	;
	v1214 = int32(0)
	goto L410
L412:
	;
	goto L413
L413:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v1105 == v1094 {
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v1214 = int32(1)
	goto L410
L415:
	;
	goto L416
L416:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v1109 <= int32(0) {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	v1214 = v1206
	goto L410
L418:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v1113 == int32(0) {
		v1206 = int32(0)
		goto L417
	} else {
		goto L421
	}
L419:
	;
	goto L420
L420:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v1177 = int32(0)
	v1179 = v1109 - int32(1)
	goto L440
L421:
	;
	v1118 = v1113
	goto L422
L422:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1118)+20))
	if v1123 == int32(4) {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	v1206 = int32(0)
	goto L417
L424:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1118)+80))
	if v1170 != 0 {
		v1118 = v1170
		goto L422
	} else {
		goto L439
	}
L425:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1118)))
	if v1126 == int32(0) {
		goto L424
	} else {
		goto L426
	}
L426:
	;
	v1129 = int32(1)
	if v1094 == v1126 {
		v1206 = v1129
		goto L417
	} else {
		goto L427
	}
L427:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1118)+52))
	v1133 = v1131 - int32(1)
	if v1133 < int32(0) {
		goto L424
	} else {
		goto L428
	}
L428:
	;
	v1138 = int32(0)
	v1140 = v1133
	goto L429
L429:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1118)+48))
	v1146 = int32(2)
	v1147 = base.I32_div_s(v1140-v1138, v1146)
	v1148 = v1147 + v1138
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1144+v1148<<(uint(v1146)%32))))
	if v1152 == v1094 {
		v1206 = v1129
		goto L417
	} else {
		goto L431
	}
L430:
	;
	goto L424
L431:
	;
	v1156 = F_TransactionIdPrecedes(m, v1152, v1094)
	mBase = m.M
	if v1156 != 0 {
		goto L432
	} else {
		goto L433
	}
L432:
	;
	v1157 = v1148 + int32(1)
	goto L434
L433:
	;
	v1157 = v1138
	goto L434
L434:
	;
	if v1156 != 0 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1160 = v1140
	goto L437
L436:
	;
	v1160 = v1148 - int32(1)
	goto L437
L437:
	;
	if v1157 <= v1160 {
		v1138 = v1157
		v1140 = v1160
		goto L429
	} else {
		goto L438
	}
L438:
	;
	goto L430
L439:
	;
	goto L423
L440:
	;
	v1184 = int32(2)
	v1185 = base.I32_div_s(v1179-v1177, v1184)
	v1186 = v1185 + v1177
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1175+v1186<<(uint(v1184)%32))))
	v1191 = base.B2i32(v1190 == v1094)
	if v1190 == v1094 {
		v1206 = v1191
		goto L417
	} else {
		goto L442
	}
L441:
	;
	v1206 = v1191
	goto L417
L442:
	;
	v1194 = base.B2i32(base.Ui32(v1190) < base.Ui32(v1094))
	if base.Ui32(v1190) < base.Ui32(v1094) {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v1195 = v1186 + int32(1)
	goto L445
L444:
	;
	v1195 = v1177
	goto L445
L445:
	;
	if base.Ui32(v1190) < base.Ui32(v1094) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1198 = v1179
	goto L448
L447:
	;
	v1198 = v1186 - int32(1)
	goto L448
L448:
	;
	if v1195 <= v1198 {
		v1177 = v1195
		v1179 = v1198
		goto L440
	} else {
		goto L449
	}
L449:
	;
	goto L441
L450:
	;
	v3640 = int32(0)
	goto L1
L451:
	;
	goto L452
L452:
	;
	v1216 = F_TransactionIdIsInProgress(m, v1094)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L64
	} else {
		goto L453
	}
L453:
	;
	if v1216 != 0 {
		goto L401
	} else {
		goto L454
	}
L454:
	;
	v1218 = F_TransactionIdDidCommit(m, v1094)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L64
	} else {
		goto L455
	}
L455:
	;
	v1220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)))
	if v1218 == int32(0) {
		goto L406
	} else {
		goto L456
	}
L456:
	;
	v1224 = v1220 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)) = uint16(v1224)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L64
	} else {
		goto L457
	}
L457:
	;
	v3640 = int32(0)
	goto L1
L458:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+8))
	if base.Ui32(v1232) < base.Ui32(int32(3)) {
		goto L462
	} else {
		goto L463
	}
L459:
	;
	goto L460
L460:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1084)))
	if base.Ui32(v1372) < base.Ui32(int32(3)) {
		goto L513
	} else {
		goto L514
	}
L461:
	;
	if v1352 != 0 {
		goto L401
	} else {
		goto L501
	}
L462:
	;
	v1352 = int32(0)
	goto L461
L463:
	;
	goto L464
L464:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v1243 == v1232 {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v1352 = int32(1)
	goto L461
L466:
	;
	goto L467
L467:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v1247 <= int32(0) {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	v1352 = v1344
	goto L461
L469:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v1251 == int32(0) {
		v1344 = int32(0)
		goto L468
	} else {
		goto L472
	}
L470:
	;
	goto L471
L471:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v1315 = int32(0)
	v1317 = v1247 - int32(1)
	goto L491
L472:
	;
	v1256 = v1251
	goto L473
L473:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+20))
	if v1261 == int32(4) {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	v1344 = int32(0)
	goto L468
L475:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+80))
	if v1308 != 0 {
		v1256 = v1308
		goto L473
	} else {
		goto L490
	}
L476:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1256)))
	if v1264 == int32(0) {
		goto L475
	} else {
		goto L477
	}
L477:
	;
	v1267 = int32(1)
	if v1232 == v1264 {
		v1344 = v1267
		goto L468
	} else {
		goto L478
	}
L478:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+52))
	v1271 = v1269 - int32(1)
	if v1271 < int32(0) {
		goto L475
	} else {
		goto L479
	}
L479:
	;
	v1276 = int32(0)
	v1278 = v1271
	goto L480
L480:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+48))
	v1284 = int32(2)
	v1285 = base.I32_div_s(v1278-v1276, v1284)
	v1286 = v1285 + v1276
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v1282+v1286<<(uint(v1284)%32))))
	if v1290 == v1232 {
		v1344 = v1267
		goto L468
	} else {
		goto L482
	}
L481:
	;
	goto L475
L482:
	;
	v1294 = F_TransactionIdPrecedes(m, v1290, v1232)
	mBase = m.M
	if v1294 != 0 {
		goto L483
	} else {
		goto L484
	}
L483:
	;
	v1295 = v1286 + int32(1)
	goto L485
L484:
	;
	v1295 = v1276
	goto L485
L485:
	;
	if v1294 != 0 {
		goto L486
	} else {
		goto L487
	}
L486:
	;
	v1298 = v1278
	goto L488
L487:
	;
	v1298 = v1286 - int32(1)
	goto L488
L488:
	;
	if v1295 <= v1298 {
		v1276 = v1295
		v1278 = v1298
		goto L480
	} else {
		goto L489
	}
L489:
	;
	goto L481
L490:
	;
	goto L474
L491:
	;
	v1322 = int32(2)
	v1323 = base.I32_div_s(v1317-v1315, v1322)
	v1324 = v1323 + v1315
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1313+v1324<<(uint(v1322)%32))))
	v1329 = base.B2i32(v1328 == v1232)
	if v1328 == v1232 {
		v1344 = v1329
		goto L468
	} else {
		goto L493
	}
L492:
	;
	v1344 = v1329
	goto L468
L493:
	;
	v1332 = base.B2i32(base.Ui32(v1328) < base.Ui32(v1232))
	if base.Ui32(v1328) < base.Ui32(v1232) {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v1333 = v1324 + int32(1)
	goto L496
L495:
	;
	v1333 = v1315
	goto L496
L496:
	;
	if base.Ui32(v1328) < base.Ui32(v1232) {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v1336 = v1317
	goto L499
L498:
	;
	v1336 = v1324 - int32(1)
	goto L499
L499:
	;
	if v1333 <= v1336 {
		v1315 = v1333
		v1317 = v1336
		goto L491
	} else {
		goto L500
	}
L500:
	;
	goto L492
L501:
	;
	v1353 = F_TransactionIdIsInProgress(m, v1232)
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L64
	} else {
		goto L502
	}
L502:
	;
	if v1353 != 0 {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	v3640 = int32(0)
	goto L1
L504:
	;
	goto L505
L505:
	;
	v1356 = F_TransactionIdDidCommit(m, v1232)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L64
	} else {
		goto L506
	}
L506:
	;
	v1358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)))
	if v1356 != 0 {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v1360 = v1358 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)) = uint16(v1360)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L64
	} else {
		goto L510
	}
L508:
	;
	goto L509
L509:
	;
	v1366 = v1358 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)) = uint16(v1366)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L64
	} else {
		goto L511
	}
L510:
	;
	goto L401
L511:
	;
	v3640 = int32(0)
	goto L1
L512:
	;
	if v1492 != 0 {
		goto L552
	} else {
		goto L553
	}
L513:
	;
	v1492 = int32(0)
	goto L512
L514:
	;
	goto L515
L515:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v1383 == v1372 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v1492 = int32(1)
	goto L512
L517:
	;
	goto L518
L518:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v1387 <= int32(0) {
		goto L520
	} else {
		goto L521
	}
L519:
	;
	v1492 = v1484
	goto L512
L520:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v1391 == int32(0) {
		v1484 = int32(0)
		goto L519
	} else {
		goto L523
	}
L521:
	;
	goto L522
L522:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v1455 = int32(0)
	v1457 = v1387 - int32(1)
	goto L542
L523:
	;
	v1396 = v1391
	goto L524
L524:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+20))
	if v1401 == int32(4) {
		goto L526
	} else {
		goto L527
	}
L525:
	;
	v1484 = int32(0)
	goto L519
L526:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+80))
	if v1448 != 0 {
		v1396 = v1448
		goto L524
	} else {
		goto L541
	}
L527:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1396)))
	if v1404 == int32(0) {
		goto L526
	} else {
		goto L528
	}
L528:
	;
	v1407 = int32(1)
	if v1372 == v1404 {
		v1484 = v1407
		goto L519
	} else {
		goto L529
	}
L529:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+52))
	v1411 = v1409 - int32(1)
	if v1411 < int32(0) {
		goto L526
	} else {
		goto L530
	}
L530:
	;
	v1416 = int32(0)
	v1418 = v1411
	goto L531
L531:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+48))
	v1424 = int32(2)
	v1425 = base.I32_div_s(v1418-v1416, v1424)
	v1426 = v1425 + v1416
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1422+v1426<<(uint(v1424)%32))))
	if v1430 == v1372 {
		v1484 = v1407
		goto L519
	} else {
		goto L533
	}
L532:
	;
	goto L526
L533:
	;
	v1434 = F_TransactionIdPrecedes(m, v1430, v1372)
	mBase = m.M
	if v1434 != 0 {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v1435 = v1426 + int32(1)
	goto L536
L535:
	;
	v1435 = v1416
	goto L536
L536:
	;
	if v1434 != 0 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v1438 = v1418
	goto L539
L538:
	;
	v1438 = v1426 - int32(1)
	goto L539
L539:
	;
	if v1435 <= v1438 {
		v1416 = v1435
		v1418 = v1438
		goto L531
	} else {
		goto L540
	}
L540:
	;
	goto L532
L541:
	;
	goto L525
L542:
	;
	v1462 = int32(2)
	v1463 = base.I32_div_s(v1457-v1455, v1462)
	v1464 = v1463 + v1455
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1453+v1464<<(uint(v1462)%32))))
	v1469 = base.B2i32(v1468 == v1372)
	if v1468 == v1372 {
		v1484 = v1469
		goto L519
	} else {
		goto L544
	}
L543:
	;
	v1484 = v1469
	goto L519
L544:
	;
	v1472 = base.B2i32(base.Ui32(v1468) < base.Ui32(v1372))
	if base.Ui32(v1468) < base.Ui32(v1372) {
		goto L545
	} else {
		goto L546
	}
L545:
	;
	v1473 = v1464 + int32(1)
	goto L547
L546:
	;
	v1473 = v1455
	goto L547
L547:
	;
	if base.Ui32(v1468) < base.Ui32(v1372) {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	v1476 = v1457
	goto L550
L549:
	;
	v1476 = v1464 - int32(1)
	goto L550
L550:
	;
	if v1473 <= v1476 {
		v1455 = v1473
		v1457 = v1476
		goto L542
	} else {
		goto L551
	}
L551:
	;
	goto L543
L552:
	;
	v1493 = int32(1)
	v1494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)))
	if v1494&int32(2048) != 0 {
		v3640 = v1493
		goto L1
	} else {
		goto L555
	}
L553:
	;
	goto L554
L554:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1084)))
	v1756 = F_TransactionIdIsInProgress(m, v1755)
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L64
	} else {
		goto L643
	}
L555:
	;
	if v1494&int32(128) != 0 {
		v3640 = v1493
		goto L1
	} else {
		goto L556
	}
L556:
	;
	if v1494&int32(4176) == int32(64) {
		v3640 = v1493
		goto L1
	} else {
		goto L557
	}
L557:
	;
	if v1494&int32(4096) != 0 {
		goto L558
	} else {
		goto L559
	}
L558:
	;
	v1505 = F_HeapTupleGetUpdateXid(m, v1084)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L64
	} else {
		goto L561
	}
L559:
	;
	goto L560
L560:
	;
	v1629 = int32(0)
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+4))
	if base.Ui32(v1630) < base.Ui32(int32(3)) {
		goto L603
	} else {
		goto L604
	}
L561:
	;
	if base.Ui32(v1505) < base.Ui32(int32(3)) {
		goto L563
	} else {
		goto L564
	}
L562:
	;
	v3640 = v1626 ^ int32(1)
	goto L1
L563:
	;
	v1626 = int32(0)
	goto L562
L564:
	;
	goto L565
L565:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v1517 == v1505 {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v1626 = int32(1)
	goto L562
L567:
	;
	goto L568
L568:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v1521 <= int32(0) {
		goto L570
	} else {
		goto L571
	}
L569:
	;
	v1626 = v1618
	goto L562
L570:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v1525 == int32(0) {
		v1618 = int32(0)
		goto L569
	} else {
		goto L573
	}
L571:
	;
	goto L572
L572:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v1589 = int32(0)
	v1591 = v1521 - int32(1)
	goto L592
L573:
	;
	v1530 = v1525
	goto L574
L574:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+20))
	if v1535 == int32(4) {
		goto L576
	} else {
		goto L577
	}
L575:
	;
	v1618 = int32(0)
	goto L569
L576:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+80))
	if v1582 != 0 {
		v1530 = v1582
		goto L574
	} else {
		goto L591
	}
L577:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1530)))
	if v1538 == int32(0) {
		goto L576
	} else {
		goto L578
	}
L578:
	;
	v1541 = int32(1)
	if v1505 == v1538 {
		v1618 = v1541
		goto L569
	} else {
		goto L579
	}
L579:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+52))
	v1545 = v1543 - int32(1)
	if v1545 < int32(0) {
		goto L576
	} else {
		goto L580
	}
L580:
	;
	v1550 = int32(0)
	v1552 = v1545
	goto L581
L581:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+48))
	v1558 = int32(2)
	v1559 = base.I32_div_s(v1552-v1550, v1558)
	v1560 = v1559 + v1550
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1556+v1560<<(uint(v1558)%32))))
	if v1564 == v1505 {
		v1618 = v1541
		goto L569
	} else {
		goto L583
	}
L582:
	;
	goto L576
L583:
	;
	v1568 = F_TransactionIdPrecedes(m, v1564, v1505)
	mBase = m.M
	if v1568 != 0 {
		goto L584
	} else {
		goto L585
	}
L584:
	;
	v1569 = v1560 + int32(1)
	goto L586
L585:
	;
	v1569 = v1550
	goto L586
L586:
	;
	if v1568 != 0 {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	v1572 = v1552
	goto L589
L588:
	;
	v1572 = v1560 - int32(1)
	goto L589
L589:
	;
	if v1569 <= v1572 {
		v1550 = v1569
		v1552 = v1572
		goto L581
	} else {
		goto L590
	}
L590:
	;
	goto L582
L591:
	;
	goto L575
L592:
	;
	v1596 = int32(2)
	v1597 = base.I32_div_s(v1591-v1589, v1596)
	v1598 = v1597 + v1589
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1587+v1598<<(uint(v1596)%32))))
	v1603 = base.B2i32(v1602 == v1505)
	if v1602 == v1505 {
		v1618 = v1603
		goto L569
	} else {
		goto L594
	}
L593:
	;
	v1618 = v1603
	goto L569
L594:
	;
	v1606 = base.B2i32(base.Ui32(v1602) < base.Ui32(v1505))
	if base.Ui32(v1602) < base.Ui32(v1505) {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v1607 = v1598 + int32(1)
	goto L597
L596:
	;
	v1607 = v1589
	goto L597
L597:
	;
	if base.Ui32(v1602) < base.Ui32(v1505) {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	v1610 = v1591
	goto L600
L599:
	;
	v1610 = v1598 - int32(1)
	goto L600
L600:
	;
	if v1607 <= v1610 {
		v1589 = v1607
		v1591 = v1610
		goto L592
	} else {
		goto L601
	}
L601:
	;
	goto L593
L602:
	;
	if v1750 != 0 {
		v3640 = v1629
		goto L1
	} else {
		goto L642
	}
L603:
	;
	v1750 = int32(0)
	goto L602
L604:
	;
	goto L605
L605:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v1641 == v1630 {
		goto L606
	} else {
		goto L607
	}
L606:
	;
	v1750 = int32(1)
	goto L602
L607:
	;
	goto L608
L608:
	;
	v1645 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v1645 <= int32(0) {
		goto L610
	} else {
		goto L611
	}
L609:
	;
	v1750 = v1742
	goto L602
L610:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v1649 == int32(0) {
		v1742 = v1629
		goto L609
	} else {
		goto L613
	}
L611:
	;
	goto L612
L612:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v1713 = int32(0)
	v1715 = v1645 - int32(1)
	goto L632
L613:
	;
	v1654 = v1649
	goto L614
L614:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1654)+20))
	if v1659 == int32(4) {
		goto L616
	} else {
		goto L617
	}
L615:
	;
	v1742 = int32(0)
	goto L609
L616:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1654)+80))
	if v1706 != 0 {
		v1654 = v1706
		goto L614
	} else {
		goto L631
	}
L617:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1654)))
	if v1662 == int32(0) {
		goto L616
	} else {
		goto L618
	}
L618:
	;
	v1665 = int32(1)
	if v1630 == v1662 {
		v1742 = v1665
		goto L609
	} else {
		goto L619
	}
L619:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v1654)+52))
	v1669 = v1667 - int32(1)
	if v1669 < int32(0) {
		goto L616
	} else {
		goto L620
	}
L620:
	;
	v1674 = int32(0)
	v1676 = v1669
	goto L621
L621:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1654)+48))
	v1682 = int32(2)
	v1683 = base.I32_div_s(v1676-v1674, v1682)
	v1684 = v1683 + v1674
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1680+v1684<<(uint(v1682)%32))))
	if v1688 == v1630 {
		v1742 = v1665
		goto L609
	} else {
		goto L623
	}
L622:
	;
	goto L616
L623:
	;
	v1692 = F_TransactionIdPrecedes(m, v1688, v1630)
	mBase = m.M
	if v1692 != 0 {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	v1693 = v1684 + int32(1)
	goto L626
L625:
	;
	v1693 = v1674
	goto L626
L626:
	;
	if v1692 != 0 {
		goto L627
	} else {
		goto L628
	}
L627:
	;
	v1696 = v1676
	goto L629
L628:
	;
	v1696 = v1684 - int32(1)
	goto L629
L629:
	;
	if v1693 <= v1696 {
		v1674 = v1693
		v1676 = v1696
		goto L621
	} else {
		goto L630
	}
L630:
	;
	goto L622
L631:
	;
	goto L615
L632:
	;
	v1720 = int32(2)
	v1721 = base.I32_div_s(v1715-v1713, v1720)
	v1722 = v1721 + v1713
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1711+v1722<<(uint(v1720)%32))))
	v1727 = base.B2i32(v1726 == v1630)
	if v1726 == v1630 {
		v1742 = v1727
		goto L609
	} else {
		goto L634
	}
L633:
	;
	v1742 = v1727
	goto L609
L634:
	;
	v1730 = base.B2i32(base.Ui32(v1726) < base.Ui32(v1630))
	if base.Ui32(v1726) < base.Ui32(v1630) {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	v1731 = v1722 + int32(1)
	goto L637
L636:
	;
	v1731 = v1713
	goto L637
L637:
	;
	if base.Ui32(v1726) < base.Ui32(v1630) {
		goto L638
	} else {
		goto L639
	}
L638:
	;
	v1734 = v1715
	goto L640
L639:
	;
	v1734 = v1722 - int32(1)
	goto L640
L640:
	;
	if v1731 <= v1734 {
		v1713 = v1731
		v1715 = v1734
		goto L632
	} else {
		goto L641
	}
L641:
	;
	goto L633
L642:
	;
	v1751 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)))
	v1753 = v1751 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)) = uint16(v1753)
	goto L2
L643:
	;
	if v1756 != 0 {
		goto L644
	} else {
		goto L645
	}
L644:
	;
	v3640 = int32(0)
	goto L1
L645:
	;
	goto L646
L646:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1084)))
	v1760 = F_TransactionIdDidCommit(m, v1759)
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L64
	} else {
		goto L647
	}
L647:
	;
	if v1760 != 0 {
		goto L648
	} else {
		goto L649
	}
L648:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1084)))
	F_HeapTupleSetHintBits(m, v1084, l2, int32(256), v1763)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L64
	} else {
		goto L651
	}
L649:
	;
	goto L650
L650:
	;
	v1766 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)))
	v1768 = v1766 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)) = uint16(v1768)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L64
	} else {
		goto L652
	}
L651:
	;
	goto L401
L652:
	;
	v3640 = int32(0)
	goto L1
L653:
	;
	goto L401
L654:
	;
	if v1783&int32(1024) != 0 {
		goto L655
	} else {
		goto L656
	}
L655:
	;
	v3640 = int32(base.Ui32(v1783&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v1783&int32(4176) == int32(64))
	goto L1
L656:
	;
	goto L657
L657:
	;
	if v1783&int32(4096) != 0 {
		goto L658
	} else {
		goto L659
	}
L658:
	;
	if v1783&int32(128) != 0 {
		v3640 = v1782
		goto L1
	} else {
		goto L661
	}
L659:
	;
	goto L660
L660:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+4))
	if base.Ui32(v1935) < base.Ui32(int32(3)) {
		goto L709
	} else {
		goto L710
	}
L661:
	;
	if v1783&int32(4176) == int32(64) {
		v3640 = v1782
		goto L1
	} else {
		goto L662
	}
L662:
	;
	v1806 = F_HeapTupleGetUpdateXid(m, v1084)
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L64
	} else {
		goto L663
	}
L663:
	;
	if base.Ui32(v1806) < base.Ui32(int32(3)) {
		goto L665
	} else {
		goto L666
	}
L664:
	;
	if v1927 != 0 {
		v3640 = int32(0)
		goto L1
	} else {
		goto L704
	}
L665:
	;
	v1927 = int32(0)
	goto L664
L666:
	;
	goto L667
L667:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v1818 == v1806 {
		goto L668
	} else {
		goto L669
	}
L668:
	;
	v1927 = int32(1)
	goto L664
L669:
	;
	goto L670
L670:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v1822 <= int32(0) {
		goto L672
	} else {
		goto L673
	}
L671:
	;
	v1927 = v1919
	goto L664
L672:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v1826 == int32(0) {
		v1919 = int32(0)
		goto L671
	} else {
		goto L675
	}
L673:
	;
	goto L674
L674:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v1890 = int32(0)
	v1892 = v1822 - int32(1)
	goto L694
L675:
	;
	v1831 = v1826
	goto L676
L676:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1831)+20))
	if v1836 == int32(4) {
		goto L678
	} else {
		goto L679
	}
L677:
	;
	v1919 = int32(0)
	goto L671
L678:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1831)+80))
	if v1883 != 0 {
		v1831 = v1883
		goto L676
	} else {
		goto L693
	}
L679:
	;
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1831)))
	if v1839 == int32(0) {
		goto L678
	} else {
		goto L680
	}
L680:
	;
	v1842 = int32(1)
	if v1806 == v1839 {
		v1919 = v1842
		goto L671
	} else {
		goto L681
	}
L681:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1831)+52))
	v1846 = v1844 - int32(1)
	if v1846 < int32(0) {
		goto L678
	} else {
		goto L682
	}
L682:
	;
	v1851 = int32(0)
	v1853 = v1846
	goto L683
L683:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1831)+48))
	v1859 = int32(2)
	v1860 = base.I32_div_s(v1853-v1851, v1859)
	v1861 = v1860 + v1851
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1857+v1861<<(uint(v1859)%32))))
	if v1865 == v1806 {
		v1919 = v1842
		goto L671
	} else {
		goto L685
	}
L684:
	;
	goto L678
L685:
	;
	v1869 = F_TransactionIdPrecedes(m, v1865, v1806)
	mBase = m.M
	if v1869 != 0 {
		goto L686
	} else {
		goto L687
	}
L686:
	;
	v1870 = v1861 + int32(1)
	goto L688
L687:
	;
	v1870 = v1851
	goto L688
L688:
	;
	if v1869 != 0 {
		goto L689
	} else {
		goto L690
	}
L689:
	;
	v1873 = v1853
	goto L691
L690:
	;
	v1873 = v1861 - int32(1)
	goto L691
L691:
	;
	if v1870 <= v1873 {
		v1851 = v1870
		v1853 = v1873
		goto L683
	} else {
		goto L692
	}
L692:
	;
	goto L684
L693:
	;
	goto L677
L694:
	;
	v1897 = int32(2)
	v1898 = base.I32_div_s(v1892-v1890, v1897)
	v1899 = v1898 + v1890
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1888+v1899<<(uint(v1897)%32))))
	v1904 = base.B2i32(v1903 == v1806)
	if v1903 == v1806 {
		v1919 = v1904
		goto L671
	} else {
		goto L696
	}
L695:
	;
	v1919 = v1904
	goto L671
L696:
	;
	v1907 = base.B2i32(base.Ui32(v1903) < base.Ui32(v1806))
	if base.Ui32(v1903) < base.Ui32(v1806) {
		goto L697
	} else {
		goto L698
	}
L697:
	;
	v1908 = v1899 + int32(1)
	goto L699
L698:
	;
	v1908 = v1890
	goto L699
L699:
	;
	if base.Ui32(v1903) < base.Ui32(v1806) {
		goto L700
	} else {
		goto L701
	}
L700:
	;
	v1911 = v1892
	goto L702
L701:
	;
	v1911 = v1899 - int32(1)
	goto L702
L702:
	;
	if v1908 <= v1911 {
		v1890 = v1908
		v1892 = v1911
		goto L694
	} else {
		goto L703
	}
L703:
	;
	goto L695
L704:
	;
	v1929 = F_TransactionIdIsInProgress(m, v1806)
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L64
	} else {
		goto L705
	}
L705:
	;
	if v1929 != 0 {
		v3640 = int32(1)
		goto L1
	} else {
		goto L706
	}
L706:
	;
	v1931 = F_TransactionIdDidCommit(m, v1806)
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L64
	} else {
		goto L707
	}
L707:
	;
	v3640 = v1931 ^ int32(1)
	goto L1
L708:
	;
	if v2055 != 0 {
		goto L748
	} else {
		goto L749
	}
L709:
	;
	v2055 = int32(0)
	goto L708
L710:
	;
	goto L711
L711:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v1946 == v1935 {
		goto L712
	} else {
		goto L713
	}
L712:
	;
	v2055 = int32(1)
	goto L708
L713:
	;
	goto L714
L714:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v1950 <= int32(0) {
		goto L716
	} else {
		goto L717
	}
L715:
	;
	v2055 = v2047
	goto L708
L716:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v1954 == int32(0) {
		v2047 = int32(0)
		goto L715
	} else {
		goto L719
	}
L717:
	;
	goto L718
L718:
	;
	v2016 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v2018 = int32(0)
	v2020 = v1950 - int32(1)
	goto L738
L719:
	;
	v1959 = v1954
	goto L720
L720:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+20))
	if v1964 == int32(4) {
		goto L722
	} else {
		goto L723
	}
L721:
	;
	v2047 = int32(0)
	goto L715
L722:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+80))
	if v2011 != 0 {
		v1959 = v2011
		goto L720
	} else {
		goto L737
	}
L723:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1959)))
	if v1967 == int32(0) {
		goto L722
	} else {
		goto L724
	}
L724:
	;
	v1970 = int32(1)
	if v1935 == v1967 {
		v2047 = v1970
		goto L715
	} else {
		goto L725
	}
L725:
	;
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+52))
	v1974 = v1972 - int32(1)
	if v1974 < int32(0) {
		goto L722
	} else {
		goto L726
	}
L726:
	;
	v1979 = int32(0)
	v1981 = v1974
	goto L727
L727:
	;
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v1959)+48))
	v1987 = int32(2)
	v1988 = base.I32_div_s(v1981-v1979, v1987)
	v1989 = v1988 + v1979
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1985+v1989<<(uint(v1987)%32))))
	if v1993 == v1935 {
		v2047 = v1970
		goto L715
	} else {
		goto L729
	}
L728:
	;
	goto L722
L729:
	;
	v1997 = F_TransactionIdPrecedes(m, v1993, v1935)
	mBase = m.M
	if v1997 != 0 {
		goto L730
	} else {
		goto L731
	}
L730:
	;
	v1998 = v1989 + int32(1)
	goto L732
L731:
	;
	v1998 = v1979
	goto L732
L732:
	;
	if v1997 != 0 {
		goto L733
	} else {
		goto L734
	}
L733:
	;
	v2001 = v1981
	goto L735
L734:
	;
	v2001 = v1989 - int32(1)
	goto L735
L735:
	;
	if v1998 <= v2001 {
		v1979 = v1998
		v1981 = v2001
		goto L727
	} else {
		goto L736
	}
L736:
	;
	goto L728
L737:
	;
	goto L721
L738:
	;
	v2025 = int32(2)
	v2026 = base.I32_div_s(v2020-v2018, v2025)
	v2027 = v2026 + v2018
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v2016+v2027<<(uint(v2025)%32))))
	v2032 = base.B2i32(v2031 == v1935)
	if v2031 == v1935 {
		v2047 = v2032
		goto L715
	} else {
		goto L740
	}
L739:
	;
	v2047 = v2032
	goto L715
L740:
	;
	v2035 = base.B2i32(base.Ui32(v2031) < base.Ui32(v1935))
	if base.Ui32(v2031) < base.Ui32(v1935) {
		goto L741
	} else {
		goto L742
	}
L741:
	;
	v2036 = v2027 + int32(1)
	goto L743
L742:
	;
	v2036 = v2018
	goto L743
L743:
	;
	if base.Ui32(v2031) < base.Ui32(v1935) {
		goto L744
	} else {
		goto L745
	}
L744:
	;
	v2039 = v2020
	goto L746
L745:
	;
	v2039 = v2027 - int32(1)
	goto L746
L746:
	;
	if v2036 <= v2039 {
		v2018 = v2036
		v2020 = v2039
		goto L738
	} else {
		goto L747
	}
L747:
	;
	goto L739
L748:
	;
	v2056 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)))
	v3640 = int32(base.Ui32(v2056&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v2056&int32(4176) == int32(64))
	goto L1
L749:
	;
	goto L750
L750:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+4))
	v2067 = F_TransactionIdIsInProgress(m, v2066)
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L64
	} else {
		goto L751
	}
L751:
	;
	if v2067 != 0 {
		v3640 = v1782
		goto L1
	} else {
		goto L752
	}
L752:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+4))
	v2070 = F_TransactionIdDidCommit(m, v2069)
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L64
	} else {
		goto L753
	}
L753:
	;
	v2072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)))
	if v2070 == int32(0) {
		goto L754
	} else {
		goto L755
	}
L754:
	;
	v2076 = v2072 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)) = uint16(v2076)
	goto L2
L755:
	;
	goto L756
L756:
	;
	v2080 = int32(0)
	if base.B2i32(v2072&int32(128) == v2080)&base.B2i32(v2072&int32(4176) != int32(64)) == v2080 {
		goto L757
	} else {
		goto L758
	}
L757:
	;
	v2090 = v2072 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v1084)+20)) = uint16(v2090)
	goto L2
L758:
	;
	goto L759
L759:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+4))
	F_HeapTupleSetHintBits(m, v1084, l2, int32(1024), v2093)
	mBase = m.M
	v2095 = m.ExcPending
	if v2095 != 0 {
		goto L64
	} else {
		goto L760
	}
L760:
	;
	v3640 = int32(0)
	goto L1
L761:
	;
	v3640 = int32(1)
	goto L1
L762:
	;
	v2102 = base.I32_extend16_s(v2099)
	if v2102&int32(512) != 0 {
		goto L763
	} else {
		goto L764
	}
L763:
	;
	v3640 = int32(0)
	goto L1
L764:
	;
	goto L765
L765:
	;
	if v2102&int32(16384) != 0 {
		goto L767
	} else {
		goto L768
	}
L766:
	;
	v2385 = v2383 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2098)+20)) = uint16(v2385)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L64
	} else {
		goto L870
	}
L767:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v2098)+8))
	if base.Ui32(v2108) < base.Ui32(int32(3)) {
		goto L771
	} else {
		goto L772
	}
L768:
	;
	goto L769
L769:
	;
	if v2102 < int32(0) {
		goto L818
	} else {
		goto L819
	}
L770:
	;
	if v2228 != 0 {
		goto L810
	} else {
		goto L811
	}
L771:
	;
	v2228 = int32(0)
	goto L770
L772:
	;
	goto L773
L773:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v2119 == v2108 {
		goto L774
	} else {
		goto L775
	}
L774:
	;
	v2228 = int32(1)
	goto L770
L775:
	;
	goto L776
L776:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v2123 <= int32(0) {
		goto L778
	} else {
		goto L779
	}
L777:
	;
	v2228 = v2220
	goto L770
L778:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v2127 == int32(0) {
		v2220 = int32(0)
		goto L777
	} else {
		goto L781
	}
L779:
	;
	goto L780
L780:
	;
	v2189 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v2191 = int32(0)
	v2193 = v2123 - int32(1)
	goto L800
L781:
	;
	v2132 = v2127
	goto L782
L782:
	;
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+20))
	if v2137 == int32(4) {
		goto L784
	} else {
		goto L785
	}
L783:
	;
	v2220 = int32(0)
	goto L777
L784:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+80))
	if v2184 != 0 {
		v2132 = v2184
		goto L782
	} else {
		goto L799
	}
L785:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2132)))
	if v2140 == int32(0) {
		goto L784
	} else {
		goto L786
	}
L786:
	;
	v2143 = int32(1)
	if v2108 == v2140 {
		v2220 = v2143
		goto L777
	} else {
		goto L787
	}
L787:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+52))
	v2147 = v2145 - int32(1)
	if v2147 < int32(0) {
		goto L784
	} else {
		goto L788
	}
L788:
	;
	v2152 = int32(0)
	v2154 = v2147
	goto L789
L789:
	;
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+48))
	v2160 = int32(2)
	v2161 = base.I32_div_s(v2154-v2152, v2160)
	v2162 = v2161 + v2152
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2158+v2162<<(uint(v2160)%32))))
	if v2166 == v2108 {
		v2220 = v2143
		goto L777
	} else {
		goto L791
	}
L790:
	;
	goto L784
L791:
	;
	v2170 = F_TransactionIdPrecedes(m, v2166, v2108)
	mBase = m.M
	if v2170 != 0 {
		goto L792
	} else {
		goto L793
	}
L792:
	;
	v2171 = v2162 + int32(1)
	goto L794
L793:
	;
	v2171 = v2152
	goto L794
L794:
	;
	if v2170 != 0 {
		goto L795
	} else {
		goto L796
	}
L795:
	;
	v2174 = v2154
	goto L797
L796:
	;
	v2174 = v2162 - int32(1)
	goto L797
L797:
	;
	if v2171 <= v2174 {
		v2152 = v2171
		v2154 = v2174
		goto L789
	} else {
		goto L798
	}
L798:
	;
	goto L790
L799:
	;
	goto L783
L800:
	;
	v2198 = int32(2)
	v2199 = base.I32_div_s(v2193-v2191, v2198)
	v2200 = v2199 + v2191
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v2189+v2200<<(uint(v2198)%32))))
	v2205 = base.B2i32(v2204 == v2108)
	if v2204 == v2108 {
		v2220 = v2205
		goto L777
	} else {
		goto L802
	}
L801:
	;
	v2220 = v2205
	goto L777
L802:
	;
	v2208 = base.B2i32(base.Ui32(v2204) < base.Ui32(v2108))
	if base.Ui32(v2204) < base.Ui32(v2108) {
		goto L803
	} else {
		goto L804
	}
L803:
	;
	v2209 = v2200 + int32(1)
	goto L805
L804:
	;
	v2209 = v2191
	goto L805
L805:
	;
	if base.Ui32(v2204) < base.Ui32(v2108) {
		goto L806
	} else {
		goto L807
	}
L806:
	;
	v2212 = v2193
	goto L808
L807:
	;
	v2212 = v2200 - int32(1)
	goto L808
L808:
	;
	if v2209 <= v2212 {
		v2191 = v2209
		v2193 = v2212
		goto L800
	} else {
		goto L809
	}
L809:
	;
	goto L801
L810:
	;
	v3640 = int32(0)
	goto L1
L811:
	;
	goto L812
L812:
	;
	v2230 = F_TransactionIdIsInProgress(m, v2108)
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		goto L64
	} else {
		goto L813
	}
L813:
	;
	if v2230 != 0 {
		goto L761
	} else {
		goto L814
	}
L814:
	;
	v2232 = F_TransactionIdDidCommit(m, v2108)
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L64
	} else {
		goto L815
	}
L815:
	;
	v2234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2098)+20)))
	if v2232 == int32(0) {
		v2383 = v2234
		goto L766
	} else {
		goto L816
	}
L816:
	;
	v2238 = v2234 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2098)+20)) = uint16(v2238)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L64
	} else {
		goto L817
	}
L817:
	;
	v3640 = int32(0)
	goto L1
L818:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2098)+8))
	if base.Ui32(v2246) < base.Ui32(int32(3)) {
		goto L822
	} else {
		goto L823
	}
L819:
	;
	goto L820
L820:
	;
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2098)))
	if v2381 != 0 {
		goto L761
	} else {
		goto L869
	}
L821:
	;
	if v2366 != 0 {
		goto L761
	} else {
		goto L861
	}
L822:
	;
	v2366 = int32(0)
	goto L821
L823:
	;
	goto L824
L824:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v2257 == v2246 {
		goto L825
	} else {
		goto L826
	}
L825:
	;
	v2366 = int32(1)
	goto L821
L826:
	;
	goto L827
L827:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v2261 <= int32(0) {
		goto L829
	} else {
		goto L830
	}
L828:
	;
	v2366 = v2358
	goto L821
L829:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v2265 == int32(0) {
		v2358 = int32(0)
		goto L828
	} else {
		goto L832
	}
L830:
	;
	goto L831
L831:
	;
	v2327 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v2329 = int32(0)
	v2331 = v2261 - int32(1)
	goto L851
L832:
	;
	v2270 = v2265
	goto L833
L833:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v2270)+20))
	if v2275 == int32(4) {
		goto L835
	} else {
		goto L836
	}
L834:
	;
	v2358 = int32(0)
	goto L828
L835:
	;
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v2270)+80))
	if v2322 != 0 {
		v2270 = v2322
		goto L833
	} else {
		goto L850
	}
L836:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2270)))
	if v2278 == int32(0) {
		goto L835
	} else {
		goto L837
	}
L837:
	;
	v2281 = int32(1)
	if v2246 == v2278 {
		v2358 = v2281
		goto L828
	} else {
		goto L838
	}
L838:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v2270)+52))
	v2285 = v2283 - int32(1)
	if v2285 < int32(0) {
		goto L835
	} else {
		goto L839
	}
L839:
	;
	v2290 = int32(0)
	v2292 = v2285
	goto L840
L840:
	;
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v2270)+48))
	v2298 = int32(2)
	v2299 = base.I32_div_s(v2292-v2290, v2298)
	v2300 = v2299 + v2290
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v2296+v2300<<(uint(v2298)%32))))
	if v2304 == v2246 {
		v2358 = v2281
		goto L828
	} else {
		goto L842
	}
L841:
	;
	goto L835
L842:
	;
	v2308 = F_TransactionIdPrecedes(m, v2304, v2246)
	mBase = m.M
	if v2308 != 0 {
		goto L843
	} else {
		goto L844
	}
L843:
	;
	v2309 = v2300 + int32(1)
	goto L845
L844:
	;
	v2309 = v2290
	goto L845
L845:
	;
	if v2308 != 0 {
		goto L846
	} else {
		goto L847
	}
L846:
	;
	v2312 = v2292
	goto L848
L847:
	;
	v2312 = v2300 - int32(1)
	goto L848
L848:
	;
	if v2309 <= v2312 {
		v2290 = v2309
		v2292 = v2312
		goto L840
	} else {
		goto L849
	}
L849:
	;
	goto L841
L850:
	;
	goto L834
L851:
	;
	v2336 = int32(2)
	v2337 = base.I32_div_s(v2331-v2329, v2336)
	v2338 = v2337 + v2329
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v2327+v2338<<(uint(v2336)%32))))
	v2343 = base.B2i32(v2342 == v2246)
	if v2342 == v2246 {
		v2358 = v2343
		goto L828
	} else {
		goto L853
	}
L852:
	;
	v2358 = v2343
	goto L828
L853:
	;
	v2346 = base.B2i32(base.Ui32(v2342) < base.Ui32(v2246))
	if base.Ui32(v2342) < base.Ui32(v2246) {
		goto L854
	} else {
		goto L855
	}
L854:
	;
	v2347 = v2338 + int32(1)
	goto L856
L855:
	;
	v2347 = v2329
	goto L856
L856:
	;
	if base.Ui32(v2342) < base.Ui32(v2246) {
		goto L857
	} else {
		goto L858
	}
L857:
	;
	v2350 = v2331
	goto L859
L858:
	;
	v2350 = v2338 - int32(1)
	goto L859
L859:
	;
	if v2347 <= v2350 {
		v2329 = v2347
		v2331 = v2350
		goto L851
	} else {
		goto L860
	}
L860:
	;
	goto L852
L861:
	;
	v2367 = F_TransactionIdIsInProgress(m, v2246)
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L64
	} else {
		goto L862
	}
L862:
	;
	if v2367 != 0 {
		goto L863
	} else {
		goto L864
	}
L863:
	;
	v3640 = int32(0)
	goto L1
L864:
	;
	goto L865
L865:
	;
	v2370 = F_TransactionIdDidCommit(m, v2246)
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L64
	} else {
		goto L866
	}
L866:
	;
	v2372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2098)+20)))
	if v2370 != 0 {
		v2383 = v2372
		goto L766
	} else {
		goto L867
	}
L867:
	;
	v2374 = v2372 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2098)+20)) = uint16(v2374)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2378 = m.ExcPending
	if v2378 != 0 {
		goto L64
	} else {
		goto L868
	}
L868:
	;
	v3640 = int32(0)
	goto L1
L869:
	;
	v3640 = int32(0)
	goto L1
L870:
	;
	goto L761
L871:
	;
	v3100 = int32(1)
	v3101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	if v3101&int32(2048) != 0 {
		v3640 = v3100
		goto L1
	} else {
		goto L1121
	}
L872:
	;
	v2401 = base.I32_extend16_s(v2398)
	if v2401&int32(512) != 0 {
		v3640 = v4
		goto L1
	} else {
		goto L873
	}
L873:
	;
	if v2401&int32(16384) != 0 {
		goto L875
	} else {
		goto L876
	}
L874:
	;
	v3093 = v2531 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)) = uint16(v3093)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v3097 = m.ExcPending
	if v3097 != 0 {
		goto L64
	} else {
		goto L1120
	}
L875:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+8))
	if base.Ui32(v2406) < base.Ui32(int32(3)) {
		goto L879
	} else {
		goto L880
	}
L876:
	;
	goto L877
L877:
	;
	if v2401 < int32(0) {
		goto L924
	} else {
		goto L925
	}
L878:
	;
	if v2526 != 0 {
		v3640 = v4
		goto L1
	} else {
		goto L918
	}
L879:
	;
	v2526 = int32(0)
	goto L878
L880:
	;
	goto L881
L881:
	;
	v2417 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v2417 == v2406 {
		goto L882
	} else {
		goto L883
	}
L882:
	;
	v2526 = int32(1)
	goto L878
L883:
	;
	goto L884
L884:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v2421 <= int32(0) {
		goto L886
	} else {
		goto L887
	}
L885:
	;
	v2526 = v2518
	goto L878
L886:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v2425 == int32(0) {
		v2518 = int32(0)
		goto L885
	} else {
		goto L889
	}
L887:
	;
	goto L888
L888:
	;
	v2487 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v2489 = int32(0)
	v2491 = v2421 - int32(1)
	goto L908
L889:
	;
	v2430 = v2425
	goto L890
L890:
	;
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2430)+20))
	if v2435 == int32(4) {
		goto L892
	} else {
		goto L893
	}
L891:
	;
	v2518 = int32(0)
	goto L885
L892:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v2430)+80))
	if v2482 != 0 {
		v2430 = v2482
		goto L890
	} else {
		goto L907
	}
L893:
	;
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(v2430)))
	if v2438 == int32(0) {
		goto L892
	} else {
		goto L894
	}
L894:
	;
	v2441 = int32(1)
	if v2406 == v2438 {
		v2518 = v2441
		goto L885
	} else {
		goto L895
	}
L895:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v2430)+52))
	v2445 = v2443 - int32(1)
	if v2445 < int32(0) {
		goto L892
	} else {
		goto L896
	}
L896:
	;
	v2450 = int32(0)
	v2452 = v2445
	goto L897
L897:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2430)+48))
	v2458 = int32(2)
	v2459 = base.I32_div_s(v2452-v2450, v2458)
	v2460 = v2459 + v2450
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v2456+v2460<<(uint(v2458)%32))))
	if v2464 == v2406 {
		v2518 = v2441
		goto L885
	} else {
		goto L899
	}
L898:
	;
	goto L892
L899:
	;
	v2468 = F_TransactionIdPrecedes(m, v2464, v2406)
	mBase = m.M
	if v2468 != 0 {
		goto L900
	} else {
		goto L901
	}
L900:
	;
	v2469 = v2460 + int32(1)
	goto L902
L901:
	;
	v2469 = v2450
	goto L902
L902:
	;
	if v2468 != 0 {
		goto L903
	} else {
		goto L904
	}
L903:
	;
	v2472 = v2452
	goto L905
L904:
	;
	v2472 = v2460 - int32(1)
	goto L905
L905:
	;
	if v2469 <= v2472 {
		v2450 = v2469
		v2452 = v2472
		goto L897
	} else {
		goto L906
	}
L906:
	;
	goto L898
L907:
	;
	goto L891
L908:
	;
	v2496 = int32(2)
	v2497 = base.I32_div_s(v2491-v2489, v2496)
	v2498 = v2497 + v2489
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v2487+v2498<<(uint(v2496)%32))))
	v2503 = base.B2i32(v2502 == v2406)
	if v2502 == v2406 {
		v2518 = v2503
		goto L885
	} else {
		goto L910
	}
L909:
	;
	v2518 = v2503
	goto L885
L910:
	;
	v2506 = base.B2i32(base.Ui32(v2502) < base.Ui32(v2406))
	if base.Ui32(v2502) < base.Ui32(v2406) {
		goto L911
	} else {
		goto L912
	}
L911:
	;
	v2507 = v2498 + int32(1)
	goto L913
L912:
	;
	v2507 = v2489
	goto L913
L913:
	;
	if base.Ui32(v2502) < base.Ui32(v2406) {
		goto L914
	} else {
		goto L915
	}
L914:
	;
	v2510 = v2491
	goto L916
L915:
	;
	v2510 = v2498 - int32(1)
	goto L916
L916:
	;
	if v2507 <= v2510 {
		v2489 = v2507
		v2491 = v2510
		goto L908
	} else {
		goto L917
	}
L917:
	;
	goto L909
L918:
	;
	v2527 = F_TransactionIdIsInProgress(m, v2406)
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L64
	} else {
		goto L919
	}
L919:
	;
	if v2527 != 0 {
		goto L871
	} else {
		goto L920
	}
L920:
	;
	v2529 = F_TransactionIdDidCommit(m, v2406)
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L64
	} else {
		goto L921
	}
L921:
	;
	v2531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	if v2529 == int32(0) {
		goto L874
	} else {
		goto L922
	}
L922:
	;
	v2535 = v2531 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)) = uint16(v2535)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L64
	} else {
		goto L923
	}
L923:
	;
	v3640 = v4
	goto L1
L924:
	;
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+8))
	if base.Ui32(v2542) < base.Ui32(int32(3)) {
		goto L928
	} else {
		goto L929
	}
L925:
	;
	goto L926
L926:
	;
	v2680 = *(*int32)(unsafe.Add(mBase, uint32(v2393)))
	if base.Ui32(v2680) < base.Ui32(int32(3)) {
		goto L977
	} else {
		goto L978
	}
L927:
	;
	if v2662 != 0 {
		goto L871
	} else {
		goto L967
	}
L928:
	;
	v2662 = int32(0)
	goto L927
L929:
	;
	goto L930
L930:
	;
	v2553 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v2553 == v2542 {
		goto L931
	} else {
		goto L932
	}
L931:
	;
	v2662 = int32(1)
	goto L927
L932:
	;
	goto L933
L933:
	;
	v2557 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v2557 <= int32(0) {
		goto L935
	} else {
		goto L936
	}
L934:
	;
	v2662 = v2654
	goto L927
L935:
	;
	v2561 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v2561 == int32(0) {
		v2654 = int32(0)
		goto L934
	} else {
		goto L938
	}
L936:
	;
	goto L937
L937:
	;
	v2623 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v2625 = int32(0)
	v2627 = v2557 - int32(1)
	goto L957
L938:
	;
	v2566 = v2561
	goto L939
L939:
	;
	v2571 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+20))
	if v2571 == int32(4) {
		goto L941
	} else {
		goto L942
	}
L940:
	;
	v2654 = int32(0)
	goto L934
L941:
	;
	v2618 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+80))
	if v2618 != 0 {
		v2566 = v2618
		goto L939
	} else {
		goto L956
	}
L942:
	;
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(v2566)))
	if v2574 == int32(0) {
		goto L941
	} else {
		goto L943
	}
L943:
	;
	v2577 = int32(1)
	if v2542 == v2574 {
		v2654 = v2577
		goto L934
	} else {
		goto L944
	}
L944:
	;
	v2579 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+52))
	v2581 = v2579 - int32(1)
	if v2581 < int32(0) {
		goto L941
	} else {
		goto L945
	}
L945:
	;
	v2586 = int32(0)
	v2588 = v2581
	goto L946
L946:
	;
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(v2566)+48))
	v2594 = int32(2)
	v2595 = base.I32_div_s(v2588-v2586, v2594)
	v2596 = v2595 + v2586
	v2600 = *(*int32)(unsafe.Add(mBase, uint32(v2592+v2596<<(uint(v2594)%32))))
	if v2600 == v2542 {
		v2654 = v2577
		goto L934
	} else {
		goto L948
	}
L947:
	;
	goto L941
L948:
	;
	v2604 = F_TransactionIdPrecedes(m, v2600, v2542)
	mBase = m.M
	if v2604 != 0 {
		goto L949
	} else {
		goto L950
	}
L949:
	;
	v2605 = v2596 + int32(1)
	goto L951
L950:
	;
	v2605 = v2586
	goto L951
L951:
	;
	if v2604 != 0 {
		goto L952
	} else {
		goto L953
	}
L952:
	;
	v2608 = v2588
	goto L954
L953:
	;
	v2608 = v2596 - int32(1)
	goto L954
L954:
	;
	if v2605 <= v2608 {
		v2586 = v2605
		v2588 = v2608
		goto L946
	} else {
		goto L955
	}
L955:
	;
	goto L947
L956:
	;
	goto L940
L957:
	;
	v2632 = int32(2)
	v2633 = base.I32_div_s(v2627-v2625, v2632)
	v2634 = v2633 + v2625
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(v2623+v2634<<(uint(v2632)%32))))
	v2639 = base.B2i32(v2638 == v2542)
	if v2638 == v2542 {
		v2654 = v2639
		goto L934
	} else {
		goto L959
	}
L958:
	;
	v2654 = v2639
	goto L934
L959:
	;
	v2642 = base.B2i32(base.Ui32(v2638) < base.Ui32(v2542))
	if base.Ui32(v2638) < base.Ui32(v2542) {
		goto L960
	} else {
		goto L961
	}
L960:
	;
	v2643 = v2634 + int32(1)
	goto L962
L961:
	;
	v2643 = v2625
	goto L962
L962:
	;
	if base.Ui32(v2638) < base.Ui32(v2542) {
		goto L963
	} else {
		goto L964
	}
L963:
	;
	v2646 = v2627
	goto L965
L964:
	;
	v2646 = v2634 - int32(1)
	goto L965
L965:
	;
	if v2643 <= v2646 {
		v2625 = v2643
		v2627 = v2646
		goto L957
	} else {
		goto L966
	}
L966:
	;
	goto L958
L967:
	;
	v2663 = F_TransactionIdIsInProgress(m, v2542)
	mBase = m.M
	v2664 = m.ExcPending
	if v2664 != 0 {
		goto L64
	} else {
		goto L968
	}
L968:
	;
	if v2663 != 0 {
		v3640 = v4
		goto L1
	} else {
		goto L969
	}
L969:
	;
	v2665 = F_TransactionIdDidCommit(m, v2542)
	mBase = m.M
	v2666 = m.ExcPending
	if v2666 != 0 {
		goto L64
	} else {
		goto L970
	}
L970:
	;
	v2667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	if v2665 != 0 {
		goto L971
	} else {
		goto L972
	}
L971:
	;
	v2669 = v2667 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)) = uint16(v2669)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		goto L64
	} else {
		goto L974
	}
L972:
	;
	goto L973
L973:
	;
	v2675 = v2667 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)) = uint16(v2675)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L64
	} else {
		goto L975
	}
L974:
	;
	goto L871
L975:
	;
	v3640 = v4
	goto L1
L976:
	;
	if v2800 != 0 {
		goto L1016
	} else {
		goto L1017
	}
L977:
	;
	v2800 = int32(0)
	goto L976
L978:
	;
	goto L979
L979:
	;
	v2691 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v2691 == v2680 {
		goto L980
	} else {
		goto L981
	}
L980:
	;
	v2800 = int32(1)
	goto L976
L981:
	;
	goto L982
L982:
	;
	v2695 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v2695 <= int32(0) {
		goto L984
	} else {
		goto L985
	}
L983:
	;
	v2800 = v2792
	goto L976
L984:
	;
	v2699 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v2699 == int32(0) {
		v2792 = int32(0)
		goto L983
	} else {
		goto L987
	}
L985:
	;
	goto L986
L986:
	;
	v2761 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v2763 = int32(0)
	v2765 = v2695 - int32(1)
	goto L1006
L987:
	;
	v2704 = v2699
	goto L988
L988:
	;
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v2704)+20))
	if v2709 == int32(4) {
		goto L990
	} else {
		goto L991
	}
L989:
	;
	v2792 = int32(0)
	goto L983
L990:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v2704)+80))
	if v2756 != 0 {
		v2704 = v2756
		goto L988
	} else {
		goto L1005
	}
L991:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(v2704)))
	if v2712 == int32(0) {
		goto L990
	} else {
		goto L992
	}
L992:
	;
	v2715 = int32(1)
	if v2680 == v2712 {
		v2792 = v2715
		goto L983
	} else {
		goto L993
	}
L993:
	;
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2704)+52))
	v2719 = v2717 - int32(1)
	if v2719 < int32(0) {
		goto L990
	} else {
		goto L994
	}
L994:
	;
	v2724 = int32(0)
	v2726 = v2719
	goto L995
L995:
	;
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2704)+48))
	v2732 = int32(2)
	v2733 = base.I32_div_s(v2726-v2724, v2732)
	v2734 = v2733 + v2724
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v2730+v2734<<(uint(v2732)%32))))
	if v2738 == v2680 {
		v2792 = v2715
		goto L983
	} else {
		goto L997
	}
L996:
	;
	goto L990
L997:
	;
	v2742 = F_TransactionIdPrecedes(m, v2738, v2680)
	mBase = m.M
	if v2742 != 0 {
		goto L998
	} else {
		goto L999
	}
L998:
	;
	v2743 = v2734 + int32(1)
	goto L1000
L999:
	;
	v2743 = v2724
	goto L1000
L1000:
	;
	if v2742 != 0 {
		goto L1001
	} else {
		goto L1002
	}
L1001:
	;
	v2746 = v2726
	goto L1003
L1002:
	;
	v2746 = v2734 - int32(1)
	goto L1003
L1003:
	;
	if v2743 <= v2746 {
		v2724 = v2743
		v2726 = v2746
		goto L995
	} else {
		goto L1004
	}
L1004:
	;
	goto L996
L1005:
	;
	goto L989
L1006:
	;
	v2770 = int32(2)
	v2771 = base.I32_div_s(v2765-v2763, v2770)
	v2772 = v2771 + v2763
	v2776 = *(*int32)(unsafe.Add(mBase, uint32(v2761+v2772<<(uint(v2770)%32))))
	v2777 = base.B2i32(v2776 == v2680)
	if v2776 == v2680 {
		v2792 = v2777
		goto L983
	} else {
		goto L1008
	}
L1007:
	;
	v2792 = v2777
	goto L983
L1008:
	;
	v2780 = base.B2i32(base.Ui32(v2776) < base.Ui32(v2680))
	if base.Ui32(v2776) < base.Ui32(v2680) {
		goto L1009
	} else {
		goto L1010
	}
L1009:
	;
	v2781 = v2772 + int32(1)
	goto L1011
L1010:
	;
	v2781 = v2763
	goto L1011
L1011:
	;
	if base.Ui32(v2776) < base.Ui32(v2680) {
		goto L1012
	} else {
		goto L1013
	}
L1012:
	;
	v2784 = v2765
	goto L1014
L1013:
	;
	v2784 = v2772 - int32(1)
	goto L1014
L1014:
	;
	if v2781 <= v2784 {
		v2763 = v2781
		v2765 = v2784
		goto L1006
	} else {
		goto L1015
	}
L1015:
	;
	goto L1007
L1016:
	;
	v2801 = int32(1)
	v2802 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	if v2802&int32(2048) != 0 {
		v3640 = v2801
		goto L1
	} else {
		goto L1019
	}
L1017:
	;
	goto L1018
L1018:
	;
	v3063 = *(*int32)(unsafe.Add(mBase, uint32(v2393)))
	v3064 = F_TransactionIdIsInProgress(m, v3063)
	mBase = m.M
	v3065 = m.ExcPending
	if v3065 != 0 {
		goto L64
	} else {
		goto L1107
	}
L1019:
	;
	if v2802&int32(128) != 0 {
		v3640 = v2801
		goto L1
	} else {
		goto L1020
	}
L1020:
	;
	if v2802&int32(4176) == int32(64) {
		v3640 = v2801
		goto L1
	} else {
		goto L1021
	}
L1021:
	;
	if v2802&int32(4096) != 0 {
		goto L1022
	} else {
		goto L1023
	}
L1022:
	;
	v2813 = F_HeapTupleGetUpdateXid(m, v2393)
	mBase = m.M
	v2814 = m.ExcPending
	if v2814 != 0 {
		goto L64
	} else {
		goto L1025
	}
L1023:
	;
	goto L1024
L1024:
	;
	v2937 = int32(0)
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	if base.Ui32(v2938) < base.Ui32(int32(3)) {
		goto L1067
	} else {
		goto L1068
	}
L1025:
	;
	if base.Ui32(v2813) < base.Ui32(int32(3)) {
		goto L1027
	} else {
		goto L1028
	}
L1026:
	;
	v3640 = v2934 ^ int32(1)
	goto L1
L1027:
	;
	v2934 = int32(0)
	goto L1026
L1028:
	;
	goto L1029
L1029:
	;
	v2825 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v2825 == v2813 {
		goto L1030
	} else {
		goto L1031
	}
L1030:
	;
	v2934 = int32(1)
	goto L1026
L1031:
	;
	goto L1032
L1032:
	;
	v2829 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v2829 <= int32(0) {
		goto L1034
	} else {
		goto L1035
	}
L1033:
	;
	v2934 = v2926
	goto L1026
L1034:
	;
	v2833 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v2833 == int32(0) {
		v2926 = int32(0)
		goto L1033
	} else {
		goto L1037
	}
L1035:
	;
	goto L1036
L1036:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v2897 = int32(0)
	v2899 = v2829 - int32(1)
	goto L1056
L1037:
	;
	v2838 = v2833
	goto L1038
L1038:
	;
	v2843 = *(*int32)(unsafe.Add(mBase, uint32(v2838)+20))
	if v2843 == int32(4) {
		goto L1040
	} else {
		goto L1041
	}
L1039:
	;
	v2926 = int32(0)
	goto L1033
L1040:
	;
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(v2838)+80))
	if v2890 != 0 {
		v2838 = v2890
		goto L1038
	} else {
		goto L1055
	}
L1041:
	;
	v2846 = *(*int32)(unsafe.Add(mBase, uint32(v2838)))
	if v2846 == int32(0) {
		goto L1040
	} else {
		goto L1042
	}
L1042:
	;
	v2849 = int32(1)
	if v2813 == v2846 {
		v2926 = v2849
		goto L1033
	} else {
		goto L1043
	}
L1043:
	;
	v2851 = *(*int32)(unsafe.Add(mBase, uint32(v2838)+52))
	v2853 = v2851 - int32(1)
	if v2853 < int32(0) {
		goto L1040
	} else {
		goto L1044
	}
L1044:
	;
	v2858 = int32(0)
	v2860 = v2853
	goto L1045
L1045:
	;
	v2864 = *(*int32)(unsafe.Add(mBase, uint32(v2838)+48))
	v2866 = int32(2)
	v2867 = base.I32_div_s(v2860-v2858, v2866)
	v2868 = v2867 + v2858
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(v2864+v2868<<(uint(v2866)%32))))
	if v2872 == v2813 {
		v2926 = v2849
		goto L1033
	} else {
		goto L1047
	}
L1046:
	;
	goto L1040
L1047:
	;
	v2876 = F_TransactionIdPrecedes(m, v2872, v2813)
	mBase = m.M
	if v2876 != 0 {
		goto L1048
	} else {
		goto L1049
	}
L1048:
	;
	v2877 = v2868 + int32(1)
	goto L1050
L1049:
	;
	v2877 = v2858
	goto L1050
L1050:
	;
	if v2876 != 0 {
		goto L1051
	} else {
		goto L1052
	}
L1051:
	;
	v2880 = v2860
	goto L1053
L1052:
	;
	v2880 = v2868 - int32(1)
	goto L1053
L1053:
	;
	if v2877 <= v2880 {
		v2858 = v2877
		v2860 = v2880
		goto L1045
	} else {
		goto L1054
	}
L1054:
	;
	goto L1046
L1055:
	;
	goto L1039
L1056:
	;
	v2904 = int32(2)
	v2905 = base.I32_div_s(v2899-v2897, v2904)
	v2906 = v2905 + v2897
	v2910 = *(*int32)(unsafe.Add(mBase, uint32(v2895+v2906<<(uint(v2904)%32))))
	v2911 = base.B2i32(v2910 == v2813)
	if v2910 == v2813 {
		v2926 = v2911
		goto L1033
	} else {
		goto L1058
	}
L1057:
	;
	v2926 = v2911
	goto L1033
L1058:
	;
	v2914 = base.B2i32(base.Ui32(v2910) < base.Ui32(v2813))
	if base.Ui32(v2910) < base.Ui32(v2813) {
		goto L1059
	} else {
		goto L1060
	}
L1059:
	;
	v2915 = v2906 + int32(1)
	goto L1061
L1060:
	;
	v2915 = v2897
	goto L1061
L1061:
	;
	if base.Ui32(v2910) < base.Ui32(v2813) {
		goto L1062
	} else {
		goto L1063
	}
L1062:
	;
	v2918 = v2899
	goto L1064
L1063:
	;
	v2918 = v2906 - int32(1)
	goto L1064
L1064:
	;
	if v2915 <= v2918 {
		v2897 = v2915
		v2899 = v2918
		goto L1056
	} else {
		goto L1065
	}
L1065:
	;
	goto L1057
L1066:
	;
	if v3058 != 0 {
		v3640 = v2937
		goto L1
	} else {
		goto L1106
	}
L1067:
	;
	v3058 = int32(0)
	goto L1066
L1068:
	;
	goto L1069
L1069:
	;
	v2949 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v2949 == v2938 {
		goto L1070
	} else {
		goto L1071
	}
L1070:
	;
	v3058 = int32(1)
	goto L1066
L1071:
	;
	goto L1072
L1072:
	;
	v2953 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v2953 <= int32(0) {
		goto L1074
	} else {
		goto L1075
	}
L1073:
	;
	v3058 = v3050
	goto L1066
L1074:
	;
	v2957 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v2957 == int32(0) {
		v3050 = v2937
		goto L1073
	} else {
		goto L1077
	}
L1075:
	;
	goto L1076
L1076:
	;
	v3019 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v3021 = int32(0)
	v3023 = v2953 - int32(1)
	goto L1096
L1077:
	;
	v2962 = v2957
	goto L1078
L1078:
	;
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v2962)+20))
	if v2967 == int32(4) {
		goto L1080
	} else {
		goto L1081
	}
L1079:
	;
	v3050 = int32(0)
	goto L1073
L1080:
	;
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(v2962)+80))
	if v3014 != 0 {
		v2962 = v3014
		goto L1078
	} else {
		goto L1095
	}
L1081:
	;
	v2970 = *(*int32)(unsafe.Add(mBase, uint32(v2962)))
	if v2970 == int32(0) {
		goto L1080
	} else {
		goto L1082
	}
L1082:
	;
	v2973 = int32(1)
	if v2938 == v2970 {
		v3050 = v2973
		goto L1073
	} else {
		goto L1083
	}
L1083:
	;
	v2975 = *(*int32)(unsafe.Add(mBase, uint32(v2962)+52))
	v2977 = v2975 - int32(1)
	if v2977 < int32(0) {
		goto L1080
	} else {
		goto L1084
	}
L1084:
	;
	v2982 = int32(0)
	v2984 = v2977
	goto L1085
L1085:
	;
	v2988 = *(*int32)(unsafe.Add(mBase, uint32(v2962)+48))
	v2990 = int32(2)
	v2991 = base.I32_div_s(v2984-v2982, v2990)
	v2992 = v2991 + v2982
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(v2988+v2992<<(uint(v2990)%32))))
	if v2996 == v2938 {
		v3050 = v2973
		goto L1073
	} else {
		goto L1087
	}
L1086:
	;
	goto L1080
L1087:
	;
	v3000 = F_TransactionIdPrecedes(m, v2996, v2938)
	mBase = m.M
	if v3000 != 0 {
		goto L1088
	} else {
		goto L1089
	}
L1088:
	;
	v3001 = v2992 + int32(1)
	goto L1090
L1089:
	;
	v3001 = v2982
	goto L1090
L1090:
	;
	if v3000 != 0 {
		goto L1091
	} else {
		goto L1092
	}
L1091:
	;
	v3004 = v2984
	goto L1093
L1092:
	;
	v3004 = v2992 - int32(1)
	goto L1093
L1093:
	;
	if v3001 <= v3004 {
		v2982 = v3001
		v2984 = v3004
		goto L1085
	} else {
		goto L1094
	}
L1094:
	;
	goto L1086
L1095:
	;
	goto L1079
L1096:
	;
	v3028 = int32(2)
	v3029 = base.I32_div_s(v3023-v3021, v3028)
	v3030 = v3029 + v3021
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(v3019+v3030<<(uint(v3028)%32))))
	v3035 = base.B2i32(v3034 == v2938)
	if v3034 == v2938 {
		v3050 = v3035
		goto L1073
	} else {
		goto L1098
	}
L1097:
	;
	v3050 = v3035
	goto L1073
L1098:
	;
	v3038 = base.B2i32(base.Ui32(v3034) < base.Ui32(v2938))
	if base.Ui32(v3034) < base.Ui32(v2938) {
		goto L1099
	} else {
		goto L1100
	}
L1099:
	;
	v3039 = v3030 + int32(1)
	goto L1101
L1100:
	;
	v3039 = v3021
	goto L1101
L1101:
	;
	if base.Ui32(v3034) < base.Ui32(v2938) {
		goto L1102
	} else {
		goto L1103
	}
L1102:
	;
	v3042 = v3023
	goto L1104
L1103:
	;
	v3042 = v3030 - int32(1)
	goto L1104
L1104:
	;
	if v3039 <= v3042 {
		v3021 = v3039
		v3023 = v3042
		goto L1096
	} else {
		goto L1105
	}
L1105:
	;
	goto L1097
L1106:
	;
	v3059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	v3061 = v3059 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)) = uint16(v3061)
	goto L2
L1107:
	;
	if v3064 != 0 {
		goto L1108
	} else {
		goto L1109
	}
L1108:
	;
	v3066 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+16)))
	if v3066 == int32(65534) {
		goto L1111
	} else {
		goto L1112
	}
L1109:
	;
	goto L1110
L1110:
	;
	v3078 = *(*int32)(unsafe.Add(mBase, uint32(v2393)))
	v3079 = F_TransactionIdDidCommit(m, v3078)
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L64
	} else {
		goto L1114
	}
L1111:
	;
	v3069 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+14)))
	v3070 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v3069 | v3070<<(uint(int32(16))%32)
	goto L1113
L1112:
	;
	goto L1113
L1113:
	;
	v3075 = *(*int32)(unsafe.Add(mBase, uint32(v2393)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v3075
	v3640 = int32(1)
	goto L1
L1114:
	;
	if v3079 != 0 {
		goto L1115
	} else {
		goto L1116
	}
L1115:
	;
	v3082 = *(*int32)(unsafe.Add(mBase, uint32(v2393)))
	F_HeapTupleSetHintBits(m, v2393, l2, int32(256), v3082)
	mBase = m.M
	v3084 = m.ExcPending
	if v3084 != 0 {
		goto L64
	} else {
		goto L1118
	}
L1116:
	;
	goto L1117
L1117:
	;
	v3085 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	v3087 = v3085 | int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)) = uint16(v3087)
	F_MarkBufferDirtyHint(m, l2, int32(1))
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L64
	} else {
		goto L1119
	}
L1118:
	;
	goto L871
L1119:
	;
	v3640 = v4
	goto L1
L1120:
	;
	goto L871
L1121:
	;
	if v3101&int32(1024) != 0 {
		goto L1122
	} else {
		goto L1123
	}
L1122:
	;
	v3640 = int32(base.Ui32(v3101&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v3101&int32(4176) == int32(64))
	goto L1
L1123:
	;
	goto L1124
L1124:
	;
	if v3101&int32(4096) != 0 {
		goto L1125
	} else {
		goto L1126
	}
L1125:
	;
	if v3101&int32(128) != 0 {
		v3640 = v3100
		goto L1
	} else {
		goto L1128
	}
L1126:
	;
	goto L1127
L1127:
	;
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	if base.Ui32(v3254) < base.Ui32(int32(3)) {
		goto L1178
	} else {
		goto L1179
	}
L1128:
	;
	if v3101&int32(4176) == int32(64) {
		v3640 = v3100
		goto L1
	} else {
		goto L1129
	}
L1129:
	;
	v3124 = F_HeapTupleGetUpdateXid(m, v2393)
	mBase = m.M
	v3125 = m.ExcPending
	if v3125 != 0 {
		goto L64
	} else {
		goto L1130
	}
L1130:
	;
	if base.Ui32(v3124) < base.Ui32(int32(3)) {
		goto L1132
	} else {
		goto L1133
	}
L1131:
	;
	if v3245 != 0 {
		v3640 = int32(0)
		goto L1
	} else {
		goto L1171
	}
L1132:
	;
	v3245 = int32(0)
	goto L1131
L1133:
	;
	goto L1134
L1134:
	;
	v3136 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v3136 == v3124 {
		goto L1135
	} else {
		goto L1136
	}
L1135:
	;
	v3245 = int32(1)
	goto L1131
L1136:
	;
	goto L1137
L1137:
	;
	v3140 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v3140 <= int32(0) {
		goto L1139
	} else {
		goto L1140
	}
L1138:
	;
	v3245 = v3237
	goto L1131
L1139:
	;
	v3144 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v3144 == int32(0) {
		v3237 = int32(0)
		goto L1138
	} else {
		goto L1142
	}
L1140:
	;
	goto L1141
L1141:
	;
	v3206 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v3208 = int32(0)
	v3210 = v3140 - int32(1)
	goto L1161
L1142:
	;
	v3149 = v3144
	goto L1143
L1143:
	;
	v3154 = *(*int32)(unsafe.Add(mBase, uint32(v3149)+20))
	if v3154 == int32(4) {
		goto L1145
	} else {
		goto L1146
	}
L1144:
	;
	v3237 = int32(0)
	goto L1138
L1145:
	;
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v3149)+80))
	if v3201 != 0 {
		v3149 = v3201
		goto L1143
	} else {
		goto L1160
	}
L1146:
	;
	v3157 = *(*int32)(unsafe.Add(mBase, uint32(v3149)))
	if v3157 == int32(0) {
		goto L1145
	} else {
		goto L1147
	}
L1147:
	;
	v3160 = int32(1)
	if v3124 == v3157 {
		v3237 = v3160
		goto L1138
	} else {
		goto L1148
	}
L1148:
	;
	v3162 = *(*int32)(unsafe.Add(mBase, uint32(v3149)+52))
	v3164 = v3162 - int32(1)
	if v3164 < int32(0) {
		goto L1145
	} else {
		goto L1149
	}
L1149:
	;
	v3169 = int32(0)
	v3171 = v3164
	goto L1150
L1150:
	;
	v3175 = *(*int32)(unsafe.Add(mBase, uint32(v3149)+48))
	v3177 = int32(2)
	v3178 = base.I32_div_s(v3171-v3169, v3177)
	v3179 = v3178 + v3169
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v3175+v3179<<(uint(v3177)%32))))
	if v3183 == v3124 {
		v3237 = v3160
		goto L1138
	} else {
		goto L1152
	}
L1151:
	;
	goto L1145
L1152:
	;
	v3187 = F_TransactionIdPrecedes(m, v3183, v3124)
	mBase = m.M
	if v3187 != 0 {
		goto L1153
	} else {
		goto L1154
	}
L1153:
	;
	v3188 = v3179 + int32(1)
	goto L1155
L1154:
	;
	v3188 = v3169
	goto L1155
L1155:
	;
	if v3187 != 0 {
		goto L1156
	} else {
		goto L1157
	}
L1156:
	;
	v3191 = v3171
	goto L1158
L1157:
	;
	v3191 = v3179 - int32(1)
	goto L1158
L1158:
	;
	if v3188 <= v3191 {
		v3169 = v3188
		v3171 = v3191
		goto L1150
	} else {
		goto L1159
	}
L1159:
	;
	goto L1151
L1160:
	;
	goto L1144
L1161:
	;
	v3215 = int32(2)
	v3216 = base.I32_div_s(v3210-v3208, v3215)
	v3217 = v3216 + v3208
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(v3206+v3217<<(uint(v3215)%32))))
	v3222 = base.B2i32(v3221 == v3124)
	if v3221 == v3124 {
		v3237 = v3222
		goto L1138
	} else {
		goto L1163
	}
L1162:
	;
	v3237 = v3222
	goto L1138
L1163:
	;
	v3225 = base.B2i32(base.Ui32(v3221) < base.Ui32(v3124))
	if base.Ui32(v3221) < base.Ui32(v3124) {
		goto L1164
	} else {
		goto L1165
	}
L1164:
	;
	v3226 = v3217 + int32(1)
	goto L1166
L1165:
	;
	v3226 = v3208
	goto L1166
L1166:
	;
	if base.Ui32(v3221) < base.Ui32(v3124) {
		goto L1167
	} else {
		goto L1168
	}
L1167:
	;
	v3229 = v3210
	goto L1169
L1168:
	;
	v3229 = v3217 - int32(1)
	goto L1169
L1169:
	;
	if v3226 <= v3229 {
		v3208 = v3226
		v3210 = v3229
		goto L1161
	} else {
		goto L1170
	}
L1170:
	;
	goto L1162
L1171:
	;
	v3246 = F_TransactionIdIsInProgress(m, v3124)
	mBase = m.M
	v3247 = m.ExcPending
	if v3247 != 0 {
		goto L64
	} else {
		goto L1172
	}
L1172:
	;
	if v3246 != 0 {
		goto L1173
	} else {
		goto L1174
	}
L1173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v3124
	v3640 = int32(1)
	goto L1
L1174:
	;
	goto L1175
L1175:
	;
	v3250 = F_TransactionIdDidCommit(m, v3124)
	mBase = m.M
	v3251 = m.ExcPending
	if v3251 != 0 {
		goto L64
	} else {
		goto L1176
	}
L1176:
	;
	v3640 = v3250 ^ int32(1)
	goto L1
L1177:
	;
	if v3374 != 0 {
		goto L1217
	} else {
		goto L1218
	}
L1178:
	;
	v3374 = int32(0)
	goto L1177
L1179:
	;
	goto L1180
L1180:
	;
	v3265 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v3265 == v3254 {
		goto L1181
	} else {
		goto L1182
	}
L1181:
	;
	v3374 = int32(1)
	goto L1177
L1182:
	;
	goto L1183
L1183:
	;
	v3269 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v3269 <= int32(0) {
		goto L1185
	} else {
		goto L1186
	}
L1184:
	;
	v3374 = v3366
	goto L1177
L1185:
	;
	v3273 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v3273 == int32(0) {
		v3366 = int32(0)
		goto L1184
	} else {
		goto L1188
	}
L1186:
	;
	goto L1187
L1187:
	;
	v3335 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v3337 = int32(0)
	v3339 = v3269 - int32(1)
	goto L1207
L1188:
	;
	v3278 = v3273
	goto L1189
L1189:
	;
	v3283 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+20))
	if v3283 == int32(4) {
		goto L1191
	} else {
		goto L1192
	}
L1190:
	;
	v3366 = int32(0)
	goto L1184
L1191:
	;
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+80))
	if v3330 != 0 {
		v3278 = v3330
		goto L1189
	} else {
		goto L1206
	}
L1192:
	;
	v3286 = *(*int32)(unsafe.Add(mBase, uint32(v3278)))
	if v3286 == int32(0) {
		goto L1191
	} else {
		goto L1193
	}
L1193:
	;
	v3289 = int32(1)
	if v3254 == v3286 {
		v3366 = v3289
		goto L1184
	} else {
		goto L1194
	}
L1194:
	;
	v3291 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+52))
	v3293 = v3291 - int32(1)
	if v3293 < int32(0) {
		goto L1191
	} else {
		goto L1195
	}
L1195:
	;
	v3298 = int32(0)
	v3300 = v3293
	goto L1196
L1196:
	;
	v3304 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+48))
	v3306 = int32(2)
	v3307 = base.I32_div_s(v3300-v3298, v3306)
	v3308 = v3307 + v3298
	v3312 = *(*int32)(unsafe.Add(mBase, uint32(v3304+v3308<<(uint(v3306)%32))))
	if v3312 == v3254 {
		v3366 = v3289
		goto L1184
	} else {
		goto L1198
	}
L1197:
	;
	goto L1191
L1198:
	;
	v3316 = F_TransactionIdPrecedes(m, v3312, v3254)
	mBase = m.M
	if v3316 != 0 {
		goto L1199
	} else {
		goto L1200
	}
L1199:
	;
	v3317 = v3308 + int32(1)
	goto L1201
L1200:
	;
	v3317 = v3298
	goto L1201
L1201:
	;
	if v3316 != 0 {
		goto L1202
	} else {
		goto L1203
	}
L1202:
	;
	v3320 = v3300
	goto L1204
L1203:
	;
	v3320 = v3308 - int32(1)
	goto L1204
L1204:
	;
	if v3317 <= v3320 {
		v3298 = v3317
		v3300 = v3320
		goto L1196
	} else {
		goto L1205
	}
L1205:
	;
	goto L1197
L1206:
	;
	goto L1190
L1207:
	;
	v3344 = int32(2)
	v3345 = base.I32_div_s(v3339-v3337, v3344)
	v3346 = v3345 + v3337
	v3350 = *(*int32)(unsafe.Add(mBase, uint32(v3335+v3346<<(uint(v3344)%32))))
	v3351 = base.B2i32(v3350 == v3254)
	if v3350 == v3254 {
		v3366 = v3351
		goto L1184
	} else {
		goto L1209
	}
L1208:
	;
	v3366 = v3351
	goto L1184
L1209:
	;
	v3354 = base.B2i32(base.Ui32(v3350) < base.Ui32(v3254))
	if base.Ui32(v3350) < base.Ui32(v3254) {
		goto L1210
	} else {
		goto L1211
	}
L1210:
	;
	v3355 = v3346 + int32(1)
	goto L1212
L1211:
	;
	v3355 = v3337
	goto L1212
L1212:
	;
	if base.Ui32(v3350) < base.Ui32(v3254) {
		goto L1213
	} else {
		goto L1214
	}
L1213:
	;
	v3358 = v3339
	goto L1215
L1214:
	;
	v3358 = v3346 - int32(1)
	goto L1215
L1215:
	;
	if v3355 <= v3358 {
		v3337 = v3355
		v3339 = v3358
		goto L1207
	} else {
		goto L1216
	}
L1216:
	;
	goto L1208
L1217:
	;
	v3375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	v3640 = int32(base.Ui32(v3375&int32(128))>>(uint(int32(7))%32)) | base.B2i32(v3375&int32(4176) == int32(64))
	goto L1
L1218:
	;
	goto L1219
L1219:
	;
	v3385 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	v3386 = F_TransactionIdIsInProgress(m, v3385)
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L64
	} else {
		goto L1220
	}
L1220:
	;
	if v3386 != 0 {
		goto L1221
	} else {
		goto L1222
	}
L1221:
	;
	v3388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	if v3388&int32(128) != 0 {
		v3640 = v3100
		goto L1
	} else {
		goto L1224
	}
L1222:
	;
	goto L1223
L1223:
	;
	v3397 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	v3398 = F_TransactionIdDidCommit(m, v3397)
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L64
	} else {
		goto L1226
	}
L1224:
	;
	if v3388&int32(4176) == int32(64) {
		v3640 = v3100
		goto L1
	} else {
		goto L1225
	}
L1225:
	;
	v3395 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v3395
	v3640 = v3100
	goto L1
L1226:
	;
	v3400 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	if v3398 == int32(0) {
		goto L1227
	} else {
		goto L1228
	}
L1227:
	;
	v3404 = v3400 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)) = uint16(v3404)
	goto L2
L1228:
	;
	goto L1229
L1229:
	;
	v3408 = int32(0)
	if base.B2i32(v3400&int32(128) == v3408)&base.B2i32(v3400&int32(4176) != int32(64)) == v3408 {
		goto L1230
	} else {
		goto L1231
	}
L1230:
	;
	v3418 = v3400 | int32(2048)
	*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)) = uint16(v3418)
	goto L2
L1231:
	;
	goto L1232
L1232:
	;
	v3421 = *(*int32)(unsafe.Add(mBase, uint32(v2393)+4))
	F_HeapTupleSetHintBits(m, v2393, l2, int32(1024), v3421)
	mBase = m.M
	v3423 = m.ExcPending
	if v3423 != 0 {
		goto L64
	} else {
		goto L1233
	}
L1233:
	;
	v3640 = int32(0)
	goto L1
L1234:
	;
	if v3429 == int32(512) {
		goto L1237
	} else {
		goto L1238
	}
L1235:
	;
	v3436 = int32(2)
	goto L1236
L1236:
	;
	v3437 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v3438 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v3439 = *(*int32)(unsafe.Add(mBase, uint32(v3426)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v3436
	if v3437 == int32(0) {
		goto L1242
	} else {
		goto L1243
	}
L1237:
	;
	v3640 = int32(0)
	goto L1
L1238:
	;
	goto L1239
L1239:
	;
	v3435 = *(*int32)(unsafe.Add(mBase, uint32(v3426)))
	v3436 = v3435
	goto L1236
L1240:
	;
	v3640 = int32(0)
	goto L1
L1241:
	;
	v3517 = int32(1)
	v3518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3426)+20)))
	if v3518&int32(2048) != 0 {
		v3640 = v3517
		goto L1
	} else {
		goto L1271
	}
L1242:
	;
	v3469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v3469))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v3436)) == int32(0) {
		goto L1250
	} else {
		goto L1251
	}
L1243:
	;
	v3447 = F_bsearch(m, v12+int32(12), v3438, v3437, int32(4), int32(185))
	mBase = m.M
	v3448 = m.ExcPending
	if v3448 != 0 {
		goto L64
	} else {
		goto L1244
	}
L1244:
	;
	if v3447 == int32(0) {
		goto L1242
	} else {
		goto L1245
	}
L1245:
	;
	v3451 = *(*int32)(unsafe.Add(mBase, uint32(v3426)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v3451
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(-1)
	v3456 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	v3461 = F_ResolveCminCmaxDuringDecoding(m, v3456, l1, l0, l2, v12+int32(12), v12+int32(8))
	mBase = m.M
	v3462 = m.ExcPending
	if v3462 != 0 {
		goto L64
	} else {
		goto L1246
	}
L1246:
	;
	if v3461 == int32(0) {
		goto L1240
	} else {
		goto L1247
	}
L1247:
	;
	v3465 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v3466 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if base.Ui32(v3465) < base.Ui32(v3466) {
		goto L1241
	} else {
		goto L1248
	}
L1248:
	;
	v3640 = int32(0)
	goto L1
L1249:
	;
	if v3481 != 0 {
		goto L1253
	} else {
		goto L1254
	}
L1250:
	;
	v3481 = base.B2i32(base.Ui32(v3436) < base.Ui32(v3469))
	goto L1249
L1251:
	;
	goto L1252
L1252:
	;
	v3481 = int32(base.Ui32(v3436-v3469) >> (uint(int32(31)) % 32))
	goto L1249
L1253:
	;
	v3482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3426)+21)))
	if v3482&int32(1) != 0 {
		goto L1241
	} else {
		goto L1256
	}
L1254:
	;
	goto L1255
L1255:
	;
	v3488 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v3488))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v3436)) == int32(0) {
		goto L1260
	} else {
		goto L1261
	}
L1256:
	;
	v3485 = F_TransactionIdDidCommit(m, v3436)
	mBase = m.M
	v3486 = m.ExcPending
	if v3486 != 0 {
		goto L64
	} else {
		goto L1257
	}
L1257:
	;
	if v3485 != 0 {
		goto L1241
	} else {
		goto L1258
	}
L1258:
	;
	v3640 = int32(0)
	goto L1
L1259:
	;
	if v3500 != 0 {
		goto L1263
	} else {
		goto L1264
	}
L1260:
	;
	v3500 = base.B2i32(base.Ui32(v3488) <= base.Ui32(v3436))
	goto L1259
L1261:
	;
	goto L1262
L1262:
	;
	v3500 = base.B2i32(int32(0) <= v3436-v3488)
	goto L1259
L1263:
	;
	v3640 = int32(0)
	goto L1
L1264:
	;
	goto L1265
L1265:
	;
	v3502 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3503 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v3436
	if v3502 == int32(0) {
		goto L1266
	} else {
		goto L1267
	}
L1266:
	;
	v3640 = int32(0)
	goto L1
L1267:
	;
	goto L1268
L1268:
	;
	v3512 = F_bsearch(m, v12+int32(12), v3503, v3502, int32(4), int32(185))
	mBase = m.M
	v3513 = m.ExcPending
	if v3513 != 0 {
		goto L64
	} else {
		goto L1269
	}
L1269:
	;
	if v3512 != 0 {
		goto L1241
	} else {
		goto L1270
	}
L1270:
	;
	v3640 = int32(0)
	goto L1
L1271:
	;
	if v3518&int32(128) != 0 {
		v3640 = v3517
		goto L1
	} else {
		goto L1272
	}
L1272:
	;
	if v3518&int32(4176) == int32(64) {
		v3640 = v3517
		goto L1
	} else {
		goto L1273
	}
L1273:
	;
	if v3518&int32(4096) != 0 {
		goto L1274
	} else {
		goto L1275
	}
L1274:
	;
	v3529 = F_HeapTupleGetUpdateXid(m, v3426)
	mBase = m.M
	v3530 = m.ExcPending
	if v3530 != 0 {
		goto L64
	} else {
		goto L1277
	}
L1275:
	;
	v3531 = v3439
	goto L1276
L1276:
	;
	v3532 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v3533 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v3531
	if v3532 == int32(0) {
		goto L1278
	} else {
		goto L1279
	}
L1277:
	;
	v3531 = v3529
	goto L1276
L1278:
	;
	v3562 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v3562))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v3531)) == int32(0) {
		goto L1286
	} else {
		goto L1287
	}
L1279:
	;
	v3541 = F_bsearch(m, v12+int32(12), v3533, v3532, int32(4), int32(185))
	mBase = m.M
	v3542 = m.ExcPending
	if v3542 != 0 {
		goto L64
	} else {
		goto L1280
	}
L1280:
	;
	if v3541 == int32(0) {
		goto L1278
	} else {
		goto L1281
	}
L1281:
	;
	v3545 = *(*int32)(unsafe.Add(mBase, uint32(v3426)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v3545
	v3548 = *(*int32)(unsafe.Add(mBase, _consts[80]))
	v3553 = F_ResolveCminCmaxDuringDecoding(m, v3548, l1, l0, l2, v12+int32(12), v12+int32(8))
	mBase = m.M
	v3554 = m.ExcPending
	if v3554 != 0 {
		goto L64
	} else {
		goto L1282
	}
L1282:
	;
	if v3553 == int32(0) {
		v3640 = v3517
		goto L1
	} else {
		goto L1283
	}
L1283:
	;
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v3557 == int32(-1) {
		v3640 = v3517
		goto L1
	} else {
		goto L1284
	}
L1284:
	;
	v3560 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v3640 = base.B2i32(base.Ui32(v3560) <= base.Ui32(v3557))
	goto L1
L1285:
	;
	if v3574 != 0 {
		goto L1289
	} else {
		goto L1290
	}
L1286:
	;
	v3574 = base.B2i32(base.Ui32(v3531) < base.Ui32(v3562))
	goto L1285
L1287:
	;
	goto L1288
L1288:
	;
	v3574 = int32(base.Ui32(v3531-v3562) >> (uint(int32(31)) % 32))
	goto L1285
L1289:
	;
	v3576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3426)+21)))
	if v3576&int32(4) != 0 {
		v3640 = int32(0)
		goto L1
	} else {
		goto L1292
	}
L1290:
	;
	goto L1291
L1291:
	;
	v3583 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v3583))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v3531)) == int32(0) {
		goto L1295
	} else {
		goto L1296
	}
L1292:
	;
	v3579 = F_TransactionIdDidCommit(m, v3531)
	mBase = m.M
	v3580 = m.ExcPending
	if v3580 != 0 {
		goto L64
	} else {
		goto L1293
	}
L1293:
	;
	v3640 = v3579 ^ int32(1)
	goto L1
L1294:
	;
	if v3595 != 0 {
		v3640 = v3517
		goto L1
	} else {
		goto L1298
	}
L1295:
	;
	v3595 = base.B2i32(base.Ui32(v3583) <= base.Ui32(v3531))
	goto L1294
L1296:
	;
	goto L1297
L1297:
	;
	v3595 = base.B2i32(int32(0) <= v3531-v3583)
	goto L1294
L1298:
	;
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v3597 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v3531
	if v3596 == int32(0) {
		v3640 = v3517
		goto L1
	} else {
		goto L1299
	}
L1299:
	;
	v3605 = F_bsearch(m, v12+int32(12), v3597, v3596, int32(4), int32(185))
	mBase = m.M
	v3606 = m.ExcPending
	if v3606 != 0 {
		goto L64
	} else {
		goto L1300
	}
L1300:
	;
	v3640 = base.B2i32(v3605 == int32(0))
	goto L1
L1301:
	;
	if v3614 == int32(2) {
		goto L1302
	} else {
		goto L1303
	}
L1302:
	;
	v3620 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v3621 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v3622 = F_GlobalVisTestIsRemovableXid(m, v3620, v3621)
	mBase = m.M
	v3623 = m.ExcPending
	if v3623 != 0 {
		goto L64
	} else {
		goto L1305
	}
L1303:
	;
	v3625 = v3614
	goto L1304
L1304:
	;
	v3640 = base.B2i32(v3625 != int32(0))
	goto L1
L1305:
	;
	if v3622 != 0 {
		goto L1306
	} else {
		goto L1307
	}
L1306:
	;
	v3624 = int32(0)
	goto L1308
L1307:
	;
	v3624 = int32(2)
	goto L1308
L1308:
	;
	v3625 = v3624
	goto L1304
L1309:
	;
	v3640 = v3633
	goto L1
}
