package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsvector_concat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
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
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v594 int32
	_ = v594
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v717 int32
	_ = v717
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v771 int32
	_ = v771
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v993 int32
	_ = v993
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1176 int32
	_ = v1176
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1224 int32
	_ = v1224
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1289 int32
	_ = v1289
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1326 int32
	_ = v1326
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1387 int32
	_ = v1387
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1442 int32
	_ = v1442
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1476 int32
	_ = v1476
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1641 int32
	_ = v1641
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1753 int32
	_ = v1753
	var v1756 int32
	_ = v1756
	var v1762 int32
	_ = v1762
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	v2 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = F_pg_detoast_datum(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = v29 + int32(8)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v36 = F_pg_detoast_datum(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v38 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v44 = v38
	v47 = v34
	v59 = v2
	goto L7
L5:
	;
	v239 = v2
	goto L6
L6:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v248 = int32(2)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v254 = v245 + v38 + int32(base.Ui32(v247)>>(uint(v248)%32)) + int32(base.Ui32(v251)>>(uint(v248)%32))
	v255 = F_palloc0(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L38
	}
L7:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v65&int32(1) == int32(0) {
		v212 = v59
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v239 = v212
	goto L6
L9:
	;
	v221 = v44 - int32(1)
	if v221 != 0 {
		v44 = v221
		v47 = v47 + int32(4)
		v59 = v212
		goto L7
	} else {
		goto L37
	}
L10:
	;
	v70 = int32(1)
	v81 = v34 + v38<<(uint(int32(2))%32) + (int32(base.Ui32(v65)>>(uint(v70)%32))&int32(2047)+int32(base.Ui32(v65)>>(uint(int32(12))%32))+v70)&int32(_a_F_tsvector_concat_0)
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81))))
	if v82 == int32(0) {
		v212 = v59
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v87 = v82 & int32(3)
	if v87 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v89 = v81
	v92 = int32(0)
	v94 = v82
	v105 = v59
	goto L15
L13:
	;
	v124 = v81
	v129 = v82
	v140 = v59
	goto L14
L14:
	;
	if base.Ui32(v82) < base.Ui32(int32(4)) {
		v212 = v140
		goto L9
	} else {
		goto L21
	}
L15:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+2)))
	v113 = v111 & int32(_a_F_tsvector_concat_1)
	if v113 < v105 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v124 = v117
	v129 = v119
	v140 = v115
	goto L14
L17:
	;
	v115 = v105
	goto L19
L18:
	;
	v115 = v113
	goto L19
L19:
	;
	v117 = v89 + int32(2)
	v118 = int32(1)
	v119 = v94 - v118
	v121 = v92 + v118
	if v121 != v87 {
		v89 = v117
		v92 = v121
		v94 = v119
		v105 = v115
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v149 = v124
	v154 = v129
	v165 = v140
	goto L22
L22:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+2)))
	v173 = v171 & int32(_a_F_tsvector_concat_1)
	if v173 < v165 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v212 = v190
	goto L9
L24:
	;
	v175 = v165
	goto L26
L25:
	;
	v175 = v173
	goto L26
L26:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+4)))
	v178 = v176 & int32(_a_F_tsvector_concat_1)
	if v178 < v175 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v180 = v175
	goto L29
L28:
	;
	v180 = v178
	goto L29
L29:
	;
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+6)))
	v183 = v181 & int32(_a_F_tsvector_concat_1)
	if v183 < v180 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v185 = v180
	goto L32
L31:
	;
	v185 = v183
	goto L32
L32:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+8)))
	v188 = v186 & int32(_a_F_tsvector_concat_1)
	if v188 < v185 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v190 = v185
	goto L35
L34:
	;
	v190 = v188
	goto L35
L35:
	;
	v194 = v154 - int32(4)
	if v194 != 0 {
		v149 = v149 + int32(8)
		v154 = v194
		v165 = v190
		goto L22
	} else {
		goto L36
	}
L36:
	;
	goto L23
L37:
	;
	goto L8
L38:
	;
	v257 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v254 << (uint(v257) % 32)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v262 = v260 + v261
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = v262
	v264 = int32(8)
	v265 = v36 + v264
	v268 = v265 + v245<<(uint(v257)%32)
	v270 = v255 + v264
	v273 = v270 + v262<<(uint(v257)%32)
	v276 = v34 + v38<<(uint(v257)%32)
	v277 = int32(0)
	if base.B2i32(v38 == v277)|base.B2i32(v245 == v277) == v277 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v1266 != 0 {
		goto L216
	} else {
		goto L217
	}
L40:
	;
	v285 = v270
	v287 = v2
	v288 = v34
	v290 = v265
	v295 = v38
	v298 = v245
	goto L43
L41:
	;
	goto L42
L42:
	;
	v1256 = v270
	v1258 = v2
	v1259 = v34
	v1261 = v265
	v1266 = v38
	v1269 = v245
	goto L39
L43:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v308 = int32(1)
	v310 = int32(2047)
	v311 = int32(base.Ui32(v307)>>(uint(v308)%32)) & v310
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v316 = int32(base.Ui32(v312)>>(uint(v308)%32)) & v310
	if v316 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v1256 = v1252
	v1258 = v1243
	v1259 = v1244
	v1261 = v1246
	v1266 = v1249
	v1269 = v1250
	goto L39
L45:
	;
	v1252 = v285 + int32(4)
	if v1249 == int32(0) {
		v1256 = v1252
		v1258 = v1243
		v1259 = v1244
		v1261 = v1246
		v1266 = v1249
		v1269 = v1250
		goto L39
	} else {
		goto L214
	}
L46:
	;
	v1243 = v1232
	v1244 = v288 + int32(4)
	v1246 = v1234
	v1249 = v295 - int32(1)
	v1250 = v1237
	goto L45
L47:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v750 = int32(1)
	v752 = v746&int32(-2) | (v312|v307)&v750
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v752
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v752&int32(-4095) | v756&int32(4094)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v765 = int32(base.Ui32(v761)>>(uint(v750)%32)) & int32(2047)
	if v765 != 0 {
		goto L136
	} else {
		goto L137
	}
L48:
	;
	v1243 = v739
	v1244 = v288
	v1246 = v290 + int32(4)
	v1249 = v295
	v1250 = v298 - int32(1)
	goto L45
L49:
	;
	v729 = int32(1)
	v739 = (v441+v729)&int32(-2) + v625<<(uint(v729)%32) + int32(2)
	goto L48
