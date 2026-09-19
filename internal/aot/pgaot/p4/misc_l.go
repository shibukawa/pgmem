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
	var v2161 int32
	_ = v2161
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2211 int32
	_ = v2211
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2329 int32
	_ = v2329
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2375 int32
	_ = v2375
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2386 int64
	_ = v2386
	var v2394 int32
	_ = v2394
	var v2397 int32
	_ = v2397
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2413 int32
	_ = v2413
	var v2417 int32
	_ = v2417
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2430 int32
	_ = v2430
	var v2435 int32
	_ = v2435
	var v2440 int32
	_ = v2440
	var v2446 int32
	_ = v2446
	var v2451 int32
	_ = v2451
	var v2455 int32
	_ = v2455
	var v2459 int32
	_ = v2459
	var v2464 int32
	_ = v2464
	var v2496 int32
	_ = v2496
	var v2527 int64
	_ = v2527
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2571 int64
	_ = v2571
	var v2576 int32
	_ = v2576
	var v2609 int32
	_ = v2609
	var v2612 int32
	_ = v2612
	var v2616 int32
	_ = v2616
	var v2623 int32
	_ = v2623
	var v2633 int32
	_ = v2633
	var v2666 int32
	_ = v2666
	var v2668 int32
	_ = v2668
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2685 int32
	_ = v2685
	var v2690 int32
	_ = v2690
	var v2693 int32
	_ = v2693
	var v2695 int32
	_ = v2695
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2702 int64
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2706 int32
	_ = v2706
	var v2714 int32
	_ = v2714
	var v2750 int32
	_ = v2750
	var v2753 int32
	_ = v2753
	var v2754 int32
	_ = v2754
	var v2758 int32
	_ = v2758
	var v2764 int32
	_ = v2764
	var v2768 int32
	_ = v2768
	var v2773 int32
	_ = v2773
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
	v2750 = m.ExcPending
	if v2750 != 0 {
		goto L23
	} else {
		goto L437
	}
L2:
	;
	m.G0 = v33 + int32(192)
	return v2714
L3:
	;
	v2527 = *(*int64)(unsafe.Add(mBase, uint32(v91)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+32)) = v2527 + int64(1)
	v2531 = *(*int32)(unsafe.Add(mBase, uint32(v91)+48))
	v2532 = int32(0)
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	if v2532 < v2533 {
		goto L413
	} else {
		goto L414
	}
L4:
	;
	F_LWLockRelease(m, v1098)
	mBase = m.M
	v2496 = m.ExcPending
	if v2496 != 0 {
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
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L23
	} else {
		goto L408
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
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
	v2714 = v280
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
	v2714 = int32(1)
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
	v2426 = m.ExcPending
	if v2426 != 0 {
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
		v2714 = v1874
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
	v2714 = v1874
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
	v2375 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[21])))
	if v2375 == int32(1) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+36)) = v2329
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+32)) = v2119
	F_appendStringInfo(m, v1891+int32(96), int32(_a_F_LockAcquireExtended_16), v1891+int32(32))
	mBase = m.M
	v2338 = m.ExcPending
	if v2338 != 0 {
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
	v2329 = int32(_a_F_LockAcquireExtended_17)
	goto L366
L370:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v2128)+4))
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v2128)))
	v2161 = int32(0)
	if base.B2i32(v2158&int32(1) == v2161)&base.B2i32(v2158 == v2158) == v2161 {
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
	v2211 = v2157
	goto L374
L374:
	;
	if v2211 == v2119 {
		goto L382
	} else {
		goto L383
	}
L375:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[24]))
	if v2199 != 0 {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	v2211 = v2202
	goto L374
L377:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L23
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v2202 = *(*int32)(unsafe.Add(mBase, uint32(v2128)+4))
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2128)))
	if v2203&int32(1)|base.B2i32(v2203 != v2203) != 0 {
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
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(v2128)+216))
	v2242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2241))))
	if v2242 == int32(0) {
		v2329 = int32(_a_F_LockAcquireExtended_18)
		goto L366
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	v2264 = v2131 + int32(1)
	v2266 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[22]))
	if v2264 <= v2266 {
		v2128 = v2128 + int32(408)
		v2131 = v2264
		goto L370
	} else {
		goto L388
	}
L385:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[25]))
	v2249 = F_pnstrdup(m, v2241, v2246-int32(1))
	mBase = m.M
	v2250 = m.ExcPending
	if v2250 != 0 {
		goto L23
	} else {
		goto L386
	}
L386:
	;
	v2251 = F_strlen(m, v2249)
	mBase = m.M
	v2253 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[25]))
	v2256 = F_pg_mbcliplen(m, v2249, v2251, v2253-int32(1))
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L23
	} else {
		goto L387
	}
L387:
	;
	v2259 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2256+v2249))) = uint8(v2259)
	v2329 = v2249
	goto L366
L388:
	;
	goto L371
