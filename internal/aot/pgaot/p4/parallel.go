package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecParallelHashJoin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 float64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v433 int32
	_ = v433
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v708 int32
	_ = v708
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v810 int32
	_ = v810
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v881 float64
	_ = v881
	var v885 int32
	_ = v885
	var v888 float64
	_ = v888
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v916 float64
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1081 float64
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1220 int32
	_ = v1220
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1436 int32
	_ = v1436
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1563 int32
	_ = v1563
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1597 int32
	_ = v1597
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	v23 = m.G0
	v25 = v23 - int32(48)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+128))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_MemoryContextReset(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v42 = v32 + int32(56)
	v53 = v27
	goto L3
L3:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v66 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	switch v69 - int32(1) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	case 5:
		goto L13
	default:
		goto L12
	}
L8:
	;
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+48)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(6)
	v53 = v74
	goto L3
L10:
	;
	m.G0 = v25 + int32(48)
	return v1648
L11:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1620)+72))
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1620)+16))
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+8))
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1623)+12))
	m.T0[v1624].(func(*base.Module, int32))(m, v1622)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L1
	} else {
		goto L313
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L1
	} else {
		goto L310
	}
L13:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+48))
	if int32(0) <= v1086 {
		goto L233
	} else {
		goto L234
	}
L14:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v925 = v920
	goto L203
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+168)))
	if v894 != 0 {
		goto L3
	} else {
		goto L196
	}
L16:
	;
	v680 = m.G0
	v682 = v680 - int32(16)
	m.G0 = v682
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v687 != 0 {
		goto L145
	} else {
		goto L146
	}
L17:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+44))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v296)+48))
	if v298 != 0 {
		goto L72
	} else {
		goto L73
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	v74 = F_ExecHashTableCreate(m, v31)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = v74
	v78 = F_MultiExecProcNode(m, v31)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v74)+64))
	if base.F64_ne(v80, float64(0)) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v74)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+56)) = v117
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+169)) = uint8(v119)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v121 == int32(3) {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v83 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if int32(3) < v84 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v1648 = int32(0)
	goto L10
L25:
	;
	goto L26
L26:
	;
	goto L27
L27:
	;
	v110 = int32(0)
	v112 = F_BarrierArriveAndWait(m, v42, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v1648 = v110
	goto L10
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v114 < int32(4) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v74)+44))
	if v124 < int32(2) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v292 != int32(5) {
		goto L9
	} else {
		goto L66
	}
L34:
	;
	v290 = F_BarrierArriveAndWait(m, v42, int32(134217748))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L65
	}
L35:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L36
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v129)+52))
	if v152 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v230 = int32(0)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v127)+44))
	if v231 <= v230 {
		goto L34
	} else {
		goto L60
	}
L38:
	;
	F_ExecReScan(m, v129)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v156 = m.T0[v155].(func(*base.Module, int32) int32)(m, v129)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	goto L37
L43:
	;
	if v156 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+4)))
	if v160&int32(2) != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v156
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	F_MemoryContextReset(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v167 = int32(4476144)
	v168 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v171
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)+20))
	v176 = m.T0[v175].(func(*base.Module, int32, int32, int32) int32)(m, v169, v128, v25+int32(43))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v176
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+43)))
	if v181 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v225 == int32(0) {
		goto L36
	} else {
		goto L58
	}
L49:
	;
	v184 = F_ExecFetchSlotMinimalTuple(m, v156, v25+int32(31))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v127)+44))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(32)))) = (v193 - int32(1)) & v186
	if base.Ui32(int32(2)) <= base.Ui32(v192) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v127)+144))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v25)+36))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v208+v209*int32(36))+32))
	F_sts_puttuple(m, v213, v25+int32(44), v184)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L55
	}
L52:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v206 = (v192 - int32(1)) & base.I32_rotr(v186, v203)
	goto L54
L53:
	;
	v206 = int32(0)
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(36)))) = v206
	goto L51
L55:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+31)))
	if v218 != int32(1) {
		goto L48
	} else {
		goto L56
	}
L56:
	;
	F_pfree(m, v184)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L48
L58:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L36
L60:
	;
	v235 = v230
	goto L61
L61:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v127)+144))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256+v235*int32(36))+32))
	F_sts_end_write(m, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L34
L63:
	;
	v264 = v235 + int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v127)+44))
	if v264 < v265 {
		v235 = v264
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L9
L66:
	;
	v1648 = int32(0)
	goto L10
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v413
	v537 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+168)) = uint8(v537)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v539
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(132)))) = (v545 - int32(1)) & v539
	if base.Ui32(int32(2)) <= base.Ui32(v544) {
		goto L118
	} else {
		goto L119
	}
L68:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v460 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L69:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+4)))
	if v433&int32(2) == int32(0) {
		goto L67
	} else {
		goto L99
	}
L70:
	;
	F_ExecForceStoreMinimalTuple(m, v370, v372, int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L97
	}
L71:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v296)+144))
	v403 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v399+v298*int32(36))+25)) = uint8(v403)
	goto L68
L72:
	;
	if v297 <= v298 {
		goto L71
	} else {
		goto L93
	}
L73:
	;
	if v297 != int32(1) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	if v301 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_ExecReScan(m, v28)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v305 = m.T0[v304].(func(*base.Module, int32) int32)(m, v28)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if v305 == int32(0) {
		goto L71
	} else {
		goto L80
	}
L80:
	;
	v311 = v305
	goto L81
L81:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+4)))
	if v331&int32(2) != 0 {
		goto L71
	} else {
		goto L83
	}
L82:
	;
	goto L71
L83:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v334)+12)) = v311
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v334)+20))
	F_MemoryContextReset(m, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v339 = int32(4476144)
	v340 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v334)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v343
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v341)+20))
	v348 = m.T0[v347].(func(*base.Module, int32, int32, int32) int32)(m, v341, v334, v25+int32(44))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v348
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+44)))
	if v353 == int32(0) {
		v413 = v311
		goto L69
	} else {
		goto L86
	}
L86:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	if v356 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_ExecReScan(m, v28)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v360 = m.T0[v359].(func(*base.Module, int32) int32)(m, v28)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	if v360 != 0 {
		v311 = v360
		goto L81
	} else {
		goto L92
	}
L92:
	;
	goto L82
L93:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v296)+144))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363+v298*int32(36))+32))
	v370 = F_sts_parallel_scan_next(m, v367, v25+int32(24))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v370 != 0 {
		goto L70
	} else {
		goto L95
	}
L95:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	m.T0[v374].(func(*base.Module, int32))(m, v372)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L71
L97:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v408 == int32(0) {
		goto L68
	} else {
		goto L98
	}
L98:
	;
	v413 = v408
	goto L69
L99:
	;
	goto L68
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(6)
	goto L3
L101:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+144))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v463)+48))
	v467 = v465 * int32(36)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v464+v467)))
	v472 = F_BarrierArriveAndDetachExceptLast(m, v469+int32(4))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L103
	}
L102:
	;
	if v525 == int32(0) {
		goto L100
	} else {
		goto L116
	}
L103:
	;
	if v472 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v463)+144))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v463)+48))
	v481 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v476+v477*int32(36))+26)) = uint8(v481)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v463)+144))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v483+v467)+28))
	F_sts_end_parallel_scan(m, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+61)))
	if v505 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L107:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v463)+144))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v488+v467)+32))
	F_sts_end_parallel_scan(m, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v469)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+48)) = int32(-1)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v463)+104))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v463)))
	v500 = v493 + v497<<(uint(int32(2))%32)
	if base.Ui32(v500) < base.Ui32(v496) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v502 = v496
	goto L111
