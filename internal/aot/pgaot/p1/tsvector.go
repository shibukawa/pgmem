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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
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
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
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
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v749 int32
	_ = v749
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1006 int32
	_ = v1006
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1189 int32
	_ = v1189
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1409 int32
	_ = v1409
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1461 int32
	_ = v1461
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1491 int32
	_ = v1491
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1666 int32
	_ = v1666
	var v1677 int32
	_ = v1677
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1704 int32
	_ = v1704
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1767 int32
	_ = v1767
	var v1770 int32
	_ = v1770
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1790 int32
	_ = v1790
	var v1797 int32
	_ = v1797
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1812 int32
	_ = v1812
	var v1819 int32
	_ = v1819
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1832 int32
	_ = v1832
	var v1839 int32
	_ = v1839
	var v1843 int32
	_ = v1843
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1889 int32
	_ = v1889
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1909 int32
	_ = v1909
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1916 int32
	_ = v1916
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1933 int32
	_ = v1933
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	v2 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = v30 + int32(8)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v37 = F_pg_detoast_datum(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v46 = v35
	v49 = v39
	v50 = v2
	goto L7
L5:
	;
	v235 = v2
	goto L6
L6:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v255 = int32(2)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v261 = v252 + v39 + int32(base.Ui32(v254)>>(uint(v255)%32)) + int32(base.Ui32(v258)>>(uint(v255)%32))
	v262 = F_palloc0(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L38
	}
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v67&int32(1) == int32(0) {
		v207 = v50
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v235 = v207
	goto L6
L9:
	;
	v227 = v49 - int32(1)
	if v227 != 0 {
		v46 = v46 + int32(4)
		v49 = v227
		v50 = v207
		goto L7
	} else {
		goto L37
	}
L10:
	;
	v72 = int32(1)
	v83 = v35 + v39<<(uint(int32(2))%32) + (int32(base.Ui32(v67)>>(uint(v72)%32))&int32(2047)+int32(base.Ui32(v67)>>(uint(int32(12))%32))+v72)&int32(4194302)
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83))))
	if v84 == int32(0) {
		v207 = v50
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v89 = v84 & int32(3)
	if v89 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v91 = v83
	v94 = int32(0)
	v95 = v84
	v97 = v50
	goto L15
L13:
	;
	v127 = v83
	v131 = v84
	v133 = v50
	goto L14
L14:
	;
	if base.Ui32(v84) < base.Ui32(int32(4)) {
		v207 = v133
		goto L9
	} else {
		goto L21
	}
L15:
	;
	v115 = v91 + int32(2)
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115))))
	v118 = v116 & int32(16383)
	if v118 < v97 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v127 = v115
	v131 = v122
	v133 = v120
	goto L14
L17:
	;
	v120 = v97
	goto L19
L18:
	;
	v120 = v118
	goto L19
L19:
	;
	v121 = int32(1)
	v122 = v95 - v121
	v124 = v94 + v121
	if v124 != v89 {
		v91 = v115
		v94 = v124
		v95 = v122
		v97 = v120
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v153 = v127
	v157 = v131
	v159 = v133
	goto L22
L22:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+2)))
	v178 = v176 & int32(16383)
	if v178 < v159 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v207 = v197
	goto L9
L24:
	;
	v180 = v159
	goto L26
L25:
	;
	v180 = v178
	goto L26
L26:
	;
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+4)))
	v183 = v181 & int32(16383)
	if v183 < v180 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v185 = v180
	goto L29
L28:
	;
	v185 = v183
	goto L29
L29:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+6)))
	v188 = v186 & int32(16383)
	if v188 < v185 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v190 = v185
	goto L32
L31:
	;
	v190 = v188
	goto L32
L32:
	;
	v192 = v153 + int32(8)
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192))))
	v195 = v193 & int32(16383)
	if v195 < v190 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v197 = v190
	goto L35
L34:
	;
	v197 = v195
	goto L35
L35:
	;
	v199 = v157 - int32(4)
	if v199 != 0 {
		v153 = v192
		v157 = v199
		v159 = v197
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
	v264 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v261 << (uint(v264) % 32)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v269 = v267 + v268
	*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v269
	v271 = int32(8)
	v272 = v37 + v271
	v275 = v272 + v252<<(uint(v264)%32)
	v277 = v262 + v271
	v280 = v277 + v269<<(uint(v264)%32)
	v283 = v35 + v39<<(uint(v264)%32)
	if v39 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v1281 != 0 {
		goto L221
	} else {
		goto L222
	}
L40:
	;
	v1273 = v277
	v1274 = v2
	v1276 = v35
	v1277 = v272
	v1281 = v39
	v1282 = v252
	goto L39
L41:
	;
	if v252 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v289 = v277
	v290 = v2
	v292 = v35
	v293 = v272
	v297 = v39
	v298 = v252
	goto L43
L43:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v313 = int32(1)
	v315 = int32(2047)
	v316 = int32(base.Ui32(v312)>>(uint(v313)%32)) & v315
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v321 = int32(base.Ui32(v317)>>(uint(v313)%32)) & v315
	if v321 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v1273 = v1269
	v1274 = v1257
	v1276 = v1259
	v1277 = v1260
	v1281 = v1263
	v1282 = v1264
	goto L39
L45:
	;
	v1269 = v289 + int32(4)
	if v1263 == int32(0) {
		v1273 = v1269
		v1274 = v1257
		v1276 = v1259
		v1277 = v1260
		v1281 = v1263
		v1282 = v1264
		goto L39
	} else {
		goto L219
	}
L46:
	;
	v1257 = v1245
	v1259 = v292 + int32(4)
	v1260 = v1247
	v1263 = v1254
	v1264 = v1250
	goto L45
L47:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v761 = int32(1)
	v763 = v757&int32(-2) | (v312|v317)&v761
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v763
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v763&int32(-4095) | v767&int32(4094)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v780 = int32(base.Ui32(v773)>>(uint(v761)%32)) & int32(2047)
	if v780 != 0 {
		goto L140
	} else {
		goto L141
	}
L48:
	;
	v1257 = v749
	v1259 = v292
	v1260 = v293 + int32(4)
	v1263 = v297
	v1264 = v298 - int32(1)
	goto L45
L49:
	;
	v740 = int32(1)
	v749 = (v447+v740)&int32(-2) + v631<<(uint(v740)%32) + int32(2)
	goto L48
L50:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v641 = int32(1)
	v643 = v638&int32(-2) | v317&v641
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v643
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v643&int32(-4095) | v647&int32(4094)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v660 = int32(base.Ui32(v653)>>(uint(v641)%32)) & int32(2047)
	if v660 != 0 {
		goto L123
	} else {
		goto L124
	}
L51:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v413 = int32(1)
	v415 = v410&int32(-2) | v312&v413
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v415
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v415&int32(-4095) | v419&int32(4094)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v432 = int32(base.Ui32(v425)>>(uint(v413)%32)) & int32(2047)
	if v432 != 0 {
		goto L87
	} else {
		goto L88
	}
L52:
	;
	if v402 < int32(0) {
		goto L50
	} else {
		goto L84
	}
L53:
	;
	if v316 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if v316 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v326 = int32(-1)
	goto L58
L57:
	;
	v326 = int32(0)
	goto L58
L58:
	;
	v402 = v326
	goto L52
L59:
	;
	v329 = int32(12)
	v331 = v283 + int32(base.Ui32(v317)>>(uint(v329)%32))
	v334 = v275 + int32(base.Ui32(v312)>>(uint(v329)%32))
	v335 = base.B2i32(base.Ui32(v321) < base.Ui32(v316))
	if base.Ui32(v321) < base.Ui32(v316) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v336 = v321
	goto L62
L61:
	;
	v336 = v316
	goto L62
L62:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v336) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	if v398 != 0 {
		v402 = v398
		goto L52
	} else {
		goto L81
	}
