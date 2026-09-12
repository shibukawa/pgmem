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
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v268 int32
	_ = v268
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v384 int32
	_ = v384
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
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
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
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v574 int32
	_ = v574
	var v580 int64
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
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
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v777 int32
	_ = v777
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v874 int64
	_ = v874
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v908 int64
	_ = v908
	var v911 int64
	_ = v911
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int64
	_ = v920
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int64
	_ = v964
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v988 int32
	_ = v988
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
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
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1157 int64
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int64
	_ = v1169
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1227 int32
	_ = v1227
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1369 int32
	_ = v1369
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1420 int32
	_ = v1420
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1498 int32
	_ = v1498
	var v1499 int64
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1511 int32
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
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
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
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1576 int64
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1643 int64
	_ = v1643
	var v1645 int64
	_ = v1645
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1671 int32
	_ = v1671
	var v1677 int32
	_ = v1677
	var v1686 int32
	_ = v1686
	var v1691 int32
	_ = v1691
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
	v46 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
	F_errmsg(m, int32(692768), v39+int32(48))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L398
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
	F_errmsg(m, int32(679835), v39)
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
	F_errmsg(m, int32(668495), v39+int32(32))
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
	F_errmsg(m, int32(668234), v39+int32(16))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errhint(m, int32(581948), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(485948), int32(825), int32(282146))
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
	F_errfinish(m, int32(485948), int32(830), int32(282146))
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
	F_errfinish(m, int32(485948), int32(840), int32(282146))
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
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v293)+52))
	if v518 != 0 {
		goto L134
	} else {
		goto L135
	}
L32:
	;
	v513 = v386
	v514 = v2
	v515 = v2
	v516 = v2
	v517 = int32(1)
	goto L31
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L130
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L126
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
	v268 = v128
	goto L37
L37:
	;
	v284 = int32(1)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v288 = F_bms_make_singleton(m, v284)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
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
	F_errmsg(m, int32(386153), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(485948), int32(879), int32(282146))
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
	F_errmsg(m, int32(384380), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(485948), int32(886), int32(282146))
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
	v169 = *(*int32)(unsafe.Add(mBase, _consts[282]))
	if v169 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v174 = base.B2i32(v170 == int32(0))
	goto L52
L51:
	;
	v174 = int32(1)
	goto L52
L52:
	;
	if v174 == int32(0) {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	v177 = m.G0
	v179 = v177 - int32(32)
	m.G0 = v179
	v184 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	F_hash_seq_init(m, v179+int32(12), v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L55
L55:
	;
	v225 = F_hash_seq_search(m, v179+int32(12))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	m.G0 = v179 + int32(32)
	if v225 != 0 {
		goto L34
	} else {
		goto L62
	}
L57:
	;
	if v225 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v225)+64))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+80))
	if v228 != int32(2) {
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
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+32))
	v237 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+8))
	goto L63
L63:
	;
	if v235 != v238 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+36))
	v243 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	goto L67
L65:
	;
	goto L66
L66:
	;
	v268 = v128 | int32(4)
	goto L37
L67:
	;
	if v241 != v244 {
		goto L33
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	F_ExecInitRangeTable(m, v41, v285, v286, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v293 = F_palloc0(m, int32(216))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = int32(388)
	F_ExecInitResultRelation(m, v41, v293, int32(1))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v301 = int32(0)
	F_CheckValidResultRel(m, v293, int32(3), v301, v301)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_ExecOpenIndices(m, v293, int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v309 = F_palloc0(m, int32(264))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+120)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v309)+116)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v309)+112)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v309)+104)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v309)+8)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v309))) = int64(396)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v293)+84))
	if v320 == int32(0) {
		v340 = v284
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+104)) = v340
	v342 = int32(4361588)
	v344 = *(*int32)(unsafe.Add(mBase, _consts[284]))
	*(*int32)(unsafe.Add(mBase, _consts[284])) = v344 + int32(1)
	goto L86
