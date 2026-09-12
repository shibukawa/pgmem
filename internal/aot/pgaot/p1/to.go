package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BeginCopyTo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v111 int32
	_ = v111
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v147 int32
	_ = v147
	var v163 int32
	_ = v163
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v197 int32
	_ = v197
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v233 int32
	_ = v233
	var v249 int32
	_ = v249
	var v266 int32
	_ = v266
	var v282 int32
	_ = v282
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v318 int32
	_ = v318
	var v334 int32
	_ = v334
	var v351 int32
	_ = v351
	var v367 int32
	_ = v367
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v403 int32
	_ = v403
	var v420 int32
	_ = v420
	var v436 int32
	_ = v436
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v472 int32
	_ = v472
	var v488 int32
	_ = v488
	var v505 int32
	_ = v505
	var v521 int32
	_ = v521
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v557 int32
	_ = v557
	var v574 int32
	_ = v574
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v679 int32
	_ = v679
	var v694 int32
	_ = v694
	var v710 int32
	_ = v710
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v756 int32
	_ = v756
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v791 int32
	_ = v791
	var v806 int32
	_ = v806
	var v824 int32
	_ = v824
	var v841 int32
	_ = v841
	var v857 int32
	_ = v857
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v894 int32
	_ = v894
	var v909 int32
	_ = v909
	var v925 int32
	_ = v925
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v975 int32
	_ = v975
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1006 int32
	_ = v1006
	var v1013 int32
	_ = v1013
	var v1029 int32
	_ = v1029
	var v1044 int32
	_ = v1044
	var v1060 int32
	_ = v1060
	var v1077 int32
	_ = v1077
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1107 int32
	_ = v1107
	var v1121 int32
	_ = v1121
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1282 int32
	_ = v1282
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1310 int32
	_ = v1310
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1341 int32
	_ = v1341
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1368 int32
	_ = v1368
	var v1383 int32
	_ = v1383
	var v1411 int32
	_ = v1411
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1467 int32
	_ = v1467
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1514 int32
	_ = v1514
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1578 int32
	_ = v1578
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1611 int32
	_ = v1611
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1649 int32
	_ = v1649
	var v1664 int32
	_ = v1664
	var v1680 int32
	_ = v1680
	var v1697 int32
	_ = v1697
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1761 int32
	_ = v1761
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1799 int32
	_ = v1799
	var v1804 int32
	_ = v1804
	var v1818 int32
	_ = v1818
	var v1834 int32
	_ = v1834
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1865 int32
	_ = v1865
	var v1888 int32
	_ = v1888
	var v1905 int32
	_ = v1905
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1941 int32
	_ = v1941
	var v1954 int32
	_ = v1954
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1966 int32
	_ = v1966
	var v1982 int32
	_ = v1982
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v2002 int32
	_ = v2002
	var v2016 int32
	_ = v2016
	var v2017 int32
	_ = v2017
	var v2035 int32
	_ = v2035
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2073 int32
	_ = v2073
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2107 int32
	_ = v2107
	var v2124 int32
	_ = v2124
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2194 int32
	_ = v2194
	var v2200 int32
	_ = v2200
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2214 int32
	_ = v2214
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2245 int32
	_ = v2245
	var v2252 int32
	_ = v2252
	var v2258 int32
	_ = v2258
	var v2263 int32
	_ = v2263
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2269 int32
	_ = v2269
	var v2364 int32
	_ = v2364
	var v2367 int32
	_ = v2367
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2383 int64
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2388 int32
	_ = v2388
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2441 int32
	_ = v2441
	var v2456 int32
	_ = v2456
	var v2472 int32
	_ = v2472
	var v2489 int32
	_ = v2489
	var v2505 int32
	_ = v2505
	var v2520 int32
	_ = v2520
	var v2536 int32
	_ = v2536
	var v2553 int32
	_ = v2553
	var v2569 int32
	_ = v2569
	var v2584 int32
	_ = v2584
	var v2600 int32
	_ = v2600
	var v2617 int32
	_ = v2617
	var v2644 int32
	_ = v2644
	var v2648 int32
	_ = v2648
	var v2649 int64
	_ = v2649
	var v2653 int32
	_ = v2653
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2660 int32
	_ = v2660
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2666 int32
	_ = v2666
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2679 int32
	_ = v2679
	v6 = l5
	v9 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(224)
	m.G0 = v33
	v47 = v9
	v48 = v9
	v49 = v9
	v50 = v9
	v51 = v9
	v52 = v9
	v53 = v9
	v54 = v9
	v55 = v9
	v56 = v9
	v57 = v9
	v58 = v9
	v59 = int32(-1)
	v60 = v9
	v64 = v33
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
	if v59 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	v2648 = int32(m.ExcTag)
	v2649 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2648 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2569 = m.ExcPending
	if v2569 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L272
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2505 = m.ExcPending
	if v2505 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L268
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L264
	}