L64:
	;
	v398 = int32(0)
	goto L63
L65:
	;
	v372 = v367
	v373 = v368
	v374 = v369
	goto L75
L66:
	;
	if (v331|v334)&int32(3) != 0 {
		v367 = v331
		v368 = v334
		v369 = v336
		goto L65
	} else {
		goto L69
	}
L67:
	;
	v360 = v331
	v361 = v334
	v362 = v336
	goto L68
L68:
	;
	if v362 == int32(0) {
		goto L64
	} else {
		goto L74
	}
L69:
	;
	v344 = v331
	v345 = v334
	v346 = v336
	goto L70
L70:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	if v349 != v350 {
		v367 = v344
		v368 = v345
		v369 = v346
		goto L65
	} else {
		goto L72
	}
L71:
	;
	v360 = v355
	v361 = v353
	v362 = v357
	goto L68
L72:
	;
	v352 = int32(4)
	v353 = v345 + v352
	v355 = v344 + v352
	v357 = v346 - v352
	if base.Ui32(int32(3)) < base.Ui32(v357) {
		v344 = v355
		v345 = v353
		v346 = v357
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v367 = v360
	v368 = v361
	v369 = v362
	goto L65
L75:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v377 == v378 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v398 = v377 - v378
	goto L63
L77:
	;
	v380 = int32(1)
	v385 = v374 - v380
	if v385 != 0 {
		v372 = v372 + v380
		v373 = v373 + v380
		v374 = v385
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
	if v316 == v321 {
		goto L47
	} else {
		goto L82
	}
L82:
	;
	if v335 == int32(0) {
		goto L51
	} else {
		goto L83
	}
L83:
	;
	goto L50
L84:
	;
	if v402 == int32(0) {
		goto L47
	} else {
		goto L85
	}
L85:
	;
	goto L51
L86:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v435&int32(4095) | v290<<(uint(int32(12))%32)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v443 = int32(1)
	v447 = int32(base.Ui32(v442)>>(uint(v443)%32))&int32(2047) + v290
	if v435&v443 == int32(0) {
		v749 = v447
		goto L48
	} else {
		goto L90
	}
L87:
	;
	v433 = F__emscripten_memcpy_bulkmem(m, v290+v280, v275+int32(base.Ui32(v425)>>(uint(int32(12))%32)), v432)
	mBase = m.M
	goto L89
L88:
	;
	goto L89
L89:
	;
	goto L86
L90:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v467 = int32(1)
	v476 = v262 + v460<<(uint(int32(2))%32) + (int32(base.Ui32(v464)>>(uint(int32(12))%32))+int32(base.Ui32(v464)>>(uint(v467)%32))&int32(2047)+v467)&int32(4194302)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	if v477&v467 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	if v631 != 0 {
		goto L49
	} else {
		goto L121
	}
L92:
	;
	v516 = v476 + int32(8)
	if v464&int32(1) != 0 {
		goto L97
	} else {
		goto L98
	}
L93:
	;
	v482 = int32(1)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v512 = (int32(base.Ui32(v477)>>(uint(v482)%32))&int32(2047) + int32(base.Ui32(v477)>>(uint(int32(12))%32)) + v482) & int32(4194302)
	v513 = v493
	v514 = int32(0)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v495 = int32(1)
	v505 = (int32(base.Ui32(v477)>>(uint(v495)%32))&int32(2047) + int32(base.Ui32(v477)>>(uint(int32(12))%32)) + v495) & int32(4194302)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v505+(v37+v506<<(uint(int32(2))%32)))+8)))
	v512 = v505
	v513 = v506
	v514 = v511
	goto L92
L96:
	;
	if v514 == int32(0) {
		v618 = v523
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v519 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v516))))
	v523 = v519
	goto L96
L98:
	;
	goto L99
L99:
	;
	v520 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v516))) = uint16(v520)
	v523 = v520
	goto L96
L100:
	;
	v631 = v618&int32(65535) - v523
	goto L91
L101:
	;
	if base.Ui32(int32(255)) < base.Ui32(v523) {
		v600 = v523
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if v523 == v600&int32(65535) {
		v618 = v523
		goto L100
	} else {
		goto L120
	}
L103:
	;
	v1939 = int32(10)
	v536 = int32(0)
	v537 = int32(256)
	v538 = v537 - v523
	if base.Ui32(v538) <= base.Ui32(v537) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v542 = v538
	goto L106
L105:
	;
	v542 = v536
	goto L106
L106:
	;
	v544 = v536
	v545 = v523
	v548 = v523
	goto L107
L107:
	;
	if v545 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v600 = v589
	goto L102
L109:
	;
	v559 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v516+v545<<(uint(int32(1))%32)))))
	v560 = int32(16383)
	if v559&v560 == v560 {
		v600 = v548
		goto L102
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v564 = int32(1)
	v566 = v476 + v1939 + v545<<(uint(v564)%32)
	v569 = v37 + v513<<(uint(int32(2))%32) + v512 + v1939 + v544<<(uint(v564)%32)
	v570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v569))))
	v572 = v570 & int32(-16384)
	v573 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v566))))
	v574 = int32(16383)
	v576 = v572 | v573&v574
	*(*uint16)(unsafe.Add(mBase, uint32(v566))) = uint16(v576)
	v579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v569))))
	v582 = v235 + v579&v574
	if base.Ui32(v574) <= base.Ui32(v582) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L111
L113:
	;
	v585 = v574
	goto L115
L114:
	;
	v585 = v582
	goto L115
L115:
	;
	v586 = v572 | v585
	*(*uint16)(unsafe.Add(mBase, uint32(v566))) = uint16(v586)
	v588 = int32(1)
	v589 = v545 + v588
	*(*uint16)(unsafe.Add(mBase, uint32(v516))) = uint16(v589)
	v592 = v544 + v588
	if v514 == v592 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v600 = v589
	goto L102
L117:
	;
	goto L118
L118:
	;
	if v592 != v542 {
		v544 = v592
		v545 = v589
		v548 = v589
		goto L107
	} else {
		goto L119
	}
L119:
	;
	goto L108
L120:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v611 | int32(1)
	v615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v516))))
	v618 = v615
	goto L100
L121:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v632 & int32(-2)
	v749 = v447
	goto L48
