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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
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
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v298 int32
	_ = v298
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int64
	_ = v396
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int64
	_ = v425
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int64
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v873 int32
	_ = v873
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int64
	_ = v925
	var v930 int32
	_ = v930
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v971 int64
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int64
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1044 int64
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int64
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1085 int64
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1134 int32
	_ = v1134
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1173 int64
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int64
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1210 int64
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1337 int32
	_ = v1337
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1371 int64
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1377 int64
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1483 int64
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int64
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1511 int64
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int64
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1537 int32
	_ = v1537
	var v1542 int64
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int64
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1579 int64
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1639 int32
	_ = v1639
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1669 int32
	_ = v1669
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
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1702 int32
	_ = v1702
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
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
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1859 int64
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1889 int64
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1960 int32
	_ = v1960
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1970 int32
	_ = v1970
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v2017 int64
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2027 int64
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2033 int32
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2043 int32
	_ = v2043
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2057 int32
	_ = v2057
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2067 int32
	_ = v2067
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2080 int32
	_ = v2080
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2107 int32
	_ = v2107
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2120 int32
	_ = v2120
	var v2144 int64
	_ = v2144
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
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
	var v2172 int64
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2178 int64
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2197 int32
	_ = v2197
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2211 int32
	_ = v2211
	var v2214 int32
	_ = v2214
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2222 int64
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2238 int32
	_ = v2238
	var v2241 int32
	_ = v2241
	var v2245 int32
	_ = v2245
	var v2250 int32
	_ = v2250
	var v2254 int32
	_ = v2254
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2265 int32
	_ = v2265
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2274 int64
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2277 int32
	_ = v2277
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2311 int32
	_ = v2311
	var v2314 int32
	_ = v2314
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2332 int64
	_ = v2332
	var v2333 int32
	_ = v2333
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2348 int32
	_ = v2348
	var v2351 int32
	_ = v2351
	var v2355 int64
	_ = v2355
	var v2356 int64
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2368 int64
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
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
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2441 int32
	_ = v2441
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2456 int64
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
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
	var v2478 int32
	_ = v2478
	var v2480 int32
	_ = v2480
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2492 int32
	_ = v2492
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2503 int32
	_ = v2503
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2540 int64
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2543 int32
	_ = v2543
	var v2547 int32
	_ = v2547
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2560 int32
	_ = v2560
	var v2563 int32
	_ = v2563
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2569 int32
	_ = v2569
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2581 int32
	_ = v2581
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2586 int64
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2588 int32
	_ = v2588
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2617 int32
	_ = v2617
	var v2619 int32
	_ = v2619
	var v2620 int64
	_ = v2620
	var v2627 int32
	_ = v2627
	var v2629 int32
	_ = v2629
	var v2632 int32
	_ = v2632
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2652 int32
	_ = v2652
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2665 int64
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2692 int32
	_ = v2692
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2698 int32
	_ = v2698
	var v2699 int64
	_ = v2699
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2711 int32
	_ = v2711
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2724 int32
	_ = v2724
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2736 int32
	_ = v2736
	var v2738 int32
	_ = v2738
	var v2739 int32
	_ = v2739
	var v2740 int64
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2748 int32
	_ = v2748
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
	var v2820 int32
	_ = v2820
	var v2825 int32
	_ = v2825
	var v2847 int32
	_ = v2847
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2866 int32
	_ = v2866
	var v2869 int32
	_ = v2869
	var v2876 int64
	_ = v2876
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2884 int32
	_ = v2884
	var v2885 int32
	_ = v2885
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2898 int32
	_ = v2898
	var v2903 int32
	_ = v2903
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2913 int32
	_ = v2913
	var v2915 int32
	_ = v2915
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2923 int32
	_ = v2923
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2928 int32
	_ = v2928
	var v2933 int32
	_ = v2933
	var v2938 int32
	_ = v2938
	var v2963 int32
	_ = v2963
	var v2966 int32
	_ = v2966
	var v2973 int32
	_ = v2973
	var v2981 int32
	_ = v2981
	var v2986 int32
	_ = v2986
	var v2988 int32
	_ = v2988
	var v2989 int32
	_ = v2989
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v3000 int32
	_ = v3000
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3013 int32
	_ = v3013
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3025 int64
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3030 int32
	_ = v3030
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3043 int32
	_ = v3043
	var v3046 int32
	_ = v3046
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3051 int32
	_ = v3051
	var v3052 int32
	_ = v3052
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
	var v3080 int32
	_ = v3080
	var v3083 int32
	_ = v3083
	var v3084 int32
	_ = v3084
	var v3086 int32
	_ = v3086
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3095 int32
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3100 int32
	_ = v3100
	var v3101 int32
	_ = v3101
	var v3102 int32
	_ = v3102
	var v3103 int32
	_ = v3103
	var v3104 int32
	_ = v3104
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3109 int32
	_ = v3109
	var v3124 int32
	_ = v3124
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3141 int32
	_ = v3141
	var v3149 int32
	_ = v3149
	var v3157 int32
	_ = v3157
	var v3165 int32
	_ = v3165
	var v3168 int32
	_ = v3168
	var v3171 int32
	_ = v3171
	var v3172 int32
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3179 int32
	_ = v3179
	var v3185 int32
	_ = v3185
	var v3191 int32
	_ = v3191
	var v3197 int32
	_ = v3197
	var v3200 int32
	_ = v3200
	var v3203 int32
	_ = v3203
	var v3206 int32
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3212 int32
	_ = v3212
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3222 int32
	_ = v3222
	var v3225 int32
	_ = v3225
	var v3227 int32
	_ = v3227
	var v3232 int64
	_ = v3232
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3235 int32
	_ = v3235
	var v3238 int64
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3242 int32
	_ = v3242
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3255 int32
	_ = v3255
	var v3256 int32
	_ = v3256
	var v3265 int64
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3271 int32
	_ = v3271
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3277 int32
	_ = v3277
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3294 int32
	_ = v3294
	var v3300 int32
	_ = v3300
	var v3303 int32
	_ = v3303
	var v3309 int32
	_ = v3309
	var v3313 int32
	_ = v3313
	var v3318 int32
	_ = v3318
	var v3324 int32
	_ = v3324
	var v3326 int32
	_ = v3326
	var v3329 int32
	_ = v3329
	var v3334 int32
	_ = v3334
	var v3335 int32
	_ = v3335
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3345 int64
	_ = v3345
	var v3346 int32
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3362 int32
	_ = v3362
	var v3363 int32
	_ = v3363
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3368 int32
	_ = v3368
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3372 int32
	_ = v3372
	var v3375 int32
	_ = v3375
	var v3376 int32
	_ = v3376
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3380 int32
	_ = v3380
	var v3381 int32
	_ = v3381
	var v3382 int64
	_ = v3382
	var v3389 int32
	_ = v3389
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3400 int32
	_ = v3400
	var v3403 int32
	_ = v3403
	var v3407 int32
	_ = v3407
	var v3411 int32
	_ = v3411
	var v3416 int32
	_ = v3416
	var v3420 int32
	_ = v3420
	var v3423 int32
	_ = v3423
	var v3427 int32
	_ = v3427
	var v3432 int32
	_ = v3432
	var v3436 int32
	_ = v3436
	var v3439 int32
	_ = v3439
	var v3443 int32
	_ = v3443
	var v3448 int32
	_ = v3448
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3461 int32
	_ = v3461
	var v3466 int32
	_ = v3466
	var v3468 int64
	_ = v3468
	var v3470 int32
	_ = v3470
	var v3474 int32
	_ = v3474
	var v3477 int32
	_ = v3477
	var v3478 int32
	_ = v3478
	var v3479 int32
	_ = v3479
	var v3483 int32
	_ = v3483
	var v3490 int32
	_ = v3490
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3500 int32
	_ = v3500
	var v3504 int32
	_ = v3504
	var v3507 int32
	_ = v3507
	var v3511 int32
	_ = v3511
	var v3517 int32
	_ = v3517
	var v3522 int32
	_ = v3522
	var v3523 int32
	_ = v3523
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3531 int32
	_ = v3531
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3537 int32
	_ = v3537
	var v3540 int32
	_ = v3540
	var v3544 int32
	_ = v3544
	var v3550 int32
	_ = v3550
	var v3555 int32
	_ = v3555
	var v3556 int32
	_ = v3556
	var v3557 int32
	_ = v3557
	var v3558 int32
	_ = v3558
	var v3559 int32
	_ = v3559
	var v3561 int32
	_ = v3561
	var v3562 int32
	_ = v3562
	var v3564 int32
	_ = v3564
	var v3565 int32
	_ = v3565
	var v3567 int32
	_ = v3567
	var v3570 int32
	_ = v3570
	var v3572 int32
	_ = v3572
	var v3576 int32
	_ = v3576
	var v3578 int32
	_ = v3578
	var v3580 int32
	_ = v3580
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3586 int32
	_ = v3586
	var v3587 int32
	_ = v3587
	var v3589 int32
	_ = v3589
	var v3590 int32
	_ = v3590
	var v3592 int32
	_ = v3592
	var v3593 int32
	_ = v3593
	var v3595 int32
	_ = v3595
	var v3596 int32
	_ = v3596
	var v3600 int32
	_ = v3600
	var v3601 int32
	_ = v3601
	var v3604 int32
	_ = v3604
	var v3607 int32
	_ = v3607
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3615 int32
	_ = v3615
	var v3616 int32
	_ = v3616
	var v3617 int32
	_ = v3617
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3625 int32
	_ = v3625
	var v3626 int32
	_ = v3626
	var v3628 int32
	_ = v3628
	var v3629 int32
	_ = v3629
	var v3630 int32
	_ = v3630
	var v3631 int32
	_ = v3631
	var v3632 int32
	_ = v3632
	var v3634 int32
	_ = v3634
	var v3635 int32
	_ = v3635
	var v3636 int32
	_ = v3636
	var v3637 int32
	_ = v3637
	var v3638 int32
	_ = v3638
	var v3639 int32
	_ = v3639
	var v3641 int32
	_ = v3641
	var v3643 int32
	_ = v3643
	var v3645 int32
	_ = v3645
	var v3646 int32
	_ = v3646
	var v3647 int32
	_ = v3647
	var v3652 int32
	_ = v3652
	var v3653 int32
	_ = v3653
	var v3654 int32
	_ = v3654
	var v3663 int32
	_ = v3663
	var v3664 int32
	_ = v3664
	var v3668 int32
	_ = v3668
	var v3672 int32
	_ = v3672
	var v3677 int32
	_ = v3677
	var v3678 int32
	_ = v3678
	var v3683 int32
	_ = v3683
	var v3684 int32
	_ = v3684
	var v3685 int32
	_ = v3685
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3691 int32
	_ = v3691
	var v3695 int32
	_ = v3695
	var v3697 int32
	_ = v3697
	var v3698 int32
	_ = v3698
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3701 int32
	_ = v3701
	var v3706 int32
	_ = v3706
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3710 int32
	_ = v3710
	var v3711 int32
	_ = v3711
	var v3716 int32
	_ = v3716
	var v3717 int32
	_ = v3717
	var v3719 int32
	_ = v3719
	var v3720 int32
	_ = v3720
	var v3723 int32
	_ = v3723
	var v3724 int32
	_ = v3724
	var v3726 int32
	_ = v3726
	var v3730 int32
	_ = v3730
	var v3731 int32
	_ = v3731
	var v3732 int32
	_ = v3732
	var v3736 int32
	_ = v3736
	var v3737 int32
	_ = v3737
	var v3740 int32
	_ = v3740
	var v3741 int32
	_ = v3741
	var v3742 int32
	_ = v3742
	var v3744 int32
	_ = v3744
	var v3745 int32
	_ = v3745
	var v3747 int32
	_ = v3747
	var v3748 int32
	_ = v3748
	var v3749 int32
	_ = v3749
	var v3752 int32
	_ = v3752
	var v3753 int32
	_ = v3753
	var v3756 int32
	_ = v3756
	var v3758 int32
	_ = v3758
	var v3763 int64
	_ = v3763
	var v3764 int32
	_ = v3764
	var v3765 int32
	_ = v3765
	var v3766 int32
	_ = v3766
	var v3769 int64
	_ = v3769
	var v3770 int32
	_ = v3770
	var v3771 int32
	_ = v3771
	var v3774 int32
	_ = v3774
	var v3776 int32
	_ = v3776
	var v3779 int32
	_ = v3779
	var v3780 int32
	_ = v3780
	var v3782 int32
	_ = v3782
	var v3785 int32
	_ = v3785
	var v3786 int32
	_ = v3786
	var v3787 int32
	_ = v3787
	var v3791 int32
	_ = v3791
	var v3792 int32
	_ = v3792
	var v3794 int32
	_ = v3794
	var v3796 int32
	_ = v3796
	var v3797 int32
	_ = v3797
	var v3798 int32
	_ = v3798
	var v3799 int32
	_ = v3799
	var v3803 int32
	_ = v3803
	var v3805 int64
	_ = v3805
	var v3809 int32
	_ = v3809
	var v3810 int32
	_ = v3810
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
	var v3820 int32
	_ = v3820
	var v3821 int32
	_ = v3821
	var v3823 int32
	_ = v3823
	var v3825 int32
	_ = v3825
	var v3827 int32
	_ = v3827
	var v3829 int32
	_ = v3829
	var v3831 int64
	_ = v3831
	var v3834 int64
	_ = v3834
	var v3836 int32
	_ = v3836
	var v3837 int32
	_ = v3837
	var v3838 int32
	_ = v3838
	var v3842 int32
	_ = v3842
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3851 int32
	_ = v3851
	var v3855 int32
	_ = v3855
	var v3856 int32
	_ = v3856
	var v3859 int32
	_ = v3859
	var v3860 int32
	_ = v3860
	var v3862 int32
	_ = v3862
	var v3863 int32
	_ = v3863
	var v3865 int32
	_ = v3865
	var v3866 int32
	_ = v3866
	var v3867 int32
	_ = v3867
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3875 int32
	_ = v3875
	var v3877 int32
	_ = v3877
	var v3882 int32
	_ = v3882
	var v3885 int32
	_ = v3885
	var v3886 int32
	_ = v3886
	var v3891 int32
	_ = v3891
	var v3894 int32
	_ = v3894
	var v3899 int32
	_ = v3899
	var v3900 int32
	_ = v3900
	var v3904 int32
	_ = v3904
	var v3909 int32
	_ = v3909
	var v3910 int32
	_ = v3910
	var v3911 int32
	_ = v3911
	var v3916 int32
	_ = v3916
	var v3920 int32
	_ = v3920
	var v3925 int32
	_ = v3925
	var v3929 int32
	_ = v3929
	var v3933 int32
	_ = v3933
	var v3938 int32
	_ = v3938
	var v3942 int32
	_ = v3942
	var v3943 int32
	_ = v3943
	var v3949 int32
	_ = v3949
	var v3954 int32
	_ = v3954
	var v3958 int32
	_ = v3958
	var v3961 int32
	_ = v3961
	var v3962 int32
	_ = v3962
	var v3966 int32
	_ = v3966
	var v3972 int32
	_ = v3972
	var v3977 int32
	_ = v3977
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3983 int32
	_ = v3983
	var v3984 int32
	_ = v3984
	var v3991 int32
	_ = v3991
	var v3996 int32
	_ = v3996
	var v4000 int32
	_ = v4000
	var v4004 int32
	_ = v4004
	var v4009 int32
	_ = v4009
	var v4013 int32
	_ = v4013
	var v4017 int32
	_ = v4017
	var v4022 int32
	_ = v4022
	var v4026 int32
	_ = v4026
	var v4029 int32
	_ = v4029
	var v4033 int32
	_ = v4033
	var v4038 int32
	_ = v4038
	var v4042 int32
	_ = v4042
	var v4045 int32
	_ = v4045
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
	var v4081 int32
	_ = v4081
	var v4086 int32
	_ = v4086
	var v4090 int32
	_ = v4090
	var v4093 int32
	_ = v4093
	var v4097 int32
	_ = v4097
	var v4102 int32
	_ = v4102
	var v4106 int32
	_ = v4106
	var v4109 int32
	_ = v4109
	var v4115 int32
	_ = v4115
	var v4120 int32
	_ = v4120
	var v4124 int32
	_ = v4124
	var v4127 int32
	_ = v4127
	var v4131 int32
	_ = v4131
	var v4136 int32
	_ = v4136
	var v4140 int32
	_ = v4140
	var v4143 int32
	_ = v4143
	var v4147 int32
	_ = v4147
	var v4152 int32
	_ = v4152
	var v4156 int32
	_ = v4156
	var v4158 int32
	_ = v4158
	var v4159 int32
	_ = v4159
	var v4160 int32
	_ = v4160
	var v4166 int32
	_ = v4166
	var v4171 int32
	_ = v4171
	var v4175 int32
	_ = v4175
	var v4178 int32
	_ = v4178
	var v4182 int32
	_ = v4182
	var v4187 int32
	_ = v4187
	var v4191 int32
	_ = v4191
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4196 int32
	_ = v4196
	var v4202 int32
	_ = v4202
	var v4207 int32
	_ = v4207
	var v4211 int32
	_ = v4211
	var v4214 int32
	_ = v4214
	var v4215 int32
	_ = v4215
	var v4216 int32
	_ = v4216
	var v4223 int32
	_ = v4223
	var v4228 int32
	_ = v4228
	var v4232 int32
	_ = v4232
	var v4235 int32
	_ = v4235
	var v4239 int32
	_ = v4239
	var v4244 int32
	_ = v4244
	var v4248 int32
	_ = v4248
	var v4251 int32
	_ = v4251
	var v4255 int32
	_ = v4255
	var v4260 int32
	_ = v4260
	var v4264 int32
	_ = v4264
	var v4267 int32
	_ = v4267
	var v4271 int32
	_ = v4271
	var v4276 int32
	_ = v4276
	var v4280 int32
	_ = v4280
	var v4283 int32
	_ = v4283
	var v4287 int32
	_ = v4287
	var v4292 int32
	_ = v4292
	var v4296 int32
	_ = v4296
	var v4299 int32
	_ = v4299
	var v4303 int32
	_ = v4303
	var v4308 int32
	_ = v4308
	var v4312 int32
	_ = v4312
	var v4315 int32
	_ = v4315
	var v4319 int32
	_ = v4319
	var v4324 int32
	_ = v4324
	var v4328 int32
	_ = v4328
	var v4331 int32
	_ = v4331
	var v4335 int32
	_ = v4335
	var v4340 int32
	_ = v4340
	var v4344 int32
	_ = v4344
	var v4347 int32
	_ = v4347
	var v4351 int32
	_ = v4351
	var v4356 int32
	_ = v4356
	var v4360 int32
	_ = v4360
	var v4363 int32
	_ = v4363
	var v4367 int32
	_ = v4367
	var v4372 int32
	_ = v4372
	var v4376 int32
	_ = v4376
	var v4379 int32
	_ = v4379
	var v4383 int32
	_ = v4383
	var v4388 int32
	_ = v4388
	var v4392 int32
	_ = v4392
	var v4393 int32
	_ = v4393
	var v4394 int32
	_ = v4394
	var v4401 int32
	_ = v4401
	var v4406 int32
	_ = v4406
	var v4408 int32
	_ = v4408
	var v4412 int32
	_ = v4412
	var v4416 int32
	_ = v4416
	var v4421 int32
	_ = v4421
	var v4425 int32
	_ = v4425
	var v4429 int32
	_ = v4429
	var v4434 int32
	_ = v4434
	var v4438 int32
	_ = v4438
	var v4441 int32
	_ = v4441
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
	var v4495 int32
	_ = v4495
	var v4502 int32
	_ = v4502
	var v4507 int32
	_ = v4507
	var v4511 int32
	_ = v4511
	var v4514 int32
	_ = v4514
	var v4521 int32
	_ = v4521
	var v4526 int32
	_ = v4526
	var v4530 int32
	_ = v4530
	var v4533 int32
	_ = v4533
	var v4540 int32
	_ = v4540
	var v4545 int32
	_ = v4545
	var v4549 int32
	_ = v4549
	var v4552 int32
	_ = v4552
	var v4559 int32
	_ = v4559
	var v4564 int32
	_ = v4564
	var v4568 int32
	_ = v4568
	var v4571 int32
	_ = v4571
	var v4578 int32
	_ = v4578
	var v4583 int32
	_ = v4583
	var v4587 int32
	_ = v4587
	var v4590 int32
	_ = v4590
	var v4597 int32
	_ = v4597
	var v4602 int32
	_ = v4602
	var v4606 int32
	_ = v4606
	var v4609 int32
	_ = v4609
	var v4616 int32
	_ = v4616
	var v4621 int32
	_ = v4621
	var v4625 int32
	_ = v4625
	var v4626 int32
	_ = v4626
	var v4632 int32
	_ = v4632
	var v4637 int32
	_ = v4637
	var v4641 int32
	_ = v4641
	var v4644 int32
	_ = v4644
	var v4648 int32
	_ = v4648
	var v4653 int32
	_ = v4653
	var v4657 int32
	_ = v4657
	var v4660 int32
	_ = v4660
	var v4664 int32
	_ = v4664
	var v4669 int32
	_ = v4669
	var v4673 int32
	_ = v4673
	var v4676 int32
	_ = v4676
	var v4682 int32
	_ = v4682
	var v4687 int32
	_ = v4687
	var v4691 int32
	_ = v4691
	var v4694 int32
	_ = v4694
	var v4698 int32
	_ = v4698
	var v4703 int32
	_ = v4703
	var v4707 int32
	_ = v4707
	var v4710 int32
	_ = v4710
	var v4714 int32
	_ = v4714
	var v4719 int32
	_ = v4719
	var v4723 int32
	_ = v4723
	var v4725 int32
	_ = v4725
	var v4726 int32
	_ = v4726
	var v4727 int32
	_ = v4727
	var v4733 int32
	_ = v4733
	var v4738 int32
	_ = v4738
	var v4742 int32
	_ = v4742
	var v4745 int32
	_ = v4745
	var v4746 int32
	_ = v4746
	var v4752 int32
	_ = v4752
	var v4757 int32
	_ = v4757
	var v4761 int32
	_ = v4761
	var v4764 int32
	_ = v4764
	var v4770 int32
	_ = v4770
	var v4775 int32
	_ = v4775
	var v4779 int32
	_ = v4779
	var v4782 int32
	_ = v4782
	var v4786 int32
	_ = v4786
	var v4791 int32
	_ = v4791
	var v4795 int32
	_ = v4795
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4805 int32
	_ = v4805
	var v4810 int32
	_ = v4810
	var v4814 int32
	_ = v4814
	var v4817 int32
	_ = v4817
	var v4823 int32
	_ = v4823
	var v4828 int32
	_ = v4828
	var v4832 int32
	_ = v4832
	var v4863 int32
	_ = v4863
	var v4867 int32
	_ = v4867
	var v4868 int32
	_ = v4868
	var v4873 int32
	_ = v4873
	var v4876 int32
	_ = v4876
	var v4879 int32
	_ = v4879
	var v4880 int32
	_ = v4880
	var v4883 int32
	_ = v4883
	var v4884 int32
	_ = v4884
	var v4887 int32
	_ = v4887
	var v4894 int32
	_ = v4894
	var v4895 int32
	_ = v4895
	var v4897 int32
	_ = v4897
	var v4904 int32
	_ = v4904
	var v4928 int64
	_ = v4928
	var v4931 int32
	_ = v4931
	var v4932 int32
	_ = v4932
	var v4936 int32
	_ = v4936
	var v4937 int32
	_ = v4937
	var v4940 int32
	_ = v4940
	var v4941 int32
	_ = v4941
	var v4945 int32
	_ = v4945
	var v4946 int32
	_ = v4946
	var v4951 int32
	_ = v4951
	var v4954 int32
	_ = v4954
	var v4957 int32
	_ = v4957
	var v4958 int32
	_ = v4958
	var v4961 int32
	_ = v4961
	var v4962 int32
	_ = v4962
	var v4965 int32
	_ = v4965
	var v4972 int32
	_ = v4972
	var v4973 int32
	_ = v4973
	var v5004 int32
	_ = v5004
	var v5036 int32
	_ = v5036
	var v5039 int32
	_ = v5039
	var v5040 int32
	_ = v5040
	var v5041 int32
	_ = v5041
	var v5044 int64
	_ = v5044
	var v5045 int32
	_ = v5045
	var v5046 int32
	_ = v5046
	var v5049 int32
	_ = v5049
	var v5053 int32
	_ = v5053
	var v5056 int32
	_ = v5056
	var v5058 int32
	_ = v5058
	var v5063 int32
	_ = v5063
	var v5069 int32
	_ = v5069
	var v5103 int32
	_ = v5103
	var v5104 int32
	_ = v5104
	var v5105 int32
	_ = v5105
	var v5106 int32
	_ = v5106
	var v5110 int32
	_ = v5110
	var v5113 int32
	_ = v5113
	var v5117 int32
	_ = v5117
	var v5121 int32
	_ = v5121
	var v5126 int32
	_ = v5126
	var v5187 int32
	_ = v5187
	var v5188 int32
	_ = v5188
	var v5189 int32
	_ = v5189
	var v5190 int32
	_ = v5190
	var v5220 int32
	_ = v5220
	var v5222 int32
	_ = v5222
	var v5223 int32
	_ = v5223
	var v5226 int32
	_ = v5226
	var v5229 int32
	_ = v5229
	var v5231 int32
	_ = v5231
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
	var v5315 int32
	_ = v5315
	v3 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(640)
	m.G0 = v32
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v32 + int32(640)
	return v5315
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v34
	v5315 = int32(0)
	goto L1