L389:
	;
	v2340 = v2084 + int32(1)
	v2342 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[18]))
	if v2340 < v2342 {
		v2084 = v2340
		goto L363
	} else {
		goto L390
	}
L390:
	;
	goto L364
L391:
	;
	v2380 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[10]))
	v2383 = F_pgstat_prep_pending_entry(m, int32(1), v2380, int64(0), int32(0))
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
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
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L23
	} else {
		goto L395
	}
L394:
	;
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v2383)+12))
	v2386 = *(*int64)(unsafe.Add(mBase, uint32(v2385)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v2385)+144)) = v2386 + int64(1)
	goto L393
L395:
	;
	F_errcode(m, int32(16908292))
	mBase = m.M
	v2397 = m.ExcPending
	if v2397 != 0 {
		goto L23
	} else {
		goto L396
	}
L396:
	;
	F_errmsg(m, int32(_a_F_LockAcquireExtended_19), int32(0))
	mBase = m.M
	v2401 = m.ExcPending
	if v2401 != 0 {
		goto L23
	} else {
		goto L397
	}
L397:
	;
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v1891)+16)) = v2402
	F_errdetail_internal(m, int32(_a_F_LockAcquireExtended_20), v1891+int32(16))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L23
	} else {
		goto L398
	}
L398:
	;
	v2409 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v1891))) = v2409
	F_errdetail_log(m, int32(_a_F_LockAcquireExtended_20), v1891)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L23
	} else {
		goto L399
	}
L399:
	;
	F_errhint(m, int32(_a_F_LockAcquireExtended_21), int32(0))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L23
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_22), int32(1138), int32(_a_F_LockAcquireExtended_23))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_LockAcquireExtended_24), int32(0))
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L23
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(1140), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2435 = m.ExcPending
	if v2435 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_LockAcquireExtended_25), v33+int32(16))
	mBase = m.M
	v2446 = m.ExcPending
	if v2446 != 0 {
		goto L23
	} else {
		goto L406
	}
L406:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(861), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_LockAcquireExtended_26), v33)
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L23
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(858), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
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
	v2666 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[15])) = v2666
	v2668 = int32(1)
	if v304 == v2666 {
		v2714 = v2668
		goto L2
	} else {
		goto L430
	}
L413:
	;
	v2537 = v2532
	goto L416
L414:
	;
	v2609 = int32(0)
	goto L415
L415:
	;
	v2612 = v2531 + v2609<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v2612)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2612))) = v136
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+40)) = v2616 + int32(1)
	if v136 == int32(0) {
		goto L412
	} else {
		goto L422
	}
L416:
	;
	v2568 = v2531 + v2537<<(uint(int32(4))%32)
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(v2568)))
	if v136 == v2569 {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	v2609 = v2533
	goto L415
L418:
	;
	v2571 = *(*int64)(unsafe.Add(mBase, uint32(v2568)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2568)+8)) = v2571 + int64(1)
	goto L412
L419:
	;
	goto L420
L420:
	;
	v2576 = v2537 + int32(1)
	if v2576 != v2533 {
		v2537 = v2576
		goto L416
	} else {
		goto L421
	}
L421:
	;
	goto L417
L422:
	;
	v2623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+18)))
	if base.Ui32(v2623) <= base.Ui32(int32(15)) {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	goto L412
L424:
	;
	if v2623 != int32(15) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v83+v2623<<(uint(int32(2))%32))+292)) = v91
	goto L429
L428:
	;
	goto L429
L429:
	;
	v2633 = v2623 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+18)) = uint8(v2633)
	goto L426
L430:
	;
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2673 = m.G0
	v2675 = v2673 - int32(16)
	m.G0 = v2675
	v2677 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v2678 = m.ExcPending
	if v2678 != 0 {
		goto L23
	} else {
		goto L431
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2675)+8)) = v2672
	*(*int32)(unsafe.Add(mBase, uint32(v2675)+4)) = v2671
	*(*int32)(unsafe.Add(mBase, uint32(v2675))) = v2677
	*(*int32)(unsafe.Add(mBase, uint32(v2675)+12)) = int32(1)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2685 = m.ExcPending
	if v2685 != 0 {
		goto L23
	} else {
		goto L432
	}
L432:
	;
	F_XLogRegisterData(m, v2675+int32(12), int32(4))
	mBase = m.M
	v2690 = m.ExcPending
	if v2690 != 0 {
		goto L23
	} else {
		goto L433
	}
L433:
	;
	F_XLogRegisterData(m, v2675, int32(12))
	mBase = m.M
	v2693 = m.ExcPending
	if v2693 != 0 {
		goto L23
	} else {
		goto L434
	}
L434:
	;
	v2695 = int32(_a_F_LockAcquireExtended_27)
	v2697 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[26])))
	v2698 = v2697 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_LockAcquireExtended[26])) = uint8(v2698)
	goto L435
L435:
	;
	v2702 = F_XLogInsert(m, int32(8), int32(0))
	mBase = m.M
	v2703 = m.ExcPending
	if v2703 != 0 {
		goto L23
	} else {
		goto L436
	}
