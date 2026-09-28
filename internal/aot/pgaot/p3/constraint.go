package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateConstraintEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 int32, l28 int32, l29 int32, l30 int32, l31 int32, l32 int32) int32 {
	mBase := m.M
	_ = mBase
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v138 int64
	_ = v138
	var v141 int32
	_ = v141
	var v148 int64
	_ = v148
	var v151 int32
	_ = v151
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v213 int32
	_ = v213
	var v248 int32
	_ = v248
	var v262 int32
	_ = v262
	var v265 int64
	_ = v265
	var v270 int32
	_ = v270
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int64
	_ = v436
	var v439 int32
	_ = v439
	var v446 int64
	_ = v446
	var v449 int32
	_ = v449
	var v456 int64
	_ = v456
	var v459 int32
	_ = v459
	var v466 int64
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v521 int32
	_ = v521
	var v556 int32
	_ = v556
	var v570 int32
	_ = v570
	var v573 int64
	_ = v573
	var v578 int32
	_ = v578
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v638 int32
	_ = v638
	var v674 int32
	_ = v674
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v690 int64
	_ = v690
	var v693 int32
	_ = v693
	var v700 int64
	_ = v700
	var v703 int32
	_ = v703
	var v710 int64
	_ = v710
	var v713 int32
	_ = v713
	var v720 int64
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v775 int32
	_ = v775
	var v810 int32
	_ = v810
	var v827 int64
	_ = v827
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v892 int32
	_ = v892
	var v928 int32
	_ = v928
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v944 int64
	_ = v944
	var v947 int32
	_ = v947
	var v954 int64
	_ = v954
	var v957 int32
	_ = v957
	var v964 int64
	_ = v964
	var v967 int32
	_ = v967
	var v974 int64
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v1029 int32
	_ = v1029
	var v1064 int32
	_ = v1064
	var v1081 int64
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1182 int32
	_ = v1182
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1198 int64
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1208 int64
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1218 int64
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1228 int64
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1283 int32
	_ = v1283
	var v1318 int32
	_ = v1318
	var v1335 int64
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1402 int32
	_ = v1402
	var v1438 int32
	_ = v1438
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1454 int64
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1464 int64
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1474 int64
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1484 int64
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1493 int32
	_ = v1493
	var v1539 int32
	_ = v1539
	var v1574 int32
	_ = v1574
	var v1588 int32
	_ = v1588
	var v1591 int64
	_ = v1591
	var v1596 int32
	_ = v1596
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1708 int32
	_ = v1708
	var v1743 int32
	_ = v1743
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1760 int64
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1770 int64
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1780 int64
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1790 int64
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1845 int32
	_ = v1845
	var v1867 int32
	_ = v1867
	var v1897 int64
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v2001 int64
	_ = v2001
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2079 int32
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2087 int32
	_ = v2087
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2101 int32
	_ = v2101
	var v2103 int32
	_ = v2103
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2127 int32
	_ = v2127
	var v2179 int32
	_ = v2179
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2196 int32
	_ = v2196
	var v2251 int32
	_ = v2251
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2266 int32
	_ = v2266
	var v2318 int32
	_ = v2318
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2335 int32
	_ = v2335
	var v2382 int32
	_ = v2382
	var v2397 int32
	_ = v2397
	var v2400 int32
	_ = v2400
	var v2405 int32
	_ = v2405
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2476 int32
	_ = v2476
	var v2478 int32
	_ = v2478
	var v2527 int32
	_ = v2527
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2542 int32
	_ = v2542
	v34 = int32(0)
	v47 = m.G0
	v49 = v47 - int32(352)
	m.G0 = v49
	v53 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v60 = F_strncpy(m, v49+int32(32), l0, int32(64))
	mBase = m.M
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+63)) = uint8(v61)
	goto L3
L3:
	;
	if int32(0) < l10 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v65 = int32(3)
	v66 = l10 & v65
	v69 = F_palloc(m, l10<<(uint(v65)%32))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v363 = v34
	goto L6
