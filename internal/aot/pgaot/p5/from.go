package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyFrom(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v264 int32
	_ = v264
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
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
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
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
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
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
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v577 int32
	_ = v577
	var v581 int64
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v647 int32
	_ = v647
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v733 int32
	_ = v733
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v792 int32
	_ = v792
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int64
	_ = v887
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v931 int64
	_ = v931
	var v934 int64
	_ = v934
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 int64
	_ = v943
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
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
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v987 int64
	_ = v987
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
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
	var v1044 int32
	_ = v1044
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1192 int64
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1204 int64
	_ = v1204
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1262 int32
	_ = v1262
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1404 int32
	_ = v1404
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1453 int32
	_ = v1453
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1533 int32
	_ = v1533
	var v1534 int64
	_ = v1534
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1616 int64
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1683 int64
	_ = v1683
	var v1685 int64
	_ = v1685
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1727 int32
	_ = v1727
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	v2 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(144)
	m.G0 = v39
	v41 = F_CreateExecutorState(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0]))
	v48 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+120)) = int32(0)
	v52 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+112)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v39)+104)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v39)+96)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v39)+88)) = v52
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+48))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+119)))
	switch v62 - int32(102) {
	case 0, 10:
		v128 = v2
		goto L5
	default:
		goto L7
	case 12:
		goto L6
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v77
	F_errmsg(m, int32(_a_F_CopyFrom_0), v39+int32(48))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L1
	} else {
		goto L394
	}
L5:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+61)))
	if v129 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L6:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	if v120 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L7:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+76))
	if v65 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	switch v62 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L6
	default:
		v128 = v2
		goto L5
	}
L9:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+10)))
	if v66 != 0 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+48))
	v77 = v75 + int32(4)
	switch v62 - int32(109) {
	case 0:
		goto L17
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L15
	case 9:
		goto L18
	default:
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v77
	F_errmsg(m, int32(_a_F_CopyFrom_1), v39)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L25
	}
L16:
	;
	if v62 == int32(83) {
		goto L4
	} else {
		goto L24
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v77
	F_errmsg(m, int32(_a_F_CopyFrom_2), v39+int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L22
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v77
	F_errmsg(m, int32(_a_F_CopyFrom_3), v39+int32(16))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errhint(m, int32(_a_F_CopyFrom_4), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(825), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(830), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	goto L15
L25:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(840), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	if v123 == int32(0) {
		v128 = v2
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v128 = int32(2)
	goto L5
L30:
	;
	goto L29
L31:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v289)+52))
	if v519 != 0 {
		goto L131
	} else {
		goto L132
	}
L32:
	;
	v514 = v2
	v515 = v386
	v516 = int32(1)
	v517 = v2
	v518 = v2
	goto L31
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L127
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L123
	}
L35:
	;
	switch v62 - int32(102) {
	case 0:
		goto L39
	default:
		goto L38
	case 10:
		goto L40
	}
L36:
	;
	v264 = v128
	goto L37
L37:
	;
	v280 = int32(1)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v284 = F_bms_make_singleton(m, v280)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L69
	}
L38:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L49
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L45
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_CopyFrom_7), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(879), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_CopyFrom_8), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(886), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[1]))
	if v169 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v172 = v170
	goto L52
L51:
	;
	v172 = int32(0)
	goto L52
L52:
	;
	if v172 != 0 {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	v173 = m.G0
	v175 = v173 - int32(32)
	m.G0 = v175
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[2]))
	F_hash_seq_init(m, v175+int32(12), v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L55
L55:
	;
	v221 = F_hash_seq_search(m, v175+int32(12))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	m.G0 = v175 + int32(32)
	if v221 != 0 {
		goto L34
	} else {
		goto L62
	}
L57:
	;
	if v221 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)+64))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+80))
	if v224 != int32(2) {
		goto L55
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L56
L61:
	;
	goto L60
L62:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+32))
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[3]))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
	goto L63
L63:
	;
	if v231 != v234 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+36))
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[3]))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	goto L67
L65:
	;
	goto L66
L66:
	;
	v264 = v128 | int32(4)
	goto L37
