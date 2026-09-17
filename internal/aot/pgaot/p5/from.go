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
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v923 int64
	_ = v923
	var v926 int64
	_ = v926
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int64
	_ = v935
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v979 int64
	_ = v979
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1005 int32
	_ = v1005
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
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
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
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
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1176 int64
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int64
	_ = v1188
	var v1196 int32
	_ = v1196
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1246 int32
	_ = v1246
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1388 int32
	_ = v1388
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1437 int32
	_ = v1437
	var v1467 int32
	_ = v1467
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1517 int32
	_ = v1517
	var v1518 int64
	_ = v1518
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
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
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1600 int64
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1667 int64
	_ = v1667
	var v1669 int64
	_ = v1669
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1697 int32
	_ = v1697
	var v1703 int32
	_ = v1703
	var v1712 int32
	_ = v1712
	var v1717 int32
	_ = v1717
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
	v1712 = m.ExcPending
	if v1712 != 0 {
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
	if v1533&int32(1) != 0 {
		goto L354
	} else {
		goto L355
	}
L159:
	;
	v1533 = v1528
	v1534 = v1527
	v1535 = v1525
	v1536 = v1526
	v1539 = v1532
	v1541 = v1531
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
	v1525 = v601
	v1526 = v551
	v1527 = v552
	v1528 = v556
	v1531 = v577
	v1532 = int32(0)
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
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	v1158 = int32(0)
	if v517|base.B2i32(v1157 == v1158) == v1158 {
		goto L287
	} else {
		goto L288
	}
L201:
	;
	v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v601)+4)))
	v950 = v948 & int32(_a_F_CopyFrom_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v601)+4)) = uint16(v950)
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v601)+12))
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v952)))
	*(*uint16)(unsafe.Add(mBase, uint32(v601)+6)) = uint16(v953)
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
	v923 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v923 <= int64(0) {
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
	*(*int64)(unsafe.Add(mBase, uint32(v890+int32(48))+232)) = v887
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	*(*int32)(unsafe.Add(mBase, uint32(v890))) = v913 + v902
	v919 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v919 - v902
	goto L205
L208:
	;
	v926 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	if base.Ui64(v926) <= base.Ui64(v923) {
		goto L139
	} else {
		goto L209
	}
L209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v935 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+64)) = v935
	F_errmsg(m, int32(_a_F_CopyFrom_16), v39-int32(-64))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(1172), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
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
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v955)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v601)+36)) = v956
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v46
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v960 == int32(0) {
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
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v964 == int32(0) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v532)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v968
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v964)+20))
	v973 = m.T0[v972].(func(*base.Module, int32, int32, int32) int32)(m, v964, v532, v39+int32(143))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v46
	if v973 != 0 {
		goto L215
	} else {
		goto L219
	}
L219:
	;
	v979 = v581 + int64(1)
	v982 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[7]))
	if v982 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v581 = v979
	goto L139
L221:
	;
	goto L220
L222:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyFrom[8])))
	if v986&int32(1) == int32(0) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v991 = int32(_a_F_CopyFrom_15)
	v993 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	v994 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v993 + v994
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v982)))
	*(*int32)(unsafe.Add(mBase, uint32(v982))) = v997 + v994
	*(*int64)(unsafe.Add(mBase, uint32(v982+int32(24))+232)) = v979
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v982)))
	*(*int32)(unsafe.Add(mBase, uint32(v982))) = v1005 + v994
	v1011 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1011 - v994
	goto L221
L224:
	;
	v1152 = int32(1)
	v1153 = F_ExecBRInsertTriggers(m, v41, v1146, v1145)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L285
	}
L225:
	;
	v1016 = F_ExecFindPartition(m, v305, v289, v359, v601, v41)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
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
	if v1016 != v577 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+52))
	if v1019 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	v1088 = v552
	v1089 = v556
	v1090 = v557
	v1091 = v577
	goto L231
L231:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v1092 != 0 {
		goto L259
	} else {
		goto L260
	}
L232:
	;
	v1028 = int32(1)
	if v1027&v1028|(v1026&v1028|base.B2i32(v516 != int32(2))) != 0 {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	v1022 = int32(0)
	v1026 = v1022
	v1027 = v1022
	goto L232
L234:
	;
	goto L235
L235:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019)+8)))
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019)+10)))
	v1026 = v1024
	v1027 = v1025
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
	v1064 = int32(0)
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	if v515|base.B2i32(v1065 == v1064) != 0 {
		v1077 = v1064
		goto L236
	} else {
		goto L250
	}
