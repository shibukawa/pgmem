package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LCS_asString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v5 = l0 - int32(1)
	if base.Ui32(v5) <= base.Ui32(int32(3)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5<<(uint(int32(2))%32))+uint32(_consts[392])))
		v13 = v12
	} else {
		v13 = int32(373969)
	}
	return v13
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
						F_sequence_close(m, v18, int32(1))
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v187 int32
	_ = v187
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int64
	_ = v365
	var v367 int32
	_ = v367
	var v378 int32
	_ = v378
	var v394 int64
	_ = v394
	var v398 int32
	_ = v398
	var v400 int64
	_ = v400
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int64
	_ = v414
	var v416 int32
	_ = v416
	var v418 int64
	_ = v418
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int64
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int64
	_ = v449
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v476 int64
	_ = v476
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int64
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int64
	_ = v545
	var v550 int32
	_ = v550
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v604 int32
	_ = v604
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v745 int32
	_ = v745
	var v752 int32
	_ = v752
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int64
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v817 int64
	_ = v817
	var v824 int32
	_ = v824
	var v826 int64
	_ = v826
	var v835 int64
	_ = v835
	var v840 int32
	_ = v840
	var v842 int64
	_ = v842
	var v851 int64
	_ = v851
	var v854 int64
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int64
	_ = v860
	var v861 int64
	_ = v861
	var v867 int64
	_ = v867
	var v870 int64
	_ = v870
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int64
	_ = v906
	var v911 int32
	_ = v911
	var v913 int64
	_ = v913
	var v917 int64
	_ = v917
	var v921 int64
	_ = v921
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int64
	_ = v957
	var v962 int32
	_ = v962
	var v964 int64
	_ = v964
	var v968 int64
	_ = v968
	var v972 int64
	_ = v972
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int64
	_ = v1008
	var v1017 int32
	_ = v1017
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1085 int32
	_ = v1085
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1139 int32
	_ = v1139
	var v1140 int64
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int64
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1278 int32
	_ = v1278
	var v1285 int32
	_ = v1285
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1325 int32
	_ = v1325
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int64
	_ = v1397
	var v1399 int64
	_ = v1399
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1405 int64
	_ = v1405
	var v1407 int64
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1448 int32
	_ = v1448
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1471 int32
	_ = v1471
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1509 int32
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1524 int32
	_ = v1524
	var v1534 int32
	_ = v1534
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1647 int64
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1703 int32
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1747 int32
	_ = v1747
	var v1754 int32
	_ = v1754
	var v1759 int32
	_ = v1759
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1806 int32
	_ = v1806
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1832 int32
	_ = v1832
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
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1876 int32
	_ = v1876
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1947 int32
	_ = v1947
	var v1951 int32
	_ = v1951
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2020 int32
	_ = v2020
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2050 int32
	_ = v2050
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2100 int32
	_ = v2100
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
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2175 int32
	_ = v2175
	var v2212 int32
	_ = v2212
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2231 int32
	_ = v2231
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2270 int32
	_ = v2270
	var v2275 int32
	_ = v2275
	var v2277 int32
	_ = v2277
	var v2338 int32
	_ = v2338
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2351 int32
	_ = v2351
	var v2383 int32
	_ = v2383
	var v2388 int32
	_ = v2388
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int64
	_ = v2394
	var v2402 int32
	_ = v2402
	var v2405 int32
	_ = v2405
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2421 int32
	_ = v2421
	var v2425 int32
	_ = v2425
	var v2430 int32
	_ = v2430
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2443 int32
	_ = v2443
	var v2448 int32
	_ = v2448
	var v2454 int32
	_ = v2454
	var v2459 int32
	_ = v2459
	var v2463 int32
	_ = v2463
	var v2467 int32
	_ = v2467
	var v2472 int32
	_ = v2472
	var v2503 int32
	_ = v2503
	var v2533 int64
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2543 int32
	_ = v2543
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2576 int64
	_ = v2576
	var v2581 int32
	_ = v2581
	var v2613 int32
	_ = v2613
	var v2616 int32
	_ = v2616
	var v2620 int32
	_ = v2620
	var v2627 int32
	_ = v2627
	var v2637 int32
	_ = v2637
	var v2669 int32
	_ = v2669
	var v2671 int32
	_ = v2671
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2681 int32
	_ = v2681
	var v2688 int32
	_ = v2688
	var v2693 int32
	_ = v2693
	var v2696 int32
	_ = v2696
	var v2698 int32
	_ = v2698
	var v2700 int32
	_ = v2700
	var v2701 int32
	_ = v2701
	var v2705 int64
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2723 int32
	_ = v2723
	var v2752 int32
	_ = v2752
	var v2755 int32
	_ = v2755
	var v2756 int32
	_ = v2756
	var v2760 int32
	_ = v2760
	var v2766 int32
	_ = v2766
	var v2770 int32
	_ = v2770
	var v2775 int32
	_ = v2775
	v7 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(192)
	m.G0 = v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v34-int32(3))&int32(255)) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2752 = m.ExcPending
	if v2752 != 0 {
		goto L24
	} else {
		goto L443
	}
L2:
	;
	m.G0 = v32 + int32(192)
	return v2723
L3:
	;
	v2533 = *(*int64)(unsafe.Add(mBase, uint32(v89)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+32)) = v2533 + int64(1)
	v2537 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v2538 = int32(0)
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	if v2538 < v2539 {
		goto L419
	} else {
		goto L420
	}
L4:
	;
	F_LWLockRelease(m, v1092)
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L24
	} else {
		goto L417
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
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L24
	} else {
		goto L414
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L24
	} else {
		goto L411
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(2))%32))+uint32(_consts[1183])))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v48 < l1 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
	if v52 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+176)) = v75
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+168)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v32)+184)) = l1
	v81 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v83 = *(*int32)(unsafe.Add(mBase, _consts[1184]))
	v89 = F_hash_search(m, v83, v32+int32(168), int32(1), v32+int32(167))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	if v62 == int32(0) {
		goto L11
	} else {
		goto L16
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+316))
	v60 = base.B2i32(v58 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v60)
	v62 = v60
	goto L15
L14:
	;
	v62 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[158])))
	if v66 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v67 != int32(8) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if base.Ui32(l1) < base.Ui32(int32(4)) {
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
		goto L23
	}
L21:
	;
	if v67 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	goto L1
L23:
	;
	goto L11
L24:
	;
	return int32(0)
L25:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+167)))
	if v93 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	if l4 != 0 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v89)+24)) = int64(0)
	v99 = *(*int32)(unsafe.Add(mBase, _consts[1114]))
	v102 = F_get_hash_value(m, v99, v32+int32(168))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L24
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v89)+44))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	if v120 < v119 {
		goto L26
	} else {
		goto L32
	}
L30:
	;
	v104 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v89)+52)) = uint16(v104)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v89)+48)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v89)+40)) = int64(34359738368)
	v114 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v116 = F_MemoryContextAlloc(m, v114, int32(128))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+48)) = v116
	goto L26
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v125 = F_repalloc(m, v122, v119<<(uint(int32(5))%32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+44)) = v119 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+48)) = v125
	goto L26
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v89
	goto L36
L35:
	;
	goto L36
L36:
	;
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v136 = int32(0)
	goto L39
L38:
	;
	v136 = v81
	goto L39
L39:
	;
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v89)+32))
	if int64(0) < v137 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v89)+32)) = v137 + int64(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v144 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	if v144 < v145 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(l1) < base.Ui32(int32(8)) {
		v301 = v7
		goto L64
	} else {
		goto L65
	}
L43:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+53)))
	if v276 != 0 {
		goto L61
	} else {
		goto L62
	}
L44:
	;
	v148 = v144
	goto L47
L45:
	;
	v219 = int32(0)
	goto L46
L46:
	;
	v222 = v143 + v219<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v222)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v136
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+40)) = v226 + int32(1)
	if v136 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L47:
	;
	v179 = v143 + v148<<(uint(int32(4))%32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if v136 == v180 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v219 = v145
	goto L46
L49:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v179)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v179)+8)) = v182 + int64(1)
	goto L43
L50:
	;
	goto L51
L51:
	;
	v187 = v148 + int32(1)
	if v187 != v145 {
		v148 = v187
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)))
	if base.Ui32(v233) <= base.Ui32(int32(15)) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L43
L55:
	;
	if v233 != int32(15) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136+v233<<(uint(int32(2))%32))+292)) = v89
	goto L60
L59:
	;
	goto L60
L60:
	;
	v243 = v233 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)) = uint8(v243)
	goto L57
L61:
	;
	v277 = int32(3)
	goto L63
L62:
	;
	v277 = int32(2)
	goto L63
L63:
	;
	v2723 = v277
	goto L2
L64:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v302 != int32(1) {
		goto L77
	} else {
		goto L78
	}
L65:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v280 != 0 {
		v301 = v7
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
	if v283 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v293 != 0 {
		v301 = v7
		goto L64
	} else {
		goto L71
	}
L68:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+316))
	v291 = base.B2i32(v289 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v291)
	v293 = v291
	goto L70
L69:
	;
	v293 = int32(0)
	goto L70
L70:
	;
	goto L67
L71:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v295 <= int32(0) {
		v301 = v7
		goto L64
	} else {
		goto L72
	}
L72:
	;
	v298 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L24
	} else {
		goto L73
	}
L73:
	;
	v301 = int32(1)
	goto L64
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v1098
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1098)))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v1213
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+20))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1216+l1<<(uint(int32(2))%32))))
	if v1215&v1220 != 0 {
		goto L222
	} else {
		goto L223
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L24
	} else {
		goto L217
	}
L76:
	;
	F_LWLockRelease(m, v724)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L24
	} else {
		goto L202
	}
L77:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v1092 = v1085 + v133&int32(15)<<(uint(int32(7))%32) + int32(23296)
	v1094 = F_LWLockAcquire(m, v1092, int32(0))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L24
	} else {
		goto L184
	}
L78:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v305 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	if v671 != int32(1) {
		goto L77
	} else {
		goto L130
	}
L80:
	;
	if base.Ui32(int32(3)) < base.Ui32(l1) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v309 != v310 {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	if v309 == int32(0) {
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _consts[329]))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32((v315-int32(1))&(v318*int32(49157))<<(uint(int32(2))%32))+uint32(_consts[1185])))
	if int32(15) < v326 {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v334 = F_LWLockAcquire(m, v330+int32(584), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L24
	} else {
		goto L85
	}
L85:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v337+v133&int32(1023)<<(uint(int32(2))%32))+4))
	if v343 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _consts[329]))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v354 = (v348 - int32(1)) & (v351 * int32(49157))
	v355 = int32(4)
	v356 = v354 << (uint(v355) % 32)
	v358 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+600))
	v364 = v359 + v354&int32(268435455)<<(uint(int32(3))%32)
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v364)))
	v367 = v348 << (uint(v355) % 32)
	v378 = v367
	v394 = int64(0)
	goto L91
L87:
	;
	goto L88
L88:
	;
	v637 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	F_LWLockRelease(m, v637+int32(584))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L24
	} else {
		goto L129
	}
L89:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	F_LWLockRelease(m, v493+int32(584))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L24
	} else {
		goto L107
	}
L90:
	;
	v480 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v364))) = v365 | int64(1)<<(uint(base.I64_extend_i32_u(l1+base.I32_wrap_i64(v476)-v480))%64)
	v491 = v480
	goto L89
L91:
	;
	v398 = v356 + base.I32_wrap_i64(v394)
	v400 = v394 * int64(3)
	if int64(base.Ui64(v365)>>(uint(v400)%64))&int64(7) == int64(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	if base.Ui32(v430) < base.Ui32(v367) {
		goto L104
	} else {
		goto L105
	}
L93:
	;
	v414 = v394 | int64(1)
	v416 = v356 + base.I32_wrap_i64(v414)
	v418 = v414 * int64(3)
	if int64(base.Ui64(v365)>>(uint(v418)%64))&int64(7) == int64(0) {
		goto L99
	} else {
		goto L100
	}
L94:
	;
	v412 = v398
	goto L93
L95:
	;
	goto L96
L96:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v358)+604))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v406+v398<<(uint(int32(2))%32))))
	if v410 == v351 {
		v476 = v400
		goto L90
	} else {
		goto L97
	}
L97:
	;
	v412 = v378
	goto L93
L98:
	;
	v432 = v394 + int64(2)
	if v432 != int64(16) {
		v378 = v430
		v394 = v432
		goto L91
	} else {
		goto L103
	}
L99:
	;
	v430 = v416
	goto L98
L100:
	;
	goto L101
L101:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v358)+604))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v424+v416<<(uint(int32(2))%32))))
	if v428 == v351 {
		v476 = v418
		goto L90
	} else {
		goto L102
	}
L102:
	;
	v430 = v412
	goto L98
L103:
	;
	goto L92
L104:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v358)+604))
	v437 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v436+v430<<(uint(v437)%32)))) = v351
	v442 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+600))
	v444 = int32(1)
	v448 = v443 + int32(base.Ui32(v430)>>(uint(v444)%32))&int32(2147483640)
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v448)))
	*(*int64)(unsafe.Add(mBase, uint32(v448))) = v449 | int64(1)<<(uint(base.I64_extend_i32_u(l1+v430&int32(15)*int32(3)-v444))%64)
	v463 = v354 << (uint(v437) % 32)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v463)+uint32(_consts[1185])))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+uint32(_consts[1185]))) = v466 + v444
	v473 = v444
	goto L106
L105:
	;
	v473 = int32(0)
	goto L106
L106:
	;
	v491 = v473
	goto L89
L107:
	;
	if v491 == int32(0) {
		goto L79
	} else {
		goto L108
	}
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v89)+24)) = int64(0)
	v502 = int32(0)
	v503 = *(*int64)(unsafe.Add(mBase, uint32(v89)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+32)) = v503 + int64(1)
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	if v502 < v508 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v2723 = int32(1)
	goto L2
L110:
	;
	v511 = v502
	goto L113
L111:
	;
	v582 = int32(0)
	goto L112
