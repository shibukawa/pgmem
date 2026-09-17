package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatCacheRemoveCTup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v10)
	F_CatCacheRemoveCList(m, l0, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+37)))
	if v19 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	F_pfree(m, l1)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v22 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v33 = int32(0)
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v44 = v33 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(48)+v44)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v39<<(uint(int32(4))%32)+v46*int32(100))+2)))
	if v50 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(8)+v44)))
	F_pfree(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v58 = v33 + int32(1)
	if v58 != v22 {
		v33 = v58
		goto L9
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	goto L10
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v71 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v70 - v71
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_CatCacheRemoveCTup[0]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v76 - v71
	return
}
func F_SearchCatCache2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = int32(0)
	v7 = F_SearchCatCacheInternal(m, l0, int32(2), l1, l2, v5, v5)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_SearchCatCacheList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v160 int32
	_ = v160
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v321 int32
	_ = v321
	var v335 int32
	_ = v335
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v510 int32
	_ = v510
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v564 int32
	_ = v564
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v636 int32
	_ = v636
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
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
	var v733 int32
	_ = v733
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v803 int32
	_ = v803
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v931 int32
	_ = v931
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v956 int32
	_ = v956
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v985 int32
	_ = v985
	var v992 int32
	_ = v992
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1125 int32
	_ = v1125
	var v1141 int32
	_ = v1141
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1236 int32
	_ = v1236
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1283 int32
	_ = v1283
	var v1291 int32
	_ = v1291
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1417 int32
	_ = v1417
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1454 int32
	_ = v1454
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1569 int32
	_ = v1569
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1759 int32
	_ = v1759
	var v1773 int32
	_ = v1773
	var v1788 int32
	_ = v1788
	var v1821 int32
	_ = v1821
	var v1822 int64
	_ = v1822
	var v1826 int32
	_ = v1826
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
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
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	v6 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(544)
	m.G0 = v35
	v44 = l0
	v45 = l1
	v46 = l2
	v47 = l3
	v48 = l4
	v49 = v35
	v50 = v6
	v52 = v6
	v53 = v6
	v54 = v6
	v55 = v6
	v56 = v6
	v57 = v6
	v58 = v6
	v59 = v6
	v60 = v6
	v61 = v6
	v63 = int32(-1)
	v69 = l0 + int32(104)
	v70 = v35 + int32(405)
	v71 = l0 + int32(32)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v63 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	goto L3
L6:
	;
	v1821 = int32(m.ExcTag)
	v1822 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1821 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L6
	} else {
		goto L184
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v1712
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v1711
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v1708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v1709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v1704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v1706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v1705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v1703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v1707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v1710
	v1738 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[0]))
	F_ResourceOwnerRemember(m, v1738, v1713, int32(_a_F_SearchCatCacheList_0))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L6
	} else {
		goto L183
	}
L9:
	;
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	if v510 != v1658 {
		goto L176
	} else {
		goto L177
	}
L10:
	;
	v79 = v44 + int32(8)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v80 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v701 = v50
	v703 = v52
	v704 = v53
	v705 = v54
	v706 = v55
	v707 = v56
	v708 = v57
	v709 = v58
	v710 = v59
	goto L12
L12:
	;
	if v701 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	F_CatalogCacheInitializeCache(m, v44)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v95 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+428)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v49)+424)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v49)+420)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v49)+416)) = v46
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	if v100 == v95 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v416 = int32(0)
	switch v45 - int32(1) {
	case 0:
		v470 = v416
		goto L46
	case 1:
		v452 = v416
		goto L47
	case 2:
		v435 = v416
		goto L48
	case 3:
		goto L49
	default:
		goto L7
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[1]))
	v116 = F_MemoryContextAllocZero(m, v114, int32(128))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	if v121 <= v122<<(uint(int32(1))%32) {
		goto L17
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+76)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v116
	goto L17
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	v138 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	if v138 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v44)+84))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v44)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v140
	F_errmsg_internal(m, int32(_a_F_SearchCatCacheList_1), v49+int32(16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[1]))
	v194 = F_MemoryContextAllocZero(m, v191, v179<<(uint(int32(4))%32))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	F_errfinish(m, int32(_a_F_SearchCatCacheList_2), int32(1030), int32(_a_F_SearchCatCacheList_3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v197 = v179 << (uint(int32(1)) % 32)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	if int32(0) < v198 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v223 = v198
	v224 = int32(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	F_pfree(m, v369)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L45
	}
L33:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v239 = v236 + v224<<(uint(int32(3))%32)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	v241 = int32(0)
	if base.B2i32(v240 == v241)|base.B2i32(v240 == v239) == v241 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v253 = v240
	goto L38
L36:
	;
	v321 = v223
	goto L37
L37:
	;
	v335 = v224 + int32(1)
	if v335 < v321 {
		v223 = v321
		v224 = v335
		goto L33
	} else {
		goto L44
	}
L38:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v253-int32(4))))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v285
	v290 = v194 + v281&(v197-int32(1))<<(uint(int32(3))%32)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	if v291 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v321 = v301
	goto L37
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v290
	v295 = v290
	goto L42