L6:
	;
	if l19 <= int32(0) {
		v1686 = v34
		v1687 = v34
		v1688 = v34
		v1690 = v34
		v1692 = v34
		goto L20
	} else {
		goto L21
	}
L7:
	;
	v71 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l10) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v319 = F_construct_array_builtin(m, v69, l10, int32(21))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v76 = v71
	v112 = v34
	goto L12
L10:
	;
	v167 = v71
	goto L11
L11:
	;
	v213 = v167
	v248 = v34
	goto L16
L12:
	;
	v122 = int32(3)
	v125 = int32(1)
	v128 = int64(*(*int16)(unsafe.Add(mBase, uint32(l9+v76<<(uint(v125)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v76<<(uint(v122)%32)))) = v128
	v131 = v76 | v125
	v138 = int64(*(*int16)(unsafe.Add(mBase, uint32(l9+v131<<(uint(v125)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v131<<(uint(v122)%32)))) = v138
	v141 = v76 | int32(2)
	v148 = int64(*(*int16)(unsafe.Add(mBase, uint32(l9+v141<<(uint(v125)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v141<<(uint(v122)%32)))) = v148
	v151 = v76 | v122
	v158 = int64(*(*int16)(unsafe.Add(mBase, uint32(l9+v151<<(uint(v125)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v151<<(uint(v122)%32)))) = v158
	v160 = int32(4)
	v161 = v76 + v160
	v163 = v112 + v160
	if v163 != l10&int32(2147483644) {
		v76 = v161
		v112 = v163
		goto L12
	} else {
		goto L14
	}
L13:
	;
	if v66 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v167 = v161
	goto L11
L16:
	;
	v262 = int32(1)
	v265 = int64(*(*int16)(unsafe.Add(mBase, uint32(l9+v213<<(uint(v262)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v213<<(uint(int32(3))%32)))) = v265
	v270 = v248 + v262
	if v270 != v66 {
		v213 = v213 + v262
		v248 = v270
		goto L16
	} else {
		goto L18
	}
L17:
	;
	goto L8
L18:
	;
	goto L17
L19:
	;
	v363 = v319
	goto L6
L20:
	;
	v1693 = int32(0)
	if l25 != 0 {
		goto L87
	} else {
		goto L88
	}
L21:
	;
	v370 = l19 & int32(3)
	if l23 < l19 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v373 = l19
	goto L24
L23:
	;
	v373 = l23
	goto L24
L24:
	;
	v376 = F_palloc(m, v373<<(uint(int32(3))%32))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v378 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l19) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v627 = l19 & int32(3)
	v630 = F_construct_array_builtin(m, v376, l19, int32(21))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L37
	}
L27:
	;
	v384 = v378
	v420 = int32(0)
	goto L30
L28:
	;
	v475 = v378
	goto L29
L29:
	;
	v521 = v475
	v556 = int32(0)
	goto L34
L30:
	;
	v430 = int32(3)
	v433 = int32(1)
	v436 = int64(*(*int16)(unsafe.Add(mBase, uint32(l15+v384<<(uint(v433)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v384<<(uint(v430)%32)))) = v436
	v439 = v384 | v433
	v446 = int64(*(*int16)(unsafe.Add(mBase, uint32(l15+v439<<(uint(v433)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v439<<(uint(v430)%32)))) = v446
	v449 = v384 | int32(2)
	v456 = int64(*(*int16)(unsafe.Add(mBase, uint32(l15+v449<<(uint(v433)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v449<<(uint(v430)%32)))) = v456
	v459 = v384 | v430
	v466 = int64(*(*int16)(unsafe.Add(mBase, uint32(l15+v459<<(uint(v433)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v459<<(uint(v430)%32)))) = v466
	v468 = int32(4)
	v469 = v384 + v468
	v471 = v420 + v468
	if v471 != l19&int32(2147483644) {
		v384 = v469
		v420 = v471
		goto L30
	} else {
		goto L32
	}
L31:
	;
	if v370 == int32(0) {
		goto L26
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v475 = v469
	goto L29
L34:
	;
	v570 = int32(1)
	v573 = int64(*(*int16)(unsafe.Add(mBase, uint32(l15+v521<<(uint(v570)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v521<<(uint(int32(3))%32)))) = v573
	v578 = v556 + v570
	if v578 != v370 {
		v521 = v521 + v570
		v556 = v578
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L26
L36:
	;
	goto L35
L37:
	;
	v632 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l19) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v881 = l19 & int32(3)
	v884 = F_construct_array_builtin(m, v376, l19, int32(26))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L1
	} else {
		goto L49
	}
L39:
	;
	v638 = v632
	v674 = int32(0)
	goto L42
L40:
	;
	v729 = v632
	goto L41
L41:
	;
	v775 = v729
	v810 = int32(0)
	goto L46
L42:
	;
	v684 = int32(3)
	v687 = int32(2)
	v690 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l16+v638<<(uint(v687)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v638<<(uint(v684)%32)))) = v690
	v693 = v638 | int32(1)
	v700 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l16+v693<<(uint(v687)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v693<<(uint(v684)%32)))) = v700
	v703 = v638 | v687
	v710 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l16+v703<<(uint(v687)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v703<<(uint(v684)%32)))) = v710
	v713 = v638 | v684
	v720 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l16+v713<<(uint(v687)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v713<<(uint(v684)%32)))) = v720
	v722 = int32(4)
	v723 = v638 + v722
	v725 = v674 + v722
	if v725 != l19&int32(2147483644) {
		v638 = v723
		v674 = v725
		goto L42
	} else {
		goto L44
	}
L43:
	;
	if v627 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v729 = v723
	goto L41
L46:
	;
	v827 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l16+v775<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v775<<(uint(int32(3))%32)))) = v827
	v829 = int32(1)
	v832 = v810 + v829
	if v832 != v627 {
		v775 = v775 + v829
		v810 = v832
		goto L46
	} else {
		goto L48
	}
L47:
	;
	goto L38
L48:
	;
	goto L47
L49:
	;
	v886 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l19) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v1135 = l19 & int32(3)
	v1138 = F_construct_array_builtin(m, v376, l19, int32(26))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L1
	} else {
		goto L61
	}
L51:
	;
	v892 = v886
	v928 = int32(0)
	goto L54
L52:
	;
	v983 = v886
	goto L53
L53:
	;
	v1029 = v983
	v1064 = int32(0)
	goto L58
L54:
	;
	v938 = int32(3)
	v941 = int32(2)
	v944 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l17+v892<<(uint(v941)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v892<<(uint(v938)%32)))) = v944
	v947 = v892 | int32(1)
	v954 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l17+v947<<(uint(v941)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v947<<(uint(v938)%32)))) = v954
	v957 = v892 | v941
	v964 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l17+v957<<(uint(v941)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v957<<(uint(v938)%32)))) = v964
	v967 = v892 | v938
	v974 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l17+v967<<(uint(v941)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v967<<(uint(v938)%32)))) = v974
	v976 = int32(4)
	v977 = v892 + v976
	v979 = v928 + v976
	if v979 != l19&int32(2147483644) {
		v892 = v977
		v928 = v979
		goto L54
	} else {
		goto L56
	}