L112:
	;
	v585 = v507 + v582<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v585)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v585))) = v136
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+40)) = v589 + int32(1)
	if v136 != 0 {
		goto L119
	} else {
		goto L120
	}
L113:
	;
	v542 = v507 + v511<<(uint(int32(4))%32)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	if v136 == v543 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v582 = v508
	goto L112
L115:
	;
	v545 = *(*int64)(unsafe.Add(mBase, uint32(v542)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v542)+8)) = v545 + int64(1)
	goto L109
L116:
	;
	goto L117
L117:
	;
	v550 = v511 + int32(1)
	if v550 != v508 {
		v511 = v550
		goto L113
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)))
	if base.Ui32(v594) <= base.Ui32(int32(15)) {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	goto L121
L121:
	;
	goto L109
L122:
	;
	goto L121
L123:
	;
	if v594 != int32(15) {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	goto L122
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136+v594<<(uint(int32(2))%32))+292)) = v89
	goto L128
L127:
	;
	goto L128
L128:
	;
	v604 = v594 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)) = uint8(v604)
	goto L125
L129:
	;
	goto L79
L130:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	if v674 != 0 {
		goto L77
	} else {
		goto L131
	}
L131:
	;
	if base.Ui32(l1) < base.Ui32(int32(5)) {
		goto L77
	} else {
		goto L132
	}
L132:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v677 == int32(0) {
		goto L77
	} else {
		goto L133
	}
L133:
	;
	v681 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	*(*int32)(unsafe.Add(mBase, uint32(v681))) = int32(1)
	if v682 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v688 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	F_s_lock(m, v688, int32(497060), int32(1837), int32(364365))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L24
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v695 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	v700 = v695 + v133&int32(1023)<<(uint(int32(2))%32) + int32(4)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)))
	v702 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v700))) = v701 + v702
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+52)) = uint8(v702)
	*(*int32)(unsafe.Add(mBase, _consts[1186])) = v89
	v709 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v695))) = v709
	v712 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+16))
	if v713 == v709 {
		goto L77
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v717 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v724 = v717 + v133&int32(15)<<(uint(int32(7))%32) + int32(23296)
	v726 = *(*int32)(unsafe.Add(mBase, _consts[329]))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v732 = (v726 - int32(1)) & (v729 * int32(49157))
	v745 = v712
	v752 = v7
	goto L139
L139:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	v771 = v768 + v752*int32(640)
	v773 = v771 + int32(584)
	v775 = F_LWLockAcquire(m, v773, int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L24
	} else {
		goto L141
	}
L140:
	;
	goto L77
L141:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v771)+60))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v777 != v778 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	F_LWLockRelease(m, v773)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L24
	} else {
		goto L182
	}
L143:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v771)+600))
	v782 = *(*int64)(unsafe.Add(mBase, uint32(v780+v732<<(uint(int32(3))%32))))
	if v782 == int64(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v786 = v732 & int32(268435455) << (uint(int32(3)) % 32)
	v787 = v780 + v786
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v771)+604))
	v789 = v788 + v732<<(uint(int32(6))%32)
	v817 = int64(0)
	goto L145
L145:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v789+base.I32_wrap_i64(v817)<<(uint(int32(2))%32))))
	if v729 != v824 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v856 = F_LWLockAcquire(m, v724, int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L24
	} else {
		goto L156
	}
L147:
	;
	goto L146
L148:
	;
	v835 = v817 | int64(1)
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v789+base.I32_wrap_i64(v835)<<(uint(int32(2))%32))))
	if v840 == v729 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v826 = *(*int64)(unsafe.Add(mBase, uint32(v787)))
	if int64(base.Ui64(v826)>>(uint(v817*int64(3))%64))&int64(7) == int64(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v854 = v817
	goto L147
L151:
	;
	v842 = *(*int64)(unsafe.Add(mBase, uint32(v787)))
	if int64(base.Ui64(v842)>>(uint(v835*int64(3))%64))&int64(7) != int64(0) {
		v854 = v835
		goto L147
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v851 = v817 + int64(2)
	if v851 != int64(16) {
		v817 = v851
		goto L145
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	goto L142
L156:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v771)+600))
	v860 = *(*int64)(unsafe.Add(mBase, uint32(v858+v786)))
	v861 = int64(1)
	v867 = (v854*int64(12884901888) - int64(4294967296)) >> (uint(int64(32)) % 64)
	v870 = v861 << (uint(v867+v861) % 64)
	if v860&v870 != int64(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v875 = F_SetupLockInTable(m, v47, v771, l0, v133, int32(1))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L24
	} else {
		goto L160
	}
L158:
	;
	v917 = v860
	goto L159
L159:
	;
	v921 = int64(1) << (uint(v867+int64(2)) % 64)
	if v917&v921 != int64(0) {
		goto L165
	} else {
		goto L166
	}
L160:
	;
	if v875 == int32(0) {
		goto L76
	} else {
		goto L161
	}
L161:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v879)+128))
	v881 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v879)+128)) = v880 + v881
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v879)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v879)+16)) = v884 | int32(2)
	v889 = v879 + int32(92)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)))
	v892 = v890 + v881
	*(*int32)(unsafe.Add(mBase, uint32(v889))) = v892
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v879)+48))
	if v894 == v892 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v879)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v879)+20)) = v896 & int32(-3)
	goto L164
L163:
	;
	goto L164
L164:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v875)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v875)+12)) = v900 | int32(2)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v771)+600))
	v905 = v904 + v786
	v906 = *(*int64)(unsafe.Add(mBase, uint32(v905)))
	*(*int64)(unsafe.Add(mBase, uint32(v905))) = v906 & (v870 ^ int64(-1))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v771)+600))
	v913 = *(*int64)(unsafe.Add(mBase, uint32(v911+v786)))
	v917 = v913
	goto L159
L165:
	;
	v926 = F_SetupLockInTable(m, v47, v771, l0, v133, int32(2))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L24
	} else {
		goto L168
	}
L166:
	;
	v968 = v917
	goto L167
L167:
	;
	v972 = int64(1) << (uint(v867+int64(3)) % 64)
	if v968&v972 != int64(0) {
		goto L173
	} else {
		goto L174
	}
L168:
	;
	if v926 == int32(0) {
		goto L76
	} else {
		goto L169
	}
L169:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v926)))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v930)+128))
	v932 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v930)+128)) = v931 + v932
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v930)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v930)+16)) = v935 | int32(4)
	v940 = v930 + int32(96)
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v940)))
	v943 = v941 + v932
	*(*int32)(unsafe.Add(mBase, uint32(v940))) = v943
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v930)+52))
	if v945 == v943 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v930)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v930)+20)) = v947 & int32(-5)
	goto L172
L171:
	;
	goto L172
L172:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v926)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v926)+12)) = v951 | int32(4)
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v771)+600))
	v956 = v955 + v786
	v957 = *(*int64)(unsafe.Add(mBase, uint32(v956)))
	*(*int64)(unsafe.Add(mBase, uint32(v956))) = v957 & (v921 ^ int64(-1))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v771)+600))
	v964 = *(*int64)(unsafe.Add(mBase, uint32(v962+v786)))
	v968 = v964
	goto L167
L173:
	;
	v977 = F_SetupLockInTable(m, v47, v771, l0, v133, int32(3))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L24
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	F_LWLockRelease(m, v724)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L24
	} else {
		goto L181
	}
L176:
	;
	if v977 == int32(0) {
		goto L76
	} else {
		goto L177
	}
L177:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v977)))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v981)+128))
	v983 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v981)+128)) = v982 + v983
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v981)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v981)+16)) = v986 | int32(8)
	v991 = v981 + int32(100)
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v991)))
	v994 = v992 + v983
	*(*int32)(unsafe.Add(mBase, uint32(v991))) = v994
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v981)+56))
	if v996 == v994 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v981)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v981)+20)) = v998 & int32(-9)
	goto L180
L179:
	;
	goto L180
L180:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v977)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v977)+12)) = v1002 | int32(8)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v771)+600))
	v1007 = v1006 + v786
	v1008 = *(*int64)(unsafe.Add(mBase, uint32(v1007)))
	*(*int64)(unsafe.Add(mBase, uint32(v1007))) = v1008 & (v972 ^ int64(-1))
	goto L175
L181:
	;
	goto L142
L182:
	;
	v1050 = v752 + int32(1)
	v1052 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+16))
	if base.Ui32(v1050) < base.Ui32(v1053) {
		v745 = v1052
		v752 = v1050
		goto L139
	} else {
		goto L183
	}
L183:
	;
	goto L140
L184:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v1098 = F_SetupLockInTable(m, v47, v1097, l0, v133, l1)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L24
	} else {
		goto L185
	}
L185:
	;
	if v1098 != 0 {
		goto L74
	} else {
		goto L186
	}
L186:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, _consts[1186]))
	if v1101 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1103)))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1103))) = int32(1)
	if v1104 != 0 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	F_LWLockRelease(m, v1092)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L24
	} else {
		goto L194
	}
L190:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	F_s_lock(m, v1111, int32(497060), int32(1869), int32(364342))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L24
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	v1123 = v1118 + v1105&int32(1023)<<(uint(int32(2))%32) + int32(4)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1123)))
	*(*int32)(unsafe.Add(mBase, uint32(v1123))) = v1124 - int32(1)
	v1128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1101)+52)) = uint8(v1128)
	*(*int32)(unsafe.Add(mBase, _consts[1186])) = v1128
	*(*int32)(unsafe.Add(mBase, uint32(v1118))) = v1128
	goto L189
L193:
	;
	goto L192
L194:
	;
	v1140 = *(*int64)(unsafe.Add(mBase, uint32(v89)+32))
	if v1140 == int64(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	F_RemoveLocalLock(m, v89)
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L24
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	if l4 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L197
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L201
L200:
	;
	goto L201
L201:
	;
	goto L75
L202:
	;
	F_LWLockRelease(m, v773)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L24
	} else {
		goto L203
	}
L203:
	;
	F_AbortStrongLockAcquire(m)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L24
	} else {
		goto L204
	}
L204:
	;
	v1159 = *(*int64)(unsafe.Add(mBase, uint32(v89)+32))
	if v1159 == int64(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	F_RemoveLocalLock(m, v89)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L24
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	if l4 != 0 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	goto L207
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L211
L210:
	;
	goto L211
L211:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L24
	} else {
		goto L212
	}
L212:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L24
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(14090), int32(0))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L24
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = int32(255642)
	F_errhint(m, int32(650943), v32+int32(80))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L24
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(497060), int32(1042), int32(460889))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L24
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L24
	} else {
		goto L218
	}
L218:
	;
	F_errmsg(m, int32(14090), int32(0))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L24
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(255642)
	F_errhint(m, int32(650943), v32+int32(32))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L24
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(497060), int32(1080), int32(460889))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L24
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	v1253 = int32(0)
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1257 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+104)) = v1259
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+616))
	if v1261 == v1253 {
		v1325 = v1259
		goto L234
	} else {
		goto L235
	}
L223:
	;
	v1222 = F_LockCheckConflicts(m, v47, l1, v1213, v1098)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L24
	} else {
		goto L224
	}
L224:
	;
	if v1222 != 0 {
		goto L222
	} else {
		goto L225
	}
L225:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+128))
	v1225 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1213)+128)) = v1224 + v1225
	v1230 = v1213 + l1<<(uint(int32(2))%32)
	v1232 = v1230 + int32(88)
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)))
	*(*int32)(unsafe.Add(mBase, uint32(v1232))) = v1233 + v1225
	v1238 = v1225 << (uint(l1) % 32)
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1213)+16)) = v1238 | v1239
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1232)))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+44))
	if v1242 == v1243 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1213)+20)) = v1245 & (v1238 ^ int32(-1))
	goto L228
L227:
	;
	goto L228
L228:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+12)) = v1250 | v1238
	goto L4
L229:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L24
	} else {
		goto L408
	}
L230:
	;
	v1892 = m.G0
	v1894 = v1892 - int32(128)
	m.G0 = v1894
	F_initStringInfo(m, v1894+int32(112))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L24
	} else {
		goto L340
	}
L231:
	;
	F_LWLockRelease(m, v1092)
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L24
	} else {
		goto L337
	}
L232:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, _consts[1186]))
	if v1568 != 0 {
		goto L280
	} else {
		goto L281
	}
L233:
	;
	switch v1564 - int32(1) {
	case 0:
		goto L231
	case 1:
		goto L232
	default:
		goto L4
	}
L234:
	;
	v1340 = v1255 + int32(32)
	if v1325 == int32(0) {
		v1471 = v1253
		goto L244
	} else {
		goto L245
	}
L235:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+28))
	if v1264 == int32(0) {
		v1325 = v1259
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1268 = v1255 + int32(24)
	if v1264 == v1268 {
		v1325 = v1259
		goto L234
	} else {
		goto L237
	}
L237:
	;
	v1278 = v1264
	v1285 = v1259
	goto L238
L238:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1278-int32(12))))
	if v1261 == v1301 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v1325 = v1307
	goto L234
L240:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1278-int32(8))))
	v1307 = v1305 | v1285
	goto L242
L241:
	;
	v1307 = v1285
	goto L242
L242:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+4))
	if v1308 != v1268 {
		v1278 = v1308
		v1285 = v1307
		goto L238
	} else {
		goto L243
	}
L243:
	;
	goto L239
L244:
	;
	if l3 != 0 {
		goto L270
	} else {
		goto L271
	}
L245:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+40))
	if v1343 == int32(0) {
		v1471 = v1253
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+36))
	if v1346 == int32(0) {
		v1471 = v1253
		goto L244
	} else {
		goto L247
	}
L247:
	;
	if v1346 == v1340 {
		v1471 = v1253
		goto L244
	} else {
		goto L248
	}
L248:
	;
	v1359 = int32(0)
	v1361 = v1346
	goto L249
L249:
	;
	if v1261 != 0 {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	v1471 = int32(0)
	goto L244
L251:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+4))
	if v1458 != v1340 {
		v1359 = v1455
		v1361 = v1458
		goto L249
	} else {
		goto L269
	}
