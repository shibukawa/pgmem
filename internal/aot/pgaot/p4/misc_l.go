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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
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
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), l0)
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = F_table_open(m, int32(2995), int32(1))
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = int32(1)
			v24 = F_systable_beginscan(m, v18, int32(2996), v21, int32(0), v21, v7)
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = F_systable_getnext(m, v24)
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_systable_endscan(m, v24)
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_relation_close(m, v18, int32(1))
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(48)
							return base.B2i32(v26 != int32(0))
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
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v398 int64
	_ = v398
	var v402 int32
	_ = v402
	var v404 int64
	_ = v404
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int64
	_ = v418
	var v420 int32
	_ = v420
	var v422 int64
	_ = v422
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int64
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int64
	_ = v453
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v476 int32
	_ = v476
	var v479 int64
	_ = v479
	var v483 int32
	_ = v483
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int64
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int64
	_ = v549
	var v554 int32
	_ = v554
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v609 int32
	_ = v609
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v748 int32
	_ = v748
	var v761 int32
	_ = v761
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int64
	_ = v790
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v826 int64
	_ = v826
	var v833 int32
	_ = v833
	var v835 int64
	_ = v835
	var v844 int64
	_ = v844
	var v849 int32
	_ = v849
	var v851 int64
	_ = v851
	var v860 int64
	_ = v860
	var v863 int64
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int64
	_ = v869
	var v870 int64
	_ = v870
	var v876 int64
	_ = v876
	var v879 int64
	_ = v879
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int64
	_ = v913
	var v918 int32
	_ = v918
	var v920 int64
	_ = v920
	var v924 int64
	_ = v924
	var v928 int64
	_ = v928
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int64
	_ = v962
	var v967 int32
	_ = v967
	var v969 int64
	_ = v969
	var v973 int64
	_ = v973
	var v977 int64
	_ = v977
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int64
	_ = v1011
	var v1021 int32
	_ = v1021
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1091 int32
	_ = v1091
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1143 int32
	_ = v1143
	var v1144 int64
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1163 int64
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1318 int32
	_ = v1318
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1355 int32
	_ = v1355
	var v1366 int32
	_ = v1366
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int64
	_ = v1400
	var v1402 int64
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1408 int64
	_ = v1408
	var v1410 int64
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1451 int32
	_ = v1451
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1466 int32
	_ = v1466
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1513 int32
	_ = v1513
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1529 int32
	_ = v1529
	var v1540 int32
	_ = v1540
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1649 int32
	_ = v1649
	var v1650 int64
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1715 int32
	_ = v1715
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1743 int32
	_ = v1743
	var v1750 int32
	_ = v1750
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1874 int32
	_ = v1874
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1910 int32
	_ = v1910
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1937 int32
	_ = v1937
	var v1941 int32
	_ = v1941
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2023 int32
	_ = v2023
	var v2027 int32
	_ = v2027
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2084 int32
	_ = v2084
	var v2110 int32
	_ = v2110
	var v2115 int32
	_ = v2115
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2125 int32
	_ = v2125
	var v2128 int32
	_ = v2128
	var v2131 int32
	_ = v2131
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2166 int32
	_ = v2166
	var v2171 int32
	_ = v2171
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2219 int32
	_ = v2219
	var v2222 int32
	_ = v2222
	var v2227 int32
	_ = v2227
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2262 int32
	_ = v2262
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2275 int32
	_ = v2275
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2345 int32
	_ = v2345
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2391 int32
	_ = v2391
	var v2396 int32
	_ = v2396
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2402 int64
	_ = v2402
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2429 int32
	_ = v2429
	var v2433 int32
	_ = v2433
	var v2438 int32
	_ = v2438
	var v2442 int32
	_ = v2442
	var v2446 int32
	_ = v2446
	var v2451 int32
	_ = v2451
	var v2456 int32
	_ = v2456
	var v2462 int32
	_ = v2462
	var v2467 int32
	_ = v2467
	var v2471 int32
	_ = v2471
	var v2475 int32
	_ = v2475
	var v2480 int32
	_ = v2480
	var v2512 int32
	_ = v2512
	var v2543 int64
	_ = v2543
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2587 int64
	_ = v2587
	var v2592 int32
	_ = v2592
	var v2625 int32
	_ = v2625
	var v2628 int32
	_ = v2628
	var v2632 int32
	_ = v2632
	var v2639 int32
	_ = v2639
	var v2649 int32
	_ = v2649
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2691 int32
	_ = v2691
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2701 int32
	_ = v2701
	var v2706 int32
	_ = v2706
	var v2709 int32
	_ = v2709
	var v2711 int32
	_ = v2711
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2718 int64
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2720 int32
	_ = v2720
	var v2722 int32
	_ = v2722
	var v2730 int32
	_ = v2730
	var v2766 int32
	_ = v2766
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2774 int32
	_ = v2774
	var v2780 int32
	_ = v2780
	var v2784 int32
	_ = v2784
	var v2789 int32
	_ = v2789
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
	v2766 = m.ExcPending
	if v2766 != 0 {
		goto L23
	} else {
		goto L437
	}
L2:
	;
	m.G0 = v33 + int32(192)
	return v2730
L3:
	;
	v2543 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+32)) = v2543 + int64(1)
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
	v2548 = int32(0)
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	if v2548 < v2549 {
		goto L413
	} else {
		goto L414
	}
L4:
	;
	F_LWLockRelease(m, v1098)
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L23
	} else {
		goto L411
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
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L23
	} else {
		goto L408
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L23
	} else {
		goto L405
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+316))
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
	*(*int32)(unsafe.Add(mBase, uint32(v83+v235<<(uint(int32(2))%32))+292)) = v91
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
	v2730 = v280
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
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+316))
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
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v1104
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1104)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+24)) = v1217
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1217)+20))
	v1221 = l1 << (uint(int32(2)) % 32)
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1221+v1222)))
	if v1219&v1224 != 0 {
		goto L218
	} else {
		goto L219
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L23
	} else {
		goto L213
	}
L75:
	;
	F_LWLockRelease(m, v731)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L23
	} else {
		goto L198
	}
L76:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[9]))
	v1098 = v1091 + v133&int32(15)<<(uint(int32(7))%32) + int32(_a_F_LockAcquireExtended_0)
	v1100 = F_LWLockAcquire(m, v1098, int32(0))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L23
	} else {
		goto L180
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
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v678 != int32(1) {
		goto L76
	} else {
		goto L127
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
	if int32(15) < v329 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v337 = F_LWLockAcquire(m, v333+int32(584), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v340+v133&int32(1023)<<(uint(int32(2))%32))+4))
	if v346 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[11]))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v357 = (v351 - int32(1)) & (v354 * int32(_a_F_LockAcquireExtended_1))
	v358 = int32(4)
	v359 = v357 << (uint(v358) % 32)
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+600))
	v367 = v362 + v357&int32(268435455)<<(uint(int32(3))%32)
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v367)))
	v370 = v351 << (uint(v358) % 32)
	v373 = v370
	v398 = int64(0)
	goto L88
L84:
	;
	goto L85
L85:
	;
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	F_LWLockRelease(m, v643+int32(584))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L23
	} else {
		goto L126
	}
L86:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	F_LWLockRelease(m, v496+int32(584))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L23
	} else {
		goto L104
	}
L87:
	;
	v483 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v367))) = v368 | int64(1)<<(uint(base.I64_extend_i32_u(l1+base.I32_wrap_i64(v479)-v483))%64)
	v494 = v483
	goto L86
L88:
	;
	v402 = v359 + base.I32_wrap_i64(v398)
	v404 = v398 * int64(3)
	if int64(base.Ui64(v368)>>(uint(v404)%64))&int64(7) == int64(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	if base.Ui32(v434) < base.Ui32(v370) {
		goto L101
	} else {
		goto L102
	}
L90:
	;
	v418 = v398 | int64(1)
	v420 = v359 + base.I32_wrap_i64(v418)
	v422 = v418 * int64(3)
	if int64(base.Ui64(v368)>>(uint(v422)%64))&int64(7) == int64(0) {
		goto L96
	} else {
		goto L97
	}
L91:
	;
	v416 = v402
	goto L90
L92:
	;
	goto L93
L93:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v361)+604))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v410+v402<<(uint(int32(2))%32))))
	if v414 == v354 {
		v479 = v404
		goto L87
	} else {
		goto L94
	}
L94:
	;
	v416 = v373
	goto L90
L95:
	;
	v436 = v398 + int64(2)
	if v436 != int64(16) {
		v373 = v434
		v398 = v436
		goto L88
	} else {
		goto L100
	}
L96:
	;
	v434 = v420
	goto L95
L97:
	;
	goto L98
L98:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v361)+604))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v428+v420<<(uint(int32(2))%32))))
	if v432 == v354 {
		v479 = v422
		goto L87
	} else {
		goto L99
	}
L99:
	;
	v434 = v416
	goto L95
L100:
	;
	goto L89
L101:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v361)+604))
	v441 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v440+v434<<(uint(v441)%32)))) = v354
	v446 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+600))
	v448 = int32(1)
	v452 = v447 + int32(base.Ui32(v434)>>(uint(v448)%32))&int32(2147483640)
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v452)))
	*(*int64)(unsafe.Add(mBase, uint32(v452))) = v453 | int64(1)<<(uint(base.I64_extend_i32_u(l1+v434&int32(15)*int32(3)-v448))%64)
	v467 = v357 << (uint(v441) % 32)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_LockAcquireExtended[12])))
	*(*int32)(unsafe.Add(mBase, uint32(v467)+uint32(_c_F_LockAcquireExtended[12]))) = v468 + v448
	v476 = v448
	goto L103
L102:
	;
	v476 = int32(0)
	goto L103
L103:
	;
	v494 = v476
	goto L86
L104:
	;
	if v494 == int32(0) {
		goto L78
	} else {
		goto L105
	}
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91)+24)) = int64(0)
	v505 = int32(0)
	v506 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+32)) = v506 + int64(1)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	if v505 < v511 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v2730 = int32(1)
	goto L2
L107:
	;
	v514 = v505
	goto L110
L108:
	;
	v587 = int32(0)
	goto L109
L109:
	;
	v590 = v510 + v587<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v590)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v590))) = v136
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+40)) = v594 + int32(1)
	if v136 != 0 {
		goto L116
	} else {
		goto L117
	}
L110:
	;
	v546 = v510 + v514<<(uint(int32(4))%32)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	if v136 == v547 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v587 = v511
	goto L109
L112:
	;
	v549 = *(*int64)(unsafe.Add(mBase, uint32(v546)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v546)+8)) = v549 + int64(1)
	goto L106
L113:
	;
	goto L114
L114:
	;
	v554 = v514 + int32(1)
	if v554 != v511 {
		v514 = v554
		goto L110
	} else {
		goto L115
	}
L115:
	;
	goto L111
L116:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)))
	if base.Ui32(v599) <= base.Ui32(int32(15)) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	goto L118
L118:
	;
	goto L106
L119:
	;
	goto L118
L120:
	;
	if v599 != int32(15) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	goto L119
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136+v599<<(uint(int32(2))%32))+292)) = v91
	goto L125
L124:
	;
	goto L125
L125:
	;
	v609 = v599 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)) = uint8(v609)
	goto L122
L126:
	;
	goto L78
L127:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v681|base.B2i32(base.Ui32(l1) < base.Ui32(int32(5))) != 0 {
		goto L76
	} else {
		goto L128
	}
L128:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v685 == int32(0) {
		goto L76
	} else {
		goto L129
	}
L129:
	;
	v689 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	v692 = base.AtomicRmwXchg32(m, v689, int32(0), int32(1))
	if v692 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	F_s_lock(m, v694, int32(_a_F_LockAcquireExtended_2), int32(1837), int32(_a_F_LockAcquireExtended_3))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L23
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	v706 = v701 + v133&int32(1023)<<(uint(int32(2))%32)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	v708 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v706)+4)) = v707 + v708
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+52)) = uint8(v708)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[15])) = v91
	v715 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v701))), uint32(v715))
	v719 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[16]))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+16))
	if v720 == v715 {
		goto L76
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	v724 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[9]))
	v731 = v724 + v133&int32(15)<<(uint(int32(7))%32) + int32(_a_F_LockAcquireExtended_0)
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[11]))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v739 = (v733 - int32(1)) & (v736 * int32(_a_F_LockAcquireExtended_1))
	v748 = v719
	v761 = v7
	goto L135
