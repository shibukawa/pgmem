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
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v634 int64
	_ = v634
	var v638 int32
	_ = v638
	var v640 int64
	_ = v640
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int64
	_ = v658
	var v669 int32
	_ = v669
	var v673 int64
	_ = v673
	var v677 int64
	_ = v677
	var v679 int32
	_ = v679
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v859 int64
	_ = v859
	var v861 int64
	_ = v861
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v879 int32
	_ = v879
	var v885 int32
	_ = v885
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v918 int32
	_ = v918
	var v921 int64
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v938 int64
	_ = v938
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v951 int32
	_ = v951
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v993 int64
	_ = v993
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1053 int64
	_ = v1053
	var v1054 int64
	_ = v1054
	var v1058 int64
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1241 int32
	_ = v1241
	var v1270 int32
	_ = v1270
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1288 int64
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1370 int32
	_ = v1370
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1461 int32
	_ = v1461
	var v1466 int32
	_ = v1466
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1579 int32
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1596 int32
	_ = v1596
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1679 int32
	_ = v1679
	var v1686 int32
	_ = v1686
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1708 int32
	_ = v1708
	var v1717 int32
	_ = v1717
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1733 int64
	_ = v1733
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1745 int64
	_ = v1745
	var v1746 int64
	_ = v1746
	var v1750 int64
	_ = v1750
	var v1776 int32
	_ = v1776
	var v1778 int64
	_ = v1778
	var v1780 int64
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1785 int64
	_ = v1785
	var v1788 int32
	_ = v1788
	var v1790 int64
	_ = v1790
	var v1797 int32
	_ = v1797
	var v1801 int64
	_ = v1801
	var v1805 int32
	_ = v1805
	var v1812 int32
	_ = v1812
	var v1817 int64
	_ = v1817
	var v1821 int32
	_ = v1821
	var v1834 int32
	_ = v1834
	var v1838 int64
	_ = v1838
	var v1842 int64
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1855 int64
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1860 int64
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1865 int64
	_ = v1865
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
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
	v506 = v78
	goto L11
L11:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v538)+22)))
	if v539 != 0 {
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
	v506 = v433
	goto L11
L14:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L5
	} else {
		goto L83
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
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
	v433 = v90
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
	v433 = v426
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
	v426 = F_pgaio_io_acquire_nb(m, v85, v77)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
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
	v240 = v159
	goto L33
L33:
	;
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240)+22)))
	if v268 != 0 {
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
	if int32(0) < v226 {
		goto L30
	} else {
		goto L41
	}
L36:
	;
	v212 = int32(0)
	v215 = base.AtomicRmwOr32(m, v212, int32(_a_F_AsyncReadBuffers_3), v212)
	F_pgaio_io_reclaim(m, v208)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L39
	}
L37:
	;
	v226 = v172
	v227 = v173
	v228 = v180
	v229 = v181
	goto L38
L38:
	;
	v231 = v171 + int32(1)
	if v231 < v228 {
		v171 = v231
		v172 = v226
		v173 = v227
		v180 = v228
		v181 = v229
		goto L34
	} else {
		goto L40
	}
L39:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[4]))
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v226 = v172 + int32(1)
	v227 = v223
	v228 = v219
	v229 = v221
	goto L38
L40:
	;
	goto L35
L41:
	;
	v240 = v227
	goto L33
L42:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v273 = v240
	goto L44
L44:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	if v274 != 0 {
		goto L30
	} else {
		goto L46
	}
L45:
	;
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v273 = v272
	goto L44
L46:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v273)+160))
	if v275 == int32(0) {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v273)+156))
	v280 = v278 - int32(24)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if base.Ui32(int32(7)) < base.Ui32(v281) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	if v390 == int32(0) {
		goto L14
	} else {
		goto L76
	}
L49:
	;
	if int32(1)<<(uint(v281)%32)&int32(48) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v376 = int32(0)
	v379 = base.AtomicRmwOr32(m, v376, int32(_a_F_AsyncReadBuffers_3), v376)
	F_pgaio_io_reclaim(m, v280)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L5
	} else {
		goto L75
	}