L122:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v663&int32(4095) | v290<<(uint(int32(12))%32)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v671 = int32(1)
	v674 = int32(base.Ui32(v670)>>(uint(v671)%32)) & int32(2047)
	v675 = v674 + v290
	if v663&v671 != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v661 = F__emscripten_memcpy_bulkmem(m, v290+v280, v283+int32(base.Ui32(v653)>>(uint(int32(12))%32)), v660)
	mBase = m.M
	goto L125
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	v678 = int32(2)
	v679 = int32(1)
	v682 = (v675 + v679) & int32(-2)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v696 = v35 + v685<<(uint(v678)%32) + (int32(base.Ui32(v670)>>(uint(int32(12))%32))+v674+v679)&int32(4194302)
	if v670&v679 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v733 = v675
	goto L128
L128:
	;
	v1245 = v733
	v1247 = v293
	v1250 = v298
	v1254 = v297 - int32(1)
	goto L46
L129:
	;
	v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v696))))
	v704 = v699<<(uint(int32(1))%32) + int32(2)
	goto L131
L130:
	;
	v704 = v678
	goto L131
L131:
	;
	if v704 != 0 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	if v707&int32(1) != 0 {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v705 = F__emscripten_memcpy_bulkmem(m, v280+v682, v696, v704)
	mBase = m.M
	goto L135
L134:
	;
	goto L135
L135:
	;
	goto L132
L136:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v711 = int32(2)
	v714 = int32(1)
	v726 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+v710<<(uint(v711)%32)+(int32(base.Ui32(v707)>>(uint(v714)%32))&int32(2047)+int32(base.Ui32(v707)>>(uint(int32(12))%32))+v714)&int32(4194302)))))
	v731 = v726<<(uint(v714)%32) + v711
	goto L138
L137:
	;
	v731 = v678
	goto L138
L138:
	;
	v733 = v731 + v682
	goto L128
L139:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v783&int32(4095) | v290<<(uint(int32(12))%32)
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v791 = int32(1)
	v794 = int32(base.Ui32(v790)>>(uint(v791)%32)) & int32(2047)
	v795 = v794 + v290
	if v783&v791 == int32(0) {
		v1236 = v795
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v781 = F__emscripten_memcpy_bulkmem(m, v290+v280, v283+int32(base.Ui32(v773)>>(uint(int32(12))%32)), v780)
	mBase = m.M
	goto L142
L141:
	;
	goto L142
L142:
	;
	goto L139
L143:
	;
	v1239 = int32(1)
	v1245 = v1236
	v1247 = v293 + int32(4)
	v1250 = v298 - v1239
	v1254 = v297 - v1239
	goto L46
L144:
	;
	if v790&int32(1) != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v802 = int32(2)
	v803 = int32(1)
	v806 = (v795 + v803) & int32(-2)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v819 = v35 + v808<<(uint(v802)%32) + (int32(base.Ui32(v790)>>(uint(int32(12))%32))+v794+v803)&int32(4194302)
	v820 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v819))))
	v824 = v820<<(uint(v803)%32) + v802
	if v824 != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	goto L147
L147:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v1056 = int32(1)
	v1065 = v262 + v1049<<(uint(int32(2))%32) + (int32(base.Ui32(v1053)>>(uint(int32(12))%32))+int32(base.Ui32(v1053)>>(uint(v1056)%32))&int32(2047)+v1056)&int32(4194302)
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	if v1066&v1056 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L148:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	if v827&int32(1) != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v825 = F__emscripten_memcpy_bulkmem(m, v280+v806, v819, v824)
	mBase = m.M
	goto L151
L150:
	;
	goto L151
L151:
	;
	goto L148
L152:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v831 = int32(2)
	v834 = int32(1)
	v846 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+v830<<(uint(v831)%32)+(int32(base.Ui32(v827)>>(uint(v834)%32))&int32(2047)+int32(base.Ui32(v827)>>(uint(int32(12))%32))+v834)&int32(4194302)))))
	v851 = v846<<(uint(v834)%32) + v831
	goto L154
L153:
	;
	v851 = v802
	goto L154
L154:
	;
	v852 = v851 + v806
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	if v853&int32(1) == int32(0) {
		v1236 = v852
		goto L143
	} else {
		goto L155
	}
L155:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v873 = int32(1)
	v882 = v262 + v866<<(uint(int32(2))%32) + (int32(base.Ui32(v870)>>(uint(int32(12))%32))+int32(base.Ui32(v870)>>(uint(v873)%32))&int32(2047)+v873)&int32(4194302)
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	if v883&v873 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v1236 = (v1024&int32(65535)-v929)<<(uint(int32(1))%32) + v852
	goto L143
L157:
	;
	v922 = v882 + int32(8)
	if v870&int32(1) != 0 {
		goto L162
	} else {
		goto L163
	}
L158:
	;
	v888 = int32(1)
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v918 = (int32(base.Ui32(v883)>>(uint(v888)%32))&int32(2047) + int32(base.Ui32(v883)>>(uint(int32(12))%32)) + v888) & int32(4194302)
	v919 = v899
	v920 = int32(0)
	goto L157
L159:
	;
	goto L160
L160:
	;
	v901 = int32(1)
	v911 = (int32(base.Ui32(v883)>>(uint(v901)%32))&int32(2047) + int32(base.Ui32(v883)>>(uint(int32(12))%32)) + v901) & int32(4194302)
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v917 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v911+(v37+v912<<(uint(int32(2))%32)))+8)))
	v918 = v911
	v919 = v912
	v920 = v917
	goto L157
L161:
	;
	if v920 == int32(0) {
		v1024 = v929
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v925 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v922))))
	v929 = v925
	goto L161
L163:
	;
	goto L164
L164:
	;
	v926 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v922))) = uint16(v926)
	v929 = v926
	goto L161
L165:
	;
	goto L156
L166:
	;
	if base.Ui32(int32(255)) < base.Ui32(v929) {
		v1006 = v929
		goto L167
	} else {
		goto L168
	}
L167:
	;
	if v929 == v1006&int32(65535) {
		v1024 = v929
		goto L165
	} else {
		goto L185
	}
L168:
	;
	v1940 = int32(10)
	v942 = int32(0)
	v943 = int32(256)
	v944 = v943 - v929
	if base.Ui32(v944) <= base.Ui32(v943) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v948 = v944
	goto L171
L170:
	;
	v948 = v942
	goto L171
L171:
	;
	v950 = v942
	v951 = v929
	v954 = v929
	goto L172
L172:
	;
	if v951 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v1006 = v995
	goto L167
L174:
	;
	v965 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v922+v951<<(uint(int32(1))%32)))))
	v966 = int32(16383)
	if v965&v966 == v966 {
		v1006 = v954
		goto L167
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v970 = int32(1)
	v972 = v882 + v1940 + v951<<(uint(v970)%32)
	v975 = v37 + v919<<(uint(int32(2))%32) + v918 + v1940 + v950<<(uint(v970)%32)
	v976 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v975))))
	v978 = v976 & int32(-16384)
	v979 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v972))))
	v980 = int32(16383)
	v982 = v978 | v979&v980
	*(*uint16)(unsafe.Add(mBase, uint32(v972))) = uint16(v982)
	v985 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v975))))
	v988 = v235 + v985&v980
	if base.Ui32(v980) <= base.Ui32(v988) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	goto L176
L178:
	;
	v991 = v980
	goto L180
