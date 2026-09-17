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
								F_errdetail(m, int32(_a_F_GetConfigOptionByName_2), v8)
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v434 int32
	_ = v434
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v545 int32
	_ = v545
	var v554 int32
	_ = v554
	var v565 int32
	_ = v565
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v654 int32
	_ = v654
	var v676 int32
	_ = v676
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v707 int32
	_ = v707
	var v722 int32
	_ = v722
	var v736 int32
	_ = v736
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v831 int32
	_ = v831
	var v842 int32
	_ = v842
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v877 int32
	_ = v877
	var v886 int32
	_ = v886
	var v897 int32
	_ = v897
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v1000 int32
	_ = v1000
	var v1011 int32
	_ = v1011
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1046 int32
	_ = v1046
	var v1055 int32
	_ = v1055
	var v1066 int32
	_ = v1066
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1169 int32
	_ = v1169
	var v1180 int32
	_ = v1180
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1237 int32
	_ = v1237
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1290 int32
	_ = v1290
	var v1301 int32
	_ = v1301
	var v1310 int32
	_ = v1310
	var v1324 int32
	_ = v1324
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1376 int32
	_ = v1376
	var v1380 int32
	_ = v1380
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1405 int32
	_ = v1405
	var v1413 int32
	_ = v1413
	var v1427 int32
	_ = v1427
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1496 int32
	_ = v1496
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1525 int32
	_ = v1525
	var v1538 int32
	_ = v1538
	var v1552 int32
	_ = v1552
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1702 int32
	_ = v1702
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1755 int32
	_ = v1755
	var v1764 int32
	_ = v1764
	var v1786 int32
	_ = v1786
	var v1816 int32
	_ = v1816
	var v1817 int64
	_ = v1817
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1850 int32
	_ = v1850
	v7 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(288)
	m.G0 = v30
	v33 = l2 + int32(1)
	v39 = l2
	v44 = v7
	v45 = int32(-1)
	v46 = v7
	v47 = v7
	v48 = v7
	v49 = v7
	v50 = v7
	v51 = v7
	v52 = v7
	v53 = v7
	v62 = v7
	v63 = v7
	goto L1
L1:
	;
	if v45 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[0])) = v84
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1])) = v83
	v1850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+243)))
	m.G0 = v30 + int32(288)
	return v1850
L3:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+243)) = uint8(v66)
	v68 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+76)) = v68
	v76 = v30 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v30 + int32(72)
	goto L6
L4:
	;
	v82 = v44
	v83 = v62
	v84 = v63
	goto L5
L5:
	;
	goto L7
L6:
	;
	v82 = v68
	v83 = v70
	v84 = v72
	goto L5
L7:
	;
	if v82 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	goto L2
L9:
	;
	v1816 = int32(m.ExcTag)
	v1817 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1816 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L10:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
	if v1622 != 0 {
		goto L259
	} else {
		goto L260
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v270 = F_GUC_yy_create_buffer(m, l0, v95)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L9
	} else {
		goto L37
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[2])) = int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v230 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L33
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[0])) = v30 + int32(80)
	v95 = F_emscripten_builtin_malloc(m, int32(92))
	mBase = m.M
	if v95 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v111 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L17
	}
L16:
	;
	base.MemoryFill(m, v95, int32(0), int32(92))
	goto L11
L17:
	;
	if v111 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[3]))
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+56)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v30)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v123
	F_errmsg_internal(m, int32(_a_F_ParseConfigFp_0), v30+int32(48))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[3]))
	v163 = F_palloc(m, int32(28))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(373), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v176 = F_pstrdup(m, v161)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v176
	if l1 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v188 = F_pstrdup(m, l1)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	v191 = v53
	v192 = int32(0)
	goto L27
L27:
	;
	v193 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = v193
	v195 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v163)+20)) = uint16(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = v192
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v199 == v193 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v191 = v188
	v192 = v188
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v163
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+243)) = uint8(v206)
	v1597 = v39
	v1602 = v163
	v1604 = v46
	v1605 = v47
	v1606 = v48
	v1607 = v49
	v1608 = v50
	v1609 = v51
	v1610 = v52
	v1611 = v191
	goto L10
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v163
	goto L29
L31:
	;
	goto L32
L32:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+24)) = v163
	goto L29
