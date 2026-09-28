package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_LCS_asString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = l0 - int32(1)
	if base.Ui32(v3) <= base.Ui32(int32(3)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_LCS_asString[0])))
		v10 = v8
	} else {
		v10 = int32(_a_F_LCS_asString_0)
	}
	return v10
}
func F_LargeObjectExists(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v19 = F_table_open(m, int32(2995), int32(1))
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = int32(1)
			v25 = F_systable_beginscan(m, v19, int32(2996), v22, int32(0), v22, v7)
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_systable_getnext(m, v25)
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_systable_endscan(m, v25)
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_relation_close(m, v19, int32(1))
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 - int32(-64)
							return base.B2i32(v27 != int32(0))
						}
					}
				}
			}
		}
	}
}
func F_LockAcquireExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int64
	_ = v183
	var v188 int32
	_ = v188
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v396 int64
	_ = v396
	var v400 int32
	_ = v400
	var v402 int64
	_ = v402
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int64
	_ = v416
	var v418 int32
	_ = v418
	var v420 int64
	_ = v420
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int64
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int64
	_ = v451
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v477 int64
	_ = v477
	var v481 int32
	_ = v481
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int64
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int64
	_ = v547
	var v552 int32
	_ = v552
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v607 int32
	_ = v607
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int64
	_ = v643
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v757 int32
	_ = v757
	var v768 int32
	_ = v768
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int64
	_ = v799
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v835 int64
	_ = v835
	var v842 int32
	_ = v842
	var v844 int64
	_ = v844
	var v853 int64
	_ = v853
	var v858 int32
	_ = v858
	var v860 int64
	_ = v860
	var v869 int64
	_ = v869
	var v872 int64
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int64
	_ = v878
	var v879 int64
	_ = v879
	var v885 int64
	_ = v885
	var v888 int64
	_ = v888
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int64
	_ = v922
	var v927 int32
	_ = v927
	var v929 int64
	_ = v929
	var v933 int64
	_ = v933
	var v937 int64
	_ = v937
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int64
	_ = v971
	var v976 int32
	_ = v976
	var v978 int64
	_ = v978
	var v982 int64
	_ = v982
	var v986 int64
	_ = v986
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int64
	_ = v1020
	var v1030 int32
	_ = v1030
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1100 int32
	_ = v1100
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1143 int32
	_ = v1143
	var v1150 int32
	_ = v1150
	var v1151 int64
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int64
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1194 int32
	_ = v1194
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1284 int32
	_ = v1284
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1367 int32
	_ = v1367
	var v1382 int32
	_ = v1382
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1410 int64
	_ = v1410
	var v1412 int64
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1418 int64
	_ = v1418
	var v1420 int64
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1461 int32
	_ = v1461
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1487 int32
	_ = v1487
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1531 int32
	_ = v1531
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1557 int32
	_ = v1557
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1619 int32
	_ = v1619
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1665 int64
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1725 int32
	_ = v1725
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1758 int32
	_ = v1758
	var v1765 int32
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1889 int32
	_ = v1889
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1930 int32
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2052 int32
	_ = v2052
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2099 int32
	_ = v2099
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2181 int32
	_ = v2181
	var v2186 int32
	_ = v2186
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2242 int32
	_ = v2242
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2277 int32
	_ = v2277
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2290 int32
	_ = v2290
	var v2295 int32
	_ = v2295
	var v2297 int32
	_ = v2297
	var v2360 int32
	_ = v2360
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2406 int32
	_ = v2406
	var v2411 int32
	_ = v2411
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2417 int64
	_ = v2417
	var v2425 int32
	_ = v2425
	var v2428 int32
	_ = v2428
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2444 int32
	_ = v2444
	var v2448 int32
	_ = v2448
	var v2453 int32
	_ = v2453
	var v2457 int32
	_ = v2457
	var v2461 int32
	_ = v2461
	var v2466 int32
	_ = v2466
	var v2471 int32
	_ = v2471
	var v2477 int32
	_ = v2477
	var v2482 int32
	_ = v2482
	var v2486 int32
	_ = v2486
	var v2490 int32
	_ = v2490
	var v2495 int32
	_ = v2495
	var v2527 int32
	_ = v2527
	var v2558 int64
	_ = v2558
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2568 int32
	_ = v2568
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2602 int64
	_ = v2602
	var v2607 int32
	_ = v2607
	var v2640 int32
	_ = v2640
	var v2643 int32
	_ = v2643
	var v2647 int32
	_ = v2647
	var v2654 int32
	_ = v2654
	var v2664 int32
	_ = v2664
	var v2697 int32
	_ = v2697
	var v2699 int32
	_ = v2699
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2706 int32
	_ = v2706
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2716 int32
	_ = v2716
	var v2721 int32
	_ = v2721
	var v2724 int32
	_ = v2724
	var v2726 int32
	_ = v2726
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2733 int64
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2737 int32
	_ = v2737
	var v2745 int32
	_ = v2745
	var v2781 int32
	_ = v2781
	var v2784 int32
	_ = v2784
	var v2785 int32
	_ = v2785
	var v2789 int32
	_ = v2789
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2804 int32
	_ = v2804
	v7 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(192)
	m.G0 = v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v35-int32(3))&int32(255)) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2781 = m.ExcPending
	if v2781 != 0 {
		goto L23
	} else {
		goto L438
	}
L2:
	;
	m.G0 = v33 + int32(192)
	return v2745
L3:
	;
	v2558 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+32)) = v2558 + int64(1)
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
	v2563 = int32(0)
	v2564 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	if v2563 < v2564 {
		goto L414
	} else {
		goto L415
	}
L4:
	;
	F_LWLockRelease(m, v1107)
	mBase = m.M
	v2527 = m.ExcPending
	if v2527 != 0 {
		goto L23
	} else {
		goto L412
	}
L5:
	;
	if l1 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2486 = m.ExcPending
	if v2486 != 0 {
		goto L23
	} else {
		goto L409
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L23
	} else {
		goto L406
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(int32(2))%32))+uint32(_c_F_LockAcquireExtended[0])))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v47 < l1 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[1])))
	if v51 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v33)+168)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v33)+176)) = v77
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[2]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[3]))
	v87 = v33 + int32(168)
	v91 = F_hash_search(m, v85, v87, int32(1), v33+int32(167))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if v61 == int32(0) {
		goto L11
	} else {
		goto L16
	}
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[4]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+308))
	v59 = base.B2i32(v57 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[1])) = uint8(v59)
	v61 = v59
	goto L15
L14:
	;
	v61 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[5])))
	if v65&int32(1) != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v68 != int32(8) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if base.B2i32(base.Ui32(l1) < base.Ui32(int32(4)))|v68 != 0 {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L1
L22:
	;
	goto L11
L23:
	;
	return int32(0)
L24:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+167)))
	if v95 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	if l4 != 0 {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91)+24)) = int64(0)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[6]))
	v102 = F_get_hash_value(m, v101, v87)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L23
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v91)+44))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	if v120 < v119 {
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v104 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+52)) = uint16(v104)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+20)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v91)+48)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v91)+40)) = int64(34359738368)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[7]))
	v116 = F_MemoryContextAlloc(m, v114, int32(128))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+48)) = v116
	goto L25
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
	v125 = F_repalloc(m, v122, v119<<(uint(int32(5))%32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+44)) = v119 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+48)) = v125
	goto L25
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v91
	goto L35
L34:
	;
	goto L35
L35:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v136 = int32(0)
	goto L38
L37:
	;
	v136 = v83
	goto L38
L38:
	;
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	if int64(0) < v137 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91)+32)) = v137 + int64(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
	v144 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	if v144 < v145 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(l1) < base.Ui32(int32(8)) {
		v304 = v7
		goto L63
	} else {
		goto L64
	}
L42:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+53)))
	if v279 != 0 {
		goto L60
	} else {
		goto L61
	}
L43:
	;
	v148 = v144
	goto L46
L44:
	;
	v221 = int32(0)
	goto L45
L45:
	;
	v224 = v143 + v221<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v224)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v136
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+40)) = v228 + int32(1)
	if v136 == int32(0) {
		goto L42
	} else {
		goto L52
	}
L46:
	;
	v180 = v143 + v148<<(uint(int32(4))%32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if v136 == v181 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v221 = v145
	goto L45
L48:
	;
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v180)+8)) = v183 + int64(1)
	goto L42
L49:
	;
	goto L50
L50:
	;
	v188 = v148 + int32(1)
	if v188 != v145 {
		v148 = v188
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+18)))
	if base.Ui32(v235) <= base.Ui32(int32(15)) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L42
L54:
	;
	if v235 != int32(15) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+v235<<(uint(int32(2))%32))+548)) = v91
	goto L59
L58:
	;
	goto L59
L59:
	;
	v245 = v235 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+18)) = uint8(v245)
	goto L56
L60:
	;
	v280 = int32(3)
	goto L62
L61:
	;
	v280 = int32(2)
	goto L62
L62:
	;
	v2745 = v280
	goto L2
L63:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v305 != int32(1) {
		goto L76
	} else {
		goto L77
	}
L64:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v283 != 0 {
		v304 = v7
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[1])))
	if v286 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v296 != 0 {
		v304 = v7
		goto L63
	} else {
		goto L70
	}
L67:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[4]))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+308))
	v294 = base.B2i32(v292 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[1])) = uint8(v294)
	v296 = v294
	goto L69
L68:
	;
	v296 = int32(0)
	goto L69
L69:
	;
	goto L66
L70:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[8]))
	if v298 <= int32(0) {
		v304 = v7
		goto L63
	} else {
		goto L71
	}
L71:
	;
	v301 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L23
	} else {
		goto L72
	}
L72:
	;
	v304 = int32(1)
	goto L63
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v1113
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1113)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+24)) = v1224
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+20))
	v1228 = l1 << (uint(int32(2)) % 32)
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1228+v1229)))
	if v1226&v1231 != 0 {
		goto L219
	} else {
		goto L220
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L23
	} else {
		goto L214
	}
L75:
	;
	F_LWLockRelease(m, v740)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L23
	} else {
		goto L199
	}
L76:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[9]))
	v1107 = v1100 + v133&int32(15)<<(uint(int32(7))%32) + int32(_a_F_LockAcquireExtended_0)
	v1109 = F_LWLockAcquire(m, v1107, int32(0))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L23
	} else {
		goto L181
	}
L77:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v308|base.B2i32(base.Ui32(int32(3)) < base.Ui32(l1)) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v689 != int32(1) {
		goto L76
	} else {
		goto L128
	}
L79:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[10]))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.B2i32(v313 != v314)|base.B2i32(v313 == int32(0)) != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[11]))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v329 = *(*int32)(unsafe.Add(mBase, uint32((v320-int32(1))&(v323*int32(_a_F_LockAcquireExtended_1))<<(uint(int32(2))%32))+uint32(_c_F_LockAcquireExtended[12])))
	if v329 <= int32(15) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	F_LWLockRelease(m, v654+int32(548))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L23
	} else {
		goto L127
	}
L82:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v337 = F_LWLockAcquire(m, v333+int32(548), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L23
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+14)))
	v642 = v640 * int32(24)
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v642)+uint32(_c_F_LockAcquireExtended[14])))
	*(*int64)(unsafe.Add(mBase, uint32(v642)+uint32(_c_F_LockAcquireExtended[14]))) = v643 + int64(1)
	v648 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[15])) = uint8(v648)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[16])) = uint8(v648)
	goto L78
L85:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v340+v133&int32(1023)<<(uint(int32(2))%32))+4))
	if v346 != 0 {
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[11]))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v355 = (v349 - int32(1)) & (v352 * int32(_a_F_LockAcquireExtended_1))
	v356 = int32(4)
	v357 = v355 << (uint(v356) % 32)
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+564))
	v365 = v360 + v355&int32(268435455)<<(uint(int32(3))%32)
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v365)))
	v368 = v349 << (uint(v356) % 32)
	v371 = v368
	v396 = int64(0)
	goto L89
L87:
	;
	v494 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	F_LWLockRelease(m, v494+int32(548))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L23
	} else {
		goto L105
	}
L88:
	;
	v481 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v365))) = v366 | int64(1)<<(uint(base.I64_extend_i32_u(l1+base.I32_wrap_i64(v477)-v481))%64)
	v492 = v481
	goto L87
L89:
	;
	v400 = v357 + base.I32_wrap_i64(v396)
	v402 = v396 * int64(3)
	if int64(base.Ui64(v366)>>(uint(v402)%64))&int64(7) == int64(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	if base.Ui32(v432) < base.Ui32(v368) {
		goto L102
	} else {
		goto L103
	}
L91:
	;
	v416 = v396 | int64(1)
	v418 = v357 + base.I32_wrap_i64(v416)
	v420 = v416 * int64(3)
	if int64(base.Ui64(v366)>>(uint(v420)%64))&int64(7) == int64(0) {
		goto L97
	} else {
		goto L98
	}
L92:
	;
	v414 = v400
	goto L91
L93:
	;
	goto L94
L94:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v359)+568))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v408+v400<<(uint(int32(2))%32))))
	if v412 == v352 {
		v477 = v402
		goto L88
	} else {
		goto L95
	}
L95:
	;
	v414 = v371
	goto L91
L96:
	;
	v434 = v396 + int64(2)
	if v434 != int64(16) {
		v371 = v432
		v396 = v434
		goto L89
	} else {
		goto L101
	}
L97:
	;
	v432 = v418
	goto L96
L98:
	;
	goto L99
L99:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v359)+568))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v426+v418<<(uint(int32(2))%32))))
	if v430 == v352 {
		v477 = v420
		goto L88
	} else {
		goto L100
	}
L100:
	;
	v432 = v414
	goto L96
L101:
	;
	goto L90
L102:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v359)+568))
	v439 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v438+v432<<(uint(v439)%32)))) = v352
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+564))
	v446 = int32(1)
	v450 = v445 + int32(base.Ui32(v432)>>(uint(v446)%32))&int32(2147483640)
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v450)))
	*(*int64)(unsafe.Add(mBase, uint32(v450))) = v451 | int64(1)<<(uint(base.I64_extend_i32_u(l1+v432&int32(15)*int32(3)-v446))%64)
	v465 = v355 << (uint(v439) % 32)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+uint32(_c_F_LockAcquireExtended[12])))
	*(*int32)(unsafe.Add(mBase, uint32(v465)+uint32(_c_F_LockAcquireExtended[12]))) = v466 + v446
	v474 = v446
	goto L104
L103:
	;
	v474 = int32(0)
	goto L104
L104:
	;
	v492 = v474
	goto L87
L105:
	;
	if v492 == int32(0) {
		goto L78
	} else {
		goto L106
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91)+24)) = int64(0)
	v503 = int32(0)
	v504 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+32)) = v504 + int64(1)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	if v503 < v509 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v2745 = int32(1)
	goto L2
L108:
	;
	v512 = v503
	goto L111
L109:
	;
	v585 = int32(0)
	goto L110
L110:
	;
	v588 = v508 + v585<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v588)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v588))) = v136
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+40)) = v592 + int32(1)
	if v136 != 0 {
		goto L117
	} else {
		goto L118
	}
L111:
	;
	v544 = v508 + v512<<(uint(int32(4))%32)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)))
	if v136 == v545 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v585 = v509
	goto L110
L113:
	;
	v547 = *(*int64)(unsafe.Add(mBase, uint32(v544)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v544)+8)) = v547 + int64(1)
	goto L107