L51:
	;
	if v281 == int32(6) {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v278)+24))
	v318 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L5
	} else {
		goto L58
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+24))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = (v280 - v298) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_4), v88+int32(16))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(840), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
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
	if v318 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_errhidestmt(m)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	F_pgaio_io_wait(m, v280, v315)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L74
	}
L62:
	;
	F_errhidecontext(m)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+24))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+2)))
	if base.Ui32(v331) <= base.Ui32(int32(2)) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v339<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[6])))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	goto L68
L65:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v331<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[7])))
	v338 = v336
	goto L67
L66:
	;
	v338 = int32(0)
	goto L67
L67:
	;
	goto L64
L68:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if base.Ui32(v344) <= base.Ui32(int32(7)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v344<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[8])))
	v350 = v349
	goto L71
L70:
	;
	v350 = int32(0)
	goto L71
L71:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v88+int32(48)))) = v353
	*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = (v280 - v327) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v88)+40)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v88)+44)) = v350
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_5), v88+int32(32))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(847), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
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
	if v426 == int32(0) {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	goto L21
L79:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_6), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+12))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v475)+160))
	v478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v475)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v477
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v476
	F_errdetail_internal(m, int32(_a_F_AsyncReadBuffers_7), v88)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(818), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_8), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(881), int32(_a_F_AsyncReadBuffers_2))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
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
	return v596
L87:
	;
	if v596 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L88:
	;
	v596 = v594
	goto L87
L89:
	;
	if v536 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L91
L91:
	;
	if v536 < int32(0) {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L5
	} else {
		goto L100
	}
L93:
	;
	v542 = int32(1)
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[9]))
	v551 = F_StartLocalBufferIO(m, v544+(v536^int32(-1))<<(uint(int32(6))%32), v542)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L5
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v555 = int32(1)
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[10]))
	v565 = F_StartBufferIO(m, v557+v536<<(uint(int32(6))%32)+int32(-64), v555, v555)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L5
	} else {
		goto L98
	}
L96:
	;
	if v551 == int32(0) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v594 = v542
	goto L88
L98:
	;
	if v565 != 0 {
		v594 = v555
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
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[9]))
	v581 = F_StartLocalBufferIO(m, v574+(v536^int32(-1))<<(uint(int32(6))%32), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L5
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[10]))
	v592 = F_StartBufferIO(m, v584+v536<<(uint(int32(6))%32)+int32(-64), int32(1), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L5
	} else {
		goto L105
	}
L104:
	;
	v596 = v581
	goto L87
L105:
	;
	v594 = v592
	goto L88
L106:
	;
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)))
	v600 = int32(1)
	v601 = v599 + v600
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)) = uint16(v601)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v600
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)+16))
	if v607 == v506 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v701 < int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v606)+16)) = int32(0)
	F_pgaio_io_reclaim(m, v506)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
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
	v616 = m.ExcPending
	if v616 != 0 {
		goto L5
	} else {
		goto L114
	}
L113:
	;
	goto L109
L114:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_9), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(258), int32(_a_F_AsyncReadBuffers_10))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
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
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v644 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v632 = int32(_a_F_AsyncReadBuffers_11)
	v634 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[11])) = v634 + int64(1)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v638 = int32(_a_F_AsyncReadBuffers_12)
	v640 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[12])) = v640 + int64(1)
	goto L118
L122:
	;
	v669 = v64*int32(320) + v63<<(uint(int32(6))%32)
	v673 = *(*int64)(unsafe.Add(mBase, uint32(v669)+uint32(_c_F_AsyncReadBuffers[13])))
	*(*int64)(unsafe.Add(mBase, uint32(v669)+uint32(_c_F_AsyncReadBuffers[13]))) = v673 + int64(1)
	v677 = *(*int64)(unsafe.Add(mBase, uint32(v669)+uint32(_c_F_AsyncReadBuffers[14])))
	*(*int64)(unsafe.Add(mBase, uint32(v669)+uint32(_c_F_AsyncReadBuffers[14]))) = v677
	v679 = int32(1)
	F_pgstat_count_backend_io_op(m, v64, v63, int32(2), v679, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[15])) = uint8(v679)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[16])) = uint8(v679)
	goto L129