L33:
	;
	if v230 == int32(0) {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_errmsg_internal(m, int32(_a_F_ParseConfigFp_3), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(388), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L11
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+76)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_GUC_yyensure_buffer_stack(m, v95)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	if v285 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v347 = v39
	v354 = v46
	v355 = v47
	v356 = v48
	v357 = v49
	v358 = v50
	v359 = v51
	v360 = v52
	v368 = int32(0)
	goto L45
L40:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v285+v288<<(uint(int32(2))%32))))
	if v292 == v270 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if v292 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v294))) = uint8(v295)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v299 = int32(2)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v297+v298<<(uint(v299)%32))))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v302)+8)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v305+v306<<(uint(v299)%32))))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v310)+16)) = v311
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v315 = v314
	v316 = v313
	goto L44
L43:
	;
	v315 = v285
	v316 = v288
	goto L44
L44:
	;
	v317 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v315+v316<<(uint(v317)%32)))) = v270
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v325 = v321 + v322<<(uint(v317)%32)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+36)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v95)+80)) = v330
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v334
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)) = uint8(v336)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+48)) = int32(1)
	goto L39
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v381 = F_GUC_yylex(m, v95)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L9
	} else {
		goto L50
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1271 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L9
	} else {
		goto L206
	}
L48:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v95)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v395 = F_pstrdup(m, v385)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L9
	} else {
		goto L52
	}
L49:
	;
	if v381 == int32(99) {
		goto L45
	} else {
		goto L51
	}
L50:
	;
	switch v381 {
	case 0:
		v1597 = v347
		v1602 = v95
		v1604 = v354
		v1605 = v355
		v1606 = v356
		v1607 = v357
		v1608 = v358
		v1609 = v359
		v1610 = v360
		v1611 = v53
		goto L10
	case 1, 7:
		goto L48
	case 2, 3, 4, 5, 6:
		v1254 = v347
		v1255 = v381
		v1256 = v358
		v1257 = v359
		v1258 = v360
		goto L47
	default:
		goto L49
	}
L51:
	;
	v1254 = v347
	v1255 = v381
	v1256 = v358
	v1257 = v359
	v1258 = v360
	goto L47
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v406 = F_GUC_yylex(m, v95)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	if v406 == int32(5) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v419 = F_GUC_yylex(m, v95)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L9
	} else {
		goto L57
	}
L55:
	;
	v421 = v406
	v422 = v360
	goto L56
L56:
	;
	v423 = int32(0)
	if base.Ui32(int32(6)) < base.Ui32(v421) {
		v476 = v347
		v477 = v421
		v478 = v358
		v479 = v359
		v480 = v423
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v421 = v419
	v422 = v419
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v523 = v395
	v524 = int32(_a_F_ParseConfigFp_4)
	goto L80
L59:
	;
	v505 = int32(_a_F_ParseConfigFp_5)
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1])) = v507 + int32(1)
	goto L58
L60:
	;
	if v395 != 0 {
		goto L72
	} else {
		goto L73
	}
L61:
	;
	if int32(1)<<(uint(v421)%32)&int32(90) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v470 = F_GUC_yylex(m, v95)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L9
	} else {
		goto L69
	}
L63:
	;
	if v421 != int32(2) {
		v476 = v347
		v477 = v421
		v478 = v358
		v479 = v359
		v480 = v423
		goto L60
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v95)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v456 = F_pstrdup(m, v446)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L9
	} else {
		goto L68
	}
L66:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v95)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v444 = F_DeescapeQuotedString(m, v434)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v458 = v358
	v459 = v444
	v460 = v444
	goto L62
L68:
	;
	v458 = v456
	v459 = v359
	v460 = v456
	goto L62
L69:
	;
	if v470 == int32(99) {
		goto L58
	} else {
		goto L70
	}
L70:
	;
	if v470 == int32(0) {
		goto L59
	} else {
		goto L71
	}