L436:
	;
	v2704 = int32(_a_F_LockAcquireExtended_28)
	v2706 = *(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[27]))
	*(*int32)(unsafe.Add(mBase, _c_F_LockAcquireExtended[27])) = v2706 | int32(2)
	m.G0 = v2675 + int32(16)
	v2714 = v2668
	goto L2
L437:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2753 = m.ExcPending
	if v2753 != 0 {
		goto L23
	} else {
		goto L438
	}
L438:
	;
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v2754+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v2758
	F_errmsg(m, int32(_a_F_LockAcquireExtended_29), v33+int32(96))
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		goto L23
	} else {
		goto L439
	}
L439:
	;
	F_errhint(m, int32(_a_F_LockAcquireExtended_30), int32(0))
	mBase = m.M
	v2768 = m.ExcPending
	if v2768 != 0 {
		goto L23
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(_a_F_LockAcquireExtended_2), int32(871), int32(_a_F_LockAcquireExtended_9))
	mBase = m.M
	v2773 = m.ExcPending
	if v2773 != 0 {
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
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int64
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v309 int64
	_ = v309
	var v312 int64
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int64
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int64
	_ = v459
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v608 int64
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v787 int32
	_ = v787
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v952 int32
	_ = v952
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1202 int32
	_ = v1202
	var v1216 int32
	_ = v1216
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1250 int32
	_ = v1250
	var v1284 int32
	_ = v1284
	var v1314 int32
	_ = v1314
	var v1320 int32
	_ = v1320
	var v1328 int32
	_ = v1328
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1351 int32
	_ = v1351
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1431 int32
	_ = v1431
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1535 int32
	_ = v1535
	var v1543 int32
	_ = v1543
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1589 int32
	_ = v1589
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1607 int32
	_ = v1607
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1643 int32
	_ = v1643
	var v1654 int32
	_ = v1654
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1725 int32
	_ = v1725
	var v1766 int32
	_ = v1766
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1780 int32
	_ = v1780
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1803 int32
	_ = v1803
	var v1808 int32
	_ = v1808
	var v1817 int32
	_ = v1817
	var v1826 int32
	_ = v1826
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1882 int32
	_ = v1882
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1897 int32
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int64
	_ = v1911
	var v1917 int32
	_ = v1917
	var v1922 int32
	_ = v1922
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1965 int32
	_ = v1965
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
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
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1959 != 0 {
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
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v277 == int32(0) {
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
	goto L28
L27:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v259 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v258 + v259
	v262 = int32(_a_F_lazy_vacuum_1)
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v264 - v259
	goto L23
L28:
	;
	goto L30
L30:
	;
	goto L31
L31:
	;
	v223 = int32(0)
	v226 = int32(0)
	goto L36
L36:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(40)+v226<<(uint(int32(2))%32))))
	v236 = int32(3)
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v34+int32(_a_F_lazy_vacuum_2)+v226<<(uint(v236)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v113+int32(232)+v235<<(uint(v236)%32)))) = v242
	v244 = int32(1)
	v247 = v223 + v244
	if v247 != int32(2) {
		v223 = v247
		v226 = v226 + v244
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
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v457 = v455 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v457
	v459 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[7]))) = v459
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[8]))) = v459
	*(*int64)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[9]))) = base.I64_extend_i32_s(v457)
	goto L59
L40:
	;
	v309 = int64(0)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v277)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v410)+16)) = base.F64_convert_i32_s(base.I32_trunc_sat_f32_s(v84))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v277)+16))
	v415 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v414)+24)) = uint8(v415)
	F_parallel_vacuum_process_all_indexes(m, v277, v409, v415)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L8
	} else {
		goto L55
	}
L43:
	;
	v312 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	v313 = base.B2i32(v312 <= v309)
	if v312 <= v309 {
		v429 = v313
		goto L39
	} else {
		goto L45
	}
L44:
	;
	v429 = v313
	goto L39
L45:
	;
	v316 = base.I32_wrap_i64(v309) << (uint(int32(2)) % 32)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v316+v317)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v320+v316)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v34)+64)) = base.F64_promote_f32(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = int32(13)
	v328 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+58)) = uint8(v328)
	v330 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+56)) = uint16(v330)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v324
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v322)+48))
	v338 = F_pstrdup(m, v335+int32(4))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v338
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v342 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v342)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(2)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v354 = F_vac_bulkdel_one_index(m, v34+int32(48), v319, v352, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v347
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v341)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v344
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	F_pfree(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	v362 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v362
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v364+v316))) = v354
	v369 = v309 + int64(1)
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	if v372 == v362 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v405 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L8
	} else {
		goto L53
	}
L50:
	;
	goto L49