L55:
	;
	if v881 == int32(0) {
		goto L50
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v983 = v977
	goto L53
L58:
	;
	v1081 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l17+v1029<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1029<<(uint(int32(3))%32)))) = v1081
	v1083 = int32(1)
	v1086 = v1064 + v1083
	if v1086 != v881 {
		v1029 = v1029 + v1083
		v1064 = v1086
		goto L58
	} else {
		goto L60
	}
L59:
	;
	goto L50
L60:
	;
	goto L59
L61:
	;
	v1140 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l19) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v1389 = F_construct_array_builtin(m, v376, l19, int32(26))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L73
	}
L63:
	;
	v1146 = v1140
	v1182 = int32(0)
	goto L66
L64:
	;
	v1237 = v1140
	goto L65
L65:
	;
	v1283 = v1237
	v1318 = int32(0)
	goto L70
L66:
	;
	v1192 = int32(3)
	v1195 = int32(2)
	v1198 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l18+v1146<<(uint(v1195)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1146<<(uint(v1192)%32)))) = v1198
	v1201 = v1146 | int32(1)
	v1208 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l18+v1201<<(uint(v1195)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1201<<(uint(v1192)%32)))) = v1208
	v1211 = v1146 | v1195
	v1218 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l18+v1211<<(uint(v1195)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1211<<(uint(v1192)%32)))) = v1218
	v1221 = v1146 | v1192
	v1228 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l18+v1221<<(uint(v1195)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1221<<(uint(v1192)%32)))) = v1228
	v1230 = int32(4)
	v1231 = v1146 + v1230
	v1233 = v1182 + v1230
	if v1233 != l19&int32(2147483644) {
		v1146 = v1231
		v1182 = v1233
		goto L66
	} else {
		goto L68
	}