L179:
	;
	v991 = v988
	goto L180
L180:
	;
	v992 = v978 | v991
	*(*uint16)(unsafe.Add(mBase, uint32(v972))) = uint16(v992)
	v994 = int32(1)
	v995 = v951 + v994
	*(*uint16)(unsafe.Add(mBase, uint32(v922))) = uint16(v995)
	v998 = v950 + v994
	if v920 == v998 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1006 = v995
	goto L167
L182:
	;
	goto L183
L183:
	;
	if v998 != v948 {
		v950 = v998
		v951 = v995
		v954 = v995
		goto L172
	} else {
		goto L184
	}
L184:
	;
	goto L173
L185:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v1017 | int32(1)
	v1021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v922))))
	v1024 = v1021
	goto L165
L186:
	;
	if v1220 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L187:
	;
	v1105 = v1065 + int32(8)
	if v1053&int32(1) != 0 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v1071 = int32(1)
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v1101 = (int32(base.Ui32(v1066)>>(uint(v1071)%32))&int32(2047) + int32(base.Ui32(v1066)>>(uint(int32(12))%32)) + v1071) & int32(4194302)
	v1102 = v1082
	v1103 = int32(0)
	goto L187
L189:
	;
	goto L190
L190:
	;
	v1084 = int32(1)
	v1094 = (int32(base.Ui32(v1066)>>(uint(v1084)%32))&int32(2047) + int32(base.Ui32(v1066)>>(uint(int32(12))%32)) + v1084) & int32(4194302)
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v1100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1094+(v37+v1095<<(uint(int32(2))%32)))+8)))
	v1101 = v1094
	v1102 = v1095
	v1103 = v1100
	goto L187
L191:
	;
	if v1103 == int32(0) {
		v1207 = v1112
		goto L195
	} else {
		goto L196
	}
L192:
	;
	v1108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105))))
	v1112 = v1108
	goto L191
L193:
	;
	goto L194
L194:
	;
	v1109 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1105))) = uint16(v1109)
	v1112 = v1109
	goto L191
L195:
	;
	v1220 = v1207&int32(65535) - v1112
	goto L186
L196:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1112) {
		v1189 = v1112
		goto L197
	} else {
		goto L198
	}
L197:
	;
	if v1112 == v1189&int32(65535) {
		v1207 = v1112
		goto L195
	} else {
		goto L215
	}
L198:
	;
	v1941 = int32(10)
	v1125 = int32(0)
	v1126 = int32(256)
	v1127 = v1126 - v1112
	if base.Ui32(v1127) <= base.Ui32(v1126) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1131 = v1127
	goto L201
L200:
	;
	v1131 = v1125
	goto L201
L201:
	;
	v1133 = v1125
	v1134 = v1112
	v1137 = v1112
	goto L202
L202:
	;
	if v1134 != 0 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v1189 = v1178
	goto L197
L204:
	;
	v1148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105+v1134<<(uint(int32(1))%32)))))
	v1149 = int32(16383)
	if v1148&v1149 == v1149 {
		v1189 = v1137
		goto L197
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v1153 = int32(1)
	v1155 = v1065 + v1941 + v1134<<(uint(v1153)%32)
	v1158 = v37 + v1102<<(uint(int32(2))%32) + v1101 + v1941 + v1133<<(uint(v1153)%32)
	v1159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1158))))
	v1161 = v1159 & int32(-16384)
	v1162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1155))))
	v1163 = int32(16383)
	v1165 = v1161 | v1162&v1163
	*(*uint16)(unsafe.Add(mBase, uint32(v1155))) = uint16(v1165)
	v1168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1158))))
	v1171 = v235 + v1168&v1163
	if base.Ui32(v1163) <= base.Ui32(v1171) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	goto L206
L208:
	;
	v1174 = v1163
	goto L210
L209:
	;
	v1174 = v1171
	goto L210
L210:
	;
	v1175 = v1161 | v1174
	*(*uint16)(unsafe.Add(mBase, uint32(v1155))) = uint16(v1175)
	v1177 = int32(1)
	v1178 = v1134 + v1177
	*(*uint16)(unsafe.Add(mBase, uint32(v1105))) = uint16(v1178)
	v1181 = v1133 + v1177
	if v1103 == v1181 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1189 = v1178
	goto L197
L212:
	;
	goto L213
L213:
	;
	if v1181 != v1131 {
		v1133 = v1181
		v1134 = v1178
		v1137 = v1178
		goto L202
	} else {
		goto L214
	}
L214:
	;
	goto L203
L215:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v1200 | int32(1)
	v1204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105))))
	v1207 = v1204
	goto L195
L216:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v1223 & int32(-2)
	v1236 = v795
	goto L143
L217:
	;
	goto L218
L218:
	;
	v1227 = int32(1)
	v1236 = (v795+v1227)&int32(-2) + v1220<<(uint(v1227)%32) + int32(2)
	goto L143
L219:
	;
	if v1264 != 0 {
		v289 = v1269
		v290 = v1257
		v292 = v1259
		v293 = v1260
		v297 = v1263
		v298 = v1264
		goto L43
	} else {
		goto L220
	}
L220:
	;
	goto L44
L221:
	;
	v1297 = v1273
	v1298 = v1274
	v1300 = v1276
	v1305 = v1281
	goto L224
L222:
	;
	v1428 = v1273
	v1429 = v1274
	goto L223
L223:
	;
	if v1282 != 0 {
		goto L244
	} else {
		goto L245
	}
L224:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1297)))
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1300)))
	v1324 = int32(1)
	v1326 = v1320&int32(-2) | v1323&v1324
	*(*int32)(unsafe.Add(mBase, uint32(v1297))) = v1326
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1300)))
	*(*int32)(unsafe.Add(mBase, uint32(v1297))) = v1326&int32(-4095) | v1330&int32(4094)
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1300)))
	v1343 = int32(base.Ui32(v1336)>>(uint(v1324)%32)) & int32(2047)
	if v1343 != 0 {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	v1428 = v1424
	v1429 = v1416
	goto L223
L226:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1297)))
	*(*int32)(unsafe.Add(mBase, uint32(v1297))) = v1346&int32(4095) | v1298<<(uint(int32(12))%32)
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1300)))
	v1354 = int32(1)
	v1357 = int32(base.Ui32(v1353)>>(uint(v1354)%32)) & int32(2047)
	v1358 = v1357 + v1298
	if v1346&v1354 != 0 {
		goto L230
	} else {
		goto L231
	}
L227:
	;
	v1344 = F__emscripten_memcpy_bulkmem(m, v1298+v280, v283+int32(base.Ui32(v1336)>>(uint(int32(12))%32)), v1343)
	mBase = m.M
	goto L229
L228:
	;
	goto L229
L229:
	;
	goto L226