L50:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v635 = int32(1)
	v637 = v632&int32(-2) | v312&v635
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v637
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v637&int32(-4095) | v641&int32(4094)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v650 = int32(base.Ui32(v646)>>(uint(v635)%32)) & int32(2047)
	if v650 != 0 {
		goto L121
	} else {
		goto L122
	}
L51:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v408 = int32(1)
	v410 = v405&int32(-2) | v307&v408
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v410
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v410&int32(-4095) | v414&int32(4094)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v423 = int32(base.Ui32(v419)>>(uint(v408)%32)) & int32(2047)
	if v423 != 0 {
		goto L86
	} else {
		goto L87
	}
L52:
	;
	if v398 < int32(0) {
		goto L50
	} else {
		goto L84
	}
L53:
	;
	if v311 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if v311 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v321 = int32(-1)
	goto L58
L57:
	;
	v321 = int32(0)
	goto L58
L58:
	;
	v398 = v321
	goto L52
L59:
	;
	v324 = int32(12)
	v326 = v276 + int32(base.Ui32(v312)>>(uint(v324)%32))
	v329 = v268 + int32(base.Ui32(v307)>>(uint(v324)%32))
	v330 = base.B2i32(base.Ui32(v316) < base.Ui32(v311))
	if base.Ui32(v316) < base.Ui32(v311) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v331 = v316
	goto L62
L61:
	;
	v331 = v311
	goto L62
L62:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v331) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	if v393 != 0 {
		v398 = v393
		goto L52
	} else {
		goto L81
	}
L64:
	;
	v393 = int32(0)
	goto L63
L65:
	;
	v367 = v362
	v368 = v363
	v369 = v364
	goto L75
L66:
	;
	if (v326|v329)&int32(3) != 0 {
		v362 = v326
		v363 = v329
		v364 = v331
		goto L65
	} else {
		goto L69
	}
L67:
	;
	v355 = v326
	v356 = v329
	v357 = v331
	goto L68
L68:
	;
	if v357 == int32(0) {
		goto L64
	} else {
		goto L74
	}
L69:
	;
	v339 = v326
	v340 = v329
	v341 = v331
	goto L70
L70:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	if v344 != v345 {
		v362 = v339
		v363 = v340
		v364 = v341
		goto L65
	} else {
		goto L72
	}
L71:
	;
	v355 = v350
	v356 = v348
	v357 = v352
	goto L68
L72:
	;
	v347 = int32(4)
	v348 = v340 + v347
	v350 = v339 + v347
	v352 = v341 - v347
	if base.Ui32(int32(3)) < base.Ui32(v352) {
		v339 = v350
		v340 = v348
		v341 = v352
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v362 = v355
	v363 = v356
	v364 = v357
	goto L65
L75:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v372 == v373 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v393 = v372 - v373
	goto L63
L77:
	;
	v375 = int32(1)
	v380 = v369 - v375
	if v380 != 0 {
		v367 = v367 + v375
		v368 = v368 + v375
		v369 = v380
		goto L75
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L76
L80:
	;
	goto L64
L81:
	;
	if v316 == v311 {
		goto L47
	} else {
		goto L82
	}
L82:
	;
	if v330 == int32(0) {
		goto L51
	} else {
		goto L83
	}
L83:
	;
	goto L50
L84:
	;
	if v398 == int32(0) {
		goto L47
	} else {
		goto L85
	}
L85:
	;
	goto L51
L86:
	;
	base.MemoryCopy(m, v287+v273, v268+int32(base.Ui32(v419)>>(uint(int32(12))%32)), v423)
	goto L88
L87:
	;
	goto L88
L88:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v429&int32(4095) | v287<<(uint(int32(12))%32)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v437 = int32(1)
	v441 = int32(base.Ui32(v436)>>(uint(v437)%32))&int32(2047) + v287
	if v429&v437 == int32(0) {
		v739 = v441
		goto L48
	} else {
		goto L89
	}
L89:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v461 = int32(1)
	v470 = v255 + v454<<(uint(int32(2))%32) + (int32(base.Ui32(v458)>>(uint(int32(12))%32))+int32(base.Ui32(v458)>>(uint(v461)%32))&int32(2047)+v461)&int32(_a_F_tsvector_concat_0)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v471&v461 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	if v625 != 0 {
		goto L49
	} else {
		goto L120
	}
L91:
	;
	v510 = v470 + int32(8)
	if v458&int32(1) != 0 {
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v476 = int32(1)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v506 = (int32(base.Ui32(v471)>>(uint(v476)%32))&int32(2047) + int32(base.Ui32(v471)>>(uint(int32(12))%32)) + v476) & int32(_a_F_tsvector_concat_0)
	v507 = v487
	v508 = int32(0)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v489 = int32(1)
	v499 = (int32(base.Ui32(v471)>>(uint(v489)%32))&int32(2047) + int32(base.Ui32(v471)>>(uint(int32(12))%32)) + v489) & int32(_a_F_tsvector_concat_0)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v505 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v499+(v36+v500<<(uint(int32(2))%32)))+8)))
	v506 = v499
	v507 = v500
	v508 = v505
	goto L91
L95:
	;
	if v508 == int32(0) {
		v612 = v517
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v510))))
	v517 = v513
	goto L95
L97:
	;
	goto L98
L98:
	;
	v514 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v510))) = uint16(v514)
	v517 = v514
	goto L95
L99:
	;
	v625 = v612&int32(_a_F_tsvector_concat_2) - v517
	goto L90
L100:
	;
	if base.Ui32(int32(255)) < base.Ui32(v517) {
		v594 = v517
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if v517 == v594&int32(_a_F_tsvector_concat_2) {
		v612 = v517
		goto L99
	} else {
		goto L119
	}
L102:
	;
	v1768 = int32(10)
	v530 = int32(0)
	v531 = int32(256)
	v532 = v531 - v517
	if base.Ui32(v532) <= base.Ui32(v531) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v536 = v532
	goto L105
L104:
	;
	v536 = v530
	goto L105
L105:
	;
	v538 = v530
	v539 = v517
	v542 = v517
	goto L106
L106:
	;
	if v539 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v594 = v583
	goto L101
L108:
	;
	v553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v510+v539<<(uint(int32(1))%32)))))
	v554 = int32(_a_F_tsvector_concat_1)
	if v553&v554 == v554 {
		v594 = v542
		goto L101
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v558 = int32(1)
	v560 = v470 + v1768 + v539<<(uint(v558)%32)
	v563 = v36 + v507<<(uint(int32(2))%32) + v506 + v1768 + v538<<(uint(v558)%32)
	v564 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563))))
	v566 = v564 & int32(-16384)
	v567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v560))))
	v568 = int32(_a_F_tsvector_concat_1)
	v570 = v566 | v567&v568
	*(*uint16)(unsafe.Add(mBase, uint32(v560))) = uint16(v570)
	v573 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v563))))
	v576 = v239 + v573&v568
	if base.Ui32(v568) <= base.Ui32(v576) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	goto L110
