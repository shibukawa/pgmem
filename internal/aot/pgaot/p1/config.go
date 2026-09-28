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
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_ConfigOptionIsVisible[0]))
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
	var v52 int32
	_ = v52
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
							F_errmsg(m, int32(_a_F_GetConfigOptionByName_0), v8+int32(16))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_GetConfigOptionByName_1)
								v52 = F_errdetail(m, int32(_a_F_GetConfigOptionByName_2), v8)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_GetConfigOptionByName_3), int32(_a_F_GetConfigOptionByName_4), int32(_a_F_GetConfigOptionByName_5))
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v138 int32
	_ = v138
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v199 int32
	_ = v199
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v307 int32
	_ = v307
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v545 int32
	_ = v545
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v599 int32
	_ = v599
	var v608 int32
	_ = v608
	var v619 int32
	_ = v619
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v707 int32
	_ = v707
	var v731 int32
	_ = v731
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v762 int32
	_ = v762
	var v778 int32
	_ = v778
	var v792 int32
	_ = v792
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
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
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v889 int32
	_ = v889
	var v900 int32
	_ = v900
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v935 int32
	_ = v935
	var v944 int32
	_ = v944
	var v955 int32
	_ = v955
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1059 int32
	_ = v1059
	var v1070 int32
	_ = v1070
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1105 int32
	_ = v1105
	var v1114 int32
	_ = v1114
	var v1125 int32
	_ = v1125
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1229 int32
	_ = v1229
	var v1240 int32
	_ = v1240
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1297 int32
	_ = v1297
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1351 int32
	_ = v1351
	var v1362 int32
	_ = v1362
	var v1371 int32
	_ = v1371
	var v1385 int32
	_ = v1385
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1466 int32
	_ = v1466
	var v1474 int32
	_ = v1474
	var v1488 int32
	_ = v1488
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1586 int32
	_ = v1586
	var v1599 int32
	_ = v1599
	var v1613 int32
	_ = v1613
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1826 int32
	_ = v1826
	var v1834 int32
	_ = v1834
	var v1858 int32
	_ = v1858
	var v1889 int32
	_ = v1889
	var v1890 int64
	_ = v1890
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1917 int32
	_ = v1917
	var v1951 int32
	_ = v1951
	v7 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(288)
	m.G0 = v31
	v34 = l2 + int32(1)
	v40 = l2
	v45 = v7
	v46 = v7
	v47 = v7
	v48 = int32(-1)
	v49 = v7
	v50 = v7
	v51 = v7
	v52 = v7
	v53 = v7
	v54 = v7
	v64 = v7
	v65 = v7
	goto L1
L1:
	;
	if v48 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[0])) = v88
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1])) = v87
	v1951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+243)))
	m.G0 = v31 + int32(288)
	return v1951
L3:
	;
	v68 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+243)) = uint8(v68)
	v70 = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+76)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v70
	v80 = v31 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v31 + int32(68)
	goto L6
L4:
	;
	v86 = v45
	v87 = v64
	v88 = v65
	goto L5
L5:
	;
	goto L8
L6:
	;
	v86 = v70
	v87 = v72
	v88 = v74
	goto L5
L7:
	;
	goto L2
L8:
	;
	if v86 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	goto L7
L10:
	;
	v1889 = int32(m.ExcTag)
	v1890 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1889 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L11:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	if v1686 == int32(0) {
		goto L7
	} else {
		goto L266
	}
L12:
	;
	base.MemoryFill(m, v99, int32(0), int32(92))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+76)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v275 = F_GUC_yy_create_buffer(m, l0, v99)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L10
	} else {
		goto L37
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[0])) = v31 + int32(80)
	v99 = F_emscripten_builtin_malloc(m, int32(92))
	mBase = m.M
	if v99 != 0 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v163 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L10
	} else {
		goto L21
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[2])) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v122 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if v122 == int32(0) {
		v1660 = v40
		v1666 = v46
		v1667 = v47
		v1669 = v49
		v1670 = v50
		v1671 = v51
		v1672 = v52
		v1673 = v53
		v1674 = v54
		goto L11
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errmsg_internal(m, int32(_a_F_ParseConfigFp_0), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(390), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v1660 = v40
	v1666 = v46
	v1667 = v47
	v1669 = v49
	v1670 = v50
	v1671 = v51
	v1672 = v52
	v1673 = v53
	v1674 = v54
	goto L11
L21:
	;
	if v163 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[3]))
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+56)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v175
	F_errmsg_internal(m, int32(_a_F_ParseConfigFp_3), v31+int32(48))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[3]))
	v215 = F_palloc(m, int32(28))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(374), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v215))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v228 = F_pstrdup(m, v213)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+8)) = v228
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v240 = F_pstrdup(m, l1)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L10
	} else {
		goto L32
	}