L67:
	;
	if v1135 == int32(0) {
		goto L62
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	v1237 = v1231
	goto L65
L70:
	;
	v1335 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l18+v1283<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1283<<(uint(int32(3))%32)))) = v1335
	v1337 = int32(1)
	v1340 = v1318 + v1337
	if v1340 != v1135 {
		v1283 = v1283 + v1337
		v1318 = v1340
		goto L70
	} else {
		goto L72
	}
L71:
	;
	goto L62
L72:
	;
	goto L71
L73:
	;
	if l23 <= int32(0) {
		v1686 = v1389
		v1687 = v884
		v1688 = v1138
		v1690 = v34
		v1692 = v630
		goto L20
	} else {
		goto L74
	}
L74:
	;
	v1394 = l23 & int32(3)
	v1395 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l23) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v1645 = F_construct_array_builtin(m, v376, l23, int32(21))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L1
	} else {
		goto L86
	}
L76:
	;
	v1402 = v1395
	v1438 = int32(0)
	goto L79
L77:
	;
	v1493 = v1395
	goto L78
L78:
	;
	v1539 = v1493
	v1574 = v1395
	goto L83
L79:
	;
	v1448 = int32(3)
	v1451 = int32(1)
	v1454 = int64(*(*int16)(unsafe.Add(mBase, uint32(l22+v1402<<(uint(v1451)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1402<<(uint(v1448)%32)))) = v1454
	v1457 = v1402 | v1451
	v1464 = int64(*(*int16)(unsafe.Add(mBase, uint32(l22+v1457<<(uint(v1451)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1457<<(uint(v1448)%32)))) = v1464
	v1467 = v1402 | int32(2)
	v1474 = int64(*(*int16)(unsafe.Add(mBase, uint32(l22+v1467<<(uint(v1451)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1467<<(uint(v1448)%32)))) = v1474
	v1477 = v1402 | v1448
	v1484 = int64(*(*int16)(unsafe.Add(mBase, uint32(l22+v1477<<(uint(v1451)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1477<<(uint(v1448)%32)))) = v1484
	v1486 = int32(4)
	v1487 = v1402 + v1486
	v1489 = v1438 + v1486
	if v1489 != l23&int32(2147483644) {
		v1402 = v1487
		v1438 = v1489
		goto L79
	} else {
		goto L81
	}
L80:
	;
	if v1394 == int32(0) {
		goto L75
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v1493 = v1487
	goto L78
L83:
	;
	v1588 = int32(1)
	v1591 = int64(*(*int16)(unsafe.Add(mBase, uint32(l22+v1539<<(uint(v1588)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v376+v1539<<(uint(int32(3))%32)))) = v1591
	v1596 = v1574 + v1588
	if v1596 != v1394 {
		v1539 = v1539 + v1588
		v1574 = v1596
		goto L83
	} else {
		goto L85
	}
L84:
	;
	goto L75
L85:
	;
	goto L84
L86:
	;
	v1686 = v1389
	v1687 = v884
	v1688 = v1138
	v1690 = v1645
	v1692 = v630
	goto L20
L87:
	;
	v1696 = F_palloc(m, l10<<(uint(int32(3))%32))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	v1953 = v1693
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+344)) = int32(0)
	v2001 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+336)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+328)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+320)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+256)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+264)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+272)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+280)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+288)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+296)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+304)) = v2001
	*(*int64)(unsafe.Add(mBase, uint32(v49)+312)) = v2001
	v2025 = F_GetNewOidWithIndex(m, v53, int32(2667), int32(1))
	mBase = m.M
	v2026 = m.ExcPending
	if v2026 != 0 {
		goto L1
	} else {
		goto L104
	}
