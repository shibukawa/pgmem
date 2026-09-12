package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_exec_stmts(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v423 int64
	_ = v423
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
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
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
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
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v799 int32
	_ = v799
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v851 int32
	_ = v851
	var v865 int32
	_ = v865
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v916 int64
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v999 int32
	_ = v999
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
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
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1324 int32
	_ = v1324
	var v1351 int32
	_ = v1351
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1634 int32
	_ = v1634
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2020 int32
	_ = v2020
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
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2181 int32
	_ = v2181
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2195 int32
	_ = v2195
	var v2198 int32
	_ = v2198
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2228 int32
	_ = v2228
	var v2233 int32
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2248 int32
	_ = v2248
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2260 int32
	_ = v2260
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2270 int32
	_ = v2270
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2297 int32
	_ = v2297
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2318 int32
	_ = v2318
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2331 int32
	_ = v2331
	var v2334 int32
	_ = v2334
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2352 int32
	_ = v2352
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
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
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2409 int32
	_ = v2409
	var v2411 int32
	_ = v2411
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2424 int32
	_ = v2424
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2437 int32
	_ = v2437
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
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2488 int32
	_ = v2488
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2512 int32
	_ = v2512
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2523 int32
	_ = v2523
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2544 int32
	_ = v2544
	var v2551 int32
	_ = v2551
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2557 int32
	_ = v2557
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2583 int64
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2614 int32
	_ = v2614
	var v2616 int32
	_ = v2616
	var v2618 int32
	_ = v2618
	var v2619 int64
	_ = v2619
	var v2626 int32
	_ = v2626
	var v2628 int32
	_ = v2628
	var v2631 int32
	_ = v2631
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2651 int32
	_ = v2651
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2691 int32
	_ = v2691
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2697 int32
	_ = v2697
	var v2699 int32
	_ = v2699
	var v2700 int64
	_ = v2700
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2708 int32
	_ = v2708
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2741 int64
	_ = v2741
	var v2744 int32
	_ = v2744
	var v2745 int32
	_ = v2745
	var v2749 int32
	_ = v2749
	var v2751 int32
	_ = v2751
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2762 int32
	_ = v2762
	var v2765 int32
	_ = v2765
	var v2769 int32
	_ = v2769
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2778 int32
	_ = v2778
	var v2783 int32
	_ = v2783
	var v2784 int32
	_ = v2784
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2792 int32
	_ = v2792
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2797 int32
	_ = v2797
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2810 int32
	_ = v2810
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2821 int32
	_ = v2821
	var v2823 int32
	_ = v2823
	var v2848 int32
	_ = v2848
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2867 int32
	_ = v2867
	var v2870 int32
	_ = v2870
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2879 int32
	_ = v2879
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2885 int32
	_ = v2885
	var v2886 int32
	_ = v2886
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2899 int32
	_ = v2899
	var v2904 int32
	_ = v2904
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2914 int32
	_ = v2914
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2921 int32
	_ = v2921
	var v2924 int32
	_ = v2924
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2929 int32
	_ = v2929
	var v2934 int32
	_ = v2934
	var v2947 int32
	_ = v2947
	var v2965 int32
	_ = v2965
	var v2968 int32
	_ = v2968
	var v2975 int32
	_ = v2975
	var v2983 int32
	_ = v2983
	var v2989 int32
	_ = v2989
	var v2998 int32
	_ = v2998
	var v2999 int32
	_ = v2999
	var v3002 int32
	_ = v3002
	var v3003 int32
	_ = v3003
	var v3004 int32
	_ = v3004
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3008 int32
	_ = v3008
	var v3009 int32
	_ = v3009
	var v3016 int32
	_ = v3016
	var v3020 int32
	_ = v3020
	var v3021 int32
	_ = v3021
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3033 int32
	_ = v3033
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3045 int32
	_ = v3045
	var v3046 int32
	_ = v3046
	var v3047 int32
	_ = v3047
	var v3048 int32
	_ = v3048
	var v3051 int32
	_ = v3051
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3060 int32
	_ = v3060
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3065 int32
	_ = v3065
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3068 int32
	_ = v3068
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3085 int32
	_ = v3085
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3091 int32
	_ = v3091
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3097 int32
	_ = v3097
	var v3100 int32
	_ = v3100
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3119 int32
	_ = v3119
	var v3122 int32
	_ = v3122
	var v3123 int32
	_ = v3123
	var v3125 int32
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3128 int32
	_ = v3128
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3143 int32
	_ = v3143
	var v3144 int32
	_ = v3144
	var v3145 int32
	_ = v3145
	var v3147 int32
	_ = v3147
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3171 int32
	_ = v3171
	var v3178 int32
	_ = v3178
	var v3181 int32
	_ = v3181
	var v3187 int32
	_ = v3187
	var v3191 int32
	_ = v3191
	var v3196 int32
	_ = v3196
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3206 int32
	_ = v3206
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3249 int32
	_ = v3249
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3255 int32
	_ = v3255
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3259 int64
	_ = v3259
	var v3266 int32
	_ = v3266
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3277 int32
	_ = v3277
	var v3280 int32
	_ = v3280
	var v3284 int32
	_ = v3284
	var v3288 int32
	_ = v3288
	var v3293 int32
	_ = v3293
	var v3297 int32
	_ = v3297
	var v3300 int32
	_ = v3300
	var v3304 int32
	_ = v3304
	var v3309 int32
	_ = v3309
	var v3313 int32
	_ = v3313
	var v3316 int32
	_ = v3316
	var v3320 int32
	_ = v3320
	var v3325 int32
	_ = v3325
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3338 int32
	_ = v3338
	var v3343 int32
	_ = v3343
	var v3345 int64
	_ = v3345
	var v3347 int32
	_ = v3347
	var v3351 int32
	_ = v3351
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3360 int32
	_ = v3360
	var v3367 int32
	_ = v3367
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3375 int32
	_ = v3375
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3381 int32
	_ = v3381
	var v3384 int32
	_ = v3384
	var v3388 int32
	_ = v3388
	var v3394 int32
	_ = v3394
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3404 int32
	_ = v3404
	var v3405 int32
	_ = v3405
	var v3408 int32
	_ = v3408
	var v3409 int32
	_ = v3409
	var v3410 int32
	_ = v3410
	var v3414 int32
	_ = v3414
	var v3417 int32
	_ = v3417
	var v3421 int32
	_ = v3421
	var v3427 int32
	_ = v3427
	var v3432 int32
	_ = v3432
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3435 int32
	_ = v3435
	var v3436 int32
	_ = v3436
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3444 int32
	_ = v3444
	var v3447 int32
	_ = v3447
	var v3449 int32
	_ = v3449
	var v3453 int32
	_ = v3453
	var v3455 int32
	_ = v3455
	var v3457 int32
	_ = v3457
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3463 int32
	_ = v3463
	var v3464 int32
	_ = v3464
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3477 int32
	_ = v3477
	var v3478 int32
	_ = v3478
	var v3481 int32
	_ = v3481
	var v3484 int32
	_ = v3484
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3492 int32
	_ = v3492
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3497 int32
	_ = v3497
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3502 int32
	_ = v3502
	var v3503 int32
	_ = v3503
	var v3504 int32
	_ = v3504
	var v3505 int32
	_ = v3505
	var v3507 int32
	_ = v3507
	var v3508 int32
	_ = v3508
	var v3509 int32
	_ = v3509
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3514 int32
	_ = v3514
	var v3515 int32
	_ = v3515
	var v3516 int32
	_ = v3516
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3530 int32
	_ = v3530
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3544 int32
	_ = v3544
	var v3548 int32
	_ = v3548
	var v3553 int32
	_ = v3553
	var v3554 int32
	_ = v3554
	var v3559 int32
	_ = v3559
	var v3560 int32
	_ = v3560
	var v3561 int32
	_ = v3561
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3567 int32
	_ = v3567
	var v3571 int32
	_ = v3571
	var v3573 int32
	_ = v3573
	var v3574 int32
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3576 int32
	_ = v3576
	var v3577 int32
	_ = v3577
	var v3582 int32
	_ = v3582
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3591 int32
	_ = v3591
	var v3592 int32
	_ = v3592
	var v3594 int32
	_ = v3594
	var v3595 int32
	_ = v3595
	var v3598 int32
	_ = v3598
	var v3599 int32
	_ = v3599
	var v3601 int32
	_ = v3601
	var v3605 int32
	_ = v3605
	var v3606 int32
	_ = v3606
	var v3607 int32
	_ = v3607
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3619 int32
	_ = v3619
	var v3620 int32
	_ = v3620
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3624 int32
	_ = v3624
	var v3627 int32
	_ = v3627
	var v3628 int32
	_ = v3628
	var v3631 int32
	_ = v3631
	var v3640 int32
	_ = v3640
	var v3641 int32
	_ = v3641
	var v3644 int32
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3648 int32
	_ = v3648
	var v3649 int32
	_ = v3649
	var v3650 int32
	_ = v3650
	var v3653 int32
	_ = v3653
	var v3655 int32
	_ = v3655
	var v3656 int32
	_ = v3656
	var v3658 int32
	_ = v3658
	var v3661 int32
	_ = v3661
	var v3663 int32
	_ = v3663
	var v3664 int32
	_ = v3664
	var v3666 int32
	_ = v3666
	var v3667 int32
	_ = v3667
	var v3671 int32
	_ = v3671
	var v3672 int32
	_ = v3672
	var v3674 int32
	_ = v3674
	var v3676 int32
	_ = v3676
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3679 int32
	_ = v3679
	var v3683 int32
	_ = v3683
	var v3685 int64
	_ = v3685
	var v3689 int32
	_ = v3689
	var v3690 int32
	_ = v3690
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3697 int32
	_ = v3697
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3703 int32
	_ = v3703
	var v3705 int32
	_ = v3705
	var v3707 int32
	_ = v3707
	var v3709 int32
	_ = v3709
	var v3711 int64
	_ = v3711
	var v3714 int64
	_ = v3714
	var v3716 int32
	_ = v3716
	var v3717 int32
	_ = v3717
	var v3718 int32
	_ = v3718
	var v3722 int32
	_ = v3722
	var v3728 int32
	_ = v3728
	var v3729 int32
	_ = v3729
	var v3730 int32
	_ = v3730
	var v3734 int32
	_ = v3734
	var v3735 int32
	_ = v3735
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3741 int32
	_ = v3741
	var v3742 int32
	_ = v3742
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3746 int32
	_ = v3746
	var v3749 int32
	_ = v3749
	var v3750 int32
	_ = v3750
	var v3754 int32
	_ = v3754
	var v3756 int32
	_ = v3756
	var v3761 int32
	_ = v3761
	var v3764 int32
	_ = v3764
	var v3767 int32
	_ = v3767
	var v3772 int32
	_ = v3772
	var v3775 int32
	_ = v3775
	var v3782 int32
	_ = v3782
	var v3783 int32
	_ = v3783
	var v3787 int32
	_ = v3787
	var v3792 int32
	_ = v3792
	var v3793 int32
	_ = v3793
	var v3794 int32
	_ = v3794
	var v3799 int32
	_ = v3799
	var v3803 int32
	_ = v3803
	var v3808 int32
	_ = v3808
	var v3812 int32
	_ = v3812
	var v3816 int32
	_ = v3816
	var v3821 int32
	_ = v3821
	var v3825 int32
	_ = v3825
	var v3826 int32
	_ = v3826
	var v3832 int32
	_ = v3832
	var v3837 int32
	_ = v3837
	var v3841 int32
	_ = v3841
	var v3844 int32
	_ = v3844
	var v3845 int32
	_ = v3845
	var v3849 int32
	_ = v3849
	var v3855 int32
	_ = v3855
	var v3860 int32
	_ = v3860
	var v3864 int32
	_ = v3864
	var v3865 int32
	_ = v3865
	var v3866 int32
	_ = v3866
	var v3867 int32
	_ = v3867
	var v3874 int32
	_ = v3874
	var v3879 int32
	_ = v3879
	var v3883 int32
	_ = v3883
	var v3887 int32
	_ = v3887
	var v3892 int32
	_ = v3892
	var v3896 int32
	_ = v3896
	var v3900 int32
	_ = v3900
	var v3905 int32
	_ = v3905
	var v3909 int32
	_ = v3909
	var v3912 int32
	_ = v3912
	var v3916 int32
	_ = v3916
	var v3921 int32
	_ = v3921
	var v3925 int32
	_ = v3925
	var v3928 int32
	_ = v3928
	var v3932 int32
	_ = v3932
	var v3937 int32
	_ = v3937
	var v3941 int32
	_ = v3941
	var v3944 int32
	_ = v3944
	var v3948 int32
	_ = v3948
	var v3953 int32
	_ = v3953
	var v3957 int32
	_ = v3957
	var v3960 int32
	_ = v3960
	var v3964 int32
	_ = v3964
	var v3969 int32
	_ = v3969
	var v3973 int32
	_ = v3973
	var v3976 int32
	_ = v3976
	var v3980 int32
	_ = v3980
	var v3985 int32
	_ = v3985
	var v3989 int32
	_ = v3989
	var v3992 int32
	_ = v3992
	var v3998 int32
	_ = v3998
	var v4003 int32
	_ = v4003
	var v4007 int32
	_ = v4007
	var v4010 int32
	_ = v4010
	var v4014 int32
	_ = v4014
	var v4019 int32
	_ = v4019
	var v4023 int32
	_ = v4023
	var v4026 int32
	_ = v4026
	var v4030 int32
	_ = v4030
	var v4035 int32
	_ = v4035
	var v4039 int32
	_ = v4039
	var v4041 int32
	_ = v4041
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4049 int32
	_ = v4049
	var v4054 int32
	_ = v4054
	var v4058 int32
	_ = v4058
	var v4061 int32
	_ = v4061
	var v4065 int32
	_ = v4065
	var v4070 int32
	_ = v4070
	var v4074 int32
	_ = v4074
	var v4077 int32
	_ = v4077
	var v4078 int32
	_ = v4078
	var v4079 int32
	_ = v4079
	var v4085 int32
	_ = v4085
	var v4090 int32
	_ = v4090
	var v4094 int32
	_ = v4094
	var v4097 int32
	_ = v4097
	var v4098 int32
	_ = v4098
	var v4099 int32
	_ = v4099
	var v4106 int32
	_ = v4106
	var v4111 int32
	_ = v4111
	var v4115 int32
	_ = v4115
	var v4118 int32
	_ = v4118
	var v4122 int32
	_ = v4122
	var v4127 int32
	_ = v4127
	var v4131 int32
	_ = v4131
	var v4134 int32
	_ = v4134
	var v4138 int32
	_ = v4138
	var v4143 int32
	_ = v4143
	var v4147 int32
	_ = v4147
	var v4150 int32
	_ = v4150
	var v4154 int32
	_ = v4154
	var v4159 int32
	_ = v4159
	var v4163 int32
	_ = v4163
	var v4166 int32
	_ = v4166
	var v4170 int32
	_ = v4170
	var v4175 int32
	_ = v4175
	var v4179 int32
	_ = v4179
	var v4182 int32
	_ = v4182
	var v4186 int32
	_ = v4186
	var v4191 int32
	_ = v4191
	var v4195 int32
	_ = v4195
	var v4198 int32
	_ = v4198
	var v4202 int32
	_ = v4202
	var v4207 int32
	_ = v4207
	var v4211 int32
	_ = v4211
	var v4214 int32
	_ = v4214
	var v4218 int32
	_ = v4218
	var v4223 int32
	_ = v4223
	var v4227 int32
	_ = v4227
	var v4230 int32
	_ = v4230
	var v4234 int32
	_ = v4234
	var v4239 int32
	_ = v4239
	var v4243 int32
	_ = v4243
	var v4246 int32
	_ = v4246
	var v4250 int32
	_ = v4250
	var v4255 int32
	_ = v4255
	var v4259 int32
	_ = v4259
	var v4260 int32
	_ = v4260
	var v4261 int32
	_ = v4261
	var v4268 int32
	_ = v4268
	var v4273 int32
	_ = v4273
	var v4275 int32
	_ = v4275
	var v4279 int32
	_ = v4279
	var v4283 int32
	_ = v4283
	var v4288 int32
	_ = v4288
	var v4292 int32
	_ = v4292
	var v4296 int32
	_ = v4296
	var v4301 int32
	_ = v4301
	var v4305 int32
	_ = v4305
	var v4308 int32
	_ = v4308
	var v4312 int32
	_ = v4312
	var v4317 int32
	_ = v4317
	var v4321 int32
	_ = v4321
	var v4324 int32
	_ = v4324
	var v4331 int32
	_ = v4331
	var v4336 int32
	_ = v4336
	var v4340 int32
	_ = v4340
	var v4343 int32
	_ = v4343
	var v4350 int32
	_ = v4350
	var v4355 int32
	_ = v4355
	var v4359 int32
	_ = v4359
	var v4362 int32
	_ = v4362
	var v4369 int32
	_ = v4369
	var v4374 int32
	_ = v4374
	var v4378 int32
	_ = v4378
	var v4381 int32
	_ = v4381
	var v4388 int32
	_ = v4388
	var v4393 int32
	_ = v4393
	var v4397 int32
	_ = v4397
	var v4400 int32
	_ = v4400
	var v4407 int32
	_ = v4407
	var v4412 int32
	_ = v4412
	var v4416 int32
	_ = v4416
	var v4419 int32
	_ = v4419
	var v4426 int32
	_ = v4426
	var v4431 int32
	_ = v4431
	var v4435 int32
	_ = v4435
	var v4438 int32
	_ = v4438
	var v4445 int32
	_ = v4445
	var v4450 int32
	_ = v4450
	var v4454 int32
	_ = v4454
	var v4457 int32
	_ = v4457
	var v4464 int32
	_ = v4464
	var v4469 int32
	_ = v4469
	var v4473 int32
	_ = v4473
	var v4476 int32
	_ = v4476
	var v4483 int32
	_ = v4483
	var v4488 int32
	_ = v4488
	var v4492 int32
	_ = v4492
	var v4493 int32
	_ = v4493
	var v4499 int32
	_ = v4499
	var v4504 int32
	_ = v4504
	var v4508 int32
	_ = v4508
	var v4511 int32
	_ = v4511
	var v4515 int32
	_ = v4515
	var v4520 int32
	_ = v4520
	var v4524 int32
	_ = v4524
	var v4527 int32
	_ = v4527
	var v4531 int32
	_ = v4531
	var v4536 int32
	_ = v4536
	var v4540 int32
	_ = v4540
	var v4543 int32
	_ = v4543
	var v4549 int32
	_ = v4549
	var v4554 int32
	_ = v4554
	var v4558 int32
	_ = v4558
	var v4561 int32
	_ = v4561
	var v4565 int32
	_ = v4565
	var v4570 int32
	_ = v4570
	var v4574 int32
	_ = v4574
	var v4577 int32
	_ = v4577
	var v4581 int32
	_ = v4581
	var v4586 int32
	_ = v4586
	var v4590 int32
	_ = v4590
	var v4592 int32
	_ = v4592
	var v4593 int32
	_ = v4593
	var v4594 int32
	_ = v4594
	var v4600 int32
	_ = v4600
	var v4605 int32
	_ = v4605
	var v4609 int32
	_ = v4609
	var v4612 int32
	_ = v4612
	var v4613 int32
	_ = v4613
	var v4619 int32
	_ = v4619
	var v4624 int32
	_ = v4624
	var v4628 int32
	_ = v4628
	var v4631 int32
	_ = v4631
	var v4637 int32
	_ = v4637
	var v4642 int32
	_ = v4642
	var v4646 int32
	_ = v4646
	var v4649 int32
	_ = v4649
	var v4653 int32
	_ = v4653
	var v4658 int32
	_ = v4658
	var v4662 int32
	_ = v4662
	var v4665 int32
	_ = v4665
	var v4666 int32
	_ = v4666
	var v4672 int32
	_ = v4672
	var v4677 int32
	_ = v4677
	var v4681 int32
	_ = v4681
	var v4684 int32
	_ = v4684
	var v4690 int32
	_ = v4690
	var v4695 int32
	_ = v4695
	var v4697 int32
	_ = v4697
	var v4711 int32
	_ = v4711
	var v4712 int32
	_ = v4712
	var v4715 int32
	_ = v4715
	var v4716 int32
	_ = v4716
	var v4717 int32
	_ = v4717
	var v4718 int32
	_ = v4718
	var v4719 int32
	_ = v4719
	var v4720 int32
	_ = v4720
	var v4721 int32
	_ = v4721
	var v4722 int32
	_ = v4722
	var v4733 int32
	_ = v4733
	var v4736 int32
	_ = v4736
	var v4737 int32
	_ = v4737
	var v4738 int32
	_ = v4738
	var v4739 int32
	_ = v4739
	var v4741 int32
	_ = v4741
	var v4742 int32
	_ = v4742
	var v4750 int32
	_ = v4750
	var v4758 int32
	_ = v4758
	var v4766 int32
	_ = v4766
	var v4774 int32
	_ = v4774
	var v4777 int32
	_ = v4777
	var v4780 int32
	_ = v4780
	var v4781 int32
	_ = v4781
	var v4782 int32
	_ = v4782
	var v4783 int32
	_ = v4783
	var v4785 int32
	_ = v4785
	var v4786 int32
	_ = v4786
	var v4788 int32
	_ = v4788
	var v4794 int32
	_ = v4794
	var v4800 int32
	_ = v4800
	var v4806 int32
	_ = v4806
	var v4809 int32
	_ = v4809
	var v4812 int32
	_ = v4812
	var v4815 int32
	_ = v4815
	var v4818 int32
	_ = v4818
	var v4821 int32
	_ = v4821
	var v4826 int32
	_ = v4826
	var v4828 int32
	_ = v4828
	var v4860 int32
	_ = v4860
	var v4864 int32
	_ = v4864
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4875 int32
	_ = v4875
	var v4876 int32
	_ = v4876
	var v4879 int32
	_ = v4879
	var v4880 int32
	_ = v4880
	var v4883 int32
	_ = v4883
	var v4890 int32
	_ = v4890
	var v4891 int32
	_ = v4891
	var v4894 int32
	_ = v4894
	var v4897 int32
	_ = v4897
	var v4902 int32
	_ = v4902
	var v4908 int32
	_ = v4908
	var v4929 int32
	_ = v4929
	var v4930 int32
	_ = v4930
	var v4934 int32
	_ = v4934
	var v4935 int32
	_ = v4935
	var v4938 int32
	_ = v4938
	var v4939 int32
	_ = v4939
	var v4943 int32
	_ = v4943
	var v4944 int32
	_ = v4944
	var v4949 int32
	_ = v4949
	var v4950 int32
	_ = v4950
	var v4954 int32
	_ = v4954
	var v4955 int32
	_ = v4955
	var v4958 int32
	_ = v4958
	var v4959 int32
	_ = v4959
	var v4962 int32
	_ = v4962
	var v4969 int32
	_ = v4969
	var v4970 int32
	_ = v4970
	var v5002 int32
	_ = v5002
	var v5035 int32
	_ = v5035
	var v5038 int32
	_ = v5038
	var v5039 int32
	_ = v5039
	var v5040 int32
	_ = v5040
	var v5043 int32
	_ = v5043
	var v5044 int32
	_ = v5044
	var v5047 int32
	_ = v5047
	var v5051 int32
	_ = v5051
	var v5053 int32
	_ = v5053
	var v5055 int32
	_ = v5055
	var v5057 int32
	_ = v5057
	var v5059 int32
	_ = v5059
	var v5064 int32
	_ = v5064
	var v5099 int32
	_ = v5099
	var v5100 int32
	_ = v5100
	var v5101 int32
	_ = v5101
	var v5102 int32
	_ = v5102
	var v5106 int32
	_ = v5106
	var v5109 int32
	_ = v5109
	var v5113 int32
	_ = v5113
	var v5117 int32
	_ = v5117
	var v5122 int32
	_ = v5122
	var v5185 int32
	_ = v5185
	var v5186 int32
	_ = v5186
	var v5187 int32
	_ = v5187
	var v5188 int32
	_ = v5188
	var v5219 int32
	_ = v5219
	var v5221 int32
	_ = v5221
	var v5222 int32
	_ = v5222
	var v5225 int32
	_ = v5225
	var v5228 int32
	_ = v5228
	var v5230 int32
	_ = v5230
	var v5234 int32
	_ = v5234
	var v5262 int32
	_ = v5262
	var v5263 int32
	_ = v5263
	var v5266 int32
	_ = v5266
	var v5270 int32
	_ = v5270
	var v5272 int32
	_ = v5272
	var v5273 int32
	_ = v5273
	var v5278 int32
	_ = v5278
	var v5279 int32
	_ = v5279
	var v5317 int32
	_ = v5317
	v3 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(640)
	m.G0 = v33
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v33 + int32(640)
	return v5317
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v35
	v5317 = int32(0)
	goto L1
L3:
	;
	v61 = v48
	v84 = v3
	goto L11
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v36 <= int32(0) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v50 == int32(0) {
		v5317 = v3
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v40 = l0 + int32(20)
	v42 = l0 + int32(16)
	v44 = l0 + int32(12)
	v46 = v33 + int32(616)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
	goto L3
L8:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v5317 = v3
	goto L1
L11:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v84<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v93 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v35
	v5317 = v5234
	goto L1
L13:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v103 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	if v96 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	m.T0[v96].(func(*base.Module, int32, int32))(m, l0, v91)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	switch v106 {
	case 0:
		goto L89
	case 1:
		goto L116
	case 2:
		goto L112
	case 3:
		goto L111
	case 4:
		goto L110
	case 5:
		goto L109
	case 6:
		goto L108
	case 7:
		goto L107
	case 8:
		goto L106
	case 9:
		goto L105
	case 10:
		goto L104
	case 11:
		goto L103
	case 12:
		goto L102
	case 13:
		goto L101
	case 14:
		goto L100
	case 15:
		goto L99
	case 16:
		goto L98
	case 17:
		goto L97
	case 18:
		goto L96
	case 19:
		goto L113
	case 20:
		goto L95
	case 21:
		goto L94
	case 22:
		goto L93
	case 23:
		goto L115
	case 24:
		goto L114
	case 25:
		goto L92
	case 26:
		goto L91
	default:
		goto L90
	}
L20:
	;
	goto L19
L21:
	;
	v5262 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
	v5263 = *(*int32)(unsafe.Add(mBase, uint32(v5262)))
	if v5263 == int32(0) {
		v5273 = v5262
		goto L1517
	} else {
		goto L1518
	}
L22:
	;
	v5219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v5219 != 0 {
		goto L1511
	} else {
		goto L1512
	}
L23:
	;
	v5186 = *(*int32)(unsafe.Add(mBase, uint32(v5185)))
	v5187 = F_exec_stmts(m, l0, v5186)
	mBase = m.M
	v5188 = m.ExcPending
	if v5188 != 0 {
		goto L9
	} else {
		goto L1510
	}
L24:
	;
	v5185 = v91 + int32(24)
	goto L23
L25:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v5106 = m.ExcPending
	if v5106 != 0 {
		goto L9
	} else {
		goto L1505
	}
L26:
	;
	v5100 = *(*int32)(unsafe.Add(mBase, uint32(v5099)))
	v5101 = F_exec_stmts(m, l0, v5100)
	mBase = m.M
	v5102 = m.ExcPending
	if v5102 != 0 {
		goto L9
	} else {
		goto L1504
	}
L27:
	;
	if v1115 != 0 {
		goto L1491
	} else {
		goto L1492
	}
L28:
	;
	v5002 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v5002
	v5234 = v5002
	goto L21
L29:
	;
	v4939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4939 == int32(0) {
		goto L1478
	} else {
		goto L1479
	}
L30:
	;
	v4929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v4930 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v4934 = *(*int32)(unsafe.Add(mBase, uint32(v4929+v4930<<(uint(int32(2))%32))))
	v4935 = int32(0)
	F_assign_simple_var(m, l0, v4934, v4908, v4935, v4935)
	mBase = m.M
	v4938 = m.ExcPending
	if v4938 != 0 {
		goto L9
	} else {
		goto L1477
	}
L31:
	;
	v4897 = int32(0)
	v4902 = v4897
	v4908 = v4897
	goto L30
L32:
	;
	v4860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4860 == int32(0) {
		goto L1460
	} else {
		goto L1461
	}
L33:
	;
	if v4712 == int32(0) {
		goto L1410
	} else {
		goto L1411
	}
L34:
	;
	F_plpgsql_create_econtext(m, l0)
	mBase = m.M
	v4697 = m.ExcPending
	if v4697 != 0 {
		goto L9
	} else {
		goto L1409
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4681 = m.ExcPending
	if v4681 != 0 {
		goto L9
	} else {
		goto L1405
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4662 = m.ExcPending
	if v4662 != 0 {
		goto L9
	} else {
		goto L1401
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4646 = m.ExcPending
	if v4646 != 0 {
		goto L9
	} else {
		goto L1397
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4628 = m.ExcPending
	if v4628 != 0 {
		goto L9
	} else {
		goto L1393
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4609 = m.ExcPending
	if v4609 != 0 {
		goto L9
	} else {
		goto L1389
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4590 = m.ExcPending
	if v4590 != 0 {
		goto L9
	} else {
		goto L1385
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4574 = m.ExcPending
	if v4574 != 0 {
		goto L9
	} else {
		goto L1381
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4558 = m.ExcPending
	if v4558 != 0 {
		goto L9
	} else {
		goto L1377
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4540 = m.ExcPending
	if v4540 != 0 {
		goto L9
	} else {
		goto L1373
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4524 = m.ExcPending
	if v4524 != 0 {
		goto L9
	} else {
		goto L1369
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4508 = m.ExcPending
	if v4508 != 0 {
		goto L9
	} else {
		goto L1365
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4492 = m.ExcPending
	if v4492 != 0 {
		goto L9
	} else {
		goto L1362
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4473 = m.ExcPending
	if v4473 != 0 {
		goto L9
	} else {
		goto L1358
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4454 = m.ExcPending
	if v4454 != 0 {
		goto L9
	} else {
		goto L1354
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4435 = m.ExcPending
	if v4435 != 0 {
		goto L9
	} else {
		goto L1350
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4416 = m.ExcPending
	if v4416 != 0 {
		goto L9
	} else {
		goto L1346
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4397 = m.ExcPending
	if v4397 != 0 {
		goto L9
	} else {
		goto L1342
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4378 = m.ExcPending
	if v4378 != 0 {
		goto L9
	} else {
		goto L1338
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4359 = m.ExcPending
	if v4359 != 0 {
		goto L9
	} else {
		goto L1334
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4340 = m.ExcPending
	if v4340 != 0 {
		goto L9
	} else {
		goto L1330
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4321 = m.ExcPending
	if v4321 != 0 {
		goto L9
	} else {
		goto L1326
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4305 = m.ExcPending
	if v4305 != 0 {
		goto L9
	} else {
		goto L1322
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4292 = m.ExcPending
	if v4292 != 0 {
		goto L9
	} else {
		goto L1319
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4279 = m.ExcPending
	if v4279 != 0 {
		goto L9
	} else {
		goto L1316
	}
L59:
	;
	F_ReThrowError(m, v2758)
	mBase = m.M
	v4275 = m.ExcPending
	if v4275 != 0 {
		goto L9
	} else {
		goto L1315
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4259 = m.ExcPending
	if v4259 != 0 {
		goto L9
	} else {
		goto L1311
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4243 = m.ExcPending
	if v4243 != 0 {
		goto L9
	} else {
		goto L1307
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4227 = m.ExcPending
	if v4227 != 0 {
		goto L9
	} else {
		goto L1303
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4211 = m.ExcPending
	if v4211 != 0 {
		goto L9
	} else {
		goto L1299
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4195 = m.ExcPending
	if v4195 != 0 {
		goto L9
	} else {
		goto L1295
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4179 = m.ExcPending
	if v4179 != 0 {
		goto L9
	} else {
		goto L1291
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4163 = m.ExcPending
	if v4163 != 0 {
		goto L9
	} else {
		goto L1287
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4147 = m.ExcPending
	if v4147 != 0 {
		goto L9
	} else {
		goto L1283
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L9
	} else {
		goto L1279
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4115 = m.ExcPending
	if v4115 != 0 {
		goto L9
	} else {
		goto L1275
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4094 = m.ExcPending
	if v4094 != 0 {
		goto L9
	} else {
		goto L1271
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4074 = m.ExcPending
	if v4074 != 0 {
		goto L9
	} else {
		goto L1266
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4058 = m.ExcPending
	if v4058 != 0 {
		goto L9
	} else {
		goto L1262
	}
L73:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4039 = m.ExcPending
	if v4039 != 0 {
		goto L9
	} else {
		goto L1258
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4023 = m.ExcPending
	if v4023 != 0 {
		goto L9
	} else {
		goto L1254
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v4007 = m.ExcPending
	if v4007 != 0 {
		goto L9
	} else {
		goto L1250
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3989 = m.ExcPending
	if v3989 != 0 {
		goto L9
	} else {
		goto L1246
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3973 = m.ExcPending
	if v3973 != 0 {
		goto L9
	} else {
		goto L1242
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3957 = m.ExcPending
	if v3957 != 0 {
		goto L9
	} else {
		goto L1238
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3941 = m.ExcPending
	if v3941 != 0 {
		goto L9
	} else {
		goto L1234
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3925 = m.ExcPending
	if v3925 != 0 {
		goto L9
	} else {
		goto L1230
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3909 = m.ExcPending
	if v3909 != 0 {
		goto L9
	} else {
		goto L1226
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3896 = m.ExcPending
	if v3896 != 0 {
		goto L9
	} else {
		goto L1223
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3883 = m.ExcPending
	if v3883 != 0 {
		goto L9
	} else {
		goto L1220
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3864 = m.ExcPending
	if v3864 != 0 {
		goto L9
	} else {
		goto L1216
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3841 = m.ExcPending
	if v3841 != 0 {
		goto L9
	} else {
		goto L1212
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3825 = m.ExcPending
	if v3825 != 0 {
		goto L9
	} else {
		goto L1209
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3812 = m.ExcPending
	if v3812 != 0 {
		goto L9
	} else {
		goto L1206
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3799 = m.ExcPending
	if v3799 != 0 {
		goto L9
	} else {
		goto L1203
	}
L89:
	;
	v3793 = F_exec_stmt_block(m, l0, v91)
	mBase = m.M
	v3794 = m.ExcPending
	if v3794 != 0 {
		goto L9
	} else {
		goto L1202
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v35
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3782 = m.ExcPending
	if v3782 != 0 {
		goto L9
	} else {
		goto L1199
	}
L91:
	;
	v3767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+12)))
	if v3767 == int32(1) {
		goto L1194
	} else {
		goto L1195
	}
L92:
	;
	v3756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+12)))
	if v3756 == int32(1) {
		goto L1188
	} else {
		goto L1189
	}
L93:
	;
	v3729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3730 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v3734 = *(*int32)(unsafe.Add(mBase, uint32(v3729+v3730<<(uint(int32(2))%32))))
	v3735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3734)+44)))
	if v3735 == int32(1) {
		goto L36
	} else {
		goto L1182
	}
L94:
	;
	v3606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3607 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v3611 = *(*int32)(unsafe.Add(mBase, uint32(v3606+v3607<<(uint(int32(2))%32))))
	v3612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3611)+44)))
	if v3612 == int32(1) {
		goto L39
	} else {
		goto L1146
	}
L95:
	;
	v3470 = int32(0)
	v3472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3473 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v3477 = *(*int32)(unsafe.Add(mBase, uint32(v3472+v3473<<(uint(int32(2))%32))))
	v3478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3477)+44)))
	if v3478 == v3470 {
		goto L1089
	} else {
		goto L1090
	}
L96:
	;
	v3459 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v3460 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v3463 = F_exec_dynquery_with_params(m, l0, v3459, v3460, int32(0), int32(4))
	mBase = m.M
	v3464 = m.ExcPending
	if v3464 != 0 {
		goto L9
	} else {
		goto L1086
	}
L97:
	;
	v3203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v3203 == int32(0) {
		goto L1001
	} else {
		goto L1002
	}
L98:
	;
	F_exec_stmt_execsql(m, l0, v91)
	mBase = m.M
	v3201 = m.ExcPending
	if v3201 != 0 {
		goto L9
	} else {
		goto L1000
	}
L99:
	;
	v3097 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1352])))
	if v3097 != int32(1) {
		goto L969
	} else {
		goto L970
	}
L100:
	;
	v2755 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v2755 != 0 {
		goto L863
	} else {
		goto L864
	}
L101:
	;
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v2563 == int32(0) {
		goto L807
	} else {
		goto L808
	}
L102:
	;
	v2294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v2294 == int32(0) {
		goto L67
	} else {
		goto L717
	}
L103:
	;
	v2187 = int32(2)
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v2188 != 0 {
		v5234 = v2187
		goto L21
	} else {
		goto L684
	}
L104:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2148 != 0 {
		goto L666
	} else {
		goto L667
	}
L105:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v1871 = F_exec_eval_expr(m, l0, v1864, v33+int32(596), v33+int32(608), v33+int32(636))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L9
	} else {
		goto L587
	}
L106:
	;
	v1721 = int32(0)
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1723+v1724<<(uint(int32(2))%32))))
	v1729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+44)))
	if v1729 == v1721 {
		goto L530
	} else {
		goto L531
	}
L107:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v1712 = F_exec_run_select(m, l0, v1709, v33+int32(608))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L9
	} else {
		goto L527
	}
L108:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1456)+4))
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1455+v1457<<(uint(int32(2))%32))))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v1469 = F_exec_eval_expr(m, l0, v1462, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L9
	} else {
		goto L434
	}
L109:
	;
	goto L391
L110:
	;
	goto L376
L111:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v1066 == int32(0) {
		goto L326
	} else {
		goto L327
	}
L112:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v965 = F_exec_eval_expr(m, l0, v958, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L9
	} else {
		goto L291
	}
L113:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+12)))
	if v455 == int32(1) {
		goto L190
	} else {
		goto L191
	}
L114:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+24))
	if v147 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L115:
	;
	v117 = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v120 = F_exec_run_select(m, l0, v118, v117)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L118
	}
L116:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	F_exec_assign_expr(m, l0, v112, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L9
	} else {
		goto L117
	}
L117:
	;
	v5234 = int32(0)
	goto L21
L118:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122+v123<<(uint(int32(2))%32))))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v131 = int32(0)
	F_assign_simple_var(m, l0, v127, base.B2i32(v128 != int64(0)), v131, v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v135 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_SPI_freetuptable(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v138 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v140 == v138 {
		v5234 = v117
		goto L21
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)+20))
	F_MemoryContextReset(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	v5234 = v117
	goto L21
L126:
	;
	F_exec_prepare_plan(m, l0, v146, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L9
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+16)))
	if v153 != int32(1) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	goto L128
L130:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+56))
	v394 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+616)) = v394
	*(*int64)(unsafe.Add(mBase, uint32(v33)+624)) = v394
	*(*int64)(unsafe.Add(mBase, uint32(v33)+608)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v33)+608)) = v390
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+612)) = uint8(v401)
	v403 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+613)) = uint8(v403)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+628)) = v405
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v146)+24))
	v410 = F_SPI_execute_plan_extended(m, v407, v33+int32(608))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L9
	} else {
		goto L169
	}