L110:
	;
	v502 = v500
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v463)+104)) = v502
	v525 = int32(0)
	goto L102
L112:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v463)+144))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v463)+48))
	v513 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v508+v509*int32(36))+26)) = uint8(v513)
	F_ExecHashTableDetachBatch(m, v463)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = int64(0)
	v525 = int32(1)
	goto L102
L115:
	;
	v525 = int32(0)
	goto L102
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(5)
	goto L3
L117:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+24)))
	if v566 != int32(1) {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v558 = (v544 - int32(1)) & base.I32_rotr(v539, v555)
	goto L120
L119:
	;
	v558 = v537
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(20)))) = v558
	goto L117
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v607
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v53)+48))
	if v611 == v612 {
		goto L131
	} else {
		goto L132
	}
L122:
	;
	v607 = int32(-1)
	goto L121
L123:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v53)+32))
	v572 = v570 - int32(1)
	v573 = v560 & v572
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v569+v573<<(uint(int32(2))%32))))
	if v577 == int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v580 = v573
	v583 = v577
	goto L125
L125:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	if v560 == v586 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L122
L127:
	;
	v607 = v580
	goto L121
L128:
	;
	goto L129
L129:
	;
	v590 = (v580 + int32(1)) & v572
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v569+v590<<(uint(int32(2))%32))))
	if v594 != 0 {
		v580 = v590
		v583 = v594
		goto L125
	} else {
		goto L130
	}
L130:
	;
	goto L126
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(3)
	goto L16
L132:
	;
	if v607 != int32(-1) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v618 = F_ExecFetchSlotMinimalTuple(m, v413, v25+int32(36))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v53)+92))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v622
	v626 = v620 + v621<<(uint(int32(2))%32)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	if v627 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v630 = int32(4476144)
	v631 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v53)+124))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v633
	v636 = F_BufFileCreateTemp(m, int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	v641 = v627
	goto L137
L137:
	;
	F_BufFileWrite(m, v641, v25+int32(44), int32(4))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = v636
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v631
	v641 = v636
	goto L137
L139:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v618)))
	F_BufFileWrite(m, v641, v618, v648)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+36)))
	if v651 != int32(1) {
		goto L3
	} else {
		goto L141
	}
L141:
	;
	F_pfree(m, v618)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	goto L3
L143:
	;
	m.G0 = v682 + int32(16)
	if v810 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L144:
	;
	if v701 != 0 {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v685)+136))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v687)))
	v690 = F_dsa_get_address(m, v688, v689)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v685)+136))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v685)+20))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v693+v694<<(uint(int32(2))%32))))
	v699 = F_dsa_get_address(m, v692, v698)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	v701 = v690
	goto L144
L149:
	;
	v701 = v699
	goto L144
L150:
	;
	v708 = v701
	goto L153
L151:
	;
	goto L152
L152:
	;
	v810 = int32(0)
	goto L143
L153:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	if v724 != v684 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L152
L155:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v685)+136))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v708)))
	v763 = F_dsa_get_address(m, v761, v762)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L166
	}
L156:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v730 = F_ExecStoreMinimalTuple(m, v708+int32(8), v728, int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v730
	if v686 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v708
	v810 = int32(1)
	goto L143
L159:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_MemoryContextReset(m, v735)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v738 = int32(4476144)
	v739 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v741
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v686)+20))
	v746 = m.T0[v745].(func(*base.Module, int32, int32, int32) int32)(m, v686, v33, v682+int32(15))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	goto L158
L163:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v739
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_MemoryContextReset(m, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	if v746 == int32(0) {
		goto L155
	} else {
		goto L165
	}
L165:
	;
	goto L158
L166:
	;
	if v763 != 0 {
		v708 = v763
		goto L153
	} else {
		goto L167
	}
L167:
	;
	goto L154
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(4)
	goto L3
L169:
	;
	goto L170
L170:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v818 == int32(6) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v822 = int32(*(*int16)(unsafe.Add(mBase, uint32(v821)+18)))
	if v822 < int32(0) {
		goto L3
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	if v30 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L173
L175:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v885 == int32(0) {
		goto L3
	} else {
		goto L195
	}
L176:
	;
	v825 = int32(4476144)
	v826 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v828
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v833 = m.T0[v832].(func(*base.Module, int32, int32, int32) int32)(m, v30, v33, v25+int32(44))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v841 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+168)) = uint8(v841)
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v845 = v843 + int32(18)
	v846 = int32(*(*int16)(unsafe.Add(mBase, uint32(v845))))
	if int32(0) <= v846 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v826
	if v833 == int32(0) {
		goto L175
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v850 = v846 | int32(32768)
	*(*uint16)(unsafe.Add(mBase, uint32(v845))) = uint16(v850)
	goto L183
L182:
	;
	goto L183
L183:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v852 == int32(5) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	goto L3
L185:
	;
	goto L186
L186:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v857 == int32(1) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	goto L189
L188:
	;
	goto L189
L189:
	;
	if v852 == int32(7) {
		goto L3
	} else {
		goto L190
	}
L190:
	;
	if v29 == int32(0) {
		goto L11
	} else {
		goto L191
	}
L191:
	;
	v866 = int32(4476144)
	v867 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v869
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v874 = m.T0[v873].(func(*base.Module, int32, int32, int32) int32)(m, v29, v33, v25+int32(44))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v867
	if v874 != 0 {
		goto L11
	} else {
		goto L193
	}
L193:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v878 == int32(0) {
		goto L3
	} else {
		goto L194
	}
L194:
	;
	v881 = *(*float64)(unsafe.Add(mBase, uint32(v878)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v878)+248)) = base.F64_add(v881, float64(1))
	goto L3
L195:
	;
	v888 = *(*float64)(unsafe.Add(mBase, uint32(v885)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v885)+240)) = base.F64_add(v888, float64(1))
	goto L3
L196:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v895 == int32(0) {
		goto L3
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v895
	if v29 == int32(0) {
		goto L11
	} else {
		goto L198
	}
L198:
	;
	v901 = int32(4476144)
	v902 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v904
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v909 = m.T0[v908].(func(*base.Module, int32, int32, int32) int32)(m, v29, v33, v25+int32(44))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v902
	if v909 != 0 {
		goto L11
	} else {
		goto L200
	}
L200:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v913 == int32(0) {
		goto L3
	} else {
		goto L201
	}
L201:
	;
	v916 = *(*float64)(unsafe.Add(mBase, uint32(v913)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v913)+248)) = base.F64_add(v916, float64(1))
	goto L3
L202:
	;
	if v1057 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L203:
	;
	if v925 != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	if v964 != 0 {
		goto L212
	} else {
		goto L213
	}
L206:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v921)+136))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v925)))
	v946 = F_dsa_get_address(m, v944, v945)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v921)))
	if v950 <= v949 {
		v1057 = int32(0)
		goto L202
	} else {
		goto L210
	}