L77:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v320)+76))
	if v323 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	m.T0[v323].(func(*base.Module, int32, int32))(m, v309, v293)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	v329 = v320
	goto L80
L80:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+60))
	if v330 == int32(0) {
		v340 = v284
		goto L76
	} else {
		goto L83
	}
L81:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v293)+84))
	if v326 == int32(0) {
		v340 = v284
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v329 = v326
	goto L80
L83:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v329)+56))
	if v333 == int32(0) {
		v340 = v284
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v336 = m.T0[v330].(func(*base.Module, int32) int32)(m, v293)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v340 = v336
	goto L76
L86:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+76))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v348)+56))
	v352 = F_MakeTransitionCaptureState(m, v349, v350, int32(3))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+204)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(l0)+260)) = v352
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+48))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+119)))
	if v358 == int32(112) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v361 = F_ExecSetupPartitionTupleRouting(m, v41, v356)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	v363 = v2
	goto L90
L90:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v364 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v363 = v361
	goto L90
L92:
	;
	v365 = F_ExecInitQual(m, v364, v309)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v293)+52))
	if v368 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v365
	goto L94
L96:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	v438 = F_table_slot_create(m, v435, v41+int32(104))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L124
	}
L97:
	;
	v369 = int32(1)
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+8)))
	if v370 != 0 {
		v434 = v369
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v293)+84))
	if v373 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+10)))
	if v371 != 0 {
		v434 = v369
		goto L96
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v374 = int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v293)+104))
	if v375 == v374 {
		v434 = v374
		goto L96
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if v363 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	v386 = int32(1)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+244)))
	if v387 != 0 {
		v434 = v386
		goto L96
	} else {
		goto L110
	}
L107:
	;
	if v368 == int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+25)))
	if v384 != 0 {
		v434 = int32(1)
		goto L96
	} else {
		goto L109
	}
L109:
	;
	goto L106
L110:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v389 = F_contain_volatile_functions(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v389 != 0 {
		v434 = v386
		goto L96
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+120)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v39)+116)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v39)+108)) = l0
	v395 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+104)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v39)+96)) = int64(0)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+48))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+119)))
	if v402 != int32(112) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v406 = F_palloc(m, int32(12016))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v363 == int32(0) {
		goto L32
	} else {
		goto L123
	}
L116:
	;
	v411 = F__emscripten_memset_bulkmem(m, v406, base.I32_extend8_s(int32(0)), int32(4000))
	mBase = m.M
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v411)+4000)) = v293
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v293)+84))
	if v413 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v416 = F_GetBulkInsertState(m)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	v418 = v395
	goto L120
L120:
	;
	v419 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v411)+4008)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v411)+4004)) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v293)+208)) = v411
	v424 = F_lappend(m, v419, v411)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L122
	}
L121:
	;
	v418 = v416
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v424
	goto L115
L123:
	;
	v434 = int32(0)
	goto L96
L124:
	;
	v440 = F_GetBulkInsertState(m)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v513 = v434
	v514 = v440
	v515 = v434
	v516 = v438
	v517 = v2
	goto L31
L126:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errmsg(m, int32(9890), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(485948), int32(900), int32(282146))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
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
	F_errcode(m, int32(325))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(250027), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(485948), int32(906), int32(282146))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
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
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+8)))
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+10)))
	v522 = v519
	v523 = v520
	goto L136
L135:
	;
	v522 = v2
	v523 = int32(0)
	goto L136
L136:
	;
	F_ExecBSInsertTriggers(m, v41, v293)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v41)+152))
	if v526 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v529 = F_MakePerTupleExprContext(m, v41)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	v531 = v526
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+136)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = int32(517)
	v535 = int32(4457400)
	v536 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v39 + int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v536
	v548 = v293
	v553 = v2
	v556 = v522
	v562 = v523
	v574 = v2
	v580 = int64(0)
	goto L142
L141:
	;
	v531 = v529
	goto L140