L123:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v644)+272))
	if v647 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644)+268)))
	if v650 != int32(1) {
		goto L122
	} else {
		goto L127
	}
L125:
	;
	v657 = v647
	goto L126
L126:
	;
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v657)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v657)+120)) = v658 + int64(1)
	goto L122
L127:
	;
	F_pgstat_assoc_relation(m, v644)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)+272))
	v657 = v656
	goto L126
L129:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[17])))
	if v689 != int32(1) {
		goto L86
	} else {
		goto L130
	}
L130:
	;
	v692 = int32(_a_F_AsyncReadBuffers_13)
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18]))
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[19]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18])) = v694 + v696
	goto L86
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v719
	v721 = int32(1)
	v723 = v45 + v721
	v724 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	if v724 <= v723 {
		v819 = v721
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[20]))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v705+(v701^int32(-1))<<(uint(int32(2))%32))))
	v719 = v711
	goto L131
L133:
	;
	goto L134
L134:
	;
	v713 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[21]))
	v719 = v713 + v701<<(uint(int32(13))%32) + int32(-8192)
	goto L131
L135:
	;
	v851 = l0 + int32(36)
	v853 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v853)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v851))) = (v506 - v854) >> (uint(int32(7)) % 32)
	v859 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v506)+52)))
	*(*uint32)(unsafe.Add(mBase, uint32(v851)+4)) = uint32(v859)
	v861 = *(*int64)(unsafe.Add(mBase, uint32(v506)+48))
	*(*uint32)(unsafe.Add(mBase, uint32(v851)+8)) = uint32(v861)
	goto L152
L136:
	;
	v728 = v721
	v730 = v723
	goto L137
L137:
	;
	v761 = v48 + v730<<(uint(int32(2))%32)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	if v762 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v819 = v812
	goto L135
L139:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	if v791 < int32(0) {
		goto L148
	} else {
		goto L149
	}
L140:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[9]))
	v773 = F_StartLocalBufferIO(m, v766+(v762^int32(-1))<<(uint(int32(6))%32), int32(1))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L5
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v776 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[10]))
	v782 = int32(1)
	v784 = F_StartBufferIO(m, v776+v762<<(uint(int32(6))%32)+int32(-64), v782, v782)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	if v773 != 0 {
		goto L139
	} else {
		goto L144
	}
L144:
	;
	v819 = v728
	goto L135
L145:
	;
	if v784 == int32(0) {
		v819 = v728
		goto L135
	} else {
		goto L146
	}
L146:
	;
	goto L139
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36+v728<<(uint(int32(2))%32)))) = v809
	v811 = int32(1)
	v812 = v728 + v811
	v814 = v730 + v811
	v815 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	if v814 < v815 {
		v728 = v812
		v730 = v814
		goto L137
	} else {
		goto L151
	}
L148:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[20]))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v795+(v791^int32(-1))<<(uint(int32(2))%32))))
	v809 = v801
	goto L147
L149:
	;
	goto L150
L150:
	;
	v803 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[21]))
	v809 = v803 + v791<<(uint(int32(13))%32) + int32(-8192)
	goto L147
L151:
	;
	goto L138
L152:
	;
	v863 = int32(0)
	v866 = v819 & int32(255)
	if v866 == v863 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+13)) = uint8(v866)
	if v51 == int32(116) {
		goto L162
	} else {
		goto L163
	}
L154:
	;
	if v866 != int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v879 = v863
	v885 = v863
	goto L158
L156:
	;
	v951 = v863
	goto L157
L157:
	;
	v981 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v981)+16))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v506)+76))
	v984 = int32(3)
	v993 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65+v951<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v982+v983<<(uint(v984)%32)+v951<<(uint(v984)%32)))) = v993
	goto L153
