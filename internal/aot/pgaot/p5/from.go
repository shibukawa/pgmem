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
	var v721 int32
	_ = v721
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
	var v758 int64
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int64
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v797 int32
	_ = v797
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
	var v834 int64
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
	var v981 int64
	_ = v981
	var v982 int32
	_ = v982
	var v989 int64
	_ = v989
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
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
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
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
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int64
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int64
	_ = v1207
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1223 int64
	_ = v1223
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1282 int32
	_ = v1282
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1424 int32
	_ = v1424
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1473 int32
	_ = v1473
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1553 int32
	_ = v1553
	var v1554 int64
	_ = v1554
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
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1636 int64
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1701 int64
	_ = v1701
	var v1703 int64
	_ = v1703
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1745 int32
	_ = v1745
	var v1754 int32
	_ = v1754
	var v1759 int32
	_ = v1759
	v2 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(160)
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
	*(*int32)(unsafe.Add(mBase, uint32(v39)+136)) = int32(0)
	v52 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+128)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v39)+120)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v39)+112)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v39)+104)) = v52
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
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L1
	} else {
		goto L399
	}
L5:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
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
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(827), int32(_a_F_CopyFrom_6))
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
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(832), int32(_a_F_CopyFrom_6))
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
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(842), int32(_a_F_CopyFrom_6))
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
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+268))
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
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(881), int32(_a_F_CopyFrom_6))
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
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(888), int32(_a_F_CopyFrom_6))
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
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = int32(394)
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
	*(*int64)(unsafe.Add(mBase, uint32(v305))) = int64(402)
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+276)) = v348
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
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v361
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
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+260)))
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
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
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
	*(*int32)(unsafe.Add(mBase, uint32(v39)+136)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v39)+124)) = l0
	v395 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+120)) = v395
	*(*int64)(unsafe.Add(mBase, uint32(v39)+112)) = int64(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v423
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
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(902), int32(_a_F_CopyFrom_6))
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
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(908), int32(_a_F_CopyFrom_6))
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
	*(*int32)(unsafe.Add(mBase, uint32(v39)+152)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v39)+148)) = int32(556)
	v536 = int32(_a_F_CopyFrom_13)
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[5])) = v39 + int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+144)) = v537
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
	v621 = v619 >> (uint(int32(13)) % 32)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v624 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+216)))
	v625 = base.I32_extend16_s(v617)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	if v626&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v621)) == int32(0) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	if v1569&int32(1) != 0 {
		goto L359
	} else {
		goto L360
	}
L159:
	;
	v1569 = v1564
	v1570 = v1563
	v1571 = v1561
	v1572 = v1562
	v1575 = v1568
	v1577 = v1567
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
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
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
	v716 = int32(0)
	v721 = v706
	goto L190
L188:
	;
	v797 = v706
	goto L189
L189:
	;
	v824 = int32(2)
	v826 = v623 + v797<<(uint(v824)%32)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v622+v827<<(uint(v824)%32))))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v831)+24))
	v834 = m.T0[v833].(func(*base.Module, int32, int32, int32) int64)(m, v831, v532, v827+v614)
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
	v750 = v623 + v721<<(uint(v748)%32)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v622+v751<<(uint(v748)%32))))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v755)+24))
	v758 = m.T0[v757].(func(*base.Module, int32, int32, int32) int64)(m, v755, v532, v614+v751)
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
	*(*int64)(unsafe.Add(mBase, uint32(v626+v760<<(uint(int32(3))%32)))) = v758
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v622+v765<<(uint(int32(2))%32))))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v769)+24))
	v772 = m.T0[v771].(func(*base.Module, int32, int32, int32) int64)(m, v769, v532, v614+v765)
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
	*(*int64)(unsafe.Add(mBase, uint32(v626+v774<<(uint(int32(3))%32)))) = v772
	v779 = int32(2)
	v780 = v721 + v779
	v782 = v716 + v779
	if v782 != v703&int32(-2) {
		v716 = v782
		v721 = v780
		goto L190
	} else {
		goto L194
	}
L194:
	;
	goto L191
L195:
	;
	v797 = v780
	goto L189
L196:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v826)))
	*(*int64)(unsafe.Add(mBase, uint32(v626+v836<<(uint(int32(3))%32)))) = v834
	goto L183
L197:
	;
	v1561 = v601
	v1562 = v551
	v1563 = v552
	v1564 = v556
	v1567 = v577
	v1568 = int32(0)
	goto L159
