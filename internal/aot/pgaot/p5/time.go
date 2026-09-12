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
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int64
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int64
	_ = v586
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v719 int32
	_ = v719
	var v726 float64
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 float64
	_ = v732
	var v734 float64
	_ = v734
	var v738 int64
	_ = v738
	var v740 int64
	_ = v740
	var v743 int64
	_ = v743
	var v748 int64
	_ = v748
	var v750 int64
	_ = v750
	var v755 int64
	_ = v755
	var v757 int64
	_ = v757
	var v761 int64
	_ = v761
	var v767 int32
	_ = v767
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v984 int32
	_ = v984
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1003 int32
	_ = v1003
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1091 int32
	_ = v1091
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1207 int32
	_ = v1207
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1250 int32
	_ = v1250
	var v1268 int32
	_ = v1268
	var v1283 int32
	_ = v1283
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1360 int32
	_ = v1360
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1408 int32
	_ = v1408
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1540 int32
	_ = v1540
	var v1560 int32
	_ = v1560
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1591 int32
	_ = v1591
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1652 int32
	_ = v1652
	var v1662 int32
	_ = v1662
	var v1668 int32
	_ = v1668
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1693 int32
	_ = v1693
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1713 int32
	_ = v1713
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1737 int32
	_ = v1737
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1770 int32
	_ = v1770
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1817 int32
	_ = v1817
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1833 int32
	_ = v1833
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1899 int32
	_ = v1899
	var v1901 int32
	_ = v1901
	var v1903 int32
	_ = v1903
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
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
	m.G0 = v1938 + int32(80)
	return v1939
L5:
	;
	v1535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1509)+59)))
	if v1529 != 0 {
		goto L369
	} else {
		goto L370
	}
L6:
	;
	v1509 = v37
	v1512 = v9
	v1519 = int32(2)
	v1520 = v9
	v1527 = v9
	v1529 = v9
	v1531 = v9
	v1532 = v9
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
		v1938 = v37
		v1939 = v102
		goto L4
	}
L10:
	;
	if v1476 != 0 {
		v1938 = v37
		v1939 = int32(-1)
		goto L4
	} else {
		goto L367
	}
L11:
	;
	v1498 = v85 + int32(1)
	if v1498 != l2 {
		v79 = v1474
		v81 = v1476
		v85 = v1498
		v86 = v1481
		v87 = v1482
		v94 = v1489
		v96 = v1491
		v98 = v1493
		v99 = v1494
		goto L9
	} else {
		goto L366
	}
L12:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	if v1459&v79 != 0 {
		goto L363
	} else {
		goto L364
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(32)
	v1438 = v81
	v1443 = v86
	v1444 = v1408
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1408 = v332
	goto L13
L15:
	;
	v1027 = l0 + v104
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	v1035 = F_DecodeTimezoneAbbrev(m, v85, v1028, v37-int32(-64), v37+int32(60), v37+int32(52), l7)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L28
	} else {
		goto L289
	}
L16:
	;
	if v81 != 0 {
		goto L128
	} else {
		goto L129
	}
L17:
	;
	if l6 == int32(0) {
		v1938 = v37
		v1939 = v102
		goto L4
	} else {
		goto L100
	}
L18:
	;
	switch v81 {
	case 0, 3:
		goto L94
	default:
		v1938 = v37
		v1939 = v102
		goto L4
	}
L19:
	;
	if l6 == int32(0) {
		v1938 = v37
		v1939 = v102
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
		v1438 = v81
		v1443 = v86
		v1444 = v87
		v1451 = v94
		v1453 = v96
		v1455 = v98
		v1456 = v99
		goto L12
	} else {
		goto L30
	}
L30:
	;
	v1938 = v37
	v1939 = v120
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
	v332 = F_pg_tzset(m, v127)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L28
	} else {
		goto L92
	}
L34:
	;
	v1938 = v37
	v1939 = int32(-1)
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
	v1938 = v37
	v1939 = int32(-1)
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
		v1938 = v37
		v1939 = v249
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
	v165 = int32(4634788)
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v172 = F_strtoint(m, v147+int32(1), v159+int32(12))
	mBase = m.M
	v173 = int32(-5)
	v175 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
	v183 = int32(4634788)
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v190 = F_strtoint(m, v178+int32(1), v159+int32(12))
	mBase = m.M
	v192 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
	v199 = int32(4634788)
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v206 = F_strtoint(m, v195+int32(1), v159+int32(12))
	mBase = m.M
	v208 = *(*int32)(unsafe.Add(mBase, _consts[86]))
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
	if v258&int32(3) == v256 {
		v282 = v258
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v322 = F_DecodeNumberField(m, v315, v258, v79|int32(14), v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L28
	} else {
		goto L88
	}
L72:
	;
	v315 = v307 - v258
	goto L71
L73:
	;
	v286 = v282
	goto L82
L74:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	if v266 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v315 = int32(0)
	goto L71
L76:
	;
	goto L77
L77:
	;
	v271 = v258
	goto L78
L78:
	;
	v275 = v271 + int32(1)
	if v275&int32(3) == int32(0) {
		v282 = v275
		goto L73
	} else {
		goto L80
	}
L79:
	;
	v307 = v275
	goto L72
L80:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	if v280 != 0 {
		v271 = v275
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v295 = int32(-2139062144)
	if (int32(16843008)-v292|v292)&v295 == v295 {
		v286 = v286 + int32(4)
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v301 = v286
	goto L85
L84:
	;
	goto L83
L85:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v305 != 0 {
		v301 = v301 + int32(1)
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v307 = v301
	goto L72
L87:
	;
	goto L86
L88:
	;
	if v322 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v322
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v328 | int32(32)
	v1438 = v81
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L92:
	;
	if v332 != 0 {
		goto L14
	} else {
		goto L93
	}
L93:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v334
	v1938 = v37
	v1939 = int32(-6)
	goto L4
L94:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v344 = F_DecodeTimeCommon(m, v338, int32(32767), v37+int32(68), v37+int32(8))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L28
	} else {
		goto L95
	}
L95:
	;
	if v344 != 0 {
		v1938 = v37
		v1939 = v344
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v37)+24))
	if int64(2147483648) <= v346 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v1938 = v37
	v1939 = int32(-2)
	goto L4
L98:
	;
	goto L99
L99:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l4)+8)) = uint32(v346)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v351
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v353
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v355
	v1438 = int32(0)
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L100:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v364 = int32(0)
	v370 = m.G0
	v372 = v370 - int32(16)
	m.G0 = v372
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	switch v375 - int32(43) {
	case 0, 2:
		goto L103
	default:
		v462 = int32(-1)
		goto L102
	}