L3:
	;
	v60 = v47
	v79 = v3
	goto L11
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 <= int32(0) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[0]))
	if v49 == int32(0) {
		v5315 = v3
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v39 = l0 + int32(28)
	v41 = l0 + int32(24)
	v43 = l0 + int32(16)
	v45 = v32 + int32(624)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[1]))
	goto L3
L8:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v5315 = v3
	goto L1
L11:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v79<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v91 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v34
	v5315 = v5234
	goto L1
L13:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[0]))
	if v101 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	if v94 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	m.T0[v94].(func(*base.Module, int32, int32))(m, l0, v89)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
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
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	switch v104 {
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
	v5262 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[1]))
	v5263 = *(*int32)(unsafe.Add(mBase, uint32(v5262)))
	if v5263 == int32(0) {
		v5273 = v5262
		goto L1491
	} else {
		goto L1492
	}
L22:
	;
	v5220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v5220 != 0 {
		goto L1485
	} else {
		goto L1486
	}
L23:
	;
	v5188 = *(*int32)(unsafe.Add(mBase, uint32(v5187)))
	v5189 = F_exec_stmts(m, l0, v5188)
	mBase = m.M
	v5190 = m.ExcPending
	if v5190 != 0 {
		goto L9
	} else {
		goto L1484
	}
L24:
	;
	v5187 = v89 + int32(24)
	goto L23
L25:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v5110 = m.ExcPending
	if v5110 != 0 {
		goto L9
	} else {
		goto L1479
	}
L26:
	;
	v5104 = *(*int32)(unsafe.Add(mBase, uint32(v5103)))
	v5105 = F_exec_stmts(m, l0, v5104)
	mBase = m.M
	v5106 = m.ExcPending
	if v5106 != 0 {
		goto L9
	} else {
		goto L1478
	}
L27:
	;
	if v1122 != 0 {
		goto L1465
	} else {
		goto L1466
	}
L28:
	;
	v5004 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v5004
	v5234 = v5004
	goto L21
L29:
	;
	v4941 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4941 == int32(0) {
		goto L1453
	} else {
		goto L1454
	}
L30:
	;
	v4931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v4932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v4936 = *(*int32)(unsafe.Add(mBase, uint32(v4931+v4932<<(uint(int32(2))%32))))
	v4937 = int32(0)
	F_assign_simple_var(m, l0, v4936, v4928, v4937, v4937)
	mBase = m.M
	v4940 = m.ExcPending
	if v4940 != 0 {
		goto L9
	} else {
		goto L1452
	}
L31:
	;
	v4904 = int32(0)
	v4928 = int64(0)
	goto L30
L32:
	;
	v4863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4863 == int32(0) {
		goto L1440
	} else {
		goto L1441
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = int64(0)
	F_plpgsql_create_econtext(m, l0)
	mBase = m.M
	v4832 = m.ExcPending
	if v4832 != 0 {
		goto L9
	} else {
		goto L1439
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4814 = m.ExcPending
	if v4814 != 0 {
		goto L9
	} else {
		goto L1435
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4795 = m.ExcPending
	if v4795 != 0 {
		goto L9
	} else {
		goto L1431
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4779 = m.ExcPending
	if v4779 != 0 {
		goto L9
	} else {
		goto L1427
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4761 = m.ExcPending
	if v4761 != 0 {
		goto L9
	} else {
		goto L1423
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4742 = m.ExcPending
	if v4742 != 0 {
		goto L9
	} else {
		goto L1419
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4723 = m.ExcPending
	if v4723 != 0 {
		goto L9
	} else {
		goto L1415
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4707 = m.ExcPending
	if v4707 != 0 {
		goto L9
	} else {
		goto L1411
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4691 = m.ExcPending
	if v4691 != 0 {
		goto L9
	} else {
		goto L1407
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4673 = m.ExcPending
	if v4673 != 0 {
		goto L9
	} else {
		goto L1403
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4657 = m.ExcPending
	if v4657 != 0 {
		goto L9
	} else {
		goto L1399
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4641 = m.ExcPending
	if v4641 != 0 {
		goto L9
	} else {
		goto L1395
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4625 = m.ExcPending
	if v4625 != 0 {
		goto L9
	} else {
		goto L1392
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4606 = m.ExcPending
	if v4606 != 0 {
		goto L9
	} else {
		goto L1388
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4587 = m.ExcPending
	if v4587 != 0 {
		goto L9
	} else {
		goto L1384
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4568 = m.ExcPending
	if v4568 != 0 {
		goto L9
	} else {
		goto L1380
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4549 = m.ExcPending
	if v4549 != 0 {
		goto L9
	} else {
		goto L1376
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4530 = m.ExcPending
	if v4530 != 0 {
		goto L9
	} else {
		goto L1372
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4511 = m.ExcPending
	if v4511 != 0 {
		goto L9
	} else {
		goto L1368
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4492 = m.ExcPending
	if v4492 != 0 {
		goto L9
	} else {
		goto L1364
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4473 = m.ExcPending
	if v4473 != 0 {
		goto L9
	} else {
		goto L1360
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4454 = m.ExcPending
	if v4454 != 0 {
		goto L9
	} else {
		goto L1356
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4438 = m.ExcPending
	if v4438 != 0 {
		goto L9
	} else {
		goto L1352
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4425 = m.ExcPending
	if v4425 != 0 {
		goto L9
	} else {
		goto L1349
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4412 = m.ExcPending
	if v4412 != 0 {
		goto L9
	} else {
		goto L1346
	}
L58:
	;
	F_ReThrowError(m, v2758)
	mBase = m.M
	v4408 = m.ExcPending
	if v4408 != 0 {
		goto L9
	} else {
		goto L1345
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4392 = m.ExcPending
	if v4392 != 0 {
		goto L9
	} else {
		goto L1341
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4376 = m.ExcPending
	if v4376 != 0 {
		goto L9
	} else {
		goto L1337
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4360 = m.ExcPending
	if v4360 != 0 {
		goto L9
	} else {
		goto L1333
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4344 = m.ExcPending
	if v4344 != 0 {
		goto L9
	} else {
		goto L1329
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		goto L9
	} else {
		goto L1325
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4312 = m.ExcPending
	if v4312 != 0 {
		goto L9
	} else {
		goto L1321
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4296 = m.ExcPending
	if v4296 != 0 {
		goto L9
	} else {
		goto L1317
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4280 = m.ExcPending
	if v4280 != 0 {
		goto L9
	} else {
		goto L1313
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4264 = m.ExcPending
	if v4264 != 0 {
		goto L9
	} else {
		goto L1309
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4248 = m.ExcPending
	if v4248 != 0 {
		goto L9
	} else {
		goto L1305
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4232 = m.ExcPending
	if v4232 != 0 {
		goto L9
	} else {
		goto L1301
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4211 = m.ExcPending
	if v4211 != 0 {
		goto L9
	} else {
		goto L1297
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4191 = m.ExcPending
	if v4191 != 0 {
		goto L9
	} else {
		goto L1292
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4175 = m.ExcPending
	if v4175 != 0 {
		goto L9
	} else {
		goto L1288
	}
L73:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4156 = m.ExcPending
	if v4156 != 0 {
		goto L9
	} else {
		goto L1284
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4140 = m.ExcPending
	if v4140 != 0 {
		goto L9
	} else {
		goto L1280
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4124 = m.ExcPending
	if v4124 != 0 {
		goto L9
	} else {
		goto L1276
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4106 = m.ExcPending
	if v4106 != 0 {
		goto L9
	} else {
		goto L1272
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4090 = m.ExcPending
	if v4090 != 0 {
		goto L9
	} else {
		goto L1268
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4074 = m.ExcPending
	if v4074 != 0 {
		goto L9
	} else {
		goto L1264
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4058 = m.ExcPending
	if v4058 != 0 {
		goto L9
	} else {
		goto L1260
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4042 = m.ExcPending
	if v4042 != 0 {
		goto L9
	} else {
		goto L1256
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4026 = m.ExcPending
	if v4026 != 0 {
		goto L9
	} else {
		goto L1252
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4013 = m.ExcPending
	if v4013 != 0 {
		goto L9
	} else {
		goto L1249
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v4000 = m.ExcPending
	if v4000 != 0 {
		goto L9
	} else {
		goto L1246
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3981 = m.ExcPending
	if v3981 != 0 {
		goto L9
	} else {
		goto L1242
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3958 = m.ExcPending
	if v3958 != 0 {
		goto L9
	} else {
		goto L1238
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3942 = m.ExcPending
	if v3942 != 0 {
		goto L9
	} else {
		goto L1235
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3929 = m.ExcPending
	if v3929 != 0 {
		goto L9
	} else {
		goto L1232
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3916 = m.ExcPending
	if v3916 != 0 {
		goto L9
	} else {
		goto L1229
	}
L89:
	;
	v3910 = F_exec_stmt_block(m, l0, v89)
	mBase = m.M
	v3911 = m.ExcPending
	if v3911 != 0 {
		goto L9
	} else {
		goto L1228
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v34
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3899 = m.ExcPending
	if v3899 != 0 {
		goto L9
	} else {
		goto L1225
	}
L91:
	;
	v3886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v3886 == int32(1) {
		goto L1220
	} else {
		goto L1221
	}
L92:
	;
	v3877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v3877 == int32(1) {
		goto L1214
	} else {
		goto L1215
	}
L93:
	;
	v3850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v3851 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3855 = *(*int32)(unsafe.Add(mBase, uint32(v3850+v3851<<(uint(int32(2))%32))))
	v3856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3855)+48)))
	if v3856 == int32(1) {
		goto L35
	} else {
		goto L1208
	}
L94:
	;
	v3731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v3732 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v3736 = *(*int32)(unsafe.Add(mBase, uint32(v3731+v3732<<(uint(int32(2))%32))))
	v3737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3736)+48)))
	if v3737 == int32(1) {
		goto L38
	} else {
		goto L1169
	}
L95:
	;
	v3593 = int32(0)
	v3595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v3596 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3600 = *(*int32)(unsafe.Add(mBase, uint32(v3595+v3596<<(uint(int32(2))%32))))
	v3601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3600)+48)))
	if v3601 == v3593 {
		goto L1112
	} else {
		goto L1113
	}
L96:
	;
	v3582 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v3583 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v3586 = F_exec_dynquery_with_params(m, l0, v3582, v3583, int32(0), int32(4))
	mBase = m.M
	v3587 = m.ExcPending
	if v3587 != 0 {
		goto L9
	} else {
		goto L1109
	}
L97:
	;
	v3326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3326 == int32(0) {
		goto L1024
	} else {
		goto L1025
	}
L98:
	;
	F_exec_stmt_execsql(m, l0, v89)
	mBase = m.M
	v3324 = m.ExcPending
	if v3324 != 0 {
		goto L9
	} else {
		goto L1023
	}
L99:
	;
	v3222 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[2])))
	if v3222 != int32(1) {
		goto L995
	} else {
		goto L996
	}
L100:
	;
	v2755 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v2755 != 0 {
		goto L838
	} else {
		goto L839
	}
L101:
	;
	v2566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2566 == int32(0) {
		goto L782
	} else {
		goto L783
	}
L102:
	;
	v2311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v2311 == int32(0) {
		goto L67
	} else {
		goto L695
	}
L103:
	;
	v2203 = int32(2)
	v2204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v2204 != 0 {
		v5234 = v2203
		goto L21
	} else {
		goto L663
	}
L104:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v2165 != 0 {
		goto L646
	} else {
		goto L647
	}
L105:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1889 = F_exec_eval_expr(m, l0, v1882, v32+int32(604), v32+int32(592), v32+int32(608))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L9
	} else {
		goto L571
	}
L106:
	;
	v1735 = int32(0)
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1737+v1738<<(uint(int32(2))%32))))
	v1743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1742)+48)))
	if v1743 == v1735 {
		goto L514
	} else {
		goto L515
	}
L107:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1727 = F_exec_run_select(m, l0, v1724, v32+int32(616))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L9
	} else {
		goto L511
	}
L108:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+4))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1469+v1471<<(uint(int32(2))%32))))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v1478 = v32 + int32(608)
	v1483 = F_exec_eval_expr(m, l0, v1476, v1478, v32+int32(616), v32+int32(592))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L9
	} else {
		goto L420
	}
L109:
	;
	goto L380
L110:
	;
	goto L366
L111:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1075 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L112:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v966 = v32 + int32(608)
	v971 = F_exec_eval_expr(m, l0, v964, v966, v32+int32(616), v32+int32(592))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L9
	} else {
		goto L289
	}
L113:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v457 == int32(1) {
		goto L190
	} else {
		goto L191
	}
L114:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	if v146 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L115:
	;
	v115 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v118 = F_exec_run_select(m, l0, v116, v115)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L118
	}
L116:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105+v106<<(uint(int32(2))%32))))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	F_exec_assign_expr(m, l0, v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120+v121<<(uint(int32(2))%32))))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v130 = int32(0)
	F_assign_simple_var(m, l0, v125, base.I64_extend_i32_u(base.B2i32(v126 != int64(0))), v130, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v134 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_SPI_freetuptable(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L9
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v137 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v139 == v137 {
		v5234 = v115
		goto L21
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+20))
	F_MemoryContextReset(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	v5234 = v115
	goto L21
L126:
	;
	F_exec_prepare_plan(m, l0, v145, int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+16)))
	if v152 != int32(1) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	goto L128
L130:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[3]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+44))
	v396 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v396
	*(*int64)(unsafe.Add(mBase, uint32(v32)+624)) = v396
	*(*int64)(unsafe.Add(mBase, uint32(v32)+632)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v32)+616)) = v392
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+620)) = uint8(v403)
	v405 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+621)) = uint8(v405)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+636)) = v407
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	v412 = F_SPI_execute_plan_extended(m, v409, v32+int32(616))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L9
	} else {
		goto L169
	}
L131:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+20)) = v145
	v392 = v390
	goto L130
L132:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	if v359 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L133:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v145)+28))
	if v357 != 0 {
		goto L131
	} else {
		goto L160
	}
L134:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v155 != 0 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v156 = int32(_a_F_exec_stmts_1)
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	v163 = F_SPI_plan_get_cached_plan(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L136
	}
L136:
	;
	if v163 == int32(0) {
		goto L88
	} else {
		goto L137
	}
L137:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v167 == int32(0) {
		goto L88
	} else {
		goto L138
	}
L138:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v170 != int32(1) {
		goto L88
	} else {
		goto L139
	}
L139:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+100))
	if v175 == int32(0) {
		goto L87
	} else {
		goto L140
	}
L140:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	if v178 != int32(213) {
		goto L87
	} else {
		goto L141
	}
L141:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	v183 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v182)+4)))
	v184 = F_SearchSysCache1(m, int32(47), v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L9
	} else {
		goto L142
	}
L142:
	;
	if v184 == int32(0) {
		goto L86
	} else {
		goto L143
	}
L143:
	;
	v194 = F_get_func_arg_info(m, v184, v32+int32(616), v32+int32(592), v32+int32(608))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	F_ReleaseCatCache(m, v184)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L145
	}
L145:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+48))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v200
	v203 = F_palloc0(m, int32(40))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = int32(_a_F_exec_stmts_2)
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = int32(1)
	v212 = F_palloc_mul(m, int32(4), v194)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+36)) = v212
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v217
	v219 = int32(0)
	if v219 < v194 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v225 = int32(0)
	v230 = v219
	goto L151
L149:
	;
	v298 = v219
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+28)) = v298
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[5]))
	F_ReleaseCachedPlan(m, v163, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L9
	} else {
		goto L159
	}
L151:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	if v252 == int32(0) {
		v285 = v230
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v298 = v285
	goto L150
L153:
	;
	v289 = v225 + int32(1)
	if v289 != v194 {
		v225 = v289
		v230 = v285
		goto L151
	} else {
		goto L158
	}
L154:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225+v252))))
	v258 = v256 - int32(98)
	v259 = int32(0)
	if base.B2i32(v258 == v259)|base.B2i32(v258 == int32(13)) == v259 {
		v285 = v230
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v267 = v230 << (uint(int32(2)) % 32)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267+v269)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if v272 != int32(8) {
		goto L132
	} else {
		goto L156
	}
L156:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)+8))
	v277 = v275 - int32(1)
	F_exec_check_assignable(m, l0, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v203)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v280+v267))) = v277
	v285 = v230 + int32(1)
	goto L153
L158:
	;
	goto L152
L159:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v203
	goto L133
L160:
	;
	v392 = int32(0)
	goto L130
L161:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L9
	} else {
		goto L165
	}
L162:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v359+v225<<(uint(int32(2))%32))))
	if v365 == int32(0) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	if v368 != 0 {
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
	v376 = m.ExcPending
	if v376 != 0 {
		goto L9
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v225 + int32(1)
	F_errmsg(m, int32(_a_F_exec_stmts_3), v32+int32(48))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L9
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2419), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
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
	if v412 < int32(0) {
		goto L84
	} else {
		goto L170
	}
L170:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[3]))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+44))
	if v418 != v395 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = int64(0)
	F_plpgsql_create_econtext(m, l0)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L9
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v425 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmts[6]))
	if base.Ui64(int64(1)) < base.Ui64(v425) {
		goto L82
	} else {
		goto L175
	}
L174:
	;
	goto L173
L175:
	;
	if base.I32_wrap_i64(v425) == int32(1) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+16)))
	if v431 == int32(0) {
		goto L83
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v443 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+4))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	F_exec_move_row(m, l0, v434, v438, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L9
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	F_SPI_freetuptable(m, v443)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L9
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v446 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v446
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v449 != 0 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	goto L183
L185:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+20))
	F_MemoryContextReset(m, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L9
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v454 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	F_SPI_freetuptable(m, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L9
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	v5234 = v446
	goto L21
L190:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v460 == int32(0) {
		goto L81
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v463 == int32(0) {
		goto L22
	} else {
		goto L194
	}
L193:
	;
	goto L192
L194:
	;
	v466 = int32(0)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v467 <= v466 {
		goto L22
	} else {
		goto L195
	}
L195:
	;
	v472 = v466
	goto L196
L196:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v463)+12))
	v501 = int32(2)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v500+v472<<(uint(v501)%32))))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v499+v505<<(uint(v501)%32))))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	switch v510 {
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
	v961 = v472 + int32(1)
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v961 < v962 {
		v472 = v961
		goto L196
	} else {
		goto L288
	}
L199:
	;
	v925 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	F_exec_assign_value(m, l0, v509, v925, int32(0), int32(20), int32(-1))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L9
	} else {
		goto L287
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L9
	} else {
		goto L284
	}
L201:
	;
	v760 = int32(_a_F_exec_stmts_1)
	v761 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v763)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v764
	v767 = int32(_a_F_exec_stmts_6)
	v769 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[8]))
	v770 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[8])) = v769 + v770
	v773 = int32(_a_F_exec_stmts_7)
	v775 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9]))
	v777 = v775 + v770
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9])) = v777
	if v777 < int32(5) {
		goto L266
	} else {
		goto L267
	}
L202:
	;
	v740 = int32(_a_F_exec_stmts_1)
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)+60))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v746
	if v743 != 0 {
		goto L260
	} else {
		goto L261
	}
L203:
	;
	v720 = int32(_a_F_exec_stmts_1)
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+64))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v725)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v726
	if v723 != 0 {
		goto L255
	} else {
		goto L256
	}
L204:
	;
	v700 = int32(_a_F_exec_stmts_1)
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)+32))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v706
	if v703 != 0 {
		goto L250
	} else {
		goto L251
	}
L205:
	;
	v680 = int32(_a_F_exec_stmts_1)
	v681 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)+72))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v685)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v686
	if v683 != 0 {
		goto L245
	} else {
		goto L246
	}
L206:
	;
	v660 = int32(_a_F_exec_stmts_1)
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v662)+76))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v665)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v666
	if v663 != 0 {
		goto L240
	} else {
		goto L241
	}
L207:
	;
	v640 = int32(_a_F_exec_stmts_1)
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+68))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v645)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v646
	if v643 != 0 {
		goto L235
	} else {
		goto L236
	}
L208:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+28))
	v580 = int32(_a_F_exec_stmts_8)
	v581 = int32(63)
	v583 = int32(48)
	v584 = v579&v581 + v583
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[10])) = uint8(v584)
	v592 = int32(base.Ui32(v579)>>(uint(int32(24))%32))&v581 + v583
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[11])) = uint8(v592)
	v600 = int32(base.Ui32(v579)>>(uint(int32(18))%32))&v581 + v583
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[12])) = uint8(v600)
	v608 = int32(base.Ui32(v579)>>(uint(int32(12))%32))&v581 + v583
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[13])) = uint8(v608)
	v616 = int32(base.Ui32(v579)>>(uint(int32(6))%32))&v581 + v583
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[14])) = uint8(v616)
	v619 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[15])) = uint8(v619)
	goto L229
L209:
	;
	v558 = int32(_a_F_exec_stmts_1)
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+44))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v564
	if v561 != 0 {
		goto L224
	} else {
		goto L225
	}
L210:
	;
	v538 = int32(_a_F_exec_stmts_1)
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+36))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v544
	if v541 != 0 {
		goto L219
	} else {
		goto L220
	}
L211:
	;
	v518 = int32(_a_F_exec_stmts_1)
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)+48))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v524
	if v521 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v512 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v511)+36)))
	F_exec_assign_value(m, l0, v509, v512, int32(0), int32(26), int32(-1))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L9
	} else {
		goto L213
	}
L213:
	;
	goto L198
L214:
	;
	v527 = v521
	goto L216
L215:
	;
	v527 = int32(_a_F_exec_stmts_9)
	goto L216
L216:
	;
	v528 = F_cstring_to_text(m, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L9
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v519
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v528), int32(0), int32(25), int32(-1))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L9
	} else {
		goto L218
	}
L218:
	;
	goto L198
L219:
	;
	v547 = v541
	goto L221
L220:
	;
	v547 = int32(_a_F_exec_stmts_9)
	goto L221
L221:
	;
	v548 = F_cstring_to_text(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L9
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v539
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v548), int32(0), int32(25), int32(-1))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L9
	} else {
		goto L223
	}
L223:
	;
	goto L198
L224:
	;
	v567 = v561
	goto L226
L225:
	;
	v567 = int32(_a_F_exec_stmts_9)
	goto L226
L226:
	;
	v568 = F_cstring_to_text(m, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L9
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v559
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v568), int32(0), int32(25), int32(-1))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L9
	} else {
		goto L228
	}
L228:
	;
	goto L198
L229:
	;
	v622 = int32(_a_F_exec_stmts_1)
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v625)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v626
	goto L230
L230:
	;
	goto L232
L232:
	;
	v630 = F_cstring_to_text(m, v580)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L9
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v623
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v630), int32(0), int32(25), int32(-1))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L9
	} else {
		goto L234
	}
L234:
	;
	goto L198
L235:
	;
	v649 = v643
	goto L237
L236:
	;
	v649 = int32(_a_F_exec_stmts_9)
	goto L237
L237:
	;
	v650 = F_cstring_to_text(m, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L9
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v641
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v650), int32(0), int32(25), int32(-1))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L9
	} else {
		goto L239
	}
L239:
	;
	goto L198
L240:
	;
	v669 = v663
	goto L242
L241:
	;
	v669 = int32(_a_F_exec_stmts_9)
	goto L242
L242:
	;
	v670 = F_cstring_to_text(m, v669)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L9
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v661
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v670), int32(0), int32(25), int32(-1))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L9
	} else {
		goto L244
	}
L244:
	;
	goto L198
L245:
	;
	v689 = v683
	goto L247
L246:
	;
	v689 = int32(_a_F_exec_stmts_9)
	goto L247
L247:
	;
	v690 = F_cstring_to_text(m, v689)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L9
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v681
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v690), int32(0), int32(25), int32(-1))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L9
	} else {
		goto L249
	}
L249:
	;
	goto L198
L250:
	;
	v709 = v703
	goto L252
L251:
	;
	v709 = int32(_a_F_exec_stmts_9)
	goto L252
L252:
	;
	v710 = F_cstring_to_text(m, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v701
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v710), int32(0), int32(25), int32(-1))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L9
	} else {
		goto L254
	}
L254:
	;
	goto L198
L255:
	;
	v729 = v723
	goto L257
L256:
	;
	v729 = int32(_a_F_exec_stmts_9)
	goto L257
L257:
	;
	v730 = F_cstring_to_text(m, v729)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L9
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v721
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v730), int32(0), int32(25), int32(-1))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L9
	} else {
		goto L259
	}
L259:
	;
	goto L198
L260:
	;
	v749 = v743
	goto L262
L261:
	;
	v749 = int32(_a_F_exec_stmts_9)
	goto L262
L262:
	;
	v750 = F_cstring_to_text(m, v749)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L9
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v741
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v750), int32(0), int32(25), int32(-1))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L9
	} else {
		goto L264
	}
L264:
	;
	goto L198
L265:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v894)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v895
	if v851 != 0 {
		goto L279
	} else {
		goto L280
	}
L266:
	;
	v781 = int32(100)
	v782 = v777 * v781
	base.MemoryFill(m, v782+int32(_a_F_exec_stmts_10), int32(0), v781)
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v782)+uint32(_c_F_exec_stmts[17]))) = v791
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v782)+uint32(_c_F_exec_stmts[18]))) = v796
	v800 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[19]))
	if v800 != 0 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9])) = int32(-1)
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L9
	} else {
		goto L276
	}
