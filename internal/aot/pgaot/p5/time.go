package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeTimeOnly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int64
	_ = v530
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 float64
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 float64
	_ = v620
	var v622 float64
	_ = v622
	var v626 int64
	_ = v626
	var v628 int64
	_ = v628
	var v631 int64
	_ = v631
	var v636 int64
	_ = v636
	var v638 int64
	_ = v638
	var v643 int64
	_ = v643
	var v645 int64
	_ = v645
	var v649 int64
	_ = v649
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v811 int32
	_ = v811
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v927 int32
	_ = v927
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v970 int32
	_ = v970
	var v988 int32
	_ = v988
	var v1003 int32
	_ = v1003
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1070 int32
	_ = v1070
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1128 int32
	_ = v1128
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1372 int32
	_ = v1372
	var v1382 int32
	_ = v1382
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1413 int32
	_ = v1413
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1433 int32
	_ = v1433
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1457 int32
	_ = v1457
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1490 int32
	_ = v1490
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	v9 = int32(0)
	v35 = m.G0
	v37 = v35 - int32(80)
	m.G0 = v37
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+59)) = uint8(v9)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(-1)
	if l6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v54 = l4 + int32(8)
	if l2 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v1658 + int32(80)
	return v1659
L5:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1229)+59)))
	if v1249 != 0 {
		goto L284
	} else {
		goto L285
	}
L6:
	;
	v1229 = v37
	v1232 = v9
	v1239 = int32(2)
	v1240 = v9
	v1247 = v9
	v1249 = v9
	v1251 = v9
	v1252 = v9
	goto L5
L7:
	;
	goto L8
L8:
	;
	v58 = int32(4)
	v60 = int32(2)
	v64 = l1 + l2<<(uint(v60)%32) - v58
	v66 = base.B2i32(l2 == int32(1))
	v79 = v9
	v81 = v9
	v85 = v9
	v86 = v60
	v87 = v9
	v94 = v9
	v96 = v9
	v98 = v9
	v99 = v9
	goto L9
L9:
	;
	v102 = int32(-1)
	v104 = v85 << (uint(int32(2)) % 32)
	v105 = l1 + v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	switch v106 {
	case 0:
		goto L16
	case 1, 6:
		goto L15
	case 2:
		goto L19
	case 3:
		goto L18
	case 4:
		goto L17
	default:
		v1658 = v37
		v1659 = v102
		goto L4
	}
L10:
	;
	if v1196 != 0 {
		v1658 = v37
		v1659 = int32(-1)
		goto L4
	} else {
		goto L282
	}
L11:
	;
	v1218 = v85 + int32(1)
	if v1218 != l2 {
		v79 = v1194
		v81 = v1196
		v85 = v1218
		v86 = v1201
		v87 = v1202
		v94 = v1209
		v96 = v1211
		v98 = v1213
		v99 = v1214
		goto L9
	} else {
		goto L281
	}
L12:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	if v1179&v79 != 0 {
		goto L278
	} else {
		goto L279
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(32)
	v1158 = v81
	v1163 = v86
	v1164 = v1128
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1128 = v276
	goto L13
L15:
	;
	v747 = l0 + v104
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v755 = F_DecodeTimezoneAbbrev(m, v85, v748, v37-int32(-64), v37+int32(60), v37+int32(52), l7)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L28
	} else {
		goto L204
	}
L16:
	;
	if v81 != 0 {
		goto L111
	} else {
		goto L112
	}
L17:
	;
	if l6 == int32(0) {
		v1658 = v37
		v1659 = v102
		goto L4
	} else {
		goto L83
	}
L18:
	;
	switch v81 {
	case 0, 3:
		goto L77
	default:
		v1658 = v37
		v1659 = v102
		goto L4
	}
