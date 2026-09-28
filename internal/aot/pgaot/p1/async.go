package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AsyncReadBuffers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int64
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v593 int64
	_ = v593
	var v597 int32
	_ = v597
	var v599 int64
	_ = v599
	var v609 int32
	_ = v609
	var v613 int64
	_ = v613
	var v617 int64
	_ = v617
	var v619 int32
	_ = v619
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int64
	_ = v652
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v827 int64
	_ = v827
	var v829 int64
	_ = v829
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v885 int32
	_ = v885
	var v888 int64
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v905 int64
	_ = v905
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v919 int32
	_ = v919
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v959 int64
	_ = v959
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1016 int64
	_ = v1016
	var v1017 int64
	_ = v1017
	var v1021 int64
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1203 int32
	_ = v1203
	var v1229 int32
	_ = v1229
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int64
	_ = v1247
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1329 int32
	_ = v1329
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
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
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1419 int32
	_ = v1419
	var v1424 int32
	_ = v1424
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1536 int32
	_ = v1536
	var v1541 int32
	_ = v1541
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1635 int32
	_ = v1635
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1673 int32
	_ = v1673
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1689 int64
	_ = v1689
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1701 int64
	_ = v1701
	var v1702 int64
	_ = v1702
	var v1706 int64
	_ = v1706
	var v1732 int32
	_ = v1732
	var v1734 int64
	_ = v1734
	var v1736 int64
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1741 int64
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1746 int64
	_ = v1746
	var v1753 int32
	_ = v1753
	var v1757 int64
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1768 int32
	_ = v1768
	var v1773 int64
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1790 int32
	_ = v1790
	var v1794 int64
	_ = v1794
	var v1798 int64
	_ = v1798
	var v1803 int32
	_ = v1803
	var v1811 int64
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1816 int64
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1821 int64
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1833 int32
	_ = v1833
	var v1842 int32
	_ = v1842
	v33 = m.G0
	v35 = v33 - int32(512)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v48 = base.B2i32(v46 == int32(116))
	if v48 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v53 = F_IOContextForStrategy(m, v52)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v57 = int32(3)
	v58 = int32(1)
	goto L3
L3:
	;
	v59 = v38<<(uint(int32(2))%32) + v41
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[0])))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[1])))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	F_pgstat_prepare_report_checksum_failure(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v57 = v53
	v58 = int32(0)
	goto L3
L6:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[2]))
	v71 = l0 + int32(48)
	v72 = F_pgaio_io_acquire_nb(m, v69, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v72 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	v494 = v72
	goto L10
L10:
	;
	v524 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v524)
	v527 = l0 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v527))) = int32(-1)
	goto L85
L11:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[2]))
	v80 = m.G0
	v82 = v80 - int32(80)
	m.G0 = v82
	v84 = F_pgaio_io_acquire_nb(m, v79, v71)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L15
	}
L12:
	;
	v494 = v422
	goto L10
L13:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L82
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L78
	}
L15:
	;
	if v84 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	goto L19
L17:
	;
	v422 = v84
	goto L18
L18:
	;
	m.G0 = v82 + int32(80)
	goto L12
L19:
	;
	v124 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	v422 = v416
	goto L18
L21:
	;
	if v124 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_errhidestmt(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[4]))
	if int32(0) < v154 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	F_errhidecontext(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)+160))
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+64)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v82)+68)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v82)+72)) = v132
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_0), v82-int32(-64))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(768), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	v416 = F_pgaio_io_acquire_nb(m, v79, v71)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L76
	}
L30:
	;
	v157 = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v163 = v157
	v165 = v154
	v166 = v152
	v167 = v157
	v172 = v159
	goto L33
L31:
	;
	v232 = v152
	goto L32
L32:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232)+22)))
	if v259 != 0 {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v172)+24))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v195 = int32(7)
	v200 = v193 + v194<<(uint(v195)%32) + v163<<(uint(v195)%32)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v201 == int32(6) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if int32(0) < v220 {
		goto L29
	} else {
		goto L40
	}
L35:
	;
	v204 = int32(0)
	v207 = base.AtomicRmwOr32(m, v204, int32(_a_F_AsyncReadBuffers_3), v204)
	F_pgaio_io_reclaim(m, v200)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	v218 = v165
	v219 = v166
	v220 = v167
	v221 = v172
	goto L37
L37:
	;
	v223 = v163 + int32(1)
	if v223 < v218 {
		v163 = v223
		v165 = v218
		v166 = v219
		v167 = v220
		v172 = v221
		goto L33
	} else {
		goto L39
	}