L67:
	;
	if v237 != v240 {
		goto L33
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	F_ExecInitRangeTable(m, v41, v281, v282, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v289 = F_palloc0(m, int32(216))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = int32(388)
	F_ExecInitResultRelation(m, v41, v289, int32(1))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v297 = int32(0)
	F_CheckValidResultRel(m, v289, int32(3), v297, v297)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_ExecOpenIndices(m, v289, int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v305 = F_palloc0(m, int32(264))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+120)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v305)+116)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v305)+112)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+104)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+8)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v305))) = int64(396)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v289)+84))
	if v316 == int32(0) {
		v336 = v280
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289)+104)) = v336
	v338 = int32(_a_F_CopyFrom_9)
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[4])) = v340 + int32(1)
	goto L86
L77:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v316)+76))
	if v319 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	m.T0[v319].(func(*base.Module, int32, int32))(m, v305, v289)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	v325 = v316
	goto L80
L80:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+60))
	if v326 == int32(0) {
		v336 = v280
		goto L76
	} else {
		goto L83
	}
L81:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v289)+84))
	if v322 == int32(0) {
		v336 = v280
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v325 = v322
	goto L80
L83:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v325)+56))
	if v329 == int32(0) {
		v336 = v280
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v332 = m.T0[v326].(func(*base.Module, int32) int32)(m, v289)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v336 = v332
	goto L76
L86:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+76))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v344)+56))
	v348 = F_MakeTransitionCaptureState(m, v345, v346, int32(3))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+204)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v348
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+48))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+119)))
	if v354 == int32(112) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v357 = F_ExecSetupPartitionTupleRouting(m, v41, v352)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	v359 = v2
	goto L90
L90:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v360 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v359 = v357
	goto L90
L92:
	;
	v361 = F_ExecInitQual(m, v360, v305)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v289)+52))
	if v364 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v361
	goto L94
L96:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	v439 = F_table_slot_create(m, v436, v41+int32(104))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L121
	}
L97:
	;
	v365 = int32(1)
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+8)))
	if v366 != 0 {
		v434 = v365
		v435 = v2
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v289)+84))
	if v369 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+10)))
	if v367 != 0 {
		v434 = v365
		v435 = v2
		goto L96
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v377 = int32(0)
	if base.B2i32(v359 == v377)|base.B2i32(v364 == v377) != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v372 = int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v289)+104))
	if v373 != v372 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v434 = v372
	v435 = v2
	goto L96
L105:
	;
	v386 = int32(1)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+244)))
	if v387 != 0 {
		v434 = v386
		v435 = v2
		goto L96
	} else {
		goto L108
	}
L106:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+25)))
	if v382 == int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v434 = int32(1)
	v435 = v2
	goto L96
L108:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v389 = F_contain_volatile_functions(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	if v389 != 0 {
		v434 = v386
		v435 = v2
		goto L96
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+120)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v39)+116)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v39)+108)) = l0
	v395 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+104)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v39)+96)) = int64(0)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+48))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+119)))
	if v402 != int32(112) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v406 = F_palloc(m, int32(_a_F_CopyFrom_10))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v359 == int32(0) {
		goto L32
	} else {
		goto L120
	}
L114:
	;
	v408 = int32(0)
	base.MemoryFill(m, v406, v408, int32(4000))
	*(*int32)(unsafe.Add(mBase, uint32(v406)+4000)) = v289
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v289)+84))
	if v412 == v408 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v415 = F_GetBulkInsertState(m)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	v417 = v395
	goto L117
L117:
	;
	v418 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v406)+4008)) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v406)+4004)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(v289)+208)) = v406
	v423 = F_lappend(m, v418, v406)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	v417 = v415
	goto L117
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v423
	goto L113
L120:
	;
	v434 = int32(0)
	v435 = int32(2)
	goto L96
L121:
	;
	v441 = F_GetBulkInsertState(m)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v514 = v441
	v515 = v434
	v516 = v435
	v517 = v434
	v518 = v439
	goto L31
L123:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_CopyFrom_11), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(900), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errmsg(m, int32(_a_F_CopyFrom_12), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(906), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+8)))
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+10)))
	v523 = v520
	v524 = v521
	goto L133