L30:
	;
	v243 = v54
	v244 = int32(0)
	goto L31
L31:
	;
	v245 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+24)) = v245
	v247 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+20)) = uint16(v247)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v244
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v251 == v245 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v243 = v240
	v244 = v240
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v215
	v258 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+243)) = uint8(v258)
	v1660 = v40
	v1666 = v46
	v1667 = v47
	v1669 = v49
	v1670 = v50
	v1671 = v51
	v1672 = v52
	v1673 = v53
	v1674 = v243
	goto L11
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v215
	goto L33
L35:
	;
	goto L36
L36:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+24)) = v215
	goto L33
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v275
	if v275 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v291 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L10
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	F_GUC_yyensure_buffer_stack(m, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L10
	} else {
		goto L45
	}
L41:
	;
	if v291 == int32(0) {
		v1660 = v40
		v1666 = v46
		v1667 = v47
		v1669 = v49
		v1670 = v50
		v1671 = v51
		v1672 = v52
		v1673 = v53
		v1674 = v54
		goto L11
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errmsg_internal(m, int32(_a_F_ParseConfigFp_4), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(399), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	v1660 = v40
	v1666 = v46
	v1667 = v47
	v1669 = v49
	v1670 = v50
	v1671 = v51
	v1672 = v52
	v1673 = v53
	v1674 = v54
	goto L11
L45:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v332)+20))
	if v335 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v396 = v40
	v402 = v46
	v403 = v47
	v405 = v49
	v406 = v50
	v407 = v51
	v408 = v52
	v409 = v53
	v418 = int32(0)
	goto L52
L47:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v335+v338<<(uint(int32(2))%32))))
	if v342 == v331 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	if v342 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v332)+36))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v344))) = uint8(v345)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v332)+20))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	v349 = int32(2)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v347+v348<<(uint(v349)%32))))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v332)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v352)+8)) = v353
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v332)+20))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v355+v356<<(uint(v349)%32))))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v332)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+16)) = v361
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v332)+20))
	v365 = v364
	v366 = v363
	goto L51
L50:
	;
	v365 = v335
	v366 = v338
	goto L51
L51:
	;
	v367 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v365+v366<<(uint(v367)%32)))) = v331
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v332)+20))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	v375 = v371 + v372<<(uint(v367)%32)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+28)) = v377
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+36)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v332)+80)) = v380
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	*(*int32)(unsafe.Add(mBase, uint32(v332)+4)) = v384
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	*(*uint8)(unsafe.Add(mBase, uint32(v332)+24)) = uint8(v386)
	*(*int32)(unsafe.Add(mBase, uint32(v332)+48)) = int32(1)
	goto L46
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v432 = F_GUC_yylex(m, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L10
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1332 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L10
	} else {
		goto L213
	}
L55:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v99)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v446 = F_pstrdup(m, v436)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L10
	} else {
		goto L59
	}
L56:
	;
	if v432 == int32(99) {
		goto L52
	} else {
		goto L58
	}
L57:
	;
	switch v432 {
	case 0:
		v1660 = v396
		v1666 = v402
		v1667 = v403
		v1669 = v405
		v1670 = v406
		v1671 = v407
		v1672 = v408
		v1673 = v409
		v1674 = v54
		goto L11
	case 1, 7:
		goto L55
	case 2, 3, 4, 5, 6:
		v1315 = v396
		v1316 = v432
		v1317 = v407
		v1318 = v408
		v1319 = v409
		goto L54
	default:
		goto L56
	}
