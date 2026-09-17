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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
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
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
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
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int64
	_ = v542
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v626 float64
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 float64
	_ = v633
	var v636 int64
	_ = v636
	var v639 int64
	_ = v639
	var v644 int64
	_ = v644
	var v646 int64
	_ = v646
	var v651 int64
	_ = v651
	var v653 int64
	_ = v653
	var v657 int64
	_ = v657
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v820 int32
	_ = v820
	var v832 int32
	_ = v832
	var v845 int32
	_ = v845
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v937 int32
	_ = v937
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v991 int32
	_ = v991
	var v1000 int32
	_ = v1000
	var v1013 int32
	_ = v1013
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1140 int32
	_ = v1140
	var v1169 int32
	_ = v1169
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1293 int32
	_ = v1293
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1385 int32
	_ = v1385
	var v1395 int32
	_ = v1395
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1445 int32
	_ = v1445
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1464 int32
	_ = v1464
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1504 int32
	_ = v1504
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1556 int32
	_ = v1556
	var v1561 int32
	_ = v1561
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
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1611 int64
	_ = v1611
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1643 int32
	_ = v1643
	var v1645 int64
	_ = v1645
	var v1646 int64
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1650 int64
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1665 int32
	_ = v1665
	var v1671 int32
	_ = v1671
	var v1676 int32
	_ = v1676
	var v1686 int32
	_ = v1686
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1782 int32
	_ = v1782
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1799 int32
	_ = v1799
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1890 int32
	_ = v1890
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1902 int64
	_ = v1902
	var v1921 int32
	_ = v1921
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1934 int32
	_ = v1934
	var v1936 int64
	_ = v1936
	var v1937 int64
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1941 int64
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1956 int32
	_ = v1956
	var v1962 int32
	_ = v1962
	var v1967 int32
	_ = v1967
	var v1977 int32
	_ = v1977
	var v1990 int32
	_ = v1990
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
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
	m.G0 = v2007 + int32(80)
	return v2008
L5:
	;
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1241)+59)))
	if v1261|base.B2i32(v1246&int32(4) == int32(0)) != 0 {
		goto L270
	} else {
		goto L271
	}
L6:
	;
	v1241 = v37
	v1246 = v9
	v1251 = int32(2)
	v1252 = v9
	v1259 = v9
	v1261 = v9
	v1263 = v9
	v1264 = v9
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
	v80 = v9
	v81 = v9
	v83 = v9
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
	v104 = v83 << (uint(int32(2)) % 32)
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
		v2007 = v37
		v2008 = v102
		goto L4
	}
L10:
	;
	if v1207 != 0 {
		v2007 = v37
		v2008 = int32(-1)
		goto L4
	} else {
		goto L268
	}
L11:
	;
	v1230 = v83 + int32(1)
	if v1230 != l2 {
		v80 = v1207
		v81 = v1208
		v83 = v1230
		v86 = v1213
		v87 = v1214
		v94 = v1221
		v96 = v1223
		v98 = v1225
		v99 = v1226
		goto L9
	} else {
		goto L267
	}
L12:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	if v1191&v81 != 0 {
		goto L264
	} else {
		goto L265
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(32)
	v1169 = v80
	v1175 = v86
	v1176 = v1140
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1140 = v281
	goto L13
L15:
	;
	v757 = l0 + v104
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	v765 = F_DecodeTimezoneAbbrev(m, v83, v758, v37-int32(-64), v37+int32(60), v37+int32(52), l7)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L27
	} else {
		goto L192
	}
L16:
	;
	if v80 != 0 {
		goto L106
	} else {
		goto L107
	}
L17:
	;
	if l6 == int32(0) {
		v2007 = v37
		v2008 = v102
		goto L4
	} else {
		goto L80
	}
L18:
	;
	switch v80 {
	case 0, 3:
		goto L74
	default:
		v2007 = v37
		v2008 = v102
		goto L4
	}
L19:
	;
	if l6 == int32(0) {
		v2007 = v37
		v2008 = v102
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if v83|v66 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v127 = l0 + v104
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if base.Ui32((v129-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v110 != int32(2) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v113 != int32(3) {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v121 = F_DecodeDate(m, v116, v81, v37+int32(68), v37+int32(59), l4)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	return int32(0)
L28:
	;
	if v121 == int32(0) {
		v1169 = v80
		v1175 = v86
		v1176 = v87
		v1183 = v94
		v1185 = v96
		v1187 = v98
		v1188 = v99
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v2007 = v37
	v2008 = v121
	goto L4
L30:
	;
	v136 = int32(_a_F_DecodeTimeOnly_0)
	if v81&v136 == v136 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v281 = F_pg_tzset(m, v128)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L27
	} else {
		goto L72
	}
L33:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L34:
	;
	goto L35
L35:
	;
	v141 = int32(45)
	v142 = F___strchrnul(m, v128, v141)
	mBase = m.M
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v144 == v141 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v148 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v148 = v142
	goto L39
L38:
	;
	v148 = int32(0)
	goto L39
L39:
	;
	goto L36
L40:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L41:
	;
	goto L42
L42:
	;
	v152 = int32(0)
	v159 = m.G0
	v161 = v159 - int32(16)
	m.G0 = v161
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	switch v164 - int32(43) {
	case 0, 2:
		goto L45
	default:
		v254 = int32(-1)
		goto L44
	}
L43:
	;
	if v254 != 0 {
		v2007 = v37
		v2008 = v254
		goto L4
	} else {
		goto L67
	}
L44:
	;
	m.G0 = v161 + int32(16)
	goto L43
L45:
	;
	v167 = int32(_a_F_DecodeTimeOnly_1)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0])) = int32(0)
	v174 = F_strtoint(m, v148+int32(1), v161+int32(12))
	mBase = m.M
	v175 = int32(-5)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0]))
	if v177 == int32(68) {
		v254 = v175
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v181 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v227 = int32(59)
	if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v221))|base.B2i32(base.Ui32(v227) < base.Ui32(v220))|base.B2i32(base.Ui32(v227) < base.Ui32(v223)) != 0 {
		v254 = v175
		goto L44
	} else {
		goto L60
	}
L48:
	;
	if v181 != int32(58) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v211 = F_strlen(m, v148)
	mBase = m.M
	if base.Ui32(v211) < base.Ui32(int32(4)) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v220 = int32(0)
	v221 = v174
	v223 = v152
	goto L47
L52:
	;
	goto L53
L53:
	;
	v185 = int32(_a_F_DecodeTimeOnly_1)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0])) = int32(0)
	v191 = v161 + int32(12)
	v192 = F_strtoint(m, v180+int32(1), v191)
	mBase = m.M
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0]))
	if v194 == int32(68) {
		v254 = v175
		goto L44
	} else {
		goto L54
	}
L54:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v198 != int32(58) {
		v220 = v192
		v221 = v174
		v223 = v152
		goto L47
	} else {
		goto L55
	}
L55:
	;
	v201 = int32(_a_F_DecodeTimeOnly_1)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0])) = int32(0)
	v206 = F_strtoint(m, v197+int32(1), v191)
	mBase = m.M
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0]))
	if v208 != int32(68) {
		v220 = v192
		v221 = v174
		v223 = v206
		goto L47
	} else {
		goto L56
	}
