package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DecodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int64
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v523 int32
	_ = v523
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int64
	_ = v765
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v849 float64
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 float64
	_ = v855
	var v857 float64
	_ = v857
	var v861 int64
	_ = v861
	var v863 int64
	_ = v863
	var v866 int64
	_ = v866
	var v871 int64
	_ = v871
	var v873 int64
	_ = v873
	var v878 int64
	_ = v878
	var v880 int64
	_ = v880
	var v884 int64
	_ = v884
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1040 int32
	_ = v1040
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1157 int32
	_ = v1157
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1200 int32
	_ = v1200
	var v1220 int32
	_ = v1220
	var v1236 int32
	_ = v1236
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1272 int32
	_ = v1272
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1378 int32
	_ = v1378
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1504 int32
	_ = v1504
	var v1512 int32
	_ = v1512
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1535 int32
	_ = v1535
	var v1543 int32
	_ = v1543
	var v1552 int32
	_ = v1552
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1582 int32
	_ = v1582
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1633 int32
	_ = v1633
	var v1659 int32
	_ = v1659
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1790 int32
	_ = v1790
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1826 int32
	_ = v1826
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1842 int32
	_ = v1842
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1857 int32
	_ = v1857
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1882 int32
	_ = v1882
	var v1892 int32
	_ = v1892
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1923 int32
	_ = v1923
	var v1933 int32
	_ = v1933
	var v1936 int32
	_ = v1936
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1956 int32
	_ = v1956
	var v1960 int32
	_ = v1960
	var v1966 int32
	_ = v1966
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1980 int32
	_ = v1980
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1989 int32
	_ = v1989
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2020 int32
	_ = v2020
	var v2028 int32
	_ = v2028
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	v9 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(112)
	m.G0 = v38
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+59)) = uint8(v9)
	v42 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v42
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
	v56 = l4 + int32(8)
	if int32(0) < l2 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v2049 + int32(112)
	return v2048
L5:
	;
	v62 = l4 + int32(20)
	v64 = l4 + int32(12)
	v66 = l4 + int32(16)
	v78 = v9
	v79 = v9
	v86 = v42
	v88 = v9
	v89 = v9
	v93 = v9
	v94 = v9
	v95 = v9
	v98 = v9
	v99 = v9
	goto L8
L6:
	;
	v1735 = v38
	v1738 = v9
	v1745 = v42
	v1748 = v9
	v1752 = v9
	v1754 = v9
	v1757 = v9
	v1758 = v9
	goto L7
L7:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1761 != int32(2) {
		goto L364
	} else {
		goto L365
	}
L8:
	;
	v102 = int32(-1)
	v104 = v88 << (uint(int32(2)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1+v104)))
	switch v106 {
	case 0:
		goto L14
	case 1, 6:
		goto L13
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	default:
		v2048 = v102
		v2049 = v38
		goto L4
	}
L9:
	;
	if v1698 != 0 {
		v2048 = int32(-1)
		v2049 = v38
		goto L4
	} else {
		goto L363
	}
L10:
	;
	v1723 = v88 + int32(1)
	if v1723 != l2 {
		v78 = v1698
		v79 = v1699
		v86 = v1706
		v88 = v1723
		v89 = v1709
		v93 = v1713
		v94 = v1714
		v95 = v1715
		v98 = v1718
		v99 = v1719
		goto L8
	} else {
		goto L362
	}
L11:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v1683&v79 != 0 {
		goto L359
	} else {
		goto L360
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(32)
	v1659 = v78
	v1667 = v86
	v1670 = v1633
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L13:
	;
	v976 = l0 + v104
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)))
	v984 = F_DecodeTimezoneAbbrev(m, v88, v977, v38-int32(-64), v38+int32(60), v38+int32(52), l7)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L105
	} else {
		goto L244
	}
L14:
	;
	if v78 != 0 {
		goto L159
	} else {
		goto L160
	}
L15:
	;
	if l6 == int32(0) {
		v2048 = v102
		v2049 = v38
		goto L4
	} else {
		goto L131
	}
L16:
	;
	switch v78 {
	case 0, 3:
		goto L116
	default:
		v2048 = v102
		v2049 = v38
		goto L4
	}
L17:
	;
	if v78 == int32(31) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if l6 == int32(0) {
		v2048 = v102
		v2049 = v38
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v301 = int32(0)
	v303 = int32(10)
	if base.B2i32(v78 == v301)&base.B2i32(v79&v303 != v303) == v301 {
		goto L56
	} else {
		goto L57
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v119 = F_strtol(m, v115, v38+int32(72), int32(10))
	mBase = m.M
	goto L22
L22:
	;
	v120 = int32(-2)
	v122 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v122 == int32(68) {
		v2048 = v120
		v2049 = v38
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v119 < int32(0) {
		v2048 = v120
		v2049 = v38
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v128 = v119 + int32(32044)
	v129 = int32(146097)
	v130 = base.I32_div_u_s(v128, v129)
	v131 = int32(3)
	v137 = int32(2)
	v142 = base.I32_div_u_s((v130*int32(1073595727)+v128)<<(uint(v137)%32)|v131, v129)
	v145 = v119 + v130*v131 + v142 + int32(32104)
	v146 = int32(1461)
	v147 = base.I32_div_u_s(v145, v146)
	v150 = v147*int32(-1461) + v145
	v152 = v150 << (uint(v137) % 32)
	if base.Ui32(v146) <= base.Ui32(v152) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v165 = base.I32_div_u_s(v152, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v165 + v147<<(uint(int32(2))%32) - int32(4800)
	v173 = v163 + int32(123)
	v176 = int32(16)
	v177 = int32(base.Ui32(v173*int32(2141)) >> (uint(v176) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v173 - int32(base.Ui32(v177*int32(7834))>>(uint(int32(8))%32))
	v187 = base.I32_rem_u_s(v177+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v187 + int32(1)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v192 = int32(0)
	v198 = m.G0
	v200 = v198 - v176
	m.G0 = v200
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	switch v203 - int32(43) {
	case 0, 2:
		goto L31
	default:
		v290 = int32(-1)
		goto L30
	}
L26:
	;
	v158 = base.I32_rem_u_s(v150+int32(305), int32(365))
	v163 = v158
	goto L25
L27:
	;
	goto L28
L28:
	;
	v162 = base.I32_rem_u_s(v150+int32(306), int32(366))
	v163 = v162
	goto L25
L29:
	;
	if v290 != 0 {
		v2048 = v290
		v2049 = v38
		goto L4
	} else {
		goto L55
	}
L30:
	;
	m.G0 = v200 + int32(16)
	goto L29
L31:
	;
	v206 = int32(4654024)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v213 = F_strtoint(m, v191+int32(1), v200+int32(12))
	mBase = m.M
	v214 = int32(-5)
	v216 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v216 == int32(68) {
		v290 = v214
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v220 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if base.Ui32(int32(15)) < base.Ui32(v262) {
		v290 = v214
		goto L30
	} else {
		goto L46
	}
L34:
	;
	if v220 != int32(58) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v252 = F_strlen(m, v191)
	mBase = m.M
	if base.Ui32(v252) < base.Ui32(int32(4)) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v261 = int32(0)
	v262 = v213
	v264 = v192
	goto L33
L38:
	;
	goto L39
L39:
	;
	v224 = int32(4654024)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v231 = F_strtoint(m, v219+int32(1), v200+int32(12))
	mBase = m.M
	v233 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v233 == int32(68) {
		v290 = v214
		goto L30
	} else {
		goto L40
	}
L40:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
	if v237 != int32(58) {
		v261 = v231
		v262 = v213
		v264 = v192
		goto L33
	} else {
		goto L41
	}
L41:
	;
	v240 = int32(4654024)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v247 = F_strtoint(m, v236+int32(1), v200+int32(12))
	mBase = m.M
	v249 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v249 != int32(68) {
		v261 = v231
		v262 = v213
		v264 = v247
		goto L33
	} else {
		goto L42
	}
L42:
	;
	v290 = v214
	goto L30
L43:
	;
	v261 = int32(0)
	v262 = v213
	v264 = v192
	goto L33
L44:
	;
	goto L45
L45:
	;
	v256 = int32(100)
	v257 = base.I32_div_s(v213, v256)
	v261 = v213 - v257*v256
	v262 = v257
	v264 = v192
	goto L33
L46:
	;
	if base.Ui32(int32(59)) < base.Ui32(v261) {
		v290 = v214
		goto L30
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(int32(59)) < base.Ui32(v264) {
		v290 = v214
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v271 = int32(60)
	v276 = (v262*v271+v261)*v271 + v264
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v279 == int32(45) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v282 = v276
	goto L51
L50:
	;
	v282 = int32(0) - v276
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v282
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v287 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v288 = int32(-1)
	goto L54
L53:
	;
	v288 = int32(0)
	goto L54
L54:
	;
	v290 = v288
	goto L30
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31790)
	v1659 = int32(0)
	v1667 = v86
	v1670 = v89
	v1674 = int32(1)
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L56:
	;
	if l6 == int32(0) {
		v2048 = v102
		v2049 = v38
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v486 = F_DecodeDate(m, v481, v79, v38+int32(68), v38+int32(59), l4)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L105
	} else {
		goto L114
	}
L59:
	;
	v312 = l0 + v104
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v78 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v470 = F_pg_tzset(m, v313)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L105
	} else {
		goto L110
	}
L61:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	if base.Ui32(int32(9)) < base.Ui32((v316-int32(48))&int32(255)) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if base.B2i32(v78 != int32(3))&base.B2i32(v78 != int32(0)) != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L66:
	;
	goto L67
L67:
	;
	v329 = int32(31744)
	if v79&v329 == v329 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L69:
	;
	goto L70
L70:
	;
	v334 = int32(45)
	v335 = F___strchrnul(m, v313, v334)
	mBase = m.M
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	if v337 == v334 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v341 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v341 = v335
	goto L74
L73:
	;
	v341 = int32(0)
	goto L74
L74:
	;
	goto L71
L75:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L76:
	;
	goto L77
L77:
	;
	v345 = int32(0)
	v351 = m.G0
	v353 = v351 - int32(16)
	m.G0 = v353
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	switch v356 - int32(43) {
	case 0, 2:
		goto L80
	default:
		v443 = int32(-1)
		goto L79
	}
L78:
	;
	if v443 != 0 {
		v2048 = v443
		v2049 = v38
		goto L4
	} else {
		goto L104
	}
L79:
	;
	m.G0 = v353 + int32(16)
	goto L78
L80:
	;
	v359 = int32(4654024)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v366 = F_strtoint(m, v341+int32(1), v353+int32(12))
	mBase = m.M
	v367 = int32(-5)
	v369 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v369 == int32(68) {
		v443 = v367
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	if v373 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if base.Ui32(int32(15)) < base.Ui32(v415) {
		v443 = v367
		goto L79
	} else {
		goto L95
	}
L83:
	;
	if v373 != int32(58) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v405 = F_strlen(m, v341)
	mBase = m.M
	if base.Ui32(v405) < base.Ui32(int32(4)) {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	v414 = int32(0)
	v415 = v366
	v417 = v345
	goto L82
L87:
	;
	goto L88
L88:
	;
	v377 = int32(4654024)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v384 = F_strtoint(m, v372+int32(1), v353+int32(12))
	mBase = m.M
	v386 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v386 == int32(68) {
		v443 = v367
		goto L79
	} else {
		goto L89
	}
L89:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	if v390 != int32(58) {
		v414 = v384
		v415 = v366
		v417 = v345
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v393 = int32(4654024)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v400 = F_strtoint(m, v389+int32(1), v353+int32(12))
	mBase = m.M
	v402 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v402 != int32(68) {
		v414 = v384
		v415 = v366
		v417 = v400
		goto L82
	} else {
		goto L91
	}
L91:
	;
	v443 = v367
	goto L79
L92:
	;
	v414 = int32(0)
	v415 = v366
	v417 = v345
	goto L82
L93:
	;
	goto L94
L94:
	;
	v409 = int32(100)
	v410 = base.I32_div_s(v366, v409)
	v414 = v366 - v410*v409
	v415 = v410
	v417 = v345
	goto L82
L95:
	;
	if base.Ui32(int32(59)) < base.Ui32(v414) {
		v443 = v367
		goto L79
	} else {
		goto L96
	}
L96:
	;
	if base.Ui32(int32(59)) < base.Ui32(v417) {
		v443 = v367
		goto L79
	} else {
		goto L97
	}
L97:
	;
	v424 = int32(60)
	v429 = (v415*v424+v414)*v424 + v417
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if v432 == int32(45) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v435 = v429
	goto L100
L99:
	;
	v435 = int32(0) - v429
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v435
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v353)+12))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439))))
	if v440 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v441 = int32(-1)
	goto L103
L102:
	;
	v441 = int32(0)
	goto L103
L103:
	;
	v443 = v441
	goto L79
L104:
	;
	v450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v341))) = uint8(v450)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	v453 = F_strlen(m, v452)
	mBase = m.M
	v458 = F_DecodeNumberField(m, v453, v452, v79, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	return int32(0)
L106:
	;
	if v458 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L108:
	;
	goto L109
L109:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v465 | int32(32)
	v1659 = int32(0)
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L110:
	;
	if v470 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v474
	v2048 = int32(-6)
	v2049 = v38
	goto L4
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(32)
	v1659 = int32(0)
	v1667 = v86
	v1670 = v470
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L114:
	;
	if v486 != 0 {
		v2048 = v486
		v2049 = v38
		goto L4
	} else {
		goto L115
	}
L115:
	;
	v1659 = int32(0)
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L116:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v496 = F_DecodeTimeCommon(m, v490, int32(32767), v38+int32(68), v38+int32(72))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L105
	} else {
		goto L117
	}
L117:
	;
	if v496 != 0 {
		v2048 = v496
		v2049 = v38
		goto L4
	} else {
		goto L118
	}
L118:
	;
	v498 = *(*int64)(unsafe.Add(mBase, uint32(v38)+88))
	if int64(2147483648) <= v498 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v2048 = int32(-2)
	v2049 = v38
	goto L4
L120:
	;
	goto L121
L121:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l4)+8)) = uint32(v498)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v503
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v505
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v507
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v513 = int32(1)
	if base.Ui32(int32(24)) < base.Ui32(v509) {
		v535 = v513
		goto L123
	} else {
		goto L124
	}