L198:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
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
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v39)+116))
	v1176 = int32(0)
	if v517|base.B2i32(v1175 == v1176) == v1176 {
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
	v880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
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
	v887 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	v890 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[7]))
	if v890 == v884 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v931 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
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
	v934 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
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
	v943 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
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
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(1174), int32(_a_F_CopyFrom_6))
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
	*(*int32)(unsafe.Add(mBase, uint32(v601)+40)) = v964
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v46
	v968 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
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
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v972 == int32(0) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v532)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v976
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v972)+24))
	v981 = m.T0[v980].(func(*base.Module, int32, int32, int32) int64)(m, v972, v532, v39+int32(159))
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
	if v981 != int64(0) {
		goto L215
	} else {
		goto L219
	}
L219:
	;
	v989 = v581 + int64(1)
	v992 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[7]))
	if v992 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v581 = v989
	goto L139
L221:
	;
	goto L220
L222:
	;
	v996 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyFrom[8])))
	if v996&int32(1) == int32(0) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1001 = int32(_a_F_CopyFrom_15)
	v1003 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	v1004 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1003 + v1004
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	*(*int32)(unsafe.Add(mBase, uint32(v992))) = v1007 + v1004
	v1011 = int32(0)
	v1013 = int32(_a_F_CopyFrom_16)
	v1014 = base.AtomicRmwOr32(m, v1011, v1013, v1011)
	*(*int64)(unsafe.Add(mBase, uint32(v992+int32(24))+232)) = v989
	v1022 = base.AtomicRmwOr32(m, v1011, v1013, v1011)
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	*(*int32)(unsafe.Add(mBase, uint32(v992))) = v1023 + v1004
	v1029 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1029 - v1004
	goto L221
L224:
	;
	v1170 = int32(1)
	v1171 = F_ExecBRInsertTriggers(m, v41, v1164, v1163)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L285
	}
L225:
	;
	v1034 = F_ExecFindPartition(m, v305, v289, v359, v601, v41)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
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
	if v1034 != v577 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+52))
	if v1037 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	v1106 = v552
	v1107 = v556
	v1108 = v557
	v1109 = v577
	goto L231
L231:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v1110 != 0 {
		goto L259
	} else {
		goto L260
	}
L232:
	;
	v1046 = int32(1)
	if v1045&v1046|(v1044&v1046|base.B2i32(v516 != int32(2))) != 0 {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	v1040 = int32(0)
	v1044 = v1040
	v1045 = v1040
	goto L232
L234:
	;
	goto L235
L235:
	;
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037)+8)))
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037)+10)))
	v1044 = v1042
	v1045 = v1043
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
	v1082 = int32(0)
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v39)+116))
	if v515|base.B2i32(v1083 == v1082) != 0 {
		v1095 = v1082
		goto L236
	} else {
		goto L250
	}
L238:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+84))
	if v1054 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+104))
	if v1055 < int32(2) {
		goto L237
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1058 = int32(1)
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+208))
	if v1059 != 0 {
		v1095 = v1058
		goto L236
	} else {
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	v1062 = F_palloc(m, int32(_a_F_CopyFrom_10))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	v1064 = int32(0)
	base.MemoryFill(m, v1062, v1064, int32(4000))
	*(*int32)(unsafe.Add(mBase, uint32(v1062)+4000)) = v1034
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+84))
	if v1068 == v1064 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1071 = F_GetBulkInsertState(m)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L248
	}
L246:
	;
	v1073 = int32(0)
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1062)+4008)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1062)+4004)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+208)) = v1062
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v39)+112))
	v1079 = F_lappend(m, v1078, v1062)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L249
	}
L248:
	;
	v1073 = v1071
	goto L247
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v1079
	v1095 = v1058
	goto L236
L250:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(112), v1034, v39+int32(104))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	v1095 = v1082
	goto L236
L252:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	if v1096 != 0 {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	goto L254
L254:
	;
	v1106 = v1095
	v1107 = v1045
	v1108 = v1044
	v1109 = v1034
	goto L231
L255:
	;
	F_ReleaseBuffer(m, v1096)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
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
	if v1108&int32(1) != 0 {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	goto L261
L261:
	;
	v1116 = F_ExecGetRootToChildMap(m, v1034, v41)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L265
	}
L262:
	;
	v1114 = int32(0)
	goto L264
L263:
	;
	v1114 = v601
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+4)) = v1114
	goto L261
