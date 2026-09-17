package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AsyncReadBuffers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v626 int64
	_ = v626
	var v630 int32
	_ = v630
	var v632 int64
	_ = v632
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int64
	_ = v650
	var v661 int32
	_ = v661
	var v665 int64
	_ = v665
	var v669 int64
	_ = v669
	var v671 int32
	_ = v671
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int64
	_ = v851
	var v853 int64
	_ = v853
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v913 int64
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v930 int64
	_ = v930
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v943 int32
	_ = v943
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v985 int64
	_ = v985
	var v1025 int32
	_ = v1025
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1045 int64
	_ = v1045
	var v1046 int64
	_ = v1046
	var v1050 int64
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1233 int32
	_ = v1233
	var v1262 int32
	_ = v1262
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int64
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1362 int32
	_ = v1362
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1443 int32
	_ = v1443
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1571 int32
	_ = v1571
	var v1576 int32
	_ = v1576
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1671 int32
	_ = v1671
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1700 int32
	_ = v1700
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1725 int64
	_ = v1725
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1737 int64
	_ = v1737
	var v1738 int64
	_ = v1738
	var v1742 int64
	_ = v1742
	var v1768 int32
	_ = v1768
	var v1770 int64
	_ = v1770
	var v1772 int64
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1777 int64
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1782 int64
	_ = v1782
	var v1789 int32
	_ = v1789
	var v1793 int64
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1804 int32
	_ = v1804
	var v1809 int64
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1826 int32
	_ = v1826
	var v1830 int64
	_ = v1830
	var v1834 int64
	_ = v1834
	var v1839 int32
	_ = v1839
	var v1847 int64
	_ = v1847
	var v1850 int32
	_ = v1850
	var v1852 int64
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1857 int64
	_ = v1857
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1868 int32
	_ = v1868
	v34 = m.G0
	v36 = v34 - int32(512)
	m.G0 = v36
	v38 = int32(3)
	v39 = int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v44 = int32(base.Ui32(v40)>>(uint(v38)%32)) & v39
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+34)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v51 == int32(116) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v65 = v45<<(uint(int32(2))%32) + v48
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[0])))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[1])))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	F_pgstat_prepare_report_checksum_failure(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L7
	}
L2:
	;
	v62 = v44 | int32(2)
	v63 = v38
	v64 = v39
	goto L1
L3:
	;
	goto L4
L4:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v58 = F_IOContextForStrategy(m, v57)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v62 = v44
	v63 = v58
	v64 = int32(0)
	goto L1
L7:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[2]))
	v77 = l0 + int32(48)
	v78 = F_pgaio_io_acquire_nb(m, v75, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v78 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	v498 = v78
	goto L11
L11:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v530)+22)))
	if v531 != 0 {
		goto L89
	} else {
		goto L90
	}
L12:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[2]))
	v86 = m.G0
	v88 = v86 - int32(80)
	m.G0 = v88
	v90 = F_pgaio_io_acquire_nb(m, v85, v77)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L16
	}
L13:
	;
	v498 = v425
	goto L11
L14:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L5
	} else {
		goto L83
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L5
	} else {
		goto L79
	}
L16:
	;
	if v90 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	goto L20
L18:
	;
	v425 = v90
	goto L19
L19:
	;
	m.G0 = v88 + int32(80)
	goto L13
L20:
	;
	v131 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L22
	}
L21:
	;
	v425 = v418
	goto L19
L22:
	;
	if v131 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_errhidestmt(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[4]))
	if int32(0) < v161 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	F_errhidecontext(m)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+160))
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v88)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v88)+72)) = v139
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_0), v88-int32(-64))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(768), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v418 = F_pgaio_io_acquire_nb(m, v85, v77)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L77
	}
L31:
	;
	v164 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v171 = v164
	v172 = v164
	v173 = v159
	v180 = v161
	v181 = v166
	goto L34
L32:
	;
	v236 = v159
	goto L33
L33:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236)+22)))
	if v264 != 0 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v181)+24))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v203 = int32(7)
	v208 = v201 + v202<<(uint(v203)%32) + v171<<(uint(v203)%32)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v209 == int32(6) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if int32(0) < v222 {
		goto L30
	} else {
		goto L41
	}
L36:
	;
	F_pgaio_io_reclaim(m, v208)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L5
	} else {
		goto L39
	}
L37:
	;
	v222 = v172
	v223 = v173
	v224 = v180
	v225 = v181
	goto L38