L135:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v748)))
	v779 = v776 + v761*int32(640)
	v781 = v779 + int32(584)
	v783 = F_LWLockAcquire(m, v781, int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L23
	} else {
		goto L137
	}
L136:
	;
	goto L76
L137:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v779)+60))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v785 != v786 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	F_LWLockRelease(m, v781)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L23
	} else {
		goto L178
	}
L139:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v779)+600))
	v790 = *(*int64)(unsafe.Add(mBase, uint32(v788+v739<<(uint(int32(3))%32))))
	if v790 == int64(0) {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v794 = v739 & int32(268435455) << (uint(int32(3)) % 32)
	v795 = v788 + v794
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v779)+604))
	v797 = v796 + v739<<(uint(int32(6))%32)
	v826 = int64(0)
	goto L141
L141:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v797+base.I32_wrap_i64(v826)<<(uint(int32(2))%32))))
	if v736 != v833 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v865 = F_LWLockAcquire(m, v731, int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L23
	} else {
		goto L152
	}
L143:
	;
	goto L142
L144:
	;
	v844 = v826 | int64(1)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v797+base.I32_wrap_i64(v844)<<(uint(int32(2))%32))))
	if v849 == v736 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v835 = *(*int64)(unsafe.Add(mBase, uint32(v795)))
	if int64(base.Ui64(v835)>>(uint(v826*int64(3))%64))&int64(7) == int64(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v863 = v826
	goto L143
L147:
	;
	v851 = *(*int64)(unsafe.Add(mBase, uint32(v795)))
	if int64(base.Ui64(v851)>>(uint(v844*int64(3))%64))&int64(7) != int64(0) {
		v863 = v844
		goto L143
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v860 = v826 + int64(2)
	if v860 != int64(16) {
		v826 = v860
		goto L141
	} else {
		goto L151
	}
L150:
	;
	goto L149
L151:
	;
	goto L138
L152:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v779)+600))
	v869 = *(*int64)(unsafe.Add(mBase, uint32(v867+v794)))
	v870 = int64(1)
	v876 = (v863*int64(12884901888) - int64(4294967296)) >> (uint(int64(32)) % 64)
	v879 = v870 << (uint(v876+v870) % 64)
	if v869&v879 != int64(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v884 = F_SetupLockInTable(m, v46, v779, l0, v133, int32(1))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L23
	} else {
		goto L156
	}
L154:
	;
	v924 = v869
	goto L155
L155:
	;
	v928 = int64(1) << (uint(v876+int64(2)) % 64)
	if v924&v928 != int64(0) {
		goto L161
	} else {
		goto L162
	}
L156:
	;
	if v884 == int32(0) {
		goto L75
	} else {
		goto L157
	}
L157:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v884)))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v888)+128))
	v890 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v888)+128)) = v889 + v890
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v888)+92))
	v895 = v893 + v890
	*(*int32)(unsafe.Add(mBase, uint32(v888)+92)) = v895
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v888)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v888)+16)) = v897 | int32(2)
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v888)+48))
	if v901 == v895 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v888)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v888)+20)) = v903 & int32(-3)
	goto L160
L159:
	;
	goto L160
L160:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v884)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v884)+12)) = v907 | int32(2)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v779)+600))
	v912 = v911 + v794
	v913 = *(*int64)(unsafe.Add(mBase, uint32(v912)))
	*(*int64)(unsafe.Add(mBase, uint32(v912))) = v913 & (v879 ^ int64(-1))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v779)+600))
	v920 = *(*int64)(unsafe.Add(mBase, uint32(v918+v794)))
	v924 = v920
	goto L155
L161:
	;
	v933 = F_SetupLockInTable(m, v46, v779, l0, v133, int32(2))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L23
	} else {
		goto L164
	}
L162:
	;
	v973 = v924
	goto L163
L163:
	;
	v977 = int64(1) << (uint(v876+int64(3)) % 64)
	if v973&v977 != int64(0) {
		goto L169
	} else {
		goto L170
	}
L164:
	;
	if v933 == int32(0) {
		goto L75
	} else {
		goto L165
	}
L165:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v933)))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v937)+128))
	v939 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v937)+128)) = v938 + v939
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v937)+96))
	v944 = v942 + v939
	*(*int32)(unsafe.Add(mBase, uint32(v937)+96)) = v944
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v937)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v937)+16)) = v946 | int32(4)
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v937)+52))
	if v950 == v944 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v937)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v937)+20)) = v952 & int32(-5)
	goto L168
L167:
	;
	goto L168
L168:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v933)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v933)+12)) = v956 | int32(4)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v779)+600))
	v961 = v960 + v794
	v962 = *(*int64)(unsafe.Add(mBase, uint32(v961)))
	*(*int64)(unsafe.Add(mBase, uint32(v961))) = v962 & (v928 ^ int64(-1))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v779)+600))
	v969 = *(*int64)(unsafe.Add(mBase, uint32(v967+v794)))
	v973 = v969
	goto L163
L169:
	;
	v982 = F_SetupLockInTable(m, v46, v779, l0, v133, int32(3))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L23
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	F_LWLockRelease(m, v731)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L23
	} else {
		goto L177
	}
L172:
	;
	if v982 == int32(0) {
		goto L75
	} else {
		goto L173
	}
L173:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v982)))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v986)+128))
	v988 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v986)+128)) = v987 + v988
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v986)+100))
	v993 = v991 + v988
	*(*int32)(unsafe.Add(mBase, uint32(v986)+100)) = v993
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v986)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v986)+16)) = v995 | int32(8)
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v986)+56))
	if v999 == v993 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v986)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v986)+20)) = v1001 & int32(-9)
	goto L176
L175:
	;
	goto L176
L176:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v982)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v982)+12)) = v1005 | int32(8)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v779)+600))
	v1010 = v1009 + v794
	v1011 = *(*int64)(unsafe.Add(mBase, uint32(v1010)))
	*(*int64)(unsafe.Add(mBase, uint32(v1010))) = v1011 & (v977 ^ int64(-1))
	goto L171
L177:
	;
	goto L138
L178:
	;
	v1055 = v761 + int32(1)
	v1057 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[16]))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+16))
	if base.Ui32(v1055) < base.Ui32(v1058) {
		v748 = v1057
		v761 = v1055
		goto L135
	} else {
		goto L179
	}
L179:
	;
	goto L136
L180:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v1104 = F_SetupLockInTable(m, v46, v1103, l0, v133, l1)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L23
	} else {
		goto L181
	}
L181:
	;
	if v1104 != 0 {
		goto L73
	} else {
		goto L182
	}
L182:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[15]))
	if v1107 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+20))
	v1112 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	v1115 = base.AtomicRmwXchg32(m, v1112, int32(0), int32(1))
	if v1115 != 0 {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	goto L185
L185:
	;
	F_LWLockRelease(m, v1098)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L23
	} else {
		goto L190
	}
L186:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	F_s_lock(m, v1117, int32(_a_F_LockAcquireExtended_2), int32(1869), int32(_a_F_LockAcquireExtended_4))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L23
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	v1127 = v1124 + v1108&int32(1023)<<(uint(int32(2))%32)
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1127)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1127)+4)) = v1128 - int32(1)
	v1132 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1107)+52)) = uint8(v1132)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[15])) = v1132
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1124))), uint32(v1132))
	goto L185
L189:
	;
	goto L188
L190:
	;
	v1144 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	if v1144 == int64(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	F_RemoveLocalLock(m, v91)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L23
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	if l4 != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L193
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L197
L196:
	;
	goto L197
L197:
	;
	goto L74
L198:
	;
	F_LWLockRelease(m, v781)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L23
	} else {
		goto L199
	}
L199:
	;
	F_AbortStrongLockAcquire(m)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L23
	} else {
		goto L200
	}
L200:
	;
	v1163 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	if v1163 == int64(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	F_RemoveLocalLock(m, v91)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L23
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	if l4 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	goto L203
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L207
L206:
	;
	goto L207
L207:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L23
	} else {
		goto L208
	}
L208:
	;
	F_errcode(m, int32(_a_F_LockAcquireExtended_5))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L23
	} else {
		goto L209
	}
L209:
	;
	F_errmsg(m, int32(_a_F_LockAcquireExtended_6), int32(0))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L23
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = int32(_a_F_LockAcquireExtended_7)
	F_errhint(m, int32(_a_F_LockAcquireExtended_8), v33+int32(80))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L23
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(1042), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L23
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	F_errcode(m, int32(_a_F_LockAcquireExtended_5))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L23
	} else {
		goto L214
	}
L214:
	;
	F_errmsg(m, int32(_a_F_LockAcquireExtended_6), int32(0))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L23
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = int32(_a_F_LockAcquireExtended_7)
	F_errhint(m, int32(_a_F_LockAcquireExtended_8), v33+int32(32))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L23
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(1080), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L23
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	v1252 = int32(0)
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v1256 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1256)+104)) = v1258
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+616))
	if v1260 == v1252 {
		v1318 = v1258
		goto L230
	} else {
		goto L231
	}
L219:
	;
	v1226 = F_LockCheckConflicts(m, v46, l1, v1217, v1104)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L23
	} else {
		goto L220
	}
L220:
	;
	if v1226 != 0 {
		goto L218
	} else {
		goto L221
	}
L221:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1217)+128))
	v1229 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+128)) = v1228 + v1229
	v1232 = v1221 + v1217
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+88))
	v1235 = v1233 + v1229
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+88)) = v1235
	v1238 = v1229 << (uint(l1) % 32)
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1217)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+16)) = v1238 | v1239
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+44))
	if v1242 == v1235 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v1217)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+20)) = v1244 & (v1238 ^ int32(-1))
	goto L224
L223:
	;
	goto L224
L224:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1104)+12)) = v1249 | v1238
	goto L4
L225:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L23
	} else {
		goto L402
	}
L226:
	;
	v1889 = m.G0
	v1891 = v1889 - int32(128)
	m.G0 = v1891
	v1894 = v1891 + int32(112)
	F_initStringInfo(m, v1894)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L23
	} else {
		goto L335
	}
L227:
	;
	F_LWLockRelease(m, v1098)
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L23
	} else {
		goto L332
	}
L228:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[15]))
	if v1575 != 0 {
		goto L275
	} else {
		goto L276
	}
L229:
	;
	switch v1571 - int32(1) {
	case 0:
		goto L227
	case 1:
		goto L228
	default:
		goto L4
	}
L230:
	;
	v1341 = v1254 + int32(32)
	if v1318 == int32(0) {
		v1466 = v1252
		goto L240
	} else {
		goto L241
	}
L231:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+28))
	if v1263 == int32(0) {
		v1318 = v1258
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v1267 = v1254 + int32(24)
	if v1263 == v1267 {
		v1318 = v1258
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v1277 = v1258
	v1282 = v1263
	goto L234
L234:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1282-int32(12))))
	if v1260 == v1301 {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v1318 = v1307
	goto L230
L236:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1282-int32(8))))
	v1307 = v1305 | v1277
	goto L238
L237:
	;
	v1307 = v1277
	goto L238
L238:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1282)+4))
	if v1308 != v1267 {
		v1277 = v1307
		v1282 = v1308
		goto L234
	} else {
		goto L239
	}
L239:
	;
	goto L235
L240:
	;
	if l3 != 0 {
		goto L265
	} else {
		goto L266
	}
L241:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+40))
	if v1344 == int32(0) {
		v1466 = v1252
		goto L240
	} else {
		goto L242
	}
L242:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+36))
	if base.B2i32(v1347 == int32(0))|base.B2i32(v1347 == v1341) != 0 {
		v1466 = v1252
		goto L240
	} else {
		goto L243
	}
L243:
	;
	v1355 = v1347
	v1366 = int32(0)
	goto L244