L56:
	;
	v254 = v175
	goto L44
L57:
	;
	v220 = int32(0)
	v221 = v174
	v223 = v152
	goto L47
L58:
	;
	goto L59
L59:
	;
	v215 = int32(100)
	v216 = base.I32_div_s(v174, v215)
	v220 = v174 - v216*v215
	v221 = v216
	v223 = v152
	goto L47
L60:
	;
	v233 = int32(60)
	v238 = (v221*v233+v220)*v233 + v223
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v241 == int32(45) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v244 = v238
	goto L63
L62:
	;
	v244 = int32(0) - v238
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v244
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v249 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v250 = int32(-1)
	goto L66
L65:
	;
	v250 = int32(0)
	goto L66
L66:
	;
	v254 = v250
	goto L44
L67:
	;
	v261 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v261)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v264 = F_strlen(m, v263)
	mBase = m.M
	v271 = F_DecodeNumberField(m, v264, v263, v81|int32(14), v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L27
	} else {
		goto L68
	}
L68:
	;
	if v271 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v271
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v277 | int32(32)
	v1169 = v80
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L72:
	;
	if v281 != 0 {
		goto L14
	} else {
		goto L73
	}
L73:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v283
	v2007 = v37
	v2008 = int32(-6)
	goto L4
L74:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v293 = F_DecodeTimeCommon(m, v287, int32(_a_F_DecodeTimeOnly_2), v37+int32(68), v37+int32(8))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L27
	} else {
		goto L75
	}
L75:
	;
	if v293 != 0 {
		v2007 = v37
		v2008 = v293
		goto L4
	} else {
		goto L76
	}
L76:
	;
	v295 = *(*int64)(unsafe.Add(mBase, uint32(v37)+24))
	if int64(2147483648) <= v295 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v2007 = v37
	v2008 = int32(-2)
	goto L4
L78:
	;
	goto L79
L79:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l4)+8)) = uint32(v295)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v300
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v304
	v1169 = int32(0)
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L80:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v313 = int32(0)
	v320 = m.G0
	v322 = v320 - int32(16)
	m.G0 = v322
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	switch v325 - int32(43) {
	case 0, 2:
		goto L83
	default:
		v415 = int32(-1)
		goto L82
	}
L81:
	;
	if v415 != 0 {
		v2007 = v37
		v2008 = v415
		goto L4
	} else {
		goto L105
	}
L82:
	;
	m.G0 = v322 + int32(16)
	goto L81
L83:
	;
	v328 = int32(_a_F_DecodeTimeOnly_1)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0])) = int32(0)
	v335 = F_strtoint(m, v310+int32(1), v322+int32(12))
	mBase = m.M
	v336 = int32(-5)
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0]))
	if v338 == int32(68) {
		v415 = v336
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v342 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v388 = int32(59)
	if base.B2i32(base.Ui32(int32(15)) < base.Ui32(v382))|base.B2i32(base.Ui32(v388) < base.Ui32(v381))|base.B2i32(base.Ui32(v388) < base.Ui32(v384)) != 0 {
		v415 = v336
		goto L82
	} else {
		goto L98
	}
L86:
	;
	if v342 != int32(58) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v372 = F_strlen(m, v310)
	mBase = m.M
	if base.Ui32(v372) < base.Ui32(int32(4)) {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v381 = int32(0)
	v382 = v335
	v384 = v313
	goto L85
L90:
	;
	goto L91
L91:
	;
	v346 = int32(_a_F_DecodeTimeOnly_1)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0])) = int32(0)
	v352 = v322 + int32(12)
	v353 = F_strtoint(m, v341+int32(1), v352)
	mBase = m.M
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0]))
	if v355 == int32(68) {
		v415 = v336
		goto L82
	} else {
		goto L92
	}
L92:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
	if v359 != int32(58) {
		v381 = v353
		v382 = v335
		v384 = v313
		goto L85
	} else {
		goto L93
	}
L93:
	;
	v362 = int32(_a_F_DecodeTimeOnly_1)
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0])) = int32(0)
	v367 = F_strtoint(m, v358+int32(1), v352)
	mBase = m.M
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0]))
	if v369 != int32(68) {
		v381 = v353
		v382 = v335
		v384 = v367
		goto L85
	} else {
		goto L94
	}
L94:
	;
	v415 = v336
	goto L82
L95:
	;
	v381 = int32(0)
	v382 = v335
	v384 = v313
	goto L85
L96:
	;
	goto L97
L97:
	;
	v376 = int32(100)
	v377 = base.I32_div_s(v335, v376)
	v381 = v335 - v377*v376
	v382 = v377
	v384 = v313
	goto L85
L98:
	;
	v394 = int32(60)
	v399 = (v382*v394+v381)*v394 + v384
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if v402 == int32(45) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v405 = v399
	goto L101
L100:
	;
	v405 = int32(0) - v399
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+int32(8)))) = v405
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if v410 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v411 = int32(-1)
	goto L104
L103:
	;
	v411 = int32(0)
	goto L104
L104:
	;
	v415 = v411
	goto L82
L105:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v422
	v1140 = v87
	goto L13
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0])) = int32(0)
	v427 = l0 + v104
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v432 = F_strtol(m, v428, v37+int32(72), int32(10))
	mBase = m.M
	goto L109
L107:
	;
	goto L108
L108:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v692 = F_strlen(m, v691)
	mBase = m.M
	v693 = int32(46)
	v694 = F___strchrnul(m, v691, v693)
	mBase = m.M
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694))))
	if v696 == v693 {
		goto L165
	} else {
		goto L166
	}
