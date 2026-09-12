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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v214 int32
	_ = v214
	var v249 int32
	_ = v249
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v512 int32
	_ = v512
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v754 int32
	_ = v754
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v998 int32
	_ = v998
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1238 int32
	_ = v1238
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1495 int32
	_ = v1495
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1680 int32
	_ = v1680
	var v1718 int32
	_ = v1718
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1805 int32
	_ = v1805
	var v1827 int32
	_ = v1827
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1962 int64
	_ = v1962
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v2003 int32
	_ = v2003
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2021 int32
	_ = v2021
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2028 int32
	_ = v2028
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2052 int32
	_ = v2052
	var v2105 int32
	_ = v2105
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2122 int32
	_ = v2122
	var v2178 int32
	_ = v2178
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2193 int32
	_ = v2193
	var v2246 int32
	_ = v2246
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2263 int32
	_ = v2263
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2331 int32
	_ = v2331
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2399 int32
	_ = v2399
	var v2405 int32
	_ = v2405
	var v2407 int32
	_ = v2407
	var v2460 int32
	_ = v2460
	var v2462 int32
	_ = v2462
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	v34 = int32(0)
	v48 = m.G0
	v50 = v48 - int32(240)
	m.G0 = v50
	v54 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v61 = F_strncpy(m, v50+int32(32), l0, int32(64))
	mBase = m.M
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+63)) = uint8(v62)
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
	v67 = l10 & int32(3)
	v70 = F_palloc(m, l10<<(uint(int32(2))%32))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v367 = v34
	goto L6
L6:
	;
	if l19 <= int32(0) {
		v1658 = v34
		v1659 = v34
		v1662 = v34
		v1663 = v34
		v1664 = v34
		goto L21
	} else {
		goto L22
	}
L7:
	;
	v72 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l10) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v77 = v72
	v115 = v34
	goto L11
L9:
	;
	v167 = v72
	goto L10
L10:
	;
	if v67 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v124 = int32(2)
	v127 = int32(1)
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v77<<(uint(v127)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v70+v77<<(uint(v124)%32)))) = v130
	v133 = v77 | v127
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v133<<(uint(v127)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v70+v133<<(uint(v124)%32)))) = v140
	v143 = v77 | v124
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v143<<(uint(v127)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v70+v143<<(uint(v124)%32)))) = v150
	v153 = v77 | int32(3)
	v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v153<<(uint(v127)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v70+v153<<(uint(v124)%32)))) = v160
	v162 = int32(4)
	v163 = v77 + v162
	v165 = v115 + v162
	if v165 != l10&int32(2147483644) {
		v77 = v163
		v115 = v165
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v167 = v163
	goto L10
L13:
	;
	goto L12
L14:
	;
	v214 = v167
	v249 = v34
	goto L17
L15:
	;
	goto L16
L16:
	;
	v322 = F_construct_array_builtin(m, v70, l10, int32(21))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L20
	}
L17:
	;
	v264 = int32(1)
	v267 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v214<<(uint(v264)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v70+v214<<(uint(int32(2))%32)))) = v267
	v272 = v249 + v264
	if v272 != v67 {
		v214 = v214 + v264
		v249 = v272
		goto L17
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	goto L18
L20:
	;
	v367 = v322
	goto L6
L21:
	;
	v1665 = int32(0)
	if l25 != 0 {
		goto L98
	} else {
		goto L99
	}
L22:
	;
	v374 = l19 & int32(3)
	if l23 < l19 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v377 = l19
	goto L25
L24:
	;
	v377 = l23
	goto L25
L25:
	;
	v380 = F_palloc(m, v377<<(uint(int32(2))%32))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v382 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l19) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v422 = v382
	v426 = int32(0)
	goto L30
L28:
	;
	v512 = v382
	goto L29