L244:
	;
	if v1260 != 0 {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	v1466 = int32(0)
	goto L240
L246:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+4))
	if v1461 != v1341 {
		v1355 = v1461
		v1366 = v1458
		goto L244
	} else {
		goto L264
	}
L247:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+616))
	if v1260 == v1383 {
		v1458 = v1366
		goto L246
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+100))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1385+v1386<<(uint(int32(2))%32))))
	if v1390&v1318 != 0 {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	goto L249
L251:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1385+v1253<<(uint(int32(2))%32))))
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+104))
	if v1395&v1396 != 0 {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	goto L253
L253:
	;
	v1458 = int32(1)<<(uint(v1386)%32) | v1366
	goto L246
L254:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v1400 = *(*int64)(unsafe.Add(mBase, uint32(v1254)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1399)+8)) = v1400
	v1402 = *(*int64)(unsafe.Add(mBase, uint32(v1254)))
	*(*int64)(unsafe.Add(mBase, uint32(v1399))) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v1399)+16)) = v1253
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1256)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1399)+20)) = v1405
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+92))
	v1408 = *(*int64)(unsafe.Add(mBase, uint32(v1407)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1399)+32)) = v1408
	v1410 = *(*int64)(unsafe.Add(mBase, uint32(v1407)))
	*(*int64)(unsafe.Add(mBase, uint32(v1399)+24)) = v1410
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v1399)+40)) = v1412
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1355)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1399)+44)) = v1414
	v1417 = int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18])) = v1417
	v1571 = v1417
	goto L229
L255:
	;
	goto L256
L256:
	;
	if v1395&v1366 != 0 {
		v1466 = v1355
		goto L240
	} else {
		goto L257
	}
L257:
	;
	v1421 = F_LockCheckConflicts(m, v46, v1253, v1254, v1257)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L23
	} else {
		goto L258
	}
L258:
	;
	if v1421 != 0 {
		v1466 = v1355
		goto L240
	} else {
		goto L259
	}
L259:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+128))
	v1426 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+128)) = v1425 + v1426
	v1431 = v1254 + v1253<<(uint(int32(2))%32)
	v1433 = v1431 + int32(88)
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1433)))
	*(*int32)(unsafe.Add(mBase, uint32(v1433))) = v1434 + v1426
	v1439 = v1426 << (uint(v1253) % 32)
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+16)) = v1439 | v1440
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1433)))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+44))
	if v1443 == v1444 {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v1571 = int32(0)
	goto L229
L261:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+20)) = v1446 & (v1439 ^ int32(-1))
	goto L263
L262:
	;
	goto L263
L263:
	;
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+12)) = v1451 | v1439
	goto L260
L264:
	;
	goto L245
L265:
	;
	v1540 = int32(2)
	goto L267
L266:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	if v1466 != 0 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	v1571 = v1540
	goto L229
L268:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+40))
	v1520 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+40)) = v1519 + v1520
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+20)) = v1523 | v1520<<(uint(v1253)%32)
	v1529 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+100)) = v1253
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+96)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+92)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+104)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v1529)+16)) = v1520
	v1540 = v1520
	goto L267
L269:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1466)))
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+4)) = v1466
	*(*int32)(unsafe.Add(mBase, uint32(v1496))) = v1497
	*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1496
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1496)))
	*(*int32)(unsafe.Add(mBase, uint32(v1501)+4)) = v1496
	goto L268
L270:
	;
	goto L271
L271:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+36))
	if v1503 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+36)) = v1341
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+32)) = v1254 + int32(32)
	goto L274
L273:
	;
	goto L274
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1496)+4)) = v1341
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1341)))
	*(*int32)(unsafe.Add(mBase, uint32(v1496))) = v1513
	*(*int32)(unsafe.Add(mBase, uint32(v1513)+4)) = v1496
	*(*int32)(unsafe.Add(mBase, uint32(v1341))) = v1496
	goto L268
L275:
	;
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1575)+20))
	v1580 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	v1583 = base.AtomicRmwXchg32(m, v1580, int32(0), int32(1))
	if v1583 != 0 {
		goto L278
	} else {
		goto L279
	}
L276:
	;
	goto L277
L277:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+12))
	if v1610 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L278:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	F_s_lock(m, v1585, int32(_a_F_LockAcquireExtended_2), int32(1869), int32(_a_F_LockAcquireExtended_4))
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L23
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[14]))
	v1595 = v1592 + v1576&int32(1023)<<(uint(int32(2))%32)
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1595)+4)) = v1596 - int32(1)
	v1600 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1575)+52)) = uint8(v1600)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[15])) = v1600
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1592))), uint32(v1600))
	goto L277
L281:
	;
	goto L280
L282:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+4))
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+20))
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1614)+4)) = v1615
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1615))) = v1617
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+28))
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1619)+4)) = v1620
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1104)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1620))) = v1622
	v1625 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[19]))
	v1631 = F_hash_search_with_hash_value(m, v1625, v1104, v1613<<(uint(int32(4))%32)^v133, int32(2), int32(0))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L23
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1217)+84))
	v1638 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1217)+84)) = v1637 - v1638
	v1643 = v1217 + l1<<(uint(int32(2))%32)
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1643)+44)) = v1644 - v1638
	F_LWLockRelease(m, v1098)
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L23
	} else {
		goto L287
	}
L285:
	;
	if v1631 == int32(0) {
		goto L225
	} else {
		goto L286
	}
L286:
	;
	goto L284
L287:
	;
	v1650 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	if v1650 == int64(0) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	F_RemoveLocalLock(m, v91)
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L23
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	if l3 == int32(0) {
		goto L226
	} else {
		goto L292
	}
L291:
	;
	goto L290
L292:
	;
	if l5 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = int32(0)
	v1660 = v33 + int32(148)
	F_initStringInfo(m, v1660)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L23
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v1874 = int32(0)
	if l4 == v1874 {
		v2730 = v1874
		goto L2
	} else {
		goto L331
	}
L296:
	;
	v1664 = v33 + int32(132)
	F_initStringInfo(m, v1664)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L23
	} else {
		goto L297
	}
L297:
	;
	v1668 = v33 + int32(116)
	F_initStringInfo(m, v1668)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L23
	} else {
		goto L298
	}
L298:
	;
	F_DescribeLockTag(m, v1660, v91)
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L23
	} else {
		goto L299
	}
L299:
	;
	v1673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+15)))
	v1674 = int32(2)
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1673<<(uint(v1674)%32))+uint32(_c_F_LockAcquireExtended[0])))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1676)+8))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1677+l1<<(uint(v1674)%32))))
	v1683 = F_LWLockAcquire(m, v1098, int32(1))
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L23
	} else {
		goto L300
	}
L300:
	;
	v1685 = m.G0
	v1687 = v1685 - int32(48)
	m.G0 = v1687
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v1691 = v33 + int32(112)
	v1692 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1691))) = v1692
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+28))
	if v1694 == v1692 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	m.G0 = v1687 + int32(48)
	F_LWLockRelease(m, v1098)
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L23
	} else {
		goto L320
	}
L302:
	;
	v1698 = v1689 + int32(24)
	if v1694 == v1698 {
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1700 = int32(1)
	v1703 = v1694
	v1704 = v1700
	v1715 = v1700
	goto L304
L304:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1703-int32(16))))
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+44))
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+96))
	if v1736 == v1703-int32(20) {
		goto L307
	} else {
		goto L308
	}
L305:
	;
	goto L301
L306:
	;
	v1767 = *(*int32)(unsafe.Add(mBase, uint32(v1703)+4))
	if v1767 != v1698 {
		v1703 = v1767
		v1704 = v1765
		v1715 = v1766
		goto L304
	} else {
		goto L319
	}
L307:
	;
	if v1704 != 0 {
		goto L310
	} else {
		goto L311
	}
L308:
	;
	goto L309
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1687)+32)) = v1735
	if v1715 != 0 {
		goto L315
	} else {
		goto L316
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1687))) = v1735
	F_appendStringInfo(m, v1664, int32(_a_F_LockAcquireExtended_10), v1687)
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L23
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1687)+16)) = v1735
	F_appendStringInfo(m, v1664, int32(_a_F_LockAcquireExtended_11), v1687+int32(16))
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L23
	} else {
		goto L314
	}
L313:
	;
	v1765 = int32(0)
	v1766 = v1715
	goto L306
L314:
	;
	v1765 = int32(0)
	v1766 = v1715
	goto L306
L315:
	;
	v1755 = int32(_a_F_LockAcquireExtended_10)
	goto L317
L316:
	;
	v1755 = int32(_a_F_LockAcquireExtended_11)
	goto L317
L317:
	;
	F_appendStringInfo(m, v1668, v1755, v1687+int32(32))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L23
	} else {
		goto L318
	}
L318:
	;
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1691)))
	*(*int32)(unsafe.Add(mBase, uint32(v1691))) = v1760 + int32(1)
	v1765 = v1704
	v1766 = int32(0)
	goto L306
L319:
	;
	goto L305
L320:
	;
	v1806 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L23
	} else {
		goto L321
	}
L321:
	;
	if v1806 != 0 {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+68)) = v1681
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v1809
	v1812 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v1812
	F_errmsg(m, int32(_a_F_LockAcquireExtended_12), v33-int32(-64))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L23
	} else {
		goto L325
	}
L323:
	;
	goto L324
L324:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v33)+148))
	F_pfree(m, v1835)
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L23
	} else {
		goto L328
	}
L325:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v33)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v1819
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v33)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = v1821
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	F_errdetail_log_plural(m, int32(_a_F_LockAcquireExtended_13), int32(_a_F_LockAcquireExtended_14), v1825, v33+int32(48))
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L23
	} else {
		goto L326
	}
L326:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(1192), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L23
	} else {
		goto L327
	}
L327:
	;
	goto L324
L328:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v33)+116))
	F_pfree(m, v1838)
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L23
	} else {
		goto L329
	}
L329:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v33)+132))
	F_pfree(m, v1841)
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L23
	} else {
		goto L330
	}
L330:
	;
	goto L295
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v2730 = v1874
	goto L2
L332:
	;
	v1881 = F_WaitOnLock(m, v91, v136)
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L23
	} else {
		goto L333
	}
L333:
	;
	if v1881 != int32(2) {
		goto L3
	} else {
		goto L334
	}
L334:
	;
	goto L226
L335:
	;
	F_initStringInfo(m, v1891+int32(96))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L23
	} else {
		goto L336
	}
L336:
	;
	v1902 = v1891 + int32(80)
	F_initStringInfo(m, v1902)
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L23
	} else {
		goto L337
	}
L337:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18]))
	if v1906 <= int32(0) {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+112))
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+116))
	F_appendBinaryStringInfo(m, v1891+int32(96), v2071, v2072)
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L23
	} else {
		goto L359
	}
L339:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	if v1906 == int32(1) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v1915 = int32(20)
	goto L342
L341:
	;
	v1915 = int32(44)
	goto L342
L342:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1910+v1915)))
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1902)))
	v1919 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1918))) = uint8(v1919)
	*(*int32)(unsafe.Add(mBase, uint32(v1902)+12)) = v1919
	*(*int32)(unsafe.Add(mBase, uint32(v1902)+4)) = v1919
	goto L343
L343:
	;
	F_DescribeLockTag(m, v1902, v1910)
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L23
	} else {
		goto L344
	}
L344:
	;
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1910)+20))
	v1928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1910)+15)))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1910)+16))
	v1930 = int32(2)
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1928<<(uint(v1930)%32))+uint32(_c_F_LockAcquireExtended[0])))
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+8))
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1933+v1929<<(uint(v1930)%32))))
	goto L345
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+64)) = v1927
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+68)) = v1937
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+76)) = v1917
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+72)) = v1941
	F_appendStringInfo(m, v1894, int32(_a_F_LockAcquireExtended_15), v1891-int32(-64))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L23
	} else {
		goto L346
	}
L346:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18]))
	if v1949 < int32(2) {
		goto L338
	} else {
		goto L347
	}
L347:
	;
	v1954 = int32(1)
	v1956 = v1949
	goto L348