L38:
	;
	v227 = v171 + int32(1)
	if v227 < v224 {
		v171 = v227
		v172 = v222
		v173 = v223
		v180 = v224
		v181 = v225
		goto L34
	} else {
		goto L40
	}
L39:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[4]))
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v222 = v172 + int32(1)
	v223 = v219
	v224 = v215
	v225 = v217
	goto L38
L40:
	;
	goto L35
L41:
	;
	v236 = v223
	goto L33
L42:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v269 = v236
	goto L44
L44:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	if v270 != 0 {
		goto L30
	} else {
		goto L46
	}
L45:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v269 = v268
	goto L44
L46:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v269)+160))
	if v271 == int32(0) {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v269)+156))
	v276 = v274 - int32(24)
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if base.Ui32(int32(7)) < base.Ui32(v277) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+12))
	if v382 == int32(0) {
		goto L14
	} else {
		goto L76
	}
L49:
	;
	if int32(1)<<(uint(v277)%32)&int32(48) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_pgaio_io_reclaim(m, v276)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L75
	}
L51:
	;
	if v277 == int32(6) {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v274)+24))
	v314 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L58
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+24))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = (v276 - v294) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_3), v88+int32(16))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(840), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	if v314 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_errhidestmt(m)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	F_pgaio_io_wait(m, v276, v311)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L74
	}
L62:
	;
	F_errhidecontext(m)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+24))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+2)))
	if base.Ui32(v327) <= base.Ui32(int32(2)) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v335<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[6])))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	goto L68
L65:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v327<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[7])))
	v334 = v332
	goto L67
L66:
	;
	v334 = int32(0)
	goto L67
L67:
	;
	goto L64
L68:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if base.Ui32(v340) <= base.Ui32(int32(7)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v340<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[8])))
	v346 = v345
	goto L71
L70:
	;
	v346 = int32(0)
	goto L71
L71:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v88+int32(48)))) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = (v276 - v323) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v88)+40)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v88)+44)) = v346
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_4), v88+int32(32))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(847), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	goto L61
L74:
	;
	goto L48
L75:
	;
	goto L48
L76:
	;
	goto L30
L77:
	;
	if v418 == int32(0) {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	goto L21
L79:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_5), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+12))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v467)+160))
	v470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v468
	F_errdetail_internal(m, int32(_a_F_AsyncReadBuffers_6), v88)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(818), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_7), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(881), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	m.G0 = v36 + int32(512)
	return v588
L87:
	;
	if v588 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L88:
	;
	v588 = v586
	goto L87
L89:
	;
	if v528 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L91
L91:
	;
	if v528 < int32(0) {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L5
	} else {
		goto L100
	}
L93:
	;
	v534 = int32(1)
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[9]))
	v543 = F_StartLocalBufferIO(m, v536+(v528^int32(-1))<<(uint(int32(6))%32), v534)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L5
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v547 = int32(1)
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[10]))
	v557 = F_StartBufferIO(m, v549+v528<<(uint(int32(6))%32)+int32(-64), v547, v547)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L5
	} else {
		goto L98
	}
L96:
	;
	if v543 == int32(0) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v586 = v534
	goto L88
L98:
	;
	if v557 != 0 {
		v586 = v547
		goto L88
	} else {
		goto L99
	}
L99:
	;
	goto L92
L100:
	;
	goto L91
L101:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[9]))
	v573 = F_StartLocalBufferIO(m, v566+(v528^int32(-1))<<(uint(int32(6))%32), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L5
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[10]))
	v584 = F_StartBufferIO(m, v576+v528<<(uint(int32(6))%32)+int32(-64), int32(1), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L5
	} else {
		goto L105
	}
L104:
	;
	v588 = v573
	goto L87
L105:
	;
	v586 = v584
	goto L88
L106:
	;
	v591 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)))
	v592 = int32(1)
	v593 = v591 + v592
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)) = uint16(v593)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v592
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v598)+16))
	if v599 == v498 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v693 < int32(0) {
		goto L132
	} else {
		goto L133
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
	goto L117
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v598)+16)) = int32(0)
	F_pgaio_io_reclaim(m, v498)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L5
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L5
	} else {
		goto L114
	}
L113:
	;
	goto L109