L10:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v2143)))
	if v2155 != 0 {
		goto L232
	} else {
		goto L233
	}
L11:
	;
	v70 = int32(16)
	v71 = v64 - v70
	m.G0 = v71
	v74 = v71 - v70
	m.G0 = v74
	v77 = v74 - int32(96)
	m.G0 = v77
	v80 = v77 - int32(160)
	m.G0 = v80
	*(*int64)(unsafe.Add(mBase, uint32(v71))) = int64(21474836484)
	v85 = *(*int64)(unsafe.Add(mBase, _consts[357]))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+8)) = v85
	v88 = *(*int64)(unsafe.Add(mBase, _consts[358]))
	*(*int64)(unsafe.Add(mBase, uint32(v74))) = v88
	if l1 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v1736 = v47
	v1737 = v48
	v1738 = v49
	v1739 = v50
	v1740 = v51
	v1741 = v52
	v1742 = v53
	v1743 = v54
	v1744 = v55
	v1745 = v56
	v1746 = v57
	v1747 = v58
	v1749 = v60
	v1753 = v64
	goto L13
L13:
	;
	if v1749 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v589 = F_palloc0(m, int32(184))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L51
	}
L15:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+119)))
	switch v93 - int32(83) {
	case 0:
		goto L18
	default:
		goto L16
	case 19:
		goto L19
	case 26:
		goto L20
	case 29:
		goto L17
	case 31:
		goto L14
	case 35:
		goto L21
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L47
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L42
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L38
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L33
	}
L20:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+129)))
	if v181 != 0 {
		goto L14
	} else {
		goto L27
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(151027844))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v127 + int32(4)
	F_errmsg(m, int32(657148), v33+int32(96))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errhint(m, int32(542300), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(654), int32(231318))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L3
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(1088))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v213 + int32(4)
	F_errmsg(m, int32(657535), v33+int32(112))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errhint(m, int32(606749), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(662), int32(231318))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L32
	}
L32:
	;
	goto L3
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(151027844))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = v298 + int32(4)
	F_errmsg(m, int32(679301), v33+int32(128))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errhint(m, int32(542300), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(669), int32(231318))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L37
	}
L37:
	;
	goto L3
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(151027844))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+144)) = v383 + int32(4)
	F_errmsg(m, int32(681690), v33+int32(144))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(674), int32(231318))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L41
	}
L41:
	;
	goto L3
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(151027844))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+160)) = v452 + int32(4)
	F_errmsg(m, int32(680665), v33+int32(160))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errhint(m, int32(542300), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(680), int32(231318))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L46
	}
L46:
	;
	goto L3
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(151027844))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L48
	}
L48:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v537 + int32(4)
	F_errmsg(m, int32(668733), v33+int32(80))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(685), int32(231318))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L50
	}
L50:
	;
	goto L3
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v604 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v609 = F_AllocSetContextCreateInternal(m, v604, int32(486959), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589)+164)) = v609
	v612 = int32(4443856)
	v613 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v609
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	v629 = v589 + int32(48)
	F_ProcessCopyOptions(m, l0, v629, int32(0), l7)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+54)))
	if v633 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v639 = int32(1576224)
	goto L56
L55:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+52)))
	if v637 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589))) = v639
	if l1 != 0 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v638 = int32(1576240)
	goto L59
L58:
	;
	v638 = int32(1576256)
	goto L59
L59:
	;
	v639 = v638
	goto L56