L29:
	;
	if v374 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v435 = int32(2)
	v438 = int32(1)
	v441 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v422<<(uint(v438)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v422<<(uint(v435)%32)))) = v441
	v444 = v422 | v438
	v451 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v444<<(uint(v438)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v444<<(uint(v435)%32)))) = v451
	v454 = v422 | v435
	v461 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v454<<(uint(v438)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v454<<(uint(v435)%32)))) = v461
	v464 = v422 | int32(3)
	v471 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v464<<(uint(v438)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v464<<(uint(v435)%32)))) = v471
	v473 = int32(4)
	v474 = v422 + v473
	v476 = v426 + v473
	if v476 != l19&int32(2147483644) {
		v422 = v474
		v426 = v476
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v512 = v474
	goto L29
L32:
	;
	goto L31
L33:
	;
	v559 = v512
	v560 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	v633 = l19 & int32(3)
	v636 = F_construct_array_builtin(m, v380, l19, int32(21))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L39
	}
L36:
	;
	v575 = int32(1)
	v578 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v559<<(uint(v575)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v559<<(uint(int32(2))%32)))) = v578
	v583 = v560 + v575
	if v583 != v374 {
		v559 = v559 + v575
		v560 = v583
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	goto L37
L39:
	;
	v638 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l19) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v678 = v638
	v680 = int32(0)
	goto L43
L41:
	;
	v754 = v638
	goto L42
L42:
	;
	if v633 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v692 = v678 << (uint(int32(2)) % 32)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l16+v692)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v692))) = v695
	v697 = int32(4)
	v698 = v692 | v697
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l16+v698)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v698))) = v701
	v704 = v692 | int32(8)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l16+v704)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v704))) = v707
	v710 = v692 | int32(12)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l16+v710)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v710))) = v713
	v716 = v678 + v697
	v718 = v680 + v697
	if v718 != l19&int32(2147483644) {
		v678 = v716
		v680 = v718
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v754 = v716
	goto L42
L45:
	;
	goto L44
L46:
	;
	v801 = v754
	v805 = int32(0)
	goto L49
L47:
	;
	goto L48
L48:
	;
	v872 = int32(1)
	if l19 <= v872 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v815 = v801 << (uint(int32(2)) % 32)
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l16+v815)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v815))) = v818
	v820 = int32(1)
	v823 = v805 + v820
	if v823 != v633 {
		v801 = v801 + v820
		v805 = v823
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L48
L51:
	;
	goto L50
L52:
	;
	v875 = v872
	goto L54
L53:
	;
	v875 = l19
	goto L54
L54:
	;
	v877 = v875 & int32(3)
	v880 = F_construct_array_builtin(m, v380, l19, int32(26))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v882 = int32(0)
	if int32(4) <= l19 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v922 = v882
	v924 = int32(0)
	goto L59
L57:
	;
	v998 = v882
	goto L58
L58:
	;
	if v877 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v936 = v922 << (uint(int32(2)) % 32)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l17+v936)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v936))) = v939
	v941 = int32(4)
	v942 = v936 | v941
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l17+v942)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v942))) = v945
	v948 = v936 | int32(8)
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l17+v948)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v948))) = v951
	v954 = v936 | int32(12)
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l17+v954)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v954))) = v957
	v960 = v922 + v941
	v962 = v924 + v941
	if v962 != v875&int32(2147483644) {
		v922 = v960
		v924 = v962
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v998 = v960
	goto L58
L61:
	;
	goto L60
L62:
	;
	v1045 = v998
	v1049 = int32(0)
	goto L65
L63:
	;
	goto L64
L64:
	;
	v1117 = v875 & int32(3)
	v1120 = F_construct_array_builtin(m, v380, l19, int32(26))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L68
	}
