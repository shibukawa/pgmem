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
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
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
	var v249 int32
	_ = v249
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v364 int32
	_ = v364
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
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v466 int32
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
	var v557 int32
	_ = v557
	var v570 int32
	_ = v570
	var v573 int32
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
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v761 int32
	_ = v761
	var v799 int32
	_ = v799
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v912 int32
	_ = v912
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v999 int32
	_ = v999
	var v1037 int32
	_ = v1037
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1114 int32
	_ = v1114
	var v1150 int32
	_ = v1150
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1237 int32
	_ = v1237
	var v1275 int32
	_ = v1275
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1355 int32
	_ = v1355
	var v1393 int32
	_ = v1393
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1492 int32
	_ = v1492
	var v1528 int32
	_ = v1528
	var v1541 int32
	_ = v1541
	var v1544 int32
	_ = v1544
	var v1549 int32
	_ = v1549
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1661 int32
	_ = v1661
	var v1699 int32
	_ = v1699
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1738 int32
	_ = v1738
	var v1784 int32
	_ = v1784
	var v1806 int32
	_ = v1806
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1938 int64
	_ = v1938
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2028 int32
	_ = v2028
	var v2080 int32
	_ = v2080
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2097 int32
	_ = v2097
	var v2152 int32
	_ = v2152
	var v2157 int32
	_ = v2157
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2167 int32
	_ = v2167
	var v2219 int32
	_ = v2219
	var v2224 int32
	_ = v2224
	var v2226 int32
	_ = v2226
	var v2236 int32
	_ = v2236
	var v2283 int32
	_ = v2283
	var v2298 int32
	_ = v2298
	var v2301 int32
	_ = v2301
	var v2306 int32
	_ = v2306
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2358 int32
	_ = v2358
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2377 int32
	_ = v2377
	var v2379 int32
	_ = v2379
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2437 int32
	_ = v2437
	var v2439 int32
	_ = v2439
	var v2443 int32
	_ = v2443
	v34 = int32(0)
	v47 = m.G0
	v49 = v47 - int32(240)
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
	v66 = l10 & int32(3)
	v69 = F_palloc(m, l10<<(uint(int32(2))%32))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v364 = v34
	goto L6