L230:
	;
	v1361 = int32(2)
	v1362 = int32(1)
	v1365 = (v1358 + v1362) & int32(-2)
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v1379 = v35 + v1368<<(uint(v1361)%32) + (int32(base.Ui32(v1353)>>(uint(int32(12))%32))+v1357+v1362)&int32(4194302)
	if v1353&v1362 != 0 {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	v1416 = v1358
	goto L232
L232:
	;
	v1421 = int32(4)
	v1424 = v1297 + v1421
	v1426 = v1305 - int32(1)
	if v1426 != 0 {
		v1297 = v1424
		v1298 = v1416
		v1300 = v1300 + v1421
		v1305 = v1426
		goto L224
	} else {
		goto L243
	}
L233:
	;
	v1382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1379))))
	v1387 = v1382<<(uint(int32(1))%32) + int32(2)
	goto L235
L234:
	;
	v1387 = v1361
	goto L235
L235:
	;
	if v1387 != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1300)))
	if v1390&int32(1) != 0 {
		goto L240
	} else {
		goto L241
	}
L237:
	;
	v1388 = F__emscripten_memcpy_bulkmem(m, v280+v1365, v1379, v1387)
	mBase = m.M
	goto L239
L238:
	;
	goto L239
L239:
	;
	goto L236
L240:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v1394 = int32(2)
	v1397 = int32(1)
	v1409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+v1393<<(uint(v1394)%32)+(int32(base.Ui32(v1390)>>(uint(v1397)%32))&int32(2047)+int32(base.Ui32(v1390)>>(uint(int32(12))%32))+v1397)&int32(4194302)))))
	v1414 = v1409<<(uint(v1397)%32) + v1394
	goto L242
L241:
	;
	v1414 = v1361
	goto L242
L242:
	;
	v1416 = v1414 + v1365
	goto L232
L243:
	;
	goto L225
L244:
	;
	v1452 = v1428
	v1453 = v1429
	v1456 = v1277
	v1461 = v1282
	goto L247
L245:
	;
	v1722 = v1428
	v1723 = v1429
	goto L246
L246:
	;
	if v1723 < int32(1048576) {
		goto L289
	} else {
		goto L290
	}
L247:
	;
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1452)))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1456)))
	v1479 = int32(1)
	v1481 = v1475&int32(-2) | v1478&v1479
	*(*int32)(unsafe.Add(mBase, uint32(v1452))) = v1481
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1456)))
	*(*int32)(unsafe.Add(mBase, uint32(v1452))) = v1481&int32(-4095) | v1485&int32(4094)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1456)))
	v1498 = int32(base.Ui32(v1491)>>(uint(v1479)%32)) & int32(2047)
	if v1498 != 0 {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	v1722 = v1718
	v1723 = v1713
	goto L246
L249:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1452)))
	*(*int32)(unsafe.Add(mBase, uint32(v1452))) = v1501&int32(4095) | v1453<<(uint(int32(12))%32)
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1456)))
	v1509 = int32(1)
	v1513 = int32(base.Ui32(v1508)>>(uint(v1509)%32))&int32(2047) + v1453
	if v1501&v1509 == int32(0) {
		v1713 = v1513
		goto L253
	} else {
		goto L254
	}
L250:
	;
	v1499 = F__emscripten_memcpy_bulkmem(m, v1453+v280, v275+int32(base.Ui32(v1491)>>(uint(int32(12))%32)), v1498)
	mBase = m.M
	goto L252
L251:
	;
	goto L252
L252:
	;
	goto L249
L253:
	;
	v1715 = int32(4)
	v1718 = v1452 + v1715
	v1720 = v1461 - int32(1)
	if v1720 != 0 {
		v1452 = v1718
		v1453 = v1713
		v1456 = v1456 + v1715
		v1461 = v1720
		goto L247
	} else {
		goto L288
	}
L254:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1452)))
	v1533 = int32(1)
	v1542 = v262 + v1526<<(uint(int32(2))%32) + (int32(base.Ui32(v1530)>>(uint(int32(12))%32))+int32(base.Ui32(v1530)>>(uint(v1533)%32))&int32(2047)+v1533)&int32(4194302)
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1456)))
	if v1543&v1533 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L255:
	;
	if v1697 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L256:
	;
	v1582 = v1542 + int32(8)
	if v1530&int32(1) != 0 {
		goto L261
	} else {
		goto L262
	}
L257:
	;
	v1548 = int32(1)
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v1578 = (int32(base.Ui32(v1543)>>(uint(v1548)%32))&int32(2047) + int32(base.Ui32(v1543)>>(uint(int32(12))%32)) + v1548) & int32(4194302)
	v1579 = v1559
	v1580 = int32(0)
	goto L256
L258:
	;
	goto L259
L259:
	;
	v1561 = int32(1)
	v1571 = (int32(base.Ui32(v1543)>>(uint(v1561)%32))&int32(2047) + int32(base.Ui32(v1543)>>(uint(int32(12))%32)) + v1561) & int32(4194302)
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v1577 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1571+(v37+v1572<<(uint(int32(2))%32)))+8)))
	v1578 = v1571
	v1579 = v1572
	v1580 = v1577
	goto L256
L260:
	;
	if v1580 == int32(0) {
		v1684 = v1589
		goto L264
	} else {
		goto L265
	}
L261:
	;
	v1585 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1582))))
	v1589 = v1585
	goto L260
L262:
	;
	goto L263
L263:
	;
	v1586 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1582))) = uint16(v1586)
	v1589 = v1586
	goto L260
L264:
	;
	v1697 = v1684&int32(65535) - v1589
	goto L255
L265:
	;
	if base.Ui32(int32(255)) < base.Ui32(v1589) {
		v1666 = v1589
		goto L266
	} else {
		goto L267
	}
L266:
	;
	if v1589 == v1666&int32(65535) {
		v1684 = v1589
		goto L264
	} else {
		goto L284
	}
L267:
	;
	v1942 = int32(10)
	v1602 = int32(0)
	v1603 = int32(256)
	v1604 = v1603 - v1589
	if base.Ui32(v1604) <= base.Ui32(v1603) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v1608 = v1604
	goto L270
L269:
	;
	v1608 = v1602
	goto L270
L270:
	;
	v1610 = v1602
	v1611 = v1589
	v1614 = v1589
	goto L271
L271:
	;
	if v1611 != 0 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	v1666 = v1655
	goto L266
L273:
	;
	v1625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1582+v1611<<(uint(int32(1))%32)))))
	v1626 = int32(16383)
	if v1625&v1626 == v1626 {
		v1666 = v1614
		goto L266
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1630 = int32(1)
	v1632 = v1542 + v1942 + v1611<<(uint(v1630)%32)
	v1635 = v37 + v1579<<(uint(int32(2))%32) + v1578 + v1942 + v1610<<(uint(v1630)%32)
	v1636 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1635))))
	v1638 = v1636 & int32(-16384)
	v1639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1632))))
	v1640 = int32(16383)
	v1642 = v1638 | v1639&v1640
	*(*uint16)(unsafe.Add(mBase, uint32(v1632))) = uint16(v1642)
	v1645 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1635))))
	v1648 = v235 + v1645&v1640
	if base.Ui32(v1640) <= base.Ui32(v1648) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	goto L275
L277:
	;
	v1651 = v1640
	goto L279
L278:
	;
	v1651 = v1648
	goto L279
L279:
	;
	v1652 = v1638 | v1651
	*(*uint16)(unsafe.Add(mBase, uint32(v1632))) = uint16(v1652)
	v1654 = int32(1)
	v1655 = v1611 + v1654
	*(*uint16)(unsafe.Add(mBase, uint32(v1582))) = uint16(v1655)
	v1658 = v1610 + v1654
	if v1580 == v1658 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1666 = v1655
	goto L266