L19:
	;
	if l6 == int32(0) {
		v1658 = v37
		v1659 = v102
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if l2 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v126 = l0 + v104
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if base.Ui32((v128-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	if v85 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v109 != int32(2) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v112 != int32(3) {
		goto L21
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v120 = F_DecodeDate(m, v115, v79, v37+int32(68), v37+int32(59), l4)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	return int32(0)
L29:
	;
	if v120 == int32(0) {
		v1158 = v81
		v1163 = v86
		v1164 = v87
		v1171 = v94
		v1173 = v96
		v1175 = v98
		v1176 = v99
		goto L12
	} else {
		goto L30
	}
L30:
	;
	v1658 = v37
	v1659 = v120
	goto L4
L31:
	;
	v135 = int32(31744)
	if v79&v135 == v135 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v276 = F_pg_tzset(m, v127)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L28
	} else {
		goto L75
	}
L34:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L35:
	;
	goto L36
L36:
	;
	v140 = int32(45)
	v141 = F___strchrnul(m, v127, v140)
	mBase = m.M
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v143 == v140 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v147 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v147 = v141
	goto L40
L39:
	;
	v147 = int32(0)
	goto L40
L40:
	;
	goto L37
L41:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L42:
	;
	goto L43
L43:
	;
	v151 = int32(0)
	v157 = m.G0
	v159 = v157 - int32(16)
	m.G0 = v159
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	switch v162 - int32(43) {
	case 0, 2:
		goto L46
	default:
		v249 = int32(-1)
		goto L45
	}
L44:
	;
	if v249 != 0 {
		v1658 = v37
		v1659 = v249
		goto L4
	} else {
		goto L70
	}
L45:
	;
	m.G0 = v159 + int32(16)
	goto L44
L46:
	;
	v165 = int32(4735052)
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v172 = F_strtoint(m, v147+int32(1), v159+int32(12))
	mBase = m.M
	v173 = int32(-5)
	v175 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v175 == int32(68) {
		v249 = v173
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v179 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if base.Ui32(int32(15)) < base.Ui32(v221) {
		v249 = v173
		goto L45
	} else {
		goto L61
	}
L49:
	;
	if v179 != int32(58) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v211 = F_strlen(m, v147)
	mBase = m.M
	if base.Ui32(v211) < base.Ui32(int32(4)) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v220 = int32(0)
	v221 = v172
	v223 = v151
	goto L48
L53:
	;
	goto L54
L54:
	;
	v183 = int32(4735052)
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v190 = F_strtoint(m, v178+int32(1), v159+int32(12))
	mBase = m.M
	v192 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v192 == int32(68) {
		v249 = v173
		goto L45
	} else {
		goto L55
	}
L55:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v196 != int32(58) {
		v220 = v190
		v221 = v172
		v223 = v151
		goto L48
	} else {
		goto L56
	}
L56:
	;
	v199 = int32(4735052)
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v206 = F_strtoint(m, v195+int32(1), v159+int32(12))
	mBase = m.M
	v208 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v208 != int32(68) {
		v220 = v190
		v221 = v172
		v223 = v206
		goto L48
	} else {
		goto L57
	}
L57:
	;
	v249 = v173
	goto L45
L58:
	;
	v220 = int32(0)
	v221 = v172
	v223 = v151
	goto L48
L59:
	;
	goto L60
L60:
	;
	v215 = int32(100)
	v216 = base.I32_div_s(v172, v215)
	v220 = v172 - v216*v215
	v221 = v216
	v223 = v151
	goto L48
L61:
	;
	if base.Ui32(int32(59)) < base.Ui32(v220) {
		v249 = v173
		goto L45
	} else {
		goto L62
	}
L62:
	;
	if base.Ui32(int32(59)) < base.Ui32(v223) {
		v249 = v173
		goto L45
	} else {
		goto L63
	}
L63:
	;
	v230 = int32(60)
	v235 = (v221*v230+v220)*v230 + v223
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v238 == int32(45) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v241 = v235
	goto L66
L65:
	;
	v241 = int32(0) - v235
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v241
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v246 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v247 = int32(-1)
	goto L69
L68:
	;
	v247 = int32(0)
	goto L69
L69:
	;
	v249 = v247
	goto L45
L70:
	;
	v256 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v256)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v259 = F_strlen(m, v258)
	mBase = m.M
	v266 = F_DecodeNumberField(m, v259, v258, v79|int32(14), v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L28
	} else {
		goto L71
	}
L71:
	;
	if v266 < int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v266
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v272 | int32(32)
	v1158 = v81
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L75:
	;
	if v276 != 0 {
		goto L14
	} else {
		goto L76
	}
L76:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v278
	v1658 = v37
	v1659 = int32(-6)
	goto L4
L77:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v288 = F_DecodeTimeCommon(m, v282, int32(32767), v37+int32(68), v37+int32(8))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L28
	} else {
		goto L78
	}
L78:
	;
	if v288 != 0 {
		v1658 = v37
		v1659 = v288
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v37)+24))
	if int64(2147483648) <= v290 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v1658 = v37
	v1659 = int32(-2)
	goto L4
L81:
	;
	goto L82
L82:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l4)+8)) = uint32(v290)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v295
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v297
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v299
	v1158 = int32(0)
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L83:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v308 = int32(0)
	v314 = m.G0
	v316 = v314 - int32(16)
	m.G0 = v316
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	switch v319 - int32(43) {
	case 0, 2:
		goto L86
	default:
		v406 = int32(-1)
		goto L85
	}
L84:
	;
	if v406 != 0 {
		v1658 = v37
		v1659 = v406
		goto L4
	} else {
		goto L110
	}
L85:
	;
	m.G0 = v316 + int32(16)
	goto L84
L86:
	;
	v322 = int32(4735052)
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v329 = F_strtoint(m, v305+int32(1), v316+int32(12))
	mBase = m.M
	v330 = int32(-5)
	v332 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v332 == int32(68) {
		v406 = v330
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v316)+12))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	if v336 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if base.Ui32(int32(15)) < base.Ui32(v378) {
		v406 = v330
		goto L85
	} else {
		goto L101
	}
L89:
	;
	if v336 != int32(58) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v368 = F_strlen(m, v305)
	mBase = m.M
	if base.Ui32(v368) < base.Ui32(int32(4)) {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	v377 = int32(0)
	v378 = v329
	v380 = v308
	goto L88
L93:
	;
	goto L94
L94:
	;
	v340 = int32(4735052)
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v347 = F_strtoint(m, v335+int32(1), v316+int32(12))
	mBase = m.M
	v349 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v349 == int32(68) {
		v406 = v330
		goto L85
	} else {
		goto L95
	}
L95:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v316)+12))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	if v353 != int32(58) {
		v377 = v347
		v378 = v329
		v380 = v308
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v356 = int32(4735052)
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v363 = F_strtoint(m, v352+int32(1), v316+int32(12))
	mBase = m.M
	v365 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v365 != int32(68) {
		v377 = v347
		v378 = v329
		v380 = v363
		goto L88
	} else {
		goto L97
	}
L97:
	;
	v406 = v330
	goto L85
L98:
	;
	v377 = int32(0)
	v378 = v329
	v380 = v308
	goto L88
L99:
	;
	goto L100
L100:
	;
	v372 = int32(100)
	v373 = base.I32_div_s(v329, v372)
	v377 = v329 - v373*v372
	v378 = v373
	v380 = v308
	goto L88
L101:
	;
	if base.Ui32(int32(59)) < base.Ui32(v377) {
		v406 = v330
		goto L85
	} else {
		goto L102
	}
L102:
	;
	if base.Ui32(int32(59)) < base.Ui32(v380) {
		v406 = v330
		goto L85
	} else {
		goto L103
	}
L103:
	;
	v387 = int32(60)
	v392 = (v378*v387+v377)*v387 + v380
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if v395 == int32(45) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v398 = v392
	goto L106
L105:
	;
	v398 = int32(0) - v392
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+int32(8)))) = v398
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v316)+12))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	if v403 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v404 = int32(-1)
	goto L109
L108:
	;
	v404 = int32(0)
	goto L109
L109:
	;
	v406 = v404
	goto L85
L110:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v413
	v1128 = v87
	goto L13