L60:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1196)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	v1210 = v589 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	v1212 = F_CopyGetAttnums(m, v1197, v1192, l6)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L126
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589)+24)) = l1
	v1192 = l1
	v1196 = l1 + int32(52)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v642 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v589)+24)) = v642
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v660 = F_pg_analyze_and_rewrite_fixedparams(m, l2, v644, v642, v642, v642)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L64
	}
L64:
	;
	if v660 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v660)+12))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v660)+4))
	if int32(1) < v729 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(1088))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errmsg(m, int32(486796), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(739), int32(231318))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L3
L72:
	;
	v756 = int32(0)
	goto L75
L73:
	;
	goto L74
L74:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v773)+28))
	if v774 != 0 {
		goto L79
	} else {
		goto L80
	}
L75:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v728+v756<<(uint(int32(2))%32))))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v766)+8))
	switch v767 - int32(3) {
	case 0:
		goto L8
	case 1:
		goto L9
	default:
		goto L77
	}
L76:
	;
	goto L7
L77:
	;
	v771 = v756 + int32(1)
	if v729 != v771 {
		v756 = v771
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v773)+4))
	if v875 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(1088))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L83
	}
L83:
	;
	if v775 == int32(242) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errmsg(m, int32(423498), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errmsg(m, int32(408703), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(772), int32(231318))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L88
	}
L88:
	;
	goto L3
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(778), int32(231318))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L90
	}
L90:
	;
	goto L3
L91:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v958 = F_pg_plan_query(m, v773, v943, int32(2048), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L98
	}
L92:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v773)+96))
	if v878 != 0 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(1088))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errmsg(m, int32(343276), int32(0))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(794), int32(231318))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L97
	}
L97:
	;
	goto L3
L98:
	;
	if l3 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1092 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1092)))
	goto L119
L100:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v958)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v975 = int32(0)
	if v962 == v975 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v1013 != 0 {
		goto L99
	} else {
		goto L114
	}
L102:
	;
	v1013 = int32(0)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v962)+4))
	if v981 <= int32(0) {
		v1006 = v975
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v1013 = v1006
	goto L101
L106:
	;
	v984 = int32(0)
	if v984 < v981 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v987 = v981
	goto L109
L108:
	;
	v987 = v984
	goto L109
L109:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v962)+12))
	v990 = int32(0)
	goto L110
L110:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v988+v990<<(uint(int32(2))%32))))
	v999 = base.B2i32(v998 == l3)
	if v998 == l3 {
		v1006 = v999
		goto L105
	} else {
		goto L112
	}
L111:
	;
	v1006 = v999
	goto L105
L112:
	;
	v1001 = v990 + int32(1)
	if v1001 != v987 {
		v990 = v1001
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(325))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errmsg(m, int32(439877), int32(0))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(823), int32(231318))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L118
	}
L118:
	;
	goto L3
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_PushCopiedSnapshot(m, v1093)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1135 = F_CreateDestReceiver(m, int32(8))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1135)+20)) = v589
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1152 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1152)))
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1166 = int32(0)
	v1170 = F_CreateQueryDesc(m, v958, v1138, v1153, v1166, v1135, v1166, v1166, v1166)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589)+28)) = v1170
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_ExecutorStart(m, v1170, int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L125
	}
L125:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v589)+24))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v589)+28))
	v1192 = v1188
	v1196 = v1189 + int32(36)
	goto L60
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589)+32)) = v1212
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1197)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1228 = F_palloc0(m, v1215)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589)+100)) = v1228
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+96)))
	if v1231 == int32(1) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v629)))
	if v1467 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L129:
	;
	v1236 = F__emscripten_memset_bulkmem(m, v1228, base.I32_extend8_s(int32(1)), v1215)
	mBase = m.M
	goto L132
L130:
	;
	goto L131
L131:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v589)+92))
	if v1237 == int32(0) {
		goto L128
	} else {
		goto L133
	}
L132:
	;
	goto L128
L133:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1210)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1253 = F_CopyGetAttnums(m, v1197, v1240, v1237)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L134
	}
L134:
	;
	if v1253 == int32(0) {
		goto L128
	} else {
		goto L135
	}
L135:
	;
	v1257 = int32(0)
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+4))
	if v1258 <= v1257 {
		goto L128
	} else {
		goto L136
	}