L112:
	;
	v579 = v568
	goto L114
L113:
	;
	v579 = v576
	goto L114
L114:
	;
	v580 = v566 | v579
	*(*uint16)(unsafe.Add(mBase, uint32(v560))) = uint16(v580)
	v582 = int32(1)
	v583 = v539 + v582
	*(*uint16)(unsafe.Add(mBase, uint32(v510))) = uint16(v583)
	v586 = v538 + v582
	if v508 != v586 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if v586 == v536 {
		v594 = v583
		goto L101
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	goto L107
L118:
	;
	v538 = v586
	v539 = v583
	v542 = v583
	goto L106
L119:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v605 | int32(1)
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v510))))
	v612 = v609
	goto L99
L120:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v626 & int32(-2)
	v739 = v441
	goto L48
L121:
	;
	base.MemoryCopy(m, v287+v273, v276+int32(base.Ui32(v646)>>(uint(int32(12))%32)), v650)
	goto L123
L122:
	;
	goto L123
L123:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v656&int32(4095) | v287<<(uint(int32(12))%32)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v664 = int32(1)
	v667 = int32(base.Ui32(v663)>>(uint(v664)%32)) & int32(2047)
	v668 = v667 + v287
	if v656&v664 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v671 = int32(2)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v679 = int32(1)
	v683 = v34 + v672<<(uint(v671)%32) + (int32(base.Ui32(v663)>>(uint(int32(12))%32))+v667+v679)&int32(_a_F_tsvector_concat_0)
	v687 = (v668 + v679) & int32(-2)
	if v663&v679 != 0 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v726 = v668
	goto L126
L126:
	;
	v1232 = v726
	v1234 = v290
	v1237 = v298
	goto L46
L127:
	;
	v690 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v683))))
	v695 = v690<<(uint(int32(1))%32) + int32(2)
	goto L129
L128:
	;
	v695 = v671
	goto L129
L129:
	;
	if v695 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	base.MemoryCopy(m, v687+v273, v683, v695)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v698&int32(1) != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v702 = int32(2)
	v705 = int32(1)
	v717 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v701<<(uint(v702)%32)+(int32(base.Ui32(v698)>>(uint(v705)%32))&int32(2047)+int32(base.Ui32(v698)>>(uint(int32(12))%32))+v705)&int32(_a_F_tsvector_concat_0)))))
	v723 = v717<<(uint(v705)%32) + v702
	goto L135
L134:
	;
	v723 = int32(2)
	goto L135
L135:
	;
	v726 = v723 + v687
	goto L126
L136:
	;
	base.MemoryCopy(m, v287+v273, v276+int32(base.Ui32(v761)>>(uint(int32(12))%32)), v765)
	goto L138
L137:
	;
	goto L138
L138:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v771&int32(4095) | v287<<(uint(int32(12))%32)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v779 = int32(1)
	v782 = int32(base.Ui32(v778)>>(uint(v779)%32)) & int32(2047)
	v783 = v782 + v287
	if v771&v779 == int32(0) {
		v1224 = v783
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v1232 = v1224
	v1234 = v290 + int32(4)
	v1237 = v298 - int32(1)
	goto L46
L140:
	;
	if v778&int32(1) != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v790 = int32(1)
	v793 = (v783 + v790) & int32(-2)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v795 = int32(2)
	v805 = v34 + v794<<(uint(v795)%32) + (int32(base.Ui32(v778)>>(uint(int32(12))%32))+v782+v790)&int32(_a_F_tsvector_concat_0)
	v806 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v805))))
	v810 = v806<<(uint(v790)%32) + v795
	if v810 != 0 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	goto L143
L143:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v1043 = int32(1)
	v1052 = v255 + v1036<<(uint(int32(2))%32) + (int32(base.Ui32(v1040)>>(uint(int32(12))%32))+int32(base.Ui32(v1040)>>(uint(v1043)%32))&int32(2047)+v1043)&int32(_a_F_tsvector_concat_0)
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v1053&v1043 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L144:
	;
	base.MemoryCopy(m, v793+v273, v805, v810)
	goto L146
L145:
	;
	goto L146
L146:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v813&int32(1) != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v817 = int32(2)
	v820 = int32(1)
	v832 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v816<<(uint(v817)%32)+(int32(base.Ui32(v813)>>(uint(v820)%32))&int32(2047)+int32(base.Ui32(v813)>>(uint(int32(12))%32))+v820)&int32(_a_F_tsvector_concat_0)))))
	v838 = v832<<(uint(v820)%32) + v817
	goto L149
L148:
	;
	v838 = int32(2)
	goto L149
L149:
	;
	v839 = v838 + v793
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	if v840&int32(1) == int32(0) {
		v1224 = v839
		goto L139
	} else {
		goto L150
	}
L150:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v860 = int32(1)
	v869 = v255 + v853<<(uint(int32(2))%32) + (int32(base.Ui32(v857)>>(uint(int32(12))%32))+int32(base.Ui32(v857)>>(uint(v860)%32))&int32(2047)+v860)&int32(_a_F_tsvector_concat_0)
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v870&v860 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v1224 = (v1011&int32(_a_F_tsvector_concat_2)-v916)<<(uint(int32(1))%32) + v839
	goto L139
L152:
	;
	v909 = v869 + int32(8)
	if v857&int32(1) != 0 {
		goto L157
	} else {
		goto L158
	}
L153:
	;
	v875 = int32(1)
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v905 = (int32(base.Ui32(v870)>>(uint(v875)%32))&int32(2047) + int32(base.Ui32(v870)>>(uint(int32(12))%32)) + v875) & int32(_a_F_tsvector_concat_0)
	v906 = v886
	v907 = int32(0)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v888 = int32(1)
	v898 = (int32(base.Ui32(v870)>>(uint(v888)%32))&int32(2047) + int32(base.Ui32(v870)>>(uint(int32(12))%32)) + v888) & int32(_a_F_tsvector_concat_0)
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v898+(v36+v899<<(uint(int32(2))%32)))+8)))
	v905 = v898
	v906 = v899
	v907 = v904
	goto L152
L156:
	;
	if v907 == int32(0) {
		v1011 = v916
		goto L160
	} else {
		goto L161
	}