L122:
	;
	if v535 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L123:
	;
	goto L122
L124:
	;
	if base.Ui32(int32(59)) < base.Ui32(v510) {
		v535 = v513
		goto L123
	} else {
		goto L125
	}
L125:
	;
	if base.Ui32(int32(60)) < base.Ui32(v511) {
		v535 = v513
		goto L123
	} else {
		goto L126
	}
L126:
	;
	if base.Ui32(int32(1000000)) < base.Ui32(v507) {
		v535 = v513
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v523 = int32(60)
	v535 = base.B2i32(base.Ui64(int64(86400000000)) < base.Ui64(base.I64_extend_i32_u(v507)+base.I64_extend_i32_u((v509*v523+v510)*v523+v511)*int64(1000000)))
	goto L123
L128:
	;
	v1659 = int32(0)
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L129:
	;
	goto L130
L130:
	;
	v2048 = int32(-2)
	v2049 = v38
	goto L4
L131:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v546 = int32(0)
	v552 = m.G0
	v554 = v552 - int32(16)
	m.G0 = v554
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	switch v557 - int32(43) {
	case 0, 2:
		goto L134
	default:
		v644 = int32(-1)
		goto L133
	}
L132:
	;
	if v644 != 0 {
		v2048 = v644
		v2049 = v38
		goto L4
	} else {
		goto L158
	}
L133:
	;
	m.G0 = v554 + int32(16)
	goto L132
L134:
	;
	v560 = int32(4654024)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v567 = F_strtoint(m, v543+int32(1), v554+int32(12))
	mBase = m.M
	v568 = int32(-5)
	v570 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v570 == int32(68) {
		v644 = v568
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v554)+12))
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573))))
	if v574 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if base.Ui32(int32(15)) < base.Ui32(v616) {
		v644 = v568
		goto L133
	} else {
		goto L149
	}
L137:
	;
	if v574 != int32(58) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	v606 = F_strlen(m, v543)
	mBase = m.M
	if base.Ui32(v606) < base.Ui32(int32(4)) {
		goto L146
	} else {
		goto L147
	}
L140:
	;
	v615 = int32(0)
	v616 = v567
	v618 = v546
	goto L136
L141:
	;
	goto L142
L142:
	;
	v578 = int32(4654024)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v585 = F_strtoint(m, v573+int32(1), v554+int32(12))
	mBase = m.M
	v587 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v587 == int32(68) {
		v644 = v568
		goto L133
	} else {
		goto L143
	}
L143:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v554)+12))
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if v591 != int32(58) {
		v615 = v585
		v616 = v567
		v618 = v546
		goto L136
	} else {
		goto L144
	}
L144:
	;
	v594 = int32(4654024)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v601 = F_strtoint(m, v590+int32(1), v554+int32(12))
	mBase = m.M
	v603 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v603 != int32(68) {
		v615 = v585
		v616 = v567
		v618 = v601
		goto L136
	} else {
		goto L145
	}
L145:
	;
	v644 = v568
	goto L133
L146:
	;
	v615 = int32(0)
	v616 = v567
	v618 = v546
	goto L136
L147:
	;
	goto L148
L148:
	;
	v610 = int32(100)
	v611 = base.I32_div_s(v567, v610)
	v615 = v567 - v611*v610
	v616 = v611
	v618 = v546
	goto L136
L149:
	;
	if base.Ui32(int32(59)) < base.Ui32(v615) {
		v644 = v568
		goto L133
	} else {
		goto L150
	}
L150:
	;
	if base.Ui32(int32(59)) < base.Ui32(v618) {
		v644 = v568
		goto L133
	} else {
		goto L151
	}
L151:
	;
	v625 = int32(60)
	v630 = (v616*v625+v615)*v625 + v618
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543))))
	if v633 == int32(45) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v636 = v630
	goto L154
L153:
	;
	v636 = int32(0) - v630
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(72)))) = v636
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v554)+12))
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if v641 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v642 = int32(-1)
	goto L157
L156:
	;
	v642 = int32(0)
	goto L157
L157:
	;
	v644 = v642
	goto L133
L158:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v651
	v1633 = v89
	goto L12
L159:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v656 = l0 + v104
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v661 = F_strtol(m, v657, v38+int32(4), int32(10))
	mBase = m.M
	goto L162
L160:
	;
	goto L161
L161:
	;
	v916 = v79 & int32(14)
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0+v104)))
	v919 = F_strlen(m, v918)
	mBase = m.M
	v920 = int32(46)
	v921 = F___strchrnul(m, v918, v920)
	mBase = m.M
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921))))
	if v923 == v920 {
		goto L222
	} else {
		goto L223
	}
L162:
	;
	v662 = int32(-2)
	v664 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v664 == int32(68) {
		v2048 = v662
		v2049 = v38
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	if v668 == int32(46) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	if v78 != int32(3) {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	if v668 == int32(0) {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	v1659 = int32(0)
	v1667 = v86
	v1670 = v89
	v1674 = v910
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L168:
	;
	if v78 != int32(31) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L170
L170:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	v891 = F_strlen(m, v890)
	mBase = m.M
	v898 = F_DecodeNumberField(m, v891, v890, v79|int32(14), v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L105
	} else {
		goto L215
	}
L171:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L172:
	;
	goto L173
L173:
	;
	if v661 < int32(0) {
		v2048 = v662
		v2049 = v38
		goto L4
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	v684 = v661 + int32(32044)
	v685 = int32(146097)
	v686 = base.I32_div_u_s(v684, v685)
	v687 = int32(3)
	v693 = int32(2)
	v698 = base.I32_div_u_s((v686*int32(1073595727)+v684)<<(uint(v693)%32)|v687, v685)
	v701 = v661 + v686*v687 + v698 + int32(32104)
	v702 = int32(1461)
	v703 = base.I32_div_u_s(v701, v702)
	v706 = v703*int32(-1461) + v701
	v708 = v706 << (uint(v693) % 32)
	if base.Ui32(v702) <= base.Ui32(v708) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v721 = base.I32_div_u_s(v708, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v721 + v703<<(uint(int32(2))%32) - int32(4800)
	v729 = v719 + int32(123)
	v733 = int32(base.Ui32(v729*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v729 - int32(base.Ui32(v733*int32(7834))>>(uint(int32(8))%32))
	v740 = int32(1)
	v744 = base.I32_rem_u_s(v733+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v744 + v740
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	if v748 != int32(46) {
		v910 = v740
		goto L167
	} else {
		goto L179
	}
L176:
	;
	v714 = base.I32_rem_u_s(v706+int32(305), int32(365))
	v719 = v714
	goto L175
L177:
	;
	goto L178
L178:
	;
	v718 = base.I32_rem_u_s(v706+int32(306), int32(366))
	v719 = v718
	goto L175
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v667
	v753 = v667 + int32(1)
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753))))
	if v754 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L181:
	;
	v857 = base.F64_mul(v855, float64(8.64e+10))
	if base.F64_lt(base.F64_abs(v857), float64(9.223372036854776e+18)) != 0 {
		goto L211
	} else {
		goto L212
	}
L182:
	;
	v855 = float64(0)
	goto L181
L183:
	;
	goto L184
L184:
	;
	v758 = int32(546909)
	v762 = m.G0
	v764 = v762 - int32(32)
	v765 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v764)+24)) = v765
	*(*int64)(unsafe.Add(mBase, uint32(v764)+16)) = v765
	*(*int64)(unsafe.Add(mBase, uint32(v764)+8)) = v765
	*(*int64)(unsafe.Add(mBase, uint32(v764))) = v765
	v773 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1069])))
	if v773 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v842 = F_strlen(m, v753)
	mBase = m.M
	if v841 != v842 {
		goto L180
	} else {
		goto L206
	}
L186:
	;
	v841 = int32(0)
	goto L185
L187:
	;
	goto L188
L188:
	;
	v777 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1070])))
	if v777 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v781 = v753
	goto L192
L190:
	;
	goto L191
L191:
	;
	v791 = v758
	v792 = v773
	goto L195
L192:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781))))
	if v787 == v773 {
		v781 = v781 + int32(1)
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v841 = v781 - v753
	goto L185
L194:
	;
	goto L193
L195:
	;
	v799 = v764 + int32(base.Ui32(v792)>>(uint(int32(3))%32))&int32(28)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	v801 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v799))) = v800 | v801<<(uint(v792)%32)
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791)+1)))
	if v805 != 0 {
		v791 = v791 + v801
		v792 = v805
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753))))
	if v808 == int32(0) {
		v833 = v753
		goto L198
	} else {
		goto L199
	}
L197:
	;
	goto L196
L198:
	;
	v841 = v833 - v753
	goto L185
L199:
	;
	v812 = v753
	v813 = v808
	goto L200
L200:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v764+int32(base.Ui32(v813)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v821)>>(uint(v813)%32))&int32(1) == int32(0) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v833 = v829
	goto L198
L202:
	;
	v833 = v812
	goto L198
L203:
	;
	goto L204
L204:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812)+1)))
	v829 = v812 + int32(1)
	if v827 != 0 {
		v812 = v829
		v813 = v827
		goto L200
	} else {
		goto L205
	}
L205:
	;
	goto L201
L206:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
	v849 = F_strtod(m, v667, v38+int32(72))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L105
	} else {
		goto L207
	}
L207:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851))))
	if v852 != 0 {
		goto L180
	} else {
		goto L208
	}
L208:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v854 != 0 {
		goto L180
	} else {
		goto L209
	}
L209:
	;
	v855 = v849
	goto L181
L210:
	;
	v866 = base.I64_div_s(v863, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v56))) = uint32(v866)
	v871 = base.I64_extend32_s(v866)*int64(-3600000000) + v863
	v873 = base.I64_div_s(v871, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4+int32(4)))) = uint32(v873)
	v878 = base.I64_extend32_s(v873)*int64(-60000000) + v871
	v880 = base.I64_div_s(v878, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l4))) = uint32(v880)
	v884 = v880*int64(4293967296) + v878
	*(*uint32)(unsafe.Add(mBase, uint32(l5))) = uint32(v884)
	goto L214
L211:
	;
	v861 = base.I64_trunc_f64_s(v857)
	v863 = v861
	goto L210
L212:
	;
	goto L213
L213:
	;
	v863 = int64(-9223372036854775807 - 1)
	goto L210
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31758)
	v910 = v740
	goto L167
L215:
	;
	if v898 < int32(0) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L217:
	;
	goto L218
L218:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v904 != int32(31744) {
		v2048 = int32(-1)
		v2049 = v38
		goto L4
	} else {
		goto L219
	}
L219:
	;
	v910 = v93
	goto L167
L220:
	;
	if v927 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L221:
	;
	if v927 == int32(0) {
		goto L220
	} else {
		goto L225
	}
L222:
	;
	v927 = v921
	goto L224
L223:
	;
	v927 = int32(0)
	goto L224
L224:
	;
	goto L221
L225:
	;
	if v916 != 0 {
		goto L220
	} else {
		goto L226
	}
L226:
	;
	v934 = F_DecodeDate(m, v918, v79, v38+int32(68), v38+int32(59), l4)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L105
	} else {
		goto L227
	}
L227:
	;
	if v934 != 0 {
		v2048 = v934
		v2049 = v38
		goto L4
	} else {
		goto L228
	}
L228:
	;
	v1659 = int32(0)
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L229:
	;
	if v919 < int32(6) {
		goto L234
	} else {
		goto L235
	}
L230:
	;
	v939 = F_strlen(m, v927)
	mBase = m.M
	if base.Ui32(v919-v939) < base.Ui32(int32(3)) {
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v948 = F_DecodeNumberField(m, v919, v918, v79, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L105
	} else {
		goto L232
	}
L232:
	;
	if int32(0) <= v948 {
		v1659 = int32(0)
		v1667 = v86
		v1670 = v89
		v1674 = v93
		v1675 = v94
		v1676 = v95
		v1679 = v98
		v1680 = v99
		goto L11
	} else {
		goto L233
	}
L233:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L234:
	;
	v973 = F_DecodeNumber(m, v919, v918, v94, v79, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L105
	} else {
		goto L242
	}
L235:
	;
	if v79&int32(31744) != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v958 = v916
	goto L238
L237:
	;
	v958 = int32(0)
	goto L238
L238:
	;
	if v958 != 0 {
		goto L234
	} else {
		goto L239
	}
L239:
	;
	v964 = F_DecodeNumberField(m, v919, v918, v79, v38+int32(68), l4, l5, v38+int32(59))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L105
	} else {
		goto L240
	}