L90:
	;
	if l10 <= int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v1951 = F_construct_array_builtin(m, v1696, l10, int32(26))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L1
	} else {
		goto L103
	}
L92:
	;
	v1701 = l10 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(l10) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v1708 = v1693
	v1743 = int32(0)
	goto L96
L94:
	;
	v1799 = v1693
	goto L95
L95:
	;
	v1845 = v1799
	v1867 = int32(0)
	goto L100
L96:
	;
	v1754 = int32(3)
	v1757 = int32(2)
	v1760 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l25+v1708<<(uint(v1757)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1696+v1708<<(uint(v1754)%32)))) = v1760
	v1763 = v1708 | int32(1)
	v1770 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l25+v1763<<(uint(v1757)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1696+v1763<<(uint(v1754)%32)))) = v1770
	v1773 = v1708 | v1757
	v1780 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l25+v1773<<(uint(v1757)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1696+v1773<<(uint(v1754)%32)))) = v1780
	v1783 = v1708 | v1754
	v1790 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l25+v1783<<(uint(v1757)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1696+v1783<<(uint(v1754)%32)))) = v1790
	v1792 = int32(4)
	v1793 = v1708 + v1792
	v1795 = v1743 + v1792
	if v1795 != l10&int32(2147483644) {
		v1708 = v1793
		v1743 = v1795
		goto L96
	} else {
		goto L98
	}
L97:
	;
	if v1701 == int32(0) {
		goto L91
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	v1799 = v1793
	goto L95
L100:
	;
	v1897 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l25+v1845<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1696+v1845<<(uint(int32(3))%32)))) = v1897
	v1899 = int32(1)
	v1902 = v1867 + v1899
	if v1902 != v1701 {
		v1845 = v1845 + v1899
		v1867 = v1902
		goto L100
	} else {
		goto L102
	}
L101:
	;
	goto L91
L102:
	;
	goto L101
L103:
	;
	v1953 = v1951
	goto L89
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+248)) = base.I64_extend_i32_u(l31)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+240)) = base.I64_extend_i32_u(l30)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+232)) = base.I64_extend_i32_s(l29)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+224)) = base.I64_extend_i32_u(l28)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+216)) = base.I64_extend_i32_s(l24)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+208)) = base.I64_extend_i32_s(l21)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+200)) = base.I64_extend_i32_s(l20)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+192)) = base.I64_extend_i32_u(l14)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+184)) = base.I64_extend_i32_u(l7)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+176)) = base.I64_extend_i32_u(l13)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+168)) = base.I64_extend_i32_u(l12)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+160)) = base.I64_extend_i32_u(l8)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+152)) = base.I64_extend_i32_u(l6)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+144)) = base.I64_extend_i32_u(l5)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+136)) = base.I64_extend_i32_u(l4)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+128)) = base.I64_extend_i32_u(l3)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+120)) = base.I64_extend_i32_s(l2)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+112)) = base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+104)) = base.I64_extend_i32_u(v49 + int32(32))
	*(*int64)(unsafe.Add(mBase, uint32(v49)+96)) = base.I64_extend_i32_u(v2025)
	if v363 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v1692 != 0 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+256)) = base.I64_extend_i32_u(v363)
	goto L105
L107:
	;
	goto L108
L108:
	;
	v2071 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+340)) = uint8(v2071)
	goto L105
L109:
	;
	if v1687 != 0 {
		goto L114
	} else {
		goto L115
	}
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+264)) = base.I64_extend_i32_u(v1692)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v2075 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+341)) = uint8(v2075)
	goto L109
