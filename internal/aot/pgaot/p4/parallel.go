package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v52 int32
	_ = v52
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
	var v202 int32
	_ = v202
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
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v436 int32
	_ = v436
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
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
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
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
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
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
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v878 float64
	_ = v878
	var v882 int32
	_ = v882
	var v885 float64
	_ = v885
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 float64
	_ = v913
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1078 float64
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
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
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1168 int32
	_ = v1168
	var v1180 int32
	_ = v1180
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1216 int32
	_ = v1216
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1417 int32
	_ = v1417
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1463 int32
	_ = v1463
	var v1467 int32
	_ = v1467
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1585 int32
	_ = v1585
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
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
	v52 = v27
	goto L3
L3:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[0]))
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
	v52 = v74
	goto L3
L10:
	;
	m.G0 = v25 + int32(48)
	return v1636
L11:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+72))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+16))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+8))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+12))
	m.T0[v1612].(func(*base.Module, int32))(m, v1610)
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L1
	} else {
		goto L306
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L1
	} else {
		goto L303
	}
L13:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+48))
	if int32(0) <= v1083 {
		goto L233
	} else {
		goto L234
	}
L14:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v922 = v917
	goto L203
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+168)))
	if v891 != 0 {
		goto L3
	} else {
		goto L196
	}
L16:
	;
	v682 = m.G0
	v684 = v682 - int32(16)
	m.G0 = v684
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v689 != 0 {
		goto L145
	} else {
		goto L146
	}
L17:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+44))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v296)+48))
	if v298|base.B2i32(v297 != int32(1)) == int32(0) {
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
	v1636 = int32(0)
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
	v1636 = v110
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
	v167 = int32(_a_F_ExecParallelHashJoin_0)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v171
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
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v176
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+43)))
	if v181 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[0]))
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
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v206 = (v192 - int32(1)) & base.I32_rotr(v186, v202)
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
	v1636 = int32(0)
	goto L10
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v416
	v540 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+168)) = uint8(v540)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v542
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v52)+44))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(132)))) = (v548 - int32(1)) & v542
	if base.Ui32(int32(2)) <= base.Ui32(v547) {
		goto L118
	} else {
		goto L119
	}
L68:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v463 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L69:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+4)))
	if v436&int32(2) == int32(0) {
		goto L67
	} else {
		goto L99
	}
L70:
	;
	F_ExecForceStoreMinimalTuple(m, v373, v375, int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L97
	}
L71:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v296)+144))
	v406 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v402+v298*int32(36))+25)) = uint8(v406)
	goto L68
L72:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	if v304 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	if v297 <= v298 {
		goto L71
	} else {
		goto L93
	}
L75:
	;
	F_ExecReScan(m, v28)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v308 = m.T0[v307].(func(*base.Module, int32) int32)(m, v28)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if v308 == int32(0) {
		goto L71
	} else {
		goto L80
	}
L80:
	;
	v314 = v308
	goto L81
L81:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+4)))
	if v334&int32(2) != 0 {
		goto L71
	} else {
		goto L83
	}
L82:
	;
	goto L71
L83:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v337)+12)) = v314
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v337)+20))
	F_MemoryContextReset(m, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v342 = int32(_a_F_ExecParallelHashJoin_0)
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v337)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v346
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v344)+20))
	v351 = m.T0[v350].(func(*base.Module, int32, int32, int32) int32)(m, v344, v337, v25+int32(44))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v351
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+44)))
	if v356 == int32(0) {
		v416 = v314
		goto L69
	} else {
		goto L86
	}
L86:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v28)+52))
	if v359 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_ExecReScan(m, v28)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v363 = m.T0[v362].(func(*base.Module, int32) int32)(m, v28)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	if v363 != 0 {
		v314 = v363
		goto L81
	} else {
		goto L92
	}
L92:
	;
	goto L82
L93:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v296)+144))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366+v298*int32(36))+32))
	v373 = F_sts_parallel_scan_next(m, v370, v25+int32(24))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v373 != 0 {
		goto L70
	} else {
		goto L95
	}
L95:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+8))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	m.T0[v377].(func(*base.Module, int32))(m, v375)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L71
L97:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v411 == int32(0) {
		goto L68
	} else {
		goto L98
	}
L98:
	;
	v416 = v411
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
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+144))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v466)+48))
	v470 = v468 * int32(36)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v467+v470)))
	v475 = F_BarrierArriveAndDetachExceptLast(m, v472+int32(4))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L103
	}
L102:
	;
	if v529 == int32(0) {
		goto L100
	} else {
		goto L116
	}
L103:
	;
	if v475 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v466)+144))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v466)+48))
	v484 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v479+v480*int32(36))+26)) = uint8(v484)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v466)+144))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v486+v470)+28))
	F_sts_end_parallel_scan(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+61)))
	if v508 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L107:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v466)+144))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v491+v470)+32))
	F_sts_end_parallel_scan(m, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v472)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v466)+48)) = int32(-1)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v466)+104))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v466)))
	v503 = v496 + v500<<(uint(int32(2))%32)
	if base.Ui32(v503) < base.Ui32(v499) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v505 = v499
	goto L111
L110:
	;
	v505 = v503
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v466)+104)) = v505
	v529 = int32(0)
	goto L102
L112:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v466)+144))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v466)+48))
	v516 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v511+v512*int32(36))+26)) = uint8(v516)
	F_ExecHashTableDetachBatch(m, v466)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
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
	v529 = int32(1)
	goto L102
L115:
	;
	v529 = int32(0)
	goto L102
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(5)
	goto L3
L117:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+24)))
	if v567 != int32(1) {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v561 = (v547 - int32(1)) & base.I32_rotr(v542, v557)
	goto L120
L119:
	;
	v561 = int32(0)
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+int32(20)))) = v561
	goto L117
L121:
	;
	v607 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v607
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v606
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	if base.B2i32(v610 == v611)|base.B2i32(v606 != int32(-1)) == v607 {
		goto L131
	} else {
		goto L132
	}