L131:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+20)) = v146
	v390 = v388
	goto L130
L132:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	if v357 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L133:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v146)+28))
	if v355 != 0 {
		goto L131
	} else {
		goto L160
	}
L134:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v156 != 0 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v157 = int32(4515600)
	v158 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v146)+24))
	v164 = F_SPI_plan_get_cached_plan(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L136
	}
L136:
	;
	if v164 == int32(0) {
		goto L88
	} else {
		goto L137
	}
L137:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v168 == int32(0) {
		goto L88
	} else {
		goto L138
	}
L138:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v171 != int32(1) {
		goto L88
	} else {
		goto L139
	}
L139:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+88))
	if v176 == int32(0) {
		goto L87
	} else {
		goto L140
	}
L140:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v179 != int32(213) {
		goto L87
	} else {
		goto L141
	}
L141:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v185 = F_SearchSysCache1(m, int32(47), v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L9
	} else {
		goto L142
	}
L142:
	;
	if v185 == int32(0) {
		goto L86
	} else {
		goto L143
	}
L143:
	;
	v195 = F_get_func_arg_info(m, v185, v33+int32(608), v33+int32(636), v33+int32(600))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	F_ReleaseCatCache(m, v185)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+48))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v201
	v204 = F_palloc0(m, int32(40))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v204)+8)) = int32(670932)
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = int32(1)
	v214 = F_palloc(m, v195<<(uint(int32(2))%32))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L9
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+36)) = v214
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v219
	v221 = int32(0)
	if v221 < v195 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v228 = int32(0)
	v230 = v221
	goto L151
L149:
	;
	v292 = v221
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+28)) = v292
	v319 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	F_ReleaseCachedPlan(m, v164, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L9
	} else {
		goto L159
	}
L151:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v33)+600))
	if v255 == int32(0) {
		v282 = v230
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v292 = v282
	goto L150
L153:
	;
	v285 = v228 + int32(1)
	if v285 != v195 {
		v228 = v285
		v230 = v282
		goto L151
	} else {
		goto L158
	}
L154:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228+v255))))
	switch v259 - int32(98) {
	case 0, 13:
		goto L155
	default:
		v282 = v230
		goto L153
	}
L155:
	;
	v263 = v230 << (uint(int32(2)) % 32)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v263+v265)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v268 != int32(8) {
		goto L132
	} else {
		goto L156
	}
L156:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	v273 = v271 - int32(1)
	F_exec_check_assignable(m, l0, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v276+v263))) = v273
	v282 = v230 + int32(1)
	goto L153
L158:
	;
	goto L152
L159:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v204
	goto L133
L160:
	;
	v390 = int32(0)
	goto L130
L161:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L9
	} else {
		goto L165
	}
L162:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v357+v228<<(uint(int32(2))%32))))
	if v363 == int32(0) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	if v366 != 0 {
		goto L85
	} else {
		goto L164
	}
L164:
	;
	goto L161
L165:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L9
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v228 + int32(1)
	F_errmsg(m, int32(391891), v33+int32(48))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(500315), int32(2388), int32(108168))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L9
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	if v410 < int32(0) {
		goto L84
	} else {
		goto L170
	}
L170:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _consts[295]))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+56))
	if v416 != v393 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
	F_plpgsql_create_econtext(m, l0)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L9
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v423 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	if base.Ui64(int64(1)) < base.Ui64(v423) {
		goto L82
	} else {
		goto L175
	}
L174:
	;
	goto L173
L175:
	;
	if base.I32_wrap_i64(v423) == int32(1) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+16)))
	if v429 == int32(0) {
		goto L83
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v441 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v434 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+4))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	F_exec_move_row(m, l0, v432, v436, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	F_SPI_freetuptable(m, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L9
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v444 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v444
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v447 != 0 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	goto L183
L185:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+20))
	F_MemoryContextReset(m, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L9
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	F_SPI_freetuptable(m, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L9
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	v5234 = v444
	goto L21
L190:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v458 == int32(0) {
		goto L81
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v461 == int32(0) {
		goto L22
	} else {
		goto L194
	}
L193:
	;
	goto L192
L194:
	;
	v464 = int32(0)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v465 <= v464 {
		goto L22
	} else {
		goto L195
	}
L195:
	;
	v471 = v464
	goto L196
L196:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v500 = int32(2)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v499+v471<<(uint(v500)%32))))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+4))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v498+v504<<(uint(v500)%32))))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	switch v509 {
	case 0:
		goto L199
	case 1:
		goto L212
	case 2:
		goto L201
	case 3:
		goto L211
	case 4:
		goto L210
	case 5:
		goto L209
	case 6:
		goto L208
	case 7:
		goto L207
	case 8:
		goto L206
	case 9:
		goto L205
	case 10:
		goto L204
	case 11:
		goto L203
	case 12:
		goto L202
	default:
		goto L200
	}
L197:
	;
	goto L22
L198:
	;
	v955 = v471 + int32(1)
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v955 < v956 {
		v471 = v955
		goto L196
	} else {
		goto L290
	}
L199:
	;
	v916 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v917 = F_Int64GetDatum(m, v916)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L9
	} else {
		goto L288
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L9
	} else {
		goto L285
	}
L201:
	;
	v749 = int32(4515600)
	v750 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v752)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v753
	v756 = int32(4508524)
	v758 = *(*int32)(unsafe.Add(mBase, _consts[1145]))
	v759 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1145])) = v758 + v759
	v762 = int32(4122340)
	v764 = *(*int32)(unsafe.Add(mBase, _consts[1146]))
	v766 = v764 + v759
	*(*int32)(unsafe.Add(mBase, _consts[1146])) = v766
	if v766 < int32(5) {
		goto L266
	} else {
		goto L267
	}
L202:
	;
	v730 = int32(4515600)
	v731 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)+60))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v736
	if v733 != 0 {
		goto L260
	} else {
		goto L261
	}
L203:
	;
	v711 = int32(4515600)
	v712 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+64))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v717
	if v714 != 0 {
		goto L255
	} else {
		goto L256
	}
L204:
	;
	v692 = int32(4515600)
	v693 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+32))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v697)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v698
	if v695 != 0 {
		goto L250
	} else {
		goto L251
	}
L205:
	;
	v673 = int32(4515600)
	v674 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v675)+72))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v678)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v679
	if v676 != 0 {
		goto L245
	} else {
		goto L246
	}
L206:
	;
	v654 = int32(4515600)
	v655 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)+76))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v660
	if v657 != 0 {
		goto L240
	} else {
		goto L241
	}
L207:
	;
	v635 = int32(4515600)
	v636 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+68))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v641
	if v638 != 0 {
		goto L235
	} else {
		goto L236
	}
L208:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)+28))
	v576 = int32(4509336)
	v577 = int32(63)
	v579 = int32(48)
	v580 = v575&v577 + v579
	*(*uint8)(unsafe.Add(mBase, _consts[1158])) = uint8(v580)
	v588 = int32(base.Ui32(v575)>>(uint(int32(24))%32))&v577 + v579
	*(*uint8)(unsafe.Add(mBase, _consts[1159])) = uint8(v588)
	v596 = int32(base.Ui32(v575)>>(uint(int32(18))%32))&v577 + v579
	*(*uint8)(unsafe.Add(mBase, _consts[1160])) = uint8(v596)
	v604 = int32(base.Ui32(v575)>>(uint(int32(12))%32))&v577 + v579
	*(*uint8)(unsafe.Add(mBase, _consts[1161])) = uint8(v604)
	v612 = int32(base.Ui32(v575)>>(uint(int32(6))%32))&v577 + v579
	*(*uint8)(unsafe.Add(mBase, _consts[1162])) = uint8(v612)
	v615 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1163])) = uint8(v615)
	goto L229
L209:
	;
	v555 = int32(4515600)
	v556 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v557)+44))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v561
	if v558 != 0 {
		goto L224
	} else {
		goto L225
	}
L210:
	;
	v536 = int32(4515600)
	v537 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+36))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v542
	if v539 != 0 {
		goto L219
	} else {
		goto L220
	}
L211:
	;
	v517 = int32(4515600)
	v518 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+48))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v523
	if v520 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v510)+36))
	F_exec_assign_value(m, l0, v508, v511, int32(0), int32(26), int32(-1))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L9
	} else {
		goto L213
	}
L213:
	;
	goto L198
L214:
	;
	v526 = v520
	goto L216
L215:
	;
	v526 = int32(757603)
	goto L216
L216:
	;
	v527 = F_cstring_to_text(m, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L9
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v518
	F_exec_assign_value(m, l0, v508, v527, int32(0), int32(25), int32(-1))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L9
	} else {
		goto L218
	}
L218:
	;
	goto L198
L219:
	;
	v545 = v539
	goto L221
L220:
	;
	v545 = int32(757603)
	goto L221
L221:
	;
	v546 = F_cstring_to_text(m, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L9
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v537
	F_exec_assign_value(m, l0, v508, v546, int32(0), int32(25), int32(-1))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L9
	} else {
		goto L223
	}
L223:
	;
	goto L198
L224:
	;
	v564 = v558
	goto L226
L225:
	;
	v564 = int32(757603)
	goto L226
L226:
	;
	v565 = F_cstring_to_text(m, v564)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L9
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v556
	F_exec_assign_value(m, l0, v508, v565, int32(0), int32(25), int32(-1))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L9
	} else {
		goto L228
	}
L228:
	;
	goto L198
L229:
	;
	v618 = int32(4515600)
	v619 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v622
	goto L230
L230:
	;
	goto L232
L232:
	;
	v626 = F_cstring_to_text(m, v576)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L9
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v619
	F_exec_assign_value(m, l0, v508, v626, int32(0), int32(25), int32(-1))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L9
	} else {
		goto L234
	}
L234:
	;
	goto L198
L235:
	;
	v644 = v638
	goto L237
L236:
	;
	v644 = int32(757603)
	goto L237
L237:
	;
	v645 = F_cstring_to_text(m, v644)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L9
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v636
	F_exec_assign_value(m, l0, v508, v645, int32(0), int32(25), int32(-1))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L9
	} else {
		goto L239
	}
L239:
	;
	goto L198
L240:
	;
	v663 = v657
	goto L242
L241:
	;
	v663 = int32(757603)
	goto L242
L242:
	;
	v664 = F_cstring_to_text(m, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L9
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v655
	F_exec_assign_value(m, l0, v508, v664, int32(0), int32(25), int32(-1))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L9
	} else {
		goto L244
	}
L244:
	;
	goto L198
L245:
	;
	v682 = v676
	goto L247
L246:
	;
	v682 = int32(757603)
	goto L247
L247:
	;
	v683 = F_cstring_to_text(m, v682)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L9
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v674
	F_exec_assign_value(m, l0, v508, v683, int32(0), int32(25), int32(-1))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L9
	} else {
		goto L249
	}
L249:
	;
	goto L198
L250:
	;
	v701 = v695
	goto L252
L251:
	;
	v701 = int32(757603)
	goto L252
L252:
	;
	v702 = F_cstring_to_text(m, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v693
	F_exec_assign_value(m, l0, v508, v702, int32(0), int32(25), int32(-1))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L9
	} else {
		goto L254
	}
L254:
	;
	goto L198
L255:
	;
	v720 = v714
	goto L257
L256:
	;
	v720 = int32(757603)
	goto L257
L257:
	;
	v721 = F_cstring_to_text(m, v720)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L9
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v712
	F_exec_assign_value(m, l0, v508, v721, int32(0), int32(25), int32(-1))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L9
	} else {
		goto L259
	}
L259:
	;
	goto L198
L260:
	;
	v739 = v733
	goto L262
L261:
	;
	v739 = int32(757603)
	goto L262
L262:
	;
	v740 = F_cstring_to_text(m, v739)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L9
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v731
	F_exec_assign_value(m, l0, v508, v740, int32(0), int32(25), int32(-1))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L9
	} else {
		goto L264
	}
L264:
	;
	goto L198
L265:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v887
	if v851 != 0 {
		goto L280
	} else {
		goto L281
	}
L266:
	;
	v770 = int32(100)
	v771 = v766 * v770
	v777 = F__emscripten_memset_bulkmem(m, v771+int32(4508528), base.I32_extend8_s(int32(0)), v770)
	mBase = m.M
	goto L269
L267:
	;
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1146])) = int32(-1)
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L9
	} else {
		goto L277
	}
L269:
	;
	v781 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	*(*int32)(unsafe.Add(mBase, uint32(v771)+uint32(_consts[1209]))) = v781
	v786 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v771)+uint32(_consts[1148]))) = v786
	v790 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	if v790 != 0 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v799 = v790
	goto L273