L6:
	;
	if l19 <= int32(0) {
		v1635 = v34
		v1639 = v34
		v1641 = v34
		v1642 = v34
		v1645 = v34
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
	v114 = v34
	goto L12
L10:
	;
	v167 = v71
	goto L11
L11:
	;
	v213 = v167
	v249 = v34
	goto L16
L12:
	;
	v122 = int32(2)
	v125 = int32(1)
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v76<<(uint(v125)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v69+v76<<(uint(v122)%32)))) = v128
	v131 = v76 | v125
	v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v131<<(uint(v125)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v69+v131<<(uint(v122)%32)))) = v138
	v141 = v76 | v122
	v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v141<<(uint(v125)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v69+v141<<(uint(v122)%32)))) = v148
	v151 = v76 | int32(3)
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v151<<(uint(v125)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v69+v151<<(uint(v122)%32)))) = v158
	v160 = int32(4)
	v161 = v76 + v160
	v163 = v114 + v160
	if v163 != l10&int32(2147483644) {
		v76 = v161
		v114 = v163
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
	v265 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v213<<(uint(v262)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v69+v213<<(uint(int32(2))%32)))) = v265
	v270 = v249 + v262
	if v270 != v66 {
		v213 = v213 + v262
		v249 = v270
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
	v364 = v319
	goto L6
L20:
	;
	v1646 = int32(0)
	if l25 != 0 {
		goto L89
	} else {
		goto L90
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
	v376 = F_palloc(m, v373<<(uint(int32(2))%32))
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
	v422 = int32(0)
	goto L30
L28:
	;
	v475 = v378
	goto L29
L29:
	;
	v521 = v475
	v557 = int32(0)
	goto L34
L30:
	;
	v430 = int32(2)
	v433 = int32(1)
	v436 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v384<<(uint(v433)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v384<<(uint(v430)%32)))) = v436
	v439 = v384 | v433
	v446 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v439<<(uint(v433)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v439<<(uint(v430)%32)))) = v446
	v449 = v384 | v430
	v456 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v449<<(uint(v433)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v449<<(uint(v430)%32)))) = v456
	v459 = v384 | int32(3)
	v466 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v459<<(uint(v433)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v459<<(uint(v430)%32)))) = v466
	v468 = int32(4)
	v469 = v384 + v468
	v471 = v422 + v468
	if v471 != l19&int32(2147483644) {
		v384 = v469
		v422 = v471
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
	v573 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v521<<(uint(v570)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v521<<(uint(int32(2))%32)))) = v573
	v578 = v557 + v570
	if v578 != v370 {
		v521 = v521 + v570
		v557 = v578
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
	v865 = l19 & int32(3)
	v868 = F_construct_array_builtin(m, v376, l19, int32(26))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
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
	v715 = v632
	goto L41
L41:
	;
	v761 = v715
	v799 = int32(0)
	goto L46
L42:
	;
	v685 = v638 << (uint(int32(2)) % 32)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l16+v685)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v685))) = v688
	v690 = int32(4)
	v691 = v685 | v690
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l16+v691)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v691))) = v694
	v697 = v685 | int32(8)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l16+v697)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v697))) = v700
	v703 = v685 | int32(12)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l16+v703)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v703))) = v706
	v709 = v638 + v690
	v711 = v674 + v690
	if v711 != l19&int32(2147483644) {
		v638 = v709
		v674 = v711
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
	v715 = v709
	goto L41
L46:
	;
	v808 = v761 << (uint(int32(2)) % 32)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(l16+v808)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v808))) = v811
	v813 = int32(1)
	v816 = v799 + v813
	if v816 != v627 {
		v761 = v761 + v813
		v799 = v816
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
	v870 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l19) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v1103 = l19 & int32(3)
	v1106 = F_construct_array_builtin(m, v376, l19, int32(26))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L61
	}
L51:
	;
	v876 = v870
	v912 = int32(0)
	goto L54
L52:
	;
	v953 = v870
	goto L53
L53:
	;
	v999 = v953
	v1037 = int32(0)
	goto L58
L54:
	;
	v923 = v876 << (uint(int32(2)) % 32)
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l17+v923)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v923))) = v926
	v928 = int32(4)
	v929 = v923 | v928
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l17+v929)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v929))) = v932
	v935 = v923 | int32(8)
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l17+v935)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v935))) = v938
	v941 = v923 | int32(12)
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l17+v941)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v941))) = v944
	v947 = v876 + v928
	v949 = v912 + v928
	if v949 != l19&int32(2147483644) {
		v876 = v947
		v912 = v949
		goto L54
	} else {
		goto L56
	}
L55:
	;
	if v865 == int32(0) {
		goto L50
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v953 = v947
	goto L53
L58:
	;
	v1046 = v999 << (uint(int32(2)) % 32)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l17+v1046)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1046))) = v1049
	v1051 = int32(1)
	v1054 = v1037 + v1051
	if v1054 != v865 {
		v999 = v999 + v1051
		v1037 = v1054
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
	v1108 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l19) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v1341 = F_construct_array_builtin(m, v376, l19, int32(26))
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L1
	} else {
		goto L73
	}
L63:
	;
	v1114 = v1108
	v1150 = int32(0)
	goto L66
L64:
	;
	v1191 = v1108
	goto L65
L65:
	;
	v1237 = v1191
	v1275 = int32(0)
	goto L70
L66:
	;
	v1161 = v1114 << (uint(int32(2)) % 32)
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1161)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1161))) = v1164
	v1166 = int32(4)
	v1167 = v1161 | v1166
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1167)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1167))) = v1170
	v1173 = v1161 | int32(8)
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1173)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1173))) = v1176
	v1179 = v1161 | int32(12)
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1179)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1179))) = v1182
	v1185 = v1114 + v1166
	v1187 = v1150 + v1166
	if v1187 != l19&int32(2147483644) {
		v1114 = v1185
		v1150 = v1187
		goto L66
	} else {
		goto L68
	}