L136:
	;
	v1282 = v1257
	goto L137
L137:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+12))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1291+v1282<<(uint(int32(2))%32))))
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1197)))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v589)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1310 = int32(0)
	if v1297 == v1310 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L128
L139:
	;
	v1350 = v1295 - int32(1)
	if v1348 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L140:
	;
	v1348 = int32(0)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+4))
	if v1316 <= int32(0) {
		v1341 = v1310
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v1348 = v1341
	goto L139
L144:
	;
	v1319 = int32(0)
	if v1319 < v1316 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v1322 = v1316
	goto L147
L146:
	;
	v1322 = v1319
	goto L147
L147:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1297)+12))
	v1325 = int32(0)
	goto L148
L148:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1323+v1325<<(uint(int32(2))%32))))
	v1334 = base.B2i32(v1333 == v1295)
	if v1333 == v1295 {
		v1341 = v1334
		goto L143
	} else {
		goto L150
	}
L149:
	;
	v1341 = v1334
	goto L143
L150:
	;
	v1336 = v1325 + int32(1)
	if v1336 != v1322 {
		v1325 = v1336
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v589)+100))
	v1431 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1429+v1350))) = uint8(v1431)
	v1434 = v1282 + v1431
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1253)+4))
	if v1434 < v1435 {
		v1282 = v1434
		goto L137
	} else {
		goto L159
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(393348))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = int32(515511)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+68)) = v1197 + v1296<<(uint(int32(4))%32) + v1350*int32(100) + int32(24)
	F_errmsg(m, int32(486673), v33-int32(-64))
	mBase = m.M
	v1411 = m.ExcPending
	if v1411 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(881), int32(231318))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L158
	}
L158:
	;
	goto L3
L159:
	;
	goto L138
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1483 = *(*int32)(unsafe.Add(mBase, _consts[359]))
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+4))
	goto L163
L161:
	;
	v1485 = v58
	v1486 = v1467
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589)+16)) = v1486
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1501 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1501)+4))
	goto L164
L163:
	;
	v1485 = v1484
	v1486 = v1484
	goto L162
L164:
	;
	v1503 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v589)+4)) = v1503
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v589)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v589)+21)) = uint8(base.B2i32(base.Ui32(v1505-int32(35)) < base.Ui32(int32(7))))
	v1514 = base.B2i32(v1486 != v1502) & base.B2i32(v1505 != v1503)
	*(*uint8)(unsafe.Add(mBase, uint32(v589)+20)) = uint8(v1514)
	if l4 == v1503 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v74)+8)) = int64(3)
	v1521 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	if v1521 == int32(2) {
		v2134 = v589
		v2135 = v74
		v2136 = v49
		v2137 = v77
		v2138 = v80
		v2139 = v52
		v2140 = v53
		v2141 = v71
		v2142 = v55
		v2143 = v1210
		v2144 = v613
		v2145 = v1485
		goto L10
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1539 = F_pstrdup(m, l4)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L169
	}
L168:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	*(*int32)(unsafe.Add(mBase, uint32(v589)+8)) = v1525
	v2134 = v589
	v2135 = v74
	v2136 = v49
	v2137 = v77
	v2138 = v80
	v2139 = v52
	v2140 = v53
	v2141 = v71
	v2142 = v55
	v2143 = v1210
	v2144 = v613
	v2145 = v1485
	goto L10
L169:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v589)+40)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v589)+36)) = v1539
	v1544 = v589 + int32(36)
	if v6 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v74)+8)) = int64(2)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1560 = F_OpenPipeStream(m, v1539, int32(30419))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v74)+8)) = int64(1)
	v1631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v1631 != int32(47) {
		goto L179
	} else {
		goto L180
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v589)+8)) = v1560
	if v1560 != 0 {
		v2134 = v589
		v2135 = v74
		v2136 = v1544
		v2137 = v77
		v2138 = v80
		v2139 = v52
		v2140 = v53
		v2141 = v71
		v2142 = v55
		v2143 = v1210
		v2144 = v613
		v2145 = v1485
		goto L10
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode_for_file_access(m)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L176
	}
L176:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1544)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v1593
	F_errmsg(m, int32(286362), v33+int32(48))
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(934), int32(231318))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L178
	}