L132:
	;
	v523 = v2
	v524 = int32(0)
	goto L133
L133:
	;
	F_ExecBSInsertTriggers(m, v41, v289)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v41)+152))
	if v527 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v530 = F_MakePerTupleExprContext(m, v41)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	v532 = v527
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+136)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = int32(517)
	v536 = int32(_a_F_CopyFrom_13)
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[5])) = v39 + int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v537
	v551 = v289
	v552 = v2
	v556 = v524
	v557 = v523
	v577 = v2
	v581 = int64(0)
	goto L139
L138:
	;
	v532 = v530
	goto L137
L139:
	;
	v583 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[6]))
	if v583 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v41)+152))
	if v586 != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L143
L145:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v586)+20))
	F_MemoryContextReset(m, v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	if v517|base.B2i32(v359 != int32(0)) != 0 {
		v601 = v518
		goto L149
	} else {
		goto L150
	}
L148:
	;
	goto L147
L149:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v41)+152))
	if v604 != 0 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v551)+208))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+4008))
	v594 = v590 + v591<<(uint(int32(2))%32)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
	if v595 != 0 {
		v601 = v595
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v551)+8))
	v598 = F_table_slot_create(m, v596, int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v594))) = v598
	v601 = v598
	goto L149
L153:
	;
	v607 = v604
	goto L155
L154:
	;
	v605 = F_MakePerTupleExprContext(m, v41)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L156
	}
L155:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v608
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v601)+8))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)+12))
	m.T0[v611].(func(*base.Module, int32))(m, v601)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	v607 = v605
	goto L155
L157:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v601)+20))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)+52))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)))
	v619 = v617 << (uint(int32(16)) % 32)
	v621 = v619 >> (uint(int32(14)) % 32)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v624 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+208)))
	v625 = base.I32_extend16_s(v617)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	if v626&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v621)) == int32(0) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	if v1549&int32(1) != 0 {
		goto L354
	} else {
		goto L355
	}
L159:
	;
	v1549 = v1544
	v1550 = v1543
	v1551 = v1541
	v1552 = v1542
	v1555 = v1548
	v1557 = v1547
	goto L158
L160:
	;
	v658 = int32(0)
	v659 = base.B2i32(v625 == v658)
	if v659 == v658 {
		goto L170
	} else {
		goto L171
	}
L161:
	;
	if v619 == int32(0) {
		goto L160
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if v621 == int32(0) {
		goto L160
	} else {
		goto L169
	}
L164:
	;
	v638 = v626 + v621
	v640 = v626 + int32(4)
	if base.Ui32(v640) < base.Ui32(v638) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v642 = v638
	goto L167
L166:
	;
	v642 = v640
	goto L167
L167:
	;
	v647 = (v626^int32(-1)+v642)&int32(-4) + int32(4)
	if v647 == int32(0) {
		goto L160
	} else {
		goto L168
	}
L168:
	;
	base.MemoryFill(m, v626, int32(0), v647)
	goto L160
L169:
	;
	base.MemoryFill(m, v626, int32(0), v621)
	goto L160
L170:
	;
	base.MemoryFill(m, v614, int32(1), v625)
	goto L172
L171:
	;
	goto L172
L172:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v665 = int32(3)
	if v664&v665|v617&v665|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v625)) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v697)+8))
	v699 = m.T0[v698].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v532, v626, v614)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L184
	}
L174:
	;
	if v619 == int32(0) {
		goto L173
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	if v625 == v658 {
		goto L173
	} else {
		goto L182
	}
L177:
	;
	v679 = v625 + v664
	v681 = v664 + int32(4)
	if base.Ui32(v681) < base.Ui32(v679) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v683 = v679
	goto L180
L179:
	;
	v683 = v681
	goto L180
L180:
	;
	v688 = (v664^int32(-1)+v683)&int32(-4) + int32(4)
	if v688 == int32(0) {
		goto L173
	} else {
		goto L181
	}
L181:
	;
	base.MemoryFill(m, v664, int32(0), v688)
	goto L173
L182:
	;
	base.MemoryFill(m, v664, int32(0), v625)
	goto L173
L183:
	;
	if v699 != 0 {
		goto L198
	} else {
		goto L199
	}
L184:
	;
	if v699 == int32(0) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v703 = base.I32_extend16_s(v624)
	if v703 <= int32(0) {
		goto L183
	} else {
		goto L186
	}
L186:
	;
	v706 = int32(0)
	if v703 != int32(1) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v716 = v706
	v733 = int32(0)
	goto L190
L188:
	;
	v792 = v706
	goto L189
L189:
	;
	v824 = int32(2)
	v826 = v623 + v792<<(uint(v824)%32)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v622+v827<<(uint(v824)%32))))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v831)+20))
	v834 = m.T0[v833].(func(*base.Module, int32, int32, int32) int32)(m, v831, v532, v827+v614)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L196
	}