L109:
	;
	v433 = int32(-2)
	v435 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0]))
	if v435 == int32(68) {
		v2007 = v37
		v2008 = v433
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	v442 = int32(0)
	if base.B2i32(v439 == int32(46))|base.B2i32(v439 == v442) == v442 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L112:
	;
	goto L113
L113:
	;
	if v80 != int32(3) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	v1169 = int32(0)
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v685
	v1187 = v98
	v1188 = v99
	goto L12
L115:
	;
	if v80 != int32(31) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	v664 = F_strlen(m, v663)
	mBase = m.M
	v671 = F_DecodeNumberField(m, v664, v663, v81|int32(14), v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L27
	} else {
		goto L159
	}
L118:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L119:
	;
	goto L120
L120:
	;
	if l6 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L122:
	;
	goto L123
L123:
	;
	if v432 < int32(0) {
		v2007 = v37
		v2008 = v433
		goto L4
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(14)
	v461 = v432 + int32(_a_F_DecodeTimeOnly_3)
	v462 = int32(_a_F_DecodeTimeOnly_4)
	v463 = base.I32_div_u_s(v461, v462)
	v464 = int32(3)
	v470 = int32(2)
	v475 = base.I32_div_u_s((v463*int32(1073595727)+v461)<<(uint(v470)%32)|v464, v462)
	v478 = v432 + v463*v464 + v475 + int32(_a_F_DecodeTimeOnly_5)
	v479 = int32(1461)
	v480 = base.I32_div_u_s(v478, v479)
	v483 = v480*int32(-1461) + v478
	v485 = v483 << (uint(v470) % 32)
	if base.Ui32(v479) <= base.Ui32(v485) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v498 = base.I32_div_u_s(v485, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v498 + v480<<(uint(int32(2))%32) - int32(_a_F_DecodeTimeOnly_6)
	v505 = int32(1)
	v507 = v496 + int32(123)
	v511 = int32(base.Ui32(v507*int32(2141)) >> (uint(int32(16)) % 32))
	v515 = base.I32_rem_u_s(v511+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v515 + v505
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v507 - int32(base.Ui32(v511*int32(_a_F_DecodeTimeOnly_7))>>(uint(int32(8))%32))
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v525 != int32(46) {
		v685 = v505
		goto L114
	} else {
		goto L129
	}
L126:
	;
	v491 = base.I32_rem_u_s(v483+int32(305), int32(365))
	v496 = v491
	goto L125
L127:
	;
	goto L128
L128:
	;
	v495 = base.I32_rem_u_s(v483+int32(306), int32(366))
	v496 = v495
	goto L125
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v438
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+1)))
	if v529 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L131:
	;
	v636 = base.I64_trunc_sat_f64_s(base.F64_mul(v633, float64(8.64e+10)))
	v639 = base.I64_div_s(v636, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v54))) = uint32(v639)
	v644 = base.I64_extend32_s(v639)*int64(-3600000000) + v636
	v646 = base.I64_div_s(v644, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4+v58))) = uint32(v646)
	v651 = base.I64_extend32_s(v646)*int64(-60000000) + v644
	v653 = base.I64_div_s(v651, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4))) = uint32(v653)
	v657 = v653*int64(4293967296) + v651
	*(*uint32)(unsafe.Add(mBase, uint32(l5))) = uint32(v657)
	goto L158
L132:
	;
	v633 = float64(0)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v534 = v438 + int32(1)
	v535 = int32(_a_F_DecodeTimeOnly_8)
	v539 = m.G0
	v541 = v539 - int32(32)
	v542 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v541)+24)) = v542
	*(*int64)(unsafe.Add(mBase, uint32(v541)+16)) = v542
	*(*int64)(unsafe.Add(mBase, uint32(v541)+8)) = v542
	*(*int64)(unsafe.Add(mBase, uint32(v541))) = v542
	v550 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[1])))
	if v550 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v619 = F_strlen(m, v534)
	mBase = m.M
	if v618 != v619 {
		goto L130
	} else {
		goto L154
	}
L136:
	;
	v618 = int32(0)
	goto L135
L137:
	;
	goto L138
L138:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[2])))
	if v554 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v558 = v534
	goto L142
L140:
	;
	goto L141
L141:
	;
	v568 = v535
	v569 = v550
	goto L145
L142:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	if v564 == v550 {
		v558 = v558 + int32(1)
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v618 = v558 - v534
	goto L135
L144:
	;
	goto L143
L145:
	;
	v576 = v541 + int32(base.Ui32(v569)>>(uint(int32(3))%32))&int32(28)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	v578 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v576))) = v577 | v578<<(uint(v569)%32)
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568)+1)))
	if v582 != 0 {
		v568 = v568 + v578
		v569 = v582
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	if v585 == int32(0) {
		v608 = v534
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L146
L148:
	;
	v618 = v608 - v534
	goto L135
L149:
	;
	v589 = v534
	v590 = v585
	goto L150
L150:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v541+int32(base.Ui32(v590)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v598)>>(uint(v590)%32))&int32(1) == int32(0) {
		v608 = v589
		goto L148
	} else {
		goto L152
	}
L151:
	;
	v608 = v606
	goto L148
L152:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+1)))
	v606 = v589 + int32(1)
	if v604 != 0 {
		v589 = v606
		v590 = v604
		goto L150
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0])) = int32(0)
	v626 = F_strtod(m, v438, v37+int32(8))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L27
	} else {
		goto L155
	}
L155:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628))))
	if v629 != 0 {
		goto L130
	} else {
		goto L156
	}
L156:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[0]))
	if v631 != 0 {
		goto L130
	} else {
		goto L157
	}
L157:
	;
	v633 = v626
	goto L131
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(_a_F_DecodeTimeOnly_9)
	v685 = v505
	goto L114
L159:
	;
	if v671 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L161:
	;
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v671
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	if v678 != int32(_a_F_DecodeTimeOnly_0) {
		v2007 = v37
		v2008 = int32(-1)
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v685 = v96
	goto L114
L164:
	;
	if v700 != 0 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	v700 = v694
	goto L167
L166:
	;
	v700 = int32(0)
	goto L167
L167:
	;
	goto L164
L168:
	;
	if v83|v66 != 0 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L170
L170:
	;
	v733 = v81 | int32(14)
	if int32(5) <= v692 {
		goto L183
	} else {
		goto L184
	}
L171:
	;
	v714 = F_strlen(m, v700)
	mBase = m.M
	if base.Ui32(v692-v714) < base.Ui32(int32(3)) {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v702 != int32(2) {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v710 = F_DecodeDate(m, v691, v81, v37+int32(68), v37+int32(59), l4)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L27
	} else {
		goto L174
	}
L174:
	;
	if v710 == int32(0) {
		v1169 = int32(0)
		v1175 = v86
		v1176 = v87
		v1183 = v94
		v1185 = v96
		v1187 = v98
		v1188 = v99
		goto L12
	} else {
		goto L175
	}
L175:
	;
	v2007 = v37
	v2008 = v710
	goto L4
L176:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L177:
	;
	goto L178
L178:
	;
	v725 = F_DecodeNumberField(m, v692, v691, v81|int32(14), v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L27
	} else {
		goto L179
	}
L179:
	;
	if v725 < int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L181:
	;
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v725
	v1169 = int32(0)
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L183:
	;
	v740 = F_DecodeNumberField(m, v692, v691, v733, v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L27
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v747 = int32(0)
	v753 = F_DecodeNumber(m, v692, v691, v747, v733, v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L27
	} else {
		goto L190
	}
L186:
	;
	if v740 < int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L188:
	;
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v740
	v1169 = int32(0)
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L190:
	;
	if v753 == int32(0) {
		v1169 = v747
		v1175 = v86
		v1176 = v87
		v1183 = v94
		v1185 = v96
		v1187 = v98
		v1188 = v99
		goto L12
	} else {
		goto L191
	}
L191:
	;
	v2007 = v37
	v2008 = v753
	goto L4
L192:
	;
	if v765 != 0 {
		v2007 = v37
		v2008 = v765
		goto L4
	} else {
		goto L193
	}
L193:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v37)+64))
	if v767 == int32(31) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_c_F_DecodeTimeOnly[3])))
	if v771 != 0 {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	v1013 = v767
	goto L196
L196:
	;
	if v1013 == int32(8) {
		v1207 = v80
		v1208 = v81
		v1213 = v86
		v1214 = v87
		v1221 = v94
		v1223 = v96
		v1225 = v98
		v1226 = v99
		goto L11
	} else {
		goto L242
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+64)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = v991
	v1013 = v1000
	goto L196
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_c_F_DecodeTimeOnly[3]))) = v937
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v937)+12))
	v965 = int32(*(*int8)(unsafe.Add(mBase, uint32(v937)+11)))
	v991 = v964
	v1000 = v965
	goto L197