L101:
	;
	if v462 != 0 {
		v1938 = v37
		v1939 = v462
		goto L4
	} else {
		goto L127
	}
L102:
	;
	m.G0 = v372 + int32(16)
	goto L101
L103:
	;
	v378 = int32(4634788)
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v385 = F_strtoint(m, v361+int32(1), v372+int32(12))
	mBase = m.M
	v386 = int32(-5)
	v388 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v388 == int32(68) {
		v462 = v386
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	if v392 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if base.Ui32(int32(15)) < base.Ui32(v434) {
		v462 = v386
		goto L102
	} else {
		goto L118
	}
L106:
	;
	if v392 != int32(58) {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	v424 = F_strlen(m, v361)
	mBase = m.M
	if base.Ui32(v424) < base.Ui32(int32(4)) {
		goto L115
	} else {
		goto L116
	}
L109:
	;
	v433 = int32(0)
	v434 = v385
	v436 = v364
	goto L105
L110:
	;
	goto L111
L111:
	;
	v396 = int32(4634788)
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v403 = F_strtoint(m, v391+int32(1), v372+int32(12))
	mBase = m.M
	v405 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v405 == int32(68) {
		v462 = v386
		goto L102
	} else {
		goto L112
	}
L112:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408))))
	if v409 != int32(58) {
		v433 = v403
		v434 = v385
		v436 = v364
		goto L105
	} else {
		goto L113
	}
L113:
	;
	v412 = int32(4634788)
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v419 = F_strtoint(m, v408+int32(1), v372+int32(12))
	mBase = m.M
	v421 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v421 != int32(68) {
		v433 = v403
		v434 = v385
		v436 = v419
		goto L105
	} else {
		goto L114
	}
L114:
	;
	v462 = v386
	goto L102
L115:
	;
	v433 = int32(0)
	v434 = v385
	v436 = v364
	goto L105
L116:
	;
	goto L117
L117:
	;
	v428 = int32(100)
	v429 = base.I32_div_s(v385, v428)
	v433 = v385 - v429*v428
	v434 = v429
	v436 = v364
	goto L105
L118:
	;
	if base.Ui32(int32(59)) < base.Ui32(v433) {
		v462 = v386
		goto L102
	} else {
		goto L119
	}
L119:
	;
	if base.Ui32(int32(59)) < base.Ui32(v436) {
		v462 = v386
		goto L102
	} else {
		goto L120
	}
L120:
	;
	v443 = int32(60)
	v448 = (v434*v443+v433)*v443 + v436
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v451 == int32(45) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v454 = v448
	goto L123
L122:
	;
	v454 = int32(0) - v448
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+int32(8)))) = v454
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if v459 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v460 = int32(-1)
	goto L126
L125:
	;
	v460 = int32(0)
	goto L126
L126:
	;
	v462 = v460
	goto L102
L127:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v469
	v1408 = v87
	goto L13
L128:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v474 = l0 + v104
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	v479 = F_strtol(m, v475, v37+int32(72), int32(10))
	mBase = m.M
	goto L131
L129:
	;
	goto L130
L130:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	if v850&int32(3) == int32(0) {
		v874 = v850
		goto L228
	} else {
		goto L229
	}
L131:
	;
	v480 = int32(-2)
	v482 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v482 == int32(68) {
		v1938 = v37
		v1939 = v480
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if v486 == int32(46) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	if v81 != int32(3) {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	if v486 == int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	v1438 = int32(0)
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v844
	v1455 = v98
	v1456 = v99
	goto L12
L137:
	;
	if v81 != int32(31) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v474)))
	if v767&int32(3) == int32(0) {
		v791 = v767
		goto L206
	} else {
		goto L207
	}
L140:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L141:
	;
	goto L142
L142:
	;
	if l6 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L144:
	;
	goto L145