L158:
	;
	v908 = int32(_a_F_AsyncReadBuffers_14)
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)+16))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v506)+76))
	v912 = int32(3)
	v918 = int32(2)
	v921 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65+v879<<(uint(v918)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v910+v911<<(uint(v912)%32)+v879<<(uint(v912)%32)))) = v921
	v924 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)+16))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v506)+76))
	v931 = v879 | int32(1)
	v938 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65+v931<<(uint(v918)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v925+v926<<(uint(v912)%32)+v931<<(uint(v912)%32)))) = v938
	v941 = v879 + v918
	v943 = v885 + v918
	if v943 != v866&int32(254) {
		v879 = v941
		v885 = v943
		goto L158
	} else {
		goto L160
	}
L159:
	;
	if v866&int32(1) == int32(0) {
		goto L153
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	v951 = v941
	goto L157
L162:
	;
	v1033 = int32(3)
	goto L164
L163:
	;
	v1033 = int32(2)
	goto L164
L164:
	;
	F_pgaio_io_register_callbacks(m, v506, v1033, (v40|v67|v69<<(uint(int32(2))%32))&int32(255))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+3)))
	v1041 = v1040 | v62
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+3)) = uint8(v1041)
	goto L166
L166:
	;
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[22])))
	v1047 = m.G0
	v1049 = v1047 - int32(16)
	m.G0 = v1049
	if v1044 != 0 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1063 = int32(_a_F_AsyncReadBuffers_15)
	v1065 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1065 + int32(1)
	v1069 = int32(0)
	v1070 = m.G0
	v1072 = v1070 - int32(16)
	m.G0 = v1072
	v1074 = v45 + v50
	v1077 = F__mdfd_getseg(m, v1062, v49, v1074, v1069, int32(9))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L5
	} else {
		goto L171
	}
L168:
	;
	F___clock_gettime(m, int32(1), v1049)
	mBase = m.M
	v1053 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1049)+8)))
	v1054 = *(*int64)(unsafe.Add(mBase, uint32(v1049)))
	v1058 = v1053 + v1054*int64(1000000000)
	goto L170
L169:
	;
	v1058 = int64(0)
	goto L170
L170:
	;
	m.G0 = v1049 + int32(16)
	goto L167
L171:
	;
	v1081 = v1074 & int32(_a_F_AsyncReadBuffers_16)
	v1082 = int32(_a_F_AsyncReadBuffers_17) - v1081
	if base.Ui32(v819) <= base.Ui32(v1082) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v1723 = int32(_a_F_AsyncReadBuffers_15)
	v1725 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	v1726 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1725 - v1726
	v1733 = base.I64_extend_i32_s(v819 << (uint(int32(13)) % 32))
	v1737 = m.G0
	v1739 = v1737 - int32(16)
	m.G0 = v1739
	if v1058 != int64(0) {
		goto L284
	} else {
		goto L285
	}
L173:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L5
	} else {
		goto L278
	}
L174:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+12))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v506)+76))
	v1090 = v1086 + v1087<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1072)+12)) = v1090
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+4)) = int32(_a_F_AsyncReadBuffers_18)
	*(*int32)(unsafe.Add(mBase, uint32(v1090))) = v1092
	v1097 = int32(1)
	if base.Ui32(v819) < base.Ui32(v1082) {
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
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L5
	} else {
		goto L275
	}
L177:
	;
	v1270 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[24])))
	if v1270&int32(1) == int32(0) {
		goto L200
	} else {
		goto L201
	}
L178:
	;
	v1099 = v819
	goto L180
L179:
	;
	v1099 = v1082
	goto L180
L180:
	;
	if base.Ui32(v1099) < base.Ui32(int32(2)) {
		v1241 = v1097
		goto L177
	} else {
		goto L181
	}
L181:
	;
	v1102 = int32(1)
	if v1099 != int32(2) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v1105 = int32(1)
	v1106 = v1099 - v1105
	v1111 = v1092
	v1116 = v1097
	v1117 = v1090
	v1118 = v1102
	v1119 = v1069
	goto L185