L199:
	;
	goto L204
L200:
	;
	goto L201
L201:
	;
	v820 = int32(*(*int8)(unsafe.Add(mBase, uint32(v770))))
	v832 = int32(_a_F_DecodeTimeOnly_10)
	v845 = int32(_a_F_DecodeTimeOnly_11)
	goto L216
L202:
	;
	if v809-v810 == int32(0) {
		v937 = v771
		goto L198
	} else {
		goto L215
	}
L204:
	;
	goto L205
L205:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
	if v778 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v779 = v770
	v780 = v771
	v781 = int32(10)
	v782 = v778
	goto L210
L207:
	;
	v805 = v771
	v809 = int32(0)
	goto L208
L208:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v805))))
	goto L202
L209:
	;
	v805 = v800
	v809 = v802
	goto L208
L210:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780))))
	if base.B2i32(v782 != v784)|base.B2i32(v784 == int32(0)) != 0 {
		v800 = v780
		v802 = v782
		goto L209
	} else {
		goto L212
	}
L211:
	;
	v800 = v794
	v802 = int32(0)
	goto L209
L212:
	;
	v790 = v781 - int32(1)
	if v790 == int32(0) {
		v800 = v780
		v802 = v782
		goto L209
	} else {
		goto L213
	}
L213:
	;
	v793 = int32(1)
	v794 = v780 + v793
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779)+1)))
	if v795 != 0 {
		v779 = v779 + v793
		v780 = v794
		v781 = v790
		v782 = v795
		goto L210
	} else {
		goto L214
	}
L214:
	;
	goto L211
L215:
	;
	goto L201
L216:
	;
	v862 = v832 + (v845-v832)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v863 = int32(*(*int8)(unsafe.Add(mBase, uint32(v862))))
	v864 = v820 - v863
	if v864 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v991 = v916
	v1000 = int32(31)
	goto L197
L218:
	;
	goto L223
L219:
	;
	v915 = v864
	goto L220
L220:
	;
	v916 = int32(0)
	v920 = base.B2i32(v915 < v916)
	if v915 < v916 {
		goto L235
	} else {
		goto L236
	}
L221:
	;
	if v906 == int32(0) {
		v937 = v862
		goto L198
	} else {
		goto L234
	}
L223:
	;
	goto L224
L224:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
	if v873 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v874 = v770
	v875 = v862
	v876 = int32(10)
	v877 = v873
	goto L229
L226:
	;
	v900 = v862
	v904 = int32(0)
	goto L227
L227:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v900))))
	v906 = v904 - v905
	goto L221
L228:
	;
	v900 = v895
	v904 = v897
	goto L227
L229:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875))))
	if base.B2i32(v877 != v879)|base.B2i32(v879 == int32(0)) != 0 {
		v895 = v875
		v897 = v877
		goto L228
	} else {
		goto L231
	}
L230:
	;
	v895 = v889
	v897 = int32(0)
	goto L228
L231:
	;
	v885 = v876 - int32(1)
	if v885 == int32(0) {
		v895 = v875
		v897 = v877
		goto L228
	} else {
		goto L232
	}
L232:
	;
	v888 = int32(1)
	v889 = v875 + v888
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874)+1)))
	if v890 != 0 {
		v874 = v874 + v888
		v875 = v889
		v876 = v885
		v877 = v890
		goto L229
	} else {
		goto L233
	}
L233:
	;
	goto L230
L234:
	;
	v915 = v906
	goto L220
L235:
	;
	v921 = v862 - int32(16)
	goto L237
L236:
	;
	v921 = v845
	goto L237
L237:
	;
	if v915 < v916 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v924 = v832
	goto L240
L239:
	;
	v924 = v862 + int32(16)
	goto L240
L240:
	;
	if base.Ui32(v924) <= base.Ui32(v921) {
		v832 = v924
		v845 = v921
		goto L216
	} else {
		goto L241
	}
L241:
	;
	goto L217