L41:
	;
	v295 = v291
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v290)+4)) = v253
	if v283 != v239 {
		v253 = v283
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	goto L34
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v44)+76)) = v197
	goto L17
L46:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	v483 = m.T0[v472].(func(*base.Module, int32) int32)(m, v46)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L6
	} else {
		goto L53
	}
L47:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	v465 = m.T0[v454].(func(*base.Module, int32) int32)(m, v47)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L6
	} else {
		goto L52
	}
L48:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	v447 = m.T0[v436].(func(*base.Module, int32) int32)(m, v48)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	v431 = m.T0[v419].(func(*base.Module, int32) int32)(m, int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	v435 = base.I32_rotl(v431, int32(24))
	goto L48
L51:
	;
	v452 = base.I32_rotl(v447, int32(16)) ^ v435
	goto L47
L52:
	;
	v470 = base.I32_rotl(v465, int32(8)) ^ v452
	goto L46
L53:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v486 = v470 ^ v483
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v493 = v485 + v486&(v487-int32(1))<<(uint(int32(3))%32)
	v495 = v493 + int32(4)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	v497 = int32(0)
	if base.B2i32(v496 == v497)|base.B2i32(v496 == v493) == v497 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v510 = v496
	goto L57
L55:
	;
	goto L56
L56:
	;
	v670 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+412)) = v670
	v673 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+404)) = uint16(v673)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+400)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v49)+396)) = v44
	v677 = int32(_a_F_SearchCatCacheList_4)
	v678 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[2])) = v49 + int32(396)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+408)) = v678
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[3]))
	v687 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[4]))
	goto L69
L57:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+28)))
	if v535 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L56
L59:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	if v636 != v493 {
		v510 = v636
		goto L57
	} else {
		goto L68
	}
L60:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v510-int32(4))))
	if v538 != v486 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v540 = int32(*(*int16)(unsafe.Add(mBase, uint32(v510)+30)))
	if v45 != v540 {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v564 = int32(0)
	goto L63
L63:
	;
	v578 = v564 << (uint(int32(2)) % 32)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v510+int32(8)+v578)))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v578+v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(416)+v578)))
	v597 = m.T0[v582].(func(*base.Module, int32, int32) int32)(m, v580, v596)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L6
	} else {
		goto L65
	}
L64:
	;
	goto L9
L65:
	;
	if v597 == int32(0) {
		goto L59
	} else {
		goto L66
	}
L66:
	;
	v602 = v564 + int32(1)
	if v45 != v602 {
		v564 = v602
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	goto L58
L69:
	;
	v689 = v49 + int32(240)
	*(*int32)(unsafe.Add(mBase, uint32(v689)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v689))) = v49 + int32(44)
	goto L72
L70:
	;
	v701 = v670
	v703 = v493
	v704 = v70
	v705 = v495
	v706 = v678
	v707 = v486
	v708 = v687
	v709 = v685
	v710 = v79
	goto L12
L72:
	;
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[3])) = v49 + int32(240)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v44)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v745 = F_table_open(m, v733, int32(1))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L6
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[4])) = v708
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[3])) = v709
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[2])) = v706
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v49)+412))
	if v1543 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L76:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v44)+64))
	v749 = v747 * int32(48)
	if v749 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	base.MemoryCopy(m, v49+int32(48), v69, v749)
	goto L79
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+236)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+188)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v49)+140)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v49)+92)) = v46
	v774 = v60
	v775 = v61
	goto L80