L183:
	;
	v1188 = v1092
	v1193 = v1097
	v1194 = v1090
	v1195 = v1102
	goto L184
L184:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v36+v1195<<(uint(int32(2))%32))))
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1194)+4))
	if v1224 != v1188+v1225 {
		goto L197
	} else {
		goto L198
	}
L185:
	;
	v1146 = v36 + v1118<<(uint(int32(2))%32)
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1146)))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+4))
	if v1147 == v1111+v1148 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	if v1106&v1105 == int32(0) {
		v1241 = v1179
		goto L177
	} else {
		goto L196
	}
L187:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1146)+4))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+4))
	if v1164 != v1161+v1165 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1117)+4)) = v1148 - int32(-8192)
	v1161 = v1111
	v1162 = v1116
	v1163 = v1117
	goto L187
L189:
	;
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1117)+12)) = int32(_a_F_AsyncReadBuffers_18)
	*(*int32)(unsafe.Add(mBase, uint32(v1117)+8)) = v1147
	v1161 = v1147
	v1162 = v1116 + int32(1)
	v1163 = v1117 + int32(8)
	goto L187
L191:
	;
	v1181 = int32(2)
	v1182 = v1118 + v1181
	v1184 = v1119 + v1181
	if v1184 != v1106&int32(-2) {
		v1111 = v1178
		v1116 = v1179
		v1117 = v1180
		v1118 = v1182
		v1119 = v1184
		goto L185
	} else {
		goto L195
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+12)) = int32(_a_F_AsyncReadBuffers_18)
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+8)) = v1164
	v1178 = v1164
	v1179 = v1162 + int32(1)
	v1180 = v1163 + int32(8)
	goto L191
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1163)+4)) = v1165 - int32(-8192)
	v1178 = v1161
	v1179 = v1162
	v1180 = v1163
	goto L191
L195:
	;
	goto L186
L196:
	;
	v1188 = v1178
	v1193 = v1179
	v1194 = v1180
	v1195 = v1182
	goto L184
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1194)+12)) = int32(_a_F_AsyncReadBuffers_18)
	*(*int32)(unsafe.Add(mBase, uint32(v1194)+8)) = v1224
	v1241 = v1193 + int32(1)
	goto L177
L198:
	;
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1194)+4)) = v1225 - int32(-8192)
	v1241 = v1193
	goto L177
L200:
	;
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+3)))
	v1277 = v1276 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+3)) = uint8(v1277)
	goto L203
L201:
	;
	goto L202
L202:
	;
	v1283 = v506 + int32(104)
	goto L204
L203:
	;
	goto L202
L204:
	;
	v1284 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+1)) = uint8(v1284)
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1062)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1283)+8)) = v1286
	v1288 = *(*int64)(unsafe.Add(mBase, uint32(v1062)))
	*(*int64)(unsafe.Add(mBase, uint32(v1283))) = v1288
	*(*int32)(unsafe.Add(mBase, uint32(v1283)+16)) = v819
	*(*int32)(unsafe.Add(mBase, uint32(v1283)+12)) = v1074
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1283)+21)))
	v1297 = v49&int32(255) | v1294<<(uint(int32(8))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1283)+20)) = uint16(v1297)
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1062)+12))
	if v1303 != int32(-1) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1306 = int32(256)
	goto L207
L206:
	;
	v1306 = int32(0)
	goto L207
L207:
	;
	v1307 = v1297&int32(_a_F_AsyncReadBuffers_19) | v1306
	*(*uint16)(unsafe.Add(mBase, uint32(v1283)+20)) = uint16(v1307)
	v1310 = v1307 & int32(_a_F_AsyncReadBuffers_20)
	*(*uint16)(unsafe.Add(mBase, uint32(v1283)+20)) = uint16(v1310)
	F_pgaio_io_register_callbacks(m, v506, int32(1), int32(0))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L5
	} else {
		goto L208
	}
L208:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	v1317 = F_FileAccess(m, v1316)
	mBase = m.M
	v1318 = m.ExcPending
	if v1318 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	if v1317 < int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v1679 = int32(-1)
	goto L212