L209:
	;
	v964 = v946
	goto L205
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v949 + int32(1)
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v921)+136))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v921)+20))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v956+v949<<(uint(int32(2))%32))))
	v961 = F_dsa_get_address(m, v955, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v964 = v961
	goto L205
L212:
	;
	v968 = v964
	goto L215
L213:
	;
	goto L214
L214:
	;
	v1028 = int32(0)
	v1030 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v1030 == v1028 {
		v925 = v1028
		goto L203
	} else {
		goto L224
	}
L215:
	;
	v987 = int32(*(*int16)(unsafe.Add(mBase, uint32(v968)+18)))
	if int32(0) <= v987 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	goto L214
L217:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v994 = F_ExecStoreMinimalTuple(m, v968+int32(8), v992, int32(0))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v921)+136))
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v968)))
	v1004 = F_dsa_get_address(m, v1002, v1003)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v994
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_MemoryContextReset(m, v997)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v968
	v1057 = int32(1)
	goto L202
L222:
	;
	if v1004 != 0 {
		v968 = v1004
		goto L215
	} else {
		goto L223
	}
L223:
	;
	goto L216
L224:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v925 = v1028
	goto L203
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(6)
	goto L3
L227:
	;
	goto L228
L228:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v1062
	if v29 == int32(0) {
		goto L11
	} else {
		goto L229
	}
L229:
	;
	v1066 = int32(4476144)
	v1067 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1069
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v1074 = m.T0[v1073].(func(*base.Module, int32, int32, int32) int32)(m, v29, v33, v25+int32(44))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1067
	if v1074 != 0 {
		goto L11
	} else {
		goto L231
	}
L231:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1078 == int32(0) {
		goto L3
	} else {
		goto L232
	}
L232:
	;
	v1081 = *(*float64)(unsafe.Add(mBase, uint32(v1078)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v1078)+248)) = base.F64_add(v1081, float64(1))
	goto L3
L233:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+144))
	v1093 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1089+v1086*int32(36))+26)) = uint8(v1093)
	F_ExecHashTableDetachBatch(m, v1085)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+140))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v1097)+164)) = v1098 + int32(1)
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+44))
	v1103 = base.I32_rem_u_s(v1098, v1102)
	v1106 = v1103
	goto L237
L236:
	;
	goto L235
L237:
	;
	v1127 = v1106 * int32(36)
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+144))
	v1129 = v1127 + v1128
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+26)))
	if v1130 != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1648 = int32(0)
	goto L10
L239:
	;
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+44))
	v1581 = base.I32_rem_s(v1106+int32(1), v1580)
	if v1581 != v1103 {
		v1106 = v1581
		goto L237
	} else {
		goto L309
	}
L240:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1129)))
	v1133 = v1131 + int32(4)
	v1134 = F_BarrierAttach(m, v1133)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L1
	} else {
		goto L248
	}
L241:
	;
	F_ExecParallelHashTableSetCurrentBatch(m, v1085, v1106)
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L1
	} else {
		goto L307
	}
L242:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L1
	} else {
		goto L304
	}
L243:
	;
	F_BarrierDetach(m, v1133)
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L303
	}
L244:
	;
	F_ExecParallelHashTableSetCurrentBatch(m, v1085, v1106)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L1
	} else {
		goto L301
	}
L245:
	;
	F_ExecParallelHashTableSetCurrentBatch(m, v1085, v1106)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L1
	} else {
		goto L266
	}
L246:
	;
	v1316 = F_BarrierArriveAndWait(m, v1133, int32(134217742))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L1
	} else {
		goto L265
	}
L247:
	;
	v1137 = F_BarrierArriveAndWait(m, v1133, int32(134217743))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L249
	}
L248:
	;
	switch v1134 {
	case 0:
		goto L247
	case 1:
		goto L246
	case 2:
		goto L245
	case 3:
		goto L244
	case 4:
		goto L241
	case 5:
		goto L243
	default:
		goto L242
	}
L249:
	;
	if v1137 == int32(0) {
		goto L246
	} else {
		goto L250
	}
L250:
	;
	v1141 = int32(0)
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+144))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1142+v1106*int32(36))))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+136))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+140))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1148)+16))
	v1153 = F_dsa_allocate_extended(m, v1147, v1149<<(uint(int32(2))%32), v1141)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1146))) = v1153
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+136))
	v1157 = F_dsa_get_address(m, v1156, v1153)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	if v1149 <= int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	goto L246
L254:
	;
	v1162 = v1149 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v1149) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1174 = v1141
	v1177 = int32(0)
	goto L258
L256:
	;
	v1220 = v1141
	goto L257
L257:
	;
	if v1162 == int32(0) {
		goto L253
	} else {
		goto L261
	}
L258:
	;
	v1192 = v1157 + v1174<<(uint(int32(2))%32)
	v1193 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1192))) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+4)) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+8)) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+12)) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+16)) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+20)) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+24)) = v1193
	*(*int32)(unsafe.Add(mBase, uint32(v1192)+28)) = v1193
	v1209 = int32(8)
	v1210 = v1174 + v1209
	v1212 = v1177 + v1209
	if v1212 != v1149&int32(-8) {
		v1174 = v1210
		v1177 = v1212
		goto L258
	} else {
		goto L260
	}
L259:
	;
	v1220 = v1210
	goto L257
L260:
	;
	goto L259
L261:
	;
	v1242 = int32(0)
	v1245 = v1220
	goto L262
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1157+v1245<<(uint(int32(2))%32)))) = int32(0)
	v1266 = int32(1)
	v1269 = v1242 + v1266
	if v1269 != v1162 {
		v1242 = v1269
		v1245 = v1245 + v1266
		goto L262
	} else {
		goto L264
	}
L263:
	;
	goto L253
L264:
	;
	goto L263
L265:
	;
	goto L245
L266:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+144))
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1342+v1127)+28))
	F_sts_begin_parallel_scan(m, v1344)
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	v1349 = F_sts_parallel_scan_next(m, v1344, v25+int32(44))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	if v1349 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1354 = v1349
	goto L272
L270:
	;
	goto L271
L271:
	;
	F_sts_end_parallel_scan(m, v1344)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L1
	} else {
		goto L299
	}
L272:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	F_ExecForceStoreMinimalTuple(m, v1354, v1373, int32(0))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L1
	} else {
		goto L274
	}
L273:
	;
	goto L271
L274:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v1379 = m.G0
	v1381 = v1379 - int32(16)
	m.G0 = v1381
	v1385 = F_ExecFetchSlotMinimalTuple(m, v1377, v1381+int32(15))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1085)))
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1385)))
	v1389 = int32(8)
	v1393 = F_ExecParallelHashTupleAlloc(m, v1085, v1388+v1389, v1381+v1389)
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1393)+4)) = v1378
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1385)))
	if v1398 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1402 = v1393 + int32(18)
	v1403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1402))))
	v1405 = v1403 & int32(32767)
	*(*uint16)(unsafe.Add(mBase, uint32(v1402))) = uint16(v1405)
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1381)+8))
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+20))
	v1414 = v1408 + (v1387-int32(1))&v1378<<(uint(int32(2))%32)
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1414)))
	*(*int32)(unsafe.Add(mBase, uint32(v1393))) = v1415
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1414)))
	v1418 = base.B2i32(v1417 == v1415)
	if v1417 == v1415 {
		goto L281
	} else {
		goto L282
	}
L278:
	;
	v1399 = F__emscripten_memcpy_bulkmem(m, v1393+int32(8), v1385, v1398)
	mBase = m.M
	goto L280
L279:
	;
	goto L280
L280:
	;
	goto L277
L281:
	;
	v1419 = v1407
	goto L283
