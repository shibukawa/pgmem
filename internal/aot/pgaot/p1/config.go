package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ConfigOptionIsVisible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v2&int32(4) != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[31]))
		v9 = F_has_privs_of_role(m, v7, int32(3374))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				v17 = int32(0)
			} else {
				v17 = int32(1)
			}
			return v17
		}
	} else {
		v17 = int32(1)
		return v17
	}
}
func F_GetConfigOptionByName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_find_option(m, l0, v4, l2, int32(21))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			if l1 == int32(0) {
				v31 = v4
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
				v31 = v4
			}
			m.G0 = v8 + int32(32)
			return v31
		} else {
			v22 = F_ConfigOptionIsVisible(m, v12)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
							F_errmsg(m, int32(698653), v8+int32(16))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(155461)
								F_errdetail(m, int32(593301), v8)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499158), int32(5457), int32(381493))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
					if l1 != 0 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
					} else {
					}
					v29 = F_ShowGUCOption(m, v12, int32(1))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = v29
						m.G0 = v8 + int32(32)
						return v31
					}
				}
			}
		}
	}
}
func F_ParseConfigFp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
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
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v276 int32
	_ = v276
	var v290 int32
	_ = v290
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
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
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v404 int32
	_ = v404
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v470 int32
	_ = v470
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v528 int32
	_ = v528
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v606 int32
	_ = v606
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v691 int32
	_ = v691
	var v719 int32
	_ = v719
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v750 int32
	_ = v750
	var v771 int32
	_ = v771
	var v785 int32
	_ = v785
	var v806 int32
	_ = v806
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v887 int32
	_ = v887
	var v898 int32
	_ = v898
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v937 int32
	_ = v937
	var v946 int32
	_ = v946
	var v957 int32
	_ = v957
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1061 int32
	_ = v1061
	var v1072 int32
	_ = v1072
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1111 int32
	_ = v1111
	var v1120 int32
	_ = v1120
	var v1131 int32
	_ = v1131
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1235 int32
	_ = v1235
	var v1246 int32
	_ = v1246
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1303 int32
	_ = v1303
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1362 int32
	_ = v1362
	var v1373 int32
	_ = v1373
	var v1382 int32
	_ = v1382
	var v1396 int32
	_ = v1396
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1489 int32
	_ = v1489
	var v1497 int32
	_ = v1497
	var v1511 int32
	_ = v1511
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1580 int32
	_ = v1580
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1609 int32
	_ = v1609
	var v1622 int32
	_ = v1622
	var v1636 int32
	_ = v1636
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1751 int32
	_ = v1751
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1857 int32
	_ = v1857
	var v1866 int32
	_ = v1866
	var v1894 int32
	_ = v1894
	var v1930 int32
	_ = v1930
	var v1931 int64
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1958 int32
	_ = v1958
	var v1964 int32
	_ = v1964
	v7 = int32(0)
	v34 = m.G0
	v36 = v34 - int32(112)
	m.G0 = v36
	v39 = l2 + int32(1)
	v45 = l2
	v50 = v7
	v51 = int32(-1)
	v52 = v7
	v53 = v7
	v54 = v7
	v55 = v7
	v56 = v7
	v57 = v7
	v58 = v7
	v59 = v7
	v65 = v7
	v66 = v7
	v67 = v36
	v69 = v7
	v70 = v7
	v71 = v7
	v74 = v7
	v75 = v7
	goto L1
L1:
	;
	if v51 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1215])) = v114
	*(*int32)(unsafe.Add(mBase, _consts[1216])) = v115
	v1964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	m.G0 = v36 + int32(112)
	return v1964
L3:
	;
	v78 = int32(16)
	v79 = v67 - v78
	m.G0 = v79
	v82 = v79 - v78
	m.G0 = v82
	v85 = v82 - v78
	m.G0 = v85
	v88 = v85 - int32(160)
	m.G0 = v88
	v91 = v88 - v78
	m.G0 = v91
	v93 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v93)
	v95 = int32(0)
	v97 = *(*int32)(unsafe.Add(mBase, _consts[1215]))
	v99 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v36 + int32(72)
	goto L6