L211:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[25]))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1323+v1316*int32(48))))
	*(*int64)(unsafe.Add(mBase, uint32(v506)+96)) = base.I64_extend_i32_u(v1081 << (uint(int32(13)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+88)) = v1327
	*(*uint16)(unsafe.Add(mBase, uint32(v506)+92)) = uint16(v1241)
	v1331 = m.G0
	v1333 = v1331 - int32(32)
	m.G0 = v1333
	v1335 = int32(1)
	v1336 = int32(_a_F_AsyncReadBuffers_15)
	v1338 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1338 + v1335
	*(*int32)(unsafe.Add(mBase, uint32(v506)+20)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v506)+2)) = uint8(v1335)
	F_pgaio_io_update_state(m, v506, int32(2))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L5
	} else {
		goto L213
	}
L212:
	;
	if v1679 != 0 {
		goto L173
	} else {
		goto L274
	}
L213:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v1350)+16)) = int32(0)
	v1353 = m.G0
	v1355 = v1353 - int32(32)
	m.G0 = v1355
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+4)))
	if v1357 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1370 = v1357
	goto L217
L215:
	;
	goto L216
L216:
	;
	m.G0 = v1355 + int32(32)
	F_pgaio_io_update_state(m, v506, int32(3))
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L5
	} else {
		goto L242
	}
L217:
	;
	v1404 = v1370 - int32(1)
	v1406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506+int32(5)+v1404))))
	v1408 = v1406 << (uint(int32(3)) % 32)
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+uint32(_c_F_AsyncReadBuffers[26])))
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1409)))
	if v1410 != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	goto L216
L219:
	;
	v1412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1404+(v506+int32(9))))))
	v1415 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L5
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1370) {
		v1370 = v1404
		goto L217
	} else {
		goto L241
	}
L222:
	;
	if v1415 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	F_errhidestmt(m)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L5
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1409)))
	m.T0[v1471].(func(*base.Module, int32, int32))(m, v506, v1412)
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L5
	} else {
		goto L240
	}
L226:
	;
	F_errhidecontext(m)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L5
	} else {
		goto L227
	}
L227:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1422)+24))
	goto L228
L228:
	;
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+2)))
	if base.Ui32(v1427) <= base.Ui32(int32(2)) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+1)))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1435<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[6])))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+8))
	goto L233
L230:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1427<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[7])))
	v1434 = v1432
	goto L232
L231:
	;
	v1434 = int32(0)
	goto L232
L232:
	;
	goto L229
L233:
	;
	v1440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if base.Ui32(v1440) <= base.Ui32(int32(7)) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1355+int32(28)))) = v1412
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1408)+uint32(_c_F_AsyncReadBuffers[27])))
	*(*int32)(unsafe.Add(mBase, uint32(v1355+int32(24)))) = v1451
	*(*int32)(unsafe.Add(mBase, uint32(v1355+int32(20)))) = v1406
	*(*int32)(unsafe.Add(mBase, uint32(v1355+int32(16)))) = v1370
	*(*int32)(unsafe.Add(mBase, uint32(v1355)+12)) = v1447
	*(*int32)(unsafe.Add(mBase, uint32(v1355)+8)) = v1439
	*(*int32)(unsafe.Add(mBase, uint32(v1355)+4)) = v1434
	*(*int32)(unsafe.Add(mBase, uint32(v1355))) = (v506 - v1423) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_21), v1355)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L5
	} else {
		goto L238
	}
L235:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1440<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[8])))
	v1447 = v1445
	goto L237
L236:
	;
	v1447 = int32(0)
	goto L237
L237:
	;
	goto L234