L282:
	;
	v1419 = v1417
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1414))) = v1419
	if v1418 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1436 = v1417
	goto L287
L285:
	;
	goto L286
L286:
	;
	v1475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381)+15)))
	if v1475 == int32(1) {
		goto L293
	} else {
		goto L294
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1393))) = v1436
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1414)))
	*(*int32)(unsafe.Add(mBase, uint32(v1393))) = v1446
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1414)))
	if v1448 == v1446 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	goto L286
L289:
	;
	v1450 = v1407
	goto L291
L290:
	;
	v1450 = v1448
	goto L291
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1414))) = v1450
	if v1448 != v1446 {
		v1436 = v1448
		goto L287
	} else {
		goto L292
	}
L292:
	;
	goto L288
L293:
	;
	F_pfree(m, v1385)
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L1
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	m.G0 = v1381 + int32(16)
	v1485 = F_sts_parallel_scan_next(m, v1344, v25+int32(44))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L1
	} else {
		goto L297
	}
L296:
	;
	goto L295
L297:
	;
	if v1485 != 0 {
		v1354 = v1485
		goto L272
	} else {
		goto L298
	}
L298:
	;
	goto L273
L299:
	;
	v1512 = F_BarrierArriveAndWait(m, v1133, int32(134217744))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	goto L244
L301:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+144))
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v1538+v1127)+32))
	F_sts_begin_parallel_scan(m, v1540)
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	goto L3
L303:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+144))
	v1549 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1547+v1127)+26)) = uint8(v1549)
	*(*int32)(unsafe.Add(mBase, uint32(v1085)+48)) = int32(-1)
	goto L239
L304:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v1557
	F_errmsg_internal(m, int32(470058), v25+int32(16))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(491075), int32(1392), int32(322466))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L307:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+144))
	v1573 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1571+v1127)+26)) = uint8(v1573)
	F_ExecHashTableDetachBatch(m, v1085)
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	goto L239
L309:
	;
	goto L238
L310:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v1588
	F_errmsg_internal(m, int32(478421), v25)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(491075), int32(672), int32(298034))
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	v1627 = int32(4476144)
	v1628 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1621)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1630
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1620)+24))
	v1636 = m.T0[v1635].(func(*base.Module, int32, int32, int32) int32)(m, v1620+int32(4), v1621, int32(0))
	mBase = m.M
	v1637 = m.ExcPending
	if v1637 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v1628
	v1640 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1622)+4)))
	v1642 = v1640 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v1622)+4)) = uint16(v1642)
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+12))
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1644)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1622)+6)) = uint16(v1645)
	v1648 = v1622
	goto L10
}
func F_ExecParallelInitializeWorker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
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
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int64
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	if l0 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v10 - int32(397) {
		case 0:
			v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+36)))
			if v140 != int32(1) {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			} else {
				v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v145 = int64(*(*int32)(unsafe.Add(mBase, uint32(v144)+40)))
				v147 = F_shm_toc_lookup(m, v143, v145, int32(0))
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(697)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v147
					v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v257 = m.ExcPending
					if v257 != 0 {
						return int32(0)
					} else {
						return v256
					}
				}
			}
		default:
			v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
			mBase = m.M
			v257 = m.ExcPending
			if v257 != 0 {
				return int32(0)
			} else {
				return v256
			}
		case 6:
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+36)))
			if v14 != int32(1) {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+40)))
				v21 = F_shm_toc_lookup(m, v17, v19, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
					v26 = F_table_beginscan_parallel(m, v25, v21)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v26
						v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							return v256
						}
					}
				}
			}
		case 8:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+36)))
			if v29|v31&int32(1) == int32(0) {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v38 = int64(*(*int32)(unsafe.Add(mBase, uint32(v30)+40)))
				v40 = F_shm_toc_lookup(m, v37, v38, int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					if v29 != 0 {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v40 + v42
					} else {
					}
					if v31&int32(1) == int32(0) {
						v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							return v256
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
						v55 = F_index_beginscan_parallel(m, v49, v50, l0+int32(168), v53, v54, v40)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v55
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
							if v58 != 0 {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
								if v59 != int32(1) {
									v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int32(0)
									} else {
										return v256
									}
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
									F_index_rescan(m, v55, v62, v63, v64, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return int32(0)
										} else {
											return v256
										}
									}
								}
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
								F_index_rescan(m, v55, v62, v63, v64, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int32(0)
									} else {
										return v256
									}
								}
							}
						}
					}
				}
			}
		case 9:
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+36)))
			if v70|v72&int32(1) == int32(0) {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v79 = int64(*(*int32)(unsafe.Add(mBase, uint32(v71)+40)))
				v81 = F_shm_toc_lookup(m, v78, v79, int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					if v70 != 0 {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v81)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v81 + v83
					} else {
					}
					if v72&int32(1) == int32(0) {
						v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							return v256
						}
					} else {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
						v96 = F_index_beginscan_parallel(m, v90, v91, l0+int32(160), v94, v95, v81)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v96
							v99 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v96)+28)) = uint8(v99)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
							if v101 != 0 {
								v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
								if v102 != int32(1) {
									v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int32(0)
									} else {
										return v256
									}
								} else {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
									v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
									F_index_rescan(m, v105, v106, v107, v108, v109)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return int32(0)
										} else {
											return v256
										}
									}
								}
							} else {
								v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
								F_index_rescan(m, v105, v106, v107, v108, v109)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int32(0)
									} else {
										return v256
									}
								}
							}
						}
					}
				}
			}
		case 10:
			v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v114 != 0 {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v117 = int64(*(*int32)(unsafe.Add(mBase, uint32(v116)+40)))
				v119 = F_shm_toc_lookup(m, v115, v117, int32(0))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v119
					v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v257 = m.ExcPending
					if v257 != 0 {
						return int32(0)
					} else {
						return v256
					}
				}
			} else {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			}
		case 11:
			v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+36)))
			if v170 != int32(1) {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			} else {
				v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v175 = int64(*(*int32)(unsafe.Add(mBase, uint32(v174)+40)))
				v177 = F_shm_toc_lookup(m, v173, v175, int32(0))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v177
					v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v180 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v177 + int32(24)
					} else {
					}
					v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v257 = m.ExcPending
					if v257 != 0 {
						return int32(0)
					} else {
						return v256
					}
				}
			}
		case 21:
			v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+36)))
			if v123 != int32(1) {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			} else {
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+156))
				if v127 != 0 {
					v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v130 = int64(*(*int32)(unsafe.Add(mBase, uint32(v129)+40)))
					v132 = F_shm_toc_lookup(m, v128, v130, int32(0))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v126)+156))
						m.T0[v135].(func(*base.Module, int32, int32, int32))(m, l0, v134, v132)
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								return v256
							}
						}
					}
				} else {
					v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v257 = m.ExcPending
					if v257 != 0 {
						return int32(0)
					} else {
						return v256
					}
				}
			}
		case 22:
			v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+36)))
			if v153 != int32(1) {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			} else {
				v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+40))
				if v157 != 0 {
					v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v160 = int64(*(*int32)(unsafe.Add(mBase, uint32(v159)+40)))
					v162 = F_shm_toc_lookup(m, v158, v160, int32(0))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v165 = *(*int32)(unsafe.Add(mBase, uint32(v156)+40))
						m.T0[v165].(func(*base.Module, int32, int32, int32))(m, l0, v164, v162)
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int32(0)
						} else {
							v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								return v256
							}
						}
					}
				} else {
					v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v257 = m.ExcPending
					if v257 != 0 {
						return int32(0)
					} else {
						return v256
					}
				}
			}
		case 26:
			v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+36)))
			if v185 != int32(1) {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			} else {
				v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v190 = int64(*(*int32)(unsafe.Add(mBase, uint32(v189)+40)))
				v192 = F_shm_toc_lookup(m, v188, v190, int32(0))
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return int32(0)
				} else {
					v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_SharedFileSetAttach(m, v192+int32(168), v196)
					mBase = m.M
					v198 = m.ExcPending
					if v198 != 0 {
						return int32(0)
					} else {
						v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v199)+128)) = v192
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(632)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(720)
						v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							return v256
						}
					}
				}
			}
		case 28:
			v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v247 = int64(*(*int32)(unsafe.Add(mBase, uint32(v246)+40)))
			v249 = F_shm_toc_lookup(m, v245, v247, int32(1))
			mBase = m.M
			v250 = m.ExcPending
			if v250 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v249
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			}
		case 29:
			v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v222 = int64(*(*int32)(unsafe.Add(mBase, uint32(v221)+40)))
			v224 = F_shm_toc_lookup(m, v220, v222, int32(1))
			mBase = m.M
			v225 = m.ExcPending
			if v225 != 0 {
				return int32(0)
			} else {
				v226 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)) = uint8(v226)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v224
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			}
		case 30:
			v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v231 = int64(*(*int32)(unsafe.Add(mBase, uint32(v230)+40)))
			v233 = F_shm_toc_lookup(m, v229, v231, int32(1))
			mBase = m.M
			v234 = m.ExcPending
			if v234 != 0 {
				return int32(0)
			} else {
				v235 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)) = uint8(v235)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v233
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			}
		case 32:
			v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v239)+40)))
			v242 = F_shm_toc_lookup(m, v238, v240, int32(1))
			mBase = m.M
			v243 = m.ExcPending
			if v243 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v242
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			}
		case 37:
			v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v205 != 0 {
				v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v208 = int64(*(*int32)(unsafe.Add(mBase, uint32(v207)+40)))
				v210 = F_shm_toc_lookup(m, v206, v208, int32(0))
				mBase = m.M
				v211 = m.ExcPending
				if v211 != 0 {
					return int32(0)
				} else {
					v213 = *(*int32)(unsafe.Add(mBase, _consts[172]))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v210 + v213*int32(20) + int32(4)
					v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v257 = m.ExcPending
					if v257 != 0 {
						return int32(0)
					} else {
						return v256
					}
				}
			} else {
				v256 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					return v256
				}
			}
		}
	}
}
func F_ExecParallelReportInstrumentation(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 float64
	_ = v84
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v101 float64
	_ = v101
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v129 float64
	_ = v129
	var v130 float64
	_ = v130
	var v133 int32
	_ = v133
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v192 int64
	_ = v192
	var v193 int64
	_ = v193
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v200 int32
	_ = v200
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v215 int64
	_ = v215
	var v216 int64
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_InstrEndLoop(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v19 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v63 = int32(416)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v70 = l1 + v59 + v61*v26*v63 + v67*v63
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)))
	if v75 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v26 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v26<<(uint(int32(2))%32))))
	if v34 == v13 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v37 = v26 + int32(1)
	if v37 != v19 {
		v26 = v37
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	F_errmsg_internal(m, int32(419240), v10)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(492034), int32(1312), int32(256402))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v220 = F_planstate_tree_walker_impl(m, l0, int32(631), l1)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L28
	}