L242:
	;
	v1040 = int32(1) << (uint(v1013) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1040
	v1042 = int32(-1)
	switch v1013 {
	case 0:
		goto L252
	default:
		v2007 = v37
		v2008 = v1042
		goto L4
	case 5:
		goto L249
	case 6:
		goto L250
	case 7:
		goto L248
	case 9:
		goto L247
	case 17:
		goto L245
	case 18:
		goto L246
	case 23:
		goto L244
	case 28:
		goto L251
	case 31:
		goto L243
	}
L243:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	v1117 = F_pg_tzset(m, v1116)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L27
	} else {
		goto L262
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(0)
	if v80 != 0 {
		v2007 = v37
		v2008 = v1042
		goto L4
	} else {
		goto L261
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(0)
	if v80 != 0 {
		v2007 = v37
		v2008 = v1042
		goto L4
	} else {
		goto L260
	}
L246:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1169 = v80
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = base.B2i32(v1107 == int32(1))
	goto L12
L247:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1169 = v80
	v1175 = v1106
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1040 | int32(32)
	if l6 == int32(0) {
		v2007 = v37
		v2008 = v1042
		goto L4
	} else {
		goto L259
	}
L249:
	;
	v1087 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1087
	if l6 == v1087 {
		v2007 = v37
		v2008 = v1042
		goto L4
	} else {
		goto L258
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1040 | int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v2007 = v37
		v2008 = v1042
		goto L4
	} else {
		goto L257
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1040 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v2007 = v37
		v2008 = v1042
		goto L4
	} else {
		goto L256
	}
L252:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	switch v1043 - int32(12) {
	case 0:
		goto L254
	default:
		v2007 = v37
		v2008 = v1042
		goto L4
	case 4:
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(_a_F_DecodeTimeOnly_12)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(3)
	v1057 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v1057
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1057
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
	v1169 = v80
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(_a_F_DecodeTimeOnly_0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(3)
	F_GetCurrentTimeUsec(m, l4, l5, int32(0))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L27
	} else {
		goto L255
	}
L255:
	;
	v1169 = v80
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L256:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1070 - v1071
	v1169 = v80
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L257:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1082
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1169 = v80
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L258:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1092
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1169 = v80
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L259:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v37)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1169 = v80
	v1175 = v86
	v1176 = v87
	v1183 = v1103
	v1185 = v96
	v1187 = v1102
	v1188 = v99
	goto L12
L260:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1169 = v1112
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L261:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1169 = v1115
	v1175 = v86
	v1176 = v87
	v1183 = v94
	v1185 = v96
	v1187 = v98
	v1188 = v99
	goto L12
L262:
	;
	if v1117 != 0 {
		v1140 = v1117
		goto L13
	} else {
		goto L263
	}
L263:
	;
	v2007 = v37
	v2008 = v1042
	goto L4
L264:
	;
	v2007 = v37
	v2008 = int32(-1)
	goto L4
L265:
	;
	goto L266
L266:
	;
	v1207 = v1169
	v1208 = v1191 | v81
	v1213 = v1175
	v1214 = v1176
	v1221 = v1183
	v1223 = v1185
	v1225 = v1187
	v1226 = v1188
	goto L11
L267:
	;
	goto L10
L268:
	;
	v1241 = v37
	v1246 = v1208
	v1251 = v1213
	v1252 = v1214
	v1259 = v1221
	v1261 = v1223
	v1263 = v1225
	v1264 = v1226
	goto L5
L269:
	;
	if v1435 != 0 {
		v2007 = v1241
		v2008 = v1435
		goto L4
	} else {
		goto L306
	}
L270:
	;
	if v1246&int32(_a_F_DecodeTimeOnly_13) != 0 {
		goto L287
	} else {
		goto L288
	}
L271:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1264 != 0 {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1293
	goto L270
L273:
	;
	v1293 = int32(1) - v1273
	goto L272
L274:
	;
	if int32(0) < v1273 {
		goto L273
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	if v1267 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1435 = int32(-2)
	goto L269
L278:
	;
	if v1273 < int32(0) {
		goto L281
	} else {
		goto L282
	}
L279:
	;
	goto L280
L280:
	;
	if int32(0) < v1273 {
		goto L270
	} else {
		goto L286
	}
L281:
	;
	v1435 = int32(-2)
	goto L269
L282:
	;
	goto L283
L283:
	;
	if base.Ui32(v1273) <= base.Ui32(int32(69)) {
		v1293 = v1273 + int32(2000)
		goto L272
	} else {
		goto L284
	}
L284:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1273) {
		goto L270
	} else {
		goto L285
	}
L285:
	;
	v1293 = v1273 + int32(1900)
	goto L272
L286:
	;
	v1435 = int32(-2)
	goto L269
L287:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v1304 = v1299 + int32(_a_F_DecodeTimeOnly_14)
	v1306 = base.I32_div_s(v1304, int32(4))
	v1309 = base.I32_div_s(v1304, int32(-100))
	v1312 = base.I32_div_s(v1304, int32(400))
	v1313 = v1298 + v1299*int32(365) + v1306 + v1309 + v1312
	v1315 = v1313 + int32(_a_F_DecodeTimeOnly_15)
	v1316 = int32(_a_F_DecodeTimeOnly_4)
	v1317 = base.I32_div_u_s(v1315, v1316)
	v1318 = int32(3)
	v1324 = int32(2)
	v1329 = base.I32_div_u_s((v1317*int32(1073595727)+v1315)<<(uint(v1324)%32)|v1318, v1316)
	v1332 = v1313 + v1317*v1318 + v1329 + int32(_a_F_DecodeTimeOnly_16)
	v1333 = int32(1461)
	v1334 = base.I32_div_u_s(v1332, v1333)
	v1337 = v1334*int32(-1461) + v1332
	v1339 = v1337 << (uint(v1324) % 32)
	if base.Ui32(v1333) <= base.Ui32(v1339) {
		goto L291
	} else {
		goto L292
	}
L288:
	;
	goto L289
L289:
	;
	if v1246&int32(2) == int32(0) {
		goto L294
	} else {
		goto L295
	}
L290:
	;
	v1352 = base.I32_div_u_s(v1339, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1352 + v1334<<(uint(int32(2))%32) - int32(_a_F_DecodeTimeOnly_6)
	v1360 = v1350 + int32(123)
	v1364 = int32(base.Ui32(v1360*int32(2141)) >> (uint(int32(16)) % 32))
	v1368 = base.I32_rem_u_s(v1364+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v1368 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v1360 - int32(base.Ui32(v1364*int32(_a_F_DecodeTimeOnly_7))>>(uint(int32(8))%32))
	goto L289
L291:
	;
	v1345 = base.I32_rem_u_s(v1337+int32(305), int32(365))
	v1350 = v1345
	goto L290
L292:
	;
	goto L293
L293:
	;
	v1349 = base.I32_rem_u_s(v1337+int32(306), int32(366))
	v1350 = v1349
	goto L290
L294:
	;
	if v1246&int32(8) == int32(0) {
		goto L297
	} else {
		goto L298
	}
L295:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if base.Ui32(int32(-12)) <= base.Ui32(v1385-int32(13)) {
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1435 = int32(-3)
	goto L269
L297:
	;
	v1401 = int32(14)
	if v1246&v1401 != v1401 {
		goto L300
	} else {
		goto L301
	}
L298:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if base.Ui32(int32(-31)) <= base.Ui32(v1395-int32(32)) {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1435 = int32(-3)
	goto L269
L300:
	;
	v1435 = int32(0)
	goto L269
L301:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1407&int32(3) != 0 {
		v1417 = int32(0)
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1417*int32(52)+v1420<<(uint(int32(2))%32))+uint32(_c_F_DecodeTimeOnly[4])))
	if v1405 <= v1426 {
		goto L300
	} else {
		goto L305
	}
L303:
	;
	v1412 = base.I32_rem_s(v1407, int32(100))
	if v1412 != 0 {
		v1417 = int32(1)
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1414 = base.I32_rem_s(v1407, int32(400))
	v1417 = base.B2i32(v1414 == int32(0))
	goto L302
L305:
	;
	v1435 = int32(-2)
	goto L269
L306:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v1251 == int32(2) {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1464 = int32(60)
	goto L319
L308:
	;
	v1453 = v1436
	goto L307
L309:
	;
	goto L310
L310:
	;
	if int32(12) < v1436 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v2007 = v1241
	v2008 = int32(-2)
	goto L4
L312:
	;
	goto L313
L313:
	;
	switch v1251 {
	case 0:
		goto L316
	case 1:
		goto L315
	default:
		v1453 = v1436
		goto L307
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v1451
	v1453 = v1451
	goto L307
L315:
	;
	v1445 = int32(12)
	if v1436 == v1445 {
		v1453 = v1445
		goto L307
	} else {
		goto L318
	}
L316:
	;
	if v1436 == int32(12) {
		v1451 = int32(0)
		goto L314
	} else {
		goto L317
	}
L317:
	;
	v1453 = v1436
	goto L307
L318:
	;
	v1451 = v1436 + int32(12)
	goto L314
L319:
	;
	if base.B2i32(base.Ui32(int32(24)) < base.Ui32(v1453))|base.B2i32(base.Ui32(int32(59)) < base.Ui32(v1454))|(base.B2i32(base.Ui32(v1464) < base.Ui32(v1455))|base.B2i32(base.Ui32(int32(_a_F_DecodeTimeOnly_17)) < base.Ui32(v1456)))|base.B2i32(base.Ui64(int64(86400000000)) < base.Ui64(base.I64_extend_i32_u(v1456)+base.I64_extend_i32_u((v1453*v1464+v1454)*v1464+v1455)*int64(1000000))) != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v2007 = v1241
	v2008 = int32(-2)
	goto L4
L321:
	;
	goto L322
L322:
	;
	v1485 = int32(-1)
	v1486 = int32(_a_F_DecodeTimeOnly_0)
	if v1246&v1486 != v1486 {
		v2007 = v1241
		v2008 = v1485
		goto L4
	} else {
		goto L323
	}
L323:
	;
	if v1252 != 0 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v2007 = v1241
	v2008 = v1998
	goto L4
L325:
	;
	if v1246&int32(268435456) != 0 {
		v1998 = v1485
		goto L324
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	if v1259 != 0 {
		goto L377
	} else {
		goto L378
	}
L328:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+uint32(_c_F_DecodeTimeOnly[5])))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1252)+264))
	if v1498 < int32(2) {
		goto L331
	} else {
		goto L332
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1703
	goto L327
L330:
	;
	if v1530 != 0 {
		goto L339
	} else {
		goto L340
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1241+int32(72)))) = v1497
	v1530 = int32(1)
	goto L330
L332:
	;
	v1504 = int32(1)
	goto L333
L333:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1252+int32(_a_F_DecodeTimeOnly_18)+v1504<<(uint(int32(4))%32))))
	if v1497 == v1512 {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	v1530 = int32(0)
	goto L330
L335:
	;
	v1515 = v1504 + int32(1)
	if v1498 != v1515 {
		v1504 = v1515
		goto L333
	} else {
		goto L338
	}
L336:
	;
	goto L337
L337:
	;
	goto L334
L338:
	;
	goto L331
L339:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+72))
	v1703 = int32(0) - v1532
	goto L329