L240:
	;
	if int32(0) <= v964 {
		v1659 = int32(0)
		v1667 = v86
		v1670 = v89
		v1674 = v93
		v1675 = v94
		v1676 = v95
		v1679 = v98
		v1680 = v99
		goto L11
	} else {
		goto L241
	}
L241:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L242:
	;
	if v973 != 0 {
		v2048 = v973
		v2049 = v38
		goto L4
	} else {
		goto L243
	}
L243:
	;
	v1659 = int32(0)
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L244:
	;
	if v984 != 0 {
		v2048 = v984
		v2049 = v38
		goto L4
	} else {
		goto L245
	}
L245:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	if v986 == int32(31) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v976)))
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[1068])))
	if v992 != 0 {
		goto L251
	} else {
		goto L252
	}
L247:
	;
	v1236 = v986
	goto L248
L248:
	;
	if v1236 == int32(8) {
		v1698 = v78
		v1699 = v79
		v1706 = v86
		v1709 = v89
		v1713 = v93
		v1714 = v94
		v1715 = v95
		v1718 = v98
		v1719 = v99
		goto L10
	} else {
		goto L296
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v1220
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v1200
	v1236 = v1220
	goto L248
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_consts[1068]))) = v1157
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+12))
	v1184 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1157)+11)))
	v1200 = v1183
	v1220 = v1184
	goto L249
L251:
	;
	goto L256
L252:
	;
	goto L253
L253:
	;
	v1040 = int32(*(*int8)(unsafe.Add(mBase, uint32(v989))))
	v1051 = int32(1642016)
	v1056 = int32(1643152)
	goto L269
L254:
	;
	if v1029-v1030 == int32(0) {
		v1157 = v992
		goto L250
	} else {
		goto L268
	}
L256:
	;
	goto L257
L257:
	;
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989))))
	if v999 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1000 = v989
	v1001 = v992
	v1002 = int32(10)
	v1003 = v999
	goto L262
L259:
	;
	v1025 = v992
	v1029 = int32(0)
	goto L260
L260:
	;
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1025))))
	goto L254
L261:
	;
	v1025 = v1020
	v1029 = v1022
	goto L260
L262:
	;
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001))))
	if v1003 != v1005 {
		v1020 = v1001
		v1022 = v1003
		goto L261
	} else {
		goto L264
	}
L263:
	;
	v1020 = v1014
	v1022 = int32(0)
	goto L261
L264:
	;
	if v1005 == int32(0) {
		v1020 = v1001
		v1022 = v1003
		goto L261
	} else {
		goto L265
	}
L265:
	;
	v1010 = v1002 - int32(1)
	if v1010 == int32(0) {
		v1020 = v1001
		v1022 = v1003
		goto L261
	} else {
		goto L266
	}
L266:
	;
	v1013 = int32(1)
	v1014 = v1001 + v1013
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000)+1)))
	if v1015 != 0 {
		v1000 = v1000 + v1013
		v1001 = v1014
		v1002 = v1010
		v1003 = v1015
		goto L262
	} else {
		goto L267
	}
L267:
	;
	goto L263
L268:
	;
	goto L253
L269:
	;
	v1083 = v1051 + (v1056-v1051)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v1084 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1083))))
	v1085 = v1040 - v1084
	if v1085 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1200 = v1136
	v1220 = int32(31)
	goto L249
L271:
	;
	goto L276
L272:
	;
	v1135 = v1085
	goto L273
L273:
	;
	v1136 = int32(0)
	v1140 = base.B2i32(v1135 < v1136)
	if v1135 < v1136 {
		goto L289
	} else {
		goto L290
	}
L274:
	;
	if v1126 == int32(0) {
		v1157 = v1083
		goto L250
	} else {
		goto L288
	}
L276:
	;
	goto L277
L277:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v989))))
	if v1094 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1095 = v989
	v1096 = v1083
	v1097 = int32(10)
	v1098 = v1094
	goto L282
L279:
	;
	v1120 = v1083
	v1124 = int32(0)
	goto L280
L280:
	;
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1120))))
	v1126 = v1124 - v1125
	goto L274
L281:
	;
	v1120 = v1115
	v1124 = v1117
	goto L280
L282:
	;
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096))))
	if v1098 != v1100 {
		v1115 = v1096
		v1117 = v1098
		goto L281
	} else {
		goto L284
	}
L283:
	;
	v1115 = v1109
	v1117 = int32(0)
	goto L281
L284:
	;
	if v1100 == int32(0) {
		v1115 = v1096
		v1117 = v1098
		goto L281
	} else {
		goto L285
	}
L285:
	;
	v1105 = v1097 - int32(1)
	if v1105 == int32(0) {
		v1115 = v1096
		v1117 = v1098
		goto L281
	} else {
		goto L286
	}
L286:
	;
	v1108 = int32(1)
	v1109 = v1096 + v1108
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1095)+1)))
	if v1110 != 0 {
		v1095 = v1095 + v1108
		v1096 = v1109
		v1097 = v1105
		v1098 = v1110
		goto L282
	} else {
		goto L287
	}
L287:
	;
	goto L283
L288:
	;
	v1135 = v1126
	goto L273
L289:
	;
	v1141 = v1083 - int32(16)
	goto L291
L290:
	;
	v1141 = v1056
	goto L291
L291:
	;
	if v1135 < v1136 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1144 = v1051
	goto L294
L293:
	;
	v1144 = v1083 + int32(16)
	goto L294
L294:
	;
	if base.Ui32(v1144) <= base.Ui32(v1141) {
		v1051 = v1144
		v1056 = v1141
		goto L269
	} else {
		goto L295
	}
L295:
	;
	goto L270
L296:
	;
	v1261 = int32(1) << (uint(v1236) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1261
	v1263 = int32(-1)
	switch v1236 {
	case 0:
		goto L308
	case 1:
		goto L307
	default:
		v2048 = v1263
		v2049 = v38
		goto L4
	case 5:
		goto L304
	case 6:
		goto L305
	case 7:
		goto L303
	case 9:
		goto L302
	case 16:
		goto L300
	case 17:
		goto L299
	case 18:
		goto L301
	case 23:
		goto L298
	case 28:
		goto L306
	case 31:
		goto L297
	}
L297:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v976)))
	v1609 = F_pg_tzset(m, v1608)
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L105
	} else {
		goto L357
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(0)
	v1603 = int32(14)
	if v79&v1603 != v1603 {
		v2048 = v1263
		v2049 = v38
		goto L4
	} else {
		goto L355
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(0)
	if v78 != 0 {
		v2048 = v1263
		v2049 = v38
		goto L4
	} else {
		goto L354
	}
L300:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+24)) = v1596
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L301:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = base.B2i32(v1593 == int32(1))
	goto L11
L302:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1659 = v78
	v1667 = v1592
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1261 | int32(32)
	if l6 == int32(0) {
		v2048 = v1263
		v2049 = v38
		goto L4
	} else {
		goto L353
	}
L304:
	;
	v1577 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v1577
	if l6 == v1577 {
		v2048 = v1263
		v2049 = v38
		goto L4
	} else {
		goto L352
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1261 | int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v2048 = v1263
		v2049 = v38
		goto L4
	} else {
		goto L351
	}
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v1261 | int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = int32(1)
	if l6 == int32(0) {
		v2048 = v1263
		v2049 = v38
		goto L4
	} else {
		goto L350
	}
L307:
	;
	if v94|base.B2i32(v79&int32(2) == int32(0)) != 0 {
		goto L346
	} else {
		goto L347
	}
L308:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	switch v1264 - int32(9) {
	case 0, 1, 2:
		goto L310
	case 3:
		goto L315
	case 4:
		goto L314
	case 5:
		goto L313
	case 6:
		goto L312
	case 7:
		goto L311
	default:
		goto L309
	}
L309:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L105
	} else {
		goto L343
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31790)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1264
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31776)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	v1512 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v1512
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = int64(0)
	if l6 == v1512 {
		v1659 = v78
		v1667 = v86
		v1670 = v89
		v1674 = v93
		v1675 = v94
		v1676 = v95
		v1679 = v98
		v1680 = v99
		goto L11
	} else {
		goto L342
	}
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, v38+int32(8), v38+int32(72), int32(0))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L105
	} else {
		goto L330
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, v38+int32(8), v38+int32(72), int32(0))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L105
	} else {
		goto L329
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(14)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, v38+int32(8), v38+int32(72), int32(0))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L105
	} else {
		goto L317
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(31790)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2)
	F_GetCurrentTimeUsec(m, l4, l5, l6)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L105
	} else {
		goto L316
	}
L316:
	;
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L317:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v1290 = base.B2i32(int32(2) < v1288)
	if int32(2) < v1288 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1291 = int32(4800)
	goto L320
L319:
	;
	v1291 = int32(4799)
	goto L320
L320:
	;
	v1292 = v1285 + v1291
	v1297 = base.I32_div_s(v1292, int32(4))
	v1300 = base.I32_div_s(v1292, int32(-100))
	v1303 = base.I32_div_s(v1292, int32(400))
	if int32(2) < v1288 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1307 = int32(1)
	goto L323
L322:
	;
	v1307 = int32(13)
	goto L323
L323:
	;
	v1312 = base.I32_div_s((v1307+v1288)*int32(7834), int32(256))
	v1315 = v1284 + v1292*int32(365) + v1297 + v1300 + v1303 + v1312 - int32(32168)
	v1319 = v1315 + int32(32044)
	v1320 = int32(146097)
	v1321 = base.I32_div_u_s(v1319, v1320)
	v1322 = int32(3)
	v1328 = int32(2)
	v1333 = base.I32_div_u_s((v1321*int32(1073595727)+v1319)<<(uint(v1328)%32)|v1322, v1320)
	v1336 = v1315 + v1321*v1322 + v1333 + int32(32104)
	v1337 = int32(1461)
	v1338 = base.I32_div_u_s(v1336, v1337)
	v1341 = v1338*int32(-1461) + v1336
	v1343 = v1341 << (uint(v1328) % 32)
	if base.Ui32(v1337) <= base.Ui32(v1343) {
		goto L326
	} else {
		goto L327
	}
L324:
	;
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L325:
	;
	v1356 = base.I32_div_u_s(v1343, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v1356 + v1338<<(uint(int32(2))%32) - int32(4800)
	v1364 = v1354 + int32(123)
	v1368 = int32(base.Ui32(v1364*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1364 - int32(base.Ui32(v1368*int32(7834))>>(uint(int32(8))%32))
	v1378 = base.I32_rem_u_s(v1368+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1378 + int32(1)
	goto L324
L326:
	;
	v1349 = base.I32_rem_u_s(v1341+int32(305), int32(365))
	v1354 = v1349
	goto L325
L327:
	;
	goto L328
L328:
	;
	v1353 = base.I32_rem_u_s(v1341+int32(306), int32(366))
	v1354 = v1353
	goto L325
L329:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v1393
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1395
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1397
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L330:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v1416 = base.B2i32(int32(2) < v1414)
	if int32(2) < v1414 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1417 = int32(4800)
	goto L333
L332:
	;
	v1417 = int32(4799)
	goto L333
L333:
	;
	v1418 = v1411 + v1417
	v1423 = base.I32_div_s(v1418, int32(4))
	v1426 = base.I32_div_s(v1418, int32(-100))
	v1429 = base.I32_div_s(v1418, int32(400))
	if int32(2) < v1414 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1433 = int32(1)
	goto L336
L335:
	;
	v1433 = int32(13)
	goto L336
L336:
	;
	v1438 = base.I32_div_s((v1433+v1414)*int32(7834), int32(256))
	v1441 = v1410 + v1418*int32(365) + v1423 + v1426 + v1429 + v1438 - int32(32166)
	v1445 = v1441 + int32(32044)
	v1446 = int32(146097)
	v1447 = base.I32_div_u_s(v1445, v1446)
	v1448 = int32(3)
	v1454 = int32(2)
	v1459 = base.I32_div_u_s((v1447*int32(1073595727)+v1445)<<(uint(v1454)%32)|v1448, v1446)
	v1462 = v1441 + v1447*v1448 + v1459 + int32(32104)
	v1463 = int32(1461)
	v1464 = base.I32_div_u_s(v1462, v1463)
	v1467 = v1464*int32(-1461) + v1462
	v1469 = v1467 << (uint(v1454) % 32)
	if base.Ui32(v1463) <= base.Ui32(v1469) {
		goto L339
	} else {
		goto L340
	}
L337:
	;
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L338:
	;
	v1482 = base.I32_div_u_s(v1469, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v1482 + v1464<<(uint(int32(2))%32) - int32(4800)
	v1490 = v1480 + int32(123)
	v1494 = int32(base.Ui32(v1490*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1490 - int32(base.Ui32(v1494*int32(7834))>>(uint(int32(8))%32))
	v1504 = base.I32_rem_u_s(v1494+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1504 + int32(1)
	goto L337
L339:
	;
	v1475 = base.I32_rem_u_s(v1467+int32(305), int32(365))
	v1480 = v1475
	goto L338
L340:
	;
	goto L341
L341:
	;
	v1479 = base.I32_rem_u_s(v1467+int32(306), int32(366))
	v1480 = v1479
	goto L338
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v1264
	F_errmsg_internal(m, int32(482562), v38)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L105
	} else {
		goto L344
	}
L344:
	;
	F_errfinish(m, int32(498289), int32(1390), int32(375522))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L105
	} else {
		goto L345
	}
L345:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L346:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v1552
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = int32(1)
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L347:
	;
	if v79&int32(8) != 0 {
		goto L346
	} else {
		goto L348
	}
L348:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if base.Ui32(int32(30)) < base.Ui32(v1543-int32(1)) {
		goto L346
	} else {
		goto L349
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v1543
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(8)
	goto L346
L350:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1562 - v1563
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L351:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1574
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L352:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0) - v1582
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L353:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v976)))
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v1659 = v78
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v1591
	v1679 = v1590
	v1680 = v99
	goto L11
L354:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1659 = v1600
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L355:
	;
	if v78 != 0 {
		v2048 = v1263
		v2049 = v38
		goto L4
	} else {
		goto L356
	}
L356:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1659 = v1607
	v1667 = v86
	v1670 = v89
	v1674 = v93
	v1675 = v94
	v1676 = v95
	v1679 = v98
	v1680 = v99
	goto L11
L357:
	;
	if v1609 != 0 {
		v1633 = v1609
		goto L12
	} else {
		goto L358
	}
L358:
	;
	v2048 = v1263
	v2049 = v38
	goto L4
L359:
	;
	v2048 = int32(-1)
	v2049 = v38
	goto L4
L360:
	;
	goto L361
L361:
	;
	v1698 = v1659
	v1699 = v1683 | v79
	v1706 = v1667
	v1709 = v1670
	v1713 = v1674
	v1714 = v1675
	v1715 = v1676
	v1718 = v1679
	v1719 = v1680
	goto L10
L362:
	;
	goto L9
L363:
	;
	v1735 = v38
	v1738 = v1699
	v1745 = v1706
	v1748 = v1709
	v1752 = v1713
	v1754 = v1715
	v1757 = v1718
	v1758 = v1719
	goto L7
L364:
	;
	v2048 = int32(0)
	v2049 = v1735
	goto L4
L365:
	;
	goto L366
L366:
	;
	v1765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1735)+59)))
	if v1752 != 0 {
		goto L368
	} else {
		goto L369
	}