L111:
	;
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v418 = l0 + v104
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v423 = F_strtol(m, v419, v37+int32(72), int32(10))
	mBase = m.M
	goto L114
L112:
	;
	goto L113
L113:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v683 = F_strlen(m, v682)
	mBase = m.M
	v684 = int32(46)
	v685 = F___strchrnul(m, v682, v684)
	mBase = m.M
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685))))
	if v687 == v684 {
		goto L176
	} else {
		goto L177
	}
L114:
	;
	v424 = int32(-2)
	v426 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v426 == int32(68) {
		v1658 = v37
		v1659 = v424
		goto L4
	} else {
		goto L115
	}
L115:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	if v430 == int32(46) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if v81 != int32(3) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	if v430 == int32(0) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	v1158 = int32(0)
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v676
	v1175 = v98
	v1176 = v99
	goto L12
L120:
	;
	if v81 != int32(31) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v656 = F_strlen(m, v655)
	mBase = m.M
	v663 = F_DecodeNumberField(m, v656, v655, v79|int32(14), v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L28
	} else {
		goto L170
	}
L123:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L124:
	;
	goto L125
L125:
	;
	if l6 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L127:
	;
	goto L128
L128:
	;
	if v423 < int32(0) {
		v1658 = v37
		v1659 = v424
		goto L4
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(14)
	v449 = v423 + int32(32044)
	v450 = int32(146097)
	v451 = base.I32_div_u_s(v449, v450)
	v452 = int32(3)
	v458 = int32(2)
	v463 = base.I32_div_u_s((v451*int32(1073595727)+v449)<<(uint(v458)%32)|v452, v450)
	v466 = v423 + v451*v452 + v463 + int32(32104)
	v467 = int32(1461)
	v468 = base.I32_div_u_s(v466, v467)
	v471 = v468*int32(-1461) + v466
	v473 = v471 << (uint(v458) % 32)
	if base.Ui32(v467) <= base.Ui32(v473) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v486 = base.I32_div_u_s(v473, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v486 + v468<<(uint(int32(2))%32) - int32(4800)
	v493 = int32(1)
	v495 = v484 + int32(123)
	v499 = int32(base.Ui32(v495*int32(2141)) >> (uint(int32(16)) % 32))
	v503 = base.I32_rem_u_s(v499+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v503 + v493
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v495 - int32(base.Ui32(v499*int32(7834))>>(uint(int32(8))%32))
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	if v513 != int32(46) {
		v676 = v493
		goto L119
	} else {
		goto L134
	}
L131:
	;
	v479 = base.I32_rem_u_s(v471+int32(305), int32(365))
	v484 = v479
	goto L130
L132:
	;
	goto L133
L133:
	;
	v483 = base.I32_rem_u_s(v471+int32(306), int32(366))
	v484 = v483
	goto L130
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v429
	v518 = v429 + int32(1)
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	if v519 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L136:
	;
	v622 = base.F64_mul(v620, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v622), float64(9.223372036854776e+18)) != 0 {
		goto L166
	} else {
		goto L167
	}
L137:
	;
	v620 = float64(0)
	goto L136
L138:
	;
	goto L139
L139:
	;
	v523 = int32(578958)
	v527 = m.G0
	v529 = v527 - int32(32)
	v530 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v529)+24)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v529)+16)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v529)+8)) = v530
	*(*int64)(unsafe.Add(mBase, uint32(v529))) = v530
	v538 = int32(*(*uint8)(unsafe.Add(mBase, _consts[179])))
	if v538 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v607 = F_strlen(m, v518)
	mBase = m.M
	if v606 != v607 {
		goto L135
	} else {
		goto L161
	}
L141:
	;
	v606 = int32(0)
	goto L140
L142:
	;
	goto L143
L143:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v542 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v546 = v518
	goto L147
L145:
	;
	goto L146
L146:
	;
	v556 = v523
	v557 = v538
	goto L150
L147:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546))))
	if v552 == v538 {
		v546 = v546 + int32(1)
		goto L147
	} else {
		goto L149
	}
L148:
	;
	v606 = v546 - v518
	goto L140
L149:
	;
	goto L148
L150:
	;
	v564 = v529 + int32(base.Ui32(v557)>>(uint(int32(3))%32))&int32(28)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v566 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v565 | v566<<(uint(v557)%32)
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)))
	if v570 != 0 {
		v556 = v556 + v566
		v557 = v570
		goto L150
	} else {
		goto L152
	}
L151:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	if v573 == int32(0) {
		v598 = v518
		goto L153
	} else {
		goto L154
	}
L152:
	;
	goto L151
L153:
	;
	v606 = v598 - v518
	goto L140
L154:
	;
	v577 = v518
	v578 = v573
	goto L155
L155:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v529+int32(base.Ui32(v578)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v586)>>(uint(v578)%32))&int32(1) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v598 = v594
	goto L153
L157:
	;
	v598 = v577
	goto L153
L158:
	;
	goto L159
L159:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577)+1)))
	v594 = v577 + int32(1)
	if v592 != 0 {
		v577 = v594
		v578 = v592
		goto L155
	} else {
		goto L160
	}
L160:
	;
	goto L156
L161:
	;
	*(*int32)(unsafe.Add(mBase, _consts[85])) = int32(0)
	v614 = F_strtod(m, v429, v37+int32(8))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L28
	} else {
		goto L162
	}
L162:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616))))
	if v617 != 0 {
		goto L135
	} else {
		goto L163
	}
L163:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v619 != 0 {
		goto L135
	} else {
		goto L164
	}
L164:
	;
	v620 = v614
	goto L136
L165:
	;
	v631 = base.I64_div_s(v628, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v54))) = uint32(v631)
	v636 = base.I64_extend32_s(v631)*int64(-3600000000) + v628
	v638 = base.I64_div_s(v636, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4+v58))) = uint32(v638)
	v643 = base.I64_extend32_s(v638)*int64(-60000000) + v636
	v645 = base.I64_div_s(v643, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4))) = uint32(v645)
	v649 = v645*int64(4293967296) + v643
	*(*uint32)(unsafe.Add(mBase, uint32(l5))) = uint32(v649)
	goto L169