L71:
	;
	v476 = v470
	v477 = v470
	v478 = v458
	v479 = v459
	v480 = v460
	goto L60
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_pfree(m, v395)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L9
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v480 == int32(0) {
		v1254 = v476
		v1255 = v477
		v1256 = v478
		v1257 = v479
		v1258 = v422
		goto L47
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_pfree(m, v480)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	v1254 = v476
	v1255 = v477
	v1256 = v478
	v1257 = v479
	v1258 = v422
	goto L47
L78:
	;
	if v470 == int32(0) {
		v1597 = v470
		v1602 = v95
		v1604 = v354
		v1605 = v355
		v1606 = v356
		v1607 = v1237
		v1608 = v458
		v1609 = v459
		v1610 = v422
		v1611 = v53
		goto L10
	} else {
		goto L205
	}
L79:
	;
	if v565 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L80:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	if v528 != 0 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v565 = base.I32_extend8_s(v545) - base.I32_extend8_s(v554)
	goto L79
L82:
	;
	v533 = int32(1)
	if base.Ui32((v528-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L83:
	;
	if v527 != 0 {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v527 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v565 = int32(1)
	goto L79
L87:
	;
	v532 = int32(-1)
	goto L89
L88:
	;
	v532 = int32(0)
	goto L89
L89:
	;
	v565 = v532
	goto L79
L90:
	;
	v545 = v528 | int32(32)
	goto L92
L91:
	;
	v545 = v528
	goto L92
L92:
	;
	if base.Ui32((v527-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v554 = v527 | int32(32)
	goto L95
L94:
	;
	v554 = v527
	goto L95
L95:
	;
	if v545 == v554&int32(255) {
		v523 = v523 + v533
		v524 = v524 + v533
		goto L80
	} else {
		goto L96
	}
L96:
	;
	goto L81
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v578 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v583 = F_GetConfFilesInDir(m, v460, l1, l3, v30+int32(244), v30+int32(248))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L9
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v855 = v395
	v856 = int32(_a_F_ParseConfigFp_6)
	goto L132
L100:
	;
	v586 = v578 - int32(1)
	if v583 != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
	F_GUC_yyensure_buffer_stack(m, v95)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L9
	} else {
		goto L122
	}
L102:
	;
	v722 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+243)) = uint8(v722)
	v736 = v707
	goto L101
L103:
	;
	v654 = v587
	goto L117
L104:
	;
	v587 = int32(0)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v30)+244))
	if v587 < v588 {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v30)+248))
	v602 = F_palloc(m, int32(28))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L9
	} else {
		goto L108
	}
L107:
	;
	v736 = v357
	goto L101
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v602))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v615 = F_pstrdup(m, v600)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v602)+8)) = v615
	if l1 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v627 = F_pstrdup(m, l1)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L9
	} else {
		goto L113
	}
L111:
	;
	v630 = v357
	v631 = int32(0)
	goto L112
L112:
	;
	v632 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v602)+24)) = v632
	v634 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v602)+20)) = uint16(v634)
	*(*int32)(unsafe.Add(mBase, uint32(v602)+16)) = v586
	*(*int32)(unsafe.Add(mBase, uint32(v602)+12)) = v631
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v638 == v632 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v630 = v627
	v631 = v627
	goto L112
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v602
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v602
	v707 = v630
	goto L102
L115:
	;
	goto L116
L116:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v643)+24)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v602
	v707 = v630
	goto L102
L117:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v583+v654<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v687 = F_ParseConfigFile(m, v676, int32(1), l1, v586, v33, l3, l4, l5)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L9
	} else {
		goto L119
	}
L118:
	;
	v736 = v357
	goto L101
L119:
	;
	if v687 == int32(0) {
		v707 = v357
		goto L102
	} else {
		goto L120
	}
L120:
	;
	v692 = v654 + int32(1)
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v30)+244))
	if v692 < v693 {
		v654 = v692
		goto L117
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	if v763 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_pfree(m, v395)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L9
	} else {
		goto L129
	}
L124:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v763+v766<<(uint(int32(2))%32))))
	if v770 == v760 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	if v770 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v772))) = uint8(v773)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v777 = int32(2)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v775+v776<<(uint(v777)%32))))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v780)+8)) = v781
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v783+v784<<(uint(v777)%32))))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v788)+16)) = v789
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v793 = v792
	v794 = v791
	goto L128
L127:
	;
	v793 = v763
	v794 = v766
	goto L128