L38:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[4]))
	v218 = v217
	v219 = v211
	v220 = v167 + int32(1)
	v221 = v213
	goto L37
L39:
	;
	goto L34
L40:
	;
	v232 = v219
	goto L32
L41:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	v264 = v232
	goto L43
L43:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	if v265 != 0 {
		goto L29
	} else {
		goto L45
	}
L44:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v264 = v263
	goto L43
L45:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)+160))
	if v266 == int32(0) {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264)+156))
	v271 = v269 - int32(24)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if base.Ui32(int32(7)) < base.Ui32(v272) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	if v381 == int32(0) {
		goto L13
	} else {
		goto L75
	}
L48:
	;
	if int32(1)<<(uint(v272)%32)&int32(48) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v367 = int32(0)
	v370 = base.AtomicRmwOr32(m, v367, int32(_a_F_AsyncReadBuffers_3), v367)
	F_pgaio_io_reclaim(m, v271)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L74
	}
L50:
	;
	if v272 == int32(6) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v269)+24))
	v309 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L4
	} else {
		goto L57
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+24))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = (v271 - v289) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_4), v82+int32(16))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(840), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	if v309 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_errhidestmt(m)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	F_pgaio_io_wait(m, v271, v306)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L73
	}
L61:
	;
	F_errhidecontext(m)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+24))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+2)))
	if base.Ui32(v322) <= base.Ui32(int32(2)) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v330<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[6])))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+8))
	goto L67
L64:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v322<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[7])))
	v329 = v327
	goto L66
L65:
	;
	v329 = int32(0)
	goto L66
L66:
	;
	goto L63
L67:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if base.Ui32(v335) <= base.Ui32(int32(7)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v335<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[8])))
	v341 = v340
	goto L70
L69:
	;
	v341 = int32(0)
	goto L70
L70:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v82+int32(48)))) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v82)+32)) = (v271 - v318) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v82)+40)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v82)+44)) = v341
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_5), v82+int32(32))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(847), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	goto L60
L73:
	;
	goto L47
L74:
	;
	goto L47
L75:
	;
	goto L29
L76:
	;
	if v416 == int32(0) {
		goto L19
	} else {
		goto L77
	}
L77:
	;
	goto L20
L78:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_6), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v464)+160))
	v467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v464)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v465
	F_errdetail_internal(m, int32(_a_F_AsyncReadBuffers_7), v82)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(818), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_8), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(881), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v530 < int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	m.G0 = v35 + int32(512)
	return v1842
L87:
	;
	if v554 != int32(2) {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[9]))
	v541 = F_StartLocalBufferIO(m, v534+(v530^int32(-1))*int32(56), int32(1), v527)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[10]))
	v545 = int32(56)
	v550 = int32(1)
	v552 = F_StartSharedBufferIO(m, v544+v530*v545-v545, v550, v550, v527)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L4
	} else {
		goto L92
	}
L91:
	;
	v554 = v541
	goto L87
L92:
	;
	v554 = v552
	goto L87
L93:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)+16))
	if v559 == v494 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	v658 = v38 + v37
	v662 = int32(base.Ui32(v43)>>(uint(int32(3))%32)) & int32(1)
	if v46 == int32(116) {
		goto L121
	} else {
		goto L122
	}
L96:
	;
	v578 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v578
	if v554 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v558)+16)) = int32(0)
	F_pgaio_io_reclaim(m, v494)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L101
	}
L100:
	;
	goto L96
L101:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_9), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(258), int32(_a_F_AsyncReadBuffers_10))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)))
	v585 = v583 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v585)
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v588 == int32(116) {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	goto L106
L106:
	;
	v656 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)) = uint8(v656)
	v1842 = v578
	goto L86
L107:
	;
	v609 = v58*int32(320) + v57<<(uint(int32(6))%32)
	v613 = *(*int64)(unsafe.Add(mBase, uint32(v609)+uint32(_c_F_AsyncReadBuffers[11])))
	*(*int64)(unsafe.Add(mBase, uint32(v609)+uint32(_c_F_AsyncReadBuffers[11]))) = v613 + int64(1)
	v617 = *(*int64)(unsafe.Add(mBase, uint32(v609)+uint32(_c_F_AsyncReadBuffers[12])))
	*(*int64)(unsafe.Add(mBase, uint32(v609)+uint32(_c_F_AsyncReadBuffers[12]))) = v617
	v619 = int32(1)
	F_pgstat_count_backend_io_op(m, v58, v57, int32(2), v619, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[13])) = uint8(v619)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[14])) = uint8(v619)
	goto L111