L252:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+616))
	if v1261 == v1380 {
		v1455 = v1359
		goto L251
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+100))
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1382+v1383<<(uint(int32(2))%32))))
	if v1387&v1325 != 0 {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	goto L254
L256:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1382+v1254<<(uint(int32(2))%32))))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+104))
	if v1392&v1393 != 0 {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L258
L258:
	;
	v1455 = int32(1)<<(uint(v1383)%32) | v1359
	goto L251
L259:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	v1397 = *(*int64)(unsafe.Add(mBase, uint32(v1255)))
	*(*int64)(unsafe.Add(mBase, uint32(v1396))) = v1397
	v1399 = *(*int64)(unsafe.Add(mBase, uint32(v1255)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1396)+8)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v1396)+16)) = v1254
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1396)+20)) = v1402
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+92))
	v1405 = *(*int64)(unsafe.Add(mBase, uint32(v1404)))
	*(*int64)(unsafe.Add(mBase, uint32(v1396)+24)) = v1405
	v1407 = *(*int64)(unsafe.Add(mBase, uint32(v1404)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1396)+32)) = v1407
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v1396)+40)) = v1409
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1396)+44)) = v1411
	v1414 = int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[1175])) = v1414
	v1564 = v1414
	goto L233
L260:
	;
	goto L261
L261:
	;
	if v1359&v1392 != 0 {
		v1471 = v1361
		goto L244
	} else {
		goto L262
	}
L262:
	;
	v1418 = F_LockCheckConflicts(m, v47, v1254, v1255, v1258)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L24
	} else {
		goto L263
	}
L263:
	;
	if v1418 != 0 {
		v1471 = v1361
		goto L244
	} else {
		goto L264
	}
L264:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+128))
	v1423 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+128)) = v1422 + v1423
	v1428 = v1255 + v1254<<(uint(int32(2))%32)
	v1430 = v1428 + int32(88)
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v1430)))
	*(*int32)(unsafe.Add(mBase, uint32(v1430))) = v1431 + v1423
	v1436 = v1423 << (uint(v1254) % 32)
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+16)) = v1436 | v1437
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1430)))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+44))
	if v1440 == v1441 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v1564 = int32(0)
	goto L233
L266:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+20)) = v1443 & (v1436 ^ int32(-1))
	goto L268
L267:
	;
	goto L268
L268:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+12)) = v1448 | v1436
	goto L265
L269:
	;
	goto L250
L270:
	;
	v1534 = int32(2)
	goto L272
L271:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	if v1471 != 0 {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	v1564 = v1534
	goto L233
L273:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+40))
	v1515 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+40)) = v1514 + v1515
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+20)) = v1518 | v1515<<(uint(v1254)%32)
	v1524 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v1524)+100)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v1524)+96)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v1524)+92)) = v1255
	*(*int32)(unsafe.Add(mBase, uint32(v1524)+104)) = v1259
	*(*int32)(unsafe.Add(mBase, uint32(v1524)+16)) = v1515
	v1534 = v1515
	goto L272
L274:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1471)))
	*(*int32)(unsafe.Add(mBase, uint32(v1492)+4)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v1492))) = v1493
	*(*int32)(unsafe.Add(mBase, uint32(v1471))) = v1492
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1492)))
	*(*int32)(unsafe.Add(mBase, uint32(v1497)+4)) = v1492
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+36))
	if v1499 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+40)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+36)) = v1340
	*(*int32)(unsafe.Add(mBase, uint32(v1255)+32)) = v1255 + int32(32)
	goto L279
L278:
	;
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1492)+4)) = v1340
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1340)))
	*(*int32)(unsafe.Add(mBase, uint32(v1492))) = v1509
	*(*int32)(unsafe.Add(mBase, uint32(v1509)+4)) = v1492
	*(*int32)(unsafe.Add(mBase, uint32(v1340))) = v1492
	goto L273
L280:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1570)))
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1570))) = int32(1)
	if v1571 != 0 {
		goto L283
	} else {
		goto L284
	}
L281:
	;
	goto L282
L282:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+12))
	if v1605 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L283:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	F_s_lock(m, v1578, int32(497060), int32(1869), int32(364342))
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L24
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, _consts[1116]))
	v1590 = v1585 + v1572&int32(1023)<<(uint(int32(2))%32) + int32(4)
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1590)))
	*(*int32)(unsafe.Add(mBase, uint32(v1590))) = v1591 - int32(1)
	v1595 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1568)+52)) = uint8(v1595)
	*(*int32)(unsafe.Add(mBase, _consts[1186])) = v1595
	*(*int32)(unsafe.Add(mBase, uint32(v1585))) = v1595
	goto L282
L286:
	;
	goto L285
L287:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+4))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+20))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1609)+4)) = v1610
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1610))) = v1612
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+28))
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1614)+4)) = v1615
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1615))) = v1617
	v1620 = *(*int32)(unsafe.Add(mBase, _consts[1115]))
	v1626 = F_hash_search_with_hash_value(m, v1620, v1098, v1608<<(uint(int32(4))%32)^v133, int32(2), int32(0))
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L24
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+84))
	v1633 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1213)+84)) = v1632 - v1633
	v1640 = v1213 + l1<<(uint(int32(2))%32) + int32(44)
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1640)))
	*(*int32)(unsafe.Add(mBase, uint32(v1640))) = v1641 - v1633
	F_LWLockRelease(m, v1092)
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L24
	} else {
		goto L292
	}
L290:
	;
	if v1626 == int32(0) {
		goto L229
	} else {
		goto L291
	}
L291:
	;
	goto L289
L292:
	;
	v1647 = *(*int64)(unsafe.Add(mBase, uint32(v89)+32))
	if v1647 == int64(0) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	F_RemoveLocalLock(m, v89)
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L24
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	if l3 == int32(0) {
		goto L230
	} else {
		goto L297
	}
L296:
	;
	goto L295
L297:
	;
	if l5 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = int32(0)
	F_initStringInfo(m, v32+int32(148))
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L24
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1876 = int32(0)
	if l4 == v1876 {
		v2723 = v1876
		goto L2
	} else {
		goto L336
	}
L301:
	;
	F_initStringInfo(m, v32+int32(132))
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L24
	} else {
		goto L302
	}
L302:
	;
	F_initStringInfo(m, v32+int32(116))
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L24
	} else {
		goto L303
	}
L303:
	;
	F_DescribeLockTag(m, v32+int32(148), v89)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L24
	} else {
		goto L304
	}
L304:
	;
	v1672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+15)))
	v1673 = int32(2)
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1672<<(uint(v1673)%32))+uint32(_consts[1183])))
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1677)+8))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1678+l1<<(uint(v1673)%32))))
	v1684 = F_LWLockAcquire(m, v1092, int32(1))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L24
	} else {
		goto L305
	}
L305:
	;
	v1689 = v32 + int32(132)
	v1690 = m.G0
	v1692 = v1690 - int32(48)
	m.G0 = v1692
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1696 = v32 + int32(112)
	v1697 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1696))) = v1697
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+28))
	if v1699 == v1697 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	m.G0 = v1692 + int32(48)
	F_LWLockRelease(m, v1092)
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L24
	} else {
		goto L325
	}
L307:
	;
	v1703 = v1694 + int32(24)
	if v1699 == v1703 {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1705 = int32(1)
	v1709 = v1705
	v1714 = v1699
	v1715 = v1705
	goto L309
L309:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1714-int32(16))))
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1738)+44))
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1738)+96))
	if v1740 == v1714-int32(20) {
		goto L312
	} else {
		goto L313
	}
L310:
	;
	goto L306
L311:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1714)+4))
	if v1771 != v1703 {
		v1709 = v1769
		v1714 = v1771
		v1715 = v1770
		goto L309
	} else {
		goto L324
	}
L312:
	;
	if v1709 != 0 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	goto L314
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+32)) = v1739
	if v1715 != 0 {
		goto L320
	} else {
		goto L321
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1692))) = v1739
	F_appendStringInfo(m, v1689, int32(487662), v1692)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L24
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+16)) = v1739
	F_appendStringInfo(m, v1689, int32(487648), v1692+int32(16))
	mBase = m.M
	v1754 = m.ExcPending
	if v1754 != 0 {
		goto L24
	} else {
		goto L319
	}
L318:
	;
	v1769 = int32(0)
	v1770 = v1715
	goto L311
L319:
	;
	v1769 = int32(0)
	v1770 = v1715
	goto L311
L320:
	;
	v1759 = int32(487662)
	goto L322
L321:
	;
	v1759 = int32(487648)
	goto L322
L322:
	;
	F_appendStringInfo(m, v32+int32(116), v1759, v1692+int32(32))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L24
	} else {
		goto L323
	}
L323:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1696)))
	*(*int32)(unsafe.Add(mBase, uint32(v1696))) = v1764 + int32(1)
	v1769 = v1709
	v1770 = int32(0)
	goto L311
L324:
	;
	goto L310
L325:
	;
	v1809 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L24
	} else {
		goto L326
	}
L326:
	;
	if v1809 != 0 {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v1682
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v1812
	v1815 = *(*int32)(unsafe.Add(mBase, _consts[712]))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v1815
	F_errmsg(m, int32(184229), v32-int32(-64))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L24
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	F_pfree(m, v1838)
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L24
	} else {
		goto L333
	}
L330:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v32)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v1822
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v32)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v1824
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
	F_errdetail_log_plural(m, int32(590651), int32(590697), v1828, v32+int32(48))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L24
	} else {
		goto L331
	}
L331:
	;
	F_errfinish(m, int32(497060), int32(1192), int32(460889))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L24
	} else {
		goto L332
	}
L332:
	;
	goto L329
L333:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v32)+116))
	F_pfree(m, v1841)
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L24
	} else {
		goto L334
	}
L334:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v32)+132))
	F_pfree(m, v1844)
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L24
	} else {
		goto L335
	}
L335:
	;
	goto L300
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v2723 = v1876
	goto L2
L337:
	;
	v1883 = F_WaitOnLock(m, v89, v136)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L24
	} else {
		goto L338
	}
L338:
	;
	if v1883 != int32(2) {
		goto L3
	} else {
		goto L339
	}
L339:
	;
	goto L230
L340:
	;
	F_initStringInfo(m, v1894+int32(96))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L24
	} else {
		goto L341
	}
L341:
	;
	F_initStringInfo(m, v1894+int32(80))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L24
	} else {
		goto L342
	}
L342:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, _consts[1175]))
	if v1909 <= int32(0) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+112))
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+116))
	F_appendBinaryStringInfo(m, v1894+int32(96), v2087, v2088)
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L24
	} else {
		goto L364
	}
L344:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	if v1909 == int32(1) {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1920 = v1913 + int32(20)
	goto L347
L346:
	;
	v1920 = v1913 + int32(44)
	goto L347
L347:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1920)))
	v1923 = v1894 + int32(80)
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1923)))
	v1925 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1924))) = uint8(v1925)
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+12)) = v1925
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+4)) = v1925
	goto L348
L348:
	;
	F_DescribeLockTag(m, v1894+int32(80), v1913)
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L24
	} else {
		goto L349
	}
L349:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+20))
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1913)+15)))
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+16))
	v1938 = int32(2)
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1936<<(uint(v1938)%32))+uint32(_consts[1183])))
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1942)+8))
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1943+v1937<<(uint(v1938)%32))))
	goto L350
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+64)) = v1935
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+68)) = v1947
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+76)) = v1921
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+72)) = v1951
	F_appendStringInfo(m, v1894+int32(112), int32(637384), v1894-int32(-64))
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L24
	} else {
		goto L351
	}
L351:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, _consts[1175]))
	if v1961 < int32(2) {
		goto L343
	} else {
		goto L352
	}
L352:
	;
	v1966 = int32(1)
	v1968 = v1961
	goto L353
L353:
	;
	v1995 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	v1998 = v1995 + v1966*int32(24)
	if v1966 < v1968-int32(1) {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	goto L343
L355:
	;
	v2006 = v1998 + int32(44)
	goto L357
L356:
	;
	v2006 = v1995 + int32(20)
	goto L357
L357:
	;
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v2006)))
	v2009 = v1894 + int32(80)
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v2009)))
	v2011 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2010))) = uint8(v2011)
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+12)) = v2011
	*(*int32)(unsafe.Add(mBase, uint32(v2009)+4)) = v2011
	goto L358
L358:
	;
	F_DescribeLockTag(m, v1894+int32(80), v1998)
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L24
	} else {
		goto L359
	}
L359:
	;
	F_appendStringInfoChar(m, v1894+int32(112), int32(10))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L24
	} else {
		goto L360
	}
L360:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v1998)+20))
	v2027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1998)+15)))
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v1998)+16))
	v2029 = int32(2)
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v2027<<(uint(v2029)%32))+uint32(_consts[1183])))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v2033)+8))
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v2034+v2028<<(uint(v2029)%32))))
	goto L361
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+48)) = v2026
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+52)) = v2038
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+60)) = v2007
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+56)) = v2042
	F_appendStringInfo(m, v1894+int32(112), int32(637384), v1894+int32(48))
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L24
	} else {
		goto L362
	}
L362:
	;
	v2052 = v1966 + int32(1)
	v2054 = *(*int32)(unsafe.Add(mBase, _consts[1175]))
	if v2052 < v2054 {
		v1966 = v2052
		v1968 = v2054
		goto L353
	} else {
		goto L363
	}
L363:
	;
	goto L354
L364:
	;
	v2092 = *(*int32)(unsafe.Add(mBase, _consts[1175]))
	if int32(0) < v2092 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v2100 = int32(0)
	goto L368
L366:
	;
	goto L367
L367:
	;
	v2383 = int32(*(*uint8)(unsafe.Add(mBase, _consts[384])))
	if v2383 == int32(1) {
		goto L397
	} else {
		goto L398
	}
L368:
	;
	v2125 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	F_appendStringInfoChar(m, v1894+int32(96), int32(10))
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L24
	} else {
		goto L370
	}
L369:
	;
	goto L367
L370:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2125+v2100*int32(24))+20))
	v2136 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	if int32(0) < v2136 {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+36)) = v2338
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+32)) = v2134
	F_appendStringInfo(m, v1894+int32(96), int32(204015), v1894+int32(32))
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L24
	} else {
		goto L395
	}