L4:
	;
	v107 = v50
	v108 = v65
	v109 = v66
	v110 = v67
	v111 = v69
	v112 = v70
	v113 = v71
	v114 = v74
	v115 = v75
	goto L5
L5:
	;
	goto L7
L6:
	;
	v107 = v95
	v108 = v85
	v109 = v91
	v110 = v91
	v111 = v82
	v112 = v79
	v113 = v88
	v114 = v97
	v115 = v99
	goto L5
L7:
	;
	if v107 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	goto L2
L9:
	;
	v1930 = int32(m.ExcTag)
	v1931 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1930 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L10:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v1718 != 0 {
		goto L261
	} else {
		goto L262
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v300 = F_GUC_yy_create_buffer(m, l0, v124)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L9
	} else {
		goto L38
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v260 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L9
	} else {
		goto L34
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1216])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1215])) = v113
	v124 = F_emscripten_builtin_malloc(m, int32(92))
	mBase = m.M
	if v124 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v141 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	v130 = F__emscripten_memset_bulkmem(m, v124, base.I32_extend8_s(int32(0)), int32(92))
	mBase = m.M
	goto L17
L17:
	;
	goto L11
L18:
	;
	if v141 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v153 = *(*int32)(unsafe.Add(mBase, _consts[1217]))
	v155 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+56)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = v153
	F_errmsg_internal(m, int32(51624), v36+int32(48))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v189 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v191 = *(*int32)(unsafe.Add(mBase, _consts[1217]))
	v193 = F_palloc(m, int32(28))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_errfinish(m, int32(314169), int32(373), int32(238797))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v193))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v206 = F_pstrdup(m, v191)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+8)) = v206
	if l1 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v218 = F_pstrdup(m, l1)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L29
	}
L27:
	;
	v221 = v59
	v222 = int32(0)
	goto L28
L28:
	;
	v223 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v193)+24)) = v223
	v225 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v193)+20)) = uint16(v225)
	*(*int32)(unsafe.Add(mBase, uint32(v193)+16)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v193)+12)) = v222
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v229 == v223 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v221 = v218
	v222 = v218
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v193
	v236 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v236)
	v1687 = v45
	v1692 = v193
	v1694 = v52
	v1695 = v53
	v1696 = v54
	v1697 = v55
	v1698 = v56
	v1699 = v57
	v1700 = v58
	v1701 = v221
	goto L10
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v193
	goto L30
L32:
	;
	goto L33
L33:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v233)+24)) = v193
	goto L30
L34:
	;
	if v260 == int32(0) {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_errmsg_internal(m, int32(295243), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_errfinish(m, int32(314169), int32(388), int32(238797))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L11
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_GUC_yyensure_buffer_stack(m, v124)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	if v315 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v377 = v45
	v384 = v52
	v385 = v53
	v386 = v54
	v387 = v55
	v388 = v56
	v389 = v57
	v390 = v58
	v404 = int32(0)
	goto L46
L41:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v315+v318<<(uint(int32(2))%32))))
	if v322 == v300 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if v322 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v124)+36))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v324))) = uint8(v325)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v329 = int32(2)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v327+v328<<(uint(v329)%32))))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v124)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+8)) = v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v335+v336<<(uint(v329)%32))))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v124)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v340)+16)) = v341
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v345 = v344
	v346 = v343
	goto L45
L44:
	;
	v345 = v315
	v346 = v318
	goto L45
L45:
	;
	v347 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v345+v346<<(uint(v347)%32)))) = v300
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v355 = v351 + v352<<(uint(v347)%32)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+28)) = v357
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+36)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v124)+80)) = v360
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v364
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)) = uint8(v366)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+48)) = int32(1)
	goto L40
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v417 = F_GUC_yylex(m, v124)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L9
	} else {
		goto L51
	}
L48:
	;
	if v1327 != 0 {
		goto L208
	} else {
		goto L209
	}
L49:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v124)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v431 = F_pstrdup(m, v421)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L9
	} else {
		goto L53
	}
L50:
	;
	if v417 == int32(99) {
		goto L46
	} else {
		goto L52
	}