L108:
	;
	v591 = int32(_a_F_AsyncReadBuffers_11)
	v593 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[15]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[15])) = v593 + int64(1)
	goto L107
L109:
	;
	goto L110
L110:
	;
	v597 = int32(_a_F_AsyncReadBuffers_12)
	v599 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[16]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[16])) = v599 + int64(1)
	goto L107
L111:
	;
	v628 = int32(0)
	v630 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[17])))
	if v630 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v633 = int32(_a_F_AsyncReadBuffers_13)
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18]))
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18])) = v635 + v637
	goto L114
L113:
	;
	goto L114
L114:
	;
	if v587 == int32(0) {
		v1842 = v628
		goto L86
	} else {
		goto L115
	}
L115:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v587)+272))
	if v642 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587)+268)))
	if v645 != int32(1) {
		v1842 = v628
		goto L86
	} else {
		goto L119
	}
L117:
	;
	v651 = v642
	goto L118
L118:
	;
	v652 = *(*int64)(unsafe.Add(mBase, uint32(v651)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v651)+120)) = v652 + int64(1)
	v1842 = v628
	goto L86
L119:
	;
	F_pgstat_assoc_relation(m, v587)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v587)+272))
	v651 = v650
	goto L118
L121:
	;
	v665 = v662 | int32(2)
	goto L123
L122:
	;
	v665 = v662
	goto L123
L123:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v670 < int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v688
	v690 = int32(1)
	v692 = v38 + v690
	v693 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	if v693 <= v692 {
		v791 = v690
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[20]))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v674+(v670^int32(-1))<<(uint(int32(2))%32))))
	v688 = v680
	goto L124
L126:
	;
	goto L127
L127:
	;
	v682 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[21]))
	v688 = v682 + v670<<(uint(int32(13))%32) + int32(-8192)
	goto L124
L128:
	;
	v821 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v527))) = (v494 - v822) >> (uint(int32(7)) % 32)
	v827 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v494)+52)))
	*(*uint32)(unsafe.Add(mBase, uint32(v527)+4)) = uint32(v827)
	v829 = *(*int64)(unsafe.Add(mBase, uint32(v494)+48))
	*(*uint32)(unsafe.Add(mBase, uint32(v527)+8)) = uint32(v829)
	goto L144
L129:
	;
	v698 = v690
	v699 = v692
	goto L130
L130:
	;
	v729 = v41 + v699<<(uint(int32(2))%32)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if v730 < int32(0) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v791 = v783
	goto L128
L132:
	;
	if v756 != int32(2) {
		v791 = v698
		goto L128
	} else {
		goto L138
	}
L133:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[9]))
	v740 = int32(0)
	v742 = F_StartLocalBufferIO(m, v734+(v730^int32(-1))*int32(56), v740, v740)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L4
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[10]))
	v746 = int32(56)
	v752 = int32(0)
	v754 = F_StartSharedBufferIO(m, v745+v730*v746-v746, int32(1), v752, v752)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L4
	} else {
		goto L137
	}
L136:
	;
	v756 = v742
	goto L132
L137:
	;
	v756 = v754
	goto L132
L138:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	if v762 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35+v698<<(uint(int32(2))%32)))) = v780
	v782 = int32(1)
	v783 = v698 + v782
	v785 = v699 + v782
	v786 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+30)))
	if v785 < v786 {
		v698 = v783
		v699 = v785
		goto L130
	} else {
		goto L143
	}
L140:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[20]))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v766+(v762^int32(-1))<<(uint(int32(2))%32))))
	v780 = v772
	goto L139
L141:
	;
	goto L142
L142:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[21]))
	v780 = v774 + v762<<(uint(int32(13))%32) + int32(-8192)
	goto L139
L143:
	;
	goto L131
L144:
	;
	v831 = int32(0)
	v834 = v791 & int32(255)
	if v834 == v831 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v494)+13)) = uint8(v834)
	if v46 == int32(116) {
		goto L154
	} else {
		goto L155
	}
L146:
	;
	if v834 != int32(1) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v847 = v831
	v848 = v831
	goto L150
L148:
	;
	v919 = v831
	goto L149
L149:
	;
	v947 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v947)+16))
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v494)+76))
	v950 = int32(3)
	v959 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v59+v919<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v948+v949<<(uint(v950)%32)+v919<<(uint(v950)%32)))) = v959
	goto L145