L178:
	;
	goto L3
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	v1712 = int32(4337620)
	v1713 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	*(*int32)(unsafe.Add(mBase, _consts[361])) = int32(18)
	v1716 = F___syscall_ret(m, v1713)
	mBase = m.M
	goto L186
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(33579140))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errmsg(m, int32(369878), int32(0))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1544
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(950), int32(231318))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L185
	}
L185:
	;
	goto L3
L186:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	v1721 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v33 + int32(172)
	goto L190
L188:
	;
	v1736 = v589
	v1737 = v74
	v1738 = v1544
	v1739 = v77
	v1740 = v80
	v1741 = v1721
	v1742 = v1719
	v1743 = v71
	v1744 = v1716
	v1745 = v1210
	v1746 = v613
	v1747 = v1485
	v1749 = int32(0)
	v1753 = v80
	goto L13
L190:
	;
	goto L188
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(v1804)+76))
	if v1954 < int32(0) {
		goto L210
	} else {
		goto L211
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1740
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1738)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	v1775 = F_AllocateFile(m, v1761, int32(30419))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1742
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	v1923 = int32(4337620)
	v1924 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	*(*int32)(unsafe.Add(mBase, _consts[361])) = v1744
	v1927 = F___syscall_ret(m, v1924)
	mBase = m.M
	goto L206
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1736)+8)) = v1775
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1742
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	v1795 = int32(4337620)
	v1796 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	*(*int32)(unsafe.Add(mBase, _consts[361])) = v1744
	v1799 = F___syscall_ret(m, v1796)
	mBase = m.M
	goto L196
L196:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1742
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1741
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1736)+8))
	if v1804 != 0 {
		goto L191
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	v1818 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errcode_for_file_access(m)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L199
	}
L199:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1738)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v1849
	F_errmsg(m, int32(280789), v33)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L200
	}
L200:
	;
	if base.B2i32(v1818 != int32(44))&base.B2i32(v1818 != int32(2)) == int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errhint(m, int32(535800), int32(0))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errfinish(m, int32(474514), int32(973), int32(231318))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	goto L3
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_pg_re_throw(m)
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L207
	}
L207:
	;
	goto L3
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	if v1966 < int32(0) {
		goto L217
	} else {
		goto L218
	}
L209:
	;
	if v1959 < int32(0) {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1804)+60))
	v1959 = v1957
	goto L209
L211:
	;
	goto L212
L212:
	;
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1804)+60))
	v1959 = v1958
	goto L209
L213:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(8)
	v1966 = int32(-1)
	goto L215
L214:
	;
	v1966 = v1959
	goto L215
L215:
	;
	goto L208
L216:
	;
	if v1986 != 0 {
		goto L220
	} else {
		goto L221
	}
L217:
	;
	v1982 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v1986 = v1982
	goto L216
L218:
	;
	goto L219
L219:
	;
	v1985 = F___fstatat(m, v1966, int32(717063), v1739, int32(4096))
	mBase = m.M
	v1986 = v1985
	goto L216
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+4))
	if v2053&int32(61440) != int32(16384) {
		v2134 = v1736
		v2135 = v1737
		v2136 = v1738
		v2137 = v1739
		v2138 = v1740
		v2139 = v1741
		v2140 = v1742
		v2141 = v1743
		v2142 = v1744
		v2143 = v1745
		v2144 = v1746
		v2145 = v1747
		goto L10
	} else {
		goto L227
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errcode_for_file_access(m)
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L224
	}
L224:
	;
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1738)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v2017
	F_errmsg(m, int32(284292), v33+int32(32))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errfinish(m, int32(474514), int32(980), int32(231318))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L226
	}
L226:
	;
	goto L3
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errcode(m, int32(151027844))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L229
	}
L229:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v1738)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v2089
	F_errmsg(m, int32(12569), v33+int32(16))
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v1744
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v1738
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v1747
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v1745
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v1746
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v1740
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v1737
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v1743
	F_errfinish(m, int32(474514), int32(985), int32(231318))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		v2644 = v1753
		goto L6
	} else {
		goto L231
	}
L231:
	;
	goto L3
L232:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v2155)+56))
	v2158 = v2156
	goto L234