L15:
	;
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v70)+16))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+16)) = v97 + v98
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v71)+32))
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v70)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+32)) = base.F64_add(v101, v102)
	v105 = *(*float64)(unsafe.Add(mBase, uint32(v71)+200))
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v70)+200))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+200)) = base.F64_add(v105, v106)
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v71)+208))
	v110 = *(*float64)(unsafe.Add(mBase, uint32(v70)+208))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+208)) = base.F64_add(v109, v110)
	v113 = *(*float64)(unsafe.Add(mBase, uint32(v71)+216))
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v70)+216))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+216)) = base.F64_add(v113, v114)
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v71)+224))
	v118 = *(*float64)(unsafe.Add(mBase, uint32(v70)+224))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+224)) = base.F64_add(v117, v118)
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v71)+232))
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v70)+232))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+232)) = base.F64_add(v121, v122)
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v71)+240))
	v126 = *(*float64)(unsafe.Add(mBase, uint32(v70)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+240)) = base.F64_add(v125, v126)
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v71)+248))
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v70)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+248)) = base.F64_add(v129, v130)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v133 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	if v74&int32(1) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v74&int32(1) == int32(0) {
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v82 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)) = uint8(v82)
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v71)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v70)+24)) = v84
	goto L15
L20:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v71)+24))
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v70)+24))
	if base.F64_lt(v90, v91) == int32(0) {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v70)+24)) = v90
	goto L15
L22:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v70)+256))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v71)+256))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+256)) = v136 + v137
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v70)+264))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v71)+264))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+264)) = v140 + v141
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v70)+272))
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v71)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+272)) = v144 + v145
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v70)+280))
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v71)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+280)) = v148 + v149
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v70)+288))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v71)+288))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+288)) = v152 + v153
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v70)+296))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v71)+296))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+296)) = v156 + v157
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v70)+304))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v71)+304))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+304)) = v160 + v161
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v70)+312))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v71)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+312)) = v164 + v165
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v70)+320))
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v71)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+320)) = v168 + v169
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v70)+328))
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v71)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+328)) = v172 + v173
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v70)+336))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v71)+336))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+336)) = v176 + v177
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v70)+344))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v71)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+344)) = v180 + v181
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v70)+352))
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v71)+352))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+352)) = v184 + v185
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v70)+360))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v71)+360))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+360)) = v188 + v189
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v70)+368))
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v71)+368))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+368)) = v192 + v193
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v70)+376))
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v71)+376))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+376)) = v196 + v197
	goto L24
L23:
	;
	goto L24
L24:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+2)))
	if v200 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v70)+400))
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v71)+400))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+400)) = v203 + v204
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v70)+384))
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v71)+384))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+384)) = v207 + v208
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v70)+392))
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v71)+392))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+392)) = v211 + v212
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v70)+408))
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v71)+408))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+408)) = v215 + v216
	goto L27
L26:
	;
	goto L27
L27:
	;
	goto L14