L51:
	;
	switch v417 {
	case 0:
		v1687 = v377
		v1692 = v124
		v1694 = v384
		v1695 = v385
		v1696 = v386
		v1697 = v387
		v1698 = v388
		v1699 = v389
		v1700 = v390
		v1701 = v59
		goto L10
	case 1, 7:
		goto L49
	case 2, 3, 4, 5, 6:
		v1326 = v377
		v1327 = v417
		v1328 = v388
		v1329 = v389
		v1330 = v390
		goto L48
	default:
		goto L50
	}
L52:
	;
	v1326 = v377
	v1327 = v417
	v1328 = v388
	v1329 = v389
	v1330 = v390
	goto L48
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v442 = F_GUC_yylex(m, v124)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	if v442 == int32(5) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v455 = F_GUC_yylex(m, v124)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L9
	} else {
		goto L58
	}
L56:
	;
	v457 = v442
	v458 = v390
	goto L57
L57:
	;
	v459 = int32(0)
	if base.Ui32(int32(6)) < base.Ui32(v457) {
		v513 = v377
		v514 = v457
		v515 = v388
		v516 = v389
		v517 = v459
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v457 = v455
	v458 = v455
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v560 = v431
	v561 = int32(213100)
	goto L81
L60:
	;
	v542 = int32(4488036)
	v544 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	*(*int32)(unsafe.Add(mBase, _consts[1216])) = v544 + int32(1)
	goto L59
L61:
	;
	if v431 != 0 {
		goto L73
	} else {
		goto L74
	}
L62:
	;
	if int32(1)<<(uint(v457)%32)&int32(90) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v507 = F_GUC_yylex(m, v124)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L9
	} else {
		goto L70
	}
L64:
	;
	if v457 != int32(2) {
		v513 = v377
		v514 = v457
		v515 = v388
		v516 = v389
		v517 = v459
		goto L61
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v124)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v492 = F_pstrdup(m, v482)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L9
	} else {
		goto L69
	}
L67:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v124)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v480 = F_DeescapeQuotedString(m, v470)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v495 = v388
	v496 = v480
	v497 = v480
	goto L63
L69:
	;
	v495 = v492
	v496 = v389
	v497 = v492
	goto L63
L70:
	;
	if v507 == int32(99) {
		goto L59
	} else {
		goto L71
	}
L71:
	;
	if v507 == int32(0) {
		goto L60
	} else {
		goto L72
	}
L72:
	;
	v513 = v507
	v514 = v507
	v515 = v495
	v516 = v496
	v517 = v497
	goto L61
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_pfree(m, v431)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L9
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if v517 == int32(0) {
		v1326 = v513
		v1327 = v514
		v1328 = v515
		v1329 = v516
		v1330 = v458
		goto L48
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_pfree(m, v517)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v1326 = v513
	v1327 = v514
	v1328 = v515
	v1329 = v516
	v1330 = v458
	goto L48
L79:
	;
	if v507 == int32(0) {
		v1687 = v507
		v1692 = v124
		v1694 = v384
		v1695 = v385
		v1696 = v386
		v1697 = v1303
		v1698 = v495
		v1699 = v496
		v1700 = v458
		v1701 = v59
		goto L10
	} else {
		goto L206
	}
L80:
	;
	if v606 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L81:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561))))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560))))
	if v565 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v606 = base.I32_extend8_s(v586) - base.I32_extend8_s(v595)
	goto L80