L238:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_22), int32(215), int32(_a_F_AsyncReadBuffers_23))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
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
	v1520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+3)))
	if v1520&int32(1) != 0 {
		v1532 = v1335
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1535 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v1536 = m.ExcPending
	if v1536 != 0 {
		goto L5
	} else {
		goto L247
	}
L244:
	;
	v1523 = int32(0)
	v1525 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[28]))
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1525)+16))
	if v1526 == v1523 {
		v1532 = v1523
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v1529 = m.T0[v1526].(func(*base.Module, int32) int32)(m, v506)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	v1532 = v1529
	goto L243
L247:
	;
	if v1535 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	F_errhidestmt(m)
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L5
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	if v1532 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L251:
	;
	F_errhidecontext(m)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L5
	} else {
		goto L252
	}
L252:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[5]))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1542)+24))
	v1547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+2)))
	if base.Ui32(v1547) <= base.Ui32(int32(2)) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+1)))
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1555<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[6])))
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+8))
	goto L257
L254:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1547<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[7])))
	v1554 = v1552
	goto L256
L255:
	;
	v1554 = int32(0)
	goto L256
L256:
	;
	goto L253
L257:
	;
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if base.Ui32(v1561) <= base.Ui32(int32(7)) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1561<<(uint(int32(2))%32))+uint32(_c_F_AsyncReadBuffers[8])))
	v1567 = v1566
	goto L260
L259:
	;
	v1567 = int32(0)
	goto L260
L260:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1569)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v1333)+16)) = v1532
	*(*int32)(unsafe.Add(mBase, uint32(v1333)+20)) = v1570
	*(*int32)(unsafe.Add(mBase, uint32(v1333))) = (v506 - v1543) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1333)+4)) = v1554
	*(*int32)(unsafe.Add(mBase, uint32(v1333)+8)) = v1559
	*(*int32)(unsafe.Add(mBase, uint32(v1333)+12)) = v1567
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_24), v1333)
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_1), int32(459), int32(_a_F_AsyncReadBuffers_25))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L5
	} else {
		goto L262
	}
L262:
	;
	goto L250
L263:
	;
	v1636 = int32(_a_F_AsyncReadBuffers_15)
	v1638 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[23])) = v1638 - int32(1)
	m.G0 = v1333 + int32(32)
	v1679 = int32(0)
	goto L212
L264:
	;
	v1592 = int32(_a_F_AsyncReadBuffers_26)
	v1593 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1594 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1593)+22)))
	v1596 = v1594 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1593)+22)) = uint16(v1596)
	*(*int32)(unsafe.Add(mBase, uint32(v1593+v1594<<(uint(int32(2))%32))+24)) = v506
	v1603 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1603)+20)))
	if v1604 != 0 {
		goto L263
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	F_pgaio_io_update_state(m, v506, int32(4))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L5
	} else {
		goto L269
	}
L267:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L5
	} else {
		goto L268
	}
L268:
	;
	goto L263
L269:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[3]))
	v1613 = v1611 + int32(152)
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+156))
	if v1614 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1611)+160)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1611)+156)) = v1613
	*(*int32)(unsafe.Add(mBase, uint32(v1611)+152)) = v1613
	goto L272
L271:
	;
	goto L272
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+28)) = v1613
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+24)) = v1622
	v1625 = v506 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+4)) = v1625
	*(*int32)(unsafe.Add(mBase, uint32(v1611)+152)) = v1625
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1611)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v1611)+160)) = v1628 + int32(1)
	F_pgaio_io_perform_synchronously(m, v506)
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L5
	} else {
		goto L273
	}
L273:
	;
	goto L263
L274:
	;
	m.G0 = v1072 + int32(16)
	goto L172
L275:
	;
	F_errmsg_internal(m, int32(_a_F_AsyncReadBuffers_27), int32(0))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L5
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_28), int32(1009), int32(_a_F_AsyncReadBuffers_29))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
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
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L5
	} else {
		goto L279
	}
L279:
	;
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	v1704 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[25]))
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1704+v1702*int32(48))+32))
	goto L280
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1072)+8)) = v1708
	*(*int32)(unsafe.Add(mBase, uint32(v1072))) = v1074
	*(*int32)(unsafe.Add(mBase, uint32(v1072)+4)) = v1074 + v1099 - int32(1)
	F_errmsg(m, int32(_a_F_AsyncReadBuffers_30), v1072)
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L5
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_AsyncReadBuffers_28), int32(1037), int32(_a_F_AsyncReadBuffers_29))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
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
	v1855 = base.I64_extend_i32_s(v819)
	if v51 == int32(116) {
		goto L301
	} else {
		goto L302
	}