L157:
	;
	v912 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v909))))
	v916 = v912
	goto L156
L158:
	;
	goto L159
L159:
	;
	v913 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v909))) = uint16(v913)
	v916 = v913
	goto L156
L160:
	;
	goto L151
L161:
	;
	if base.Ui32(int32(255)) < base.Ui32(v916) {
		v993 = v916
		goto L162
	} else {
		goto L163
	}
L162:
	;
	if v916 == v993&int32(_a_F_tsvector_concat_2) {
		v1011 = v916
		goto L160
	} else {
		goto L180
	}
L163:
	;
	v1769 = int32(10)
	v929 = int32(0)
	v930 = int32(256)
	v931 = v930 - v916
	if base.Ui32(v931) <= base.Ui32(v930) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v935 = v931
	goto L166
L165:
	;
	v935 = v929
	goto L166
L166:
	;
	v937 = v929
	v938 = v916
	v941 = v916
	goto L167
L167:
	;
	if v938 != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v993 = v982
	goto L162
L169:
	;
	v952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v909+v938<<(uint(int32(1))%32)))))
	v953 = int32(_a_F_tsvector_concat_1)
	if v952&v953 == v953 {
		v993 = v941
		goto L162
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v957 = int32(1)
	v959 = v869 + v1769 + v938<<(uint(v957)%32)
	v962 = v36 + v906<<(uint(int32(2))%32) + v905 + v1769 + v937<<(uint(v957)%32)
	v963 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v962))))
	v965 = v963 & int32(-16384)
	v966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v959))))
	v967 = int32(_a_F_tsvector_concat_1)
	v969 = v965 | v966&v967
	*(*uint16)(unsafe.Add(mBase, uint32(v959))) = uint16(v969)
	v972 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v962))))
	v975 = v239 + v972&v967
	if base.Ui32(v967) <= base.Ui32(v975) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	goto L171
L173:
	;
	v978 = v967
	goto L175
L174:
	;
	v978 = v975
	goto L175
L175:
	;
	v979 = v965 | v978
	*(*uint16)(unsafe.Add(mBase, uint32(v959))) = uint16(v979)
	v981 = int32(1)
	v982 = v938 + v981
	*(*uint16)(unsafe.Add(mBase, uint32(v909))) = uint16(v982)
	v985 = v937 + v981
	if v907 != v985 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	if v985 == v935 {
		v993 = v982
		goto L162
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	goto L168
L179:
	;
	v937 = v985
	v938 = v982
	v941 = v982
	goto L167
L180:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v1004 | int32(1)
	v1008 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v909))))
	v1011 = v1008
	goto L160
L181:
	;
	if v1207 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L182:
	;
	v1092 = v1052 + int32(8)
	if v1040&int32(1) != 0 {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v1058 = int32(1)
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v1088 = (int32(base.Ui32(v1053)>>(uint(v1058)%32))&int32(2047) + int32(base.Ui32(v1053)>>(uint(int32(12))%32)) + v1058) & int32(_a_F_tsvector_concat_0)
	v1089 = v1069
	v1090 = int32(0)
	goto L182
L184:
	;
	goto L185
L185:
	;
	v1071 = int32(1)
	v1081 = (int32(base.Ui32(v1053)>>(uint(v1071)%32))&int32(2047) + int32(base.Ui32(v1053)>>(uint(int32(12))%32)) + v1071) & int32(_a_F_tsvector_concat_0)
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v1087 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1081+(v36+v1082<<(uint(int32(2))%32)))+8)))
	v1088 = v1081
	v1089 = v1082
	v1090 = v1087
	goto L182
L186:
	;
	if v1090 == int32(0) {
		v1194 = v1099
		goto L190
	} else {
		goto L191
	}
L187:
	;
	v1095 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1092))))
	v1099 = v1095
	goto L186
L188:
	;
	goto L189
L189:
	;
	v1096 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1092))) = uint16(v1096)
	v1099 = v1096
	goto L186
L190:
	;
	v1207 = v1194&int32(_a_F_tsvector_concat_2) - v1099
	goto L181
L191:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1099) {
		v1176 = v1099
		goto L192
	} else {
		goto L193
	}
L192:
	;
	if v1099 == v1176&int32(_a_F_tsvector_concat_2) {
		v1194 = v1099
		goto L190
	} else {
		goto L210
	}
L193:
	;
	v1770 = int32(10)
	v1112 = int32(0)
	v1113 = int32(256)
	v1114 = v1113 - v1099
	if base.Ui32(v1114) <= base.Ui32(v1113) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1118 = v1114
	goto L196
L195:
	;
	v1118 = v1112
	goto L196
L196:
	;
	v1120 = v1112
	v1121 = v1099
	v1124 = v1099
	goto L197
L197:
	;
	if v1121 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v1176 = v1165
	goto L192
L199:
	;
	v1135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1092+v1121<<(uint(int32(1))%32)))))
	v1136 = int32(_a_F_tsvector_concat_1)
	if v1135&v1136 == v1136 {
		v1176 = v1124
		goto L192
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1140 = int32(1)
	v1142 = v1052 + v1770 + v1121<<(uint(v1140)%32)
	v1145 = v36 + v1089<<(uint(int32(2))%32) + v1088 + v1770 + v1120<<(uint(v1140)%32)
	v1146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145))))
	v1148 = v1146 & int32(-16384)
	v1149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1142))))
	v1150 = int32(_a_F_tsvector_concat_1)
	v1152 = v1148 | v1149&v1150
	*(*uint16)(unsafe.Add(mBase, uint32(v1142))) = uint16(v1152)
	v1155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145))))
	v1158 = v239 + v1155&v1150
	if base.Ui32(v1150) <= base.Ui32(v1158) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	goto L201
L203:
	;
	v1161 = v1150
	goto L205
L204:
	;
	v1161 = v1158
	goto L205
L205:
	;
	v1162 = v1148 | v1161
	*(*uint16)(unsafe.Add(mBase, uint32(v1142))) = uint16(v1162)
	v1164 = int32(1)
	v1165 = v1121 + v1164
	*(*uint16)(unsafe.Add(mBase, uint32(v1092))) = uint16(v1165)
	v1168 = v1120 + v1164
	if v1090 != v1168 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	if v1168 == v1118 {
		v1176 = v1165
		goto L192
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	goto L198
L209:
	;
	v1120 = v1168
	v1121 = v1165
	v1124 = v1165
	goto L197
L210:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v1187 | int32(1)
	v1191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1092))))
	v1194 = v1191
	goto L190