L113:
	;
	if v1688 != 0 {
		goto L118
	} else {
		goto L119
	}
L114:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+272)) = base.I64_extend_i32_u(v1687)
	goto L113
L115:
	;
	goto L116
L116:
	;
	v2079 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+342)) = uint8(v2079)
	goto L113
L117:
	;
	if v1686 != 0 {
		goto L122
	} else {
		goto L123
	}
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+280)) = base.I64_extend_i32_u(v1688)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v2083 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+343)) = uint8(v2083)
	goto L117
L121:
	;
	if v1690 != 0 {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+288)) = base.I64_extend_i32_u(v1686)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v2087 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+344)) = uint8(v2087)
	goto L121
L125:
	;
	if v1953 != 0 {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+296)) = base.I64_extend_i32_u(v1690)
	goto L125
L127:
	;
	goto L128
L128:
	;
	v2091 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+345)) = uint8(v2091)
	goto L125
L129:
	;
	if l27 != 0 {
		goto L134
	} else {
		goto L135
	}
L130:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+304)) = base.I64_extend_i32_u(v1953)
	goto L129
L131:
	;
	goto L132
L132:
	;
	v2095 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+346)) = uint8(v2095)
	goto L129
L133:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	v2108 = F_heap_form_tuple(m, v2103, v49+int32(96), v49+int32(320))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L1
	} else {
		goto L138
	}
L134:
	;
	v2097 = F_cstring_to_text(m, l27)
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v2101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+347)) = uint8(v2101)
	goto L133
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+312)) = base.I64_extend_i32_u(v2097)
	goto L133
L138:
	;
	F_CatalogTupleInsert(m, v53, v2108)
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v2025
	*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = int32(2606)
	F_relation_close(m, v53, int32(3))
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v2120 = F_new_object_addresses(m)
	mBase = m.M
	v2121 = m.ExcPending
	if v2121 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	if l8 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	if l12 != 0 {
		goto L152
	} else {
		goto L153
	}
L143:
	;
	if int32(0) < l11 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v2127 = int32(0)
	goto L147
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	F_add_exact_object_address(m, v49+int32(8), v2120)
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L1
	} else {
		goto L151
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	v2179 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v2127<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v2179
	F_add_exact_object_address(m, v49+int32(8), v2120)
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L1
	} else {
		goto L149
	}
L148:
	;
	goto L142
L149:
	;
	v2186 = v2127 + int32(1)
	if v2186 != l11 {
		v2127 = v2186
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	goto L142
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1247)
	F_add_exact_object_address(m, v49+int32(8), v2120)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	F_record_object_address_dependencies(m, v49+int32(20), v2120, int32(97))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L1
	} else {
		goto L156
	}
L155:
	;
	goto L154
L156:
	;
	F_free_object_addresses(m, v2120)
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v2259 = F_new_object_addresses(m)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	if l14 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v2382 = int32(0)
	if base.B2i32(l13 == v2382)|base.B2i32(l2 != int32(102)) == v2382 {
		goto L169
	} else {
		goto L170
	}
L160:
	;
	if int32(0) < l19 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v2266 = int32(0)
	goto L164
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	F_add_exact_object_address(m, v49+int32(8), v2259)
	mBase = m.M
	v2335 = m.ExcPending
	if v2335 != 0 {
		goto L1
	} else {
		goto L168
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	v2318 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v2266<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v2318
	F_add_exact_object_address(m, v49+int32(8), v2259)
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L1
	} else {
		goto L166
	}
L165:
	;
	goto L159