L367:
	;
	if v1933 != 0 {
		v2048 = v1933
		v2049 = v1735
		goto L4
	} else {
		goto L405
	}
L368:
	;
	if v1738&int32(32768) != 0 {
		goto L386
	} else {
		goto L387
	}
L369:
	;
	if v1738&int32(4) == int32(0) {
		goto L368
	} else {
		goto L370
	}
L370:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1758 != 0 {
		goto L373
	} else {
		goto L374
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1790
	goto L368
L372:
	;
	v1790 = int32(1) - v1770
	goto L371
L373:
	;
	if int32(0) < v1770 {
		goto L372
	} else {
		goto L376
	}
L374:
	;
	goto L375
L375:
	;
	if v1765 != 0 {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	v1933 = int32(-2)
	goto L367
L377:
	;
	if v1770 < int32(0) {
		goto L380
	} else {
		goto L381
	}
L378:
	;
	goto L379
L379:
	;
	if int32(0) < v1770 {
		goto L368
	} else {
		goto L385
	}
L380:
	;
	v1933 = int32(-2)
	goto L367
L381:
	;
	goto L382
L382:
	;
	if base.Ui32(v1770) <= base.Ui32(int32(69)) {
		v1790 = v1770 + int32(2000)
		goto L371
	} else {
		goto L383
	}
L383:
	;
	if base.Ui32(int32(99)) < base.Ui32(v1770) {
		goto L368
	} else {
		goto L384
	}
L384:
	;
	v1790 = v1770 + int32(1900)
	goto L371
L385:
	;
	v1933 = int32(-2)
	goto L367
L386:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v1801 = v1796 + int32(4799)
	v1803 = base.I32_div_s(v1801, int32(4))
	v1806 = base.I32_div_s(v1801, int32(-100))
	v1809 = base.I32_div_s(v1801, int32(400))
	v1810 = v1795 + v1796*int32(365) + v1803 + v1806 + v1809
	v1812 = v1810 + int32(1751940)
	v1813 = int32(146097)
	v1814 = base.I32_div_u_s(v1812, v1813)
	v1815 = int32(3)
	v1821 = int32(2)
	v1826 = base.I32_div_u_s((v1814*int32(1073595727)+v1812)<<(uint(v1821)%32)|v1815, v1813)
	v1829 = v1810 + v1814*v1815 + v1826 + int32(1752000)
	v1830 = int32(1461)
	v1831 = base.I32_div_u_s(v1829, v1830)
	v1834 = v1831*int32(-1461) + v1829
	v1836 = v1834 << (uint(v1821) % 32)
	if base.Ui32(v1830) <= base.Ui32(v1836) {
		goto L390
	} else {
		goto L391
	}
L387:
	;
	goto L388
L388:
	;
	if v1738&int32(2) == int32(0) {
		goto L393
	} else {
		goto L394
	}
L389:
	;
	v1849 = base.I32_div_u_s(v1836, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v1849 + v1831<<(uint(int32(2))%32) - int32(4800)
	v1857 = v1847 + int32(123)
	v1861 = int32(base.Ui32(v1857*int32(2141)) >> (uint(int32(16)) % 32))
	v1865 = base.I32_rem_u_s(v1861+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v1865 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v1857 - int32(base.Ui32(v1861*int32(7834))>>(uint(int32(8))%32))
	goto L388
L390:
	;
	v1842 = base.I32_rem_u_s(v1834+int32(305), int32(365))
	v1847 = v1842
	goto L389
L391:
	;
	goto L392
L392:
	;
	v1846 = base.I32_rem_u_s(v1834+int32(306), int32(366))
	v1847 = v1846
	goto L389
L393:
	;
	if v1738&int32(8) == int32(0) {
		goto L396
	} else {
		goto L397
	}
L394:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if base.Ui32(int32(-12)) <= base.Ui32(v1882-int32(13)) {
		goto L393
	} else {
		goto L395
	}
L395:
	;
	v1933 = int32(-3)
	goto L367
L396:
	;
	v1898 = int32(14)
	if v1738&v1898 != v1898 {
		goto L399
	} else {
		goto L400
	}
L397:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if base.Ui32(int32(-31)) <= base.Ui32(v1892-int32(32)) {
		goto L396
	} else {
		goto L398
	}
L398:
	;
	v1933 = int32(-3)
	goto L367
L399:
	;
	v1933 = int32(0)
	goto L367
L400:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v1904&int32(3) != 0 {
		v1914 = int32(0)
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v1914*int32(52)+v1917<<(uint(int32(2))%32))+uint32(_consts[1071])))
	if v1902 <= v1923 {
		goto L399
	} else {
		goto L404
	}
L402:
	;
	v1909 = base.I32_rem_s(v1904, int32(100))
	if v1909 != 0 {
		v1914 = int32(1)
		goto L401
	} else {
		goto L403
	}
L403:
	;
	v1911 = base.I32_rem_s(v1904, int32(400))
	v1914 = base.B2i32(v1911 == int32(0))
	goto L401
L404:
	;
	v1933 = int32(-2)
	goto L367
L405:
	;
	if v1745 == int32(2) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1950 = int32(14)
	if v1738&v1950 != v1950 {
		goto L416
	} else {
		goto L417
	}
L407:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if int32(12) < v1936 {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v2048 = int32(-2)
	v2049 = v1735
	goto L4
L409:
	;
	goto L410
L410:
	;
	switch v1745 {
	case 0:
		goto L413
	case 1:
		goto L412
	default:
		goto L406
	}
L411:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v1947
	goto L406
L412:
	;
	if v1936 == int32(12) {
		goto L406
	} else {
		goto L415
	}
L413:
	;
	if v1936 == int32(12) {
		v1947 = int32(0)
		goto L411
	} else {
		goto L414
	}
L414:
	;
	goto L406
L415:
	;
	v1947 = v1936 + int32(12)
	goto L411
L416:
	;
	v1956 = int32(31744)
	if v1738&v1956 == v1956 {
		goto L419
	} else {
		goto L420
	}
L417:
	;
	goto L418
L418:
	;
	if v1748 != 0 {
		goto L422
	} else {
		goto L423
	}
L419:
	;
	v1960 = int32(1)
	goto L421
L420:
	;
	v1960 = int32(-1)
	goto L421
L421:
	;
	v2048 = v1960
	v2049 = v1735
	goto L4
L422:
	;
	if v1738&int32(268435456) != 0 {
		goto L425
	} else {
		goto L426
	}
L423:
	;
	goto L424
L424:
	;
	if v1754 != 0 {
		goto L428
	} else {
		goto L429
	}
L425:
	;
	v2048 = int32(-1)
	v2049 = v1735
	goto L4
L426:
	;
	goto L427
L427:
	;
	v1966 = F_DetermineTimeZoneOffsetInternal(m, l4, v1748, v1735+int32(72))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1966
	goto L424
L428:
	;
	if v1738&int32(268435456) != 0 {
		goto L431
	} else {
		goto L432
	}
L429:
	;
	goto L430
L430:
	;
	if l6 == int32(0) {
		goto L444
	} else {
		goto L445
	}
L431:
	;
	v2048 = int32(-1)
	v2049 = v1735
	goto L4
L432:
	;
	goto L433
L433:
	;
	v1974 = m.G0
	v1976 = v1974 - int32(288)
	m.G0 = v1976
	v1980 = F_DetermineTimeZoneOffsetInternal(m, l4, v1754, v1976+int32(280))
	mBase = m.M
	v1984 = F_strlcpy(m, v1976+int32(16), v1757, int32(256))
	mBase = m.M
	v1985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1976)+16)))
	if v1985 != 0 {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2020
	goto L430
L435:
	;
	v1989 = v1976 + int32(16)
	v1993 = v1985
	goto L438
L436:
	;
	goto L437
L437:
	;
	v2013 = F_pg_interpret_timezone_abbrev(m, v1976+int32(16), v1976+int32(280), v1976+int32(12), v1976+int32(8), v1754)
	mBase = m.M
	if v2013 != 0 {
		goto L441
	} else {
		goto L442
	}
L438:
	;
	v1994 = F_pg_toupper(m, v1993)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v1989))) = uint8(v1994)
	v1997 = v1989 + int32(1)
	v1998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1997))))
	if v1998 != 0 {
		v1989 = v1997
		v1993 = v1998
		goto L438
	} else {
		goto L440
	}
L439:
	;
	goto L437
L440:
	;
	goto L439
L441:
	;
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v1976)+12))
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1976)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+32)) = v2015
	v2020 = int32(0) - v2014
	goto L443
L442:
	;
	v2020 = v1980
	goto L443
L443:
	;
	m.G0 = v1976 + int32(288)
	goto L434
L444:
	;
	v2048 = int32(0)
	v2049 = v1735
	goto L4
L445:
	;
	goto L446
L446:
	;
	v2028 = int32(0)
	if v1738&int32(32) != 0 {
		v2048 = v2028
		v2049 = v1735
		goto L4
	} else {
		goto L447
	}