L284:
	;
	F___clock_gettime(m, int32(1), v1739)
	mBase = m.M
	v1745 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1739)+8)))
	v1746 = *(*int64)(unsafe.Add(mBase, uint32(v1739)))
	v1750 = v1745 + (v1746*int64(1000000000) - v1058)
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
	v1834 = v64*int32(320) + v63<<(uint(int32(6))%32)
	v1838 = *(*int64)(unsafe.Add(mBase, uint32(v1834)+uint32(_c_F_AsyncReadBuffers[29])))
	*(*int64)(unsafe.Add(mBase, uint32(v1834)+uint32(_c_F_AsyncReadBuffers[29]))) = v1838 + base.I64_extend_i32_u(v1726)
	v1842 = *(*int64)(unsafe.Add(mBase, uint32(v1834)+uint32(_c_F_AsyncReadBuffers[30])))
	*(*int64)(unsafe.Add(mBase, uint32(v1834)+uint32(_c_F_AsyncReadBuffers[30]))) = v1842 + v1733
	F_pgstat_count_backend_io_op(m, v64, v63, int32(6), v1726, v1733)
	mBase = m.M
	v1847 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[15])) = uint8(v1847)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[16])) = uint8(v1847)
	m.G0 = v1739 + int32(16)
	goto L283
L287:
	;
	v1797 = v64*int32(320) + v63<<(uint(int32(6))%32)
	v1801 = *(*int64)(unsafe.Add(mBase, uint32(v1797)+uint32(_c_F_AsyncReadBuffers[31])))
	*(*int64)(unsafe.Add(mBase, uint32(v1797)+uint32(_c_F_AsyncReadBuffers[31]))) = v1801 + v1750
	v1805 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[32]))
	v1812 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v1805))|base.B2i32(int32(1)<<(uint(v1805)%32)&int32(_a_F_AsyncReadBuffers_31) == v1812) == v1812 {
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
	v1776 = int32(_a_F_AsyncReadBuffers_32)
	v1778 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[33]))
	v1780 = base.I64_div_s(v1750, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[33])) = v1778 + v1780
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
	v1788 = int32(_a_F_AsyncReadBuffers_33)
	v1790 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[34]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[34])) = v1790 + v1750
	goto L287
L296:
	;
	v1783 = int32(_a_F_AsyncReadBuffers_34)
	v1785 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[35]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[35])) = v1785 + v1750
	goto L287
L297:
	;
	v1817 = *(*int64)(unsafe.Add(mBase, uint32(v1797)+uint32(_c_F_AsyncReadBuffers[36])))
	*(*int64)(unsafe.Add(mBase, uint32(v1797)+uint32(_c_F_AsyncReadBuffers[36]))) = v1817 + v1750
	v1821 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[15])) = uint8(v1821)
	*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[37])) = uint8(v1821)
	goto L299
L298:
	;
	goto L299
L299:
	;
	goto L286
L300:
	;
	v1869 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[17])))
	if v1869 == int32(1) {
		goto L304
	} else {
		goto L305
	}
L301:
	;
	v1858 = int32(_a_F_AsyncReadBuffers_35)
	v1860 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[38]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[38])) = v1860 + v1855
	goto L300
L302:
	;
	goto L303
L303:
	;
	v1863 = int32(_a_F_AsyncReadBuffers_36)
	v1865 = *(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[39]))
	*(*int64)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[39])) = v1865 + v1855
	goto L300
L304:
	;
	v1872 = int32(_a_F_AsyncReadBuffers_13)
	v1874 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18]))
	v1876 = *(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[40]))
	*(*int32)(unsafe.Add(mBase, _c_F_AsyncReadBuffers[18])) = v1874 + v1876*v819
	goto L306
L305:
	;
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v819
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