L114:
	;
	goto L115
L115:
	;
	v552 = v512 + int32(1)
	if v552 != v509 {
		v512 = v552
		goto L111
	} else {
		goto L116
	}
L116:
	;
	goto L112
L117:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)))
	if base.Ui32(v597) <= base.Ui32(int32(15)) {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	goto L119
L119:
	;
	goto L107
L120:
	;
	goto L119
L121:
	;
	if v597 != int32(15) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	goto L120
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136+v597<<(uint(int32(2))%32))+548)) = v91
	goto L126
L125:
	;
	goto L126
L126:
	;
	v607 = v597 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)) = uint8(v607)
	goto L123
L127:
	;
	goto L78
L128:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v692|base.B2i32(base.Ui32(l1) < base.Ui32(int32(5))) != 0 {
		goto L76
	} else {
		goto L129
	}
L129:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v696 == int32(0) {
		goto L76
	} else {
		goto L130
	}
L130:
	;
	v700 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v703 = base.AtomicRmwXchg32(m, v700, int32(0), int32(1))
	if v703 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	F_s_lock(m, v700, int32(_a_F_LockAcquireExtended_2))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L23
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v707 = int32(_a_F_LockAcquireExtended_3)
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v713 = v708 + v133&int32(1023)<<(uint(int32(2))%32)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	v715 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v713)+4)) = v714 + v715
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+52)) = uint8(v715)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18])) = v91
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v724 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v723))), uint32(v724))
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[19]))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)+16))
	if v729 == v724 {
		goto L76
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[9]))
	v740 = v733 + v133&int32(15)<<(uint(int32(7))%32) + int32(_a_F_LockAcquireExtended_0)
	v742 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[11]))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v748 = (v742 - int32(1)) & (v745 * int32(_a_F_LockAcquireExtended_1))
	v757 = v728
	v768 = v7
	goto L136
L136:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v757)))
	v788 = v785 + v768*int32(768)
	v790 = v788 + int32(548)
	v792 = F_LWLockAcquire(m, v790, int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L23
	} else {
		goto L138
	}
L137:
	;
	goto L76
L138:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v788)+20))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v794 != v795 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	F_LWLockRelease(m, v790)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L23
	} else {
		goto L179
	}
L140:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v788)+564))
	v799 = *(*int64)(unsafe.Add(mBase, uint32(v797+v748<<(uint(int32(3))%32))))
	if v799 == int64(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v803 = v748 & int32(268435455) << (uint(int32(3)) % 32)
	v804 = v797 + v803
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v788)+568))
	v806 = v805 + v748<<(uint(int32(6))%32)
	v835 = int64(0)
	goto L142
L142:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v806+base.I32_wrap_i64(v835)<<(uint(int32(2))%32))))
	if v745 != v842 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v874 = F_LWLockAcquire(m, v740, int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L23
	} else {
		goto L153
	}
L144:
	;
	goto L143
L145:
	;
	v853 = v835 | int64(1)
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v806+base.I32_wrap_i64(v853)<<(uint(int32(2))%32))))
	if v858 == v745 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v844 = *(*int64)(unsafe.Add(mBase, uint32(v804)))
	if int64(base.Ui64(v844)>>(uint(v835*int64(3))%64))&int64(7) == int64(0) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v872 = v835
	goto L144
L148:
	;
	v860 = *(*int64)(unsafe.Add(mBase, uint32(v804)))
	if int64(base.Ui64(v860)>>(uint(v853*int64(3))%64))&int64(7) != int64(0) {
		v872 = v853
		goto L144
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v869 = v835 + int64(2)
	if v869 != int64(16) {
		v835 = v869
		goto L142
	} else {
		goto L152
	}
L151:
	;
	goto L150
L152:
	;
	goto L139
L153:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v788)+564))
	v878 = *(*int64)(unsafe.Add(mBase, uint32(v876+v803)))
	v879 = int64(1)
	v885 = (v872*int64(12884901888) - int64(4294967296)) >> (uint(int64(32)) % 64)
	v888 = v879 << (uint(v885+v879) % 64)
	if v878&v888 != int64(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v893 = F_SetupLockInTable(m, v46, v788, l0, v133, int32(1))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L23
	} else {
		goto L157
	}
L155:
	;
	v933 = v878
	goto L156
L156:
	;
	v937 = int64(1) << (uint(v885+int64(2)) % 64)
	if v933&v937 != int64(0) {
		goto L162
	} else {
		goto L163
	}
L157:
	;
	if v893 == int32(0) {
		goto L75
	} else {
		goto L158
	}
L158:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v893)))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)+128))
	v899 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v897)+128)) = v898 + v899
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v897)+92))
	v904 = v902 + v899
	*(*int32)(unsafe.Add(mBase, uint32(v897)+92)) = v904
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v897)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v897)+16)) = v906 | int32(2)
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v897)+48))
	if v910 == v904 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v897)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v897)+20)) = v912 & int32(-3)
	goto L161
L160:
	;
	goto L161
L161:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v893)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v893)+12)) = v916 | int32(2)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v788)+564))
	v921 = v920 + v803
	v922 = *(*int64)(unsafe.Add(mBase, uint32(v921)))
	*(*int64)(unsafe.Add(mBase, uint32(v921))) = v922 & (v888 ^ int64(-1))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v788)+564))
	v929 = *(*int64)(unsafe.Add(mBase, uint32(v927+v803)))
	v933 = v929
	goto L156
L162:
	;
	v942 = F_SetupLockInTable(m, v46, v788, l0, v133, int32(2))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L23
	} else {
		goto L165
	}
L163:
	;
	v982 = v933
	goto L164
L164:
	;
	v986 = int64(1) << (uint(v885+int64(3)) % 64)
	if v982&v986 != int64(0) {
		goto L170
	} else {
		goto L171
	}
L165:
	;
	if v942 == int32(0) {
		goto L75
	} else {
		goto L166
	}
L166:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v942)))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v946)+128))
	v948 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v946)+128)) = v947 + v948
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v946)+96))
	v953 = v951 + v948
	*(*int32)(unsafe.Add(mBase, uint32(v946)+96)) = v953
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v946)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v946)+16)) = v955 | int32(4)
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v946)+52))
	if v959 == v953 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v946)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v946)+20)) = v961 & int32(-5)
	goto L169
L168:
	;
	goto L169
L169:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v942)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v942)+12)) = v965 | int32(4)
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v788)+564))
	v970 = v969 + v803
	v971 = *(*int64)(unsafe.Add(mBase, uint32(v970)))
	*(*int64)(unsafe.Add(mBase, uint32(v970))) = v971 & (v937 ^ int64(-1))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v788)+564))
	v978 = *(*int64)(unsafe.Add(mBase, uint32(v976+v803)))
	v982 = v978
	goto L164
L170:
	;
	v991 = F_SetupLockInTable(m, v46, v788, l0, v133, int32(3))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L23
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	F_LWLockRelease(m, v740)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L23
	} else {
		goto L178
	}
L173:
	;
	if v991 == int32(0) {
		goto L75
	} else {
		goto L174
	}
L174:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v991)))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v995)+128))
	v997 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v995)+128)) = v996 + v997
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v995)+100))
	v1002 = v1000 + v997
	*(*int32)(unsafe.Add(mBase, uint32(v995)+100)) = v1002
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v995)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v995)+16)) = v1004 | int32(8)
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v995)+56))
	if v1008 == v1002 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v995)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v995)+20)) = v1010 & int32(-9)
	goto L177
L176:
	;
	goto L177
L177:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v991)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v991)+12)) = v1014 | int32(8)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v788)+564))
	v1019 = v1018 + v803
	v1020 = *(*int64)(unsafe.Add(mBase, uint32(v1019)))
	*(*int64)(unsafe.Add(mBase, uint32(v1019))) = v1020 & (v986 ^ int64(-1))
	goto L172
L178:
	;
	goto L139
L179:
	;
	v1064 = v768 + int32(1)
	v1066 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[19]))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1066)+16))
	if base.Ui32(v1064) < base.Ui32(v1067) {
		v757 = v1066
		v768 = v1064
		goto L136
	} else {
		goto L180
	}
L180:
	;
	goto L137
L181:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v1113 = F_SetupLockInTable(m, v46, v1112, l0, v133, l1)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L23
	} else {
		goto L182
	}
L182:
	;
	if v1113 != 0 {
		goto L73
	} else {
		goto L183
	}
L183:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18]))
	if v1116 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+20))
	v1121 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v1124 = base.AtomicRmwXchg32(m, v1121, int32(0), int32(1))
	if v1124 != 0 {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	goto L186
L186:
	;
	F_LWLockRelease(m, v1107)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L23
	} else {
		goto L191
	}
L187:
	;
	F_s_lock(m, v1121, int32(_a_F_LockAcquireExtended_2))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L23
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v1128 = int32(_a_F_LockAcquireExtended_3)
	v1129 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v1132 = v1129 + v1117&int32(1023)<<(uint(int32(2))%32)
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1132)+4)) = v1133 - int32(1)
	v1137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1116)+52)) = uint8(v1137)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18])) = v1137
	v1143 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1143))), uint32(v1137))
	goto L186
L190:
	;
	goto L189
L191:
	;
	v1151 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	if v1151 == int64(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	F_RemoveLocalLock(m, v91)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L23
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	if l4 != 0 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	goto L194
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L198
L197:
	;
	goto L198
L198:
	;
	goto L74
L199:
	;
	F_LWLockRelease(m, v790)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L23
	} else {
		goto L200
	}
L200:
	;
	F_AbortStrongLockAcquire(m)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L23
	} else {
		goto L201
	}
L201:
	;
	v1170 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	if v1170 == int64(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	F_RemoveLocalLock(m, v91)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L23
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	if l4 != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	goto L204
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L208
L207:
	;
	goto L208
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L23
	} else {
		goto L209
	}
L209:
	;
	F_errcode(m, int32(_a_F_LockAcquireExtended_4))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L23
	} else {
		goto L210
	}
L210:
	;
	F_errmsg(m, int32(_a_F_LockAcquireExtended_5), int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L23
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = int32(_a_F_LockAcquireExtended_6)
	F_errhint(m, int32(_a_F_LockAcquireExtended_7), v33+int32(80))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L23
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_8), int32(1083), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L23
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errcode(m, int32(_a_F_LockAcquireExtended_4))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L23
	} else {
		goto L215
	}
L215:
	;
	F_errmsg(m, int32(_a_F_LockAcquireExtended_5), int32(0))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L23
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = int32(_a_F_LockAcquireExtended_6)
	F_errhint(m, int32(_a_F_LockAcquireExtended_7), v33+int32(32))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L23
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_8), int32(1121), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L23
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	v1259 = int32(0)
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v1263 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1263)+404)) = v1265
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1263)+364))
	if v1267 == v1259 {
		v1319 = v1265
		goto L231
	} else {
		goto L232
	}
L220:
	;
	v1233 = F_LockCheckConflicts(m, v46, l1, v1224, v1113)
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L23
	} else {
		goto L221
	}
L221:
	;
	if v1233 != 0 {
		goto L219
	} else {
		goto L222
	}
L222:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+128))
	v1236 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+128)) = v1235 + v1236
	v1239 = v1228 + v1224
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+88))
	v1242 = v1240 + v1236
	*(*int32)(unsafe.Add(mBase, uint32(v1239)+88)) = v1242
	v1245 = v1236 << (uint(l1) % 32)
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+16)) = v1245 | v1246
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+44))
	if v1249 == v1242 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+20)) = v1251 & (v1245 ^ int32(-1))
	goto L225
L224:
	;
	goto L225
L225:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1113)+12)) = v1256 | v1245
	goto L4
L226:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L23
	} else {
		goto L403
	}
L227:
	;
	v1904 = m.G0
	v1906 = v1904 - int32(128)
	m.G0 = v1906
	v1909 = v1906 + int32(112)
	F_initStringInfo(m, v1909)
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L23
	} else {
		goto L336
	}
L228:
	;
	F_LWLockRelease(m, v1107)
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L23
	} else {
		goto L333
	}
L229:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18]))
	if v1592 != 0 {
		goto L276
	} else {
		goto L277
	}
L230:
	;
	switch v1588 - int32(1) {
	case 0:
		goto L228
	case 1:
		goto L229
	default:
		goto L4
	}
L231:
	;
	v1348 = v1261 + int32(32)
	if v1319 == int32(0) {
		v1487 = v1259
		goto L241
	} else {
		goto L242
	}
L232:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+28))
	if v1270 == int32(0) {
		v1319 = v1265
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v1274 = v1261 + int32(24)
	if v1270 == v1274 {
		v1319 = v1265
		goto L231
	} else {
		goto L234
	}
L234:
	;
	v1278 = v1265
	v1284 = v1270
	goto L235
L235:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1284-int32(12))))
	if v1267 == v1308 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1319 = v1314
	goto L231
L237:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1284-int32(8))))
	v1314 = v1312 | v1278
	goto L239
L238:
	;
	v1314 = v1278
	goto L239
L239:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+4))
	if v1315 != v1274 {
		v1278 = v1314
		v1284 = v1315
		goto L235
	} else {
		goto L240
	}
L240:
	;
	goto L236
L241:
	;
	if l3 != 0 {
		goto L266
	} else {
		goto L267
	}
L242:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+40))
	if v1351 == int32(0) {
		v1487 = v1259
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+36))
	if base.B2i32(v1354 == int32(0))|base.B2i32(v1354 == v1348) != 0 {
		v1487 = v1259
		goto L241
	} else {
		goto L244
	}
L244:
	;
	v1367 = v1354
	v1382 = v7
	goto L245
L245:
	;
	if v1267 != 0 {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	v1487 = int32(0)
	goto L241
L247:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+4))
	if v1471 != v1348 {
		v1367 = v1471
		v1382 = v1469
		goto L245
	} else {
		goto L265
	}
L248:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1367-int32(24))))
	if v1267 == v1391 {
		v1469 = v1382
		goto L247
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+12))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1393+v1394<<(uint(int32(2))%32))))
	if v1398&v1319 != 0 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	goto L250
L252:
	;
	v1401 = v1367 - int32(388)
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1393+v1260<<(uint(int32(2))%32))))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+16))
	if v1405&v1406 != 0 {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	goto L254
L254:
	;
	v1469 = int32(1)<<(uint(v1394)%32) | v1382
	goto L247
L255:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[20]))
	v1410 = *(*int64)(unsafe.Add(mBase, uint32(v1261)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1409)+8)) = v1410
	v1412 = *(*int64)(unsafe.Add(mBase, uint32(v1261)))
	*(*int64)(unsafe.Add(mBase, uint32(v1409))) = v1412
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+16)) = v1260
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1263)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+20)) = v1415
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+384))
	v1418 = *(*int64)(unsafe.Add(mBase, uint32(v1417)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1409)+32)) = v1418
	v1420 = *(*int64)(unsafe.Add(mBase, uint32(v1417)))
	*(*int64)(unsafe.Add(mBase, uint32(v1409)+24)) = v1420
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+400))
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+40)) = v1422
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+44)) = v1424
	v1427 = int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[21])) = v1427
	v1588 = v1427
	goto L230
L256:
	;
	goto L257
L257:
	;
	if v1405&v1382 != 0 {
		v1487 = v1401
		goto L241
	} else {
		goto L258
	}
L258:
	;
	v1431 = F_LockCheckConflicts(m, v46, v1260, v1261, v1264)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L23
	} else {
		goto L259
	}