L114:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_8), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(258), int32(_a_F_AsyncReadBuffers_9))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	if v51 == int32(116) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v636 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v624 = int32(_a_F_AsyncReadBuffers_10)
	v626 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[11])) = v626 + int64(1)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v630 = int32(_a_F_AsyncReadBuffers_11)
	v632 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[12])) = v632 + int64(1)
	goto L118
L122:
	;
	v661 = v64*int32(320) + v63<<(uint(int32(6))%32)
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v661)+uint32(_c_F_AsyncReadBuffers[13])))
	*(*int64)(unsafe.Add(mBase, uint32(v661)+uint32(_c_F_AsyncReadBuffers[13]))) = v665 + int64(1)
	v669 = *(*int64)(unsafe.Add(mBase, uint32(v661)+uint32(_c_F_AsyncReadBuffers[14])))
	*(*int64)(unsafe.Add(mBase, uint32(v661)+uint32(_c_F_AsyncReadBuffers[14]))) = v669
	v671 = int32(1)
	F_pgstat_count_backend_io_op(m, v64, v63, int32(2), v671, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[15])) = uint8(v671)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[16])) = uint8(v671)
	goto L129
L123:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v636)+272))
	if v639 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636)+268)))
	if v642 != int32(1) {
		goto L122
	} else {
		goto L127
	}
L125:
	;
	v649 = v639
	goto L126
L126:
	;
	v650 = *(*int64)(unsafe.Add(mBase, uint32(v649)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v649)+120)) = v650 + int64(1)
	goto L122
L127:
	;
	F_pgstat_assoc_relation(m, v636)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)+272))
	v649 = v648
	goto L126
L129:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[17])))
	if v681 != int32(1) {
		goto L86
	} else {
		goto L130
	}
L130:
	;
	v684 = int32(_a_F_AsyncReadBuffers_12)
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18]))
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18])) = v686 + v688
	goto L86
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v711
	v713 = int32(1)
	v715 = v45 + v713
	v716 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	if v716 <= v715 {
		v811 = v713
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[20]))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v697+(v693^int32(-1))<<(uint(int32(2))%32))))
	v711 = v703
	goto L131
L133:
	;
	goto L134
L134:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[21]))
	v711 = v705 + v693<<(uint(int32(13))%32) + int32(-8192)
	goto L131
L135:
	;
	v843 = l0 + int32(36)
	v845 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v845)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v843))) = (v498 - v846) >> (uint(int32(7)) % 32)
	v851 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v498)+52)))
	*(*uint32)(unsafe.Add(mBase, uint32(v843)+4)) = uint32(v851)
	v853 = *(*int64)(unsafe.Add(mBase, uint32(v498)+48))
	*(*uint32)(unsafe.Add(mBase, uint32(v843)+8)) = uint32(v853)
	goto L152
L136:
	;
	v720 = v713
	v722 = v715
	goto L137
L137:
	;
	v753 = v48 + v722<<(uint(int32(2))%32)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	if v754 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v811 = v804
	goto L135
L139:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	if v783 < int32(0) {
		goto L148
	} else {
		goto L149
	}
L140:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[9]))
	v765 = F_StartLocalBufferIO(m, v758+(v754^int32(-1))<<(uint(int32(6))%32), int32(1))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L5
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v768 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[10]))
	v774 = int32(1)
	v776 = F_StartBufferIO(m, v768+v754<<(uint(int32(6))%32)+int32(-64), v774, v774)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	if v765 != 0 {
		goto L139
	} else {
		goto L144
	}
L144:
	;
	v811 = v720
	goto L135
L145:
	;
	if v776 == int32(0) {
		v811 = v720
		goto L135
	} else {
		goto L146
	}
L146:
	;
	goto L139
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36+v720<<(uint(int32(2))%32)))) = v801
	v803 = int32(1)
	v804 = v720 + v803
	v806 = v722 + v803
	v807 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	if v806 < v807 {
		v720 = v804
		v722 = v806
		goto L137
	} else {
		goto L151
	}
L148:
	;
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[20]))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v787+(v783^int32(-1))<<(uint(int32(2))%32))))
	v801 = v793
	goto L147
L149:
	;
	goto L150
L150:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[21]))
	v801 = v795 + v783<<(uint(int32(13))%32) + int32(-8192)
	goto L147
L151:
	;
	goto L138
L152:
	;
	v855 = int32(0)
	v858 = v811 & int32(255)
	if v858 == v855 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+13)) = uint8(v858)
	if v51 == int32(116) {
		goto L162
	} else {
		goto L163
	}