L150:
	;
	v875 = int32(_a_F_AsyncReadBuffers_14)
	v876 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v876)+16))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v494)+76))
	v879 = int32(3)
	v885 = int32(2)
	v888 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v59+v848<<(uint(v885)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v877+v878<<(uint(v879)%32)+v848<<(uint(v879)%32)))) = v888
	v891 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v891)+16))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v494)+76))
	v898 = v848 | int32(1)
	v905 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v59+v898<<(uint(v885)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v892+v893<<(uint(v879)%32)+v898<<(uint(v879)%32)))) = v905
	v908 = v848 + v885
	v910 = v847 + v885
	if v910 != v834&int32(254) {
		v847 = v910
		v848 = v908
		goto L150
	} else {
		goto L152
	}
L151:
	;
	if v834&int32(1) == int32(0) {
		goto L145
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	v919 = v908
	goto L149
L154:
	;
	v998 = int32(3)
	goto L156
L155:
	;
	v998 = int32(2)
	goto L156
L156:
	;
	F_pgaio_io_register_callbacks(m, v494, v998, (v43|v61|v63<<(uint(int32(2))%32))&int32(255))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+3)))
	v1004 = v1003 | v665
	*(*uint8)(unsafe.Add(mBase, uint32(v494)+3)) = uint8(v1004)
	goto L158
L158:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[22])))
	v1010 = m.G0
	v1012 = v1010 - int32(16)
	m.G0 = v1012
	if v1007 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1026 = int32(_a_F_AsyncReadBuffers_15)
	v1028 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1028 + int32(1)
	v1032 = int32(0)
	v1033 = m.G0
	v1035 = v1033 - int32(16)
	m.G0 = v1035
	v1039 = F__mdfd_getseg(m, v1025, v42, v658, v1032, int32(9))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L4
	} else {
		goto L163
	}
L160:
	;
	F___clock_gettime(m, int32(1), v1012)
	mBase = m.M
	v1016 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1012)+8)))
	v1017 = *(*int64)(unsafe.Add(mBase, uint32(v1012)))
	v1021 = v1016 + v1017*int64(1000000000)
	goto L162
L161:
	;
	v1021 = int64(0)
	goto L162
L162:
	;
	m.G0 = v1012 + int32(16)
	goto L159
L163:
	;
	v1043 = v658 & int32(_a_F_AsyncReadBuffers_16)
	v1044 = int32(_a_F_AsyncReadBuffers_17) - v1043
	if base.Ui32(v791) <= base.Ui32(v1044) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v1679 = int32(_a_F_AsyncReadBuffers_15)
	v1681 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	v1682 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1681 - v1682
	v1689 = base.I64_extend_i32_s(v791 << (uint(int32(13)) % 32))
	v1693 = m.G0
	v1695 = v1693 - int32(16)
	m.G0 = v1695
	if v1021 != int64(0) {
		goto L276
	} else {
		goto L277
	}
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L4
	} else {
		goto L270
	}
L166:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+12))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v494)+76))
	v1052 = v1048 + v1049<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1035)+12)) = v1052
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v1052)+4)) = int32(_a_F_AsyncReadBuffers_18)
	*(*int32)(unsafe.Add(mBase, uint32(v1052))) = v1054
	v1059 = int32(1)
	if base.Ui32(v791) < base.Ui32(v1044) {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	goto L168
L168:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L4
	} else {
		goto L267
	}
L169:
	;
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[24])))
	if v1229&int32(1) == int32(0) {
		goto L192
	} else {
		goto L193
	}
L170:
	;
	v1061 = v791
	goto L172
L171:
	;
	v1061 = v1044
	goto L172
L172:
	;
	if base.Ui32(v1061) < base.Ui32(int32(2)) {
		v1203 = v1059
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v1064 = int32(1)
	if v1061 != int32(2) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v1067 = int32(1)
	v1068 = v1061 - v1067
	v1073 = v1054
	v1078 = v1064
	v1079 = v1052
	v1080 = v1059
	v1083 = v1032
	goto L177
L175:
	;
	v1149 = v1054
	v1154 = v1064
	v1155 = v1052
	v1156 = v1059
	goto L176
L176:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v35+v1154<<(uint(int32(2))%32))))
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1155)+4))
	if v1184 != v1149+v1185 {
		goto L189
	} else {
		goto L190
	}