L67:
	;
	if v1103 == int32(0) {
		goto L62
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	v1191 = v1185
	goto L65
L70:
	;
	v1284 = v1237 << (uint(int32(2)) % 32)
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(l18+v1284)))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1284))) = v1287
	v1289 = int32(1)
	v1292 = v1275 + v1289
	if v1292 != v1103 {
		v1237 = v1237 + v1289
		v1275 = v1292
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
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v1635 = int32(0)
	v1639 = v1341
	v1641 = v630
	v1642 = v1106
	v1645 = v868
	goto L20
L75:
	;
	goto L76
L76:
	;
	v1347 = l23 & int32(3)
	v1348 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l23) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v1598 = F_construct_array_builtin(m, v376, l23, int32(21))
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L1
	} else {
		goto L88
	}
L78:
	;
	v1355 = v1348
	v1393 = int32(0)
	goto L81
L79:
	;
	v1446 = v1348
	goto L80
L80:
	;
	v1492 = v1446
	v1528 = v1348
	goto L85
L81:
	;
	v1401 = int32(2)
	v1404 = int32(1)
	v1407 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1355<<(uint(v1404)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1355<<(uint(v1401)%32)))) = v1407
	v1410 = v1355 | v1404
	v1417 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1410<<(uint(v1404)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1410<<(uint(v1401)%32)))) = v1417
	v1420 = v1355 | v1401
	v1427 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1420<<(uint(v1404)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1420<<(uint(v1401)%32)))) = v1427
	v1430 = v1355 | int32(3)
	v1437 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1430<<(uint(v1404)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1430<<(uint(v1401)%32)))) = v1437
	v1439 = int32(4)
	v1440 = v1355 + v1439
	v1442 = v1393 + v1439
	if v1442 != l23&int32(2147483644) {
		v1355 = v1440
		v1393 = v1442
		goto L81
	} else {
		goto L83
	}
L82:
	;
	if v1347 == int32(0) {
		goto L77
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v1446 = v1440
	goto L80
L85:
	;
	v1541 = int32(1)
	v1544 = int32(*(*int16)(unsafe.Add(mBase, uint32(l22+v1492<<(uint(v1541)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v376+v1492<<(uint(int32(2))%32)))) = v1544
	v1549 = v1528 + v1541
	if v1549 != v1347 {
		v1492 = v1492 + v1541
		v1528 = v1549
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L77
L87:
	;
	goto L86
L88:
	;
	v1635 = v1598
	v1639 = v1341
	v1641 = v630
	v1642 = v1106
	v1645 = v868
	goto L20
L89:
	;
	v1649 = F_palloc(m, l10<<(uint(int32(2))%32))
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	v1890 = v1646
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+232)) = int32(0)
	v1938 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+224)) = v1938
	*(*int64)(unsafe.Add(mBase, uint32(v49)+216)) = v1938
	*(*int64)(unsafe.Add(mBase, uint32(v49)+208)) = v1938
	*(*int64)(unsafe.Add(mBase, uint32(v49)+176)) = v1938
	*(*int64)(unsafe.Add(mBase, uint32(v49)+184)) = v1938
	*(*int64)(unsafe.Add(mBase, uint32(v49)+192)) = v1938
	*(*int64)(unsafe.Add(mBase, uint32(v49)+200)) = v1938
	v1954 = F_GetNewOidWithIndex(m, v53, int32(2667), int32(1))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L1
	} else {
		goto L106
	}
L92:
	;
	if l10 <= int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v1888 = F_construct_array_builtin(m, v1649, l10, int32(26))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L1
	} else {
		goto L105
	}
L94:
	;
	v1654 = l10 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(l10) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v1661 = v1646
	v1699 = int32(0)
	goto L98
L96:
	;
	v1738 = v1646
	goto L97
L97:
	;
	v1784 = v1738
	v1806 = int32(0)
	goto L102