L265:
	;
	v1118 = int32(0)
	if base.B2i32(v516 != v1118)&v1106 == v1118 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+8))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1150)+40)) = v1153
	if v1108&int32(1) != 0 {
		v1163 = v1150
		v1164 = v1034
		v1165 = v1106
		v1166 = v1107
		v1169 = v1109
		goto L224
	} else {
		goto L283
	}
L267:
	;
	if v1116 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+208))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+4008))
	v1133 = v1129 + v1130<<(uint(int32(2))%32)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1133)))
	if v1134 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L270:
	;
	v1150 = v601
	goto L266
L271:
	;
	goto L272
L272:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+8))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+204))
	v1127 = F_execute_attr_map_slot(m, v1125, v601, v1126)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v1150 = v1127
	goto L266
L274:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+8))
	v1139 = F_table_slot_create(m, v1137, int32(0))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L1
	} else {
		goto L277
	}
L275:
	;
	v1142 = v1134
	goto L276
L276:
	;
	if v1116 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1133))) = v1139
	v1142 = v1139
	goto L276
L278:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+8))
	v1144 = F_execute_attr_map_slot(m, v1143, v601, v1142)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L1
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+8))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+32))
	m.T0[v1147].(func(*base.Module, int32, int32))(m, v1142, v601)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L282
	}
L281:
	;
	v1150 = v1144
	goto L266
L282:
	;
	v1150 = v1142
	goto L266
L283:
	;
	v1561 = v1150
	v1562 = v1034
	v1563 = v1106
	v1564 = v1107
	v1567 = v1109
	v1568 = int32(0)
	goto L159
L284:
	;
	v1163 = v601
	v1164 = v551
	v1165 = v552
	v1166 = v556
	v1169 = v577
	goto L224
L285:
	;
	if v1171 == int32(0) {
		v551 = v1164
		v552 = v1165
		v556 = v1166
		v557 = v1170
		v577 = v1169
		goto L139
	} else {
		goto L286
	}
L286:
	;
	v1569 = v1166
	v1570 = v1165
	v1571 = v1163
	v1572 = v1164
	v1575 = v1170
	v1577 = v1169
	goto L158
L287:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(112), int32(0), v39+int32(104))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L1
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v39)+144))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[5])) = v1189
	v1191 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if v1191 == int64(0) {
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
		goto L304
	} else {
		goto L305
	}
L292:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v1194 < int32(0) {
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	switch v1198 - int32(1) {
	case 0:
		goto L296
	case 1:
		goto L295
	default:
		goto L291
	}
L294:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), v1234, int32(_a_F_CopyFrom_6))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L1
	} else {
		goto L303
	}
L295:
	;
	v1219 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L1
	} else {
		goto L300
	}
L296:
	;
	v1203 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	if v1203 == int32(0) {
		goto L291
	} else {
		goto L298
	}
L298:
	;
	v1207 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+80)) = v1207
	F_errmsg_plural(m, int32(_a_F_CopyFrom_18), int32(_a_F_CopyFrom_19), base.I32_wrap_i64(v1207), v39+int32(80))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	v1234 = int32(1476)
	goto L294
L300:
	;
	if v1219 == int32(0) {
		goto L291
	} else {
		goto L301
	}
L301:
	;
	v1223 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v39)+96)) = v1223
	F_errmsg_plural(m, int32(_a_F_CopyFrom_20), int32(_a_F_CopyFrom_21), base.I32_wrap_i64(v1223), v39+int32(96))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	v1234 = int32(1482)
	goto L294
L303:
	;
	goto L291
L304:
	;
	F_FreeBulkInsertState(m, v514)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L1
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[0])) = v46
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_ExecASInsertTriggers(m, v41, v289, v1243)
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L1
	} else {
		goto L308
	}
L307:
	;
	goto L306
L308:
	;
	F_AfterTriggerEndQuery(m, v41)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v41)+104))
	F_ExecResetTupleTable(m, v1248, int32(0))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v289)+84))
	if v1252 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	if v517 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L312:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+80))
	if v1255 == int32(0) {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	m.T0[v1255].(func(*base.Module, int32, int32))(m, v41, v289)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	goto L311
L315:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v39)+112))
	if v1263 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L316:
	;
	goto L317
L317:
	;
	if v359 != 0 {
		goto L342
	} else {
		goto L343
	}
L318:
	;
	F_list_free(m, v1263)
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L1
	} else {
		goto L341
	}
L319:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1263)+4))
	if v1266 <= int32(0) {
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v39)+136))
	v1282 = int32(0)
	goto L321