L269:
	;
	v805 = v800
	goto L272
L270:
	;
	v849 = v775
	v851 = int32(0)
	v873 = v769
	goto L271
L271:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[8])) = v873
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9])) = v849
	goto L265
L272:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v805)+8))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v805)+4))
	m.T0[v831].(func(*base.Module, int32))(m, v830)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L9
	} else {
		goto L274
	}
L273:
	;
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[9]))
	v837 = int32(1)
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v782)+uint32(_c_F_exec_stmts[20])))
	v841 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[8]))
	v849 = v836 - v837
	v851 = v839
	v873 = v841 - v837
	goto L271
L274:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v805)))
	if v834 != 0 {
		v805 = v834
		goto L272
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_11), int32(0))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L9
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_12), int32(783), int32(_a_F_exec_stmts_13))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L9
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	v898 = v851
	goto L281
L280:
	;
	v898 = int32(_a_F_exec_stmts_9)
	goto L281
L281:
	;
	v899 = F_cstring_to_text(m, v898)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L9
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v761
	F_exec_assign_value(m, l0, v509, base.I64_extend_i32_u(v899), int32(0), int32(25), int32(-1))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L9
	} else {
		goto L283
	}
L283:
	;
	goto L198
L284:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = v913
	F_errmsg_internal(m, int32(_a_F_exec_stmts_14), v32+int32(80))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L9
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2541), int32(_a_F_exec_stmts_15))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L9
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L287:
	;
	goto L198
L288:
	;
	goto L197
L289:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v977 = F_exec_cast_value(m, l0, v971, v966, v973, v974, int32(16), int32(-1))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L9
	} else {
		goto L290
	}
L290:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v979 != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	F_SPI_freetuptable(m, v979)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L9
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v984 != 0 {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	goto L293
L295:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)+20))
	F_MemoryContextReset(m, v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L9
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	if v990|base.B2i32(v977 == int64(0)) == int32(0) {
		v5187 = v89 + int32(16)
		goto L23
	} else {
		goto L299
	}
L298:
	;
	goto L297
L299:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v996 == int32(0) {
		goto L24
	} else {
		goto L300
	}
L300:
	;
	v999 = int32(0)
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v996)+4))
	if v1000 <= v999 {
		goto L24
	} else {
		goto L301
	}
L301:
	;
	v1005 = v999
	goto L302
L302:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v996)+12))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1032+v1005<<(uint(int32(2))%32))))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+4))
	v1039 = v32 + int32(608)
	v1044 = F_exec_eval_expr(m, l0, v1037, v1039, v32+int32(616), v32+int32(592))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L9
	} else {
		goto L304
	}
L303:
	;
	v5187 = v1036 + int32(8)
	goto L23
L304:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v1050 = F_exec_cast_value(m, l0, v1044, v1039, v1046, v1047, int32(16), int32(-1))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L9
	} else {
		goto L305
	}
L305:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1052 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	F_SPI_freetuptable(m, v1052)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L9
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1057 != 0 {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	goto L308
L310:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+20))
	F_MemoryContextReset(m, v1058)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L9
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	v1062 = int32(0)
	if base.B2i32(v1061 == v1062)&base.B2i32(v1050 != int64(0)) == v1062 {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	goto L312
L314:
	;
	v1070 = v1005 + int32(1)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v996)+4))
	if v1071 <= v1070 {
		goto L24
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	goto L303
L317:
	;
	v1005 = v1070
	goto L302
L318:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v1125 == int32(0) {
		goto L27
	} else {
		goto L336
	}
L319:
	;
	v1122 = int32(0)
	goto L318
L320:
	;
	goto L321
L321:
	;
	v1085 = F_exec_eval_expr(m, l0, v1075, v32+int32(608), v32+int32(616), v32+int32(592))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L9
	} else {
		goto L322
	}
L322:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1089+v1090<<(uint(int32(2))%32))))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+24))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+4))
	if v1088 == v1096 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	F_exec_assign_value(m, l0, v1094, v1085, v1106, v1088, v1087)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L9
	} else {
		goto L329
	}
L324:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+24))
	if v1098 == v1087 {
		goto L323
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+44))
	v1103 = F_plpgsql_build_datatype(m, v1088, v1087, v1101, int32(0))
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L9
	} else {
		goto L328
	}
L327:
	;
	goto L326
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1094)+24)) = v1103
	goto L323
L329:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1109 != 0 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	F_SPI_freetuptable(m, v1109)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L9
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v1112
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1114 == v1112 {
		v1122 = v1094
		goto L318
	} else {
		goto L334
	}
L333:
	;
	goto L332
L334:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+20))
	F_MemoryContextReset(m, v1117)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L9
	} else {
		goto L335
	}
L335:
	;
	v1122 = v1094
	goto L318
L336:
	;
	v1128 = int32(0)
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1129 <= v1128 {
		goto L27
	} else {
		goto L337
	}
L337:
	;
	v1134 = v1128
	goto L338
L338:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+12))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1161+v1134<<(uint(int32(2))%32))))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+4))
	v1168 = v32 + int32(608)
	v1173 = F_exec_eval_expr(m, l0, v1166, v1168, v32+int32(616), v32+int32(592))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L9
	} else {
		goto L340
	}
L339:
	;
	if v1122 != 0 {
		goto L354
	} else {
		goto L355
	}
L340:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v1179 = F_exec_cast_value(m, l0, v1173, v1168, v1175, v1176, int32(16), int32(-1))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L9
	} else {
		goto L341
	}
L341:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1181 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	F_SPI_freetuptable(m, v1181)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L9
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1186 != 0 {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	goto L344
L346:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+20))
	F_MemoryContextReset(m, v1187)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L9
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	v1191 = int32(0)
	if base.B2i32(v1190 == v1191)&base.B2i32(v1179 != int64(0)) == v1191 {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	goto L348
L350:
	;
	v1199 = v1134 + int32(1)
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+4))
	if v1200 <= v1199 {
		goto L27
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	goto L339
L353:
	;
	v1134 = v1199
	goto L338
L354:
	;
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122)+49)))
	if v1202 != int32(1) {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	goto L356
L356:
	;
	v5103 = v1165 + int32(8)
	goto L26
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1122)+52)) = int32(0)
	v1229 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1122)+48)) = uint16(v1229)
	*(*int64)(unsafe.Add(mBase, uint32(v1122)+40)) = int64(0)
	goto L356
L358:
	;
	v1205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122)+48)))
	if v1205 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+40))
	F_pfree(m, v1222)
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L9
	} else {
		goto L365
	}
L360:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+24))
	v1207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1206)+12)))
	if v1207 != int32(_a_F_exec_stmts_16) {
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v1210 = *(*int64)(unsafe.Add(mBase, uint32(v1122)+40))
	v1211 = base.I32_wrap_i64(v1210)
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
	if v1212 != int32(1) {
		goto L359
	} else {
		goto L362
	}
L362:
	;
	v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211)+1)))
	if v1215 != int32(3) {
		goto L359
	} else {
		goto L363
	}
L363:
	;
	F_DeleteExpandedObject(m, v1210)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L9
	} else {
		goto L364
	}
L364:
	;
	goto L357
L365:
	;
	goto L357
L366:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1267 = F_exec_stmts(m, l0, v1266)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L9
	} else {
		goto L369
	}
L368:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1271 == int32(0) {
		goto L366
	} else {
		goto L370
	}
L369:
	;
	switch v1267 - int32(1) {
	case 0:
		goto L29
	case 1:
		v5234 = v1267
		goto L21
	case 2:
		goto L368
	default:
		goto L366
	}
L370:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1274 == int32(0) {
		v5234 = v1267
		goto L21
	} else {
		goto L371
	}
L371:
	;
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274))))
	v1282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271))))
	if base.B2i32(v1279 == int32(0))|base.B2i32(v1279 != v1282) != 0 {
		v1300 = v1279
		v1301 = v1282
		goto L373
	} else {
		goto L374
	}
L372:
	;
	if v1300-v1301 != 0 {
		v5234 = v1267
		goto L21
	} else {
		goto L379
	}
L373:
	;
	goto L372
L374:
	;
	v1285 = v1274
	v1286 = v1271
	goto L375
L375:
	;
	v1289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1286)+1)))
	v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1285)+1)))
	if v1290 == int32(0) {
		v1300 = v1290
		v1301 = v1289
		goto L373
	} else {
		goto L377
	}
L376:
	;
	v1300 = v1290
	v1301 = v1289
	goto L373
L377:
	;
	v1293 = int32(1)
	if v1290 == v1289 {
		v1285 = v1285 + v1293
		v1286 = v1286 + v1293
		goto L375
	} else {
		goto L378
	}
L378:
	;
	goto L376
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	goto L366
L380:
	;
	v1337 = int32(0)
	goto L382
L382:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1366 = v32 + int32(608)
	v1371 = F_exec_eval_expr(m, l0, v1364, v1366, v32+int32(616), v32+int32(592))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L9
	} else {
		goto L384
	}
L384:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v1377 = F_exec_cast_value(m, l0, v1371, v1366, v1373, v1374, int32(16), int32(-1))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L9
	} else {
		goto L385
	}
L385:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1379 != 0 {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	F_SPI_freetuptable(m, v1379)
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L9
	} else {
		goto L389
	}
L387:
	;
	goto L388
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1384 != 0 {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	goto L388
L390:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1384)+20))
	F_MemoryContextReset(m, v1385)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L9
	} else {
		goto L393
	}
L391:
	;
	goto L392
L392:
	;
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	if v1388|base.B2i32(v1377 == int64(0)) != 0 {
		v5234 = v1337
		goto L21
	} else {
		goto L394
	}
L393:
	;
	goto L392
L394:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v1393 = F_exec_stmts(m, l0, v1392)
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L9
	} else {
		goto L397
	}
L395:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1433 == int32(0) {
		goto L380
	} else {
		goto L410
	}
L396:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1397 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	switch v1393 - int32(1) {
	case 0:
		goto L396
	case 1:
		v5234 = v1393
		goto L21
	case 2:
		goto L395
	default:
		v1337 = v1393
		goto L382
	}
L398:
	;
	v5234 = int32(0)
	goto L21
L399:
	;
	goto L400
L400:
	;
	v1401 = int32(1)
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1402 == int32(0) {
		v5234 = v1401
		goto L21
	} else {
		goto L401
	}
L401:
	;
	v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1402))))
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1397))))
	if base.B2i32(v1407 == int32(0))|base.B2i32(v1407 != v1410) != 0 {
		v1428 = v1407
		v1429 = v1410
		goto L403
	} else {
		goto L404
	}
L402:
	;
	if v1428-v1429 == int32(0) {
		goto L28
	} else {
		goto L409
	}
L403:
	;
	goto L402
L404:
	;
	v1413 = v1402
	v1414 = v1397
	goto L405
L405:
	;
	v1417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1414)+1)))
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1413)+1)))
	if v1418 == int32(0) {
		v1428 = v1418
		v1429 = v1417
		goto L403
	} else {
		goto L407
	}
L406:
	;
	v1428 = v1418
	v1429 = v1417
	goto L403
L407:
	;
	v1421 = int32(1)
	if v1418 == v1417 {
		v1413 = v1413 + v1421
		v1414 = v1414 + v1421
		goto L405
	} else {
		goto L408
	}
L408:
	;
	goto L406
L409:
	;
	v5234 = v1401
	goto L21
L410:
	;
	v1436 = int32(3)
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1437 == int32(0) {
		v5234 = v1436
		goto L21
	} else {
		goto L411
	}
L411:
	;
	v1442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1437))))
	v1445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1433))))
	if base.B2i32(v1442 == int32(0))|base.B2i32(v1442 != v1445) != 0 {
		v1463 = v1442
		v1464 = v1445
		goto L413
	} else {
		goto L414
	}
L412:
	;
	if v1463-v1464 != 0 {
		v5234 = v1436
		goto L21
	} else {
		goto L419
	}
L413:
	;
	goto L412
L414:
	;
	v1448 = v1437
	v1449 = v1433
	goto L415
L415:
	;
	v1452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1449)+1)))
	v1453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1448)+1)))
	if v1453 == int32(0) {
		v1463 = v1453
		v1464 = v1452
		goto L413
	} else {
		goto L417
	}
L416:
	;
	v1463 = v1453
	v1464 = v1452
	goto L413
L417:
	;
	v1456 = int32(1)
	if v1453 == v1452 {
		v1448 = v1448 + v1456
		v1449 = v1449 + v1456
		goto L415
	} else {
		goto L418
	}
L418:
	;
	goto L416
L419:
	;
	v1466 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1466
	v1337 = v1466
	goto L382
L420:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+24))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1487)+4))
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1487)+24))
	v1490 = F_exec_cast_value(m, l0, v1483, v1478, v1485, v1486, v1488, v1489)
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L9
	} else {
		goto L421
	}
L421:
	;
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	if v1492 == int32(1) {
		goto L80
	} else {
		goto L422
	}
L422:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1495 != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	F_SPI_freetuptable(m, v1495)
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L9
	} else {
		goto L426
	}
L424:
	;
	goto L425
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1500 != 0 {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	goto L425
L427:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1500)+20))
	F_MemoryContextReset(m, v1501)
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L9
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1506 = v32 + int32(608)
	v1511 = F_exec_eval_expr(m, l0, v1504, v1506, v32+int32(616), v32+int32(592))
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L9
	} else {
		goto L431
	}
L430:
	;
	goto L429
L431:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+24))
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+4))
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+24))
	v1518 = F_exec_cast_value(m, l0, v1511, v1506, v1513, v1514, v1516, v1517)
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L9
	} else {
		goto L432
	}
L432:
	;
	v1520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	if v1520 == int32(1) {
		goto L79
	} else {
		goto L433
	}
L433:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1523 != 0 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	F_SPI_freetuptable(m, v1523)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L9
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1528 != 0 {
		goto L438
	} else {
		goto L439
	}
L437:
	;
	goto L436
L438:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1528)+20))
	F_MemoryContextReset(m, v1529)
	mBase = m.M
	v1531 = m.ExcPending
	if v1531 != 0 {
		goto L9
	} else {
		goto L441
	}
L439:
	;
	goto L440
L440:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v1532 == int32(0) {
		goto L443
	} else {
		goto L444
	}
L441:
	;
	goto L440
L442:
	;
	v1569 = base.I32_wrap_i64(v1490)
	v1570 = base.I32_wrap_i64(v1518)
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	if v1571 != 0 {
		goto L459
	} else {
		goto L460
	}
L443:
	;
	v1568 = int32(1)
	goto L442
L444:
	;
	goto L445
L445:
	;
	v1537 = v32 + int32(608)
	v1542 = F_exec_eval_expr(m, l0, v1532, v1537, v32+int32(616), v32+int32(592))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L9
	} else {
		goto L446
	}
L446:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1475)+24))
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1546)+4))
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1546)+24))
	v1549 = F_exec_cast_value(m, l0, v1542, v1537, v1544, v1545, v1547, v1548)
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L9
	} else {
		goto L447
	}
L447:
	;
	v1551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	if v1551 == int32(1) {
		goto L78
	} else {
		goto L448
	}
L448:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1554 != 0 {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	F_SPI_freetuptable(m, v1554)
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L9
	} else {
		goto L452
	}
L450:
	;
	goto L451
L451:
	;
	v1557 = base.I32_wrap_i64(v1549)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1560 != 0 {
		goto L453
	} else {
		goto L454
	}
L452:
	;
	goto L451
L453:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+20))
	F_MemoryContextReset(m, v1561)
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L9
	} else {
		goto L456
	}
L454:
	;
	goto L455
L455:
	;
	if v1557 <= int32(0) {
		goto L77
	} else {
		goto L457
	}
L456:
	;
	goto L455
L457:
	;
	v1568 = v1557
	goto L442
L458:
	;
	v1575 = int32(0)
	F_assign_simple_var(m, l0, v1475, base.I64_extend32_s(v1490), v1575, v1575)
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L9
	} else {
		goto L464
	}
L459:
	;
	if v1570 <= v1569 {
		goto L458
	} else {
		goto L462
	}
L460:
	;
	goto L461
L461:
	;
	if v1570 < v1569 {
		goto L31
	} else {
		goto L463
	}
L462:
	;
	goto L31
L463:
	;
	goto L458
L464:
	;
	v1579 = int64(1)
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	v1581 = F_exec_stmts(m, l0, v1580)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L9
	} else {
		goto L467
	}
L465:
	;
	v1626 = v1568 ^ int32(2147483647)
	v1628 = v1568 | int32(-2147483648)
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	if v1629 != 0 {
		goto L479
	} else {
		goto L480
	}
L466:
	;
	v1585 = int32(0)
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1586 == v1585 {
		v1622 = v1585
		goto L465
	} else {
		goto L468
	}
L467:
	;
	switch v1581 - int32(1) {
	case 0:
		goto L32
	case 1:
		v4904 = v1581
		v4928 = v1579
		goto L30
	case 2:
		goto L466
	default:
		v1622 = v1581
		goto L465
	}
L468:
	;
	v1589 = int32(3)
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1590 == int32(0) {
		v4904 = v1589
		v4928 = v1579
		goto L30
	} else {
		goto L469
	}
L469:
	;
	v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1590))))
	v1598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1586))))
	if base.B2i32(v1595 == int32(0))|base.B2i32(v1595 != v1598) != 0 {
		v1616 = v1595
		v1617 = v1598
		goto L471
	} else {
		goto L472
	}
L470:
	;
	if v1616-v1617 != 0 {
		v4904 = v1589
		v4928 = v1579
		goto L30
	} else {
		goto L477
	}
L471:
	;
	goto L470
L472:
	;
	v1601 = v1590
	v1602 = v1586
	goto L473
L473:
	;
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1602)+1)))
	v1606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1601)+1)))
	if v1606 == int32(0) {
		v1616 = v1606
		v1617 = v1605
		goto L471
	} else {
		goto L475
	}
L474:
	;
	v1616 = v1606
	v1617 = v1605
	goto L471
L475:
	;
	v1609 = int32(1)
	if v1606 == v1605 {
		v1601 = v1601 + v1609
		v1602 = v1602 + v1609
		goto L473
	} else {
		goto L476
	}
L476:
	;
	goto L474
L477:
	;
	v1619 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1619
	v1622 = v1619
	goto L465
L478:
	;
	v1639 = v1622
	v1644 = v1636
	v1647 = v1635
	goto L484
L479:
	;
	if v1569 < v1628 {
		v4904 = v1622
		v4928 = v1579
		goto L30
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	if v1626 < v1569 {
		v4904 = v1622
		v4928 = v1579
		goto L30
	} else {
		goto L483
	}
L482:
	;
	v1635 = int32(1)
	v1636 = v1569 - v1568
	goto L478
L483:
	;
	v1635 = v1629
	v1636 = v1569 + v1568
	goto L478
L484:
	;
	if v1647 != 0 {
		goto L487
	} else {
		goto L488
	}
L486:
	;
	v1669 = int32(0)
	F_assign_simple_var(m, l0, v1475, base.I64_extend_i32_s(v1644), v1669, v1669)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L9
	} else {
		goto L492
	}
L487:
	;
	if v1570 <= v1644 {
		goto L486
	} else {
		goto L490
	}
L488:
	;
	goto L489
L489:
	;
	if v1570 < v1644 {
		v4904 = v1639
		v4928 = v1579
		goto L30
	} else {
		goto L491
	}
L490:
	;
	v4904 = v1639
	v4928 = v1579
	goto L30
L491:
	;
	goto L486
L492:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	v1674 = F_exec_stmts(m, l0, v1673)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L9
	} else {
		goto L495
	}
L493:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	if v1718 != 0 {
		goto L506
	} else {
		goto L507
	}
L494:
	;
	v1678 = int32(0)
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1679 == v1678 {
		v1715 = v1678
		goto L493
	} else {
		goto L496
	}
L495:
	;
	switch v1674 - int32(1) {
	case 0:
		goto L32
	case 1:
		v4904 = v1674
		v4928 = v1579
		goto L30
	case 2:
		goto L494
	default:
		v1715 = v1674
		goto L493
	}
L496:
	;
	v1682 = int32(3)
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v1683 == int32(0) {
		v4904 = v1682
		v4928 = v1579
		goto L30
	} else {
		goto L497
	}
L497:
	;
	v1688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1683))))
	v1691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1679))))
	if base.B2i32(v1688 == int32(0))|base.B2i32(v1688 != v1691) != 0 {
		v1709 = v1688
		v1710 = v1691
		goto L499
	} else {
		goto L500
	}
L498:
	;
	if v1709-v1710 != 0 {
		v4904 = v1682
		v4928 = v1579
		goto L30
	} else {
		goto L505
	}
L499:
	;
	goto L498
L500:
	;
	v1694 = v1683
	v1695 = v1679
	goto L501
L501:
	;
	v1698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1695)+1)))
	v1699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694)+1)))
	if v1699 == int32(0) {
		v1709 = v1699
		v1710 = v1698
		goto L499
	} else {
		goto L503
	}
L502:
	;
	v1709 = v1699
	v1710 = v1698
	goto L499
L503:
	;
	v1702 = int32(1)
	if v1699 == v1698 {
		v1694 = v1694 + v1702
		v1695 = v1695 + v1702
		goto L501
	} else {
		goto L504
	}
L504:
	;
	goto L502
L505:
	;
	v1712 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1712
	v1715 = v1712
	goto L493
L506:
	;
	if v1644 < v1628 {
		v4904 = v1715
		v4928 = v1579
		goto L30
	} else {
		goto L509
	}
L507:
	;
	goto L508
L508:
	;
	if v1626 < v1644 {
		v4904 = v1715
		v4928 = v1579
		goto L30
	} else {
		goto L510
	}
L509:
	;
	v1639 = v1715
	v1644 = v1644 - v1568
	v1647 = int32(1)
	goto L484
L510:
	;
	v1639 = v1715
	v1644 = v1568 + v1644
	v1647 = v1718
	goto L484
L511:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v1731 = F_exec_for_query(m, l0, v89, v1729, int32(1))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L9
	} else {
		goto L512
	}
L512:
	;
	F_SPI_cursor_close(m, v1729)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L9
	} else {
		goto L513
	}
L513:
	;
	v5234 = v1731
	goto L21
L514:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1746 == int32(0) {
		goto L517
	} else {
		goto L518
	}
L515:
	;
	v1770 = v1735
	v1771 = v1735
	goto L516
L516:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v1772 != 0 {
		goto L525
	} else {
		goto L526
	}
L517:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v1754 = F_AllocSetContextCreateInternal(m, v1749, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L9
	} else {
		goto L520
	}
L518:
	;
	v1757 = v1746
	goto L519
L519:
	;
	v1758 = int32(_a_F_exec_stmts_1)
	v1759 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1757
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+40))
	v1763 = F_text_to_cstring(m, v1762)
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L9
	} else {
		goto L521
	}
L520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1754
	v1757 = v1754
	goto L519
L521:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1759
	v1767 = F_GetPortalByName(m, v1763)
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L9
	} else {
		goto L522
	}
L522:
	;
	if v1767 != 0 {
		goto L76
	} else {
		goto L523
	}
L523:
	;
	v1770 = v1763
	v1771 = v1757
	goto L516
L524:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+28))
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1802)+24))
	if v1803 == int32(0) {
		goto L531
	} else {
		goto L532
	}
L525:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+32))
	if v1773 < int32(0) {
		goto L75
	} else {
		goto L528
	}
L526:
	;
	goto L527
L527:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+32))
	if int32(0) <= v1797 {
		goto L74
	} else {
		goto L530
	}
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+616)) = int32(16)
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v1783 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+634)) = uint8(v1783)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+628)) = v1772
	*(*int32)(unsafe.Add(mBase, uint32(v32)+620)) = v1782
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1787+v1773<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+636)) = v1791
	F_exec_stmt_execsql(m, l0, v32+int32(616))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L9
	} else {
		goto L529
	}
L529:
	;
	goto L524
L530:
	;
	goto L524
L531:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+36))
	F_exec_prepare_plan(m, l0, v1802, v1806)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L9
	} else {
		goto L534
	}
L532:
	;
	goto L533
L533:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1802)+28))
	if v1809 == int32(0) {
		goto L536
	} else {
		goto L537
	}
L534:
	;
	goto L533
L535:
	;
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1802)+24))
	v1817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	v1818 = F_SPI_cursor_open_internal(m, v1770, v1816, v1815, v1817)
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L9
	} else {
		goto L539
	}
L536:
	;
	v1815 = int32(0)
	goto L535
L537:
	;
	goto L538
L538:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v1813)+20)) = v1802
	v1815 = v1813
	goto L535
L539:
	;
	if v1818 == int32(0) {
		goto L73
	} else {
		goto L540
	}
L540:
	;
	if v1770 == int32(0) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	F_exec_check_assignable(m, l0, v1824)
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L9
	} else {
		goto L544
	}
L542:
	;
	goto L543
L543:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1835 != 0 {
		goto L547
	} else {
		goto L548
	}
L544:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1828 = F_cstring_to_text(m, v1827)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L9
	} else {
		goto L545
	}
L545:
	;
	F_assign_simple_var(m, l0, v1742, base.I64_extend_i32_u(v1828), int32(0), int32(1))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L9
	} else {
		goto L546
	}
L546:
	;
	goto L543