L259:
	;
	if v1431 != 0 {
		v1487 = v1401
		goto L241
	} else {
		goto L260
	}
L260:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+128))
	v1436 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1261)+128)) = v1435 + v1436
	v1441 = v1261 + v1260<<(uint(int32(2))%32)
	v1443 = v1441 + int32(88)
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1443)))
	*(*int32)(unsafe.Add(mBase, uint32(v1443))) = v1444 + v1436
	v1449 = v1436 << (uint(v1260) % 32)
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1261)+16)) = v1449 | v1450
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1443)))
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+44))
	if v1453 == v1454 {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v1588 = int32(0)
	goto L230
L262:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1261)+20)) = v1456 & (v1449 ^ int32(-1))
	goto L264
L263:
	;
	goto L264
L264:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1264)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1264)+12)) = v1461 | v1449
	goto L261
L265:
	;
	goto L246
L266:
	;
	v1557 = int32(2)
	goto L268
L267:
	;
	if v1487 != 0 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v1588 = v1557
	goto L230
L269:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+40))
	v1539 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1261)+40)) = v1538 + v1539
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1261)+20)) = v1542 | v1539<<(uint(v1260)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1536)+416)) = v1539
	*(*int32)(unsafe.Add(mBase, uint32(v1536)+400)) = v1260
	*(*int32)(unsafe.Add(mBase, uint32(v1536)+396)) = v1264
	*(*int32)(unsafe.Add(mBase, uint32(v1536)+384)) = v1261
	*(*int32)(unsafe.Add(mBase, uint32(v1536)+404)) = v1265
	v1557 = v1539
	goto L268
L270:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1487)+388))
	v1507 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v1508 = int32(388)
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+392)) = v1487 + v1508
	*(*int32)(unsafe.Add(mBase, uint32(v1507)+388)) = v1505
	v1513 = v1507 + v1508
	*(*int32)(unsafe.Add(mBase, uint32(v1487)+388)) = v1513
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+388))
	*(*int32)(unsafe.Add(mBase, uint32(v1515)+4)) = v1513
	v1536 = v1507
	goto L269
L271:
	;
	goto L272
L272:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v1520 = v1518 + int32(388)
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1261)+36))
	if v1521 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1261)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1261)+36)) = v1348
	*(*int32)(unsafe.Add(mBase, uint32(v1261)+32)) = v1261 + int32(32)
	goto L275
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1518)+392)) = v1348
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1348)))
	*(*int32)(unsafe.Add(mBase, uint32(v1518)+388)) = v1531
	*(*int32)(unsafe.Add(mBase, uint32(v1531)+4)) = v1520
	*(*int32)(unsafe.Add(mBase, uint32(v1348))) = v1520
	v1536 = v1518
	goto L269
L276:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1592)+20))
	v1597 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v1600 = base.AtomicRmwXchg32(m, v1597, int32(0), int32(1))
	if v1600 != 0 {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	goto L278
L278:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+12))
	if v1625 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L279:
	;
	F_s_lock(m, v1597, int32(_a_F_LockAcquireExtended_2))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L23
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v1604 = int32(_a_F_LockAcquireExtended_3)
	v1605 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v1608 = v1605 + v1593&int32(1023)<<(uint(int32(2))%32)
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1608)+4)) = v1609 - int32(1)
	v1613 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1592)+52)) = uint8(v1613)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18])) = v1613
	v1619 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1619))), uint32(v1613))
	goto L278
L282:
	;
	goto L281
L283:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+4))
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+20))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+4)) = v1630
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1630))) = v1632
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+28))
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1634)+4)) = v1635
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1635))) = v1637
	v1640 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[22]))
	v1646 = F_hash_search_with_hash_value(m, v1640, v1113, v1628<<(uint(int32(4))%32)^v133, int32(2), int32(0))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L23
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1224)+84))
	v1653 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+84)) = v1652 - v1653
	v1658 = v1224 + l1<<(uint(int32(2))%32)
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1658)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1658)+44)) = v1659 - v1653
	F_LWLockRelease(m, v1107)
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L23
	} else {
		goto L288
	}
L286:
	;
	if v1646 == int32(0) {
		goto L226
	} else {
		goto L287
	}
L287:
	;
	goto L285
L288:
	;
	v1665 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	if v1665 == int64(0) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	F_RemoveLocalLock(m, v91)
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L23
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	if l3 == int32(0) {
		goto L227
	} else {
		goto L293
	}
L292:
	;
	goto L291
L293:
	;
	if l5 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = int32(0)
	v1675 = v33 + int32(148)
	F_initStringInfo(m, v1675)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L23
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v1889 = int32(0)
	if l4 == v1889 {
		v2745 = v1889
		goto L2
	} else {
		goto L332
	}
L297:
	;
	v1679 = v33 + int32(132)
	F_initStringInfo(m, v1679)
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L23
	} else {
		goto L298
	}
L298:
	;
	v1683 = v33 + int32(116)
	F_initStringInfo(m, v1683)
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L23
	} else {
		goto L299
	}
L299:
	;
	F_DescribeLockTag(m, v1675, v91)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L23
	} else {
		goto L300
	}
L300:
	;
	v1688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+15)))
	v1689 = int32(2)
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1688<<(uint(v1689)%32))+uint32(_c_F_LockAcquireExtended[0])))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1691)+8))
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1692+l1<<(uint(v1689)%32))))
	v1698 = F_LWLockAcquire(m, v1107, int32(1))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L23
	} else {
		goto L301
	}
L301:
	;
	v1700 = m.G0
	v1702 = v1700 - int32(48)
	m.G0 = v1702
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v1706 = v33 + int32(112)
	v1707 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1706))) = v1707
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1704)+28))
	if v1709 == v1707 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	m.G0 = v1702 + int32(48)
	F_LWLockRelease(m, v1107)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L23
	} else {
		goto L321
	}
L303:
	;
	v1713 = v1704 + int32(24)
	if v1709 == v1713 {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1715 = int32(1)
	v1718 = v1709
	v1719 = v1715
	v1725 = v1715
	goto L305
L305:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1718-int32(16))))
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+12))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1749)+396))
	if v1751 == v1718-int32(20) {
		goto L308
	} else {
		goto L309
	}
L306:
	;
	goto L302
L307:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+4))
	if v1782 != v1713 {
		v1718 = v1782
		v1719 = v1780
		v1725 = v1781
		goto L305
	} else {
		goto L320
	}
L308:
	;
	if v1719 != 0 {
		goto L311
	} else {
		goto L312
	}
L309:
	;
	goto L310
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1702)+32)) = v1750
	if v1725 != 0 {
		goto L316
	} else {
		goto L317
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1702))) = v1750
	F_appendStringInfo(m, v1679, int32(_a_F_LockAcquireExtended_10), v1702)
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L23
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1702)+16)) = v1750
	F_appendStringInfo(m, v1679, int32(_a_F_LockAcquireExtended_11), v1702+int32(16))
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L23
	} else {
		goto L315
	}
L314:
	;
	v1780 = int32(0)
	v1781 = v1725
	goto L307
L315:
	;
	v1780 = int32(0)
	v1781 = v1725
	goto L307
L316:
	;
	v1770 = int32(_a_F_LockAcquireExtended_10)
	goto L318
L317:
	;
	v1770 = int32(_a_F_LockAcquireExtended_11)
	goto L318
L318:
	;
	F_appendStringInfo(m, v1683, v1770, v1702+int32(32))
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L23
	} else {
		goto L319
	}
L319:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1706)))
	*(*int32)(unsafe.Add(mBase, uint32(v1706))) = v1775 + int32(1)
	v1780 = v1719
	v1781 = int32(0)
	goto L307
L320:
	;
	goto L306
L321:
	;
	v1821 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L23
	} else {
		goto L322
	}
L322:
	;
	if v1821 != 0 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+68)) = v1696
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v1824
	v1827 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v1827
	F_errmsg(m, int32(_a_F_LockAcquireExtended_12), v33-int32(-64))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L23
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	F_pfree(m, v1850)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L23
	} else {
		goto L329
	}
L326:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v33)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v1834
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v33)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = v1836
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	F_errdetail_log_plural(m, int32(_a_F_LockAcquireExtended_13), int32(_a_F_LockAcquireExtended_14), v1840, v33+int32(48))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L23
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_8), int32(1233), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L23
	} else {
		goto L328
	}
L328:
	;
	goto L325
L329:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v33)+116))
	F_pfree(m, v1853)
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L23
	} else {
		goto L330
	}
L330:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v33)+132))
	F_pfree(m, v1856)
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L23
	} else {
		goto L331
	}
L331:
	;
	goto L296
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v2745 = v1889
	goto L2
L333:
	;
	v1896 = F_WaitOnLock(m, v91, v136)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L23
	} else {
		goto L334
	}
L334:
	;
	if v1896 != int32(2) {
		goto L3
	} else {
		goto L335
	}
L335:
	;
	goto L227
L336:
	;
	F_initStringInfo(m, v1906+int32(96))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L23
	} else {
		goto L337
	}
L337:
	;
	v1917 = v1906 + int32(80)
	F_initStringInfo(m, v1917)
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L23
	} else {
		goto L338
	}
L338:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[21]))
	if v1921 <= int32(0) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+112))
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+116))
	F_appendBinaryStringInfo(m, v1906+int32(96), v2086, v2087)
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L23
	} else {
		goto L360
	}
L340:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[20]))
	if v1921 == int32(1) {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1930 = int32(20)
	goto L343
L342:
	;
	v1930 = int32(44)
	goto L343
L343:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1925+v1930)))
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	v1934 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1933))) = uint8(v1934)
	*(*int32)(unsafe.Add(mBase, uint32(v1917)+12)) = v1934
	*(*int32)(unsafe.Add(mBase, uint32(v1917)+4)) = v1934
	goto L344
L344:
	;
	F_DescribeLockTag(m, v1917, v1925)
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L23
	} else {
		goto L345
	}
L345:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+20))
	v1943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1925)+15)))
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+16))
	v1945 = int32(2)
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1943<<(uint(v1945)%32))+uint32(_c_F_LockAcquireExtended[0])))
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1947)+8))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1948+v1944<<(uint(v1945)%32))))
	goto L346
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+64)) = v1942
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+68)) = v1952
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+76)) = v1932
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+72)) = v1956
	F_appendStringInfo(m, v1909, int32(_a_F_LockAcquireExtended_15), v1906-int32(-64))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L23
	} else {
		goto L347
	}
L347:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[21]))
	if v1964 < int32(2) {
		goto L339
	} else {
		goto L348
	}
L348:
	;
	v1969 = int32(1)
	v1971 = v1964
	goto L349
L349:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[20]))
	v2002 = v1999 + v1969*int32(24)
	if v1969 < v1971-int32(1) {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	goto L339
L351:
	;
	v2010 = v2002 + int32(44)
	goto L353
L352:
	;
	v2010 = v1999 + int32(20)
	goto L353
L353:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v2010)))
	v2013 = v1906 + int32(80)
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v2013)))
	v2015 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2014))) = uint8(v2015)
	*(*int32)(unsafe.Add(mBase, uint32(v2013)+12)) = v2015
	*(*int32)(unsafe.Add(mBase, uint32(v2013)+4)) = v2015
	goto L354
L354:
	;
	F_DescribeLockTag(m, v2013, v2002)
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L23
	} else {
		goto L355
	}
L355:
	;
	v2024 = v1906 + int32(112)
	F_appendStringInfoChar(m, v2024, int32(10))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L23
	} else {
		goto L356
	}
L356:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2002)+20))
	v2029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2002)+15)))
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v2002)+16))
	v2031 = int32(2)
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v2029<<(uint(v2031)%32))+uint32(_c_F_LockAcquireExtended[0])))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v2033)+8))
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2034+v2030<<(uint(v2031)%32))))
	goto L357
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+48)) = v2028
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+52)) = v2038
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+60)) = v2011
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+56)) = v2042
	F_appendStringInfo(m, v2024, int32(_a_F_LockAcquireExtended_15), v1906+int32(48))
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L23
	} else {
		goto L358
	}
L358:
	;
	v2050 = v1969 + int32(1)
	v2052 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[21]))
	if v2050 < v2052 {
		v1969 = v2050
		v1971 = v2052
		goto L349
	} else {
		goto L359
	}
L359:
	;
	goto L350
L360:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[21]))
	if int32(0) < v2091 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v2099 = int32(0)
	goto L364
L362:
	;
	goto L363
L363:
	;
	v2406 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[24])))
	if v2406 == int32(1) {
		goto L392
	} else {
		goto L393
	}
L364:
	;
	v2125 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[20]))
	F_appendStringInfoChar(m, v1906+int32(96), int32(10))
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L23
	} else {
		goto L366
	}
L365:
	;
	goto L363
L366:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2125+v2099*int32(24))+20))
	v2136 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[25]))
	if int32(0) < v2136 {
		goto L368
	} else {
		goto L369
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+36)) = v2360
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+32)) = v2134
	F_appendStringInfo(m, v1906+int32(96), int32(_a_F_LockAcquireExtended_16), v1906+int32(32))
	mBase = m.M
	v2369 = m.ExcPending
	if v2369 != 0 {
		goto L23
	} else {
		goto L390
	}
L368:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[26]))
	v2143 = v2140
	v2146 = int32(1)
	goto L371
L369:
	;
	goto L370
L370:
	;
	v2360 = int32(_a_F_LockAcquireExtended_17)
	goto L367
L371:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2143)))
	v2173 = int32(0)
	v2175 = int32(_a_F_LockAcquireExtended_18)
	v2176 = base.AtomicRmwOr32(m, v2173, v2175, v2173)
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+4))
	v2181 = base.AtomicRmwOr32(m, v2173, v2175, v2173)
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v2143)))
	if base.B2i32(v2172&int32(1) == v2173)&base.B2i32(v2186 == v2172) == v2173 {
		goto L373
	} else {
		goto L374
	}
L372:
	;
	goto L370
L373:
	;
	goto L376
L374:
	;
	v2242 = v2177
	goto L375
L375:
	;
	if v2242 == v2134 {
		goto L383
	} else {
		goto L384
	}
L376:
	;
	v2222 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[27]))
	if v2222 != 0 {
		goto L378
	} else {
		goto L379
	}
L377:
	;
	v2242 = v2230
	goto L375
L378:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L23
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2143)))
	v2226 = int32(0)
	v2228 = int32(_a_F_LockAcquireExtended_18)
	v2229 = base.AtomicRmwOr32(m, v2226, v2228, v2226)
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+4))
	v2234 = base.AtomicRmwOr32(m, v2226, v2228, v2226)
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v2143)))
	if v2225&int32(1)|base.B2i32(v2237 != v2225) != 0 {
		goto L376
	} else {
		goto L382
	}
L381:
	;
	goto L380
L382:
	;
	goto L377
L383:
	;
	v2272 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+216))
	v2273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2272))))
	if v2273 == int32(0) {
		v2360 = int32(_a_F_LockAcquireExtended_19)
		goto L367
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	v2295 = v2146 + int32(1)
	v2297 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[25]))
	if v2295 <= v2297 {
		v2143 = v2143 + int32(408)
		v2146 = v2295
		goto L371
	} else {
		goto L389
	}
L386:
	;
	v2277 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[28]))
	v2280 = F_pnstrdup(m, v2272, v2277-int32(1))
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L23
	} else {
		goto L387
	}