L233:
	;
	v2158 = int32(0)
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v2139
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v2140
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v2142
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v2136
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v2145
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v2143
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v2144
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v2134
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v2138
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v2137
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v2135
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v2141
	v2174 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v2174 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v2139
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v2140
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v2142
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v2136
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v2145
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v2143
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v2144
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v2134
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v2138
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v2137
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v2135
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v2141
	v2245 = int32(0)
	v2252 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v2252 == v2245 {
		goto L248
	} else {
		goto L249
	}
L236:
	;
	goto L235
L237:
	;
	v2178 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v2178 != int32(1) {
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v2181 = int32(4438516)
	v2183 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2184 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2183 + v2184
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v2174)))
	*(*int32)(unsafe.Add(mBase, uint32(v2174))) = v2187 + v2184
	*(*int32)(unsafe.Add(mBase, uint32(v2174)+220)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v2174)+224)) = v2158
	v2194 = v2174 + int32(232)
	if v2194&int32(3) == int32(0) {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v2174)))
	v2221 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2174))) = v2220 + v2221
	v2224 = int32(4438516)
	v2226 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2226 - v2221
	goto L236
L240:
	;
	v2200 = v2174 + int32(392)
	if base.Ui32(v2200) <= base.Ui32(v2194) {
		goto L239
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v2217 = F___memset(m, v2194, int32(0), int32(160))
	mBase = m.M
	goto L239
L243:
	;
	v2204 = v2174 + int32(236)
	if base.Ui32(v2204) < base.Ui32(v2200) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v2206 = v2200
	goto L246
L245:
	;
	v2206 = v2204
	goto L246
L246:
	;
	v2214 = F___memset(m, v2194, int32(0), (v2206-v2174-int32(233))&int32(-4)+int32(4))
	mBase = m.M
	goto L239
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2134)+176)) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v2144
	m.G0 = v33 + int32(224)
	return v2134
L248:
	;
	goto L247
L249:
	;
	goto L250
L250:
	;
	v2258 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v2258&int32(1) == int32(0) {
		goto L248
	} else {
		goto L251
	}
L251:
	;
	v2263 = int32(4438516)
	v2265 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2266 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2265 + v2266
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v2252)))
	*(*int32)(unsafe.Add(mBase, uint32(v2252))) = v2269 + v2266
	goto L253
L252:
	;
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(v2252)))
	v2400 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2252))) = v2399 + v2400
	v2403 = int32(4438516)
	v2405 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v2405 - v2400
	goto L248
L253:
	;
	goto L255
L255:
	;
	goto L256
L256:
	;
	goto L260
L260:
	;
	v2364 = int32(0)
	v2367 = v2245
	goto L261
L261:
	;
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v2141+v2367<<(uint(int32(2))%32))))
	v2377 = int32(3)
	v2383 = *(*int64)(unsafe.Add(mBase, uint32(v2135+v2367<<(uint(v2377)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v2252+int32(232)+v2376<<(uint(v2377)%32)))) = v2383
	v2385 = int32(1)
	v2388 = v2364 + v2385
	if v2388 != int32(2) {
		v2364 = v2388
		v2367 = v2367 + v2385
		goto L261
	} else {
		goto L263
	}
L262:
	;
	goto L252
L263:
	;
	goto L262
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(1088))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errmsg(m, int32(486755), int32(0))
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(757), int32(231318))
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L267
	}
L267:
	;
	goto L3
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(1088))
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errmsg(m, int32(486908), int32(0))
	mBase = m.M
	v2536 = m.ExcPending
	if v2536 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(753), int32(231318))
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L271
	}
L271:
	;
	goto L3
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errcode(m, int32(1088))
	mBase = m.M
	v2584 = m.ExcPending
	if v2584 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errmsg(m, int32(486848), int32(0))
	mBase = m.M
	v2600 = m.ExcPending
	if v2600 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+180)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+176)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+184)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+188)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+192)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v33)+196)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v33)+200)) = v613
	*(*int32)(unsafe.Add(mBase, uint32(v33)+204)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v33)+208)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v33)+212)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v33)+216)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v33)+220)) = v71
	F_errfinish(m, int32(474514), int32(762), int32(231318))
	mBase = m.M
	v2617 = m.ExcPending
	if v2617 != 0 {
		v2644 = v80
		goto L6
	} else {
		goto L275
	}