L190:
	;
	v748 = int32(2)
	v750 = v623 + v716<<(uint(v748)%32)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v622+v751<<(uint(v748)%32))))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v755)+20))
	v758 = m.T0[v757].(func(*base.Module, int32, int32, int32) int32)(m, v755, v532, v614+v751)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L192
	}
L191:
	;
	if v703&int32(1) == int32(0) {
		goto L183
	} else {
		goto L195
	}
L192:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v761 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v626+v760<<(uint(v761)%32)))) = v758
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v622+v765<<(uint(v761)%32))))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v769)+20))
	v772 = m.T0[v771].(func(*base.Module, int32, int32, int32) int32)(m, v769, v532, v614+v765)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
	v775 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v626+v774<<(uint(v775)%32)))) = v772
	v780 = v716 + v775
	v782 = v733 + v775
	if v782 != v703&int32(-2) {
		v716 = v780
		v733 = v782
		goto L190
	} else {
		goto L194
	}
L194:
	;
	goto L191
L195:
	;
	v792 = v780
	goto L189
L196:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v826)))
	*(*int32)(unsafe.Add(mBase, uint32(v626+v836<<(uint(int32(2))%32)))) = v834
	goto L183
L197:
	;
	v1541 = v601
	v1542 = v551
	v1543 = v552
	v1544 = v556
	v1547 = v577
	v1548 = int32(0)
	goto L159
L198:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v877 != int32(1) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	goto L200
L200:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	v1174 = int32(0)
	if v517|base.B2i32(v1173 == v1174) == v1174 {
		goto L287
	} else {
		goto L288
	}
L201:
	;
	v956 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v601)+4)))
	v958 = v956 & int32(_a_F_CopyFrom_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v601)+4)) = uint16(v958)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v601)+12))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v960)))
	*(*uint16)(unsafe.Add(mBase, uint32(v601)+6)) = uint16(v961)
	goto L214
L202:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880)+4)))
	if v881 != int32(1) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v884 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v880)+4)) = uint8(v884)
	v887 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	v890 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[7]))
	if v890 == v884 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v931 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v931 <= int64(0) {
		goto L139
	} else {
		goto L208
	}
L205:
	;
	goto L204
L206:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyFrom[8])))
	if v894&int32(1) == int32(0) {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v899 = int32(_a_F_CopyFrom_15)
	v901 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	v902 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v901 + v902
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	*(*int32)(unsafe.Add(mBase, uint32(v890))) = v905 + v902
	v909 = int32(0)
	v911 = int32(_a_F_CopyFrom_16)
	v912 = base.AtomicRmwOr32(m, v909, v911, v909)
	*(*int64)(unsafe.Add(mBase, uint32(v890+int32(48))+232)) = v887
	v920 = base.AtomicRmwOr32(m, v909, v911, v909)
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	*(*int32)(unsafe.Add(mBase, uint32(v890))) = v921 + v902
	v927 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v927 - v902
	goto L205
L208:
	;
	v934 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	if base.Ui64(v934) <= base.Ui64(v931) {
		goto L139
	} else {
		goto L209
	}
L209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v943 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+64)) = v943
	F_errmsg(m, int32(_a_F_CopyFrom_17), v39-int32(-64))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(1172), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
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
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v963)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v601)+36)) = v964
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v46
	v968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v968 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	if v359 != 0 {
		goto L225
	} else {
		goto L226
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v532)+4)) = v601
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v972 == int32(0) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v532)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v976
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v972)+20))
	v981 = m.T0[v980].(func(*base.Module, int32, int32, int32) int32)(m, v972, v532, v39+int32(143))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v46
	if v981 != 0 {
		goto L215
	} else {
		goto L219
	}