L166:
	;
	v626 = base.I64_trunc_f64_s(v622)
	v628 = v626
	goto L165
L167:
	;
	goto L168
L168:
	;
	v628 = int64(-9223372036854775807 - 1)
	goto L165
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(31758)
	v676 = v493
	goto L119
L170:
	;
	if v663 < int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v663
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	if v670 != int32(31744) {
		v1658 = v37
		v1659 = int32(-1)
		goto L4
	} else {
		goto L174
	}
L174:
	;
	v676 = v96
	goto L119
L175:
	;
	if v691 != 0 {
		goto L179
	} else {
		goto L180
	}
L176:
	;
	v691 = v685
	goto L178
L177:
	;
	v691 = int32(0)
	goto L178
L178:
	;
	goto L175
L179:
	;
	if l2 == int32(1) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	goto L181
L181:
	;
	v723 = v79 | int32(14)
	if int32(5) <= v683 {
		goto L195
	} else {
		goto L196
	}
L182:
	;
	v704 = F_strlen(m, v691)
	mBase = m.M
	if base.Ui32(v683-v704) < base.Ui32(int32(3)) {
		goto L188
	} else {
		goto L189
	}
L183:
	;
	if v85 != 0 {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v692 != int32(2) {
		goto L182
	} else {
		goto L185
	}
L185:
	;
	v700 = F_DecodeDate(m, v682, v79, v37+int32(68), v37+int32(59), l4)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L28
	} else {
		goto L186
	}
L186:
	;
	if v700 == int32(0) {
		v1158 = int32(0)
		v1163 = v86
		v1164 = v87
		v1171 = v94
		v1173 = v96
		v1175 = v98
		v1176 = v99
		goto L12
	} else {
		goto L187
	}
L187:
	;
	v1658 = v37
	v1659 = v700
	goto L4
L188:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L189:
	;
	goto L190
L190:
	;
	v715 = F_DecodeNumberField(m, v683, v682, v79|int32(14), v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L28
	} else {
		goto L191
	}
L191:
	;
	if v715 < int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v715
	v1158 = int32(0)
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L195:
	;
	v730 = F_DecodeNumberField(m, v683, v682, v723, v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L28
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v737 = int32(0)
	v743 = F_DecodeNumber(m, v683, v682, v737, v723, v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L28
	} else {
		goto L202
	}
L198:
	;
	if v730 < int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L200:
	;
	goto L201
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v730
	v1158 = int32(0)
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L202:
	;
	if v743 == int32(0) {
		v1158 = v737
		v1163 = v86
		v1164 = v87
		v1171 = v94
		v1173 = v96
		v1175 = v98
		v1176 = v99
		goto L12
	} else {
		goto L203
	}
L203:
	;
	v1658 = v37
	v1659 = v743
	goto L4
L204:
	;
	if v755 != 0 {
		v1658 = v37
		v1659 = v755
		goto L4
	} else {
		goto L205
	}
L205:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v37)+64))
	if v757 == int32(31) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[1029])))
	if v763 != 0 {
		goto L211
	} else {
		goto L212
	}
L207:
	;
	v1003 = v757
	goto L208
L208:
	;
	if v1003 == int32(8) {
		v1194 = v79
		v1196 = v81
		v1201 = v86
		v1202 = v87
		v1209 = v94
		v1211 = v96
		v1213 = v98
		v1214 = v99
		goto L11
	} else {
		goto L256
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+64)) = v988
	*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = v970
	v1003 = v988
	goto L208
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[1029]))) = v927
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v927)+12))
	v953 = int32(*(*int8)(unsafe.Add(mBase, uint32(v927)+11)))
	v970 = v952
	v988 = v953
	goto L209
L211:
	;
	goto L216
L212:
	;
	goto L213
L213:
	;
	v811 = int32(*(*int8)(unsafe.Add(mBase, uint32(v760))))
	v823 = int32(1695264)
	v826 = int32(1696400)
	goto L229
L214:
	;
	if v800-v801 == int32(0) {
		v927 = v763
		goto L210
	} else {
		goto L228
	}
L216:
	;
	goto L217
L217:
	;
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760))))
	if v770 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v771 = v760
	v772 = v763
	v773 = int32(10)
	v774 = v770
	goto L222
L219:
	;
	v796 = v763
	v800 = int32(0)
	goto L220
L220:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796))))
	goto L214
L221:
	;
	v796 = v791
	v800 = v793
	goto L220
L222:
	;
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772))))
	if v774 != v776 {
		v791 = v772
		v793 = v774
		goto L221
	} else {
		goto L224
	}
L223:
	;
	v791 = v785
	v793 = int32(0)
	goto L221
L224:
	;
	if v776 == int32(0) {
		v791 = v772
		v793 = v774
		goto L221
	} else {
		goto L225
	}
L225:
	;
	v781 = v773 - int32(1)
	if v781 == int32(0) {
		v791 = v772
		v793 = v774
		goto L221
	} else {
		goto L226
	}
L226:
	;
	v784 = int32(1)
	v785 = v772 + v784
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771)+1)))
	if v786 != 0 {
		v771 = v771 + v784
		v772 = v785
		v773 = v781
		v774 = v786
		goto L222
	} else {
		goto L227
	}
L227:
	;
	goto L223
L228:
	;
	goto L213
L229:
	;
	v853 = v823 + (v826-v823)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v854 = int32(*(*int8)(unsafe.Add(mBase, uint32(v853))))
	v855 = v811 - v854
	if v855 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	v970 = v906
	v988 = int32(31)
	goto L209
L231:
	;
	goto L236
L232:
	;
	v905 = v855
	goto L233
L233:
	;
	v906 = int32(0)
	v910 = base.B2i32(v905 < v906)
	if v905 < v906 {
		goto L249
	} else {
		goto L250
	}
L234:
	;
	if v896 == int32(0) {
		v927 = v853
		goto L210
	} else {
		goto L248
	}