L154:
	;
	if v858 != int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v871 = v855
	v877 = v855
	goto L158
L156:
	;
	v943 = v855
	goto L157
L157:
	;
	v973 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v973)+16))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v498)+76))
	v976 = int32(3)
	v985 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65+v943<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v974+v975<<(uint(v976)%32)+v943<<(uint(v976)%32)))) = v985
	goto L153
L158:
	;
	v900 = int32(_a_F_AsyncReadBuffers_13)
	v901 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v901)+16))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v498)+76))
	v904 = int32(3)
	v910 = int32(2)
	v913 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65+v871<<(uint(v910)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v902+v903<<(uint(v904)%32)+v871<<(uint(v904)%32)))) = v913
	v916 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v916)+16))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v498)+76))
	v923 = v871 | int32(1)
	v930 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65+v923<<(uint(v910)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v917+v918<<(uint(v904)%32)+v923<<(uint(v904)%32)))) = v930
	v933 = v871 + v910
	v935 = v877 + v910
	if v935 != v858&int32(254) {
		v871 = v933
		v877 = v935
		goto L158
	} else {
		goto L160
	}
L159:
	;
	if v858&int32(1) == int32(0) {
		goto L153
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	v943 = v933
	goto L157
L162:
	;
	v1025 = int32(3)
	goto L164
L163:
	;
	v1025 = int32(2)
	goto L164
L164:
	;
	F_pgaio_io_register_callbacks(m, v498, v1025, (v40|v67|v69<<(uint(int32(2))%32))&int32(255))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	v1032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+3)))
	v1033 = v1032 | v62
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+3)) = uint8(v1033)
	goto L166
L166:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[22])))
	v1039 = m.G0
	v1041 = v1039 - int32(16)
	m.G0 = v1041
	if v1036 != 0 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1055 = int32(_a_F_AsyncReadBuffers_14)
	v1057 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1057 + int32(1)
	v1061 = int32(0)
	v1062 = m.G0
	v1064 = v1062 - int32(16)
	m.G0 = v1064
	v1066 = v45 + v50
	v1069 = F__mdfd_getseg(m, v1054, v49, v1066, v1061, int32(9))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L5
	} else {
		goto L171
	}
L168:
	;
	F___clock_gettime(m, int32(1), v1041)
	mBase = m.M
	v1045 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1041)+8)))
	v1046 = *(*int64)(unsafe.Add(mBase, uint32(v1041)))
	v1050 = v1045 + v1046*int64(1000000000)
	goto L170
L169:
	;
	v1050 = int64(0)
	goto L170
L170:
	;
	m.G0 = v1041 + int32(16)
	goto L167
L171:
	;
	v1073 = v1066 & int32(_a_F_AsyncReadBuffers_15)
	v1074 = int32(_a_F_AsyncReadBuffers_16) - v1073
	if base.Ui32(v811) <= base.Ui32(v1074) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v1715 = int32(_a_F_AsyncReadBuffers_14)
	v1717 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	v1718 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1717 - v1718
	v1725 = base.I64_extend_i32_s(v811 << (uint(int32(13)) % 32))
	v1729 = m.G0
	v1731 = v1729 - int32(16)
	m.G0 = v1731
	if v1050 != int64(0) {
		goto L284
	} else {
		goto L285
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L5
	} else {
		goto L278
	}
L174:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+12))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v498)+76))
	v1082 = v1078 + v1079<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1064)+12)) = v1082
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v1082)+4)) = int32(_a_F_AsyncReadBuffers_17)
	*(*int32)(unsafe.Add(mBase, uint32(v1082))) = v1084
	v1089 = int32(1)
	if base.Ui32(v811) < base.Ui32(v1074) {
		goto L178
	} else {
		goto L179
	}
L175:
	;
	goto L176
L176:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L5
	} else {
		goto L275
	}
L177:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[24])))
	if v1262&int32(1) == int32(0) {
		goto L200
	} else {
		goto L201
	}
L178:
	;
	v1091 = v811
	goto L180
L179:
	;
	v1091 = v1074
	goto L180
L180:
	;
	if base.Ui32(v1091) < base.Ui32(int32(2)) {
		v1233 = v1089
		goto L177
	} else {
		goto L181
	}
L181:
	;
	v1094 = int32(1)
	if v1091 != int32(2) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1097 = int32(1)
	v1098 = v1091 - v1097
	v1103 = v1084
	v1108 = v1089
	v1109 = v1082
	v1110 = v1094
	v1111 = v1061
	goto L185