L122:
	;
	v606 = int32(-1)
	goto L121
L123:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v52)+28))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v573 = v571 - int32(1)
	v574 = v563 & v573
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v570+v574<<(uint(int32(2))%32))))
	if v578 == int32(0) {
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v581 = v574
	v583 = v578
	goto L125
L125:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	if v563 == v586 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L122
L127:
	;
	v606 = v581
	goto L121
L128:
	;
	goto L129
L129:
	;
	v590 = (v581 + int32(1)) & v573
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v570+v590<<(uint(int32(2))%32))))
	if v594 != 0 {
		v581 = v590
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
	v620 = F_ExecFetchSlotMinimalTuple(m, v416, v25+int32(36))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(3)
	goto L16
L134:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v52)+92))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v624
	v628 = v622 + v623<<(uint(int32(2))%32)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)))
	if v629 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v632 = int32(_a_F_ExecParallelHashJoin_0)
	v633 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v52)+124))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v635
	v638 = F_BufFileCreateTemp(m, int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	v643 = v629
	goto L137
L137:
	;
	F_BufFileWrite(m, v643, v25+int32(44), int32(4))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v628))) = v638
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v633
	v643 = v638
	goto L137
L139:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	F_BufFileWrite(m, v643, v620, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+36)))
	if v653 != int32(1) {
		goto L3
	} else {
		goto L141
	}
L141:
	;
	F_pfree(m, v620)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	goto L3
L143:
	;
	m.G0 = v684 + int32(16)
	if v810 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L144:
	;
	if v703 != 0 {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v687)+136))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v689)))
	v692 = F_dsa_get_address(m, v690, v691)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v687)+136))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v687)+20))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v695+v696<<(uint(int32(2))%32))))
	v701 = F_dsa_get_address(m, v694, v700)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	v703 = v692
	goto L144
L149:
	;
	v703 = v701
	goto L144
L150:
	;
	v708 = v703
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
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v708)+4))
	if v726 != v686 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L152
L155:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v687)+136))
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
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v732 = F_ExecStoreMinimalTuple(m, v708+int32(8), v730, int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v732
	if v688 == int32(0) {
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
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_MemoryContextReset(m, v737)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v740 = int32(_a_F_ExecParallelHashJoin_0)
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v743
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v688)+20))
	v748 = m.T0[v747].(func(*base.Module, int32, int32, int32) int32)(m, v688, v33, v684+int32(15))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L163
	}
L162:
	;
	goto L158
L163:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v741
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_MemoryContextReset(m, v752)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	if v748 == int32(0) {
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
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v882 == int32(0) {
		goto L3
	} else {
		goto L195
	}
L176:
	;
	v825 = int32(_a_F_ExecParallelHashJoin_0)
	v826 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v828
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
	v840 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+168)) = uint8(v840)
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v843 = int32(*(*int16)(unsafe.Add(mBase, uint32(v842)+18)))
	if int32(0) <= v843 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v826
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
	v847 = v843 | int32(_a_F_ExecParallelHashJoin_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v842)+18)) = uint16(v847)
	goto L183
L182:
	;
	goto L183
L183:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v849 == int32(5) {
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
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+108)))
	if v854 == int32(1) {
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
	if v849 == int32(7) {
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
	v863 = int32(_a_F_ExecParallelHashJoin_0)
	v864 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v866
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v871 = m.T0[v870].(func(*base.Module, int32, int32, int32) int32)(m, v29, v33, v25+int32(44))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v864
	if v871 != 0 {
		goto L11
	} else {
		goto L193
	}
L193:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v875 == int32(0) {
		goto L3
	} else {
		goto L194
	}
L194:
	;
	v878 = *(*float64)(unsafe.Add(mBase, uint32(v875)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v875)+248)) = base.F64_add(v878, float64(1))
	goto L3
L195:
	;
	v885 = *(*float64)(unsafe.Add(mBase, uint32(v882)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v882)+240)) = base.F64_add(v885, float64(1))
	goto L3
L196:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v892 == int32(0) {
		goto L3
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v892
	if v29 == int32(0) {
		goto L11
	} else {
		goto L198
	}
L198:
	;
	v898 = int32(_a_F_ExecParallelHashJoin_0)
	v899 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v901
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v906 = m.T0[v905].(func(*base.Module, int32, int32, int32) int32)(m, v29, v33, v25+int32(44))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v899
	if v906 != 0 {
		goto L11
	} else {
		goto L200
	}
L200:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v910 == int32(0) {
		goto L3
	} else {
		goto L201
	}
L201:
	;
	v913 = *(*float64)(unsafe.Add(mBase, uint32(v910)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v910)+248)) = base.F64_add(v913, float64(1))
	goto L3
L202:
	;
	if v1054 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L203:
	;
	if v922 != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	if v961 != 0 {
		goto L212
	} else {
		goto L213
	}
L206:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v918)+136))
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v922)))
	v943 = F_dsa_get_address(m, v941, v942)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L1
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v918)))
	if v947 <= v946 {
		v1054 = int32(0)
		goto L202
	} else {
		goto L210
	}
L209:
	;
	v961 = v943
	goto L205
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v946 + int32(1)
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v918)+136))
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v918)+20))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v953+v946<<(uint(int32(2))%32))))
	v958 = F_dsa_get_address(m, v952, v957)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v961 = v958
	goto L205
L212:
	;
	v965 = v961
	goto L215
L213:
	;
	goto L214
L214:
	;
	v1025 = int32(0)
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[0]))
	if v1027 == v1025 {
		v922 = v1025
		goto L203
	} else {
		goto L224
	}
L215:
	;
	v984 = int32(*(*int16)(unsafe.Add(mBase, uint32(v965)+18)))
	if int32(0) <= v984 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	goto L214