L271:
	;
	v842 = v764
	v851 = int32(0)
	v865 = v758
	goto L272
L272:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1145])) = v865
	*(*int32)(unsafe.Add(mBase, _consts[1146])) = v842
	goto L265
L273:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v799)+8))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v799)+4))
	m.T0[v822].(func(*base.Module, int32))(m, v821)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L9
	} else {
		goto L275
	}
L274:
	;
	v827 = *(*int32)(unsafe.Add(mBase, _consts[1146]))
	v828 = int32(1)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v777)+48))
	v832 = *(*int32)(unsafe.Add(mBase, _consts[1145]))
	v842 = v827 - v828
	v851 = v830
	v865 = v832 - v828
	goto L272
L275:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	if v825 != 0 {
		v799 = v825
		goto L273
	} else {
		goto L276
	}
L276:
	;
	goto L274
L277:
	;
	F_errmsg_internal(m, int32(462813), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L9
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(498425), int32(762), int32(12159))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L9
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	v890 = v851
	goto L282
L281:
	;
	v890 = int32(757603)
	goto L282
L282:
	;
	v891 = F_cstring_to_text(m, v890)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L9
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v750
	F_exec_assign_value(m, l0, v508, v891, int32(0), int32(25), int32(-1))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L9
	} else {
		goto L284
	}
L284:
	;
	goto L198
L285:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v904
	F_errmsg_internal(m, int32(487632), v33+int32(80))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L9
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(500315), int32(2510), int32(338060))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L9
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	F_exec_assign_value(m, l0, v508, v917, int32(0), int32(20), int32(-1))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L9
	} else {
		goto L289
	}
L289:
	;
	goto L198
L290:
	;
	goto L197
L291:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v973 = F_exec_cast_value(m, l0, v965, v33+int32(600), v969, v970, int32(16), int32(-1))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L9
	} else {
		goto L292
	}
L292:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v975 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	F_SPI_freetuptable(m, v975)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L9
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v980 != 0 {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	goto L295
L297:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v980)+20))
	F_MemoryContextReset(m, v981)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L9
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v984 != 0 {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	goto L299
L301:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v989 == int32(0) {
		goto L24
	} else {
		goto L304
	}
L302:
	;
	if v973 == int32(0) {
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v5185 = v91 + int32(16)
	goto L23
L304:
	;
	v992 = int32(0)
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v989)+4))
	if v993 <= v992 {
		goto L24
	} else {
		goto L305
	}
L305:
	;
	v999 = v992
	goto L306
L306:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v989)+12))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1026+v999<<(uint(int32(2))%32))))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1030)+4))
	v1038 = F_exec_eval_expr(m, l0, v1031, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L9
	} else {
		goto L308
	}
L307:
	;
	v5185 = v1030 + int32(8)
	goto L23
L308:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v1046 = F_exec_cast_value(m, l0, v1038, v33+int32(600), v1042, v1043, int32(16), int32(-1))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L9
	} else {
		goto L309
	}
L309:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1048 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	F_SPI_freetuptable(m, v1048)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L9
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1053 != 0 {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	goto L312
L314:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+20))
	F_MemoryContextReset(m, v1054)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L9
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v1046 != 0 {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	goto L316
L318:
	;
	v1059 = v1057
	goto L320
L319:
	;
	v1059 = int32(1)
	goto L320
L320:
	;
	if v1059 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1061 = v999 + int32(1)
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v989)+4))
	if v1062 <= v1061 {
		goto L24
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	goto L307
L324:
	;
	v999 = v1061
	goto L306
L325:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v1116 == int32(0) {
		goto L27
	} else {
		goto L343
	}
L326:
	;
	v1115 = int32(0)
	goto L325
L327:
	;
	goto L328
L328:
	;
	v1076 = F_exec_eval_expr(m, l0, v1066, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L9
	} else {
		goto L329
	}
L329:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1080+v1081<<(uint(int32(2))%32))))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+24))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+4))
	if v1079 == v1087 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	F_exec_assign_value(m, l0, v1085, v1076, v1097, v1079, v1078)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L9
	} else {
		goto L336
	}
L331:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+24))
	if v1089 == v1078 {
		goto L330
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+44))
	v1094 = F_plpgsql_build_datatype(m, v1079, v1078, v1092, int32(0))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L9
	} else {
		goto L335
	}
L334:
	;
	goto L333
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1085)+24)) = v1094
	goto L330
L336:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1100 != 0 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	F_SPI_freetuptable(m, v1100)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L9
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1103
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1105 == v1103 {
		v1115 = v1085
		goto L325
	} else {
		goto L341
	}
L340:
	;
	goto L339
L341:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+20))
	F_MemoryContextReset(m, v1108)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L9
	} else {
		goto L342
	}
L342:
	;
	v1115 = v1085
	goto L325
L343:
	;
	v1119 = int32(0)
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+4))
	if v1120 <= v1119 {
		goto L27
	} else {
		goto L344
	}
L344:
	;
	v1126 = v1119
	goto L345
L345:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+12))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1153+v1126<<(uint(int32(2))%32))))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+4))
	v1165 = F_exec_eval_expr(m, l0, v1158, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L9
	} else {
		goto L347
	}
L346:
	;
	if v1115 != 0 {
		goto L364
	} else {
		goto L365
	}
L347:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v1173 = F_exec_cast_value(m, l0, v1165, v33+int32(600), v1169, v1170, int32(16), int32(-1))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L9
	} else {
		goto L348
	}
L348:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1175 != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	F_SPI_freetuptable(m, v1175)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L9
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1180 != 0 {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	goto L351
L353:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1180)+20))
	F_MemoryContextReset(m, v1181)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L9
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v1173 != 0 {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	goto L355
L357:
	;
	v1186 = v1184
	goto L359
L358:
	;
	v1186 = int32(1)
	goto L359
L359:
	;
	if v1186 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1188 = v1126 + int32(1)
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+4))
	if v1189 <= v1188 {
		goto L27
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	goto L346
L363:
	;
	v1126 = v1188
	goto L345
L364:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115)+45)))
	if v1191 != int32(1) {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	goto L366
L366:
	;
	v5099 = v1157 + int32(8)
	goto L26
L367:
	;
	v1213 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1115)+48)) = v1213
	v1215 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1115)+44)) = uint16(v1215)
	*(*int32)(unsafe.Add(mBase, uint32(v1115)+40)) = v1213
	goto L366
L368:
	;
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115)+44)))
	if v1194 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+40))
	F_pfree(m, v1209)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L9
	} else {
		goto L375
	}
L370:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+24))
	v1196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1195)+12)))
	if v1196 != int32(65535) {
		goto L369
	} else {
		goto L371
	}
L371:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+40))
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1199))))
	if v1200 != int32(1) {
		goto L369
	} else {
		goto L372
	}
L372:
	;
	v1203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1199)+1)))
	if v1203 != int32(3) {
		goto L369
	} else {
		goto L373
	}
L373:
	;
	F_DeleteExpandedObject(m, v1199)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L9
	} else {
		goto L374
	}
L374:
	;
	goto L367
L375:
	;
	goto L367
L376:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v1253 = F_exec_stmts(m, l0, v1252)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L9
	} else {
		goto L379
	}
L378:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1257 == int32(0) {
		goto L376
	} else {
		goto L380
	}
L379:
	;
	switch v1253 - int32(1) {
	case 0:
		goto L29
	case 1:
		v5234 = v1253
		goto L21
	case 2:
		goto L378
	default:
		goto L376
	}
L380:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v1260 == int32(0) {
		v5234 = v1253
		goto L21
	} else {
		goto L381
	}
L381:
	;
	v1265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1257))))
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1260))))
	if v1266 == int32(0) {
		v1285 = v1265
		v1286 = v1266
		goto L383
	} else {
		goto L384
	}
L382:
	;
	if v1286-v1285 != 0 {
		v5234 = v1253
		goto L21
	} else {
		goto L390
	}
L383:
	;
	goto L382
L384:
	;
	if v1265 != v1266 {
		v1285 = v1265
		v1286 = v1266
		goto L383
	} else {
		goto L385
	}
L385:
	;
	v1270 = v1260
	v1271 = v1257
	goto L386
L386:
	;
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271)+1)))
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1270)+1)))
	if v1275 == int32(0) {
		v1285 = v1274
		v1286 = v1275
		goto L383
	} else {
		goto L388
	}
L387:
	;
	v1285 = v1274
	v1286 = v1275
	goto L383
L388:
	;
	v1278 = int32(1)
	if v1274 == v1275 {
		v1270 = v1270 + v1278
		v1271 = v1271 + v1278
		goto L386
	} else {
		goto L389
	}
L389:
	;
	goto L387
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
	goto L376
L391:
	;
	v1324 = int32(0)
	goto L393
L393:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v1358 = F_exec_eval_expr(m, l0, v1351, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L9
	} else {
		goto L395
	}
L395:
	;
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v1366 = F_exec_cast_value(m, l0, v1358, v33+int32(600), v1362, v1363, int32(16), int32(-1))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L9
	} else {
		goto L396
	}
L396:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1368 != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	F_SPI_freetuptable(m, v1368)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L9
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1373 != 0 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	goto L399
L401:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+20))
	F_MemoryContextReset(m, v1374)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L9
	} else {
		goto L404
	}
L402:
	;
	goto L403
L403:
	;
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v1377 != 0 {
		v5234 = v1324
		goto L21
	} else {
		goto L405
	}
L404:
	;
	goto L403
L405:
	;
	if v1366 == int32(0) {
		v5234 = v1324
		goto L21
	} else {
		goto L406
	}
L406:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v1381 = F_exec_stmts(m, l0, v1380)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L9
	} else {
		goto L409
	}
L407:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1420 == int32(0) {
		goto L391
	} else {
		goto L423
	}
L408:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1385 == int32(0) {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	switch v1381 - int32(1) {
	case 0:
		goto L408
	case 1:
		v5234 = v1381
		goto L21
	case 2:
		goto L407
	default:
		v1324 = v1381
		goto L393
	}
L410:
	;
	v5234 = int32(0)
	goto L21
L411:
	;
	goto L412
L412:
	;
	v1389 = int32(1)
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v1390 == int32(0) {
		v5234 = v1389
		goto L21
	} else {
		goto L413
	}
L413:
	;
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1385))))
	v1396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1390))))
	if v1396 == int32(0) {
		v1415 = v1395
		v1416 = v1396
		goto L415
	} else {
		goto L416
	}
L414:
	;
	if v1416-v1415 == int32(0) {
		goto L28
	} else {
		goto L422
	}
L415:
	;
	goto L414
L416:
	;
	if v1395 != v1396 {
		v1415 = v1395
		v1416 = v1396
		goto L415
	} else {
		goto L417
	}
L417:
	;
	v1400 = v1390
	v1401 = v1385
	goto L418
L418:
	;
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1401)+1)))
	v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400)+1)))
	if v1405 == int32(0) {
		v1415 = v1404
		v1416 = v1405
		goto L415
	} else {
		goto L420
	}
L419:
	;
	v1415 = v1404
	v1416 = v1405
	goto L415
L420:
	;
	v1408 = int32(1)
	if v1404 == v1405 {
		v1400 = v1400 + v1408
		v1401 = v1401 + v1408
		goto L418
	} else {
		goto L421
	}
L421:
	;
	goto L419
L422:
	;
	v5234 = v1389
	goto L21
L423:
	;
	v1423 = int32(3)
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v1424 == int32(0) {
		v5234 = v1423
		goto L21
	} else {
		goto L424
	}
L424:
	;
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1420))))
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1424))))
	if v1430 == int32(0) {
		v1449 = v1429
		v1450 = v1430
		goto L426
	} else {
		goto L427
	}
L425:
	;
	if v1450-v1449 != 0 {
		v5234 = v1423
		goto L21
	} else {
		goto L433
	}
L426:
	;
	goto L425
L427:
	;
	if v1429 != v1430 {
		v1449 = v1429
		v1450 = v1430
		goto L426
	} else {
		goto L428
	}
L428:
	;
	v1434 = v1424
	v1435 = v1420
	goto L429
L429:
	;
	v1438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1435)+1)))
	v1439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434)+1)))
	if v1439 == int32(0) {
		v1449 = v1438
		v1450 = v1439
		goto L426
	} else {
		goto L431
	}
L430:
	;
	v1449 = v1438
	v1450 = v1439
	goto L426
L431:
	;
	v1442 = int32(1)
	if v1438 == v1439 {
		v1434 = v1434 + v1442
		v1435 = v1435 + v1442
		goto L429
	} else {
		goto L432
	}
L432:
	;
	goto L430
L433:
	;
	v1452 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1452
	v1324 = v1452
	goto L393
L434:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+24))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+4))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+24))
	v1478 = F_exec_cast_value(m, l0, v1469, v33+int32(600), v1473, v1474, v1476, v1477)
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L9
	} else {
		goto L435
	}
L435:
	;
	v1480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v1480 == int32(1) {
		goto L80
	} else {
		goto L436
	}
L436:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1483 != 0 {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	F_SPI_freetuptable(m, v1483)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L9
	} else {
		goto L440
	}
L438:
	;
	goto L439
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1488 != 0 {
		goto L441
	} else {
		goto L442
	}
L440:
	;
	goto L439
L441:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1488)+20))
	F_MemoryContextReset(m, v1489)
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L9
	} else {
		goto L444
	}
L442:
	;
	goto L443
L443:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v1499 = F_exec_eval_expr(m, l0, v1492, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L9
	} else {
		goto L445
	}
L444:
	;
	goto L443
L445:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+24))
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+4))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1505)+24))
	v1508 = F_exec_cast_value(m, l0, v1499, v33+int32(600), v1503, v1504, v1506, v1507)
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L9
	} else {
		goto L446
	}
L446:
	;
	v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v1510 == int32(1) {
		goto L79
	} else {
		goto L447
	}
L447:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1513 != 0 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	F_SPI_freetuptable(m, v1513)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L9
	} else {
		goto L451
	}
L449:
	;
	goto L450
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1518 != 0 {
		goto L452
	} else {
		goto L453
	}
L451:
	;
	goto L450
L452:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1518)+20))
	F_MemoryContextReset(m, v1519)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L9
	} else {
		goto L455
	}
L453:
	;
	goto L454
L454:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	if v1522 == int32(0) {
		goto L457
	} else {
		goto L458
	}
L455:
	;
	goto L454
L456:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	if v1559 != 0 {
		goto L473
	} else {
		goto L474
	}
L457:
	;
	v1558 = int32(1)
	goto L456
L458:
	;
	goto L459
L459:
	;
	v1532 = F_exec_eval_expr(m, l0, v1522, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L9
	} else {
		goto L460
	}
L460:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+24))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1538)+4))
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v1538)+24))
	v1541 = F_exec_cast_value(m, l0, v1532, v33+int32(600), v1536, v1537, v1539, v1540)
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L9
	} else {
		goto L461
	}
L461:
	;
	v1543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v1543 == int32(1) {
		goto L78
	} else {
		goto L462
	}
L462:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1546 != 0 {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	F_SPI_freetuptable(m, v1546)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L9
	} else {
		goto L466
	}
L464:
	;
	goto L465
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1551 != 0 {
		goto L467
	} else {
		goto L468
	}
L466:
	;
	goto L465
L467:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1551)+20))
	F_MemoryContextReset(m, v1552)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L9
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	if v1541 <= int32(0) {
		goto L77
	} else {
		goto L471
	}
L470:
	;
	goto L469
L471:
	;
	v1558 = v1541
	goto L456
L472:
	;
	v1562 = int32(0)
	F_assign_simple_var(m, l0, v1461, v1478, v1562, v1562)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L9
	} else {
		goto L478
	}
L473:
	;
	if v1508 <= v1478 {
		goto L472
	} else {
		goto L476
	}
L474:
	;
	goto L475
L475:
	;
	if v1508 < v1478 {
		goto L31
	} else {
		goto L477
	}
L476:
	;
	goto L31
L477:
	;
	goto L472
L478:
	;
	v1566 = int32(1)
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v1568 = F_exec_stmts(m, l0, v1567)
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L9
	} else {
		goto L481
	}
L479:
	;
	v1612 = v1558 ^ int32(2147483647)
	v1614 = v1558 | int32(-2147483648)
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	if v1615 != 0 {
		goto L494
	} else {
		goto L495
	}
L480:
	;
	v1572 = int32(0)
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1573 == v1572 {
		v1608 = v1572
		goto L479
	} else {
		goto L482
	}
L481:
	;
	switch v1568 - int32(1) {
	case 0:
		goto L32
	case 1:
		v4902 = v1568
		v4908 = v1566
		goto L30
	case 2:
		goto L480
	default:
		v1608 = v1568
		goto L479
	}
L482:
	;
	v1576 = int32(3)
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v1577 == int32(0) {
		v4902 = v1576
		v4908 = v1566
		goto L30
	} else {
		goto L483
	}
L483:
	;
	v1582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1573))))
	v1583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577))))
	if v1583 == int32(0) {
		v1602 = v1582
		v1603 = v1583
		goto L485
	} else {
		goto L486
	}
L484:
	;
	if v1603-v1602 != 0 {
		v4902 = v1576
		v4908 = v1566
		goto L30
	} else {
		goto L492
	}
L485:
	;
	goto L484
L486:
	;
	if v1582 != v1583 {
		v1602 = v1582
		v1603 = v1583
		goto L485
	} else {
		goto L487
	}
L487:
	;
	v1587 = v1577
	v1588 = v1573
	goto L488
L488:
	;
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1588)+1)))
	v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587)+1)))
	if v1592 == int32(0) {
		v1602 = v1591
		v1603 = v1592
		goto L485
	} else {
		goto L490
	}
L489:
	;
	v1602 = v1591
	v1603 = v1592
	goto L485
L490:
	;
	v1595 = int32(1)
	if v1591 == v1592 {
		v1587 = v1587 + v1595
		v1588 = v1588 + v1595
		goto L488
	} else {
		goto L491
	}
L491:
	;
	goto L489
L492:
	;
	v1605 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1605
	v1608 = v1605
	goto L479
L493:
	;
	v1626 = v1608
	v1628 = v1622
	v1634 = v1621
	goto L499
L494:
	;
	if v1478 < v1614 {
		v4902 = v1608
		v4908 = v1566
		goto L30
	} else {
		goto L497
	}
L495:
	;
	goto L496
L496:
	;
	if v1612 < v1478 {
		v4902 = v1608
		v4908 = v1566
		goto L30
	} else {
		goto L498
	}
L497:
	;
	v1621 = int32(1)
	v1622 = v1478 - v1558
	goto L493
L498:
	;
	v1621 = v1615
	v1622 = v1478 + v1558
	goto L493
L499:
	;
	if v1634 != 0 {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	v1655 = int32(0)
	F_assign_simple_var(m, l0, v1461, v1628, v1655, v1655)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L9
	} else {
		goto L507
	}
L502:
	;
	if v1508 <= v1628 {
		goto L501
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	if v1508 < v1628 {
		v4902 = v1626
		v4908 = v1566
		goto L30
	} else {
		goto L506
	}
L505:
	;
	v4902 = v1626
	v4908 = v1566
	goto L30
L506:
	;
	goto L501
L507:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v91)+36))
	v1660 = F_exec_stmts(m, l0, v1659)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L9
	} else {
		goto L510
	}
L508:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	if v1703 != 0 {
		goto L522
	} else {
		goto L523
	}
L509:
	;
	v1664 = int32(0)
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1665 == v1664 {
		v1700 = v1664
		goto L508
	} else {
		goto L511
	}
L510:
	;
	switch v1660 - int32(1) {
	case 0:
		goto L32
	case 1:
		v4902 = v1660
		v4908 = v1566
		goto L30
	case 2:
		goto L509
	default:
		v1700 = v1660
		goto L508
	}
L511:
	;
	v1668 = int32(3)
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v1669 == int32(0) {
		v4902 = v1668
		v4908 = v1566
		goto L30
	} else {
		goto L512
	}
L512:
	;
	v1674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1665))))
	v1675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1669))))
	if v1675 == int32(0) {
		v1694 = v1674
		v1695 = v1675
		goto L514
	} else {
		goto L515
	}
L513:
	;
	if v1695-v1694 != 0 {
		v4902 = v1668
		v4908 = v1566
		goto L30
	} else {
		goto L521
	}
L514:
	;
	goto L513
L515:
	;
	if v1674 != v1675 {
		v1694 = v1674
		v1695 = v1675
		goto L514
	} else {
		goto L516
	}
L516:
	;
	v1679 = v1669
	v1680 = v1665
	goto L517
L517:
	;
	v1683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1680)+1)))
	v1684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1679)+1)))
	if v1684 == int32(0) {
		v1694 = v1683
		v1695 = v1684
		goto L514
	} else {
		goto L519
	}
L518:
	;
	v1694 = v1683
	v1695 = v1684
	goto L514
L519:
	;
	v1687 = int32(1)
	if v1683 == v1684 {
		v1679 = v1679 + v1687
		v1680 = v1680 + v1687
		goto L517
	} else {
		goto L520
	}
L520:
	;
	goto L518
L521:
	;
	v1697 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1697
	v1700 = v1697
	goto L508
L522:
	;
	if v1628 < v1614 {
		v4902 = v1700
		v4908 = v1566
		goto L30
	} else {
		goto L525
	}
L523:
	;
	goto L524
L524:
	;
	if v1612 < v1628 {
		v4902 = v1700
		v4908 = v1566
		goto L30
	} else {
		goto L526
	}
L525:
	;
	v1626 = v1700
	v1628 = v1628 - v1558
	v1634 = int32(1)
	goto L499
L526:
	;
	v1626 = v1700
	v1628 = v1628 + v1558
	v1634 = v1703
	goto L499
L527:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v1716 = F_exec_for_query(m, l0, v91, v1714, int32(1))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L9
	} else {
		goto L528
	}
L528:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	F_SPI_cursor_close(m, v1718)
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L9
	} else {
		goto L529
	}
L529:
	;
	v5234 = v1716
	goto L21
L530:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v1732 == int32(0) {
		goto L533
	} else {
		goto L534
	}
L531:
	;
	v1755 = v1721
	v1757 = v1721
	goto L532
L532:
	;
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	if v1758 != 0 {
		goto L541
	} else {
		goto L542
	}
L533:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v1740 = F_AllocSetContextCreateInternal(m, v1735, int32(504671), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L9
	} else {
		goto L536
	}
L534:
	;
	v1743 = v1732
	goto L535
L535:
	;
	v1744 = int32(4515600)
	v1745 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1743
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+40))
	v1749 = F_text_to_cstring(m, v1748)
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L9
	} else {
		goto L537
	}
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v1740
	v1743 = v1740
	goto L535
L537:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1745
	v1753 = F_GetPortalByName(m, v1749)
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L9
	} else {
		goto L538
	}
L538:
	;
	if v1753 != 0 {
		goto L76
	} else {
		goto L539
	}
L539:
	;
	v1755 = v1749
	v1757 = v1743
	goto L532
L540:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+28))
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+24))
	if v1789 == int32(0) {
		goto L547
	} else {
		goto L548
	}
L541:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+32))
	if v1759 < int32(0) {
		goto L75
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+32))
	if int32(0) <= v1783 {
		goto L74
	} else {
		goto L546
	}
L544:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+608)) = int32(16)
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v1769 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+626)) = uint8(v1769)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+620)) = v1758
	*(*int32)(unsafe.Add(mBase, uint32(v33)+612)) = v1768
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1773+v1759<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+628)) = v1777
	F_exec_stmt_execsql(m, l0, v33+int32(608))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L9
	} else {
		goto L545
	}
L545:
	;
	goto L540
L546:
	;
	goto L540
L547:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+36))
	F_exec_prepare_plan(m, l0, v1788, v1792)
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L9
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+28))
	if v1795 == int32(0) {
		goto L552
	} else {
		goto L553
	}
L550:
	;
	goto L549
L551:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+24))
	v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	v1804 = F_SPI_cursor_open_with_paramlist(m, v1755, v1802, v1801, v1803)
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L9
	} else {
		goto L555
	}
L552:
	;
	v1801 = int32(0)
	goto L551
L553:
	;
	goto L554
L554:
	;
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1799)+20)) = v1788
	v1801 = v1799
	goto L551
L555:
	;
	if v1804 == int32(0) {
		goto L73
	} else {
		goto L556
	}
L556:
	;
	if v1755 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L557:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	F_exec_check_assignable(m, l0, v1810)
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L9
	} else {
		goto L560
	}
L558:
	;
	goto L559
L559:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1820 != 0 {
		goto L563
	} else {
		goto L564
	}
L560:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1804)))
	v1814 = F_cstring_to_text(m, v1813)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L9
	} else {
		goto L561
	}
L561:
	;
	F_assign_simple_var(m, l0, v1728, v1814, int32(0), int32(1))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L9
	} else {
		goto L562
	}
L562:
	;
	goto L559
L563:
	;
	F_SPI_freetuptable(m, v1820)
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L9
	} else {
		goto L566
	}
L564:
	;
	goto L565
L565:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1825 != 0 {
		goto L567
	} else {
		goto L568
	}
L566:
	;
	goto L565
L567:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1825)+20))
	F_MemoryContextReset(m, v1826)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L9
	} else {
		goto L570
	}
L568:
	;
	goto L569
L569:
	;
	if v1757 != 0 {
		goto L571
	} else {
		goto L572
	}
L570:
	;
	goto L569
L571:
	;
	F_MemoryContextReset(m, v1757)
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L9
	} else {
		goto L574
	}
L572:
	;
	goto L573
L573:
	;
	v1832 = F_exec_for_query(m, l0, v91, v1804, int32(0))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L9
	} else {
		goto L575
	}
L574:
	;
	goto L573
L575:
	;
	F_SPI_cursor_close(m, v1804)
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L9
	} else {
		goto L576
	}
L576:
	;
	if v1755 != 0 {
		v5234 = v1832
		goto L21
	} else {
		goto L577
	}