L219:
	;
	v987 = v581 + int64(1)
	v990 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[7]))
	if v990 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v581 = v987
	goto L139
L221:
	;
	goto L220
L222:
	;
	v994 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyFrom[8])))
	if v994&int32(1) == int32(0) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v999 = int32(_a_F_CopyFrom_15)
	v1001 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	v1002 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1001 + v1002
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v990)))
	*(*int32)(unsafe.Add(mBase, uint32(v990))) = v1005 + v1002
	v1009 = int32(0)
	v1011 = int32(_a_F_CopyFrom_16)
	v1012 = base.AtomicRmwOr32(m, v1009, v1011, v1009)
	*(*int64)(unsafe.Add(mBase, uint32(v990+int32(24))+232)) = v987
	v1020 = base.AtomicRmwOr32(m, v1009, v1011, v1009)
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v990)))
	*(*int32)(unsafe.Add(mBase, uint32(v990))) = v1021 + v1002
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1027 - v1002
	goto L221
L224:
	;
	v1168 = int32(1)
	v1169 = F_ExecBRInsertTriggers(m, v41, v1162, v1161)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L285
	}
L225:
	;
	v1032 = F_ExecFindPartition(m, v305, v289, v359, v601, v41)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	if v557&int32(1) == int32(0) {
		goto L197
	} else {
		goto L284
	}
L228:
	;
	if v1032 != v577 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+52))
	if v1035 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	v1104 = v552
	v1105 = v556
	v1106 = v557
	v1107 = v577
	goto L231
L231:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v1108 != 0 {
		goto L259
	} else {
		goto L260
	}
L232:
	;
	v1044 = int32(1)
	if v1043&v1044|(v1042&v1044|base.B2i32(v516 != int32(2))) != 0 {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	v1038 = int32(0)
	v1042 = v1038
	v1043 = v1038
	goto L232
L234:
	;
	goto L235
L235:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035)+8)))
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035)+10)))
	v1042 = v1040
	v1043 = v1041
	goto L232
L236:
	;
	if v514 != 0 {
		goto L252
	} else {
		goto L253
	}
L237:
	;
	v1080 = int32(0)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	if v515|base.B2i32(v1081 == v1080) != 0 {
		v1093 = v1080
		goto L236
	} else {
		goto L250
	}
L238:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+84))
	if v1052 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+104))
	if v1053 < int32(2) {
		goto L237
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1056 = int32(1)
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+208))
	if v1057 != 0 {
		v1093 = v1056
		goto L236
	} else {
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	v1060 = F_palloc(m, int32(_a_F_CopyFrom_10))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v1062 = int32(0)
	base.MemoryFill(m, v1060, v1062, int32(4000))
	*(*int32)(unsafe.Add(mBase, uint32(v1060)+4000)) = v1032
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+84))
	if v1066 == v1062 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1069 = F_GetBulkInsertState(m)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	v1071 = int32(0)
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1060)+4008)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1060)+4004)) = v1071
	*(*int32)(unsafe.Add(mBase, uint32(v1032)+208)) = v1060
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	v1077 = F_lappend(m, v1076, v1060)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L249
	}
L248:
	;
	v1071 = v1069
	goto L247
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v1077
	v1093 = v1056
	goto L236
L250:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(96), v1032, v39+int32(88))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1093 = v1080
	goto L236
L252:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v1094 != 0 {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	goto L254
L254:
	;
	v1104 = v1093
	v1105 = v1043
	v1106 = v1042
	v1107 = v1032
	goto L231
L255:
	;
	F_ReleaseBuffer(m, v1094)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v514)+12)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v514)+4)) = int64(-4294967296)
	goto L254
L258:
	;
	goto L257
L259:
	;
	if v1106&int32(1) != 0 {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	goto L261
L261:
	;
	v1114 = F_ExecGetRootToChildMap(m, v1032, v41)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L265
	}