L348:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	v1987 = v1984 + v1954*int32(24)
	if v1954 < v1956-int32(1) {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	goto L338
L350:
	;
	v1995 = v1987 + int32(44)
	goto L352
L351:
	;
	v1995 = v1984 + int32(20)
	goto L352
L352:
	;
	v1996 = *(*int32)(unsafe.Add(mBase, uint32(v1995)))
	v1998 = v1891 + int32(80)
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1998)))
	v2000 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1999))) = uint8(v2000)
	*(*int32)(unsafe.Add(mBase, uint32(v1998)+12)) = v2000
	*(*int32)(unsafe.Add(mBase, uint32(v1998)+4)) = v2000
	goto L353
L353:
	;
	F_DescribeLockTag(m, v1998, v1987)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L23
	} else {
		goto L354
	}
L354:
	;
	v2009 = v1891 + int32(112)
	F_appendStringInfoChar(m, v2009, int32(10))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L23
	} else {
		goto L355
	}
L355:
	;
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+20))
	v2014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1987)+15)))
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+16))
	v2016 = int32(2)
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v2014<<(uint(v2016)%32))+uint32(_c_F_LockAcquireExtended[0])))
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(v2018)+8))
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v2019+v2015<<(uint(v2016)%32))))
	goto L356
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+48)) = v2013
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+52)) = v2023
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+60)) = v1996
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+56)) = v2027
	F_appendStringInfo(m, v2009, int32(_a_F_LockAcquireExtended_15), v1891+int32(48))
	mBase = m.M
	v2033 = m.ExcPending
	if v2033 != 0 {
		goto L23
	} else {
		goto L357
	}
L357:
	;
	v2035 = v1954 + int32(1)
	v2037 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18]))
	if v2035 < v2037 {
		v1954 = v2035
		v1956 = v2037
		goto L348
	} else {
		goto L358
	}
L358:
	;
	goto L349
L359:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18]))
	if int32(0) < v2076 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v2084 = int32(0)
	goto L363
L361:
	;
	goto L362
L362:
	;
	v2391 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[21])))
	if v2391 == int32(1) {
		goto L391
	} else {
		goto L392
	}
L363:
	;
	v2110 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[17]))
	F_appendStringInfoChar(m, v1891+int32(96), int32(10))
	mBase = m.M
	v2115 = m.ExcPending
	if v2115 != 0 {
		goto L23
	} else {
		goto L365
	}
L364:
	;
	goto L362
L365:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v2110+v2084*int32(24))+20))
	v2121 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[22]))
	if int32(0) < v2121 {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+36)) = v2345
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+32)) = v2119
	F_appendStringInfo(m, v1891+int32(96), int32(_a_F_LockAcquireExtended_16), v1891+int32(32))
	mBase = m.M
	v2354 = m.ExcPending
	if v2354 != 0 {
		goto L23
	} else {
		goto L389
	}
L367:
	;
	v2125 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[23]))
	v2128 = v2125
	v2131 = int32(1)
	goto L370
L368:
	;
	goto L369
L369:
	;
	v2345 = int32(_a_F_LockAcquireExtended_17)
	goto L366
L370:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v2128)))
	v2158 = int32(0)
	v2160 = int32(_a_F_LockAcquireExtended_18)
	v2161 = base.AtomicRmwOr32(m, v2158, v2160, v2158)
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v2128)+4))
	v2166 = base.AtomicRmwOr32(m, v2158, v2160, v2158)
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2128)))
	if base.B2i32(v2157&int32(1) == v2158)&base.B2i32(v2171 == v2157) == v2158 {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	goto L369
L372:
	;
	goto L375
L373:
	;
	v2227 = v2162
	goto L374
L374:
	;
	if v2227 == v2119 {
		goto L382
	} else {
		goto L383
	}
L375:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[24]))
	if v2207 != 0 {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	v2227 = v2215
	goto L374
L377:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L23
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v2128)))
	v2211 = int32(0)
	v2213 = int32(_a_F_LockAcquireExtended_18)
	v2214 = base.AtomicRmwOr32(m, v2211, v2213, v2211)
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v2128)+4))
	v2219 = base.AtomicRmwOr32(m, v2211, v2213, v2211)
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v2128)))
	if v2210&int32(1)|base.B2i32(v2222 != v2210) != 0 {
		goto L375
	} else {
		goto L381
	}
L380:
	;
	goto L379
L381:
	;
	goto L376
L382:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v2128)+216))
	v2258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2257))))
	if v2258 == int32(0) {
		v2345 = int32(_a_F_LockAcquireExtended_19)
		goto L366
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	v2280 = v2131 + int32(1)
	v2282 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[22]))
	if v2280 <= v2282 {
		v2128 = v2128 + int32(408)
		v2131 = v2280
		goto L370
	} else {
		goto L388
	}
L385:
	;
	v2262 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[25]))
	v2265 = F_pnstrdup(m, v2257, v2262-int32(1))
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L23
	} else {
		goto L386
	}
L386:
	;
	v2267 = F_strlen(m, v2265)
	mBase = m.M
	v2269 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[25]))
	v2272 = F_pg_mbcliplen(m, v2265, v2267, v2269-int32(1))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L23
	} else {
		goto L387
	}
L387:
	;
	v2275 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2272+v2265))) = uint8(v2275)
	v2345 = v2265
	goto L366
L388:
	;
	goto L371
L389:
	;
	v2356 = v2084 + int32(1)
	v2358 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18]))
	if v2356 < v2358 {
		v2084 = v2356
		goto L363
	} else {
		goto L390
	}
L390:
	;
	goto L364
L391:
	;
	v2396 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[10]))
	v2399 = F_pgstat_prep_pending_entry(m, int32(1), v2396, int64(0), int32(0))
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		goto L23
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L23
	} else {
		goto L395
	}
L394:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2399)+12))
	v2402 = *(*int64)(unsafe.Add(mBase, uint32(v2401)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v2401)+144)) = v2402 + int64(1)
	goto L393
L395:
	;
	F_errcode(m, int32(16908292))
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L23
	} else {
		goto L396
	}
L396:
	;
	F_errmsg(m, int32(_a_F_LockAcquireExtended_20), int32(0))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L23
	} else {
		goto L397
	}
L397:
	;
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+16)) = v2418
	F_errdetail_internal(m, int32(_a_F_LockAcquireExtended_21), v1891+int32(16))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L23
	} else {
		goto L398
	}
L398:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v1891))) = v2425
	F_errdetail_log(m, int32(_a_F_LockAcquireExtended_21), v1891)
	mBase = m.M
	v2429 = m.ExcPending
	if v2429 != 0 {
		goto L23
	} else {
		goto L399
	}
L399:
	;
	F_errhint(m, int32(_a_F_LockAcquireExtended_22), int32(0))
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L23
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_23), int32(1138), int32(_a_F_LockAcquireExtended_24))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L23
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	F_errmsg_internal(m, int32(_a_F_LockAcquireExtended_25), int32(0))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L23
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(1140), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
		goto L23
	} else {
		goto L404
	}
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_LockAcquireExtended_26), v33+int32(16))
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L23
	} else {
		goto L406
	}
L406:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(861), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L23
	} else {
		goto L407
	}
L407:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v35
	F_errmsg_internal(m, int32(_a_F_LockAcquireExtended_27), v33)
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L23
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(858), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2480 = m.ExcPending
	if v2480 != 0 {
		goto L23
	} else {
		goto L410
	}
L410:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L411:
	;
	goto L3
L412:
	;
	v2682 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[15])) = v2682
	v2684 = int32(1)
	if v304 == v2682 {
		v2730 = v2684
		goto L2
	} else {
		goto L430
	}
L413:
	;
	v2553 = v2548
	goto L416
L414:
	;
	v2625 = int32(0)
	goto L415
L415:
	;
	v2628 = v2547 + v2625<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v2628)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2628))) = v136
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+40)) = v2632 + int32(1)
	if v136 == int32(0) {
		goto L412
	} else {
		goto L422
	}
L416:
	;
	v2584 = v2547 + v2553<<(uint(int32(4))%32)
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(v2584)))
	if v136 == v2585 {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	v2625 = v2549
	goto L415
L418:
	;
	v2587 = *(*int64)(unsafe.Add(mBase, uint32(v2584)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2584)+8)) = v2587 + int64(1)
	goto L412
L419:
	;
	goto L420
L420:
	;
	v2592 = v2553 + int32(1)
	if v2592 != v2549 {
		v2553 = v2592
		goto L416
	} else {
		goto L421
	}
L421:
	;
	goto L417
L422:
	;
	v2639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+18)))
	if base.Ui32(v2639) <= base.Ui32(int32(15)) {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	goto L412
L424:
	;
	if v2639 != int32(15) {
		goto L427
	} else {
		goto L428
	}
L425:
	;
	goto L426
L426:
	;
	goto L423
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+v2639<<(uint(int32(2))%32))+292)) = v91
	goto L429
L428:
	;
	goto L429
L429:
	;
	v2649 = v2639 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+18)) = uint8(v2649)
	goto L426
L430:
	;
	v2687 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2689 = m.G0
	v2691 = v2689 - int32(16)
	m.G0 = v2691
	v2693 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v2694 = m.ExcPending
	if v2694 != 0 {
		goto L23
	} else {
		goto L431
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2691)+8)) = v2688
	*(*int32)(unsafe.Add(mBase, uint32(v2691)+4)) = v2687
	*(*int32)(unsafe.Add(mBase, uint32(v2691))) = v2693
	*(*int32)(unsafe.Add(mBase, uint32(v2691)+12)) = int32(1)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2701 = m.ExcPending
	if v2701 != 0 {
		goto L23
	} else {
		goto L432
	}
L432:
	;
	F_XLogRegisterData(m, v2691+int32(12), int32(4))
	mBase = m.M
	v2706 = m.ExcPending
	if v2706 != 0 {
		goto L23
	} else {
		goto L433
	}
L433:
	;
	F_XLogRegisterData(m, v2691, int32(12))
	mBase = m.M
	v2709 = m.ExcPending
	if v2709 != 0 {
		goto L23
	} else {
		goto L434
	}
L434:
	;
	v2711 = int32(_a_F_LockAcquireExtended_28)
	v2713 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[26])))
	v2714 = v2713 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[26])) = uint8(v2714)
	goto L435
L435:
	;
	v2718 = F_XLogInsert(m, int32(8), int32(0))
	mBase = m.M
	v2719 = m.ExcPending
	if v2719 != 0 {
		goto L23
	} else {
		goto L436
	}
L436:
	;
	v2720 = int32(_a_F_LockAcquireExtended_29)
	v2722 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[27])) = v2722 | int32(2)
	m.G0 = v2691 + int32(16)
	v2730 = v2684
	goto L2
L437:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2769 = m.ExcPending
	if v2769 != 0 {
		goto L23
	} else {
		goto L438
	}
L438:
	;
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(v2770+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v2774
	F_errmsg(m, int32(_a_F_LockAcquireExtended_30), v33+int32(96))
	mBase = m.M
	v2780 = m.ExcPending
	if v2780 != 0 {
		goto L23
	} else {
		goto L439
	}
L439:
	;
	F_errhint(m, int32(_a_F_LockAcquireExtended_31), int32(0))
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L23
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(871), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L23
	} else {
		goto L441
	}