L83:
	;
	v574 = int32(1)
	if base.Ui32((v565-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	if v564&int32(255) != 0 {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if v564&int32(255) != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v606 = int32(1)
	goto L80
L88:
	;
	v573 = int32(-1)
	goto L90
L89:
	;
	v573 = int32(0)
	goto L90
L90:
	;
	v606 = v573
	goto L80
L91:
	;
	v586 = v565 | int32(32)
	goto L93
L92:
	;
	v586 = v565
	goto L93
L93:
	;
	if base.Ui32((v564-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v595 = v564 | int32(32)
	goto L96
L95:
	;
	v595 = v564
	goto L96
L96:
	;
	if v586 == v595&int32(255) {
		v560 = v560 + v574
		v561 = v561 + v574
		goto L81
	} else {
		goto L97
	}
L97:
	;
	goto L82
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v619 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v620 = F_GetConfFilesInDir(m, v497, l1, l3, v111, v112)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L9
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v911 = v431
	v912 = int32(115502)
	goto L133
L101:
	;
	v623 = v619 - int32(1)
	if v620 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_GUC_yyensure_buffer_stack(m, v124)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L9
	} else {
		goto L123
	}
L103:
	;
	v771 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v771)
	v785 = v750
	goto L102
L104:
	;
	v691 = v624
	goto L118
L105:
	;
	v624 = int32(0)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v624 < v625 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v639 = F_palloc(m, int32(28))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L9
	} else {
		goto L109
	}
L108:
	;
	v785 = v387
	goto L102
L109:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v639))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v652 = F_pstrdup(m, v628)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+8)) = v652
	if l1 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v664 = F_pstrdup(m, l1)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L9
	} else {
		goto L114
	}
L112:
	;
	v667 = v387
	v668 = int32(0)
	goto L113
L113:
	;
	v669 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v639)+24)) = v669
	v671 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v639)+20)) = uint16(v671)
	*(*int32)(unsafe.Add(mBase, uint32(v639)+16)) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v639)+12)) = v668
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v675 == v669 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v667 = v664
	v668 = v664
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v639
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v639
	v750 = v667
	goto L103
L116:
	;
	goto L117
L117:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v680)+24)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v639
	v750 = v667
	goto L103
L118:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v620+v691<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v730 = F_ParseConfigFile(m, v719, int32(1), l1, v623, v39, l3, l4, l5)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L9
	} else {
		goto L120
	}
L119:
	;
	v785 = v387
	goto L102
L120:
	;
	if v730 == int32(0) {
		v750 = v387
		goto L103
	} else {
		goto L121
	}
L121:
	;
	v735 = v691 + int32(1)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v735 < v736 {
		v691 = v735
		goto L118
	} else {
		goto L122
	}
L122:
	;
	goto L119
L123:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	if v818 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_pfree(m, v431)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L9
	} else {
		goto L130
	}
L125:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v818+v821<<(uint(int32(2))%32))))
	if v825 == v806 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	if v825 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v124)+36))
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v827))) = uint8(v828)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v832 = int32(2)
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v830+v831<<(uint(v832)%32))))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v124)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v835)+8)) = v836
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v838+v839<<(uint(v832)%32))))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v124)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v843)+16)) = v844
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v848 = v847
	v849 = v846
	goto L129
L128:
	;
	v848 = v818
	v849 = v821
	goto L129
L129:
	;
	v850 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v848+v849<<(uint(v850)%32)))) = v806
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v858 = v854 + v855<<(uint(v850)%32)
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v859)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+28)) = v860
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v862)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+36)) = v863
	*(*int32)(unsafe.Add(mBase, uint32(v124)+80)) = v863
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v867
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)) = uint8(v869)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+48)) = int32(1)
	goto L124
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_pfree(m, v497)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L9
	} else {
		goto L131
	}
L131:
	;
	v1303 = v785
	goto L79
L132:
	;
	if v957 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L133:
	;
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912))))
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911))))
	if v916 != 0 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v957 = base.I32_extend8_s(v937) - base.I32_extend8_s(v946)
	goto L132