L236:
	;
	goto L237
L237:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760))))
	if v864 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v865 = v760
	v866 = v853
	v867 = int32(10)
	v868 = v864
	goto L242
L239:
	;
	v890 = v853
	v894 = int32(0)
	goto L240
L240:
	;
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v890))))
	v896 = v894 - v895
	goto L234
L241:
	;
	v890 = v885
	v894 = v887
	goto L240
L242:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866))))
	if v868 != v870 {
		v885 = v866
		v887 = v868
		goto L241
	} else {
		goto L244
	}
L243:
	;
	v885 = v879
	v887 = int32(0)
	goto L241
L244:
	;
	if v870 == int32(0) {
		v885 = v866
		v887 = v868
		goto L241
	} else {
		goto L245
	}
L245:
	;
	v875 = v867 - int32(1)
	if v875 == int32(0) {
		v885 = v866
		v887 = v868
		goto L241
	} else {
		goto L246
	}
L246:
	;
	v878 = int32(1)
	v879 = v866 + v878
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v865)+1)))
	if v880 != 0 {
		v865 = v865 + v878
		v866 = v879
		v867 = v875
		v868 = v880
		goto L242
	} else {
		goto L247
	}
L247:
	;
	goto L243
L248:
	;
	v905 = v896
	goto L233
L249:
	;
	v911 = v853 - int32(16)
	goto L251
L250:
	;
	v911 = v826
	goto L251
L251:
	;
	if v905 < v906 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v914 = v823
	goto L254
L253:
	;
	v914 = v853 + int32(16)
	goto L254
L254:
	;
	if base.Ui32(v914) <= base.Ui32(v911) {
		v823 = v914
		v826 = v911
		goto L229
	} else {
		goto L255
	}
L255:
	;
	goto L230