L51:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[5])))
	if v376&int32(1) == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v381 = int32(_a_F_lazy_vacuum_1)
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	v384 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v383 + v384
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	*(*int32)(unsafe.Add(mBase, uint32(v372))) = v387 + v384
	*(*int64)(unsafe.Add(mBase, uint32(v372+int32(72))+232)) = v369
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	*(*int32)(unsafe.Add(mBase, uint32(v372))) = v395 + v384
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v401 - v384
	goto L50
L53:
	;
	if v405 == int32(0) {
		v309 = v369
		goto L43
	} else {
		goto L54
	}
L54:
	;
	goto L44
L55:
	;
	v420 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v429 = v420 ^ int32(1)
	goto L39
L57:
	;
	if v429 == int32(0) {
		goto L13
	} else {
		goto L74
	}
L58:
	;
	goto L57
L59:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	if v479 == int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[5])))
	if v483&int32(1) == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v488 = int32(_a_F_lazy_vacuum_1)
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	v491 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v490 + v491
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	*(*int32)(unsafe.Add(mBase, uint32(v479))) = v494 + v491
	goto L63
L62:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	v625 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v479))) = v624 + v625
	v628 = int32(_a_F_lazy_vacuum_1)
	v630 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v630 - v625
	goto L58
L63:
	;
	goto L65
L65:
	;
	goto L66
L66:
	;
	v589 = int32(0)
	v592 = int32(0)
	goto L71
L71:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v34+int32(24)+v592<<(uint(int32(2))%32))))
	v602 = int32(3)
	v608 = *(*int64)(unsafe.Add(mBase, uint32(v34+int32(_a_F_lazy_vacuum_3)+v592<<(uint(v602)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v479+int32(232)+v601<<(uint(v602)%32)))) = v608
	v610 = int32(1)
	v613 = v589 + v610
	if v613 != int32(3) {
		v589 = v613
		v592 = v592 + v610
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
	v645 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v645
	v651 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	if v651 == v645 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(3)
	v687 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v688 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v688)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v695 = F_palloc0(m, int32(16))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L8
	} else {
		goto L79
	}
L76:
	;
	goto L75
L77:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[5])))
	if v655&int32(1) == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v660 = int32(_a_F_lazy_vacuum_1)
	v662 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	v663 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v662 + v663
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v651)))
	*(*int32)(unsafe.Add(mBase, uint32(v651))) = v666 + v663
	*(*int64)(unsafe.Add(mBase, uint32(v651+int32(0))+232)) = int64(3)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v651)))
	*(*int32)(unsafe.Add(mBase, uint32(v651))) = v674 + v663
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v680 - v663
	goto L76
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v695))) = v693
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v693)+4))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v693)+8))
	if v699 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v695)+4)) = v746
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v758 = F_read_stream_begin_relation(m, int32(9), v753, v754, int32(0), int32(190), v695, int32(8))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L8
	} else {
		goto L87
	}
L81:
	;
	v701 = F_palloc0(m, int32(120))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L8
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v730 = F_palloc0(m, int32(88))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L8
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v701))) = v698
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v698)))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+24))
	v707 = F_dsa_get_address(m, v704, v706)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v709)))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)+48))
	v713 = base.I32_div_s(v711, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v701)+104)) = v713
	*(*int32)(unsafe.Add(mBase, uint32(v701)+100)) = v713
	v717 = v701 + int32(4)
	v718 = int32(12)
	v720 = v717 + v713*v718
	*(*int32)(unsafe.Add(mBase, uint32(v720)+4)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v720))) = v706
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v701)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v717+v723*v718)+8)) = int32(0)
	v746 = v701
	goto L80
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v698
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v698)))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v733)+24))
	v737 = base.I32_div_s(v735, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v730)+72)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v730)+68)) = v737
	v742 = v730 + v737<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v742)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v742)+4)) = v734
	v746 = v730
	goto L80
L87:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v765 = F_read_stream_next_buffer(m, v758, v34+int32(40))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	if v765 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v771 = v765
	v787 = v2
	goto L93
L91:
	;
	v1882 = v2
	goto L92
L92:
	;
	F_read_stream_end(m, v758)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L8
	} else {
		goto L267
	}
L93:
	;
	if v771 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v1882 = v1857
	goto L92
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v816
	v819 = v34 + int32(48)
	v820 = int32(0)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v822)+4))
	v824 = int32(*(*int8)(unsafe.Add(mBase, uint32(v823)+1)))
	if v824 != 0 {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[10]))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v801+(v771^int32(-1))<<(uint(int32(6))%32))+16))
	v816 = v807
	goto L95
L97:
	;
	goto L98
L98:
	;
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[11]))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v809+v771<<(uint(int32(6))%32)+int32(-64))+16))
	v816 = v815
	goto L95
L99:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_visibilitymap_pin(m, v987, v816, v34+int32(24))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L8
	} else {
		goto L121
	}
L100:
	;
	v837 = v824
	v838 = v820
	v841 = v820
	goto L106
L101:
	;
	if int32(0) < v824 {
		goto L100
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v828 = int32(0)
	v829 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823)+2)))
	if v829 == v828 {
		v986 = v828
		goto L99
	} else {
		goto L105
	}