L142:
	;
	v582 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v582 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v41)+152))
	if v585 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L146
L148:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)+20))
	F_MemoryContextReset(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	if v515|base.B2i32(v363 != int32(0)) != 0 {
		v600 = v516
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L150
L152:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v41)+152))
	if v603 != 0 {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v548)+208))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)+4008))
	v593 = v589 + v590<<(uint(int32(2))%32)
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)))
	if v594 != 0 {
		v600 = v594
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v548)+8))
	v597 = F_table_slot_create(m, v595, int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v593))) = v597
	v600 = v597
	goto L152
L156:
	;
	v606 = v603
	goto L158
L157:
	;
	v604 = F_MakePerTupleExprContext(m, v41)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L159
	}
L158:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v607
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v600)+8))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)+12))
	m.T0[v610].(func(*base.Module, int32))(m, v600)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L160
	}
L159:
	;
	v606 = v604
	goto L158
L160:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v600)+20))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)+52))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	v620 = v616 << (uint(int32(16)) % 32) >> (uint(int32(14)) % 32)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+208)))
	v624 = base.I32_extend16_s(v616)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v600)+16))
	if v625&int32(3) != 0 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	if v1514&int32(1) != 0 {
		goto L358
	} else {
		goto L359
	}
L162:
	;
	v1514 = v1511
	v1515 = v1505
	v1516 = v1506
	v1517 = v1508
	v1520 = v1513
	v1522 = v1512
	goto L161
L163:
	;
	v653 = F__emscripten_memset_bulkmem(m, v613, base.I32_extend8_s(int32(1)), v624)
	mBase = m.M
	goto L173
L164:
	;
	v648 = F__emscripten_memset_bulkmem(m, v625, base.I32_extend8_s(int32(0)), v620)
	mBase = m.M
	goto L172
L165:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v620) {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v630 = v620 + v625
	if base.Ui32(v630) <= base.Ui32(v625) {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v636 = v625 + int32(4)
	if base.Ui32(v636) < base.Ui32(v630) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v638 = v630
	goto L170
L169:
	;
	v638 = v636
	goto L170
L170:
	;
	v645 = F__emscripten_memset_bulkmem(m, v625, base.I32_extend8_s(int32(0)), (v625^int32(-1)+v638)&int32(-4)+int32(4))
	mBase = m.M
	goto L171
L171:
	;
	goto L163
L172:
	;
	goto L163
L173:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v654&int32(3) != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v682)+8))
	v684 = m.T0[v683].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v531, v625, v653)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L186
	}
L175:
	;
	v679 = F__emscripten_memset_bulkmem(m, v654, base.I32_extend8_s(int32(0)), v624)
	mBase = m.M
	goto L184
L176:
	;
	if v616&int32(3) != 0 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v624) {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	v661 = v624 + v654
	if base.Ui32(v661) <= base.Ui32(v654) {
		goto L174
	} else {
		goto L179
	}
L179:
	;
	v667 = v654 + int32(4)
	if base.Ui32(v667) < base.Ui32(v661) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v669 = v661
	goto L182
L181:
	;
	v669 = v667
	goto L182
L182:
	;
	v676 = F__emscripten_memset_bulkmem(m, v654, base.I32_extend8_s(int32(0)), (v654^int32(-1)+v669)&int32(-4)+int32(4))
	mBase = m.M
	goto L183
L183:
	;
	goto L174
L184:
	;
	goto L174
L185:
	;
	if v684 != 0 {
		goto L200
	} else {
		goto L201
	}
L186:
	;
	if v684 == int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v688 = base.I32_extend16_s(v623)
	if v688 <= int32(0) {
		goto L185
	} else {
		goto L188
	}
L188:
	;
	v691 = int32(0)
	if v688 != int32(1) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v703 = v691
	v709 = int32(0)
	goto L192
L190:
	;
	v777 = v691
	goto L191
L191:
	;
	if v688&int32(1) == int32(0) {
		goto L185
	} else {
		goto L197
	}
L192:
	;
	v733 = int32(2)
	v735 = v622 + v703<<(uint(v733)%32)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v621+v736<<(uint(v733)%32))))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v740)+20))
	v743 = m.T0[v742].(func(*base.Module, int32, int32, int32) int32)(m, v740, v531, v736+v653)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L194
	}