L65:
	;
	v1059 = v1045 << (uint(int32(2)) % 32)
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(l17+v1059)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1059))) = v1062
	v1064 = int32(1)
	v1067 = v1049 + v1064
	if v1067 != v877 {
		v1045 = v1045 + v1064
		v1049 = v1067
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	goto L66
L68:
	;
	v1122 = int32(0)
	if int32(4) <= l19 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v1162 = v1122
	v1164 = int32(0)
	goto L72
L70:
	;
	v1238 = v1122
	goto L71
L71:
	;
	if v1117 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v1176 = v1162 << (uint(int32(2)) % 32)
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1176)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1176))) = v1179
	v1181 = int32(4)
	v1182 = v1176 | v1181
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1182)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1182))) = v1185
	v1188 = v1176 | int32(8)
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1188)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1188))) = v1191
	v1194 = v1176 | int32(12)
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1194)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1194))) = v1197
	v1200 = v1162 + v1181
	v1202 = v1164 + v1181
	if v1202 != v875&int32(2147483644) {
		v1162 = v1200
		v1164 = v1202
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v1238 = v1200
	goto L71
L74:
	;
	goto L73
L75:
	;
	v1285 = v1238
	v1289 = int32(0)
	goto L78
L76:
	;
	goto L77
L77:
	;
	v1357 = F_construct_array_builtin(m, v380, l19, int32(26))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	v1299 = v1285 << (uint(int32(2)) % 32)
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1299)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1299))) = v1302
	v1304 = int32(1)
	v1307 = v1289 + v1304
	if v1307 != v1117 {
		v1285 = v1285 + v1304
		v1289 = v1307
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L77
L80:
	;
	goto L79
L81:
	;
	if l23 <= int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v1658 = int32(0)
	v1659 = v1357
	v1662 = v1120
	v1663 = v880
	v1664 = v636
	goto L21
L83:
	;
	goto L84
L84:
	;
	v1363 = l23 & int32(3)
	v1364 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l23) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v1405 = v1364
	v1409 = int32(0)
	goto L88
L86:
	;
	v1495 = v1364
	goto L87
L87:
	;
	if v1363 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v1418 = int32(2)
	v1421 = int32(1)
	v1424 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1405<<(uint(v1421)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1405<<(uint(v1418)%32)))) = v1424
	v1427 = v1405 | v1421
	v1434 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1427<<(uint(v1421)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1427<<(uint(v1418)%32)))) = v1434
	v1437 = v1405 | v1418
	v1444 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1437<<(uint(v1421)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1437<<(uint(v1418)%32)))) = v1444
	v1447 = v1405 | int32(3)
	v1454 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1447<<(uint(v1421)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1447<<(uint(v1418)%32)))) = v1454
	v1456 = int32(4)
	v1457 = v1405 + v1456
	v1459 = v1409 + v1456
	if v1459 != l23&int32(2147483644) {
		v1405 = v1457
		v1409 = v1459
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v1495 = v1457
	goto L87
L90:
	;
	goto L89
L91:
	;
	v1542 = v1495
	v1543 = v1364
	goto L94
L92:
	;
	goto L93
L93:
	;
	v1616 = F_construct_array_builtin(m, v380, l23, int32(21))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L1
	} else {
		goto L97
	}
L94:
	;
	v1558 = int32(1)
	v1561 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1542<<(uint(v1558)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v1542<<(uint(int32(2))%32)))) = v1561
	v1566 = v1543 + v1558
	if v1566 != v1363 {
		v1542 = v1542 + v1558
		v1543 = v1566
		goto L94
	} else {
		goto L96
	}
L95:
	;
	goto L93
L96:
	;
	goto L95
L97:
	;
	v1658 = v1616
	v1659 = v1357
	v1662 = v1120
	v1663 = v880
	v1664 = v636
	goto L21
L98:
	;
	v1668 = F_palloc(m, l10<<(uint(int32(2))%32))
	mBase = m.M
	v1669 = m.ExcPending
	if v1669 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	v1913 = v1665
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+232)) = int32(0)
	v1962 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+224)) = v1962
	*(*int64)(unsafe.Add(mBase, uint32(v50)+184)) = v1962
	*(*int64)(unsafe.Add(mBase, uint32(v50)+192)) = v1962
	*(*int64)(unsafe.Add(mBase, uint32(v50)+200)) = v1962
	*(*int64)(unsafe.Add(mBase, uint32(v50)+216)) = v1962
	*(*int64)(unsafe.Add(mBase, uint32(v50)+208)) = v1962
	*(*int64)(unsafe.Add(mBase, uint32(v50)+176)) = v1962
	v1978 = F_GetNewOidWithIndex(m, v54, int32(2667), int32(1))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L1
	} else {
		goto L115
	}
L101:
	;
	if l10 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v1911 = F_construct_array_builtin(m, v1668, l10, int32(26))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L1
	} else {
		goto L114
	}
L103:
	;
	v1673 = l10 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(l10) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v1680 = v1665
	v1718 = int32(0)
	goto L107
L105:
	;
	v1756 = v1665
	goto L106
L106:
	;
	if v1673 == int32(0) {
		goto L102
	} else {
		goto L110
	}