L387:
	;
	v2282 = F_strlen(m, v2280)
	mBase = m.M
	v2284 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[28]))
	v2287 = F_pg_mbcliplen(m, v2280, v2282, v2284-int32(1))
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L23
	} else {
		goto L388
	}
L388:
	;
	v2290 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2287+v2280))) = uint8(v2290)
	v2360 = v2280
	goto L367
L389:
	;
	goto L372
L390:
	;
	v2371 = v2099 + int32(1)
	v2373 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[21]))
	if v2371 < v2373 {
		v2099 = v2371
		goto L364
	} else {
		goto L391
	}
L391:
	;
	goto L365
L392:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[10]))
	v2414 = F_pgstat_prep_pending_entry(m, int32(1), v2411, int64(0), int32(0))
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L23
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2425 = m.ExcPending
	if v2425 != 0 {
		goto L23
	} else {
		goto L396
	}
L395:
	;
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(v2414)+12))
	v2417 = *(*int64)(unsafe.Add(mBase, uint32(v2416)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v2416)+144)) = v2417 + int64(1)
	goto L394
L396:
	;
	F_errcode(m, int32(16908292))
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L23
	} else {
		goto L397
	}
L397:
	;
	F_errmsg(m, int32(_a_F_LockAcquireExtended_20), int32(0))
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L23
	} else {
		goto L398
	}
L398:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+16)) = v2433
	F_errdetail_internal(m, int32(_a_F_LockAcquireExtended_21), v1906+int32(16))
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L23
	} else {
		goto L399
	}
L399:
	;
	v2440 = *(*int32)(unsafe.Add(mBase, uint32(v1906)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v1906))) = v2440
	F_errdetail_log(m, int32(_a_F_LockAcquireExtended_21), v1906)
	mBase = m.M
	v2444 = m.ExcPending
	if v2444 != 0 {
		goto L23
	} else {
		goto L400
	}
L400:
	;
	F_errhint(m, int32(_a_F_LockAcquireExtended_22), int32(0))
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L23
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_23), int32(1138), int32(_a_F_LockAcquireExtended_24))
	mBase = m.M
	v2453 = m.ExcPending
	if v2453 != 0 {
		goto L23
	} else {
		goto L402
	}
L402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L403:
	;
	F_errmsg_internal(m, int32(_a_F_LockAcquireExtended_25), int32(0))
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L23
	} else {
		goto L404
	}
L404:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_8), int32(1181), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2466 = m.ExcPending
	if v2466 != 0 {
		goto L23
	} else {
		goto L405
	}
L405:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_LockAcquireExtended_26), v33+int32(16))
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L23
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_8), int32(891), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L23
	} else {
		goto L408
	}
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v35
	F_errmsg_internal(m, int32(_a_F_LockAcquireExtended_27), v33)
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L23
	} else {
		goto L410
	}
L410:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_8), int32(888), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2495 = m.ExcPending
	if v2495 != 0 {
		goto L23
	} else {
		goto L411
	}
L411:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L412:
	;
	goto L3
L413:
	;
	v2697 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18])) = v2697
	v2699 = int32(1)
	if v304 == v2697 {
		v2745 = v2699
		goto L2
	} else {
		goto L431
	}
L414:
	;
	v2568 = v2563
	goto L417
L415:
	;
	v2640 = int32(0)
	goto L416
L416:
	;
	v2643 = v2562 + v2640<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v2643)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2643))) = v136
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+40)) = v2647 + int32(1)
	if v136 == int32(0) {
		goto L413
	} else {
		goto L423
	}
L417:
	;
	v2599 = v2562 + v2568<<(uint(int32(4))%32)
	v2600 = *(*int32)(unsafe.Add(mBase, uint32(v2599)))
	if v136 == v2600 {
		goto L419
	} else {
		goto L420
	}
L418:
	;
	v2640 = v2564
	goto L416
L419:
	;
	v2602 = *(*int64)(unsafe.Add(mBase, uint32(v2599)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2599)+8)) = v2602 + int64(1)
	goto L413
L420:
	;
	goto L421
L421:
	;
	v2607 = v2568 + int32(1)
	if v2607 != v2564 {
		v2568 = v2607
		goto L417
	} else {
		goto L422
	}
L422:
	;
	goto L418
L423:
	;
	v2654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+18)))
	if base.Ui32(v2654) <= base.Ui32(int32(15)) {
		goto L425
	} else {
		goto L426
	}
L424:
	;
	goto L413
L425:
	;
	if v2654 != int32(15) {
		goto L428
	} else {
		goto L429
	}
L426:
	;
	goto L427
L427:
	;
	goto L424
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+v2654<<(uint(int32(2))%32))+548)) = v91
	goto L430
L429:
	;
	goto L430
L430:
	;
	v2664 = v2654 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+18)) = uint8(v2664)
	goto L427
L431:
	;
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2704 = m.G0
	v2706 = v2704 - int32(16)
	m.G0 = v2706
	v2708 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v2709 = m.ExcPending
	if v2709 != 0 {
		goto L23
	} else {
		goto L432
	}
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2706)+8)) = v2703
	*(*int32)(unsafe.Add(mBase, uint32(v2706)+4)) = v2702
	*(*int32)(unsafe.Add(mBase, uint32(v2706))) = v2708
	*(*int32)(unsafe.Add(mBase, uint32(v2706)+12)) = int32(1)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2716 = m.ExcPending
	if v2716 != 0 {
		goto L23
	} else {
		goto L433
	}
L433:
	;
	F_XLogRegisterData(m, v2706+int32(12), int32(4))
	mBase = m.M
	v2721 = m.ExcPending
	if v2721 != 0 {
		goto L23
	} else {
		goto L434
	}
L434:
	;
	F_XLogRegisterData(m, v2706, int32(12))
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L23
	} else {
		goto L435
	}
L435:
	;
	v2726 = int32(_a_F_LockAcquireExtended_28)
	v2728 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[29])))
	v2729 = v2728 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[29])) = uint8(v2729)
	goto L436
L436:
	;
	v2733 = F_XLogInsert(m, int32(8), int32(0))
	mBase = m.M
	v2734 = m.ExcPending
	if v2734 != 0 {
		goto L23
	} else {
		goto L437
	}
L437:
	;
	v2735 = int32(_a_F_LockAcquireExtended_29)
	v2737 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[30]))
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[30])) = v2737 | int32(2)
	m.G0 = v2706 + int32(16)
	v2745 = v2699
	goto L2
L438:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L23
	} else {
		goto L439
	}
L439:
	;
	v2785 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v2785+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v2789
	F_errmsg(m, int32(_a_F_LockAcquireExtended_30), v33+int32(96))
	mBase = m.M
	v2795 = m.ExcPending
	if v2795 != 0 {
		goto L23
	} else {
		goto L440
	}
L440:
	;
	F_errhint(m, int32(_a_F_LockAcquireExtended_31), int32(0))
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L23
	} else {
		goto L441
	}
L441:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_8), int32(901), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2804 = m.ExcPending
	if v2804 != 0 {
		goto L23
	} else {
		goto L442
	}
L442:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__ltq_extract_regex(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14225(m, l0, int32(_a_F__ltq_extract_regex_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_lappend_int(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14324(m, l0, l1, int64(4294967775))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_launch_sync_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_launch_sync_worker[0]))
	if v8 <= l1 {
		return
	} else {
		v13 = m.G0
		v14 = int32(16)
		v15 = v13 - v14
		m.G0 = v15
		F_gettimeofday(m, v15)
		mBase = m.M
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
		v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+8)))
		m.G0 = v15 + v14
		v27 = v19 + v18*int64(1000000) - int64(946684800000000)
		v28 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
		if v28 != int64(0) {
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_launch_sync_worker[1]))
			if base.B2i32(base.I64_extend_i32_s(v32)*int64(1000) <= v27-v28) == int32(0) {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l3))) = v27
				v42 = *(*int32)(unsafe.Add(mBase, _c_F_launch_sync_worker[2]))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_launch_sync_worker[3]))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
				v49 = int32(0)
				v51 = F_logicalrep_worker_launch(m, l0, v43, v46, v47, v48, l2, v49, v49)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = v27
			v42 = *(*int32)(unsafe.Add(mBase, _c_F_launch_sync_worker[2]))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
			v45 = *(*int32)(unsafe.Add(mBase, _c_F_launch_sync_worker[3]))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
			v49 = int32(0)
			v51 = F_logicalrep_worker_launch(m, l0, v43, v46, v47, v48, l2, v49, v49)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_lazy_vacuum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 float32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int64
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v302 int64
	_ = v302
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int64
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v433 int32
	_ = v433
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int64
	_ = v466
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v619 int64
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v797 int32
	_ = v797
	var v814 int32
	_ = v814
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v991 int32
	_ = v991
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1223 int32
	_ = v1223
	var v1233 int32
	_ = v1233
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1326 int32
	_ = v1326
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1377 int32
	_ = v1377
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1465 int32
	_ = v1465
	var v1476 int32
	_ = v1476
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1559 int32
	_ = v1559
	var v1594 int32
	_ = v1594
	var v1609 int32
	_ = v1609
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1644 int32
	_ = v1644
	var v1684 int32
	_ = v1684
	var v1716 int32
	_ = v1716
	var v1722 int32
	_ = v1722
	var v1730 int32
	_ = v1730
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1753 int32
	_ = v1753
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1767 int32
	_ = v1767
	var v1772 int32
	_ = v1772
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1807 int64
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1860 int32
	_ = v1860
	var v1870 int32
	_ = v1870
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1884 int32
	_ = v1884
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1907 int32
	_ = v1907
	var v1912 int32
	_ = v1912
	var v1921 int32
	_ = v1921
	var v1930 int32
	_ = v1930
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1951 int32
	_ = v1951
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1990 int32
	_ = v1990
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2019 int64
	_ = v2019
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2120 int32
	_ = v2120
	v2 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(_a_F_lazy_vacuum_0)
	m.G0 = v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+23)))
	if v40 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v2071 + int32(1)
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v2076 = F_TidStoreMemoryUsage(m, v2075)
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L7
	} else {
		goto L278
	}
L2:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v43 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+48))
	v65 = *(*float32)(unsafe.Add(mBase, uint32(v64)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+88)) = int64(34359738368)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v69
	v72 = *(*int64)(unsafe.Add(mBase, _c_F_lazy_vacuum[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+72)) = v72
	v74 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L10
	}
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v46 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if base.Ui32(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_u(v46), float64(0.02)))) <= base.Ui32(v49) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v56 = F_TidStoreMemoryUsage(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	if base.Ui32(int32(33554431)) < base.Ui32(v56) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v60 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+23)) = uint8(v60)
	goto L1
L10:
	;
	if v74 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = int64(2)
	v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+56)) = v78
	goto L14
L12:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v266 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	goto L12
L14:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[2]))
	if v94 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[3])))
	if v98&int32(1) == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v103 = int32(_a_F_lazy_vacuum_1)
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	v106 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v105 + v106
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v109 + v106
	v113 = int32(0)
	v116 = base.AtomicRmwOr32(m, v113, int32(_a_F_lazy_vacuum_2), v113)
	goto L18
L17:
	;
	v243 = int32(0)
	v246 = base.AtomicRmwOr32(m, v243, int32(_a_F_lazy_vacuum_2), v243)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v248 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v247 + v248
	v251 = int32(_a_F_lazy_vacuum_1)
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v253 - v248
	goto L13
L18:
	;
	goto L20
L20:
	;
	goto L21
L21:
	;
	v208 = int32(0)
	v211 = int32(0)
	goto L26
L26:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(88)+v211<<(uint(int32(2))%32))))
	v221 = int32(3)
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v38+int32(48)+v211<<(uint(v221)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v94+int32(232)+v220<<(uint(v221)%32)))) = v227
	v229 = int32(1)
	v232 = v208 + v229
	if v232 != int32(2) {
		v208 = v232
		v211 = v211 + v229
		goto L26
	} else {
		goto L28
	}
L27:
	;
	goto L17
L28:
	;
	goto L27
L29:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v464 = v462 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v464
	v466 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[5]))) = v466
	*(*int64)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[6]))) = v466
	*(*int64)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[7]))) = base.I64_extend_i32_s(v464)
	goto L49
L30:
	;
	v302 = int64(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v411)+16)) = base.F64_convert_i32_s(base.I32_trunc_sat_f32_s(v65))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	v416 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v415)+24)) = uint8(v416)
	F_parallel_vacuum_process_all_indexes(m, v266, v410, v416, l0+int32(184))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L7
	} else {
		goto L45
	}
L33:
	;
	v305 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	v306 = base.B2i32(v305 <= v302)
	if v305 <= v302 {
		v433 = v306
		goto L29
	} else {
		goto L35
	}
L34:
	;
	v433 = v306
	goto L29
L35:
	;
	v309 = base.I32_wrap_i64(v302) << (uint(int32(2)) % 32)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v309+v310)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v313+v309)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v315
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v38)+112)) = base.F64_promote_f32(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = int32(13)
	v321 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+106)) = uint8(v321)
	v323 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+104)) = uint16(v323)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v317
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+120)) = v326
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v315)+48))
	v331 = F_pstrdup(m, v328+int32(4))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v331
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v335 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v335)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(2)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v347 = F_vac_bulkdel_one_index(m, v38+int32(96), v312, v345, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v340
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v334)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v337
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	F_pfree(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v355 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v357+v309))) = v347
	v362 = v302 + int64(1)
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[2]))
	if v365 == v355 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v406 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L7
	} else {
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[3])))
	if v369&int32(1) == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v374 = int32(_a_F_lazy_vacuum_1)
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	v377 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v376 + v377
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v380 + v377
	v384 = int32(0)
	v386 = int32(_a_F_lazy_vacuum_2)
	v387 = base.AtomicRmwOr32(m, v384, v386, v384)
	*(*int64)(unsafe.Add(mBase, uint32(v365+int32(72))+232)) = v362
	v395 = base.AtomicRmwOr32(m, v384, v386, v384)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v396 + v377
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v402 - v377
	goto L40
L43:
	;
	if v406 == int32(0) {
		v302 = v362
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L34
L45:
	;
	v423 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v433 = v423 ^ int32(1)
	goto L29
L47:
	;
	if v433 == int32(0) {
		goto L1
	} else {
		goto L64
	}
L48:
	;
	goto L47
L49:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[2]))
	if v486 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[3])))
	if v490&int32(1) == int32(0) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v495 = int32(_a_F_lazy_vacuum_1)
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	v498 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v497 + v498
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	*(*int32)(unsafe.Add(mBase, uint32(v486))) = v501 + v498
	v505 = int32(0)
	v508 = base.AtomicRmwOr32(m, v505, int32(_a_F_lazy_vacuum_2), v505)
	goto L53
L52:
	;
	v635 = int32(0)
	v638 = base.AtomicRmwOr32(m, v635, int32(_a_F_lazy_vacuum_2), v635)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	v640 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v486))) = v639 + v640
	v643 = int32(_a_F_lazy_vacuum_1)
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v645 - v640
	goto L48
L53:
	;
	goto L55
L55:
	;
	goto L56
L56:
	;
	v600 = int32(0)
	v603 = int32(0)
	goto L61