L28:
	;
	m.G0 = v10 + int32(16)
	return v220
}
func F_ParallelQueryMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v159 int32
	_ = v159
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
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int64
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v340 int64
	_ = v340
	var v344 int64
	_ = v344
	var v348 int64
	_ = v348
	var v352 int64
	_ = v352
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v359 int64
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int64
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v402 int64
	_ = v402
	var v404 int64
	_ = v404
	var v408 int64
	_ = v408
	var v410 int64
	_ = v410
	var v414 int64
	_ = v414
	var v416 int64
	_ = v416
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int64
	_ = v439
	var v441 int64
	_ = v441
	var v443 int64
	_ = v443
	var v445 int64
	_ = v445
	var v447 int64
	_ = v447
	var v449 int64
	_ = v449
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v25 = F_shm_toc_lookup(m, l1, int64(-2305843009213693951), v3)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = F_shm_toc_lookup(m, l1, int64(-2305843009213693947), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v35 = v29 + v32<<(uint(int32(16))%32)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	F_shm_mq_set_sender(m, v35, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = F_shm_mq_attach(m, v35, l0)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = F_CreateTupleQueueDestReceiver(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v47 = F_shm_toc_lookup(m, l1, int64(-2305843009213693946), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v47 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v50 = v49
	goto L10
L9:
	;
	v50 = int32(0)
	goto L10
L10:
	;
	v53 = F_shm_toc_lookup(m, l1, int64(-2305843009213693943), int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v57 = F_shm_toc_lookup(m, l1, int64(-2305843009213693944), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v61 = F_shm_toc_lookup(m, l1, int64(-2305843009213693950), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v63 = F_stringToNode(m, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v67 = F_shm_toc_lookup(m, l1, int64(-2305843009213693949), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v67
	v71 = v21 + int32(8)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 + int32(4)
	v81 = F_palloc(m, v73*int32(12)+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v83 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v81)+28)) = v73
	v86 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = int32(815)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = v81
	if v86 < v73 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v110 = v3
	goto L20
L18:
	;
	goto L19
L19:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[254]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	goto L24
L20:
	;
	v117 = v81 + int32(32) + v110*int32(12)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v122 = int32(4)
	v123 = v121 + v122
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v123
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	*(*uint16)(unsafe.Add(mBase, uint32(v117)+6)) = uint16(v125)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v127 + int32(2)
	v133 = F_datumRestore(m, v71, v117+v122)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v133
	v137 = v110 + int32(1)
	if v137 != v73 {
		v110 = v137
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v161 = int32(0)
	v163 = F_CreateQueryDesc(m, v63, v57, v160, v161, v42, v81, v161, v50)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	*(*int32)(unsafe.Add(mBase, _consts[57])) = v165
	F_pgstat_report_activity(m, int32(3), v165)
	mBase = m.M
	v171 = F_shm_toc_lookup(m, l1, int64(-2305843009213693945), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v173 = F_dsa_attach_in_place(m, v171, l0)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+32)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	F_ExecutorStart(m, v163, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v163)+44))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v182)+172)) = v173
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v184 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v163)+44))
	v259 = F_ExecParallelInitializeWorker(m, v256, v21+int32(8))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L37
	}
L30:
	;
	v187 = F_dsa_get_address(m, v173, v184)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v163)+40))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v187 + int32(4)
	if v190 <= int32(0) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v201 = int32(0)
	goto L33
L33:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v215 + int32(4)
	v220 = int32(8)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v189)+92))
	v225 = v222 + v216*int32(12)
	v228 = F_datumRestore(m, v21+v220, v225+v220)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L29
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v228
	v234 = v201 + int32(1)
	if v234 != v190 {
		v201 = v234
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v163)+44))
	v265 = v262
	goto L40
L38:
	;
	v337 = F___memcpy(m, int32(4374568), int32(4374408), int32(128))
	mBase = m.M
	v340 = *(*int64)(unsafe.Add(mBase, _consts[24]))
	*(*int64)(unsafe.Add(mBase, _consts[492])) = v340
	v344 = *(*int64)(unsafe.Add(mBase, _consts[21]))
	*(*int64)(unsafe.Add(mBase, _consts[493])) = v344
	v348 = *(*int64)(unsafe.Add(mBase, _consts[23]))
	*(*int64)(unsafe.Add(mBase, _consts[494])) = v348
	v352 = *(*int64)(unsafe.Add(mBase, _consts[22]))
	*(*int64)(unsafe.Add(mBase, _consts[495])) = v352
	goto L65
L39:
	;
	goto L38
L40:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	switch v267 - int32(394) {
	case 0:
		goto L46
	default:
		goto L39
	case 3:
		goto L49
	case 4:
		goto L50
	case 17:
		goto L45
	case 32:
		goto L48
	case 33:
		goto L47
	case 38, 39:
		goto L44
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v265)+120)) = v261
	v326 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+117)) = uint8(v326)
	goto L39
L42:
	;
	goto L41
L43:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v265+v322)))
	v265 = v324
	goto L40
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v265)+112)) = v261
	v322 = int32(36)
	goto L43
L45:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v265)+32))
	if v317 == int32(0) {
		v322 = int32(116)
		goto L43
	} else {
		goto L64
	}
L46:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v265)+36))
	if v315 != 0 {
		v265 = v315
		goto L40
	} else {
		goto L63
	}
L47:
	;
	if v261 < int64(0) {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	if int64(0) <= v261 {
		goto L42
	} else {
		goto L59
	}
L49:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v265)+108))
	if v287 <= int32(0) {
		goto L39
	} else {
		goto L55
	}
L50:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v265)+108))
	if v270 <= int32(0) {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	v276 = int32(0)
	goto L52
L52:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v265)+104))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277+v276<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v261, v281)
	mBase = m.M
	v284 = v276 + int32(1)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v265)+108))
	if v284 < v285 {
		v276 = v284
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L39
L54:
	;
	goto L53
L55:
	;
	v293 = int32(0)
	goto L56
L56:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v265)+104))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294+v293<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v261, v298)
	mBase = m.M
	v301 = v293 + int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v265)+108))
	if v301 < v302 {
		v293 = v301
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L39
L58:
	;
	goto L57
L59:
	;
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+117)) = uint8(v306)
	goto L38
L60:
	;
	v310 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+116)) = uint8(v310)
	goto L38
L61:
	;
	goto L62
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v265)+120)) = v261
	v313 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+116)) = uint8(v313)
	goto L38
L63:
	;
	goto L39
L64:
	;
	goto L39
L65:
	;
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v356 = int64(0)
	if v356 < v355 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v359 = v355
	goto L68
L67:
	;
	v359 = v356
	goto L68
L68:
	;
	F_ExecutorRun(m, v163, int32(1), v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_ExecutorFinish(m, v163)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v366 = F_shm_toc_lookup(m, l1, int64(-2305843009213693948), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v370 = F_shm_toc_lookup(m, l1, int64(-2305843009213693942), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v379 = v370 + v373<<(uint(int32(5))%32)
	v384 = F___memset(m, v366+v373<<(uint(int32(7))%32), int32(0), int32(128))
	mBase = m.M
	F_BufferUsageAccumDiff(m, v384, int32(4374568))
	mBase = m.M
	v388 = v379 + int32(24)
	v389 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v388))) = v389
	v392 = v379 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v392))) = v389
	v396 = v379 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v396))) = v389
	*(*int64)(unsafe.Add(mBase, uint32(v379))) = v389
	v402 = *(*int64)(unsafe.Add(mBase, _consts[21]))
	v404 = *(*int64)(unsafe.Add(mBase, _consts[493]))
	*(*int64)(unsafe.Add(mBase, uint32(v392))) = v402 - v404
	v408 = *(*int64)(unsafe.Add(mBase, _consts[22]))
	v410 = *(*int64)(unsafe.Add(mBase, _consts[495]))
	*(*int64)(unsafe.Add(mBase, uint32(v379))) = v408 - v410
	v414 = *(*int64)(unsafe.Add(mBase, _consts[23]))
	v416 = *(*int64)(unsafe.Add(mBase, _consts[494]))
	*(*int64)(unsafe.Add(mBase, uint32(v396))) = v414 - v416
	v420 = *(*int64)(unsafe.Add(mBase, _consts[24]))
	v422 = *(*int64)(unsafe.Add(mBase, _consts[492]))
	*(*int64)(unsafe.Add(mBase, uint32(v388))) = v420 - v422
	goto L73