L547:
	;
	F_SPI_freetuptable(m, v1835)
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L9
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1840 != 0 {
		goto L551
	} else {
		goto L552
	}
L550:
	;
	goto L549
L551:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+20))
	F_MemoryContextReset(m, v1841)
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L9
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	if v1771 != 0 {
		goto L555
	} else {
		goto L556
	}
L554:
	;
	goto L553
L555:
	;
	F_MemoryContextReset(m, v1771)
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L9
	} else {
		goto L558
	}
L556:
	;
	goto L557
L557:
	;
	v1847 = F_exec_for_query(m, l0, v89, v1818, int32(0))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L9
	} else {
		goto L559
	}
L558:
	;
	goto L557
L559:
	;
	F_SPI_cursor_close(m, v1818)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L9
	} else {
		goto L560
	}
L560:
	;
	if v1770 != 0 {
		v5234 = v1847
		goto L21
	} else {
		goto L561
	}
L561:
	;
	v1851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1742)+49)))
	if v1851 != int32(1) {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1742)+52)) = int32(0)
	v1878 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1742)+48)) = uint16(v1878)
	*(*int64)(unsafe.Add(mBase, uint32(v1742)+40)) = int64(0)
	v5234 = v1847
	goto L21
L563:
	;
	v1854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1742)+48)))
	if v1854 != 0 {
		goto L564
	} else {
		goto L565
	}
L564:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+40))
	F_pfree(m, v1871)
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L9
	} else {
		goto L570
	}
L565:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+24))
	v1856 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1855)+12)))
	if v1856 != int32(_a_F_exec_stmts_16) {
		goto L564
	} else {
		goto L566
	}
L566:
	;
	v1859 = *(*int64)(unsafe.Add(mBase, uint32(v1742)+40))
	v1860 = base.I32_wrap_i64(v1859)
	v1861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1860))))
	if v1861 != int32(1) {
		goto L564
	} else {
		goto L567
	}
L567:
	;
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1860)+1)))
	if v1864 != int32(3) {
		goto L564
	} else {
		goto L568
	}
L568:
	;
	F_DeleteExpandedObject(m, v1859)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L9
	} else {
		goto L569
	}
L569:
	;
	goto L562
L570:
	;
	goto L562
L571:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v1889
	v1892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+604)))
	if v1892 == int32(1) {
		goto L72
	} else {
		goto L572
	}
L572:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1895 == int32(0) {
		goto L573
	} else {
		goto L574
	}
L573:
	;
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v1903 = F_AllocSetContextCreateInternal(m, v1898, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L9
	} else {
		goto L576
	}
L574:
	;
	v1905 = v1895
	goto L575
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v1905
	v1909 = int32(_a_F_exec_stmts_1)
	v1910 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1905
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v1914 = F_get_element_type(m, v1913)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L9
	} else {
		goto L577
	}
L576:
	;
	v1905 = v1903
	goto L575
L577:
	;
	if v1914 == int32(0) {
		goto L71
	} else {
		goto L578
	}
L578:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v1919 = F_pg_detoast_datum_copy(m, v1918)
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L9
	} else {
		goto L579
	}
L579:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1921 != 0 {
		goto L580
	} else {
		goto L581
	}
L580:
	;
	F_SPI_freetuptable(m, v1921)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L9
	} else {
		goto L583
	}
L581:
	;
	goto L582
L582:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v1926 != 0 {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	goto L582
L584:
	;
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+20))
	F_MemoryContextReset(m, v1927)
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L9
	} else {
		goto L587
	}
L585:
	;
	goto L586
L586:
	;
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v1930 < int32(0) {
		goto L70
	} else {
		goto L588
	}
L587:
	;
	goto L586
L588:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1919)+4))
	if v1933 < v1930 {
		goto L70
	} else {
		goto L589
	}
L589:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1938 = int32(2)
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1936+v1937<<(uint(v1938)%32))))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1941)))
	if base.Ui32(v1938) <= base.Ui32(v1942-int32(1)) {
		goto L590
	} else {
		goto L591
	}
L590:
	;
	v1947 = F_plpgsql_exec_get_datum_type(m, l0, v1941)
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L9
	} else {
		goto L593
	}
L591:
	;
	v1952 = v1930
	v1953 = int32(0)
	goto L592
L592:
	;
	v1954 = int32(0)
	if base.B2i32(v1953 == v1954)&base.B2i32(v1954 < v1952) != 0 {
		goto L69
	} else {
		goto L595
	}
L593:
	;
	v1949 = F_get_element_type(m, v1947)
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L9
	} else {
		goto L594
	}
L594:
	;
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v1952 = v1951
	v1953 = v1949
	goto L592
L595:
	;
	if v1953 != 0 {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	v1960 = v1952
	goto L598
L597:
	;
	v1960 = int32(1)
	goto L598
L598:
	;
	if v1960 == int32(0) {
		goto L68
	} else {
		goto L599
	}
L599:
	;
	v1964 = F_array_create_iterator(m, v1919, v1952, int32(0))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L9
	} else {
		goto L600
	}
L600:
	;
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if int32(0) < v1970 {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	v1973 = v32 + int32(592)
	goto L603
L602:
	;
	v1973 = v1919 + int32(12)
	goto L603
L603:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v1973)))
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v1982 = F_array_iterate(m, v1964, v32+int32(616), v32+int32(604))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L9
	} else {
		goto L605
	}
L604:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1910
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2149
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(v2149)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v2151
	F_MemoryContextReset(m, v1905)
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L9
	} else {
		goto L644
	}
L605:
	;
	if v1982 == int32(0) {
		v2120 = int32(0)
		v2144 = int64(0)
		goto L604
	} else {
		goto L606
	}
L606:
	;
	goto L607
L607:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1910
	v2017 = *(*int64)(unsafe.Add(mBase, uint32(v32)+616))
	v2018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+604)))
	F_exec_assign_value(m, l0, v1941, v2017, v2018, v1974, v1976)
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L9
	} else {
		goto L609
	}
L608:
	;
	v2120 = v2107
	v2144 = v2027
	goto L604
L609:
	;
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if int32(0) < v2021 {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	F_pfree(m, v2024)
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L9
	} else {
		goto L613
	}
L611:
	;
	goto L612
L612:
	;
	v2027 = int64(1)
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v2029 = F_exec_stmts(m, l0, v2028)
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L9
	} else {
		goto L617
	}
L613:
	;
	goto L612
L614:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v1905
	v2116 = F_array_iterate(m, v1964, v32+int32(616), v32+int32(604))
	mBase = m.M
	v2117 = m.ExcPending
	if v2117 != 0 {
		goto L9
	} else {
		goto L642
	}
L615:
	;
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2070 == int32(0) {
		goto L630
	} else {
		goto L631
	}
L616:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2033 == int32(0) {
		goto L618
	} else {
		goto L619
	}
L617:
	;
	switch v2029 - int32(1) {
	case 0:
		goto L616
	case 1:
		v2120 = v2029
		v2144 = v2027
		goto L604
	case 2:
		goto L615
	default:
		v2107 = v2029
		goto L614
	}
L618:
	;
	v2120 = int32(0)
	v2144 = v2027
	goto L604
L619:
	;
	goto L620
L620:
	;
	v2037 = int32(1)
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2038 == int32(0) {
		v2120 = v2037
		v2144 = v2027
		goto L604
	} else {
		goto L621
	}
L621:
	;
	v2043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2038))))
	v2046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2033))))
	if base.B2i32(v2043 == int32(0))|base.B2i32(v2043 != v2046) != 0 {
		v2064 = v2043
		v2065 = v2046
		goto L623
	} else {
		goto L624
	}
L622:
	;
	if v2064-v2065 != 0 {
		v2120 = v2037
		v2144 = v2027
		goto L604
	} else {
		goto L629
	}
L623:
	;
	goto L622
L624:
	;
	v2049 = v2038
	v2050 = v2033
	goto L625
L625:
	;
	v2053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2050)+1)))
	v2054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2049)+1)))
	if v2054 == int32(0) {
		v2064 = v2054
		v2065 = v2053
		goto L623
	} else {
		goto L627
	}
L626:
	;
	v2064 = v2054
	v2065 = v2053
	goto L623
L627:
	;
	v2057 = int32(1)
	if v2054 == v2053 {
		v2049 = v2049 + v2057
		v2050 = v2050 + v2057
		goto L625
	} else {
		goto L628
	}
L628:
	;
	goto L626
L629:
	;
	v2067 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v2067
	v2120 = v2067
	v2144 = v2027
	goto L604
L630:
	;
	v2107 = int32(0)
	goto L614
L631:
	;
	goto L632
L632:
	;
	v2074 = int32(3)
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2075 == int32(0) {
		v2120 = v2074
		v2144 = v2027
		goto L604
	} else {
		goto L633
	}
L633:
	;
	v2080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2075))))
	v2083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2070))))
	if base.B2i32(v2080 == int32(0))|base.B2i32(v2080 != v2083) != 0 {
		v2101 = v2080
		v2102 = v2083
		goto L635
	} else {
		goto L636
	}
L634:
	;
	if v2101-v2102 != 0 {
		v2120 = v2074
		v2144 = v2027
		goto L604
	} else {
		goto L641
	}
L635:
	;
	goto L634
L636:
	;
	v2086 = v2075
	v2087 = v2070
	goto L637
L637:
	;
	v2090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2087)+1)))
	v2091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2086)+1)))
	if v2091 == int32(0) {
		v2101 = v2091
		v2102 = v2090
		goto L635
	} else {
		goto L639
	}
L638:
	;
	v2101 = v2091
	v2102 = v2090
	goto L635
L639:
	;
	v2094 = int32(1)
	if v2091 == v2090 {
		v2086 = v2086 + v2094
		v2087 = v2087 + v2094
		goto L637
	} else {
		goto L640
	}
L640:
	;
	goto L638
L641:
	;
	v2104 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v2104
	v2107 = v2104
	goto L614
L642:
	;
	if v2116 != 0 {
		goto L607
	} else {
		goto L643
	}
L643:
	;
	goto L608
L644:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v2155+v2156<<(uint(int32(2))%32))))
	v2161 = int32(0)
	F_assign_simple_var(m, l0, v2160, v2144, v2161, v2161)
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L9
	} else {
		goto L645
	}
L645:
	;
	v5234 = v2120
	goto L21
L646:
	;
	v2167 = v32 + int32(608)
	v2172 = F_exec_eval_expr(m, l0, v2165, v2167, v32+int32(616), v32+int32(592))
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L9
	} else {
		goto L649
	}
L647:
	;
	goto L648
L648:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v2197
	v2201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+12)))
	if v2201 != 0 {
		goto L660
	} else {
		goto L661
	}
L649:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v2178 = F_exec_cast_value(m, l0, v2172, v2167, v2174, v2175, int32(16), int32(-1))
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L9
	} else {
		goto L650
	}
L650:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v2180 != 0 {
		goto L651
	} else {
		goto L652
	}
L651:
	;
	F_SPI_freetuptable(m, v2180)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L9
	} else {
		goto L654
	}
L652:
	;
	goto L653
L653:
	;
	v2183 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v2183
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v2186 != 0 {
		goto L655
	} else {
		goto L656
	}
L654:
	;
	goto L653
L655:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v2186)+20))
	F_MemoryContextReset(m, v2187)
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L9
	} else {
		goto L658
	}
L656:
	;
	goto L657
L657:
	;
	v2190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	if v2190|base.B2i32(v2178 == int64(0)) != 0 {
		v5234 = v2183
		goto L21
	} else {
		goto L659
	}
L658:
	;
	goto L657
L659:
	;
	goto L648
L660:
	;
	v2202 = int32(1)
	goto L662
L661:
	;
	v2202 = int32(3)
	goto L662
L662:
	;
	v5234 = v2202
	goto L21
L663:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v43))) = int64(0)
	v2207 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v2207)
	v2209 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v2209
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v2209 <= v2211 {
		goto L664
	} else {
		goto L665
	}
L664:
	;
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v2214+v2211<<(uint(int32(2))%32))))
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v2218)))
	switch v2219 {
	case 0:
		goto L669
	case 1, 2:
		goto L668
	default:
		goto L667
	case 4:
		goto L670
	}
L665:
	;
	goto L666
L666:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2271 != 0 {
		goto L681
	} else {
		goto L682
	}
L667:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L9
	} else {
		goto L678
	}
L668:
	;
	F_exec_eval_datum(m, l0, v2218, v39, v32+int32(616), v43, v41)
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L9
	} else {
		goto L677
	}
L669:
	;
	v2222 = *(*int64)(unsafe.Add(mBase, uint32(v2218)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v2222
	v2224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2218)+48)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v2224)
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2218)+24))
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v2226)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v2227
	v2229 = int32(1)
	v2231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v2224&v2229|base.B2i32(v2231 != v2229) != 0 {
		v5234 = v2203
		goto L21
	} else {
		goto L672
	}
L670:
	;
	F_plpgsql_fulfill_promise(m, l0, v2218)
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L9
	} else {
		goto L671
	}
L671:
	;
	goto L669
L672:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L9
	} else {
		goto L673
	}
L673:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v2241 = m.ExcPending
	if v2241 != 0 {
		goto L9
	} else {
		goto L674
	}
L674:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_20), int32(0))
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		goto L9
	} else {
		goto L675
	}
L675:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3286), int32(_a_F_exec_stmts_21))
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L9
	} else {
		goto L676
	}
L676:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L677:
	;
	v5234 = v2203
	goto L21
L678:
	;
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v2218)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+160)) = v2259
	F_errmsg_internal(m, int32(_a_F_exec_stmts_22), v32+int32(160))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L9
	} else {
		goto L679
	}
L679:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3306), int32(_a_F_exec_stmts_21))
	mBase = m.M
	v2270 = m.ExcPending
	if v2270 != 0 {
		goto L9
	} else {
		goto L680
	}
L680:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L681:
	;
	v2274 = F_exec_eval_expr(m, l0, v2271, v41, v39, v32+int32(616))
	mBase = m.M
	v2275 = m.ExcPending
	if v2275 != 0 {
		goto L9
	} else {
		goto L684
	}
L682:
	;
	goto L683
L683:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2300 != int32(2278) {
		v5234 = v2203
		goto L21
	} else {
		goto L693
	}
L684:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v2274
	v2277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v2277 != int32(1) {
		v5234 = v2203
		goto L21
	} else {
		goto L685
	}
L685:
	;
	v2280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v2280 != 0 {
		v5234 = v2203
		goto L21
	} else {
		goto L686
	}
L686:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v2282 = F_type_is_rowtype(m, v2281)
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L9
	} else {
		goto L687
	}
L687:
	;
	if v2282 != 0 {
		v5234 = v2203
		goto L21
	} else {
		goto L688
	}
L688:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L9
	} else {
		goto L689
	}
L689:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L9
	} else {
		goto L690
	}
L690:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_20), int32(0))
	mBase = m.M
	v2294 = m.ExcPending
	if v2294 != 0 {
		goto L9
	} else {
		goto L691
	}
L691:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3329), int32(_a_F_exec_stmts_21))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L9
	} else {
		goto L692
	}
L692:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L693:
	;
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2303)+65)))
	if v2304 == int32(112) {
		v5234 = v2203
		goto L21
	} else {
		goto L694
	}
L694:
	;
	v2307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v2307)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(2278)
	v5234 = v2203
	goto L21
L695:
	;
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v2314 == int32(0) {
		goto L696
	} else {
		goto L697
	}
L696:
	;
	F_exec_init_tuple_store(m, l0)
	mBase = m.M
	v2318 = m.ExcPending
	if v2318 != 0 {
		goto L9
	} else {
		goto L699
	}
L697:
	;
	goto L698
L698:
	;
	v2319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v2319)))
	v2321 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if int32(0) <= v2321 {
		goto L701
	} else {
		goto L702
	}
L699:
	;
	goto L698
L700:
	;
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v2554 != 0 {
		goto L776
	} else {
		goto L777
	}
L701:
	;
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v2328 = *(*int32)(unsafe.Add(mBase, uint32(v2324+v2321<<(uint(int32(2))%32))))
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v2328)))
	switch v2329 {
	case 0:
		goto L707
	case 1:
		goto L705
	case 2:
		goto L706
	default:
		goto L704
	case 4:
		goto L708
	}
L702:
	;
	goto L703
L703:
	;
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2447 == int32(0) {
		goto L62
	} else {
		goto L745
	}
L704:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L9
	} else {
		goto L742
	}
L705:
	;
	v2416 = int32(_a_F_exec_stmts_1)
	v2417 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(v2419)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2420
	v2422 = F_make_tuple_from_row(m, l0, v2328, v2319)
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L9
	} else {
		goto L739
	}
L706:
	;
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v2328)+36))
	if v2376 == int32(0) {
		goto L720
	} else {
		goto L721
	}
L707:
	;
	v2332 = *(*int64)(unsafe.Add(mBase, uint32(v2328)+40))
	v2333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2328)+48)))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+592)) = uint8(v2333)
	if v2320 != int32(1) {
		goto L66
	} else {
		goto L710
	}
L708:
	;
	F_plpgsql_fulfill_promise(m, l0, v2328)
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L9
	} else {
		goto L709
	}
L709:
	;
	goto L707
L710:
	;
	if v2333&int32(1) != 0 {
		v2356 = v2332
		goto L711
	} else {
		goto L712
	}
L711:
	;
	v2358 = v32 + int32(592)
	v2359 = *(*int32)(unsafe.Add(mBase, uint32(v2328)+24))
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2359)+4))
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v2359)+24))
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v2319)))
	v2365 = v2319 + v2362<<(uint(int32(3))%32)
	v2366 = *(*int32)(unsafe.Add(mBase, uint32(v2365)+96))
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v2365)+104))
	v2368 = F_exec_cast_value(m, l0, v2356, v2358, v2360, v2361, v2366, v2367)
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		goto L9
	} else {
		goto L718
	}
L712:
	;
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v2328)+24))
	v2340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2339)+12)))
	if v2340 != int32(_a_F_exec_stmts_16) {
		v2356 = v2332
		goto L711
	} else {
		goto L713
	}
L713:
	;
	v2344 = base.I32_wrap_i64(v2332)
	v2345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2344))))
	if v2345 != int32(1) {
		v2355 = v2332
		goto L715
	} else {
		goto L716
	}
L714:
	;
	v2356 = v2355
	goto L711
L715:
	;
	goto L714
L716:
	;
	v2348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2344)+1)))
	if v2348 != int32(3) {
		v2355 = v2332
		goto L715
	} else {
		goto L717
	}
L717:
	;
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v2344)+2))
	v2355 = base.I64_extend_i32_u(v2351 + int32(18))
	goto L715
L718:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v2368
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_tuplestore_putvalues(m, v2371, v2319, v32+int32(616), v2358)
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L9
	} else {
		goto L719
	}
L719:
	;
	goto L700
L720:
	;
	F_instantiate_empty_record_variable(m, l0, v2328)
	mBase = m.M
	v2380 = m.ExcPending
	if v2380 != 0 {
		goto L9
	} else {
		goto L723
	}
L721:
	;
	v2382 = v2376
	goto L722
L722:
	;
	v2383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2382)+28)))
	if v2383&int32(5) == int32(0) {
		goto L724
	} else {
		goto L725
	}
L723:
	;
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2328)+36))
	v2382 = v2381
	goto L722
L724:
	;
	F_deconstruct_expanded_record(m, v2382)
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L9
	} else {
		goto L727
	}
L725:
	;
	v2391 = v2382
	goto L726
L726:
	;
	v2392 = int32(_a_F_exec_stmts_1)
	v2393 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2395)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2396
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(v2391)+44))
	if v2398 != 0 {
		goto L728
	} else {
		goto L729
	}
L727:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v2328)+36))
	v2391 = v2390
	goto L726
L728:
	;
	v2401 = v2398
	goto L730
L729:
	;
	v2399 = F_expanded_record_fetch_tupdesc(m, v2391)
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		goto L9
	} else {
		goto L731
	}
L730:
	;
	v2403 = F_convert_tuples_by_position(m, v2401, v2319, int32(_a_F_exec_stmts_23))
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L9
	} else {
		goto L732
	}
L731:
	;
	v2401 = v2399
	goto L730
L732:
	;
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(v2328)+36))
	v2406 = F_expanded_record_get_tuple(m, v2405)
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L9
	} else {
		goto L733
	}
L733:
	;
	if v2403 != 0 {
		goto L734
	} else {
		goto L735
	}
L734:
	;
	v2408 = F_execute_attr_map_tuple(m, v2406, v2403)
	mBase = m.M
	v2409 = m.ExcPending
	if v2409 != 0 {
		goto L9
	} else {
		goto L737
	}
L735:
	;
	v2410 = v2406
	goto L736
L736:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_tuplestore_puttuple(m, v2411, v2410)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L9
	} else {
		goto L738
	}
L737:
	;
	v2410 = v2408
	goto L736
L738:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2393
	goto L700
L739:
	;
	if v2422 == int32(0) {
		goto L65
	} else {
		goto L740
	}
L740:
	;
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_tuplestore_puttuple(m, v2426, v2422)
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L9
	} else {
		goto L741
	}
L741:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2417
	goto L700
L742:
	;
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2328)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+176)) = v2435
	F_errmsg_internal(m, int32(_a_F_exec_stmts_22), v32+int32(176))
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		goto L9
	} else {
		goto L743
	}
L743:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3476), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L9
	} else {
		goto L744
	}
L744:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L745:
	;
	v2456 = F_exec_eval_expr(m, l0, v2447, v32+int32(615), v32+int32(608), v32+int32(604))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L9
	} else {
		goto L746
	}
L746:
	;
	v2458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v2458 == int32(1) {
		goto L747
	} else {
		goto L748
	}
L747:
	;
	v2461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+615)))
	if v2461 == int32(0) {
		goto L750
	} else {
		goto L751
	}
L748:
	;
	goto L749
L749:
	;
	if v2320 != int32(1) {
		goto L63
	} else {
		goto L773
	}
L750:
	;
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v2465 = F_type_is_rowtype(m, v2464)
	mBase = m.M
	v2466 = m.ExcPending
	if v2466 != 0 {
		goto L9
	} else {
		goto L753
	}
L751:
	;
	goto L752
L752:
	;
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v2513)+20))
	v2517 = F_MemoryContextAllocZero(m, v2514, v2320<<(uint(int32(3))%32))
	mBase = m.M
	v2518 = m.ExcPending
	if v2518 != 0 {
		goto L9
	} else {
		goto L767
	}
L753:
	;
	if v2465 == int32(0) {
		goto L64
	} else {
		goto L754
	}
L754:
	;
	v2469 = int32(_a_F_exec_stmts_1)
	v2470 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2472)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2473
	v2476 = F_pg_detoast_datum(m, base.I32_wrap_i64(v2456))
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L9
	} else {
		goto L755
	}
L755:
	;
	v2478 = *(*int32)(unsafe.Add(mBase, uint32(v2476)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+632)) = v2476
	v2480 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+628)) = v2480
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+624)) = uint16(v2480)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+620)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+616)) = int32(base.Ui32(v2478) >> (uint(int32(2)) % 32))
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v2476)+8))
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v2476)+4))
	v2491 = F_lookup_rowtype_tupdesc(m, v2489, v2490)
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L9
	} else {
		goto L756
	}
L756:
	;
	v2494 = F_convert_tuples_by_position(m, v2491, v2319, int32(_a_F_exec_stmts_25))
	mBase = m.M
	v2495 = m.ExcPending
	if v2495 != 0 {
		goto L9
	} else {
		goto L757
	}
L757:
	;
	if v2494 != 0 {
		goto L758
	} else {
		goto L759
	}
L758:
	;
	v2498 = F_execute_attr_map_tuple(m, v32+int32(616), v2494)
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L9
	} else {
		goto L761
	}
L759:
	;
	v2502 = v32 + int32(616)
	goto L760
L760:
	;
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_tuplestore_puttuple(m, v2503, v2502)
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		goto L9
	} else {
		goto L762
	}
L761:
	;
	v2502 = v2498
	goto L760
L762:
	;
	v2506 = *(*int32)(unsafe.Add(mBase, uint32(v2491)+12))
	if int32(0) <= v2506 {
		goto L763
	} else {
		goto L764
	}
L763:
	;
	F_DecrTupleDescRefCount(m, v2491)
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L9
	} else {
		goto L766
	}
L764:
	;
	goto L765
L765:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2470
	goto L700
L766:
	;
	goto L765
L767:
	;
	v2519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(v2519)+20))
	v2521 = F_MemoryContextAlloc(m, v2520, v2320)
	mBase = m.M
	v2522 = m.ExcPending
	if v2522 != 0 {
		goto L9
	} else {
		goto L768
	}
L768:
	;
	if v2320 != 0 {
		goto L769
	} else {
		goto L770
	}
L769:
	;
	base.MemoryFill(m, v2521, int32(1), v2320)
	goto L771
L770:
	;
	goto L771
L771:
	;
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_tuplestore_putvalues(m, v2525, v2319, v2517, v2521)
	mBase = m.M
	v2527 = m.ExcPending
	if v2527 != 0 {
		goto L9
	} else {
		goto L772
	}
L772:
	;
	goto L700
L773:
	;
	v2531 = v32 + int32(615)
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(v32)+604))
	v2534 = *(*int32)(unsafe.Add(mBase, uint32(v2319)))
	v2537 = v2319 + v2534<<(uint(int32(3))%32)
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v2537)+96))
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v2537)+104))
	v2540 = F_exec_cast_value(m, l0, v2456, v2531, v2532, v2533, v2538, v2539)
	mBase = m.M
	v2541 = m.ExcPending
	if v2541 != 0 {
		goto L9
	} else {
		goto L774
	}