L217:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v991 = F_ExecStoreMinimalTuple(m, v965+int32(8), v989, int32(0))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L1
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v918)+136))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v965)))
	v1001 = F_dsa_get_address(m, v999, v1000)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v991
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_MemoryContextReset(m, v994)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v965
	v1054 = int32(1)
	goto L202
L222:
	;
	if v1001 != 0 {
		v965 = v1001
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
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v922 = v1025
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
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v1059
	if v29 == int32(0) {
		goto L11
	} else {
		goto L229
	}
L229:
	;
	v1063 = int32(_a_F_ExecParallelHashJoin_0)
	v1064 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1066
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v1071 = m.T0[v1070].(func(*base.Module, int32, int32, int32) int32)(m, v29, v33, v25+int32(44))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1064
	if v1071 != 0 {
		goto L11
	} else {
		goto L231
	}
L231:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1075 == int32(0) {
		goto L3
	} else {
		goto L232
	}
L232:
	;
	v1078 = *(*float64)(unsafe.Add(mBase, uint32(v1075)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v1075)+248)) = base.F64_add(v1078, float64(1))
	goto L3
L233:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+144))
	v1090 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1086+v1083*int32(36))+26)) = uint8(v1090)
	F_ExecHashTableDetachBatch(m, v1082)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+140))
	v1097 = base.AtomicRmwAdd32(m, v1094, int32(164), int32(1))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+44))
	v1099 = base.I32_rem_u_s(v1097, v1098)
	v1102 = v1099
	goto L237
L236:
	;
	goto L235
L237:
	;
	v1123 = v1102 * int32(36)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+144))
	v1125 = v1123 + v1124
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125)+26)))
	if v1126 != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1636 = int32(0)
	goto L10
L239:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+44))
	v1569 = base.I32_rem_s(v1102+int32(1), v1568)
	if v1569 != v1099 {
		v1102 = v1569
		goto L237
	} else {
		goto L302
	}
L240:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1125)))
	v1129 = v1127 + int32(4)
	v1130 = F_BarrierAttach(m, v1129)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L248
	}
L241:
	;
	F_ExecParallelHashTableSetCurrentBatch(m, v1082, v1102)
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L1
	} else {
		goto L300
	}
L242:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L1
	} else {
		goto L297
	}
L243:
	;
	F_BarrierDetach(m, v1129)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L1
	} else {
		goto L296
	}
L244:
	;
	F_ExecParallelHashTableSetCurrentBatch(m, v1082, v1102)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L294
	}
L245:
	;
	F_ExecParallelHashTableSetCurrentBatch(m, v1082, v1102)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L266
	}
L246:
	;
	v1312 = F_BarrierArriveAndWait(m, v1129, int32(134217742))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L265
	}
L247:
	;
	v1133 = F_BarrierArriveAndWait(m, v1129, int32(134217743))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L1
	} else {
		goto L249
	}
L248:
	;
	switch v1130 {
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
	if v1133 == int32(0) {
		goto L246
	} else {
		goto L250
	}
L250:
	;
	v1137 = int32(0)
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+144))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1138+v1102*int32(36))))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+136))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+140))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+16))
	v1149 = F_dsa_allocate_extended(m, v1143, v1145<<(uint(int32(2))%32), v1137)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1142))) = v1149
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+136))
	v1153 = F_dsa_get_address(m, v1152, v1149)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	if v1145 <= int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	goto L246
L254:
	;
	v1158 = v1145 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v1145) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1168 = v1137
	v1180 = int32(0)
	goto L258
L256:
	;
	v1216 = v1137
	goto L257
L257:
	;
	v1238 = int32(0)
	v1239 = v1216
	goto L262
L258:
	;
	v1188 = v1153 + v1168<<(uint(int32(2))%32)
	v1189 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1188))) = v1189
	*(*int32)(unsafe.Add(mBase, uint32(v1188)+4)) = v1189
	*(*int32)(unsafe.Add(mBase, uint32(v1188)+8)) = v1189
	*(*int32)(unsafe.Add(mBase, uint32(v1188)+12)) = v1189
	*(*int32)(unsafe.Add(mBase, uint32(v1188)+16)) = v1189
	*(*int32)(unsafe.Add(mBase, uint32(v1188)+20)) = v1189
	*(*int32)(unsafe.Add(mBase, uint32(v1188)+24)) = v1189
	*(*int32)(unsafe.Add(mBase, uint32(v1188)+28)) = v1189
	v1205 = int32(8)
	v1206 = v1168 + v1205
	v1208 = v1180 + v1205
	if v1208 != v1145&int32(-8) {
		v1168 = v1206
		v1180 = v1208
		goto L258
	} else {
		goto L260
	}
L259:
	;
	if v1158 == int32(0) {
		goto L253
	} else {
		goto L261
	}
L260:
	;
	goto L259
L261:
	;
	v1216 = v1206
	goto L257
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1153+v1239<<(uint(int32(2))%32)))) = int32(0)
	v1262 = int32(1)
	v1265 = v1238 + v1262
	if v1265 != v1158 {
		v1238 = v1265
		v1239 = v1239 + v1262
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
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+144))
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1338+v1123)+28))
	F_sts_begin_parallel_scan(m, v1340)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	v1345 = F_sts_parallel_scan_next(m, v1340, v25+int32(44))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	if v1345 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1350 = v1345
	goto L272
L270:
	;
	goto L271
L271:
	;
	F_sts_end_parallel_scan(m, v1340)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L1
	} else {
		goto L292
	}
L272:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	F_ExecForceStoreMinimalTuple(m, v1350, v1369, int32(0))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L1
	} else {
		goto L274
	}
L273:
	;
	goto L271
L274:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v1375 = m.G0
	v1377 = v1375 - int32(16)
	m.G0 = v1377
	v1381 = F_ExecFetchSlotMinimalTuple(m, v1373, v1377+int32(15))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1082)))
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1381)))
	v1385 = int32(8)
	v1389 = F_ExecParallelHashTupleAlloc(m, v1082, v1384+v1385, v1377+v1385)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1389)+4)) = v1374
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1381)))
	if v1392 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	base.MemoryCopy(m, v1389+int32(8), v1381, v1392)
	goto L279