L58:
	;
	v1315 = v396
	v1316 = v432
	v1317 = v407
	v1318 = v408
	v1319 = v409
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v458 = F_GUC_yylex(m, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	if v458 == int32(5) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v472 = F_GUC_yylex(m, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L10
	} else {
		goto L64
	}
L62:
	;
	v474 = v458
	v475 = v409
	goto L63
L63:
	;
	v476 = int32(0)
	if base.Ui32(int32(6)) < base.Ui32(v474) {
		v530 = v396
		v531 = v474
		v532 = v407
		v533 = v408
		v534 = v476
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v474 = v472
	v475 = v472
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v577 = v446
	v578 = int32(_a_F_ParseConfigFp_5)
	goto L87
L66:
	;
	v559 = int32(_a_F_ParseConfigFp_6)
	v561 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1])) = v561 + int32(1)
	goto L65
L67:
	;
	if v446 != 0 {
		goto L79
	} else {
		goto L80
	}
L68:
	;
	if int32(1)<<(uint(v474)%32)&int32(90) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v524 = F_GUC_yylex(m, v523)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L10
	} else {
		goto L76
	}
L70:
	;
	if v474 != int32(2) {
		v530 = v396
		v531 = v474
		v532 = v407
		v533 = v408
		v534 = v476
		goto L67
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v99)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v509 = F_pstrdup(m, v499)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L10
	} else {
		goto L75
	}
L73:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v99)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v396
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v497 = F_DeescapeQuotedString(m, v487)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	v511 = v407
	v512 = v497
	v513 = v497
	goto L69
L75:
	;
	v511 = v509
	v512 = v408
	v513 = v509
	goto L69
L76:
	;
	if v524 == int32(99) {
		goto L65
	} else {
		goto L77
	}
L77:
	;
	if v524 == int32(0) {
		goto L66
	} else {
		goto L78
	}
L78:
	;
	v530 = v524
	v531 = v524
	v532 = v511
	v533 = v512
	v534 = v513
	goto L67
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_pfree(m, v446)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L10
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if v534 == int32(0) {
		v1315 = v530
		v1316 = v531
		v1317 = v532
		v1318 = v533
		v1319 = v475
		goto L54
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_pfree(m, v534)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v1315 = v530
	v1316 = v531
	v1317 = v532
	v1318 = v533
	v1319 = v475
	goto L54
L85:
	;
	if v524 == int32(0) {
		v1660 = v524
		v1666 = v402
		v1667 = v403
		v1669 = v405
		v1670 = v1297
		v1671 = v511
		v1672 = v512
		v1673 = v475
		v1674 = v54
		goto L11
	} else {
		goto L212
	}
L86:
	;
	if v619 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L87:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	if v582 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v619 = base.I32_extend8_s(v599) - base.I32_extend8_s(v608)
	goto L86
L89:
	;
	v587 = int32(1)
	if base.Ui32((v582-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L97
	} else {
		goto L98
	}
L90:
	;
	if v581 != 0 {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v581 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v619 = int32(1)
	goto L86
L94:
	;
	v586 = int32(-1)
	goto L96
L95:
	;
	v586 = int32(0)
	goto L96
L96:
	;
	v619 = v586
	goto L86
L97:
	;
	v599 = v582 | int32(32)
	goto L99
L98:
	;
	v599 = v582
	goto L99
L99:
	;
	if base.Ui32((v581-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v608 = v581 | int32(32)
	goto L102
L101:
	;
	v608 = v581
	goto L102
L102:
	;
	if v599 == v608&int32(255) {
		v577 = v577 + v587
		v578 = v578 + v587
		goto L87
	} else {
		goto L103
	}
L103:
	;
	goto L88
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v637 = F_GetConfFilesInDir(m, v513, l1, l3, v31+int32(244), v31+int32(248))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L10
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v913 = v446
	v914 = int32(_a_F_ParseConfigFp_7)
	goto L139
L107:
	;
	v640 = v632 - int32(1)
	if v637 != 0 {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v792
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	F_GUC_yyensure_buffer_stack(m, v818)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L10
	} else {
		goto L129
	}
L109:
	;
	v778 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+243)) = uint8(v778)
	v792 = v762
	goto L108
L110:
	;
	v707 = v641
	goto L124
L111:
	;
	v641 = int32(0)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v31)+244))
	if v641 < v642 {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v31)+248))
	v656 = F_palloc(m, int32(28))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L10
	} else {
		goto L115
	}