L145:
	;
	if v479 < int32(0) {
		v1938 = v37
		v1939 = v480
		goto L4
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(14)
	v505 = v479 + int32(32044)
	v506 = int32(146097)
	v507 = base.I32_div_u_s(v505, v506)
	v508 = int32(3)
	v514 = int32(2)
	v519 = base.I32_div_u_s((v507*int32(1073595727)+v505)<<(uint(v514)%32)|v508, v506)
	v522 = v479 + v507*v508 + v519 + int32(32104)
	v523 = int32(1461)
	v524 = base.I32_div_u_s(v522, v523)
	v527 = v524*int32(-1461) + v522
	v529 = v527 << (uint(v514) % 32)
	if base.Ui32(v523) <= base.Ui32(v529) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v542 = base.I32_div_u_s(v529, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v542 + v524<<(uint(int32(2))%32) - int32(4800)
	v549 = int32(1)
	v551 = v540 + int32(123)
	v555 = int32(base.Ui32(v551*int32(2141)) >> (uint(int32(16)) % 32))
	v559 = base.I32_rem_u_s(v555+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v559 + v549
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v551 - int32(base.Ui32(v555*int32(7834))>>(uint(int32(8))%32))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	if v569 != int32(46) {
		v844 = v549
		goto L136
	} else {
		goto L151
	}
L148:
	;
	v535 = base.I32_rem_u_s(v527+int32(305), int32(365))
	v540 = v535
	goto L147
L149:
	;
	goto L150
L150:
	;
	v539 = base.I32_rem_u_s(v527+int32(306), int32(366))
	v540 = v539
	goto L147
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v485
	v574 = v485 + int32(1)
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	if v575 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L153:
	;
	v734 = base.F64_mul(v732, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v734), float64(9.223372036854776e+18)) != 0 {
		goto L200
	} else {
		goto L201
	}
L154:
	;
	v732 = float64(0)
	goto L153
L155:
	;
	goto L156
L156:
	;
	v579 = int32(539317)
	v583 = m.G0
	v585 = v583 - int32(32)
	v586 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v585)+24)) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v585)+16)) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v585)+8)) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v585))) = v586
	v594 = int32(*(*uint8)(unsafe.Add(mBase, _consts[180])))
	if v594 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	if v574&int32(3) == int32(0) {
		v686 = v574
		goto L180
	} else {
		goto L181
	}
L158:
	;
	v662 = int32(0)
	goto L157
L159:
	;
	goto L160
L160:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, _consts[181])))
	if v598 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v602 = v574
	goto L164
L162:
	;
	goto L163
L163:
	;
	v612 = v579
	v613 = v594
	goto L167
L164:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602))))
	if v608 == v594 {
		v602 = v602 + int32(1)
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v662 = v602 - v574
	goto L157
L166:
	;
	goto L165
L167:
	;
	v620 = v585 + int32(base.Ui32(v613)>>(uint(int32(3))%32))&int32(28)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v622 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v621 | v622<<(uint(v613)%32)
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+1)))
	if v626 != 0 {
		v612 = v612 + v622
		v613 = v626
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	if v629 == int32(0) {
		v654 = v574
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L168
L170:
	;
	v662 = v654 - v574
	goto L157
L171:
	;
	v633 = v574
	v634 = v629
	goto L172
L172:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v585+int32(base.Ui32(v634)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v642)>>(uint(v634)%32))&int32(1) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v654 = v650
	goto L170
L174:
	;
	v654 = v633
	goto L170
L175:
	;
	goto L176
L176:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633)+1)))
	v650 = v633 + int32(1)
	if v648 != 0 {
		v633 = v650
		v634 = v648
		goto L172
	} else {
		goto L177
	}
L177:
	;
	goto L173
L178:
	;
	if v662 != v719 {
		goto L152
	} else {
		goto L195
	}
L179:
	;
	v719 = v711 - v574
	goto L178
L180:
	;
	v690 = v686
	goto L189
L181:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	if v670 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v719 = int32(0)
	goto L178
L183:
	;
	goto L184
L184:
	;
	v675 = v574
	goto L185
L185:
	;
	v679 = v675 + int32(1)
	if v679&int32(3) == int32(0) {
		v686 = v679
		goto L180
	} else {
		goto L187
	}
L186:
	;
	v711 = v679
	goto L179
L187:
	;
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	if v684 != 0 {
		v675 = v679
		goto L185
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v690)))
	v699 = int32(-2139062144)
	if (int32(16843008)-v696|v696)&v699 == v699 {
		v690 = v690 + int32(4)
		goto L189
	} else {
		goto L191
	}
L190:
	;
	v705 = v690
	goto L192
L191:
	;
	goto L190
L192:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705))))
	if v709 != 0 {
		v705 = v705 + int32(1)
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v711 = v705
	goto L179
L194:
	;
	goto L193
L195:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v726 = F_strtod(m, v485, v37+int32(8))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L28
	} else {
		goto L196
	}
L196:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728))))
	if v729 != 0 {
		goto L152
	} else {
		goto L197
	}
L197:
	;
	v731 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v731 != 0 {
		goto L152
	} else {
		goto L198
	}
L198:
	;
	v732 = v726
	goto L153
L199:
	;
	v743 = base.I64_div_s(v740, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v54))) = uint32(v743)
	v748 = base.I64_extend32_s(v743)*int64(-3600000000) + v740
	v750 = base.I64_div_s(v748, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4+v58))) = uint32(v750)
	v755 = base.I64_extend32_s(v750)*int64(-60000000) + v748
	v757 = base.I64_div_s(v755, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4))) = uint32(v757)
	v761 = v757*int64(4293967296) + v755
	*(*uint32)(unsafe.Add(mBase, uint32(l5))) = uint32(v761)
	goto L203
L200:
	;
	v738 = base.I64_trunc_f64_s(v734)
	v740 = v738
	goto L199
L201:
	;
	goto L202
L202:
	;
	v740 = int64(-9223372036854775807 - 1)
	goto L199
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(31758)
	v844 = v549
	goto L136
L204:
	;
	v831 = F_DecodeNumberField(m, v824, v767, v79|int32(14), v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L28
	} else {
		goto L221
	}
L205:
	;
	v824 = v816 - v767
	goto L204
L206:
	;
	v795 = v791
	goto L215
L207:
	;
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v767))))
	if v775 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v824 = int32(0)
	goto L204
L209:
	;
	goto L210
L210:
	;
	v780 = v767
	goto L211
L211:
	;
	v784 = v780 + int32(1)
	if v784&int32(3) == int32(0) {
		v791 = v784
		goto L206
	} else {
		goto L213
	}
L212:
	;
	v816 = v784
	goto L205