L128:
	;
	v795 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v793+v794<<(uint(v795)%32)))) = v760
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v803 = v799 + v800<<(uint(v795)%32)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v804)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v805
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v803)))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v807)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+36)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v95)+80)) = v808
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v803)))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v811)))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v812
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v808))))
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)) = uint8(v814)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+48)) = int32(1)
	goto L123
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_pfree(m, v460)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L9
	} else {
		goto L130
	}
L130:
	;
	v1237 = v736
	goto L78
L131:
	;
	if v897 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L132:
	;
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856))))
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v855))))
	if v860 != 0 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v897 = base.I32_extend8_s(v877) - base.I32_extend8_s(v886)
	goto L131
L134:
	;
	v865 = int32(1)
	if base.Ui32((v860-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L142
	} else {
		goto L143
	}
L135:
	;
	if v859 != 0 {
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	if v859 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v897 = int32(1)
	goto L131
L139:
	;
	v864 = int32(-1)
	goto L141
L140:
	;
	v864 = int32(0)
	goto L141
L141:
	;
	v897 = v864
	goto L131
L142:
	;
	v877 = v860 | int32(32)
	goto L144
L143:
	;
	v877 = v860
	goto L144
L144:
	;
	if base.Ui32((v859-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v886 = v859 | int32(32)
	goto L147
L146:
	;
	v886 = v859
	goto L147
L147:
	;
	if v877 == v886&int32(255) {
		v855 = v855 + v865
		v856 = v856 + v865
		goto L132
	} else {
		goto L148
	}
L148:
	;
	goto L133
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v911 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v914 = F_ParseConfigFile(m, v460, int32(0), l1, v911-int32(1), v33, l3, l4, l5)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L9
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1024 = v395
	v1025 = int32(_a_F_ParseConfigFp_7)
	goto L166
L152:
	;
	if v914 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v918 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+243)) = uint8(v918)
	goto L155
L154:
	;
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
	F_GUC_yyensure_buffer_stack(m, v95)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L9
	} else {
		goto L156
	}
L156:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	if v932 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_pfree(m, v395)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L9
	} else {
		goto L163
	}
L158:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v932+v935<<(uint(int32(2))%32))))
	if v939 == v929 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	if v939 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v941))) = uint8(v942)
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v946 = int32(2)
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v944+v945<<(uint(v946)%32))))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v949)+8)) = v950
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v952+v953<<(uint(v946)%32))))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v957)+16)) = v958
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v962 = v961
	v963 = v960
	goto L162
L161:
	;
	v962 = v932
	v963 = v935
	goto L162
L162:
	;
	v964 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v962+v963<<(uint(v964)%32)))) = v929
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v972 = v968 + v969<<(uint(v964)%32)
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v972)))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v973)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v974
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v972)))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+36)) = v977
	*(*int32)(unsafe.Add(mBase, uint32(v95)+80)) = v977
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v972)))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v980)))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v981
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977))))
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)) = uint8(v983)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+48)) = int32(1)
	goto L157
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_pfree(m, v460)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L9
	} else {
		goto L164
	}
L164:
	;
	v1237 = v357
	goto L78
L165:
	;
	if v1066 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L166:
	;
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1025))))
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024))))
	if v1029 != 0 {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	v1066 = base.I32_extend8_s(v1046) - base.I32_extend8_s(v1055)
	goto L165
L168:
	;
	v1034 = int32(1)
	if base.Ui32((v1029-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	if v1028 != 0 {
		goto L168
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	if v1028 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v1066 = int32(1)
	goto L165
L173:
	;
	v1033 = int32(-1)
	goto L175
L174:
	;
	v1033 = int32(0)
	goto L175
L175:
	;
	v1066 = v1033
	goto L165
L176:
	;
	v1046 = v1029 | int32(32)
	goto L178
L177:
	;
	v1046 = v1029
	goto L178
L178:
	;
	if base.Ui32((v1028-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1055 = v1028 | int32(32)
	goto L181
L180:
	;
	v1055 = v1028
	goto L181
L181:
	;
	if v1046 == v1055&int32(255) {
		v1024 = v1024 + v1034
		v1025 = v1025 + v1034
		goto L166
	} else {
		goto L182
	}
L182:
	;
	goto L167
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1078 = int32(1)
	v1080 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v1083 = F_ParseConfigFile(m, v460, v1078, l1, v1080-v1078, v33, l3, l4, l5)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L9
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1191 = F_palloc(m, int32(28))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L9
	} else {
		goto L199
	}
L186:
	;
	if v1083 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1087 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+243)) = uint8(v1087)
	goto L189
L188:
	;
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
	F_GUC_yyensure_buffer_stack(m, v95)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L9
	} else {
		goto L190
	}
L190:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	if v1101 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_pfree(m, v395)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L9
	} else {
		goto L197
	}
L192:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1101+v1104<<(uint(int32(2))%32))))
	if v1108 == v1098 {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	if v1108 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1110))) = uint8(v1111)
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v1115 = int32(2)
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1113+v1114<<(uint(v1115)%32))))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1118)+8)) = v1119
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1121+v1122<<(uint(v1115)%32))))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1126)+16)) = v1127
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v1131 = v1130
	v1132 = v1129
	goto L196