L211:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = v1210 & int32(-2)
	v1224 = v783
	goto L139
L212:
	;
	goto L213
L213:
	;
	v1214 = int32(1)
	v1224 = (v783+v1214)&int32(-2) + v1207<<(uint(v1214)%32) + int32(2)
	goto L139
L214:
	;
	if v1250 != 0 {
		v285 = v1252
		v287 = v1243
		v288 = v1244
		v290 = v1246
		v295 = v1249
		v298 = v1250
		goto L43
	} else {
		goto L215
	}
L215:
	;
	goto L44
L216:
	;
	v1279 = v1256
	v1281 = v1258
	v1282 = v1259
	v1289 = v1266
	goto L219
L217:
	;
	v1406 = v1256
	v1408 = v1258
	goto L218
L218:
	;
	if v1269 != 0 {
		goto L237
	} else {
		goto L238
	}
L219:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1279)))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	v1305 = int32(1)
	v1307 = v1301&int32(-2) | v1304&v1305
	*(*int32)(unsafe.Add(mBase, uint32(v1279))) = v1307
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	*(*int32)(unsafe.Add(mBase, uint32(v1279))) = v1307&int32(-4095) | v1311&int32(4094)
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	v1320 = int32(base.Ui32(v1316)>>(uint(v1305)%32)) & int32(2047)
	if v1320 != 0 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v1406 = v1402
	v1408 = v1396
	goto L218
L221:
	;
	base.MemoryCopy(m, v1281+v273, v276+int32(base.Ui32(v1316)>>(uint(int32(12))%32)), v1320)
	goto L223
L222:
	;
	goto L223
L223:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1279)))
	*(*int32)(unsafe.Add(mBase, uint32(v1279))) = v1326&int32(4095) | v1281<<(uint(int32(12))%32)
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	v1334 = int32(1)
	v1337 = int32(base.Ui32(v1333)>>(uint(v1334)%32)) & int32(2047)
	v1338 = v1337 + v1281
	if v1326&v1334 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1341 = int32(2)
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v1349 = int32(1)
	v1353 = v34 + v1342<<(uint(v1341)%32) + (int32(base.Ui32(v1333)>>(uint(int32(12))%32))+v1337+v1349)&int32(_a_F_tsvector_concat_0)
	v1357 = (v1338 + v1349) & int32(-2)
	if v1333&v1349 != 0 {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	v1396 = v1338
	goto L226
L226:
	;
	v1399 = int32(4)
	v1402 = v1279 + v1399
	v1404 = v1289 - int32(1)
	if v1404 != 0 {
		v1279 = v1402
		v1281 = v1396
		v1282 = v1282 + v1399
		v1289 = v1404
		goto L219
	} else {
		goto L236
	}
L227:
	;
	v1360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1353))))
	v1365 = v1360<<(uint(int32(1))%32) + int32(2)
	goto L229
L228:
	;
	v1365 = v1341
	goto L229
L229:
	;
	if v1365 != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	base.MemoryCopy(m, v1357+v273, v1353, v1365)
	goto L232
L231:
	;
	goto L232
L232:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1282)))
	if v1368&int32(1) != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v1372 = int32(2)
	v1375 = int32(1)
	v1387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v1371<<(uint(v1372)%32)+(int32(base.Ui32(v1368)>>(uint(v1375)%32))&int32(2047)+int32(base.Ui32(v1368)>>(uint(int32(12))%32))+v1375)&int32(_a_F_tsvector_concat_0)))))
	v1393 = v1387<<(uint(v1375)%32) + v1372
	goto L235
L234:
	;
	v1393 = int32(2)
	goto L235
L235:
	;
	v1396 = v1393 + v1357
	goto L226
L236:
	;
	goto L220
L237:
	;
	v1429 = v1406
	v1431 = v1408
	v1434 = v1261
	v1442 = v1269
	goto L240
L238:
	;
	v1697 = v1406
	v1699 = v1408
	goto L239
L239:
	;
	if v1699 < int32(_a_F_tsvector_concat_3) {
		goto L281
	} else {
		goto L282
	}
L240:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1429)))
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1434)))
	v1455 = int32(1)
	v1457 = v1451&int32(-2) | v1454&v1455
	*(*int32)(unsafe.Add(mBase, uint32(v1429))) = v1457
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1434)))
	*(*int32)(unsafe.Add(mBase, uint32(v1429))) = v1457&int32(-4095) | v1461&int32(4094)
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1434)))
	v1470 = int32(base.Ui32(v1466)>>(uint(v1455)%32)) & int32(2047)
	if v1470 != 0 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1697 = v1693
	v1699 = v1688
	goto L239
L242:
	;
	base.MemoryCopy(m, v1431+v273, v268+int32(base.Ui32(v1466)>>(uint(int32(12))%32)), v1470)
	goto L244
L243:
	;
	goto L244
L244:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1429)))
	*(*int32)(unsafe.Add(mBase, uint32(v1429))) = v1476&int32(4095) | v1431<<(uint(int32(12))%32)
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1434)))
	v1484 = int32(1)
	v1488 = int32(base.Ui32(v1483)>>(uint(v1484)%32))&int32(2047) + v1431
	if v1476&v1484 == int32(0) {
		v1688 = v1488
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1690 = int32(4)
	v1693 = v1429 + v1690
	v1695 = v1442 - int32(1)
	if v1695 != 0 {
		v1429 = v1693
		v1431 = v1688
		v1434 = v1434 + v1690
		v1442 = v1695
		goto L240
	} else {
		goto L280
	}
L246:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1429)))
	v1508 = int32(1)
	v1517 = v255 + v1501<<(uint(int32(2))%32) + (int32(base.Ui32(v1505)>>(uint(int32(12))%32))+int32(base.Ui32(v1505)>>(uint(v1508)%32))&int32(2047)+v1508)&int32(_a_F_tsvector_concat_0)
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1434)))
	if v1518&v1508 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	if v1672 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L248:
	;
	v1557 = v1517 + int32(8)
	if v1505&int32(1) != 0 {
		goto L253
	} else {
		goto L254
	}
L249:
	;
	v1523 = int32(1)
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v1553 = (int32(base.Ui32(v1518)>>(uint(v1523)%32))&int32(2047) + int32(base.Ui32(v1518)>>(uint(int32(12))%32)) + v1523) & int32(_a_F_tsvector_concat_0)
	v1554 = v1534
	v1555 = int32(0)
	goto L248
L250:
	;
	goto L251