L372:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	v2143 = v2140
	v2146 = int32(1)
	goto L375
L373:
	;
	goto L374
L374:
	;
	v2338 = int32(545444)
	goto L371
L375:
	;
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+4))
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2143)))
	v2175 = int32(0)
	if base.B2i32(v2172&int32(1) == v2175)&base.B2i32(v2172 == v2172) == v2175 {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	goto L374
L377:
	;
	goto L380
L378:
	;
	v2231 = v2171
	goto L379
L379:
	;
	if v2134 == v2231 {
		goto L388
	} else {
		goto L389
	}
L380:
	;
	v2212 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v2212 != 0 {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	v2231 = v2215
	goto L379
L382:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L24
	} else {
		goto L385
	}
L383:
	;
	goto L384
L384:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+4))
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v2143)))
	if v2216&int32(1) != 0 {
		goto L380
	} else {
		goto L386
	}
L385:
	;
	goto L384
L386:
	;
	if v2216 != v2216 {
		goto L380
	} else {
		goto L387
	}
L387:
	;
	goto L381
L388:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(v2143)+216))
	v2253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2252))))
	if v2253 == int32(0) {
		v2338 = int32(545515)
		goto L371
	} else {
		goto L391
	}
L389:
	;
	goto L390
L390:
	;
	v2275 = v2146 + int32(1)
	v2277 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	if v2275 <= v2277 {
		v2143 = v2143 + int32(408)
		v2146 = v2275
		goto L375
	} else {
		goto L394
	}
L391:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, _consts[782]))
	v2260 = F_pnstrdup(m, v2252, v2257-int32(1))
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L24
	} else {
		goto L392
	}
L392:
	;
	v2262 = F_strlen(m, v2260)
	mBase = m.M
	v2264 = *(*int32)(unsafe.Add(mBase, _consts[782]))
	v2267 = F_pg_mbcliplen(m, v2260, v2262, v2264-int32(1))
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L24
	} else {
		goto L393
	}
L393:
	;
	v2270 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2267+v2260))) = uint8(v2270)
	v2338 = v2260
	goto L371
L394:
	;
	goto L376
L395:
	;
	v2349 = v2100 + int32(1)
	v2351 = *(*int32)(unsafe.Add(mBase, _consts[1175]))
	if v2349 < v2351 {
		v2100 = v2349
		goto L368
	} else {
		goto L396
	}
L396:
	;
	goto L369
L397:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v2391 = F_pgstat_prep_pending_entry(m, int32(1), v2388, int64(0), int32(0))
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L24
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L24
	} else {
		goto L401
	}
L400:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v2391)+12))
	v2394 = *(*int64)(unsafe.Add(mBase, uint32(v2393)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v2393)+144)) = v2394 + int64(1)
	goto L399
L401:
	;
	F_errcode(m, int32(16908292))
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L24
	} else {
		goto L402
	}
L402:
	;
	F_errmsg(m, int32(446352), int32(0))
	mBase = m.M
	v2409 = m.ExcPending
	if v2409 != 0 {
		goto L24
	} else {
		goto L403
	}
L403:
	;
	v2410 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+16)) = v2410
	F_errdetail_internal(m, int32(205923), v1894+int32(16))
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L24
	} else {
		goto L404
	}
L404:
	;
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v1894))) = v2417
	F_errdetail_log(m, int32(205923), v1894)
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		goto L24
	} else {
		goto L405
	}
L405:
	;
	F_errhint(m, int32(578649), int32(0))
	mBase = m.M
	v2425 = m.ExcPending
	if v2425 != 0 {
		goto L24
	} else {
		goto L406
	}
L406:
	;
	F_errfinish(m, int32(497047), int32(1138), int32(80807))
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L24
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
	F_errmsg_internal(m, int32(444360), int32(0))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L24
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(497060), int32(1140), int32(460889))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L24
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
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = l1
	F_errmsg_internal(m, int32(486071), v32+int32(16))
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L24
	} else {
		goto L412
	}
L412:
	;
	F_errfinish(m, int32(497060), int32(861), int32(460889))
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L24
	} else {
		goto L413
	}
L413:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
	F_errmsg_internal(m, int32(486525), v32)
	mBase = m.M
	v2467 = m.ExcPending
	if v2467 != 0 {
		goto L24
	} else {
		goto L415
	}
L415:
	;
	F_errfinish(m, int32(497060), int32(858), int32(460889))
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L24
	} else {
		goto L416
	}
L416:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L417:
	;
	goto L3
L418:
	;
	v2669 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1186])) = v2669
	v2671 = int32(1)
	if v301 == v2669 {
		v2723 = v2671
		goto L2
	} else {
		goto L436
	}
L419:
	;
	v2543 = v2538
	goto L422
L420:
	;
	v2613 = int32(0)
	goto L421
L421:
	;
	v2616 = v2537 + v2613<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v2616)+8)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2616))) = v136
	v2620 = *(*int32)(unsafe.Add(mBase, uint32(v89)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+40)) = v2620 + int32(1)
	if v136 == int32(0) {
		goto L418
	} else {
		goto L428
	}
L422:
	;
	v2573 = v2537 + v2543<<(uint(int32(4))%32)
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(v2573)))
	if v136 == v2574 {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	v2613 = v2539
	goto L421
L424:
	;
	v2576 = *(*int64)(unsafe.Add(mBase, uint32(v2573)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2573)+8)) = v2576 + int64(1)
	goto L418
L425:
	;
	goto L426
L426:
	;
	v2581 = v2543 + int32(1)
	if v2581 != v2539 {
		v2543 = v2581
		goto L422
	} else {
		goto L427
	}
L427:
	;
	goto L423
L428:
	;
	v2627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)))
	if base.Ui32(v2627) <= base.Ui32(int32(15)) {
		goto L430
	} else {
		goto L431
	}
L429:
	;
	goto L418
L430:
	;
	if v2627 != int32(15) {
		goto L433
	} else {
		goto L434
	}
L431:
	;
	goto L432
L432:
	;
	goto L429
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136+v2627<<(uint(int32(2))%32))+292)) = v89
	goto L435
L434:
	;
	goto L435
L435:
	;
	v2637 = v2627 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+18)) = uint8(v2637)
	goto L432
L436:
	;
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2676 = m.G0
	v2678 = v2676 - int32(16)
	m.G0 = v2678
	v2680 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v2681 = m.ExcPending
	if v2681 != 0 {
		goto L24
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2678)+8)) = v2675
	*(*int32)(unsafe.Add(mBase, uint32(v2678)+4)) = v2674
	*(*int32)(unsafe.Add(mBase, uint32(v2678))) = v2680
	*(*int32)(unsafe.Add(mBase, uint32(v2678)+12)) = int32(1)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2688 = m.ExcPending
	if v2688 != 0 {
		goto L24
	} else {
		goto L438
	}
L438:
	;
	F_XLogRegisterData(m, v2678+int32(12), int32(4))
	mBase = m.M
	v2693 = m.ExcPending
	if v2693 != 0 {
		goto L24
	} else {
		goto L439
	}
L439:
	;
	F_XLogRegisterData(m, v2678, int32(12))
	mBase = m.M
	v2696 = m.ExcPending
	if v2696 != 0 {
		goto L24
	} else {
		goto L440
	}
L440:
	;
	v2698 = int32(4385284)
	v2700 = int32(*(*uint8)(unsafe.Add(mBase, _consts[94])))
	v2701 = v2700 | int32(2)
	*(*uint8)(unsafe.Add(mBase, _consts[94])) = uint8(v2701)
	goto L441
L441:
	;
	v2705 = F_XLogInsert(m, int32(8), int32(0))
	mBase = m.M
	v2706 = m.ExcPending
	if v2706 != 0 {
		goto L24
	} else {
		goto L442
	}
L442:
	;
	v2707 = int32(4384564)
	v2709 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	*(*int32)(unsafe.Add(mBase, _consts[173])) = v2709 | int32(2)
	m.G0 = v2678 + int32(16)
	v2723 = v2671
	goto L2
L443:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2755 = m.ExcPending
	if v2755 != 0 {
		goto L24
	} else {
		goto L444
	}
L444:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v2756+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v2760
	F_errmsg(m, int32(127820), v32+int32(96))
	mBase = m.M
	v2766 = m.ExcPending
	if v2766 != 0 {
		goto L24
	} else {
		goto L445
	}
L445:
	;
	F_errhint(m, int32(556924), int32(0))
	mBase = m.M
	v2770 = m.ExcPending
	if v2770 != 0 {
		goto L24
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(497060), int32(871), int32(460889))
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L24
	} else {
		goto L447
	}
L447:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__ltq_extract_regex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v22 = F_array_iterator(m, v12, int32(5617), v18, v9+int32(12))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v26 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v30 != v18 {
								F_pfree(m, v18)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v34 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
									v56 = int32(0)
									m.G0 = v9 + int32(16)
									return v56
								}
							} else {
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v56 = int32(0)
								m.G0 = v9 + int32(16)
								return v56
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v30 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
								v56 = int32(0)
								m.G0 = v9 + int32(16)
								return v56
							}
						} else {
							v34 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
							v56 = int32(0)
							m.G0 = v9 + int32(16)
							return v56
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v40 = F_palloc0(m, int32(base.Ui32(v37)>>(uint(int32(2))%32)))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
						v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
						if v45 != 0 {
							v46 = F__emscripten_memcpy_bulkmem(m, v40, v42, v45)
							mBase = m.M
						} else {
						}
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v48 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v18 == v52 {
									v56 = v40
									m.G0 = v9 + int32(16)
									return v56
								} else {
									F_pfree(m, v18)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = v40
										m.G0 = v9 + int32(16)
										return v56
									}
								}
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v18 == v52 {
								v56 = v40
								m.G0 = v9 + int32(16)
								return v56
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v40
									m.G0 = v9 + int32(16)
									return v56
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_lappend_int(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
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
	if l0 == int32(0) {
		v9 = F_palloc(m, int32(32))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(4294967767)
			v18 = v9 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v18
			v71 = v9
			v72 = v18
			v73 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
			return v71
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v22 <= v21 {
			v24 = int32(1)
			v26 = int32(16)
			v28 = v21 + v24
			if v28 <= v26 {
				v31 = v26
			} else {
				v31 = v28
			}
			if v31&(v31-int32(1)) != 0 {
				v38 = v24 << (uint(int32(32)-base.I32_clz(v31)) % 32)
			} else {
				v38 = v31
			}
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v41 = l0 + int32(16)
			if v39 == v41 {
				v43 = F_GetMemoryChunkContext(m, l0)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v47 = F_MemoryContextAlloc(m, v43, v38<<(uint(int32(2))%32))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v52 = v50 << (uint(int32(2)) % 32)
						if v52 != 0 {
							v53 = F__emscripten_memcpy_bulkmem(m, v47, v41, v52)
							mBase = m.M
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v66 = v62
						v68 = v66 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v71 = l0
						v72 = v70
						v73 = v68
						*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
						return v71
					}
				}
			} else {
				v57 = F_repalloc(m, v39, v38<<(uint(int32(2))%32))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v57
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v38
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v66 = v62
					v68 = v66 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v71 = l0
					v72 = v70
					v73 = v68
					*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
					return v71
				}
			}
		} else {
			v66 = v21
			v68 = v66 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v68
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v71 = l0
			v72 = v70
			v73 = v68
			*(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4)))) = l1
			return v71
		}
	}
}
func F_lazy_vacuum(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 float64
	_ = v69
	var v75 int32
	_ = v75
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
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 float32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int64
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v249 int64
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v284 int32
	_ = v284
	var v315 int64
	_ = v315
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
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
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int64
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int64
	_ = v467
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v616 int64
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
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
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v791 int32
	_ = v791
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v956 int32
	_ = v956
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1199 int32
	_ = v1199
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1346 int32
	_ = v1346
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1451 int32
	_ = v1451
	var v1459 int32
	_ = v1459
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1523 int32
	_ = v1523
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1568 int32
	_ = v1568
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1637 int32
	_ = v1637
	var v1677 int32
	_ = v1677
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1714 int32
	_ = v1714
	var v1719 int32
	_ = v1719
	var v1728 int32
	_ = v1728
	var v1739 int32
	_ = v1739
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1793 int32
	_ = v1793
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int64
	_ = v1823
	var v1829 int32
	_ = v1829
	var v1834 int32
	_ = v1834
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	v2 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(4768)
	m.G0 = v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+23)))
	if v35 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v33 + int32(4768)
	return
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v38 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
	if v61 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	F_parallel_vacuum_reset_dead_items(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_TidStoreDestroy(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L11
	}
L8:
	;
	return
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(104)))) = v44 + int32(56)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v48
	goto L1
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v55 = F_TidStoreCreateLocal(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+8)) = int64(0)
	goto L1
L13:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1870 != 0 {
		goto L289
	} else {
		goto L290
	}
L14:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	v91 = *(*float32)(unsafe.Add(mBase, uint32(v90)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+40)) = int64(34359738368)
	v95 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v95
	v98 = *(*int64)(unsafe.Add(mBase, _consts[102]))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v98
	v100 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L24
	}
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v64 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v69 = base.F64_mul(base.F64_convert_i32_u(v64), float64(0.02))
	if base.F64_lt(v69, float64(4.294967296e+09))&base.F64_ge(v69, float64(0)) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if base.Ui32(v77) <= base.Ui32(v78) {
		goto L14
	} else {
		goto L21
	}
L18:
	;
	v75 = base.I32_trunc_f64_u(v69)
	v77 = v75
	goto L17
L19:
	;
	goto L20
L20:
	;
	v77 = int32(0)
	goto L17
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v81 = F_TidStoreMemoryUsage(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32(int32(33554431)) < base.Ui32(v81) {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v85 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+23)) = uint8(v85)
	goto L13
L24:
	;
	if v100 != 0 {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[103]))) = int64(2)
	v104 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[104]))) = v104
	v111 = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v118 == v111 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v284 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L27:
	;
	goto L26
L28:
	;
	goto L29
L29:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v124&int32(1) == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v129 = int32(4484100)
	v131 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v132 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v131 + v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v135 + v132
	goto L32