L278:
	;
	goto L279
L279:
	;
	v1396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1389)+18)))
	v1398 = v1396 & int32(_a_F_ExecParallelHashJoin_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v1389)+18)) = uint16(v1398)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+8))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+20))
	v1407 = v1401 + (v1383-int32(1))&v1374<<(uint(int32(2))%32)
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1407)))
	*(*int32)(unsafe.Add(mBase, uint32(v1389))) = v1408
	v1411 = base.AtomicRmwCmpxchg32(m, v1407, int32(0), v1408, v1400)
	if v1408 != v1411 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1417 = v1411
	goto L283
L281:
	;
	goto L282
L282:
	;
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1377)+15)))
	if v1463 == int32(1) {
		goto L286
	} else {
		goto L287
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1389))) = v1417
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1407)))
	*(*int32)(unsafe.Add(mBase, uint32(v1389))) = v1436
	v1439 = base.AtomicRmwCmpxchg32(m, v1407, int32(0), v1436, v1400)
	if v1436 != v1439 {
		v1417 = v1439
		goto L283
	} else {
		goto L285
	}
L284:
	;
	goto L282
L285:
	;
	goto L284
L286:
	;
	F_pfree(m, v1381)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	m.G0 = v1377 + int32(16)
	v1473 = F_sts_parallel_scan_next(m, v1340, v25+int32(44))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L1
	} else {
		goto L290
	}
L289:
	;
	goto L288
L290:
	;
	if v1473 != 0 {
		v1350 = v1473
		goto L272
	} else {
		goto L291
	}
L291:
	;
	goto L273
L292:
	;
	v1500 = F_BarrierArriveAndWait(m, v1129, int32(134217744))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	goto L244
L294:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+144))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1526+v1123)+32))
	F_sts_begin_parallel_scan(m, v1528)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(2)
	goto L3
L296:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+144))
	v1537 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1535+v1123)+26)) = uint8(v1537)
	*(*int32)(unsafe.Add(mBase, uint32(v1082)+48)) = int32(-1)
	goto L239
L297:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v1545
	F_errmsg_internal(m, int32(_a_F_ExecParallelHashJoin_3), v25+int32(16))
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_ExecParallelHashJoin_4), int32(1392), int32(_a_F_ExecParallelHashJoin_5))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1082)+144))
	v1561 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1559+v1123)+26)) = uint8(v1561)
	F_ExecHashTableDetachBatch(m, v1082)
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	goto L239
L302:
	;
	goto L238