L774:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+592)) = v2540
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_tuplestore_putvalues(m, v2543, v2319, v32+int32(592), v2531)
	mBase = m.M
	v2547 = m.ExcPending
	if v2547 != 0 {
		goto L9
	} else {
		goto L775
	}
L775:
	;
	goto L700
L776:
	;
	F_SPI_freetuptable(m, v2554)
	mBase = m.M
	v2556 = m.ExcPending
	if v2556 != 0 {
		goto L9
	} else {
		goto L779
	}
L777:
	;
	goto L778
L778:
	;
	v2557 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v2557
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v2560 == v2557 {
		v5234 = v2557
		goto L21
	} else {
		goto L780
	}
L779:
	;
	goto L778
L780:
	;
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v2560)+20))
	F_MemoryContextReset(m, v2563)
	mBase = m.M
	v2565 = m.ExcPending
	if v2565 != 0 {
		goto L9
	} else {
		goto L781
	}
L781:
	;
	v5234 = v2557
	goto L21
L782:
	;
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v2574 = F_AllocSetContextCreateInternal(m, v2569, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L9
	} else {
		goto L785
	}
L783:
	;
	v2577 = v2566
	goto L784
L784:
	;
	v2578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v2578 == int32(0) {
		goto L61
	} else {
		goto L786
	}
L785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2574
	v2577 = v2574
	goto L784
L786:
	;
	v2581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v2581 != 0 {
		goto L787
	} else {
		goto L788
	}
L787:
	;
	v2585 = v2581
	goto L789
L788:
	;
	F_exec_init_tuple_store(m, l0)
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L9
	} else {
		goto L790
	}
L789:
	;
	v2586 = *(*int64)(unsafe.Add(mBase, uint32(v2585)+40))
	v2587 = int32(_a_F_exec_stmts_1)
	v2588 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2577
	v2592 = F_CreateDestReceiver(m, int32(6))
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L9
	} else {
		goto L791
	}
L790:
	;
	v2584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2585 = v2584
	goto L789
L791:
	;
	v2594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v2596 = int32(0)
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v2592)+36)) = int32(_a_F_exec_stmts_26)
	*(*int32)(unsafe.Add(mBase, uint32(v2592)+32)) = v2597
	*(*uint8)(unsafe.Add(mBase, uint32(v2592)+28)) = uint8(v2596)
	*(*int32)(unsafe.Add(mBase, uint32(v2592)+24)) = v2595
	*(*int32)(unsafe.Add(mBase, uint32(v2592)+20)) = v2594
	goto L792
L792:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2588
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v2606 != 0 {
		goto L794
	} else {
		goto L795
	}
L793:
	;
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(v2592)+12))
	m.T0[v2724].(func(*base.Module, int32))(m, v2592)
	mBase = m.M
	v2726 = m.ExcPending
	if v2726 != 0 {
		goto L9
	} else {
		goto L827
	}
L794:
	;
	v2607 = *(*int32)(unsafe.Add(mBase, uint32(v2606)+24))
	if v2607 == int32(0) {
		goto L797
	} else {
		goto L798
	}
L795:
	;
	goto L796
L796:
	;
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v2665 = F_exec_eval_expr(m, l0, v2658, v32+int32(615), v32+int32(592), v32+int32(608))
	mBase = m.M
	v2666 = m.ExcPending
	if v2666 != 0 {
		goto L9
	} else {
		goto L811
	}
L797:
	;
	F_exec_prepare_plan(m, l0, v2606, int32(2048))
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L9
	} else {
		goto L800
	}
L798:
	;
	goto L799
L799:
	;
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v2606)+28))
	if v2613 == int32(0) {
		goto L802
	} else {
		goto L803
	}
L800:
	;
	goto L799
L801:
	;
	v2620 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v2620
	*(*int64)(unsafe.Add(mBase, uint32(v32)+632)) = v2620
	*(*int64)(unsafe.Add(mBase, uint32(v32)+624)) = v2620
	*(*int32)(unsafe.Add(mBase, uint32(v32)+616)) = v2619
	v2627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+632)) = v2592
	v2629 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+622)) = uint8(v2629)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+620)) = uint8(v2627)
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2606)+24))
	v2635 = F_SPI_execute_plan_extended(m, v2632, v32+int32(616))
	mBase = m.M
	v2636 = m.ExcPending
	if v2636 != 0 {
		goto L9
	} else {
		goto L805
	}
L802:
	;
	v2619 = int32(0)
	goto L801
L803:
	;
	goto L804
L804:
	;
	v2617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2617)+20)) = v2606
	v2619 = v2617
	goto L801
L805:
	;
	if int32(0) <= v2635 {
		goto L793
	} else {
		goto L806
	}
L806:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L9
	} else {
		goto L807
	}
L807:
	;
	v2643 = *(*int32)(unsafe.Add(mBase, uint32(v2606)))
	v2644 = F_SPI_result_code_string(m, v2635)
	mBase = m.M
	v2645 = m.ExcPending
	if v2645 != 0 {
		goto L9
	} else {
		goto L808
	}
L808:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+212)) = v2644
	*(*int32)(unsafe.Add(mBase, uint32(v32)+208)) = v2643
	F_errmsg_internal(m, int32(_a_F_exec_stmts_27), v32+int32(208))
	mBase = m.M
	v2652 = m.ExcPending
	if v2652 != 0 {
		goto L9
	} else {
		goto L809
	}
L809:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3641), int32(_a_F_exec_stmts_28))
	mBase = m.M
	v2657 = m.ExcPending
	if v2657 != 0 {
		goto L9
	} else {
		goto L810
	}
L810:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L811:
	;
	v2667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+615)))
	if v2667 == int32(1) {
		goto L60
	} else {
		goto L812
	}
L812:
	;
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v2671 = int32(_a_F_exec_stmts_1)
	v2672 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(v2674)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2675
	F_getTypeOutputInfo(m, v2670, v32+int32(616), v32+int32(604))
	mBase = m.M
	v2682 = m.ExcPending
	if v2682 != 0 {
		goto L9
	} else {
		goto L813
	}
L813:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v2684 = F_OidOutputFunctionCall(m, v2683, v2665)
	mBase = m.M
	v2685 = m.ExcPending
	if v2685 != 0 {
		goto L9
	} else {
		goto L814
	}
L814:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2672
	v2688 = F_MemoryContextStrdup(m, v2577, v2684)
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L9
	} else {
		goto L815
	}
L815:
	;
	v2690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v2690 != 0 {
		goto L816
	} else {
		goto L817
	}
L816:
	;
	F_SPI_freetuptable(m, v2690)
	mBase = m.M
	v2692 = m.ExcPending
	if v2692 != 0 {
		goto L9
	} else {
		goto L819
	}
L817:
	;
	goto L818
L818:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v2695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v2695 != 0 {
		goto L820
	} else {
		goto L821
	}
L819:
	;
	goto L818
L820:
	;
	v2696 = *(*int32)(unsafe.Add(mBase, uint32(v2695)+20))
	F_MemoryContextReset(m, v2696)
	mBase = m.M
	v2698 = m.ExcPending
	if v2698 != 0 {
		goto L9
	} else {
		goto L823
	}
L821:
	;
	goto L822
L822:
	;
	v2699 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+632)) = v2699
	*(*int64)(unsafe.Add(mBase, uint32(v32)+624)) = v2699
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v2699
	v2705 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v2706 = F_exec_eval_using_params(m, l0, v2705)
	mBase = m.M
	v2707 = m.ExcPending
	if v2707 != 0 {
		goto L9
	} else {
		goto L824
	}
L823:
	;
	goto L822
L824:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+616)) = v2706
	v2709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+632)) = v2592
	v2711 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+622)) = uint8(v2711)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+620)) = uint8(v2709)
	v2716 = F_SPI_execute_extended(m, v2688, v32+int32(616))
	mBase = m.M
	v2717 = m.ExcPending
	if v2717 != 0 {
		goto L9
	} else {
		goto L825
	}
L825:
	;
	if v2716 < int32(0) {
		goto L59
	} else {
		goto L826
	}
L826:
	;
	goto L793
L827:
	;
	v2727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v2727 != 0 {
		goto L828
	} else {
		goto L829
	}
L828:
	;
	F_SPI_freetuptable(m, v2727)
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L9
	} else {
		goto L831
	}
L829:
	;
	goto L830
L830:
	;
	v2730 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v2730
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v2733 != 0 {
		goto L832
	} else {
		goto L833
	}
L831:
	;
	goto L830
L832:
	;
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v2733)+20))
	F_MemoryContextReset(m, v2734)
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L9
	} else {
		goto L835
	}
L833:
	;
	goto L834
L834:
	;
	F_MemoryContextReset(m, v2577)
	mBase = m.M
	v2738 = m.ExcPending
	if v2738 != 0 {
		goto L9
	} else {
		goto L836
	}
L835:
	;
	goto L834
L836:
	;
	v2739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2740 = *(*int64)(unsafe.Add(mBase, uint32(v2739)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v2740 - v2586
	v2743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v2744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v2743+v2744<<(uint(int32(2))%32))))
	v2751 = int32(0)
	F_assign_simple_var(m, l0, v2748, base.I64_extend_i32_u(base.B2i32(v2586 != v2740)), v2751, v2751)
	mBase = m.M
	v2754 = m.ExcPending
	if v2754 != 0 {
		goto L9
	} else {
		goto L837
	}
L837:
	;
	v5234 = v2730
	goto L21
L838:
	;
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2775 == int32(0) {
		goto L847
	} else {
		goto L848
	}
L839:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v2756 != 0 {
		goto L838
	} else {
		goto L840
	}
L840:
	;
	v2757 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v2757 != 0 {
		goto L838
	} else {
		goto L841
	}
L841:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v2758 != 0 {
		goto L58
	} else {
		goto L842
	}
L842:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v2762 = m.ExcPending
	if v2762 != 0 {
		goto L9
	} else {
		goto L843
	}
L843:
	;
	F_errcode(m, int32(33557120))
	mBase = m.M
	v2765 = m.ExcPending
	if v2765 != 0 {
		goto L9
	} else {
		goto L844
	}
L844:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_29), int32(0))
	mBase = m.M
	v2769 = m.ExcPending
	if v2769 != 0 {
		goto L9
	} else {
		goto L845
	}
L845:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3782), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v2774 = m.ExcPending
	if v2774 != 0 {
		goto L9
	} else {
		goto L846
	}
L846:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L847:
	;
	v2778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v2783 = F_AllocSetContextCreateInternal(m, v2778, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L9
	} else {
		goto L850
	}
L848:
	;
	v2787 = v2755
	v2788 = v2775
	goto L849
L849:
	;
	v2789 = int32(0)
	if v2787 == v2789 {
		goto L852
	} else {
		goto L853
	}
L850:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2783
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v2787 = v2786
	v2788 = v2783
	goto L849
L851:
	;
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v2802 != 0 {
		goto L857
	} else {
		goto L858
	}
L852:
	;
	v2792 = int32(0)
	v2800 = v2792
	v2801 = v2792
	goto L851
L853:
	;
	goto L854
L854:
	;
	v2795 = F_plpgsql_recognize_err_condition(m, v2787, int32(1))
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L9
	} else {
		goto L855
	}
L855:
	;
	v2797 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v2798 = F_MemoryContextStrdup(m, v2788, v2797)
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L9
	} else {
		goto L856
	}
L856:
	;
	v2800 = v2795
	v2801 = v2798
	goto L851
L857:
	;
	v2803 = int32(_a_F_exec_stmts_1)
	v2804 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2788
	F_initStringInfo(m, v32+int32(616))
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L9
	} else {
		goto L860
	}
L858:
	;
	v2938 = v2789
	goto L859
L859:
	;
	v2963 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v2963 == int32(0) {
		goto L897
	} else {
		goto L898
	}
L860:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2804
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	if v2813 != 0 {
		goto L861
	} else {
		goto L862
	}
L861:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, uint32(v2813)+12))
	v2816 = v2814
	goto L863
L862:
	;
	v2816 = int32(0)
	goto L863
L863:
	;
	v2817 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v2820 = v2817
	v2825 = v2816
	goto L864
L864:
	;
	v2847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2820))))
	if v2847 != int32(37) {
		goto L868
	} else {
		goto L869
	}
L865:
	;
	if v2825 != 0 {
		goto L56
	} else {
		goto L895
	}
L866:
	;
	goto L865
L867:
	;
	v2820 = v2926 + int32(1)
	v2825 = v2928
	goto L864
L868:
	;
	if v2847 == int32(0) {
		goto L866
	} else {
		goto L871
	}
L869:
	;
	goto L870
L870:
	;
	v2857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2820)+1)))
	if v2857 == int32(37) {
		goto L873
	} else {
		goto L874
	}
L871:
	;
	F_appendStringInfoChar(m, v32+int32(616), base.I32_extend8_s(v2847))
	mBase = m.M
	v2856 = m.ExcPending
	if v2856 != 0 {
		goto L9
	} else {
		goto L872
	}
L872:
	;
	v2926 = v2820
	v2928 = v2825
	goto L867
L873:
	;
	F_appendStringInfoChar(m, v32+int32(616), int32(37))
	mBase = m.M
	v2866 = m.ExcPending
	if v2866 != 0 {
		goto L9
	} else {
		goto L876
	}
L874:
	;
	goto L875
L875:
	;
	if v2825 == int32(0) {
		goto L57
	} else {
		goto L877
	}
L876:
	;
	v2926 = v2820 + int32(1)
	v2928 = v2825
	goto L867
L877:
	;
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2825)))
	v2876 = F_exec_eval_expr(m, l0, v2869, v32+int32(603), v32+int32(608), v32+int32(604))
	mBase = m.M
	v2877 = m.ExcPending
	if v2877 != 0 {
		goto L9
	} else {
		goto L878
	}
L878:
	;
	v2878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+603)))
	if v2878 != 0 {
		goto L880
	} else {
		goto L881
	}
L879:
	;
	F_appendStringInfoString(m, v32+int32(616), v2898)
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L9
	} else {
		goto L885
	}
L880:
	;
	v2898 = int32(_a_F_exec_stmts_31)
	goto L879
L881:
	;
	goto L882
L882:
	;
	v2880 = *(*int32)(unsafe.Add(mBase, uint32(v32)+608))
	v2881 = int32(_a_F_exec_stmts_1)
	v2882 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v2884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v2885 = *(*int32)(unsafe.Add(mBase, uint32(v2884)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2885
	F_getTypeOutputInfo(m, v2880, v32+int32(592), v32+int32(615))
	mBase = m.M
	v2892 = m.ExcPending
	if v2892 != 0 {
		goto L9
	} else {
		goto L883
	}
L883:
	;
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v2894 = F_OidOutputFunctionCall(m, v2893, v2876)
	mBase = m.M
	v2895 = m.ExcPending
	if v2895 != 0 {
		goto L9
	} else {
		goto L884
	}
L884:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v2882
	v2898 = v2894
	goto L879
L885:
	;
	v2905 = v2825 + int32(4)
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v2906)+12))
	v2908 = *(*int32)(unsafe.Add(mBase, uint32(v2906)+4))
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v2913 != 0 {
		goto L886
	} else {
		goto L887
	}
L886:
	;
	F_SPI_freetuptable(m, v2913)
	mBase = m.M
	v2915 = m.ExcPending
	if v2915 != 0 {
		goto L9
	} else {
		goto L889
	}
L887:
	;
	goto L888
L888:
	;
	if base.Ui32(v2905) < base.Ui32(v2907+v2908<<(uint(int32(2))%32)) {
		goto L890
	} else {
		goto L891
	}
L889:
	;
	goto L888
L890:
	;
	v2917 = v2905
	goto L892
L891:
	;
	v2917 = int32(0)
	goto L892
L892:
	;
	v2918 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v2918
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v2920 == v2918 {
		v2926 = v2820
		v2928 = v2917
		goto L867
	} else {
		goto L893
	}
L893:
	;
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(v2920)+20))
	F_MemoryContextReset(m, v2923)
	mBase = m.M
	v2925 = m.ExcPending
	if v2925 != 0 {
		goto L9
	} else {
		goto L894
	}
L894:
	;
	v2926 = v2820
	v2928 = v2917
	goto L867
L895:
	;
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v2938 = v2933
	goto L859
L896:
	;
	if v3100 == int32(0) {
		goto L945
	} else {
		goto L946
	}
L897:
	;
	v2966 = int32(0)
	v3095 = v2938
	v3096 = v2801
	v3100 = v2800
	v3101 = v2966
	v3102 = v2966
	v3103 = v2966
	v3104 = v2966
	v3107 = v2966
	v3108 = v2966
	v3109 = v2966
	goto L896
L898:
	;
	goto L899
L899:
	;
	v2973 = int32(0)
	v2981 = *(*int32)(unsafe.Add(mBase, uint32(v2963)+4))
	if v2981 <= v2973 {
		v3095 = v2938
		v3096 = v2801
		v3100 = v2800
		v3101 = v2973
		v3102 = v2973
		v3103 = v2973
		v3104 = v2973
		v3107 = v2973
		v3108 = v2973
		v3109 = v2973
		goto L896
	} else {
		goto L900
	}
L900:
	;
	v2986 = v2973
	v2988 = v2938
	v2989 = v2801
	v2993 = v2800
	v2994 = v2973
	v2995 = v2973
	v2996 = v2973
	v2997 = v2973
	v3000 = v2973
	v3001 = v2973
	v3002 = v2973
	goto L901
L901:
	;
	v3013 = *(*int32)(unsafe.Add(mBase, uint32(v2963)+12))
	v3017 = *(*int32)(unsafe.Add(mBase, uint32(v3013+v2986<<(uint(int32(2))%32))))
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(v3017)+4))
	v3025 = F_exec_eval_expr(m, l0, v3018, v32+int32(615), v32+int32(592), v32+int32(608))
	mBase = m.M
	v3026 = m.ExcPending
	if v3026 != 0 {
		goto L9
	} else {
		goto L903
	}
L902:
	;
	v3095 = v3068
	v3096 = v3069
	v3100 = v3070
	v3101 = v3071
	v3102 = v3072
	v3103 = v3073
	v3104 = v3074
	v3107 = v3075
	v3108 = v3076
	v3109 = v3077
	goto L896
L903:
	;
	v3027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+615)))
	if v3027 != 0 {
		goto L55
	} else {
		goto L904
	}
L904:
	;
	v3028 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v3029 = int32(_a_F_exec_stmts_1)
	v3030 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v3032)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3033
	F_getTypeOutputInfo(m, v3028, v32+int32(616), v32+int32(604))
	mBase = m.M
	v3040 = m.ExcPending
	if v3040 != 0 {
		goto L9
	} else {
		goto L905
	}
L905:
	;
	v3041 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v3042 = F_OidOutputFunctionCall(m, v3041, v3025)
	mBase = m.M
	v3043 = m.ExcPending
	if v3043 != 0 {
		goto L9
	} else {
		goto L906
	}
L906:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3030
	v3046 = *(*int32)(unsafe.Add(mBase, uint32(v3017)))
	switch v3046 {
	case 0:
		goto L916
	case 1:
		goto L915
	case 2:
		goto L914
	case 3:
		goto L913
	case 4:
		goto L912
	case 5:
		goto L911
	case 6:
		goto L910
	case 7:
		goto L909
	case 8:
		goto L908
	default:
		goto L45
	}
L907:
	;
	v3078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3078 != 0 {
		goto L936
	} else {
		goto L937
	}
L908:
	;
	if v2995 != 0 {
		goto L46
	} else {
		goto L934
	}
L909:
	;
	if v2996 != 0 {
		goto L47
	} else {
		goto L932
	}
L910:
	;
	if v2994 != 0 {
		goto L48
	} else {
		goto L930
	}
L911:
	;
	if v2997 != 0 {
		goto L49
	} else {
		goto L928
	}
L912:
	;
	if v3000 != 0 {
		goto L50
	} else {
		goto L926
	}
L913:
	;
	if v3001 != 0 {
		goto L51
	} else {
		goto L924
	}
L914:
	;
	if v3002 != 0 {
		goto L52
	} else {
		goto L922
	}
L915:
	;
	if v2988 != 0 {
		goto L53
	} else {
		goto L920
	}
L916:
	;
	if v2993 != 0 {
		goto L54
	} else {
		goto L917
	}
L917:
	;
	v3048 = F_plpgsql_recognize_err_condition(m, v3042, int32(1))
	mBase = m.M
	v3049 = m.ExcPending
	if v3049 != 0 {
		goto L9
	} else {
		goto L918
	}
L918:
	;
	v3050 = F_MemoryContextStrdup(m, v2788, v3042)
	mBase = m.M
	v3051 = m.ExcPending
	if v3051 != 0 {
		goto L9
	} else {
		goto L919
	}
L919:
	;
	v3068 = v2988
	v3069 = v3050
	v3070 = v3048
	v3071 = v2994
	v3072 = v2995
	v3073 = v2996
	v3074 = v2997
	v3075 = v3000
	v3076 = v3001
	v3077 = v3002
	goto L907
L920:
	;
	v3052 = F_MemoryContextStrdup(m, v2788, v3042)
	mBase = m.M
	v3053 = m.ExcPending
	if v3053 != 0 {
		goto L9
	} else {
		goto L921
	}
L921:
	;
	v3068 = v3052
	v3069 = v2989
	v3070 = v2993
	v3071 = v2994
	v3072 = v2995
	v3073 = v2996
	v3074 = v2997
	v3075 = v3000
	v3076 = v3001
	v3077 = v3002
	goto L907
L922:
	;
	v3054 = F_MemoryContextStrdup(m, v2788, v3042)
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L9
	} else {
		goto L923
	}
L923:
	;
	v3068 = v2988
	v3069 = v2989
	v3070 = v2993
	v3071 = v2994
	v3072 = v2995
	v3073 = v2996
	v3074 = v2997
	v3075 = v3000
	v3076 = v3001
	v3077 = v3054
	goto L907
L924:
	;
	v3056 = F_MemoryContextStrdup(m, v2788, v3042)
	mBase = m.M
	v3057 = m.ExcPending
	if v3057 != 0 {
		goto L9
	} else {
		goto L925
	}
L925:
	;
	v3068 = v2988
	v3069 = v2989
	v3070 = v2993
	v3071 = v2994
	v3072 = v2995
	v3073 = v2996
	v3074 = v2997
	v3075 = v3000
	v3076 = v3056
	v3077 = v3002
	goto L907
L926:
	;
	v3058 = F_MemoryContextStrdup(m, v2788, v3042)
	mBase = m.M
	v3059 = m.ExcPending
	if v3059 != 0 {
		goto L9
	} else {
		goto L927
	}
L927:
	;
	v3068 = v2988
	v3069 = v2989
	v3070 = v2993
	v3071 = v2994
	v3072 = v2995
	v3073 = v2996
	v3074 = v2997
	v3075 = v3058
	v3076 = v3001
	v3077 = v3002
	goto L907
L928:
	;
	v3060 = F_MemoryContextStrdup(m, v2788, v3042)
	mBase = m.M
	v3061 = m.ExcPending
	if v3061 != 0 {
		goto L9
	} else {
		goto L929
	}
L929:
	;
	v3068 = v2988
	v3069 = v2989
	v3070 = v2993
	v3071 = v2994
	v3072 = v2995
	v3073 = v2996
	v3074 = v3060
	v3075 = v3000
	v3076 = v3001
	v3077 = v3002
	goto L907
L930:
	;
	v3062 = F_MemoryContextStrdup(m, v2788, v3042)
	mBase = m.M
	v3063 = m.ExcPending
	if v3063 != 0 {
		goto L9
	} else {
		goto L931
	}
L931:
	;
	v3068 = v2988
	v3069 = v2989
	v3070 = v2993
	v3071 = v3062
	v3072 = v2995
	v3073 = v2996
	v3074 = v2997
	v3075 = v3000
	v3076 = v3001
	v3077 = v3002
	goto L907
L932:
	;
	v3064 = F_MemoryContextStrdup(m, v2788, v3042)
	mBase = m.M
	v3065 = m.ExcPending
	if v3065 != 0 {
		goto L9
	} else {
		goto L933
	}
L933:
	;
	v3068 = v2988
	v3069 = v2989
	v3070 = v2993
	v3071 = v2994
	v3072 = v2995
	v3073 = v3064
	v3074 = v2997
	v3075 = v3000
	v3076 = v3001
	v3077 = v3002
	goto L907
L934:
	;
	v3066 = F_MemoryContextStrdup(m, v2788, v3042)
	mBase = m.M
	v3067 = m.ExcPending
	if v3067 != 0 {
		goto L9
	} else {
		goto L935
	}
L935:
	;
	v3068 = v2988
	v3069 = v2989
	v3070 = v2993
	v3071 = v2994
	v3072 = v3066
	v3073 = v2996
	v3074 = v2997
	v3075 = v3000
	v3076 = v3001
	v3077 = v3002
	goto L907
L936:
	;
	F_SPI_freetuptable(m, v3078)
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L9
	} else {
		goto L939
	}
L937:
	;
	goto L938
L938:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v3083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v3083 != 0 {
		goto L940
	} else {
		goto L941
	}
L939:
	;
	goto L938