L195:
	;
	v1131 = v1101
	v1132 = v1104
	goto L196
L196:
	;
	v1133 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1131+v1132<<(uint(v1133)%32)))) = v1098
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v1141 = v1137 + v1138<<(uint(v1133)%32)
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1141)))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+28)) = v1143
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1141)))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+36)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(v95)+80)) = v1146
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1141)))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1149)))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v1150
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1146))))
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+24)) = uint8(v1152)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+48)) = int32(1)
	goto L191
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_pfree(m, v460)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L9
	} else {
		goto L198
	}
L198:
	;
	v1237 = v357
	goto L78
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1191)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1191)+4)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(v1191))) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v458
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1206 = F_pstrdup(m, l1)
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L9
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1191)+12)) = v1206
	v1210 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v1211 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1191)+24)) = v1211
	*(*uint16)(unsafe.Add(mBase, uint32(v1191)+20)) = uint16(v1211)
	*(*int32)(unsafe.Add(mBase, uint32(v1191)+16)) = v1210 - int32(1)
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v1218 == v1211 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1191
	v1237 = v357
	goto L78
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1191
	goto L201
L203:
	;
	goto L204
L204:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v1222)+24)) = v1191
	goto L201
L205:
	;
	v347 = v470
	v357 = v1237
	v358 = v458
	v359 = v459
	v360 = v422
	goto L45
L206:
	;
	if v1255 != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1487
	v1491 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+243)) = uint8(v1491)
	if base.B2i32(l3 < int32(15)) == v1491 {
		goto L245
	} else {
		goto L246
	}
L208:
	;
	v1276 = base.B2i32(v1255 != int32(99))
	goto L210
L209:
	;
	v1276 = int32(0)
	goto L210
L210:
	;
	if v1276 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	if v1271 != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	goto L213
L213:
	;
	if v1271 != 0 {
		goto L229
	} else {
		goto L230
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L9
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1335 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v1337 = F_palloc(m, int32(28))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L9
	} else {
		goto L220
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1301 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v1301 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = l1
	F_errmsg(m, int32(_a_F_ParseConfigFp_8), v30+int32(32))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L9
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(519), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L9
	} else {
		goto L219
	}
L219:
	;
	goto L216
L220:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1337))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1351 = F_pstrdup(m, int32(_a_F_ParseConfigFp_9))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L9
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1337)+8)) = v1351
	if l1 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1363 = F_pstrdup(m, l1)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L9
	} else {
		goto L225
	}
L223:
	;
	v1366 = v356
	v1367 = int32(0)
	goto L224
L224:
	;
	v1368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1337)+24)) = v1368
	v1370 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1337)+20)) = uint16(v1370)
	*(*int32)(unsafe.Add(mBase, uint32(v1337)+16)) = v1335 - v1370
	*(*int32)(unsafe.Add(mBase, uint32(v1337)+12)) = v1367
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v1376 == v1368 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v1366 = v1363
	v1367 = v1363
	goto L224
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1337
	v1485 = v355
	v1486 = v1366
	v1487 = v1337
	goto L207
L227:
	;
	goto L228
L228:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v1380)+24)) = v1337
	v1485 = v355
	v1486 = v1366
	v1487 = v1337
	goto L207
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L9
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1440 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	v1442 = F_palloc(m, int32(28))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L9
	} else {
		goto L235
	}