L107:
	;
	v1728 = v1680 << (uint(int32(2)) % 32)
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1728)))
	*(*int32)(unsafe.Add(mBase, uint32(v1668+v1728))) = v1731
	v1733 = int32(4)
	v1734 = v1728 | v1733
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1734)))
	*(*int32)(unsafe.Add(mBase, uint32(v1668+v1734))) = v1737
	v1740 = v1728 | int32(8)
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1740)))
	*(*int32)(unsafe.Add(mBase, uint32(v1668+v1740))) = v1743
	v1746 = v1728 | int32(12)
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1746)))
	*(*int32)(unsafe.Add(mBase, uint32(v1668+v1746))) = v1749
	v1752 = v1680 + v1733
	v1754 = v1718 + v1733
	if v1754 != l10&int32(2147483644) {
		v1680 = v1752
		v1718 = v1754
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v1756 = v1752
	goto L106
L109:
	;
	goto L108
L110:
	;
	v1805 = v1756
	v1827 = int32(0)
	goto L111
L111:
	;
	v1853 = v1805 << (uint(int32(2)) % 32)
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1853)))
	*(*int32)(unsafe.Add(mBase, uint32(v1668+v1853))) = v1856
	v1858 = int32(1)
	v1861 = v1827 + v1858
	if v1861 != v1673 {
		v1805 = v1805 + v1858
		v1827 = v1861
		goto L111
	} else {
		goto L113
	}
L112:
	;
	goto L102
L113:
	;
	goto L112
L114:
	;
	v1913 = v1911
	goto L100
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+172)) = l31
	*(*int32)(unsafe.Add(mBase, uint32(v50)+168)) = l30
	*(*int32)(unsafe.Add(mBase, uint32(v50)+164)) = l29
	*(*int32)(unsafe.Add(mBase, uint32(v50)+160)) = l28
	*(*int32)(unsafe.Add(mBase, uint32(v50)+156)) = l24
	*(*int32)(unsafe.Add(mBase, uint32(v50)+152)) = l21
	*(*int32)(unsafe.Add(mBase, uint32(v50)+148)) = l20
	*(*int32)(unsafe.Add(mBase, uint32(v50)+144)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v50)+140)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v50)+136)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v50)+132)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v50)+128)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v50)+124)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v50)+120)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v50)+116)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v50)+112)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v50)+108)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v50)+104)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v50)+96)) = v1978
	*(*int32)(unsafe.Add(mBase, uint32(v50)+100)) = v50 + int32(32)
	if v367 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v1664 != 0 {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+176)) = v367
	goto L116
L118:
	;
	goto L119
L119:
	;
	v2003 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+228)) = uint8(v2003)
	goto L116
L120:
	;
	if v1663 != 0 {
		goto L125
	} else {
		goto L126
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+180)) = v1664
	goto L120
L122:
	;
	goto L123
L123:
	;
	v2006 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+229)) = uint8(v2006)
	goto L120
L124:
	;
	if v1662 != 0 {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+184)) = v1663
	goto L124
L126:
	;
	goto L127
L127:
	;
	v2009 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+230)) = uint8(v2009)
	goto L124
L128:
	;
	if v1659 != 0 {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+188)) = v1662
	goto L128
L130:
	;
	goto L131
L131:
	;
	v2012 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+231)) = uint8(v2012)
	goto L128
L132:
	;
	if v1658 != 0 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+192)) = v1659
	goto L132
L134:
	;
	goto L135
L135:
	;
	v2015 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+232)) = uint8(v2015)
	goto L132
L136:
	;
	if v1913 != 0 {
		goto L141
	} else {
		goto L142
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+196)) = v1658
	goto L136
L138:
	;
	goto L139
L139:
	;
	v2018 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+233)) = uint8(v2018)
	goto L136
L140:
	;
	if l27 != 0 {
		goto L145
	} else {
		goto L146
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+200)) = v1913
	goto L140
L142:
	;
	goto L143
L143:
	;
	v2021 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+234)) = uint8(v2021)
	goto L140
L144:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v2033 = F_heap_form_tuple(m, v2028, v50+int32(96), v50+int32(208))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L1
	} else {
		goto L149
	}