L281:
	;
	goto L282
L282:
	;
	if v1658 != v1608 {
		v1610 = v1658
		v1611 = v1655
		v1614 = v1655
		goto L271
	} else {
		goto L283
	}
L283:
	;
	goto L272
L284:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1452)))
	*(*int32)(unsafe.Add(mBase, uint32(v1452))) = v1677 | int32(1)
	v1681 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1582))))
	v1684 = v1681
	goto L264
L285:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1452)))
	*(*int32)(unsafe.Add(mBase, uint32(v1452))) = v1700 & int32(-2)
	v1713 = v1513
	goto L253
L286:
	;
	goto L287
L287:
	;
	v1704 = int32(1)
	v1713 = (v1513+v1704)&int32(-2) + v1697<<(uint(v1704)%32) + int32(2)
	goto L253
L288:
	;
	goto L248
L289:
	;
	v1748 = int32(2)
	v1749 = (v1722 - v277) >> (uint(v1748) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v1749
	if v1749 != v269 {
		goto L292
	} else {
		goto L293
	}
L290:
	;
	goto L291
L291:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L1
	} else {
		goto L349
	}
L292:
	;
	v1756 = v277 + v1749<<(uint(int32(2))%32)
	if v1756 == v280 {
		goto L296
	} else {
		goto L297
	}
L293:
	;
	v1902 = v269
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v1723<<(uint(v1748)%32) + v1902<<(uint(int32(4))%32) + int32(32)
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1909 != v30 {
		goto L341
	} else {
		goto L342
	}
L295:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v1902 = v1901
	goto L294
L296:
	;
	goto L295
L297:
	;
	v1760 = v1756 + v1723
	if base.Ui32(v280-v1760) <= base.Ui32(int32(0)-v1723<<(uint(int32(1))%32)) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1767 = F___memcpy(m, v1756, v280, v1723)
	mBase = m.M
	goto L295
L299:
	;
	goto L300
L300:
	;
	v1770 = (v1756 ^ v280) & int32(3)
	if base.Ui32(v1756) < base.Ui32(v280) {
		goto L303
	} else {
		goto L304
	}
L301:
	;
	if v1872 == int32(0) {
		goto L296
	} else {
		goto L337
	}
L302:
	;
	if base.Ui32(v1850) <= base.Ui32(int32(3)) {
		v1871 = v1849
		v1872 = v1850
		v1873 = v1851
		goto L301
	} else {
		goto L333
	}
L303:
	;
	if v1770 != 0 {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	goto L305
L305:
	;
	if v1770 != 0 {
		v1832 = v1723
		goto L316
	} else {
		goto L317
	}
L306:
	;
	v1871 = v280
	v1872 = v1723
	v1873 = v1756
	goto L301
L307:
	;
	goto L308
L308:
	;
	if v1756&int32(3) == int32(0) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1849 = v280
	v1850 = v1723
	v1851 = v1756
	goto L302
L310:
	;
	goto L311
L311:
	;
	v1777 = v280
	v1778 = v1723
	v1779 = v1756
	goto L312
L312:
	;
	if v1778 == int32(0) {
		goto L296
	} else {
		goto L314
	}
L313:
	;
	v1849 = v1786
	v1850 = v1788
	v1851 = v1790
	goto L302
L314:
	;
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1777))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1779))) = uint8(v1783)
	v1785 = int32(1)
	v1786 = v1777 + v1785
	v1788 = v1778 - v1785
	v1790 = v1779 + v1785
	if v1790&int32(3) != 0 {
		v1777 = v1786
		v1778 = v1788
		v1779 = v1790
		goto L312
	} else {
		goto L315
	}
L315:
	;
	goto L313
L316:
	;
	if v1832 == int32(0) {
		goto L296
	} else {
		goto L329
	}
L317:
	;
	if v1760&int32(3) != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1797 = v1723
	goto L321
L319:
	;
	v1812 = v1723
	goto L320
L320:
	;
	if base.Ui32(v1812) <= base.Ui32(int32(3)) {
		v1832 = v1812
		goto L316
	} else {
		goto L325
	}
L321:
	;
	if v1797 == int32(0) {
		goto L296
	} else {
		goto L323
	}
L322:
	;
	v1812 = v1803
	goto L320
L323:
	;
	v1803 = v1797 - int32(1)
	v1804 = v1756 + v1803
	v1806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v1803))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1804))) = uint8(v1806)
	if v1804&int32(3) != 0 {
		v1797 = v1803
		goto L321
	} else {
		goto L324
	}
L324:
	;
	goto L322
L325:
	;
	v1819 = v1812
	goto L326
L326:
	;
	v1823 = v1819 - int32(4)
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v280+v1823)))
	*(*int32)(unsafe.Add(mBase, uint32(v1756+v1823))) = v1826
	if base.Ui32(int32(3)) < base.Ui32(v1823) {
		v1819 = v1823
		goto L326
	} else {
		goto L328
	}
L327:
	;
	v1832 = v1823
	goto L316
L328:
	;
	goto L327
L329:
	;
	v1839 = v1832
	goto L330
L330:
	;
	v1843 = v1839 - int32(1)
	v1846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v1843))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1756+v1843))) = uint8(v1846)
	if v1843 != 0 {
		v1839 = v1843
		goto L330
	} else {
		goto L332
	}
L331:
	;
	goto L296
L332:
	;
	goto L331
L333:
	;
	v1856 = v1849
	v1857 = v1850
	v1858 = v1851
	goto L334
L334:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1856)))
	*(*int32)(unsafe.Add(mBase, uint32(v1858))) = v1860
	v1862 = int32(4)
	v1863 = v1856 + v1862
	v1865 = v1858 + v1862
	v1867 = v1857 - v1862
	if base.Ui32(int32(3)) < base.Ui32(v1867) {
		v1856 = v1863
		v1857 = v1867
		v1858 = v1865
		goto L334
	} else {
		goto L336
	}
L335:
	;
	v1871 = v1863
	v1872 = v1867
	v1873 = v1865
	goto L301
L336:
	;
	goto L335
L337:
	;
	v1878 = v1871
	v1879 = v1872
	v1880 = v1873
	goto L338
L338:
	;
	v1882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1878))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1880))) = uint8(v1882)
	v1884 = int32(1)
	v1889 = v1879 - v1884
	if v1889 != 0 {
		v1878 = v1878 + v1884
		v1879 = v1889
		v1880 = v1880 + v1884
		goto L338
	} else {
		goto L340
	}
L339:
	;
	goto L296
L340:
	;
	goto L339
L341:
	;
	F_pfree(m, v30)
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L1
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1913 != v37 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	goto L343
L345:
	;
	F_pfree(m, v37)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L1
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	m.G0 = v27 + int32(16)
	return v262
L348:
	;
	goto L347