L303:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v1576
	F_errmsg_internal(m, int32(_a_F_ExecParallelHashJoin_6), v25)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(_a_F_ExecParallelHashJoin_4), int32(672), int32(_a_F_ExecParallelHashJoin_7))
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L306:
	;
	v1615 = int32(_a_F_ExecParallelHashJoin_0)
	v1616 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1]))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1609)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1618
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+24))
	v1624 = m.T0[v1623].(func(*base.Module, int32, int32, int32) int32)(m, v1608+int32(4), v1609, int32(0))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashJoin[1])) = v1616
	v1628 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1610)+4)))
	v1630 = v1628 & int32(_a_F_ExecParallelHashJoin_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v1610)+4)) = uint16(v1630)
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+12))
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1632)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1610)+6)) = uint16(v1633)
	v1636 = v1610
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	if l0 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v10 - int32(397) {
		case 0:
			v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+36)))
			if v138 != int32(1) {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			} else {
				v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v143 = int64(*(*int32)(unsafe.Add(mBase, uint32(v142)+40)))
				v145 = F_shm_toc_lookup(m, v141, v143, int32(0))
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = int32(697)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v145
					v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v255 = m.ExcPending
					if v255 != 0 {
						return int32(0)
					} else {
						return v254
					}
				}
			}
		default:
			v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
			mBase = m.M
			v255 = m.ExcPending
			if v255 != 0 {
				return int32(0)
			} else {
				return v254
			}
		case 6:
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+36)))
			if v14 != int32(1) {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
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
						v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
						mBase = m.M
						v255 = m.ExcPending
						if v255 != 0 {
							return int32(0)
						} else {
							return v254
						}
					}
				}
			}
		case 8:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+36)))
			if v29|v31&int32(1) == int32(0) {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
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
						v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
						mBase = m.M
						v255 = m.ExcPending
						if v255 != 0 {
							return int32(0)
						} else {
							return v254
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
									v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
									mBase = m.M
									v255 = m.ExcPending
									if v255 != 0 {
										return int32(0)
									} else {
										return v254
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
										v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
										mBase = m.M
										v255 = m.ExcPending
										if v255 != 0 {
											return int32(0)
										} else {
											return v254
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
									v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
									mBase = m.M
									v255 = m.ExcPending
									if v255 != 0 {
										return int32(0)
									} else {
										return v254
									}
								}
							}
						}
					}
				}
			}
		case 9:
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+36)))
			if v69|v71&int32(1) == int32(0) {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v70)+40)))
				v80 = F_shm_toc_lookup(m, v77, v78, int32(0))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					if v69 != 0 {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v80 + v82
					} else {
					}
					if v71&int32(1) == int32(0) {
						v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
						mBase = m.M
						v255 = m.ExcPending
						if v255 != 0 {
							return int32(0)
						} else {
							return v254
						}
					} else {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
						v95 = F_index_beginscan_parallel(m, v89, v90, l0+int32(160), v93, v94, v80)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v95
							v98 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v95)+28)) = uint8(v98)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
							if v100 != 0 {
								v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
								if v101 != int32(1) {
									v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
									mBase = m.M
									v255 = m.ExcPending
									if v255 != 0 {
										return int32(0)
									} else {
										return v254
									}
								} else {
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
									v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
									v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
									F_index_rescan(m, v104, v105, v106, v107, v108)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
										mBase = m.M
										v255 = m.ExcPending
										if v255 != 0 {
											return int32(0)
										} else {
											return v254
										}
									}
								}
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
								v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
								F_index_rescan(m, v104, v105, v106, v107, v108)
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
									mBase = m.M
									v255 = m.ExcPending
									if v255 != 0 {
										return int32(0)
									} else {
										return v254
									}
								}
							}
						}
					}
				}
			}
		case 10:
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v112 != 0 {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v115 = int64(*(*int32)(unsafe.Add(mBase, uint32(v114)+40)))
				v117 = F_shm_toc_lookup(m, v113, v115, int32(0))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v117
					v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v255 = m.ExcPending
					if v255 != 0 {
						return int32(0)
					} else {
						return v254
					}
				}
			} else {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			}
		case 11:
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+36)))
			if v168 != int32(1) {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			} else {
				v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v173 = int64(*(*int32)(unsafe.Add(mBase, uint32(v172)+40)))
				v175 = F_shm_toc_lookup(m, v171, v173, int32(0))
				mBase = m.M
				v176 = m.ExcPending
				if v176 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v175
					v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v178 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v175 + int32(24)
					} else {
					}
					v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v255 = m.ExcPending
					if v255 != 0 {
						return int32(0)
					} else {
						return v254
					}
				}
			}
		case 21:
			v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+36)))
			if v121 != int32(1) {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			} else {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+156))
				if v125 != 0 {
					v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v128 = int64(*(*int32)(unsafe.Add(mBase, uint32(v127)+40)))
					v130 = F_shm_toc_lookup(m, v126, v128, int32(0))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v133 = *(*int32)(unsafe.Add(mBase, uint32(v124)+156))
						m.T0[v133].(func(*base.Module, int32, int32, int32))(m, l0, v132, v130)
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return int32(0)
						} else {
							v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return int32(0)
							} else {
								return v254
							}
						}
					}
				} else {
					v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v255 = m.ExcPending
					if v255 != 0 {
						return int32(0)
					} else {
						return v254
					}
				}
			}
		case 22:
			v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+36)))
			if v151 != int32(1) {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			} else {
				v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+40))
				if v155 != 0 {
					v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v158 = int64(*(*int32)(unsafe.Add(mBase, uint32(v157)+40)))
					v160 = F_shm_toc_lookup(m, v156, v158, int32(0))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int32(0)
					} else {
						v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v163 = *(*int32)(unsafe.Add(mBase, uint32(v154)+40))
						m.T0[v163].(func(*base.Module, int32, int32, int32))(m, l0, v162, v160)
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return int32(0)
						} else {
							v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return int32(0)
							} else {
								return v254
							}
						}
					}
				} else {
					v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v255 = m.ExcPending
					if v255 != 0 {
						return int32(0)
					} else {
						return v254
					}
				}
			}
		case 26:
			v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+36)))
			if v183 != int32(1) {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			} else {
				v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v188 = int64(*(*int32)(unsafe.Add(mBase, uint32(v187)+40)))
				v190 = F_shm_toc_lookup(m, v186, v188, int32(0))
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_SharedFileSetAttach(m, v190+int32(168), v194)
					mBase = m.M
					v196 = m.ExcPending
					if v196 != 0 {
						return int32(0)
					} else {
						v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v197)+128)) = v190
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(632)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(720)
						v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
						mBase = m.M
						v255 = m.ExcPending
						if v255 != 0 {
							return int32(0)
						} else {
							return v254
						}
					}
				}
			}
		case 28:
			v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v245 = int64(*(*int32)(unsafe.Add(mBase, uint32(v244)+40)))
			v247 = F_shm_toc_lookup(m, v243, v245, int32(1))
			mBase = m.M
			v248 = m.ExcPending
			if v248 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v247
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			}
		case 29:
			v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v220 = int64(*(*int32)(unsafe.Add(mBase, uint32(v219)+40)))
			v222 = F_shm_toc_lookup(m, v218, v220, int32(1))
			mBase = m.M
			v223 = m.ExcPending
			if v223 != 0 {
				return int32(0)
			} else {
				v224 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)) = uint8(v224)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v222
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			}
		case 30:
			v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v229 = int64(*(*int32)(unsafe.Add(mBase, uint32(v228)+40)))
			v231 = F_shm_toc_lookup(m, v227, v229, int32(1))
			mBase = m.M
			v232 = m.ExcPending
			if v232 != 0 {
				return int32(0)
			} else {
				v233 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)) = uint8(v233)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v231
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			}
		case 32:
			v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v238 = int64(*(*int32)(unsafe.Add(mBase, uint32(v237)+40)))
			v240 = F_shm_toc_lookup(m, v236, v238, int32(1))
			mBase = m.M
			v241 = m.ExcPending
			if v241 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v240
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
				}
			}
		case 37:
			v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v203 != 0 {
				v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v206 = int64(*(*int32)(unsafe.Add(mBase, uint32(v205)+40)))
				v208 = F_shm_toc_lookup(m, v204, v206, int32(0))
				mBase = m.M
				v209 = m.ExcPending
				if v209 != 0 {
					return int32(0)
				} else {
					v211 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelInitializeWorker[0]))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v208 + v211*int32(20) + int32(4)
					v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
					mBase = m.M
					v255 = m.ExcPending
					if v255 != 0 {
						return int32(0)
					} else {
						return v254
					}
				}
			} else {
				v254 = F_planstate_tree_walker_impl(m, l0, int32(630), l1)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					return v254
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
	var v27 int32
	_ = v27
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
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelReportInstrumentation[0]))
	v70 = l1 + v59 + v61*v27*v63 + v67*v63
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
	v27 = int32(0)
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
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v27<<(uint(int32(2))%32))))
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
	v37 = v27 + int32(1)
	if v37 != v19 {
		v27 = v37
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
	F_errmsg_internal(m, int32(_a_F_ExecParallelReportInstrumentation_0), v10)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_ExecParallelReportInstrumentation_1), int32(1312), int32(_a_F_ExecParallelReportInstrumentation_2))
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v87 int32
	_ = v87
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v158 int32
	_ = v158
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
	var v170 int32
	_ = v170
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
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v339 int64
	_ = v339
	var v343 int64
	_ = v343
	var v347 int64
	_ = v347
	var v351 int64
	_ = v351
	var v354 int64
	_ = v354
	var v355 int64
	_ = v355
	var v358 int64
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v384 int64
	_ = v384
	var v393 int64
	_ = v393
	var v395 int64
	_ = v395
	var v399 int64
	_ = v399
	var v401 int64
	_ = v401
	var v405 int64
	_ = v405
	var v407 int64
	_ = v407
	var v411 int64
	_ = v411
	var v413 int64
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v435 int64
	_ = v435
	var v437 int64
	_ = v437
	var v439 int64
	_ = v439
	var v441 int64
	_ = v441
	var v443 int64
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
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
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[0]))
	v35 = v29 + v32<<(uint(int32(16))%32)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[1]))
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
	v46 = F_shm_toc_lookup(m, l1, int64(-2305843009213693946), int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v46 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v49 = v48
	goto L10
L9:
	;
	v49 = v3
	goto L10
L10:
	;
	v52 = F_shm_toc_lookup(m, l1, int64(-2305843009213693943), int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v56 = F_shm_toc_lookup(m, l1, int64(-2305843009213693944), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v60 = F_shm_toc_lookup(m, l1, int64(-2305843009213693950), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v62 = F_stringToNode(m, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v66 = F_shm_toc_lookup(m, l1, int64(-2305843009213693949), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v66
	v70 = v21 + int32(8)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v71 + int32(4)
	v80 = F_palloc(m, v72*int32(12)+int32(32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v80))) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v72
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = int32(815)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v80
	if v87 < v72 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v106 = v3
	goto L20
L18:
	;
	goto L19
L19:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[2]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	goto L24
L20:
	;
	v116 = v80 + int32(32) + v106*int32(12)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v121 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v120 + v121
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)) = uint16(v124)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v126 + int32(2)
	v132 = F_datumRestore(m, v70, v116+v121)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v132
	v136 = v106 + int32(1)
	if v136 != v72 {
		v106 = v136
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v160 = int32(0)
	v162 = F_CreateQueryDesc(m, v62, v56, v159, v160, v42, v80, v160, v49)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[3])) = v164
	F_pgstat_report_activity(m, int32(3), v164)
	mBase = m.M
	v170 = F_shm_toc_lookup(m, l1, int64(-2305843009213693945), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v172 = F_dsa_attach_in_place(m, v170, l0)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v174)+32)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	F_ExecutorStart(m, v162, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v162)+44))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+172)) = v172
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v183 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v162)+44))
	v258 = F_ExecParallelInitializeWorker(m, v255, v21+int32(8))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L37
	}