L114:
	;
	v792 = v406
	goto L108
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v656))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v669 = F_pstrdup(m, v654)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L10
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v656)+8)) = v669
	if l1 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v681 = F_pstrdup(m, l1)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L10
	} else {
		goto L120
	}
L118:
	;
	v684 = v406
	v685 = int32(0)
	goto L119
L119:
	;
	v686 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v656)+24)) = v686
	v688 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v656)+20)) = uint16(v688)
	*(*int32)(unsafe.Add(mBase, uint32(v656)+16)) = v640
	*(*int32)(unsafe.Add(mBase, uint32(v656)+12)) = v685
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v692 == v686 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v684 = v681
	v685 = v681
	goto L119
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v656
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v656
	v762 = v684
	goto L109
L122:
	;
	goto L123
L123:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v697)+24)) = v656
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v656
	v762 = v684
	goto L109
L124:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v637+v707<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v742 = F_ParseConfigFile(m, v731, int32(1), l1, v640, v34, l3, l4, l5)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L10
	} else {
		goto L126
	}
L125:
	;
	v792 = v406
	goto L108
L126:
	;
	if v742 == int32(0) {
		v762 = v406
		goto L109
	} else {
		goto L127
	}
L127:
	;
	v747 = v707 + int32(1)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v31)+244))
	if v747 < v748 {
		v707 = v747
		goto L124
	} else {
		goto L128
	}
L128:
	;
	goto L125
L129:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	if v821 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v792
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_pfree(m, v446)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L10
	} else {
		goto L136
	}
L131:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v821+v824<<(uint(int32(2))%32))))
	if v828 == v817 {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	if v828 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v818)+36))
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v830))) = uint8(v831)
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v835 = int32(2)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v833+v834<<(uint(v835)%32))))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v818)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v838)+8)) = v839
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v841+v842<<(uint(v835)%32))))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v818)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v846)+16)) = v847
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v851 = v850
	v852 = v849
	goto L135
L134:
	;
	v851 = v821
	v852 = v824
	goto L135
L135:
	;
	v853 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v851+v852<<(uint(v853)%32)))) = v817
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v818)+20))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v818)+12))
	v861 = v857 + v858<<(uint(v853)%32)
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v861)))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v862)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v818)+28)) = v863
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v861)))
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v865)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v818)+36)) = v866
	*(*int32)(unsafe.Add(mBase, uint32(v818)+80)) = v866
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v861)))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v869)))
	*(*int32)(unsafe.Add(mBase, uint32(v818)+4)) = v870
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866))))
	*(*uint8)(unsafe.Add(mBase, uint32(v818)+24)) = uint8(v872)
	*(*int32)(unsafe.Add(mBase, uint32(v818)+48)) = int32(1)
	goto L130
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v792
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_pfree(m, v513)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L10
	} else {
		goto L137
	}
L137:
	;
	v1297 = v792
	goto L85
L138:
	;
	if v955 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L139:
	;
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v913))))
	if v918 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v955 = base.I32_extend8_s(v935) - base.I32_extend8_s(v944)
	goto L138
L141:
	;
	v923 = int32(1)
	if base.Ui32((v918-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L149
	} else {
		goto L150
	}
L142:
	;
	if v917 != 0 {
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if v917 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v955 = int32(1)
	goto L138
L146:
	;
	v922 = int32(-1)
	goto L148
L147:
	;
	v922 = int32(0)
	goto L148
L148:
	;
	v955 = v922
	goto L138
L149:
	;
	v935 = v918 | int32(32)
	goto L151
L150:
	;
	v935 = v918
	goto L151
L151:
	;
	if base.Ui32((v917-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v944 = v917 | int32(32)
	goto L154
L153:
	;
	v944 = v917
	goto L154
L154:
	;
	if v935 == v944&int32(255) {
		v913 = v913 + v923
		v914 = v914 + v923
		goto L139
	} else {
		goto L155
	}
L155:
	;
	goto L140
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v969 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v972 = F_ParseConfigFile(m, v513, int32(0), l1, v969-int32(1), v34, l3, l4, l5)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L10
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1083 = v446
	v1084 = int32(_a_F_ParseConfigFp_8)
	goto L173
L159:
	;
	if v972 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v976 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+243)) = uint8(v976)
	goto L162
L161:
	;
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	F_GUC_yyensure_buffer_stack(m, v988)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v988)+20))
	if v991 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_pfree(m, v446)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L10
	} else {
		goto L170
	}