L349:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(1048575)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v1723
	F_errmsg(m, int32(658192), v27)
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	F_errfinish(m, int32(495198), int32(1126), int32(112288))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsvector_eq(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = l0 + int32(28)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v29 = int32(2)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v33 = int32(base.Ui32(v31) >> (uint(v29) % 32))
	if base.Ui32(v30) < base.Ui32(v33) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v224 != v7 {
		goto L67
	} else {
		goto L68
	}
L5:
	;
	v223 = int32(-1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v36 = int32(1)
	if base.Ui32(v33) < base.Ui32(v30) {
		v199 = v36
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v223 = v199
	goto L4
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v38 < v39 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v223 = int32(-1)
	goto L4
L11:
	;
	goto L12
L12:
	;
	if v39 < v38 {
		v199 = v36
		goto L8
	} else {
		goto L13
	}
L13:
	;
	if v38 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v223 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	v46 = int32(8)
	v47 = v14 + v46
	v48 = int32(2)
	v50 = v47 + v39<<(uint(v48)%32)
	v52 = v7 + v46
	v55 = v52 + v38<<(uint(v48)%32)
	v63 = v47
	v64 = v52
	v68 = int32(0)
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v71 = int32(1)
	v72 = v70 & v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v75 = v73 & v71
	if v72 != v75 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v199 = int32(0)
	goto L8
L19:
	;
	if base.Ui32(v75) < base.Ui32(v72) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v81 = int32(1)
	v83 = int32(2047)
	v84 = int32(base.Ui32(v73)>>(uint(v81)%32)) & v83
	v85 = int32(12)
	v86 = int32(base.Ui32(v73) >> (uint(v85) % 32))
	v88 = int32(base.Ui32(v70) >> (uint(v85) % 32))
	v92 = int32(base.Ui32(v70)>>(uint(v81)%32)) & v83
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v80 = int32(-1)
	goto L24
L23:
	;
	v80 = int32(1)
	goto L24
L24:
	;
	v223 = v80
	goto L4
L25:
	;
	if v72 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	if v84 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v84 == int32(0) {
		goto L25
	} else {
		goto L40
	}
L29:
	;
	v223 = int32(1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v98 = base.B2i32(base.Ui32(v92) < base.Ui32(v84))
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v92
	goto L34
L33:
	;
	v99 = v84
	goto L34
L34:
	;
	v100 = F_memcmp(m, v88+v55, v86+v50, v99)
	mBase = m.M
	if v100 != 0 {
		v199 = v100
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v92 == v84 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	if base.Ui32(v92) < base.Ui32(v84) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = int32(-1)
	goto L39
L38:
	;
	v104 = int32(1)
	goto L39
L39:
	;
	v223 = v104
	goto L4
L40:
	;
	v223 = int32(-1)
	goto L4
L41:
	;
	v187 = int32(4)
	v193 = v68 + int32(1)
	if v193 != v38 {
		v63 = v63 + v187
		v64 = v64 + v187
		v68 = v193
		goto L17
	} else {
		goto L66
	}
L42:
	;
	v113 = int32(1)
	v115 = int32(4194302)
	v117 = v50 + (v84+v86+v113)&v115
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117))))
	v124 = v55 + (v92+v88+v113)&v115
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v118 == v125 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v132 = v124
	v133 = v117
	v134 = int32(0)
	goto L51
L44:
	;
	if v125 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v118) < base.Ui32(v125) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L41
L48:
	;
	v131 = int32(-1)
	goto L50
L49:
	;
	v131 = int32(1)
	goto L50
L50:
	;
	v223 = v131
	goto L4
L51:
	;
	v146 = int32(2)
	v147 = v132 + v146
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147))))
	v149 = int32(16383)
	v150 = v148 & v149
	v152 = v133 + v146
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v155 = v153 & v149
	if v150 != v155 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(v164) < base.Ui32(v162) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	if base.Ui32(v155) < base.Ui32(v150) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v161 = int32(14)
	v162 = int32(base.Ui32(v148) >> (uint(v161) % 32))
	v164 = int32(base.Ui32(v153) >> (uint(v161) % 32))
	if v162 == v164 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v160 = int32(-1)
	goto L58
L57:
	;
	v160 = int32(1)
	goto L58
L58:
	;
	v223 = v160
	goto L4