L577:
	;
	v1836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+45)))
	if v1836 != int32(1) {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v1858 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1728)+48)) = v1858
	v1860 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1728)+44)) = uint16(v1860)
	*(*int32)(unsafe.Add(mBase, uint32(v1728)+40)) = v1858
	v5234 = v1832
	goto L21
L579:
	;
	v1839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1728)+44)))
	if v1839 != 0 {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+40))
	F_pfree(m, v1854)
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L9
	} else {
		goto L586
	}
L581:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+24))
	v1841 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1840)+12)))
	if v1841 != int32(65535) {
		goto L580
	} else {
		goto L582
	}
L582:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+40))
	v1845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1844))))
	if v1845 != int32(1) {
		goto L580
	} else {
		goto L583
	}
L583:
	;
	v1848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1844)+1)))
	if v1848 != int32(3) {
		goto L580
	} else {
		goto L584
	}
L584:
	;
	F_DeleteExpandedObject(m, v1844)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L9
	} else {
		goto L585
	}
L585:
	;
	goto L578
L586:
	;
	goto L578
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+600)) = v1871
	v1874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+596)))
	if v1874 == int32(1) {
		goto L72
	} else {
		goto L588
	}
L588:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v1877 == int32(0) {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v1885 = F_AllocSetContextCreateInternal(m, v1880, int32(504671), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L9
	} else {
		goto L592
	}
L590:
	;
	v1887 = v1877
	goto L591
L591:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1887
	v1891 = int32(4515600)
	v1892 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1887
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v1896 = F_get_element_type(m, v1895)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L9
	} else {
		goto L593
	}
L592:
	;
	v1887 = v1885
	goto L591
L593:
	;
	if v1896 == int32(0) {
		goto L71
	} else {
		goto L594
	}
L594:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v33)+600))
	v1901 = F_pg_detoast_datum_copy(m, v1900)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L9
	} else {
		goto L595
	}
L595:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1903 != 0 {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	F_SPI_freetuptable(m, v1903)
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L9
	} else {
		goto L599
	}
L597:
	;
	goto L598
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v1908 != 0 {
		goto L600
	} else {
		goto L601
	}
L599:
	;
	goto L598
L600:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v1908)+20))
	F_MemoryContextReset(m, v1909)
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L9
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v1912 < int32(0) {
		goto L70
	} else {
		goto L604
	}
L603:
	;
	goto L602
L604:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+4))
	if v1915 < v1912 {
		goto L70
	} else {
		goto L605
	}
L605:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v1920 = int32(2)
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1918+v1919<<(uint(v1920)%32))))
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1923)))
	if base.Ui32(v1920) <= base.Ui32(v1924-int32(1)) {
		goto L606
	} else {
		goto L607
	}
L606:
	;
	v1929 = F_plpgsql_exec_get_datum_type(m, l0, v1923)
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L9
	} else {
		goto L609
	}
L607:
	;
	v1934 = v1912
	v1935 = int32(0)
	goto L608
L608:
	;
	v1936 = int32(0)
	if base.B2i32(v1935 == v1936)&base.B2i32(v1936 < v1934) != 0 {
		goto L69
	} else {
		goto L611
	}
L609:
	;
	v1931 = F_get_element_type(m, v1929)
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L9
	} else {
		goto L610
	}
L610:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v1934 = v1933
	v1935 = v1931
	goto L608
L611:
	;
	if v1935 != 0 {
		goto L612
	} else {
		goto L613
	}
L612:
	;
	v1942 = v1934
	goto L614
L613:
	;
	v1942 = int32(1)
	goto L614
L614:
	;
	if v1942 == int32(0) {
		goto L68
	} else {
		goto L615
	}
L615:
	;
	v1946 = F_array_create_iterator(m, v1901, v1934, int32(0))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L9
	} else {
		goto L616
	}
L616:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if int32(0) < v1952 {
		goto L617
	} else {
		goto L618
	}
L617:
	;
	v1955 = v33 + int32(608)
	goto L619
L618:
	;
	v1955 = v1901 + int32(12)
	goto L619
L619:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1955)))
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v1958 = int32(0)
	v1964 = F_array_iterate(m, v1946, v33+int32(600), v33+int32(596))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L9
	} else {
		goto L621
	}
L620:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1892
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v2132
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2132)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v2134
	F_MemoryContextReset(m, v1887)
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L9
	} else {
		goto L664
	}
L621:
	;
	if v1964 == int32(0) {
		v2103 = v1958
		v2105 = v1958
		goto L620
	} else {
		goto L622
	}
L622:
	;
	goto L623
L623:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1892
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(v33)+600))
	v2001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+596)))
	F_exec_assign_value(m, l0, v1923, v2000, v2001, v1956, v1957)
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L9
	} else {
		goto L625
	}
L624:
	;
	v2103 = v2089
	v2105 = v2004
	goto L620
L625:
	;
	v2004 = int32(1)
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if int32(0) < v2005 {
		goto L626
	} else {
		goto L627
	}
L626:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v33)+600))
	F_pfree(m, v2008)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L9
	} else {
		goto L629
	}
L627:
	;
	goto L628
L628:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v2012 = F_exec_stmts(m, l0, v2011)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L9
	} else {
		goto L633
	}
L629:
	;
	goto L628
L630:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v1887
	v2098 = F_array_iterate(m, v1946, v33+int32(600), v33+int32(596))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L9
	} else {
		goto L662
	}
L631:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2053 == int32(0) {
		goto L649
	} else {
		goto L650
	}
L632:
	;
	v2016 = int32(0)
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2017 == v2016 {
		v2103 = v2016
		v2105 = v2004
		goto L620
	} else {
		goto L634
	}
L633:
	;
	switch v2012 - int32(1) {
	case 0:
		goto L632
	case 1:
		v2103 = v2012
		v2105 = v2004
		goto L620
	case 2:
		goto L631
	default:
		v2089 = v2012
		goto L630
	}
L634:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v2020 == int32(0) {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	v2103 = int32(1)
	v2105 = v2004
	goto L620
L636:
	;
	goto L637
L637:
	;
	v2026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2017))))
	v2027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2020))))
	if v2027 == int32(0) {
		v2046 = v2026
		v2047 = v2027
		goto L639
	} else {
		goto L640
	}
L638:
	;
	if v2047-v2046 != 0 {
		goto L646
	} else {
		goto L647
	}
L639:
	;
	goto L638
L640:
	;
	if v2026 != v2027 {
		v2046 = v2026
		v2047 = v2027
		goto L639
	} else {
		goto L641
	}
L641:
	;
	v2031 = v2020
	v2032 = v2017
	goto L642
L642:
	;
	v2035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2032)+1)))
	v2036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2031)+1)))
	if v2036 == int32(0) {
		v2046 = v2035
		v2047 = v2036
		goto L639
	} else {
		goto L644
	}
L643:
	;
	v2046 = v2035
	v2047 = v2036
	goto L639
L644:
	;
	v2039 = int32(1)
	if v2035 == v2036 {
		v2031 = v2031 + v2039
		v2032 = v2032 + v2039
		goto L642
	} else {
		goto L645
	}
L645:
	;
	goto L643
L646:
	;
	v2103 = int32(1)
	v2105 = v2004
	goto L620
L647:
	;
	goto L648
L648:
	;
	v2050 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2050
	v2103 = v2050
	v2105 = v2004
	goto L620
L649:
	;
	v2089 = int32(0)
	goto L630
L650:
	;
	goto L651
L651:
	;
	v2057 = int32(3)
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v2058 == int32(0) {
		v2103 = v2057
		v2105 = v2004
		goto L620
	} else {
		goto L652
	}
L652:
	;
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2053))))
	v2064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2058))))
	if v2064 == int32(0) {
		v2083 = v2063
		v2084 = v2064
		goto L654
	} else {
		goto L655
	}
L653:
	;
	if v2084-v2083 != 0 {
		v2103 = v2057
		v2105 = v2004
		goto L620
	} else {
		goto L661
	}
L654:
	;
	goto L653
L655:
	;
	if v2063 != v2064 {
		v2083 = v2063
		v2084 = v2064
		goto L654
	} else {
		goto L656
	}
L656:
	;
	v2068 = v2058
	v2069 = v2053
	goto L657
L657:
	;
	v2072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2069)+1)))
	v2073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2068)+1)))
	if v2073 == int32(0) {
		v2083 = v2072
		v2084 = v2073
		goto L654
	} else {
		goto L659
	}
L658:
	;
	v2083 = v2072
	v2084 = v2073
	goto L654
L659:
	;
	v2076 = int32(1)
	if v2072 == v2073 {
		v2068 = v2068 + v2076
		v2069 = v2069 + v2076
		goto L657
	} else {
		goto L660
	}
L660:
	;
	goto L658
L661:
	;
	v2086 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2086
	v2089 = v2086
	goto L630
L662:
	;
	if v2098 != 0 {
		goto L623
	} else {
		goto L663
	}
L663:
	;
	goto L624
L664:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v2138+v2139<<(uint(int32(2))%32))))
	v2144 = int32(0)
	F_assign_simple_var(m, l0, v2143, v2105, v2144, v2144)
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L9
	} else {
		goto L665
	}
L665:
	;
	v5234 = v2103
	goto L21
L666:
	;
	v2155 = F_exec_eval_expr(m, l0, v2148, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v2156 = m.ExcPending
	if v2156 != 0 {
		goto L9
	} else {
		goto L669
	}
L667:
	;
	goto L668
L668:
	;
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2181
	v2185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+12)))
	if v2185 != 0 {
		goto L681
	} else {
		goto L682
	}
L669:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v2163 = F_exec_cast_value(m, l0, v2155, v33+int32(600), v2159, v2160, int32(16), int32(-1))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L9
	} else {
		goto L670
	}
L670:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2165 != 0 {
		goto L671
	} else {
		goto L672
	}
L671:
	;
	F_SPI_freetuptable(m, v2165)
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L9
	} else {
		goto L674
	}
L672:
	;
	goto L673
L673:
	;
	v2168 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2168
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2171 != 0 {
		goto L675
	} else {
		goto L676
	}
L674:
	;
	goto L673
L675:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2171)+20))
	F_MemoryContextReset(m, v2172)
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L9
	} else {
		goto L678
	}
L676:
	;
	goto L677
L677:
	;
	v2175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v2175 != 0 {
		v5234 = v2168
		goto L21
	} else {
		goto L679
	}
L678:
	;
	goto L677
L679:
	;
	if v2163 == int32(0) {
		v5234 = v2168
		goto L21
	} else {
		goto L680
	}
L680:
	;
	goto L668
L681:
	;
	v2186 = int32(1)
	goto L683
L682:
	;
	v2186 = int32(3)
	goto L683
L683:
	;
	v5234 = v2186
	goto L21
L684:
	;
	v2189 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v2189
	v2191 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v2191)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v2189
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v2189 <= v2195 {
		goto L685
	} else {
		goto L686
	}
L685:
	;
	v2198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v2198+v2195<<(uint(int32(2))%32))))
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2202)))
	switch v2203 {
	case 0:
		goto L690
	case 1, 2:
		goto L689
	default:
		goto L688
	case 4:
		goto L691
	}
L686:
	;
	goto L687
L687:
	;
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v2254 != 0 {
		goto L703
	} else {
		goto L704
	}
L688:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v2241 = m.ExcPending
	if v2241 != 0 {
		goto L9
	} else {
		goto L700
	}
L689:
	;
	F_exec_eval_datum(m, l0, v2202, v40, v33+int32(608), v44, v42)
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L9
	} else {
		goto L699
	}
L690:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v2202)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2206
	v2208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2202)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v2208)
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v2202)+24))
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v2210)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2211
	v2213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2213 != int32(1) {
		v5234 = v2187
		goto L21
	} else {
		goto L693
	}
L691:
	;
	F_plpgsql_fulfill_promise(m, l0, v2202)
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L9
	} else {
		goto L692
	}
L692:
	;
	goto L690
L693:
	;
	if v2208&int32(1) != 0 {
		v5234 = v2187
		goto L21
	} else {
		goto L694
	}
L694:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L9
	} else {
		goto L695
	}
L695:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L9
	} else {
		goto L696
	}
L696:
	;
	F_errmsg(m, int32(369766), int32(0))
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L9
	} else {
		goto L697
	}
L697:
	;
	F_errfinish(m, int32(500315), int32(3255), int32(244944))
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L9
	} else {
		goto L698
	}
L698:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L699:
	;
	v5234 = v2187
	goto L21
L700:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v2202)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v2242
	F_errmsg_internal(m, int32(484534), v33+int32(160))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L9
	} else {
		goto L701
	}
L701:
	;
	F_errfinish(m, int32(500315), int32(3275), int32(244944))
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L9
	} else {
		goto L702
	}
L702:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L703:
	;
	v2257 = F_exec_eval_expr(m, l0, v2254, v42, v40, v33+int32(608))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L9
	} else {
		goto L706
	}
L704:
	;
	goto L705
L705:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v2283 != int32(2278) {
		v5234 = v2187
		goto L21
	} else {
		goto L715
	}
L706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2257
	v2260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2260 != int32(1) {
		v5234 = v2187
		goto L21
	} else {
		goto L707
	}
L707:
	;
	v2263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v2263 != 0 {
		v5234 = v2187
		goto L21
	} else {
		goto L708
	}
L708:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v2265 = F_type_is_rowtype(m, v2264)
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L9
	} else {
		goto L709
	}
L709:
	;
	if v2265 != 0 {
		v5234 = v2187
		goto L21
	} else {
		goto L710
	}
L710:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v2270 = m.ExcPending
	if v2270 != 0 {
		goto L9
	} else {
		goto L711
	}
L711:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L9
	} else {
		goto L712
	}
L712:
	;
	F_errmsg(m, int32(369766), int32(0))
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		goto L9
	} else {
		goto L713
	}
L713:
	;
	F_errfinish(m, int32(500315), int32(3298), int32(244944))
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L9
	} else {
		goto L714
	}
L714:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L715:
	;
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2286)+65)))
	if v2287 == int32(112) {
		v5234 = v2187
		goto L21
	} else {
		goto L716
	}
L716:
	;
	v2290 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v2290)
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(2278)
	v5234 = v2187
	goto L21
L717:
	;
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2297 == int32(0) {
		goto L718
	} else {
		goto L719
	}
L718:
	;
	F_exec_init_tuple_store(m, l0)
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L9
	} else {
		goto L721
	}
L719:
	;
	goto L720
L720:
	;
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(v2302)))
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if int32(0) <= v2304 {
		goto L723
	} else {
		goto L724
	}
L721:
	;
	goto L720
L722:
	;
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2551 != 0 {
		goto L801
	} else {
		goto L802
	}
L723:
	;
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2311 = *(*int32)(unsafe.Add(mBase, uint32(v2307+v2304<<(uint(int32(2))%32))))
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v2311)))
	switch v2312 {
	case 0:
		v2316 = v2303
		goto L729
	case 1:
		goto L727
	case 2:
		goto L728
	default:
		goto L726
	case 4:
		goto L730
	}
L724:
	;
	goto L725
L725:
	;
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v2430 != 0 {
		goto L768
	} else {
		goto L769
	}
L726:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L9
	} else {
		goto L764
	}
L727:
	;
	v2399 = int32(4515600)
	v2400 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2402)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2403
	v2405 = F_make_tuple_from_row(m, l0, v2311, v2302)
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L9
	} else {
		goto L761
	}
L728:
	;
	v2359 = *(*int32)(unsafe.Add(mBase, uint32(v2311)+36))
	if v2359 == int32(0) {
		goto L742
	} else {
		goto L743
	}
L729:
	;
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v2311)+40))
	v2318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2311)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+636)) = uint8(v2318)
	if v2303 != int32(1) {
		goto L66
	} else {
		goto L732
	}
L730:
	;
	F_plpgsql_fulfill_promise(m, l0, v2311)
	mBase = m.M
	v2314 = m.ExcPending
	if v2314 != 0 {
		goto L9
	} else {
		goto L731
	}
L731:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2302)))
	v2316 = v2315
	goto L729
L732:
	;
	v2324 = v2302 + v2316<<(uint(int32(4))%32)
	if v2318&int32(1) != 0 {
		v2341 = v2317
		goto L733
	} else {
		goto L734
	}
L733:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(v2311)+24))
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v2344)+4))
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v2344)+24))
	v2347 = *(*int32)(unsafe.Add(mBase, uint32(v2324)+88))
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(v2324)+96))
	v2349 = F_exec_cast_value(m, l0, v2341, v33+int32(636), v2345, v2346, v2347, v2348)
	mBase = m.M
	v2350 = m.ExcPending
	if v2350 != 0 {
		goto L9
	} else {
		goto L740
	}
L734:
	;
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2311)+24))
	v2328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2327)+12)))
	if v2328 != int32(65535) {
		v2341 = v2317
		goto L733
	} else {
		goto L735
	}
L735:
	;
	v2331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2317))))
	if v2331 != int32(1) {
		v2340 = v2317
		goto L737
	} else {
		goto L738
	}
L736:
	;
	v2341 = v2340
	goto L733
L737:
	;
	goto L736
L738:
	;
	v2334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2317)+1)))
	if v2334 != int32(3) {
		v2340 = v2317
		goto L737
	} else {
		goto L739
	}
L739:
	;
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(v2317)+2))
	v2340 = v2337 + int32(18)
	goto L737
L740:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+608)) = v2349
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_putvalues(m, v2352, v2302, v33+int32(608), v33+int32(636))
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
		goto L9
	} else {
		goto L741
	}
L741:
	;
	goto L722
L742:
	;
	F_instantiate_empty_record_variable(m, l0, v2311)
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		goto L9
	} else {
		goto L745
	}
L743:
	;
	v2365 = v2359
	goto L744
L744:
	;
	v2366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2365)+28)))
	if v2366&int32(5) == int32(0) {
		goto L746
	} else {
		goto L747
	}
L745:
	;
	v2364 = *(*int32)(unsafe.Add(mBase, uint32(v2311)+36))
	v2365 = v2364
	goto L744
L746:
	;
	F_deconstruct_expanded_record(m, v2365)
	mBase = m.M
	v2372 = m.ExcPending
	if v2372 != 0 {
		goto L9
	} else {
		goto L749
	}
L747:
	;
	v2374 = v2365
	goto L748
L748:
	;
	v2375 = int32(4515600)
	v2376 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2378)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2379
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2374)+44))
	if v2381 != 0 {
		goto L750
	} else {
		goto L751
	}
L749:
	;
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2311)+36))
	v2374 = v2373
	goto L748
L750:
	;
	v2384 = v2381
	goto L752
L751:
	;
	v2382 = F_expanded_record_fetch_tupdesc(m, v2374)
	mBase = m.M
	v2383 = m.ExcPending
	if v2383 != 0 {
		goto L9
	} else {
		goto L753
	}
L752:
	;
	v2386 = F_convert_tuples_by_position(m, v2384, v2302, int32(517189))
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		goto L9
	} else {
		goto L754
	}
L753:
	;
	v2384 = v2382
	goto L752
L754:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2311)+36))
	v2389 = F_expanded_record_get_tuple(m, v2388)
	mBase = m.M
	v2390 = m.ExcPending
	if v2390 != 0 {
		goto L9
	} else {
		goto L755
	}
L755:
	;
	if v2386 != 0 {
		goto L756
	} else {
		goto L757
	}
L756:
	;
	v2391 = F_execute_attr_map_tuple(m, v2389, v2386)
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L9
	} else {
		goto L759
	}
L757:
	;
	v2393 = v2389
	goto L758
L758:
	;
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_puttuple(m, v2394, v2393)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L9
	} else {
		goto L760
	}
L759:
	;
	v2393 = v2391
	goto L758
L760:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2376
	goto L722
L761:
	;
	if v2405 == int32(0) {
		goto L65
	} else {
		goto L762
	}
L762:
	;
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_puttuple(m, v2409, v2405)
	mBase = m.M
	v2411 = m.ExcPending
	if v2411 != 0 {
		goto L9
	} else {
		goto L763
	}
L763:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2400
	goto L722
L764:
	;
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(v2311)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v2418
	F_errmsg_internal(m, int32(484534), v33+int32(176))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L9
	} else {
		goto L765
	}
L765:
	;
	F_errfinish(m, int32(500315), int32(3444), int32(63483))
	mBase = m.M
	v2429 = m.ExcPending
	if v2429 != 0 {
		goto L9
	} else {
		goto L766
	}
L766:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L767:
	;
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2530 = *(*int32)(unsafe.Add(mBase, uint32(v2529)+20))
	v2533 = F_MemoryContextAllocZero(m, v2530, v2303<<(uint(int32(2))%32))
	mBase = m.M
	v2534 = m.ExcPending
	if v2534 != 0 {
		goto L9
	} else {
		goto L797
	}
L768:
	;
	v2437 = F_exec_eval_expr(m, l0, v2430, v33+int32(607), v33+int32(600), v33+int32(596))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L9
	} else {
		goto L771
	}
L769:
	;
	goto L770
L770:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v2516 = m.ExcPending
	if v2516 != 0 {
		goto L9
	} else {
		goto L793
	}
L771:
	;
	v2439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v2439 == int32(1) {
		goto L772
	} else {
		goto L773
	}
L772:
	;
	v2442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+607)))
	if v2442 != 0 {
		goto L767
	} else {
		goto L775
	}
L773:
	;
	goto L774
L774:
	;
	if v2303 != int32(1) {
		goto L63
	} else {
		goto L790
	}
L775:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v33)+600))
	v2444 = F_type_is_rowtype(m, v2443)
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L9
	} else {
		goto L776
	}
L776:
	;
	if v2444 == int32(0) {
		goto L64
	} else {
		goto L777
	}
L777:
	;
	v2448 = int32(4515600)
	v2449 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2452 = *(*int32)(unsafe.Add(mBase, uint32(v2451)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2452
	v2454 = F_pg_detoast_datum(m, v2437)
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L9
	} else {
		goto L778
	}
L778:
	;
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2454)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+624)) = v2454
	v2458 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+620)) = v2458
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+616)) = uint16(v2458)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+612)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+608)) = int32(base.Ui32(v2456) >> (uint(int32(2)) % 32))
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v2454)+8))
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2454)+4))
	v2469 = F_lookup_rowtype_tupdesc(m, v2467, v2468)
	mBase = m.M
	v2470 = m.ExcPending
	if v2470 != 0 {
		goto L9
	} else {
		goto L779
	}
L779:
	;
	v2472 = F_convert_tuples_by_position(m, v2469, v2302, int32(370752))
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L9
	} else {
		goto L780
	}
L780:
	;
	if v2472 != 0 {
		goto L781
	} else {
		goto L782
	}
L781:
	;
	v2476 = F_execute_attr_map_tuple(m, v33+int32(608), v2472)
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L9
	} else {
		goto L784
	}
L782:
	;
	v2480 = v33 + int32(608)
	goto L783
L783:
	;
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_puttuple(m, v2481, v2480)
	mBase = m.M
	v2483 = m.ExcPending
	if v2483 != 0 {
		goto L9
	} else {
		goto L785
	}
L784:
	;
	v2480 = v2476
	goto L783
L785:
	;
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(v2469)+12))
	if int32(0) <= v2484 {
		goto L786
	} else {
		goto L787
	}
L786:
	;
	F_DecrTupleDescRefCount(m, v2469)
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L9
	} else {
		goto L789
	}
L787:
	;
	goto L788
L788:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2449
	goto L722
L789:
	;
	goto L788
L790:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v33)+600))
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(v33)+596))
	v2497 = *(*int32)(unsafe.Add(mBase, uint32(v2302)))
	v2500 = v2302 + v2497<<(uint(int32(4))%32)
	v2501 = *(*int32)(unsafe.Add(mBase, uint32(v2500)+88))
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(v2500)+96))
	v2503 = F_exec_cast_value(m, l0, v2437, v33+int32(607), v2495, v2496, v2501, v2502)
	mBase = m.M
	v2504 = m.ExcPending
	if v2504 != 0 {
		goto L9
	} else {
		goto L791
	}
L791:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+636)) = v2503
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_putvalues(m, v2506, v2302, v33+int32(636), v33+int32(607))
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L9
	} else {
		goto L792
	}
L792:
	;
	goto L722
L793:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		goto L9
	} else {
		goto L794
	}
L794:
	;
	F_errmsg(m, int32(216873), int32(0))
	mBase = m.M
	v2523 = m.ExcPending
	if v2523 != 0 {
		goto L9
	} else {
		goto L795
	}
L795:
	;
	F_errfinish(m, int32(500315), int32(3529), int32(63483))
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L9
	} else {
		goto L796
	}
L796:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L797:
	;
	v2535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(v2535)+20))
	v2537 = F_MemoryContextAlloc(m, v2536, v2303)
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L9
	} else {
		goto L798
	}
L798:
	;
	v2541 = F__emscripten_memset_bulkmem(m, v2537, base.I32_extend8_s(int32(1)), v2303)
	mBase = m.M
	goto L799
L799:
	;
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_tuplestore_putvalues(m, v2542, v2302, v2533, v2541)
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L9
	} else {
		goto L800
	}
L800:
	;
	goto L722
L801:
	;
	F_SPI_freetuptable(m, v2551)
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L9
	} else {
		goto L804
	}
L802:
	;
	goto L803
L803:
	;
	v2554 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2554
	v2557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2557 == v2554 {
		v5234 = v2554
		goto L21
	} else {
		goto L805
	}
L804:
	;
	goto L803
L805:
	;
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v2557)+20))
	F_MemoryContextReset(m, v2560)
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L9
	} else {
		goto L806
	}
L806:
	;
	v5234 = v2554
	goto L21
L807:
	;
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v2571 = F_AllocSetContextCreateInternal(m, v2566, int32(504671), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		goto L9
	} else {
		goto L810
	}
L808:
	;
	v2574 = v2563
	goto L809
L809:
	;
	v2575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v2575 == int32(0) {
		goto L62
	} else {
		goto L811
	}
L810:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v2571
	v2574 = v2571
	goto L809