L213:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784))))
	if v789 != 0 {
		v780 = v784
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v795)))
	v804 = int32(-2139062144)
	if (int32(16843008)-v801|v801)&v804 == v804 {
		v795 = v795 + int32(4)
		goto L215
	} else {
		goto L217
	}
L216:
	;
	v810 = v795
	goto L218
L217:
	;
	goto L216
L218:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810))))
	if v814 != 0 {
		v810 = v810 + int32(1)
		goto L218
	} else {
		goto L220
	}
L219:
	;
	v816 = v810
	goto L205
L220:
	;
	goto L219
L221:
	;
	if v831 < int32(0) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L223:
	;
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v831
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	if v838 != int32(31744) {
		v1938 = v37
		v1939 = int32(-1)
		goto L4
	} else {
		goto L225
	}
L225:
	;
	v844 = v96
	goto L136
L226:
	;
	v908 = int32(46)
	v909 = F___strchrnul(m, v850, v908)
	mBase = m.M
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909))))
	if v911 == v908 {
		goto L244
	} else {
		goto L245
	}
L227:
	;
	v907 = v899 - v850
	goto L226
L228:
	;
	v878 = v874
	goto L237
L229:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850))))
	if v858 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v907 = int32(0)
	goto L226
L231:
	;
	goto L232
L232:
	;
	v863 = v850
	goto L233
L233:
	;
	v867 = v863 + int32(1)
	if v867&int32(3) == int32(0) {
		v874 = v867
		goto L228
	} else {
		goto L235
	}
L234:
	;
	v899 = v867
	goto L227
L235:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867))))
	if v872 != 0 {
		v863 = v867
		goto L233
	} else {
		goto L236
	}
L236:
	;
	goto L234
L237:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v878)))
	v887 = int32(-2139062144)
	if (int32(16843008)-v884|v884)&v887 == v887 {
		v878 = v878 + int32(4)
		goto L237
	} else {
		goto L239
	}
L238:
	;
	v893 = v878
	goto L240
L239:
	;
	goto L238
L240:
	;
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893))))
	if v897 != 0 {
		v893 = v893 + int32(1)
		goto L240
	} else {
		goto L242
	}
L241:
	;
	v899 = v893
	goto L227
L242:
	;
	goto L241
L243:
	;
	if v915 != 0 {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	v915 = v909
	goto L246
L245:
	;
	v915 = int32(0)
	goto L246
L246:
	;
	goto L243
L247:
	;
	if l2 == int32(1) {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	goto L249
L249:
	;
	v1003 = v79 | int32(14)
	if int32(5) <= v907 {
		goto L280
	} else {
		goto L281
	}
L250:
	;
	if v915&int32(3) == int32(0) {
		v951 = v915
		goto L258
	} else {
		goto L259
	}
L251:
	;
	if v85 != 0 {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v916 != int32(2) {
		goto L250
	} else {
		goto L253
	}
L253:
	;
	v924 = F_DecodeDate(m, v850, v79, v37+int32(68), v37+int32(59), l4)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L28
	} else {
		goto L254
	}
L254:
	;
	if v924 == int32(0) {
		v1438 = int32(0)
		v1443 = v86
		v1444 = v87
		v1451 = v94
		v1453 = v96
		v1455 = v98
		v1456 = v99
		goto L12
	} else {
		goto L255
	}
L255:
	;
	v1938 = v37
	v1939 = v924
	goto L4
L256:
	;
	if base.Ui32(v907-v984) < base.Ui32(int32(3)) {
		goto L273
	} else {
		goto L274
	}
L257:
	;
	v984 = v976 - v915
	goto L256
L258:
	;
	v955 = v951
	goto L267
L259:
	;
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915))))
	if v935 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v984 = int32(0)
	goto L256
L261:
	;
	goto L262
L262:
	;
	v940 = v915
	goto L263
L263:
	;
	v944 = v940 + int32(1)
	if v944&int32(3) == int32(0) {
		v951 = v944
		goto L258
	} else {
		goto L265
	}
L264:
	;
	v976 = v944
	goto L257
L265:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v944))))
	if v949 != 0 {
		v940 = v944
		goto L263
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v955)))
	v964 = int32(-2139062144)
	if (int32(16843008)-v961|v961)&v964 == v964 {
		v955 = v955 + int32(4)
		goto L267
	} else {
		goto L269
	}
L268:
	;
	v970 = v955
	goto L270
L269:
	;
	goto L268
L270:
	;
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970))))
	if v974 != 0 {
		v970 = v970 + int32(1)
		goto L270
	} else {
		goto L272
	}
L271:
	;
	v976 = v970
	goto L257
L272:
	;
	goto L271
L273:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L274:
	;
	goto L275
L275:
	;
	v995 = F_DecodeNumberField(m, v907, v850, v79|int32(14), v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L28
	} else {
		goto L276
	}
L276:
	;
	if v995 < int32(0) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L278:
	;
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v995
	v1438 = int32(0)
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L280:
	;
	v1010 = F_DecodeNumberField(m, v907, v850, v1003, v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L28
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1017 = int32(0)
	v1023 = F_DecodeNumber(m, v907, v850, v1017, v1003, v37+int32(68), l4, l5, v37+int32(59))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L28
	} else {
		goto L287
	}
L283:
	;
	if v1010 < int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L285:
	;
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v1010
	v1438 = int32(0)
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L287:
	;
	if v1023 == int32(0) {
		v1438 = v1017
		v1443 = v86
		v1444 = v87
		v1451 = v94
		v1453 = v96
		v1455 = v98
		v1456 = v99
		goto L12
	} else {
		goto L288
	}
L288:
	;
	v1938 = v37
	v1939 = v1023
	goto L4
L289:
	;
	if v1035 != 0 {
		v1938 = v37
		v1939 = v1035
		goto L4
	} else {
		goto L290
	}