L441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__ltq_extract_regex(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13844(m, l0, int32(_a_F__ltq_extract_regex_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_lappend_int(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13939(m, l0, l1, int64(4294967767))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_lazy_vacuum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 float32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int64
	_ = v97
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int64
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v285 int32
	_ = v285
	var v317 int64
	_ = v317
	var v320 int64
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v377 int64
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int64
	_ = v475
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v628 int64
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v819 int32
	_ = v819
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v984 int32
	_ = v984
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1242 int32
	_ = v1242
	var v1256 int32
	_ = v1256
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1290 int32
	_ = v1290
	var v1324 int32
	_ = v1324
	var v1354 int32
	_ = v1354
	var v1360 int32
	_ = v1360
	var v1368 int32
	_ = v1368
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1391 int32
	_ = v1391
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1471 int32
	_ = v1471
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1575 int32
	_ = v1575
	var v1583 int32
	_ = v1583
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1647 int32
	_ = v1647
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1765 int32
	_ = v1765
	var v1806 int32
	_ = v1806
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1857 int32
	_ = v1857
	var v1866 int32
	_ = v1866
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1887 int32
	_ = v1887
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1922 int32
	_ = v1922
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int64
	_ = v1951
	var v1957 int32
	_ = v1957
	var v1962 int32
	_ = v1962
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	v2 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(_a_F_lazy_vacuum_0)
	m.G0 = v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+23)))
	if v36 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v34 + int32(_a_F_lazy_vacuum_0)
	return
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v62 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	F_parallel_vacuum_reset_dead_items(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_TidStoreDestroy(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L11
	}
L8:
	;
	return
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(104)))) = v45 + int32(56)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v49
	goto L1
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v56 = F_TidStoreCreateLocal(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = int64(0)
	goto L1
L13:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1999 != 0 {
		goto L280
	} else {
		goto L281
	}
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	v84 = *(*float32)(unsafe.Add(mBase, uint32(v83)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+40)) = int64(34359738368)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v88
	v91 = *(*int64)(unsafe.Add(mBase, _c_F_lazy_vacuum[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v91
	v93 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L20
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v65 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if base.Ui32(base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i32_u(v65), float64(0.02)))) <= base.Ui32(v68) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v75 = F_TidStoreMemoryUsage(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(int32(33554431)) < base.Ui32(v75) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+23)) = uint8(v79)
	goto L13
L20:
	;
	if v93 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[2]))) = int64(2)
	v97 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[3]))) = v97
	goto L24
L22:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v285 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	goto L22
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	if v113 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[5])))
	if v117&int32(1) == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v122 = int32(_a_F_lazy_vacuum_1)
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	v125 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v124 + v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v128 + v125
	v132 = int32(0)
	v135 = base.AtomicRmwOr32(m, v132, int32(_a_F_lazy_vacuum_2), v132)
	goto L28
L27:
	;
	v262 = int32(0)
	v265 = base.AtomicRmwOr32(m, v262, int32(_a_F_lazy_vacuum_2), v262)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v267 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v266 + v267
	v270 = int32(_a_F_lazy_vacuum_1)
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v272 - v267
	goto L23
L28:
	;
	goto L30
L30:
	;
	goto L31
L31:
	;
	v227 = int32(0)
	v230 = int32(0)
	goto L36
L36:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(40)+v230<<(uint(int32(2))%32))))
	v240 = int32(3)
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v34+int32(_a_F_lazy_vacuum_3)+v230<<(uint(v240)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v113+int32(232)+v239<<(uint(v240)%32)))) = v246
	v248 = int32(1)
	v251 = v227 + v248
	if v251 != int32(2) {
		v227 = v251
		v230 = v230 + v248
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L27
L38:
	;
	goto L37
L39:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v473 = v471 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v473
	v475 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[7]))) = v475
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[8]))) = v475
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[9]))) = base.I64_extend_i32_s(v473)
	goto L59
L40:
	;
	v317 = int64(0)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v426)+16)) = base.F64_convert_i32_s(base.I32_trunc_sat_f32_s(v84))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v285)+16))
	v431 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+24)) = uint8(v431)
	F_parallel_vacuum_process_all_indexes(m, v285, v425, v431)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L8
	} else {
		goto L55
	}
L43:
	;
	v320 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	v321 = base.B2i32(v320 <= v317)
	if v320 <= v317 {
		v445 = v321
		goto L39
	} else {
		goto L45
	}
L44:
	;
	v445 = v321
	goto L39
L45:
	;
	v324 = base.I32_wrap_i64(v317) << (uint(int32(2)) % 32)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v324+v325)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v328+v324)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v34)+64)) = base.F64_promote_f32(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = int32(13)
	v336 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+58)) = uint8(v336)
	v338 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+56)) = uint16(v338)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v332
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v341
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v330)+48))
	v346 = F_pstrdup(m, v343+int32(4))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v346
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v350 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v350)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(2)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v362 = F_vac_bulkdel_one_index(m, v34+int32(48), v327, v360, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v355
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v349)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v352
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	F_pfree(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v370 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v370
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v372+v324))) = v362
	v377 = v317 + int64(1)
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	if v380 == v370 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v421 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L8
	} else {
		goto L53
	}
L50:
	;
	goto L49
L51:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[5])))
	if v384&int32(1) == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v389 = int32(_a_F_lazy_vacuum_1)
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	v392 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v391 + v392
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = v395 + v392
	v399 = int32(0)
	v401 = int32(_a_F_lazy_vacuum_2)
	v402 = base.AtomicRmwOr32(m, v399, v401, v399)
	*(*int64)(unsafe.Add(mBase, uint32(v380+int32(72))+232)) = v377
	v410 = base.AtomicRmwOr32(m, v399, v401, v399)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = v411 + v392
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v417 - v392
	goto L50
L53:
	;
	if v421 == int32(0) {
		v317 = v377
		goto L43
	} else {
		goto L54
	}
L54:
	;
	goto L44
L55:
	;
	v436 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v445 = v436 ^ int32(1)
	goto L39
L57:
	;
	if v445 == int32(0) {
		goto L13
	} else {
		goto L74
	}
L58:
	;
	goto L57
L59:
	;
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	if v495 == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[5])))
	if v499&int32(1) == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v504 = int32(_a_F_lazy_vacuum_1)
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	v507 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v506 + v507
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v510 + v507
	v514 = int32(0)
	v517 = base.AtomicRmwOr32(m, v514, int32(_a_F_lazy_vacuum_2), v514)
	goto L63
L62:
	;
	v644 = int32(0)
	v647 = base.AtomicRmwOr32(m, v644, int32(_a_F_lazy_vacuum_2), v644)
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	v649 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v648 + v649
	v652 = int32(_a_F_lazy_vacuum_1)
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v654 - v649
	goto L58
L63:
	;
	goto L65
L65:
	;
	goto L66
L66:
	;
	v609 = int32(0)
	v612 = int32(0)
	goto L71
L71:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(24)+v612<<(uint(int32(2))%32))))
	v622 = int32(3)
	v628 = *(*int64)(unsafe.Add(mBase, uint32(v34+int32(_a_F_lazy_vacuum_4)+v612<<(uint(v622)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v495+int32(232)+v621<<(uint(v622)%32)))) = v628
	v630 = int32(1)
	v633 = v609 + v630
	if v633 != int32(3) {
		v609 = v633
		v612 = v612 + v630
		goto L71
	} else {
		goto L73
	}
L72:
	;
	goto L62
L73:
	;
	goto L72
L74:
	;
	v669 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v669
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	if v675 == v669 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(3)
	v719 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v720 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v720)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v727 = F_palloc0(m, int32(16))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L8
	} else {
		goto L79
	}
L76:
	;
	goto L75
L77:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[5])))
	if v679&int32(1) == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v684 = int32(_a_F_lazy_vacuum_1)
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	v687 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v686 + v687
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	*(*int32)(unsafe.Add(mBase, uint32(v675))) = v690 + v687
	v694 = int32(0)
	v696 = int32(_a_F_lazy_vacuum_2)
	v697 = base.AtomicRmwOr32(m, v694, v696, v694)
	*(*int64)(unsafe.Add(mBase, uint32(v675+v694)+232)) = int64(3)
	v705 = base.AtomicRmwOr32(m, v694, v696, v694)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	*(*int32)(unsafe.Add(mBase, uint32(v675))) = v706 + v687
	v712 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v712 - v687
	goto L76
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v727))) = v725
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v725)+4))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v725)+8))
	if v731 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v727)+4)) = v778
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v790 = F_read_stream_begin_relation(m, int32(9), v785, v786, int32(0), int32(190), v727, int32(8))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L8
	} else {
		goto L87
	}
L81:
	;
	v733 = F_palloc0(m, int32(120))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L8
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v762 = F_palloc0(m, int32(88))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L8
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v733))) = v730
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v730)+4))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)+24))
	v739 = F_dsa_get_address(m, v736, v738)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v733)))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)+48))
	v745 = base.I32_div_s(v743, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v733)+104)) = v745
	*(*int32)(unsafe.Add(mBase, uint32(v733)+100)) = v745
	v749 = v733 + int32(4)
	v750 = int32(12)
	v752 = v749 + v745*v750
	*(*int32)(unsafe.Add(mBase, uint32(v752)+4)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v752))) = v738
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v733)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v749+v755*v750)+8)) = int32(0)
	v778 = v733
	goto L80
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v762))) = v730
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v765)))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v765)+24))
	v769 = base.I32_div_s(v767, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v762)+72)) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v762)+68)) = v769
	v774 = v762 + v769<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v774)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v774)+4)) = v766
	v778 = v762
	goto L80
L87:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v797 = F_read_stream_next_buffer(m, v790, v34+int32(40))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	if v797 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v803 = v797
	v819 = v2
	goto L93
L91:
	;
	v1922 = v2
	goto L92
L92:
	;
	F_read_stream_end(m, v790)
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L8
	} else {
		goto L267
	}
L93:
	;
	if v803 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v1922 = v1897
	goto L92
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v848
	v851 = v34 + int32(48)
	v852 = int32(0)
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+4))
	v856 = int32(*(*int8)(unsafe.Add(mBase, uint32(v855)+1)))
	if v856 != 0 {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[10]))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v833+(v803^int32(-1))<<(uint(int32(6))%32))+16))
	v848 = v839
	goto L95
L97:
	;
	goto L98
L98:
	;
	v841 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[11]))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v841+v803<<(uint(int32(6))%32)+int32(-64))+16))
	v848 = v847
	goto L95
L99:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_visibilitymap_pin(m, v1019, v848, v34+int32(24))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L8
	} else {
		goto L121
	}
L100:
	;
	v869 = v856
	v870 = v852
	v873 = v852
	goto L106
L101:
	;
	if int32(0) < v856 {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v860 = int32(0)
	v861 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v855)+2)))
	if v861 == v860 {
		v1018 = v860
		goto L99
	} else {
		goto L105
	}
L104:
	;
	v1018 = int32(0)
	goto L99
L105:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v851))) = uint16(v861)
	v1018 = int32(1)
	goto L99
L106:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v855+int32(4)+v873<<(uint(int32(2))%32))))
	if v902 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v1018 = v954
	goto L99
L108:
	;
	v906 = v873 << (uint(int32(5)) % 32)
	v907 = v870
	v913 = v902
	goto L111
L109:
	;
	v953 = v869
	v954 = v870
	goto L110
L110:
	;
	v984 = v873 + int32(1)
	if v984 < base.I32_extend8_s(v953) {
		v869 = v953
		v870 = v954
		v873 = v984
		goto L106
	} else {
		goto L120
	}
L111:
	;
	if v913&int32(1) != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v855)+1)))
	v953 = v951
	v954 = v946
	goto L110
L113:
	;
	if v907 < int32(2048) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v946 = v907
	goto L115
L115:
	;
	v947 = int32(1)
	v950 = int32(base.Ui32(v913) >> (uint(v947) % 32))
	if v950 != 0 {
		v906 = v906 + v947
		v907 = v946
		v913 = v950
		goto L111
	} else {
		goto L119
	}
L116:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v851+v907<<(uint(int32(1))%32)))) = uint16(v906)
	goto L118
L117:
	;
	goto L118
L118:
	;
	v946 = v907 + int32(1)
	goto L115
L119:
	;
	goto L112
L120:
	;
	goto L107
L121:
	;
	F_LockBuffer(m, v803, int32(2))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	v1027 = int32(0)
	v1028 = base.B2i32(v1027 <= v803)
	if v1028 == v1027 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v1052 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	if v1052 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[12]))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1032+(v803^int32(-1))<<(uint(int32(2))%32))))
	v1046 = v1038
	goto L123
L125:
	;
	goto L126
L126:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[13]))
	v1046 = v1040 + v803<<(uint(int32(13))%32) + int32(-8192)
	goto L123
L127:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(3)
	v1096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v1097 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1097)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v848
	v1102 = int32(_a_F_lazy_vacuum_1)
	v1104 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v1104 + int32(1)
	if v1097 < v1018 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	goto L127