L811:
	;
	v2578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2578 != 0 {
		goto L812
	} else {
		goto L813
	}
L812:
	;
	v2582 = v2578
	goto L814
L813:
	;
	F_exec_init_tuple_store(m, l0)
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L9
	} else {
		goto L815
	}
L814:
	;
	v2583 = *(*int64)(unsafe.Add(mBase, uint32(v2582)+40))
	v2584 = int32(4515600)
	v2585 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2574
	v2589 = F_CreateDestReceiver(m, int32(6))
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L9
	} else {
		goto L816
	}
L815:
	;
	v2581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v2582 = v2581
	goto L814
L816:
	;
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2593 = int32(0)
	v2594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v2589)+36)) = int32(368154)
	*(*int32)(unsafe.Add(mBase, uint32(v2589)+32)) = v2594
	*(*uint8)(unsafe.Add(mBase, uint32(v2589)+28)) = uint8(v2593)
	*(*int32)(unsafe.Add(mBase, uint32(v2589)+24)) = v2592
	*(*int32)(unsafe.Add(mBase, uint32(v2589)+20)) = v2591
	goto L817
L817:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2585
	v2603 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v2603 != 0 {
		goto L819
	} else {
		goto L820
	}
L818:
	;
	v2725 = *(*int32)(unsafe.Add(mBase, uint32(v2589)+12))
	m.T0[v2725].(func(*base.Module, int32))(m, v2589)
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L9
	} else {
		goto L852
	}
L819:
	;
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v2603)+24))
	if v2604 == int32(0) {
		goto L822
	} else {
		goto L823
	}
L820:
	;
	goto L821
L821:
	;
	v2657 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v2664 = F_exec_eval_expr(m, l0, v2657, v33+int32(607), v33+int32(636), v33+int32(600))
	mBase = m.M
	v2665 = m.ExcPending
	if v2665 != 0 {
		goto L9
	} else {
		goto L836
	}
L822:
	;
	F_exec_prepare_plan(m, l0, v2603, int32(2048))
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L9
	} else {
		goto L825
	}
L823:
	;
	goto L824
L824:
	;
	v2610 = *(*int32)(unsafe.Add(mBase, uint32(v2603)+28))
	if v2610 == int32(0) {
		goto L827
	} else {
		goto L828
	}
L825:
	;
	goto L824
L826:
	;
	v2618 = v33 + int32(624)
	v2619 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2618))) = v2619
	*(*int64)(unsafe.Add(mBase, uint32(v33)+616)) = v2619
	*(*int64)(unsafe.Add(mBase, uint32(v33)+608)) = v2619
	*(*int32)(unsafe.Add(mBase, uint32(v33)+608)) = v2616
	v2626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	*(*int32)(unsafe.Add(mBase, uint32(v2618))) = v2589
	v2628 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+614)) = uint8(v2628)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+612)) = uint8(v2626)
	v2631 = *(*int32)(unsafe.Add(mBase, uint32(v2603)+24))
	v2634 = F_SPI_execute_plan_extended(m, v2631, v33+int32(608))
	mBase = m.M
	v2635 = m.ExcPending
	if v2635 != 0 {
		goto L9
	} else {
		goto L830
	}
L827:
	;
	v2616 = int32(0)
	goto L826
L828:
	;
	goto L829
L829:
	;
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v2614)+20)) = v2603
	v2616 = v2614
	goto L826
L830:
	;
	if int32(0) <= v2634 {
		goto L818
	} else {
		goto L831
	}
L831:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L9
	} else {
		goto L832
	}
L832:
	;
	v2642 = *(*int32)(unsafe.Add(mBase, uint32(v2603)))
	v2643 = F_SPI_result_code_string(m, v2634)
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L9
	} else {
		goto L833
	}
L833:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v2643
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v2642
	F_errmsg_internal(m, int32(205222), v33+int32(208))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L9
	} else {
		goto L834
	}
L834:
	;
	F_errfinish(m, int32(500315), int32(3608), int32(15796))
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L9
	} else {
		goto L835
	}
L835:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L836:
	;
	v2666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+607)))
	if v2666 == int32(1) {
		goto L61
	} else {
		goto L837
	}
L837:
	;
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v2670 = int32(4515600)
	v2671 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(v2673)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2674
	F_getTypeOutputInfo(m, v2669, v33+int32(608), v33+int32(596))
	mBase = m.M
	v2681 = m.ExcPending
	if v2681 != 0 {
		goto L9
	} else {
		goto L838
	}
L838:
	;
	v2682 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v2683 = F_OidOutputFunctionCall(m, v2682, v2664)
	mBase = m.M
	v2684 = m.ExcPending
	if v2684 != 0 {
		goto L9
	} else {
		goto L839
	}
L839:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2671
	v2687 = F_MemoryContextStrdup(m, v2574, v2683)
	mBase = m.M
	v2688 = m.ExcPending
	if v2688 != 0 {
		goto L9
	} else {
		goto L840
	}
L840:
	;
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2689 != 0 {
		goto L841
	} else {
		goto L842
	}
L841:
	;
	F_SPI_freetuptable(m, v2689)
	mBase = m.M
	v2691 = m.ExcPending
	if v2691 != 0 {
		goto L9
	} else {
		goto L844
	}
L842:
	;
	goto L843
L843:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2694 != 0 {
		goto L845
	} else {
		goto L846
	}
L844:
	;
	goto L843
L845:
	;
	v2695 = *(*int32)(unsafe.Add(mBase, uint32(v2694)+20))
	F_MemoryContextReset(m, v2695)
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L9
	} else {
		goto L848
	}
L846:
	;
	goto L847
L847:
	;
	v2699 = v33 + int32(624)
	v2700 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2699))) = v2700
	*(*int64)(unsafe.Add(mBase, uint32(v33)+616)) = v2700
	*(*int64)(unsafe.Add(mBase, uint32(v33)+608)) = v2700
	v2706 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2707 = F_exec_eval_using_params(m, l0, v2706)
	mBase = m.M
	v2708 = m.ExcPending
	if v2708 != 0 {
		goto L9
	} else {
		goto L849
	}
L848:
	;
	goto L847
L849:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+608)) = v2707
	v2710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	*(*int32)(unsafe.Add(mBase, uint32(v2699))) = v2589
	v2712 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+614)) = uint8(v2712)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+612)) = uint8(v2710)
	v2717 = F_SPI_execute_extended(m, v2687, v33+int32(608))
	mBase = m.M
	v2718 = m.ExcPending
	if v2718 != 0 {
		goto L9
	} else {
		goto L850
	}
L850:
	;
	if v2717 < int32(0) {
		goto L60
	} else {
		goto L851
	}
L851:
	;
	goto L818
L852:
	;
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2728 != 0 {
		goto L853
	} else {
		goto L854
	}
L853:
	;
	F_SPI_freetuptable(m, v2728)
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L9
	} else {
		goto L856
	}
L854:
	;
	goto L855
L855:
	;
	v2731 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2731
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2734 != 0 {
		goto L857
	} else {
		goto L858
	}
L856:
	;
	goto L855
L857:
	;
	v2735 = *(*int32)(unsafe.Add(mBase, uint32(v2734)+20))
	F_MemoryContextReset(m, v2735)
	mBase = m.M
	v2737 = m.ExcPending
	if v2737 != 0 {
		goto L9
	} else {
		goto L860
	}
L858:
	;
	goto L859
L859:
	;
	F_MemoryContextReset(m, v2574)
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L9
	} else {
		goto L861
	}
L860:
	;
	goto L859
L861:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v2741 = *(*int64)(unsafe.Add(mBase, uint32(v2740)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v2741 - v2583
	v2744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v2749 = *(*int32)(unsafe.Add(mBase, uint32(v2744+v2745<<(uint(int32(2))%32))))
	v2751 = int32(0)
	F_assign_simple_var(m, l0, v2749, base.B2i32(v2583 != v2741), v2751, v2751)
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		goto L9
	} else {
		goto L862
	}
L862:
	;
	v5234 = v2731
	goto L21
L863:
	;
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v2775 == int32(0) {
		goto L872
	} else {
		goto L873
	}
L864:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2756 != 0 {
		goto L863
	} else {
		goto L865
	}
L865:
	;
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	if v2757 != 0 {
		goto L863
	} else {
		goto L866
	}
L866:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2758 != 0 {
		goto L59
	} else {
		goto L867
	}
L867:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v2762 = m.ExcPending
	if v2762 != 0 {
		goto L9
	} else {
		goto L868
	}
L868:
	;
	F_errcode(m, int32(33557120))
	mBase = m.M
	v2765 = m.ExcPending
	if v2765 != 0 {
		goto L9
	} else {
		goto L869
	}
L869:
	;
	F_errmsg(m, int32(219784), int32(0))
	mBase = m.M
	v2769 = m.ExcPending
	if v2769 != 0 {
		goto L9
	} else {
		goto L870
	}
L870:
	;
	F_errfinish(m, int32(500315), int32(3749), int32(361640))
	mBase = m.M
	v2774 = m.ExcPending
	if v2774 != 0 {
		goto L9
	} else {
		goto L871
	}
L871:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L872:
	;
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v2783 = F_AllocSetContextCreateInternal(m, v2778, int32(504671), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L9
	} else {
		goto L875
	}
L873:
	;
	v2787 = v2755
	v2788 = v2775
	goto L874
L874:
	;
	v2789 = int32(0)
	if v2787 == v2789 {
		goto L877
	} else {
		goto L878
	}
L875:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v2783
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v2787 = v2786
	v2788 = v2783
	goto L874
L876:
	;
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v2802 != 0 {
		goto L882
	} else {
		goto L883
	}
L877:
	;
	v2792 = int32(0)
	v2800 = v2792
	v2801 = v2792
	goto L876
L878:
	;
	goto L879
L879:
	;
	v2795 = F_plpgsql_recognize_err_condition(m, v2787, int32(1))
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L9
	} else {
		goto L880
	}
L880:
	;
	v2797 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v2798 = F_MemoryContextStrdup(m, v2788, v2797)
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L9
	} else {
		goto L881
	}
L881:
	;
	v2800 = v2795
	v2801 = v2798
	goto L876
L882:
	;
	v2803 = int32(4515600)
	v2804 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2788
	F_initStringInfo(m, v33+int32(608))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L9
	} else {
		goto L885
	}
L883:
	;
	v2947 = v2789
	goto L884
L884:
	;
	v2965 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	if v2965 == int32(0) {
		goto L921
	} else {
		goto L922
	}
L885:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2804
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	if v2813 != 0 {
		goto L886
	} else {
		goto L887
	}
L886:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v2813)+12))
	v2816 = v2814
	goto L888
L887:
	;
	v2816 = int32(0)
	goto L888
L888:
	;
	v2817 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v2821 = v2817
	v2823 = v2816
	goto L889
L889:
	;
	v2848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2821))))
	if v2848 != int32(37) {
		goto L893
	} else {
		goto L894
	}
L890:
	;
	if v2823 != 0 {
		goto L57
	} else {
		goto L920
	}
L891:
	;
	goto L890
L892:
	;
	v2821 = v2927 + int32(1)
	v2823 = v2929
	goto L889
L893:
	;
	if v2848 == int32(0) {
		goto L891
	} else {
		goto L896
	}
L894:
	;
	goto L895
L895:
	;
	v2858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2821)+1)))
	if v2858 == int32(37) {
		goto L898
	} else {
		goto L899
	}
L896:
	;
	F_appendStringInfoChar(m, v33+int32(608), base.I32_extend8_s(v2848))
	mBase = m.M
	v2857 = m.ExcPending
	if v2857 != 0 {
		goto L9
	} else {
		goto L897
	}
L897:
	;
	v2927 = v2821
	v2929 = v2823
	goto L892
L898:
	;
	F_appendStringInfoChar(m, v33+int32(608), int32(37))
	mBase = m.M
	v2867 = m.ExcPending
	if v2867 != 0 {
		goto L9
	} else {
		goto L901
	}
L899:
	;
	goto L900
L900:
	;
	if v2823 == int32(0) {
		goto L58
	} else {
		goto L902
	}
L901:
	;
	v2927 = v2821 + int32(1)
	v2929 = v2823
	goto L892
L902:
	;
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(v2823)))
	v2877 = F_exec_eval_expr(m, l0, v2870, v33+int32(595), v33+int32(600), v33+int32(596))
	mBase = m.M
	v2878 = m.ExcPending
	if v2878 != 0 {
		goto L9
	} else {
		goto L903
	}
L903:
	;
	v2879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+595)))
	if v2879 != 0 {
		goto L905
	} else {
		goto L906
	}
L904:
	;
	F_appendStringInfoString(m, v33+int32(608), v2899)
	mBase = m.M
	v2904 = m.ExcPending
	if v2904 != 0 {
		goto L9
	} else {
		goto L910
	}
L905:
	;
	v2899 = int32(546551)
	goto L904
L906:
	;
	goto L907
L907:
	;
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v33)+600))
	v2882 = int32(4515600)
	v2883 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2886 = *(*int32)(unsafe.Add(mBase, uint32(v2885)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2886
	F_getTypeOutputInfo(m, v2881, v33+int32(636), v33+int32(607))
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L9
	} else {
		goto L908
	}
L908:
	;
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v2895 = F_OidOutputFunctionCall(m, v2894, v2877)
	mBase = m.M
	v2896 = m.ExcPending
	if v2896 != 0 {
		goto L9
	} else {
		goto L909
	}
L909:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v2883
	v2899 = v2895
	goto L904
L910:
	;
	v2906 = v2823 + int32(4)
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v2908 = *(*int32)(unsafe.Add(mBase, uint32(v2907)+12))
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v2907)+4))
	v2914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2914 != 0 {
		goto L911
	} else {
		goto L912
	}
L911:
	;
	F_SPI_freetuptable(m, v2914)
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L9
	} else {
		goto L914
	}
L912:
	;
	goto L913
L913:
	;
	if base.Ui32(v2906) < base.Ui32(v2908+v2909<<(uint(int32(2))%32)) {
		goto L915
	} else {
		goto L916
	}
L914:
	;
	goto L913
L915:
	;
	v2918 = v2906
	goto L917
L916:
	;
	v2918 = int32(0)
	goto L917
L917:
	;
	v2919 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2919
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2921 == v2919 {
		v2927 = v2821
		v2929 = v2918
		goto L892
	} else {
		goto L918
	}
L918:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2921)+20))
	F_MemoryContextReset(m, v2924)
	mBase = m.M
	v2926 = m.ExcPending
	if v2926 != 0 {
		goto L9
	} else {
		goto L919
	}
L919:
	;
	v2927 = v2821
	v2929 = v2918
	goto L892
L920:
	;
	v2934 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v2947 = v2934
	goto L884
L921:
	;
	v2968 = int32(0)
	v4711 = v2947
	v4712 = v2800
	v4715 = v2801
	v4716 = v2968
	v4717 = v2968
	v4718 = v2968
	v4719 = v2968
	v4720 = v2968
	v4721 = v2968
	v4722 = v2968
	goto L33
L922:
	;
	goto L923
L923:
	;
	v2975 = int32(0)
	v2983 = *(*int32)(unsafe.Add(mBase, uint32(v2965)+4))
	if v2983 <= v2975 {
		v4711 = v2947
		v4712 = v2800
		v4715 = v2801
		v4716 = v2975
		v4717 = v2975
		v4718 = v2975
		v4719 = v2975
		v4720 = v2975
		v4721 = v2975
		v4722 = v2975
		goto L33
	} else {
		goto L924
	}
L924:
	;
	v2989 = v2975
	v2998 = v2947
	v2999 = v2800
	v3002 = v2801
	v3003 = v2975
	v3004 = v2975
	v3005 = v2975
	v3006 = v2975
	v3007 = v2975
	v3008 = v2975
	v3009 = v2975
	goto L925
L925:
	;
	v3016 = *(*int32)(unsafe.Add(mBase, uint32(v2965)+12))
	v3020 = *(*int32)(unsafe.Add(mBase, uint32(v3016+v2989<<(uint(int32(2))%32))))
	v3021 = *(*int32)(unsafe.Add(mBase, uint32(v3020)+4))
	v3028 = F_exec_eval_expr(m, l0, v3021, v33+int32(607), v33+int32(636), v33+int32(600))
	mBase = m.M
	v3029 = m.ExcPending
	if v3029 != 0 {
		goto L9
	} else {
		goto L927
	}
L926:
	;
	v4711 = v3073
	v4712 = v3074
	v4715 = v3075
	v4716 = v3076
	v4717 = v3077
	v4718 = v3078
	v4719 = v3079
	v4720 = v3080
	v4721 = v3081
	v4722 = v3082
	goto L33
L927:
	;
	v3030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+607)))
	if v3030 == int32(1) {
		goto L56
	} else {
		goto L928
	}
L928:
	;
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v3034 = int32(4515600)
	v3035 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v3037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v3038 = *(*int32)(unsafe.Add(mBase, uint32(v3037)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3038
	F_getTypeOutputInfo(m, v3033, v33+int32(608), v33+int32(596))
	mBase = m.M
	v3045 = m.ExcPending
	if v3045 != 0 {
		goto L9
	} else {
		goto L929
	}
L929:
	;
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v3047 = F_OidOutputFunctionCall(m, v3046, v3028)
	mBase = m.M
	v3048 = m.ExcPending
	if v3048 != 0 {
		goto L9
	} else {
		goto L930
	}
L930:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3035
	v3051 = *(*int32)(unsafe.Add(mBase, uint32(v3020)))
	switch v3051 {
	case 0:
		goto L940
	case 1:
		goto L939
	case 2:
		goto L938
	case 3:
		goto L937
	case 4:
		goto L936
	case 5:
		goto L935
	case 6:
		goto L934
	case 7:
		goto L933
	case 8:
		goto L932
	default:
		goto L46
	}
L931:
	;
	v3083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3083 != 0 {
		goto L960
	} else {
		goto L961
	}
L932:
	;
	if v3005 != 0 {
		goto L47
	} else {
		goto L958
	}
L933:
	;
	if v3004 != 0 {
		goto L48
	} else {
		goto L956
	}
L934:
	;
	if v3003 != 0 {
		goto L49
	} else {
		goto L954
	}
L935:
	;
	if v3006 != 0 {
		goto L50
	} else {
		goto L952
	}
L936:
	;
	if v3007 != 0 {
		goto L51
	} else {
		goto L950
	}
L937:
	;
	if v3008 != 0 {
		goto L52
	} else {
		goto L948
	}
L938:
	;
	if v3009 != 0 {
		goto L53
	} else {
		goto L946
	}
L939:
	;
	if v2998 != 0 {
		goto L54
	} else {
		goto L944
	}
L940:
	;
	if v2999 != 0 {
		goto L55
	} else {
		goto L941
	}
L941:
	;
	v3053 = F_plpgsql_recognize_err_condition(m, v3047, int32(1))
	mBase = m.M
	v3054 = m.ExcPending
	if v3054 != 0 {
		goto L9
	} else {
		goto L942
	}
L942:
	;
	v3055 = F_MemoryContextStrdup(m, v2788, v3047)
	mBase = m.M
	v3056 = m.ExcPending
	if v3056 != 0 {
		goto L9
	} else {
		goto L943
	}
L943:
	;
	v3073 = v2998
	v3074 = v3053
	v3075 = v3055
	v3076 = v3003
	v3077 = v3004
	v3078 = v3005
	v3079 = v3006
	v3080 = v3007
	v3081 = v3008
	v3082 = v3009
	goto L931
L944:
	;
	v3057 = F_MemoryContextStrdup(m, v2788, v3047)
	mBase = m.M
	v3058 = m.ExcPending
	if v3058 != 0 {
		goto L9
	} else {
		goto L945
	}
L945:
	;
	v3073 = v3057
	v3074 = v2999
	v3075 = v3002
	v3076 = v3003
	v3077 = v3004
	v3078 = v3005
	v3079 = v3006
	v3080 = v3007
	v3081 = v3008
	v3082 = v3009
	goto L931
L946:
	;
	v3059 = F_MemoryContextStrdup(m, v2788, v3047)
	mBase = m.M
	v3060 = m.ExcPending
	if v3060 != 0 {
		goto L9
	} else {
		goto L947
	}
L947:
	;
	v3073 = v2998
	v3074 = v2999
	v3075 = v3002
	v3076 = v3003
	v3077 = v3004
	v3078 = v3005
	v3079 = v3006
	v3080 = v3007
	v3081 = v3008
	v3082 = v3059
	goto L931
L948:
	;
	v3061 = F_MemoryContextStrdup(m, v2788, v3047)
	mBase = m.M
	v3062 = m.ExcPending
	if v3062 != 0 {
		goto L9
	} else {
		goto L949
	}
L949:
	;
	v3073 = v2998
	v3074 = v2999
	v3075 = v3002
	v3076 = v3003
	v3077 = v3004
	v3078 = v3005
	v3079 = v3006
	v3080 = v3007
	v3081 = v3061
	v3082 = v3009
	goto L931
L950:
	;
	v3063 = F_MemoryContextStrdup(m, v2788, v3047)
	mBase = m.M
	v3064 = m.ExcPending
	if v3064 != 0 {
		goto L9
	} else {
		goto L951
	}
L951:
	;
	v3073 = v2998
	v3074 = v2999
	v3075 = v3002
	v3076 = v3003
	v3077 = v3004
	v3078 = v3005
	v3079 = v3006
	v3080 = v3063
	v3081 = v3008
	v3082 = v3009
	goto L931
L952:
	;
	v3065 = F_MemoryContextStrdup(m, v2788, v3047)
	mBase = m.M
	v3066 = m.ExcPending
	if v3066 != 0 {
		goto L9
	} else {
		goto L953
	}
L953:
	;
	v3073 = v2998
	v3074 = v2999
	v3075 = v3002
	v3076 = v3003
	v3077 = v3004
	v3078 = v3005
	v3079 = v3065
	v3080 = v3007
	v3081 = v3008
	v3082 = v3009
	goto L931
L954:
	;
	v3067 = F_MemoryContextStrdup(m, v2788, v3047)
	mBase = m.M
	v3068 = m.ExcPending
	if v3068 != 0 {
		goto L9
	} else {
		goto L955
	}
L955:
	;
	v3073 = v2998
	v3074 = v2999
	v3075 = v3002
	v3076 = v3067
	v3077 = v3004
	v3078 = v3005
	v3079 = v3006
	v3080 = v3007
	v3081 = v3008
	v3082 = v3009
	goto L931
L956:
	;
	v3069 = F_MemoryContextStrdup(m, v2788, v3047)
	mBase = m.M
	v3070 = m.ExcPending
	if v3070 != 0 {
		goto L9
	} else {
		goto L957
	}
L957:
	;
	v3073 = v2998
	v3074 = v2999
	v3075 = v3002
	v3076 = v3003
	v3077 = v3069
	v3078 = v3005
	v3079 = v3006
	v3080 = v3007
	v3081 = v3008
	v3082 = v3009
	goto L931
L958:
	;
	v3071 = F_MemoryContextStrdup(m, v2788, v3047)
	mBase = m.M
	v3072 = m.ExcPending
	if v3072 != 0 {
		goto L9
	} else {
		goto L959
	}
L959:
	;
	v3073 = v2998
	v3074 = v2999
	v3075 = v3002
	v3076 = v3003
	v3077 = v3004
	v3078 = v3071
	v3079 = v3006
	v3080 = v3007
	v3081 = v3008
	v3082 = v3009
	goto L931
L960:
	;
	F_SPI_freetuptable(m, v3083)
	mBase = m.M
	v3085 = m.ExcPending
	if v3085 != 0 {
		goto L9
	} else {
		goto L963
	}
L961:
	;
	goto L962
L962:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v3088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3088 != 0 {
		goto L964
	} else {
		goto L965
	}
L963:
	;
	goto L962
L964:
	;
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v3088)+20))
	F_MemoryContextReset(m, v3089)
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L9
	} else {
		goto L967
	}
L965:
	;
	goto L966
L966:
	;
	v3093 = v2989 + int32(1)
	v3094 = *(*int32)(unsafe.Add(mBase, uint32(v2965)+4))
	if v3093 < v3094 {
		v2989 = v3093
		v2998 = v3073
		v2999 = v3074
		v3002 = v3075
		v3003 = v3076
		v3004 = v3077
		v3005 = v3078
		v3006 = v3079
		v3007 = v3080
		v3008 = v3081
		v3009 = v3082
		goto L925
	} else {
		goto L968
	}
L967:
	;
	goto L966
L968:
	;
	goto L926
L969:
	;
	v5234 = int32(0)
	goto L21