L290:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v37)+64))
	if v1037 == int32(31) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[1041])))
	if v1043 != 0 {
		goto L296
	} else {
		goto L297
	}
L292:
	;
	v1283 = v1037
	goto L293
L293:
	;
	if v1283 == int32(8) {
		v1474 = v79
		v1476 = v81
		v1481 = v86
		v1482 = v87
		v1489 = v94
		v1491 = v96
		v1493 = v98
		v1494 = v99
		goto L11
	} else {
		goto L341
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+64)) = v1268
	*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = v1250
	v1283 = v1268
	goto L293
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[1041]))) = v1207
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+12))
	v1233 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1207)+11)))
	v1250 = v1232
	v1268 = v1233
	goto L294
L296:
	;
	goto L301
L297:
	;
	goto L298
L298:
	;
	v1091 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1040))))
	v1103 = int32(1633824)
	v1106 = int32(1634960)
	goto L314
L299:
	;
	if v1080-v1081 == int32(0) {
		v1207 = v1043
		goto L295
	} else {
		goto L313
	}
L301:
	;
	goto L302
L302:
	;
	v1050 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040))))
	if v1050 != 0 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1051 = v1040
	v1052 = v1043
	v1053 = int32(10)
	v1054 = v1050
	goto L307
L304:
	;
	v1076 = v1043
	v1080 = int32(0)
	goto L305
L305:
	;
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1076))))
	goto L299
L306:
	;
	v1076 = v1071
	v1080 = v1073
	goto L305
L307:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052))))
	if v1054 != v1056 {
		v1071 = v1052
		v1073 = v1054
		goto L306
	} else {
		goto L309
	}
L308:
	;
	v1071 = v1065
	v1073 = int32(0)
	goto L306
L309:
	;
	if v1056 == int32(0) {
		v1071 = v1052
		v1073 = v1054
		goto L306
	} else {
		goto L310
	}
L310:
	;
	v1061 = v1053 - int32(1)
	if v1061 == int32(0) {
		v1071 = v1052
		v1073 = v1054
		goto L306
	} else {
		goto L311
	}
L311:
	;
	v1064 = int32(1)
	v1065 = v1052 + v1064
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1051)+1)))
	if v1066 != 0 {
		v1051 = v1051 + v1064
		v1052 = v1065
		v1053 = v1061
		v1054 = v1066
		goto L307
	} else {
		goto L312
	}
L312:
	;
	goto L308
L313:
	;
	goto L298
L314:
	;
	v1133 = v1103 + (v1106-v1103)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v1134 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1133))))
	v1135 = v1091 - v1134
	if v1135 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	v1250 = v1186
	v1268 = int32(31)
	goto L294
L316:
	;
	goto L321
L317:
	;
	v1185 = v1135
	goto L318
L318:
	;
	v1186 = int32(0)
	v1190 = base.B2i32(v1185 < v1186)
	if v1185 < v1186 {
		goto L334
	} else {
		goto L335
	}
L319:
	;
	if v1176 == int32(0) {
		v1207 = v1133
		goto L295
	} else {
		goto L333
	}
L321:
	;
	goto L322
L322:
	;
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040))))
	if v1144 != 0 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1145 = v1040
	v1146 = v1133
	v1147 = int32(10)
	v1148 = v1144
	goto L327
L324:
	;
	v1170 = v1133
	v1174 = int32(0)
	goto L325
L325:
	;
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1170))))
	v1176 = v1174 - v1175
	goto L319
L326:
	;
	v1170 = v1165
	v1174 = v1167
	goto L325
L327:
	;
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1146))))
	if v1148 != v1150 {
		v1165 = v1146
		v1167 = v1148
		goto L326
	} else {
		goto L329
	}
L328:
	;
	v1165 = v1159
	v1167 = int32(0)
	goto L326
L329:
	;
	if v1150 == int32(0) {
		v1165 = v1146
		v1167 = v1148
		goto L326
	} else {
		goto L330
	}
L330:
	;
	v1155 = v1147 - int32(1)
	if v1155 == int32(0) {
		v1165 = v1146
		v1167 = v1148
		goto L326
	} else {
		goto L331
	}
L331:
	;
	v1158 = int32(1)
	v1159 = v1146 + v1158
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145)+1)))
	if v1160 != 0 {
		v1145 = v1145 + v1158
		v1146 = v1159
		v1147 = v1155
		v1148 = v1160
		goto L327
	} else {
		goto L332
	}
L332:
	;
	goto L328
L333:
	;
	v1185 = v1176
	goto L318
L334:
	;
	v1191 = v1133 - int32(16)
	goto L336
L335:
	;
	v1191 = v1106
	goto L336