L145:
	;
	v2023 = F_cstring_to_text(m, l27)
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v2026 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+235)) = uint8(v2026)
	goto L144
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+204)) = v2023
	goto L144
L149:
	;
	F_CatalogTupleInsert(m, v54, v2033)
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v1978
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = int32(2606)
	F_sequence_close(m, v54, int32(3))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v2045 = F_new_object_addresses(m)
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	if l8 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if l12 != 0 {
		goto L163
	} else {
		goto L164
	}
L154:
	;
	if int32(0) < l11 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v2052 = int32(0)
	goto L158
L156:
	;
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(1259)
	F_add_exact_object_address(m, v50+int32(8), v2045)
	mBase = m.M
	v2122 = m.ExcPending
	if v2122 != 0 {
		goto L1
	} else {
		goto L162
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(1259)
	v2105 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v2052<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v2105
	F_add_exact_object_address(m, v50+int32(8), v2045)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L1
	} else {
		goto L160
	}
L159:
	;
	goto L153
L160:
	;
	v2112 = v2052 + int32(1)
	if v2112 != l11 {
		v2052 = v2112
		goto L158
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	goto L153
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(1247)
	F_add_exact_object_address(m, v50+int32(8), v2045)
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	F_record_object_address_dependencies(m, v50+int32(20), v2045, int32(97))
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L1
	} else {
		goto L167
	}
L166:
	;
	goto L165
L167:
	;
	F_free_object_addresses(m, v2045)
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v2186 = F_new_object_addresses(m)
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	if l14 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	if l2 != int32(102) {
		goto L180
	} else {
		goto L181
	}
L171:
	;
	if int32(0) < l19 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v2193 = int32(0)
	goto L175
L173:
	;
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(1259)
	F_add_exact_object_address(m, v50+int32(8), v2186)
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L1
	} else {
		goto L179
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(1259)
	v2246 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v2193<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v2246
	F_add_exact_object_address(m, v50+int32(8), v2186)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L1
	} else {
		goto L177
	}
L176:
	;
	goto L170
L177:
	;
	v2253 = v2193 + int32(1)
	if v2253 != l19 {
		v2193 = v2253
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	goto L170
L180:
	;
	if int32(0) < l19 {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	if l13 == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(1259)
	F_add_exact_object_address(m, v50+int32(8), v2186)
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	goto L180
L184:
	;
	v2326 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v2326
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(2617)
	v2331 = v2326
	goto L187
L185:
	;
	goto L186
L186:
	;
	F_record_object_address_dependencies(m, v50+int32(20), v2186, int32(110))
	mBase = m.M
	v2460 = m.ExcPending
	if v2460 != 0 {
		goto L1
	} else {
		goto L199
	}
L187:
	;
	v2379 = v2331 << (uint(int32(2)) % 32)
	v2380 = l16 + v2379
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2380)))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v2381
	F_add_exact_object_address(m, v50+int32(8), v2186)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L1
	} else {
		goto L189
	}
L188:
	;
	goto L186
L189:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(l17+v2379)))
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v2380)))
	if v2388 != v2389 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v2388
	F_add_exact_object_address(m, v50+int32(8), v2186)
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L1
	} else {
		goto L193
	}
L191:
	;
	v2397 = v2388
	goto L192
L192:
	;
	v2399 = *(*int32)(unsafe.Add(mBase, uint32(l18+v2379)))
	if v2397 != v2399 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2380)))
	v2397 = v2396
	goto L192
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v2399
	F_add_exact_object_address(m, v50+int32(8), v2186)
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L1
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v2407 = v2331 + int32(1)
	if v2407 != l19 {
		v2331 = v2407
		goto L187
	} else {
		goto L198
	}
L197:
	;
	goto L196
L198:
	;
	goto L188
L199:
	;
	F_free_object_addresses(m, v2186)
	mBase = m.M
	v2462 = m.ExcPending
	if v2462 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	if l26 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	F_recordDependencyOnSingleRelExpr(m, v50+int32(20), l26, l8, int32(110), int32(0))
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L1
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	if v2470 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	goto L203
L205:
	;
	F_RunObjectPostCreateHook(m, int32(2606), v1978, int32(0), l32)
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	m.G0 = v50 + int32(240)
	return v1978
L208:
	;
	goto L207
}