L321:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1263)+12))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1307+v1282<<(uint(int32(2))%32))))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4000))
	v1313 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1312)+208)) = v1313
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+84))
	if v1315 == v1313 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	goto L318
L323:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4004))
	F_FreeBulkInsertState(m, v1318)
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L1
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1323 = int32(0)
	goto L327
L326:
	;
	goto L325
L327:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1311+v1323<<(uint(int32(2))%32))))
	if v1361 != 0 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+84))
	if v1369 != 0 {
		goto L334
	} else {
		goto L335
	}
L329:
	;
	F_ExecDropSingleTupleTableSlot(m, v1361)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	goto L328
L332:
	;
	v1365 = v1323 + int32(1)
	if v1365 != int32(1000) {
		v1323 = v1365
		goto L327
	} else {
		goto L333
	}
L333:
	;
	goto L331
L334:
	;
	F_pfree(m, v1311)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L1
	} else {
		goto L339
	}
L335:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1312)+8))
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1370)+188))
	if v1371 == int32(0) {
		goto L334
	} else {
		goto L336
	}
L336:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1371)+108))
	if v1374 == int32(0) {
		goto L334
	} else {
		goto L337
	}
L337:
	;
	m.T0[v1374].(func(*base.Module, int32, int32))(m, v1370, v1269)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	goto L334
L339:
	;
	v1384 = v1282 + int32(1)
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1263)+4))
	if v1384 < v1385 {
		v1282 = v1384
		goto L321
	} else {
		goto L340
	}
L340:
	;
	goto L322
L341:
	;
	goto L317
L342:
	;
	F_ExecCleanupTupleRouting(m, v305, v359)
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L1
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	F_ExecCloseResultRelations(m, v41)
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L1
	} else {
		goto L346
	}
L345:
	;
	goto L344
L346:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v1466 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1467 = int32(0)
	v1473 = v1466
	goto L350
L348:
	;
	goto L349
L349:
	;
	F_FreeExecutorState(m, v41)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L357
	}
L350:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1503+v1467<<(uint(int32(2))%32))))
	if v1507 != 0 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	goto L349
L352:
	;
	F_relation_close(m, v1507, int32(0))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L1
	} else {
		goto L355
	}
L353:
	;
	v1512 = v1473
	goto L354
L354:
	;
	v1514 = v1467 + int32(1)
	if base.Ui32(v1514) < base.Ui32(v1512) {
		v1467 = v1514
		v1473 = v1512
		goto L350
	} else {
		goto L356
	}
L355:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v1512 = v1511
	goto L354
L356:
	;
	goto L351
L357:
	;
	v1554 = *(*int64)(unsafe.Add(mBase, uint32(v39)+104))
	m.G0 = v39 + int32(160)
	return v1554
L358:
	;
	v1701 = *(*int64)(unsafe.Add(mBase, uint32(v39)+104))
	v1703 = v1701 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+104)) = v1703
	v1708 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[7]))
	if v1708 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L359:
	;
	v1580 = F_ExecIRInsertTriggers(m, v41, v1572, v1571)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L1
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+8))
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+52))
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+24))
	if v1584 == int32(0) {
		v1594 = v1582
		goto L363
	} else {
		goto L364
	}
L362:
	;
	goto L358
L363:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+84))
	if v1595 != 0 {
		v1603 = v1594
		goto L367
	} else {
		goto L368
	}
L364:
	;
	v1587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1584)+17)))
	if v1587 != int32(1) {
		v1594 = v1582
		goto L363
	} else {
		goto L365
	}
L365:
	;
	F_ExecComputeStoredGenerated(m, v1572, v41, v1571, int32(3))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L1
	} else {
		goto L366
	}
L366:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+8))
	v1594 = v1593
	goto L363
L367:
	;
	v1604 = int32(0)
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1603)+48))
	v1610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1609)+131)))
	if base.B2i32(base.B2i32(v359 == v1604)|v1575 == v1604)|base.B2i32(v1610 != int32(1)) == v1604 {
		goto L371
	} else {
		goto L372
	}
L368:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1594)+52))
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1596)+24))
	if v1597 == int32(0) {
		v1603 = v1594
		goto L367
	} else {
		goto L369
	}
L369:
	;
	F_ExecConstraints(m, v1572, v1571, v41)
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+8))
	v1603 = v1602
	goto L367