L61:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(72)+v603<<(uint(int32(2))%32))))
	v613 = int32(3)
	v619 = *(*int64)(unsafe.Add(mBase, uint32(v38+int32(_a_F_lazy_vacuum_3)+v603<<(uint(v613)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v486+int32(232)+v612<<(uint(v613)%32)))) = v619
	v621 = int32(1)
	v624 = v600 + v621
	if v624 != int32(3) {
		v600 = v624
		v603 = v603 + v621
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L52
L63:
	;
	goto L62
L64:
	;
	v660 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v660
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[2]))
	if v666 == v660 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(3)
	v710 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v711 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v711)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v718 = F_palloc0(m, int32(16))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L7
	} else {
		goto L69
	}
L66:
	;
	goto L65
L67:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[3])))
	if v670&int32(1) == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v675 = int32(_a_F_lazy_vacuum_1)
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	v678 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v677 + v678
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v666)))
	*(*int32)(unsafe.Add(mBase, uint32(v666))) = v681 + v678
	v685 = int32(0)
	v687 = int32(_a_F_lazy_vacuum_2)
	v688 = base.AtomicRmwOr32(m, v685, v687, v685)
	*(*int64)(unsafe.Add(mBase, uint32(v666+v685)+232)) = int64(3)
	v696 = base.AtomicRmwOr32(m, v685, v687, v685)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v666)))
	*(*int32)(unsafe.Add(mBase, uint32(v666))) = v697 + v678
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v703 - v678
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v716
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v716)+4))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v716)+8))
	if v722 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v718)+4)) = v769
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v781 = F_read_stream_begin_relation(m, int32(9), v776, v777, int32(0), int32(192), v718, int32(8))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L7
	} else {
		goto L77
	}
L71:
	;
	v724 = F_palloc0(m, int32(120))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L7
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v753 = F_palloc0(m, int32(88))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L7
	} else {
		goto L76
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v724))) = v721
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v721)+4))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)+24))
	v730 = F_dsa_get_address(m, v727, v729)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v724)))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)+48))
	v736 = base.I32_div_s(v734, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+104)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v724)+100)) = v736
	v740 = v724 + int32(4)
	v741 = int32(12)
	v743 = v740 + v736*v741
	*(*int32)(unsafe.Add(mBase, uint32(v743)+4)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v743))) = v729
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v724)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v740+v746*v741)+8)) = int32(0)
	v769 = v724
	goto L70
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v753))) = v721
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v756)+24))
	v760 = base.I32_div_s(v758, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v753)+72)) = v760
	*(*int32)(unsafe.Add(mBase, uint32(v753)+68)) = v760
	v765 = v753 + v760<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v765)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v765)+4)) = v757
	v769 = v753
	goto L70
L77:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L7
	} else {
		goto L78
	}
L78:
	;
	v788 = F_read_stream_next_buffer(m, v781, v38+int32(88))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	if v788 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v797 = v788
	v814 = v2
	goto L83
L81:
	;
	v1990 = v2
	goto L82
L82:
	;
	F_read_stream_end(m, v781)
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L7
	} else {
		goto L265
	}
L83:
	;
	if v797 < int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v1990 = v1961
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v843
	v846 = v38 + int32(96)
	v847 = int32(0)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v38)+88))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v849)+4))
	v851 = int32(*(*int8)(unsafe.Add(mBase, uint32(v850)+1)))
	if v851 != 0 {
		goto L91
	} else {
		goto L92
	}
L86:
	;
	v828 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[8]))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v828+(v797^int32(-1))*int32(56))+16))
	v843 = v834
	goto L85
L87:
	;
	goto L88
L88:
	;
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[9]))
	v837 = int32(56)
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v836+v797*v837-v837)+16))
	v843 = v842
	goto L85
L89:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_visibilitymap_pin(m, v1030, v843, v38+int32(72))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L7
	} else {
		goto L111
	}
L90:
	;
	v864 = v851
	v867 = v847
	v868 = v847
	goto L96
L91:
	;
	if int32(0) < v851 {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v855 = int32(0)
	v856 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v850)+2)))
	if v856 == v855 {
		v1029 = v855
		goto L89
	} else {
		goto L95
	}
L94:
	;
	v1029 = int32(0)
	goto L89
L95:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v846))) = uint16(v856)
	v1029 = int32(1)
	goto L89
L96:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v850+int32(4)+v867<<(uint(int32(2))%32))))
	if v901 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v1029 = v960
	goto L89
L98:
	;
	v905 = v867 << (uint(int32(5)) % 32)
	v907 = v901
	v909 = v868
	goto L101
L99:
	;
	v956 = v864
	v960 = v868
	goto L100
L100:
	;
	v991 = v867 + int32(1)
	if v991 < base.I32_extend8_s(v956) {
		v864 = v956
		v867 = v991
		v868 = v960
		goto L96
	} else {
		goto L110
	}
L101:
	;
	if v907&int32(1) != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850)+1)))
	v956 = v954
	v960 = v949
	goto L100
L103:
	;
	if v909 < int32(2048) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v949 = v909
	goto L105
L105:
	;
	v950 = int32(1)
	v953 = int32(base.Ui32(v907) >> (uint(v950) % 32))
	if v953 != 0 {
		v905 = v905 + v950
		v907 = v953
		v909 = v949
		goto L101
	} else {
		goto L109
	}
L106:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v846+v909<<(uint(int32(1))%32)))) = uint16(v905)
	goto L108
L107:
	;
	goto L108
L108:
	;
	v949 = v909 + int32(1)
	goto L105
L109:
	;
	goto L102
L110:
	;
	goto L97
L111:
	;
	F_LockBufferInternal(m, v797, int32(3))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	v1038 = int32(0)
	v1039 = base.B2i32(v1038 <= v797)
	if v1039 == v1038 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v1063 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[2]))
	if v1063 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L114:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[10]))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1043+(v797^int32(-1))<<(uint(int32(2))%32))))
	v1057 = v1049
	goto L113
L115:
	;
	goto L116
L116:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[11]))
	v1057 = v1051 + v797<<(uint(int32(13))%32) + int32(-8192)
	goto L113
L117:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(3)
	v1107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v1108 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1108)
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v843
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1039 == v1108 {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	goto L117
L119:
	;
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[3])))
	if v1067&int32(1) == int32(0) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v1072 = int32(_a_F_lazy_vacuum_1)
	v1074 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	v1075 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v1074 + v1075
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	*(*int32)(unsafe.Add(mBase, uint32(v1063))) = v1078 + v1075
	v1082 = int32(0)
	v1084 = int32(_a_F_lazy_vacuum_2)
	v1085 = base.AtomicRmwOr32(m, v1082, v1084, v1082)
	*(*int64)(unsafe.Add(mBase, uint32(v1063+int32(24))+232)) = base.I64_extend_i32_u(v843)
	v1093 = base.AtomicRmwOr32(m, v1082, v1084, v1082)
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	*(*int32)(unsafe.Add(mBase, uint32(v1063))) = v1094 + v1075
	v1100 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v1100 - v1075
	goto L118
L121:
	;
	if v797 < int32(0) {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[10]))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1117+(v797^int32(-1))<<(uint(int32(2))%32))))
	v1131 = v1123
	goto L121
L123:
	;
	goto L124
L124:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[11]))
	v1131 = v1125 + v797<<(uint(int32(13))%32) + int32(-8192)
	goto L121
L125:
	;
	v1151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1131)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1151) {
		goto L132
	} else {
		goto L133
	}
L126:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[8]))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1135+(v797^int32(-1))*int32(56))+16))
	v1150 = v1141
	goto L125
L127:
	;
	goto L128
L128:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[9]))
	v1144 = int32(56)
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1143+v797*v1144-v1144)+16))
	v1150 = v1149
	goto L125
L129:
	;
	v1442 = int32(0)
	v1443 = int32(_a_F_lazy_vacuum_1)
	v1445 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v1445 + int32(1)
	if v1442 < v1029 {
		goto L186
	} else {
		goto L187
	}
L130:
	;
	F_LockBufferInternal(m, v1058, int32(3))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L7
	} else {
		goto L185
	}
L131:
	;
	v1167 = int32(base.Ui32(v1150) >> (uint(int32(16)) % 32))
	v1170 = int32(1)
	v1171 = int32(0)
	v1175 = v1171
	v1177 = v1170
	v1180 = v1171
	v1183 = v1170
	goto L138
L132:
	;
	v1159 = int32(base.Ui32(v1151+int32(_a_F_lazy_vacuum_4))>>(uint(int32(2))%32)) & int32(_a_F_lazy_vacuum_5)
	if v1159 != 0 {
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v1161 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1161)
	v1369 = v1161
	v1377 = int32(1)
	v1403 = int32(3)
	goto L130
L135:
	;
	goto L134
L136:
	;
	v1361 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1361)
	v1364 = int32(1)
	if v1338&v1364 != 0 {
		goto L182
	} else {
		goto L183
	}
L137:
	;
	v1357 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1357)
	v1408 = v1351
	v1413 = v1357
	v1416 = v1357
	goto L129
L138:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1177)
	v1211 = v1177 & int32(_a_F_lazy_vacuum_5)
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1131+int32(20)+v1211<<(uint(int32(2))%32))))
	switch int32(base.Ui32(v1215)>>(uint(int32(15))%32)) & int32(3) {
	case 0, 2:
		v1335 = v1175
		v1337 = v1180
		v1338 = v1183
		goto L140
	default:
		goto L141
	}
L139:
	;
	if base.Ui32(v1335) < base.Ui32(int32(3)) {
		goto L136
	} else {
		goto L179
	}
L140:
	;
	v1341 = v1177 + int32(1)
	if base.Ui32(v1341&int32(_a_F_lazy_vacuum_5)) <= base.Ui32(v1159) {
		v1175 = v1335
		v1177 = v1341
		v1180 = v1337
		v1183 = v1338
		goto L138
	} else {
		goto L178
	}
L141:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[5]))) = uint16(v1177)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[12]))) = uint16(v1150)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[13]))) = uint16(v1167)
	v1223 = int32(_a_F_lazy_vacuum_6)
	if v1215&v1223 == v1223 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	if v1029 <= v1180 {
		v1351 = v1175
		goto L137
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[6]))) = int32(base.Ui32(v1215) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[7]))) = v1131 + v1215&int32(_a_F_lazy_vacuum_7)
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[14]))) = v1244
	v1250 = F_HeapTupleSatisfiesVacuumHorizon(m, v38+int32(_a_F_lazy_vacuum_3), v797, v38+int32(48))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L7
	} else {
		goto L147
	}
L145:
	;
	v1233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+int32(96)+v1180<<(uint(int32(1))%32)))))
	if v1233 != v1211 {
		v1351 = v1175
		goto L137
	} else {
		goto L146
	}
L146:
	;
	v1335 = v1175
	v1337 = v1180 + int32(1)
	v1338 = v1183
	goto L140
L147:
	;
	if v1250 != int32(1) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	if base.B2i32(v1250 != int32(1))&base.B2i32(base.Ui32(v1250) <= base.Ui32(int32(4))) != 0 {
		v1351 = v1175
		goto L137
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_lazy_vacuum[7])))
	v1273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1272)+20)))
	if v1273&int32(256) == int32(0) {
		v1351 = v1175
		goto L137
	} else {
		goto L155
	}
L151:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L7
	} else {
		goto L152
	}
L152:
	;
	F_errmsg_internal(m, int32(_a_F_lazy_vacuum_8), int32(0))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L7
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_lazy_vacuum_9), int32(3754), int32(_a_F_lazy_vacuum_10))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L7
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	v1278 = int32(768)
	if v1273&v1278 == v1278 {
		v1297 = v1175
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v1301 = int32(0)
	if v1183&int32(1) == v1301 {
		v1335 = v1297
		v1337 = v1180
		v1338 = v1301
		goto L140
	} else {
		goto L164
	}
L157:
	;
	v1282 = int32(3)
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1272)))
	if base.B2i32(base.Ui32(v1175) < base.Ui32(v1282))|base.B2i32(base.Ui32(v1284) < base.Ui32(v1282)) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v1297 = v1284
	goto L156
L159:
	;
	if int32(0) < v1284-v1175 {
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	if base.B2i32(base.Ui32(v1284) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v1284) <= base.Ui32(v1175)) != 0 {
		v1297 = v1175
		goto L156
	} else {
		goto L163
	}
L162:
	;
	v1297 = v1175
	goto L156
L163:
	;
	goto L158
L164:
	;
	v1306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1272)+20)))
	v1307 = int32(768)
	if v1306&v1307 == v1307 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v1335 = v1297
	v1337 = v1180
	v1338 = v1332 ^ int32(1)
	goto L140
L166:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+4))
	if v1306&int32(_a_F_lazy_vacuum_11) != 0 {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1272)))
	if base.Ui32(v1311) <= base.Ui32(int32(2)) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v1332 = int32(1)
	goto L165
L169:
	;
	if base.Ui32(v1306) < base.Ui32(int32(_a_F_lazy_vacuum_12)) {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	if v1315 == int32(0) {
		goto L169
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	if base.Ui32(v1315) <= base.Ui32(int32(2)) {
		goto L169
	} else {
		goto L174
	}
L173:
	;
	v1332 = int32(1)
	goto L165
L174:
	;
	v1332 = int32(1)
	goto L165
L175:
	;
	v1332 = int32(0)
	goto L165
L176:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+8))
	if base.Ui32(v1326) <= base.Ui32(int32(2)) {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	v1332 = int32(1)
	goto L165
L178:
	;
	goto L139
L179:
	;
	v1347 = F_GlobalVisTestXidConsideredRunning(m, v1112, v1335)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L7
	} else {
		goto L180
	}
L180:
	;
	if v1347 == int32(0) {
		goto L136
	} else {
		goto L181
	}
L181:
	;
	v1351 = v1335
	goto L137
L182:
	;
	v1367 = int32(3)
	goto L184
L183:
	;
	v1367 = v1364
	goto L184
L184:
	;
	v1369 = v1335
	v1377 = v1338
	v1403 = v1367
	goto L130
L185:
	;
	v1408 = v1369
	v1413 = v1403
	v1416 = v1377
	goto L129
L186:
	;
	v1453 = v1029 & int32(3)
	v1455 = v1057 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v1029) {
		goto L190
	} else {
		goto L191
	}
L187:
	;
	v1684 = v1442
	goto L188
L188:
	;
	v1716 = int32(0)
	v1722 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1057)+12)))
	if base.Ui32(v1722) < base.Ui32(int32(25)) {
		goto L201
	} else {
		goto L202
	}
L189:
	;
	v1684 = v1029
	goto L188
L190:
	;
	v1465 = v1442
	v1476 = int32(0)
	goto L193
L191:
	;
	v1559 = v1442
	goto L192
L192:
	;
	v1594 = v1559
	v1609 = int32(0)
	goto L197
L193:
	;
	v1498 = v1465 << (uint(int32(1)) % 32)
	v1500 = v38 + int32(96)
	v1502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1498+v1500))))
	v1503 = int32(2)
	v1506 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1455+v1502<<(uint(v1503)%32)))) = v1506
	v1509 = v38 + int32(_a_F_lazy_vacuum_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v1509+v1498))) = uint16(v1502)
	v1513 = v1498 | v1503
	v1515 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1500+v1513))))
	*(*int32)(unsafe.Add(mBase, uint32(v1455+v1515<<(uint(v1503)%32)))) = v1506
	*(*uint16)(unsafe.Add(mBase, uint32(v1513+v1509))) = uint16(v1515)
	v1523 = int32(4)
	v1524 = v1498 | v1523
	v1526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1500+v1524))))
	*(*int32)(unsafe.Add(mBase, uint32(v1455+v1526<<(uint(v1503)%32)))) = v1506
	*(*uint16)(unsafe.Add(mBase, uint32(v1509+v1524))) = uint16(v1526)
	v1537 = v1498 | int32(6)
	v1539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1500+v1537))))
	*(*int32)(unsafe.Add(mBase, uint32(v1455+v1539<<(uint(v1503)%32)))) = v1506
	*(*uint16)(unsafe.Add(mBase, uint32(v1509+v1537))) = uint16(v1539)
	v1550 = v1465 + v1523
	v1552 = v1476 + v1523
	if v1552 != v1029&int32(2147483644) {
		v1465 = v1550
		v1476 = v1552
		goto L193
	} else {
		goto L195
	}