L135:
	;
	v925 = int32(1)
	if base.Ui32((v916-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L143
	} else {
		goto L144
	}
L136:
	;
	if v915&int32(255) != 0 {
		goto L135
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	if v915&int32(255) != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v957 = int32(1)
	goto L132
L140:
	;
	v924 = int32(-1)
	goto L142
L141:
	;
	v924 = int32(0)
	goto L142
L142:
	;
	v957 = v924
	goto L132
L143:
	;
	v937 = v916 | int32(32)
	goto L145
L144:
	;
	v937 = v916
	goto L145
L145:
	;
	if base.Ui32((v915-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v946 = v915 | int32(32)
	goto L148
L147:
	;
	v946 = v915
	goto L148
L148:
	;
	if v937 == v946&int32(255) {
		v911 = v911 + v925
		v912 = v912 + v925
		goto L133
	} else {
		goto L149
	}
L149:
	;
	goto L134
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v971 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v974 = F_ParseConfigFile(m, v497, int32(0), l1, v971-int32(1), v39, l3, l4, l5)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L9
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1085 = v431
	v1086 = int32(410260)
	goto L167
L153:
	;
	if v974 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v978 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v978)
	goto L156
L155:
	;
	goto L156
L156:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_GUC_yyensure_buffer_stack(m, v124)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L9
	} else {
		goto L157
	}
L157:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	if v992 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_pfree(m, v431)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L9
	} else {
		goto L164
	}
L159:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v992+v995<<(uint(int32(2))%32))))
	if v999 == v980 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	if v999 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v124)+36))
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1001))) = uint8(v1002)
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v1006 = int32(2)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1004+v1005<<(uint(v1006)%32))))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v124)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1009)+8)) = v1010
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1012+v1013<<(uint(v1006)%32))))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v124)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1017)+16)) = v1018
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v1022 = v1021
	v1023 = v1020
	goto L163
L162:
	;
	v1022 = v992
	v1023 = v995
	goto L163
L163:
	;
	v1024 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1022+v1023<<(uint(v1024)%32)))) = v980
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v1032 = v1028 + v1029<<(uint(v1024)%32)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1032)))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+28)) = v1034
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1032)))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+36)) = v1037
	*(*int32)(unsafe.Add(mBase, uint32(v124)+80)) = v1037
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1032)))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1040)))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v1041
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1037))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)) = uint8(v1043)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+48)) = int32(1)
	goto L158
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_pfree(m, v497)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L9
	} else {
		goto L165
	}
L165:
	;
	v1303 = v387
	goto L79
L166:
	;
	if v1131 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L167:
	;
	v1089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1086))))
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1085))))
	if v1090 != 0 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v1131 = base.I32_extend8_s(v1111) - base.I32_extend8_s(v1120)
	goto L166
L169:
	;
	v1099 = int32(1)
	if base.Ui32((v1090-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L177
	} else {
		goto L178
	}
L170:
	;
	if v1089&int32(255) != 0 {
		goto L169
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	if v1089&int32(255) != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v1131 = int32(1)
	goto L166
L174:
	;
	v1098 = int32(-1)
	goto L176
L175:
	;
	v1098 = int32(0)
	goto L176
L176:
	;
	v1131 = v1098
	goto L166
L177:
	;
	v1111 = v1090 | int32(32)
	goto L179
L178:
	;
	v1111 = v1090
	goto L179
L179:
	;
	if base.Ui32((v1089-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1120 = v1089 | int32(32)
	goto L182
L181:
	;
	v1120 = v1089
	goto L182
L182:
	;
	if v1111 == v1120&int32(255) {
		v1085 = v1085 + v1099
		v1086 = v1086 + v1099
		goto L167
	} else {
		goto L183
	}
L183:
	;
	goto L168
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1143 = int32(1)
	v1145 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v1148 = F_ParseConfigFile(m, v497, v1143, l1, v1145-v1143, v39, l3, l4, l5)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L9
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1257 = F_palloc(m, int32(28))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L9
	} else {
		goto L200
	}
L187:
	;
	if v1148 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v1152)
	goto L190
L189:
	;
	goto L190
L190:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_GUC_yyensure_buffer_stack(m, v124)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L9
	} else {
		goto L191
	}
L191:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	if v1166 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_pfree(m, v431)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L9
	} else {
		goto L198
	}
L193:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1166+v1169<<(uint(int32(2))%32))))
	if v1173 == v1154 {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	if v1173 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v124)+36))
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1175))) = uint8(v1176)
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v1180 = int32(2)
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1178+v1179<<(uint(v1180)%32))))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v124)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1183)+8)) = v1184
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1186+v1187<<(uint(v1180)%32))))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v124)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1191)+16)) = v1192
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v1196 = v1195
	v1197 = v1194
	goto L197