L238:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+84))
	if v1036 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+104))
	if v1037 < int32(2) {
		goto L237
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1040 = int32(1)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+208))
	if v1041 != 0 {
		v1077 = v1040
		goto L236
	} else {
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	v1044 = F_palloc(m, int32(_a_F_CopyFrom_10))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v1046 = int32(0)
	base.MemoryFill(m, v1044, v1046, int32(4000))
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+4000)) = v1016
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+84))
	if v1050 == v1046 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1053 = F_GetBulkInsertState(m)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	v1055 = int32(0)
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+4008)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+4004)) = v1055
	*(*int32)(unsafe.Add(mBase, uint32(v1016)+208)) = v1044
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	v1061 = F_lappend(m, v1060, v1044)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L1
	} else {
		goto L249
	}
L248:
	;
	v1055 = v1053
	goto L247
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v1061
	v1077 = v1040
	goto L236
L250:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(96), v1016, v39+int32(88))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1077 = v1064
	goto L236
L252:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v1078 != 0 {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	goto L254
L254:
	;
	v1088 = v1077
	v1089 = v1027
	v1090 = v1026
	v1091 = v1016
	goto L231
L255:
	;
	F_ReleaseBuffer(m, v1078)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
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
	if v1090&int32(1) != 0 {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	goto L261
L261:
	;
	v1098 = F_ExecGetRootToChildMap(m, v1016, v41)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L1
	} else {
		goto L265
	}
L262:
	;
	v1096 = int32(0)
	goto L264
L263:
	;
	v1096 = v601
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1092)+4)) = v1096
	goto L261
L265:
	;
	v1100 = int32(0)
	if base.B2i32(v516 != v1100)&v1088 == v1100 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+8))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+36)) = v1135
	if v1090&int32(1) != 0 {
		v1145 = v1132
		v1146 = v1016
		v1147 = v1088
		v1148 = v1089
		v1151 = v1091
		goto L224
	} else {
		goto L283
	}
L267:
	;
	if v1098 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+208))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+4008))
	v1115 = v1111 + v1112<<(uint(int32(2))%32)
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1115)))
	if v1116 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L270:
	;
	v1132 = v601
	goto L266
L271:
	;
	goto L272
L272:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+8))
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+204))
	v1109 = F_execute_attr_map_slot(m, v1107, v601, v1108)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v1132 = v1109
	goto L266
L274:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+8))
	v1121 = F_table_slot_create(m, v1119, int32(0))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	v1124 = v1116
	goto L276
L276:
	;
	if v1098 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1115))) = v1121
	v1124 = v1121
	goto L276
L278:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+8))
	v1126 = F_execute_attr_map_slot(m, v1125, v601, v1124)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+8))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+32))
	m.T0[v1129].(func(*base.Module, int32, int32))(m, v1124, v601)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L282
	}
L281:
	;
	v1132 = v1126
	goto L266
L282:
	;
	v1132 = v1124
	goto L266
L283:
	;
	v1525 = v1132
	v1526 = v1016
	v1527 = v1088
	v1528 = v1089
	v1531 = v1091
	v1532 = int32(0)
	goto L159
L284:
	;
	v1145 = v601
	v1146 = v551
	v1147 = v552
	v1148 = v556
	v1151 = v577
	goto L224
L285:
	;
	if v1153 == int32(0) {
		v551 = v1146
		v552 = v1147
		v556 = v1148
		v557 = v1152
		v577 = v1151
		goto L139
	} else {
		goto L286
	}
L286:
	;
	v1533 = v1148
	v1534 = v1147
	v1535 = v1145
	v1536 = v1146
	v1539 = v1152
	v1541 = v1151
	goto L158
L287:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(96), int32(0), v39+int32(88))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v39)+128))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[5])) = v1171
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v1173 == int32(0) {
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
	v1176 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	if v1176 == int64(0) {
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v1179 < int32(0) {
		goto L291
	} else {
		goto L294
	}
L294:
	;
	v1184 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	if v1184 == int32(0) {
		goto L291
	} else {
		goto L296
	}
L296:
	;
	v1188 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+80)) = v1188
	F_errmsg_plural(m, int32(_a_F_CopyFrom_17), int32(_a_F_CopyFrom_18), base.I32_wrap_i64(v1188), v39+int32(80))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(1477), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
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
	v1204 = m.ExcPending
	if v1204 != 0 {
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
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	F_ExecASInsertTriggers(m, v41, v289, v1207)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
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
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
	F_ExecResetTupleTable(m, v1212, int32(0))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v289)+84))
	if v1216 == int32(0) {
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
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1216)+80))
	if v1219 == int32(0) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	m.T0[v1219].(func(*base.Module, int32, int32))(m, v41, v289)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	goto L306