L166:
	;
	v2325 = v2266 + int32(1)
	if v2325 != l19 {
		v2266 = v2325
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	goto L159
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	F_add_exact_object_address(m, v49+int32(8), v2259)
	mBase = m.M
	v2397 = m.ExcPending
	if v2397 != 0 {
		goto L1
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	if int32(0) < l19 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	goto L171
L173:
	;
	v2400 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v2400
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(2617)
	v2405 = v2400
	goto L176
L174:
	;
	goto L175
L175:
	;
	v2527 = v49 + int32(20)
	F_record_object_address_dependencies(m, v2527, v2259, int32(110))
	mBase = m.M
	v2530 = m.ExcPending
	if v2530 != 0 {
		goto L1
	} else {
		goto L188
	}
L176:
	;
	v2452 = v2405 << (uint(int32(2)) % 32)
	v2453 = l16 + v2452
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2453)))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v2454
	v2457 = v49 + int32(8)
	F_add_exact_object_address(m, v2457, v2259)
	mBase = m.M
	v2459 = m.ExcPending
	if v2459 != 0 {
		goto L1
	} else {
		goto L178
	}
L177:
	;
	goto L175
L178:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v2452+l17)))
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2453)))
	if v2461 != v2462 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v2461
	F_add_exact_object_address(m, v2457, v2259)
	mBase = m.M
	v2466 = m.ExcPending
	if v2466 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	v2468 = v2461
	goto L181
L181:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v2452+l18)))
	if v2468 != v2470 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v2453)))
	v2468 = v2467
	goto L181
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v2470
	F_add_exact_object_address(m, v49+int32(8), v2259)
	mBase = m.M
	v2476 = m.ExcPending
	if v2476 != 0 {
		goto L1
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v2478 = v2405 + int32(1)
	if v2478 != l19 {
		v2405 = v2478
		goto L176
	} else {
		goto L187
	}
L186:
	;
	goto L185
L187:
	;
	goto L177
L188:
	;
	F_free_object_addresses(m, v2259)
	mBase = m.M
	v2532 = m.ExcPending
	if v2532 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	if l26 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	F_recordDependencyOnSingleRelExpr(m, v2527, l26, l8, int32(110), int32(0))
	mBase = m.M
	v2536 = m.ExcPending
	if v2536 != 0 {
		goto L1
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v2538 = *(*int32)(unsafe.Add(mBase, _c_F_CreateConstraintEntry[0]))
	if v2538 != 0 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L192
L194:
	;
	F_RunObjectPostCreateHook(m, int32(2606), v2025, int32(0), l32)
	mBase = m.M
	v2542 = m.ExcPending
	if v2542 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	m.G0 = v49 + int32(352)
	return v2025
L197:
	;
	goto L196
}
func F_transformConstraintAttrs(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l1 == v3 {
		goto L12
	} else {
		goto L13
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L38
	} else {
		goto L126
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L38
	} else {
		goto L121
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L38
	} else {
		goto L116
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L38
	} else {
		goto L111
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L38
	} else {
		goto L106
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L38
	} else {
		goto L101
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L38
	} else {
		goto L96
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L38
	} else {
		goto L91
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L38
	} else {
		goto L86
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L38
	} else {
		goto L81
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L38
	} else {
		goto L78
	}
L12:
	;
	m.G0 = v16 + int32(16)
	return
L13:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v27 = v3
	v29 = v3
	v30 = v3
	v32 = v3
	v35 = v3
	goto L15
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v35<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 != int32(161) {
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v44 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	switch v47 - int32(10) {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	case 5:
		goto L22
	default:
		v219 = v44
		v220 = v44
		v221 = v44
		v222 = v40
		goto L18
	}
L18:
	;
	v224 = v35 + int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v224 < v225 {
		v27 = v219
		v29 = v220
		v30 = v221
		v32 = v222
		v35 = v224
		goto L15
	} else {
		goto L77
	}
L19:
	;
	v219 = v216
	v220 = v217
	v221 = v218
	v222 = v32
	goto L18
L20:
	;
	v216 = v214
	v217 = v215
	v218 = v30
	goto L19
L21:
	;
	v214 = v213
	v215 = v29
	goto L20
L22:
	;
	if v32 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L23:
	;
	if v32 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L24:
	;
	if v32 == int32(0) {
		goto L4
	} else {
		goto L56
	}
L25:
	;
	if v32 == int32(0) {
		goto L6
	} else {
		goto L44
	}
L26:
	;
	if v32 == int32(0) {
		goto L8
	} else {
		goto L31
	}
L27:
	;
	if v32 == int32(0) {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v52-int32(6)) {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v30&int32(1) != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+12)) = uint8(v59)
	v216 = v27
	v217 = v29
	v218 = v59
	goto L19
L31:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v64-int32(6)) {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	if v30&int32(1) != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+12)) = uint8(v71)
	v73 = int32(1)
	if v29&v73 == v71 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v216 = v27
	v217 = v44
	v218 = v73
	goto L19