L183:
	;
	v1180 = v1084
	v1185 = v1089
	v1186 = v1082
	v1187 = v1094
	goto L184
L184:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v36+v1187<<(uint(int32(2))%32))))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+4))
	if v1216 != v1180+v1217 {
		goto L197
	} else {
		goto L198
	}
L185:
	;
	v1138 = v36 + v1110<<(uint(int32(2))%32)
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1138)))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1109)+4))
	if v1139 == v1103+v1140 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	if v1098&v1097 == int32(0) {
		v1233 = v1171
		goto L177
	} else {
		goto L196
	}
L187:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1138)+4))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1155)+4))
	if v1156 != v1153+v1157 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+4)) = v1140 - int32(-8192)
	v1153 = v1103
	v1154 = v1108
	v1155 = v1109
	goto L187
L189:
	;
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+12)) = int32(_a_F_AsyncReadBuffers_17)
	*(*int32)(unsafe.Add(mBase, uint32(v1109)+8)) = v1139
	v1153 = v1139
	v1154 = v1108 + int32(1)
	v1155 = v1109 + int32(8)
	goto L187
L191:
	;
	v1173 = int32(2)
	v1174 = v1110 + v1173
	v1176 = v1111 + v1173
	if v1176 != v1098&int32(-2) {
		v1103 = v1170
		v1108 = v1171
		v1109 = v1172
		v1110 = v1174
		v1111 = v1176
		goto L185
	} else {
		goto L195
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1155)+12)) = int32(_a_F_AsyncReadBuffers_17)
	*(*int32)(unsafe.Add(mBase, uint32(v1155)+8)) = v1156
	v1170 = v1156
	v1171 = v1154 + int32(1)
	v1172 = v1155 + int32(8)
	goto L191
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1155)+4)) = v1157 - int32(-8192)
	v1170 = v1153
	v1171 = v1154
	v1172 = v1155
	goto L191
L195:
	;
	goto L186
L196:
	;
	v1180 = v1170
	v1185 = v1171
	v1186 = v1172
	v1187 = v1174
	goto L184
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1186)+12)) = int32(_a_F_AsyncReadBuffers_17)
	*(*int32)(unsafe.Add(mBase, uint32(v1186)+8)) = v1216
	v1233 = v1185 + int32(1)
	goto L177
L198:
	;
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1186)+4)) = v1217 - int32(-8192)
	v1233 = v1185
	goto L177
L200:
	;
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+3)))
	v1269 = v1268 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+3)) = uint8(v1269)
	goto L203
L201:
	;
	goto L202
L202:
	;
	v1275 = v498 + int32(104)
	goto L204
L203:
	;
	goto L202
L204:
	;
	v1276 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+1)) = uint8(v1276)
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+8)) = v1278
	v1280 = *(*int64)(unsafe.Add(mBase, uint32(v1054)))
	*(*int64)(unsafe.Add(mBase, uint32(v1275))) = v1280
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+16)) = v811
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+12)) = v1066
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1275)+21)))
	v1289 = v49&int32(255) | v1286<<(uint(int32(8))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1275)+20)) = uint16(v1289)
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+12))
	if v1295 != int32(-1) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1298 = int32(256)
	goto L207
L206:
	;
	v1298 = int32(0)
	goto L207
L207:
	;
	v1299 = v1289&int32(_a_F_AsyncReadBuffers_18) | v1298
	*(*uint16)(unsafe.Add(mBase, uint32(v1275)+20)) = uint16(v1299)
	v1302 = v1299 & int32(_a_F_AsyncReadBuffers_19)
	*(*uint16)(unsafe.Add(mBase, uint32(v1275)+20)) = uint16(v1302)
	F_pgaio_io_register_callbacks(m, v498, int32(1), int32(0))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1069)))
	v1309 = F_FileAccess(m, v1308)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	if v1309 < int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1671 = int32(-1)
	goto L212
L211:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[25]))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1315+v1308*int32(48))))
	*(*int64)(unsafe.Add(mBase, uint32(v498)+96)) = base.I64_extend_i32_u(v1073 << (uint(int32(13)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+88)) = v1319
	*(*uint16)(unsafe.Add(mBase, uint32(v498)+92)) = uint16(v1233)
	v1323 = m.G0
	v1325 = v1323 - int32(32)
	m.G0 = v1325
	v1327 = int32(1)
	v1328 = int32(_a_F_AsyncReadBuffers_14)
	v1330 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1330 + v1327
	*(*int32)(unsafe.Add(mBase, uint32(v498)+20)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+2)) = uint8(v1327)
	F_pgaio_io_update_state(m, v498, int32(2))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L5
	} else {
		goto L213
	}