L129:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[5])))
	if v1056&int32(1) == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v1061 = int32(_a_F_lazy_vacuum_1)
	v1063 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	v1064 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v1063 + v1064
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1052)))
	*(*int32)(unsafe.Add(mBase, uint32(v1052))) = v1067 + v1064
	v1071 = int32(0)
	v1073 = int32(_a_F_lazy_vacuum_2)
	v1074 = base.AtomicRmwOr32(m, v1071, v1073, v1071)
	*(*int64)(unsafe.Add(mBase, uint32(v1052+int32(24))+232)) = base.I64_extend_i32_u(v848)
	v1082 = base.AtomicRmwOr32(m, v1071, v1073, v1071)
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1052)))
	*(*int32)(unsafe.Add(mBase, uint32(v1052))) = v1083 + v1064
	v1089 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v1089 - v1064
	goto L128
L131:
	;
	v1111 = v1018 & int32(3)
	v1113 = v1046 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v1018) {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v1324 = v1097
	goto L133
L133:
	;
	v1354 = int32(0)
	v1360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046)+12)))
	if base.Ui32(v1360) < base.Ui32(int32(25)) {
		goto L146
	} else {
		goto L147
	}
L134:
	;
	v1324 = v1018
	goto L133
L135:
	;
	v1121 = v1097
	v1125 = int32(0)
	goto L138
L136:
	;
	v1211 = v1097
	goto L137
L137:
	;
	v1242 = v1211
	v1256 = int32(0)
	goto L142
L138:
	;
	v1152 = v1121 << (uint(int32(1)) % 32)
	v1154 = v34 + int32(48)
	v1156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1152+v1154))))
	v1157 = int32(2)
	v1160 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1113+v1156<<(uint(v1157)%32)))) = v1160
	v1163 = v34 + int32(_a_F_lazy_vacuum_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1163+v1152))) = uint16(v1156)
	v1167 = v1152 | v1157
	v1169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1154+v1167))))
	*(*int32)(unsafe.Add(mBase, uint32(v1113+v1169<<(uint(v1157)%32)))) = v1160
	*(*uint16)(unsafe.Add(mBase, uint32(v1167+v1163))) = uint16(v1169)
	v1177 = int32(4)
	v1178 = v1152 | v1177
	v1180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1154+v1178))))
	*(*int32)(unsafe.Add(mBase, uint32(v1113+v1180<<(uint(v1157)%32)))) = v1160
	*(*uint16)(unsafe.Add(mBase, uint32(v1163+v1178))) = uint16(v1180)
	v1191 = v1152 | int32(6)
	v1193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1154+v1191))))
	*(*int32)(unsafe.Add(mBase, uint32(v1113+v1193<<(uint(v1157)%32)))) = v1160
	*(*uint16)(unsafe.Add(mBase, uint32(v1163+v1191))) = uint16(v1193)
	v1204 = v1121 + v1177
	v1206 = v1125 + v1177
	if v1206 != v1018&int32(2147483644) {
		v1121 = v1204
		v1125 = v1206
		goto L138
	} else {
		goto L140
	}
L139:
	;
	if v1111 == int32(0) {
		goto L134
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v1211 = v1204
	goto L137
L142:
	;
	v1272 = int32(1)
	v1273 = v1242 << (uint(v1272) % 32)
	v1277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1273+(v34+int32(48))))))
	*(*int32)(unsafe.Add(mBase, uint32(v1113+v1277<<(uint(int32(2))%32)))) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_lazy_vacuum_4)+v1273))) = uint16(v1277)
	v1290 = v1256 + v1272
	if v1290 != v1111 {
		v1242 = v1242 + v1272
		v1256 = v1290
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L134
L144:
	;
	goto L143
L145:
	;
	F_MarkBufferDirty(m, v803)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L8
	} else {
		goto L162
	}
L146:
	;
	v1425 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046)+10)))
	v1427 = v1425 & int32(_a_F_lazy_vacuum_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v1046)+10)) = uint16(v1427)
	goto L145
L147:
	;
	v1368 = int32(base.Ui32(v1360+int32(_a_F_lazy_vacuum_6))>>(uint(int32(2))%32)) & int32(_a_F_lazy_vacuum_7)
	if v1368 == int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v1374 = v1368
	v1376 = v1354
	v1379 = v1354
	goto L150
L149:
	;
	if int32(0) < v1403 {
		goto L158
	} else {
		goto L159
	}
L150:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1046+int32(20)+v1374<<(uint(int32(2))%32))))
	v1385 = v1383 & int32(_a_F_lazy_vacuum_8)
	if base.B2i32(v1374 == int32(1))|v1379 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v1403 = v1397
	v1405 = int32(0)
	goto L149
L152:
	;
	v1400 = v1374 - int32(1)
	if v1400 != 0 {
		v1374 = v1400
		v1376 = v1397
		v1379 = v1398
		goto L150
	} else {
		goto L157
	}
L153:
	;
	v1391 = int32(0)
	v1397 = v1376 + base.B2i32(v1385 == v1391)
	v1398 = base.B2i32(v1385 != v1391)
	goto L152
L154:
	;
	goto L155
L155:
	;
	if v1385 != 0 {
		v1397 = v1376
		v1398 = v1379
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v1403 = v1376
	v1405 = int32(1)
	goto L149
L157:
	;
	goto L151
L158:
	;
	v1410 = v1360 - v1403<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1046)+12)) = uint16(v1410)
	goto L160
L159:
	;
	goto L160
L160:
	;
	if v1405 == int32(0) {
		goto L146
	} else {
		goto L161
	}
L161:
	;
	v1414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046)+10)))
	v1416 = v1414 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1046)+10)) = uint16(v1416)
	goto L145
L162:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+48))
	v1440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1439)+118)))
	if v1440 != int32(112) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v1462 = int32(_a_F_lazy_vacuum_1)
	v1464 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v1464 - int32(1)
	if v1028 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L164:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[14]))
	if v1444 <= int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+32))
	if v1447 != 0 {
		goto L163
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v1449 = int32(0)
	F_log_heap_prune_and_freeze(m, v1438, v803, v1449, v1449, int32(2), v1449, v1449, v1449, v1449, v1449, v1449, v34+int32(_a_F_lazy_vacuum_4), v1324)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L8
	} else {
		goto L170
	}
L168:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+40))
	if v1448 != 0 {
		goto L163
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	goto L163
L171:
	;
	if v803 < int32(0) {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[12]))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1471+(v803^int32(-1))<<(uint(int32(2))%32))))
	v1485 = v1477
	goto L171
L173:
	;
	goto L174
L174:
	;
	v1479 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[13]))
	v1485 = v1479 + v803<<(uint(int32(13))%32) + int32(-8192)
	goto L171
L175:
	;
	v1505 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1485)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1505) {
		goto L182
	} else {
		goto L183
	}
L176:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[10]))
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1489+(v803^int32(-1))<<(uint(int32(6))%32))+16))
	v1504 = v1495
	goto L175
L177:
	;
	goto L178
L178:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[11]))
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1497+v803<<(uint(int32(6))%32)+int32(-64))+16))
	v1504 = v1503
	goto L175
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v1093
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1096)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1100
	if v1028 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L180:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v1765 + int32(1)
	goto L179
L181:
	;
	v1532 = int32(base.Ui32(v1504) >> (uint(int32(16)) % 32))
	v1535 = int32(1)
	v1539 = v1535
	v1543 = int32(0)
	v1546 = v1535
	goto L187
L182:
	;
	v1513 = int32(base.Ui32(v1505+int32(_a_F_lazy_vacuum_6))>>(uint(int32(2))%32)) & int32(_a_F_lazy_vacuum_7)
	if v1513 != 0 {
		goto L181
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v1515 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1515)
	v1517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046)+10)))
	v1519 = v1517 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1046)+10)) = uint16(v1519)
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1525 = F_visibilitymap_set(m, v1521, v848, v803, int64(0), v1047, v1515, int32(3))
	mBase = m.M
	v1526 = m.ExcPending
	if v1526 != 0 {
		goto L8
	} else {
		goto L186
	}
L185:
	;
	goto L184
L186:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1527 + int32(1)
	goto L180
L187:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1539)
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1485+int32(20)+v1539&int32(_a_F_lazy_vacuum_7)<<(uint(int32(2))%32))))
	switch int32(base.Ui32(v1575)>>(uint(int32(15))%32)) & int32(3) {
	case 0, 2:
		v1704 = v1543
		v1705 = v1546
		goto L189
	default:
		goto L190
	}
L188:
	;
	v1713 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1713)
	v1715 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046)+10)))
	v1717 = v1715 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1046)+10)) = uint16(v1717)
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1722 = int32(1)
	v1724 = v1705 & v1722
	if v1724 != 0 {
		goto L234
	} else {
		goto L235
	}
L189:
	;
	v1709 = v1539 + int32(1)
	if base.Ui32(v1709&int32(_a_F_lazy_vacuum_7)) <= base.Ui32(v1513) {
		v1539 = v1709
		v1543 = v1704
		v1546 = v1705
		goto L187
	} else {
		goto L233
	}
L190:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[3]))) = uint16(v1539)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[15]))) = uint16(v1504)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[16]))) = uint16(v1532)
	v1583 = int32(_a_F_lazy_vacuum_8)
	if v1575&v1583 == v1583 {
		goto L179
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[2]))) = int32(base.Ui32(v1575) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[17]))) = v1485 + v1575&int32(_a_F_lazy_vacuum_9)
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1594)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[18]))) = v1595
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1600 = F_HeapTupleSatisfiesVacuum(m, v34+int32(_a_F_lazy_vacuum_3), v1599, v803)
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L8
	} else {
		goto L192
	}
L192:
	;
	if v1600 != int32(1) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if base.B2i32(v1600 != int32(1))&base.B2i32(base.Ui32(v1600) <= base.Ui32(int32(4))) != 0 {
		goto L179
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[17])))
	v1623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1622)+20)))
	if v1623&int32(256) == int32(0) {
		goto L179
	} else {
		goto L200
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L8
	} else {
		goto L197
	}
L197:
	;
	F_errmsg_internal(m, int32(_a_F_lazy_vacuum_10), int32(0))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L8
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_lazy_vacuum_11), int32(3708), int32(_a_F_lazy_vacuum_12))
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L8
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	v1629 = int32(768)
	if v1623&v1629 != v1629 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1622)))
	v1634 = v1633
	goto L203
L202:
	;
	v1634 = int32(2)
	goto L203
L203:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1635))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1634)) == int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	if v1647 == int32(0) {
		goto L179
	} else {
		goto L208
	}
L205:
	;
	v1647 = base.B2i32(base.Ui32(v1634) < base.Ui32(v1635))
	goto L204
L206:
	;
	goto L207
L207:
	;
	v1647 = int32(base.Ui32(v1634-v1635) >> (uint(int32(31)) % 32))
	goto L204
L208:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1543))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1634)) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	if v1661 != 0 {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v1661 = base.B2i32(base.Ui32(v1543) < base.Ui32(v1634))
	goto L209
L211:
	;
	goto L212
L212:
	;
	v1661 = base.B2i32(int32(0) < v1634-v1543)
	goto L209
L213:
	;
	v1662 = v1634
	goto L215
L214:
	;
	v1662 = v1543
	goto L215
L215:
	;
	if base.Ui32(int32(2)) < base.Ui32(v1634) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1665 = v1662
	goto L218
L217:
	;
	v1665 = v1543
	goto L218
L218:
	;
	v1668 = int32(0)
	if v1546&int32(1) == v1668 {
		v1704 = v1665
		v1705 = v1668
		goto L189
	} else {
		goto L219
	}
L219:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[17])))
	v1674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1671)+20)))
	v1675 = int32(768)
	if v1674&v1675 == v1675 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v1704 = v1665
	v1705 = v1700 ^ int32(1)
	goto L189
L221:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1671)+4))
	if v1674&int32(_a_F_lazy_vacuum_13) != 0 {
		goto L225
	} else {
		goto L226
	}
L222:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1671)))
	if base.Ui32(v1679) <= base.Ui32(int32(2)) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1700 = int32(1)
	goto L220