L35:
	;
	goto L36
L36:
	;
	v78 = int32(1)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+13)))
	if v79 != v78 {
		v219 = v27
		v220 = v78
		v221 = v73
		v222 = v32
		goto L18
	} else {
		goto L37
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return
L39:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_0), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(3981), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v103-int32(6)) {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	if v29&int32(1) != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v110 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+13)) = uint8(v110)
	if v30&v110 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+12)) = uint8(v116)
	v216 = v27
	v217 = v116
	v218 = v44
	goto L19
L48:
	;
	goto L49
L49:
	;
	v119 = int32(1)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+12)))
	if v121 != 0 {
		v219 = v27
		v220 = v119
		v221 = v119
		v222 = v32
		goto L18
	} else {
		goto L50
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L38
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L38
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_0), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L38
	} else {
		goto L53
	}
L53:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L38
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(4007), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L38
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v143-int32(6)) {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	if v29&int32(1) != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+13)) = uint8(v150)
	v214 = v27
	v215 = int32(1)
	goto L20
L59:
	;
	if v27&int32(1) != 0 {
		goto L2
	} else {
		goto L67
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L38
	} else {
		goto L62
	}
L61:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	switch v155 - int32(5) {
	case 0, 4:
		goto L59
	default:
		goto L60
	}
L62:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L38
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_3), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L38
	} else {
		goto L64
	}
L64:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L38
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(4032), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v179 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+14)) = uint8(v179)
	v213 = v179
	goto L21
L68:
	;
	if v27&int32(1) != 0 {
		goto L1
	} else {
		goto L76
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L38
	} else {
		goto L71
	}
L70:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	switch v184 - int32(5) {
	case 0, 4:
		goto L68
	default:
		goto L69
	}
L71:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L38
	} else {
		goto L72
	}
L72:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_4), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L38
	} else {
		goto L73
	}
L73:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L38
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(4049), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L38
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+16)) = uint8(v208)
	v210 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)) = uint16(v210)
	v213 = int32(1)
	goto L21
L77:
	;
	goto L16
L78:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v247
	F_errmsg_internal(m, int32(_a_F_transformConstraintAttrs_5), v16)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L38
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(3945), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L38
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L38
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_6), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L38
	} else {
		goto L83
	}
L83:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L38
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(3953), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L38
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L38
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_7), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L38
	} else {
		goto L88
	}
L88:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L38
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(3958), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L38
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L38
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_8), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L38
	} else {
		goto L93
	}
L93:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L38
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(3968), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L38
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L38
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_7), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L38
	} else {
		goto L98
	}
L98:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L38
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(3973), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L38
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L38
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_9), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L38
	} else {
		goto L103
	}
L103:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L38
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(3989), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L38
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L38
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_10), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L38
	} else {
		goto L108
	}
L108:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L38
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(3994), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L38
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L38
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_11), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L38
	} else {
		goto L113
	}
L113:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L38
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(4015), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L38
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L38
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_10), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L38
	} else {
		goto L118
	}
L118:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L38
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(4020), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L38
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L38
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_12), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L38
	} else {
		goto L123
	}
L123:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L38
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(4037), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L38
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L38
	} else {
		goto L127
	}
L127:
	;
	F_errmsg(m, int32(_a_F_transformConstraintAttrs_12), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L38
	} else {
		goto L128
	}
L128:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v40)+104))
	F_parser_errposition(m, l0, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L38
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_transformConstraintAttrs_1), int32(4054), int32(_a_F_transformConstraintAttrs_2))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L38
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