L165:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v988)+12))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v991+v994<<(uint(int32(2))%32))))
	if v998 == v987 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	if v998 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v988)+36))
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1000))) = uint8(v1001)
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v988)+20))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v988)+12))
	v1005 = int32(2)
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1003+v1004<<(uint(v1005)%32))))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v988)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1008)+8)) = v1009
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v988)+20))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v988)+12))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1011+v1012<<(uint(v1005)%32))))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v988)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1016)+16)) = v1017
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v988)+12))
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v988)+20))
	v1021 = v1020
	v1022 = v1019
	goto L169
L168:
	;
	v1021 = v991
	v1022 = v994
	goto L169
L169:
	;
	v1023 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1021+v1022<<(uint(v1023)%32)))) = v987
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v988)+20))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v988)+12))
	v1031 = v1027 + v1028<<(uint(v1023)%32)
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1031)))
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1032)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v988)+28)) = v1033
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1031)))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1035)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v988)+36)) = v1036
	*(*int32)(unsafe.Add(mBase, uint32(v988)+80)) = v1036
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1031)))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1039)))
	*(*int32)(unsafe.Add(mBase, uint32(v988)+4)) = v1040
	v1042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1036))))
	*(*uint8)(unsafe.Add(mBase, uint32(v988)+24)) = uint8(v1042)
	*(*int32)(unsafe.Add(mBase, uint32(v988)+48)) = int32(1)
	goto L164
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_pfree(m, v513)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	v1297 = v406
	goto L85
L172:
	;
	if v1125 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L173:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1084))))
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083))))
	if v1088 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v1125 = base.I32_extend8_s(v1105) - base.I32_extend8_s(v1114)
	goto L172
L175:
	;
	v1093 = int32(1)
	if base.Ui32((v1088-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L183
	} else {
		goto L184
	}
L176:
	;
	if v1087 != 0 {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	if v1087 != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v1125 = int32(1)
	goto L172
L180:
	;
	v1092 = int32(-1)
	goto L182
L181:
	;
	v1092 = int32(0)
	goto L182
L182:
	;
	v1125 = v1092
	goto L172
L183:
	;
	v1105 = v1088 | int32(32)
	goto L185
L184:
	;
	v1105 = v1088
	goto L185
L185:
	;
	if base.Ui32((v1087-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v1114 = v1087 | int32(32)
	goto L188
L187:
	;
	v1114 = v1087
	goto L188
L188:
	;
	if v1105 == v1114&int32(255) {
		v1083 = v1083 + v1093
		v1084 = v1084 + v1093
		goto L173
	} else {
		goto L189
	}
L189:
	;
	goto L174
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1137 = int32(1)
	v1139 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v1142 = F_ParseConfigFile(m, v513, v1137, l1, v1139-v1137, v34, l3, l4, l5)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L10
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1251 = F_palloc(m, int32(28))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L10
	} else {
		goto L206
	}
L193:
	;
	if v1142 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1146 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+243)) = uint8(v1146)
	goto L196
L195:
	;
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	F_GUC_yyensure_buffer_stack(m, v1158)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+20))
	if v1161 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_pfree(m, v446)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L10
	} else {
		goto L204
	}
L199:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+12))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1161+v1164<<(uint(int32(2))%32))))
	if v1168 == v1157 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	if v1168 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+36))
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1158)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1170))) = uint8(v1171)
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+20))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+12))
	v1175 = int32(2)
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v1173+v1174<<(uint(v1175)%32))))
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1178)+8)) = v1179
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+20))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+12))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1181+v1182<<(uint(v1175)%32))))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1186)+16)) = v1187
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+12))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+20))
	v1191 = v1190
	v1192 = v1189
	goto L203