L193:
	;
	v777 = v767
	goto L191
L194:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v735)))
	v746 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v625+v745<<(uint(v746)%32)))) = v743
	v751 = v735 + int32(4)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v621+v752<<(uint(v746)%32))))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v756)+20))
	v759 = m.T0[v758].(func(*base.Module, int32, int32, int32) int32)(m, v756, v531, v752+v653)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v751)))
	v762 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v625+v761<<(uint(v762)%32)))) = v759
	v767 = v703 + v762
	v769 = v709 + v762
	if v769 != v688&int32(-2) {
		v703 = v767
		v709 = v769
		goto L192
	} else {
		goto L196
	}
L196:
	;
	goto L193
L197:
	;
	v811 = int32(2)
	v813 = v622 + v777<<(uint(v811)%32)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v813)))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v621+v814<<(uint(v811)%32))))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v821 = m.T0[v820].(func(*base.Module, int32, int32, int32) int32)(m, v818, v531, v814+v653)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v813)))
	*(*int32)(unsafe.Add(mBase, uint32(v625+v823<<(uint(int32(2))%32)))) = v821
	goto L185
L199:
	;
	v1505 = v600
	v1506 = v548
	v1508 = v553
	v1511 = v562
	v1512 = v574
	v1513 = int32(0)
	goto L162
L200:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v864 != int32(1) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L202
L202:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	v1139 = int32(0)
	if v515|base.B2i32(v1138 == v1139) == v1139 {
		goto L291
	} else {
		goto L292
	}
L203:
	;
	v933 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v600)+4)))
	v935 = v933 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v600)+4)) = uint16(v935)
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v600)+12))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	*(*uint16)(unsafe.Add(mBase, uint32(v600)+6)) = uint16(v938)
	goto L216
L204:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+4)))
	if v868 != int32(1) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v871 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v867)+4)) = uint8(v871)
	v874 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	v877 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v877 == v871 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v908 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v908 <= int64(0) {
		goto L142
	} else {
		goto L210
	}
L207:
	;
	goto L206
L208:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
	if v881 != int32(1) {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v884 = int32(4459156)
	v886 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v887 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v886 + v887
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v877)))
	*(*int32)(unsafe.Add(mBase, uint32(v877))) = v890 + v887
	*(*int64)(unsafe.Add(mBase, uint32(v877+int32(48))+232)) = v874
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v877)))
	*(*int32)(unsafe.Add(mBase, uint32(v877))) = v898 + v887
	v904 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v904 - v887
	goto L207
L210:
	;
	v911 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	if base.Ui64(v911) <= base.Ui64(v908) {
		goto L142
	} else {
		goto L211
	}
L211:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v920 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+64)) = v920
	F_errmsg(m, int32(11221), v39-int32(-64))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(485948), int32(1172), int32(282146))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v600)+36)) = v941
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v46
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v945 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	if v363 != 0 {
		goto L227
	} else {
		goto L228
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v531)+4)) = v600
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v949 == int32(0) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v531)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v953
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v949)+20))
	v958 = m.T0[v957].(func(*base.Module, int32, int32, int32) int32)(m, v949, v531, v39+int32(143))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v46
	if v958 != 0 {
		goto L217
	} else {
		goto L221
	}
L221:
	;
	v964 = v580 + int64(1)
	v967 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v967 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v580 = v964
	goto L142
L223:
	;
	goto L222