L224:
	;
	if base.Ui32(v1674) < base.Ui32(int32(_a_F_lazy_vacuum_14)) {
		goto L230
	} else {
		goto L231
	}
L225:
	;
	if v1683 == int32(0) {
		goto L224
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	if base.Ui32(v1683) <= base.Ui32(int32(2)) {
		goto L224
	} else {
		goto L229
	}
L228:
	;
	v1700 = int32(1)
	goto L220
L229:
	;
	v1700 = int32(1)
	goto L220
L230:
	;
	v1700 = int32(0)
	goto L220
L231:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1671)+8))
	if base.Ui32(v1694) <= base.Ui32(int32(2)) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v1700 = int32(1)
	goto L220
L233:
	;
	goto L188
L234:
	;
	v1725 = int32(3)
	goto L236
L235:
	;
	v1725 = v1722
	goto L236
L236:
	;
	v1726 = F_visibilitymap_set(m, v1719, v848, v803, int64(0), v1047, v1704, v1725)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L8
	} else {
		goto L237
	}
L237:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1728 + int32(1)
	if v1724 == int32(0) {
		goto L179
	} else {
		goto L238
	}
L238:
	;
	goto L180
L239:
	;
	v1824 = int32(4)
	v1825 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1820)+14)))
	v1826 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1820)+12)))
	v1827 = v1825 - v1826
	if v1827 <= v1824 {
		goto L244
	} else {
		goto L245
	}
L240:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[12]))
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1806+(v803^int32(-1))<<(uint(int32(2))%32))))
	v1820 = v1812
	goto L239
L241:
	;
	goto L242
L242:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[13]))
	v1820 = v1814 + v803<<(uint(int32(13))%32) + int32(-8192)
	goto L239
L243:
	;
	F_UnlockReleaseBuffer(m, v803)
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L8
	} else {
		goto L262
	}
L244:
	;
	v1830 = v1824
	goto L246
L245:
	;
	v1830 = v1827
	goto L246
L246:
	;
	v1832 = v1830 - int32(4)
	if v1832 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1887 = int32(0)
	goto L243
L248:
	;
	goto L249
L249:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1826) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v1887 = v1832
	goto L243
L251:
	;
	v1843 = int32(base.Ui32(v1826+int32(_a_F_lazy_vacuum_6)) >> (uint(int32(2)) % 32))
	goto L253
L252:
	;
	v1843 = int32(0)
	goto L253
L253:
	;
	if base.Ui32(v1843&int32(_a_F_lazy_vacuum_7)) < base.Ui32(int32(291)) {
		goto L250
	} else {
		goto L254
	}
L254:
	;
	v1848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1820)+10)))
	if v1848&int32(1) == int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1887 = int32(0)
	goto L243
L256:
	;
	goto L257
L257:
	;
	v1857 = int32(1)
	goto L258
L258:
	;
	v1866 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1820+int32(20)+v1857&int32(_a_F_lazy_vacuum_7)<<(uint(int32(2))%32))+1)))
	if v1866&int32(384) == int32(0) {
		goto L250
	} else {
		goto L260
	}
L259:
	;
	v1887 = int32(0)
	goto L243
L260:
	;
	v1872 = v1857 + int32(1)
	v1873 = int32(_a_F_lazy_vacuum_7)
	if base.Ui32(v1872&v1873) <= base.Ui32(v1843&v1873) {
		v1857 = v1872
		goto L258
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_RecordPageWithFreeSpace(m, v1890, v848, v1887)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L8
	} else {
		goto L263
	}
L263:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1895 = m.ExcPending
	if v1895 != 0 {
		goto L8
	} else {
		goto L264
	}
L264:
	;
	v1897 = v819 + int32(1)
	v1900 = F_read_stream_next_buffer(m, v790, v34+int32(40))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L8
	} else {
		goto L265
	}
L265:
	;
	if v1900 != 0 {
		v803 = v1900
		v819 = v1897
		goto L93
	} else {
		goto L266
	}
L266:
	;
	goto L94
L267:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	F_pfree(m, v1935)
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L8
	} else {
		goto L268
	}
L268:
	;
	F_pfree(m, v727)
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L8
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v1942 != 0 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	F_ReleaseBuffer(m, v1942)
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L8
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1947 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L8
	} else {
		goto L274
	}
L273:
	;
	goto L272
L274:
	;
	if v1947 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1951 = *(*int64)(unsafe.Add(mBase, uint32(v1950)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v1922
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v1951
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v1949
	F_errmsg(m, int32(_a_F_lazy_vacuum_15), v34)
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L8
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v716
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v719)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v722
	goto L13
L278:
	;
	F_errfinish(m, int32(_a_F_lazy_vacuum_11), int32(2823), int32(_a_F_lazy_vacuum_16))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L8
	} else {
		goto L279
	}
L279:
	;
	goto L277
L280:
	;
	F_parallel_vacuum_reset_dead_items(m, v1999)
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L8
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_TidStoreDestroy(m, v2011)
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L8
	} else {
		goto L285
	}
L283:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v2002)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(104)))) = v2005 + int32(56)
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v2002)+24))
	goto L284
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v2009
	goto L1
L285:
	;
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v2014)))
	v2016 = F_TidStoreCreateLocal(m, v2015)
	mBase = m.M
	v2017 = m.ExcPending
	if v2017 != 0 {
		goto L8
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v2016
	v2019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v2019)+8)) = int64(0)
	goto L1
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
func F_leftmostvalue_inet(m *base.Module) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_DirectFunctionCall1Coll(m, int32(1466), int32(0), int32(_a_F_leftmostvalue_inet_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_levenshtein_less_equal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v24 = v22 & int32(1)
			if v24 != 0 {
				v25 = v14
			} else {
				v25 = v9 + int32(4)
			}
			if v22 == int32(1) {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v31 == int32(18) {
					v34 = int32(16)
				} else {
					v34 = int32(0)
				}
				if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v41 = int32(4)
				} else {
					v41 = v34
				}
				v52 = v41
			} else {
				v42 = int32(1)
				if v24 != 0 {
					v52 = int32(base.Ui32(v22)>>(uint(v42)%32)) - v42
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v53 = int32(1)
			v54 = v16 + v53
			if v18&v53 != 0 {
				v59 = v54
			} else {
				v59 = v16 + int32(4)
			}
			if v18 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
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
				if v18&v76 != 0 {
					v88 = int32(base.Ui32(v18)>>(uint(v76)%32)) - v76
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v89 = int32(1)
			v93 = F_varstr_levenshtein_less_equal(m, v25, v52, v59, v88, v89, v89, v89, v19, int32(0))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				return v93
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int64
	_ = v413
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l0 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L24
	} else {
		goto L112
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L24
	} else {
		goto L108
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[0]))
	if v22 <= l0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[1]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+l0<<(uint(int32(2))%32))))
	if v29 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+24)))
	v34 = v32 & int32(2)
	if v34 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v37 = m.G0
	v39 = v37 - int32(2224)
	m.G0 = v39
	v42 = base.I64_div_s(l1, int64(2048))
	if v34 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	m.G0 = v17 + int32(32)
	return
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L24
	} else {
		goto L104
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L24
	} else {
		goto L101
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
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
	v361 = m.ExcPending
	if v361 != 0 {
		goto L24
	} else {
		goto L93
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[3]))
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = v46
	goto L17
L16:
	;
	v50 = int32(0)
	goto L17
L17:
	;
	if v50 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v53 = int32(_a_F_lo_truncate_internal_0)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[4]))
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[4])) = v57
	if v46 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v82 = v46
	goto L20
L20:
	;
	v84 = base.I32_wrap_i64(v42)
	v86 = v39 + int32(80)
	v87 = F_CatalogOpenIndexes(m, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L30
	}
L21:
	;
	v67 = v46
	v68 = v49
	goto L23
L22:
	;
	v62 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v68 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	return
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2])) = v62
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[3]))
	v67 = v62
	v68 = v66
	goto L23
L26:
	;
	v74 = F_index_open(m, int32(2683), int32(3))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v79 = v67
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[4])) = v54
	v82 = v79
	goto L20
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[3])) = v74
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v79 = v78
	goto L28
L30:
	;
	v90 = v39 + int32(2128)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	F_ScanKeyInit(m, v90, int32(1), int32(3), int32(184), v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	F_ScanKeyInit(m, v39+int32(2176), int32(2), int32(4), int32(150), v84)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[3]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v110 = F_systable_beginscan_ordered(m, v105, v107, v108, int32(2), v90)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L24
	} else {
		goto L35
	}
L33:
	;
	F_systable_endscan_ordered(m, v110)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L24
	} else {
		goto L90
	}
L34:
	;
	v308 = F_systable_getnext_ordered(m, v110, int32(1))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L24
	} else {
		goto L83
	}
L35:
	;
	v113 = F_systable_getnext_ordered(m, v110, int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L24
	} else {
		goto L36
	}
L36:
	;
	if v113 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+20)))
	if v116&int32(1) != 0 {
		goto L9
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v236 = base.I32_wrap_i64(l1)
	v238 = v236 & int32(2047)
	if v238 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L40:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+22)))
	v120 = v115 + v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if v84 == v121 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v124 = v120 + int32(8)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)))
	v127 = v125 & int32(3)
	if v127 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	F_simple_heap_delete(m, v229, v113+int32(4))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L24
	} else {
		goto L69
	}
L44:
	;
	v128 = F_detoast_attr(m, v124)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L24
	} else {
		goto L47
	}
L45:
	;
	v130 = v124
	goto L46
L46:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v133 = int32(base.Ui32(v131) >> (uint(int32(2)) % 32))
	v135 = v133 - int32(4)
	if base.Ui32(v131-int32(_a_F_lo_truncate_internal_1)) <= base.Ui32(int32(-8197)) {
		goto L8
	} else {
		goto L48
	}
L47:
	;
	v130 = v128
	goto L46
L48:
	;
	if v135 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	base.MemoryCopy(m, v86, v130+int32(4), v135)
	goto L51
L50:
	;
	goto L51
L51:
	;
	if v127 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_pfree(m, v130)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L24
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v147 = base.I32_wrap_i64(l1) & int32(2047)
	if v147 <= v135 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+64)) = int64(0)
	v193 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v193)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+62)) = uint8(v193)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+56)) = uint16(v193)
	v199 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+58)) = uint8(v199)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v147<<(uint(int32(2))%32) + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v39 + int32(76)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+52))
	v218 = F_heap_modify_tuple(m, v113, v211, v39-int32(-64), v39+int32(60), v39+int32(56))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L24
	} else {
		goto L66
	}
L57:
	;
	v150 = v39 + int32(76)
	v151 = v150 + v133
	v152 = int32(3)
	v154 = v147 - v135
	if v151&v152|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v154))|v154&v152 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if base.Ui32(v147+v86) <= base.Ui32(v151) {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	v181 = v154
	goto L60
L60:
	;
	if v181 == int32(0) {
		goto L56
	} else {
		goto L65
	}
L61:
	;
	v168 = int32(80)
	v169 = v39 + v133 + v168
	v172 = v39 + v147 + v168
	if base.Ui32(v172) < base.Ui32(v169) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v174 = v169
	goto L64
L63:
	;
	v174 = v172
	goto L64
L64:
	;
	v181 = (v150^int32(-1)+v174-v133)&int32(-4) + int32(4)
	goto L60
L65:
	;
	base.MemoryFill(m, v151, int32(0), v181)
	goto L56
L66:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	F_CatalogTupleUpdateWithInfo(m, v221, v218+int32(4), v218, v87)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L24
	} else {
		goto L67
	}
L67:
	;
	F_pfree(m, v218)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
	v268 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v268)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+62)) = uint8(v268)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v238<<(uint(int32(2))%32) + int32(16)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v39 + int32(76)
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+52))
	v290 = F_heap_form_tuple(m, v285, v39-int32(-64), v39+int32(60))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L24
	} else {
		goto L79
	}