L80:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v49)+412))
	if v790 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v1195
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v1196
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	F_relation_close(m, v745, int32(1))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L6
	} else {
		goto L132
	}
L82:
	;
	v874 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+412)) = v874
	*(*uint8)(unsafe.Add(mBase, uint32(v704))) = uint8(v874)
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v44)+92))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	switch v880 - int32(1) {
	case 0, 1:
		v890 = v874
		goto L88
	default:
		goto L89
	case 7, 9, 10, 20:
		goto L90
	case 33:
		goto L91
	}
L83:
	;
	v793 = int32(0)
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v790)+4))
	if v794 <= v793 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v803 = v793
	goto L85
L85:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v790)+12))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v829+v803<<(uint(int32(2))%32))))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v833)+32))
	v835 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v833)+32)) = v834 - v835
	v839 = v803 + v835
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v790)+4))
	if v839 < v840 {
		v803 = v839
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L82
L87:
	;
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v775
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v904 = F_systable_beginscan(m, v745, v879, v890, int32(0), v45, v49+int32(48))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L6
	} else {
		goto L94
	}
L89:
	;
	v890 = int32(1)
	goto L88
L90:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheList[5])))
	if v886 != int32(1) {
		v890 = v874
		goto L88
	} else {
		goto L93
	}
L91:
	;
	v884 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SearchCatCacheList[6])))
	if v884 != 0 {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	v890 = v874
	goto L88
L93:
	;
	goto L89
L94:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v904)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v775
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v917 = F_systable_getnext(m, v904)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L6
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v1195
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v1196
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	F_systable_endscan(m, v904)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L6
	} else {
		goto L130
	}
L96:
	;
	if v917 == int32(0) {
		v1195 = v774
		v1196 = v775
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704))))
	if v921&int32(1) != 0 {
		v1195 = v774
		v1196 = v775
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v931 = v917
	v940 = v774
	v941 = v775
	goto L99
L99:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v44)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v940
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v941
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v967 = F_CatalogCacheComputeTupleHashValue(m, v44, v956, v931)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L6
	} else {
		goto L101
	}
L100:
	;
	v1195 = v1125
	v1196 = v1170
	goto L95
L101:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v973 = (v970 - int32(1)) & v967
	v976 = v969 + v973<<(uint(int32(3))%32)
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)+4))
	v978 = int32(0)
	if base.B2i32(v977 == v978)|base.B2i32(v977 == v976) == v978 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v941
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v1125
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v49)+412))
	v1153 = F_lappend(m, v1152, v1141)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L6
	} else {
		goto L126
	}
L103:
	;
	v1125 = v940
	v1141 = v992 - int32(24)
	goto L102
L104:
	;
	v985 = v931 + int32(4)
	v992 = v977
	goto L107
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v940
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v941
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v1103 = F_CatalogCacheCreateEntry(m, v44, v931, int32(0), v967, v973)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L6
	} else {
		goto L122
	}
L107:
	;
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+12)))
	if v1018 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L106
L109:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v992)+4))
	if v1058 != v976 {
		v992 = v1058
		goto L107
	} else {
		goto L121
	}
L110:
	;
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+13)))
	if v1019 != 0 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v992-int32(20))))
	if v1022 != v967 {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v940
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v941
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v1035 = v992 + int32(20)
	v1036 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1035)+2)))
	v1037 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1035))))
	v1038 = int32(16)
	v1041 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+2)))
	v1042 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985))))
	if v1036|v1037<<(uint(v1038)%32) == v1041|v1042<<(uint(v1038)%32) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	if v1052 == int32(0) {
		goto L109
	} else {
		goto L119
	}
L114:
	;
	goto L113
L115:
	;
	v1048 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1035)+4)))
	v1049 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v985)+4)))
	if v1048 == v1049 {
		v1052 = int32(1)
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v1052 = int32(0)
	goto L114
L118:
	;
	goto L117
L119:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v992)+36))
	if v1055 == int32(0) {
		goto L103
	} else {
		goto L120
	}
L120:
	;
	goto L109
L121:
	;
	goto L108
L122:
	;
	if v1103 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v1125 = v1103
	v1141 = v1103
	goto L102