L340:
	;
	goto L341
L341:
	;
	v1534 = int32(14)
	if v1246&v1534 != v1534 {
		v1998 = v1485
		goto L324
	} else {
		goto L342
	}
L342:
	;
	v1539 = v1241 + int32(8)
	v1547 = m.G0
	v1549 = v1547 - int32(32)
	m.G0 = v1549
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1551 <= int32(-4713) {
		goto L347
	} else {
		goto L348
	}
L343:
	;
	v1703 = v1699
	goto L329
L344:
	;
	m.G0 = v1549 + int32(32)
	goto L343
L345:
	;
	v1686 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1686
	*(*int64)(unsafe.Add(mBase, uint32(v1539))) = int64(0)
	v1699 = v1686
	goto L344
L346:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v1571 = int32(60)
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1582 = base.B2i32(int32(2) < v1567)
	if int32(2) < v1567 {
		goto L357
	} else {
		goto L358
	}
L347:
	;
	if v1551 != int32(-4713) {
		goto L345
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	if v1551 <= int32(_a_F_DecodeTimeOnly_19) {
		goto L352
	} else {
		goto L353
	}
L350:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if int32(10) < v1556 {
		v1567 = v1556
		goto L346
	} else {
		goto L351
	}
L351:
	;
	goto L345
L352:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v1567 = v1561
	goto L346
L353:
	;
	goto L354
L354:
	;
	if v1551 != int32(_a_F_DecodeTimeOnly_20) {
		goto L345
	} else {
		goto L355
	}
L355:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if int32(5) < v1564 {
		goto L345
	} else {
		goto L356
	}
L356:
	;
	v1567 = v1564
	goto L346
L357:
	;
	v1583 = int32(_a_F_DecodeTimeOnly_6)
	goto L359
L358:
	;
	v1583 = int32(_a_F_DecodeTimeOnly_14)
	goto L359
L359:
	;
	v1584 = v1583 + v1551
	v1592 = base.I32_div_u_s(v1584, int32(100))
	v1595 = base.I32_div_u_s(v1584, int32(400))
	if int32(2) < v1567 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1599 = int32(1)
	goto L362
L361:
	;
	v1599 = int32(13)
	goto L362
L362:
	;
	v1604 = base.I32_div_s((v1599+v1567)*int32(_a_F_DecodeTimeOnly_7), int32(256))
	v1607 = v1578 + v1584*int32(365) + int32(base.Ui32(v1584)>>(uint(int32(2))%32)) - v1592 + v1595 + v1604 - int32(_a_F_DecodeTimeOnly_21)
	v1611 = base.I64_extend_i32_s(v1568+(v1569+v1570*v1571)*v1571) + base.I64_extend_i32_s(v1607)*int64(86400)
	if base.B2i32(int32(0) < v1607)&base.B2i32(v1611 < int64(0)) != 0 {
		goto L345
	} else {
		goto L363
	}
L363:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1549)+24)) = v1611 - int64(86400)
	v1630 = F_pg_next_dst_boundary(m, v1549+int32(24), v1549+int32(12), v1549+int32(4), v1549+int32(16), v1549+int32(8), v1549, v1252)
	mBase = m.M
	if v1630 < int32(0) {
		goto L345
	} else {
		goto L364
	}
L364:
	;
	if v1630 == int32(0) {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1635
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v1539))) = v1611 - base.I64_extend_i32_s(v1637)
	v1699 = int32(0) - v1637
	goto L344
L366:
	;
	goto L367
L367:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+12))
	v1645 = v1611 - base.I64_extend_i32_s(v1643)
	v1646 = *(*int64)(unsafe.Add(mBase, uint32(v1549)+16))
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+8))
	v1650 = v1611 - base.I64_extend_i32_s(v1648)
	if base.B2i32(v1646 <= v1645)|base.B2i32(v1646 <= v1650) == int32(0) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1655
	*(*int64)(unsafe.Add(mBase, uint32(v1539))) = v1645
	v1699 = int32(0) - v1643
	goto L344
L369:
	;
	goto L370
L370:
	;
	if base.B2i32(v1650 < v1646)|base.B2i32(v1645 <= v1646) == int32(0) {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1549)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1665
	*(*int64)(unsafe.Add(mBase, uint32(v1539))) = v1650
	v1699 = int32(0) - v1648
	goto L344
L372:
	;
	goto L373
L373:
	;
	if v1643 < v1648 {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1671
	*(*int64)(unsafe.Add(mBase, uint32(v1539))) = v1645
	v1699 = int32(0) - v1643
	goto L344
L375:
	;
	goto L376
L376:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1549)))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1676
	*(*int64)(unsafe.Add(mBase, uint32(v1539))) = v1650
	v1699 = int32(0) - v1648
	goto L344
L377:
	;
	if v1246&int32(268435456) != 0 {
		v2007 = v1241
		v2008 = v1485
		goto L4
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v1790 = int32(0)
	if base.B2i32(l6 == v1790)|v1246&int32(32) != 0 {
		v2007 = v1241
		v2008 = v1790
		goto L4
	} else {
		goto L397
	}
L380:
	;
	v1708 = v1246 & int32(14)
	if v1708 != 0 {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+16)) = v1724
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+12)) = v1726
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+8)) = v1728
	v1731 = v1241 + int32(8)
	v1736 = m.G0
	v1738 = v1736 - int32(288)
	m.G0 = v1738
	v1742 = F_DetermineTimeZoneOffsetInternal(m, v1731, v1259, v1738+int32(280))
	mBase = m.M
	v1744 = v1738 + int32(16)
	v1746 = F_strlcpy(m, v1744, v1263, int32(256))
	mBase = m.M
	v1747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1738)+16)))
	if v1747 != 0 {
		goto L388
	} else {
		goto L389
	}
L382:
	;
	if v1708 != int32(14) {
		v2007 = v1241
		v2008 = v1485
		goto L4
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	F_GetCurrentTimeUsec(m, v1241+int32(8), v1241+int32(72), int32(0))
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L27
	} else {
		goto L386
	}