L71:
	;
	if v236&int32(3) != 0 {
		v261 = v238
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v261 == int32(0) {
		goto L70
	} else {
		goto L78
	}
L73:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v238) {
		v261 = v238
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v247 = v39 + v238 + int32(80)
	v249 = v39 + int32(84)
	if base.Ui32(v249) < base.Ui32(v247) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v251 = v247
	goto L77
L76:
	;
	v251 = v249
	goto L77
L77:
	;
	v261 = (v251-v39-int32(81))&int32(-4) + int32(4)
	goto L72
L78:
	;
	base.MemoryFill(m, v86, int32(0), v261)
	goto L70
L79:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	F_CatalogTupleInsertWithInfo(m, v293, v290, v87)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L24
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v290)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L24
	} else {
		goto L81
	}
L81:
	;
	if v113 == int32(0) {
		goto L33
	} else {
		goto L82
	}
L82:
	;
	goto L34
L83:
	;
	if v308 == int32(0) {
		goto L33
	} else {
		goto L84
	}
L84:
	;
	v315 = v308
	goto L85
L85:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_lo_truncate_internal[2]))
	F_simple_heap_delete(m, v327, v315+int32(4))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L24
	} else {
		goto L87
	}
L86:
	;
	goto L33
L87:
	;
	v333 = F_systable_getnext_ordered(m, v110, int32(1))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L24
	} else {
		goto L88
	}
L88:
	;
	if v333 != 0 {
		v315 = v333
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	F_CatalogCloseIndexes(m, v87)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L24
	} else {
		goto L91
	}
L91:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L24
	} else {
		goto L92
	}
L92:
	;
	m.G0 = v39 + int32(2224)
	goto L7
L93:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v365
	F_errmsg(m, int32(_a_F_lo_truncate_internal_2), v39)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L24
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_3), int32(770), int32(_a_F_lo_truncate_internal_4))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
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
	v381 = m.ExcPending
	if v381 != 0 {
		goto L24
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_lo_truncate_internal_5), v39+int32(16))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L24
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_3), int32(780), int32(_a_F_lo_truncate_internal_4))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_lo_truncate_internal_6), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L24
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_3), int32(810), int32(_a_F_lo_truncate_internal_4))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
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
	v412 = m.ExcPending
	if v412 != 0 {
		goto L24
	} else {
		goto L105
	}
L105:
	;
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v120)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v135
	*(*int64)(unsafe.Add(mBase, uint32(v39)+32)) = v413
	F_errmsg(m, int32(_a_F_lo_truncate_internal_7), v39+int32(32))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L24
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_3), int32(153), int32(_a_F_lo_truncate_internal_8))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
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
	v436 = m.ExcPending
	if v436 != 0 {
		goto L24
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	F_errmsg(m, int32(_a_F_lo_truncate_internal_9), v17)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L24
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_10), int32(565), int32(_a_F_lo_truncate_internal_11))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
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
	v452 = m.ExcPending
	if v452 != 0 {
		goto L24
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l0
	F_errmsg(m, int32(_a_F_lo_truncate_internal_12), v17+int32(16))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L24
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_lo_truncate_internal_10), int32(573), int32(_a_F_lo_truncate_internal_11))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(96)
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
	v43 = F_SearchSysCache1(m, int32(82), v22)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	F_ReleaseCatCache(m, v54)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L87
	}
L8:
	;
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v46 = v22
	v48 = v2
	v53 = v2
	v54 = v43
	v55 = v2
	v60 = v2
	goto L12
L10:
	;
	v325 = v22
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L84
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
	v64 = v62 + v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+79)))
	if v65 != int32(100) {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v325 = v318
	goto L11
L14:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+130)))
	v70 = v20 + int32(48)
	F_ScanKeyInit(m, v70, int32(10), int32(3), int32(184), v46)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v76 = int32(0)
	v78 = int32(1)
	v81 = F_systable_beginscan(m, v40, int32(2666), v78, v76, v78, v70)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v64)+132))
	F_ReleaseCatCache(m, v54)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L81
	}
L17:
	;
	v83 = F_systable_getnext(m, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v83 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_systable_endscan(m, v81)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v90 = v76
	v92 = v48
	v93 = v83
	v97 = v53
	v99 = v55
	goto L23
L22:
	;
	v303 = v48
	v308 = v53
	v310 = v55
	goto L16
L23:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+22)))
	v108 = v106 + v107
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+72)))
	if v109 == int32(99) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_systable_endscan(m, v81)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L71
	}
L25:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+20)))
	if v113&int32(1) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v244 = v90
	v246 = v92
	v250 = v97
	v251 = v99
	goto L27
L27:
	;
	v252 = F_systable_getnext(m, v81)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L69
	}
L28:
	;
	v179 = F_text_to_cstring(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L51
	}
L29:
	;
	v175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v121))))
	v178 = v175
	goto L28
L30:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112)+452))
	if int32(0) <= v118 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+26)))
	if v149&int32(8) != 0 {
		goto L44
	} else {
		goto L45
	}
L33:
	;
	v121 = v118 + v108
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+458)))
	if v122 != int32(1) {
		v178 = v121
		goto L28
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v147 = F_nocachegetattr(m, v93, int32(28), v112)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L43
	}
L36:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+456)))
	switch v125 - int32(1) {
	case 0:
		goto L29
	case 1:
		goto L39
	default:
		goto L37
	case 3:
		goto L38
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v178 = v129
	goto L28
L39:
	;
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121))))
	v178 = v128
	goto L28
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = base.I32_extend16_s(v125)
	F_errmsg_internal(m, int32(_a_F_load_domaintype_info_0), v20+int32(16))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_load_domaintype_info_1), int32(70), int32(_a_F_load_domaintype_info_2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v178 = v147
	goto L28
L44:
	;
	v153 = F_nocachegetattr(m, v93, int32(28), v112)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L48
	}
L47:
	;
	v178 = v153
	goto L28
L48:
	;
	v159 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v108 + v159
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v64 + v159
	F_errmsg_internal(m, int32(_a_F_load_domaintype_info_3), v20+int32(32))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_load_domaintype_info_4), int32(1172), int32(_a_F_load_domaintype_info_5))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	if v92 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v200 = int32(_a_F_load_domaintype_info_6)
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v199
	v204 = F_stringToNode(m, v179)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L58
	}
L53:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v198 = v92
	v199 = v181
	goto L52
L54:
	;
	goto L55
L55:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	v188 = F_AllocSetContextCreateInternal(m, v183, int32(_a_F_load_domaintype_info_7), int32(0), int32(1024), int32(_a_F_load_domaintype_info_8))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v191 = F_MemoryContextAlloc(m, v188, int32(12))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v193 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v191)+4)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v193
	v198 = v191
	v199 = v188
	goto L52
L58:
	;
	v206 = F_expression_planner(m, v204)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v209 = F_palloc0(m, int32(20))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v209))) = int64(4294967689)
	v215 = F_pstrdup(m, v108+int32(4))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v217 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v209)+16)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v209)+12)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = v215
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v201
	if v97 == v217 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v236+v90<<(uint(int32(2))%32)))) = v209
	v244 = v90 + int32(1)
	v246 = v198
	v250 = v236
	v251 = v237
	goto L27
L63:
	;
	v227 = F_palloc(m, int32(32))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v90 < v99 {
		v236 = v97
		v237 = v99
		goto L62
	} else {
		goto L67
	}
L66:
	;
	v236 = v227
	v237 = int32(8)
	goto L62
L67:
	;
	v232 = F_repalloc(m, v97, v99<<(uint(int32(3))%32))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v236 = v232
	v237 = v99 << (uint(int32(1)) % 32)
	goto L62
L69:
	;
	if v252 != 0 {
		v90 = v244
		v92 = v246
		v93 = v252
		v97 = v250
		v99 = v251
		goto L23
	} else {
		goto L70
	}
L70:
	;
	goto L24
L71:
	;
	if v244 <= int32(0) {
		v303 = v246
		v308 = v250
		v310 = v251
		goto L16
	} else {
		goto L72
	}
L72:
	;
	if v244 != int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_pg_qsort(m, v250, v244, int32(4), int32(1603))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v264 = int32(_a_F_load_domaintype_info_6)
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v271 = v244
	v275 = v269
	goto L77
L76:
	;
	goto L75
L77:
	;
	v288 = v271 - int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v250+v288<<(uint(int32(2))%32))))
	v293 = F_lcons(m, v292, v275)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v265
	v303 = v246
	v308 = v250
	v310 = v251
	goto L16
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = v293
	if base.Ui32(int32(1)) < base.Ui32(v271) {
		v271 = v288
		v275 = v293
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v322 = F_SearchSysCache1(m, int32(82), v318)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	if v322 != 0 {
		v46 = v318
		v48 = v303
		v53 = v308
		v54 = v322
		v55 = v310
		v60 = v60 | v68
		goto L12
	} else {
		goto L83
	}
L83:
	;
	goto L13
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v325
	F_errmsg_internal(m, int32(_a_F_load_domaintype_info_9), v20)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_load_domaintype_info_4), int32(1131), int32(_a_F_load_domaintype_info_5))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
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
	F_relation_close(m, v40, int32(1))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	if v60&int32(1) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v451 | int32(_a_F_load_domaintype_info_10)
	m.G0 = v20 + int32(96)
	return
L90:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[1]))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	if v413 != v409 {
		goto L105
	} else {
		goto L106
	}
L91:
	;
	if v48 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	goto L93
L93:
	;
	if v48 == int32(0) {
		goto L89
	} else {
		goto L103
	}
L94:
	;
	v380 = int32(_a_F_load_domaintype_info_6)
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v378
	v385 = F_palloc0(m, int32(20))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L100
	}
L95:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v378 = v361
	v379 = v48
	goto L94
L96:
	;
	goto L97
L97:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0]))
	v368 = F_AllocSetContextCreateInternal(m, v363, int32(_a_F_load_domaintype_info_7), int32(0), int32(1024), int32(_a_F_load_domaintype_info_8))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v371 = F_MemoryContextAlloc(m, v368, int32(12))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v373 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+8)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v373
	v378 = v368
	v379 = v371
	goto L94
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v385))) = int64(393)
	v390 = F_pstrdup(m, int32(_a_F_load_domaintype_info_11))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v385)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v385)+8)) = v390
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v396 = F_lcons(m, v385, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379))) = v396
	*(*int32)(unsafe.Add(mBase, _c_F_load_domaintype_info[0])) = v381
	v405 = v379
	goto L90
L103:
	;
	v405 = v48
	goto L90
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v405
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v405)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v405)+8)) = v443 + int32(1)
	goto L89
L105:
	;
	if v413 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	if v409 != 0 {
		goto L115
	} else {
		goto L116
	}
L109:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v407)+28))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v407)+24))
	if v418 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v417 == int32(0) {
		goto L108
	} else {
		goto L114
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418)+28)) = v417
	goto L110
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v413)+20)) = v417
	goto L110
L114:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v407)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v417)+24)) = v423
	goto L108
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+16)) = v409
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v409)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v407)+28)) = v430
	if v430 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v407)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v407)+16)) = int32(0)
	goto L107
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430)+24)) = v407
	goto L120
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+20)) = v407
	goto L104
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
	F_errfinish(m, int32(_a_F_load_hba_4), int32(2708), int32(_a_F_load_hba_5))
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
	F_errfinish(m, int32(_a_F_load_libraries_1), int32(1872), int32(_a_F_load_libraries_2))
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
	v71 = Fn13880(m, v65, int32(47))
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
	F_errfinish(m, int32(_a_F_load_libraries_1), int32(1890), int32(_a_F_load_libraries_2))
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
				v33 = F_expression_tree_walker_impl(m, l0, int32(1044), l1)
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
					v33 = F_expression_tree_walker_impl(m, l0, int32(1044), l1)
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
				v33 = F_expression_tree_walker_impl(m, l0, int32(1044), l1)
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
					v33 = F_expression_tree_walker_impl(m, l0, int32(1044), l1)
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
				v44 = F_query_tree_walker_impl(m, l0, int32(1044), l1, int32(0))
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
				v33 = F_expression_tree_walker_impl(m, l0, int32(1044), l1)
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
func F_lquery_in(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_parse_lquery(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v12 = v5
		} else {
			v9 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
			v12 = int32(0)
		}
		return v12
	}
}
func F_ltrim(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13857(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