L202:
	;
	v1191 = v1161
	v1192 = v1164
	goto L203
L203:
	;
	v1193 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1191+v1192<<(uint(v1193)%32)))) = v1157
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+20))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1158)+12))
	v1201 = v1197 + v1198<<(uint(v1193)%32)
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1201)))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1158)+28)) = v1203
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1201)))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1158)+36)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v1158)+80)) = v1206
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1201)))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1209)))
	*(*int32)(unsafe.Add(mBase, uint32(v1158)+4)) = v1210
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1206))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1158)+24)) = uint8(v1212)
	*(*int32)(unsafe.Add(mBase, uint32(v1158)+48)) = int32(1)
	goto L198
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_pfree(m, v513)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L10
	} else {
		goto L205
	}
L205:
	;
	v1297 = v406
	goto L85
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1251)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1251)+4)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v1251))) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1266 = F_pstrdup(m, l1)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L10
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1251)+12)) = v1266
	v1270 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v1271 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1251)+24)) = v1271
	*(*uint16)(unsafe.Add(mBase, uint32(v1251)+20)) = uint16(v1271)
	*(*int32)(unsafe.Add(mBase, uint32(v1251)+16)) = v1270 - int32(1)
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v1278 == v1271 {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1251
	v1297 = v406
	goto L85
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1251
	goto L208
L210:
	;
	goto L211
L211:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v1282)+24)) = v1251
	goto L208
L212:
	;
	v396 = v524
	v406 = v1297
	v407 = v511
	v408 = v512
	v409 = v475
	goto L52
L213:
	;
	if v1316 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1547
	v1552 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+243)) = uint8(v1552)
	if base.B2i32(l3 < int32(15)) == v1552 {
		goto L252
	} else {
		goto L253
	}
L215:
	;
	v1337 = base.B2i32(v1316 != int32(99))
	goto L217
L216:
	;
	v1337 = int32(0)
	goto L217
L217:
	;
	if v1337 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	if v1332 != 0 {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	if v1332 != 0 {
		goto L236
	} else {
		goto L237
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L10
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1396 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v1398 = F_palloc(m, int32(28))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L10
	} else {
		goto L227
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1362 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+36)) = v1362 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = l1
	F_errmsg(m, int32(_a_F_ParseConfigFp_9), v31+int32(32))
	mBase = m.M
	v1371 = m.ExcPending
	if v1371 != 0 {
		goto L10
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(529), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L10
	} else {
		goto L226
	}
L226:
	;
	goto L223
L227:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1398))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1412 = F_pstrdup(m, int32(_a_F_ParseConfigFp_10))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1398)+8)) = v1412
	if l1 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1424 = F_pstrdup(m, l1)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L10
	} else {
		goto L232
	}
L230:
	;
	v1427 = v405
	v1428 = int32(0)
	goto L231
L231:
	;
	v1429 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1398)+24)) = v1429
	v1431 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1398)+20)) = uint16(v1431)
	*(*int32)(unsafe.Add(mBase, uint32(v1398)+16)) = v1396 - v1431
	*(*int32)(unsafe.Add(mBase, uint32(v1398)+12)) = v1428
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v1437 == v1429 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v1427 = v1424
	v1428 = v1424
	goto L231
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1398
	v1546 = v402
	v1547 = v1398
	v1548 = v1427
	goto L214
L234:
	;
	goto L235
L235:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v1441)+24)) = v1398
	v1546 = v402
	v1547 = v1398
	v1548 = v1427
	goto L214
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L10
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1501 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v1503 = F_palloc(m, int32(28))
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L10
	} else {
		goto L242
	}
L239:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v99)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1466 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v1455
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v1466
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = l1
	F_errmsg(m, int32(_a_F_ParseConfigFp_11), v31+int32(16))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L10
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(539), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L10
	} else {
		goto L241
	}
L241:
	;
	goto L238
L242:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1503))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1517 = F_pstrdup(m, int32(_a_F_ParseConfigFp_10))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L10
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1503)+8)) = v1517
	if l1 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1529 = F_pstrdup(m, l1)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L10
	} else {
		goto L247
	}