L336:
	;
	if v1185 < v1186 {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1194 = v1103
	goto L339
L338:
	;
	v1194 = v1133 + int32(16)
	goto L339
L339:
	;
	if base.Ui32(v1194) <= base.Ui32(v1191) {
		v1103 = v1194
		v1106 = v1191
		goto L314
	} else {
		goto L340
	}
L340:
	;
	goto L315
L341:
	;
	v1308 = int32(1) << (uint(v1283) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1308
	v1310 = int32(-1)
	switch v1283 {
	case 0:
		goto L351
	default:
		v1938 = v37
		v1939 = v1310
		goto L4
	case 5:
		goto L348
	case 6:
		goto L349
	case 7:
		goto L347
	case 9:
		goto L346
	case 17:
		goto L344
	case 18:
		goto L345
	case 23:
		goto L343
	case 28:
		goto L350
	case 31:
		goto L342
	}
L342:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	v1385 = F_pg_tzset(m, v1384)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L28
	} else {
		goto L361
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(0)
	if v81 != 0 {
		v1938 = v37
		v1939 = v1310
		goto L4
	} else {
		goto L360
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(0)
	if v81 != 0 {
		v1938 = v37
		v1939 = v1310
		goto L4
	} else {
		goto L359
	}
L345:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1438 = v81
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = base.B2i32(v1375 == int32(1))
	goto L12
L346:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1438 = v81
	v1443 = v1374
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1308 | int32(32)
	if l6 == int32(0) {
		v1938 = v37
		v1939 = v1310
		goto L4
	} else {
		goto L358
	}
L348:
	;
	v1355 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1355
	if l6 == v1355 {
		v1938 = v37
		v1939 = v1310
		goto L4
	} else {
		goto L357
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1308 | int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v1938 = v37
		v1939 = v1310
		goto L4
	} else {
		goto L356
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = v1308 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v1938 = v37
		v1939 = v1310
		goto L4
	} else {
		goto L355
	}
L351:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	switch v1311 - int32(12) {
	case 0:
		goto L353
	default:
		v1938 = v37
		v1939 = v1310
		goto L4
	case 4:
		goto L352
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(31776)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(3)
	v1325 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1325
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
	v1438 = v81
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+68)) = int32(31744)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(3)
	F_GetCurrentTimeUsec(m, l4, l5, int32(0))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L28
	} else {
		goto L354
	}
L354:
	;
	v1438 = v81
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L355:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1338 - v1339
	v1438 = v81
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L356:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1350
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1438 = v81
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L357:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1360
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1438 = v81
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L358:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v37)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(4)
	v1438 = v81
	v1443 = v86
	v1444 = v87
	v1451 = v1371
	v1453 = v96
	v1455 = v1370
	v1456 = v99
	goto L12
L359:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1438 = v1380
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L360:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v1438 = v1383
	v1443 = v86
	v1444 = v87
	v1451 = v94
	v1453 = v96
	v1455 = v98
	v1456 = v99
	goto L12
L361:
	;
	if v1385 != 0 {
		v1408 = v1385
		goto L13
	} else {
		goto L362
	}
L362:
	;
	v1938 = v37
	v1939 = v1310
	goto L4
L363:
	;
	v1938 = v37
	v1939 = int32(-1)
	goto L4
L364:
	;
	goto L365
L365:
	;
	v1474 = v1459 | v79
	v1476 = v1438
	v1481 = v1443
	v1482 = v1444
	v1489 = v1451
	v1491 = v1453
	v1493 = v1455
	v1494 = v1456
	goto L11
L366:
	;
	goto L10
L367:
	;
	v1509 = v37
	v1512 = v1474
	v1519 = v1481
	v1520 = v1482
	v1527 = v1489
	v1529 = v1491
	v1531 = v1493
	v1532 = v1494
	goto L5
L368:
	;
	if v1703 != 0 {
		v1938 = v1509
		v1939 = v1703
		goto L4
	} else {
		goto L406
	}
L369:
	;
	if v1512&int32(32768) != 0 {
		goto L387
	} else {
		goto L388
	}
L370:
	;
	if v1512&int32(4) == int32(0) {
		goto L369
	} else {
		goto L371
	}