L262:
	;
	v1112 = int32(0)
	goto L264
L263:
	;
	v1112 = v601
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+4)) = v1112
	goto L261
L265:
	;
	v1116 = int32(0)
	if base.B2i32(v516 != v1116)&v1104 == v1116 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+8))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1148)+36)) = v1151
	if v1106&int32(1) != 0 {
		v1161 = v1148
		v1162 = v1032
		v1163 = v1104
		v1164 = v1105
		v1167 = v1107
		goto L224
	} else {
		goto L283
	}
L267:
	;
	if v1114 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+208))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4008))
	v1131 = v1127 + v1128<<(uint(int32(2))%32)
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1131)))
	if v1132 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L270:
	;
	v1148 = v601
	goto L266
L271:
	;
	goto L272
L272:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+8))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+204))
	v1125 = F_execute_attr_map_slot(m, v1123, v601, v1124)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v1148 = v1125
	goto L266
L274:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+8))
	v1137 = F_table_slot_create(m, v1135, int32(0))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	v1140 = v1132
	goto L276
L276:
	;
	if v1114 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1131))) = v1137
	v1140 = v1137
	goto L276
L278:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+8))
	v1142 = F_execute_attr_map_slot(m, v1141, v601, v1140)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1140)+8))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+32))
	m.T0[v1145].(func(*base.Module, int32, int32))(m, v1140, v601)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L1
	} else {
		goto L282
	}
L281:
	;
	v1148 = v1142
	goto L266
L282:
	;
	v1148 = v1140
	goto L266
L283:
	;
	v1541 = v1148
	v1542 = v1032
	v1543 = v1104
	v1544 = v1105
	v1547 = v1107
	v1548 = int32(0)
	goto L159
L284:
	;
	v1161 = v601
	v1162 = v551
	v1163 = v552
	v1164 = v556
	v1167 = v577
	goto L224
L285:
	;
	if v1169 == int32(0) {
		v551 = v1162
		v552 = v1163
		v556 = v1164
		v557 = v1168
		v577 = v1167
		goto L139
	} else {
		goto L286
	}
L286:
	;
	v1549 = v1164
	v1550 = v1163
	v1551 = v1161
	v1552 = v1162
	v1555 = v1168
	v1557 = v1167
	goto L158
L287:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(96), int32(0), v39+int32(88))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L1
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v39)+128))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[5])) = v1187
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v1189 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	goto L289
L291:
	;
	if v514 != 0 {
		goto L299
	} else {
		goto L300
	}
L292:
	;
	v1192 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	if v1192 == int64(0) {
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v1195 < int32(0) {
		goto L291
	} else {
		goto L294
	}
L294:
	;
	v1200 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	if v1200 == int32(0) {
		goto L291
	} else {
		goto L296
	}
L296:
	;
	v1204 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+80)) = v1204
	F_errmsg_plural(m, int32(_a_F_CopyFrom_18), int32(_a_F_CopyFrom_19), base.I32_wrap_i64(v1204), v39+int32(80))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(1477), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	goto L291
L299:
	;
	F_FreeBulkInsertState(m, v514)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L1
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v46
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	F_ExecASInsertTriggers(m, v41, v289, v1223)
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L1
	} else {
		goto L303
	}
L302:
	;
	goto L301
L303:
	;
	F_AfterTriggerEndQuery(m, v41)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
	F_ExecResetTupleTable(m, v1228, int32(0))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v289)+84))
	if v1232 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	if v517 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L307:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+80))
	if v1235 == int32(0) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	m.T0[v1235].(func(*base.Module, int32, int32))(m, v41, v289)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	goto L306
L310:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	if v1243 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L311:
	;
	goto L312
L312:
	;
	if v359 != 0 {
		goto L337
	} else {
		goto L338
	}
L313:
	;
	F_list_free(m, v1243)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L1
	} else {
		goto L336
	}
L314:
	;
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+4))
	if v1246 <= int32(0) {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v39)+120))
	v1262 = int32(0)
	goto L316
L316:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+12))
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1287+v1262<<(uint(int32(2))%32))))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+4000))
	v1293 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+208)) = v1293
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+84))
	if v1295 == v1293 {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	goto L313
L318:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1291)+4004))
	F_FreeBulkInsertState(m, v1298)
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L1
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v1303 = int32(0)
	goto L322