L940:
	;
	v3084 = *(*int32)(unsafe.Add(mBase, uint32(v3083)+20))
	F_MemoryContextReset(m, v3084)
	mBase = m.M
	v3086 = m.ExcPending
	if v3086 != 0 {
		goto L9
	} else {
		goto L943
	}
L941:
	;
	goto L942
L942:
	;
	v3088 = v2986 + int32(1)
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v2963)+4))
	if v3088 < v3089 {
		v2986 = v3088
		v2988 = v3068
		v2989 = v3069
		v2993 = v3070
		v2994 = v3071
		v2995 = v3072
		v2996 = v3073
		v2997 = v3074
		v3000 = v3075
		v3001 = v3076
		v3002 = v3077
		goto L901
	} else {
		goto L944
	}
L943:
	;
	goto L942
L944:
	;
	goto L902
L945:
	;
	v3124 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if int32(20) < v3124 {
		goto L948
	} else {
		goto L949
	}
L946:
	;
	v3128 = v3100
	goto L947
L947:
	;
	if v3095 != 0 {
		v3173 = v3095
		goto L951
	} else {
		goto L952
	}
L948:
	;
	v3127 = int32(16777248)
	goto L950
L949:
	;
	v3127 = int32(0)
	goto L950
L950:
	;
	v3128 = v3127
	goto L947
L951:
	;
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3176 = F_errstart(m, v3174, int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3177 = m.ExcPending
	if v3177 != 0 {
		goto L9
	} else {
		goto L956
	}
L952:
	;
	if v3096 != 0 {
		v3173 = v3096
		goto L951
	} else {
		goto L953
	}
L953:
	;
	v3129 = int32(_a_F_exec_stmts_8)
	v3130 = int32(63)
	v3132 = int32(48)
	v3133 = v3128&v3130 + v3132
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[10])) = uint8(v3133)
	v3141 = int32(base.Ui32(v3128)>>(uint(int32(24))%32))&v3130 + v3132
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[11])) = uint8(v3141)
	v3149 = int32(base.Ui32(v3128)>>(uint(int32(18))%32))&v3130 + v3132
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[12])) = uint8(v3149)
	v3157 = int32(base.Ui32(v3128)>>(uint(int32(12))%32))&v3130 + v3132
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[13])) = uint8(v3157)
	v3165 = int32(base.Ui32(v3128)>>(uint(int32(6))%32))&v3130 + v3132
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[14])) = uint8(v3165)
	v3168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_exec_stmts[15])) = uint8(v3168)
	goto L954
L954:
	;
	v3171 = F_MemoryContextStrdup(m, v2788, v3129)
	mBase = m.M
	v3172 = m.ExcPending
	if v3172 != 0 {
		goto L9
	} else {
		goto L955
	}
L955:
	;
	v3173 = v3171
	goto L951
L956:
	;
	if v3176 != 0 {
		goto L957
	} else {
		goto L958
	}
L957:
	;
	if v3128 != 0 {
		goto L960
	} else {
		goto L961
	}
L958:
	;
	goto L959
L959:
	;
	F_MemoryContextReset(m, v2788)
	mBase = m.M
	v3219 = m.ExcPending
	if v3219 != 0 {
		goto L9
	} else {
		goto L994
	}
L960:
	;
	F_errcode(m, v3128)
	mBase = m.M
	v3179 = m.ExcPending
	if v3179 != 0 {
		goto L9
	} else {
		goto L963
	}
L961:
	;
	goto L962
L962:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+256)) = v3173
	F_errmsg_internal(m, int32(_a_F_exec_stmts_32), v32+int32(256))
	mBase = m.M
	v3185 = m.ExcPending
	if v3185 != 0 {
		goto L9
	} else {
		goto L964
	}
L963:
	;
	goto L962
L964:
	;
	if v3109 != 0 {
		goto L965
	} else {
		goto L966
	}
L965:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+240)) = v3109
	F_errdetail_internal(m, int32(_a_F_exec_stmts_32), v32+int32(240))
	mBase = m.M
	v3191 = m.ExcPending
	if v3191 != 0 {
		goto L9
	} else {
		goto L968
	}
L966:
	;
	goto L967
L967:
	;
	if v3108 != 0 {
		goto L969
	} else {
		goto L970
	}
L968:
	;
	goto L967
L969:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+224)) = v3108
	F_errhint(m, int32(_a_F_exec_stmts_32), v32+int32(224))
	mBase = m.M
	v3197 = m.ExcPending
	if v3197 != 0 {
		goto L9
	} else {
		goto L972
	}
L970:
	;
	goto L971
L971:
	;
	if v3107 != 0 {
		goto L973
	} else {
		goto L974
	}
L972:
	;
	goto L971
L973:
	;
	F_err_generic_string(m, int32(99), v3107)
	mBase = m.M
	v3200 = m.ExcPending
	if v3200 != 0 {
		goto L9
	} else {
		goto L976
	}
L974:
	;
	goto L975
L975:
	;
	if v3104 != 0 {
		goto L977
	} else {
		goto L978
	}
L976:
	;
	goto L975
L977:
	;
	F_err_generic_string(m, int32(110), v3104)
	mBase = m.M
	v3203 = m.ExcPending
	if v3203 != 0 {
		goto L9
	} else {
		goto L980
	}
L978:
	;
	goto L979
L979:
	;
	if v3101 != 0 {
		goto L981
	} else {
		goto L982
	}
L980:
	;
	goto L979
L981:
	;
	F_err_generic_string(m, int32(100), v3101)
	mBase = m.M
	v3206 = m.ExcPending
	if v3206 != 0 {
		goto L9
	} else {
		goto L984
	}
L982:
	;
	goto L983
L983:
	;
	if v3103 != 0 {
		goto L985
	} else {
		goto L986
	}
L984:
	;
	goto L983
L985:
	;
	F_err_generic_string(m, int32(116), v3103)
	mBase = m.M
	v3209 = m.ExcPending
	if v3209 != 0 {
		goto L9
	} else {
		goto L988
	}
L986:
	;
	goto L987
L987:
	;
	if v3102 != 0 {
		goto L989
	} else {
		goto L990
	}
L988:
	;
	goto L987
L989:
	;
	F_err_generic_string(m, int32(115), v3102)
	mBase = m.M
	v3212 = m.ExcPending
	if v3212 != 0 {
		goto L9
	} else {
		goto L992
	}
L990:
	;
	goto L991
L991:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3956), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v3217 = m.ExcPending
	if v3217 != 0 {
		goto L9
	} else {
		goto L993
	}
L992:
	;
	goto L991
L993:
	;
	goto L959
L994:
	;
	v5234 = int32(0)
	goto L21
L995:
	;
	v5234 = int32(0)
	goto L21
L996:
	;
	v3225 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3227 = v32 + int32(608)
	v3232 = F_exec_eval_expr(m, l0, v3225, v3227, v32+int32(616), v32+int32(592))
	mBase = m.M
	v3233 = m.ExcPending
	if v3233 != 0 {
		goto L9
	} else {
		goto L997
	}
L997:
	;
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v3235 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v3238 = F_exec_cast_value(m, l0, v3232, v3227, v3234, v3235, int32(16), int32(-1))
	mBase = m.M
	v3239 = m.ExcPending
	if v3239 != 0 {
		goto L9
	} else {
		goto L998
	}
L998:
	;
	v3240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3240 != 0 {
		goto L999
	} else {
		goto L1000
	}
L999:
	;
	F_SPI_freetuptable(m, v3240)
	mBase = m.M
	v3242 = m.ExcPending
	if v3242 != 0 {
		goto L9
	} else {
		goto L1002
	}
L1000:
	;
	goto L1001
L1001:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v3245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v3245 != 0 {
		goto L1003
	} else {
		goto L1004
	}
L1002:
	;
	goto L1001
L1003:
	;
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(v3245)+20))
	F_MemoryContextReset(m, v3246)
	mBase = m.M
	v3248 = m.ExcPending
	if v3248 != 0 {
		goto L9
	} else {
		goto L1006
	}
L1004:
	;
	goto L1005
L1005:
	;
	v3249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	if base.B2i32(v3249 == int32(0))&base.B2i32(v3238 != int64(0)) != 0 {
		goto L995
	} else {
		goto L1007
	}
L1006:
	;
	goto L1005
L1007:
	;
	v3255 = int32(0)
	v3256 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v3256 == v3255 {
		v3294 = v3255
		goto L1008
	} else {
		goto L1009
	}
L1008:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3300 = m.ExcPending
	if v3300 != 0 {
		goto L9
	} else {
		goto L1014
	}
L1009:
	;
	v3265 = F_exec_eval_expr(m, l0, v3256, v32+int32(608), v32+int32(616), v32+int32(592))
	mBase = m.M
	v3266 = m.ExcPending
	if v3266 != 0 {
		goto L9
	} else {
		goto L1010
	}
L1010:
	;
	v3267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	if v3267 != 0 {
		v3294 = v3255
		goto L1008
	} else {
		goto L1011
	}
L1011:
	;
	v3268 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v3269 = m.G0
	v3271 = v3269 - int32(16)
	m.G0 = v3271
	v3273 = int32(_a_F_exec_stmts_1)
	v3274 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v3277 = *(*int32)(unsafe.Add(mBase, uint32(v3276)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3277
	F_getTypeOutputInfo(m, v3268, v3271+int32(12), v3271+int32(11))
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L9
	} else {
		goto L1012
	}
L1012:
	;
	v3285 = *(*int32)(unsafe.Add(mBase, uint32(v3271)+12))
	v3286 = F_OidOutputFunctionCall(m, v3285, v3265)
	mBase = m.M
	v3287 = m.ExcPending
	if v3287 != 0 {
		goto L9
	} else {
		goto L1013
	}
L1013:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3274
	m.G0 = v3271 + int32(16)
	v3294 = v3286
	goto L1008
L1014:
	;
	F_errcode(m, int32(67108896))
	mBase = m.M
	v3303 = m.ExcPending
	if v3303 != 0 {
		goto L9
	} else {
		goto L1015
	}
L1015:
	;
	if v3294 != 0 {
		goto L1017
	} else {
		goto L1018
	}
L1016:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(4001), int32(_a_F_exec_stmts_33))
	mBase = m.M
	v3318 = m.ExcPending
	if v3318 != 0 {
		goto L9
	} else {
		goto L1022
	}
L1017:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+432)) = v3294
	F_errmsg_internal(m, int32(_a_F_exec_stmts_32), v32+int32(432))
	mBase = m.M
	v3309 = m.ExcPending
	if v3309 != 0 {
		goto L9
	} else {
		goto L1020
	}
L1018:
	;
	goto L1019
L1019:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_34), int32(0))
	mBase = m.M
	v3313 = m.ExcPending
	if v3313 != 0 {
		goto L9
	} else {
		goto L1021
	}
L1020:
	;
	goto L1016
L1021:
	;
	goto L1016
L1022:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1023:
	;
	v5234 = int32(0)
	goto L21
L1024:
	;
	v3329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v3334 = F_AllocSetContextCreateInternal(m, v3329, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v3335 = m.ExcPending
	if v3335 != 0 {
		goto L9
	} else {
		goto L1027
	}
L1025:
	;
	v3337 = v3326
	goto L1026
L1026:
	;
	v3338 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3345 = F_exec_eval_expr(m, l0, v3338, v32+int32(615), v32+int32(592), v32+int32(608))
	mBase = m.M
	v3346 = m.ExcPending
	if v3346 != 0 {
		goto L9
	} else {
		goto L1028
	}
L1027:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v3334
	v3337 = v3334
	goto L1026
L1028:
	;
	v3347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+615)))
	if v3347 == int32(1) {
		goto L44
	} else {
		goto L1029
	}
L1029:
	;
	v3350 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v3351 = int32(_a_F_exec_stmts_1)
	v3352 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v3354)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3355
	F_getTypeOutputInfo(m, v3350, v32+int32(616), v32+int32(604))
	mBase = m.M
	v3362 = m.ExcPending
	if v3362 != 0 {
		goto L9
	} else {
		goto L1030
	}
L1030:
	;
	v3363 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v3364 = F_OidOutputFunctionCall(m, v3363, v3345)
	mBase = m.M
	v3365 = m.ExcPending
	if v3365 != 0 {
		goto L9
	} else {
		goto L1031
	}
L1031:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3352
	v3368 = F_MemoryContextStrdup(m, v3337, v3364)
	mBase = m.M
	v3369 = m.ExcPending
	if v3369 != 0 {
		goto L9
	} else {
		goto L1032
	}
L1032:
	;
	v3370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3370 != 0 {
		goto L1033
	} else {
		goto L1034
	}
L1033:
	;
	F_SPI_freetuptable(m, v3370)
	mBase = m.M
	v3372 = m.ExcPending
	if v3372 != 0 {
		goto L9
	} else {
		goto L1036
	}
L1034:
	;
	goto L1035
L1035:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v3375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v3375 != 0 {
		goto L1037
	} else {
		goto L1038
	}
L1036:
	;
	goto L1035
L1037:
	;
	v3376 = *(*int32)(unsafe.Add(mBase, uint32(v3375)+20))
	F_MemoryContextReset(m, v3376)
	mBase = m.M
	v3378 = m.ExcPending
	if v3378 != 0 {
		goto L9
	} else {
		goto L1040
	}
L1038:
	;
	goto L1039
L1039:
	;
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v3380 = F_exec_eval_using_params(m, l0, v3379)
	mBase = m.M
	v3381 = m.ExcPending
	if v3381 != 0 {
		goto L9
	} else {
		goto L1041
	}
L1040:
	;
	goto L1039
L1041:
	;
	v3382 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+616)) = v3382
	*(*int64)(unsafe.Add(mBase, uint32(v32)+632)) = v3382
	*(*int64)(unsafe.Add(mBase, uint32(v32)+624)) = v3382
	*(*int32)(unsafe.Add(mBase, uint32(v32)+616)) = v3380
	v3389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+620)) = uint8(v3389)
	v3393 = F_SPI_execute_extended(m, v3368, v32+int32(616))
	mBase = m.M
	v3394 = m.ExcPending
	if v3394 != 0 {
		goto L9
	} else {
		goto L1047
	}
L1042:
	;
	v3468 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmts[6]))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v3468
	v3470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+16)))
	if v3470 != int32(1) {
		goto L1065
	} else {
		goto L1066
	}
L1043:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3452 = m.ExcPending
	if v3452 != 0 {
		goto L9
	} else {
		goto L1061
	}
L1044:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3436 = m.ExcPending
	if v3436 != 0 {
		goto L9
	} else {
		goto L1057
	}
L1045:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3420 = m.ExcPending
	if v3420 != 0 {
		goto L9
	} else {
		goto L1053
	}
L1046:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3400 = m.ExcPending
	if v3400 != 0 {
		goto L9
	} else {
		goto L1048
	}
L1047:
	;
	switch v3393 + int32(8) {
	case 0:
		goto L1044
	default:
		goto L1043
	case 6:
		goto L1045
	case 8, 12, 13, 15, 16, 17, 19, 20, 21, 22, 26, 27:
		goto L1042
	case 14:
		goto L1046
	}
L1048:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3403 = m.ExcPending
	if v3403 != 0 {
		goto L9
	} else {
		goto L1049
	}
L1049:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_35), int32(0))
	mBase = m.M
	v3407 = m.ExcPending
	if v3407 != 0 {
		goto L9
	} else {
		goto L1050
	}
L1050:
	;
	F_errhint(m, int32(_a_F_exec_stmts_36), int32(0))
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		goto L9
	} else {
		goto L1051
	}
L1051:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_37), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3416 = m.ExcPending
	if v3416 != 0 {
		goto L9
	} else {
		goto L1052
	}
L1052:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1053:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3423 = m.ExcPending
	if v3423 != 0 {
		goto L9
	} else {
		goto L1054
	}
L1054:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_39), int32(0))
	mBase = m.M
	v3427 = m.ExcPending
	if v3427 != 0 {
		goto L9
	} else {
		goto L1055
	}
L1055:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_40), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3432 = m.ExcPending
	if v3432 != 0 {
		goto L9
	} else {
		goto L1056
	}
L1056:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1057:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L9
	} else {
		goto L1058
	}
L1058:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_41), int32(0))
	mBase = m.M
	v3443 = m.ExcPending
	if v3443 != 0 {
		goto L9
	} else {
		goto L1059
	}
L1059:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_42), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3448 = m.ExcPending
	if v3448 != 0 {
		goto L9
	} else {
		goto L1060
	}
L1060:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1061:
	;
	v3453 = F_SPI_result_code_string(m, v3393)
	mBase = m.M
	v3454 = m.ExcPending
	if v3454 != 0 {
		goto L9
	} else {
		goto L1062
	}
L1062:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+452)) = v3453
	*(*int32)(unsafe.Add(mBase, uint32(v32)+448)) = v3368
	F_errmsg_internal(m, int32(_a_F_exec_stmts_43), v32+int32(448))
	mBase = m.M
	v3461 = m.ExcPending
	if v3461 != 0 {
		goto L9
	} else {
		goto L1063
	}
L1063:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_44), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3466 = m.ExcPending
	if v3466 != 0 {
		goto L9
	} else {
		goto L1064
	}
L1064:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1065:
	;
	v3576 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	F_SPI_freetuptable(m, v3576)
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L9
	} else {
		goto L1107
	}
L1066:
	;
	v3474 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	if v3474 == int32(0) {
		goto L43
	} else {
		goto L1067
	}
L1067:
	;
	v3477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v3478 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v3479 = *(*int32)(unsafe.Add(mBase, uint32(v3478)+4))
	v3483 = *(*int32)(unsafe.Add(mBase, uint32(v3477+v3479<<(uint(int32(2))%32))))
	if base.Ui64(v3468) <= base.Ui64(int64(1)) {
		goto L1070
	} else {
		goto L1071
	}
L1068:
	;
	v3559 = *(*int32)(unsafe.Add(mBase, uint32(v3474)))
	F_exec_move_row(m, l0, v3483, v3558, v3559)
	mBase = m.M
	v3561 = m.ExcPending
	if v3561 != 0 {
		goto L9
	} else {
		goto L1100
	}
L1069:
	;
	v3556 = *(*int32)(unsafe.Add(mBase, uint32(v3474)+4))
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(v3556)))
	v3558 = v3557
	goto L1068
L1070:
	;
	if base.I32_wrap_i64(v3468) == int32(1) {
		goto L1069
	} else {
		goto L1073
	}
L1071:
	;
	goto L1072
L1072:
	;
	v3523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+17)))
	if v3523 != int32(1) {
		goto L1069
	} else {
		goto L1087
	}
L1073:
	;
	v3490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+17)))
	if v3490 != int32(1) {
		v3558 = int32(0)
		goto L1068
	} else {
		goto L1074
	}
L1074:
	;
	v3494 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3494)+492)))
	if v3495 == int32(1) {
		goto L1075
	} else {
		goto L1076
	}
L1075:
	;
	v3498 = F_format_preparedparamsdata(m, l0, v3380)
	mBase = m.M
	v3499 = m.ExcPending
	if v3499 != 0 {
		goto L9
	} else {
		goto L1078
	}
L1076:
	;
	v3500 = int32(0)
	goto L1077
L1077:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3504 = m.ExcPending
	if v3504 != 0 {
		goto L9
	} else {
		goto L1079
	}
L1078:
	;
	v3500 = v3498
	goto L1077
L1079:
	;
	F_errcode(m, int32(33554464))
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		goto L9
	} else {
		goto L1080
	}
L1080:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_45), int32(0))
	mBase = m.M
	v3511 = m.ExcPending
	if v3511 != 0 {
		goto L9
	} else {
		goto L1081
	}
L1081:
	;
	if v3500 != 0 {
		goto L1082
	} else {
		goto L1083
	}
L1082:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+480)) = v3500
	F_errdetail_internal(m, int32(_a_F_exec_stmts_46), v32+int32(480))
	mBase = m.M
	v3517 = m.ExcPending
	if v3517 != 0 {
		goto L9
	} else {
		goto L1085
	}
L1083:
	;
	goto L1084
L1084:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_47), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3522 = m.ExcPending
	if v3522 != 0 {
		goto L9
	} else {
		goto L1086
	}
L1085:
	;
	goto L1084
L1086:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1087:
	;
	v3527 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3527)+492)))
	if v3528 == int32(1) {
		goto L1088
	} else {
		goto L1089
	}
L1088:
	;
	v3531 = F_format_preparedparamsdata(m, l0, v3380)
	mBase = m.M
	v3532 = m.ExcPending
	if v3532 != 0 {
		goto L9
	} else {
		goto L1091
	}
L1089:
	;
	v3533 = int32(0)
	goto L1090
L1090:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_stmts_0))
	mBase = m.M
	v3537 = m.ExcPending
	if v3537 != 0 {
		goto L9
	} else {
		goto L1092
	}
L1091:
	;
	v3533 = v3531
	goto L1090
L1092:
	;
	F_errcode(m, int32(50331680))
	mBase = m.M
	v3540 = m.ExcPending
	if v3540 != 0 {
		goto L9
	} else {
		goto L1093
	}
L1093:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_48), int32(0))
	mBase = m.M
	v3544 = m.ExcPending
	if v3544 != 0 {
		goto L9
	} else {
		goto L1094
	}
L1094:
	;
	if v3533 != 0 {
		goto L1095
	} else {
		goto L1096
	}
L1095:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+464)) = v3533
	F_errdetail_internal(m, int32(_a_F_exec_stmts_46), v32+int32(464))
	mBase = m.M
	v3550 = m.ExcPending
	if v3550 != 0 {
		goto L9
	} else {
		goto L1098
	}
L1096:
	;
	goto L1097
L1097:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_49), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v3555 = m.ExcPending
	if v3555 != 0 {
		goto L9
	} else {
		goto L1099
	}
L1098:
	;
	goto L1097
L1099:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1100:
	;
	v3562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3562 != 0 {
		goto L1101
	} else {
		goto L1102
	}
L1101:
	;
	F_SPI_freetuptable(m, v3562)
	mBase = m.M
	v3564 = m.ExcPending
	if v3564 != 0 {
		goto L9
	} else {
		goto L1104
	}
L1102:
	;
	goto L1103
L1103:
	;
	v3565 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v3565
	v3567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v3567 == v3565 {
		goto L1065
	} else {
		goto L1105
	}
L1104:
	;
	goto L1103
L1105:
	;
	v3570 = *(*int32)(unsafe.Add(mBase, uint32(v3567)+20))
	F_MemoryContextReset(m, v3570)
	mBase = m.M
	v3572 = m.ExcPending
	if v3572 != 0 {
		goto L9
	} else {
		goto L1106
	}
L1106:
	;
	goto L1065
L1107:
	;
	F_MemoryContextReset(m, v3337)
	mBase = m.M
	v3580 = m.ExcPending
	if v3580 != 0 {
		goto L9
	} else {
		goto L1108
	}
L1108:
	;
	v5234 = int32(0)
	goto L21
L1109:
	;
	v3589 = F_exec_for_query(m, l0, v89, v3586, int32(1))
	mBase = m.M
	v3590 = m.ExcPending
	if v3590 != 0 {
		goto L9
	} else {
		goto L1110
	}
L1110:
	;
	F_SPI_cursor_close(m, v3586)
	mBase = m.M
	v3592 = m.ExcPending
	if v3592 != 0 {
		goto L9
	} else {
		goto L1111
	}
L1111:
	;
	v5234 = v3589
	goto L21
L1112:
	;
	v3604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v3604 == int32(0) {
		goto L1115
	} else {
		goto L1116
	}
L1113:
	;
	v3628 = v3593
	v3629 = v3593
	goto L1114
L1114:
	;
	v3630 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	if v3630 != 0 {
		goto L1123
	} else {
		goto L1124
	}
L1115:
	;
	v3607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v3612 = F_AllocSetContextCreateInternal(m, v3607, int32(_a_F_exec_stmts_17), int32(0), int32(_a_F_exec_stmts_18), int32(_a_F_exec_stmts_19))
	mBase = m.M
	v3613 = m.ExcPending
	if v3613 != 0 {
		goto L9
	} else {
		goto L1118
	}
L1116:
	;
	v3615 = v3604
	goto L1117
L1117:
	;
	v3616 = int32(_a_F_exec_stmts_1)
	v3617 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3615
	v3620 = *(*int32)(unsafe.Add(mBase, uint32(v3600)+40))
	v3621 = F_text_to_cstring(m, v3620)
	mBase = m.M
	v3622 = m.ExcPending
	if v3622 != 0 {
		goto L9
	} else {
		goto L1119
	}
L1118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v3612
	v3615 = v3612
	goto L1117
L1119:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3617
	v3625 = F_GetPortalByName(m, v3621)
	mBase = m.M
	v3626 = m.ExcPending
	if v3626 != 0 {
		goto L9
	} else {
		goto L1120
	}
L1120:
	;
	if v3625 != 0 {
		goto L42
	} else {
		goto L1121
	}
L1121:
	;
	v3628 = v3621
	v3629 = v3615
	goto L1114
L1122:
	;
	v3691 = *(*int32)(unsafe.Add(mBase, uint32(v3688)+28))
	if v3691 == int32(0) {
		goto L1148
	} else {
		goto L1149
	}
L1123:
	;
	v3631 = *(*int32)(unsafe.Add(mBase, uint32(v3630)+24))
	if v3631 != 0 {
		v3688 = v3630
		goto L1122
	} else {
		goto L1126
	}
L1124:
	;
	goto L1125
L1125:
	;
	v3635 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v3635 != 0 {
		goto L1128
	} else {
		goto L1129
	}