L177:
	;
	v1107 = v35 + v1078<<(uint(int32(2))%32)
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1107)))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+4))
	if v1108 == v1073+v1109 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	if v1068&v1067 == int32(0) {
		v1203 = v1141
		goto L169
	} else {
		goto L188
	}
L179:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+4))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+4))
	if v1125 != v1122+v1126 {
		goto L184
	} else {
		goto L185
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+4)) = v1109 - int32(-8192)
	v1122 = v1073
	v1123 = v1079
	v1124 = v1080
	goto L179
L181:
	;
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+12)) = int32(_a_F_AsyncReadBuffers_18)
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+8)) = v1108
	v1122 = v1108
	v1123 = v1079 + int32(8)
	v1124 = v1080 + int32(1)
	goto L179
L183:
	;
	v1142 = int32(2)
	v1143 = v1078 + v1142
	v1145 = v1083 + v1142
	if v1145 != v1068&int32(-2) {
		v1073 = v1139
		v1078 = v1143
		v1079 = v1140
		v1080 = v1141
		v1083 = v1145
		goto L177
	} else {
		goto L187
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1123)+12)) = int32(_a_F_AsyncReadBuffers_18)
	*(*int32)(unsafe.Add(mBase, uint32(v1123)+8)) = v1125
	v1139 = v1125
	v1140 = v1123 + int32(8)
	v1141 = v1124 + int32(1)
	goto L183
L185:
	;
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1123)+4)) = v1126 - int32(-8192)
	v1139 = v1122
	v1140 = v1123
	v1141 = v1124
	goto L183
L187:
	;
	goto L178
L188:
	;
	v1149 = v1139
	v1154 = v1143
	v1155 = v1140
	v1156 = v1141
	goto L176
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1155)+12)) = int32(_a_F_AsyncReadBuffers_18)
	*(*int32)(unsafe.Add(mBase, uint32(v1155)+8)) = v1184
	v1203 = v1156 + int32(1)
	goto L169
L190:
	;
	goto L191
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1155)+4)) = v1185 - int32(-8192)
	v1203 = v1156
	goto L169
L192:
	;
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+3)))
	v1236 = v1235 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v494)+3)) = uint8(v1236)
	goto L195
L193:
	;
	goto L194
L194:
	;
	v1242 = v494 + int32(104)
	goto L196
L195:
	;
	goto L194
L196:
	;
	v1243 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v494)+1)) = uint8(v1243)
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1242)+8)) = v1245
	v1247 = *(*int64)(unsafe.Add(mBase, uint32(v1025)))
	*(*int64)(unsafe.Add(mBase, uint32(v1242))) = v1247
	*(*int32)(unsafe.Add(mBase, uint32(v1242)+16)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v1242)+12)) = v658
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242)+21)))
	v1256 = v42&int32(255) | v1253<<(uint(int32(8))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1242)+20)) = uint16(v1256)
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+12))
	if v1262 != int32(-1) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1265 = int32(256)
	goto L199
L198:
	;
	v1265 = int32(0)
	goto L199
L199:
	;
	v1266 = v1256&int32(_a_F_AsyncReadBuffers_19) | v1265
	*(*uint16)(unsafe.Add(mBase, uint32(v1242)+20)) = uint16(v1266)
	v1269 = v1266 & int32(_a_F_AsyncReadBuffers_20)
	*(*uint16)(unsafe.Add(mBase, uint32(v1242)+20)) = uint16(v1269)
	F_pgaio_io_register_callbacks(m, v494, int32(1), int32(0))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1039)))
	v1276 = F_FileAccess(m, v1275)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L4
	} else {
		goto L201
	}
L201:
	;
	if v1276 < int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1635 = int32(-1)
	goto L204
L203:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[25]))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1282+v1275*int32(48))))
	*(*int64)(unsafe.Add(mBase, uint32(v494)+96)) = base.I64_extend_i32_u(v1043 << (uint(int32(13)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v494)+88)) = v1286
	*(*uint16)(unsafe.Add(mBase, uint32(v494)+92)) = uint16(v1203)
	v1290 = m.G0
	v1292 = v1290 - int32(32)
	m.G0 = v1292
	v1294 = int32(1)
	v1295 = int32(_a_F_AsyncReadBuffers_15)
	v1297 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1297 + v1294
	*(*int32)(unsafe.Add(mBase, uint32(v494)+20)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v494)+2)) = uint8(v1294)
	F_pgaio_io_update_state(m, v494, int32(2))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L4
	} else {
		goto L205
	}