L385:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+28)) = v1711
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+24)) = v1713
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+20)) = v1715
	goto L381
L386:
	;
	goto L381
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1782
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1787
	goto L379
L388:
	;
	v1749 = v1744
	v1753 = v1747
	goto L391
L389:
	;
	goto L390
L390:
	;
	v1775 = F_pg_interpret_timezone_abbrev(m, v1738+int32(16), v1738+int32(280), v1738+int32(12), v1738+int32(8), v1259)
	mBase = m.M
	if v1775 != 0 {
		goto L394
	} else {
		goto L395
	}
L391:
	;
	v1755 = F_pg_toupper(m, v1753)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1749))) = uint8(v1755)
	v1757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1749)+1)))
	if v1757 != 0 {
		v1749 = v1749 + int32(1)
		v1753 = v1757
		goto L391
	} else {
		goto L393
	}
L392:
	;
	goto L390
L393:
	;
	goto L392
L394:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1738)+12))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1738)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1731)+32)) = v1777
	v1782 = int32(0) - v1776
	goto L396
L395:
	;
	v1782 = v1742
	goto L396
L396:
	;
	m.G0 = v1738 + int32(288)
	goto L387
L397:
	;
	if v1246&int32(268435456) != 0 {
		goto L400
	} else {
		goto L401
	}
L398:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+16)) = v1819
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+12)) = v1821
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+8)) = v1823
	v1826 = v1241 + int32(8)
	v1828 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeTimeOnly[6]))
	v1830 = v1241 + int32(72)
	v1838 = m.G0
	v1840 = v1838 - int32(32)
	m.G0 = v1840
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1826)+20))
	if v1842 <= int32(-4713) {
		goto L409
	} else {
		goto L410
	}
L399:
	;
	F_GetCurrentTimeUsec(m, v1241+int32(8), v1241+int32(72), int32(0))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L27
	} else {
		goto L404
	}
L400:
	;
	v2007 = v1241
	v2008 = int32(-1)
	goto L4
L401:
	;
	v1799 = v1246 & int32(14)
	if v1799 == int32(0) {
		goto L399
	} else {
		goto L402
	}
L402:
	;
	if v1799 != int32(14) {
		goto L400
	} else {
		goto L403
	}
L403:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+28)) = v1804
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+24)) = v1806
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+20)) = v1808
	goto L398
L404:
	;
	goto L398
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1990
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1995
	v1998 = v1790
	goto L324
L406:
	;
	m.G0 = v1840 + int32(32)
	goto L405
L407:
	;
	v1977 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1826)+32)) = v1977
	*(*int64)(unsafe.Add(mBase, uint32(v1830))) = int64(0)
	v1990 = v1977
	goto L406
L408:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1826)))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1826)+4))
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v1826)+8))
	v1862 = int32(60)
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v1826)+12))
	v1873 = base.B2i32(int32(2) < v1858)
	if int32(2) < v1858 {
		goto L419
	} else {
		goto L420
	}
L409:
	;
	if v1842 != int32(-4713) {
		goto L407
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	if v1842 <= int32(_a_F_DecodeTimeOnly_19) {
		goto L414
	} else {
		goto L415
	}
L412:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1826)+16))
	if int32(10) < v1847 {
		v1858 = v1847
		goto L408
	} else {
		goto L413
	}
L413:
	;
	goto L407
L414:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1826)+16))
	v1858 = v1852
	goto L408
L415:
	;
	goto L416
L416:
	;
	if v1842 != int32(_a_F_DecodeTimeOnly_20) {
		goto L407
	} else {
		goto L417
	}
L417:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1826)+16))
	if int32(5) < v1855 {
		goto L407
	} else {
		goto L418
	}
L418:
	;
	v1858 = v1855
	goto L408
L419:
	;
	v1874 = int32(_a_F_DecodeTimeOnly_6)
	goto L421
L420:
	;
	v1874 = int32(_a_F_DecodeTimeOnly_14)
	goto L421
L421:
	;
	v1875 = v1874 + v1842
	v1883 = base.I32_div_u_s(v1875, int32(100))
	v1886 = base.I32_div_u_s(v1875, int32(400))
	if int32(2) < v1858 {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1890 = int32(1)
	goto L424
L423:
	;
	v1890 = int32(13)
	goto L424
L424:
	;
	v1895 = base.I32_div_s((v1890+v1858)*int32(_a_F_DecodeTimeOnly_7), int32(256))
	v1898 = v1869 + v1875*int32(365) + int32(base.Ui32(v1875)>>(uint(int32(2))%32)) - v1883 + v1886 + v1895 - int32(_a_F_DecodeTimeOnly_21)
	v1902 = base.I64_extend_i32_s(v1859+(v1860+v1861*v1862)*v1862) + base.I64_extend_i32_s(v1898)*int64(86400)
	if base.B2i32(int32(0) < v1898)&base.B2i32(v1902 < int64(0)) != 0 {
		goto L407
	} else {
		goto L425
	}
L425:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1840)+24)) = v1902 - int64(86400)
	v1921 = F_pg_next_dst_boundary(m, v1840+int32(24), v1840+int32(12), v1840+int32(4), v1840+int32(16), v1840+int32(8), v1840, v1828)
	mBase = m.M
	if v1921 < int32(0) {
		goto L407
	} else {
		goto L426
	}
L426:
	;
	if v1921 == int32(0) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1826)+32)) = v1926
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v1830))) = v1902 - base.I64_extend_i32_s(v1928)
	v1990 = int32(0) - v1928
	goto L406
L428:
	;
	goto L429
L429:
	;
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+12))
	v1936 = v1902 - base.I64_extend_i32_s(v1934)
	v1937 = *(*int64)(unsafe.Add(mBase, uint32(v1840)+16))
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+8))
	v1941 = v1902 - base.I64_extend_i32_s(v1939)
	if base.B2i32(v1937 <= v1936)|base.B2i32(v1937 <= v1941) == int32(0) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1826)+32)) = v1946
	*(*int64)(unsafe.Add(mBase, uint32(v1830))) = v1936
	v1990 = int32(0) - v1934
	goto L406
L431:
	;
	goto L432
L432:
	;
	if base.B2i32(v1941 < v1937)|base.B2i32(v1936 <= v1937) == int32(0) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1840)))
	*(*int32)(unsafe.Add(mBase, uint32(v1826)+32)) = v1956
	*(*int64)(unsafe.Add(mBase, uint32(v1830))) = v1941
	v1990 = int32(0) - v1939
	goto L406
L434:
	;
	goto L435
L435:
	;
	if v1934 < v1939 {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1826)+32)) = v1962
	*(*int64)(unsafe.Add(mBase, uint32(v1830))) = v1936
	v1990 = int32(0) - v1934
	goto L406
L437:
	;
	goto L438
L438:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1840)))
	*(*int32)(unsafe.Add(mBase, uint32(v1826)+32)) = v1967
	*(*int64)(unsafe.Add(mBase, uint32(v1830))) = v1941
	v1990 = int32(0) - v1939
	goto L406
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
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int64
	_ = v281
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	v2 = int32(0)
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(2288)
	m.G0 = v12
	if l0 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L6
	} else {
		goto L85
	}