L275:
	;
	goto L5
L276:
	;
	v2653 = int32(v2649)
	m.G0 = v2644
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(v2653)+4))
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(v2653)))
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v2656)))
	if v33+int32(172) == v2660 {
		goto L279
	} else {
		goto L280
	}
L277:
	;
	m.ExcPending = 1
	goto L285
L278:
	;
	if v2663 != 0 {
		goto L282
	} else {
		goto L283
	}
L279:
	;
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v2656)+4))
	v2663 = v2662
	goto L281
L280:
	;
	v2663 = int32(0)
	goto L281
L281:
	;
	goto L278
L282:
	;
	v2664 = *(*int32)(unsafe.Add(mBase, uint32(v33)+220))
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v33)+216))
	v2666 = *(*int32)(unsafe.Add(mBase, uint32(v33)+212))
	v2667 = *(*int32)(unsafe.Add(mBase, uint32(v33)+208))
	v2668 = *(*int32)(unsafe.Add(mBase, uint32(v33)+204))
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(v33)+200))
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v33)+196))
	v2671 = *(*int32)(unsafe.Add(mBase, uint32(v33)+192))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v33)+188))
	v2673 = *(*int32)(unsafe.Add(mBase, uint32(v33)+184))
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(v33)+180))
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(v33)+176))
	v47 = v2668
	v48 = v2665
	v49 = v2672
	v50 = v2666
	v51 = v2667
	v52 = v2674
	v53 = v2675
	v54 = v2664
	v55 = v2673
	v56 = v2670
	v57 = v2669
	v58 = v2671
	v59 = v2663
	v60 = v2655
	v64 = v2644
	goto L1
L283:
	;
	goto L284
L284:
	;
	F___wasm_longjmp(m, v2656, v2655)
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	return int32(0)
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyToBinaryOneRow(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v17 = v15
	goto L3
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	v18 = int32(8)
	v24 = v17<<(uint(v18)%32) | int32(base.Ui32(v17&int32(65280))>>(uint(v18)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+6)) = uint16(v24)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v26, v11+int32(6), int32(2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v32 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_CopySendEndOfRow(m, l0)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L20
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v35 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v42 = int32(0)
	goto L9
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v42<<(uint(int32(2))%32))))
	v53 = v51 - int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v54))))
	if v56 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L6
L11:
	;
	v115 = v42 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v115 < v116 {
		v42 = v115
		goto L9
	} else {
		goto L19
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(-1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v59, v11+int32(8), int32(4))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v53<<(uint(int32(2))%32))))
	v73 = F_SendFunctionCall(m, v13+v53*int32(28), v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v78 = int32(4)
	v79 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - v78
	v80 = int32(24)
	v82 = int32(65280)
	v84 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v79<<(uint(v80)%32) | v79&v82<<(uint(v84)%32) | (int32(base.Ui32(v79)>>(uint(v84)%32))&v82 | int32(base.Ui32(v79)>>(uint(v80)%32)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v96, v11+int32(12), v78)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v103 = int32(4)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	F_appendBinaryStringInfo(m, v102, v73+v103, int32(base.Ui32(v105)>>(uint(int32(2))%32))-v103)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	goto L10
L20:
	;
	m.G0 = v11 + int32(16)
	return
}
func F_coerce_to_specific_type_typmod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = F_exprType(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l2 == v12 {
			v24 = l1
			v25 = F_expression_returns_set(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
							F_errmsg(m, int32(99929), v10)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = F_exprLocation(m, v24)
								mBase = m.M
								F_parser_errposition(m, l0, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(477776), int32(1239), int32(403323))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
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
					m.G0 = v10 + int32(32)
					return v24
				}
			}
		} else {
			v20 = F_coerce_to_target_type(m, l0, l1, v12, l2, l3, int32(1), int32(2), int32(-1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = F_format_type_be(m, l2)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = F_format_type_be(m, v12)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v40
									*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l4
									F_errmsg(m, int32(179091), v10+int32(16))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v50 = F_exprLocation(m, l1)
										mBase = m.M
										F_parser_errposition(m, l0, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(477776), int32(1229), int32(403323))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
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
				} else {
					v24 = v20
					v25 = F_expression_returns_set(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = l4
									F_errmsg(m, int32(99929), v10)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v69 = F_exprLocation(m, v24)
										mBase = m.M
										F_parser_errposition(m, l0, v69)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(477776), int32(1239), int32(403323))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
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
							m.G0 = v10 + int32(32)
							return v24
						}
					}
				}
			}
		}
	}
}
func F_to_ascii_default(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_copy(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[356]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		v10 = F_encode_to_ascii(m, v3, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_to_bin64(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	v7 = m.G0
	v8 = int32(-64)
	v9 = v7 + v8
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v14 = v9 - v8
	v15 = v14
	v16 = v12
	goto L1
L1:
	;
	v21 = int32(1)
	v22 = v15 - v21
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v16)&v21)+uint32(_consts[1142]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v28)
	if base.Ui64(v16) < base.Ui64(int64(2)) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v36 = v14 - v22
	v38 = v36 + int32(4)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v9) < base.Ui32(v22) {
		v15 = v22
		v16 = int64(base.Ui64(v16) >> (uint(int64(1)) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(int32(2)) % 32)
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v9 - int32(-64)
	return v39
L9:
	;
	v48 = F__emscripten_memcpy_bulkmem(m, v39+int32(4), v22, v36)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
}
func F_to_oct32(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v7 = m.G0
	v8 = int32(-64)
	v9 = v7 + v8
	m.G0 = v9
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+20)))
	v13 = v9 - v8
	v14 = v13
	v15 = v11
	goto L1