L59:
	;
	v167 = v134 + int32(1)
	if v167 == v125 {
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
	v132 = v147
	v133 = v152
	v134 = v167
	goto L51
L63:
	;
	v172 = int32(-1)
	goto L65
L64:
	;
	v172 = int32(1)
	goto L65
L65:
	;
	v199 = v172
	goto L8
L66:
	;
	goto L18
L67:
	;
	F_pfree(m, v7)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v228 != v14 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	F_pfree(m, v14)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return base.B2i32(v223 == int32(0))
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v199 int32
	_ = v199
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
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
			v46 = int32(49152)
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
					v55 = F__emscripten_memcpy_bulkmem(m, v50, v19, v54)
					mBase = m.M
					v56 = v55
				} else {
					v56 = v50
				}
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
				if v57 != 0 {
					v59 = v56 + int32(8)
					v68 = v59
					v70 = v57
					for {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
						if v73&int32(1) == int32(0) {
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
							v82 = int32(1)
							v93 = v59 + v78<<(uint(int32(2))%32) + (int32(base.Ui32(v73)>>(uint(v82)%32))&int32(2047)+int32(base.Ui32(v73)>>(uint(int32(12))%32))+v82)&int32(4194302)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
							if v94 == int32(0) {
							} else {
								v99 = v94 & int32(3)
								if v99 != 0 {
									v101 = v93
									v103 = int32(0)
									v105 = v94
									for {
										v114 = v101 + int32(2)
										v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
										v118 = v115&int32(16383) | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v114))) = uint16(v118)
										v120 = int32(1)
										v121 = v105 - v120
										v123 = v103 + v120
										if v123 != v99 {
											v101 = v114
											v103 = v123
											v105 = v121
											continue
										} else {
											break
										}
										break
									}
									v126 = v114
									v130 = v121
								} else {
									v126 = v93
									v130 = v94
								}
								if base.Ui32(v94) < base.Ui32(int32(4)) {
								} else {
									v141 = v126
									v145 = v130
									for {
										v154 = v141 + int32(2)
										v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154))))
										v156 = int32(16383)
										v158 = v155&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v154))) = uint16(v158)
										v160 = int32(4)
										v161 = v141 + v160
										v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161))))
										v165 = v162&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v161))) = uint16(v165)
										v168 = v141 + int32(6)
										v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168))))
										v172 = v169&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v168))) = uint16(v172)
										v175 = v141 + int32(8)
										v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175))))
										v179 = v176&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v175))) = uint16(v179)
										v182 = v145 - v160
										if v182 != 0 {
											v141 = v175
											v145 = v182
											continue
										} else {
											break
										}
										break
									}
								}
							}
						}
						v199 = v70 - int32(1)
						if v199 != 0 {
							v68 = v68 + int32(4)
							v70 = v199
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v213 != v19 {
					F_pfree(m, v19)
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(16)
						return v56
					}
				} else {
					m.G0 = v16 + int32(16)
					return v56
				}
			}
		case 1, 33:
			v46 = int32(32768)
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
					v55 = F__emscripten_memcpy_bulkmem(m, v50, v19, v54)
					mBase = m.M
					v56 = v55
				} else {
					v56 = v50
				}
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
				if v57 != 0 {
					v59 = v56 + int32(8)
					v68 = v59
					v70 = v57
					for {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
						if v73&int32(1) == int32(0) {
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
							v82 = int32(1)
							v93 = v59 + v78<<(uint(int32(2))%32) + (int32(base.Ui32(v73)>>(uint(v82)%32))&int32(2047)+int32(base.Ui32(v73)>>(uint(int32(12))%32))+v82)&int32(4194302)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
							if v94 == int32(0) {
							} else {
								v99 = v94 & int32(3)
								if v99 != 0 {
									v101 = v93
									v103 = int32(0)
									v105 = v94
									for {
										v114 = v101 + int32(2)
										v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
										v118 = v115&int32(16383) | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v114))) = uint16(v118)
										v120 = int32(1)
										v121 = v105 - v120
										v123 = v103 + v120
										if v123 != v99 {
											v101 = v114
											v103 = v123
											v105 = v121
											continue
										} else {
											break
										}
										break
									}
									v126 = v114
									v130 = v121
								} else {
									v126 = v93
									v130 = v94
								}
								if base.Ui32(v94) < base.Ui32(int32(4)) {
								} else {
									v141 = v126
									v145 = v130
									for {
										v154 = v141 + int32(2)
										v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154))))
										v156 = int32(16383)
										v158 = v155&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v154))) = uint16(v158)
										v160 = int32(4)
										v161 = v141 + v160
										v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161))))
										v165 = v162&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v161))) = uint16(v165)
										v168 = v141 + int32(6)
										v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168))))
										v172 = v169&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v168))) = uint16(v172)
										v175 = v141 + int32(8)
										v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175))))
										v179 = v176&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v175))) = uint16(v179)
										v182 = v145 - v160
										if v182 != 0 {
											v141 = v175
											v145 = v182
											continue
										} else {
											break
										}
										break
									}
								}
							}
						}
						v199 = v70 - int32(1)
						if v199 != 0 {
							v68 = v68 + int32(4)
							v70 = v199
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v213 != v19 {
					F_pfree(m, v19)
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(16)
						return v56
					}
				} else {
					m.G0 = v16 + int32(16)
					return v56
				}
			}
		case 2, 34:
			v46 = int32(16384)
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
					v55 = F__emscripten_memcpy_bulkmem(m, v50, v19, v54)
					mBase = m.M
					v56 = v55
				} else {
					v56 = v50
				}
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
				if v57 != 0 {
					v59 = v56 + int32(8)
					v68 = v59
					v70 = v57
					for {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
						if v73&int32(1) == int32(0) {
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
							v82 = int32(1)
							v93 = v59 + v78<<(uint(int32(2))%32) + (int32(base.Ui32(v73)>>(uint(v82)%32))&int32(2047)+int32(base.Ui32(v73)>>(uint(int32(12))%32))+v82)&int32(4194302)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
							if v94 == int32(0) {
							} else {
								v99 = v94 & int32(3)
								if v99 != 0 {
									v101 = v93
									v103 = int32(0)
									v105 = v94
									for {
										v114 = v101 + int32(2)
										v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
										v118 = v115&int32(16383) | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v114))) = uint16(v118)
										v120 = int32(1)
										v121 = v105 - v120
										v123 = v103 + v120
										if v123 != v99 {
											v101 = v114
											v103 = v123
											v105 = v121
											continue
										} else {
											break
										}
										break
									}
									v126 = v114
									v130 = v121
								} else {
									v126 = v93
									v130 = v94
								}
								if base.Ui32(v94) < base.Ui32(int32(4)) {
								} else {
									v141 = v126
									v145 = v130
									for {
										v154 = v141 + int32(2)
										v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154))))
										v156 = int32(16383)
										v158 = v155&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v154))) = uint16(v158)
										v160 = int32(4)
										v161 = v141 + v160
										v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161))))
										v165 = v162&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v161))) = uint16(v165)
										v168 = v141 + int32(6)
										v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168))))
										v172 = v169&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v168))) = uint16(v172)
										v175 = v141 + int32(8)
										v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175))))
										v179 = v176&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v175))) = uint16(v179)
										v182 = v145 - v160
										if v182 != 0 {
											v141 = v175
											v145 = v182
											continue
										} else {
											break
										}
										break
									}
								}
							}
						}
						v199 = v70 - int32(1)
						if v199 != 0 {
							v68 = v68 + int32(4)
							v70 = v199
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v213 != v19 {
					F_pfree(m, v19)
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(16)
						return v56
					}
				} else {
					m.G0 = v16 + int32(16)
					return v56
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
					v55 = F__emscripten_memcpy_bulkmem(m, v50, v19, v54)
					mBase = m.M
					v56 = v55
				} else {
					v56 = v50
				}
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
				if v57 != 0 {
					v59 = v56 + int32(8)
					v68 = v59
					v70 = v57
					for {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
						if v73&int32(1) == int32(0) {
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
							v82 = int32(1)
							v93 = v59 + v78<<(uint(int32(2))%32) + (int32(base.Ui32(v73)>>(uint(v82)%32))&int32(2047)+int32(base.Ui32(v73)>>(uint(int32(12))%32))+v82)&int32(4194302)
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93))))
							if v94 == int32(0) {
							} else {
								v99 = v94 & int32(3)
								if v99 != 0 {
									v101 = v93
									v103 = int32(0)
									v105 = v94
									for {
										v114 = v101 + int32(2)
										v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
										v118 = v115&int32(16383) | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v114))) = uint16(v118)
										v120 = int32(1)
										v121 = v105 - v120
										v123 = v103 + v120
										if v123 != v99 {
											v101 = v114
											v103 = v123
											v105 = v121
											continue
										} else {
											break
										}
										break
									}
									v126 = v114
									v130 = v121
								} else {
									v126 = v93
									v130 = v94
								}
								if base.Ui32(v94) < base.Ui32(int32(4)) {
								} else {
									v141 = v126
									v145 = v130
									for {
										v154 = v141 + int32(2)
										v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154))))
										v156 = int32(16383)
										v158 = v155&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v154))) = uint16(v158)
										v160 = int32(4)
										v161 = v141 + v160
										v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161))))
										v165 = v162&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v161))) = uint16(v165)
										v168 = v141 + int32(6)
										v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168))))
										v172 = v169&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v168))) = uint16(v172)
										v175 = v141 + int32(8)
										v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175))))
										v179 = v176&v156 | v46
										*(*uint16)(unsafe.Add(mBase, uint32(v175))) = uint16(v179)
										v182 = v145 - v160
										if v182 != 0 {
											v141 = v175
											v145 = v182
											continue
										} else {
											break
										}
										break
									}
								}
							}
						}
						v199 = v70 - int32(1)
						if v199 != 0 {
							v68 = v68 + int32(4)
							v70 = v199
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v213 != v19 {
					F_pfree(m, v19)
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return int32(0)
					} else {
						m.G0 = v16 + int32(16)
						return v56
					}
				} else {
					m.G0 = v16 + int32(16)
					return v56
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
				F_errmsg_internal(m, int32(480971), v16)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495198), int32(242), int32(104506))
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