L310:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	if v1227 == int32(0) {
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
	F_list_free(m, v1227)
	mBase = m.M
	v1388 = m.ExcPending
	if v1388 != 0 {
		goto L1
	} else {
		goto L336
	}
L314:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+4))
	if v1230 <= int32(0) {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v39)+120))
	v1246 = int32(0)
	goto L316
L316:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+12))
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1271+v1246<<(uint(int32(2))%32))))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+4000))
	v1277 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1276)+208)) = v1277
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+84))
	if v1279 == v1277 {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	goto L313
L318:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+4004))
	F_FreeBulkInsertState(m, v1282)
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v1287 = int32(0)
	goto L322
L321:
	;
	goto L320
L322:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1275+v1287<<(uint(int32(2))%32))))
	if v1325 != 0 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+84))
	if v1333 != 0 {
		goto L329
	} else {
		goto L330
	}
L324:
	;
	F_ExecDropSingleTupleTableSlot(m, v1325)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
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
	v1329 = v1287 + int32(1)
	if v1329 != int32(1000) {
		v1287 = v1329
		goto L322
	} else {
		goto L328
	}
L328:
	;
	goto L326
L329:
	;
	F_pfree(m, v1275)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L1
	} else {
		goto L334
	}
L330:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+8))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1334)+188))
	if v1335 == int32(0) {
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1335)+108))
	if v1338 == int32(0) {
		goto L329
	} else {
		goto L332
	}
L332:
	;
	m.T0[v1338].(func(*base.Module, int32, int32))(m, v1334, v1233)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L333
	}
L333:
	;
	goto L329
L334:
	;
	v1348 = v1246 + int32(1)
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+4))
	if v1348 < v1349 {
		v1246 = v1348
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
	v1426 = m.ExcPending
	if v1426 != 0 {
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
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L1
	} else {
		goto L341
	}
L340:
	;
	goto L339
L341:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v1430 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1431 = int32(0)
	v1437 = v1430
	goto L345
L343:
	;
	goto L344
L344:
	;
	F_FreeExecutorState(m, v41)
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L1
	} else {
		goto L352
	}
L345:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1467+v1431<<(uint(int32(2))%32))))
	if v1471 != 0 {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	goto L344
L347:
	;
	F_relation_close(m, v1471, int32(0))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L1
	} else {
		goto L350
	}
L348:
	;
	v1476 = v1437
	goto L349
L349:
	;
	v1478 = v1431 + int32(1)
	if base.Ui32(v1478) < base.Ui32(v1476) {
		v1431 = v1478
		v1437 = v1476
		goto L345
	} else {
		goto L351
	}
L350:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v1476 = v1475
	goto L349
L351:
	;
	goto L346
L352:
	;
	v1518 = *(*int64)(unsafe.Add(mBase, uint32(v39)+88))
	m.G0 = v39 + int32(144)
	return v1518
L353:
	;
	v1667 = *(*int64)(unsafe.Add(mBase, uint32(v39)+88))
	v1669 = v1667 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+88)) = v1669
	v1674 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[7]))
	if v1674 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L354:
	;
	v1544 = F_ExecIRInsertTriggers(m, v41, v1536, v1535)
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L1
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+8))
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1546)+52))
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+16))
	if v1548 == int32(0) {
		v1558 = v1546
		goto L358
	} else {
		goto L359
	}
L357:
	;
	goto L353
L358:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+84))
	if v1559 != 0 {
		v1567 = v1558
		goto L362
	} else {
		goto L363
	}
L359:
	;
	v1551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548)+17)))
	if v1551 != int32(1) {
		v1558 = v1546
		goto L358
	} else {
		goto L360
	}
L360:
	;
	F_ExecComputeStoredGenerated(m, v1536, v41, v1535, int32(3))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+8))
	v1558 = v1557
	goto L358