L232:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v95)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1405 = *(*int32)(unsafe.Add(mBase, _c_F_ParseConfigFp[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v1405
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = l1
	F_errmsg(m, int32(_a_F_ParseConfigFp_10), v30+int32(16))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L9
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(529), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L9
	} else {
		goto L234
	}
L234:
	;
	goto L231
L235:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1442))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1456 = F_pstrdup(m, int32(_a_F_ParseConfigFp_9))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L9
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1442)+8)) = v1456
	if l1 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1468 = F_pstrdup(m, l1)
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L9
	} else {
		goto L240
	}
L238:
	;
	v1471 = v355
	v1472 = int32(0)
	goto L239
L239:
	;
	v1473 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1442)+24)) = v1473
	v1475 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1442)+20)) = uint16(v1475)
	*(*int32)(unsafe.Add(mBase, uint32(v1442)+16)) = v1440
	*(*int32)(unsafe.Add(mBase, uint32(v1442)+12)) = v1472
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v1479 == v1473 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1471 = v1468
	v1472 = v1468
	goto L239
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1442
	v1485 = v1471
	v1486 = v356
	v1487 = v1442
	goto L207
L242:
	;
	goto L243
L243:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v1483)+24)) = v1442
	v1485 = v1471
	v1486 = v356
	v1487 = v1442
	goto L207
L244:
	;
	v1561 = v1255
	v1562 = v354
	goto L254
L245:
	;
	v1496 = v368 + int32(1)
	if v1496 < int32(100) {
		goto L244
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1510 = F_errstart(m, l3, int32(0))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L9
	} else {
		goto L249
	}
L248:
	;
	goto L247
L249:
	;
	if v1510 == int32(0) {
		v1597 = v1254
		v1602 = v95
		v1604 = v354
		v1605 = v1485
		v1606 = v1486
		v1607 = v357
		v1608 = v1256
		v1609 = v1257
		v1610 = v1258
		v1611 = v53
		goto L10
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_errcode(m, int32(261))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L9
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = l1
	F_errmsg(m, int32(_a_F_ParseConfigFp_11), v30)
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L9
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	F_errfinish(m, int32(_a_F_ParseConfigFp_1), int32(549), int32(_a_F_ParseConfigFp_2))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L9
	} else {
		goto L253
	}
L253:
	;
	v1597 = v1254
	v1602 = v95
	v1604 = v354
	v1605 = v1485
	v1606 = v1486
	v1607 = v357
	v1608 = v1256
	v1609 = v1257
	v1610 = v1258
	v1611 = v53
	goto L10
L254:
	;
	if v1561 == int32(0) {
		v1597 = v1254
		v1602 = v95
		v1604 = v1562
		v1605 = v1485
		v1606 = v1486
		v1607 = v357
		v1608 = v1256
		v1609 = v1257
		v1610 = v1258
		v1611 = v53
		goto L10
	} else {
		goto L256
	}
L256:
	;
	if v1561 == int32(99) {
		v347 = v1254
		v354 = v1562
		v355 = v1485
		v356 = v1486
		v358 = v1256
		v359 = v1257
		v360 = v1258
		v368 = v1496
		goto L45
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v1562
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1254
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1256
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1257
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1258
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v53
	v1593 = F_GUC_yylex(m, v95)
	mBase = m.M
	v1594 = m.ExcPending
	if v1594 != 0 {
		goto L9
	} else {
		goto L258
	}
L258:
	;
	v1561 = v1593
	v1562 = v1593
	goto L254
L259:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	if v1623 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+256)) = v1605
	*(*int32)(unsafe.Add(mBase, uint32(v30)+252)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v30)+260)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v30)+264)) = v1607
	*(*int32)(unsafe.Add(mBase, uint32(v30)+268)) = v1597
	*(*int32)(unsafe.Add(mBase, uint32(v30)+272)) = v1608
	*(*int32)(unsafe.Add(mBase, uint32(v30)+276)) = v1609
	*(*int32)(unsafe.Add(mBase, uint32(v30)+280)) = v1610
	*(*int32)(unsafe.Add(mBase, uint32(v30)+284)) = v1611
	v1649 = int32(0)
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	if v1650 == v1649 {
		v1764 = v1649
		goto L268
	} else {
		goto L269
	}
L262:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+20))
	if v1635 != 0 {
		goto L265
	} else {
		goto L266
	}