L224:
	;
	v971 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
	if v971 != int32(1) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v974 = int32(4459156)
	v976 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v977 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v976 + v977
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v967)))
	*(*int32)(unsafe.Add(mBase, uint32(v967))) = v980 + v977
	*(*int64)(unsafe.Add(mBase, uint32(v967+int32(24))+232)) = v964
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v967)))
	*(*int32)(unsafe.Add(mBase, uint32(v967))) = v988 + v977
	v994 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v994 - v977
	goto L223
L226:
	;
	v1133 = int32(1)
	v1134 = F_ExecBRInsertTriggers(m, v41, v1126, v1125)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L1
	} else {
		goto L289
	}
L227:
	;
	v999 = F_ExecFindPartition(m, v309, v293, v363, v600, v41)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	if v556&int32(1) == int32(0) {
		goto L199
	} else {
		goto L288
	}
L230:
	;
	if v999 != v574 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v999)+52))
	if v1002 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L232:
	;
	v1069 = v553
	v1070 = v556
	v1071 = v562
	v1072 = v574
	goto L233
L233:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v1073 != 0 {
		goto L263
	} else {
		goto L264
	}
L234:
	;
	if (v1009|v513)&int32(1) != 0 {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v1005 = int32(0)
	v1009 = v1005
	v1010 = v1005
	goto L234
L236:
	;
	goto L237
L237:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002)+8)))
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002)+10)))
	v1009 = v1007
	v1010 = v1008
	goto L234
L238:
	;
	if v514 != 0 {
		goto L256
	} else {
		goto L257
	}
L239:
	;
	v1045 = int32(0)
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	if v513|base.B2i32(v1046 == v1045) != 0 {
		v1058 = v1045
		goto L238
	} else {
		goto L254
	}
L240:
	;
	if v1010&int32(1) != 0 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v999)+84))
	if v1016 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v999)+104))
	if v1017 < int32(2) {
		goto L239
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v1020 = int32(1)
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v999)+208))
	if v1021 != 0 {
		v1058 = v1020
		goto L238
	} else {
		goto L246
	}
L245:
	;
	goto L244
L246:
	;
	v1024 = F_palloc(m, int32(12016))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	v1029 = F__emscripten_memset_bulkmem(m, v1024, base.I32_extend8_s(int32(0)), int32(4000))
	mBase = m.M
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+4000)) = v999
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v999)+84))
	if v1031 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1034 = F_GetBulkInsertState(m)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L1
	} else {
		goto L252
	}
L250:
	;
	v1036 = int32(0)
	goto L251
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+4008)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+4004)) = v1036
	*(*int32)(unsafe.Add(mBase, uint32(v999)+208)) = v1029
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	v1042 = F_lappend(m, v1041, v1029)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L253
	}
L252:
	;
	v1036 = v1034
	goto L251
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = v1042
	v1058 = v1020
	goto L238
L254:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(96), v999, v39+int32(88))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v1058 = v1045
	goto L238
L256:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v1059 != 0 {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L258
L258:
	;
	v1069 = v1058
	v1070 = v1009
	v1071 = v1010
	v1072 = v999
	goto L233
L259:
	;
	F_ReleaseBuffer(m, v1059)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v514)+12)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v514)+4)) = int64(-4294967296)
	goto L258
L262:
	;
	goto L261
L263:
	;
	if v1070&int32(1) != 0 {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	goto L265
L265:
	;
	v1079 = F_ExecGetRootToChildMap(m, v999, v41)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L269
	}
L266:
	;
	v1077 = int32(0)
	goto L268
L267:
	;
	v1077 = v600
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1073)+4)) = v1077
	goto L265
L269:
	;
	if (v515|(v1069^int32(-1)))&int32(1) != 0 {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v999)+8))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1113)+36)) = v1116
	if v1070&int32(1) != 0 {
		v1125 = v1113
		v1126 = v999
		v1128 = v1069
		v1131 = v1071
		v1132 = v1072
		goto L226
	} else {
		goto L287
	}
L271:
	;
	if v1079 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	goto L273