L245:
	;
	v1532 = v402
	v1533 = int32(0)
	goto L246
L246:
	;
	v1534 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1503)+24)) = v1534
	v1536 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1503)+20)) = uint16(v1536)
	*(*int32)(unsafe.Add(mBase, uint32(v1503)+16)) = v1501
	*(*int32)(unsafe.Add(mBase, uint32(v1503)+12)) = v1533
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v1540 == v1534 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v1532 = v1529
	v1533 = v1529
	goto L246
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1503
	v1546 = v1532
	v1547 = v1503
	v1548 = v405
	goto L214
L249:
	;
	goto L250
L250:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+24)) = v1503
	v1546 = v1532
	v1547 = v1503
	v1548 = v405
	goto L214
L251:
	;
	v1621 = v1316
	v1623 = v403
	goto L261
L252:
	;
	v1557 = v418 + int32(1)
	if v1557 < int32(100) {
		goto L251
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v1546
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v1548
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1571 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L10
	} else {
		goto L256
	}
L255:
	;
	goto L254
L256:
	;
	if v1571 == int32(0) {
		v1660 = v1315
		v1666 = v1546
		v1667 = v403
		v1669 = v1548
		v1670 = v406
		v1671 = v1317
		v1672 = v1318
		v1673 = v1319
		v1674 = v54
		goto L11
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v1546
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v1548
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errcode(m, int32(261))
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		goto L10
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v1546
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v1548
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = l1
	F_errmsg(m, int32(_a_F_ParseConfigFp_12), v31)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L10
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v1546
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v1548
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(559), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L10
	} else {
		goto L260
	}
L260:
	;
	v1660 = v1315
	v1666 = v1546
	v1667 = v403
	v1669 = v1548
	v1670 = v406
	v1671 = v1317
	v1672 = v1318
	v1673 = v1319
	v1674 = v54
	goto L11
L261:
	;
	if v1621 == int32(0) {
		v1660 = v1315
		v1666 = v1546
		v1667 = v1623
		v1669 = v1548
		v1670 = v406
		v1671 = v1317
		v1672 = v1318
		v1673 = v1319
		v1674 = v54
		goto L11
	} else {
		goto L263
	}
L263:
	;
	if v1621 == int32(99) {
		v396 = v1315
		v402 = v1546
		v403 = v1623
		v405 = v1548
		v407 = v1317
		v408 = v1318
		v409 = v1319
		v418 = v1557
		goto L52
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1623
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v1546
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v1548
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1317
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1319
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v54
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v1656 = F_GUC_yylex(m, v1655)
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L10
	} else {
		goto L265
	}
L265:
	;
	v1621 = v1656
	v1623 = v1656
	goto L261
L266:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	if v1690 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+20))
	if v1691 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+252)) = v1667
	*(*int32)(unsafe.Add(mBase, uint32(v31)+256)) = v1666
	*(*int32)(unsafe.Add(mBase, uint32(v31)+260)) = v1669
	*(*int32)(unsafe.Add(mBase, uint32(v31)+264)) = v1670
	*(*int32)(unsafe.Add(mBase, uint32(v31)+268)) = v1660
	*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v1671
	*(*int32)(unsafe.Add(mBase, uint32(v31)+276)) = v1672
	*(*int32)(unsafe.Add(mBase, uint32(v31)+280)) = v1673
	*(*int32)(unsafe.Add(mBase, uint32(v31)+284)) = v1674
	v1718 = int32(0)
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+20))
	if v1720 == v1718 {
		v1834 = v1718
		goto L276
	} else {
		goto L277
	}
L270:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+20))
	if v1703 != 0 {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+12))
	v1697 = v1691 + v1694<<(uint(int32(2))%32)
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1697)))
	if v1690 != v1698 {
		goto L270
	} else {
		goto L272
	}
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1697))) = int32(0)
	goto L270
L273:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+4))
	F_emscripten_builtin_free(m, v1704)
	mBase = m.M
	goto L275
L274:
	;
	goto L275
L275:
	;
	F_emscripten_builtin_free(m, v1690)
	mBase = m.M
	goto L269