L2:
	;
	m.G0 = v12 + int32(2288)
	return v330
L3:
	;
	v17 = F_palloc(m, int32(24))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_readTimeLineHistory[0])))
	if v35 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	return int32(0)
L7:
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
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v330 = v32
	goto L2
L9:
	;
	v68 = F_AllocateFile(m, v12+int32(1264), int32(_a_F_readTimeLineHistory_0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L19
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = l0
	v40 = v12 + int32(1200)
	v45 = F_pg_snprintf(m, v40, int32(64), int32(_a_F_readTimeLineHistory_1), v12+int32(128))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = l0
	v61 = F_pg_snprintf(m, v12+int32(1264), int32(1024), int32(_a_F_readTimeLineHistory_2), v12+int32(144))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v52 = F_RestoreArchivedFile(m, v12+int32(1264), v40, int32(_a_F_readTimeLineHistory_3), int64(0), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v64 = v52
	goto L9
L15:
	;
	v64 = v2
	goto L9
L16:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	goto L68
L17:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L6
	} else {
		goto L64
	}
L18:
	;
	v105 = v2
	v108 = v2
	v110 = v9
	goto L30
L19:
	;
	if v68 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_readTimeLineHistory[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(167772219)
	v77 = F_fgets(m, v12+int32(160), int32(1024), v68)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_readTimeLineHistory[2]))
	if v84 != int32(44) {
		goto L17
	} else {
		goto L25
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_readTimeLineHistory[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(0)
	if v77 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v276 = v2
	v279 = v2
	v281 = v9
	goto L16
L25:
	;
	v88 = F_palloc(m, int32(24))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v90 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v12)+1192)) = v88
	v100 = F_list_make1_impl(m, int32(1), v12+int32(12))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v330 = v100
	goto L2
L28:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L6
	} else {
		goto L60
	}
L29:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L56
	}
L30:
	;
	v115 = v12 + int32(160)
	goto L32
L31:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L52
	}
L32:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if base.Ui32(v122-int32(9)) < base.Ui32(int32(5)) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L31
L34:
	;
	goto L33
L35:
	;
	v115 = v115 + int32(1)
	goto L32
L36:
	;
	switch v122 - int32(32) {
	case 0:
		goto L35
	case 1, 2:
		goto L38
	case 3:
		v170 = v105
		v172 = v108
		v173 = v110
		goto L37
	default:
		goto L39
	}
L37:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_readTimeLineHistory[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = int32(167772219)
	v181 = F_fgets(m, v12+int32(160), int32(1024), v68)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L6
	} else {
		goto L50
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v12 + int32(148)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+116)) = v12 + int32(152)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v12 + int32(156)
	v145 = F_sscanf(m, v12+int32(160), int32(_a_F_readTimeLineHistory_4), v12+int32(112))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	if v122 == int32(0) {
		v170 = v105
		v172 = v108
		v173 = v110
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	if v145 <= int32(0) {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	if v145 != int32(3) {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	if base.Ui32(v152) <= base.Ui32(v105) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v154 = v108
	goto L46
L45:
	;
	v154 = int32(0)
	goto L46
L46:
	;
	if v154 != 0 {
		goto L28
	} else {
		goto L47
	}
L47:
	;
	v156 = F_palloc(m, int32(24))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	*(*int64)(unsafe.Add(mBase, uint32(v156)+8)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v158
	v161 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+148)))
	v162 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+152)))
	v165 = v161 | v162<<(uint(int64(32))%64)
	*(*int64)(unsafe.Add(mBase, uint32(v156)+16)) = v165
	v167 = F_lcons(m, v156, v108)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	v170 = v152
	v172 = v167
	v173 = v165
	goto L37
L50:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_readTimeLineHistory[1]))
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v185
	if v181 == v185 {
		v276 = v170
		v279 = v172
		v281 = v173
		goto L16
	} else {
		goto L51
	}
L51:
	;
	v105 = v170
	v108 = v172
	v110 = v173
	goto L30
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v12 + int32(160)
	F_errmsg(m, int32(_a_F_readTimeLineHistory_5), v12-int32(-64))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	F_errhint(m, int32(_a_F_readTimeLineHistory_6), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_readTimeLineHistory_7), int32(164), int32(_a_F_readTimeLineHistory_8))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v12 + int32(160)
	F_errmsg(m, int32(_a_F_readTimeLineHistory_5), v12+int32(96))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	F_errhint(m, int32(_a_F_readTimeLineHistory_9), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_readTimeLineHistory_7), int32(169), int32(_a_F_readTimeLineHistory_8))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v12 + int32(160)
	F_errmsg(m, int32(_a_F_readTimeLineHistory_10), v12+int32(80))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errhint(m, int32(_a_F_readTimeLineHistory_11), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_readTimeLineHistory_7), int32(174), int32(_a_F_readTimeLineHistory_8))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(1264)
	F_errmsg(m, int32(_a_F_readTimeLineHistory_12), v12+int32(16))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_readTimeLineHistory_7), int32(111), int32(_a_F_readTimeLineHistory_8))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	if int32(base.Ui32(v282)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v306 = F_FreeFile(m, v68)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L76
	}
L72:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(1264)
	F_errmsg(m, int32(_a_F_readTimeLineHistory_13), v12+int32(48))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_readTimeLineHistory_7), int32(143), int32(_a_F_readTimeLineHistory_8))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	if base.Ui32(l0) <= base.Ui32(v276) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v310 = v279
	goto L79
L78:
	;
	v310 = int32(0)
	goto L79
L79:
	;
	if v310 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v312 = F_palloc(m, int32(24))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v312)+16)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v312)+8)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = l0
	v318 = F_lcons(m, v312, v279)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	if v64 == int32(0) {
		v330 = v318
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_KeepFileRestoredFromArchive(m, v12+int32(1264), v12+int32(1200))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	v330 = v318
	goto L2
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(1264)
	F_errmsg(m, int32(_a_F_readTimeLineHistory_14), v12+int32(32))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	F_errhint(m, int32(_a_F_readTimeLineHistory_15), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_readTimeLineHistory_7), int32(195), int32(_a_F_readTimeLineHistory_8))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_time(m *base.Module) int64 {
	var v1 float64
	_ = v1
	v1 = m.Env.Emscripten_date_now(m)
	return base.I64_trunc_sat_f64_s(base.F64_div(v1, float64(1000)))
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
	var v14 int32
	_ = v14
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v14 = v9 + int32(4)
	F_GetCurrentDateTime(m, v14)
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
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_time_timetz[0]))
		v39 = m.G0
		v40 = int32(16)
		v41 = v39 - v40
		m.G0 = v41
		v45 = F_DetermineTimeZoneOffsetInternal(m, v14, v37, v41+int32(8))
		mBase = m.M
		m.G0 = v41 + v40
		v50 = F_palloc(m, int32(16))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v45
			*(*int64)(unsafe.Add(mBase, uint32(v50))) = v12
			m.G0 = v9 + int32(48)
			return v50
		}
	}
}