L273:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v999)+208))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1092)+4008))
	v1096 = v1092 + v1093<<(uint(int32(2))%32)
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1096)))
	if v1097 == int32(0) {
		goto L278
	} else {
		goto L279
	}
L274:
	;
	v1113 = v600
	goto L270
L275:
	;
	goto L276
L276:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+8))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v999)+204))
	v1090 = F_execute_attr_map_slot(m, v1088, v600, v1089)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	v1113 = v1090
	goto L270
L278:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v999)+8))
	v1102 = F_table_slot_create(m, v1100, int32(0))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L1
	} else {
		goto L281
	}
L279:
	;
	v1105 = v1097
	goto L280
L280:
	;
	if v1079 != 0 {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1096))) = v1102
	v1105 = v1102
	goto L280
L282:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+8))
	v1107 = F_execute_attr_map_slot(m, v1106, v600, v1105)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L1
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+8))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+32))
	m.T0[v1110].(func(*base.Module, int32, int32))(m, v1105, v600)
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L286
	}
L285:
	;
	v1113 = v1107
	goto L270
L286:
	;
	v1113 = v1105
	goto L270
L287:
	;
	v1505 = v1113
	v1506 = v999
	v1508 = v1069
	v1511 = v1071
	v1512 = v1072
	v1513 = int32(0)
	goto L162
L288:
	;
	v1125 = v600
	v1126 = v548
	v1128 = v553
	v1131 = v562
	v1132 = v574
	goto L226
L289:
	;
	if v1134 == int32(0) {
		v548 = v1126
		v553 = v1128
		v556 = v1133
		v562 = v1131
		v574 = v1132
		goto L142
	} else {
		goto L290
	}
L290:
	;
	v1514 = v1131
	v1515 = v1125
	v1516 = v1126
	v1517 = v1128
	v1520 = v1133
	v1522 = v1132
	goto L161
L291:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(96), int32(0), v39+int32(88))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v39)+128))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v1152
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v1154 == int32(0) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	goto L293
L295:
	;
	if v514 != 0 {
		goto L303
	} else {
		goto L304
	}
L296:
	;
	v1157 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	if v1157 == int64(0) {
		goto L295
	} else {
		goto L297
	}
L297:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v1160 < int32(0) {
		goto L295
	} else {
		goto L298
	}
L298:
	;
	v1165 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	if v1165 == int32(0) {
		goto L295
	} else {
		goto L300
	}
L300:
	;
	v1169 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+80)) = v1169
	F_errmsg_plural(m, int32(11297), int32(11351), base.I32_wrap_i64(v1169), v39+int32(80))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(485948), int32(1477), int32(282146))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	goto L295
L303:
	;
	F_FreeBulkInsertState(m, v514)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L1
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v46
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	F_ExecASInsertTriggers(m, v41, v293, v1188)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L1
	} else {
		goto L307
	}
L306:
	;
	goto L305
L307:
	;
	F_AfterTriggerEndQuery(m, v41)
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
	F_ExecResetTupleTable(m, v1193, int32(0))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v293)+84))
	if v1197 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	if v515 == int32(0) {
		goto L314
	} else {
		goto L315
	}
L311:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+80))
	if v1200 == int32(0) {
		goto L310
	} else {
		goto L312
	}
L312:
	;
	m.T0[v1200].(func(*base.Module, int32, int32))(m, v41, v293)
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	goto L310
L314:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v39)+96))
	if v1208 == int32(0) {
		goto L317
	} else {
		goto L318
	}
L315:
	;
	goto L316
L316:
	;
	if v363 != 0 {
		goto L341
	} else {
		goto L342
	}
L317:
	;
	F_list_free(m, v1208)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L1
	} else {
		goto L340
	}
L318:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1211 <= int32(0) {
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v39)+120))
	v1227 = int32(0)
	goto L320
L320:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+12))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1252+v1227<<(uint(int32(2))%32))))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+4000))
	v1258 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+208)) = v1258
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+84))
	if v1260 == v1258 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	goto L317