L371:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1532 != 0 {
		goto L374
	} else {
		goto L375
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1560
	goto L369
L373:
	;
	v1560 = int32(1) - v1540
	goto L372
L374:
	;
	if int32(0) < v1540 {
		goto L373
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	if v1535 != 0 {
		goto L378
	} else {
		goto L379
	}
L377:
	;
	v1703 = int32(-2)
	goto L368
L378:
	;
	if v1540 < int32(0) {
		goto L381
	} else {
		goto L382
	}
L379:
	;
	goto L380
L380:
	;
	if int32(0) < v1540 {
		goto L369
	} else {
		goto L386
	}
L381:
	;
	v1703 = int32(-2)
	goto L368
L382:
	;
	goto L383
L383:
	;
	if base.Ui32(v1540) <= base.Ui32(int32(69)) {
		v1560 = v1540 + int32(2000)
		goto L372
	} else {
		goto L384
	}
L384:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1540) {
		goto L369
	} else {
		goto L385
	}
L385:
	;
	v1560 = v1540 + int32(1900)
	goto L372
L386:
	;
	v1703 = int32(-2)
	goto L368
L387:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v1571 = v1566 + int32(4799)
	v1573 = base.I32_div_s(v1571, int32(4))
	v1576 = base.I32_div_s(v1571, int32(-100))
	v1579 = base.I32_div_s(v1571, int32(400))
	v1580 = v1565 + v1566*int32(365) + v1573 + v1576 + v1579
	v1582 = v1580 + int32(1751940)
	v1583 = int32(146097)
	v1584 = base.I32_div_u_s(v1582, v1583)
	v1585 = int32(3)
	v1591 = int32(2)
	v1596 = base.I32_div_u_s((v1584*int32(1073595727)+v1582)<<(uint(v1591)%32)|v1585, v1583)
	v1599 = v1580 + v1584*v1585 + v1596 + int32(1752000)
	v1600 = int32(1461)
	v1601 = base.I32_div_u_s(v1599, v1600)
	v1604 = v1601*int32(-1461) + v1599
	v1606 = v1604 << (uint(v1591) % 32)
	if base.Ui32(v1600) <= base.Ui32(v1606) {
		goto L391
	} else {
		goto L392
	}
L388:
	;
	goto L389
L389:
	;
	if v1512&int32(2) == int32(0) {
		goto L394
	} else {
		goto L395
	}
L390:
	;
	v1619 = base.I32_div_u_s(v1606, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1619 + v1601<<(uint(int32(2))%32) - int32(4800)
	v1627 = v1617 + int32(123)
	v1631 = int32(base.Ui32(v1627*int32(2141)) >> (uint(int32(16)) % 32))
	v1635 = base.I32_rem_u_s(v1631+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v1635 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v1627 - int32(base.Ui32(v1631*int32(7834))>>(uint(int32(8))%32))
	goto L389
L391:
	;
	v1612 = base.I32_rem_u_s(v1604+int32(305), int32(365))
	v1617 = v1612
	goto L390
L392:
	;
	goto L393
L393:
	;
	v1616 = base.I32_rem_u_s(v1604+int32(306), int32(366))
	v1617 = v1616
	goto L390
L394:
	;
	if v1512&int32(8) == int32(0) {
		goto L397
	} else {
		goto L398
	}
L395:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if base.Ui32(int32(-12)) <= base.Ui32(v1652-int32(13)) {
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v1703 = int32(-3)
	goto L368
L397:
	;
	v1668 = int32(14)
	if v1512&v1668 != v1668 {
		goto L400
	} else {
		goto L401
	}
L398:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if base.Ui32(int32(-31)) <= base.Ui32(v1662-int32(32)) {
		goto L397
	} else {
		goto L399
	}
L399:
	;
	v1703 = int32(-3)
	goto L368
L400:
	;
	v1703 = int32(0)
	goto L368
L401:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1674&int32(3) != 0 {
		v1684 = int32(0)
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1684*int32(52)+v1687<<(uint(int32(2))%32))+uint32(_consts[1042])))
	if v1672 <= v1693 {
		goto L400
	} else {
		goto L405
	}
L403:
	;
	v1679 = base.I32_rem_s(v1674, int32(100))
	if v1679 != 0 {
		v1684 = int32(1)
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v1681 = base.I32_rem_s(v1674, int32(400))
	v1684 = base.B2i32(v1681 == int32(0))
	goto L402
L405:
	;
	v1703 = int32(-2)
	goto L368
L406:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v1519 == int32(2) {
		goto L408
	} else {
		goto L409
	}
L407:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v1727 = int32(1)
	if base.Ui32(int32(24)) < base.Ui32(v1721) {
		v1749 = v1727
		goto L420
	} else {
		goto L421
	}
L408:
	;
	v1721 = v1704
	goto L407
L409:
	;
	goto L410
L410:
	;
	if int32(12) < v1704 {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1938 = v1509
	v1939 = int32(-2)
	goto L4
L412:
	;
	goto L413
L413:
	;
	switch v1519 {
	case 0:
		goto L416
	case 1:
		goto L415
	default:
		v1721 = v1704
		goto L407
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v1719
	v1721 = v1719
	goto L407
L415:
	;
	v1713 = int32(12)
	if v1704 == v1713 {
		v1721 = v1713
		goto L407
	} else {
		goto L418
	}
L416:
	;
	if v1704 == int32(12) {
		v1719 = int32(0)
		goto L414
	} else {
		goto L417
	}
L417:
	;
	v1721 = v1704
	goto L407
L418:
	;
	v1719 = v1704 + int32(12)
	goto L414
L419:
	;
	if v1749 != 0 {
		goto L425
	} else {
		goto L426
	}
L420:
	;
	goto L419
L421:
	;
	if base.Ui32(int32(59)) < base.Ui32(v1723) {
		v1749 = v1727
		goto L420
	} else {
		goto L422
	}
L422:
	;
	if base.Ui32(int32(60)) < base.Ui32(v1724) {
		v1749 = v1727
		goto L420
	} else {
		goto L423
	}
L423:
	;
	if base.Ui32(int32(1000000)) < base.Ui32(v1725) {
		v1749 = v1727
		goto L420
	} else {
		goto L424
	}
L424:
	;
	v1737 = int32(60)
	v1749 = base.B2i32(base.Ui64(int64(86400000000)) < base.Ui64(base.I64_extend_i32_u(v1725)+base.I64_extend_i32_u((v1721*v1737+v1723)*v1737+v1724)*int64(1000000)))
	goto L420
L425:
	;
	v1938 = v1509
	v1939 = int32(-2)
	goto L4
L426:
	;
	goto L427
L427:
	;
	v1751 = int32(-1)
	v1752 = int32(31744)
	if v1512&v1752 != v1752 {
		v1938 = v1509
		v1939 = v1751
		goto L4
	} else {
		goto L428
	}
L428:
	;
	if v1520 != 0 {
		goto L430
	} else {
		goto L431
	}
L429:
	;
	v1938 = v1509
	v1939 = v1929
	goto L4
L430:
	;
	if v1512&int32(268435456) != 0 {
		v1929 = v1751
		goto L429
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	if v1527 != 0 {
		goto L448
	} else {
		goto L449
	}
L433:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v1520)+uint32(_consts[1043])))
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1520)+264))
	if v1764 < int32(2) {
		goto L436
	} else {
		goto L437
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1807
	goto L432
L435:
	;
	if v1796 != 0 {
		goto L444
	} else {
		goto L445
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1509+int32(72)))) = v1763
	v1796 = int32(1)
	goto L435
L437:
	;
	v1770 = int32(1)
	goto L438
L438:
	;
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1520+int32(18280)+v1770<<(uint(int32(4))%32))))
	if v1763 == v1778 {
		goto L440
	} else {
		goto L441
	}
L439:
	;
	v1796 = int32(0)
	goto L435
L440:
	;
	v1781 = v1770 + int32(1)
	if v1764 != v1781 {
		v1770 = v1781
		goto L438
	} else {
		goto L443
	}
L441:
	;
	goto L442
L442:
	;
	goto L439
L443:
	;
	goto L436
L444:
	;
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+72))
	v1807 = int32(0) - v1798
	goto L434
L445:
	;
	goto L446