L212:
	;
	if v1671 != 0 {
		goto L173
	} else {
		goto L274
	}
L213:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v1342)+16)) = int32(0)
	v1345 = m.G0
	v1347 = v1345 - int32(32)
	m.G0 = v1347
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+4)))
	if v1349 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1362 = v1349
	goto L217
L215:
	;
	goto L216
L216:
	;
	m.G0 = v1347 + int32(32)
	F_pgaio_io_update_state(m, v498, int32(3))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L5
	} else {
		goto L242
	}
L217:
	;
	v1396 = v1362 - int32(1)
	v1398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498+int32(5)+v1396))))
	v1400 = v1398 << (uint(int32(3)) % 32)
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+uint32(_c_F_AsyncReadBuffers[26])))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1401)))
	if v1402 != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	goto L216
L219:
	;
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1396+(v498+int32(9))))))
	v1407 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L5
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1362) {
		v1362 = v1396
		goto L217
	} else {
		goto L241
	}
L222:
	;
	if v1407 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	F_errhidestmt(m)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L5
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1401)))
	m.T0[v1463].(func(*base.Module, int32, int32))(m, v498, v1404)
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L5
	} else {
		goto L240
	}
L226:
	;
	F_errhidecontext(m)
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L5
	} else {
		goto L227
	}
L227:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1414)+24))
	goto L228
L228:
	;
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+2)))
	if base.Ui32(v1419) <= base.Ui32(int32(2)) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+1)))
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1427<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[6])))
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1430)+8))
	goto L233
L230:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1419<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[7])))
	v1426 = v1424
	goto L232
L231:
	;
	v1426 = int32(0)
	goto L232
L232:
	;
	goto L229
L233:
	;
	v1432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	if base.Ui32(v1432) <= base.Ui32(int32(7)) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1347+int32(28)))) = v1404
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+uint32(_c_F_AsyncReadBuffers[27])))
	*(*int32)(unsafe.Add(mBase, uint32(v1347+int32(24)))) = v1443
	*(*int32)(unsafe.Add(mBase, uint32(v1347+int32(20)))) = v1398
	*(*int32)(unsafe.Add(mBase, uint32(v1347+int32(16)))) = v1362
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+12)) = v1439
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+8)) = v1431
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+4)) = v1426
	*(*int32)(unsafe.Add(mBase, uint32(v1347))) = (v498 - v1415) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_20), v1347)
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L5
	} else {
		goto L238
	}
L235:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1432<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[8])))
	v1439 = v1437
	goto L237
L236:
	;
	v1439 = int32(0)
	goto L237
L237:
	;
	goto L234
L238:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_21), int32(215), int32(_a_F_AsyncReadBuffers_22))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		goto L5
	} else {
		goto L239
	}
L239:
	;
	goto L225
L240:
	;
	goto L221
L241:
	;
	goto L218
L242:
	;
	v1512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+3)))
	if v1512&int32(1) != 0 {
		v1524 = v1327
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1527 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L5
	} else {
		goto L247
	}
L244:
	;
	v1515 = int32(0)
	v1517 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[28]))
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1517)+16))
	if v1518 == v1515 {
		v1524 = v1515
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v1521 = m.T0[v1518].(func(*base.Module, int32) int32)(m, v498)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	v1524 = v1521
	goto L243
L247:
	;
	if v1527 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	F_errhidestmt(m)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L5
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	if v1524 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L251:
	;
	F_errhidecontext(m)
	mBase = m.M
	v1532 = m.ExcPending
	if v1532 != 0 {
		goto L5
	} else {
		goto L252
	}
L252:
	;
	v1534 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1534)+24))
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+2)))
	if base.Ui32(v1539) <= base.Ui32(int32(2)) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+1)))
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v1547<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[6])))
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(v1550)+8))
	goto L257
L254:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1539<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[7])))
	v1546 = v1544
	goto L256
L255:
	;
	v1546 = int32(0)
	goto L256
L256:
	;
	goto L253
L257:
	;
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	if base.Ui32(v1553) <= base.Ui32(int32(7)) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1553<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[8])))
	v1559 = v1558
	goto L260