L104:
	;
	v986 = int32(0)
	goto L99
L105:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v819))) = uint16(v829)
	v986 = int32(1)
	goto L99
L106:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v823+int32(4)+v841<<(uint(int32(2))%32))))
	if v870 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v986 = v922
	goto L99
L108:
	;
	v874 = v841 << (uint(int32(5)) % 32)
	v875 = v838
	v881 = v870
	goto L111
L109:
	;
	v921 = v837
	v922 = v838
	goto L110
L110:
	;
	v952 = v841 + int32(1)
	if v952 < base.I32_extend8_s(v921) {
		v837 = v921
		v838 = v922
		v841 = v952
		goto L106
	} else {
		goto L120
	}
L111:
	;
	if v881&int32(1) != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823)+1)))
	v921 = v919
	v922 = v914
	goto L110
L113:
	;
	if v875 < int32(2048) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v914 = v875
	goto L115
L115:
	;
	v915 = int32(1)
	v918 = int32(base.Ui32(v881) >> (uint(v915) % 32))
	if v918 != 0 {
		v874 = v874 + v915
		v875 = v914
		v881 = v918
		goto L111
	} else {
		goto L119
	}
L116:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v819+v875<<(uint(int32(1))%32)))) = uint16(v874)
	goto L118
L117:
	;
	goto L118
L118:
	;
	v914 = v875 + int32(1)
	goto L115
L119:
	;
	goto L112
L120:
	;
	goto L107
L121:
	;
	F_LockBuffer(m, v771, int32(2))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	v995 = int32(0)
	v996 = base.B2i32(v995 <= v771)
	if v996 == v995 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v1020 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[4]))
	if v1020 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[12]))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1000+(v771^int32(-1))<<(uint(int32(2))%32))))
	v1014 = v1006
	goto L123
L125:
	;
	goto L126
L126:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[13]))
	v1014 = v1008 + v771<<(uint(int32(13))%32) + int32(-8192)
	goto L123
L127:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(3)
	v1056 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v1057 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1057)
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v816
	v1062 = int32(_a_F_lazy_vacuum_1)
	v1064 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v1064 + int32(1)
	if v1057 < v986 {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	goto L127
L129:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_lazy_vacuum[5])))
	if v1024&int32(1) == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v1029 = int32(_a_F_lazy_vacuum_1)
	v1031 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	v1032 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v1031 + v1032
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	*(*int32)(unsafe.Add(mBase, uint32(v1020))) = v1035 + v1032
	*(*int64)(unsafe.Add(mBase, uint32(v1020+int32(24))+232)) = base.I64_extend_i32_u(v816)
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1020)))
	*(*int32)(unsafe.Add(mBase, uint32(v1020))) = v1043 + v1032
	v1049 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v1049 - v1032
	goto L128
L131:
	;
	v1071 = v986 & int32(3)
	v1073 = v1014 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v986) {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v1284 = v1057
	goto L133
L133:
	;
	v1314 = int32(0)
	v1320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1014)+12)))
	if base.Ui32(v1320) < base.Ui32(int32(25)) {
		goto L146
	} else {
		goto L147
	}
L134:
	;
	v1284 = v986
	goto L133
L135:
	;
	v1081 = v1057
	v1085 = int32(0)
	goto L138
L136:
	;
	v1171 = v1057
	goto L137
L137:
	;
	v1202 = v1171
	v1216 = int32(0)
	goto L142
L138:
	;
	v1112 = v1081 << (uint(int32(1)) % 32)
	v1114 = v34 + int32(48)
	v1116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1112+v1114))))
	v1117 = int32(2)
	v1120 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1073+v1116<<(uint(v1117)%32)))) = v1120
	v1123 = v34 + int32(_a_F_lazy_vacuum_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v1123+v1112))) = uint16(v1116)
	v1127 = v1112 | v1117
	v1129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1114+v1127))))
	*(*int32)(unsafe.Add(mBase, uint32(v1073+v1129<<(uint(v1117)%32)))) = v1120
	*(*uint16)(unsafe.Add(mBase, uint32(v1127+v1123))) = uint16(v1129)
	v1137 = int32(4)
	v1138 = v1112 | v1137
	v1140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1114+v1138))))
	*(*int32)(unsafe.Add(mBase, uint32(v1073+v1140<<(uint(v1117)%32)))) = v1120
	*(*uint16)(unsafe.Add(mBase, uint32(v1123+v1138))) = uint16(v1140)
	v1151 = v1112 | int32(6)
	v1153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1114+v1151))))
	*(*int32)(unsafe.Add(mBase, uint32(v1073+v1153<<(uint(v1117)%32)))) = v1120
	*(*uint16)(unsafe.Add(mBase, uint32(v1123+v1151))) = uint16(v1153)
	v1164 = v1081 + v1137
	v1166 = v1085 + v1137
	if v1166 != v986&int32(2147483644) {
		v1081 = v1164
		v1085 = v1166
		goto L138
	} else {
		goto L140
	}