L321:
	;
	goto L320
L322:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1291+v1303<<(uint(int32(2))%32))))
	if v1341 != 0 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+84))
	if v1349 != 0 {
		goto L329
	} else {
		goto L330
	}
L324:
	;
	F_ExecDropSingleTupleTableSlot(m, v1341)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L1
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	goto L323
L327:
	;
	v1345 = v1303 + int32(1)
	if v1345 != int32(1000) {
		v1303 = v1345
		goto L322
	} else {
		goto L328
	}
L328:
	;
	goto L326
L329:
	;
	F_pfree(m, v1291)
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L1
	} else {
		goto L334
	}
L330:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+8))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1350)+188))
	if v1351 == int32(0) {
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1351)+108))
	if v1354 == int32(0) {
		goto L329
	} else {
		goto L332
	}
L332:
	;
	m.T0[v1354].(func(*base.Module, int32, int32))(m, v1350, v1249)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	goto L329
L334:
	;
	v1364 = v1262 + int32(1)
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1243)+4))
	if v1364 < v1365 {
		v1262 = v1364
		goto L316
	} else {
		goto L335
	}
L335:
	;
	goto L317
L336:
	;
	goto L312
L337:
	;
	F_ExecCleanupTupleRouting(m, v305, v359)
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L1
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	F_ExecCloseResultRelations(m, v41)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L1
	} else {
		goto L341
	}
L340:
	;
	goto L339
L341:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v1446 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1447 = int32(0)
	v1453 = v1446
	goto L345
L343:
	;
	goto L344
L344:
	;
	F_FreeExecutorState(m, v41)
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L1
	} else {
		goto L352
	}
L345:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1483+v1447<<(uint(int32(2))%32))))
	if v1487 != 0 {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	goto L344
L347:
	;
	F_relation_close(m, v1487, int32(0))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L1
	} else {
		goto L350
	}
L348:
	;
	v1492 = v1453
	goto L349
L349:
	;
	v1494 = v1447 + int32(1)
	if base.Ui32(v1494) < base.Ui32(v1492) {
		v1447 = v1494
		v1453 = v1492
		goto L345
	} else {
		goto L351
	}
L350:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v1492 = v1491
	goto L349
L351:
	;
	goto L346
L352:
	;
	v1534 = *(*int64)(unsafe.Add(mBase, uint32(v39)+88))
	m.G0 = v39 + int32(144)
	return v1534
L353:
	;
	v1683 = *(*int64)(unsafe.Add(mBase, uint32(v39)+88))
	v1685 = v1683 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+88)) = v1685
	v1690 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[7]))
	if v1690 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L354:
	;
	v1560 = F_ExecIRInsertTriggers(m, v41, v1552, v1551)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L1
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+8))
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1562)+52))
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1563)+16))
	if v1564 == int32(0) {
		v1574 = v1562
		goto L358
	} else {
		goto L359
	}
L357:
	;
	goto L353
L358:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+84))
	if v1575 != 0 {
		v1583 = v1574
		goto L362
	} else {
		goto L363
	}
L359:
	;
	v1567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564)+17)))
	if v1567 != int32(1) {
		v1574 = v1562
		goto L358
	} else {
		goto L360
	}
L360:
	;
	F_ExecComputeStoredGenerated(m, v1552, v41, v1551, int32(3))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+8))
	v1574 = v1573
	goto L358
L362:
	;
	v1584 = int32(0)
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+48))
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589)+131)))
	if base.B2i32(base.B2i32(v359 == v1584)|v1555 == v1584)|base.B2i32(v1590 != int32(1)) == v1584 {
		goto L366
	} else {
		goto L367
	}
L363:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+52))
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1576)+16))
	if v1577 == int32(0) {
		v1583 = v1574
		goto L362
	} else {
		goto L364
	}
L364:
	;
	F_ExecConstraints(m, v1552, v1551, v41)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+8))
	v1583 = v1582
	goto L362