L196:
	;
	v1196 = v1166
	v1197 = v1169
	goto L197
L197:
	;
	v1198 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1196+v1197<<(uint(v1198)%32)))) = v1154
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v124)+20))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v1206 = v1202 + v1203<<(uint(v1198)%32)
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1206)))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+28)) = v1208
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1206)))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+36)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v124)+80)) = v1211
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1206)))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1214)))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v1215
	v1217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)) = uint8(v1217)
	*(*int32)(unsafe.Add(mBase, uint32(v124)+48)) = int32(1)
	goto L192
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_pfree(m, v497)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L9
	} else {
		goto L199
	}
L199:
	;
	v1303 = v387
	goto L79
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+4)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v1257))) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1272 = F_pstrdup(m, l1)
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L9
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+12)) = v1272
	v1276 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v1277 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+24)) = v1277
	*(*uint16)(unsafe.Add(mBase, uint32(v1257)+20)) = uint16(v1277)
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+16)) = v1276 - int32(1)
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v1284 == v1277 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1257
	v1303 = v387
	goto L79
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1257
	goto L202
L204:
	;
	goto L205
L205:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v1288)+24)) = v1257
	goto L202
L206:
	;
	v377 = v507
	v387 = v1303
	v388 = v495
	v389 = v496
	v390 = v458
	goto L46
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1571
	v1575 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v1575)
	if base.B2i32(l3 < int32(15)) == v1575 {
		goto L247
	} else {
		goto L248
	}
L208:
	;
	v1336 = base.B2i32(v1327 != int32(99))
	goto L210
L209:
	;
	v1336 = int32(0)
	goto L210
L210:
	;
	if v1336 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1349 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L9
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1464 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L9
	} else {
		goto L230
	}
L214:
	;
	if v1349 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L9
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1407 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v1409 = F_palloc(m, int32(28))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L9
	} else {
		goto L221
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1373 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v1373 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = l1
	F_errmsg(m, int32(373131), v36+int32(32))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L9
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_errfinish(m, int32(314169), int32(519), int32(238797))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L9
	} else {
		goto L220
	}
L220:
	;
	goto L217
L221:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1409))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1423 = F_pstrdup(m, int32(211912))
	mBase = m.M
	v1424 = m.ExcPending
	if v1424 != 0 {
		goto L9
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+8)) = v1423
	if l1 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1435 = F_pstrdup(m, l1)
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L9
	} else {
		goto L226
	}
L224:
	;
	v1438 = v386
	v1439 = int32(0)
	goto L225
L225:
	;
	v1440 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+24)) = v1440
	v1442 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1409)+20)) = uint16(v1442)
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+16)) = v1407 - v1442
	*(*int32)(unsafe.Add(mBase, uint32(v1409)+12)) = v1439
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v1448 == v1440 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v1438 = v1435
	v1439 = v1435
	goto L225
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1409
	v1569 = v384
	v1570 = v1438
	v1571 = v1409
	goto L207
L228:
	;
	goto L229
L229:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v1452)+24)) = v1409
	v1569 = v384
	v1570 = v1438
	v1571 = v1409
	goto L207
L230:
	;
	if v1464 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L9
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1524 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	v1526 = F_palloc(m, int32(28))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L9
	} else {
		goto L237
	}
L234:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v124)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1489 = *(*int32)(unsafe.Add(mBase, _consts[1216]))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v1489
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = l1
	F_errmsg(m, int32(696413), v36+int32(16))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L9
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_errfinish(m, int32(314169), int32(529), int32(238797))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L9
	} else {
		goto L236
	}
L236:
	;
	goto L233
L237:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1526))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1540 = F_pstrdup(m, int32(211912))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L9
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1526)+8)) = v1540
	if l1 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1552 = F_pstrdup(m, l1)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L9
	} else {
		goto L242
	}
L240:
	;
	v1555 = v384
	v1556 = int32(0)
	goto L241