L204:
	;
	if v1635 != 0 {
		goto L165
	} else {
		goto L266
	}
L205:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v1309)+16)) = int32(0)
	v1312 = m.G0
	v1314 = v1312 - int32(32)
	m.G0 = v1314
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+4)))
	if v1316 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1329 = v1316
	goto L209
L207:
	;
	goto L208
L208:
	;
	m.G0 = v1314 + int32(32)
	F_pgaio_io_update_state(m, v494, int32(3))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L4
	} else {
		goto L234
	}
L209:
	;
	v1362 = v1329 - int32(1)
	v1364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494+int32(5)+v1362))))
	v1366 = v1364 << (uint(int32(3)) % 32)
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1366)+uint32(_c_F_AsyncReadBuffers[26])))
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	if v1368 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	goto L208
L211:
	;
	v1370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1362+(v494+int32(9))))))
	v1373 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L4
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1329) {
		v1329 = v1362
		goto L209
	} else {
		goto L233
	}
L214:
	;
	if v1373 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	F_errhidestmt(m)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L4
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	m.T0[v1429].(func(*base.Module, int32, int32))(m, v494, v1370)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L4
	} else {
		goto L232
	}
L218:
	;
	F_errhidecontext(m)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L4
	} else {
		goto L219
	}
L219:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+24))
	goto L220
L220:
	;
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+2)))
	if base.Ui32(v1385) <= base.Ui32(int32(2)) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+1)))
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1393<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[6])))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+8))
	goto L225
L222:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1385<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[7])))
	v1392 = v1390
	goto L224
L223:
	;
	v1392 = int32(0)
	goto L224
L224:
	;
	goto L221
L225:
	;
	v1398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	if base.Ui32(v1398) <= base.Ui32(int32(7)) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1314+int32(28)))) = v1370
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1366)+uint32(_c_F_AsyncReadBuffers[27])))
	*(*int32)(unsafe.Add(mBase, uint32(v1314+int32(24)))) = v1409
	*(*int32)(unsafe.Add(mBase, uint32(v1314+int32(20)))) = v1364
	*(*int32)(unsafe.Add(mBase, uint32(v1314+int32(16)))) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v1314)+12)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v1314)+8)) = v1397
	*(*int32)(unsafe.Add(mBase, uint32(v1314)+4)) = v1392
	*(*int32)(unsafe.Add(mBase, uint32(v1314))) = (v494 - v1381) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_21), v1314)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L4
	} else {
		goto L230
	}
L227:
	;
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1398<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[8])))
	v1405 = v1403
	goto L229
L228:
	;
	v1405 = int32(0)
	goto L229
L229:
	;
	goto L226
L230:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_22), int32(215), int32(_a_F_AsyncReadBuffers_23))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L4
	} else {
		goto L231
	}
L231:
	;
	goto L217
L232:
	;
	goto L213
L233:
	;
	goto L210
L234:
	;
	v1477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+3)))
	if v1477&int32(1) != 0 {
		v1489 = v1294
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1492 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L4
	} else {
		goto L239
	}
L236:
	;
	v1480 = int32(0)
	v1482 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[28]))
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1482)+28))
	if v1483 == v1480 {
		v1489 = v1480
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v1486 = m.T0[v1483].(func(*base.Module, int32) int32)(m, v494)
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L4
	} else {
		goto L238
	}
L238:
	;
	v1489 = v1486
	goto L235
L239:
	;
	if v1492 != 0 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	F_errhidestmt(m)
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L4
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	if v1489 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L243:
	;
	F_errhidecontext(m)
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L4
	} else {
		goto L244
	}
L244:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1499)+24))
	v1504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+2)))
	if base.Ui32(v1504) <= base.Ui32(int32(2)) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v1512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+1)))
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1512<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[6])))
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1515)+8))
	goto L249
L246:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1504<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[7])))
	v1511 = v1509
	goto L248
L247:
	;
	v1511 = int32(0)
	goto L248
L248:
	;
	goto L245
L249:
	;
	v1518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	if base.Ui32(v1518) <= base.Ui32(int32(7)) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1518<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[8])))
	v1524 = v1523
	goto L252
L251:
	;
	v1524 = int32(0)
	goto L252
L252:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1526)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+16)) = v1489
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+20)) = v1527
	*(*int32)(unsafe.Add(mBase, uint32(v1292))) = (v494 - v1500) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+4)) = v1511
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+8)) = v1516
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+12)) = v1524
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_24), v1292)
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L4
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(459), int32(_a_F_AsyncReadBuffers_25))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	goto L242
L255:
	;
	v1593 = int32(_a_F_AsyncReadBuffers_15)
	v1595 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1595 - int32(1)
	m.G0 = v1292 + int32(32)
	v1635 = int32(0)
	goto L204