L259:
	;
	v1559 = int32(0)
	goto L260
L260:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1561)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v1325)+16)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v1325)+20)) = v1562
	*(*int32)(unsafe.Add(mBase, uint32(v1325))) = (v498 - v1535) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1325)+4)) = v1546
	*(*int32)(unsafe.Add(mBase, uint32(v1325)+8)) = v1551
	*(*int32)(unsafe.Add(mBase, uint32(v1325)+12)) = v1559
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_23), v1325)
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(459), int32(_a_F_AsyncReadBuffers_24))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L5
	} else {
		goto L262
	}
L262:
	;
	goto L250
L263:
	;
	v1628 = int32(_a_F_AsyncReadBuffers_14)
	v1630 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1630 - int32(1)
	m.G0 = v1325 + int32(32)
	v1671 = int32(0)
	goto L212
L264:
	;
	v1584 = int32(_a_F_AsyncReadBuffers_25)
	v1585 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1586 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1585)+22)))
	v1588 = v1586 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1585)+22)) = uint16(v1588)
	*(*int32)(unsafe.Add(mBase, uint32(v1585+v1586<<(uint(int32(2))%32))+24)) = v498
	v1595 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1595)+20)))
	if v1596 != 0 {
		goto L263
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	F_pgaio_io_update_state(m, v498, int32(4))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L5
	} else {
		goto L269
	}
L267:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L5
	} else {
		goto L268
	}
L268:
	;
	goto L263
L269:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1605 = v1603 + int32(152)
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v1603)+156))
	if v1606 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+160)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+156)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+152)) = v1605
	goto L272
L271:
	;
	goto L272
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v498)+28)) = v1605
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1603)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+24)) = v1614
	v1617 = v498 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v1614)+4)) = v1617
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+152)) = v1617
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1603)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v1603)+160)) = v1620 + int32(1)
	F_pgaio_io_perform_synchronously(m, v498)
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L5
	} else {
		goto L273
	}
L273:
	;
	goto L263
L274:
	;
	m.G0 = v1064 + int32(16)
	goto L172
L275:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_26), int32(0))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L5
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_27), int32(1009), int32(_a_F_AsyncReadBuffers_28))
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L5
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L5
	} else {
		goto L279
	}
L279:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1069)))
	v1696 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[25]))
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1696+v1694*int32(48))+32))
	goto L280
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1064)+8)) = v1700
	*(*int32)(unsafe.Add(mBase, uint32(v1064))) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v1064)+4)) = v1066 + v1091 - int32(1)
	F_errmsg(m, int32(_a_F_AsyncReadBuffers_29), v1064)
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L5
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_27), int32(1037), int32(_a_F_AsyncReadBuffers_28))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L5
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	v1847 = base.I64_extend_i32_s(v811)
	if v51 == int32(116) {
		goto L301
	} else {
		goto L302
	}
L284:
	;
	F___clock_gettime(m, int32(1), v1731)
	mBase = m.M
	v1737 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1731)+8)))
	v1738 = *(*int64)(unsafe.Add(mBase, uint32(v1731)))
	v1742 = v1737 + (v1738*int64(1000000000) - v1050)
	if v64 == int32(2) {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	goto L286
L286:
	;
	v1826 = v64*int32(320) + v63<<(uint(int32(6))%32)
	v1830 = *(*int64)(unsafe.Add(mBase, uint32(v1826)+uint32(_c_F_AsyncReadBuffers[29])))
	*(*int64)(unsafe.Add(mBase, uint32(v1826)+uint32(_c_F_AsyncReadBuffers[29]))) = v1830 + base.I64_extend_i32_u(v1718)
	v1834 = *(*int64)(unsafe.Add(mBase, uint32(v1826)+uint32(_c_F_AsyncReadBuffers[30])))
	*(*int64)(unsafe.Add(mBase, uint32(v1826)+uint32(_c_F_AsyncReadBuffers[30]))) = v1834 + v1725
	F_pgstat_count_backend_io_op(m, v64, v63, int32(6), v1718, v1725)
	mBase = m.M
	v1839 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[15])) = uint8(v1839)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[16])) = uint8(v1839)
	m.G0 = v1731 + int32(16)
	goto L283