L970:
	;
	v3100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v3107 = F_exec_eval_expr(m, l0, v3100, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v3108 = m.ExcPending
	if v3108 != 0 {
		goto L9
	} else {
		goto L971
	}
L971:
	;
	v3111 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v3112 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v3115 = F_exec_cast_value(m, l0, v3107, v33+int32(600), v3111, v3112, int32(16), int32(-1))
	mBase = m.M
	v3116 = m.ExcPending
	if v3116 != 0 {
		goto L9
	} else {
		goto L972
	}
L972:
	;
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3117 != 0 {
		goto L973
	} else {
		goto L974
	}
L973:
	;
	F_SPI_freetuptable(m, v3117)
	mBase = m.M
	v3119 = m.ExcPending
	if v3119 != 0 {
		goto L9
	} else {
		goto L976
	}
L974:
	;
	goto L975
L975:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3122 != 0 {
		goto L977
	} else {
		goto L978
	}
L976:
	;
	goto L975
L977:
	;
	v3123 = *(*int32)(unsafe.Add(mBase, uint32(v3122)+20))
	F_MemoryContextReset(m, v3123)
	mBase = m.M
	v3125 = m.ExcPending
	if v3125 != 0 {
		goto L9
	} else {
		goto L980
	}
L978:
	;
	goto L979
L979:
	;
	v3126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v3115 != 0 {
		goto L981
	} else {
		goto L982
	}
L980:
	;
	goto L979
L981:
	;
	v3128 = v3126
	goto L983
L982:
	;
	v3128 = int32(1)
	goto L983
L983:
	;
	if v3128 == int32(0) {
		goto L969
	} else {
		goto L984
	}
L984:
	;
	v3131 = int32(0)
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	if v3132 == v3131 {
		v3171 = v3131
		goto L985
	} else {
		goto L986
	}
L985:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3178 = m.ExcPending
	if v3178 != 0 {
		goto L9
	} else {
		goto L991
	}
L986:
	;
	v3141 = F_exec_eval_expr(m, l0, v3132, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v3142 = m.ExcPending
	if v3142 != 0 {
		goto L9
	} else {
		goto L987
	}
L987:
	;
	v3143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v3143 != 0 {
		v3171 = v3131
		goto L985
	} else {
		goto L988
	}
L988:
	;
	v3144 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v3145 = m.G0
	v3147 = v3145 - int32(16)
	m.G0 = v3147
	v3149 = int32(4515600)
	v3150 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v3152)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3153
	F_getTypeOutputInfo(m, v3144, v3147+int32(12), v3147+int32(11))
	mBase = m.M
	v3160 = m.ExcPending
	if v3160 != 0 {
		goto L9
	} else {
		goto L989
	}
L989:
	;
	v3161 = *(*int32)(unsafe.Add(mBase, uint32(v3147)+12))
	v3162 = F_OidOutputFunctionCall(m, v3161, v3141)
	mBase = m.M
	v3163 = m.ExcPending
	if v3163 != 0 {
		goto L9
	} else {
		goto L990
	}
L990:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3150
	m.G0 = v3147 + int32(16)
	v3171 = v3162
	goto L985
L991:
	;
	F_errcode(m, int32(67108896))
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		goto L9
	} else {
		goto L992
	}
L992:
	;
	if v3171 != 0 {
		goto L994
	} else {
		goto L995
	}
L993:
	;
	F_errfinish(m, int32(500315), int32(3968), int32(81505))
	mBase = m.M
	v3196 = m.ExcPending
	if v3196 != 0 {
		goto L9
	} else {
		goto L999
	}
L994:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+432)) = v3171
	F_errmsg_internal(m, int32(206200), v33+int32(432))
	mBase = m.M
	v3187 = m.ExcPending
	if v3187 != 0 {
		goto L9
	} else {
		goto L997
	}
L995:
	;
	goto L996
L996:
	;
	F_errmsg(m, int32(454416), int32(0))
	mBase = m.M
	v3191 = m.ExcPending
	if v3191 != 0 {
		goto L9
	} else {
		goto L998
	}
L997:
	;
	goto L993
L998:
	;
	goto L993
L999:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1000:
	;
	v5234 = int32(0)
	goto L21
L1001:
	;
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3211 = F_AllocSetContextCreateInternal(m, v3206, int32(504671), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v3212 = m.ExcPending
	if v3212 != 0 {
		goto L9
	} else {
		goto L1004
	}
L1002:
	;
	v3214 = v3203
	goto L1003
L1003:
	;
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v3222 = F_exec_eval_expr(m, l0, v3215, v33+int32(607), v33+int32(636), v33+int32(600))
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		goto L9
	} else {
		goto L1005
	}
L1004:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v3211
	v3214 = v3211
	goto L1003
L1005:
	;
	v3224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+607)))
	if v3224 == int32(1) {
		goto L45
	} else {
		goto L1006
	}
L1006:
	;
	v3227 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v3228 = int32(4515600)
	v3229 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v3231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v3232 = *(*int32)(unsafe.Add(mBase, uint32(v3231)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3232
	F_getTypeOutputInfo(m, v3227, v33+int32(608), v33+int32(596))
	mBase = m.M
	v3239 = m.ExcPending
	if v3239 != 0 {
		goto L9
	} else {
		goto L1007
	}
L1007:
	;
	v3240 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v3241 = F_OidOutputFunctionCall(m, v3240, v3222)
	mBase = m.M
	v3242 = m.ExcPending
	if v3242 != 0 {
		goto L9
	} else {
		goto L1008
	}
L1008:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3229
	v3245 = F_MemoryContextStrdup(m, v3214, v3241)
	mBase = m.M
	v3246 = m.ExcPending
	if v3246 != 0 {
		goto L9
	} else {
		goto L1009
	}
L1009:
	;
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3247 != 0 {
		goto L1010
	} else {
		goto L1011
	}
L1010:
	;
	F_SPI_freetuptable(m, v3247)
	mBase = m.M
	v3249 = m.ExcPending
	if v3249 != 0 {
		goto L9
	} else {
		goto L1013
	}
L1011:
	;
	goto L1012
L1012:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v3252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3252 != 0 {
		goto L1014
	} else {
		goto L1015
	}
L1013:
	;
	goto L1012
L1014:
	;
	v3253 = *(*int32)(unsafe.Add(mBase, uint32(v3252)+20))
	F_MemoryContextReset(m, v3253)
	mBase = m.M
	v3255 = m.ExcPending
	if v3255 != 0 {
		goto L9
	} else {
		goto L1017
	}
L1015:
	;
	goto L1016
L1016:
	;
	v3256 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v3257 = F_exec_eval_using_params(m, l0, v3256)
	mBase = m.M
	v3258 = m.ExcPending
	if v3258 != 0 {
		goto L9
	} else {
		goto L1018
	}
L1017:
	;
	goto L1016
L1018:
	;
	v3259 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+624)) = v3259
	*(*int64)(unsafe.Add(mBase, uint32(v33)+616)) = v3259
	*(*int64)(unsafe.Add(mBase, uint32(v33)+608)) = v3259
	*(*int32)(unsafe.Add(mBase, uint32(v33)+608)) = v3257
	v3266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+612)) = uint8(v3266)
	v3270 = F_SPI_execute_extended(m, v3245, v33+int32(608))
	mBase = m.M
	v3271 = m.ExcPending
	if v3271 != 0 {
		goto L9
	} else {
		goto L1024
	}
L1019:
	;
	v3345 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v3345
	v3347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+16)))
	if v3347 != int32(1) {
		goto L1042
	} else {
		goto L1043
	}
L1020:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3329 = m.ExcPending
	if v3329 != 0 {
		goto L9
	} else {
		goto L1038
	}
L1021:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3313 = m.ExcPending
	if v3313 != 0 {
		goto L9
	} else {
		goto L1034
	}
L1022:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3297 = m.ExcPending
	if v3297 != 0 {
		goto L9
	} else {
		goto L1030
	}
L1023:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3277 = m.ExcPending
	if v3277 != 0 {
		goto L9
	} else {
		goto L1025
	}
L1024:
	;
	switch v3270 + int32(8) {
	case 0:
		goto L1021
	default:
		goto L1020
	case 6:
		goto L1022
	case 8, 12, 13, 15, 16, 17, 19, 20, 21, 22, 26, 27:
		goto L1019
	case 14:
		goto L1023
	}
L1025:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L9
	} else {
		goto L1026
	}
L1026:
	;
	F_errmsg(m, int32(446056), int32(0))
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L9
	} else {
		goto L1027
	}
L1027:
	;
	F_errhint(m, int32(652552), int32(0))
	mBase = m.M
	v3288 = m.ExcPending
	if v3288 != 0 {
		goto L9
	} else {
		goto L1028
	}
L1028:
	;
	F_errfinish(m, int32(500315), int32(4517), int32(348903))
	mBase = m.M
	v3293 = m.ExcPending
	if v3293 != 0 {
		goto L9
	} else {
		goto L1029
	}
L1029:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1030:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3300 = m.ExcPending
	if v3300 != 0 {
		goto L9
	} else {
		goto L1031
	}
L1031:
	;
	F_errmsg(m, int32(532418), int32(0))
	mBase = m.M
	v3304 = m.ExcPending
	if v3304 != 0 {
		goto L9
	} else {
		goto L1032
	}
L1032:
	;
	F_errfinish(m, int32(500315), int32(4524), int32(348903))
	mBase = m.M
	v3309 = m.ExcPending
	if v3309 != 0 {
		goto L9
	} else {
		goto L1033
	}
L1033:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1034:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3316 = m.ExcPending
	if v3316 != 0 {
		goto L9
	} else {
		goto L1035
	}
L1035:
	;
	F_errmsg(m, int32(445770), int32(0))
	mBase = m.M
	v3320 = m.ExcPending
	if v3320 != 0 {
		goto L9
	} else {
		goto L1036
	}
L1036:
	;
	F_errfinish(m, int32(500315), int32(4530), int32(348903))
	mBase = m.M
	v3325 = m.ExcPending
	if v3325 != 0 {
		goto L9
	} else {
		goto L1037
	}
L1037:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1038:
	;
	v3330 = F_SPI_result_code_string(m, v3270)
	mBase = m.M
	v3331 = m.ExcPending
	if v3331 != 0 {
		goto L9
	} else {
		goto L1039
	}
L1039:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+452)) = v3330
	*(*int32)(unsafe.Add(mBase, uint32(v33)+448)) = v3245
	F_errmsg_internal(m, int32(205280), v33+int32(448))
	mBase = m.M
	v3338 = m.ExcPending
	if v3338 != 0 {
		goto L9
	} else {
		goto L1040
	}
L1040:
	;
	F_errfinish(m, int32(500315), int32(4535), int32(348903))
	mBase = m.M
	v3343 = m.ExcPending
	if v3343 != 0 {
		goto L9
	} else {
		goto L1041
	}
L1041:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1042:
	;
	v3453 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	F_SPI_freetuptable(m, v3453)
	mBase = m.M
	v3455 = m.ExcPending
	if v3455 != 0 {
		goto L9
	} else {
		goto L1084
	}
L1043:
	;
	v3351 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	if v3351 == int32(0) {
		goto L44
	} else {
		goto L1044
	}
L1044:
	;
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v3356 = *(*int32)(unsafe.Add(mBase, uint32(v3355)+4))
	v3360 = *(*int32)(unsafe.Add(mBase, uint32(v3354+v3356<<(uint(int32(2))%32))))
	if base.Ui64(v3345) <= base.Ui64(int64(1)) {
		goto L1047
	} else {
		goto L1048
	}
L1045:
	;
	v3436 = *(*int32)(unsafe.Add(mBase, uint32(v3351)))
	F_exec_move_row(m, l0, v3360, v3435, v3436)
	mBase = m.M
	v3438 = m.ExcPending
	if v3438 != 0 {
		goto L9
	} else {
		goto L1077
	}
L1046:
	;
	v3433 = *(*int32)(unsafe.Add(mBase, uint32(v3351)+4))
	v3434 = *(*int32)(unsafe.Add(mBase, uint32(v3433)))
	v3435 = v3434
	goto L1045
L1047:
	;
	if base.I32_wrap_i64(v3345) == int32(1) {
		goto L1046
	} else {
		goto L1050
	}
L1048:
	;
	goto L1049
L1049:
	;
	v3400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+17)))
	if v3400 != int32(1) {
		goto L1046
	} else {
		goto L1064
	}
L1050:
	;
	v3367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+17)))
	if v3367 != int32(1) {
		v3435 = int32(0)
		goto L1045
	} else {
		goto L1051
	}
L1051:
	;
	v3371 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3371)+492)))
	if v3372 == int32(1) {
		goto L1052
	} else {
		goto L1053
	}
L1052:
	;
	v3375 = F_format_preparedparamsdata(m, l0, v3257)
	mBase = m.M
	v3376 = m.ExcPending
	if v3376 != 0 {
		goto L9
	} else {
		goto L1055
	}
L1053:
	;
	v3377 = int32(0)
	goto L1054
L1054:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3381 = m.ExcPending
	if v3381 != 0 {
		goto L9
	} else {
		goto L1056
	}
L1055:
	;
	v3377 = v3375
	goto L1054
L1056:
	;
	F_errcode(m, int32(33554464))
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L9
	} else {
		goto L1057
	}
L1057:
	;
	F_errmsg(m, int32(113700), int32(0))
	mBase = m.M
	v3388 = m.ExcPending
	if v3388 != 0 {
		goto L9
	} else {
		goto L1058
	}
L1058:
	;
	if v3377 != 0 {
		goto L1059
	} else {
		goto L1060
	}
L1059:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+480)) = v3377
	F_errdetail_internal(m, int32(200210), v33+int32(480))
	mBase = m.M
	v3394 = m.ExcPending
	if v3394 != 0 {
		goto L9
	} else {
		goto L1062
	}
L1060:
	;
	goto L1061
L1061:
	;
	F_errfinish(m, int32(500315), int32(4577), int32(348903))
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L9
	} else {
		goto L1063
	}
L1062:
	;
	goto L1061
L1063:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1064:
	;
	v3404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3404)+492)))
	if v3405 == int32(1) {
		goto L1065
	} else {
		goto L1066
	}
L1065:
	;
	v3408 = F_format_preparedparamsdata(m, l0, v3257)
	mBase = m.M
	v3409 = m.ExcPending
	if v3409 != 0 {
		goto L9
	} else {
		goto L1068
	}
L1066:
	;
	v3410 = int32(0)
	goto L1067
L1067:
	;
	F_errstart_cold(m, int32(21), int32(556508))
	mBase = m.M
	v3414 = m.ExcPending
	if v3414 != 0 {
		goto L9
	} else {
		goto L1069
	}
L1068:
	;
	v3410 = v3408
	goto L1067
L1069:
	;
	F_errcode(m, int32(50331680))
	mBase = m.M
	v3417 = m.ExcPending
	if v3417 != 0 {
		goto L9
	} else {
		goto L1070
	}
L1070:
	;
	F_errmsg(m, int32(30502), int32(0))
	mBase = m.M
	v3421 = m.ExcPending
	if v3421 != 0 {
		goto L9
	} else {
		goto L1071
	}
L1071:
	;
	if v3410 != 0 {
		goto L1072
	} else {
		goto L1073
	}
L1072:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+464)) = v3410
	F_errdetail_internal(m, int32(200210), v33+int32(464))
	mBase = m.M
	v3427 = m.ExcPending
	if v3427 != 0 {
		goto L9
	} else {
		goto L1075
	}
L1073:
	;
	goto L1074
L1074:
	;
	F_errfinish(m, int32(500315), int32(4596), int32(348903))
	mBase = m.M
	v3432 = m.ExcPending
	if v3432 != 0 {
		goto L9
	} else {
		goto L1076
	}
L1075:
	;
	goto L1074
L1076:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1077:
	;
	v3439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3439 != 0 {
		goto L1078
	} else {
		goto L1079
	}
L1078:
	;
	F_SPI_freetuptable(m, v3439)
	mBase = m.M
	v3441 = m.ExcPending
	if v3441 != 0 {
		goto L9
	} else {
		goto L1081
	}
L1079:
	;
	goto L1080
L1080:
	;
	v3442 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v3442
	v3444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3444 == v3442 {
		goto L1042
	} else {
		goto L1082
	}
L1081:
	;
	goto L1080
L1082:
	;
	v3447 = *(*int32)(unsafe.Add(mBase, uint32(v3444)+20))
	F_MemoryContextReset(m, v3447)
	mBase = m.M
	v3449 = m.ExcPending
	if v3449 != 0 {
		goto L9
	} else {
		goto L1083
	}
L1083:
	;
	goto L1042
L1084:
	;
	F_MemoryContextReset(m, v3214)
	mBase = m.M
	v3457 = m.ExcPending
	if v3457 != 0 {
		goto L9
	} else {
		goto L1085
	}
L1085:
	;
	v5234 = int32(0)
	goto L21
L1086:
	;
	v3466 = F_exec_for_query(m, l0, v91, v3463, int32(1))
	mBase = m.M
	v3467 = m.ExcPending
	if v3467 != 0 {
		goto L9
	} else {
		goto L1087
	}
L1087:
	;
	F_SPI_cursor_close(m, v3463)
	mBase = m.M
	v3469 = m.ExcPending
	if v3469 != 0 {
		goto L9
	} else {
		goto L1088
	}
L1088:
	;
	v5234 = v3466
	goto L21
L1089:
	;
	v3481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v3481 == int32(0) {
		goto L1092
	} else {
		goto L1093
	}
L1090:
	;
	v3504 = v3470
	v3505 = v3470
	goto L1091
L1091:
	;
	v3507 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	if v3507 != 0 {
		goto L1100
	} else {
		goto L1101
	}
L1092:
	;
	v3484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v3489 = F_AllocSetContextCreateInternal(m, v3484, int32(504671), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v3490 = m.ExcPending
	if v3490 != 0 {
		goto L9
	} else {
		goto L1095
	}
L1093:
	;
	v3492 = v3481
	goto L1094
L1094:
	;
	v3493 = int32(4515600)
	v3494 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3492
	v3497 = *(*int32)(unsafe.Add(mBase, uint32(v3477)+40))
	v3498 = F_text_to_cstring(m, v3497)
	mBase = m.M
	v3499 = m.ExcPending
	if v3499 != 0 {
		goto L9
	} else {
		goto L1096
	}
L1095:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v3489
	v3492 = v3489
	goto L1094
L1096:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3494
	v3502 = F_GetPortalByName(m, v3498)
	mBase = m.M
	v3503 = m.ExcPending
	if v3503 != 0 {
		goto L9
	} else {
		goto L1097
	}
L1097:
	;
	if v3502 != 0 {
		goto L43
	} else {
		goto L1098
	}
L1098:
	;
	v3504 = v3498
	v3505 = v3492
	goto L1091
L1099:
	;
	v3567 = *(*int32)(unsafe.Add(mBase, uint32(v3564)+28))
	if v3567 == int32(0) {
		goto L1125
	} else {
		goto L1126
	}
L1100:
	;
	v3508 = *(*int32)(unsafe.Add(mBase, uint32(v3507)+24))
	if v3508 != 0 {
		v3564 = v3507
		goto L1099
	} else {
		goto L1103
	}
L1101:
	;
	goto L1102
L1102:
	;
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	if v3512 != 0 {
		goto L1105
	} else {
		goto L1106
	}
L1103:
	;
	v3509 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	F_exec_prepare_plan(m, l0, v3507, v3509)
	mBase = m.M
	v3511 = m.ExcPending
	if v3511 != 0 {
		goto L9
	} else {
		goto L1104
	}
L1104:
	;
	v3564 = v3507
	goto L1099
L1105:
	;
	v3513 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v3514 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v3515 = F_exec_dynquery_with_params(m, l0, v3512, v3513, v3504, v3514)
	mBase = m.M
	v3516 = m.ExcPending
	if v3516 != 0 {
		goto L9
	} else {
		goto L1108
	}
L1106:
	;
	goto L1107
L1107:
	;
	v3529 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if v3529 != 0 {
		goto L1116
	} else {
		goto L1117
	}
L1108:
	;
	if v3504 != 0 {
		goto L1109
	} else {
		goto L1110
	}
L1109:
	;
	v5234 = int32(0)
	goto L21
L1110:
	;
	goto L1111
L1111:
	;
	v3518 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	F_exec_check_assignable(m, l0, v3518)
	mBase = m.M
	v3520 = m.ExcPending
	if v3520 != 0 {
		goto L9
	} else {
		goto L1112
	}
L1112:
	;
	v3522 = *(*int32)(unsafe.Add(mBase, uint32(v3515)))
	v3523 = F_cstring_to_text(m, v3522)
	mBase = m.M
	v3524 = m.ExcPending
	if v3524 != 0 {
		goto L9
	} else {
		goto L1113
	}
L1113:
	;
	F_assign_simple_var(m, l0, v3477, v3523, int32(0), int32(1))
	mBase = m.M
	v3528 = m.ExcPending
	if v3528 != 0 {
		goto L9
	} else {
		goto L1114
	}
L1114:
	;
	v5234 = int32(0)
	goto L21
L1115:
	;
	v3559 = *(*int32)(unsafe.Add(mBase, uint32(v3477)+28))
	v3560 = *(*int32)(unsafe.Add(mBase, uint32(v3559)+24))
	if v3560 != 0 {
		v3564 = v3559
		goto L1099
	} else {
		goto L1122
	}
L1116:
	;
	v3530 = *(*int32)(unsafe.Add(mBase, uint32(v3477)+32))
	if v3530 < int32(0) {
		goto L42
	} else {
		goto L1119
	}
L1117:
	;
	goto L1118
L1118:
	;
	v3554 = *(*int32)(unsafe.Add(mBase, uint32(v3477)+32))
	if int32(0) <= v3554 {
		goto L41
	} else {
		goto L1121
	}
L1119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+608)) = int32(16)
	v3539 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v3540 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+626)) = uint8(v3540)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+620)) = v3529
	*(*int32)(unsafe.Add(mBase, uint32(v33)+612)) = v3539
	v3544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3548 = *(*int32)(unsafe.Add(mBase, uint32(v3544+v3530<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+628)) = v3548
	F_exec_stmt_execsql(m, l0, v33+int32(608))
	mBase = m.M
	v3553 = m.ExcPending
	if v3553 != 0 {
		goto L9
	} else {
		goto L1120
	}
L1120:
	;
	goto L1115
L1121:
	;
	goto L1115
L1122:
	;
	v3561 = *(*int32)(unsafe.Add(mBase, uint32(v3477)+36))
	F_exec_prepare_plan(m, l0, v3559, v3561)
	mBase = m.M
	v3563 = m.ExcPending
	if v3563 != 0 {
		goto L9
	} else {
		goto L1123
	}
L1123:
	;
	v3564 = v3559
	goto L1099
L1124:
	;
	v3574 = *(*int32)(unsafe.Add(mBase, uint32(v3564)+24))
	v3575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	v3576 = F_SPI_cursor_open_with_paramlist(m, v3504, v3574, v3573, v3575)
	mBase = m.M
	v3577 = m.ExcPending
	if v3577 != 0 {
		goto L9
	} else {
		goto L1128
	}
L1125:
	;
	v3573 = int32(0)
	goto L1124
L1126:
	;
	goto L1127
L1127:
	;
	v3571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3571)+20)) = v3564
	v3573 = v3571
	goto L1124
L1128:
	;
	if v3576 == int32(0) {
		goto L40
	} else {
		goto L1129
	}
L1129:
	;
	if v3504 == int32(0) {
		goto L1130
	} else {
		goto L1131
	}
L1130:
	;
	v3582 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	F_exec_check_assignable(m, l0, v3582)
	mBase = m.M
	v3584 = m.ExcPending
	if v3584 != 0 {
		goto L9
	} else {
		goto L1133
	}
L1131:
	;
	goto L1132
L1132:
	;
	v3592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3592 != 0 {
		goto L1136
	} else {
		goto L1137
	}
L1133:
	;
	v3585 = *(*int32)(unsafe.Add(mBase, uint32(v3576)))
	v3586 = F_cstring_to_text(m, v3585)
	mBase = m.M
	v3587 = m.ExcPending
	if v3587 != 0 {
		goto L9
	} else {
		goto L1134
	}
L1134:
	;
	F_assign_simple_var(m, l0, v3477, v3586, int32(0), int32(1))
	mBase = m.M
	v3591 = m.ExcPending
	if v3591 != 0 {
		goto L9
	} else {
		goto L1135
	}
L1135:
	;
	goto L1132
L1136:
	;
	F_SPI_freetuptable(m, v3592)
	mBase = m.M
	v3594 = m.ExcPending
	if v3594 != 0 {
		goto L9
	} else {
		goto L1139
	}
L1137:
	;
	goto L1138
L1138:
	;
	v3595 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v3595
	v3598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3598 != 0 {
		goto L1140
	} else {
		goto L1141
	}
L1139:
	;
	goto L1138
L1140:
	;
	v3599 = *(*int32)(unsafe.Add(mBase, uint32(v3598)+20))
	F_MemoryContextReset(m, v3599)
	mBase = m.M
	v3601 = m.ExcPending
	if v3601 != 0 {
		goto L9
	} else {
		goto L1143
	}
L1141:
	;
	goto L1142
L1142:
	;
	if v3505 == int32(0) {
		v5234 = v3595
		goto L21
	} else {
		goto L1144
	}
L1143:
	;
	goto L1142
L1144:
	;
	F_MemoryContextReset(m, v3505)
	mBase = m.M
	v3605 = m.ExcPending
	if v3605 != 0 {
		goto L9
	} else {
		goto L1145
	}
L1145:
	;
	v5234 = v3595
	goto L21
L1146:
	;
	v3615 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v3616 = int32(4515600)
	v3617 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v3619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v3620 = *(*int32)(unsafe.Add(mBase, uint32(v3619)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3620
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v3611)+40))
	v3623 = F_text_to_cstring(m, v3622)
	mBase = m.M
	v3624 = m.ExcPending
	if v3624 != 0 {
		goto L9
	} else {
		goto L1147
	}
L1147:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3617
	v3627 = F_GetPortalByName(m, v3623)
	mBase = m.M
	v3628 = m.ExcPending
	if v3628 != 0 {
		goto L9
	} else {
		goto L1148
	}
L1148:
	;
	if v3627 == int32(0) {
		goto L38
	} else {
		goto L1149
	}
L1149:
	;
	v3631 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	if v3631 == int32(0) {
		v3664 = v3615
		goto L1150
	} else {
		goto L1151
	}
L1150:
	;
	v3666 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v3667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+32)))
	if v3667 == int32(0) {
		goto L1162
	} else {
		goto L1163
	}
L1151:
	;
	v3640 = F_exec_eval_expr(m, l0, v3631, v33+int32(600), v33+int32(608), v33+int32(636))
	mBase = m.M
	v3641 = m.ExcPending
	if v3641 != 0 {
		goto L9
	} else {
		goto L1152
	}
L1152:
	;
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v33)+608))
	v3645 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v3648 = F_exec_cast_value(m, l0, v3640, v33+int32(600), v3644, v3645, int32(23), int32(-1))
	mBase = m.M
	v3649 = m.ExcPending
	if v3649 != 0 {
		goto L9
	} else {
		goto L1153
	}
L1153:
	;
	v3650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+600)))
	if v3650 == int32(1) {
		goto L37
	} else {
		goto L1154
	}