L256:
	;
	v1549 = int32(_a_F_AsyncReadBuffers_26)
	v1550 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1551 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1550)+22)))
	v1553 = v1551 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1550)+22)) = uint16(v1553)
	*(*int32)(unsafe.Add(mBase, uint32(v1550+v1551<<(uint(int32(2))%32))+24)) = v494
	v1560 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1560)+20)))
	if v1561 != 0 {
		goto L255
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	F_pgaio_io_update_state(m, v494, int32(4))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L4
	} else {
		goto L261
	}
L259:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	goto L255
L261:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1570 = v1568 + int32(152)
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+156))
	if v1571 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+160)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+156)) = v1570
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+152)) = v1570
	goto L264
L263:
	;
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v494)+28)) = v1570
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v494)+24)) = v1579
	v1582 = v494 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v1579)+4)) = v1582
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+152)) = v1582
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+160)) = v1585 + int32(1)
	F_pgaio_io_perform_synchronously(m, v494)
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	goto L255
L266:
	;
	m.G0 = v1035 + int32(16)
	goto L164
L267:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_27), int32(0))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_28), int32(1020), int32(_a_F_AsyncReadBuffers_29))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L4
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1039)))
	v1660 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[25]))
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1660+v1658*int32(48))+32))
	goto L272
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1035)+8)) = v1664
	*(*int32)(unsafe.Add(mBase, uint32(v1035))) = v658
	*(*int32)(unsafe.Add(mBase, uint32(v1035)+4)) = v1061 + v658 - int32(1)
	F_errmsg(m, int32(_a_F_AsyncReadBuffers_30), v1035)
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_28), int32(1048), int32(_a_F_AsyncReadBuffers_29))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	v1811 = base.I64_extend_i32_s(v791)
	if v46 == int32(116) {
		goto L293
	} else {
		goto L294
	}
L276:
	;
	F___clock_gettime(m, int32(1), v1695)
	mBase = m.M
	v1701 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1695)+8)))
	v1702 = *(*int64)(unsafe.Add(mBase, uint32(v1695)))
	v1706 = v1701 + (v1702*int64(1000000000) - v1021)
	if v58 == int32(2) {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	goto L278
L278:
	;
	v1790 = v58*int32(320) + v57<<(uint(int32(6))%32)
	v1794 = *(*int64)(unsafe.Add(mBase, uint32(v1790)+uint32(_c_F_AsyncReadBuffers[29])))
	*(*int64)(unsafe.Add(mBase, uint32(v1790)+uint32(_c_F_AsyncReadBuffers[29]))) = v1794 + base.I64_extend_i32_u(v1682)
	v1798 = *(*int64)(unsafe.Add(mBase, uint32(v1790)+uint32(_c_F_AsyncReadBuffers[30])))
	*(*int64)(unsafe.Add(mBase, uint32(v1790)+uint32(_c_F_AsyncReadBuffers[30]))) = v1798 + v1689
	F_pgstat_count_backend_io_op(m, v58, v57, int32(6), v1682, v1689)
	mBase = m.M
	v1803 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[13])) = uint8(v1803)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[14])) = uint8(v1803)
	m.G0 = v1695 + int32(16)
	goto L275
L279:
	;
	v1753 = v58*int32(320) + v57<<(uint(int32(6))%32)
	v1757 = *(*int64)(unsafe.Add(mBase, uint32(v1753)+uint32(_c_F_AsyncReadBuffers[31])))
	*(*int64)(unsafe.Add(mBase, uint32(v1753)+uint32(_c_F_AsyncReadBuffers[31]))) = v1757 + v1706
	v1761 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[32]))
	v1768 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1761))|base.B2i32(int32(1)<<(uint(v1761)%32)&int32(_a_F_AsyncReadBuffers_31) == v1768) == v1768 {
		goto L289
	} else {
		goto L290
	}
L280:
	;
	goto L282
L282:
	;
	goto L283
L283:
	;
	goto L286
L286:
	;
	v1732 = int32(_a_F_AsyncReadBuffers_32)
	v1734 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[33]))
	v1736 = base.I64_div_s(v1706, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[33])) = v1734 + v1736
	switch v58 {
	case 0:
		goto L288
	case 1:
		goto L287
	default:
		goto L279
	}