L322:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+4004))
	F_FreeBulkInsertState(m, v1263)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L1
	} else {
		goto L325
	}
L323:
	;
	goto L324
L324:
	;
	v1268 = int32(0)
	goto L326
L325:
	;
	goto L324
L326:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1256+v1268<<(uint(int32(2))%32))))
	if v1306 != 0 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+84))
	if v1314 != 0 {
		goto L333
	} else {
		goto L334
	}
L328:
	;
	F_ExecDropSingleTupleTableSlot(m, v1306)
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L1
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	goto L327
L331:
	;
	v1310 = v1268 + int32(1)
	if v1310 != int32(1000) {
		v1268 = v1310
		goto L326
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	F_pfree(m, v1256)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L1
	} else {
		goto L338
	}
L334:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+8))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+188))
	if v1316 == int32(0) {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+108))
	if v1319 == int32(0) {
		goto L333
	} else {
		goto L336
	}
L336:
	;
	m.T0[v1319].(func(*base.Module, int32, int32))(m, v1315, v1214)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	goto L333
L338:
	;
	v1329 = v1227 + int32(1)
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+4))
	if v1329 < v1330 {
		v1227 = v1329
		goto L320
	} else {
		goto L339
	}
L339:
	;
	goto L321
L340:
	;
	goto L316
L341:
	;
	F_ExecCleanupTupleRouting(m, v309, v363)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L1
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	F_ExecCloseResultRelations(m, v41)
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L1
	} else {
		goto L345
	}
L344:
	;
	goto L343
L345:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v1411 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1412 = int32(0)
	v1420 = v1411
	goto L349
L347:
	;
	goto L348
L348:
	;
	F_FreeExecutorState(m, v41)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L356
	}
L349:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1448+v1412<<(uint(int32(2))%32))))
	if v1452 != 0 {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	goto L348
L351:
	;
	F_sequence_close(m, v1452, int32(0))
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L1
	} else {
		goto L354
	}
L352:
	;
	v1457 = v1420
	goto L353
L353:
	;
	v1459 = v1412 + int32(1)
	if base.Ui32(v1459) < base.Ui32(v1457) {
		v1412 = v1459
		v1420 = v1457
		goto L349
	} else {
		goto L355
	}
L354:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v1457 = v1456
	goto L353
L355:
	;
	goto L350
L356:
	;
	v1499 = *(*int64)(unsafe.Add(mBase, uint32(v39)+88))
	m.G0 = v39 + int32(144)
	return v1499
L357:
	;
	v1643 = *(*int64)(unsafe.Add(mBase, uint32(v39)+88))
	v1645 = v1643 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+88)) = v1645
	v1650 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v1650 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L358:
	;
	v1525 = F_ExecIRInsertTriggers(m, v41, v1516, v1515)
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L1
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+8))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1527)+52))
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1528)+16))
	if v1529 == int32(0) {
		v1539 = v1527
		goto L362
	} else {
		goto L363
	}
L361:
	;
	goto L357
L362:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+84))
	if v1540 != 0 {
		v1548 = v1539
		goto L366
	} else {
		goto L367
	}
L363:
	;
	v1532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1529)+17)))
	if v1532 != int32(1) {
		v1539 = v1527
		goto L362
	} else {
		goto L364
	}
L364:
	;
	F_ExecComputeStoredGenerated(m, v1516, v41, v1515, int32(3))
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+8))
	v1539 = v1538
	goto L362
L366:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1548)+48))
	v1550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549)+131)))
	if v1550 != int32(1) {
		goto L370
	} else {
		goto L371
	}
L367:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1539)+52))
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1541)+16))
	if v1542 == int32(0) {
		v1548 = v1539
		goto L366
	} else {
		goto L368
	}
L368:
	;
	F_ExecConstraints(m, v1516, v1515, v41)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L369
	}
L369:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+8))
	v1548 = v1547
	goto L366
L370:
	;
	if (v1517|v517)&int32(1) != 0 {
		goto L374
	} else {
		goto L375
	}