L139:
	;
	if v1071 == int32(0) {
		goto L134
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v1171 = v1164
	goto L137
L142:
	;
	v1232 = int32(1)
	v1233 = v1202 << (uint(v1232) % 32)
	v1237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1233+(v34+int32(48))))))
	*(*int32)(unsafe.Add(mBase, uint32(v1073+v1237<<(uint(int32(2))%32)))) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v34+int32(_a_F_lazy_vacuum_3)+v1233))) = uint16(v1237)
	v1250 = v1216 + v1232
	if v1250 != v1071 {
		v1202 = v1202 + v1232
		v1216 = v1250
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
	F_MarkBufferDirty(m, v771)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L8
	} else {
		goto L162
	}
L146:
	;
	v1385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1014)+10)))
	v1387 = v1385 & int32(_a_F_lazy_vacuum_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1014)+10)) = uint16(v1387)
	goto L145
L147:
	;
	v1328 = int32(base.Ui32(v1320+int32(_a_F_lazy_vacuum_5))>>(uint(int32(2))%32)) & int32(_a_F_lazy_vacuum_6)
	if v1328 == int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v1334 = v1328
	v1336 = v1314
	v1339 = v1314
	goto L150
L149:
	;
	if int32(0) < v1363 {
		goto L158
	} else {
		goto L159
	}
L150:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1014+int32(20)+v1334<<(uint(int32(2))%32))))
	v1345 = v1343 & int32(_a_F_lazy_vacuum_7)
	if base.B2i32(v1334 == int32(1))|v1339 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v1363 = v1357
	v1365 = int32(0)
	goto L149
L152:
	;
	v1360 = v1334 - int32(1)
	if v1360 != 0 {
		v1334 = v1360
		v1336 = v1357
		v1339 = v1358
		goto L150
	} else {
		goto L157
	}
L153:
	;
	v1351 = int32(0)
	v1357 = v1336 + base.B2i32(v1345 == v1351)
	v1358 = base.B2i32(v1345 != v1351)
	goto L152
L154:
	;
	goto L155
L155:
	;
	if v1345 != 0 {
		v1357 = v1336
		v1358 = v1339
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v1363 = v1336
	v1365 = int32(1)
	goto L149
L157:
	;
	goto L151
L158:
	;
	v1370 = v1320 - v1363<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1014)+12)) = uint16(v1370)
	goto L160
L159:
	;
	goto L160
L160:
	;
	if v1365 == int32(0) {
		goto L146
	} else {
		goto L161
	}
L161:
	;
	v1374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1014)+10)))
	v1376 = v1374 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1014)+10)) = uint16(v1376)
	goto L145
L162:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+48))
	v1400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1399)+118)))
	if v1400 != int32(112) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v1422 = int32(_a_F_lazy_vacuum_1)
	v1424 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[6])) = v1424 - int32(1)
	if v996 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L164:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[14]))
	if v1404 <= int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+32))
	if v1407 != 0 {
		goto L163
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v1409 = int32(0)
	F_log_heap_prune_and_freeze(m, v1398, v771, v1409, v1409, int32(2), v1409, v1409, v1409, v1409, v1409, v1409, v34+int32(_a_F_lazy_vacuum_3), v1284)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L8
	} else {
		goto L170
	}
L168:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1398)+40))
	if v1408 != 0 {
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
	if v771 < int32(0) {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[12]))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1431+(v771^int32(-1))<<(uint(int32(2))%32))))
	v1445 = v1437
	goto L171
L173:
	;
	goto L174
L174:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[13]))
	v1445 = v1439 + v771<<(uint(int32(13))%32) + int32(-8192)
	goto L171
L175:
	;
	v1465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1445)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1465) {
		goto L182
	} else {
		goto L183
	}
L176:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[10]))
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1449+(v771^int32(-1))<<(uint(int32(6))%32))+16))
	v1464 = v1455
	goto L175
L177:
	;
	goto L178
L178:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[11]))
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1457+v771<<(uint(int32(6))%32)+int32(-64))+16))
	v1464 = v1463
	goto L175
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v1053
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1056)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1060
	if v996 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L180:
	;
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v1725 + int32(1)
	goto L179
L181:
	;
	v1492 = int32(base.Ui32(v1464) >> (uint(int32(16)) % 32))
	v1495 = int32(1)
	v1499 = v1495
	v1503 = int32(0)
	v1506 = v1495
	goto L187
L182:
	;
	v1473 = int32(base.Ui32(v1465+int32(_a_F_lazy_vacuum_5))>>(uint(int32(2))%32)) & int32(_a_F_lazy_vacuum_6)
	if v1473 != 0 {
		goto L181
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v1475 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1475)
	v1477 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1014)+10)))
	v1479 = v1477 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1014)+10)) = uint16(v1479)
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1485 = F_visibilitymap_set(m, v1481, v816, v771, int64(0), v1015, v1475, int32(3))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L8
	} else {
		goto L186
	}