L276:
	;
	F_emscripten_builtin_free(m, v1834)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1719)+20)) = int32(0)
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+60))
	F_emscripten_builtin_free(m, v1858)
	mBase = m.M
	F_emscripten_builtin_free(m, v1719)
	mBase = m.M
	goto L9
L277:
	;
	v1725 = v1720
	goto L278
L278:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+12))
	v1754 = v1725 + v1751<<(uint(int32(2))%32)
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1754)))
	if v1755 == int32(0) {
		v1834 = v1725
		goto L276
	} else {
		goto L280
	}
L279:
	;
	v1834 = v1764
	goto L276
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1754))) = int32(0)
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1755)+20))
	if v1760 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1755)+4))
	F_emscripten_builtin_free(m, v1761)
	mBase = m.M
	goto L283
L282:
	;
	goto L283
L283:
	;
	F_emscripten_builtin_free(m, v1755)
	mBase = m.M
	v1764 = int32(0)
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+20))
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1765+v1766<<(uint(int32(2))%32)))) = v1764
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+20))
	if v1773 == v1764 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+20))
	if v1826 != 0 {
		v1725 = v1826
		goto L278
	} else {
		goto L295
	}
L285:
	;
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+12))
	v1779 = v1773 + v1776<<(uint(int32(2))%32)
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v1779)))
	if v1780 == int32(0) {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1779))) = int32(0)
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+20))
	if v1785 != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+4))
	F_emscripten_builtin_free(m, v1786)
	mBase = m.M
	goto L289
L288:
	;
	goto L289
L289:
	;
	F_emscripten_builtin_free(m, v1780)
	mBase = m.M
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+20))
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1789+v1790<<(uint(int32(2))%32)))) = int32(0)
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+12))
	if v1796 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1798 = v1796 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1719)+12)) = v1798
	v1800 = v1798
	goto L292
L291:
	;
	v1800 = v1764
	goto L292
L292:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+20))
	if v1801 == int32(0) {
		goto L284
	} else {
		goto L293
	}
L293:
	;
	v1806 = v1801 + v1800<<(uint(int32(2))%32)
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1806)))
	if v1807 == int32(0) {
		goto L284
	} else {
		goto L294
	}
L294:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1807)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1719)+28)) = v1810
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1806)))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1719)+80)) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v1719)+36)) = v1813
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1806)))
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1816)))
	*(*int32)(unsafe.Add(mBase, uint32(v1719)+4)) = v1817
	v1819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1813))))
	*(*int32)(unsafe.Add(mBase, uint32(v1719)+48)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1719)+24)) = uint8(v1819)
	goto L284
L295:
	;
	goto L279
L296:
	;
	v1894 = int32(v1890)
	m.G0 = v31
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1894)+4))
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1894)))
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v1897)))
	if v31+int32(68) == v1900 {
		goto L299
	} else {
		goto L300
	}
L297:
	;
	m.ExcPending = 1
	goto L305
L298:
	;
	if v1904 != 0 {
		goto L302
	} else {
		goto L303
	}
L299:
	;
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v1897)+4))
	v1904 = v1902
	goto L301
L300:
	;
	v1904 = int32(0)
	goto L301
L301:
	;
	goto L298
L302:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v31)+284))
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v31)+280))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v31)+276))
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v31)+272))
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v31)+268))
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v31)+264))
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v31)+260))
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v31)+256))
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v31)+252))
	v40 = v1909
	v45 = v1896
	v46 = v1912
	v47 = v1913
	v48 = v1904
	v49 = v1911
	v50 = v1910
	v51 = v1908
	v52 = v1907
	v53 = v1906
	v54 = v1905
	v64 = v87
	v65 = v88
	goto L1
L303:
	;
	goto L304
L304:
	;
	F___wasm_longjmp(m, v1897, v1896)
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	return int32(0)
L306:
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
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v10
	v20 = v13
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if l1 != v21 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v8 + int32(16)
	return v20
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v23 != 0 {
		v18 = v18 + int32(12)
		v20 = v23
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	goto L1
L10:
	;
	return int32(0)
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
	F_errmsg_internal(m, int32(_a_F_config_enum_lookup_by_value_0), v8)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_config_enum_lookup_by_value_1), int32(2938), int32(_a_F_config_enum_lookup_by_value_2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
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