L1154:
	;
	v3653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3653 != 0 {
		goto L1155
	} else {
		goto L1156
	}
L1155:
	;
	F_SPI_freetuptable(m, v3653)
	mBase = m.M
	v3655 = m.ExcPending
	if v3655 != 0 {
		goto L9
	} else {
		goto L1158
	}
L1156:
	;
	goto L1157
L1157:
	;
	v3656 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v3656
	v3658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3658 == v3656 {
		v3664 = v3648
		goto L1150
	} else {
		goto L1159
	}
L1158:
	;
	goto L1157
L1159:
	;
	v3661 = *(*int32)(unsafe.Add(mBase, uint32(v3658)+20))
	F_MemoryContextReset(m, v3661)
	mBase = m.M
	v3663 = m.ExcPending
	if v3663 != 0 {
		goto L9
	} else {
		goto L1160
	}
L1160:
	;
	v3664 = v3648
	goto L1150
L1161:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v3714
	v3716 = int32(0)
	v3717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v3722 = *(*int32)(unsafe.Add(mBase, uint32(v3717+v3718<<(uint(int32(2))%32))))
	F_assign_simple_var(m, l0, v3722, base.B2i32(v3714 != int64(0)), v3716, v3716)
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L9
	} else {
		goto L1181
	}
L1162:
	;
	v3671 = F_CreateDestReceiver(m, int32(5))
	mBase = m.M
	v3672 = m.ExcPending
	if v3672 != 0 {
		goto L9
	} else {
		goto L1165
	}
L1163:
	;
	goto L1164
L1164:
	;
	v3707 = *(*int32)(unsafe.Add(mBase, _consts[921]))
	F__SPI_cursor_operation(m, v3627, v3666, v3664, v3707)
	mBase = m.M
	v3709 = m.ExcPending
	if v3709 != 0 {
		goto L9
	} else {
		goto L1180
	}
L1165:
	;
	F__SPI_cursor_operation(m, v3627, v3666, v3664, v3671)
	mBase = m.M
	v3674 = m.ExcPending
	if v3674 != 0 {
		goto L9
	} else {
		goto L1166
	}
L1166:
	;
	v3676 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	v3677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3678 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v3679 = *(*int32)(unsafe.Add(mBase, uint32(v3678)+4))
	v3683 = *(*int32)(unsafe.Add(mBase, uint32(v3677+v3679<<(uint(int32(2))%32))))
	v3685 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	if v3685 == int64(0) {
		goto L1167
	} else {
		goto L1168
	}
L1167:
	;
	v3691 = int32(0)
	goto L1169
L1168:
	;
	v3689 = *(*int32)(unsafe.Add(mBase, uint32(v3676)+4))
	v3690 = *(*int32)(unsafe.Add(mBase, uint32(v3689)))
	v3691 = v3690
	goto L1169
L1169:
	;
	v3692 = *(*int32)(unsafe.Add(mBase, uint32(v3676)))
	F_exec_move_row(m, l0, v3683, v3691, v3692)
	mBase = m.M
	v3694 = m.ExcPending
	if v3694 != 0 {
		goto L9
	} else {
		goto L1170
	}
L1170:
	;
	v3695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3695 != 0 {
		goto L1171
	} else {
		goto L1172
	}
L1171:
	;
	F_SPI_freetuptable(m, v3695)
	mBase = m.M
	v3697 = m.ExcPending
	if v3697 != 0 {
		goto L9
	} else {
		goto L1174
	}
L1172:
	;
	goto L1173
L1173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v3700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v3700 != 0 {
		goto L1175
	} else {
		goto L1176
	}
L1174:
	;
	goto L1173
L1175:
	;
	v3701 = *(*int32)(unsafe.Add(mBase, uint32(v3700)+20))
	F_MemoryContextReset(m, v3701)
	mBase = m.M
	v3703 = m.ExcPending
	if v3703 != 0 {
		goto L9
	} else {
		goto L1178
	}
L1176:
	;
	goto L1177
L1177:
	;
	F_SPI_freetuptable(m, v3676)
	mBase = m.M
	v3705 = m.ExcPending
	if v3705 != 0 {
		goto L9
	} else {
		goto L1179
	}
L1178:
	;
	goto L1177
L1179:
	;
	v3714 = v3685
	goto L1161
L1180:
	;
	v3711 = *(*int64)(unsafe.Add(mBase, _consts[363]))
	v3714 = v3711
	goto L1161
L1181:
	;
	v5234 = v3716
	goto L21
L1182:
	;
	v3738 = int32(4515600)
	v3739 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v3741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v3742 = *(*int32)(unsafe.Add(mBase, uint32(v3741)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3742
	v3744 = *(*int32)(unsafe.Add(mBase, uint32(v3734)+40))
	v3745 = F_text_to_cstring(m, v3744)
	mBase = m.M
	v3746 = m.ExcPending
	if v3746 != 0 {
		goto L9
	} else {
		goto L1183
	}
L1183:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v3739
	v3749 = F_GetPortalByName(m, v3745)
	mBase = m.M
	v3750 = m.ExcPending
	if v3750 != 0 {
		goto L9
	} else {
		goto L1184
	}
L1184:
	;
	if v3749 == int32(0) {
		goto L35
	} else {
		goto L1185
	}
L1185:
	;
	F_SPI_cursor_close(m, v3749)
	mBase = m.M
	v3754 = m.ExcPending
	if v3754 != 0 {
		goto L9
	} else {
		goto L1186
	}
L1186:
	;
	v5234 = int32(0)
	goto L21
L1187:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
	goto L34
L1188:
	;
	F__SPI_commit(m, int32(1))
	mBase = m.M
	v3761 = m.ExcPending
	if v3761 != 0 {
		goto L9
	} else {
		goto L1191
	}
L1189:
	;
	goto L1190
L1190:
	;
	F__SPI_commit(m, int32(0))
	mBase = m.M
	v3764 = m.ExcPending
	if v3764 != 0 {
		goto L9
	} else {
		goto L1192
	}
L1191:
	;
	goto L1187
L1192:
	;
	goto L1187
L1193:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
	goto L34
L1194:
	;
	F__SPI_rollback(m, int32(1))
	mBase = m.M
	v3772 = m.ExcPending
	if v3772 != 0 {
		goto L9
	} else {
		goto L1197
	}
L1195:
	;
	goto L1196
L1196:
	;
	F__SPI_rollback(m, int32(0))
	mBase = m.M
	v3775 = m.ExcPending
	if v3775 != 0 {
		goto L9
	} else {
		goto L1198
	}
L1197:
	;
	goto L1193
L1198:
	;
	goto L1193
L1199:
	;
	v3783 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v3783
	F_errmsg_internal(m, int32(484623), v33)
	mBase = m.M
	v3787 = m.ExcPending
	if v3787 != 0 {
		goto L9
	} else {
		goto L1200
	}
L1200:
	;
	F_errfinish(m, int32(500315), int32(2138), int32(123456))
	mBase = m.M
	v3792 = m.ExcPending
	if v3792 != 0 {
		goto L9
	} else {
		goto L1201
	}
L1201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1202:
	;
	v5234 = v3793
	goto L21
L1203:
	;
	F_errmsg_internal(m, int32(97598), int32(0))
	mBase = m.M
	v3803 = m.ExcPending
	if v3803 != 0 {
		goto L9
	} else {
		goto L1204
	}
L1204:
	;
	F_errfinish(m, int32(500315), int32(2315), int32(108168))
	mBase = m.M
	v3808 = m.ExcPending
	if v3808 != 0 {
		goto L9
	} else {
		goto L1205
	}
L1205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1206:
	;
	F_errmsg_internal(m, int32(97598), int32(0))
	mBase = m.M
	v3816 = m.ExcPending
	if v3816 != 0 {
		goto L9
	} else {
		goto L1207
	}
L1207:
	;
	F_errfinish(m, int32(500315), int32(2319), int32(108168))
	mBase = m.M
	v3821 = m.ExcPending
	if v3821 != 0 {
		goto L9
	} else {
		goto L1208
	}
L1208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1209:
	;
	v3826 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v3826
	F_errmsg_internal(m, int32(44718), v33+int32(16))
	mBase = m.M
	v3832 = m.ExcPending
	if v3832 != 0 {
		goto L9
	} else {
		goto L1210
	}
L1210:
	;
	F_errfinish(m, int32(500315), int32(2327), int32(108168))
	mBase = m.M
	v3837 = m.ExcPending
	if v3837 != 0 {
		goto L9
	} else {
		goto L1211
	}
L1211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1212:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3844 = m.ExcPending
	if v3844 != 0 {
		goto L9
	} else {
		goto L1213
	}
L1213:
	;
	v3845 = *(*int32)(unsafe.Add(mBase, uint32(v33)+636))
	v3849 = *(*int32)(unsafe.Add(mBase, uint32(v3845+v228<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v3849
	F_errmsg(m, int32(391980), v33-int32(-64))
	mBase = m.M
	v3855 = m.ExcPending
	if v3855 != 0 {
		goto L9
	} else {
		goto L1214
	}
L1214:
	;
	F_errfinish(m, int32(500315), int32(2383), int32(108168))
	mBase = m.M
	v3860 = m.ExcPending
	if v3860 != 0 {
		goto L9
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
	v3865 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v3866 = F_SPI_result_code_string(m, v410)
	mBase = m.M
	v3867 = m.ExcPending
	if v3867 != 0 {
		goto L9
	} else {
		goto L1217
	}
L1217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+36)) = v3866
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v3865
	F_errmsg_internal(m, int32(205222), v33+int32(32))
	mBase = m.M
	v3874 = m.ExcPending
	if v3874 != 0 {
		goto L9
	} else {
		goto L1218
	}
L1218:
	;
	F_errfinish(m, int32(500315), int32(2246), int32(305027))
	mBase = m.M
	v3879 = m.ExcPending
	if v3879 != 0 {
		goto L9
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
	F_errmsg_internal(m, int32(30912), int32(0))
	mBase = m.M
	v3887 = m.ExcPending
	if v3887 != 0 {
		goto L9
	} else {
		goto L1221
	}
L1221:
	;
	F_errfinish(m, int32(500315), int32(2270), int32(305027))
	mBase = m.M
	v3892 = m.ExcPending
	if v3892 != 0 {
		goto L9
	} else {
		goto L1222
	}
L1222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1223:
	;
	F_errmsg_internal(m, int32(30535), int32(0))
	mBase = m.M
	v3900 = m.ExcPending
	if v3900 != 0 {
		goto L9
	} else {
		goto L1224
	}
L1224:
	;
	F_errfinish(m, int32(500315), int32(2275), int32(305027))
	mBase = m.M
	v3905 = m.ExcPending
	if v3905 != 0 {
		goto L9
	} else {
		goto L1225
	}
L1225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1226:
	;
	F_errcode(m, int32(33557120))
	mBase = m.M
	v3912 = m.ExcPending
	if v3912 != 0 {
		goto L9
	} else {
		goto L1227
	}
L1227:
	;
	F_errmsg(m, int32(219853), int32(0))
	mBase = m.M
	v3916 = m.ExcPending
	if v3916 != 0 {
		goto L9
	} else {
		goto L1228
	}
L1228:
	;
	F_errfinish(m, int32(500315), int32(2423), int32(338060))
	mBase = m.M
	v3921 = m.ExcPending
	if v3921 != 0 {
		goto L9
	} else {
		goto L1229
	}
L1229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1230:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v3928 = m.ExcPending
	if v3928 != 0 {
		goto L9
	} else {
		goto L1231
	}
L1231:
	;
	F_errmsg(m, int32(302789), int32(0))
	mBase = m.M
	v3932 = m.ExcPending
	if v3932 != 0 {
		goto L9
	} else {
		goto L1232
	}
L1232:
	;
	F_errfinish(m, int32(500315), int32(2723), int32(319920))
	mBase = m.M
	v3937 = m.ExcPending
	if v3937 != 0 {
		goto L9
	} else {
		goto L1233
	}
L1233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1234:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L9
	} else {
		goto L1235
	}
L1235:
	;
	F_errmsg(m, int32(302828), int32(0))
	mBase = m.M
	v3948 = m.ExcPending
	if v3948 != 0 {
		goto L9
	} else {
		goto L1236
	}
L1236:
	;
	F_errfinish(m, int32(500315), int32(2739), int32(319920))
	mBase = m.M
	v3953 = m.ExcPending
	if v3953 != 0 {
		goto L9
	} else {
		goto L1237
	}
L1237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1238:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v3960 = m.ExcPending
	if v3960 != 0 {
		goto L9
	} else {
		goto L1239
	}
L1239:
	;
	F_errmsg(m, int32(302753), int32(0))
	mBase = m.M
	v3964 = m.ExcPending
	if v3964 != 0 {
		goto L9
	} else {
		goto L1240
	}
L1240:
	;
	F_errfinish(m, int32(500315), int32(2757), int32(319920))
	mBase = m.M
	v3969 = m.ExcPending
	if v3969 != 0 {
		goto L9
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v3976 = m.ExcPending
	if v3976 != 0 {
		goto L9
	} else {
		goto L1243
	}
L1243:
	;
	F_errmsg(m, int32(239943), int32(0))
	mBase = m.M
	v3980 = m.ExcPending
	if v3980 != 0 {
		goto L9
	} else {
		goto L1244
	}
L1244:
	;
	F_errfinish(m, int32(500315), int32(2763), int32(319920))
	mBase = m.M
	v3985 = m.ExcPending
	if v3985 != 0 {
		goto L9
	} else {
		goto L1245
	}
L1245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1246:
	;
	F_errcode(m, int32(50462852))
	mBase = m.M
	v3992 = m.ExcPending
	if v3992 != 0 {
		goto L9
	} else {
		goto L1247
	}
L1247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1749
	F_errmsg(m, int32(360349), v33+int32(112))
	mBase = m.M
	v3998 = m.ExcPending
	if v3998 != 0 {
		goto L9
	} else {
		goto L1248
	}
L1248:
	;
	F_errfinish(m, int32(500315), int32(2897), int32(488993))
	mBase = m.M
	v4003 = m.ExcPending
	if v4003 != 0 {
		goto L9
	} else {
		goto L1249
	}
L1249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1250:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4010 = m.ExcPending
	if v4010 != 0 {
		goto L9
	} else {
		goto L1251
	}
L1251:
	;
	F_errmsg(m, int32(121017), int32(0))
	mBase = m.M
	v4014 = m.ExcPending
	if v4014 != 0 {
		goto L9
	} else {
		goto L1252
	}
L1252:
	;
	F_errfinish(m, int32(500315), int32(2920), int32(488993))
	mBase = m.M
	v4019 = m.ExcPending
	if v4019 != 0 {
		goto L9
	} else {
		goto L1253
	}
L1253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1254:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4026 = m.ExcPending
	if v4026 != 0 {
		goto L9
	} else {
		goto L1255
	}
L1255:
	;
	F_errmsg(m, int32(210119), int32(0))
	mBase = m.M
	v4030 = m.ExcPending
	if v4030 != 0 {
		goto L9
	} else {
		goto L1256
	}
L1256:
	;
	F_errfinish(m, int32(500315), int32(2939), int32(488993))
	mBase = m.M
	v4035 = m.ExcPending
	if v4035 != 0 {
		goto L9
	} else {
		goto L1257
	}
L1257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1258:
	;
	v4041 = *(*int32)(unsafe.Add(mBase, _consts[365]))
	v4042 = F_SPI_result_code_string(m, v4041)
	mBase = m.M
	v4043 = m.ExcPending
	if v4043 != 0 {
		goto L9
	} else {
		goto L1259
	}
L1259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v4042
	F_errmsg_internal(m, int32(200811), v33+int32(96))
	mBase = m.M
	v4049 = m.ExcPending
	if v4049 != 0 {
		goto L9
	} else {
		goto L1260
	}
L1260:
	;
	F_errfinish(m, int32(500315), int32(2961), int32(488993))
	mBase = m.M
	v4054 = m.ExcPending
	if v4054 != 0 {
		goto L9
	} else {
		goto L1261
	}
L1261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1262:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4061 = m.ExcPending
	if v4061 != 0 {
		goto L9
	} else {
		goto L1263
	}
L1263:
	;
	F_errmsg(m, int32(303386), int32(0))
	mBase = m.M
	v4065 = m.ExcPending
	if v4065 != 0 {
		goto L9
	} else {
		goto L1264
	}
L1264:
	;
	F_errfinish(m, int32(500315), int32(3030), int32(507694))
	mBase = m.M
	v4070 = m.ExcPending
	if v4070 != 0 {
		goto L9
	} else {
		goto L1265
	}
L1265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1266:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4077 = m.ExcPending
	if v4077 != 0 {
		goto L9
	} else {
		goto L1267
	}
L1267:
	;
	v4078 = F_format_type_be(m, v1895)
	mBase = m.M
	v4079 = m.ExcPending
	if v4079 != 0 {
		goto L9
	} else {
		goto L1268
	}
L1268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v4078
	F_errmsg(m, int32(188715), v33+int32(128))
	mBase = m.M
	v4085 = m.ExcPending
	if v4085 != 0 {
		goto L9
	} else {
		goto L1269
	}
L1269:
	;
	F_errfinish(m, int32(500315), int32(3046), int32(507694))
	mBase = m.M
	v4090 = m.ExcPending
	if v4090 != 0 {
		goto L9
	} else {
		goto L1270
	}
L1270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1271:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v4097 = m.ExcPending
	if v4097 != 0 {
		goto L9
	} else {
		goto L1272
	}
L1272:
	;
	v4098 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	v4099 = *(*int32)(unsafe.Add(mBase, uint32(v1901)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+148)) = v4099
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v4098
	F_errmsg(m, int32(467014), v33+int32(144))
	mBase = m.M
	v4106 = m.ExcPending
	if v4106 != 0 {
		goto L9
	} else {
		goto L1273
	}
L1273:
	;
	F_errfinish(m, int32(500315), int32(3063), int32(507694))
	mBase = m.M
	v4111 = m.ExcPending
	if v4111 != 0 {
		goto L9
	} else {
		goto L1274
	}
L1274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1275:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4118 = m.ExcPending
	if v4118 != 0 {
		goto L9
	} else {
		goto L1276
	}
L1276:
	;
	F_errmsg(m, int32(367237), int32(0))
	mBase = m.M
	v4122 = m.ExcPending
	if v4122 != 0 {
		goto L9
	} else {
		goto L1277
	}
L1277:
	;
	F_errfinish(m, int32(500315), int32(3089), int32(507694))
	mBase = m.M
	v4127 = m.ExcPending
	if v4127 != 0 {
		goto L9
	} else {
		goto L1278
	}
L1278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1279:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4134 = m.ExcPending
	if v4134 != 0 {
		goto L9
	} else {
		goto L1280
	}
L1280:
	;
	F_errmsg(m, int32(367294), int32(0))
	mBase = m.M
	v4138 = m.ExcPending
	if v4138 != 0 {
		goto L9
	} else {
		goto L1281
	}
L1281:
	;
	F_errfinish(m, int32(500315), int32(3093), int32(507694))
	mBase = m.M
	v4143 = m.ExcPending
	if v4143 != 0 {
		goto L9
	} else {
		goto L1282
	}
L1282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1283:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4150 = m.ExcPending
	if v4150 != 0 {
		goto L9
	} else {
		goto L1284
	}
L1284:
	;
	F_errmsg(m, int32(254454), int32(0))
	mBase = m.M
	v4154 = m.ExcPending
	if v4154 != 0 {
		goto L9
	} else {
		goto L1285
	}
L1285:
	;
	F_errfinish(m, int32(500315), int32(3337), int32(63483))
	mBase = m.M
	v4159 = m.ExcPending
	if v4159 != 0 {
		goto L9
	} else {
		goto L1286
	}
L1286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1287:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4166 = m.ExcPending
	if v4166 != 0 {
		goto L9
	} else {
		goto L1288
	}
L1288:
	;
	F_errmsg(m, int32(517147), int32(0))
	mBase = m.M
	v4170 = m.ExcPending
	if v4170 != 0 {
		goto L9
	} else {
		goto L1289
	}
L1289:
	;
	F_errfinish(m, int32(500315), int32(3378), int32(63483))
	mBase = m.M
	v4175 = m.ExcPending
	if v4175 != 0 {
		goto L9
	} else {
		goto L1290
	}
L1290:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1291:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4182 = m.ExcPending
	if v4182 != 0 {
		goto L9
	} else {
		goto L1292
	}
L1292:
	;
	F_errmsg(m, int32(517189), int32(0))
	mBase = m.M
	v4186 = m.ExcPending
	if v4186 != 0 {
		goto L9
	} else {
		goto L1293
	}
L1293:
	;
	F_errfinish(m, int32(500315), int32(3437), int32(63483))
	mBase = m.M
	v4191 = m.ExcPending
	if v4191 != 0 {
		goto L9
	} else {
		goto L1294
	}
L1294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1295:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4198 = m.ExcPending
	if v4198 != 0 {
		goto L9
	} else {
		goto L1296
	}
L1296:
	;
	F_errmsg(m, int32(369766), int32(0))
	mBase = m.M
	v4202 = m.ExcPending
	if v4202 != 0 {
		goto L9
	} else {
		goto L1297
	}
L1297:
	;
	F_errfinish(m, int32(500315), int32(3473), int32(63483))
	mBase = m.M
	v4207 = m.ExcPending
	if v4207 != 0 {
		goto L9
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4214 = m.ExcPending
	if v4214 != 0 {
		goto L9
	} else {
		goto L1300
	}
L1300:
	;
	F_errmsg(m, int32(517147), int32(0))
	mBase = m.M
	v4218 = m.ExcPending
	if v4218 != 0 {
		goto L9
	} else {
		goto L1301
	}
L1301:
	;
	F_errfinish(m, int32(500315), int32(3510), int32(63483))
	mBase = m.M
	v4223 = m.ExcPending
	if v4223 != 0 {
		goto L9
	} else {
		goto L1302
	}
L1302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1303:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4230 = m.ExcPending
	if v4230 != 0 {
		goto L9
	} else {
		goto L1304
	}
L1304:
	;
	F_errmsg(m, int32(254406), int32(0))
	mBase = m.M
	v4234 = m.ExcPending
	if v4234 != 0 {
		goto L9
	} else {
		goto L1305
	}
L1305:
	;
	F_errfinish(m, int32(500315), int32(3557), int32(15796))
	mBase = m.M
	v4239 = m.ExcPending
	if v4239 != 0 {
		goto L9
	} else {
		goto L1306
	}
L1306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1307:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4246 = m.ExcPending
	if v4246 != 0 {
		goto L9
	} else {
		goto L1308
	}
L1308:
	;
	F_errmsg(m, int32(302458), int32(0))
	mBase = m.M
	v4250 = m.ExcPending
	if v4250 != 0 {
		goto L9
	} else {
		goto L1309
	}
L1309:
	;
	F_errfinish(m, int32(500315), int32(3630), int32(15796))
	mBase = m.M
	v4255 = m.ExcPending
	if v4255 != 0 {
		goto L9
	} else {
		goto L1310
	}
L1310:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1311:
	;
	v4260 = F_SPI_result_code_string(m, v2717)
	mBase = m.M
	v4261 = m.ExcPending
	if v4261 != 0 {
		goto L9
	} else {
		goto L1312
	}
L1312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v4260
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v2687
	F_errmsg_internal(m, int32(205280), v33+int32(192))
	mBase = m.M
	v4268 = m.ExcPending
	if v4268 != 0 {
		goto L9
	} else {
		goto L1313
	}
L1313:
	;
	F_errfinish(m, int32(500315), int32(3651), int32(15796))
	mBase = m.M
	v4273 = m.ExcPending
	if v4273 != 0 {
		goto L9
	} else {
		goto L1314
	}
L1314:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1316:
	;
	F_errmsg_internal(m, int32(320991), int32(0))
	mBase = m.M
	v4283 = m.ExcPending
	if v4283 != 0 {
		goto L9
	} else {
		goto L1317
	}
L1317:
	;
	F_errfinish(m, int32(500315), int32(3798), int32(361640))
	mBase = m.M
	v4288 = m.ExcPending
	if v4288 != 0 {
		goto L9
	} else {
		goto L1318
	}
L1318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1319:
	;
	F_errmsg_internal(m, int32(320991), int32(0))
	mBase = m.M
	v4296 = m.ExcPending
	if v4296 != 0 {
		goto L9
	} else {
		goto L1320
	}
L1320:
	;
	F_errfinish(m, int32(500315), int32(3822), int32(361640))
	mBase = m.M
	v4301 = m.ExcPending
	if v4301 != 0 {
		goto L9
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4308 = m.ExcPending
	if v4308 != 0 {
		goto L9
	} else {
		goto L1323
	}
L1323:
	;
	F_errmsg(m, int32(302867), int32(0))
	mBase = m.M
	v4312 = m.ExcPending
	if v4312 != 0 {
		goto L9
	} else {
		goto L1324
	}
L1324:
	;
	F_errfinish(m, int32(500315), int32(3843), int32(361640))
	mBase = m.M
	v4317 = m.ExcPending
	if v4317 != 0 {
		goto L9
	} else {
		goto L1325
	}
L1325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1326:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4324 = m.ExcPending
	if v4324 != 0 {
		goto L9
	} else {
		goto L1327
	}
L1327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+288)) = int32(542317)
	F_errmsg(m, int32(204257), v33+int32(288))
	mBase = m.M
	v4331 = m.ExcPending
	if v4331 != 0 {
		goto L9
	} else {
		goto L1328
	}
L1328:
	;
	F_errfinish(m, int32(500315), int32(3854), int32(361640))
	mBase = m.M
	v4336 = m.ExcPending
	if v4336 != 0 {
		goto L9
	} else {
		goto L1329
	}
L1329:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1330:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4343 = m.ExcPending
	if v4343 != 0 {
		goto L9
	} else {
		goto L1331
	}
L1331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+304)) = int32(541844)
	F_errmsg(m, int32(204257), v33+int32(304))
	mBase = m.M
	v4350 = m.ExcPending
	if v4350 != 0 {
		goto L9
	} else {
		goto L1332
	}
L1332:
	;
	F_errfinish(m, int32(500315), int32(3859), int32(361640))
	mBase = m.M
	v4355 = m.ExcPending
	if v4355 != 0 {
		goto L9
	} else {
		goto L1333
	}
L1333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1334:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4362 = m.ExcPending
	if v4362 != 0 {
		goto L9
	} else {
		goto L1335
	}
L1335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+320)) = int32(533970)
	F_errmsg(m, int32(204257), v33+int32(320))
	mBase = m.M
	v4369 = m.ExcPending
	if v4369 != 0 {
		goto L9
	} else {
		goto L1336
	}
L1336:
	;
	F_errfinish(m, int32(500315), int32(3862), int32(361640))
	mBase = m.M
	v4374 = m.ExcPending
	if v4374 != 0 {
		goto L9
	} else {
		goto L1337
	}
L1337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1338:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4381 = m.ExcPending
	if v4381 != 0 {
		goto L9
	} else {
		goto L1339
	}
L1339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+336)) = int32(518377)
	F_errmsg(m, int32(204257), v33+int32(336))
	mBase = m.M
	v4388 = m.ExcPending
	if v4388 != 0 {
		goto L9
	} else {
		goto L1340
	}