L30:
	;
	v186 = F_dsa_get_address(m, v172, v183)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v162)+40))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v186 + int32(4)
	if v189 <= int32(0) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v200 = int32(0)
	goto L33
L33:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v214 + int32(4)
	v219 = int32(8)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v188)+92))
	v224 = v221 + v215*int32(12)
	v227 = F_datumRestore(m, v21+v219, v224+v219)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L29
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+4)) = v227
	v233 = v200 + int32(1)
	if v233 != v189 {
		v200 = v233
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v162)+44))
	v264 = v261
	goto L40
L38:
	;
	base.MemoryCopy(m, int32(_a_F_ParallelQueryMain_0), int32(_a_F_ParallelQueryMain_1), int32(128))
	v339 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[5])) = v339
	v343 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[6]))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[7])) = v343
	v347 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[9])) = v347
	v351 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[10]))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[11])) = v351
	goto L65
L39:
	;
	goto L38
L40:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	switch v266 - int32(394) {
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
	*(*int64)(unsafe.Add(mBase, uint32(v264)+120)) = v260
	v325 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+117)) = uint8(v325)
	goto L39
L42:
	;
	goto L41
L43:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v264+v321)))
	v264 = v323
	goto L40
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v264)+112)) = v260
	v321 = int32(36)
	goto L43
L45:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v264)+32))
	if v316 == int32(0) {
		v321 = int32(116)
		goto L43
	} else {
		goto L64
	}
L46:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v264)+36))
	if v314 != 0 {
		v264 = v314
		goto L40
	} else {
		goto L63
	}
L47:
	;
	if v260 < int64(0) {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	if int64(0) <= v260 {
		goto L42
	} else {
		goto L59
	}
L49:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v264)+108))
	if v286 <= int32(0) {
		goto L39
	} else {
		goto L55
	}
L50:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264)+108))
	if v269 <= int32(0) {
		goto L39
	} else {
		goto L51
	}
L51:
	;
	v275 = int32(0)
	goto L52
L52:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v264)+104))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276+v275<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v260, v280)
	mBase = m.M
	v283 = v275 + int32(1)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v264)+108))
	if v283 < v284 {
		v275 = v283
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
	v292 = int32(0)
	goto L56
L56:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v264)+104))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293+v292<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, v260, v297)
	mBase = m.M
	v300 = v292 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v264)+108))
	if v300 < v301 {
		v292 = v300
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
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+117)) = uint8(v305)
	goto L38
L60:
	;
	v309 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+116)) = uint8(v309)
	goto L38
L61:
	;
	goto L62
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v264)+120)) = v260
	v312 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v264)+116)) = uint8(v312)
	goto L38
L63:
	;
	goto L39
L64:
	;
	goto L39
L65:
	;
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	v355 = int64(0)
	if v355 < v354 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v358 = v354
	goto L68
L67:
	;
	v358 = v355
	goto L68