L362:
	;
	v1568 = int32(0)
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+48))
	v1574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1573)+131)))
	if base.B2i32(base.B2i32(v359 == v1568)|v1539 == v1568)|base.B2i32(v1574 != int32(1)) == v1568 {
		goto L366
	} else {
		goto L367
	}
L363:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+52))
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+16))
	if v1561 == int32(0) {
		v1567 = v1558
		goto L362
	} else {
		goto L364
	}
L364:
	;
	F_ExecConstraints(m, v1536, v1535, v41)
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+8))
	v1567 = v1566
	goto L362
L366:
	;
	v1581 = F_ExecPartitionCheck(m, v1536, v1535, v41, int32(1))
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L1
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1583 = int32(1)
	if (base.B2i32(v516 == v1583)|v1534)&v1583 != 0 {
		goto L370
	} else {
		goto L371
	}
L369:
	;
	goto L368
L370:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1535)+8))
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1588)+28))
	m.T0[v1589].(func(*base.Module, int32))(m, v1535)
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L1
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+84))
	if v1625 != 0 {
		goto L380
	} else {
		goto L381
	}
L373:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+208))
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+4008))
	v1600 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v1593+v1594<<(uint(int32(3))%32)+int32(4016)))) = v1600
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1593)+4008))
	v1603 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1593)+4008)) = v1602 + v1603
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v1607 = v1592 + v1606
	*(*int32)(unsafe.Add(mBase, uint32(v39)+104)) = v1607
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	v1611 = v1609 + v1603
	*(*int32)(unsafe.Add(mBase, uint32(v39)+100)) = v1611
	v1613 = int32(0)
	if v1611 <= int32(999) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	if v1607 < int32(_a_F_CopyFrom_19) {
		v551 = v1536
		v552 = v1534
		v556 = v1613
		v557 = v1539
		v577 = v1541
		goto L139
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(96), v1536, v39+int32(88))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L1
	} else {
		goto L378
	}
L377:
	;
	goto L376
L378:
	;
	v551 = v1536
	v552 = v1534
	v556 = v1613
	v557 = v1539
	v577 = v1541
	goto L139
L379:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	F_ExecARInsertTriggers(m, v41, v1536, v1654, v1657, v1658)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L1
	} else {
		goto L388
	}
L380:
	;
	v1626 = int32(0)
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+52))
	v1630 = m.T0[v1629].(func(*base.Module, int32, int32, int32, int32) int32)(m, v41, v1536, v1535, v1626)
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L1
	} else {
		goto L383
	}
L381:
	;
	goto L382
L382:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+8))
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+188))
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1639)+80))
	m.T0[v1640].(func(*base.Module, int32, int32, int32, int32, int32))(m, v1638, v1535, v48, v264, v514)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L1
	} else {
		goto L385
	}
L383:
	;
	if v1630 == int32(0) {
		v551 = v1536
		v552 = v1626
		v556 = v1626
		v557 = v1539
		v577 = v1541
		goto L139
	} else {
		goto L384
	}
L384:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+8))
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1630)+36)) = v1635
	v1654 = v1630
	v1657 = int32(0)
	goto L379
L385:
	;
	v1643 = int32(0)
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+12))
	if v1644 <= v1643 {
		v1654 = v1535
		v1657 = v1643
		goto L379
	} else {
		goto L386
	}
L386:
	;
	v1647 = int32(0)
	v1652 = F_ExecInsertIndexTuples(m, v1536, v1535, v41, v1647, v1647, v1647, v1647, v1647)
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	v1654 = v1535
	v1657 = v1652
	goto L379
L388:
	;
	F_list_free(m, v1657)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	goto L353
L390:
	;
	v551 = v1536
	v552 = v1534
	v556 = v1533
	v557 = v1539
	v577 = v1541
	goto L139
L391:
	;
	goto L390
L392:
	;
	v1678 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyFrom[8])))
	if v1678&int32(1) == int32(0) {
		goto L391
	} else {
		goto L393
	}
L393:
	;
	v1683 = int32(_a_F_CopyFrom_15)
	v1685 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	v1686 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1685 + v1686
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1674)))
	*(*int32)(unsafe.Add(mBase, uint32(v1674))) = v1689 + v1686
	*(*int64)(unsafe.Add(mBase, uint32(v1674+int32(16))+232)) = v1669
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1674)))
	*(*int32)(unsafe.Add(mBase, uint32(v1674))) = v1697 + v1686
	v1703 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1703 - v1686
	goto L391
L394:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(835), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
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