L73:
	;
	if v47 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v163)+44))
	v426 = F_ExecParallelReportInstrumentation(m, v425, v47)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v163)+40))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+180))
	if v429 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	F_ExecutorEnd(m, v163)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	if v53 == int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v435 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	v438 = v53 + v435*int32(48)
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v429)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v438)+48)) = v439
	v441 = *(*int64)(unsafe.Add(mBase, uint32(v429)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v438)+40)) = v441
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v429)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v438)+32)) = v443
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v429)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v438)+24)) = v445
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v429)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v438)+16)) = v447
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v429)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v438)+8)) = v449
	goto L78
L81:
	;
	F_dsa_detach(m, v173)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_FreeQueryDesc(m, v163)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	m.T0[v458].(func(*base.Module, int32))(m, v42)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	m.G0 = v21 + int32(16)
	return
}
func F_ReinitializeParallelDSM(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	v6 = int32(4476144)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v10 = *(*int32)(unsafe.Add(mBase, _consts[144]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = F_shm_toc_lookup(m, v29, int64(-65535), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L8
	}
L2:
	;
	F_WaitForParallelWorkersToFinish(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	F_WaitForParallelWorkersToExit(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v21 == v19 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_pfree(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+60)) = int64(0)
	goto L1
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = int64(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v36 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v7
	return
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v42 = F_shm_toc_lookup(m, v39, int64(-65534), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v44 <= int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v49 = int32(0)
	goto L13
L13:
	;
	v55 = v42 + v49<<(uint(int32(14))%32)
	v57 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v57
	v63 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v55)+36)) = uint16(v63)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+24)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v55)+32)) = int32(16344)
	goto L15
L14:
	;
	goto L9
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	F_shm_mq_set_receiver(m, v55, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v77 = F_shm_mq_attach(m, v55, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v79+v49<<(uint(int32(3))%32))+4)) = v77
	v85 = v49 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v85 < v86 {
		v49 = v85
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
}
func F_parallel_vacuum_process_all_indexes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v318 int32
	_ = v318
	var v320 int64
	_ = v320
	var v321 int64
	_ = v321
	var v324 int32
	_ = v324
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v330 int32
	_ = v330
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v336 int32
	_ = v336
	var v338 int64
	_ = v338
	var v339 int64
	_ = v339
	var v342 int32
	_ = v342
	var v344 int64
	_ = v344
	var v345 int64
	_ = v345
	var v348 int32
	_ = v348
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v354 int32
	_ = v354
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v360 int32
	_ = v360
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v366 int32
	_ = v366
	var v368 int64
	_ = v368
	var v369 int64
	_ = v369
	var v372 int32
	_ = v372
	var v374 int64
	_ = v374
	var v375 int64
	_ = v375
	var v378 int32
	_ = v378
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v384 int32
	_ = v384
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v390 int32
	_ = v390
	var v392 int64
	_ = v392
	var v393 int64
	_ = v393
	var v396 int32
	_ = v396
	var v398 int64
	_ = v398
	var v399 int64
	_ = v399
	var v402 int32
	_ = v402
	var v404 int64
	_ = v404
	var v405 int64
	_ = v405
	var v408 int32
	_ = v408
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v414 int32
	_ = v414
	var v416 int64
	_ = v416
	var v417 int64
	_ = v417
	var v420 int32
	_ = v420
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v426 int32
	_ = v426
	var v428 int64
	_ = v428
	var v429 int64
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v23 = v21 - int32(1)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v23 < v25 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = int32(1)
	v21 = v14
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = int32(2)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if l1 != 0 {
		v20 = v15
		v21 = v16
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v20 = v15
	v21 = v17 + v16
	goto L1
L6:
	;
	v27 = v23
	goto L8
L7:
	;
	v27 = v25
	goto L8
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v28 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v35 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+44)) = v88
	if v88 < v27 {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v43 = v40 + v35*int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v20
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v35))))
	if v48 != int32(1) {
		v73 = int32(0)
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+4)) = uint8(v73)
	v76 = v35 + int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v76 < v77 {
		v35 = v76
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v35<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+204))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+29)))
	if l2 != 0 {
		v73 = v57 & int32(1)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v60 = int32(0)
	if v57&int32(6) == v60 {
		v73 = v60
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v67 = int32(0)
	v73 = base.B2i32(v57&int32(2) == v67) | base.B2i32(l1 <= v67)
	goto L14
L18:
	;
	goto L13
L19:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v559 {
		goto L115
	} else {
		goto L116
	}
L20:
	;
	if int32(0) < l1 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v438 != 0 {
		goto L87
	} else {
		goto L88
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ReinitializeParallelDSM(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v99 = *(*int32)(unsafe.Add(mBase, _consts[465]))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+36)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+40)) = int32(0)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	if v105 < v27 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	return
L27:
	;
	goto L25
L28:
	;
	v107 = v105
	goto L30
L29:
	;
	v107 = v27
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+16)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LaunchParallelWorkers(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	if int32(0) < v113 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[466])) = v117
	*(*int32)(unsafe.Add(mBase, _consts[465])) = v117
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, _consts[464])) = v123 + int32(40)
	*(*int32)(unsafe.Add(mBase, _consts[467])) = v123 + int32(36)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v135 = F_errstart(m, v133, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	if l2 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v177 != 0 {
		goto L52
	} else {
		goto L53
	}
L37:
	;
	F_errfinish(m, int32(492017), v171, int32(155955))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L26
	} else {
		goto L51
	}
L38:
	;
	if v135 == int32(0) {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v135 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v141
	if v141 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v148 = int32(660199)
	goto L44
L43:
	;
	v148 = int32(660129)
	goto L44
L44:
	;
	F_errmsg(m, v148, v11+int32(16))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	v171 = int32(712)
	goto L37
L46:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v157
	if v157 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v164 = int32(660062)
	goto L49
L48:
	;
	v164 = int32(659994)
	goto L49
L49:
	;
	F_errmsg(m, v164, v11+int32(32))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L26
	} else {
		goto L50
	}
L50:
	;
	v171 = int32(718)
	goto L37
L51:
	;
	goto L36
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v178 + int32(1)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v182 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v189 = int32(0)
	v190 = v182
	goto L58
L56:
	;
	goto L57
L57:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v222 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v197 = v194 + v189*int32(48)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+4)))
	if v198 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L57
L60:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v189<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v205, v197)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L26
	} else {
		goto L63
	}
L61:
	;
	v209 = v190
	goto L62
L62:
	;
	v211 = v189 + int32(1)
	if v211 < v209 {
		v189 = v211
		v190 = v209
		goto L58
	} else {
		goto L64
	}
L63:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v209 = v208
	goto L62
L64:
	;
	goto L59
L65:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v238)+44)) = v239 + int32(1)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v239 < v243 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v225 - int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v230 == int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v233 + int32(1)
	goto L65
L68:
	;
	v248 = v239
	goto L71
L69:
	;
	goto L70
L70:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v283 != 0 {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v256 = v253 + v248*int32(48)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+4)))
	if v257 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v260+v248<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v264, v256)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L26
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v267)+44)) = v268 + int32(1)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v268 < v272 {
		v248 = v268
		goto L71
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	goto L72
L78:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v284 - int32(1)
	goto L80
L79:
	;
	goto L80
L80:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_WaitForParallelWorkersToFinish(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L26
	} else {
		goto L81
	}
L81:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+20))
	if v292 <= int32(0) {
		goto L19
	} else {
		goto L82
	}