L124:
	;
	goto L125
L125:
	;
	v1105 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v704))) = uint8(v1105)
	v1195 = v1103
	v1196 = v941
	goto L95
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+412)) = v1153
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1141)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1141)+32)) = v1156 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v1125
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v941
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v1170 = F_systable_getnext(m, v904)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	if v1170 == int32(0) {
		v1195 = v1125
		v1196 = v1170
		goto L95
	} else {
		goto L128
	}
L128:
	;
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704))))
	if v1174&int32(1) == int32(0) {
		v931 = v1170
		v940 = v1125
		v941 = v1170
		goto L99
	} else {
		goto L129
	}
L129:
	;
	goto L100
L130:
	;
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704))))
	if v1223 != 0 {
		v774 = v1195
		v775 = v1196
		goto L80
	} else {
		goto L131
	}
L131:
	;
	goto L81
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v1196
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v1195
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v1248 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[0]))
	F_ResourceOwnerEnlarge(m, v1248)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	v1252 = int32(_a_F_SearchCatCacheList_5)
	v1253 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[7]))
	v1256 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[7])) = v1256
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v49)+412))
	if v1258 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+4))
	v1260 = v1259
	goto L136
L135:
	;
	v1260 = int32(0)
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v1195
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v1196
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v1275 = F_palloc(m, v1260<<(uint(int32(2))%32)+int32(48))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	if int32(0) < v45 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	v1291 = int32(0)
	goto L141
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[7])) = v1253
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[3])) = v709
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[4])) = v708
	*(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[2])) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v1275))) = int32(1383485699)
	*(*uint16)(unsafe.Add(mBase, uint32(v1275)+38)) = uint16(v45)
	v1417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1275)+37)) = uint8(base.B2i32(v906 != v1417))
	*(*uint8)(unsafe.Add(mBase, uint32(v1275)+36)) = uint8(v1417)
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+32)) = v1417
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+40)) = v1260
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+4)) = v707
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v49)+412))
	if v1426 == v1417 {
		goto L150
	} else {
		goto L151
	}
L141:
	;
	v1318 = v1291 << (uint(int32(2)) % 32)
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1318+(v49+int32(416)))))
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1283)))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1318+(v44+int32(48)))))
	v1333 = v1283 + v1323<<(uint(int32(4))%32) + v1328*int32(100) - int32(80)
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+68))
	if v1334 != int32(19) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	goto L140
L143:
	;
	v1354 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1333)+72)))
	v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1333)+82)))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v1195
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v1196
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v1367 = F_datumCopy(m, v1353, v1355, v1354)
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L6
	} else {
		goto L148
	}
L144:
	;
	v1353 = v1322
	goto L143
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v1195
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v1196
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	v1348 = v49 + int32(440)
	v1350 = F_strncpy(m, v1348, v1322, int32(64))
	mBase = m.M
	v1351 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1350)+63)) = uint8(v1351)
	goto L147
L147:
	;
	v1353 = v1348
	goto L143
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1318+(v1275+int32(16))))) = v1367
	v1371 = v1291 + int32(1)
	if v1371 != v45 {
		v1291 = v1371
		goto L141
	} else {
		goto L149
	}
L149:
	;
	goto L142
L150:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	if v1518 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L151:
	;
	v1429 = int32(0)
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+4))
	if v1430 <= v1429 {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v1454 = v1429
	goto L153
L153:
	;
	v1468 = v1454 << (uint(int32(2)) % 32)
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+12))
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1470+v1468)))
	*(*int32)(unsafe.Add(mBase, uint32(v1275+int32(48)+v1468))) = v1472
	*(*int32)(unsafe.Add(mBase, uint32(v1472)+60)) = v1275
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1472)+32)) = v1475 - int32(1)
	v1479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1472)+36)))
	if v1479 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L150
L155:
	;
	v1480 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1275)+36)) = uint8(v1480)
	goto L157
L156:
	;
	goto L157
L157:
	;
	v1483 = v1454 + int32(1)
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+4))
	if v1483 < v1484 {
		v1454 = v1483
		goto L153
	} else {
		goto L158
	}
L158:
	;
	goto L154
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v703))) = v703
	v1522 = v703
	goto L161