L241:
	;
	v1557 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1526)+24)) = v1557
	v1559 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1526)+20)) = uint16(v1559)
	*(*int32)(unsafe.Add(mBase, uint32(v1526)+16)) = v1524
	*(*int32)(unsafe.Add(mBase, uint32(v1526)+12)) = v1556
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v1563 == v1557 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1555 = v1552
	v1556 = v1552
	goto L241
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1526
	v1569 = v1555
	v1570 = v386
	v1571 = v1526
	goto L207
L244:
	;
	goto L245
L245:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v1567)+24)) = v1526
	v1569 = v1555
	v1570 = v386
	v1571 = v1526
	goto L207
L246:
	;
	v1645 = v1327
	v1647 = v385
	goto L256
L247:
	;
	v1580 = v404 + int32(1)
	if v1580 < int32(100) {
		goto L246
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1569
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1570
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1594 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L9
	} else {
		goto L251
	}
L250:
	;
	goto L249
L251:
	;
	if v1594 == int32(0) {
		v1687 = v1326
		v1692 = v124
		v1694 = v1569
		v1695 = v385
		v1696 = v1570
		v1697 = v387
		v1698 = v1328
		v1699 = v1329
		v1700 = v1330
		v1701 = v59
		goto L10
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1569
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1570
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_errcode(m, int32(261))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1569
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1570
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = l1
	F_errmsg(m, int32(701102), v36)
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L9
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1569
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1570
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	F_errfinish(m, int32(314169), int32(549), int32(238797))
	mBase = m.M
	v1636 = m.ExcPending
	if v1636 != 0 {
		goto L9
	} else {
		goto L255
	}
L255:
	;
	v1687 = v1326
	v1692 = v124
	v1694 = v1569
	v1695 = v385
	v1696 = v1570
	v1697 = v387
	v1698 = v1328
	v1699 = v1329
	v1700 = v1330
	v1701 = v59
	goto L10
L256:
	;
	if v1645 == int32(0) {
		v1687 = v1326
		v1692 = v124
		v1694 = v1569
		v1695 = v1647
		v1696 = v1570
		v1697 = v387
		v1698 = v1328
		v1699 = v1329
		v1700 = v1330
		v1701 = v59
		goto L10
	} else {
		goto L258
	}
L258:
	;
	if v1645 == int32(99) {
		v377 = v1326
		v384 = v1569
		v385 = v1647
		v386 = v1570
		v388 = v1328
		v389 = v1329
		v390 = v1330
		v404 = v1580
		goto L46
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1569
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1647
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1570
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v387
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1330
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v59
	v1683 = F_GUC_yylex(m, v124)
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L9
	} else {
		goto L260
	}
L260:
	;
	v1645 = v1683
	v1647 = v1683
	goto L256
L261:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+20))
	if v1719 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L263
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v1694
	*(*int32)(unsafe.Add(mBase, uint32(v36)+76)) = v1695
	*(*int32)(unsafe.Add(mBase, uint32(v36)+84)) = v1696
	*(*int32)(unsafe.Add(mBase, uint32(v36)+88)) = v1697
	*(*int32)(unsafe.Add(mBase, uint32(v36)+92)) = v1687
	*(*int32)(unsafe.Add(mBase, uint32(v36)+96)) = v1698
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v1699
	*(*int32)(unsafe.Add(mBase, uint32(v36)+104)) = v1700
	*(*int32)(unsafe.Add(mBase, uint32(v36)+108)) = v1701
	v1745 = int32(0)
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+20))
	if v1746 == v1745 {
		v1866 = v1745
		goto L270
	} else {
		goto L271
	}
L264:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+20))
	if v1731 != 0 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+12))
	v1725 = v1719 + v1722<<(uint(int32(2))%32)
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1725)))
	if v1718 != v1726 {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1725))) = int32(0)
	goto L264
L267:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+4))
	F_emscripten_builtin_free(m, v1732)
	mBase = m.M
	goto L269
L268:
	;
	goto L269
L269:
	;
	F_emscripten_builtin_free(m, v1718)
	mBase = m.M
	goto L263
L270:
	;
	F_emscripten_builtin_free(m, v1866)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+20)) = int32(0)
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+60))
	F_emscripten_builtin_free(m, v1894)
	mBase = m.M
	F_emscripten_builtin_free(m, v1692)
	mBase = m.M
	goto L8