L287:
	;
	v1789 = v64*int32(320) + v63<<(uint(int32(6))%32)
	v1793 = *(*int64)(unsafe.Add(mBase, uint32(v1789)+uint32(_c_F_AsyncReadBuffers[31])))
	*(*int64)(unsafe.Add(mBase, uint32(v1789)+uint32(_c_F_AsyncReadBuffers[31]))) = v1793 + v1742
	v1797 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[32]))
	v1804 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1797))|base.B2i32(int32(1)<<(uint(v1797)%32)&int32(_a_F_AsyncReadBuffers_30) == v1804) == v1804 {
		goto L297
	} else {
		goto L298
	}
L288:
	;
	goto L290
L290:
	;
	goto L291
L291:
	;
	goto L294
L294:
	;
	v1768 = int32(_a_F_AsyncReadBuffers_31)
	v1770 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[33]))
	v1772 = base.I64_div_s(v1742, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[33])) = v1770 + v1772
	switch v64 {
	case 0:
		goto L296
	case 1:
		goto L295
	default:
		goto L287
	}
L295:
	;
	v1780 = int32(_a_F_AsyncReadBuffers_32)
	v1782 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[34]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[34])) = v1782 + v1742
	goto L287
L296:
	;
	v1775 = int32(_a_F_AsyncReadBuffers_33)
	v1777 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[35]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[35])) = v1777 + v1742
	goto L287
L297:
	;
	v1809 = *(*int64)(unsafe.Add(mBase, uint32(v1789)+uint32(_c_F_AsyncReadBuffers[36])))
	*(*int64)(unsafe.Add(mBase, uint32(v1789)+uint32(_c_F_AsyncReadBuffers[36]))) = v1809 + v1742
	v1813 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[15])) = uint8(v1813)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[37])) = uint8(v1813)
	goto L299
L298:
	;
	goto L299
L299:
	;
	goto L286
L300:
	;
	v1861 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[17])))
	if v1861 == int32(1) {
		goto L304
	} else {
		goto L305
	}
L301:
	;
	v1850 = int32(_a_F_AsyncReadBuffers_34)
	v1852 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[38]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[38])) = v1852 + v1847
	goto L300
L302:
	;
	goto L303
L303:
	;
	v1855 = int32(_a_F_AsyncReadBuffers_35)
	v1857 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[39]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[39])) = v1857 + v1847
	goto L300
L304:
	;
	v1864 = int32(_a_F_AsyncReadBuffers_12)
	v1866 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18]))
	v1868 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[40]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18])) = v1866 + v1868*v811
	goto L306
L305:
	;
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v811
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
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
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
					v39 = v23 + v33<<(uint(int32(5))%32)
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+72))
					if v30 < v40 {
						v47 = v30
						v50 = v34
						v51 = v36
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
						if v30 == v40 {
							if v34 < v42 {
								v47 = v30
								v50 = v34
								v51 = v36
							} else {
								v45 = v30
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
								v47 = v45
								v50 = v42
								v51 = v46
							}
						} else {
							v45 = v40
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
							v47 = v45
							v50 = v42
							v51 = v46
						}
					}
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v39-int32(-64))))
					if v54 != int32(-1) {
						v30 = v47
						v33 = v54
						v34 = v50
						v36 = v51
						continue
					} else {
						break
					}
					break
				}
				v57 = v47
				v61 = v50
				v63 = v51
			} else {
				v57 = v26
				v61 = v25
				v63 = v24
			}
			*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v63
			*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v61
			*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v57
			v67 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
			v69 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
			F_LWLockRelease(m, v69+int32(3456))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return
			} else {
				v75 = base.I64_rem_s(v57, int64(32))
				if v67 < v57-v75 {
					F_SimpleLruTruncate(m, int32(_a_F_asyncQueueAdvanceTail_1), v57)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
						v86 = F_LWLockAcquire(m, v82+int32(3456), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v89)+32)) = v57
							v92 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
							F_LWLockRelease(m, v92+int32(3456))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
								F_LWLockRelease(m, v98+int32(_a_F_asyncQueueAdvanceTail_0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, _c_F_asyncQueueAdvanceTail[0]))
					F_LWLockRelease(m, v98+int32(_a_F_asyncQueueAdvanceTail_0))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_asyncQueuePagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	return base.B2i32(l0 < l1)
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
	if v7 != int32(301) {
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
	switch v7 - int32(287) {
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
	if v42 == int32(331) {
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
	if v25 == int32(331) {
		goto L3
	} else {
		goto L16
	}
L9:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(331) {
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+168))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+168))
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