L194:
	;
	if v1453 == int32(0) {
		goto L189
	} else {
		goto L196
	}
L195:
	;
	goto L194
L196:
	;
	v1559 = v1550
	goto L192
L197:
	;
	v1626 = int32(1)
	v1627 = v1594 << (uint(v1626) % 32)
	v1631 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1627+(v38+int32(96))))))
	*(*int32)(unsafe.Add(mBase, uint32(v1455+v1631<<(uint(int32(2))%32)))) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v38+int32(_a_F_lazy_vacuum_3)+v1627))) = uint16(v1631)
	v1644 = v1609 + v1626
	if v1644 != v1453 {
		v1594 = v1594 + v1626
		v1609 = v1644
		goto L197
	} else {
		goto L199
	}
L198:
	;
	goto L189
L199:
	;
	goto L198
L200:
	;
	if v1413 != 0 {
		goto L217
	} else {
		goto L218
	}
L201:
	;
	v1787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1057)+10)))
	v1789 = v1787 & int32(_a_F_lazy_vacuum_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v1057)+10)) = uint16(v1789)
	goto L200
L202:
	;
	v1730 = int32(base.Ui32(v1722+int32(_a_F_lazy_vacuum_4))>>(uint(int32(2))%32)) & int32(_a_F_lazy_vacuum_5)
	if v1730 == int32(0) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v1736 = v1730
	v1738 = v1716
	v1741 = v1716
	goto L205
L204:
	;
	if int32(0) < v1765 {
		goto L213
	} else {
		goto L214
	}
L205:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1057+int32(20)+v1736<<(uint(int32(2))%32))))
	v1747 = v1745 & int32(_a_F_lazy_vacuum_6)
	if base.B2i32(v1736 == int32(1))|v1741 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	v1765 = v1759
	v1767 = int32(0)
	goto L204
L207:
	;
	v1762 = v1736 - int32(1)
	if v1762 != 0 {
		v1736 = v1762
		v1738 = v1759
		v1741 = v1760
		goto L205
	} else {
		goto L212
	}
L208:
	;
	v1753 = int32(0)
	v1759 = v1738 + base.B2i32(v1747 == v1753)
	v1760 = base.B2i32(v1747 != v1753)
	goto L207
L209:
	;
	goto L210
L210:
	;
	if v1747 != 0 {
		v1759 = v1738
		v1760 = v1741
		goto L207
	} else {
		goto L211
	}
L211:
	;
	v1765 = v1738
	v1767 = int32(1)
	goto L204
L212:
	;
	goto L206
L213:
	;
	v1772 = v1722 - v1765<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1057)+12)) = uint16(v1772)
	goto L215
L214:
	;
	goto L215
L215:
	;
	if v1767 == int32(0) {
		goto L201
	} else {
		goto L216
	}
L216:
	;
	v1776 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1057)+10)))
	v1778 = v1776 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1057)+10)) = uint16(v1778)
	goto L200
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1057)+20)) = int32(0)
	v1800 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1057)+10)))
	v1802 = v1800 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1057)+10)) = uint16(v1802)
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1804)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+40)) = v1805
	v1807 = *(*int64)(unsafe.Add(mBase, uint32(v1804)))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v1807
	v1809 = F_visibilitymap_set(m, v843, v1058, v1413)
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L7
	} else {
		goto L220
	}
L218:
	;
	v1811 = v1442
	goto L219
L219:
	;
	F_MarkBufferDirty(m, v797)
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L7
	} else {
		goto L221
	}
L220:
	;
	v1811 = v1408
	goto L219
L221:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+48))
	v1817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1816)+118)))
	if v1817 != int32(112) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1840 = int32(_a_F_lazy_vacuum_1)
	v1842 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	v1843 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4])) = v1842 - v1843
	if v1413&v1843 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L223:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[15]))
	if v1821 <= int32(0) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+32))
	if v1824 != 0 {
		goto L222
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	if v1413 != 0 {
		goto L229
	} else {
		goto L230
	}
L227:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+40))
	if v1825 != 0 {
		goto L222
	} else {
		goto L228
	}
L228:
	;
	goto L226
L229:
	;
	v1827 = v1058
	goto L231
L230:
	;
	v1827 = int32(0)
	goto L231
L231:
	;
	v1828 = int32(0)
	F_log_heap_prune_and_freeze(m, v1815, v797, v1827, v1413, v1811, v1828, int32(2), v1828, v1828, v1828, v1828, v1828, v1828, v38+int32(_a_F_lazy_vacuum_3), v1684)
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L7
	} else {
		goto L232
	}
L232:
	;
	goto L222
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v1104
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1107)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1110
	if v1039 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L234:
	;
	F_UnlockBuffer(m, v1058)
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L7
	} else {
		goto L235
	}
L235:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v1853 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1852 + v1853
	if v1416&v1853 == int32(0) {
		goto L233
	} else {
		goto L236
	}
L236:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v1860 + int32(1)
	goto L233
L237:
	;
	v1888 = int32(4)
	v1889 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1884)+14)))
	v1890 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1884)+12)))
	v1891 = v1889 - v1890
	if v1891 <= v1888 {
		goto L242
	} else {
		goto L243
	}
L238:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[10]))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1870+(v797^int32(-1))<<(uint(int32(2))%32))))
	v1884 = v1876
	goto L237
L239:
	;
	goto L240
L240:
	;
	v1878 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[11]))
	v1884 = v1878 + v797<<(uint(int32(13))%32) + int32(-8192)
	goto L237
L241:
	;
	F_UnlockReleaseBuffer(m, v797)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L7
	} else {
		goto L260
	}
L242:
	;
	v1894 = v1888
	goto L244
L243:
	;
	v1894 = v1891
	goto L244
L244:
	;
	v1896 = v1894 - int32(4)
	if v1896 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1951 = int32(0)
	goto L241
L246:
	;
	goto L247
L247:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1890) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v1951 = v1896
	goto L241
L249:
	;
	v1907 = int32(base.Ui32(v1890+int32(_a_F_lazy_vacuum_4)) >> (uint(int32(2)) % 32))
	goto L251
L250:
	;
	v1907 = int32(0)
	goto L251
L251:
	;
	if base.Ui32(v1907&int32(_a_F_lazy_vacuum_5)) < base.Ui32(int32(291)) {
		goto L248
	} else {
		goto L252
	}
L252:
	;
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884)+10)))
	if v1912&int32(1) == int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1951 = int32(0)
	goto L241
L254:
	;
	goto L255
L255:
	;
	v1921 = int32(1)
	goto L256
L256:
	;
	v1930 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1884+int32(20)+v1921&int32(_a_F_lazy_vacuum_5)<<(uint(int32(2))%32))+1)))
	if v1930&int32(384) == int32(0) {
		goto L248
	} else {
		goto L258
	}
L257:
	;
	v1951 = int32(0)
	goto L241
L258:
	;
	v1936 = v1921 + int32(1)
	v1937 = int32(_a_F_lazy_vacuum_5)
	if base.Ui32(v1936&v1937) <= base.Ui32(v1907&v1937) {
		v1921 = v1936
		goto L256
	} else {
		goto L259
	}
L259:
	;
	goto L257
L260:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_RecordPageWithFreeSpace(m, v1954, v843, v1951)
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L7
	} else {
		goto L261
	}
L261:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L7
	} else {
		goto L262
	}
L262:
	;
	v1961 = v814 + int32(1)
	v1964 = F_read_stream_next_buffer(m, v781, v38+int32(88))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L7
	} else {
		goto L263
	}
L263:
	;
	if v1964 != 0 {
		v797 = v1964
		v814 = v1961
		goto L83
	} else {
		goto L264
	}
L264:
	;
	goto L84
L265:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	F_pfree(m, v2003)
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L7
	} else {
		goto L266
	}
L266:
	;
	F_pfree(m, v718)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L7
	} else {
		goto L267
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	if v2010 != 0 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	F_ReleaseBuffer(m, v2010)
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L7
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v2015 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L7
	} else {
		goto L272
	}
L271:
	;
	goto L270
L272:
	;
	if v2015 != 0 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v2019 = *(*int64)(unsafe.Add(mBase, uint32(v2018)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v1990
	*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v2019
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v2017
	F_errmsg(m, int32(_a_F_lazy_vacuum_14), v38)
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L7
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v707
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v710)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v713
	goto L1
L276:
	;
	F_errfinish(m, int32(_a_F_lazy_vacuum_9), int32(2757), int32(_a_F_lazy_vacuum_15))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L7
	} else {
		goto L277
	}
L277:
	;
	goto L275
L278:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v2076 + v2078
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v2081 != 0 {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	m.G0 = v38 + int32(_a_F_lazy_vacuum_0)
	return
L280:
	;
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+16))
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+24))
	F_TidStoreDestroy(m, v2083)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L7
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_TidStoreDestroy(m, v2112)
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L7
	} else {
		goto L288
	}
L283:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v2082)+56))
	v2087 = F_TidStoreCreateShared(m, v2086)
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L7
	} else {
		goto L284
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2081)+24)) = v2087
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v2087)+8))
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2090)))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2091)+28))
	goto L285
L285:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2093)+48)) = v2092
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+24))
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+4))
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2096)))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2097)))
	goto L286
L286:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v2081)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2099)+52)) = v2098
	*(*int64)(unsafe.Add(mBase, uint32(v2082)+64)) = int64(0)
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2103)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(104)))) = v2106 + int32(56)
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v2103)+24))
	goto L287
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v2110
	goto L279
L288:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2115)))
	v2117 = F_TidStoreCreateLocal(m, v2116)
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L7
	} else {
		goto L289
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v2117
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v2120)+8)) = int64(0)
	goto L279
}
func F_leading_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	v2 = l1
	v5 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.B2i32(l0 == v5)|base.B2i32(v10 <= v5) == v5 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v141
	goto L1
L3:
	;
	F_dopr_outchmulti(m, l0, v130, l3)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L18
	} else {
		goto L43
	}
L4:
	;
	if v2 == int32(0) {
		v130 = v10
		goto L3
	} else {
		goto L7
	}
L5:
	;
	v66 = v2
	v67 = v10
	goto L6
L6:
	;
	v71 = base.B2i32(v66 != int32(0))
	if v71 < v67 {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v19 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if base.B2i32(v18 == v19)|base.B2i32(base.Ui32(v21) < base.Ui32(v18)) == v19 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v61 = v59 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v61
	if int32(0) < v61 {
		v130 = v61
		goto L3
	} else {
		goto L21
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v26 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v49 = v21
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v49 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v2)
	goto L8
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v29 + int32(1)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v33 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v49 = v48
	goto L11
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v21 == v34 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v37 = v21 - v34
	v38 = F_fwrite(m, v34, int32(1), v37, v26)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v38 + v40
	if v37 == v38 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v44)
	goto L15
L21:
	;
	v66 = int32(0)
	v67 = v61
	goto L6
L22:
	;
	F_dopr_outchmulti(m, int32(32), v67-v71, l3)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L18
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v66 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v71
	goto L24
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v81 = int32(0)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if base.B2i32(v80 == v81)|base.B2i32(base.Ui32(v83) < base.Ui32(v80)) == v81 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v121 {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v88 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v112 = v83
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v112 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v66)
	goto L27
L31:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v91 + int32(1)
	goto L27
L32:
	;
	goto L33
L33:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v95 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v112 = v110
	goto L30
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v83 == v96 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v99 = v83 - v96
	v100 = F_fwrite(m, v96, int32(1), v99, v88)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v100 + v102
	if v99 == v100 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v106 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v106)
	goto L34
L39:
	;
	v141 = v121 - int32(1)
	goto L2
L40:
	;
	goto L41
L41:
	;
	if int32(0) <= v121 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v141 = v121 + int32(1)
	goto L2
L43:
	;
	v141 = int32(0)
	goto L2
}
func F_leadlag_common(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WinCheckAndInitializeNullTreatment(m, v12, int32(1), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v20 = v10 + int32(15)
		v21 = F_WinGetFuncArgCurrent(m, v12, int32(1), v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
			if v23 == int32(0) {
				v26 = base.I32_wrap_i64(v21)
				if l1 != 0 {
					v29 = v26
				} else {
					v29 = int32(0) - v26
				}
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v32 = int32(0)
				if v30 == v32 {
					v78 = v32
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
					if v36 == int32(0) {
						v78 = v32
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						v41 = v39 - int32(11)
						if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v41))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v41)%32))&int32(1) == int32(0))|int32(0) != 0 {
							v78 = v32
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v41<<(uint(int32(2))%32))+uint32(_c_F_leadlag_common[0])))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v36+v56)))
							if v58 == int32(0) {
								v78 = v32
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								if v61 <= int32(1) {
									v78 = v32
								} else {
									v63 = int32(1)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(4))))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
									switch v69 - int32(7) {
									case 0:
										v78 = v63
									case 1:
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
										if v72 == int32(0) {
											v78 = v63
										} else {
											v78 = int32(0)
										}
									default:
										v78 = int32(0)
									}
								}
							}
						}
					}
				}
				v81 = F_WinGetFuncArgInPartition(m, v12, v29, v78, v20, v10+int32(14))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int64(0)
				} else {
					v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
					if v83&int32(1) != 0 {
						v87 = F_WinGetFuncArgCurrent(m, v12, int32(2), v20)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int64(0)
						} else {
							v89 = v87
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v90 != int32(1) {
								v99 = v89
							} else {
								v95 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v95)
								v99 = int64(0)
							}
							m.G0 = v10 + int32(16)
							return v99
						}
					} else {
						v89 = v81
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v90 != int32(1) {
							v99 = v89
						} else {
							v95 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v95)
							v99 = int64(0)
						}
						m.G0 = v10 + int32(16)
						return v99
					}
				}
			} else {
				v95 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v95)
				v99 = int64(0)
				m.G0 = v10 + int32(16)
				return v99
			}
		}
	}
}
func F_levenshtein_less_equal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			v48 = int32(1)
			if v17&v48 != 0 {
				v52 = v48
			} else {
				v52 = int32(4)
			}
			v54 = int32(1)
			if v16&v54 != 0 {
				v58 = v54
			} else {
				v58 = int32(4)
			}
			if v16 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v88 = v75
			} else {
				v76 = int32(1)
				if v16&v76 != 0 {
					v88 = int32(base.Ui32(v16)>>(uint(v76)%32)) - v76
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v89 = int32(1)
			v94 = F_varstr_levenshtein_less_equal(m, v9+v52, v46, v14+v58, v88, v89, v89, v89, base.I32_wrap_i64(v47), int32(0))
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_s(v94)
			}
		}
	}
}
func F_lexeme_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v6 < v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	if v5 < v6 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v59
L8:
	;
	v59 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v20 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v21 = v13
	v22 = v14
	v23 = v5
	v24 = v20
	goto L15