L447:
	;
	if v1738&int32(268435456) != 0 {
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v2048 = int32(-1)
	v2049 = v1735
	goto L4
L449:
	;
	goto L450
L450:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
	v2038 = F_DetermineTimeZoneOffsetInternal(m, l4, v2035, v1735+int32(72))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v2038
	v2048 = v2028
	v2049 = v1735
	goto L4
}
func F_EncodeDateTime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
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
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
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
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1264 int32
	_ = v1264
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1384 int32
	_ = v1384
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v19 = l2 & base.B2i32(int32(0) <= v16)
	switch l5 - int32(1) {
	case 0, 3:
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v22 {
			v27 = v22
		} else {
			v27 = int32(1) - v22
		}
		v28 = int32(4)
		if base.Ui32(int32(99)) < base.Ui32(v27) {
		} else {
		}
		v42 = F_pg_ultoa_n(m, v27, l6)
		mBase = m.M
		if v28 <= v42 {
			v53 = l6 + v42
		} else {
			v45 = l6 + v28
			v47 = F_memmove(m, v45-v42, l6, v42)
			mBase = m.M
			v50 = F___memset(m, l6, int32(48), v28-v42)
			mBase = m.M
			v53 = v45
		}
		v54 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v54)
		v57 = v53 + int32(1)
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v59 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v58) {
			v73 = F_pg_ultoa_n(m, v58, v57)
			mBase = m.M
			if v59 <= v73 {
				v84 = v57 + v73
			} else {
				v76 = v53 + int32(3)
				v78 = F_memmove(m, v76-v73, v57, v73)
				mBase = m.M
				v81 = F___memset(m, v57, int32(48), v59-v73)
				mBase = m.M
				v84 = v76
			}
		} else {
			v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v57))) = uint16(v69)
			v84 = v53 + int32(3)
		}
		v85 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v84))) = uint8(v85)
		v88 = v84 + int32(1)
		v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v90 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v89) {
			v104 = F_pg_ultoa_n(m, v89, v88)
			mBase = m.M
			if v90 <= v104 {
				v115 = v88 + v104
			} else {
				v107 = v84 + int32(3)
				v109 = F_memmove(m, v107-v104, v88, v104)
				mBase = m.M
				v112 = F___memset(m, v88, int32(48), v90-v104)
				mBase = m.M
				v115 = v107
			}
		} else {
			v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v88))) = uint16(v100)
			v115 = v84 + int32(3)
		}
		if l5 == int32(1) {
			v120 = int32(32)
		} else {
			v120 = int32(84)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v120)
		v123 = v115 + int32(1)
		v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v125 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v124) {
			v139 = F_pg_ultoa_n(m, v124, v123)
			mBase = m.M
			if v125 <= v139 {
				v150 = v123 + v139
			} else {
				v142 = v115 + int32(3)
				v144 = F_memmove(m, v142-v139, v123, v139)
				mBase = m.M
				v147 = F___memset(m, v123, int32(48), v125-v139)
				mBase = m.M
				v150 = v142
			}
		} else {
			v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v123))) = uint16(v135)
			v150 = v115 + int32(3)
		}
		v151 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v151)
		v154 = v150 + int32(1)
		v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v156 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v155) {
			v170 = F_pg_ultoa_n(m, v155, v154)
			mBase = m.M
			if v156 <= v170 {
				v181 = v154 + v170
			} else {
				v173 = v150 + int32(3)
				v175 = F_memmove(m, v173-v170, v154, v170)
				mBase = m.M
				v178 = F___memset(m, v154, int32(48), v156-v170)
				mBase = m.M
				v181 = v173
			}
		} else {
			v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v154))) = uint16(v166)
			v181 = v150 + int32(3)
		}
		v182 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v182)
		v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v192 = v186 >> (uint(int32(31)) % 32)
		v196 = F_pg_ultostr_zeropad(m, v181+int32(1), v186^v192-v192, int32(2))
		mBase = m.M
		if l1 == int32(0) {
			v304 = v196
		} else {
			v201 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v201)
			v204 = l1 >> (uint(int32(31)) % 32)
			v206 = l1 ^ v204 - v204
			v208 = base.I32_div_s(v206, int32(10))
			v211 = v208*int32(-10) + v206
			if v211 != 0 {
				v213 = v211 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v196)+6)) = uint8(v213)
				v219 = v196 + int32(7)
			} else {
				v219 = v196 + int32(6)
			}
			v221 = base.I32_div_s(v206, int32(100))
			v224 = v221*int32(-10) + v208
			v225 = v211 | v224
			if v225 == int32(0) {
				v233 = v196 + int32(5)
			} else {
				v231 = v224 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v196)+5)) = uint8(v231)
				v233 = v219
			}
			v235 = base.I32_div_s(v206, int32(1000))
			v238 = v235*int32(-10) + v221
			v239 = v225 | v238
			if v239 == int32(0) {
				v247 = v196 + int32(4)
			} else {
				v245 = v238 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)) = uint8(v245)
				v247 = v233
			}
			v249 = base.I32_div_s(v206, int32(10000))
			v252 = v249*int32(-10) + v235
			v253 = v239 | v252
			if v253 == int32(0) {
				v261 = v196 + int32(3)
			} else {
				v259 = v252 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v196)+3)) = uint8(v259)
				v261 = v247
			}
			v263 = base.I32_div_s(v206, int32(100000))
			v266 = v263*int32(-10) + v249
			v267 = v253 | v266
			if v267 == int32(0) {
				v275 = v196 + int32(2)
			} else {
				v273 = v266 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v196)+2)) = uint8(v273)
				v275 = v261
			}
			v277 = base.I32_div_s(v206, int32(1000000))
			v280 = v277*int32(-10) + v263
			if v267|v280 == int32(0) {
				v289 = v196 + int32(1)
			} else {
				v287 = v280 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)) = uint8(v287)
				v289 = v275
			}
			if base.Ui32(int32(19)) <= base.Ui32(v263+int32(9)) {
				v296 = F_pg_ultostr(m, v196+int32(1), v206)
				mBase = m.M
				v297 = v296
			} else {
				v297 = v289
			}
			v304 = v297
		}
		if v19 == int32(0) {
			v1662 = v304
		} else {
			if l3 <= int32(0) {
				v314 = int32(43)
			} else {
				v314 = int32(45)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v304))) = uint8(v314)
			v317 = l3 >> (uint(int32(31)) % 32)
			v319 = l3 ^ v317 - v317
			v321 = base.I32_div_s(v319, int32(3600))
			v322 = int32(-60)
			v325 = base.I32_div_s(v319, int32(60))
			v326 = v321*v322 + v325
			v328 = v304 + int32(1)
			v331 = v325*v322 + v319
			if v331 != 0 {
				v332 = int32(2)
				v333 = F_pg_ultostr_zeropad(m, v328, v321, v332)
				mBase = m.M
				v334 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v333))) = uint8(v334)
				v339 = F_pg_ultostr_zeropad(m, v333+int32(1), v326, v332)
				mBase = m.M
				v346 = v339
				v347 = v331
				v348 = int32(58)
				*(*uint8)(unsafe.Add(mBase, uint32(v346))) = uint8(v348)
				v353 = F_pg_ultostr_zeropad(m, v346+int32(1), v347, int32(2))
				mBase = m.M
				v354 = v353
			} else {
				v341 = F_pg_ultostr_zeropad(m, v328, v321, int32(2))
				mBase = m.M
				if l5 == int32(4) {
					v346 = v341
					v347 = v326
					v348 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v346))) = uint8(v348)
					v353 = F_pg_ultostr_zeropad(m, v346+int32(1), v347, int32(2))
					mBase = m.M
					v354 = v353
				} else {
					if v326 == int32(0) {
						v354 = v341
					} else {
						v346 = v341
						v347 = v326
						v348 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v346))) = uint8(v348)
						v353 = F_pg_ultostr_zeropad(m, v346+int32(1), v347, int32(2))
						mBase = m.M
						v354 = v353
					}
				}
			}
			v1662 = v354
		}
		v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v1663 <= int32(0) {
			v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
			*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
			v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
			*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
			v1674 = v1662 + int32(3)
		} else {
			v1674 = v1662
		}
		v1675 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
		m.G0 = v14 + int32(48)
		return
	case 1:
		v359 = *(*int32)(unsafe.Add(mBase, _consts[1061]))
		v361 = base.B2i32(v359 == int32(1))
		if v359 == int32(1) {
			v362 = int32(12)
		} else {
			v362 = int32(16)
		}
		v364 = *(*int32)(unsafe.Add(mBase, uint32(l0+v362)))
		v365 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v364) {
			v379 = F_pg_ultoa_n(m, v364, l6)
			mBase = m.M
			if v365 <= v379 {
				v390 = l6 + v379
			} else {
				v382 = l6 + v365
				v384 = F_memmove(m, v382-v379, l6, v379)
				mBase = m.M
				v387 = F___memset(m, l6, int32(48), v365-v379)
				mBase = m.M
				v390 = v382
			}
		} else {
			v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v364<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v375)
			v390 = l6 + int32(2)
		}
		v391 = int32(47)
		*(*uint8)(unsafe.Add(mBase, uint32(v390))) = uint8(v391)
		v394 = v390 + int32(1)
		if v359 == int32(1) {
			v397 = int32(16)
		} else {
			v397 = int32(12)
		}
		v399 = *(*int32)(unsafe.Add(mBase, uint32(l0+v397)))
		v400 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v399) {
			v414 = F_pg_ultoa_n(m, v399, v394)
			mBase = m.M
			if v400 <= v414 {
				v425 = v394 + v414
			} else {
				v417 = v390 + int32(3)
				v419 = F_memmove(m, v417-v414, v394, v414)
				mBase = m.M
				v422 = F___memset(m, v394, int32(48), v400-v414)
				mBase = m.M
				v425 = v417
			}
		} else {
			v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v394))) = uint16(v410)
			v425 = v390 + int32(3)
		}
		v426 = int32(47)
		*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v426)
		v428 = int32(1)
		v429 = v425 + v428
		v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v430 {
			v435 = v430
		} else {
			v435 = v428 - v430
		}
		v436 = int32(4)
		if base.Ui32(int32(99)) < base.Ui32(v435) {
		} else {
		}
		v450 = F_pg_ultoa_n(m, v435, v429)
		mBase = m.M
		if v436 <= v450 {
			v461 = v429 + v450
		} else {
			v453 = v425 + int32(5)
			v455 = F_memmove(m, v453-v450, v429, v450)
			mBase = m.M
			v458 = F___memset(m, v429, int32(48), v436-v450)
			mBase = m.M
			v461 = v453
		}
		v462 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v461))) = uint8(v462)
		v465 = v461 + int32(1)
		v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v467 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v466) {
			v481 = F_pg_ultoa_n(m, v466, v465)
			mBase = m.M
			if v467 <= v481 {
				v492 = v465 + v481
			} else {
				v484 = v461 + int32(3)
				v486 = F_memmove(m, v484-v481, v465, v481)
				mBase = m.M
				v489 = F___memset(m, v465, int32(48), v467-v481)
				mBase = m.M
				v492 = v484
			}
		} else {
			v477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v466<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v465))) = uint16(v477)
			v492 = v461 + int32(3)
		}
		v493 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v492))) = uint8(v493)
		v496 = v492 + int32(1)
		v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v498 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v497) {
			v512 = F_pg_ultoa_n(m, v497, v496)
			mBase = m.M
			if v498 <= v512 {
				v523 = v496 + v512
			} else {
				v515 = v492 + int32(3)
				v517 = F_memmove(m, v515-v512, v496, v512)
				mBase = m.M
				v520 = F___memset(m, v496, int32(48), v498-v512)
				mBase = m.M
				v523 = v515
			}
		} else {
			v508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v497<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v496))) = uint16(v508)
			v523 = v492 + int32(3)
		}
		v524 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v523))) = uint8(v524)
		v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v534 = v528 >> (uint(int32(31)) % 32)
		v538 = F_pg_ultostr_zeropad(m, v523+int32(1), v528^v534-v534, int32(2))
		mBase = m.M
		if l1 == int32(0) {
			v646 = v538
		} else {
			v543 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v538))) = uint8(v543)
			v546 = l1 >> (uint(int32(31)) % 32)
			v548 = l1 ^ v546 - v546
			v550 = base.I32_div_s(v548, int32(10))
			v553 = v550*int32(-10) + v548
			if v553 != 0 {
				v555 = v553 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v538)+6)) = uint8(v555)
				v561 = v538 + int32(7)
			} else {
				v561 = v538 + int32(6)
			}
			v563 = base.I32_div_s(v548, int32(100))
			v566 = v563*int32(-10) + v550
			v567 = v553 | v566
			if v567 == int32(0) {
				v575 = v538 + int32(5)
			} else {
				v573 = v566 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v538)+5)) = uint8(v573)
				v575 = v561
			}
			v577 = base.I32_div_s(v548, int32(1000))
			v580 = v577*int32(-10) + v563
			v581 = v567 | v580
			if v581 == int32(0) {
				v589 = v538 + int32(4)
			} else {
				v587 = v580 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v538)+4)) = uint8(v587)
				v589 = v575
			}
			v591 = base.I32_div_s(v548, int32(10000))
			v594 = v591*int32(-10) + v577
			v595 = v581 | v594
			if v595 == int32(0) {
				v603 = v538 + int32(3)
			} else {
				v601 = v594 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v538)+3)) = uint8(v601)
				v603 = v589
			}
			v605 = base.I32_div_s(v548, int32(100000))
			v608 = v605*int32(-10) + v591
			v609 = v595 | v608
			if v609 == int32(0) {
				v617 = v538 + int32(2)
			} else {
				v615 = v608 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v538)+2)) = uint8(v615)
				v617 = v603
			}
			v619 = base.I32_div_s(v548, int32(1000000))
			v622 = v619*int32(-10) + v605
			if v609|v622 == int32(0) {
				v631 = v538 + int32(1)
			} else {
				v629 = v622 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v538)+1)) = uint8(v629)
				v631 = v617
			}
			if base.Ui32(int32(19)) <= base.Ui32(v605+int32(9)) {
				v638 = F_pg_ultostr(m, v538+int32(1), v548)
				mBase = m.M
				v639 = v638
			} else {
				v639 = v631
			}
			v646 = v639
		}
		if v19 == int32(0) {
			v1662 = v646
			v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v1663 <= int32(0) {
				v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
				*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
				v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
				*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
				v1674 = v1662 + int32(3)
			} else {
				v1674 = v1662
			}
			v1675 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
			m.G0 = v14 + int32(48)
			return
		} else {
			if l4 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(10)
				v655 = F_pg_sprintf(m, v646, int32(175134), v14+int32(16))
				mBase = m.M
				v656 = m.ExcPending
				if v656 != 0 {
					return
				} else {
					v657 = F_strlen(m, v646)
					mBase = m.M
					v1662 = v657 + v646
					v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v1663 <= int32(0) {
						v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
						*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
						v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
						*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
						v1674 = v1662 + int32(3)
					} else {
						v1674 = v1662
					}
					v1675 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
					m.G0 = v14 + int32(48)
					return
				}
			} else {
				if l3 <= int32(0) {
					v663 = int32(43)
				} else {
					v663 = int32(45)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v646))) = uint8(v663)
				v666 = l3 >> (uint(int32(31)) % 32)
				v668 = l3 ^ v666 - v666
				v670 = base.I32_div_s(v668, int32(3600))
				v671 = int32(-60)
				v674 = base.I32_div_s(v668, int32(60))
				v675 = v670*v671 + v674
				v677 = v646 + int32(1)
				v680 = v674*v671 + v668
				if v680 != 0 {
					v681 = int32(2)
					if base.Ui32(int32(99)) < base.Ui32(v670) {
						v695 = F_pg_ultoa_n(m, v670, v677)
						mBase = m.M
						if v681 <= v695 {
							v706 = v677 + v695
						} else {
							v698 = v646 + int32(3)
							v700 = F_memmove(m, v698-v695, v677, v695)
							mBase = m.M
							v703 = F___memset(m, v677, int32(48), v681-v695)
							mBase = m.M
							v706 = v698
						}
					} else {
						v691 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670<<(uint(int32(1))%32))+uint32(_consts[1073]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v677))) = uint16(v691)
						v706 = v646 + int32(3)
					}
					v707 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v706))) = uint8(v707)
					v710 = v706 + int32(1)
					v711 = int32(2)
					if base.Ui32(int32(99)) < base.Ui32(v675) {
						v725 = F_pg_ultoa_n(m, v675, v710)
						mBase = m.M
						if v711 <= v725 {
							v736 = v710 + v725
						} else {
							v728 = v706 + int32(3)
							v730 = F_memmove(m, v728-v725, v710, v725)
							mBase = m.M
							v733 = F___memset(m, v710, int32(48), v711-v725)
							mBase = m.M
							v736 = v728
						}
					} else {
						v721 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v675<<(uint(int32(1))%32))+uint32(_consts[1073]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v710))) = uint16(v721)
						v736 = v706 + int32(3)
					}
					v765 = v680
					v766 = v736
					v767 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v766))) = uint8(v767)
					v770 = v766 + int32(1)
					v771 = int32(2)
					if base.Ui32(int32(99)) < base.Ui32(v765) {
						v785 = F_pg_ultoa_n(m, v765, v770)
						mBase = m.M
						if v771 <= v785 {
							v796 = v770 + v785
						} else {
							v788 = v766 + int32(3)
							v790 = F_memmove(m, v788-v785, v770, v785)
							mBase = m.M
							v793 = F___memset(m, v770, int32(48), v771-v785)
							mBase = m.M
							v796 = v788
						}
					} else {
						v781 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v765<<(uint(int32(1))%32))+uint32(_consts[1073]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v770))) = uint16(v781)
						v796 = v766 + int32(3)
					}
					v1662 = v796
				} else {
					v737 = int32(2)
					if base.Ui32(int32(99)) < base.Ui32(v670) {
						v751 = F_pg_ultoa_n(m, v670, v677)
						mBase = m.M
						if v737 <= v751 {
							v762 = v677 + v751
						} else {
							v754 = v646 + int32(3)
							v756 = F_memmove(m, v754-v751, v677, v751)
							mBase = m.M
							v759 = F___memset(m, v677, int32(48), v737-v751)
							mBase = m.M
							v762 = v754
						}
					} else {
						v747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670<<(uint(int32(1))%32))+uint32(_consts[1073]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v677))) = uint16(v747)
						v762 = v646 + int32(3)
					}
					if v675 == int32(0) {
						v1662 = v762
					} else {
						v765 = v675
						v766 = v762
						v767 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v766))) = uint8(v767)
						v770 = v766 + int32(1)
						v771 = int32(2)
						if base.Ui32(int32(99)) < base.Ui32(v765) {
							v785 = F_pg_ultoa_n(m, v765, v770)
							mBase = m.M
							if v771 <= v785 {
								v796 = v770 + v785
							} else {
								v788 = v766 + int32(3)
								v790 = F_memmove(m, v788-v785, v770, v785)
								mBase = m.M
								v793 = F___memset(m, v770, int32(48), v771-v785)
								mBase = m.M
								v796 = v788
							}
						} else {
							v781 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v765<<(uint(int32(1))%32))+uint32(_consts[1073]))))
							*(*uint16)(unsafe.Add(mBase, uint32(v770))) = uint16(v781)
							v796 = v766 + int32(3)
						}
						v1662 = v796
					}
				}
				v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v1663 <= int32(0) {
					v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
					*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
					v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
					*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
					v1674 = v1662 + int32(3)
				} else {
					v1674 = v1662
				}
				v1675 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
				m.G0 = v14 + int32(48)
				return
			}
		}
	case 2:
		v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v798 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v797) {
			v812 = F_pg_ultoa_n(m, v797, l6)
			mBase = m.M
			if v798 <= v812 {
				v823 = l6 + v812
			} else {
				v815 = l6 + v798
				v817 = F_memmove(m, v815-v812, l6, v812)
				mBase = m.M
				v820 = F___memset(m, l6, int32(48), v798-v812)
				mBase = m.M
				v823 = v815
			}
		} else {
			v808 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v797<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v808)
			v823 = l6 + int32(2)
		}
		v824 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v823))) = uint8(v824)
		v827 = v823 + int32(1)
		v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v829 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v828) {
			v843 = F_pg_ultoa_n(m, v828, v827)
			mBase = m.M
			if v829 <= v843 {
				v854 = v827 + v843
			} else {
				v846 = v823 + int32(3)
				v848 = F_memmove(m, v846-v843, v827, v843)
				mBase = m.M
				v851 = F___memset(m, v827, int32(48), v829-v843)
				mBase = m.M
				v854 = v846
			}
		} else {
			v839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v828<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v827))) = uint16(v839)
			v854 = v823 + int32(3)
		}
		v855 = int32(46)
		*(*uint8)(unsafe.Add(mBase, uint32(v854))) = uint8(v855)
		v857 = int32(1)
		v858 = v854 + v857
		v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v859 {
			v864 = v859
		} else {
			v864 = v857 - v859
		}
		v865 = int32(4)
		if base.Ui32(int32(99)) < base.Ui32(v864) {
		} else {
		}
		v879 = F_pg_ultoa_n(m, v864, v858)
		mBase = m.M
		if v865 <= v879 {
			v890 = v858 + v879
		} else {
			v882 = v854 + int32(5)
			v884 = F_memmove(m, v882-v879, v858, v879)
			mBase = m.M
			v887 = F___memset(m, v858, int32(48), v865-v879)
			mBase = m.M
			v890 = v882
		}
		v891 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v890))) = uint8(v891)
		v894 = v890 + int32(1)
		v895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v896 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v895) {
			v910 = F_pg_ultoa_n(m, v895, v894)
			mBase = m.M
			if v896 <= v910 {
				v921 = v894 + v910
			} else {
				v913 = v890 + int32(3)
				v915 = F_memmove(m, v913-v910, v894, v910)
				mBase = m.M
				v918 = F___memset(m, v894, int32(48), v896-v910)
				mBase = m.M
				v921 = v913
			}
		} else {
			v906 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v895<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v894))) = uint16(v906)
			v921 = v890 + int32(3)
		}
		v922 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v921))) = uint8(v922)
		v925 = v921 + int32(1)
		v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v927 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v926) {
			v941 = F_pg_ultoa_n(m, v926, v925)
			mBase = m.M
			if v927 <= v941 {
				v952 = v925 + v941
			} else {
				v944 = v921 + int32(3)
				v946 = F_memmove(m, v944-v941, v925, v941)
				mBase = m.M
				v949 = F___memset(m, v925, int32(48), v927-v941)
				mBase = m.M
				v952 = v944
			}
		} else {
			v937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v926<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v925))) = uint16(v937)
			v952 = v921 + int32(3)
		}
		v953 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v952))) = uint8(v953)
		v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v963 = v957 >> (uint(int32(31)) % 32)
		v967 = F_pg_ultostr_zeropad(m, v952+int32(1), v957^v963-v963, int32(2))
		mBase = m.M
		if l1 == int32(0) {
			v1075 = v967
		} else {
			v972 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v967))) = uint8(v972)
			v975 = l1 >> (uint(int32(31)) % 32)
			v977 = l1 ^ v975 - v975
			v979 = base.I32_div_s(v977, int32(10))
			v982 = v979*int32(-10) + v977
			if v982 != 0 {
				v984 = v982 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v967)+6)) = uint8(v984)
				v990 = v967 + int32(7)
			} else {
				v990 = v967 + int32(6)
			}
			v992 = base.I32_div_s(v977, int32(100))
			v995 = v992*int32(-10) + v979
			v996 = v982 | v995
			if v996 == int32(0) {
				v1004 = v967 + int32(5)
			} else {
				v1002 = v995 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v967)+5)) = uint8(v1002)
				v1004 = v990
			}
			v1006 = base.I32_div_s(v977, int32(1000))
			v1009 = v1006*int32(-10) + v992
			v1010 = v996 | v1009
			if v1010 == int32(0) {
				v1018 = v967 + int32(4)
			} else {
				v1016 = v1009 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v967)+4)) = uint8(v1016)
				v1018 = v1004
			}
			v1020 = base.I32_div_s(v977, int32(10000))
			v1023 = v1020*int32(-10) + v1006
			v1024 = v1010 | v1023
			if v1024 == int32(0) {
				v1032 = v967 + int32(3)
			} else {
				v1030 = v1023 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v967)+3)) = uint8(v1030)
				v1032 = v1018
			}
			v1034 = base.I32_div_s(v977, int32(100000))
			v1037 = v1034*int32(-10) + v1020
			v1038 = v1024 | v1037
			if v1038 == int32(0) {
				v1046 = v967 + int32(2)
			} else {
				v1044 = v1037 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v967)+2)) = uint8(v1044)
				v1046 = v1032
			}
			v1048 = base.I32_div_s(v977, int32(1000000))
			v1051 = v1048*int32(-10) + v1034
			if v1038|v1051 == int32(0) {
				v1060 = v967 + int32(1)
			} else {
				v1058 = v1051 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v967)+1)) = uint8(v1058)
				v1060 = v1046
			}
			if base.Ui32(int32(19)) <= base.Ui32(v1034+int32(9)) {
				v1067 = F_pg_ultostr(m, v967+int32(1), v977)
				mBase = m.M
				v1068 = v1067
			} else {
				v1068 = v1060
			}
			v1075 = v1068
		}
		if v19 == int32(0) {
			v1662 = v1075
			v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v1663 <= int32(0) {
				v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
				*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
				v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
				*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
				v1674 = v1662 + int32(3)
			} else {
				v1674 = v1662
			}
			v1675 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
			m.G0 = v14 + int32(48)
			return
		} else {
			if l4 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(10)
				v1084 = F_pg_sprintf(m, v1075, int32(175134), v14+int32(32))
				mBase = m.M
				v1085 = m.ExcPending
				if v1085 != 0 {
					return
				} else {
					v1086 = F_strlen(m, v1075)
					mBase = m.M
					v1662 = v1086 + v1075
					v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v1663 <= int32(0) {
						v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
						*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
						v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
						*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
						v1674 = v1662 + int32(3)
					} else {
						v1674 = v1662
					}
					v1675 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
					m.G0 = v14 + int32(48)
					return
				}
			} else {
				if l3 <= int32(0) {
					v1092 = int32(43)
				} else {
					v1092 = int32(45)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v1075))) = uint8(v1092)
				v1095 = l3 >> (uint(int32(31)) % 32)
				v1097 = l3 ^ v1095 - v1095
				v1099 = base.I32_div_s(v1097, int32(3600))
				v1100 = int32(-60)
				v1103 = base.I32_div_s(v1097, int32(60))
				v1104 = v1099*v1100 + v1103
				v1106 = v1075 + int32(1)
				v1109 = v1103*v1100 + v1097
				if v1109 != 0 {
					v1110 = int32(2)
					if base.Ui32(int32(99)) < base.Ui32(v1099) {
						v1124 = F_pg_ultoa_n(m, v1099, v1106)
						mBase = m.M
						if v1110 <= v1124 {
							v1135 = v1106 + v1124
						} else {
							v1127 = v1075 + int32(3)
							v1129 = F_memmove(m, v1127-v1124, v1106, v1124)
							mBase = m.M
							v1132 = F___memset(m, v1106, int32(48), v1110-v1124)
							mBase = m.M
							v1135 = v1127
						}
					} else {
						v1120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1099<<(uint(int32(1))%32))+uint32(_consts[1073]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v1106))) = uint16(v1120)
						v1135 = v1075 + int32(3)
					}
					v1136 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v1135))) = uint8(v1136)
					v1139 = v1135 + int32(1)
					v1140 = int32(2)
					if base.Ui32(int32(99)) < base.Ui32(v1104) {
						v1154 = F_pg_ultoa_n(m, v1104, v1139)
						mBase = m.M
						if v1140 <= v1154 {
							v1165 = v1139 + v1154
						} else {
							v1157 = v1135 + int32(3)
							v1159 = F_memmove(m, v1157-v1154, v1139, v1154)
							mBase = m.M
							v1162 = F___memset(m, v1139, int32(48), v1140-v1154)
							mBase = m.M
							v1165 = v1157
						}
					} else {
						v1150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1104<<(uint(int32(1))%32))+uint32(_consts[1073]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v1139))) = uint16(v1150)
						v1165 = v1135 + int32(3)
					}
					v1194 = v1109
					v1195 = v1165
					v1196 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v1195))) = uint8(v1196)
					v1199 = v1195 + int32(1)
					v1200 = int32(2)
					if base.Ui32(int32(99)) < base.Ui32(v1194) {
						v1214 = F_pg_ultoa_n(m, v1194, v1199)
						mBase = m.M
						if v1200 <= v1214 {
							v1225 = v1199 + v1214
						} else {
							v1217 = v1195 + int32(3)
							v1219 = F_memmove(m, v1217-v1214, v1199, v1214)
							mBase = m.M
							v1222 = F___memset(m, v1199, int32(48), v1200-v1214)
							mBase = m.M
							v1225 = v1217
						}
					} else {
						v1210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1194<<(uint(int32(1))%32))+uint32(_consts[1073]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v1199))) = uint16(v1210)
						v1225 = v1195 + int32(3)
					}
					v1662 = v1225
				} else {
					v1166 = int32(2)
					if base.Ui32(int32(99)) < base.Ui32(v1099) {
						v1180 = F_pg_ultoa_n(m, v1099, v1106)
						mBase = m.M
						if v1166 <= v1180 {
							v1191 = v1106 + v1180
						} else {
							v1183 = v1075 + int32(3)
							v1185 = F_memmove(m, v1183-v1180, v1106, v1180)
							mBase = m.M
							v1188 = F___memset(m, v1106, int32(48), v1166-v1180)
							mBase = m.M
							v1191 = v1183
						}
					} else {
						v1176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1099<<(uint(int32(1))%32))+uint32(_consts[1073]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v1106))) = uint16(v1176)
						v1191 = v1075 + int32(3)
					}
					if v1104 == int32(0) {
						v1662 = v1191
					} else {
						v1194 = v1104
						v1195 = v1191
						v1196 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v1195))) = uint8(v1196)
						v1199 = v1195 + int32(1)
						v1200 = int32(2)
						if base.Ui32(int32(99)) < base.Ui32(v1194) {
							v1214 = F_pg_ultoa_n(m, v1194, v1199)
							mBase = m.M
							if v1200 <= v1214 {
								v1225 = v1199 + v1214
							} else {
								v1217 = v1195 + int32(3)
								v1219 = F_memmove(m, v1217-v1214, v1199, v1214)
								mBase = m.M
								v1222 = F___memset(m, v1199, int32(48), v1200-v1214)
								mBase = m.M
								v1225 = v1217
							}
						} else {
							v1210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1194<<(uint(int32(1))%32))+uint32(_consts[1073]))))
							*(*uint16)(unsafe.Add(mBase, uint32(v1199))) = uint16(v1210)
							v1225 = v1195 + int32(3)
						}
						v1662 = v1225
					}
				}
				v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v1663 <= int32(0) {
					v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
					*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
					v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
					*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
					v1674 = v1662 + int32(3)
				} else {
					v1674 = v1662
				}
				v1675 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
				m.G0 = v14 + int32(48)
				return
			}
		}
	default:
		v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v1227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v1230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v1232 = base.B2i32(int32(2) < v1230)
		if int32(2) < v1230 {
			v1233 = int32(4800)
		} else {
			v1233 = int32(4799)
		}
		v1234 = v1227 + v1233
		v1239 = base.I32_div_s(v1234, int32(4))
		v1242 = base.I32_div_s(v1234, int32(-100))
		v1245 = base.I32_div_s(v1234, int32(400))
		if int32(2) < v1230 {
			v1249 = int32(1)
		} else {
			v1249 = int32(13)
		}
		v1254 = base.I32_div_s((v1249+v1230)*int32(7834), int32(256))
		v1258 = int32(7)
		v1259 = base.I32_rem_s(v1226+v1234*int32(365)+v1239+v1242+v1245+v1254-int32(32166), v1258)
		if v1259 < int32(0) {
			v1264 = v1259 + v1258
		} else {
			v1264 = v1259
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1264
		v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1264<<(uint(int32(2))%32))+uint32(_consts[1074])))
		v1271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1270))))
		*(*uint16)(unsafe.Add(mBase, uint32(l6))) = uint16(v1271)
		v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1270)+2)))
		v1274 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)) = uint8(v1274)
		*(*uint8)(unsafe.Add(mBase, uint32(l6)+2)) = uint8(v1273)
		v1278 = l6 + int32(4)
		v1280 = *(*int32)(unsafe.Add(mBase, _consts[1061]))
		if v1280 == int32(1) {
			v1283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v1284 = int32(2)
			if base.Ui32(int32(99)) < base.Ui32(v1283) {
				v1298 = F_pg_ultoa_n(m, v1283, v1278)
				mBase = m.M
				if v1284 <= v1298 {
					v1309 = v1278 + v1298
				} else {
					v1301 = l6 + int32(6)
					v1303 = F_memmove(m, v1301-v1298, v1278, v1298)
					mBase = m.M
					v1306 = F___memset(m, v1278, int32(48), v1284-v1298)
					mBase = m.M
					v1309 = v1301
				}
			} else {
				v1294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1283<<(uint(int32(1))%32))+uint32(_consts[1073]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v1278))) = uint16(v1294)
				v1309 = l6 + int32(6)
			}
			v1310 = int32(32)
			*(*uint8)(unsafe.Add(mBase, uint32(v1309))) = uint8(v1310)
			v1312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1312<<(uint(int32(2))%32))+uint32(_consts[1075])))
			v1318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1317))))
			*(*uint16)(unsafe.Add(mBase, uint32(v1309)+1)) = uint16(v1318)
			v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1317)+2)))
			*(*uint8)(unsafe.Add(mBase, uint32(v1309)+3)) = uint8(v1320)
			v1368 = v1309 + int32(4)
		} else {
			v1324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v1325 = int32(2)
			v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1324<<(uint(v1325)%32))+uint32(_consts[1075])))
			v1330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1329))))
			*(*uint16)(unsafe.Add(mBase, uint32(v1278))) = uint16(v1330)
			v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+2)))
			*(*uint8)(unsafe.Add(mBase, uint32(v1278)+2)) = uint8(v1332)
			v1334 = int32(32)
			*(*uint8)(unsafe.Add(mBase, uint32(l6)+7)) = uint8(v1334)
			v1337 = l6 + int32(8)
			v1338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if base.Ui32(int32(99)) < base.Ui32(v1338) {
				v1353 = F_pg_ultoa_n(m, v1338, v1337)
				mBase = m.M
				if v1325 <= v1353 {
					v1364 = v1337 + v1353
				} else {
					v1356 = l6 + int32(10)
					v1358 = F_memmove(m, v1356-v1353, v1337, v1353)
					mBase = m.M
					v1361 = F___memset(m, v1337, int32(48), v1325-v1353)
					mBase = m.M
					v1364 = v1356
				}
			} else {
				v1349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1338<<(uint(int32(1))%32))+uint32(_consts[1073]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v1337))) = uint16(v1349)
				v1364 = l6 + int32(10)
			}
			v1368 = v1364
		}
		v1369 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v1368))) = uint8(v1369)
		v1372 = v1368 + int32(1)
		v1373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v1374 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v1373) {
			v1388 = F_pg_ultoa_n(m, v1373, v1372)
			mBase = m.M
			if v1374 <= v1388 {
				v1399 = v1372 + v1388
			} else {
				v1391 = v1368 + int32(3)
				v1393 = F_memmove(m, v1391-v1388, v1372, v1388)
				mBase = m.M
				v1396 = F___memset(m, v1372, int32(48), v1374-v1388)
				mBase = m.M
				v1399 = v1391
			}
		} else {
			v1384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1373<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v1372))) = uint16(v1384)
			v1399 = v1368 + int32(3)
		}
		v1400 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v1399))) = uint8(v1400)
		v1403 = v1399 + int32(1)
		v1404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v1405 = int32(2)
		if base.Ui32(int32(99)) < base.Ui32(v1404) {
			v1419 = F_pg_ultoa_n(m, v1404, v1403)
			mBase = m.M
			if v1405 <= v1419 {
				v1430 = v1403 + v1419
			} else {
				v1422 = v1399 + int32(3)
				v1424 = F_memmove(m, v1422-v1419, v1403, v1419)
				mBase = m.M
				v1427 = F___memset(m, v1403, int32(48), v1405-v1419)
				mBase = m.M
				v1430 = v1422
			}
		} else {
			v1415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1404<<(uint(int32(1))%32))+uint32(_consts[1073]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v1403))) = uint16(v1415)
			v1430 = v1399 + int32(3)
		}
		v1431 = int32(58)
		*(*uint8)(unsafe.Add(mBase, uint32(v1430))) = uint8(v1431)
		v1435 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v1441 = v1435 >> (uint(int32(31)) % 32)
		v1445 = F_pg_ultostr_zeropad(m, v1430+int32(1), v1435^v1441-v1441, int32(2))
		mBase = m.M
		if l1 == int32(0) {
			v1553 = v1445
		} else {
			v1450 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v1445))) = uint8(v1450)
			v1453 = l1 >> (uint(int32(31)) % 32)
			v1455 = l1 ^ v1453 - v1453
			v1457 = base.I32_div_s(v1455, int32(10))
			v1460 = v1457*int32(-10) + v1455
			if v1460 != 0 {
				v1462 = v1460 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1445)+6)) = uint8(v1462)
				v1468 = v1445 + int32(7)
			} else {
				v1468 = v1445 + int32(6)
			}
			v1470 = base.I32_div_s(v1455, int32(100))
			v1473 = v1470*int32(-10) + v1457
			v1474 = v1460 | v1473
			if v1474 == int32(0) {
				v1482 = v1445 + int32(5)
			} else {
				v1480 = v1473 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1445)+5)) = uint8(v1480)
				v1482 = v1468
			}
			v1484 = base.I32_div_s(v1455, int32(1000))
			v1487 = v1484*int32(-10) + v1470
			v1488 = v1474 | v1487
			if v1488 == int32(0) {
				v1496 = v1445 + int32(4)
			} else {
				v1494 = v1487 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1445)+4)) = uint8(v1494)
				v1496 = v1482
			}
			v1498 = base.I32_div_s(v1455, int32(10000))
			v1501 = v1498*int32(-10) + v1484
			v1502 = v1488 | v1501
			if v1502 == int32(0) {
				v1510 = v1445 + int32(3)
			} else {
				v1508 = v1501 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1445)+3)) = uint8(v1508)
				v1510 = v1496
			}
			v1512 = base.I32_div_s(v1455, int32(100000))
			v1515 = v1512*int32(-10) + v1498
			v1516 = v1502 | v1515
			if v1516 == int32(0) {
				v1524 = v1445 + int32(2)
			} else {
				v1522 = v1515 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1445)+2)) = uint8(v1522)
				v1524 = v1510
			}
			v1526 = base.I32_div_s(v1455, int32(1000000))
			v1529 = v1526*int32(-10) + v1512
			if v1516|v1529 == int32(0) {
				v1538 = v1445 + int32(1)
			} else {
				v1536 = v1529 + int32(48)
				*(*uint8)(unsafe.Add(mBase, uint32(v1445)+1)) = uint8(v1536)
				v1538 = v1524
			}
			if base.Ui32(int32(19)) <= base.Ui32(v1512+int32(9)) {
				v1545 = F_pg_ultostr(m, v1445+int32(1), v1455)
				mBase = m.M
				v1546 = v1545
			} else {
				v1546 = v1538
			}
			v1553 = v1546
		}
		v1554 = int32(32)
		*(*uint8)(unsafe.Add(mBase, uint32(v1553))) = uint8(v1554)
		v1556 = int32(1)
		v1557 = v1553 + v1556
		v1558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if int32(0) < v1558 {
			v1563 = v1558
		} else {
			v1563 = v1556 - v1558
		}
		v1564 = int32(4)
		if base.Ui32(int32(99)) < base.Ui32(v1563) {
		} else {
		}
		v1578 = F_pg_ultoa_n(m, v1563, v1557)
		mBase = m.M
		if v1564 <= v1578 {
			v1589 = v1557 + v1578
		} else {
			v1581 = v1553 + int32(5)
			v1583 = F_memmove(m, v1581-v1578, v1557, v1578)
			mBase = m.M
			v1586 = F___memset(m, v1557, int32(48), v1564-v1578)
			mBase = m.M
			v1589 = v1581
		}
		if v19 == int32(0) {
			v1662 = v1589
			v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v1663 <= int32(0) {
				v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
				*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
				v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
				*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
				v1674 = v1662 + int32(3)
			} else {
				v1674 = v1662
			}
			v1675 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
			m.G0 = v14 + int32(48)
			return
		} else {
			if l4 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(10)
				v1596 = F_pg_sprintf(m, v1589, int32(175134), v14)
				mBase = m.M
				v1597 = m.ExcPending
				if v1597 != 0 {
					return
				} else {
					v1598 = F_strlen(m, v1589)
					mBase = m.M
					v1662 = v1598 + v1589
					v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v1663 <= int32(0) {
						v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
						*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
						v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
						*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
						v1674 = v1662 + int32(3)
					} else {
						v1674 = v1662
					}
					v1675 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
					m.G0 = v14 + int32(48)
					return
				}
			} else {
				v1600 = int32(32)
				*(*uint8)(unsafe.Add(mBase, uint32(v1589))) = uint8(v1600)
				if l3 <= int32(0) {
					v1611 = int32(43)
				} else {
					v1611 = int32(45)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v1589+int32(1)))) = uint8(v1611)
				v1614 = l3 >> (uint(int32(31)) % 32)
				v1616 = l3 ^ v1614 - v1614
				v1618 = base.I32_div_s(v1616, int32(3600))
				v1619 = int32(-60)
				v1622 = base.I32_div_s(v1616, int32(60))
				v1623 = v1618*v1619 + v1622
				v1625 = v1589 + int32(2)
				v1628 = v1622*v1619 + v1616
				if v1628 != 0 {
					v1629 = int32(2)
					v1630 = F_pg_ultostr_zeropad(m, v1625, v1618, v1629)
					mBase = m.M
					v1631 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v1630))) = uint8(v1631)
					v1636 = F_pg_ultostr_zeropad(m, v1630+int32(1), v1623, v1629)
					mBase = m.M
					v1643 = v1636
					v1644 = v1628
					v1645 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v1643))) = uint8(v1645)
					v1650 = F_pg_ultostr_zeropad(m, v1643+int32(1), v1644, int32(2))
					mBase = m.M
					v1651 = v1650
				} else {
					v1638 = F_pg_ultostr_zeropad(m, v1625, v1618, int32(2))
					mBase = m.M
					if l5 == int32(4) {
						v1643 = v1638
						v1644 = v1623
						v1645 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v1643))) = uint8(v1645)
						v1650 = F_pg_ultostr_zeropad(m, v1643+int32(1), v1644, int32(2))
						mBase = m.M
						v1651 = v1650
					} else {
						if v1623 == int32(0) {
							v1651 = v1638
						} else {
							v1643 = v1638
							v1644 = v1623
							v1645 = int32(58)
							*(*uint8)(unsafe.Add(mBase, uint32(v1643))) = uint8(v1645)
							v1650 = F_pg_ultostr_zeropad(m, v1643+int32(1), v1644, int32(2))
							mBase = m.M
							v1651 = v1650
						}
					}
				}
				v1662 = v1651
				v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v1663 <= int32(0) {
					v1667 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
					*(*uint8)(unsafe.Add(mBase, uint32(v1662)+2)) = uint8(v1667)
					v1670 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
					*(*uint16)(unsafe.Add(mBase, uint32(v1662))) = uint16(v1670)
					v1674 = v1662 + int32(3)
				} else {
					v1674 = v1662
				}
				v1675 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v1674))) = uint8(v1675)
				m.G0 = v14 + int32(48)
				return
			}
		}
	}
}
func F_date_gt_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return base.B2i32(int32(0) < v65)
}
func F_date_le_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(-2147483648) {
		v17 = int64(-9223372036854775807 - 1)
		return base.B2i32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4) <= int32(0))
	} else {
		if v6 == int32(2147483647) {
			v17 = int64(9223372036854775807)
			return base.B2i32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4) <= int32(0))
		} else {
			if int32(106751982) < v6 {
				return base.B2i32(v4 == int64(9223372036854775807))
			} else {
				v17 = base.I64_extend_i32_s(v6) * int64(86400000000)
				return base.B2i32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4) <= int32(0))
			}
		}
	}
}
func F_date_mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(2)) <= base.Ui32(v3-int32(2147483647)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if base.Ui32(int32(1)) < base.Ui32(v8-int32(2147483647)) {
			return v3 - v8
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(160424), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(498052), int32(560), int32(319414))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(160424), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498052), int32(560), int32(319414))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
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
func F_date_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
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
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	v3 = m.G0
	v5 = v3 - int32(176)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v7-int32(2147483647)) <= base.Ui32(int32(1)) {
		if v7 == int32(-2147483648) {
			v15 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1056])))
			*(*uint16)(unsafe.Add(mBase, uint32(v5)+8)) = uint16(v15)
			v18 = *(*int64)(unsafe.Add(mBase, _consts[1057]))
			*(*int64)(unsafe.Add(mBase, uint32(v5))) = v18
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1058])))
			*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)) = uint8(v21)
			v24 = *(*int64)(unsafe.Add(mBase, _consts[1059]))
			*(*int64)(unsafe.Add(mBase, uint32(v5))) = v24
		}
	} else {
		v37 = v7 + int32(2483589)
		v38 = int32(146097)
		v39 = base.I32_div_u_s(v37, v38)
		v40 = int32(3)
		v46 = int32(2)
		v51 = base.I32_div_u_s((v39*int32(1073595727)+v37)<<(uint(v46)%32)|v40, v38)
		v54 = v7 + int32(2451545) + v39*v40 + v51 + int32(32104)
		v55 = int32(1461)
		v56 = base.I32_div_u_s(v54, v55)
		v59 = v56*int32(-1461) + v54
		v61 = v59 << (uint(v46) % 32)
		if base.Ui32(v55) <= base.Ui32(v61) {
			v67 = base.I32_rem_u_s(v59+int32(305), int32(365))
			v72 = v67
		} else {
			v71 = base.I32_rem_u_s(v59+int32(306), int32(366))
			v72 = v71
		}
		v74 = base.I32_div_u_s(v61, int32(1461))
		*(*int32)(unsafe.Add(mBase, uint32(v5+int32(152)))) = v74 + v56<<(uint(int32(2))%32) - int32(4800)
		v82 = v72 + int32(123)
		v86 = int32(base.Ui32(v82*int32(2141)) >> (uint(int32(16)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(v5+int32(144)))) = v82 - int32(base.Ui32(v86*int32(7834))>>(uint(int32(8))%32))
		v96 = base.I32_rem_u_s(v86+int32(10), int32(12))
		*(*int32)(unsafe.Add(mBase, uint32(v5+int32(148)))) = v96 + int32(1)
		v101 = v5 + int32(132)
		v103 = *(*int32)(unsafe.Add(mBase, _consts[1060]))
		switch v103 - int32(1) {
		case 0, 3:
			v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
			if int32(0) < v106 {
				v111 = v106
			} else {
				v111 = int32(1) - v106
			}
			v113 = F_pg_ultostr_zeropad(m, v5, v111, int32(4))
			mBase = m.M
			v114 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v113))) = uint8(v114)
			v116 = int32(1)
			v118 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
			v119 = int32(2)
			v120 = F_pg_ultostr_zeropad(m, v113+v116, v118, v119)
			mBase = m.M
			*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v114)
			v125 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
			v127 = F_pg_ultostr_zeropad(m, v120+v116, v125, v119)
			mBase = m.M
			v220 = v127
		case 1:
			v131 = *(*int32)(unsafe.Add(mBase, _consts[1061]))
			v133 = base.B2i32(v131 == int32(1))
			if v131 == int32(1) {
				v134 = int32(12)
			} else {
				v134 = int32(16)
			}
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v101+v134)))
			v138 = F_pg_ultostr_zeropad(m, v5, v136, int32(2))
			mBase = m.M
			v139 = int32(47)
			*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v139)
			if v131 == int32(1) {
				v145 = int32(16)
			} else {
				v145 = int32(12)
			}
			v147 = *(*int32)(unsafe.Add(mBase, uint32(v101+v145)))
			v149 = F_pg_ultostr_zeropad(m, v138+int32(1), v147, int32(2))
			mBase = m.M
			v150 = int32(47)
			*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v150)
			v152 = int32(1)
			v154 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
			if int32(0) < v154 {
				v159 = v154
			} else {
				v159 = v152 - v154
			}
			v161 = F_pg_ultostr_zeropad(m, v149+v152, v159, int32(4))
			mBase = m.M
			v220 = v161
		case 2:
			v162 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
			v163 = int32(2)
			v164 = F_pg_ultostr_zeropad(m, v5, v162, v163)
			mBase = m.M
			v165 = int32(46)
			*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v165)
			v167 = int32(1)
			v169 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
			v171 = F_pg_ultostr_zeropad(m, v164+v167, v169, v163)
			mBase = m.M
			*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v165)
			v176 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
			if int32(0) < v176 {
				v181 = v176
			} else {
				v181 = v167 - v176
			}
			v183 = F_pg_ultostr_zeropad(m, v171+v167, v181, int32(4))
			mBase = m.M
			v220 = v183
		default:
			v187 = *(*int32)(unsafe.Add(mBase, _consts[1061]))
			v189 = base.B2i32(v187 == int32(1))
			if v187 == int32(1) {
				v190 = int32(12)
			} else {
				v190 = int32(16)
			}
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v101+v190)))
			v194 = F_pg_ultostr_zeropad(m, v5, v192, int32(2))
			mBase = m.M
			v195 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v195)
			if v187 == int32(1) {
				v201 = int32(16)
			} else {
				v201 = int32(12)
			}
			v203 = *(*int32)(unsafe.Add(mBase, uint32(v101+v201)))
			v205 = F_pg_ultostr_zeropad(m, v194+int32(1), v203, int32(2))
			mBase = m.M
			v206 = int32(45)
			*(*uint8)(unsafe.Add(mBase, uint32(v205))) = uint8(v206)
			v208 = int32(1)
			v210 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
			if int32(0) < v210 {
				v215 = v210
			} else {
				v215 = v208 - v210
			}
			v217 = F_pg_ultostr_zeropad(m, v205+v208, v215, int32(4))
			mBase = m.M
			v220 = v217
		}
		v221 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
		if v221 <= int32(0) {
			v225 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1062])))
			*(*uint8)(unsafe.Add(mBase, uint32(v220)+2)) = uint8(v225)
			v228 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1063])))
			*(*uint16)(unsafe.Add(mBase, uint32(v220))) = uint16(v228)
			v232 = v220 + int32(3)
		} else {
			v232 = v220
		}
		v233 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v232))) = uint8(v233)
	}
	v235 = F_pstrdup(m, v5)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(176)
		return v235
	}
}
func F_date_pl_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v7 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v19 = F_Int64GetDatum(m, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_DirectFunctionCall2Coll(m, int32(1283), int32(0), v19, v3)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v23
			}
		}
	} else {
		if v7 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v19 = F_Int64GetDatum(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_DirectFunctionCall2Coll(m, int32(1283), int32(0), v19, v3)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			}
		} else {
			if int32(106751983) <= v7 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(236993), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498052), int32(658), int32(31203))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v18 = base.I64_extend_i32_s(v7) * int64(86400000000)
				v19 = F_Int64GetDatum(m, v18)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = F_DirectFunctionCall2Coll(m, int32(1283), int32(0), v19, v3)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v23
					}
				}
			}
		}
	}
}
func F_date_skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(1280)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(1281)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9223372034707292160)
	return int32(0)
}
func F_date_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 < v4 {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_date_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 == int32(-2147483648) {
		v6 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		if v2 == int32(2147483647) {
			v14 = F_Int64GetDatum(m, int64(9223372036854775807))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		} else {
			if v2 < int32(106751983) {
				v22 = F_Int64GetDatum(m, base.I64_extend_i32_s(v2)*int64(86400000000))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return v22
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(236993), int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498052), int32(658), int32(31203))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
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
	}
}
func F_date_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v133 int64
	_ = v133
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if v4 == int32(-2147483648) {
		v155 = int64(-9223372036854775807 - 1)
		m.G0 = v7 + int32(48)
		v159 = F_Int64GetDatum(m, v155)
		mBase = m.M
		v160 = m.ExcPending
		if v160 != 0 {
			return int32(0)
		} else {
			return v159
		}
	} else {
		if v4 != int32(2147483647) {
			if int32(106751983) <= v4 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(236993), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498052), int32(721), int32(31173))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v45 = v4 + int32(2483589)
				v46 = int32(146097)
				v47 = base.I32_div_u_s(v45, v46)
				v48 = int32(3)
				v54 = int32(2)
				v59 = base.I32_div_u_s((v47*int32(1073595727)+v45)<<(uint(v54)%32)|v48, v46)
				v62 = v4 + int32(2451545) + v47*v48 + v59 + int32(32104)
				v63 = int32(1461)
				v64 = base.I32_div_u_s(v62, v63)
				v67 = v64*int32(-1461) + v62
				v69 = v67 << (uint(v54) % 32)
				if base.Ui32(v63) <= base.Ui32(v69) {
					v75 = base.I32_rem_u_s(v67+int32(305), int32(365))
					v80 = v75
				} else {
					v79 = base.I32_rem_u_s(v67+int32(306), int32(366))
					v80 = v79
				}
				v82 = base.I32_div_u_s(v69, int32(1461))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v82 + v64<<(uint(int32(2))%32) - int32(4800)
				v90 = v80 + int32(123)
				v94 = int32(base.Ui32(v90*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v90 - int32(base.Ui32(v94*int32(7834))>>(uint(int32(8))%32))
				v104 = base.I32_rem_u_s(v94+int32(10), int32(12))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(20)))) = v104 + int32(1)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v115 = *(*int32)(unsafe.Add(mBase, _consts[1064]))
				v117 = m.G0
				v118 = int32(16)
				v119 = v117 - v118
				m.G0 = v119
				v123 = F_DetermineTimeZoneOffsetInternal(m, v7+int32(4), v115, v119+int32(8))
				mBase = m.M
				m.G0 = v119 + v118
				v133 = base.I64_extend_i32_s(v123)*int64(1000000) + base.I64_extend_i32_s(v4)*int64(86400000000)
				if base.Ui64(v133+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v155 = v133
					m.G0 = v7 + int32(48)
					v159 = F_Int64GetDatum(m, v155)
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return int32(0)
					} else {
						return v159
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(236993), int32(0))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(498052), int32(757), int32(31173))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
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
		} else {
			v155 = int64(9223372036854775807)
			m.G0 = v7 + int32(48)
			v159 = F_Int64GetDatum(m, v155)
			mBase = m.M
			v160 = m.ExcPending
			if v160 != 0 {
				return int32(0)
			} else {
				return v159
			}
		}
	}
}