L256:
	;
	v1028 = int32(1) << (uint(v1003) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1028
	v1030 = int32(-1)
	switch v1003 {
	case 0:
		goto L266
	default:
		v1658 = v37
		v1659 = v1030
		goto L4
	case 5:
		goto L263
	case 6:
		goto L264
	case 7:
		goto L262
	case 9:
		goto L261
	case 17:
		goto L259
	case 18:
		goto L260
	case 23:
		goto L258
	case 28:
		goto L265
	case 31:
		goto L257
	}
L257:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v1105 = F_pg_tzset(m, v1104)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L28
	} else {
		goto L276
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(0)
	if v81 != 0 {
		v1658 = v37
		v1659 = v1030
		goto L4
	} else {
		goto L275
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(0)
	if v81 != 0 {
		v1658 = v37
		v1659 = v1030
		goto L4
	} else {
		goto L274
	}
L260:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1158 = v81
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = base.B2i32(v1095 == int32(1))
	goto L12
L261:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1158 = v81
	v1163 = v1094
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1028 | int32(32)
	if l6 == int32(0) {
		v1658 = v37
		v1659 = v1030
		goto L4
	} else {
		goto L273
	}
L263:
	;
	v1075 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1075
	if l6 == v1075 {
		v1658 = v37
		v1659 = v1030
		goto L4
	} else {
		goto L272
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1028 | int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v1658 = v37
		v1659 = v1030
		goto L4
	} else {
		goto L271
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1028 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v1658 = v37
		v1659 = v1030
		goto L4
	} else {
		goto L270
	}
L266:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	switch v1031 - int32(12) {
	case 0:
		goto L268
	default:
		v1658 = v37
		v1659 = v1030
		goto L4
	case 4:
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(31776)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(3)
	v1045 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v1045
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1045
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
	v1158 = v81
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(31744)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(3)
	F_GetCurrentTimeUsec(m, l4, l5, int32(0))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L28
	} else {
		goto L269
	}
L269:
	;
	v1158 = v81
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L270:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1058 - v1059
	v1158 = v81
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L271:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1070
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1158 = v81
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L272:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1080
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1158 = v81
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L273:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v37)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1158 = v81
	v1163 = v86
	v1164 = v87
	v1171 = v1091
	v1173 = v96
	v1175 = v1090
	v1176 = v99
	goto L12
L274:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1158 = v1100
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L275:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1158 = v1103
	v1163 = v86
	v1164 = v87
	v1171 = v94
	v1173 = v96
	v1175 = v98
	v1176 = v99
	goto L12
L276:
	;
	if v1105 != 0 {
		v1128 = v1105
		goto L13
	} else {
		goto L277
	}
L277:
	;
	v1658 = v37
	v1659 = v1030
	goto L4
L278:
	;
	v1658 = v37
	v1659 = int32(-1)
	goto L4
L279:
	;
	goto L280
L280:
	;
	v1194 = v1179 | v79
	v1196 = v1158
	v1201 = v1163
	v1202 = v1164
	v1209 = v1171
	v1211 = v1173
	v1213 = v1175
	v1214 = v1176
	goto L11
L281:
	;
	goto L10
L282:
	;
	v1229 = v37
	v1232 = v1194
	v1239 = v1201
	v1240 = v1202
	v1247 = v1209
	v1249 = v1211
	v1251 = v1213
	v1252 = v1214
	goto L5
L283:
	;
	if v1423 != 0 {
		v1658 = v1229
		v1659 = v1423
		goto L4
	} else {
		goto L321
	}
L284:
	;
	if v1232&int32(32768) != 0 {
		goto L302
	} else {
		goto L303
	}
L285:
	;
	if v1232&int32(4) == int32(0) {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1252 != 0 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1280
	goto L284
L288:
	;
	v1280 = int32(1) - v1260
	goto L287
L289:
	;
	if int32(0) < v1260 {
		goto L288
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	if v1255 != 0 {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	v1423 = int32(-2)
	goto L283
L293:
	;
	if v1260 < int32(0) {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	goto L295
L295:
	;
	if int32(0) < v1260 {
		goto L284
	} else {
		goto L301
	}
L296:
	;
	v1423 = int32(-2)
	goto L283
L297:
	;
	goto L298
L298:
	;
	if base.Ui32(v1260) <= base.Ui32(int32(69)) {
		v1280 = v1260 + int32(2000)
		goto L287
	} else {
		goto L299
	}
L299:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1260) {
		goto L284
	} else {
		goto L300
	}
L300:
	;
	v1280 = v1260 + int32(1900)
	goto L287
L301:
	;
	v1423 = int32(-2)
	goto L283
L302:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v1291 = v1286 + int32(4799)
	v1293 = base.I32_div_s(v1291, int32(4))
	v1296 = base.I32_div_s(v1291, int32(-100))
	v1299 = base.I32_div_s(v1291, int32(400))
	v1300 = v1285 + v1286*int32(365) + v1293 + v1296 + v1299
	v1302 = v1300 + int32(1751940)
	v1303 = int32(146097)
	v1304 = base.I32_div_u_s(v1302, v1303)
	v1305 = int32(3)
	v1311 = int32(2)
	v1316 = base.I32_div_u_s((v1304*int32(1073595727)+v1302)<<(uint(v1311)%32)|v1305, v1303)
	v1319 = v1300 + v1304*v1305 + v1316 + int32(1752000)
	v1320 = int32(1461)
	v1321 = base.I32_div_u_s(v1319, v1320)
	v1324 = v1321*int32(-1461) + v1319
	v1326 = v1324 << (uint(v1311) % 32)
	if base.Ui32(v1320) <= base.Ui32(v1326) {
		goto L306
	} else {
		goto L307
	}
L303:
	;
	goto L304
L304:
	;
	if v1232&int32(2) == int32(0) {
		goto L309
	} else {
		goto L310
	}
L305:
	;
	v1339 = base.I32_div_u_s(v1326, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1339 + v1321<<(uint(int32(2))%32) - int32(4800)
	v1347 = v1337 + int32(123)
	v1351 = int32(base.Ui32(v1347*int32(2141)) >> (uint(int32(16)) % 32))
	v1355 = base.I32_rem_u_s(v1351+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v1355 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v1347 - int32(base.Ui32(v1351*int32(7834))>>(uint(int32(8))%32))
	goto L304
L306:
	;
	v1332 = base.I32_rem_u_s(v1324+int32(305), int32(365))
	v1337 = v1332
	goto L305
L307:
	;
	goto L308
L308:
	;
	v1336 = base.I32_rem_u_s(v1324+int32(306), int32(366))
	v1337 = v1336
	goto L305
L309:
	;
	if v1232&int32(8) == int32(0) {
		goto L312
	} else {
		goto L313
	}
L310:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if base.Ui32(int32(-12)) <= base.Ui32(v1372-int32(13)) {
		goto L309
	} else {
		goto L311
	}
L311:
	;
	v1423 = int32(-3)
	goto L283
L312:
	;
	v1388 = int32(14)
	if v1232&v1388 != v1388 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if base.Ui32(int32(-31)) <= base.Ui32(v1382-int32(32)) {
		goto L312
	} else {
		goto L314
	}
L314:
	;
	v1423 = int32(-3)
	goto L283
L315:
	;
	v1423 = int32(0)
	goto L283
L316:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1394&int32(3) != 0 {
		v1404 = int32(0)
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1404*int32(52)+v1407<<(uint(int32(2))%32))+uint32(_consts[1030])))
	if v1392 <= v1413 {
		goto L315
	} else {
		goto L320
	}
L318:
	;
	v1399 = base.I32_rem_s(v1394, int32(100))
	if v1399 != 0 {
		v1404 = int32(1)
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v1401 = base.I32_rem_s(v1394, int32(400))
	v1404 = base.B2i32(v1401 == int32(0))
	goto L317
L320:
	;
	v1423 = int32(-2)
	goto L283
L321:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v1239 == int32(2) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1447 = int32(1)
	if base.Ui32(int32(24)) < base.Ui32(v1441) {
		v1469 = v1447
		goto L335
	} else {
		goto L336
	}
L323:
	;
	v1441 = v1424
	goto L322
L324:
	;
	goto L325
L325:
	;
	if int32(12) < v1424 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1658 = v1229
	v1659 = int32(-2)
	goto L4
L327:
	;
	goto L328
L328:
	;
	switch v1239 {
	case 0:
		goto L331
	case 1:
		goto L330
	default:
		v1441 = v1424
		goto L322
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v1439
	v1441 = v1439
	goto L322
L330:
	;
	v1433 = int32(12)
	if v1424 == v1433 {
		v1441 = v1433
		goto L322
	} else {
		goto L333
	}
L331:
	;
	if v1424 == int32(12) {
		v1439 = int32(0)
		goto L329
	} else {
		goto L332
	}
L332:
	;
	v1441 = v1424
	goto L322
L333:
	;
	v1439 = v1424 + int32(12)
	goto L329
L334:
	;
	if v1469 != 0 {
		goto L340
	} else {
		goto L341
	}
L335:
	;
	goto L334
L336:
	;
	if base.Ui32(int32(59)) < base.Ui32(v1443) {
		v1469 = v1447
		goto L335
	} else {
		goto L337
	}
L337:
	;
	if base.Ui32(int32(60)) < base.Ui32(v1444) {
		v1469 = v1447
		goto L335
	} else {
		goto L338
	}
L338:
	;
	if base.Ui32(int32(1000000)) < base.Ui32(v1445) {
		v1469 = v1447
		goto L335
	} else {
		goto L339
	}
L339:
	;
	v1457 = int32(60)
	v1469 = base.B2i32(base.Ui64(int64(86400000000)) < base.Ui64(base.I64_extend_i32_u(v1445)+base.I64_extend_i32_u((v1441*v1457+v1443)*v1457+v1444)*int64(1000000)))
	goto L335
L340:
	;
	v1658 = v1229
	v1659 = int32(-2)
	goto L4
L341:
	;
	goto L342
L342:
	;
	v1471 = int32(-1)
	v1472 = int32(31744)
	if v1232&v1472 != v1472 {
		v1658 = v1229
		v1659 = v1471
		goto L4
	} else {
		goto L343
	}
L343:
	;
	if v1240 != 0 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v1658 = v1229
	v1659 = v1649
	goto L4
L345:
	;
	if v1232&int32(268435456) != 0 {
		v1649 = v1471
		goto L344
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	if v1247 != 0 {
		goto L363
	} else {
		goto L364
	}
L348:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+uint32(_consts[1031])))
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1240)+264))
	if v1484 < int32(2) {
		goto L351
	} else {
		goto L352
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1527
	goto L347
L350:
	;
	if v1516 != 0 {
		goto L359
	} else {
		goto L360
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1229+int32(72)))) = v1483
	v1516 = int32(1)
	goto L350
L352:
	;
	v1490 = int32(1)
	goto L353
L353:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1240+int32(18280)+v1490<<(uint(int32(4))%32))))
	if v1483 == v1498 {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v1516 = int32(0)
	goto L350
L355:
	;
	v1501 = v1490 + int32(1)
	if v1484 != v1501 {
		v1490 = v1501
		goto L353
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	goto L354
L358:
	;
	goto L351
L359:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+72))
	v1527 = int32(0) - v1518
	goto L349