L185:
	;
	goto L184
L186:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1487 + int32(1)
	goto L180
L187:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1499)
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1445+int32(20)+v1499&int32(_a_F_lazy_vacuum_6)<<(uint(int32(2))%32))))
	switch int32(base.Ui32(v1535)>>(uint(int32(15))%32)) & int32(3) {
	case 0, 2:
		v1664 = v1503
		v1665 = v1506
		goto L189
	default:
		goto L190
	}
L188:
	;
	v1673 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1673)
	v1675 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1014)+10)))
	v1677 = v1675 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1014)+10)) = uint16(v1677)
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1682 = int32(1)
	v1684 = v1665 & v1682
	if v1684 != 0 {
		goto L234
	} else {
		goto L235
	}
L189:
	;
	v1669 = v1499 + int32(1)
	if base.Ui32(v1669&int32(_a_F_lazy_vacuum_6)) <= base.Ui32(v1473) {
		v1499 = v1669
		v1503 = v1664
		v1506 = v1665
		goto L187
	} else {
		goto L233
	}
L190:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[3]))) = uint16(v1499)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[15]))) = uint16(v1464)
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[16]))) = uint16(v1492)
	v1543 = int32(_a_F_lazy_vacuum_7)
	if v1535&v1543 == v1543 {
		goto L179
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[2]))) = int32(base.Ui32(v1535) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[17]))) = v1445 + v1535&int32(_a_F_lazy_vacuum_8)
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(v1554)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[18]))) = v1555
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1560 = F_HeapTupleSatisfiesVacuum(m, v34+int32(_a_F_lazy_vacuum_2), v1559, v771)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L8
	} else {
		goto L192
	}
L192:
	;
	if v1560 != int32(1) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	if base.B2i32(v1560 != int32(1))&base.B2i32(base.Ui32(v1560) <= base.Ui32(int32(4))) != 0 {
		goto L179
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[17])))
	v1583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1582)+20)))
	if v1583&int32(256) == int32(0) {
		goto L179
	} else {
		goto L200
	}
L196:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L8
	} else {
		goto L197
	}
L197:
	;
	F_errmsg_internal(m, int32(_a_F_lazy_vacuum_9), int32(0))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L8
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_lazy_vacuum_10), int32(3708), int32(_a_F_lazy_vacuum_11))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
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
	v1589 = int32(768)
	if v1583&v1589 != v1589 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1582)))
	v1594 = v1593
	goto L203
L202:
	;
	v1594 = int32(2)
	goto L203
L203:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1595))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1594)) == int32(0) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	if v1607 == int32(0) {
		goto L179
	} else {
		goto L208
	}
L205:
	;
	v1607 = base.B2i32(base.Ui32(v1594) < base.Ui32(v1595))
	goto L204
L206:
	;
	goto L207
L207:
	;
	v1607 = int32(base.Ui32(v1594-v1595) >> (uint(int32(31)) % 32))
	goto L204
L208:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1503))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1594)) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	if v1621 != 0 {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v1621 = base.B2i32(base.Ui32(v1503) < base.Ui32(v1594))
	goto L209
L211:
	;
	goto L212
L212:
	;
	v1621 = base.B2i32(int32(0) < v1594-v1503)
	goto L209
L213:
	;
	v1622 = v1594
	goto L215
L214:
	;
	v1622 = v1503
	goto L215
L215:
	;
	if base.Ui32(int32(2)) < base.Ui32(v1594) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1625 = v1622
	goto L218
L217:
	;
	v1625 = v1503
	goto L218
L218:
	;
	v1628 = int32(0)
	if v1506&int32(1) == v1628 {
		v1664 = v1625
		v1665 = v1628
		goto L189
	} else {
		goto L219
	}
L219:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v34)+uint32(_c_F_lazy_vacuum[17])))
	v1634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1631)+20)))
	v1635 = int32(768)
	if v1634&v1635 == v1635 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v1664 = v1625
	v1665 = v1660 ^ int32(1)
	goto L189
L221:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+4))
	if v1634&int32(_a_F_lazy_vacuum_12) != 0 {
		goto L225
	} else {
		goto L226
	}
L222:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1631)))
	if base.Ui32(v1639) <= base.Ui32(int32(2)) {
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1660 = int32(1)
	goto L220
L224:
	;
	if base.Ui32(v1634) < base.Ui32(int32(_a_F_lazy_vacuum_13)) {
		goto L230
	} else {
		goto L231
	}
L225:
	;
	if v1643 == int32(0) {
		goto L224
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	if base.Ui32(v1643) <= base.Ui32(int32(2)) {
		goto L224
	} else {
		goto L229
	}
L228:
	;
	v1660 = int32(1)
	goto L220
L229:
	;
	v1660 = int32(1)
	goto L220
L230:
	;
	v1660 = int32(0)
	goto L220
L231:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+8))
	if base.Ui32(v1654) <= base.Ui32(int32(2)) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v1660 = int32(1)
	goto L220