L1340:
	;
	F_errfinish(m, int32(500315), int32(3865), int32(361640))
	mBase = m.M
	v4393 = m.ExcPending
	if v4393 != 0 {
		goto L9
	} else {
		goto L1341
	}
L1341:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1342:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4400 = m.ExcPending
	if v4400 != 0 {
		goto L9
	} else {
		goto L1343
	}
L1343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+352)) = int32(530561)
	F_errmsg(m, int32(204257), v33+int32(352))
	mBase = m.M
	v4407 = m.ExcPending
	if v4407 != 0 {
		goto L9
	} else {
		goto L1344
	}
L1344:
	;
	F_errfinish(m, int32(500315), int32(3868), int32(361640))
	mBase = m.M
	v4412 = m.ExcPending
	if v4412 != 0 {
		goto L9
	} else {
		goto L1345
	}
L1345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1346:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4419 = m.ExcPending
	if v4419 != 0 {
		goto L9
	} else {
		goto L1347
	}
L1347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+368)) = int32(518505)
	F_errmsg(m, int32(204257), v33+int32(368))
	mBase = m.M
	v4426 = m.ExcPending
	if v4426 != 0 {
		goto L9
	} else {
		goto L1348
	}
L1348:
	;
	F_errfinish(m, int32(500315), int32(3871), int32(361640))
	mBase = m.M
	v4431 = m.ExcPending
	if v4431 != 0 {
		goto L9
	} else {
		goto L1349
	}
L1349:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1350:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4438 = m.ExcPending
	if v4438 != 0 {
		goto L9
	} else {
		goto L1351
	}
L1351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+384)) = int32(539915)
	F_errmsg(m, int32(204257), v33+int32(384))
	mBase = m.M
	v4445 = m.ExcPending
	if v4445 != 0 {
		goto L9
	} else {
		goto L1352
	}
L1352:
	;
	F_errfinish(m, int32(500315), int32(3874), int32(361640))
	mBase = m.M
	v4450 = m.ExcPending
	if v4450 != 0 {
		goto L9
	} else {
		goto L1353
	}
L1353:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1354:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4457 = m.ExcPending
	if v4457 != 0 {
		goto L9
	} else {
		goto L1355
	}
L1355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+400)) = int32(541182)
	F_errmsg(m, int32(204257), v33+int32(400))
	mBase = m.M
	v4464 = m.ExcPending
	if v4464 != 0 {
		goto L9
	} else {
		goto L1356
	}
L1356:
	;
	F_errfinish(m, int32(500315), int32(3877), int32(361640))
	mBase = m.M
	v4469 = m.ExcPending
	if v4469 != 0 {
		goto L9
	} else {
		goto L1357
	}
L1357:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1358:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4476 = m.ExcPending
	if v4476 != 0 {
		goto L9
	} else {
		goto L1359
	}
L1359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+416)) = int32(545552)
	F_errmsg(m, int32(204257), v33+int32(416))
	mBase = m.M
	v4483 = m.ExcPending
	if v4483 != 0 {
		goto L9
	} else {
		goto L1360
	}
L1360:
	;
	F_errfinish(m, int32(500315), int32(3880), int32(361640))
	mBase = m.M
	v4488 = m.ExcPending
	if v4488 != 0 {
		goto L9
	} else {
		goto L1361
	}
L1361:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1362:
	;
	v4493 = *(*int32)(unsafe.Add(mBase, uint32(v3020)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+272)) = v4493
	F_errmsg_internal(m, int32(483199), v33+int32(272))
	mBase = m.M
	v4499 = m.ExcPending
	if v4499 != 0 {
		goto L9
	} else {
		goto L1363
	}
L1363:
	;
	F_errfinish(m, int32(500315), int32(3883), int32(361640))
	mBase = m.M
	v4504 = m.ExcPending
	if v4504 != 0 {
		goto L9
	} else {
		goto L1364
	}
L1364:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1365:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4511 = m.ExcPending
	if v4511 != 0 {
		goto L9
	} else {
		goto L1366
	}
L1366:
	;
	F_errmsg(m, int32(302458), int32(0))
	mBase = m.M
	v4515 = m.ExcPending
	if v4515 != 0 {
		goto L9
	} else {
		goto L1367
	}
L1367:
	;
	F_errfinish(m, int32(500315), int32(4461), int32(348903))
	mBase = m.M
	v4520 = m.ExcPending
	if v4520 != 0 {
		goto L9
	} else {
		goto L1368
	}
L1368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1369:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4527 = m.ExcPending
	if v4527 != 0 {
		goto L9
	} else {
		goto L1370
	}
L1370:
	;
	F_errmsg(m, int32(504885), int32(0))
	mBase = m.M
	v4531 = m.ExcPending
	if v4531 != 0 {
		goto L9
	} else {
		goto L1371
	}
L1371:
	;
	F_errfinish(m, int32(500315), int32(4553), int32(348903))
	mBase = m.M
	v4536 = m.ExcPending
	if v4536 != 0 {
		goto L9
	} else {
		goto L1372
	}
L1372:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1373:
	;
	F_errcode(m, int32(50462852))
	mBase = m.M
	v4543 = m.ExcPending
	if v4543 != 0 {
		goto L9
	} else {
		goto L1374
	}
L1374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+512)) = v3498
	F_errmsg(m, int32(360349), v33+int32(512))
	mBase = m.M
	v4549 = m.ExcPending
	if v4549 != 0 {
		goto L9
	} else {
		goto L1375
	}
L1375:
	;
	F_errfinish(m, int32(500315), int32(4685), int32(281788))
	mBase = m.M
	v4554 = m.ExcPending
	if v4554 != 0 {
		goto L9
	} else {
		goto L1376
	}
L1376:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1377:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4561 = m.ExcPending
	if v4561 != 0 {
		goto L9
	} else {
		goto L1378
	}
L1378:
	;
	F_errmsg(m, int32(121017), int32(0))
	mBase = m.M
	v4565 = m.ExcPending
	if v4565 != 0 {
		goto L9
	} else {
		goto L1379
	}
L1379:
	;
	F_errfinish(m, int32(500315), int32(4755), int32(281788))
	mBase = m.M
	v4570 = m.ExcPending
	if v4570 != 0 {
		goto L9
	} else {
		goto L1380
	}
L1380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1381:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4577 = m.ExcPending
	if v4577 != 0 {
		goto L9
	} else {
		goto L1382
	}
L1382:
	;
	F_errmsg(m, int32(210119), int32(0))
	mBase = m.M
	v4581 = m.ExcPending
	if v4581 != 0 {
		goto L9
	} else {
		goto L1383
	}
L1383:
	;
	F_errfinish(m, int32(500315), int32(4774), int32(281788))
	mBase = m.M
	v4586 = m.ExcPending
	if v4586 != 0 {
		goto L9
	} else {
		goto L1384
	}
L1384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1385:
	;
	v4592 = *(*int32)(unsafe.Add(mBase, _consts[365]))
	v4593 = F_SPI_result_code_string(m, v4592)
	mBase = m.M
	v4594 = m.ExcPending
	if v4594 != 0 {
		goto L9
	} else {
		goto L1386
	}
L1386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+496)) = v4593
	F_errmsg_internal(m, int32(200811), v33+int32(496))
	mBase = m.M
	v4600 = m.ExcPending
	if v4600 != 0 {
		goto L9
	} else {
		goto L1387
	}
L1387:
	;
	F_errfinish(m, int32(500315), int32(4795), int32(281788))
	mBase = m.M
	v4605 = m.ExcPending
	if v4605 != 0 {
		goto L9
	} else {
		goto L1388
	}
L1388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1389:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4612 = m.ExcPending
	if v4612 != 0 {
		goto L9
	} else {
		goto L1390
	}
L1390:
	;
	v4613 = *(*int32)(unsafe.Add(mBase, uint32(v3611)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+528)) = v4613
	F_errmsg(m, int32(302499), v33+int32(528))
	mBase = m.M
	v4619 = m.ExcPending
	if v4619 != 0 {
		goto L9
	} else {
		goto L1391
	}
L1391:
	;
	F_errfinish(m, int32(500315), int32(4840), int32(324234))
	mBase = m.M
	v4624 = m.ExcPending
	if v4624 != 0 {
		goto L9
	} else {
		goto L1392
	}
L1392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1393:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v4631 = m.ExcPending
	if v4631 != 0 {
		goto L9
	} else {
		goto L1394
	}
L1394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+544)) = v3623
	F_errmsg(m, int32(70860), v33+int32(544))
	mBase = m.M
	v4637 = m.ExcPending
	if v4637 != 0 {
		goto L9
	} else {
		goto L1395
	}
L1395:
	;
	F_errfinish(m, int32(500315), int32(4851), int32(324234))
	mBase = m.M
	v4642 = m.ExcPending
	if v4642 != 0 {
		goto L9
	} else {
		goto L1396
	}
L1396:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1397:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4649 = m.ExcPending
	if v4649 != 0 {
		goto L9
	} else {
		goto L1398
	}
L1398:
	;
	F_errmsg(m, int32(302321), int32(0))
	mBase = m.M
	v4653 = m.ExcPending
	if v4653 != 0 {
		goto L9
	} else {
		goto L1399
	}
L1399:
	;
	F_errfinish(m, int32(500315), int32(4864), int32(324234))
	mBase = m.M
	v4658 = m.ExcPending
	if v4658 != 0 {
		goto L9
	} else {
		goto L1400
	}
L1400:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1401:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4665 = m.ExcPending
	if v4665 != 0 {
		goto L9
	} else {
		goto L1402
	}
L1402:
	;
	v4666 = *(*int32)(unsafe.Add(mBase, uint32(v3734)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+560)) = v4666
	F_errmsg(m, int32(302499), v33+int32(560))
	mBase = m.M
	v4672 = m.ExcPending
	if v4672 != 0 {
		goto L9
	} else {
		goto L1403
	}
L1403:
	;
	F_errfinish(m, int32(500315), int32(4928), int32(361351))
	mBase = m.M
	v4677 = m.ExcPending
	if v4677 != 0 {
		goto L9
	} else {
		goto L1404
	}
L1404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1405:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v4684 = m.ExcPending
	if v4684 != 0 {
		goto L9
	} else {
		goto L1406
	}
L1406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+576)) = v3745
	F_errmsg(m, int32(70860), v33+int32(576))
	mBase = m.M
	v4690 = m.ExcPending
	if v4690 != 0 {
		goto L9
	} else {
		goto L1407
	}
L1407:
	;
	F_errfinish(m, int32(500315), int32(4939), int32(361351))
	mBase = m.M
	v4695 = m.ExcPending
	if v4695 != 0 {
		goto L9
	} else {
		goto L1408
	}
L1408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1409:
	;
	v5234 = int32(0)
	goto L21
L1410:
	;
	v4733 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if int32(20) < v4733 {
		goto L1413
	} else {
		goto L1414
	}
L1411:
	;
	v4737 = v4712
	goto L1412
L1412:
	;
	if v4711 != 0 {
		v4782 = v4711
		goto L1416
	} else {
		goto L1417
	}
L1413:
	;
	v4736 = int32(16777248)
	goto L1415
L1414:
	;
	v4736 = int32(0)
	goto L1415
L1415:
	;
	v4737 = v4736
	goto L1412
L1416:
	;
	v4783 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v4785 = F_errstart(m, v4783, int32(556508))
	mBase = m.M
	v4786 = m.ExcPending
	if v4786 != 0 {
		goto L9
	} else {
		goto L1421
	}
L1417:
	;
	if v4715 != 0 {
		v4782 = v4715
		goto L1416
	} else {
		goto L1418
	}
L1418:
	;
	v4738 = int32(4509336)
	v4739 = int32(63)
	v4741 = int32(48)
	v4742 = v4737&v4739 + v4741
	*(*uint8)(unsafe.Add(mBase, _consts[1158])) = uint8(v4742)
	v4750 = int32(base.Ui32(v4737)>>(uint(int32(24))%32))&v4739 + v4741
	*(*uint8)(unsafe.Add(mBase, _consts[1159])) = uint8(v4750)
	v4758 = int32(base.Ui32(v4737)>>(uint(int32(18))%32))&v4739 + v4741
	*(*uint8)(unsafe.Add(mBase, _consts[1160])) = uint8(v4758)
	v4766 = int32(base.Ui32(v4737)>>(uint(int32(12))%32))&v4739 + v4741
	*(*uint8)(unsafe.Add(mBase, _consts[1161])) = uint8(v4766)
	v4774 = int32(base.Ui32(v4737)>>(uint(int32(6))%32))&v4739 + v4741
	*(*uint8)(unsafe.Add(mBase, _consts[1162])) = uint8(v4774)
	v4777 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1163])) = uint8(v4777)
	goto L1419
L1419:
	;
	v4780 = F_MemoryContextStrdup(m, v2788, v4738)
	mBase = m.M
	v4781 = m.ExcPending
	if v4781 != 0 {
		goto L9
	} else {
		goto L1420
	}
L1420:
	;
	v4782 = v4780
	goto L1416
L1421:
	;
	if v4785 != 0 {
		goto L1422
	} else {
		goto L1423
	}
L1422:
	;
	if v4737 != 0 {
		goto L1425
	} else {
		goto L1426
	}
L1423:
	;
	goto L1424
L1424:
	;
	F_MemoryContextReset(m, v2788)
	mBase = m.M
	v4828 = m.ExcPending
	if v4828 != 0 {
		goto L9
	} else {
		goto L1459
	}
L1425:
	;
	F_errcode(m, v4737)
	mBase = m.M
	v4788 = m.ExcPending
	if v4788 != 0 {
		goto L9
	} else {
		goto L1428
	}
L1426:
	;
	goto L1427
L1427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+256)) = v4782
	F_errmsg_internal(m, int32(206200), v33+int32(256))
	mBase = m.M
	v4794 = m.ExcPending
	if v4794 != 0 {
		goto L9
	} else {
		goto L1429
	}
L1428:
	;
	goto L1427
L1429:
	;
	if v4722 != 0 {
		goto L1430
	} else {
		goto L1431
	}
L1430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+240)) = v4722
	F_errdetail_internal(m, int32(206200), v33+int32(240))
	mBase = m.M
	v4800 = m.ExcPending
	if v4800 != 0 {
		goto L9
	} else {
		goto L1433
	}
L1431:
	;
	goto L1432
L1432:
	;
	if v4721 != 0 {
		goto L1434
	} else {
		goto L1435
	}
L1433:
	;
	goto L1432
L1434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+224)) = v4721
	F_errhint(m, int32(206200), v33+int32(224))
	mBase = m.M
	v4806 = m.ExcPending
	if v4806 != 0 {
		goto L9
	} else {
		goto L1437
	}
L1435:
	;
	goto L1436
L1436:
	;
	if v4720 != 0 {
		goto L1438
	} else {
		goto L1439
	}
L1437:
	;
	goto L1436
L1438:
	;
	F_err_generic_string(m, int32(99), v4720)
	mBase = m.M
	v4809 = m.ExcPending
	if v4809 != 0 {
		goto L9
	} else {
		goto L1441
	}
L1439:
	;
	goto L1440
L1440:
	;
	if v4719 != 0 {
		goto L1442
	} else {
		goto L1443
	}
L1441:
	;
	goto L1440
L1442:
	;
	F_err_generic_string(m, int32(110), v4719)
	mBase = m.M
	v4812 = m.ExcPending
	if v4812 != 0 {
		goto L9
	} else {
		goto L1445
	}
L1443:
	;
	goto L1444
L1444:
	;
	if v4716 != 0 {
		goto L1446
	} else {
		goto L1447
	}
L1445:
	;
	goto L1444
L1446:
	;
	F_err_generic_string(m, int32(100), v4716)
	mBase = m.M
	v4815 = m.ExcPending
	if v4815 != 0 {
		goto L9
	} else {
		goto L1449
	}
L1447:
	;
	goto L1448
L1448:
	;
	if v4717 != 0 {
		goto L1450
	} else {
		goto L1451
	}
L1449:
	;
	goto L1448
L1450:
	;
	F_err_generic_string(m, int32(116), v4717)
	mBase = m.M
	v4818 = m.ExcPending
	if v4818 != 0 {
		goto L9
	} else {
		goto L1453
	}
L1451:
	;
	goto L1452
L1452:
	;
	if v4718 != 0 {
		goto L1454
	} else {
		goto L1455
	}
L1453:
	;
	goto L1452
L1454:
	;
	F_err_generic_string(m, int32(115), v4718)
	mBase = m.M
	v4821 = m.ExcPending
	if v4821 != 0 {
		goto L9
	} else {
		goto L1457
	}
L1455:
	;
	goto L1456
L1456:
	;
	F_errfinish(m, int32(500315), int32(3923), int32(361640))
	mBase = m.M
	v4826 = m.ExcPending
	if v4826 != 0 {
		goto L9
	} else {
		goto L1458
	}
L1457:
	;
	goto L1456
L1458:
	;
	goto L1424
L1459:
	;
	v5234 = int32(0)
	goto L21
L1460:
	;
	v4902 = int32(0)
	v4908 = v1566
	goto L30
L1461:
	;
	goto L1462
L1462:
	;
	v4864 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v4864 == int32(0) {
		goto L1463
	} else {
		goto L1464
	}
L1463:
	;
	v4902 = int32(1)
	v4908 = v1566
	goto L30
L1464:
	;
	goto L1465
L1465:
	;
	v4870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4860))))
	v4871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4864))))
	if v4871 == int32(0) {
		v4890 = v4870
		v4891 = v4871
		goto L1467
	} else {
		goto L1468
	}
L1466:
	;
	if v4891-v4890 != 0 {
		goto L1474
	} else {
		goto L1475
	}
L1467:
	;
	goto L1466
L1468:
	;
	if v4870 != v4871 {
		v4890 = v4870
		v4891 = v4871
		goto L1467
	} else {
		goto L1469
	}
L1469:
	;
	v4875 = v4864
	v4876 = v4860
	goto L1470
L1470:
	;
	v4879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4876)+1)))
	v4880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4875)+1)))
	if v4880 == int32(0) {
		v4890 = v4879
		v4891 = v4880
		goto L1467
	} else {
		goto L1472
	}
L1471:
	;
	v4890 = v4879
	v4891 = v4880
	goto L1467
L1472:
	;
	v4883 = int32(1)
	if v4879 == v4880 {
		v4875 = v4875 + v4883
		v4876 = v4876 + v4883
		goto L1470
	} else {
		goto L1473
	}
L1473:
	;
	goto L1471
L1474:
	;
	v4902 = int32(1)
	v4908 = v1566
	goto L30
L1475:
	;
	goto L1476
L1476:
	;
	v4894 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v4894
	v4902 = v4894
	v4908 = v1566
	goto L30
L1477:
	;
	v5234 = v4902
	goto L21
L1478:
	;
	v5234 = int32(0)
	goto L21
L1479:
	;
	goto L1480
L1480:
	;
	v4943 = int32(1)
	v4944 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v4944 == int32(0) {
		v5234 = v4943
		goto L21
	} else {
		goto L1481
	}
L1481:
	;
	v4949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4939))))
	v4950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4944))))
	if v4950 == int32(0) {
		v4969 = v4949
		v4970 = v4950
		goto L1483
	} else {
		goto L1484
	}
L1482:
	;
	if v4970-v4969 != 0 {
		v5234 = v4943
		goto L21
	} else {
		goto L1490
	}
L1483:
	;
	goto L1482
L1484:
	;
	if v4949 != v4950 {
		v4969 = v4949
		v4970 = v4950
		goto L1483
	} else {
		goto L1485
	}
L1485:
	;
	v4954 = v4944
	v4955 = v4939
	goto L1486
L1486:
	;
	v4958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4955)+1)))
	v4959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4954)+1)))
	if v4959 == int32(0) {
		v4969 = v4958
		v4970 = v4959
		goto L1483
	} else {
		goto L1488
	}
L1487:
	;
	v4969 = v4958
	v4970 = v4959
	goto L1483
L1488:
	;
	v4962 = int32(1)
	if v4958 == v4959 {
		v4954 = v4954 + v4962
		v4955 = v4955 + v4962
		goto L1486
	} else {
		goto L1489
	}
L1489:
	;
	goto L1487
L1490:
	;
	goto L28
L1491:
	;
	v5035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115)+45)))
	if v5035 != int32(1) {
		goto L1494
	} else {
		goto L1495
	}
L1492:
	;
	goto L1493
L1493:
	;
	v5064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+24)))
	if v5064 == int32(0) {
		goto L25
	} else {
		goto L1503
	}
L1494:
	;
	v5057 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1115)+48)) = v5057
	v5059 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1115)+44)) = uint16(v5059)
	*(*int32)(unsafe.Add(mBase, uint32(v1115)+40)) = v5057
	goto L1493
L1495:
	;
	v5038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115)+44)))
	if v5038 != 0 {
		goto L1496
	} else {
		goto L1497
	}
L1496:
	;
	v5053 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+40))
	F_pfree(m, v5053)
	mBase = m.M
	v5055 = m.ExcPending
	if v5055 != 0 {
		goto L9
	} else {
		goto L1502
	}
L1497:
	;
	v5039 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+24))
	v5040 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5039)+12)))
	if v5040 != int32(65535) {
		goto L1496
	} else {
		goto L1498
	}
L1498:
	;
	v5043 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+40))
	v5044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5043))))
	if v5044 != int32(1) {
		goto L1496
	} else {
		goto L1499
	}
L1499:
	;
	v5047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5043)+1)))
	if v5047 != int32(3) {
		goto L1496
	} else {
		goto L1500
	}
L1500:
	;
	F_DeleteExpandedObject(m, v5043)
	mBase = m.M
	v5051 = m.ExcPending
	if v5051 != 0 {
		goto L9
	} else {
		goto L1501
	}
L1501:
	;
	goto L1494
L1502:
	;
	goto L1494
L1503:
	;
	v5099 = v91 + int32(28)
	goto L26
L1504:
	;
	v5234 = v5101
	goto L21
L1505:
	;
	F_errcode(m, int32(2))
	mBase = m.M
	v5109 = m.ExcPending
	if v5109 != 0 {
		goto L9
	} else {
		goto L1506
	}
L1506:
	;
	F_errmsg(m, int32(424131), int32(0))
	mBase = m.M
	v5113 = m.ExcPending
	if v5113 != 0 {
		goto L9
	} else {
		goto L1507
	}
L1507:
	;
	F_errhint(m, int32(579344), int32(0))
	mBase = m.M
	v5117 = m.ExcPending
	if v5117 != 0 {
		goto L9
	} else {
		goto L1508
	}
L1508:
	;
	F_errfinish(m, int32(500315), int32(2630), int32(361796))
	mBase = m.M
	v5122 = m.ExcPending
	if v5122 != 0 {
		goto L9
	} else {
		goto L1509
	}
L1509:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1510:
	;
	v5234 = v5187
	goto L21
L1511:
	;
	F_SPI_freetuptable(m, v5219)
	mBase = m.M
	v5221 = m.ExcPending
	if v5221 != 0 {
		goto L9
	} else {
		goto L1514
	}
L1512:
	;
	goto L1513
L1513:
	;
	v5222 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v5222
	v5225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v5225 == v5222 {
		v5234 = v5222
		goto L21
	} else {
		goto L1515
	}
L1514:
	;
	goto L1513
L1515:
	;
	v5228 = *(*int32)(unsafe.Add(mBase, uint32(v5225)+20))
	F_MemoryContextReset(m, v5228)
	mBase = m.M
	v5230 = m.ExcPending
	if v5230 != 0 {
		goto L9
	} else {
		goto L1516
	}
L1516:
	;
	v5234 = v5222
	goto L21
L1517:
	;
	if v5234 == int32(0) {
		goto L1521
	} else {
		goto L1522
	}
L1518:
	;
	v5266 = *(*int32)(unsafe.Add(mBase, uint32(v5263)+16))
	if v5266 == int32(0) {
		v5273 = v5262
		goto L1517
	} else {
		goto L1519
	}
L1519:
	;
	m.T0[v5266].(func(*base.Module, int32, int32))(m, l0, v91)
	mBase = m.M
	v5270 = m.ExcPending
	if v5270 != 0 {
		goto L9
	} else {
		goto L1520
	}
L1520:
	;
	v5272 = *(*int32)(unsafe.Add(mBase, _consts[1351]))
	v5273 = v5272
	goto L1517
L1521:
	;
	v5278 = v84 + int32(1)
	v5279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5279 <= v5278 {
		goto L2
	} else {
		goto L1524
	}
L1522:
	;
	goto L1523
L1523:
	;
	goto L12
L1524:
	;
	v61 = v5273
	v84 = v5278
	goto L11
}