L360:
	;
	goto L361
L361:
	;
	v1520 = int32(14)
	if v1232&v1520 != v1520 {
		v1649 = v1471
		goto L344
	} else {
		goto L362
	}
L362:
	;
	v1526 = F_DetermineTimeZoneOffsetInternal(m, l4, v1240, v1229+int32(8))
	mBase = m.M
	v1527 = v1526
	goto L349
L363:
	;
	if v1232&int32(268435456) != 0 {
		v1658 = v1229
		v1659 = v1471
		goto L4
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1610 = int32(0)
	if l6 == v1610 {
		v1658 = v1229
		v1659 = v1610
		goto L4
	} else {
		goto L381
	}
L366:
	;
	switch v1232 & int32(14) {
	case 0:
		goto L368
	default:
		v1658 = v1229
		v1659 = v1471
		goto L4
	case 14:
		goto L369
	}
L367:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+16)) = v1546
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+12)) = v1548
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+8)) = v1550
	v1553 = v1229 + int32(8)
	v1557 = m.G0
	v1559 = v1557 - int32(288)
	m.G0 = v1559
	v1563 = F_DetermineTimeZoneOffsetInternal(m, v1553, v1247, v1559+int32(280))
	mBase = m.M
	v1567 = F_strlcpy(m, v1559+int32(16), v1251, int32(256))
	mBase = m.M
	v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1559)+16)))
	if v1568 != 0 {
		goto L372
	} else {
		goto L373
	}
L368:
	;
	F_GetCurrentTimeUsec(m, v1229+int32(8), v1229+int32(72), int32(0))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L28
	} else {
		goto L370
	}
L369:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+28)) = v1533
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+24)) = v1535
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+20)) = v1537
	goto L367
L370:
	;
	goto L367
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1603
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1608
	goto L365
L372:
	;
	v1572 = v1559 + int32(16)
	v1576 = v1568
	goto L375
L373:
	;
	goto L374
L374:
	;
	v1596 = F_pg_interpret_timezone_abbrev(m, v1559+int32(16), v1559+int32(280), v1559+int32(12), v1559+int32(8), v1247)
	mBase = m.M
	if v1596 != 0 {
		goto L378
	} else {
		goto L379
	}
L375:
	;
	v1577 = F_pg_toupper(m, v1576)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1572))) = uint8(v1577)
	v1580 = v1572 + int32(1)
	v1581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1580))))
	if v1581 != 0 {
		v1572 = v1580
		v1576 = v1581
		goto L375
	} else {
		goto L377
	}
L376:
	;
	goto L374
L377:
	;
	goto L376
L378:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1559)+12))
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1559)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1553)+32)) = v1598
	v1603 = int32(0) - v1597
	goto L380
L379:
	;
	v1603 = v1563
	goto L380
L380:
	;
	m.G0 = v1559 + int32(288)
	goto L371
L381:
	;
	if v1232&int32(32) != 0 {
		v1658 = v1229
		v1659 = v1610
		goto L4
	} else {
		goto L382
	}
L382:
	;
	if v1232&int32(268435456) != 0 {
		goto L385
	} else {
		goto L386
	}
L383:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+16)) = v1633
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+12)) = v1635
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+8)) = v1637
	v1642 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	v1645 = F_DetermineTimeZoneOffsetInternal(m, v1229+int32(8), v1642, v1229+int32(72))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1645
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1647
	v1649 = v1610
	goto L344
L384:
	;
	F_GetCurrentTimeUsec(m, v1229+int32(8), v1229+int32(72), int32(0))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L28
	} else {
		goto L388
	}
L385:
	;
	v1658 = v1229
	v1659 = int32(-1)
	goto L4
L386:
	;
	switch v1232 & int32(14) {
	case 0:
		goto L384
	default:
		goto L385
	case 14:
		goto L387
	}
L387:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+28)) = v1619
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+24)) = v1621
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1229)+20)) = v1623
	goto L383
L388:
	;
	goto L383
}
func F_extract_time(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_time_part_common(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_readTimeLineHistory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int64
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	v2 = int32(0)
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(2288)
	m.G0 = v12
	if l0 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(2288)
	return v359
L2:
	;
	v17 = F_palloc(m, int32(24))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[94])))
	if v35 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	return int32(0)
L6:
	;
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v21
	v23 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1196)) = v17
	v32 = F_list_make1_impl(m, v23, v12+int32(8))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v359 = v32
	goto L1
L8:
	;
	v69 = F_AllocateFile(m, v12+int32(1264), int32(242658))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L18
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = l0
	v45 = F_pg_snprintf(m, v12+int32(1200), int32(64), int32(12998), v12+int32(128))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = l0
	v63 = F_pg_snprintf(m, v12+int32(1264), int32(1024), int32(12991), v12+int32(144))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	v54 = F_RestoreArchivedFile(m, v12+int32(1264), v12+int32(1200), int32(534635), int64(0), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v65 = v54
	goto L8
L14:
	;
	v65 = v2
	goto L8
L15:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v69)+76))
	if v283 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L16:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L5
	} else {
		goto L63
	}