L31:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v266 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v265 + v266
	v269 = int32(4484100)
	v271 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v271 - v266
	goto L27
L32:
	;
	goto L34
L34:
	;
	goto L35
L35:
	;
	goto L39
L39:
	;
	v230 = int32(0)
	v233 = v111
	goto L40
L40:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(40)+v233<<(uint(int32(2))%32))))
	v243 = int32(3)
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v33+int32(4736)+v233<<(uint(v243)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v118+int32(232)+v242<<(uint(v243)%32)))) = v249
	v251 = int32(1)
	v254 = v230 + v251
	if v254 != int32(2) {
		v230 = v254
		v233 = v233 + v251
		goto L40
	} else {
		goto L42
	}
L41:
	;
	goto L31
L42:
	;
	goto L41
L43:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v465 = v463 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v465
	v467 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[105]))) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[106]))) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[107]))) = base.I64_extend_i32_s(v465)
	v478 = int32(0)
	v485 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v485 == v478 {
		goto L66
	} else {
		goto L67
	}
L44:
	;
	v315 = int64(0)
	goto L47
L45:
	;
	goto L46
L46:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	if base.F32_lt(base.F32_abs(v91), float32(2.1474836e+09)) != 0 {
		goto L60
	} else {
		goto L61
	}
L47:
	;
	v318 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	v319 = base.B2i32(v318 <= v315)
	if v318 <= v315 {
		v438 = v319
		goto L43
	} else {
		goto L49
	}
L48:
	;
	v438 = v319
	goto L43
L49:
	;
	v322 = base.I32_wrap_i64(v315) << (uint(int32(2)) % 32)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v322+v323)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v326+v322)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v328
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*float64)(unsafe.Add(mBase, uint32(v33)+64)) = base.F64_promote_f32(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+60)) = int32(13)
	v334 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+58)) = uint8(v334)
	v336 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+56)) = uint16(v336)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = v330
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v339
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v328)+48))
	v344 = F_pstrdup(m, v341+int32(4))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v344
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v348 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v348)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(2)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v360 = F_vac_bulkdel_one_index(m, v33+int32(48), v325, v358, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v353
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v347)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v350
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	F_pfree(m, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	v368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v368
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v370+v322))) = v360
	v375 = v315 + int64(1)
	v378 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v378 == v368 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v409 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L8
	} else {
		goto L57
	}
L54:
	;
	goto L53
L55:
	;
	v382 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v382 != int32(1) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v385 = int32(4484100)
	v387 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v388 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v387 + v388
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	*(*int32)(unsafe.Add(mBase, uint32(v378))) = v391 + v388
	*(*int64)(unsafe.Add(mBase, uint32(v378+int32(72))+232)) = v375
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	*(*int32)(unsafe.Add(mBase, uint32(v378))) = v399 + v388
	v405 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v405 - v388
	goto L54
L57:
	;
	if v409 == int32(0) {
		v315 = v375
		goto L47
	} else {
		goto L58
	}
L58:
	;
	goto L48
L59:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v414)+16)) = base.F64_convert_i32_s(v420)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	v424 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v423)+24)) = uint8(v424)
	F_parallel_vacuum_process_all_indexes(m, v284, v413, v424)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L8
	} else {
		goto L63
	}
L60:
	;
	v418 = base.I32_trunc_f32_s(v91)
	v420 = v418
	goto L59
L61:
	;
	goto L62
L62:
	;
	v420 = int32(-2147483648)
	goto L59
L63:
	;
	v429 = F_lazy_check_wraparound_failsafe(m, l0)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	v438 = v429 ^ int32(1)
	goto L43
L65:
	;
	if v438 == int32(0) {
		goto L13
	} else {
		goto L82
	}
L66:
	;
	goto L65
L67:
	;
	goto L68
L68:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v491&int32(1) == int32(0) {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v496 = int32(4484100)
	v498 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v499 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v498 + v499
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = v502 + v499
	goto L71
L70:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	v633 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = v632 + v633
	v636 = int32(4484100)
	v638 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v638 - v633
	goto L66
L71:
	;
	goto L73
L73:
	;
	goto L74
L74:
	;
	goto L78
L78:
	;
	v597 = int32(0)
	v600 = v478
	goto L79
L79:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(24)+v600<<(uint(int32(2))%32))))
	v610 = int32(3)
	v616 = *(*int64)(unsafe.Add(mBase, uint32(v33+int32(4144)+v600<<(uint(v610)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v485+int32(232)+v609<<(uint(v610)%32)))) = v616
	v618 = int32(1)
	v621 = v597 + v618
	if v621 != int32(3) {
		v597 = v621
		v600 = v600 + v618
		goto L79
	} else {
		goto L81
	}
L80:
	;
	goto L70
L81:
	;
	goto L80
L82:
	;
	v653 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v653
	v659 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v659 == v653 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(3)
	v693 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v694 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v694)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v701 = F_palloc0(m, int32(16))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L8
	} else {
		goto L88
	}
L84:
	;
	goto L83
L85:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v663 != int32(1) {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v666 = int32(4484100)
	v668 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v669 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v668 + v669
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	*(*int32)(unsafe.Add(mBase, uint32(v659))) = v672 + v669
	*(*int64)(unsafe.Add(mBase, uint32(v659+int32(0))+232)) = int64(3)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	*(*int32)(unsafe.Add(mBase, uint32(v659))) = v680 + v669
	v686 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v686 - v669
	goto L84
L87:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v764 = F_read_stream_begin_relation(m, int32(9), v760, v761, int32(190), v701, int32(8))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L8
	} else {
		goto L95
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v701))) = v699
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v699)+4))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v699)+8))
	if v705 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v707 = F_palloc0(m, int32(120))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L8
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v737 = F_palloc0(m, int32(88))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L8
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v707))) = v704
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v704)+4))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v711)+24))
	v713 = F_dsa_get_address(m, v710, v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v716)+48))
	v719 = base.I32_div_s(v717, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v707)+104)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(v707)+100)) = v719
	v723 = v707 + int32(4)
	v724 = int32(12)
	v726 = v723 + v719*v724
	*(*int32)(unsafe.Add(mBase, uint32(v726)+4)) = v713
	*(*int32)(unsafe.Add(mBase, uint32(v726))) = v712
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v707)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v723+v729*v724)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v701)+4)) = v707
	goto L87
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v737))) = v704
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v740)+24))
	v744 = base.I32_div_s(v742, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v737)+72)) = v744
	*(*int32)(unsafe.Add(mBase, uint32(v737)+68)) = v744
	v749 = v737 + v744<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v749)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v749)+4)) = v741
	*(*int32)(unsafe.Add(mBase, uint32(v701)+4)) = v737
	goto L87
L95:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	v771 = F_read_stream_next_buffer(m, v764, v33+int32(40))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	if v771 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v778 = v771
	v791 = v2
	goto L101
L99:
	;
	v1793 = v2
	goto L100
L100:
	;
	F_read_stream_end(m, v764)
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L8
	} else {
		goto L276
	}
L101:
	;
	if v778 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v1793 = v1770
	goto L100
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v821
	v824 = v33 + int32(48)
	v825 = int32(0)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v827)+4))
	v829 = int32(*(*int8)(unsafe.Add(mBase, uint32(v828)+1)))
	if v829 != 0 {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v806 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v806+(v778^int32(-1))<<(uint(int32(6))%32))+16))
	v821 = v812
	goto L103
L105:
	;
	goto L106
L106:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v814+v778<<(uint(int32(6))%32)+int32(-64))+16))
	v821 = v820
	goto L103
L107:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_visibilitymap_pin(m, v990, v821, v33+int32(24))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L8
	} else {
		goto L129
	}
L108:
	;
	v844 = v825
	v845 = v829
	v847 = v825
	goto L114
L109:
	;
	if int32(0) < v829 {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v833 = int32(0)
	v834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v828)+2)))
	if v834 == v833 {
		v989 = v833
		goto L107
	} else {
		goto L113
	}
L112:
	;
	v989 = int32(0)
	goto L107
L113:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v824))) = uint16(v834)
	v989 = int32(1)
	goto L107
L114:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v828+int32(4)+v847<<(uint(int32(2))%32))))
	if v874 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v989 = v928
	goto L107
L116:
	;
	v878 = v874
	v880 = v844
	v881 = v847 << (uint(int32(5)) % 32)
	goto L119
L117:
	;
	v928 = v844
	v929 = v845
	goto L118
L118:
	;
	v956 = v847 + int32(1)
	if v956 < base.I32_extend8_s(v929) {
		v844 = v928
		v845 = v929
		v847 = v956
		goto L114
	} else {
		goto L128
	}
L119:
	;
	if v878&int32(1) != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828)+1)))
	v928 = v917
	v929 = v924
	goto L118
L121:
	;
	if v880 < int32(2048) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	v917 = v880
	goto L123
L123:
	;
	v918 = int32(1)
	if base.Ui32(v918) < base.Ui32(v878) {
		v878 = int32(base.Ui32(v878) >> (uint(v918) % 32))
		v880 = v917
		v881 = v881 + v918
		goto L119
	} else {
		goto L127
	}
L124:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v824+v880<<(uint(int32(1))%32)))) = uint16(v881)
	goto L126
L125:
	;
	goto L126
L126:
	;
	v917 = v880 + int32(1)
	goto L123
L127:
	;
	goto L120
L128:
	;
	goto L115
L129:
	;
	F_LockBuffer(m, v778, int32(2))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L8
	} else {
		goto L130
	}
L130:
	;
	v998 = int32(0)
	v999 = base.B2i32(v998 <= v778)
	if v999 == v998 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v1023 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v1023 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1003+(v778^int32(-1))<<(uint(int32(2))%32))))
	v1017 = v1009
	goto L131
L133:
	;
	goto L134
L134:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1017 = v1011 + v778<<(uint(int32(13))%32) + int32(-8192)
	goto L131
L135:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(3)
	v1057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v1058 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1058)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v821
	v1063 = int32(4484100)
	v1065 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v1065 + int32(1)
	if v1058 < v989 {
		goto L139
	} else {
		goto L140
	}
L136:
	;
	goto L135
L137:
	;
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v1027 != int32(1) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v1030 = int32(4484100)
	v1032 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v1033 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v1032 + v1033
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	*(*int32)(unsafe.Add(mBase, uint32(v1023))) = v1036 + v1033
	*(*int64)(unsafe.Add(mBase, uint32(v1023+int32(24))+232)) = base.I64_extend_i32_u(v821)
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	*(*int32)(unsafe.Add(mBase, uint32(v1023))) = v1044 + v1033
	v1050 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v1050 - v1033
	goto L136
L139:
	;
	v1071 = int32(1)
	v1074 = v1017 + int32(24)
	if v989 != v1071 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v1199 = v1058
	goto L141
L141:
	;
	v1228 = int32(0)
	v1234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1017)+12)))
	if base.Ui32(v1234) < base.Ui32(int32(25)) {
		goto L152
	} else {
		goto L153
	}
L142:
	;
	v1081 = v1058
	v1088 = int32(0)
	goto L145
L143:
	;
	v1150 = v1058
	goto L144
L144:
	;
	if v989&v1071 != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v1111 = v1081 << (uint(int32(1)) % 32)
	v1113 = v33 + int32(48)
	v1115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1111+v1113))))
	v1116 = int32(2)
	v1119 = int32(4)
	v1121 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1115<<(uint(v1116)%32)+v1074-v1119))) = v1121
	v1124 = v33 + int32(4144)
	*(*uint16)(unsafe.Add(mBase, uint32(v1124+v1111))) = uint16(v1115)
	v1128 = v1111 | v1116
	v1132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1128+v1113))))
	*(*int32)(unsafe.Add(mBase, uint32(v1132<<(uint(v1116)%32)+v1074-v1119))) = v1121
	*(*uint16)(unsafe.Add(mBase, uint32(v1124+v1128))) = uint16(v1132)
	v1145 = v1081 + v1116
	v1147 = v1088 + v1116
	if v1147 != v989&int32(2147483646) {
		v1081 = v1145
		v1088 = v1147
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v1150 = v1145
	goto L144
L147:
	;
	goto L146
L148:
	;
	v1180 = v1150 << (uint(int32(1)) % 32)
	v1184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1180+(v33+int32(48))))))
	*(*int32)(unsafe.Add(mBase, uint32(v1184<<(uint(int32(2))%32)+v1074-int32(4)))) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v33+int32(4144)+v1180))) = uint16(v1184)
	goto L150
L149:
	;
	goto L150
L150:
	;
	v1199 = v989
	goto L141
L151:
	;
	F_MarkBufferDirty(m, v778)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L8
	} else {
		goto L168
	}
L152:
	;
	v1300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1017)+10)))
	v1302 = v1300 & int32(65534)
	*(*uint16)(unsafe.Add(mBase, uint32(v1017)+10)) = uint16(v1302)
	goto L151
L153:
	;
	v1242 = int32(base.Ui32(v1234+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v1242 == int32(0) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v1248 = v1242
	v1249 = v1228
	v1252 = v1228
	goto L156
L155:
	;
	if int32(0) < v1278 {
		goto L164
	} else {
		goto L165
	}
L156:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1248&int32(65535)<<(uint(int32(2))%32)+(v1017+int32(24))-int32(4))))
	v1263 = v1261 & int32(98304)
	if v1248 == int32(1) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v1278 = v1272
	v1280 = int32(0)
	goto L155
L158:
	;
	v1275 = v1248 - int32(1)
	if v1275 != 0 {
		v1248 = v1275
		v1249 = v1272
		v1252 = v1273
		goto L156
	} else {
		goto L163
	}
L159:
	;
	if v1263 != 0 {
		v1272 = v1249
		v1273 = v1252
		goto L158
	} else {
		goto L162
	}
L160:
	;
	if v1252 != 0 {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v1266 = int32(0)
	v1272 = v1249 + base.B2i32(v1263 == v1266)
	v1273 = base.B2i32(v1263 != v1266)
	goto L158
L162:
	;
	v1278 = v1249
	v1280 = int32(1)
	goto L155
L163:
	;
	goto L157
L164:
	;
	v1285 = v1234 - v1278<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1017)+12)) = uint16(v1285)
	goto L166