L287:
	;
	v1744 = int32(_a_F_AsyncReadBuffers_33)
	v1746 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[34]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[34])) = v1746 + v1706
	goto L279
L288:
	;
	v1739 = int32(_a_F_AsyncReadBuffers_34)
	v1741 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[35]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[35])) = v1741 + v1706
	goto L279
L289:
	;
	v1773 = *(*int64)(unsafe.Add(mBase, uint32(v1753)+uint32(_c_F_AsyncReadBuffers[36])))
	*(*int64)(unsafe.Add(mBase, uint32(v1753)+uint32(_c_F_AsyncReadBuffers[36]))) = v1773 + v1706
	v1777 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[13])) = uint8(v1777)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[37])) = uint8(v1777)
	goto L291
L290:
	;
	goto L291
L291:
	;
	goto L278
L292:
	;
	v1824 = int32(1)
	v1826 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[17])))
	if v1826 == v1824 {
		goto L296
	} else {
		goto L297
	}
L293:
	;
	v1814 = int32(_a_F_AsyncReadBuffers_35)
	v1816 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[38]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[38])) = v1816 + v1811
	goto L292
L294:
	;
	goto L295
L295:
	;
	v1819 = int32(_a_F_AsyncReadBuffers_36)
	v1821 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[39]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[39])) = v1821 + v1811
	goto L292
L296:
	;
	v1829 = int32(_a_F_AsyncReadBuffers_13)
	v1831 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18]))
	v1833 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[40]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18])) = v1831 + v1833*v791
	goto L298
L297:
	;
	goto L298
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v791
	v1842 = v1824
	goto L86
}
func F_asyncQueueAdvanceTail(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
	v13 = F_LWLockAcquire(m, v9+int32(_a_F_asyncQueueAdvanceTail_0), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
		v20 = F_LWLockAcquire(m, v16+int32(3456), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[1]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
			if v27 != int32(-1) {
				v30 = v26
				v33 = v27
				v34 = v25
				v36 = v24
				for {
					v39 = v23 + v33*int32(40)
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+80))
					if v30 < v40 {
						v47 = v30
						v50 = v34
						v51 = v36
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
						if v30 == v40 {
							if v34 < v42 {
								v47 = v30
								v50 = v34
								v51 = v36
							} else {
								v45 = v30
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+92))
								v47 = v45
								v50 = v42
								v51 = v46
							}
						} else {
							v45 = v40
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+92))
							v47 = v45
							v50 = v42
							v51 = v46
						}
					}
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
					if v52 != int32(-1) {
						v30 = v47
						v33 = v52
						v34 = v50
						v36 = v51
						continue
					} else {
						break
					}
					break
				}
				v55 = v47
				v59 = v50
				v61 = v51
			} else {
				v55 = v26
				v59 = v25
				v61 = v24
			}
			*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v61
			*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v59
			*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v55
			v65 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
			v67 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
			F_LWLockRelease(m, v67+int32(3456))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				v73 = base.I64_rem_s(v55, int64(32))
				if v65 < v55-v73 {
					F_SimpleLruTruncate(m, int32(_a_F_asyncQueueAdvanceTail_1), v55)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						v80 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
						v84 = F_LWLockAcquire(m, v80+int32(3456), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v87 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v87)+32)) = v55
							v90 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
							F_LWLockRelease(m, v90+int32(3456))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								v96 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
								F_LWLockRelease(m, v96+int32(_a_F_asyncQueueAdvanceTail_0))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
					F_LWLockRelease(m, v96+int32(_a_F_asyncQueueAdvanceTail_0))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_mark_async_capable_plan(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v5 = l1
	goto L1
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v7 != int32(303) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return int32(0)
L3:
	;
	goto L2
L4:
	;
	switch v7 - int32(290) {
	case 0:
		goto L9
	case 1:
		goto L8
	default:
		goto L3
	}
L5:
	;
	goto L6
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v42 == int32(335) {
		goto L3
	} else {
		goto L20
	}
L7:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)) = uint8(v38)
	return v38
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v25 == int32(335) {
		goto L3
	} else {
		goto L16
	}
L9:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(335) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v15 = F_trivial_subqueryscan(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v15 == int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
	v23 = F_mark_async_capable_plan(m, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v23 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+176))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+172))
	if v30 == int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v33 = m.T0[v30].(func(*base.Module, int32) int32)(m, v5)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	if v33 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	goto L7
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v5)+72))
	v5 = v45
	goto L1
}