L271:
	;
	v1751 = v1746
	goto L272
L272:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+12))
	v1785 = v1751 + v1782<<(uint(int32(2))%32)
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1785)))
	if v1786 == int32(0) {
		v1866 = v1751
		goto L270
	} else {
		goto L274
	}
L273:
	;
	v1866 = v1795
	goto L270
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1785))) = int32(0)
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1786)+20))
	if v1791 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1786)+4))
	F_emscripten_builtin_free(m, v1792)
	mBase = m.M
	goto L277
L276:
	;
	goto L277
L277:
	;
	F_emscripten_builtin_free(m, v1786)
	mBase = m.M
	v1795 = int32(0)
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+20))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1796+v1797<<(uint(int32(2))%32)))) = v1795
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+20))
	if v1804 == v1795 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+20))
	if v1857 != 0 {
		v1751 = v1857
		goto L272
	} else {
		goto L289
	}
L279:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+12))
	v1810 = v1804 + v1807<<(uint(int32(2))%32)
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1810)))
	if v1811 == int32(0) {
		goto L278
	} else {
		goto L280
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1810))) = int32(0)
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+20))
	if v1816 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1811)+4))
	F_emscripten_builtin_free(m, v1817)
	mBase = m.M
	goto L283
L282:
	;
	goto L283
L283:
	;
	F_emscripten_builtin_free(m, v1811)
	mBase = m.M
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+20))
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1820+v1821<<(uint(int32(2))%32)))) = int32(0)
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+12))
	if v1827 != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1829 = v1827 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+12)) = v1829
	v1831 = v1829
	goto L286
L285:
	;
	v1831 = v1795
	goto L286
L286:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+20))
	if v1832 == int32(0) {
		goto L278
	} else {
		goto L287
	}
L287:
	;
	v1837 = v1832 + v1831<<(uint(int32(2))%32)
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1837)))
	if v1838 == int32(0) {
		goto L278
	} else {
		goto L288
	}
L288:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1838)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+28)) = v1841
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v1837)))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+80)) = v1844
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+36)) = v1844
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1837)))
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1847)))
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+4)) = v1848
	v1850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1844))))
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+48)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1692)+24)) = uint8(v1850)
	goto L278
L289:
	;
	goto L273
L290:
	;
	v1935 = int32(v1931)
	m.G0 = v110
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v1935)+4))
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1935)))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1938)))
	if v36+int32(72) == v1942 {
		goto L293
	} else {
		goto L294
	}
L291:
	;
	m.ExcPending = 1
	goto L299
L292:
	;
	if v1945 != 0 {
		goto L296
	} else {
		goto L297
	}
L293:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1938)+4))
	v1945 = v1944
	goto L295
L294:
	;
	v1945 = int32(0)
	goto L295
L295:
	;
	goto L292
L296:
	;
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v36)+108))
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v36)+104))
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v36)+96))
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v36)+92))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v45 = v1950
	v50 = v1937
	v51 = v1945
	v52 = v1953
	v53 = v1954
	v54 = v1952
	v55 = v1951
	v56 = v1949
	v57 = v1948
	v58 = v1947
	v59 = v1946
	v65 = v108
	v66 = v109
	v67 = v110
	v69 = v111
	v70 = v112
	v71 = v113
	v74 = v114
	v75 = v115
	goto L1
L297:
	;
	goto L298
L298:
	;
	F___wasm_longjmp(m, v1938, v1937)
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	return int32(0)
L300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_config_enum_lookup_by_value(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v16 = v10
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v8 + int32(16)
	return v18
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if l1 != v21 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = v16 + int32(12)
	if v24 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L4
L9:
	;
	v16 = v24
	goto L3
L10:
	;
	return int32(0)
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
	F_errmsg_internal(m, int32(181449), v8)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(499158), int32(3036), int32(344377))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_set_config_option_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v8 = int32(0)
	v11 = F_set_config_with_handle(m, l0, v8, l1, l2, l3, l4, l5, int32(1), l6, v8)
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