L165:
	;
	goto L166
L166:
	;
	if v1280 == int32(0) {
		goto L152
	} else {
		goto L167
	}
L167:
	;
	v1289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1017)+10)))
	v1291 = v1289 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1017)+10)) = uint16(v1291)
	goto L151
L168:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+48))
	v1315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1314)+118)))
	if v1315 != int32(112) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1337 = int32(4484100)
	v1339 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v1339 - int32(1)
	if v999 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L170:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v1319 <= int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+32))
	if v1322 != 0 {
		goto L169
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v1324 = int32(0)
	F_log_heap_prune_and_freeze(m, v1313, v778, v1324, v1324, int32(2), v1324, v1324, v1324, v1324, v1324, v1324, v33+int32(4144), v1199)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L8
	} else {
		goto L176
	}
L174:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1313)+40))
	if v1323 != 0 {
		goto L169
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	goto L169
L177:
	;
	if v778 < int32(0) {
		goto L182
	} else {
		goto L183
	}
L178:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1346+(v778^int32(-1))<<(uint(int32(2))%32))))
	v1360 = v1352
	goto L177
L179:
	;
	goto L180
L180:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1360 = v1354 + v778<<(uint(int32(13))%32) + int32(-8192)
	goto L177
L181:
	;
	v1380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1360)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1380) {
		goto L188
	} else {
		goto L189
	}
L182:
	;
	v1364 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1364+(v778^int32(-1))<<(uint(int32(6))%32))+16))
	v1379 = v1370
	goto L181
L183:
	;
	goto L184
L184:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1372+v778<<(uint(int32(6))%32)+int32(-64))+16))
	v1379 = v1378
	goto L181
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v1054
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1057)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1061
	if v999 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L186:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v1637 + int32(1)
	goto L185
L187:
	;
	v1407 = int32(base.Ui32(v1379) >> (uint(int32(16)) % 32))
	v1410 = int32(1)
	v1414 = v1410
	v1417 = v1410
	v1419 = int32(0)
	goto L193
L188:
	;
	v1388 = int32(base.Ui32(v1380+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v1388 != 0 {
		goto L187
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1390 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1390)
	v1392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1017)+10)))
	v1394 = v1392 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1017)+10)) = uint16(v1394)
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1400 = F_visibilitymap_set(m, v1396, v821, v778, int64(0), v1018, v1390, int32(3))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L8
	} else {
		goto L192
	}
L191:
	;
	goto L190
L192:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1402 + int32(1)
	goto L186
L193:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1414)
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1414&int32(65535)<<(uint(int32(2))%32)+(v1360+int32(24))-int32(4))))
	switch int32(base.Ui32(v1451)>>(uint(int32(15))%32)) & int32(3) {
	case 0, 2:
		v1577 = v1417
		v1578 = v1419
		goto L195
	default:
		goto L196
	}
L194:
	;
	v1586 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v1586)
	v1588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1017)+10)))
	v1590 = v1588 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1017)+10)) = uint16(v1590)
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1595 = int32(1)
	v1597 = v1577 & v1595
	if v1597 != 0 {
		goto L243
	} else {
		goto L244
	}
L195:
	;
	v1582 = v1414 + int32(1)
	if base.Ui32(v1582&int32(65535)) <= base.Ui32(v1388) {
		v1414 = v1582
		v1417 = v1577
		v1419 = v1578
		goto L193
	} else {
		goto L242
	}
L196:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[104]))) = uint16(v1414)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[108]))) = uint16(v1379)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[109]))) = uint16(v1407)
	v1459 = int32(98304)
	if v1451&v1459 == v1459 {
		goto L185
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[103]))) = int32(base.Ui32(v1451) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[110]))) = v1360 + v1451&int32(32767)
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[111]))) = v1471
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1476 = F_HeapTupleSatisfiesVacuum(m, v33+int32(4736), v1475, v778)
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L8
	} else {
		goto L198
	}
L198:
	;
	if v1476 != int32(1) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	if base.B2i32(v1476 != int32(1))&base.B2i32(base.Ui32(v1476) <= base.Ui32(int32(4))) != 0 {
		goto L185
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[110])))
	v1499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1498)+20)))
	if v1499&int32(256) == int32(0) {
		goto L185
	} else {
		goto L206
	}
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L8
	} else {
		goto L203
	}
L203:
	;
	F_errmsg_internal(m, int32(97822), int32(0))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L8
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(491715), int32(3708), int32(390804))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L8
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	v1505 = int32(768)
	if v1499&v1505 != v1505 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1498)))
	v1510 = v1509
	goto L209
L208:
	;
	v1510 = int32(2)
	goto L209
L209:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1511))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1510)) == int32(0) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	if v1523 == int32(0) {
		goto L185
	} else {
		goto L214
	}
L211:
	;
	v1523 = base.B2i32(base.Ui32(v1510) < base.Ui32(v1511))
	goto L210
L212:
	;
	goto L213
L213:
	;
	v1523 = int32(base.Ui32(v1510-v1511) >> (uint(int32(31)) % 32))
	goto L210
L214:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1419))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1510)) == int32(0) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	if v1537 != 0 {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	v1537 = base.B2i32(base.Ui32(v1419) < base.Ui32(v1510))
	goto L215
L217:
	;
	goto L218
L218:
	;
	v1537 = base.B2i32(int32(0) < v1510-v1419)
	goto L215
L219:
	;
	v1538 = v1510
	goto L221
L220:
	;
	v1538 = v1419
	goto L221
L221:
	;
	if base.Ui32(int32(2)) < base.Ui32(v1510) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1541 = v1538
	goto L224
L223:
	;
	v1541 = v1419
	goto L224
L224:
	;
	if v1417&int32(1) == int32(0) {
		v1577 = v1417
		v1578 = v1541
		goto L195
	} else {
		goto L225
	}
L225:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[110])))
	v1548 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1547)+20)))
	v1549 = int32(768)
	if v1548&v1549 == v1549 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	if v1574 != 0 {
		goto L239
	} else {
		goto L240
	}
L227:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+4))
	if v1548&int32(4096) != 0 {
		goto L231
	} else {
		goto L232
	}
L228:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1547)))
	if base.Ui32(v1553) <= base.Ui32(int32(2)) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v1574 = int32(1)
	goto L226
L230:
	;
	if base.Ui32(v1548) < base.Ui32(int32(16384)) {
		goto L236
	} else {
		goto L237
	}
L231:
	;
	if v1557 == int32(0) {
		goto L230
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	if base.Ui32(v1557) <= base.Ui32(int32(2)) {
		goto L230
	} else {
		goto L235
	}
L234:
	;
	v1574 = int32(1)
	goto L226
L235:
	;
	v1574 = int32(1)
	goto L226
L236:
	;
	v1574 = int32(0)
	goto L226
L237:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1547)+8))
	if base.Ui32(v1568) <= base.Ui32(int32(2)) {
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1574 = int32(1)
	goto L226
L239:
	;
	v1575 = int32(0)
	goto L241
L240:
	;
	v1575 = v1417
	goto L241
L241:
	;
	v1577 = v1575
	v1578 = v1541
	goto L195
L242:
	;
	goto L194
L243:
	;
	v1598 = int32(3)
	goto L245
L244:
	;
	v1598 = v1595
	goto L245
L245:
	;
	v1599 = F_visibilitymap_set(m, v1592, v821, v778, int64(0), v1018, v1578, v1598)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L8
	} else {
		goto L246
	}
L246:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1601 + int32(1)
	if v1597 == int32(0) {
		goto L185
	} else {
		goto L247
	}
L247:
	;
	goto L186
L248:
	;
	v1695 = int32(4)
	v1696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1691)+14)))
	v1697 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1691)+12)))
	v1698 = v1696 - v1697
	if v1698 <= v1695 {
		goto L253
	} else {
		goto L254
	}
L249:
	;
	v1677 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1677+(v778^int32(-1))<<(uint(int32(2))%32))))
	v1691 = v1683
	goto L248
L250:
	;
	goto L251
L251:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v1691 = v1685 + v778<<(uint(int32(13))%32) + int32(-8192)
	goto L248
L252:
	;
	F_UnlockReleaseBuffer(m, v778)
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L8
	} else {
		goto L271
	}
L253:
	;
	v1701 = v1695
	goto L255
L254:
	;
	v1701 = v1698
	goto L255
L255:
	;
	v1703 = v1701 - int32(4)
	if v1703 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1760 = int32(0)
	goto L252
L257:
	;
	goto L258
L258:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1697) {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	v1760 = v1703
	goto L252
L260:
	;
	v1714 = int32(base.Ui32(v1697+int32(262120)) >> (uint(int32(2)) % 32))
	goto L262
L261:
	;
	v1714 = int32(0)
	goto L262
L262:
	;
	if base.Ui32(v1714&int32(65535)) < base.Ui32(int32(291)) {
		goto L259
	} else {
		goto L263
	}
L263:
	;
	v1719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1691)+10)))
	if v1719&int32(1) == int32(0) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1760 = int32(0)
	goto L252
L265:
	;
	goto L266
L266:
	;
	v1728 = int32(1)
	goto L267
L267:
	;
	v1739 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1728&int32(65535)<<(uint(int32(2))%32)+(v1691+int32(24))-int32(3)))))
	if v1739&int32(384) == int32(0) {
		goto L259
	} else {
		goto L269
	}
L268:
	;
	v1760 = int32(0)
	goto L252
L269:
	;
	v1745 = v1728 + int32(1)
	v1746 = int32(65535)
	if base.Ui32(v1745&v1746) <= base.Ui32(v1714&v1746) {
		v1728 = v1745
		goto L267
	} else {
		goto L270
	}
L270:
	;
	goto L268
L271:
	;
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_RecordPageWithFreeSpace(m, v1763, v821, v1760)
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L8
	} else {
		goto L272
	}
L272:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L8
	} else {
		goto L273
	}
L273:
	;
	v1770 = v791 + int32(1)
	v1773 = F_read_stream_next_buffer(m, v764, v33+int32(40))
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L8
	} else {
		goto L274
	}
L274:
	;
	if v1773 != 0 {
		v778 = v1773
		v791 = v1770
		goto L101
	} else {
		goto L275
	}
L275:
	;
	goto L102
L276:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	F_pfree(m, v1807)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L8
	} else {
		goto L277
	}
L277:
	;
	F_pfree(m, v701)
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L8
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(-1)
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	if v1814 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	F_ReleaseBuffer(m, v1814)
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L8
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v1819 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L8
	} else {
		goto L283
	}
L282:
	;
	goto L281
L283:
	;
	if v1819 != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1823 = *(*int64)(unsafe.Add(mBase, uint32(v1822)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v1793
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v1821
	F_errmsg(m, int32(170393), v33)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L8
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v690
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)) = uint16(v693)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v696
	goto L13
L287:
	;
	F_errfinish(m, int32(491715), int32(2823), int32(307123))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L8
	} else {
		goto L288
	}
L288:
	;
	goto L286
L289:
	;
	F_parallel_vacuum_reset_dead_items(m, v1870)
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L8
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_TidStoreDestroy(m, v1882)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L8
	} else {
		goto L294
	}
L292:
	;
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(104)))) = v1876 + int32(56)
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1873)+24))
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1880
	goto L1
L294:
	;
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1885)))
	v1887 = F_TidStoreCreateLocal(m, v1886)
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L8
	} else {
		goto L295
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1887
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v1890)+8)) = int64(0)
	goto L1
}
func F_leading_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	v2 = l1
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l0 == int32(0) {
		v61 = v2
		v62 = v8
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v135
	goto L1
L3:
	;
	F_dopr_outchmulti(m, l0, v123, l3)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L18
	} else {
		goto L43
	}
L4:
	;
	v66 = base.B2i32(v61 != int32(0))
	if v66 < v62 {
		goto L22
	} else {
		goto L23
	}