L68:
	;
	F_ExecutorRun(m, v162, int32(1), v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_ExecutorFinish(m, v162)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v365 = F_shm_toc_lookup(m, l1, int64(-2305843009213693948), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v369 = F_shm_toc_lookup(m, l1, int64(-2305843009213693942), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[0]))
	v375 = v365 + v372<<(uint(int32(7))%32)
	v378 = v369 + v372<<(uint(int32(5))%32)
	base.MemoryFill(m, v375, int32(0), int32(128))
	F_BufferUsageAccumDiff(m, v375, int32(_a_F_ParallelQueryMain_0))
	mBase = m.M
	v384 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v378)+24)) = v384
	*(*int64)(unsafe.Add(mBase, uint32(v378)+16)) = v384
	*(*int64)(unsafe.Add(mBase, uint32(v378)+8)) = v384
	*(*int64)(unsafe.Add(mBase, uint32(v378))) = v384
	v393 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[6]))
	v395 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v378)+16)) = v393 - v395
	v399 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[10]))
	v401 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v378))) = v399 - v401
	v405 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[8]))
	v407 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v378)+8)) = v405 - v407
	v411 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[4]))
	v413 = *(*int64)(unsafe.Add(mBase, _c_F_ParallelQueryMain[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v378)+24)) = v411 - v413
	goto L73
L73:
	;
	if v46 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v162)+44))
	v417 = F_ExecParallelReportInstrumentation(m, v416, v46)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v162)+40))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+180))
	v421 = int32(0)
	if base.B2i32(v420 == v421)|base.B2i32(v52 == v421) == v421 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelQueryMain[0]))
	v432 = v52 + v429*int32(48)
	v433 = *(*int64)(unsafe.Add(mBase, uint32(v420)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v432)+48)) = v433
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v420)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v432)+40)) = v435
	v437 = *(*int64)(unsafe.Add(mBase, uint32(v420)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v432)+32)) = v437
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v420)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v432)+24)) = v439
	v441 = *(*int64)(unsafe.Add(mBase, uint32(v420)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v432)+16)) = v441
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v420)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v432)+8)) = v443
	goto L80
L79:
	;
	goto L80
L80:
	;
	F_ExecutorEnd(m, v162)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_dsa_detach(m, v172)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_FreeQueryDesc(m, v162)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	m.T0[v452].(func(*base.Module, int32))(m, v42)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
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
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v6 = int32(_a_F_ReinitializeParallelDSM_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[0]))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[0])) = v10
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
	*(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[0])) = v7
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
	v57 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v55))), uint32(v57))
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+16)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v55)+4)) = v60
	v64 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v55)+36)) = uint16(v64)
	*(*int64)(unsafe.Add(mBase, uint32(v55)+24)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v55)+32)) = int32(_a_F_ReinitializeParallelDSM_1)
	goto L15
L14:
	;
	goto L9
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ReinitializeParallelDSM[2]))
	F_shm_mq_set_receiver(m, v55, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v78 = F_shm_mq_attach(m, v55, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v80+v49<<(uint(int32(3))%32))+4)) = v78
	v86 = v49 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v86 < v87 {
		v49 = v86
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
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
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
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
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
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v547 {
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
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v432 != 0 {
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
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[1]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[2])) = v117
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[1])) = v117
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0])) = v123 + int32(40)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[3])) = v123 + int32(36)
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
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v177 != 0 {
		goto L52
	} else {
		goto L53
	}
L37:
	;
	F_errfinish(m, int32(_a_F_parallel_vacuum_process_all_indexes_0), v171, int32(_a_F_parallel_vacuum_process_all_indexes_1))
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
	v148 = int32(_a_F_parallel_vacuum_process_all_indexes_2)
	goto L44
L43:
	;
	v148 = int32(_a_F_parallel_vacuum_process_all_indexes_3)
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
	v164 = int32(_a_F_parallel_vacuum_process_all_indexes_4)
	goto L49
L48:
	;
	v164 = int32(_a_F_parallel_vacuum_process_all_indexes_5)
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
	v180 = base.AtomicRmwAdd32(m, v177, int32(0), int32(1))
	goto L54
L53:
	;
	goto L54
L54:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v181 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v186 = v181
	v188 = int32(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v221 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v196 = v193 + v188*int32(48)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)))
	if v197 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L57
L60:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v188<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v204, v196)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L26
	} else {
		goto L63
	}
L61:
	;
	v208 = v186
	goto L62
L62:
	;
	v210 = v188 + int32(1)
	if v210 < v208 {
		v186 = v208
		v188 = v210
		goto L58
	} else {
		goto L64
	}
L63:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v208 = v207
	goto L62
L64:
	;
	goto L59
L65:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v238 = base.AtomicRmwAdd32(m, v235, int32(44), int32(1))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v238 < v239 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v225 = int32(0)
	v226 = base.AtomicRmwSub32(m, v221, v225, int32(1))
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v228 == v225 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v233 = base.AtomicRmwAdd32(m, v228, int32(0), int32(1))
	goto L65
L68:
	;
	v244 = v238
	goto L71
L69:
	;
	goto L70
L70:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v278 != 0 {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v252 = v249 + v244*int32(48)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+4)))
	if v253 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256+v244<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v260, v252)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L26
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v266 = base.AtomicRmwAdd32(m, v263, int32(44), int32(1))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v266 < v267 {
		v244 = v266
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
	v281 = base.AtomicRmwSub32(m, v278, int32(0), int32(1))
	goto L80
L79:
	;
	goto L80
L80:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_WaitForParallelWorkersToFinish(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L26
	} else {
		goto L81
	}
L81:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+20))
	if v286 <= int32(0) {
		goto L19
	} else {
		goto L82
	}
L82:
	;
	v293 = int32(0)
	goto L83