L12:
	;
	v47 = v14
	v51 = int32(0)
	goto L13
L13:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v59 = v51 - v52
	goto L7
L14:
	;
	v47 = v42
	v51 = v44
	goto L13
L15:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if base.B2i32(v24 != v26)|base.B2i32(v26 == int32(0)) != 0 {
		v42 = v22
		v44 = v24
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v42 = v36
	v44 = int32(0)
	goto L14
L17:
	;
	v32 = v23 - int32(1)
	if v32 == int32(0) {
		v42 = v22
		v44 = v24
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v35 = int32(1)
	v36 = v22 + v35
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v37 != 0 {
		v21 = v21 + v35
		v22 = v36
		v23 = v32
		v24 = v37
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
}
func F_lo_truncate_internal(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
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
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v198 int64
	_ = v198
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v287 int64
	_ = v287
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int64
	_ = v422
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	if l0 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L24
	} else {
		goto L112
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L24
	} else {
		goto L108
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[0]))
	if v21 <= l0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[1]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+l0<<(uint(int32(2))%32))))
	if v28 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	if v31&int32(2) == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v36 = m.G0
	v38 = v36 - int32(2256)
	m.G0 = v38
	base.MemoryFill(m, v38+int32(92), int32(0), int32(2052))
	v46 = base.I64_div_s(l1, int64(2048))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	if v47&int32(2) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	m.G0 = v16 + int32(32)
	return
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L24
	} else {
		goto L104
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L24
	} else {
		goto L101
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L24
	} else {
		goto L97
	}
L11:
	;
	if base.Ui64(int64(4398046509057)) <= base.Ui64(l1) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L24
	} else {
		goto L93
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[3]))
	if v56 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v57 = v53
	goto L17
L16:
	;
	v57 = int32(0)
	goto L17
L17:
	;
	if v57 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = int32(_a_F_lo_truncate_internal_0)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[4]))
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[4])) = v64
	if v53 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v89 = v53
	goto L20
L20:
	;
	v92 = v38 + int32(96)
	v93 = F_CatalogOpenIndexes(m, v89)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L24
	} else {
		goto L30
	}
L21:
	;
	v74 = v53
	v75 = v56
	goto L23
L22:
	;
	v69 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v75 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	return
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2])) = v69
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[3]))
	v74 = v69
	v75 = v73
	goto L23
L26:
	;
	v81 = F_index_open(m, int32(2683), int32(3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v86 = v74
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[4])) = v61
	v89 = v86
	goto L20
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[3])) = v81
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v86 = v85
	goto L28
L30:
	;
	v96 = v38 + int32(2144)
	v100 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28))))
	F_ScanKeyInit(m, v96, int32(1), int32(3), int32(184), v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	F_ScanKeyInit(m, v38+int32(2200), int32(2), int32(4), int32(150), v46)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[3]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v116 = F_systable_beginscan_ordered(m, v111, v113, v114, int32(2), v96)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L24
	} else {
		goto L35
	}
L33:
	;
	F_systable_endscan_ordered(m, v116)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L24
	} else {
		goto L90
	}
L34:
	;
	v319 = F_systable_getnext_ordered(m, v116, int32(1))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L24
	} else {
		goto L83
	}
L35:
	;
	v119 = F_systable_getnext_ordered(m, v116, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L24
	} else {
		goto L36
	}
L36:
	;
	if v119 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+20)))
	if v122&int32(1) != 0 {
		goto L9
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v246 = base.I32_wrap_i64(l1)
	v248 = v246 & int32(2047)
	if v248 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L40:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+22)))
	v126 = v121 + v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v127 == base.I32_wrap_i64(v46) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v131 = v126 + int32(8)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+8)))
	v134 = v132 & int32(3)
	if v134 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	F_simple_heap_delete(m, v239, v119+int32(4))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L24
	} else {
		goto L69
	}
L44:
	;
	v135 = F_detoast_attr(m, v131)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L24
	} else {
		goto L47
	}
L45:
	;
	v137 = v131
	goto L46
L46:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v140 = int32(base.Ui32(v138) >> (uint(int32(2)) % 32))
	v142 = v140 - int32(4)
	if base.Ui32(v140-int32(2053)) <= base.Ui32(int32(-2050)) {
		goto L8
	} else {
		goto L48
	}
L47:
	;
	v137 = v135
	goto L46
L48:
	;
	if v142 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	base.MemoryCopy(m, v92, v137+int32(4), v142)
	goto L51
L50:
	;
	goto L51
L51:
	;
	if v134 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_pfree(m, v137)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L24
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v154 = base.I32_wrap_i64(l1) & int32(2047)
	if v154 <= v142 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v198 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+64)) = v198
	*(*int64)(unsafe.Add(mBase, uint32(v38)+72)) = v198
	v202 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+60)) = uint16(v202)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+62)) = uint8(v202)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+56)) = uint16(v202)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+80)) = base.I64_extend_i32_u(v38 + int32(92))
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+58)) = uint8(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v154<<(uint(int32(2))%32) + int32(16)
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+52))
	v228 = F_heap_modify_tuple(m, v119, v221, v38-int32(-64), v38+int32(60), v38+int32(56))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L24
	} else {
		goto L66
	}
L57:
	;
	v157 = v38 + int32(92)
	v158 = v157 + v140
	v159 = int32(3)
	v161 = v154 - v142
	if v158&v159|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v161))|v161&v159 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if base.Ui32(v154+v92) <= base.Ui32(v158) {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	v188 = v161
	goto L60
L60:
	;
	if v188 == int32(0) {
		goto L56
	} else {
		goto L65
	}
L61:
	;
	v175 = int32(96)
	v176 = v38 + v140 + v175
	v179 = v38 + v154 + v175
	if base.Ui32(v179) < base.Ui32(v176) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v181 = v176
	goto L64
L63:
	;
	v181 = v179
	goto L64
L64:
	;
	v188 = (v157^int32(-1)+v181-v140)&int32(-4) + int32(4)
	goto L60
L65:
	;
	base.MemoryFill(m, v158, int32(0), v188)
	goto L56
L66:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	F_CatalogTupleUpdateWithInfo(m, v231, v228+int32(4), v228, v93)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L24
	} else {
		goto L67
	}
L67:
	;
	F_pfree(m, v228)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L24
	} else {
		goto L68
	}
L68:
	;
	goto L34
L69:
	;
	goto L39
L70:
	;
	v278 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+60)) = uint16(v278)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+62)) = uint8(v278)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v248<<(uint(int32(2))%32) + int32(16)
	v287 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28))))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+80)) = base.I64_extend_i32_u(v38 + int32(92))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+72)) = v46
	*(*int64)(unsafe.Add(mBase, uint32(v38)+64)) = v287
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+52))
	v301 = F_heap_form_tuple(m, v296, v38-int32(-64), v38+int32(60))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L24
	} else {
		goto L79
	}
L71:
	;
	if v246&int32(3) != 0 {
		v271 = v248
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v271 == int32(0) {
		goto L70
	} else {
		goto L78
	}
L73:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v248) {
		v271 = v248
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v257 = v38 + v248 + int32(96)
	v259 = v38 + int32(100)
	if base.Ui32(v259) < base.Ui32(v257) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v261 = v257
	goto L77
L76:
	;
	v261 = v259
	goto L77
L77:
	;
	v271 = (v261-v38-int32(97))&int32(-4) + int32(4)
	goto L72
L78:
	;
	base.MemoryFill(m, v92, int32(0), v271)
	goto L70
L79:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	F_CatalogTupleInsertWithInfo(m, v304, v301, v93)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L24
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v301)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L24
	} else {
		goto L81
	}
L81:
	;
	if v119 == int32(0) {
		goto L33
	} else {
		goto L82
	}
L82:
	;
	goto L34
L83:
	;
	if v319 == int32(0) {
		goto L33
	} else {
		goto L84
	}
L84:
	;
	v326 = v319
	goto L85
L85:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	F_simple_heap_delete(m, v337, v326+int32(4))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L24
	} else {
		goto L87
	}
L86:
	;
	goto L33
L87:
	;
	v343 = F_systable_getnext_ordered(m, v116, int32(1))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L24
	} else {
		goto L88
	}
L88:
	;
	if v343 != 0 {
		v326 = v343
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	F_CatalogCloseIndexes(m, v93)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L24
	} else {
		goto L91
	}
L91:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L24
	} else {
		goto L92
	}
L92:
	;
	m.G0 = v38 + int32(2256)
	goto L7
L93:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v374
	F_errmsg(m, int32(_a_F_lo_truncate_internal_1), v38)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L24
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_2), int32(766), int32(_a_F_lo_truncate_internal_3))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L24
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L24
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_lo_truncate_internal_4), v38+int32(16))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L24
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_2), int32(776), int32(_a_F_lo_truncate_internal_3))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L24
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errmsg_internal(m, int32(_a_F_lo_truncate_internal_5), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L24
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_2), int32(806), int32(_a_F_lo_truncate_internal_3))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L24
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
	F_errcode(m, int32(16779816))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L24
	} else {
		goto L105
	}
L105:
	;
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+40)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v422
	F_errmsg(m, int32(_a_F_lo_truncate_internal_6), v38+int32(32))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L24
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_2), int32(153), int32(_a_F_lo_truncate_internal_7))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L24
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L24
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg(m, int32(_a_F_lo_truncate_internal_8), v16)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L24
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_9), int32(565), int32(_a_F_lo_truncate_internal_10))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L24
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L24
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l0
	F_errmsg(m, int32(_a_F_lo_truncate_internal_11), v16+int32(16))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L24
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_9), int32(573), int32(_a_F_lo_truncate_internal_10))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L24
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_load_domaintype_info(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v151 int64
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v460 int32
	_ = v460
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v23 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L2:
	;
	v26 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v30 = v28 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v30
	if v26 < v30 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	F_MemoryContextDelete(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	goto L1
L6:
	;
	v43 = base.I64_extend_i32_u(v22)
	v44 = F_SearchSysCache1(m, int32(82), v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	F_ReleaseCatCache(m, v56)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L91
	}
L8:
	;
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v50 = v2
	v54 = v2
	v56 = v44
	v57 = v2
	v60 = v2
	v62 = v43
	goto L12
L10:
	;
	v334 = v22
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L88
	}
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)))
	v65 = v63 + v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+79)))
	if v66 != int32(100) {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v334 = v326
	goto L11
L14:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+130)))
	v71 = v20 + int32(48)
	F_ScanKeyInit(m, v71, int32(10), int32(3), int32(184), v62)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v77 = int32(0)
	v79 = int32(1)
	v82 = F_systable_beginscan(m, v40, int32(2666), v79, v77, v79, v71)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v65)+132))
	F_ReleaseCatCache(m, v56)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L85
	}
L17:
	;
	v84 = F_systable_getnext(m, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v84 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_systable_endscan(m, v82)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v91 = v77
	v93 = v84
	v94 = v50
	v98 = v54
	v101 = v57
	goto L23
L22:
	;
	v312 = v50
	v316 = v54
	v319 = v57
	goto L16
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
	v109 = v107 + v108
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+72)))
	if v110 == int32(99) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_systable_endscan(m, v82)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L75
	}
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+20)))
	if v114&int32(1) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v253 = v91
	v256 = v94
	v258 = v98
	v259 = v101
	goto L27
L27:
	;
	v260 = F_systable_getnext(m, v82)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L73
	}
L28:
	;
	if v94 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L29:
	;
	v179 = int64(*(*int8)(unsafe.Add(mBase, uint32(v122))))
	v182 = v179
	goto L28
L30:
	;
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v113)+244)))
	if int32(0) <= v119 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+26)))
	if v153&int32(8) != 0 {
		goto L48
	} else {
		goto L49
	}
L33:
	;
	v122 = v119 + v109
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+248)))
	if v123 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v151 = F_nocachegetattr(m, v93, int32(28), v113)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L47
	}
L36:
	;
	v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v113)+246)))
	if base.I32_popcnt(v126) != int32(1) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v182 = base.I64_extend_i32_u(v122)
	goto L28
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L44
	}
L40:
	;
	switch base.I32_ctz(v126) {
	case 0:
		goto L29
	case 1:
		goto L43
	case 2:
		goto L42
	case 3:
		goto L41
	default:
		goto L39
	}
L41:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v122)))
	v182 = v133
	goto L28
L42:
	;
	v132 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v122))))
	v182 = v132
	goto L28
L43:
	;
	v131 = int64(*(*int16)(unsafe.Add(mBase, uint32(v122))))
	v182 = v131
	goto L28
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v126
	F_errmsg_internal(m, int32(_a_F_load_domaintype_info_0), v20+int32(16))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_load_domaintype_info_1), int32(123), int32(_a_F_load_domaintype_info_2))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v182 = v151
	goto L28
L48:
	;
	v157 = F_nocachegetattr(m, v93, int32(28), v113)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	v182 = v157
	goto L28
L52:
	;
	v163 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v109 + v163
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v65 + v163
	F_errmsg_internal(m, int32(_a_F_load_domaintype_info_3), v20+int32(32))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_load_domaintype_info_4), int32(1177), int32(_a_F_load_domaintype_info_5))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
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
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	v191 = F_AllocSetContextCreateInternal(m, v186, int32(_a_F_load_domaintype_info_6), int32(0), int32(1024), int32(_a_F_load_domaintype_info_7))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	v202 = v94
	goto L57
L57:
	;
	v204 = F_text_to_cstring(m, base.I32_wrap_i64(v182))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	v194 = F_MemoryContextAlloc(m, v191, int32(12))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v196 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v194)+4)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v194))) = v196
	v202 = v194
	goto L57
L60:
	;
	v206 = F_stringToNode(m, v204)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v208 = F_expression_planner(m, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v210 = int32(_a_F_load_domaintype_info_8)
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v213
	v216 = F_palloc0(m, int32(20))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v216))) = int64(4294967695)
	v222 = F_pstrdup(m, v109+int32(4))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+8)) = v222
	v225 = F_copyObjectImpl(m, v208)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v227 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v216)+16)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v216)+12)) = v225
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v211
	if v98 == v227 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245+v91<<(uint(int32(2))%32)))) = v216
	v253 = v91 + int32(1)
	v256 = v202
	v258 = v245
	v259 = v246
	goto L27
L67:
	;
	v236 = F_palloc(m, int32(32))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v91 < v101 {
		v245 = v98
		v246 = v101
		goto L66
	} else {
		goto L71
	}
L70:
	;
	v245 = v236
	v246 = int32(8)
	goto L66
L71:
	;
	v241 = F_repalloc(m, v98, v101<<(uint(int32(3))%32))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v245 = v241
	v246 = v101 << (uint(int32(1)) % 32)
	goto L66
L73:
	;
	if v260 != 0 {
		v91 = v253
		v93 = v260
		v94 = v256
		v98 = v258
		v101 = v259
		goto L23
	} else {
		goto L74
	}
L74:
	;
	goto L24
L75:
	;
	if v253 <= int32(0) {
		v312 = v256
		v316 = v258
		v319 = v259
		goto L16
	} else {
		goto L76
	}
L76:
	;
	if v253 != int32(1) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_pg_qsort(m, v258, v253, int32(4), int32(1817))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v272 = int32(_a_F_load_domaintype_info_8)
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v279 = v253
	v285 = v277
	goto L81
L80:
	;
	goto L79