L5:
	;
	if v8 <= int32(0) {
		v61 = v2
		v62 = v8
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v2 == int32(0) {
		v123 = v8
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v16 == int32(0) {
		v44 = v15
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v56 = v54 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v56
	if int32(0) < v56 {
		v123 = v56
		goto L3
	} else {
		goto L21
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v44 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v2)
	goto L8
L10:
	;
	if base.Ui32(v15) < base.Ui32(v16) {
		v44 = v15
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v20 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v23 + int32(1)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v27 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v44 = v43
	goto L9
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v15 == v28 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v31 = v15 - v28
	v32 = F_fwrite(m, v28, int32(1), v31, v20)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v32 + v34
	if v31 == v32 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v38)
	goto L15
L21:
	;
	v61 = int32(0)
	v62 = v56
	goto L4
L22:
	;
	F_dopr_outchmulti(m, int32(32), v62-v66, l3)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L18
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v61 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v66
	goto L24
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v76 == int32(0) {
		v105 = v75
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v114 {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v105 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v61)
	goto L27
L29:
	;
	if base.Ui32(v75) < base.Ui32(v76) {
		v105 = v75
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v80 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v83 + int32(1)
	goto L27
L32:
	;
	goto L33
L33:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v87 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v105 = v103
	goto L28
L35:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v75 == v88 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v91 = v75 - v88
	v92 = F_fwrite(m, v88, int32(1), v91, v80)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v92 + v94
	if v92 == v91 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v98 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)) = uint8(v98)
	goto L34
L39:
	;
	v135 = v114 - int32(1)
	goto L2
L40:
	;
	goto L41
L41:
	;
	if int32(0) <= v114 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v135 = v114 + int32(1)
	goto L2
L43:
	;
	v135 = int32(0)
	goto L2
}
func F_leftmostvalue_inet(m *base.Module) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_DirectFunctionCall1Coll(m, int32(1482), int32(0), int32(554341))
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
				v28 = int32(4)
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v30&int32(254) == int32(2) {
					v39 = v28
				} else {
					v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
				}
				if v30 == int32(1) {
					v42 = v28
				} else {
					v42 = v39
				}
				v53 = v42
			} else {
				v43 = int32(1)
				if v24 != 0 {
					v53 = int32(base.Ui32(v22)>>(uint(v43)%32)) - v43
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v54 = int32(1)
			v55 = v16 + v54
			if v18&v54 != 0 {
				v60 = v55
			} else {
				v60 = v16 + int32(4)
			}
			if v18 == int32(1) {
				v63 = int32(4)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
				if v65&int32(254) == int32(2) {
					v74 = v63
				} else {
					v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
				}
				if v65 == int32(1) {
					v77 = v63
				} else {
					v77 = v74
				}
				v90 = v77
			} else {
				v78 = int32(1)
				if v18&v78 != 0 {
					v90 = int32(base.Ui32(v18)>>(uint(v78)%32)) - v78
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v91 = int32(1)
			v95 = F_varstr_levenshtein_less_equal(m, v25, v53, v60, v90, v91, v91, v91, v19, int32(0))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return int32(0)
			} else {
				return v95
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
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
	return v58
L8:
	;
	v58 = int32(0)
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
	v46 = v14
	v50 = int32(0)
	goto L13
L13:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	v58 = v50 - v51
	goto L7
L14:
	;
	v46 = v41
	v50 = v43
	goto L13
L15:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v24 != v26 {
		v41 = v22
		v43 = v24
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v41 = v35
	v43 = int32(0)
	goto L14
L17:
	;
	if v26 == int32(0) {
		v41 = v22
		v43 = v24
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v31 = v23 - int32(1)
	if v31 == int32(0) {
		v41 = v22
		v43 = v24
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v34 = int32(1)
	v35 = v22 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v36 != 0 {
		v21 = v21 + v34
		v22 = v35
		v23 = v31
		v24 = v36
		goto L15
	} else {
		goto L20
	}
L20:
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
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
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int64
	_ = v409
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
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
	v445 = m.ExcPending
	if v445 != 0 {
		goto L24
	} else {
		goto L112
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L24
	} else {
		goto L108
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[518]))
	if v22 <= l0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[519]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25+l0<<(uint(int32(2))%32))))
	if v29 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+24)))
	if v32&int32(2) == int32(0) {
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
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+24)))
	if v43&int32(2) != 0 {
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
	v405 = m.ExcPending
	if v405 != 0 {
		goto L24
	} else {
		goto L104
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L24
	} else {
		goto L101
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
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
	v357 = m.ExcPending
	if v357 != 0 {
		goto L24
	} else {
		goto L93
	}
L14:
	;
	v48 = base.I32_wrap_i64(v42)
	v50 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v53 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	if v53 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = v50
	goto L17
L16:
	;
	v54 = int32(0)
	goto L17
L17:
	;
	if v54 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v57 = int32(4489492)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[182]))
	v61 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	*(*int32)(unsafe.Add(mBase, _consts[182])) = v61
	if v50 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v86 = v50
	goto L20
L20:
	;
	v89 = v39 + int32(80)
	v90 = F_CatalogOpenIndexes(m, v86)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L24
	} else {
		goto L30
	}
L21:
	;
	v71 = v50
	v72 = v53
	goto L23
L22:
	;
	v66 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v72 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	return
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[521])) = v66
	v70 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v71 = v66
	v72 = v70
	goto L23
L26:
	;
	v78 = F_index_open(m, int32(2683), int32(3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v83 = v71
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[182])) = v58
	v86 = v83
	goto L20
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v78
	v82 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v83 = v82
	goto L28
L30:
	;
	v92 = int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	F_ScanKeyInit(m, v39+int32(2128), v92, int32(3), int32(184), v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	F_ScanKeyInit(m, v39+int32(2176), int32(2), int32(4), int32(150), v48)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v111 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v116 = F_systable_beginscan_ordered(m, v109, v111, v112, int32(2), v39+int32(2128))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L24
	} else {
		goto L36
	}
L33:
	;
	F_systable_endscan_ordered(m, v116)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L24
	} else {
		goto L90
	}
L34:
	;
	v304 = F_systable_getnext_ordered(m, v116, int32(1))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L24
	} else {
		goto L83
	}
L35:
	;
	v244 = base.I32_wrap_i64(l1)
	v246 = v244 & int32(2047)
	if v246 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L36:
	;
	v119 = F_systable_getnext_ordered(m, v116, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L24
	} else {
		goto L37
	}
L37:
	;
	if v119 == int32(0) {
		v242 = v92
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+20)))
	if v124&int32(1) != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+22)))
	v128 = v123 + v127
	if v128 == int32(0) {
		v242 = v92
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v48 == v131 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v134 = v128 + int32(8)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+8)))
	v137 = v135 & int32(3)
	if v137 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	F_CatalogTupleDelete(m, v237, v119+int32(4))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L24
	} else {
		goto L71
	}
L44:
	;
	v138 = F_detoast_attr(m, v134)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L24
	} else {
		goto L47
	}
L45:
	;
	v140 = v134
	goto L46
L46:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v143 = int32(base.Ui32(v141) >> (uint(int32(2)) % 32))
	v145 = v143 - int32(4)
	if base.Ui32(v141-int32(8212)) <= base.Ui32(int32(-8197)) {
		goto L8
	} else {
		goto L48
	}
L47:
	;
	v140 = v138
	goto L46
L48:
	;
	if v145 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v137 != 0 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v152 = F__emscripten_memcpy_bulkmem(m, v89, v140+int32(4), v145)
	mBase = m.M
	v153 = v152
	goto L52
L51:
	;
	v153 = v89
	goto L52
L52:
	;
	goto L49
L53:
	;
	F_pfree(m, v140)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L24
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v158 = base.I32_wrap_i64(l1) & int32(2047)
	if v158 <= v145 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	v198 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+62)) = uint8(v198)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+64)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v198)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+56)) = uint16(v198)
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+58)) = uint8(v206)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v158<<(uint(int32(2))%32) + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v39 + int32(76)
	v217 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+52))
	v225 = F_heap_modify_tuple(m, v119, v218, v39-int32(-64), v39+int32(60), v39+int32(56))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L24
	} else {
		goto L68
	}
L58:
	;
	v160 = v158 - v145
	v163 = v39 + int32(76) + v143
	if v163&int32(3) != 0 {
		v190 = v160
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v194 = F__emscripten_memset_bulkmem(m, v163, base.I32_extend8_s(int32(0)), v190)
	mBase = m.M
	goto L67
L60:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v160) {
		v190 = v160
		goto L59
	} else {
		goto L61
	}
L61:
	;
	if v160&int32(3) != 0 {
		v190 = v160
		goto L59
	} else {
		goto L62
	}
L62:
	;
	if base.Ui32(v158+v153) <= base.Ui32(v163) {
		goto L57
	} else {
		goto L63
	}
L63:
	;
	v177 = int32(80)
	v178 = v143 + v39 + v177
	v181 = v39 + v158 + v177
	if base.Ui32(v181) < base.Ui32(v178) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v183 = v178
	goto L66
L65:
	;
	v183 = v181
	goto L66
L66:
	;
	v190 = (v39+int32(76)^int32(-1)+v183-v143)&int32(-4) + int32(4)
	goto L59
L67:
	;
	goto L57
L68:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	F_CatalogTupleUpdateWithInfo(m, v228, v225+int32(4), v225, v90)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L24
	} else {
		goto L69
	}
L69:
	;
	F_pfree(m, v225)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L24
	} else {
		goto L70
	}
L70:
	;
	goto L34
L71:
	;
	v242 = int32(0)
	goto L35
L72:
	;
	v265 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+62)) = uint8(v265)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v265)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v246<<(uint(int32(2))%32) + int32(16)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v39 + int32(76)
	v281 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+52))
	v287 = F_heap_form_tuple(m, v282, v39-int32(-64), v39+int32(60))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L24
	} else {
		goto L79
	}
L73:
	;
	if v244&int32(3) != 0 {
		v262 = v246
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v264 = F__emscripten_memset_bulkmem(m, v89, base.I32_extend8_s(int32(0)), v262)
	mBase = m.M
	goto L78
L75:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v246) {
		v262 = v246
		goto L74
	} else {
		goto L76
	}
L76:
	;
	if base.Ui32(v246+v89) <= base.Ui32(v89) {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v262 = (v246-int32(1))&int32(-4) + int32(4)
	goto L74
L78:
	;
	goto L72
L79:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	F_CatalogTupleInsertWithInfo(m, v290, v287, v90)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L24
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v287)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L24
	} else {
		goto L81
	}
L81:
	;
	if v242 != 0 {
		goto L33
	} else {
		goto L82
	}
L82:
	;
	goto L34
L83:
	;
	if v304 == int32(0) {
		goto L33
	} else {
		goto L84
	}
L84:
	;
	v311 = v304
	goto L85
L85:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	F_CatalogTupleDelete(m, v323, v311+int32(4))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L24
	} else {
		goto L87
	}
L86:
	;
	goto L33
L87:
	;
	v329 = F_systable_getnext_ordered(m, v116, int32(1))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L24
	} else {
		goto L88
	}
L88:
	;
	if v329 != 0 {
		v311 = v329
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	F_CatalogCloseIndexes(m, v90)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L24
	} else {
		goto L91
	}
L91:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
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
	v360 = m.ExcPending
	if v360 != 0 {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v361
	F_errmsg(m, int32(42015), v39)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L24
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(497131), int32(770), int32(356899))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
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
	v377 = m.ExcPending
	if v377 != 0 {
		goto L24
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = l1
	F_errmsg_internal(m, int32(429822), v39+int32(16))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L24
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(497131), int32(780), int32(356899))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
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
	F_errmsg_internal(m, int32(110186), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L24
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(497131), int32(810), int32(356899))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
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
	v408 = m.ExcPending
	if v408 != 0 {
		goto L24
	} else {
		goto L105
	}
L105:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v145
	*(*int64)(unsafe.Add(mBase, uint32(v39)+32)) = v409
	F_errmsg(m, int32(474454), v39+int32(32))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L24
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(497131), int32(153), int32(430998))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
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
	v432 = m.ExcPending
	if v432 != 0 {
		goto L24
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	F_errmsg(m, int32(481361), v17)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L24
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(494386), int32(565), int32(311759))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
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
	v448 = m.ExcPending
	if v448 != 0 {
		goto L24
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l0
	F_errmsg(m, int32(328560), v17+int32(16))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L24
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(494386), int32(573), int32(311759))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v22 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L6
	}
L2:
	;
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v29 = v27 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v29
	if v25 < v29 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	F_MemoryContextDelete(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
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
	v42 = F_SearchSysCache1(m, int32(82), v21)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	F_ReleaseCatCache(m, v53)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L87
	}
L8:
	;
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v46 = v21
	v47 = v2
	v52 = v2
	v53 = v42
	v54 = v2
	v58 = v2
	goto L12
L10:
	;
	v323 = v21
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L84
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v62 = v60 + v61
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+79)))
	if v63 != int32(100) {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v323 = v315
	goto L11
L14:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+130)))
	F_ScanKeyInit(m, v19+int32(48), int32(10), int32(3), int32(184), v46)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v74 = int32(0)
	v76 = int32(1)
	v81 = F_systable_beginscan(m, v39, int32(2666), v76, v74, v76, v19+int32(48))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v62)+132))
	F_ReleaseCatCache(m, v53)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
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
	v91 = v74
	v92 = v47
	v93 = v83
	v97 = v52
	v99 = v54
	goto L23
L22:
	;
	v301 = v47
	v306 = v52
	v308 = v54
	goto L16
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+22)))
	v107 = v105 + v106
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+72)))
	if v108 == int32(99) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_systable_endscan(m, v81)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L71
	}
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)))
	if v112&int32(1) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v244 = v91
	v245 = v92
	v249 = v97
	v250 = v99
	goto L27
L27:
	;
	v251 = F_systable_getnext(m, v81)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L69
	}
L28:
	;
	v178 = F_text_to_cstring(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L51
	}
L29:
	;
	v174 = int32(*(*int8)(unsafe.Add(mBase, uint32(v120))))
	v177 = v174
	goto L28
L30:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v111)+452))
	if int32(0) <= v117 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+26)))
	if v148&int32(8) != 0 {
		goto L44
	} else {
		goto L45
	}
L33:
	;
	v120 = v117 + v107
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+458)))
	if v121 != int32(1) {
		v177 = v120
		goto L28
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v146 = F_nocachegetattr(m, v93, int32(28), v111)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L43
	}
L36:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+456)))
	switch v124 - int32(1) {
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
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v177 = v128
	goto L28
L39:
	;
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v120))))
	v177 = v127
	goto L28
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = base.I32_extend16_s(v124)
	F_errmsg_internal(m, int32(482718), v19+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(326157), int32(70), int32(67716))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
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
	v177 = v146
	goto L28
L44:
	;
	v152 = F_nocachegetattr(m, v93, int32(28), v111)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L48
	}
L47:
	;
	v177 = v152
	goto L28
L48:
	;
	v158 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v107 + v158
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v62 + v158
	F_errmsg_internal(m, int32(277172), v19+int32(32))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(498595), int32(1172), int32(241626))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
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
	v199 = int32(4489440)
	v200 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v197
	v203 = F_stringToNode(m, v178)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L58
	}
L53:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v197 = v180
	v198 = v92
	goto L52
L54:
	;
	goto L55
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v187 = F_AllocSetContextCreateInternal(m, v182, int32(119399), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v190 = F_MemoryContextAlloc(m, v187, int32(12))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v192
	v197 = v187
	v198 = v190
	goto L52
L58:
	;
	v205 = F_expression_planner(m, v203)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v208 = F_palloc0(m, int32(20))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v208))) = int64(4294967689)
	v214 = F_pstrdup(m, v107+int32(4))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v216 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v208)+16)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v208)+12)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v208)+8)) = v214
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v200
	if v97 == v216 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235+v91<<(uint(int32(2))%32)))) = v208
	v244 = v91 + int32(1)
	v245 = v198
	v249 = v235
	v250 = v236
	goto L27
L63:
	;
	v226 = F_palloc(m, int32(32))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v91 < v99 {
		v235 = v97
		v236 = v99
		goto L62
	} else {
		goto L67
	}