L98:
	;
	v1708 = v1661 << (uint(int32(2)) % 32)
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1708)))
	*(*int32)(unsafe.Add(mBase, uint32(v1649+v1708))) = v1711
	v1713 = int32(4)
	v1714 = v1708 | v1713
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1714)))
	*(*int32)(unsafe.Add(mBase, uint32(v1649+v1714))) = v1717
	v1720 = v1708 | int32(8)
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1720)))
	*(*int32)(unsafe.Add(mBase, uint32(v1649+v1720))) = v1723
	v1726 = v1708 | int32(12)
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1726)))
	*(*int32)(unsafe.Add(mBase, uint32(v1649+v1726))) = v1729
	v1732 = v1661 + v1713
	v1734 = v1699 + v1713
	if v1734 != l10&int32(2147483644) {
		v1661 = v1732
		v1699 = v1734
		goto L98
	} else {
		goto L100
	}
L99:
	;
	if v1654 == int32(0) {
		goto L93
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	v1738 = v1732
	goto L97
L102:
	;
	v1831 = v1784 << (uint(int32(2)) % 32)
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l25+v1831)))
	*(*int32)(unsafe.Add(mBase, uint32(v1649+v1831))) = v1834
	v1836 = int32(1)
	v1839 = v1806 + v1836
	if v1839 != v1654 {
		v1784 = v1784 + v1836
		v1806 = v1839
		goto L102
	} else {
		goto L104
	}
L103:
	;
	goto L93
L104:
	;
	goto L103
L105:
	;
	v1890 = v1888
	goto L91
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+172)) = l31
	*(*int32)(unsafe.Add(mBase, uint32(v49)+168)) = l30
	*(*int32)(unsafe.Add(mBase, uint32(v49)+164)) = l29
	*(*int32)(unsafe.Add(mBase, uint32(v49)+160)) = l28
	*(*int32)(unsafe.Add(mBase, uint32(v49)+156)) = l24
	*(*int32)(unsafe.Add(mBase, uint32(v49)+152)) = l21
	*(*int32)(unsafe.Add(mBase, uint32(v49)+148)) = l20
	*(*int32)(unsafe.Add(mBase, uint32(v49)+144)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v49)+140)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v49)+136)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v49)+132)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v49)+128)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v49)+124)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v49)+120)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v49)+116)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v49)+112)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v49)+108)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v49)+104)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v49)+96)) = v1954
	*(*int32)(unsafe.Add(mBase, uint32(v49)+100)) = v49 + int32(32)
	if v364 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v1641 != 0 {
		goto L112
	} else {
		goto L113
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+176)) = v364
	goto L107
L109:
	;
	goto L110
L110:
	;
	v1979 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+228)) = uint8(v1979)
	goto L107
L111:
	;
	if v1645 != 0 {
		goto L116
	} else {
		goto L117
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+180)) = v1641
	goto L111
L113:
	;
	goto L114
L114:
	;
	v1982 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+229)) = uint8(v1982)
	goto L111
L115:
	;
	if v1642 != 0 {
		goto L120
	} else {
		goto L121
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+184)) = v1645
	goto L115
L117:
	;
	goto L118
L118:
	;
	v1985 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+230)) = uint8(v1985)
	goto L115
L119:
	;
	if v1639 != 0 {
		goto L124
	} else {
		goto L125
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+188)) = v1642
	goto L119
L121:
	;
	goto L122
L122:
	;
	v1988 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+231)) = uint8(v1988)
	goto L119
L123:
	;
	if v1635 != 0 {
		goto L128
	} else {
		goto L129
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+192)) = v1639
	goto L123
L125:
	;
	goto L126
L126:
	;
	v1991 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+232)) = uint8(v1991)
	goto L123
L127:
	;
	if v1890 != 0 {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+196)) = v1635
	goto L127
L129:
	;
	goto L130
L130:
	;
	v1994 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+233)) = uint8(v1994)
	goto L127
L131:
	;
	if l27 != 0 {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+200)) = v1890
	goto L131
L133:
	;
	goto L134
L134:
	;
	v1997 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+234)) = uint8(v1997)
	goto L131
L135:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	v2009 = F_heap_form_tuple(m, v2004, v49+int32(96), v49+int32(208))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L1
	} else {
		goto L140
	}
L136:
	;
	v1999 = F_cstring_to_text(m, l27)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v2002 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+235)) = uint8(v2002)
	goto L135
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+204)) = v1999
	goto L135