L233:
	;
	goto L188
L234:
	;
	v1685 = int32(3)
	goto L236
L235:
	;
	v1685 = v1682
	goto L236
L236:
	;
	v1686 = F_visibilitymap_set(m, v1679, v816, v771, int64(0), v1015, v1664, v1685)
	mBase = m.M
	v1687 = m.ExcPending
	if v1687 != 0 {
		goto L8
	} else {
		goto L237
	}
L237:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1688 + int32(1)
	if v1684 == int32(0) {
		goto L179
	} else {
		goto L238
	}
L238:
	;
	goto L180
L239:
	;
	v1784 = int32(4)
	v1785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1780)+14)))
	v1786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1780)+12)))
	v1787 = v1785 - v1786
	if v1787 <= v1784 {
		goto L244
	} else {
		goto L245
	}
L240:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[12]))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1766+(v771^int32(-1))<<(uint(int32(2))%32))))
	v1780 = v1772
	goto L239
L241:
	;
	goto L242
L242:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, _c_F_lazy_vacuum[13]))
	v1780 = v1774 + v771<<(uint(int32(13))%32) + int32(-8192)
	goto L239
L243:
	;
	F_UnlockReleaseBuffer(m, v771)
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L8
	} else {
		goto L262
	}
L244:
	;
	v1790 = v1784
	goto L246
L245:
	;
	v1790 = v1787
	goto L246
L246:
	;
	v1792 = v1790 - int32(4)
	if v1792 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1847 = int32(0)
	goto L243
L248:
	;
	goto L249
L249:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1786) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	v1847 = v1792
	goto L243
L251:
	;
	v1803 = int32(base.Ui32(v1786+int32(_a_F_lazy_vacuum_5)) >> (uint(int32(2)) % 32))
	goto L253
L252:
	;
	v1803 = int32(0)
	goto L253
L253:
	;
	if base.Ui32(v1803&int32(_a_F_lazy_vacuum_6)) < base.Ui32(int32(291)) {
		goto L250
	} else {
		goto L254
	}
L254:
	;
	v1808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1780)+10)))
	if v1808&int32(1) == int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1847 = int32(0)
	goto L243
L256:
	;
	goto L257
L257:
	;
	v1817 = int32(1)
	goto L258
L258:
	;
	v1826 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1780+int32(20)+v1817&int32(_a_F_lazy_vacuum_6)<<(uint(int32(2))%32))+1)))
	if v1826&int32(384) == int32(0) {
		goto L250
	} else {
		goto L260
	}
L259:
	;
	v1847 = int32(0)
	goto L243
L260:
	;
	v1832 = v1817 + int32(1)
	v1833 = int32(_a_F_lazy_vacuum_6)
	if base.Ui32(v1832&v1833) <= base.Ui32(v1803&v1833) {
		v1817 = v1832
		goto L258
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_RecordPageWithFreeSpace(m, v1850, v816, v1847)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L8
	} else {
		goto L263
	}
L263:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L8
	} else {
		goto L264
	}
L264:
	;
	v1857 = v787 + int32(1)
	v1860 = F_read_stream_next_buffer(m, v758, v34+int32(40))
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L8
	} else {
		goto L265
	}
L265:
	;
	if v1860 != 0 {
		v771 = v1860
		v787 = v1857
		goto L93
	} else {
		goto L266
	}
L266:
	;
	goto L94
L267:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	F_pfree(m, v1895)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L8
	} else {
		goto L268
	}
L268:
	;
	F_pfree(m, v695)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L8
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v1902 != 0 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	F_ReleaseBuffer(m, v1902)
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L8
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1907 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L8
	} else {
		goto L274
	}
L273:
	;
	goto L272
L274:
	;
	if v1907 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1911 = *(*int64)(unsafe.Add(mBase, uint32(v1910)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v1882
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v1911
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v1909
	F_errmsg(m, int32(_a_F_lazy_vacuum_14), v34)
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L8
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v684
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v687)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v690
	goto L13
L278:
	;
	F_errfinish(m, int32(_a_F_lazy_vacuum_10), int32(2823), int32(_a_F_lazy_vacuum_15))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L8
	} else {
		goto L279
	}
L279:
	;
	goto L277
L280:
	;
	F_parallel_vacuum_reset_dead_items(m, v1959)
	mBase = m.M
	v1961 = m.ExcPending
	if v1961 != 0 {
		goto L8
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_TidStoreDestroy(m, v1971)
	mBase = m.M
	v1973 = m.ExcPending
	if v1973 != 0 {
		goto L8
	} else {
		goto L285
	}
L283:
	;
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1962)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(104)))) = v1965 + int32(56)
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1962)+24))
	goto L284
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1969
	goto L1
L285:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v1974)))
	v1976 = F_TidStoreCreateLocal(m, v1975)
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L8
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1976
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v1979)+8)) = int64(0)
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