L371:
	;
	v1617 = F_ExecPartitionCheck(m, v1572, v1571, v41, int32(1))
	mBase = m.M
	v1618 = m.ExcPending
	if v1618 != 0 {
		goto L1
	} else {
		goto L374
	}
L372:
	;
	goto L373
L373:
	;
	v1619 = int32(1)
	if (base.B2i32(v516 == v1619)|v1570)&v1619 != 0 {
		goto L375
	} else {
		goto L376
	}
L374:
	;
	goto L373
L375:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+8))
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+28))
	m.T0[v1625].(func(*base.Module, int32))(m, v1571)
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L1
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+84))
	if v1661 != 0 {
		goto L385
	} else {
		goto L386
	}
L378:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+208))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1629)+4008))
	v1636 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v1629+v1630<<(uint(int32(3))%32)+int32(4016)))) = v1636
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1629)+4008))
	v1639 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+4008)) = v1638 + v1639
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v39)+120))
	v1643 = v1628 + v1642
	*(*int32)(unsafe.Add(mBase, uint32(v39)+120)) = v1643
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v39)+116))
	v1647 = v1645 + v1639
	*(*int32)(unsafe.Add(mBase, uint32(v39)+116)) = v1647
	v1649 = int32(0)
	if v1647 <= int32(999) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	if v1643 < int32(_a_F_CopyFrom_22) {
		v551 = v1572
		v552 = v1570
		v556 = v1649
		v557 = v1575
		v577 = v1577
		goto L139
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	F_CopyMultiInsertInfoFlush(m, v39+int32(112), v1572, v39+int32(104))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L1
	} else {
		goto L383
	}
L382:
	;
	goto L381
L383:
	;
	v551 = v1572
	v552 = v1570
	v556 = v1649
	v557 = v1575
	v577 = v1577
	goto L139
L384:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_ExecARInsertTriggers(m, v41, v1572, v1688, v1691, v1692)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L1
	} else {
		goto L393
	}
L385:
	;
	v1662 = int32(0)
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1661)+52))
	v1666 = m.T0[v1665].(func(*base.Module, int32, int32, int32, int32) int32)(m, v41, v1572, v1571, v1662)
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L1
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+8))
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1674)+188))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1675)+80))
	m.T0[v1676].(func(*base.Module, int32, int32, int32, int32, int32))(m, v1674, v1571, v48, v264, v514)
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L1
	} else {
		goto L390
	}
L388:
	;
	if v1666 == int32(0) {
		v551 = v1572
		v552 = v1662
		v556 = v1662
		v557 = v1575
		v577 = v1577
		goto L139
	} else {
		goto L389
	}
L389:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+8))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1670)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v1666)+40)) = v1671
	v1688 = v1666
	v1691 = int32(0)
	goto L384
L390:
	;
	v1679 = int32(0)
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1572)+12))
	if v1680 <= v1679 {
		v1688 = v1571
		v1691 = v1679
		goto L384
	} else {
		goto L391
	}
L391:
	;
	v1683 = int32(0)
	v1686 = F_ExecInsertIndexTuples(m, v1572, v41, v1683, v1571, v1683, v1683)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	v1688 = v1571
	v1691 = v1686
	goto L384
L393:
	;
	F_list_free(m, v1691)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	goto L358
L395:
	;
	v551 = v1572
	v552 = v1570
	v556 = v1569
	v557 = v1575
	v577 = v1577
	goto L139
L396:
	;
	goto L395
L397:
	;
	v1712 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyFrom[8])))
	if v1712&int32(1) == int32(0) {
		goto L396
	} else {
		goto L398
	}
L398:
	;
	v1717 = int32(_a_F_CopyFrom_15)
	v1719 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	v1720 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1719 + v1720
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1708)))
	*(*int32)(unsafe.Add(mBase, uint32(v1708))) = v1723 + v1720
	v1727 = int32(0)
	v1729 = int32(_a_F_CopyFrom_16)
	v1730 = base.AtomicRmwOr32(m, v1727, v1729, v1727)
	*(*int64)(unsafe.Add(mBase, uint32(v1708+int32(16))+232)) = v1703
	v1738 = base.AtomicRmwOr32(m, v1727, v1729, v1727)
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1708)))
	*(*int32)(unsafe.Add(mBase, uint32(v1708))) = v1739 + v1720
	v1745 = *(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyFrom[9])) = v1745 - v1720
	goto L396
L399:
	;
	F_errfinish(m, int32(_a_F_CopyFrom_5), int32(837), int32(_a_F_CopyFrom_6))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
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