L66:
	;
	v235 = v226
	v236 = int32(8)
	goto L62
L67:
	;
	v231 = F_repalloc(m, v97, v99<<(uint(int32(3))%32))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v235 = v231
	v236 = v99 << (uint(int32(1)) % 32)
	goto L62
L69:
	;
	if v251 != 0 {
		v91 = v244
		v92 = v245
		v93 = v251
		v97 = v249
		v99 = v250
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
		v301 = v245
		v306 = v249
		v308 = v250
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
	F_pg_qsort(m, v249, v244, int32(4), int32(1619))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v263 = int32(4489440)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v270 = v268
	v271 = v244
	goto L77
L76:
	;
	goto L75
L77:
	;
	v286 = v271 - int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v249+v286<<(uint(int32(2))%32))))
	v291 = F_lcons(m, v290, v270)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v264
	v301 = v245
	v306 = v249
	v308 = v250
	goto L16
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v291
	if base.Ui32(int32(1)) < base.Ui32(v271) {
		v270 = v291
		v271 = v286
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v319 = F_SearchSysCache1(m, int32(82), v315)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	if v319 != 0 {
		v46 = v315
		v47 = v301
		v52 = v306
		v53 = v319
		v54 = v308
		v58 = v58 | v66
		goto L12
	} else {
		goto L83
	}
L83:
	;
	goto L13
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v323
	F_errmsg_internal(m, int32(50314), v19)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(498595), int32(1131), int32(241626))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
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
	F_sequence_close(m, v39, int32(1))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	if v58&int32(1) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = v448 | int32(524288)
	m.G0 = v19 + int32(96)
	return
L90:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	v405 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v403)+16))
	if v409 != v405 {
		goto L105
	} else {
		goto L106
	}
L91:
	;
	if v47 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	goto L93
L93:
	;
	if v47 == int32(0) {
		goto L89
	} else {
		goto L103
	}
L94:
	;
	v376 = int32(4489440)
	v377 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v374
	v381 = F_palloc0(m, int32(20))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L100
	}
L95:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v374 = v357
	v375 = v47
	goto L94
L96:
	;
	goto L97
L97:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v364 = F_AllocSetContextCreateInternal(m, v359, int32(119399), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v367 = F_MemoryContextAlloc(m, v364, int32(12))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v369 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v367)+8)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v367)+4)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v367))) = v369
	v374 = v364
	v375 = v367
	goto L94
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v381))) = int64(393)
	v386 = F_pstrdup(m, int32(532262))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v381)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v381)+8)) = v386
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v392 = F_lcons(m, v381, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v375))) = v392
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v377
	v401 = v375
	goto L90
L103:
	;
	v401 = v47
	goto L90
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+308)) = v401
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v401)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v401)+8)) = v440 + int32(1)
	goto L89
L105:
	;
	if v409 == int32(0) {
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
	if v405 != 0 {
		goto L115
	} else {
		goto L116
	}
L109:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v403)+28))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v403)+24))
	if v414 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v413 == int32(0) {
		goto L108
	} else {
		goto L114
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v414)+28)) = v413
	goto L110
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+20)) = v413
	goto L110
L114:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v403)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v413)+24)) = v419
	goto L108
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v403)+16)) = v405
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v405)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v403)+28)) = v426
	if v426 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v403)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v403)+16)) = int32(0)
	goto L107
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v426)+24)) = v403
	goto L120
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v405)+20)) = v403
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
	var v57 int32
	_ = v57
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
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	v1 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v1
	v17 = *(*int32)(unsafe.Add(mBase, _consts[580]))
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
	return v150
L2:
	;
	return int32(0)
L3:
	;
	if v21 == int32(0) {
		v150 = v1
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[580]))
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
	v36 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	v41 = F_AllocSetContextCreateInternal(m, v36, int32(60519), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v43 = int32(4489440)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v41
	v47 = int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v48 == int32(0) {
		v87 = v1
		v88 = v47
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v88 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v51 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v87 = v1
	v88 = v47
	goto L7
L10:
	;
	goto L11
L11:
	;
	v54 = v1
	v56 = v1
	v57 = v47
	goto L12
L12:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v54<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v68 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v87 = v78
	v88 = v79
	goto L7
L14:
	;
	v82 = v54 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v82 < v83 {
		v54 = v82
		v56 = v78
		v57 = v79
		goto L12
	} else {
		goto L23
	}
L15:
	;
	v78 = v56
	v79 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v71 = F_parse_hba_line(m, v67, int32(15))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v71 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v78 = v56
	v79 = int32(0)
	goto L14
L20:
	;
	goto L21
L21:
	;
	v76 = F_lappend(m, v56, v71)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v78 = v76
	v79 = v57
	goto L14
L23:
	;
	goto L13
L24:
	;
	F_MemoryContextDelete(m, v41)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L44
	}
L25:
	;
	v125 = F_FreeFile(m, v21)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L37
	}
L26:
	;
	if v87 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v98 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v98 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v114 = F_FreeFile(m, v21)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L35
	}
L32:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[580]))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v104
	F_errmsg(m, int32(167367), v12)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(499538), int32(2708), int32(506641))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	F_MemoryContextDelete(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v44
	*(*int32)(unsafe.Add(mBase, _consts[582])) = int32(0)
	goto L24
L37:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	F_MemoryContextDelete(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v44
	v134 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[582])) = v134
	if v88 == v134 {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[583]))
	if v139 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_MemoryContextDelete(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[584])) = v87
	*(*int32)(unsafe.Add(mBase, _consts[583])) = v41
	v150 = int32(1)
	goto L1
L43:
	;
	goto L42
L44:
	;
	v150 = int32(0)
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
	var v51 int32
	_ = v51
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
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
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
	F_errmsg(m, int32(684985), v10+int32(32))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(492607), int32(1872), int32(168046))
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
	v50 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v50 < v51 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v127 = int32(0)
	goto L19
L19:
	;
	F_list_free_deep(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L49
	}
L20:
	;
	v55 = v50
	goto L23
L21:
	;
	goto L22
L22:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v127 = v118
	goto L19
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v55<<(uint(int32(2))%32))))
	v66 = int32(0)
	if l2 == v66 {
		v88 = v65
		v89 = v66
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	F_load_file(m, v88, l2)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L37
	}
L26:
	;
	v71 = v65
	goto L28
L27:
	;
	if v81 != 0 {
		v88 = v65
		v89 = int32(0)
		goto L25
	} else {
		goto L35
	}
L28:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v73 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L27
L30:
	;
	goto L29
L31:
	;
	v81 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	if v73 == int32(47) {
		v81 = v71
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v71 = v71 + int32(1)
	goto L28
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v65
	v86 = F_psprintf(m, int32(176752), v10+int32(16))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v88 = v86
	v89 = v86
	goto L25
L37:
	;
	v94 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	if v94 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v88
	F_errmsg_internal(m, int32(676601), v10)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if v89 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	F_errfinish(m, int32(492607), int32(1890), int32(168046))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_pfree(m, v89)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v108 = v55 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v108 < v109 {
		v55 = v108
		goto L23
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	goto L24
L49:
	;
	F_pfree(m, v17)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	goto L1
}
func F_local2local(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v54 int32
	_ = v54
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
	return v54 - l0
L2:
	;
	v50 = l1
	v54 = l0
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = l1
	v17 = l2
	v20 = l0
	goto L5
L5:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v50 = v44
	v54 = v42
	goto L1
L7:
	;
	if l6 != 0 {
		v50 = v16
		v54 = v20
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v32 = base.I32_extend8_s(v25)
	if int32(0) <= v32 {
		v39 = v32
		goto L13
	} else {
		goto L14
	}
L10:
	;
	F_report_invalid_encoding(m, l3, v20, v17)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v39)
	v41 = int32(1)
	v42 = v20 + v41
	v44 = v16 + v41
	if v41 < v17 {
		v16 = v44
		v17 = v17 - v41
		v20 = v42
		goto L5
	} else {
		goto L18
	}
L14:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+(l5-int32(128))))))
	if v36 != 0 {
		v39 = v36
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if l6 != 0 {
		v50 = v16
		v54 = v20
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_report_untranslatable_char(m, l3, l4, v20, v17)
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
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
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
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
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
	v17 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17
	v21 = F_pg_snprintf(m, v12, int32(1024), int32(554917), v8)
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
	v28 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[489]))
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
	v43 = int32(327018)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[816])))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v47 == int32(0) {
		v66 = v46
		v67 = v47
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v70 = v37
	goto L11
L11:
	;
	v71 = v70 + v12
	v73 = int32(1024) - v70
	if v73 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L12:
	;
	if v67-v66 != 0 {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v46 != v47 {
		v66 = v46
		v67 = v47
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v51 = v42
	v52 = v43
	goto L16
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v55
		v67 = v56
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v66 = v55
	v67 = v56
	goto L13
L18:
	;
	v59 = int32(1)
	if v55 == v56 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v69 = v37
	goto L22
L21:
	;
	v69 = v41
	goto L22
L22:
	;
	v70 = v69
	goto L11
L23:
	;
	goto L8
L24:
	;
	v185 = F_strlen(m, v181)
	mBase = m.M
	goto L23
L25:
	;
	v181 = l1
	goto L24
L26:
	;
	goto L27
L27:
	;
	v79 = v73 - int32(1)
	if (v71^l1)&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v178)
	v181 = v174
	goto L24
L29:
	;
	v159 = v154
	v160 = v155
	v161 = v156
	goto L51
L30:
	;
	if v149 == int32(0) {
		v174 = v147
		v175 = v148
		goto L28
	} else {
		goto L50
	}
L31:
	;
	v147 = l1
	v148 = v71
	v149 = v79
	goto L30
L32:
	;
	goto L33
L33:
	;
	v83 = int32(0)
	if l1&int32(3) == v83 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v116 == int32(0) {
		v174 = v113
		v175 = v114
		goto L28
	} else {
		goto L43
	}
L35:
	;
	v113 = l1
	v114 = v71
	v115 = v79
	v116 = base.B2i32(v79 != v83)
	goto L34
L36:
	;
	if v79 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v92 = l1
	v93 = v71
	v94 = v79
	goto L38
L38:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v96)
	if v96 == int32(0) {
		v154 = v92
		v155 = v93
		v156 = v94
		goto L29
	} else {
		goto L40
	}
L39:
	;
	v113 = v107
	v114 = v101
	v115 = v103
	v116 = v105
	goto L34
L40:
	;
	v100 = int32(1)
	v101 = v93 + v100
	v103 = v94 - v100
	v104 = int32(0)
	v105 = base.B2i32(v103 != v104)
	v107 = v92 + v100
	if v107&int32(3) == v104 {
		v113 = v107
		v114 = v101
		v115 = v103
		v116 = v105
		goto L34
	} else {
		goto L41
	}
L41:
	;
	if v103 != 0 {
		v92 = v107
		v93 = v101
		v94 = v103
		goto L38
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v119 == int32(0) {
		v147 = v113
		v148 = v114
		v149 = v115
		goto L30
	} else {
		goto L44
	}
L44:
	;
	if base.Ui32(v115) < base.Ui32(int32(4)) {
		v147 = v113
		v148 = v114
		v149 = v115
		goto L30
	} else {
		goto L45
	}
L45:
	;
	v125 = v113
	v126 = v114
	v127 = v115
	goto L46
L46:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v133 = int32(-2139062144)
	if (int32(16843008)-v130|v130)&v133 != v133 {
		v154 = v125
		v155 = v126
		v156 = v127
		goto L29
	} else {
		goto L48
	}
L47:
	;
	v147 = v141
	v148 = v139
	v149 = v143
	goto L30
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v130
	v138 = int32(4)
	v139 = v126 + v138
	v141 = v125 + v138
	v143 = v127 - v138
	if base.Ui32(int32(3)) < base.Ui32(v143) {
		v125 = v141
		v126 = v139
		v127 = v143
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v154 = v147
	v155 = v148
	v156 = v149
	goto L29
L51:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v163)
	if v163 == int32(0) {
		v174 = v159
		v175 = v160
		goto L28
	} else {
		goto L53
	}
L52:
	;
	v174 = v170
	v175 = v168
	goto L28
L53:
	;
	v167 = int32(1)
	v168 = v160 + v167
	v170 = v159 + v167
	v172 = v161 - v167
	if v172 != 0 {
		v159 = v170
		v160 = v168
		v161 = v172
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
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
	v21 = int32(312586)
	goto L5
L4:
	;
	v21 = int32(312582)
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v21
	v24 = v14 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v24
	F_appendStringInfo(m, l0, int32(729748), v8+int32(32))
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
	v34 = v24 + v15
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(741336)
	F_appendStringInfo(m, l0, int32(509772), v8+int32(16))
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
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(730275)
	F_appendStringInfo(m, l0, int32(509772), v8)
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = v7 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			v20 = v18 & int32(1)
			if v20 != 0 {
				v21 = v12
			} else {
				v21 = v7 + int32(4)
			}
			if v18 == int32(1) {
				v24 = int32(4)
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
				if v26&int32(254) == int32(2) {
					v35 = v24
				} else {
					v35 = base.B2i32(v26 == int32(18)) << (uint(v24) % 32)
				}
				if v26 == int32(1) {
					v38 = v24
				} else {
					v38 = v35
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v20 != 0 {
					v49 = int32(base.Ui32(v18)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = int32(1)
			v51 = v14 + v50
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v56 = v54 & v50
			if v56 != 0 {
				v57 = v51
			} else {
				v57 = v14 + int32(4)
			}
			if v54 == int32(1) {
				v60 = int32(4)
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				if v62&int32(254) == int32(2) {
					v71 = v60
				} else {
					v71 = base.B2i32(v62 == int32(18)) << (uint(v60) % 32)
				}
				if v62 == int32(1) {
					v74 = v60
				} else {
					v74 = v71
				}
				v85 = v74
			} else {
				v75 = int32(1)
				if v56 != 0 {
					v85 = int32(base.Ui32(v54)>>(uint(v75)%32)) - v75
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v88 = F_dotrim(m, v21, v49, v57, v85, int32(1), int32(0))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				return v88
			}
		}
	}
}