L1:
	;
	v21 = v14 - int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v15)&int32(7))+uint32(_consts[1142]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v27)
	if base.Ui64(v15) < base.Ui64(int64(8)) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v35 = v13 - v21
	v37 = v35 + int32(4)
	v38 = F_palloc(m, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v9) < base.Ui32(v21) {
		v14 = v21
		v15 = int64(base.Ui64(v15) >> (uint(int64(3)) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v37 << (uint(int32(2)) % 32)
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v9 - int32(-64)
	return v38
L9:
	;
	v47 = F__emscripten_memcpy_bulkmem(m, v38+int32(4), v21, v35)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
}
func F_to_tsvector_byid(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		if v17 == int32(1) {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
			if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(2)
			} else {
				v43 = base.B2i32(v20 == int32(18)) << (uint(int32(4)) % 32)
				v45 = base.I32_div_u_s(v43, int32(6))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v45
				if base.Ui32(int32(11)) < base.Ui32(v43) {
					if base.Ui32(v43) < base.Ui32(int32(402653184)) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(67108863)
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(2)
				}
			}
		} else {
			v31 = int32(1)
			if v17&v31 != 0 {
				v43 = int32(base.Ui32(v17)>>(uint(v31)%32)) - v31
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
			}
			v45 = base.I32_div_u_s(v43, int32(6))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v45
			if base.Ui32(int32(11)) < base.Ui32(v43) {
				if base.Ui32(v43) < base.Ui32(int32(402653184)) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(67108863)
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(2)
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v62 = F_palloc(m, v59<<(uint(int32(4))%32))
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v62
			v65 = int32(1)
			v66 = v13 + v65
			v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v71 = v69 & v65
			if v71 != 0 {
				v72 = v66
			} else {
				v72 = v13 + int32(4)
			}
			if v69 == int32(1) {
				v75 = int32(4)
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
				if v77&int32(254) == int32(2) {
					v86 = v75
				} else {
					v86 = base.B2i32(v77 == int32(18)) << (uint(v75) % 32)
				}
				if v77 == int32(1) {
					v89 = v75
				} else {
					v89 = v86
				}
				v100 = v89
			} else {
				v90 = int32(1)
				if v71 != 0 {
					v100 = int32(base.Ui32(v69)>>(uint(v90)%32)) - v90
				} else {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v100 = int32(base.Ui32(v94)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			F_parsetext(m, v11, v9, v72, v100)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v103 != v13 {
					F_pfree(m, v13)
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v107 = F_make_tsvector(m, v9)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v107
						}
					}
				} else {
					v107 = F_make_tsvector(m, v9)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v107
					}
				}
			}
		}
	}
}