L371:
	;
	v1553 = int32(0)
	if base.B2i32(v363 == v1553)|v1520 == v1553 {
		goto L370
	} else {
		goto L372
	}
L372:
	;
	v1559 = F_ExecPartitionCheck(m, v1516, v1515, v41, int32(1))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	goto L370
L374:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+8))
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1564)+28))
	m.T0[v1565].(func(*base.Module, int32))(m, v1515)
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L1
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+84))
	if v1601 != 0 {
		goto L384
	} else {
		goto L385
	}
L377:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+208))
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1569)+4008))
	v1576 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v1569+v1570<<(uint(int32(3))%32)+int32(4016)))) = v1576
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1569)+4008))
	v1579 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1569)+4008)) = v1578 + v1579
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v39)+104))
	v1583 = v1568 + v1582
	*(*int32)(unsafe.Add(mBase, uint32(v39)+104)) = v1583
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	v1587 = v1585 + v1579
	*(*int32)(unsafe.Add(mBase, uint32(v39)+100)) = v1587
	v1589 = int32(0)
	if v1587 <= int32(999) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	if v1583 < int32(65535) {
		v548 = v1516
		v553 = v1517
		v556 = v1520
		v562 = v1589
		v574 = v1522
		goto L142
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(96), v1516, v39+int32(88))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L1
	} else {
		goto L382
	}
L381:
	;
	goto L380
L382:
	;
	v548 = v1516
	v553 = v1517
	v556 = v1520
	v562 = v1589
	v574 = v1522
	goto L142
L383:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	F_ExecARInsertTriggers(m, v41, v1516, v1630, v1633, v1634)
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L1
	} else {
		goto L392
	}
L384:
	;
	v1602 = int32(0)
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1601)+52))
	v1606 = m.T0[v1605].(func(*base.Module, int32, int32, int32, int32) int32)(m, v41, v1516, v1515, v1602)
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L1
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+8))
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1614)+188))
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1615)+80))
	m.T0[v1616].(func(*base.Module, int32, int32, int32, int32, int32))(m, v1614, v1515, v48, v268, v514)
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L1
	} else {
		goto L389
	}
L387:
	;
	if v1606 == int32(0) {
		v548 = v1516
		v553 = v1602
		v556 = v1520
		v562 = v1602
		v574 = v1522
		goto L142
	} else {
		goto L388
	}
L388:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+8))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1606)+36)) = v1611
	v1630 = v1606
	v1633 = int32(0)
	goto L383
L389:
	;
	v1619 = int32(0)
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1516)+12))
	if v1620 <= v1619 {
		v1630 = v1515
		v1633 = v1619
		goto L383
	} else {
		goto L390
	}
L390:
	;
	v1623 = int32(0)
	v1628 = F_ExecInsertIndexTuples(m, v1516, v1515, v41, v1623, v1623, v1623, v1623, v1623)
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	v1630 = v1515
	v1633 = v1628
	goto L383
L392:
	;
	F_list_free(m, v1633)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	goto L357
L394:
	;
	v548 = v1516
	v553 = v1517
	v556 = v1520
	v562 = v1514
	v574 = v1522
	goto L142
L395:
	;
	goto L394
L396:
	;
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, _consts[49])))
	if v1654 != int32(1) {
		goto L395
	} else {
		goto L397
	}
L397:
	;
	v1657 = int32(4459156)
	v1659 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	v1660 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v1659 + v1660
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1650)))
	*(*int32)(unsafe.Add(mBase, uint32(v1650))) = v1663 + v1660
	*(*int64)(unsafe.Add(mBase, uint32(v1650+int32(16))+232)) = v1645
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1650)))
	*(*int32)(unsafe.Add(mBase, uint32(v1650))) = v1671 + v1660
	v1677 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v1677 - v1660
	goto L395
L398:
	;
	F_errfinish(m, int32(485948), int32(835), int32(282146))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
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