L263:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+12))
	v1629 = v1623 + v1626<<(uint(int32(2))%32)
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1629)))
	if v1622 != v1630 {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1629))) = int32(0)
	goto L262
L265:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+4))
	F_emscripten_builtin_free(m, v1636)
	mBase = m.M
	goto L267
L266:
	;
	goto L267
L267:
	;
	F_emscripten_builtin_free(m, v1622)
	mBase = m.M
	goto L261
L268:
	;
	F_emscripten_builtin_free(m, v1764)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1602)+20)) = int32(0)
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+60))
	F_emscripten_builtin_free(m, v1786)
	mBase = m.M
	F_emscripten_builtin_free(m, v1602)
	mBase = m.M
	goto L8
L269:
	;
	v1655 = v1650
	goto L270
L270:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+12))
	v1683 = v1655 + v1680<<(uint(int32(2))%32)
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1683)))
	if v1684 == int32(0) {
		v1764 = v1655
		goto L268
	} else {
		goto L272
	}
L271:
	;
	v1764 = v1693
	goto L268
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1683))) = int32(0)
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+20))
	if v1689 != 0 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+4))
	F_emscripten_builtin_free(m, v1690)
	mBase = m.M
	goto L275
L274:
	;
	goto L275
L275:
	;
	F_emscripten_builtin_free(m, v1684)
	mBase = m.M
	v1693 = int32(0)
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1694+v1695<<(uint(int32(2))%32)))) = v1693
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	if v1702 == v1693 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	if v1755 != 0 {
		v1655 = v1755
		goto L270
	} else {
		goto L287
	}
L277:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+12))
	v1708 = v1702 + v1705<<(uint(int32(2))%32)
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1708)))
	if v1709 == int32(0) {
		goto L276
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1708))) = int32(0)
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+20))
	if v1714 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+4))
	F_emscripten_builtin_free(m, v1715)
	mBase = m.M
	goto L281
L280:
	;
	goto L281
L281:
	;
	F_emscripten_builtin_free(m, v1709)
	mBase = m.M
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1718+v1719<<(uint(int32(2))%32)))) = int32(0)
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+12))
	if v1725 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1727 = v1725 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1602)+12)) = v1727
	v1729 = v1727
	goto L284
L283:
	;
	v1729 = v1693
	goto L284
L284:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1602)+20))
	if v1730 == int32(0) {
		goto L276
	} else {
		goto L285
	}
L285:
	;
	v1735 = v1730 + v1729<<(uint(int32(2))%32)
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1735)))
	if v1736 == int32(0) {
		goto L276
	} else {
		goto L286
	}
L286:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1736)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1602)+28)) = v1739
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1735)))
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1741)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1602)+80)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v1602)+36)) = v1742
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1735)))
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1745)))
	*(*int32)(unsafe.Add(mBase, uint32(v1602)+4)) = v1746
	v1748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1742))))
	*(*int32)(unsafe.Add(mBase, uint32(v1602)+48)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1602)+24)) = uint8(v1748)
	goto L276
L287:
	;
	goto L271
L288:
	;
	v1821 = int32(v1817)
	m.G0 = v30
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+4))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1821)))
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1824)))
	if v30+int32(72) == v1827 {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	m.ExcPending = 1
	goto L297
L290:
	;
	if v1831 != 0 {
		goto L294
	} else {
		goto L295
	}
L291:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1824)+4))
	v1831 = v1829
	goto L293
L292:
	;
	v1831 = int32(0)
	goto L293
L293:
	;
	goto L290
L294:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v30)+284))
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v30)+280))
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v30)+276))
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v30)+272))
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v30)+268))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v30)+264))
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v30)+260))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v30)+256))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v30)+252))
	v39 = v1836
	v44 = v1823
	v45 = v1831
	v46 = v1840
	v47 = v1839
	v48 = v1838
	v49 = v1837
	v50 = v1835
	v51 = v1834
	v52 = v1833
	v53 = v1832
	v62 = v83
	v63 = v84
	goto L1
L295:
	;
	goto L296
L296:
	;
	F___wasm_longjmp(m, v1824, v1823)
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	return int32(0)
L298:
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
	F_errfinish(m, int32(_a_F_config_enum_lookup_by_value_1), int32(3036), int32(_a_F_config_enum_lookup_by_value_2))
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