L83:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v301 = v298 + v293<<(uint(int32(7))%32)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v305 = v302 + v293<<(uint(int32(5))%32)
	v306 = int32(_a_F_parallel_vacuum_process_all_indexes_6)
	v308 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[4]))
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v301)))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[4])) = v308 + v309
	v312 = int32(_a_F_parallel_vacuum_process_all_indexes_7)
	v314 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[5]))
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v301)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[5])) = v314 + v315
	v318 = int32(_a_F_parallel_vacuum_process_all_indexes_8)
	v320 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[6]))
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v301)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[6])) = v320 + v321
	v324 = int32(_a_F_parallel_vacuum_process_all_indexes_9)
	v326 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[7]))
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v301)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[7])) = v326 + v327
	v330 = int32(_a_F_parallel_vacuum_process_all_indexes_10)
	v332 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[8]))
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v301)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[8])) = v332 + v333
	v336 = int32(_a_F_parallel_vacuum_process_all_indexes_11)
	v338 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[9]))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v301)+40))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[9])) = v338 + v339
	v342 = int32(_a_F_parallel_vacuum_process_all_indexes_12)
	v344 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[10]))
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v301)+48))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[10])) = v344 + v345
	v348 = int32(_a_F_parallel_vacuum_process_all_indexes_13)
	v350 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[11]))
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v301)+56))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[11])) = v350 + v351
	v354 = int32(_a_F_parallel_vacuum_process_all_indexes_14)
	v356 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[12]))
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v301)+64))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[12])) = v356 + v357
	v360 = int32(_a_F_parallel_vacuum_process_all_indexes_15)
	v362 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[13]))
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v301)+72))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[13])) = v362 + v363
	v366 = int32(_a_F_parallel_vacuum_process_all_indexes_16)
	v368 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[14]))
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v301)+80))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[14])) = v368 + v369
	v372 = int32(_a_F_parallel_vacuum_process_all_indexes_17)
	v374 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[15]))
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v301)+88))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[15])) = v374 + v375
	v378 = int32(_a_F_parallel_vacuum_process_all_indexes_18)
	v380 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[16]))
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v301)+96))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[16])) = v380 + v381
	v384 = int32(_a_F_parallel_vacuum_process_all_indexes_19)
	v386 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[17]))
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v301)+104))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[17])) = v386 + v387
	v390 = int32(_a_F_parallel_vacuum_process_all_indexes_20)
	v392 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[18]))
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v301)+112))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[18])) = v392 + v393
	v396 = int32(_a_F_parallel_vacuum_process_all_indexes_21)
	v398 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[19]))
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v301)+120))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[19])) = v398 + v399
	v402 = int32(_a_F_parallel_vacuum_process_all_indexes_22)
	v404 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[20]))
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v305)+16))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[20])) = v404 + v405
	v408 = int32(_a_F_parallel_vacuum_process_all_indexes_23)
	v410 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[21]))
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v305)))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[21])) = v410 + v411
	v414 = int32(_a_F_parallel_vacuum_process_all_indexes_24)
	v416 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[22]))
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v305)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[22])) = v416 + v417
	v420 = int32(_a_F_parallel_vacuum_process_all_indexes_25)
	v422 = *(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[23]))
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v305)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[23])) = v422 + v423
	goto L85
L84:
	;
	goto L19
L85:
	;
	v427 = v293 + int32(1)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+20))
	if v427 < v429 {
		v293 = v427
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v435 = base.AtomicRmwAdd32(m, v432, int32(0), int32(1))
	goto L89
L88:
	;
	goto L89
L89:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v436 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v441 = v436
	v443 = int32(0)
	goto L93
L91:
	;
	goto L92
L92:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v476 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L93:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v451 = v448 + v443*int32(48)
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+4)))
	if v452 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L92
L95:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v455+v443<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v459, v451)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L26
	} else {
		goto L98
	}
L96:
	;
	v463 = v441
	goto L97
L97:
	;
	v465 = v443 + int32(1)
	if v465 < v463 {
		v441 = v463
		v443 = v465
		goto L93
	} else {
		goto L99
	}
L98:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v463 = v462
	goto L97
L99:
	;
	goto L94
L100:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v493 = base.AtomicRmwAdd32(m, v490, int32(44), int32(1))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v493 < v494 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v480 = int32(0)
	v481 = base.AtomicRmwSub32(m, v476, v480, int32(1))
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v483 == v480 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v488 = base.AtomicRmwAdd32(m, v483, int32(0), int32(1))
	goto L100
L103:
	;
	v499 = v493
	goto L106
L104:
	;
	goto L105
L105:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0]))
	if v533 == int32(0) {
		goto L19
	} else {
		goto L113
	}
L106:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v507 = v504 + v499*int32(48)
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+4)))
	if v508 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L105
L108:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v511+v499<<(uint(int32(2))%32))))
	F_parallel_vacuum_process_one_index(m, l0, v515, v507)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L26
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v521 = base.AtomicRmwAdd32(m, v518, int32(44), int32(1))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v521 < v522 {
		v499 = v521
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
	v538 = base.AtomicRmwSub32(m, v533, int32(0), int32(1))
	goto L19
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L26
	} else {
		goto L125
	}
L115:
	;
	v554 = int32(0)
	goto L118
L116:
	;
	goto L117
L117:
	;
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[3]))
	if v581 != 0 {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v562 = v559 + v554*int32(48)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	if v563 != int32(3) {
		goto L114
	} else {
		goto L120
	}
L119:
	;
	goto L117
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = int32(0)
	v569 = v554 + int32(1)
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v569 < v570 {
		v554 = v569
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v581)))
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[1])) = v583
	v586 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[3])) = v586
	*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_process_all_indexes[0])) = v586
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
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v598+v554<<(uint(int32(2))%32))))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v603 + int32(4)
	F_errmsg_internal(m, int32(_a_F_parallel_vacuum_process_all_indexes_26), v11)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L26
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_parallel_vacuum_process_all_indexes_0), int32(753), int32(_a_F_parallel_vacuum_process_all_indexes_1))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
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