L1126:
	;
	v3632 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	F_exec_prepare_plan(m, l0, v3630, v3632)
	mBase = m.M
	v3634 = m.ExcPending
	if v3634 != 0 {
		goto L9
	} else {
		goto L1127
	}
L1127:
	;
	v3688 = v3630
	goto L1122
L1128:
	;
	v3636 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	v3637 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v3638 = F_exec_dynquery_with_params(m, l0, v3635, v3636, v3628, v3637)
	mBase = m.M
	v3639 = m.ExcPending
	if v3639 != 0 {
		goto L9
	} else {
		goto L1131
	}
L1129:
	;
	goto L1130
L1130:
	;
	v3653 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if v3653 != 0 {
		goto L1139
	} else {
		goto L1140
	}
L1131:
	;
	if v3628 != 0 {
		goto L1132
	} else {
		goto L1133
	}
L1132:
	;
	v5234 = int32(0)
	goto L21
L1133:
	;
	goto L1134
L1134:
	;
	v3641 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	F_exec_check_assignable(m, l0, v3641)
	mBase = m.M
	v3643 = m.ExcPending
	if v3643 != 0 {
		goto L9
	} else {
		goto L1135
	}
L1135:
	;
	v3645 = *(*int32)(unsafe.Add(mBase, uint32(v3638)))
	v3646 = F_cstring_to_text(m, v3645)
	mBase = m.M
	v3647 = m.ExcPending
	if v3647 != 0 {
		goto L9
	} else {
		goto L1136
	}
L1136:
	;
	F_assign_simple_var(m, l0, v3600, base.I64_extend_i32_u(v3646), int32(0), int32(1))
	mBase = m.M
	v3652 = m.ExcPending
	if v3652 != 0 {
		goto L9
	} else {
		goto L1137
	}
L1137:
	;
	v5234 = int32(0)
	goto L21
L1138:
	;
	v3683 = *(*int32)(unsafe.Add(mBase, uint32(v3600)+28))
	v3684 = *(*int32)(unsafe.Add(mBase, uint32(v3683)+24))
	if v3684 != 0 {
		v3688 = v3683
		goto L1122
	} else {
		goto L1145
	}
L1139:
	;
	v3654 = *(*int32)(unsafe.Add(mBase, uint32(v3600)+32))
	if v3654 < int32(0) {
		goto L41
	} else {
		goto L1142
	}
L1140:
	;
	goto L1141
L1141:
	;
	v3678 = *(*int32)(unsafe.Add(mBase, uint32(v3600)+32))
	if int32(0) <= v3678 {
		goto L40
	} else {
		goto L1144
	}
L1142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+616)) = int32(16)
	v3663 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v3664 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+634)) = uint8(v3664)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+628)) = v3653
	*(*int32)(unsafe.Add(mBase, uint32(v32)+620)) = v3663
	v3668 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v3672 = *(*int32)(unsafe.Add(mBase, uint32(v3668+v3654<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+636)) = v3672
	F_exec_stmt_execsql(m, l0, v32+int32(616))
	mBase = m.M
	v3677 = m.ExcPending
	if v3677 != 0 {
		goto L9
	} else {
		goto L1143
	}
L1143:
	;
	goto L1138
L1144:
	;
	goto L1138
L1145:
	;
	v3685 = *(*int32)(unsafe.Add(mBase, uint32(v3600)+36))
	F_exec_prepare_plan(m, l0, v3683, v3685)
	mBase = m.M
	v3687 = m.ExcPending
	if v3687 != 0 {
		goto L9
	} else {
		goto L1146
	}
L1146:
	;
	v3688 = v3683
	goto L1122
L1147:
	;
	v3698 = *(*int32)(unsafe.Add(mBase, uint32(v3688)+24))
	v3699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	v3700 = F_SPI_cursor_open_internal(m, v3628, v3698, v3697, v3699)
	mBase = m.M
	v3701 = m.ExcPending
	if v3701 != 0 {
		goto L9
	} else {
		goto L1151
	}
L1148:
	;
	v3697 = int32(0)
	goto L1147
L1149:
	;
	goto L1150
L1150:
	;
	v3695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v3695)+20)) = v3688
	v3697 = v3695
	goto L1147
L1151:
	;
	if v3700 == int32(0) {
		goto L39
	} else {
		goto L1152
	}
L1152:
	;
	if v3628 == int32(0) {
		goto L1153
	} else {
		goto L1154
	}
L1153:
	;
	v3706 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	F_exec_check_assignable(m, l0, v3706)
	mBase = m.M
	v3708 = m.ExcPending
	if v3708 != 0 {
		goto L9
	} else {
		goto L1156
	}
L1154:
	;
	goto L1155
L1155:
	;
	v3717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3717 != 0 {
		goto L1159
	} else {
		goto L1160
	}
L1156:
	;
	v3709 = *(*int32)(unsafe.Add(mBase, uint32(v3700)))
	v3710 = F_cstring_to_text(m, v3709)
	mBase = m.M
	v3711 = m.ExcPending
	if v3711 != 0 {
		goto L9
	} else {
		goto L1157
	}
L1157:
	;
	F_assign_simple_var(m, l0, v3600, base.I64_extend_i32_u(v3710), int32(0), int32(1))
	mBase = m.M
	v3716 = m.ExcPending
	if v3716 != 0 {
		goto L9
	} else {
		goto L1158
	}
L1158:
	;
	goto L1155
L1159:
	;
	F_SPI_freetuptable(m, v3717)
	mBase = m.M
	v3719 = m.ExcPending
	if v3719 != 0 {
		goto L9
	} else {
		goto L1162
	}
L1160:
	;
	goto L1161
L1161:
	;
	v3720 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v3720
	v3723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v3723 != 0 {
		goto L1163
	} else {
		goto L1164
	}
L1162:
	;
	goto L1161
L1163:
	;
	v3724 = *(*int32)(unsafe.Add(mBase, uint32(v3723)+20))
	F_MemoryContextReset(m, v3724)
	mBase = m.M
	v3726 = m.ExcPending
	if v3726 != 0 {
		goto L9
	} else {
		goto L1166
	}
L1164:
	;
	goto L1165
L1165:
	;
	if v3629 == int32(0) {
		v5234 = v3720
		goto L21
	} else {
		goto L1167
	}
L1166:
	;
	goto L1165
L1167:
	;
	F_MemoryContextReset(m, v3629)
	mBase = m.M
	v3730 = m.ExcPending
	if v3730 != 0 {
		goto L9
	} else {
		goto L1168
	}
L1168:
	;
	v5234 = v3720
	goto L21
L1169:
	;
	v3740 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v3741 = int32(_a_F_exec_stmts_1)
	v3742 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v3744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v3745 = *(*int32)(unsafe.Add(mBase, uint32(v3744)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3745
	v3747 = *(*int32)(unsafe.Add(mBase, uint32(v3736)+40))
	v3748 = F_text_to_cstring(m, v3747)
	mBase = m.M
	v3749 = m.ExcPending
	if v3749 != 0 {
		goto L9
	} else {
		goto L1170
	}
L1170:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3742
	v3752 = F_GetPortalByName(m, v3748)
	mBase = m.M
	v3753 = m.ExcPending
	if v3753 != 0 {
		goto L9
	} else {
		goto L1171
	}
L1171:
	;
	if v3752 == int32(0) {
		goto L37
	} else {
		goto L1172
	}
L1172:
	;
	v3756 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v3756 != 0 {
		goto L1173
	} else {
		goto L1174
	}
L1173:
	;
	v3758 = v32 + int32(608)
	v3763 = F_exec_eval_expr(m, l0, v3756, v3758, v32+int32(616), v32+int32(592))
	mBase = m.M
	v3764 = m.ExcPending
	if v3764 != 0 {
		goto L9
	} else {
		goto L1176
	}
L1174:
	;
	v3785 = v3740
	goto L1175
L1175:
	;
	v3786 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v3787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+32)))
	if v3787 == int32(0) {
		goto L1188
	} else {
		goto L1189
	}
L1176:
	;
	v3765 = *(*int32)(unsafe.Add(mBase, uint32(v32)+616))
	v3766 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v3769 = F_exec_cast_value(m, l0, v3763, v3758, v3765, v3766, int32(23), int32(-1))
	mBase = m.M
	v3770 = m.ExcPending
	if v3770 != 0 {
		goto L9
	} else {
		goto L1177
	}
L1177:
	;
	v3771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+608)))
	if v3771 == int32(1) {
		goto L36
	} else {
		goto L1178
	}
L1178:
	;
	v3774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3774 != 0 {
		goto L1179
	} else {
		goto L1180
	}
L1179:
	;
	F_SPI_freetuptable(m, v3774)
	mBase = m.M
	v3776 = m.ExcPending
	if v3776 != 0 {
		goto L9
	} else {
		goto L1182
	}
L1180:
	;
	goto L1181
L1181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v3779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v3779 != 0 {
		goto L1183
	} else {
		goto L1184
	}
L1182:
	;
	goto L1181
L1183:
	;
	v3780 = *(*int32)(unsafe.Add(mBase, uint32(v3779)+20))
	F_MemoryContextReset(m, v3780)
	mBase = m.M
	v3782 = m.ExcPending
	if v3782 != 0 {
		goto L9
	} else {
		goto L1186
	}
L1184:
	;
	goto L1185
L1185:
	;
	v3785 = base.I32_wrap_i64(v3769)
	goto L1175
L1186:
	;
	goto L1185
L1187:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v3834
	v3836 = int32(0)
	v3837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v3838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v3842 = *(*int32)(unsafe.Add(mBase, uint32(v3837+v3838<<(uint(int32(2))%32))))
	F_assign_simple_var(m, l0, v3842, base.I64_extend_i32_u(base.B2i32(v3834 != int64(0))), v3836, v3836)
	mBase = m.M
	v3849 = m.ExcPending
	if v3849 != 0 {
		goto L9
	} else {
		goto L1207
	}
L1188:
	;
	v3791 = F_CreateDestReceiver(m, int32(5))
	mBase = m.M
	v3792 = m.ExcPending
	if v3792 != 0 {
		goto L9
	} else {
		goto L1191
	}
L1189:
	;
	goto L1190
L1190:
	;
	v3827 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[21]))
	F__SPI_cursor_operation(m, v3752, v3786, v3785, v3827)
	mBase = m.M
	v3829 = m.ExcPending
	if v3829 != 0 {
		goto L9
	} else {
		goto L1206
	}
L1191:
	;
	F__SPI_cursor_operation(m, v3752, v3786, v3785, v3791)
	mBase = m.M
	v3794 = m.ExcPending
	if v3794 != 0 {
		goto L9
	} else {
		goto L1192
	}
L1192:
	;
	v3796 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[7]))
	v3797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v3798 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v3799 = *(*int32)(unsafe.Add(mBase, uint32(v3798)+4))
	v3803 = *(*int32)(unsafe.Add(mBase, uint32(v3797+v3799<<(uint(int32(2))%32))))
	v3805 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmts[6]))
	if v3805 == int64(0) {
		goto L1193
	} else {
		goto L1194
	}
L1193:
	;
	v3811 = int32(0)
	goto L1195
L1194:
	;
	v3809 = *(*int32)(unsafe.Add(mBase, uint32(v3796)+4))
	v3810 = *(*int32)(unsafe.Add(mBase, uint32(v3809)))
	v3811 = v3810
	goto L1195
L1195:
	;
	v3812 = *(*int32)(unsafe.Add(mBase, uint32(v3796)))
	F_exec_move_row(m, l0, v3803, v3811, v3812)
	mBase = m.M
	v3814 = m.ExcPending
	if v3814 != 0 {
		goto L9
	} else {
		goto L1196
	}
L1196:
	;
	v3815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v3815 != 0 {
		goto L1197
	} else {
		goto L1198
	}
L1197:
	;
	F_SPI_freetuptable(m, v3815)
	mBase = m.M
	v3817 = m.ExcPending
	if v3817 != 0 {
		goto L9
	} else {
		goto L1200
	}
L1198:
	;
	goto L1199
L1199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v3820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v3820 != 0 {
		goto L1201
	} else {
		goto L1202
	}
L1200:
	;
	goto L1199
L1201:
	;
	v3821 = *(*int32)(unsafe.Add(mBase, uint32(v3820)+20))
	F_MemoryContextReset(m, v3821)
	mBase = m.M
	v3823 = m.ExcPending
	if v3823 != 0 {
		goto L9
	} else {
		goto L1204
	}
L1202:
	;
	goto L1203
L1203:
	;
	F_SPI_freetuptable(m, v3796)
	mBase = m.M
	v3825 = m.ExcPending
	if v3825 != 0 {
		goto L9
	} else {
		goto L1205
	}
L1204:
	;
	goto L1203
L1205:
	;
	v3834 = v3805
	goto L1187
L1206:
	;
	v3831 = *(*int64)(unsafe.Add(mBase, _c_F_exec_stmts[6]))
	v3834 = v3831
	goto L1187
L1207:
	;
	v5234 = v3836
	goto L21
L1208:
	;
	v3859 = int32(_a_F_exec_stmts_1)
	v3860 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4]))
	v3862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v3863 = *(*int32)(unsafe.Add(mBase, uint32(v3862)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3863
	v3865 = *(*int32)(unsafe.Add(mBase, uint32(v3855)+40))
	v3866 = F_text_to_cstring(m, v3865)
	mBase = m.M
	v3867 = m.ExcPending
	if v3867 != 0 {
		goto L9
	} else {
		goto L1209
	}
L1209:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[4])) = v3860
	v3870 = F_GetPortalByName(m, v3866)
	mBase = m.M
	v3871 = m.ExcPending
	if v3871 != 0 {
		goto L9
	} else {
		goto L1210
	}
L1210:
	;
	if v3870 == int32(0) {
		goto L34
	} else {
		goto L1211
	}
L1211:
	;
	F_SPI_cursor_close(m, v3870)
	mBase = m.M
	v3875 = m.ExcPending
	if v3875 != 0 {
		goto L9
	} else {
		goto L1212
	}
L1212:
	;
	v5234 = int32(0)
	goto L21
L1213:
	;
	goto L33
L1214:
	;
	F__SPI_commit(m, int32(1))
	mBase = m.M
	v3882 = m.ExcPending
	if v3882 != 0 {
		goto L9
	} else {
		goto L1217
	}
L1215:
	;
	goto L1216
L1216:
	;
	F__SPI_commit(m, int32(0))
	mBase = m.M
	v3885 = m.ExcPending
	if v3885 != 0 {
		goto L9
	} else {
		goto L1218
	}
L1217:
	;
	goto L1213
L1218:
	;
	goto L1213
L1219:
	;
	goto L33
L1220:
	;
	F__SPI_rollback(m, int32(1))
	mBase = m.M
	v3891 = m.ExcPending
	if v3891 != 0 {
		goto L9
	} else {
		goto L1223
	}
L1221:
	;
	goto L1222
L1222:
	;
	F__SPI_rollback(m, int32(0))
	mBase = m.M
	v3894 = m.ExcPending
	if v3894 != 0 {
		goto L9
	} else {
		goto L1224
	}
L1223:
	;
	goto L1219
L1224:
	;
	goto L1219
L1225:
	;
	v3900 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v3900
	F_errmsg_internal(m, int32(_a_F_exec_stmts_50), v32)
	mBase = m.M
	v3904 = m.ExcPending
	if v3904 != 0 {
		goto L9
	} else {
		goto L1226
	}
L1226:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2169), int32(_a_F_exec_stmts_51))
	mBase = m.M
	v3909 = m.ExcPending
	if v3909 != 0 {
		goto L9
	} else {
		goto L1227
	}
L1227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1228:
	;
	v5234 = v3910
	goto L21
L1229:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_52), int32(0))
	mBase = m.M
	v3920 = m.ExcPending
	if v3920 != 0 {
		goto L9
	} else {
		goto L1230
	}
L1230:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2346), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v3925 = m.ExcPending
	if v3925 != 0 {
		goto L9
	} else {
		goto L1231
	}
L1231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1232:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_52), int32(0))
	mBase = m.M
	v3933 = m.ExcPending
	if v3933 != 0 {
		goto L9
	} else {
		goto L1233
	}
L1233:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2350), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v3938 = m.ExcPending
	if v3938 != 0 {
		goto L9
	} else {
		goto L1234
	}
L1234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1235:
	;
	v3943 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v3943
	F_errmsg_internal(m, int32(_a_F_exec_stmts_53), v32+int32(16))
	mBase = m.M
	v3949 = m.ExcPending
	if v3949 != 0 {
		goto L9
	} else {
		goto L1236
	}
L1236:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2358), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v3954 = m.ExcPending
	if v3954 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v3961 = m.ExcPending
	if v3961 != 0 {
		goto L9
	} else {
		goto L1239
	}
L1239:
	;
	v3962 = *(*int32)(unsafe.Add(mBase, uint32(v32)+592))
	v3966 = *(*int32)(unsafe.Add(mBase, uint32(v3962+v225<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v3966
	F_errmsg(m, int32(_a_F_exec_stmts_54), v32-int32(-64))
	mBase = m.M
	v3972 = m.ExcPending
	if v3972 != 0 {
		goto L9
	} else {
		goto L1240
	}
L1240:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2414), int32(_a_F_exec_stmts_5))
	mBase = m.M
	v3977 = m.ExcPending
	if v3977 != 0 {
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
	v3982 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v3983 = F_SPI_result_code_string(m, v412)
	mBase = m.M
	v3984 = m.ExcPending
	if v3984 != 0 {
		goto L9
	} else {
		goto L1243
	}
L1243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v3983
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v3982
	F_errmsg_internal(m, int32(_a_F_exec_stmts_27), v32+int32(32))
	mBase = m.M
	v3991 = m.ExcPending
	if v3991 != 0 {
		goto L9
	} else {
		goto L1244
	}
L1244:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2277), int32(_a_F_exec_stmts_55))
	mBase = m.M
	v3996 = m.ExcPending
	if v3996 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_exec_stmts_56), int32(0))
	mBase = m.M
	v4004 = m.ExcPending
	if v4004 != 0 {
		goto L9
	} else {
		goto L1247
	}
L1247:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2301), int32(_a_F_exec_stmts_55))
	mBase = m.M
	v4009 = m.ExcPending
	if v4009 != 0 {
		goto L9
	} else {
		goto L1248
	}
L1248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1249:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_57), int32(0))
	mBase = m.M
	v4017 = m.ExcPending
	if v4017 != 0 {
		goto L9
	} else {
		goto L1250
	}
L1250:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2306), int32(_a_F_exec_stmts_55))
	mBase = m.M
	v4022 = m.ExcPending
	if v4022 != 0 {
		goto L9
	} else {
		goto L1251
	}
L1251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1252:
	;
	F_errcode(m, int32(33557120))
	mBase = m.M
	v4029 = m.ExcPending
	if v4029 != 0 {
		goto L9
	} else {
		goto L1253
	}
L1253:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_58), int32(0))
	mBase = m.M
	v4033 = m.ExcPending
	if v4033 != 0 {
		goto L9
	} else {
		goto L1254
	}
L1254:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2454), int32(_a_F_exec_stmts_15))
	mBase = m.M
	v4038 = m.ExcPending
	if v4038 != 0 {
		goto L9
	} else {
		goto L1255
	}
L1255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1256:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4045 = m.ExcPending
	if v4045 != 0 {
		goto L9
	} else {
		goto L1257
	}
L1257:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_59), int32(0))
	mBase = m.M
	v4049 = m.ExcPending
	if v4049 != 0 {
		goto L9
	} else {
		goto L1258
	}
L1258:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2754), int32(_a_F_exec_stmts_60))
	mBase = m.M
	v4054 = m.ExcPending
	if v4054 != 0 {
		goto L9
	} else {
		goto L1259
	}
L1259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1260:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4061 = m.ExcPending
	if v4061 != 0 {
		goto L9
	} else {
		goto L1261
	}
L1261:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_61), int32(0))
	mBase = m.M
	v4065 = m.ExcPending
	if v4065 != 0 {
		goto L9
	} else {
		goto L1262
	}
L1262:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2770), int32(_a_F_exec_stmts_60))
	mBase = m.M
	v4070 = m.ExcPending
	if v4070 != 0 {
		goto L9
	} else {
		goto L1263
	}
L1263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1264:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4077 = m.ExcPending
	if v4077 != 0 {
		goto L9
	} else {
		goto L1265
	}
L1265:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_62), int32(0))
	mBase = m.M
	v4081 = m.ExcPending
	if v4081 != 0 {
		goto L9
	} else {
		goto L1266
	}
L1266:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2788), int32(_a_F_exec_stmts_60))
	mBase = m.M
	v4086 = m.ExcPending
	if v4086 != 0 {
		goto L9
	} else {
		goto L1267
	}
L1267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1268:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v4093 = m.ExcPending
	if v4093 != 0 {
		goto L9
	} else {
		goto L1269
	}
L1269:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_63), int32(0))
	mBase = m.M
	v4097 = m.ExcPending
	if v4097 != 0 {
		goto L9
	} else {
		goto L1270
	}
L1270:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2794), int32(_a_F_exec_stmts_60))
	mBase = m.M
	v4102 = m.ExcPending
	if v4102 != 0 {
		goto L9
	} else {
		goto L1271
	}
L1271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1272:
	;
	F_errcode(m, int32(50462852))
	mBase = m.M
	v4109 = m.ExcPending
	if v4109 != 0 {
		goto L9
	} else {
		goto L1273
	}
L1273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v1763
	F_errmsg(m, int32(_a_F_exec_stmts_64), v32+int32(112))
	mBase = m.M
	v4115 = m.ExcPending
	if v4115 != 0 {
		goto L9
	} else {
		goto L1274
	}
L1274:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2928), int32(_a_F_exec_stmts_65))
	mBase = m.M
	v4120 = m.ExcPending
	if v4120 != 0 {
		goto L9
	} else {
		goto L1275
	}
L1275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1276:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4127 = m.ExcPending
	if v4127 != 0 {
		goto L9
	} else {
		goto L1277
	}
L1277:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_66), int32(0))
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L9
	} else {
		goto L1278
	}
L1278:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2951), int32(_a_F_exec_stmts_65))
	mBase = m.M
	v4136 = m.ExcPending
	if v4136 != 0 {
		goto L9
	} else {
		goto L1279
	}
L1279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1280:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4143 = m.ExcPending
	if v4143 != 0 {
		goto L9
	} else {
		goto L1281
	}
L1281:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_67), int32(0))
	mBase = m.M
	v4147 = m.ExcPending
	if v4147 != 0 {
		goto L9
	} else {
		goto L1282
	}
L1282:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2970), int32(_a_F_exec_stmts_65))
	mBase = m.M
	v4152 = m.ExcPending
	if v4152 != 0 {
		goto L9
	} else {
		goto L1283
	}
L1283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1284:
	;
	v4158 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[22]))
	v4159 = F_SPI_result_code_string(m, v4158)
	mBase = m.M
	v4160 = m.ExcPending
	if v4160 != 0 {
		goto L9
	} else {
		goto L1285
	}
L1285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v4159
	F_errmsg_internal(m, int32(_a_F_exec_stmts_68), v32+int32(96))
	mBase = m.M
	v4166 = m.ExcPending
	if v4166 != 0 {
		goto L9
	} else {
		goto L1286
	}
L1286:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2992), int32(_a_F_exec_stmts_65))
	mBase = m.M
	v4171 = m.ExcPending
	if v4171 != 0 {
		goto L9
	} else {
		goto L1287
	}
L1287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1288:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4178 = m.ExcPending
	if v4178 != 0 {
		goto L9
	} else {
		goto L1289
	}
L1289:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_69), int32(0))
	mBase = m.M
	v4182 = m.ExcPending
	if v4182 != 0 {
		goto L9
	} else {
		goto L1290
	}
L1290:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3061), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4187 = m.ExcPending
	if v4187 != 0 {
		goto L9
	} else {
		goto L1291
	}
L1291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1292:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4194 = m.ExcPending
	if v4194 != 0 {
		goto L9
	} else {
		goto L1293
	}
L1293:
	;
	v4195 = F_format_type_be(m, v1913)
	mBase = m.M
	v4196 = m.ExcPending
	if v4196 != 0 {
		goto L9
	} else {
		goto L1294
	}
L1294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+128)) = v4195
	F_errmsg(m, int32(_a_F_exec_stmts_71), v32+int32(128))
	mBase = m.M
	v4202 = m.ExcPending
	if v4202 != 0 {
		goto L9
	} else {
		goto L1295
	}
L1295:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3077), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4207 = m.ExcPending
	if v4207 != 0 {
		goto L9
	} else {
		goto L1296
	}
L1296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1297:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v4214 = m.ExcPending
	if v4214 != 0 {
		goto L9
	} else {
		goto L1298
	}
L1298:
	;
	v4215 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v4216 = *(*int32)(unsafe.Add(mBase, uint32(v1919)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+148)) = v4216
	*(*int32)(unsafe.Add(mBase, uint32(v32)+144)) = v4215
	F_errmsg(m, int32(_a_F_exec_stmts_72), v32+int32(144))
	mBase = m.M
	v4223 = m.ExcPending
	if v4223 != 0 {
		goto L9
	} else {
		goto L1299
	}
L1299:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3094), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4228 = m.ExcPending
	if v4228 != 0 {
		goto L9
	} else {
		goto L1300
	}
L1300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1301:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4235 = m.ExcPending
	if v4235 != 0 {
		goto L9
	} else {
		goto L1302
	}