L251:
	;
	v1536 = int32(1)
	v1546 = (int32(base.Ui32(v1518)>>(uint(v1536)%32))&int32(2047) + int32(base.Ui32(v1518)>>(uint(int32(12))%32)) + v1536) & int32(_a_F_tsvector_concat_0)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v1552 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1546+(v36+v1547<<(uint(int32(2))%32)))+8)))
	v1553 = v1546
	v1554 = v1547
	v1555 = v1552
	goto L248
L252:
	;
	if v1555 == int32(0) {
		v1659 = v1564
		goto L256
	} else {
		goto L257
	}
L253:
	;
	v1560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1557))))
	v1564 = v1560
	goto L252
L254:
	;
	goto L255
L255:
	;
	v1561 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1557))) = uint16(v1561)
	v1564 = v1561
	goto L252
L256:
	;
	v1672 = v1659&int32(_a_F_tsvector_concat_2) - v1564
	goto L247
L257:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1564) {
		v1641 = v1564
		goto L258
	} else {
		goto L259
	}
L258:
	;
	if v1564 == v1641&int32(_a_F_tsvector_concat_2) {
		v1659 = v1564
		goto L256
	} else {
		goto L276
	}
L259:
	;
	v1771 = int32(10)
	v1577 = int32(0)
	v1578 = int32(256)
	v1579 = v1578 - v1564
	if base.Ui32(v1579) <= base.Ui32(v1578) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1583 = v1579
	goto L262
L261:
	;
	v1583 = v1577
	goto L262
L262:
	;
	v1585 = v1577
	v1586 = v1564
	v1589 = v1564
	goto L263
L263:
	;
	if v1586 != 0 {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v1641 = v1630
	goto L258
L265:
	;
	v1600 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1557+v1586<<(uint(int32(1))%32)))))
	v1601 = int32(_a_F_tsvector_concat_1)
	if v1600&v1601 == v1601 {
		v1641 = v1589
		goto L258
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v1605 = int32(1)
	v1607 = v1517 + v1771 + v1586<<(uint(v1605)%32)
	v1610 = v36 + v1554<<(uint(int32(2))%32) + v1553 + v1771 + v1585<<(uint(v1605)%32)
	v1611 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1610))))
	v1613 = v1611 & int32(-16384)
	v1614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1607))))
	v1615 = int32(_a_F_tsvector_concat_1)
	v1617 = v1613 | v1614&v1615
	*(*uint16)(unsafe.Add(mBase, uint32(v1607))) = uint16(v1617)
	v1620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1610))))
	v1623 = v239 + v1620&v1615
	if base.Ui32(v1615) <= base.Ui32(v1623) {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	goto L267
L269:
	;
	v1626 = v1615
	goto L271
L270:
	;
	v1626 = v1623
	goto L271
L271:
	;
	v1627 = v1613 | v1626
	*(*uint16)(unsafe.Add(mBase, uint32(v1607))) = uint16(v1627)
	v1629 = int32(1)
	v1630 = v1586 + v1629
	*(*uint16)(unsafe.Add(mBase, uint32(v1557))) = uint16(v1630)
	v1633 = v1585 + v1629
	if v1555 != v1633 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	if v1633 == v1583 {
		v1641 = v1630
		goto L258
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	goto L264
L275:
	;
	v1585 = v1633
	v1586 = v1630
	v1589 = v1630
	goto L263
L276:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1429)))
	*(*int32)(unsafe.Add(mBase, uint32(v1429))) = v1652 | int32(1)
	v1656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1557))))
	v1659 = v1656
	goto L256
L277:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v1429)))
	*(*int32)(unsafe.Add(mBase, uint32(v1429))) = v1675 & int32(-2)
	v1688 = v1488
	goto L245
L278:
	;
	goto L279
L279:
	;
	v1679 = int32(1)
	v1688 = (v1488+v1679)&int32(-2) + v1672<<(uint(v1679)%32) + int32(2)
	goto L245
L280:
	;
	goto L241
L281:
	;
	v1721 = v1697 - v270
	v1722 = int32(2)
	v1723 = v1721 >> (uint(v1722) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = v1723
	if v1697 != v273 {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L1
	} else {
		goto L298
	}
L284:
	;
	if v1699 != 0 {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	v1731 = v1723
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v1699<<(uint(v1722)%32) + v1731<<(uint(int32(4))%32) + int32(32)
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1738 != v29 {
		goto L290
	} else {
		goto L291
	}
L287:
	;
	base.MemoryCopy(m, v1721+v270, v273, v1699)
	goto L289
L288:
	;
	goto L289
L289:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v1731 = v1730
	goto L286
L290:
	;
	F_pfree(m, v29)
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L1
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1742 != v36 {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	goto L292
L294:
	;
	F_pfree(m, v36)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	m.G0 = v26 + int32(16)
	return v255
L297:
	;
	goto L296
L298:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(_a_F_tsvector_concat_4)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v1699
	F_errmsg(m, int32(_a_F_tsvector_concat_5), v26)
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	F_errfinish(m, int32(_a_F_tsvector_concat_6), int32(1126), int32(_a_F_tsvector_concat_7))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v26 = int32(2)
	v27 = int32(base.Ui32(v25) >> (uint(v26) % 32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v30 = int32(base.Ui32(v28) >> (uint(v26) % 32))
	if base.Ui32(v27) < base.Ui32(v30) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v221 != v6 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v220 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if base.Ui32(v30) < base.Ui32(v27) {
		v195 = v33
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v220 = v195
	goto L4
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v35 < v36 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v220 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v36 < v35 {
		v195 = v33
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v35 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v220 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v43 = int32(8)
	v44 = v11 + v43
	v45 = int32(2)
	v47 = v44 + v36<<(uint(v45)%32)
	v49 = v6 + v43
	v52 = v49 + v35<<(uint(v45)%32)
	v61 = v44
	v62 = v49
	v66 = int32(0)
	goto L17
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v68 = int32(1)
	v69 = v67 & v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v72 = v70 & v68
	if v69 != v72 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v195 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v72) < base.Ui32(v69) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v78 = int32(1)
	v80 = int32(2047)
	v81 = int32(base.Ui32(v70)>>(uint(v78)%32)) & v80
	v82 = int32(12)
	v83 = int32(base.Ui32(v70) >> (uint(v82) % 32))
	v85 = int32(base.Ui32(v67) >> (uint(v82) % 32))
	v89 = int32(base.Ui32(v67)>>(uint(v78)%32)) & v80
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v77 = int32(-1)
	goto L24
L23:
	;
	v77 = int32(1)
	goto L24
L24:
	;
	v220 = v77
	goto L4
L25:
	;
	if v69 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v81 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v81 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v220 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v95 = base.B2i32(base.Ui32(v89) < base.Ui32(v81))
	if base.Ui32(v89) < base.Ui32(v81) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v96 = v89
	goto L34
L33:
	;
	v96 = v81
	goto L34
L34:
	;
	v97 = F_memcmp(m, v85+v52, v83+v47, v96)
	mBase = m.M
	if v97 != 0 {
		v195 = v97
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v81 == v89 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v89) < base.Ui32(v81) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = int32(-1)
	goto L39
L38:
	;
	v101 = int32(1)
	goto L39
L39:
	;
	v220 = v101
	goto L4
L40:
	;
	v220 = int32(-1)
	goto L4
L41:
	;
	v184 = int32(4)
	v190 = v66 + int32(1)
	if v190 != v35 {
		v61 = v61 + v184
		v62 = v62 + v184
		v66 = v190
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v110 = int32(1)
	v112 = int32(_a_F_tsvector_eq_0)
	v114 = v47 + (v81+v83+v110)&v112
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
	v121 = v52 + (v89+v85+v110)&v112
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121))))
	if v115 == v122 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v129 = v121
	v130 = v114
	v131 = int32(0)
	goto L51
L44:
	;
	if v122 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v115) < base.Ui32(v122) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v128 = int32(-1)
	goto L50
L49:
	;
	v128 = int32(1)
	goto L50
L50:
	;
	v220 = v128
	goto L4
L51:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+2)))
	v144 = int32(_a_F_tsvector_eq_1)
	v145 = v143 & v144
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+2)))
	v148 = v146 & v144
	if v145 != v148 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v157) < base.Ui32(v155) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v148) < base.Ui32(v145) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v154 = int32(14)
	v155 = int32(base.Ui32(v143) >> (uint(v154) % 32))
	v157 = int32(base.Ui32(v146) >> (uint(v154) % 32))
	if v155 == v157 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v153 = int32(-1)
	goto L58