L160:
	;
	v1522 = v1518
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+8)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+12)) = v1522
	v1526 = v1275 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v1522))) = v1526
	*(*int32)(unsafe.Add(mBase, uint32(v705))) = v1526
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	v1530 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+72)) = v1529 + v1530
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+32)) = v1533 + v1530
	v1703 = v703
	v1704 = v704
	v1705 = v705
	v1706 = v706
	v1707 = v707
	v1708 = v708
	v1709 = v709
	v1710 = v710
	v1711 = v1195
	v1712 = v1196
	v1713 = v1275
	goto L8
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	F_pg_re_throw(m)
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L6
	} else {
		goto L175
	}
L163:
	;
	v1546 = int32(0)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1543)+4))
	if v1547 <= v1546 {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v1569 = v1546
	goto L165
L165:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1543)+12))
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1582+v1569<<(uint(int32(2))%32))))
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1586)+32))
	v1588 = int32(1)
	v1589 = v1587 - v1588
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+32)) = v1589
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1586)+36)))
	if base.B2i32(v1591 != v1588)|v1589 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	goto L162
L167:
	;
	v1611 = v1569 + int32(1)
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1543)+4))
	if v1611 < v1612 {
		v1569 = v1611
		goto L165
	} else {
		goto L174
	}
L168:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1586)+60))
	if v1595 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+32))
	if v1596 != 0 {
		goto L167
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v709
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v704
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v710
	F_CatCacheRemoveCTup(m, v44, v1586)
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L6
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	goto L167
L174:
	;
	goto L166
L175:
	;
	goto L3
L176:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1660)+4)) = v1661
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	*(*int32)(unsafe.Add(mBase, uint32(v1661))) = v1663
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	if v1665 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	v1686 = *(*int32)(unsafe.Add(mBase, _c_F_SearchCatCacheList[0]))
	F_ResourceOwnerEnlarge(m, v1686)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L6
	} else {
		goto L182
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v493))) = v493
	v1669 = v493
	goto L181
L180:
	;
	v1669 = v1665
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v510))) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v510)+4)) = v1669
	*(*int32)(unsafe.Add(mBase, uint32(v1669))) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v493)+4)) = v510
	goto L178
L182:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v510)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v510)+24)) = v1691 + int32(1)
	v1703 = v493
	v1704 = v53
	v1705 = v495
	v1706 = v55
	v1707 = v486
	v1708 = v57
	v1709 = v58
	v1710 = v79
	v1711 = v60
	v1712 = v61
	v1713 = v510 - int32(8)
	goto L8
L183:
	;
	m.G0 = v49 + int32(544)
	return v1713
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v45
	F_errmsg_internal(m, int32(_a_F_SearchCatCacheList_6), v49)
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L6
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+508)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v49)+504)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v49)+512)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+516)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v49)+520)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+524)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v49)+528)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v49)+532)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v49)+536)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v49)+540)) = v79
	F_errfinish(m, int32(_a_F_SearchCatCacheList_2), int32(373), int32(_a_F_SearchCatCacheList_7))
	mBase = m.M
	v1788 = m.ExcPending
	if v1788 != 0 {
		goto L6
	} else {
		goto L186
	}
L186:
	;
	goto L5
L187:
	;
	v1826 = int32(v1822)
	m.G0 = v49
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1826)+4))
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1826)))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1829)))
	if v49+int32(44) == v1832 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	m.ExcPending = 1
	goto L196
L189:
	;
	if v1836 != 0 {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1829)+4))
	v1836 = v1834
	goto L192
L191:
	;
	v1836 = int32(0)
	goto L192
L192:
	;
	goto L189
L193:
	;
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v49)+540))
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v49)+536))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v49)+532))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v49)+528))
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v49)+524))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v49)+520))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v49)+516))
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v49)+512))
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v49)+508))
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v49)+504))
	v50 = v1828
	v52 = v1839
	v53 = v1842
	v54 = v1840
	v55 = v1841
	v56 = v1838
	v57 = v1844
	v58 = v1843
	v59 = v1837
	v60 = v1845
	v61 = v1846
	v63 = v1836
	goto L1
L194:
	;
	goto L195
L195:
	;
	F___wasm_longjmp(m, v1829, v1828)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	return int32(0)
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