L1302:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_73), int32(0))
	mBase = m.M
	v4239 = m.ExcPending
	if v4239 != 0 {
		goto L9
	} else {
		goto L1303
	}
L1303:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3120), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4244 = m.ExcPending
	if v4244 != 0 {
		goto L9
	} else {
		goto L1304
	}
L1304:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1305:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4251 = m.ExcPending
	if v4251 != 0 {
		goto L9
	} else {
		goto L1306
	}
L1306:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_74), int32(0))
	mBase = m.M
	v4255 = m.ExcPending
	if v4255 != 0 {
		goto L9
	} else {
		goto L1307
	}
L1307:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3124), int32(_a_F_exec_stmts_70))
	mBase = m.M
	v4260 = m.ExcPending
	if v4260 != 0 {
		goto L9
	} else {
		goto L1308
	}
L1308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1309:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4267 = m.ExcPending
	if v4267 != 0 {
		goto L9
	} else {
		goto L1310
	}
L1310:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_75), int32(0))
	mBase = m.M
	v4271 = m.ExcPending
	if v4271 != 0 {
		goto L9
	} else {
		goto L1311
	}
L1311:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3368), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4276 = m.ExcPending
	if v4276 != 0 {
		goto L9
	} else {
		goto L1312
	}
L1312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1313:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4283 = m.ExcPending
	if v4283 != 0 {
		goto L9
	} else {
		goto L1314
	}
L1314:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_76), int32(0))
	mBase = m.M
	v4287 = m.ExcPending
	if v4287 != 0 {
		goto L9
	} else {
		goto L1315
	}
L1315:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3409), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4292 = m.ExcPending
	if v4292 != 0 {
		goto L9
	} else {
		goto L1316
	}
L1316:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1317:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4299 = m.ExcPending
	if v4299 != 0 {
		goto L9
	} else {
		goto L1318
	}
L1318:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_23), int32(0))
	mBase = m.M
	v4303 = m.ExcPending
	if v4303 != 0 {
		goto L9
	} else {
		goto L1319
	}
L1319:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3469), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4308 = m.ExcPending
	if v4308 != 0 {
		goto L9
	} else {
		goto L1320
	}
L1320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1321:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4315 = m.ExcPending
	if v4315 != 0 {
		goto L9
	} else {
		goto L1322
	}
L1322:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_20), int32(0))
	mBase = m.M
	v4319 = m.ExcPending
	if v4319 != 0 {
		goto L9
	} else {
		goto L1323
	}
L1323:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3505), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4324 = m.ExcPending
	if v4324 != 0 {
		goto L9
	} else {
		goto L1324
	}
L1324:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1325:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v4331 = m.ExcPending
	if v4331 != 0 {
		goto L9
	} else {
		goto L1326
	}
L1326:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_76), int32(0))
	mBase = m.M
	v4335 = m.ExcPending
	if v4335 != 0 {
		goto L9
	} else {
		goto L1327
	}
L1327:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3542), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4340 = m.ExcPending
	if v4340 != 0 {
		goto L9
	} else {
		goto L1328
	}
L1328:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1329:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4347 = m.ExcPending
	if v4347 != 0 {
		goto L9
	} else {
		goto L1330
	}
L1330:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_77), int32(0))
	mBase = m.M
	v4351 = m.ExcPending
	if v4351 != 0 {
		goto L9
	} else {
		goto L1331
	}
L1331:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3562), int32(_a_F_exec_stmts_24))
	mBase = m.M
	v4356 = m.ExcPending
	if v4356 != 0 {
		goto L9
	} else {
		goto L1332
	}
L1332:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1333:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4363 = m.ExcPending
	if v4363 != 0 {
		goto L9
	} else {
		goto L1334
	}
L1334:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_78), int32(0))
	mBase = m.M
	v4367 = m.ExcPending
	if v4367 != 0 {
		goto L9
	} else {
		goto L1335
	}
L1335:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3590), int32(_a_F_exec_stmts_28))
	mBase = m.M
	v4372 = m.ExcPending
	if v4372 != 0 {
		goto L9
	} else {
		goto L1336
	}
L1336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1337:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4379 = m.ExcPending
	if v4379 != 0 {
		goto L9
	} else {
		goto L1338
	}
L1338:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_79), int32(0))
	mBase = m.M
	v4383 = m.ExcPending
	if v4383 != 0 {
		goto L9
	} else {
		goto L1339
	}
L1339:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3663), int32(_a_F_exec_stmts_28))
	mBase = m.M
	v4388 = m.ExcPending
	if v4388 != 0 {
		goto L9
	} else {
		goto L1340
	}
L1340:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1341:
	;
	v4393 = F_SPI_result_code_string(m, v2716)
	mBase = m.M
	v4394 = m.ExcPending
	if v4394 != 0 {
		goto L9
	} else {
		goto L1342
	}
L1342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+196)) = v4393
	*(*int32)(unsafe.Add(mBase, uint32(v32)+192)) = v2688
	F_errmsg_internal(m, int32(_a_F_exec_stmts_43), v32+int32(192))
	mBase = m.M
	v4401 = m.ExcPending
	if v4401 != 0 {
		goto L9
	} else {
		goto L1343
	}
L1343:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3684), int32(_a_F_exec_stmts_28))
	mBase = m.M
	v4406 = m.ExcPending
	if v4406 != 0 {
		goto L9
	} else {
		goto L1344
	}
L1344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1346:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_80), int32(0))
	mBase = m.M
	v4416 = m.ExcPending
	if v4416 != 0 {
		goto L9
	} else {
		goto L1347
	}
L1347:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3831), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4421 = m.ExcPending
	if v4421 != 0 {
		goto L9
	} else {
		goto L1348
	}
L1348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1349:
	;
	F_errmsg_internal(m, int32(_a_F_exec_stmts_80), int32(0))
	mBase = m.M
	v4429 = m.ExcPending
	if v4429 != 0 {
		goto L9
	} else {
		goto L1350
	}
L1350:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3855), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4434 = m.ExcPending
	if v4434 != 0 {
		goto L9
	} else {
		goto L1351
	}
L1351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1352:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4441 = m.ExcPending
	if v4441 != 0 {
		goto L9
	} else {
		goto L1353
	}
L1353:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_81), int32(0))
	mBase = m.M
	v4445 = m.ExcPending
	if v4445 != 0 {
		goto L9
	} else {
		goto L1354
	}
L1354:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3876), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4450 = m.ExcPending
	if v4450 != 0 {
		goto L9
	} else {
		goto L1355
	}
L1355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1356:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4457 = m.ExcPending
	if v4457 != 0 {
		goto L9
	} else {
		goto L1357
	}
L1357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+288)) = int32(_a_F_exec_stmts_82)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(288))
	mBase = m.M
	v4464 = m.ExcPending
	if v4464 != 0 {
		goto L9
	} else {
		goto L1358
	}
L1358:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3887), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4469 = m.ExcPending
	if v4469 != 0 {
		goto L9
	} else {
		goto L1359
	}
L1359:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1360:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4476 = m.ExcPending
	if v4476 != 0 {
		goto L9
	} else {
		goto L1361
	}
L1361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+304)) = int32(_a_F_exec_stmts_84)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(304))
	mBase = m.M
	v4483 = m.ExcPending
	if v4483 != 0 {
		goto L9
	} else {
		goto L1362
	}
L1362:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3892), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4488 = m.ExcPending
	if v4488 != 0 {
		goto L9
	} else {
		goto L1363
	}
L1363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1364:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4495 = m.ExcPending
	if v4495 != 0 {
		goto L9
	} else {
		goto L1365
	}
L1365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+320)) = int32(_a_F_exec_stmts_85)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(320))
	mBase = m.M
	v4502 = m.ExcPending
	if v4502 != 0 {
		goto L9
	} else {
		goto L1366
	}
L1366:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3895), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4507 = m.ExcPending
	if v4507 != 0 {
		goto L9
	} else {
		goto L1367
	}
L1367:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1368:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4514 = m.ExcPending
	if v4514 != 0 {
		goto L9
	} else {
		goto L1369
	}
L1369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+336)) = int32(_a_F_exec_stmts_86)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(336))
	mBase = m.M
	v4521 = m.ExcPending
	if v4521 != 0 {
		goto L9
	} else {
		goto L1370
	}
L1370:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3898), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4526 = m.ExcPending
	if v4526 != 0 {
		goto L9
	} else {
		goto L1371
	}
L1371:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1372:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4533 = m.ExcPending
	if v4533 != 0 {
		goto L9
	} else {
		goto L1373
	}
L1373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+352)) = int32(_a_F_exec_stmts_87)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(352))
	mBase = m.M
	v4540 = m.ExcPending
	if v4540 != 0 {
		goto L9
	} else {
		goto L1374
	}
L1374:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3901), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4545 = m.ExcPending
	if v4545 != 0 {
		goto L9
	} else {
		goto L1375
	}
L1375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1376:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4552 = m.ExcPending
	if v4552 != 0 {
		goto L9
	} else {
		goto L1377
	}
L1377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+368)) = int32(_a_F_exec_stmts_88)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(368))
	mBase = m.M
	v4559 = m.ExcPending
	if v4559 != 0 {
		goto L9
	} else {
		goto L1378
	}
L1378:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3904), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4564 = m.ExcPending
	if v4564 != 0 {
		goto L9
	} else {
		goto L1379
	}
L1379:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1380:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4571 = m.ExcPending
	if v4571 != 0 {
		goto L9
	} else {
		goto L1381
	}
L1381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = int32(_a_F_exec_stmts_89)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(384))
	mBase = m.M
	v4578 = m.ExcPending
	if v4578 != 0 {
		goto L9
	} else {
		goto L1382
	}
L1382:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3907), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4583 = m.ExcPending
	if v4583 != 0 {
		goto L9
	} else {
		goto L1383
	}
L1383:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1384:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4590 = m.ExcPending
	if v4590 != 0 {
		goto L9
	} else {
		goto L1385
	}
L1385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = int32(_a_F_exec_stmts_90)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(400))
	mBase = m.M
	v4597 = m.ExcPending
	if v4597 != 0 {
		goto L9
	} else {
		goto L1386
	}
L1386:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3910), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4602 = m.ExcPending
	if v4602 != 0 {
		goto L9
	} else {
		goto L1387
	}
L1387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1388:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4609 = m.ExcPending
	if v4609 != 0 {
		goto L9
	} else {
		goto L1389
	}
L1389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+416)) = int32(_a_F_exec_stmts_91)
	F_errmsg(m, int32(_a_F_exec_stmts_83), v32+int32(416))
	mBase = m.M
	v4616 = m.ExcPending
	if v4616 != 0 {
		goto L9
	} else {
		goto L1390
	}
L1390:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3913), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4621 = m.ExcPending
	if v4621 != 0 {
		goto L9
	} else {
		goto L1391
	}
L1391:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1392:
	;
	v4626 = *(*int32)(unsafe.Add(mBase, uint32(v3017)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+272)) = v4626
	F_errmsg_internal(m, int32(_a_F_exec_stmts_92), v32+int32(272))
	mBase = m.M
	v4632 = m.ExcPending
	if v4632 != 0 {
		goto L9
	} else {
		goto L1393
	}
L1393:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(3916), int32(_a_F_exec_stmts_30))
	mBase = m.M
	v4637 = m.ExcPending
	if v4637 != 0 {
		goto L9
	} else {
		goto L1394
	}
L1394:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1395:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4644 = m.ExcPending
	if v4644 != 0 {
		goto L9
	} else {
		goto L1396
	}
L1396:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_79), int32(0))
	mBase = m.M
	v4648 = m.ExcPending
	if v4648 != 0 {
		goto L9
	} else {
		goto L1397
	}
L1397:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_93), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v4653 = m.ExcPending
	if v4653 != 0 {
		goto L9
	} else {
		goto L1398
	}
L1398:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1399:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4660 = m.ExcPending
	if v4660 != 0 {
		goto L9
	} else {
		goto L1400
	}
L1400:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_94), int32(0))
	mBase = m.M
	v4664 = m.ExcPending
	if v4664 != 0 {
		goto L9
	} else {
		goto L1401
	}
L1401:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_95), int32(_a_F_exec_stmts_38))
	mBase = m.M
	v4669 = m.ExcPending
	if v4669 != 0 {
		goto L9
	} else {
		goto L1402
	}
L1402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1403:
	;
	F_errcode(m, int32(50462852))
	mBase = m.M
	v4676 = m.ExcPending
	if v4676 != 0 {
		goto L9
	} else {
		goto L1404
	}
L1404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+512)) = v3621
	F_errmsg(m, int32(_a_F_exec_stmts_64), v32+int32(512))
	mBase = m.M
	v4682 = m.ExcPending
	if v4682 != 0 {
		goto L9
	} else {
		goto L1405
	}
L1405:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_96), int32(_a_F_exec_stmts_97))
	mBase = m.M
	v4687 = m.ExcPending
	if v4687 != 0 {
		goto L9
	} else {
		goto L1406
	}
L1406:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1407:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4694 = m.ExcPending
	if v4694 != 0 {
		goto L9
	} else {
		goto L1408
	}
L1408:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_66), int32(0))
	mBase = m.M
	v4698 = m.ExcPending
	if v4698 != 0 {
		goto L9
	} else {
		goto L1409
	}
L1409:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_98), int32(_a_F_exec_stmts_97))
	mBase = m.M
	v4703 = m.ExcPending
	if v4703 != 0 {
		goto L9
	} else {
		goto L1410
	}
L1410:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1411:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v4710 = m.ExcPending
	if v4710 != 0 {
		goto L9
	} else {
		goto L1412
	}
L1412:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_67), int32(0))
	mBase = m.M
	v4714 = m.ExcPending
	if v4714 != 0 {
		goto L9
	} else {
		goto L1413
	}
L1413:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_99), int32(_a_F_exec_stmts_97))
	mBase = m.M
	v4719 = m.ExcPending
	if v4719 != 0 {
		goto L9
	} else {
		goto L1414
	}
L1414:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1415:
	;
	v4725 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[22]))
	v4726 = F_SPI_result_code_string(m, v4725)
	mBase = m.M
	v4727 = m.ExcPending
	if v4727 != 0 {
		goto L9
	} else {
		goto L1416
	}
L1416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+496)) = v4726
	F_errmsg_internal(m, int32(_a_F_exec_stmts_68), v32+int32(496))
	mBase = m.M
	v4733 = m.ExcPending
	if v4733 != 0 {
		goto L9
	} else {
		goto L1417
	}
L1417:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_100), int32(_a_F_exec_stmts_97))
	mBase = m.M
	v4738 = m.ExcPending
	if v4738 != 0 {
		goto L9
	} else {
		goto L1418
	}
L1418:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1419:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4745 = m.ExcPending
	if v4745 != 0 {
		goto L9
	} else {
		goto L1420
	}
L1420:
	;
	v4746 = *(*int32)(unsafe.Add(mBase, uint32(v3736)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+528)) = v4746
	F_errmsg(m, int32(_a_F_exec_stmts_101), v32+int32(528))
	mBase = m.M
	v4752 = m.ExcPending
	if v4752 != 0 {
		goto L9
	} else {
		goto L1421
	}
L1421:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_102), int32(_a_F_exec_stmts_103))
	mBase = m.M
	v4757 = m.ExcPending
	if v4757 != 0 {
		goto L9
	} else {
		goto L1422
	}
L1422:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1423:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v4764 = m.ExcPending
	if v4764 != 0 {
		goto L9
	} else {
		goto L1424
	}
L1424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v3748
	F_errmsg(m, int32(_a_F_exec_stmts_104), v32+int32(544))
	mBase = m.M
	v4770 = m.ExcPending
	if v4770 != 0 {
		goto L9
	} else {
		goto L1425
	}
L1425:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_105), int32(_a_F_exec_stmts_103))
	mBase = m.M
	v4775 = m.ExcPending
	if v4775 != 0 {
		goto L9
	} else {
		goto L1426
	}
L1426:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1427:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4782 = m.ExcPending
	if v4782 != 0 {
		goto L9
	} else {
		goto L1428
	}
L1428:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_106), int32(0))
	mBase = m.M
	v4786 = m.ExcPending
	if v4786 != 0 {
		goto L9
	} else {
		goto L1429
	}
L1429:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_107), int32(_a_F_exec_stmts_103))
	mBase = m.M
	v4791 = m.ExcPending
	if v4791 != 0 {
		goto L9
	} else {
		goto L1430
	}
L1430:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1431:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v4798 = m.ExcPending
	if v4798 != 0 {
		goto L9
	} else {
		goto L1432
	}
L1432:
	;
	v4799 = *(*int32)(unsafe.Add(mBase, uint32(v3855)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v4799
	F_errmsg(m, int32(_a_F_exec_stmts_101), v32+int32(560))
	mBase = m.M
	v4805 = m.ExcPending
	if v4805 != 0 {
		goto L9
	} else {
		goto L1433
	}
L1433:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_108), int32(_a_F_exec_stmts_109))
	mBase = m.M
	v4810 = m.ExcPending
	if v4810 != 0 {
		goto L9
	} else {
		goto L1434
	}
L1434:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1435:
	;
	F_errcode(m, int32(259))
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		goto L9
	} else {
		goto L1436
	}
L1436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+576)) = v3866
	F_errmsg(m, int32(_a_F_exec_stmts_104), v32+int32(576))
	mBase = m.M
	v4823 = m.ExcPending
	if v4823 != 0 {
		goto L9
	} else {
		goto L1437
	}
L1437:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(_a_F_exec_stmts_110), int32(_a_F_exec_stmts_109))
	mBase = m.M
	v4828 = m.ExcPending
	if v4828 != 0 {
		goto L9
	} else {
		goto L1438
	}
L1438:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1439:
	;
	v5234 = int32(0)
	goto L21
L1440:
	;
	v4904 = int32(0)
	v4928 = v1579
	goto L30
L1441:
	;
	goto L1442
L1442:
	;
	v4867 = int32(1)
	v4868 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v4868 == int32(0) {
		v4904 = v4867
		v4928 = v1579
		goto L30
	} else {
		goto L1443
	}
L1443:
	;
	v4873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4868))))
	v4876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4863))))
	if base.B2i32(v4873 == int32(0))|base.B2i32(v4873 != v4876) != 0 {
		v4894 = v4873
		v4895 = v4876
		goto L1445
	} else {
		goto L1446
	}
L1444:
	;
	if v4894-v4895 != 0 {
		v4904 = v4867
		v4928 = v1579
		goto L30
	} else {
		goto L1451
	}
L1445:
	;
	goto L1444
L1446:
	;
	v4879 = v4868
	v4880 = v4863
	goto L1447
L1447:
	;
	v4883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4880)+1)))
	v4884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4879)+1)))
	if v4884 == int32(0) {
		v4894 = v4884
		v4895 = v4883
		goto L1445
	} else {
		goto L1449
	}
L1448:
	;
	v4894 = v4884
	v4895 = v4883
	goto L1445
L1449:
	;
	v4887 = int32(1)
	if v4884 == v4883 {
		v4879 = v4879 + v4887
		v4880 = v4880 + v4887
		goto L1447
	} else {
		goto L1450
	}
L1450:
	;
	goto L1448
L1451:
	;
	v4897 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v4897
	v4904 = v4897
	v4928 = v1579
	goto L30
L1452:
	;
	v5234 = v4904
	goto L21
L1453:
	;
	v5234 = int32(0)
	goto L21
L1454:
	;
	goto L1455
L1455:
	;
	v4945 = int32(1)
	v4946 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v4946 == int32(0) {
		v5234 = v4945
		goto L21
	} else {
		goto L1456
	}
L1456:
	;
	v4951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4946))))
	v4954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4941))))
	if base.B2i32(v4951 == int32(0))|base.B2i32(v4951 != v4954) != 0 {
		v4972 = v4951
		v4973 = v4954
		goto L1458
	} else {
		goto L1459
	}
L1457:
	;
	if v4972-v4973 != 0 {
		v5234 = v4945
		goto L21
	} else {
		goto L1464
	}
L1458:
	;
	goto L1457
L1459:
	;
	v4957 = v4946
	v4958 = v4941
	goto L1460
L1460:
	;
	v4961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4958)+1)))
	v4962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4957)+1)))
	if v4962 == int32(0) {
		v4972 = v4962
		v4973 = v4961
		goto L1458
	} else {
		goto L1462
	}
L1461:
	;
	v4972 = v4962
	v4973 = v4961
	goto L1458
L1462:
	;
	v4965 = int32(1)
	if v4962 == v4961 {
		v4957 = v4957 + v4965
		v4958 = v4958 + v4965
		goto L1460
	} else {
		goto L1463
	}
L1463:
	;
	goto L1461
L1464:
	;
	goto L28
L1465:
	;
	v5036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122)+49)))
	if v5036 != int32(1) {
		goto L1468
	} else {
		goto L1469
	}
L1466:
	;
	goto L1467
L1467:
	;
	v5069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+24)))
	if v5069 == int32(0) {
		goto L25
	} else {
		goto L1477
	}
L1468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1122)+52)) = int32(0)
	v5063 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1122)+48)) = uint16(v5063)
	*(*int64)(unsafe.Add(mBase, uint32(v1122)+40)) = int64(0)
	goto L1467
L1469:
	;
	v5039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122)+48)))
	if v5039 != 0 {
		goto L1470
	} else {
		goto L1471
	}
L1470:
	;
	v5056 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+40))
	F_pfree(m, v5056)
	mBase = m.M
	v5058 = m.ExcPending
	if v5058 != 0 {
		goto L9
	} else {
		goto L1476
	}
L1471:
	;
	v5040 = *(*int32)(unsafe.Add(mBase, uint32(v1122)+24))
	v5041 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5040)+12)))
	if v5041 != int32(_a_F_exec_stmts_16) {
		goto L1470
	} else {
		goto L1472
	}
L1472:
	;
	v5044 = *(*int64)(unsafe.Add(mBase, uint32(v1122)+40))
	v5045 = base.I32_wrap_i64(v5044)
	v5046 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5045))))
	if v5046 != int32(1) {
		goto L1470
	} else {
		goto L1473
	}
L1473:
	;
	v5049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5045)+1)))
	if v5049 != int32(3) {
		goto L1470
	} else {
		goto L1474
	}
L1474:
	;
	F_DeleteExpandedObject(m, v5044)
	mBase = m.M
	v5053 = m.ExcPending
	if v5053 != 0 {
		goto L9
	} else {
		goto L1475
	}
L1475:
	;
	goto L1468
L1476:
	;
	goto L1468
L1477:
	;
	v5103 = v89 + int32(28)
	goto L26
L1478:
	;
	v5234 = v5105
	goto L21
L1479:
	;
	F_errcode(m, int32(2))
	mBase = m.M
	v5113 = m.ExcPending
	if v5113 != 0 {
		goto L9
	} else {
		goto L1480
	}
L1480:
	;
	F_errmsg(m, int32(_a_F_exec_stmts_111), int32(0))
	mBase = m.M
	v5117 = m.ExcPending
	if v5117 != 0 {
		goto L9
	} else {
		goto L1481
	}
L1481:
	;
	F_errhint(m, int32(_a_F_exec_stmts_112), int32(0))
	mBase = m.M
	v5121 = m.ExcPending
	if v5121 != 0 {
		goto L9
	} else {
		goto L1482
	}
L1482:
	;
	F_errfinish(m, int32(_a_F_exec_stmts_4), int32(2661), int32(_a_F_exec_stmts_113))
	mBase = m.M
	v5126 = m.ExcPending
	if v5126 != 0 {
		goto L9
	} else {
		goto L1483
	}
L1483:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1484:
	;
	v5234 = v5189
	goto L21
L1485:
	;
	F_SPI_freetuptable(m, v5220)
	mBase = m.M
	v5222 = m.ExcPending
	if v5222 != 0 {
		goto L9
	} else {
		goto L1488
	}
L1486:
	;
	goto L1487
L1487:
	;
	v5223 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v5223
	v5226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v5226 == v5223 {
		v5234 = v5223
		goto L21
	} else {
		goto L1489
	}
L1488:
	;
	goto L1487
L1489:
	;
	v5229 = *(*int32)(unsafe.Add(mBase, uint32(v5226)+20))
	F_MemoryContextReset(m, v5229)
	mBase = m.M
	v5231 = m.ExcPending
	if v5231 != 0 {
		goto L9
	} else {
		goto L1490
	}
L1490:
	;
	v5234 = v5223
	goto L21
L1491:
	;
	if v5234 == int32(0) {
		goto L1495
	} else {
		goto L1496
	}
L1492:
	;
	v5266 = *(*int32)(unsafe.Add(mBase, uint32(v5263)+16))
	if v5266 == int32(0) {
		v5273 = v5262
		goto L1491
	} else {
		goto L1493
	}
L1493:
	;
	m.T0[v5266].(func(*base.Module, int32, int32))(m, l0, v89)
	mBase = m.M
	v5270 = m.ExcPending
	if v5270 != 0 {
		goto L9
	} else {
		goto L1494
	}
L1494:
	;
	v5272 = *(*int32)(unsafe.Add(mBase, _c_F_exec_stmts[1]))
	v5273 = v5272
	goto L1491
L1495:
	;
	v5278 = v79 + int32(1)
	v5279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5279 <= v5278 {
		goto L2
	} else {
		goto L1498
	}
L1496:
	;
	goto L1497
L1497:
	;
	goto L12
L1498:
	;
	v60 = v5273
	v79 = v5278
	goto L11
}