L81:
	;
	v296 = v279 - int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v258+v296<<(uint(int32(2))%32))))
	v301 = F_lcons(m, v300, v285)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v273
	v312 = v256
	v316 = v258
	v319 = v259
	goto L16
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v301
	if base.Ui32(int32(1)) < base.Ui32(v279) {
		v279 = v296
		v285 = v301
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v330 = base.I64_extend_i32_u(v326)
	v331 = F_SearchSysCache1(m, int32(82), v330)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	if v331 != 0 {
		v50 = v312
		v54 = v316
		v56 = v331
		v57 = v319
		v60 = v60 | v69
		v62 = v330
		goto L12
	} else {
		goto L87
	}
L87:
	;
	goto L13
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v334
	F_errmsg_internal(m, int32(_a_F_load_domaintype_info_9), v20)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_load_domaintype_info_4), int32(1136), int32(_a_F_load_domaintype_info_5))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_relation_close(m, v40, int32(1))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	if v60&int32(1) != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v460 | int32(_a_F_load_domaintype_info_10)
	m.G0 = v20 + int32(112)
	return
L94:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[1]))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v416)+16))
	if v422 != v418 {
		goto L109
	} else {
		goto L110
	}
L95:
	;
	if v50 != 0 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	goto L97
L97:
	;
	if v50 == int32(0) {
		goto L93
	} else {
		goto L107
	}
L98:
	;
	v389 = int32(_a_F_load_domaintype_info_8)
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v387
	v394 = F_palloc0(m, int32(20))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L104
	}
L99:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v387 = v370
	v388 = v50
	goto L98
L100:
	;
	goto L101
L101:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	v377 = F_AllocSetContextCreateInternal(m, v372, int32(_a_F_load_domaintype_info_6), int32(0), int32(1024), int32(_a_F_load_domaintype_info_7))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v380 = F_MemoryContextAlloc(m, v377, int32(12))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v382 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+8)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = v382
	v387 = v377
	v388 = v380
	goto L98
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = int64(399)
	v399 = F_pstrdup(m, int32(_a_F_load_domaintype_info_11))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v394)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v394)+8)) = v399
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v405 = F_lcons(m, v394, v404)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v405
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v390
	v415 = v388
	goto L94
L107:
	;
	v415 = v50
	goto L94
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v415
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v415)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v415)+8)) = v452 + int32(1)
	goto L93
L109:
	;
	if v422 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	goto L108
L112:
	;
	if v418 != 0 {
		goto L119
	} else {
		goto L120
	}
L113:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v416)+28))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v416)+24))
	if v427 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v426 == int32(0) {
		goto L112
	} else {
		goto L118
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427)+28)) = v426
	goto L114
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422)+20)) = v426
	goto L114
L118:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v416)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v426)+24)) = v432
	goto L112
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v416)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+16)) = v418
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v418)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v416)+28)) = v439
	if v439 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v416)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v416)+16)) = int32(0)
	goto L111
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v439)+24)) = v416
	goto L124
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418)+20)) = v416
	goto L108
}
func F_load_hba(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	v1 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v1
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_load_hba[0]))
	v21 = F_open_auth_file(m, v17, int32(15), v1, v1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v153
L2:
	;
	return int32(0)
L3:
	;
	if v21 == int32(0) {
		v153 = v1
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_load_hba[0]))
	F_tokenize_auth_file(m, v28, v21, v12+int32(12), int32(15), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_load_hba[1]))
	v41 = F_AllocSetContextCreateInternal(m, v36, int32(_a_F_load_hba_0), int32(0), int32(1024), int32(_a_F_load_hba_1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v43 = int32(_a_F_load_hba_2)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_load_hba[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_load_hba[2])) = v41
	v47 = int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v48 == int32(0) {
		v87 = v47
		v92 = v1
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v94 = int32(0)
	if base.B2i32(v87 == v94)|v92 == v94 {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v51 <= int32(0) {
		v87 = v47
		v92 = v1
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v54 = v1
	v56 = v47
	v61 = v1
	goto L10
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v54<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v68 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v87 = v78
	v92 = v80
	goto L7
L12:
	;
	v82 = v54 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v82 < v83 {
		v54 = v82
		v56 = v78
		v61 = v80
		goto L10
	} else {
		goto L21
	}
L13:
	;
	v78 = int32(0)
	v80 = v61
	goto L12
L14:
	;
	goto L15
L15:
	;
	v71 = F_parse_hba_line(m, v67, int32(15))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v71 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v78 = int32(0)
	v80 = v61
	goto L12
L18:
	;
	goto L19
L19:
	;
	v76 = F_lappend(m, v61, v71)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v78 = v56
	v80 = v76
	goto L12
L21:
	;
	goto L11
L22:
	;
	F_MemoryContextDelete(m, v41)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L42
	}
L23:
	;
	v101 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v128 = F_FreeFile(m, v21)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L35
	}
L26:
	;
	if v101 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v117 = F_FreeFile(m, v21)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L33
	}
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_load_hba[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v107
	F_errmsg(m, int32(_a_F_load_hba_3), v12)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_load_hba_4), int32(2515), int32(_a_F_load_hba_5))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_load_hba[3]))
	F_MemoryContextDelete(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_load_hba[2])) = v44
	*(*int32)(unsafe.Add(mBase, _c_F_load_hba[3])) = int32(0)
	goto L22
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_load_hba[3]))
	F_MemoryContextDelete(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_load_hba[2])) = v44
	v137 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_load_hba[3])) = v137
	if v87 == v137 {
		goto L22
	} else {
		goto L37
	}
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_load_hba[4]))
	if v142 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_MemoryContextDelete(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_load_hba[5])) = v92
	*(*int32)(unsafe.Add(mBase, _c_F_load_hba[4])) = v41
	v153 = int32(1)
	goto L1
L41:
	;
	goto L40
L42:
	;
	v153 = int32(0)
	goto L1
}
func F_load_libraries(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return
L2:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = F_pstrdup(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v21 = F_SplitDirectoriesString(m, v17, v10+int32(44))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_list_free_deep(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v23 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	F_pfree(m, v17)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v32 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v32 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l1
	F_errmsg(m, int32(_a_F_load_libraries_0), v10+int32(32))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_load_libraries_1), int32(1821), int32(_a_F_load_libraries_2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L1
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if int32(0) < v50 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v117 = int32(0)
	goto L19
L19:
	;
	F_list_free_deep(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L42
	}
L20:
	;
	v55 = int32(0)
	goto L23
L21:
	;
	goto L22
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v117 = v108
	goto L19
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v55<<(uint(int32(2))%32))))
	v66 = int32(0)
	if l2 == v66 {
		v78 = v65
		v79 = v66
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	F_load_file(m, v78, l2)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L30
	}
L26:
	;
	v71 = Fn14265(m, v65, int32(47))
	mBase = m.M
	goto L27
L27:
	;
	if v71 != 0 {
		v78 = v65
		v79 = int32(0)
		goto L25
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v65
	v76 = F_psprintf(m, int32(_a_F_load_libraries_3), v10+int32(16))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v78 = v76
	v79 = v76
	goto L25
L30:
	;
	v84 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if v84 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v78
	F_errmsg_internal(m, int32(_a_F_load_libraries_4), v10)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v79 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_errfinish(m, int32(_a_F_load_libraries_1), int32(1839), int32(_a_F_load_libraries_2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	F_pfree(m, v79)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v98 = v55 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v98 < v99 {
		v55 = v98
		goto L23
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	goto L24
L42:
	;
	F_pfree(m, v17)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L1
}
func F_local2local(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	if l2 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v50))) = uint8(v59)
	return v56 - l0
L2:
	;
	v50 = l1
	v56 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v14 = l1
	v15 = l2
	v20 = l0
	goto L5
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v23 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v50 = v44
	v56 = v42
	goto L1
L7:
	;
	if l6 != 0 {
		v50 = v14
		v56 = v20
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v30 = base.I32_extend8_s(v23)
	if int32(0) <= v30 {
		v39 = v30
		goto L13
	} else {
		goto L14
	}
L10:
	;
	F_report_invalid_encoding(m, l3, v20, v15)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v39)
	v41 = int32(1)
	v42 = v20 + v41
	v44 = v14 + v41
	if v41 < v15 {
		v14 = v44
		v15 = v15 - v41
		v20 = v42
		goto L5
	} else {
		goto L18
	}
L14:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v23-int32(128)))))
	if v36 != 0 {
		v39 = v36
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if l6 != 0 {
		v50 = v14
		v56 = v20
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_report_untranslatable_char(m, l3, l4, v20, v15)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L6
}
func F_locate_agg_of_level_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v8 - int32(9) {
		case 0:
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v11 != v12 {
				v33 = F_expression_tree_walker_impl(m, l0, int32(1122), l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v33
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				if v14 < int32(0) {
					v33 = F_expression_tree_walker_impl(m, l0, int32(1122), l1)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						return v33
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
					return int32(1)
				}
			}
		case 1:
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v20 != v21 {
				v33 = F_expression_tree_walker_impl(m, l0, int32(1122), l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v33
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v23 < int32(0) {
					v33 = F_expression_tree_walker_impl(m, l0, int32(1122), l1)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						return v33
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
					return int32(1)
				}
			}
		default:
			if v8 == int32(67) {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v38 + int32(1)
				v44 = F_query_tree_walker_impl(m, l0, int32(1122), l1, int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v46 - int32(1)
					return v44
				}
			} else {
				v33 = F_expression_tree_walker_impl(m, l0, int32(1122), l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v33
				}
			}
		}
	}
}
func F_logfile_getname(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	v12 = F_palloc(m, int32(1024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_logfile_getname[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
	v21 = F_pg_snprintf(m, v12, int32(1024), int32(_a_F_logfile_getname_0), v8)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = F_strlen(m, v12)
	mBase = m.M
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_logfile_getname[1]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_logfile_getname[2]))
	v33 = F_pg_localtime(m, v8+int32(8), v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = F_pg_strftime(m, v12+v23, int32(1024)-v23, v28, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = F_strlen(m, v12)
	mBase = m.M
	if int32(5) <= v37 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	m.G0 = v8 + int32(16)
	return v12
L9:
	;
	v41 = v37 - int32(4)
	v42 = v41 + v12
	v43 = int32(_a_F_logfile_getname_1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logfile_getname[3])))
	if base.B2i32(v46 == int32(0))|base.B2i32(v46 != v49) != 0 {
		v67 = v46
		v68 = v49
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v71 = v37
	goto L11
L11:
	;
	v72 = v71 + v12
	v74 = int32(1024) - v71
	if v74 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	if v67-v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	v52 = v42
	v53 = v43
	goto L15
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v57
		v68 = v56
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v67 = v57
	v68 = v56
	goto L13
L17:
	;
	v60 = int32(1)
	if v57 == v56 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v70 = v37
	goto L21
L20:
	;
	v70 = v41
	goto L21
L21:
	;
	v71 = v70
	goto L11
L22:
	;
	goto L8
L23:
	;
	v190 = F_strlen(m, v186)
	mBase = m.M
	goto L22
L24:
	;
	v186 = l1
	goto L23
L25:
	;
	goto L26
L26:
	;
	v80 = v74 - int32(1)
	if (v72^l1)&int32(3) != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v183 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v183)
	v186 = v179
	goto L23
L28:
	;
	v164 = v159
	v165 = v160
	v166 = v161
	goto L49
L29:
	;
	if v154 == int32(0) {
		v179 = v152
		v180 = v153
		goto L27
	} else {
		goto L48
	}
L30:
	;
	v152 = l1
	v153 = v72
	v154 = v80
	goto L29
L31:
	;
	goto L32
L32:
	;
	v84 = int32(0)
	if base.B2i32(l1&int32(3) == v84)|base.B2i32(v80 == v84) == v84 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v120 == int32(0) {
		v179 = v117
		v180 = v118
		goto L27
	} else {
		goto L42
	}
L34:
	;
	v96 = l1
	v97 = v72
	v98 = v80
	goto L37
L35:
	;
	goto L36
L36:
	;
	v117 = l1
	v118 = v72
	v119 = v80
	v120 = base.B2i32(v80 != v84)
	goto L33
L37:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v100)
	if v100 == int32(0) {
		v159 = v96
		v160 = v97
		v161 = v98
		goto L28
	} else {
		goto L39
	}
L38:
	;
	v117 = v111
	v118 = v105
	v119 = v107
	v120 = v109
	goto L33
L39:
	;
	v104 = int32(1)
	v105 = v97 + v104
	v107 = v98 - v104
	v108 = int32(0)
	v109 = base.B2i32(v107 != v108)
	v111 = v96 + v104
	if v111&int32(3) == v108 {
		v117 = v111
		v118 = v105
		v119 = v107
		v120 = v109
		goto L33
	} else {
		goto L40
	}
L40:
	;
	if v107 != 0 {
		v96 = v111
		v97 = v105
		v98 = v107
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if base.B2i32(v123 == int32(0))|base.B2i32(base.Ui32(v119) < base.Ui32(int32(4))) != 0 {
		v152 = v117
		v153 = v118
		v154 = v119
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v130 = v117
	v131 = v118
	v132 = v119
	goto L44
L44:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v138 = int32(-2139062144)
	if (int32(16843008)-v135|v135)&v138 != v138 {
		v159 = v130
		v160 = v131
		v161 = v132
		goto L28
	} else {
		goto L46
	}
L45:
	;
	v152 = v146
	v153 = v144
	v154 = v148
	goto L29
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v135
	v143 = int32(4)
	v144 = v131 + v143
	v146 = v130 + v143
	v148 = v132 - v143
	if base.Ui32(int32(3)) < base.Ui32(v148) {
		v130 = v146
		v131 = v144
		v132 = v148
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v159 = v152
	v160 = v153
	v161 = v154
	goto L28
L49:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v168)
	if v168 == int32(0) {
		v179 = v164
		v180 = v165
		goto L27
	} else {
		goto L51
	}
L50:
	;
	v179 = v175
	v180 = v173
	goto L27
L51:
	;
	v172 = int32(1)
	v173 = v165 + v172
	v175 = v164 + v172
	v177 = v166 - v172
	if v177 != 0 {
		v164 = v175
		v165 = v173
		v166 = v177
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
}
func F_logicalmsg_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	if base.Ui32(int32(15)) < base.Ui32(v11) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(48)
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v17
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = int32(_a_F_logicalmsg_desc_0)
	goto L5
L4:
	;
	v21 = int32(_a_F_logicalmsg_desc_1)
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v21
	v24 = v14 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v24
	F_appendStringInfo(m, l0, int32(_a_F_logicalmsg_desc_2), v8+int32(32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v31 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v34 = v15 + v24
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a_F_logicalmsg_desc_3)
	F_appendStringInfo(m, l0, int32(_a_F_logicalmsg_desc_4), v8+int32(16))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v44) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v49 = int32(1)
	goto L11
L11:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v34))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_logicalmsg_desc_5)
	F_appendStringInfo(m, l0, int32(_a_F_logicalmsg_desc_4), v8)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L1
L13:
	;
	v62 = v49 + int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(v62) < base.Ui32(v63) {
		v49 = v62
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_lookup_rowtype_tupdesc_copy(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v3 = F_lookup_rowtype_tupdesc_internal(m, l0, l1)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_CreateTupleDescCopyConstr(m, v3)
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_lquery_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_parse_lquery(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v5)
		}
	}
}
func F_ltrim(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14239(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