L57:
	;
	v153 = int32(1)
	goto L58
L58:
	;
	v220 = v153
	goto L4
L59:
	;
	v159 = int32(2)
	v164 = v131 + int32(1)
	if v164 == v122 {
		goto L41
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L52
L62:
	;
	v129 = v129 + v159
	v130 = v130 + v159
	v131 = v164
	goto L51
L63:
	;
	v169 = int32(-1)
	goto L65
L64:
	;
	v169 = int32(1)
	goto L65
L65:
	;
	v195 = v169
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v6)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v225 != v11 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v11)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return base.B2i32(v220 == int32(0))
L74:
	;
	goto L73
}
func F_tsvector_length(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v10 != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v9
			}
		} else {
			return v9
		}
	}
}
func F_tsvector_setweight(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		switch v24&int32(255) - int32(65) {
		case 0, 32:
			v46 = int32(_a_F_tsvector_setweight_0)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v50 = F_palloc(m, int32(base.Ui32(v47)>>(uint(int32(2))%32)))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v54 = int32(base.Ui32(v52) >> (uint(int32(2)) % 32))
				if v54 != 0 {
					base.MemoryCopy(m, v50, v19, v54)
				} else {
				}
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
				if v56 != 0 {
					v58 = v50 + int32(8)
					v66 = v58
					v67 = v56
					for {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
						if v72&int32(1) == int32(0) {
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							v81 = int32(1)
							v92 = v58 + v77<<(uint(int32(2))%32) + (int32(base.Ui32(v72)>>(uint(v81)%32))&int32(2047)+int32(base.Ui32(v72)>>(uint(int32(12))%32))+v81)&int32(_a_F_tsvector_setweight_1)
							v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92))))
							if v93 == int32(0) {
							} else {
								v98 = v93 & int32(3)
								if v98 != 0 {
									v100 = v92
									v101 = v93
									v109 = int32(0)
									for {
										v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
										v115 = v112&int32(_a_F_tsvector_setweight_2) | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)) = uint16(v115)
										v117 = int32(1)
										v118 = v101 - v117
										v120 = v100 + int32(2)
										v122 = v109 + v117
										if v122 != v98 {
											v100 = v120
											v101 = v118
											v109 = v122
											continue
										} else {
											break
										}
										break
									}
									v125 = v120
									v126 = v118
								} else {
									v125 = v92
									v126 = v93
								}
								if base.Ui32(v93) < base.Ui32(int32(4)) {
								} else {
									v140 = v125
									v141 = v126
									for {
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)))
										v153 = int32(_a_F_tsvector_setweight_2)
										v155 = v152&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v155)
										v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)))
										v160 = v157&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v160)
										v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+6)))
										v165 = v162&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+6)) = uint16(v165)
										v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+8)))
										v170 = v167&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+8)) = uint16(v170)
										v175 = v141 - int32(4)
										if v175 != 0 {
											v140 = v140 + int32(8)
											v141 = v175
											continue
										} else {
											break
										}
										break
									}
								}
							}
						}
						v192 = v67 - int32(1)
						if v192 != 0 {
							v66 = v66 + int32(4)
							v67 = v192
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v206 != v19 {
					F_pfree(m, v19)
					mBase = m.M
					v209 = m.ExcPending
					if v209 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(16)
						return v50
					}
				} else {
					m.G0 = v16 + int32(16)
					return v50
				}
			}
		case 1, 33:
			v46 = int32(_a_F_tsvector_setweight_3)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v50 = F_palloc(m, int32(base.Ui32(v47)>>(uint(int32(2))%32)))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v54 = int32(base.Ui32(v52) >> (uint(int32(2)) % 32))
				if v54 != 0 {
					base.MemoryCopy(m, v50, v19, v54)
				} else {
				}
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
				if v56 != 0 {
					v58 = v50 + int32(8)
					v66 = v58
					v67 = v56
					for {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
						if v72&int32(1) == int32(0) {
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							v81 = int32(1)
							v92 = v58 + v77<<(uint(int32(2))%32) + (int32(base.Ui32(v72)>>(uint(v81)%32))&int32(2047)+int32(base.Ui32(v72)>>(uint(int32(12))%32))+v81)&int32(_a_F_tsvector_setweight_1)
							v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92))))
							if v93 == int32(0) {
							} else {
								v98 = v93 & int32(3)
								if v98 != 0 {
									v100 = v92
									v101 = v93
									v109 = int32(0)
									for {
										v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
										v115 = v112&int32(_a_F_tsvector_setweight_2) | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)) = uint16(v115)
										v117 = int32(1)
										v118 = v101 - v117
										v120 = v100 + int32(2)
										v122 = v109 + v117
										if v122 != v98 {
											v100 = v120
											v101 = v118
											v109 = v122
											continue
										} else {
											break
										}
										break
									}
									v125 = v120
									v126 = v118
								} else {
									v125 = v92
									v126 = v93
								}
								if base.Ui32(v93) < base.Ui32(int32(4)) {
								} else {
									v140 = v125
									v141 = v126
									for {
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)))
										v153 = int32(_a_F_tsvector_setweight_2)
										v155 = v152&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v155)
										v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)))
										v160 = v157&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v160)
										v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+6)))
										v165 = v162&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+6)) = uint16(v165)
										v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+8)))
										v170 = v167&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+8)) = uint16(v170)
										v175 = v141 - int32(4)
										if v175 != 0 {
											v140 = v140 + int32(8)
											v141 = v175
											continue
										} else {
											break
										}
										break
									}
								}
							}
						}
						v192 = v67 - int32(1)
						if v192 != 0 {
							v66 = v66 + int32(4)
							v67 = v192
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v206 != v19 {
					F_pfree(m, v19)
					mBase = m.M
					v209 = m.ExcPending
					if v209 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(16)
						return v50
					}
				} else {
					m.G0 = v16 + int32(16)
					return v50
				}
			}
		case 2, 34:
			v46 = int32(_a_F_tsvector_setweight_4)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v50 = F_palloc(m, int32(base.Ui32(v47)>>(uint(int32(2))%32)))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v54 = int32(base.Ui32(v52) >> (uint(int32(2)) % 32))
				if v54 != 0 {
					base.MemoryCopy(m, v50, v19, v54)
				} else {
				}
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
				if v56 != 0 {
					v58 = v50 + int32(8)
					v66 = v58
					v67 = v56
					for {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
						if v72&int32(1) == int32(0) {
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							v81 = int32(1)
							v92 = v58 + v77<<(uint(int32(2))%32) + (int32(base.Ui32(v72)>>(uint(v81)%32))&int32(2047)+int32(base.Ui32(v72)>>(uint(int32(12))%32))+v81)&int32(_a_F_tsvector_setweight_1)
							v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92))))
							if v93 == int32(0) {
							} else {
								v98 = v93 & int32(3)
								if v98 != 0 {
									v100 = v92
									v101 = v93
									v109 = int32(0)
									for {
										v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
										v115 = v112&int32(_a_F_tsvector_setweight_2) | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)) = uint16(v115)
										v117 = int32(1)
										v118 = v101 - v117
										v120 = v100 + int32(2)
										v122 = v109 + v117
										if v122 != v98 {
											v100 = v120
											v101 = v118
											v109 = v122
											continue
										} else {
											break
										}
										break
									}
									v125 = v120
									v126 = v118
								} else {
									v125 = v92
									v126 = v93
								}
								if base.Ui32(v93) < base.Ui32(int32(4)) {
								} else {
									v140 = v125
									v141 = v126
									for {
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)))
										v153 = int32(_a_F_tsvector_setweight_2)
										v155 = v152&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v155)
										v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)))
										v160 = v157&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v160)
										v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+6)))
										v165 = v162&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+6)) = uint16(v165)
										v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+8)))
										v170 = v167&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+8)) = uint16(v170)
										v175 = v141 - int32(4)
										if v175 != 0 {
											v140 = v140 + int32(8)
											v141 = v175
											continue
										} else {
											break
										}
										break
									}
								}
							}
						}
						v192 = v67 - int32(1)
						if v192 != 0 {
							v66 = v66 + int32(4)
							v67 = v192
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v206 != v19 {
					F_pfree(m, v19)
					mBase = m.M
					v209 = m.ExcPending
					if v209 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(16)
						return v50
					}
				} else {
					m.G0 = v16 + int32(16)
					return v50
				}
			}
		case 3, 35:
			v46 = int32(0)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v50 = F_palloc(m, int32(base.Ui32(v47)>>(uint(int32(2))%32)))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v54 = int32(base.Ui32(v52) >> (uint(int32(2)) % 32))
				if v54 != 0 {
					base.MemoryCopy(m, v50, v19, v54)
				} else {
				}
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
				if v56 != 0 {
					v58 = v50 + int32(8)
					v66 = v58
					v67 = v56
					for {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
						if v72&int32(1) == int32(0) {
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							v81 = int32(1)
							v92 = v58 + v77<<(uint(int32(2))%32) + (int32(base.Ui32(v72)>>(uint(v81)%32))&int32(2047)+int32(base.Ui32(v72)>>(uint(int32(12))%32))+v81)&int32(_a_F_tsvector_setweight_1)
							v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92))))
							if v93 == int32(0) {
							} else {
								v98 = v93 & int32(3)
								if v98 != 0 {
									v100 = v92
									v101 = v93
									v109 = int32(0)
									for {
										v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
										v115 = v112&int32(_a_F_tsvector_setweight_2) | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)) = uint16(v115)
										v117 = int32(1)
										v118 = v101 - v117
										v120 = v100 + int32(2)
										v122 = v109 + v117
										if v122 != v98 {
											v100 = v120
											v101 = v118
											v109 = v122
											continue
										} else {
											break
										}
										break
									}
									v125 = v120
									v126 = v118
								} else {
									v125 = v92
									v126 = v93
								}
								if base.Ui32(v93) < base.Ui32(int32(4)) {
								} else {
									v140 = v125
									v141 = v126
									for {
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)))
										v153 = int32(_a_F_tsvector_setweight_2)
										v155 = v152&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v155)
										v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)))
										v160 = v157&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v160)
										v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+6)))
										v165 = v162&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+6)) = uint16(v165)
										v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+8)))
										v170 = v167&v153 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v140)+8)) = uint16(v170)
										v175 = v141 - int32(4)
										if v175 != 0 {
											v140 = v140 + int32(8)
											v141 = v175
											continue
										} else {
											break
										}
										break
									}
								}
							}
						}
						v192 = v67 - int32(1)
						if v192 != 0 {
							v66 = v66 + int32(4)
							v67 = v192
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v206 != v19 {
					F_pfree(m, v19)
					mBase = m.M
					v209 = m.ExcPending
					if v209 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(16)
						return v50
					}
				} else {
					m.G0 = v16 + int32(16)
					return v50
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = base.I32_extend8_s(v24)
				F_errmsg_internal(m, int32(_a_F_tsvector_setweight_5), v16)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_tsvector_setweight_6), int32(242), int32(_a_F_tsvector_setweight_7))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