L446:
	;
	v1800 = int32(14)
	if v1512&v1800 != v1800 {
		v1929 = v1751
		goto L429
	} else {
		goto L447
	}
L447:
	;
	v1806 = F_DetermineTimeZoneOffsetInternal(m, l4, v1520, v1509+int32(8))
	mBase = m.M
	v1807 = v1806
	goto L434
L448:
	;
	if v1512&int32(268435456) != 0 {
		v1938 = v1509
		v1939 = v1751
		goto L4
	} else {
		goto L451
	}
L449:
	;
	goto L450
L450:
	;
	v1890 = int32(0)
	if l6 == v1890 {
		v1938 = v1509
		v1939 = v1890
		goto L4
	} else {
		goto L466
	}
L451:
	;
	switch v1512 & int32(14) {
	case 0:
		goto L453
	default:
		v1938 = v1509
		v1939 = v1751
		goto L4
	case 14:
		goto L454
	}
L452:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+16)) = v1826
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+12)) = v1828
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+8)) = v1830
	v1833 = v1509 + int32(8)
	v1837 = m.G0
	v1839 = v1837 - int32(288)
	m.G0 = v1839
	v1843 = F_DetermineTimeZoneOffsetInternal(m, v1833, v1527, v1839+int32(280))
	mBase = m.M
	v1847 = F_strlcpy(m, v1839+int32(16), v1531, int32(256))
	mBase = m.M
	v1848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1839)+16)))
	if v1848 != 0 {
		goto L457
	} else {
		goto L458
	}
L453:
	;
	F_GetCurrentTimeUsec(m, v1509+int32(8), v1509+int32(72), int32(0))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L28
	} else {
		goto L455
	}
L454:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+28)) = v1813
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+24)) = v1815
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+20)) = v1817
	goto L452
L455:
	;
	goto L452
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1883
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1888
	goto L450
L457:
	;
	v1852 = v1839 + int32(16)
	v1856 = v1848
	goto L460
L458:
	;
	goto L459
L459:
	;
	v1876 = F_pg_interpret_timezone_abbrev(m, v1839+int32(16), v1839+int32(280), v1839+int32(12), v1839+int32(8), v1527)
	mBase = m.M
	if v1876 != 0 {
		goto L463
	} else {
		goto L464
	}
L460:
	;
	v1857 = F_pg_toupper(m, v1856)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1852))) = uint8(v1857)
	v1860 = v1852 + int32(1)
	v1861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1860))))
	if v1861 != 0 {
		v1852 = v1860
		v1856 = v1861
		goto L460
	} else {
		goto L462
	}
L461:
	;
	goto L459
L462:
	;
	goto L461
L463:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+12))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1839)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1833)+32)) = v1878
	v1883 = int32(0) - v1877
	goto L465
L464:
	;
	v1883 = v1843
	goto L465
L465:
	;
	m.G0 = v1839 + int32(288)
	goto L456
L466:
	;
	if v1512&int32(32) != 0 {
		v1938 = v1509
		v1939 = v1890
		goto L4
	} else {
		goto L467
	}
L467:
	;
	if v1512&int32(268435456) != 0 {
		goto L470
	} else {
		goto L471
	}
L468:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+16)) = v1913
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+12)) = v1915
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+8)) = v1917
	v1922 = *(*int32)(unsafe.Add(mBase, _consts[329]))
	v1925 = F_DetermineTimeZoneOffsetInternal(m, v1509+int32(8), v1922, v1509+int32(72))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1925
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1927
	v1929 = v1890
	goto L429
L469:
	;
	F_GetCurrentTimeUsec(m, v1509+int32(8), v1509+int32(72), int32(0))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L28
	} else {
		goto L473
	}
L470:
	;
	v1938 = v1509
	v1939 = int32(-1)
	goto L4
L471:
	;
	switch v1512 & int32(14) {
	case 0:
		goto L469
	default:
		goto L470
	case 14:
		goto L472
	}
L472:
	;
	v1899 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+28)) = v1899
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+24)) = v1901
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+20)) = v1903
	goto L468
L473:
	;
	goto L468
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[95])))
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
	v69 = F_AllocateFile(m, v12+int32(1264), int32(227660))
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
	v45 = F_pg_snprintf(m, v12+int32(1200), int32(64), int32(12756), v12+int32(128))
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
	v63 = F_pg_snprintf(m, v12+int32(1264), int32(1024), int32(12749), v12+int32(144))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	v54 = F_RestoreArchivedFile(m, v12+int32(1264), v12+int32(1200), int32(500692), int64(0), int32(0))
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
	v72 = *(*int32)(unsafe.Add(mBase, _consts[96]))
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
	v85 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v85 != int32(44) {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[96]))
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
	v176 = *(*int32)(unsafe.Add(mBase, _consts[96]))
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
	v146 = F_sscanf(m, v12+int32(160), int32(508015), v12+int32(112))
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
	v185 = *(*int32)(unsafe.Add(mBase, _consts[96]))
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
	F_errmsg(m, int32(200921), v12-int32(-64))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_errhint(m, int32(635898), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(491163), int32(164), int32(13131))
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
	F_errmsg(m, int32(200921), v12+int32(96))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errhint(m, int32(593359), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(491163), int32(169), int32(13131))
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
	F_errmsg(m, int32(200954), v12+int32(80))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(620152), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(491163), int32(174), int32(13131))
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
	F_errmsg(m, int32(294211), v12+int32(16))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(491163), int32(111), int32(13131))
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
	F_errmsg(m, int32(691678), v12+int32(32))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	F_errhint(m, int32(635846), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(491163), int32(195), int32(13131))
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
	F_errmsg(m, int32(295130), v12+int32(48))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(491163), int32(143), int32(13131))
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
		v39 = *(*int32)(unsafe.Add(mBase, _consts[329]))
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