L17:
	;
	v108 = v2
	v109 = v2
	v111 = v9
	goto L29
L18:
	;
	if v69 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(167772219)
	v78 = F_fgets(m, v12+int32(160), int32(1024), v69)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[85]))
	if v85 != int32(44) {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = int32(0)
	if v78 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v279 = v2
	v280 = v2
	v282 = v9
	goto L15
L24:
	;
	v89 = F_palloc(m, int32(24))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v91 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v89)+16)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1192)) = v89
	v101 = F_list_make1_impl(m, int32(1), v12+int32(12))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v359 = v101
	goto L1
L27:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L59
	}
L28:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L55
	}
L29:
	;
	v116 = v12 + int32(160)
	goto L31
L30:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L5
	} else {
		goto L51
	}
L31:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if base.Ui32(v123-int32(9)) < base.Ui32(int32(5)) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L30
L33:
	;
	goto L32
L34:
	;
	v116 = v116 + int32(1)
	goto L31
L35:
	;
	switch v123 - int32(32) {
	case 0:
		goto L34
	case 1, 2:
		goto L37
	case 3:
		v172 = v108
		v173 = v109
		v174 = v111
		goto L36
	default:
		goto L38
	}
L36:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = int32(167772219)
	v182 = F_fgets(m, v12+int32(160), int32(1024), v69)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L49
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v12 + int32(148)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+116)) = v12 + int32(152)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v12 + int32(156)
	v146 = F_sscanf(m, v12+int32(160), int32(542031), v12+int32(112))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	if v123 == int32(0) {
		v172 = v108
		v173 = v109
		v174 = v111
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	if v146 <= int32(0) {
		goto L33
	} else {
		goto L41
	}
L41:
	;
	if v146 != int32(3) {
		goto L28
	} else {
		goto L42
	}
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	if base.Ui32(v153) <= base.Ui32(v109) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v155 = v108
	goto L45
L44:
	;
	v155 = int32(0)
	goto L45
L45:
	;
	if v155 != 0 {
		goto L27
	} else {
		goto L46
	}
L46:
	;
	v157 = F_palloc(m, int32(24))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+8)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v159
	v162 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+148)))
	v163 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+152)))
	v166 = v162 | v163<<(uint(int64(32))%64)
	*(*int64)(unsafe.Add(mBase, uint32(v157)+16)) = v166
	v168 = F_lcons(m, v157, v108)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v172 = v168
	v173 = v153
	v174 = v166
	goto L36
L49:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[95]))
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v186
	if v182 == v186 {
		v279 = v172
		v280 = v173
		v282 = v174
		goto L15
	} else {
		goto L50
	}
L50:
	;
	v108 = v172
	v109 = v173
	v111 = v174
	goto L29
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v12 + int32(160)
	F_errmsg(m, int32(214256), v12-int32(-64))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_errhint(m, int32(689732), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(523822), int32(164), int32(13373))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v12 + int32(160)
	F_errmsg(m, int32(214256), v12+int32(96))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errhint(m, int32(645488), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(523822), int32(169), int32(13373))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v12 + int32(160)
	F_errmsg(m, int32(214289), v12+int32(80))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(672866), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(523822), int32(174), int32(13373))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(1264)
	F_errmsg(m, int32(313019), v12+int32(16))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(523822), int32(111), int32(13373))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v341 = F_palloc(m, int32(24))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L5
	} else {
		goto L87
	}
L68:
	;
	if int32(base.Ui32(v288)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	goto L68
L70:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v288 = v286
	goto L69
L71:
	;
	goto L72
L72:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v288 = v287
	goto L69
L73:
	;
	v295 = F_FreeFile(m, v69)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L5
	} else {
		goto L83
	}
L76:
	;
	if v279 == int32(0) {
		goto L67
	} else {
		goto L77
	}
L77:
	;
	if base.Ui32(v280) < base.Ui32(l0) {
		goto L67
	} else {
		goto L78
	}
L78:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(1264)
	F_errmsg(m, int32(749794), v12+int32(32))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	F_errhint(m, int32(689680), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(523822), int32(195), int32(13373))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(1264)
	F_errmsg(m, int32(313938), v12+int32(48))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(523822), int32(143), int32(13373))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v341)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v341)+8)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = l0
	v347 = F_lcons(m, v341, v279)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	if v65 == int32(0) {
		v359 = v347
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_KeepFileRestoredFromArchive(m, v12+int32(1264), v12+int32(1200))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	v359 = v347
	goto L1
}
func F_time_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v4
		return v6
	}
}
func F_time_part(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_time_part_common(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_time_t_to_timestamptz(m *base.Module, l0 int64) int64 {
	return l0*int64(1000000) - int64(946684800000000)
}
func F_time_timetz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	F_GetCurrentDateTime(m, v9+int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = base.I64_div_s(v12, int64(3600000000))
		*(*uint32)(unsafe.Add(mBase, uint32(v9)+12)) = uint32(v20)
		v25 = v12 + base.I64_extend32_s(v20)*int64(-3600000000)
		v27 = base.I64_div_s(v25, int64(60000000))
		*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)) = uint32(v27)
		v34 = base.I64_div_s(base.I64_extend32_s(v27)*int64(-60000000)+v25, int64(1000000))
		*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v34)
		v39 = *(*int32)(unsafe.Add(mBase, _consts[326]))
		v41 = m.G0
		v42 = int32(16)
		v43 = v41 - v42
		m.G0 = v43
		v47 = F_DetermineTimeZoneOffsetInternal(m, v9+int32(4), v39, v43+int32(8))
		mBase = m.M
		m.G0 = v43 + v42
		v52 = F_palloc(m, int32(16))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v47
			*(*int64)(unsafe.Add(mBase, uint32(v52))) = v12
			m.G0 = v9 + int32(48)
			return v52
		}
	}
}