L140:
	;
	F_CatalogTupleInsert(m, v53, v2009)
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v1954
	*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = int32(2606)
	F_relation_close(m, v53, int32(3))
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v2021 = F_new_object_addresses(m)
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	if l8 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	if l12 != 0 {
		goto L154
	} else {
		goto L155
	}
L145:
	;
	if int32(0) < l11 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v2028 = int32(0)
	goto L149
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	F_add_exact_object_address(m, v49+int32(8), v2021)
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L1
	} else {
		goto L153
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	v2080 = int32(*(*int16)(unsafe.Add(mBase, uint32(l9+v2028<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v2080
	F_add_exact_object_address(m, v49+int32(8), v2021)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L1
	} else {
		goto L151
	}
L150:
	;
	goto L144
L151:
	;
	v2087 = v2028 + int32(1)
	if v2087 != l11 {
		v2028 = v2087
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	goto L144
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1247)
	F_add_exact_object_address(m, v49+int32(8), v2021)
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	F_record_object_address_dependencies(m, v49+int32(20), v2021, int32(97))
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L1
	} else {
		goto L158
	}
L157:
	;
	goto L156
L158:
	;
	F_free_object_addresses(m, v2021)
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v2160 = F_new_object_addresses(m)
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	if l14 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v2283 = int32(0)
	if base.B2i32(l13 == v2283)|base.B2i32(l2 != int32(102)) == v2283 {
		goto L171
	} else {
		goto L172
	}
L162:
	;
	if int32(0) < l19 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v2167 = int32(0)
	goto L166
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	F_add_exact_object_address(m, v49+int32(8), v2160)
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L1
	} else {
		goto L170
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	v2219 = int32(*(*int16)(unsafe.Add(mBase, uint32(l15+v2167<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v2219
	F_add_exact_object_address(m, v49+int32(8), v2160)
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L1
	} else {
		goto L168
	}
L167:
	;
	goto L161
L168:
	;
	v2226 = v2167 + int32(1)
	if v2226 != l19 {
		v2167 = v2226
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	goto L161
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(1259)
	F_add_exact_object_address(m, v49+int32(8), v2160)
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L1
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	if int32(0) < l19 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L173
L175:
	;
	v2301 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v2301
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(2617)
	v2306 = v2301
	goto L178
L176:
	;
	goto L177
L177:
	;
	v2428 = v49 + int32(20)
	F_record_object_address_dependencies(m, v2428, v2160, int32(110))
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L1
	} else {
		goto L190
	}
L178:
	;
	v2353 = v2306 << (uint(int32(2)) % 32)
	v2354 = l16 + v2353
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(v2354)))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v2355
	v2358 = v49 + int32(8)
	F_add_exact_object_address(m, v2358, v2160)
	mBase = m.M
	v2360 = m.ExcPending
	if v2360 != 0 {
		goto L1
	} else {
		goto L180
	}
L179:
	;
	goto L177
L180:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v2353+l17)))
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v2354)))
	if v2362 != v2363 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v2362
	F_add_exact_object_address(m, v2358, v2160)
	mBase = m.M
	v2367 = m.ExcPending
	if v2367 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	v2369 = v2362
	goto L183
L183:
	;
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2353+l18)))
	if v2369 != v2371 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v2354)))
	v2369 = v2368
	goto L183
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v2371
	F_add_exact_object_address(m, v49+int32(8), v2160)
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L1
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v2379 = v2306 + int32(1)
	if v2379 != l19 {
		v2306 = v2379
		goto L178
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	goto L179
L190:
	;
	F_free_object_addresses(m, v2160)
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	if l26 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	F_recordDependencyOnSingleRelExpr(m, v2428, l26, l8, int32(110), int32(0))
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		goto L1
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, _c_F_CreateConstraintEntry[0]))
	if v2439 != 0 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	goto L194
L196:
	;
	F_RunObjectPostCreateHook(m, int32(2606), v1954, int32(0), l32)
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L1
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	m.G0 = v49 + int32(240)
	return v1954
L199:
	;
	goto L198
}