L366:
	;
	v1597 = F_ExecPartitionCheck(m, v1552, v1551, v41, int32(1))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L1
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1599 = int32(1)
	if (base.B2i32(v516 == v1599)|v1550)&v1599 != 0 {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	goto L368
L370:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1551)+8))
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1604)+28))
	m.T0[v1605].(func(*base.Module, int32))(m, v1551)
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L1
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+84))
	if v1641 != 0 {
		goto L380
	} else {
		goto L381
	}
L373:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+208))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1609)+4008))
	v1616 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v1609+v1610<<(uint(int32(3))%32)+int32(4016)))) = v1616
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1609)+4008))
	v1619 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1609)+4008)) = v1618 + v1619
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v1623 = v1608 + v1622
	*(*int32)(unsafe.Add(mBase, uint32(v39)+104)) = v1623
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	v1627 = v1625 + v1619
	*(*int32)(unsafe.Add(mBase, uint32(v39)+100)) = v1627
	v1629 = int32(0)
	if v1627 <= int32(999) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	if v1623 < int32(_a_F_CopyFrom_20) {
		v551 = v1552
		v552 = v1550
		v556 = v1629
		v557 = v1555
		v577 = v1557
		goto L139
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(96), v1552, v39+int32(88))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L1
	} else {
		goto L378
	}
L377:
	;
	goto L376
L378:
	;
	v551 = v1552
	v552 = v1550
	v556 = v1629
	v557 = v1555
	v577 = v1557
	goto L139
L379:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	F_ExecARInsertTriggers(m, v41, v1552, v1670, v1673, v1674)
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L1
	} else {
		goto L388
	}
L380:
	;
	v1642 = int32(0)
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1641)+52))
	v1646 = m.T0[v1645].(func(*base.Module, int32, int32, int32, int32) int32)(m, v41, v1552, v1551, v1642)
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L1
	} else {
		goto L383
	}
L381:
	;
	goto L382
L382:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+8))
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v1654)+188))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1655)+80))
	m.T0[v1656].(func(*base.Module, int32, int32, int32, int32, int32))(m, v1654, v1551, v48, v264, v514)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L1
	} else {
		goto L385
	}
L383:
	;
	if v1646 == int32(0) {
		v551 = v1552
		v552 = v1642
		v556 = v1642
		v557 = v1555
		v577 = v1557
		goto L139
	} else {
		goto L384
	}
L384:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+8))
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1650)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1646)+36)) = v1651
	v1670 = v1646
	v1673 = int32(0)
	goto L379
L385:
	;
	v1659 = int32(0)
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+12))
	if v1660 <= v1659 {
		v1670 = v1551
		v1673 = v1659
		goto L379
	} else {
		goto L386
	}
L386:
	;
	v1663 = int32(0)
	v1668 = F_ExecInsertIndexTuples(m, v1552, v1551, v41, v1663, v1663, v1663, v1663, v1663)
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	v1670 = v1551
	v1673 = v1668
	goto L379
L388:
	;
	F_list_free(m, v1673)
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	goto L353
L390:
	;
	v551 = v1552
	v552 = v1550
	v556 = v1549
	v557 = v1555
	v577 = v1557
	goto L139
L391:
	;
	goto L390
L392:
	;
	v1694 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyFrom[8])))
	if v1694&int32(1) == int32(0) {
		goto L391
	} else {
		goto L393
	}
L393:
	;
	v1699 = int32(_a_F_CopyFrom_15)
	v1701 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	v1702 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1701 + v1702
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1690)))
	*(*int32)(unsafe.Add(mBase, uint32(v1690))) = v1705 + v1702
	v1709 = int32(0)
	v1711 = int32(_a_F_CopyFrom_16)
	v1712 = base.AtomicRmwOr32(m, v1709, v1711, v1709)
	*(*int64)(unsafe.Add(mBase, uint32(v1690+int32(16))+232)) = v1685
	v1720 = base.AtomicRmwOr32(m, v1709, v1711, v1709)
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1690)))
	*(*int32)(unsafe.Add(mBase, uint32(v1690))) = v1721 + v1702
	v1727 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1727 - v1702
	goto L391
L394:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(835), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_makeFromExpr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc0(m, int32(12))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(65)
		return v5
	}
}