L82:
	;
	v299 = int32(0)
	goto L83
L83:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v307 = v304 + v299<<(uint(int32(7))%32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v311 = v308 + v299<<(uint(int32(5))%32)
	v312 = int32(4374408)
	v314 = *(*int64)(unsafe.Add(mBase, _consts[5]))
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v307)))
	*(*int64)(unsafe.Add(mBase, _consts[5])) = v314 + v315
	v318 = int32(4374416)
	v320 = *(*int64)(unsafe.Add(mBase, _consts[6]))
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v307)+8))
	*(*int64)(unsafe.Add(mBase, _consts[6])) = v320 + v321
	v324 = int32(4374424)
	v326 = *(*int64)(unsafe.Add(mBase, _consts[7]))
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v307)+16))
	*(*int64)(unsafe.Add(mBase, _consts[7])) = v326 + v327
	v330 = int32(4374432)
	v332 = *(*int64)(unsafe.Add(mBase, _consts[8]))
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v307)+24))
	*(*int64)(unsafe.Add(mBase, _consts[8])) = v332 + v333
	v336 = int32(4374440)
	v338 = *(*int64)(unsafe.Add(mBase, _consts[9]))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v307)+32))
	*(*int64)(unsafe.Add(mBase, _consts[9])) = v338 + v339
	v342 = int32(4374448)
	v344 = *(*int64)(unsafe.Add(mBase, _consts[10]))
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v307)+40))
	*(*int64)(unsafe.Add(mBase, _consts[10])) = v344 + v345
	v348 = int32(4374456)
	v350 = *(*int64)(unsafe.Add(mBase, _consts[11]))
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v307)+48))
	*(*int64)(unsafe.Add(mBase, _consts[11])) = v350 + v351
	v354 = int32(4374464)
	v356 = *(*int64)(unsafe.Add(mBase, _consts[12]))
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v307)+56))
	*(*int64)(unsafe.Add(mBase, _consts[12])) = v356 + v357
	v360 = int32(4374472)
	v362 = *(*int64)(unsafe.Add(mBase, _consts[13]))
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v307)+64))
	*(*int64)(unsafe.Add(mBase, _consts[13])) = v362 + v363
	v366 = int32(4374480)
	v368 = *(*int64)(unsafe.Add(mBase, _consts[14]))
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v307)+72))
	*(*int64)(unsafe.Add(mBase, _consts[14])) = v368 + v369
	v372 = int32(4374488)
	v374 = *(*int64)(unsafe.Add(mBase, _consts[15]))
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v307)+80))
	*(*int64)(unsafe.Add(mBase, _consts[15])) = v374 + v375
	v378 = int32(4374496)
	v380 = *(*int64)(unsafe.Add(mBase, _consts[16]))
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v307)+88))
	*(*int64)(unsafe.Add(mBase, _consts[16])) = v380 + v381
	v384 = int32(4374504)
	v386 = *(*int64)(unsafe.Add(mBase, _consts[17]))
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v307)+96))
	*(*int64)(unsafe.Add(mBase, _consts[17])) = v386 + v387
	v390 = int32(4374512)
	v392 = *(*int64)(unsafe.Add(mBase, _consts[18]))
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v307)+104))
	*(*int64)(unsafe.Add(mBase, _consts[18])) = v392 + v393
	v396 = int32(4374520)
	v398 = *(*int64)(unsafe.Add(mBase, _consts[19]))
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v307)+112))
	*(*int64)(unsafe.Add(mBase, _consts[19])) = v398 + v399
	v402 = int32(4374528)
	v404 = *(*int64)(unsafe.Add(mBase, _consts[20]))
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v307)+120))
	*(*int64)(unsafe.Add(mBase, _consts[20])) = v404 + v405
	v408 = int32(4374552)
	v410 = *(*int64)(unsafe.Add(mBase, _consts[21]))
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v311)+16))
	*(*int64)(unsafe.Add(mBase, _consts[21])) = v410 + v411
	v414 = int32(4374536)
	v416 = *(*int64)(unsafe.Add(mBase, _consts[22]))
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v311)))
	*(*int64)(unsafe.Add(mBase, _consts[22])) = v416 + v417
	v420 = int32(4374544)
	v422 = *(*int64)(unsafe.Add(mBase, _consts[23]))
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v311)+8))
	*(*int64)(unsafe.Add(mBase, _consts[23])) = v422 + v423
	v426 = int32(4374560)
	v428 = *(*int64)(unsafe.Add(mBase, _consts[24]))
	v429 = *(*int64)(unsafe.Add(mBase, uint32(v311)+24))
	*(*int64)(unsafe.Add(mBase, _consts[24])) = v428 + v429
	goto L85
L84:
	;
	goto L19
L85:
	;
	v433 = v299 + int32(1)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+20))
	if v433 < v435 {
		v299 = v433
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)))
	*(*int32)(unsafe.Add(mBase, uint32(v438))) = v439 + int32(1)
	goto L89
L88:
	;
	goto L89
L89:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v443 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v450 = int32(0)
	v451 = v443
	goto L93
L91:
	;
	goto L92
L92:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v483 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L93:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v458 = v455 + v450*int32(48)
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458)+4)))
	if v459 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L92
L95:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v462+v450<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v466, v458)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L26
	} else {
		goto L98
	}
L96:
	;
	v470 = v451
	goto L97
L97:
	;
	v472 = v450 + int32(1)
	if v472 < v470 {
		v450 = v472
		v451 = v470
		goto L93
	} else {
		goto L99
	}
L98:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v470 = v469
	goto L97
L99:
	;
	goto L94
L100:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v499)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v499)+44)) = v500 + int32(1)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v500 < v504 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = v486 - int32(1)
	v491 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v491 == int32(0) {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	*(*int32)(unsafe.Add(mBase, uint32(v491))) = v494 + int32(1)
	goto L100
L103:
	;
	v509 = v500
	goto L106
L104:
	;
	goto L105
L105:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _consts[464]))
	if v544 == int32(0) {
		goto L19
	} else {
		goto L113
	}
L106:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v517 = v514 + v509*int32(48)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+4)))
	if v518 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L105
L108:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v521+v509<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v525, v517)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L26
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v528)+44)) = v529 + int32(1)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v529 < v533 {
		v509 = v529
		goto L106
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	goto L107
L113:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v544)))
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = v547 - int32(1)
	goto L19
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L26
	} else {
		goto L125
	}
L115:
	;
	v566 = int32(0)
	goto L118
L116:
	;
	goto L117
L117:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _consts[467]))
	if v593 != 0 {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v574 = v571 + v566*int32(48)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	if v575 != int32(3) {
		goto L114
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v574))) = int32(0)
	v581 = v566 + int32(1)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v581 < v582 {
		v566 = v581
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v593)))
	*(*int32)(unsafe.Add(mBase, _consts[465])) = v595
	v598 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[467])) = v598
	*(*int32)(unsafe.Add(mBase, _consts[464])) = v598
	goto L124
L123:
	;
	goto L124
L124:
	;
	m.G0 = v11 + int32(48)
	return
L125:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v610+v566<<(uint(int32(2))%32))))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v615 + int32(4)
	F_errmsg_internal(m, int32(441808), v11)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L26
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(492017), int32(753), int32(155955))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L26
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
