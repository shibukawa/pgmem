package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vacuum(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v146 int32
	_ = v146
	var v162 int32
	_ = v162
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v352 int32
	_ = v352
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v410 int32
	_ = v410
	var v428 int32
	_ = v428
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v463 int32
	_ = v463
	var v482 int32
	_ = v482
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v613 int32
	_ = v613
	var v631 int32
	_ = v631
	var v647 int32
	_ = v647
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v802 int32
	_ = v802
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v952 int32
	_ = v952
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1075 int32
	_ = v1075
	var v1091 int32
	_ = v1091
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1157 int32
	_ = v1157
	var v1174 int32
	_ = v1174
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1229 int32
	_ = v1229
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1293 int32
	_ = v1293
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int64
	_ = v1311
	var v1313 int64
	_ = v1313
	var v1315 int64
	_ = v1315
	var v1317 int64
	_ = v1317
	var v1319 int64
	_ = v1319
	var v1321 int64
	_ = v1321
	var v1323 int64
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1369 int32
	_ = v1369
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1498 int32
	_ = v1498
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1566 int32
	_ = v1566
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1658 int32
	_ = v1658
	var v1667 int32
	_ = v1667
	var v1684 int32
	_ = v1684
	var v1699 int32
	_ = v1699
	var v1714 int32
	_ = v1714
	var v1729 int32
	_ = v1729
	var v1738 int32
	_ = v1738
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1786 int32
	_ = v1786
	var v1797 int32
	_ = v1797
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1834 int32
	_ = v1834
	var v1843 int32
	_ = v1843
	var v1868 int32
	_ = v1868
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1900 int64
	_ = v1900
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	v6 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(128)
	m.G0 = v33
	v42 = v6
	v43 = v6
	v44 = v6
	v45 = v6
	v46 = v6
	v47 = v6
	v48 = v6
	v49 = v6
	v50 = v6
	v51 = v6
	v52 = v6
	v53 = v6
	v54 = int32(-1)
	v55 = v6
	v56 = v6
	v65 = v33
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
	if v54 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v1899 = int32(m.ExcTag)
	v1900 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1899 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L7:
	;
	if v1220 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L8:
	;
	v1206 = v42
	v1207 = v43
	v1208 = v44
	v1209 = v45
	v1210 = v46
	v1211 = v47
	v1212 = v48
	v1213 = v49
	v1214 = v50
	v1215 = v51
	v1216 = v52
	v1217 = v53
	v1218 = v55
	v1220 = v56
	v1229 = v65
	goto L7
L9:
	;
	goto L10
L10:
	;
	v68 = int32(16)
	v69 = v65 - v68
	m.G0 = v69
	v72 = v69 - v68
	m.G0 = v72
	v75 = v72 - int32(160)
	m.G0 = v75
	v78 = v75 + int32(-64)
	m.G0 = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v82 = v80 & int32(1)
	if v82 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v126)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _consts[406])))
	if v129 != 0 {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_PreventInTransactionBlock(m, l4, int32(530487))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v114 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+24))
	if base.Ui32(v115) <= base.Ui32(int32(1)) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v126 = int32(0)
	goto L11
L16:
	;
	v118 = int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)+28))
	v125 = l4 ^ v118 | base.B2i32(v118 < v120)
	goto L18
L17:
	;
	v125 = int32(1)
	goto L18
L18:
	;
	v126 = v125
	goto L11
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v201&int32(1024) != 0 {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errcode(m, int32(1088))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	if v82 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v178 = int32(530487)
	goto L26
L25:
	;
	v178 = int32(536605)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v178
	F_errmsg(m, int32(536548), v33)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errfinish(m, int32(496439), int32(538), int32(285751))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	v1122 = int32(1)
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v1123&v1122 != 0 {
		v1138 = v1122
		goto L119
	} else {
		goto L120
	}
L30:
	;
	v1104 = v48
	v1105 = v49
	v1106 = v50
	v1107 = v51
	v1108 = v52
	v1109 = v53
	v1110 = l0
	goto L29
L31:
	;
	goto L32
L32:
	;
	if l0 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v204 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v873 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L100
	}
L36:
	;
	v1104 = v48
	v1105 = v49
	v1106 = v50
	v1107 = v51
	v1108 = v52
	v1109 = v53
	v1110 = int32(0)
	goto L29
L37:
	;
	goto L38
L38:
	;
	v208 = int32(0)
	v222 = v48
	v225 = v51
	v226 = v52
	v227 = v53
	v231 = v208
	v236 = v208
	goto L39
L39:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240+v236<<(uint(int32(2))%32))))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	if v245 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v1104 = v850
	v1105 = v49
	v1106 = v50
	v1107 = v818
	v1108 = v819
	v1109 = v820
	v1110 = v850
	goto L29
L41:
	;
	v833 = int32(4489440)
	v834 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v819
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v820
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v850 = F_list_concat(m, v231, v823)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L98
	}
L42:
	;
	v246 = int32(4489440)
	v247 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v264 = F_lappend(m, int32(0), v244)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v288 = int32(0)
	v290 = F_RangeVarGetRelidExtended(m, v269, int32(1), int32(base.Ui32(v268)>>(uint(int32(3))%32))&int32(4), v288, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v247
	v818 = v225
	v819 = v226
	v820 = v264
	v823 = v264
	goto L41
L46:
	;
	if v290 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v309 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v443 = F_SearchSysCache1(m, int32(57), v290)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L62
	}
L50:
	;
	if v268&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v313 = int32(0)
	if v309 == v313 {
		v818 = v225
		v819 = v226
		v820 = v227
		v823 = v313
		goto L41
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v371 = int32(0)
	if v309 == v371 {
		v818 = v225
		v819 = v226
		v820 = v227
		v823 = v371
		goto L41
	} else {
		goto L58
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errcode(m, int32(50463045))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v333
	F_errmsg(m, int32(395136), v33+int32(32))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errfinish(m, int32(496439), int32(952), int32(307195))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v818 = v225
	v819 = v226
	v820 = v227
	v823 = v313
	goto L41
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errcode(m, int32(50463045))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v391
	F_errmsg(m, int32(395183), v33+int32(16))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errfinish(m, int32(496439), int32(957), int32(307195))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v818 = v225
	v819 = v226
	v820 = v227
	v823 = v371
	goto L41
L62:
	;
	if v443 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v443)+16))
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v516 = v501 + v502
	v517 = F_vacuum_is_permitted_for_relation(m, v290, v516, v268)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L69
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = v290
	F_errmsg_internal(m, int32(46249), v33+int32(48))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errfinish(m, int32(496439), int32(967), int32(307195))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L68
	}
L68:
	;
	goto L3
L69:
	;
	if v517 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v520 = int32(4489440)
	v521 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = l3
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v539 = F_makeVacuumRelation(m, v525, v290, v524)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L73
	}
L71:
	;
	v559 = v226
	v560 = int32(0)
	goto L72
L72:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+16)))
	if v268&int32(1) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v555 = F_lappend(m, int32(0), v539)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v521
	v559 = v555
	v560 = v555
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_ReleaseCatCache(m, v443)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L83
	}
L76:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+119)))
	if v569 != int32(112) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	if v564&int32(1) != 0 {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v589 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L79
	}
L79:
	;
	if v589 == int32(0) {
		goto L75
	} else {
		goto L80
	}
L80:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = v594
	F_errmsg(m, int32(111441), v33-int32(-64))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_errfinish(m, int32(496439), int32(993), int32(307195))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L82
	}
L82:
	;
	goto L75
L83:
	;
	if v564&int32(1) == int32(0) {
		v772 = v225
		v777 = v560
		goto L84
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_UnlockRelationOid(m, v290, int32(1))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L97
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v665 = int32(0)
	v667 = F_find_all_inheritors(m, v290, v665, v665)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L86
	}
L86:
	;
	if v667 == int32(0) {
		v772 = v225
		v777 = v560
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v671 = int32(0)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	if v672 <= v671 {
		v772 = v225
		v777 = v560
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v690 = v225
	v693 = v671
	v695 = v560
	v700 = v672
	goto L89
L89:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v667)+12))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v705+v693<<(uint(int32(2))%32))))
	if v290 != v709 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v772 = v750
	v777 = v751
	goto L84
L91:
	;
	v711 = int32(4489440)
	v712 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = l3
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v690
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v730 = F_makeVacuumRelation(m, int32(0), v709, v715)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L94
	}
L92:
	;
	v750 = v690
	v751 = v695
	v752 = v700
	goto L93
L93:
	;
	v755 = v693 + int32(1)
	if v755 < v752 {
		v690 = v750
		v693 = v755
		v695 = v751
		v700 = v752
		goto L89
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v690
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v559
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v745 = F_lappend(m, v695, v730)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v712
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	v750 = v745
	v751 = v745
	v752 = v749
	goto L93
L96:
	;
	goto L90
L97:
	;
	v818 = v772
	v819 = v559
	v820 = v227
	v823 = v777
	goto L41
L98:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v834
	v855 = v236 + int32(1)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v855 < v856 {
		v222 = v850
		v225 = v818
		v226 = v819
		v227 = v820
		v231 = v850
		v236 = v855
		goto L39
	} else {
		goto L99
	}
L99:
	;
	goto L40
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v888 = int32(0)
	v890 = F_table_beginscan_catalog(m, v873, v888, v888)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v905 = F_heap_getnext(m, v890)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L102
	}
L102:
	;
	v907 = int32(0)
	if v905 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v921 = v49
	v922 = v50
	v926 = v907
	v928 = v905
	goto L106
L104:
	;
	v1041 = v49
	v1042 = v50
	v1046 = v907
	goto L105
L105:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+188))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1041
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	m.T0[v1060].(func(*base.Module, int32))(m, v890)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L117
	}
L106:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v928)+16))
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938)+22)))
	v940 = v938 + v939
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v940)+119)))
	v943 = v941 - int32(109)
	if base.Ui32(int32(5)) < base.Ui32(v943) {
		v1008 = v922
		v1009 = v926
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v1041 = v1026
	v1042 = v1008
	v1046 = v1009
	goto L105
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v921
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1008
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v1026 = F_heap_getnext(m, v890)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L115
	}
L109:
	;
	if int32(1)<<(uint(v943)%32)&int32(41) == int32(0) {
		v1008 = v922
		v1009 = v926
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v940)))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v921
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v966 = F_vacuum_is_permitted_for_relation(m, v952, v940, v201)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L111
	}
L111:
	;
	if v966 == int32(0) {
		v1008 = v922
		v1009 = v926
		goto L108
	} else {
		goto L112
	}
L112:
	;
	v970 = int32(4489440)
	v971 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v921
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v987 = int32(0)
	v989 = F_makeVacuumRelation(m, v987, v952, v987)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v921
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v1004 = F_lappend(m, v926, v989)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v971
	v1008 = v1004
	v1009 = v1004
	goto L108
L115:
	;
	if v1026 != 0 {
		v921 = v1026
		v922 = v1008
		v926 = v1009
		v928 = v1026
		goto L106
	} else {
		goto L116
	}
L116:
	;
	goto L107
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1041
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_sequence_close(m, v873, int32(1))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L118
	}
L118:
	;
	v1104 = v48
	v1105 = v1041
	v1106 = v1042
	v1107 = v51
	v1108 = v52
	v1109 = v53
	v1110 = v1046
	goto L29
L119:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v1138)
	v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v1140 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L120:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, _consts[407]))
	if v1127 == int32(4) {
		v1138 = v1122
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v1131 != 0 {
		v1138 = int32(0)
		goto L119
	} else {
		goto L122
	}
L122:
	;
	if v1110 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v1132 = int32(1)
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+4))
	if v1132 < v1133 {
		v1138 = v1132
		goto L119
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v1138 = int32(0)
	goto L119
L126:
	;
	goto L125
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1110
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1109
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	v1157 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	goto L130
L128:
	;
	goto L129
L129:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	v1194 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	goto L136
L130:
	;
	if v1157 != int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1110
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1109
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1110
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1107
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1109
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v69
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		v1898 = v78
		goto L6
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	goto L129
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v33 + int32(72)
	goto L139
L137:
	;
	v1206 = v78
	v1207 = v72
	v1208 = v69
	v1209 = v75
	v1210 = v1194
	v1211 = v1192
	v1212 = v1104
	v1213 = v1105
	v1214 = v1106
	v1215 = v1107
	v1216 = v1108
	v1217 = v1109
	v1218 = v1110
	v1220 = int32(0)
	v1229 = v78
	goto L7
L139:
	;
	goto L137
L140:
	;
	v1233 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[406])) = uint8(v1233)
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1209
	v1238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[408])) = uint8(v1238)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_VacuumUpdateCosts(m)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1211
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1210
	v1843 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[406])) = uint8(v1843)
	*(*uint8)(unsafe.Add(mBase, _consts[409])) = uint8(v1843)
	*(*uint8)(unsafe.Add(mBase, _consts[408])) = uint8(v1843)
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v1843
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_pg_re_throw(m)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L250
	}
L143:
	;
	v1256 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[411])) = v1256
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v1256
	*(*int32)(unsafe.Add(mBase, _consts[412])) = v1256
	*(*int32)(unsafe.Add(mBase, _consts[413])) = v1256
	if v1218 == v1256 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v1210
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1211
	v1786 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[409])) = uint8(v1786)
	*(*uint8)(unsafe.Add(mBase, _consts[406])) = uint8(v1786)
	*(*uint8)(unsafe.Add(mBase, _consts[408])) = uint8(v1786)
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v1786
	v1797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1207))))
	if v1797 == int32(1) {
		goto L242
	} else {
		goto L243
	}
L145:
	;
	v1269 = int32(0)
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+4))
	if v1270 <= v1269 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v1293 = v1269
	goto L147
L147:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+12))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1303+v1293<<(uint(int32(2))%32))))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1308&int32(1) != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	goto L144
L149:
	;
	v1748 = v1293 + int32(1)
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+4))
	if v1748 < v1749 {
		v1293 = v1748
		goto L147
	} else {
		goto L241
	}
L150:
	;
	v1738 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[408])) = uint8(v1738)
	goto L149
L151:
	;
	v1311 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v1206))) = v1311
	v1313 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1206)+48)) = v1313
	v1315 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1206)+40)) = v1315
	v1317 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1206)+32)) = v1317
	v1319 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1206)+24)) = v1319
	v1321 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1206)+16)) = v1321
	v1323 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1206)+8)) = v1323
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+8))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	v1340 = F_vacuum_rel(m, v1325, v1326, v1206, l2)
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L154
	}
L152:
	;
	v1347 = v1308
	goto L153
L153:
	;
	if v1347&int32(2) == int32(0) {
		goto L150
	} else {
		goto L156
	}
L154:
	;
	if v1340 == int32(0) {
		goto L149
	} else {
		goto L155
	}
L155:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1347 = v1344
	goto L153
L156:
	;
	v1352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1207))))
	if v1352 == int32(1) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_StartTransactionCommand(m)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+12))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+4))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1307)+8))
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1208))))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	v1418 = m.G0
	v1420 = v1418 - int32(32)
	m.G0 = v1420
	v1422 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+28)) = v1422
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+24)) = v1422
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _consts[352])) = l2
	v1430 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v1430 != 0 {
		goto L164
	} else {
		goto L165
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	v1383 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_PushActiveSnapshot(m, v1383)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L162
	}
L162:
	;
	goto L159
L163:
	;
	m.G0 = v1420 + int32(32)
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1207))))
	if v1667 == int32(1) {
		goto L234
	} else {
		goto L235
	}
L164:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L167
	}
L165:
	;
	v1434 = v1426
	goto L166
L166:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1443 = F_vacuum_open_relation(m, v1403, v1402, v1434&int32(-2), int32(base.Ui32(v1437^int32(-1))>>(uint(int32(31))%32)), int32(4))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L168
	}
L167:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1434 = v1433
	goto L166
L168:
	;
	if v1443 == int32(0) {
		goto L163
	} else {
		goto L169
	}
L169:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+56))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+48))
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v1452 = F_vacuum_is_permitted_for_relation(m, v1447, v1448, v1449&int32(-2))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L170
	}
L170:
	;
	if v1452 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	F_relation_close(m, v1443, int32(4))
	mBase = m.M
	v1458 = m.ExcPending
	if v1458 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+48))
	v1460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1459)+118)))
	if v1460 != int32(116) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L163
L175:
	;
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+56))
	if v1467 == int32(2619) {
		goto L179
	} else {
		goto L180
	}
L176:
	;
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1443)+24)))
	if v1463 != 0 {
		goto L175
	} else {
		goto L177
	}
L177:
	;
	F_relation_close(m, v1443, int32(4))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L178
	}
L178:
	;
	goto L163
L179:
	;
	F_relation_close(m, v1443, int32(4))
	mBase = m.M
	v1472 = m.ExcPending
	if v1472 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v1473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1459)+119)))
	switch v1473 - int32(102) {
	case 0:
		goto L186
	default:
		goto L185
	case 7, 12:
		goto L184
	case 10:
		goto L183
	}
L182:
	;
	goto L163
L183:
	;
	v1541 = v1426&int32(4) + int32(13)
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+56))
	v1546 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v1546 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+28)) = int32(501)
	v1534 = F_RelationGetNumberOfBlocksInFork(m, v1443, int32(0))
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L207
	}
L185:
	;
	v1507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v1507&int32(1) != 0 {
		goto L200
	} else {
		goto L201
	}
L186:
	;
	v1477 = F_GetFdwRoutineForRelation(m, v1443, int32(0))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L187
	}
L187:
	;
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+128))
	if v1479 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v1484 = m.T0[v1479].(func(*base.Module, int32, int32, int32) int32)(m, v1443, v1420+int32(28), v1420+int32(24))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1488 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L193
	}
L191:
	;
	if v1484 != 0 {
		goto L183
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	if v1488 != 0 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+16)) = v1490 + int32(4)
	F_errmsg(m, int32(392510), v1420+int32(16))
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	F_relation_close(m, v1443, int32(4))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L199
	}
L197:
	;
	F_errfinish(m, int32(497756), int32(216), int32(307235))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	goto L163
L200:
	;
	F_relation_close(m, v1443, int32(4))
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L206
	}
L201:
	;
	v1512 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L202
	}
L202:
	;
	if v1512 == int32(0) {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1420))) = v1516 + int32(4)
	F_errmsg(m, int32(165751), v1420)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(497756), int32(233), int32(307235))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L205
	}
L205:
	;
	goto L200
L206:
	;
	goto L163
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+24)) = v1534
	goto L183
L208:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+48))
	v1605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1604)+119)))
	if v1605 != int32(112) {
		goto L220
	} else {
		goto L221
	}
L209:
	;
	goto L208
L210:
	;
	v1550 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v1550 != int32(1) {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v1553 = int32(4484100)
	v1555 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v1556 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1555 + v1556
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1546)))
	*(*int32)(unsafe.Add(mBase, uint32(v1546))) = v1559 + v1556
	*(*int32)(unsafe.Add(mBase, uint32(v1546)+220)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1546)+224)) = v1543
	v1566 = v1546 + int32(232)
	if v1566&int32(3) == int32(0) {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1546)))
	v1593 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1546))) = v1592 + v1593
	v1596 = int32(4484100)
	v1598 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1598 - v1593
	goto L209
L213:
	;
	v1572 = v1546 + int32(392)
	if base.Ui32(v1572) <= base.Ui32(v1566) {
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1589 = F___memset(m, v1566, int32(0), int32(160))
	mBase = m.M
	goto L212
L216:
	;
	v1576 = v1546 + int32(236)
	if base.Ui32(v1576) < base.Ui32(v1572) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1578 = v1572
	goto L219
L218:
	;
	v1578 = v1576
	goto L219
L219:
	;
	v1586 = F___memset(m, v1566, int32(0), (v1578-v1546-int32(233))&int32(-4)+int32(4))
	mBase = m.M
	goto L212
L220:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+28))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+24))
	F_do_analyze_rel(m, v1443, l1, v1401, v1608, v1609, int32(0), v1404, v1541)
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L223
	}
L221:
	;
	v1614 = v1604
	goto L222
L222:
	;
	v1615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1614)+126)))
	if v1615 == int32(1) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1443)+48))
	v1614 = v1613
	goto L222
L224:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+28))
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+24))
	F_do_analyze_rel(m, v1443, l1, v1401, v1618, v1619, int32(1), v1404, v1541)
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	F_relation_close(m, v1443, int32(0))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L228
	}
L227:
	;
	goto L226
L228:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v1628 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	goto L163
L230:
	;
	goto L229
L231:
	;
	v1632 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v1632 != int32(1) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1628)+220))
	if v1635 == int32(0) {
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v1638 = int32(4484100)
	v1640 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v1641 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1640 + v1641
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1628)))
	*(*int32)(unsafe.Add(mBase, uint32(v1628))) = v1644 + v1641
	v1648 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1628)+220)) = v1648
	*(*int32)(unsafe.Add(mBase, uint32(v1628)+224)) = v1648
	*(*int32)(unsafe.Add(mBase, uint32(v1628))) = v1644 + int32(2)
	v1658 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v1658 - v1641
	goto L230
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L240
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L239
	}
L239:
	;
	goto L150
L240:
	;
	goto L150
L241:
	;
	goto L148
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_StartTransactionCommand(m)
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v1815&int32(513) == int32(1) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	goto L244
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v1218
	*(*int32)(unsafe.Add(mBase, uint32(v33)+88)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v1214
	*(*int32)(unsafe.Add(mBase, uint32(v33)+96)) = v1212
	*(*int32)(unsafe.Add(mBase, uint32(v33)+100)) = v1215
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v1216
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v1217
	*(*int32)(unsafe.Add(mBase, uint32(v33)+112)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v33)+116)) = v1209
	*(*int32)(unsafe.Add(mBase, uint32(v33)+120)) = v1207
	*(*int32)(unsafe.Add(mBase, uint32(v33)+124)) = v1208
	F_vac_update_datfrozenxid(m)
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		v1898 = v1229
		goto L6
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	m.G0 = v33 + int32(128)
	return
L249:
	;
	goto L248
L250:
	;
	goto L5
L251:
	;
	v1904 = int32(v1900)
	m.G0 = v1898
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1904)+4))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1904)))
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1907)))
	if v33+int32(72) == v1911 {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	m.ExcPending = 1
	goto L260
L253:
	;
	if v1914 != 0 {
		goto L257
	} else {
		goto L258
	}
L254:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1907)+4))
	v1914 = v1913
	goto L256
L255:
	;
	v1914 = int32(0)
	goto L256
L256:
	;
	goto L253
L257:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v33)+124))
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v33)+120))
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v33)+116))
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v33)+108))
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v33)+104))
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v33)+100))
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v33)+96))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v33)+92))
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v42 = v1918
	v43 = v1916
	v44 = v1915
	v45 = v1917
	v46 = v1926
	v47 = v1927
	v48 = v1922
	v49 = v1924
	v50 = v1923
	v51 = v1921
	v52 = v1920
	v53 = v1919
	v54 = v1914
	v55 = v1925
	v56 = v1906
	v65 = v1898
	goto L1
L258:
	;
	goto L259
L259:
	;
	F___wasm_longjmp(m, v1907, v1906)
	mBase = m.M
	v1929 = m.ExcPending
	if v1929 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	return
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vacuum_get_cutoffs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v170 int32
	_ = v170
	var v173 float64
	_ = v173
	var v175 float64
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 float64
	_ = v206
	var v209 float64
	_ = v209
	var v211 float64
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+140))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v22
	v24 = F_GetOldestNonRemovableTransactionId(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v24
		v29 = F_GetOldestMultiXactId(m)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v29
			v32 = F_ReadNextFullTransactionId(m)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v35 = F_ReadNextMultiXactId(m)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = F_MultiXactMemberFreezeThreshold(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v37 == v35 {
							v41 = int32(1)
						} else {
							v41 = v35 - v37
						}
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						v43 = int32(3)
						v44 = base.I32_wrap_i64(v32)
						v46 = *(*int32)(unsafe.Add(mBase, _consts[414]))
						v47 = v44 - v46
						if base.Ui32(v47) <= base.Ui32(v43) {
							v50 = v43
						} else {
							v50 = v47
						}
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v50))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v42)) == int32(0) {
							v62 = base.B2i32(base.Ui32(v42) < base.Ui32(v50))
						} else {
							v62 = int32(base.Ui32(v42-v50) >> (uint(int32(31)) % 32))
						}
						if v62 == int32(0) {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
							if int32(base.Ui32(v84-v41)>>(uint(int32(31))%32)) == int32(0) {
								v111 = *(*int32)(unsafe.Add(mBase, _consts[415]))
								if v17 < int32(0) {
									v114 = v111
								} else {
									v114 = v17
								}
								v116 = *(*int32)(unsafe.Add(mBase, _consts[414]))
								v118 = base.I32_div_s(v116, int32(2))
								if v114 < v118 {
									v120 = v114
								} else {
									v120 = v118
								}
								v121 = v44 - v120
								if base.Ui32(v121) <= base.Ui32(int32(3)) {
									v124 = int32(3)
								} else {
									v124 = v121
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
									v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
								} else {
									v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
								}
								if v138 != 0 {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
								} else {
								}
								v143 = *(*int32)(unsafe.Add(mBase, _consts[416]))
								if v16 < int32(0) {
									v146 = v143
								} else {
									v146 = v16
								}
								v148 = base.I32_div_s(v37, int32(2))
								if v146 < v148 {
									v150 = v146
								} else {
									v150 = v148
								}
								if v35 == v150 {
									v153 = int32(1)
								} else {
									v153 = v35 - v150
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
								v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
								if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
									v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
								} else {
								}
								v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v164 = *(*int32)(unsafe.Add(mBase, _consts[417]))
								if v15 < int32(0) {
									v167 = v164
								} else {
									v167 = v15
								}
								v168 = base.F64_convert_i32_s(v167)
								v170 = *(*int32)(unsafe.Add(mBase, _consts[414]))
								v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
								if base.F64_lt(v168, v173) != 0 {
									v175 = v168
								} else {
									v175 = v173
								}
								if base.F64_lt(base.F64_abs(v175), float64(2.147483648e+09)) != 0 {
									v179 = base.I32_trunc_f64_s(v175)
									v181 = v179
								} else {
									v181 = int32(-2147483648)
								}
								v182 = v44 - v181
								if base.Ui32(v182) <= base.Ui32(int32(3)) {
									v185 = int32(3)
								} else {
									v185 = v182
								}
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
									v197 = base.B2i32(base.Ui32(v161) <= base.Ui32(v185))
								} else {
									v197 = base.B2i32(v161-v185 <= int32(0))
								}
								if v197 != 0 {
									v227 = int32(1)
								} else {
									v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v202 = *(*int32)(unsafe.Add(mBase, _consts[418]))
									if v14 < int32(0) {
										v205 = v202
									} else {
										v205 = v14
									}
									v206 = base.F64_convert_i32_s(v205)
									v209 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
									if base.F64_lt(v206, v209) != 0 {
										v211 = v206
									} else {
										v211 = v209
									}
									if base.F64_lt(base.F64_abs(v211), float64(2.147483648e+09)) != 0 {
										v215 = base.I32_trunc_f64_s(v211)
										v217 = v215
									} else {
										v217 = int32(-2147483648)
									}
									if v217 == v35 {
										v220 = int32(1)
									} else {
										v220 = v35 - v217
									}
									v227 = base.B2i32(v199-v220 <= int32(0))
								}
								return v227
							} else {
								v92 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if v92 == int32(0) {
										v111 = *(*int32)(unsafe.Add(mBase, _consts[415]))
										if v17 < int32(0) {
											v114 = v111
										} else {
											v114 = v17
										}
										v116 = *(*int32)(unsafe.Add(mBase, _consts[414]))
										v118 = base.I32_div_s(v116, int32(2))
										if v114 < v118 {
											v120 = v114
										} else {
											v120 = v118
										}
										v121 = v44 - v120
										if base.Ui32(v121) <= base.Ui32(int32(3)) {
											v124 = int32(3)
										} else {
											v124 = v121
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
										v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
											v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
										} else {
											v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
										}
										if v138 != 0 {
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
										} else {
										}
										v143 = *(*int32)(unsafe.Add(mBase, _consts[416]))
										if v16 < int32(0) {
											v146 = v143
										} else {
											v146 = v16
										}
										v148 = base.I32_div_s(v37, int32(2))
										if v146 < v148 {
											v150 = v146
										} else {
											v150 = v148
										}
										if v35 == v150 {
											v153 = int32(1)
										} else {
											v153 = v35 - v150
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
										v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
										if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
											v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
										} else {
										}
										v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v164 = *(*int32)(unsafe.Add(mBase, _consts[417]))
										if v15 < int32(0) {
											v167 = v164
										} else {
											v167 = v15
										}
										v168 = base.F64_convert_i32_s(v167)
										v170 = *(*int32)(unsafe.Add(mBase, _consts[414]))
										v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
										if base.F64_lt(v168, v173) != 0 {
											v175 = v168
										} else {
											v175 = v173
										}
										if base.F64_lt(base.F64_abs(v175), float64(2.147483648e+09)) != 0 {
											v179 = base.I32_trunc_f64_s(v175)
											v181 = v179
										} else {
											v181 = int32(-2147483648)
										}
										v182 = v44 - v181
										if base.Ui32(v182) <= base.Ui32(int32(3)) {
											v185 = int32(3)
										} else {
											v185 = v182
										}
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
											v197 = base.B2i32(base.Ui32(v161) <= base.Ui32(v185))
										} else {
											v197 = base.B2i32(v161-v185 <= int32(0))
										}
										if v197 != 0 {
											v227 = int32(1)
										} else {
											v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
											v202 = *(*int32)(unsafe.Add(mBase, _consts[418]))
											if v14 < int32(0) {
												v205 = v202
											} else {
												v205 = v14
											}
											v206 = base.F64_convert_i32_s(v205)
											v209 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
											if base.F64_lt(v206, v209) != 0 {
												v211 = v206
											} else {
												v211 = v209
											}
											if base.F64_lt(base.F64_abs(v211), float64(2.147483648e+09)) != 0 {
												v215 = base.I32_trunc_f64_s(v211)
												v217 = v215
											} else {
												v217 = int32(-2147483648)
											}
											if v217 == v35 {
												v220 = int32(1)
											} else {
												v220 = v35 - v217
											}
											v227 = base.B2i32(v199-v220 <= int32(0))
										}
										return v227
									} else {
										F_errmsg(m, int32(77783), int32(0))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return int32(0)
										} else {
											F_errhint(m, int32(568469), int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496439), int32(1190), int32(156704))
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return int32(0)
												} else {
													v111 = *(*int32)(unsafe.Add(mBase, _consts[415]))
													if v17 < int32(0) {
														v114 = v111
													} else {
														v114 = v17
													}
													v116 = *(*int32)(unsafe.Add(mBase, _consts[414]))
													v118 = base.I32_div_s(v116, int32(2))
													if v114 < v118 {
														v120 = v114
													} else {
														v120 = v118
													}
													v121 = v44 - v120
													if base.Ui32(v121) <= base.Ui32(int32(3)) {
														v124 = int32(3)
													} else {
														v124 = v121
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
													v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
														v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
													} else {
														v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
													}
													if v138 != 0 {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
													} else {
													}
													v143 = *(*int32)(unsafe.Add(mBase, _consts[416]))
													if v16 < int32(0) {
														v146 = v143
													} else {
														v146 = v16
													}
													v148 = base.I32_div_s(v37, int32(2))
													if v146 < v148 {
														v150 = v146
													} else {
														v150 = v148
													}
													if v35 == v150 {
														v153 = int32(1)
													} else {
														v153 = v35 - v150
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
													v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
														v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
													} else {
													}
													v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													v164 = *(*int32)(unsafe.Add(mBase, _consts[417]))
													if v15 < int32(0) {
														v167 = v164
													} else {
														v167 = v15
													}
													v168 = base.F64_convert_i32_s(v167)
													v170 = *(*int32)(unsafe.Add(mBase, _consts[414]))
													v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
													if base.F64_lt(v168, v173) != 0 {
														v175 = v168
													} else {
														v175 = v173
													}
													if base.F64_lt(base.F64_abs(v175), float64(2.147483648e+09)) != 0 {
														v179 = base.I32_trunc_f64_s(v175)
														v181 = v179
													} else {
														v181 = int32(-2147483648)
													}
													v182 = v44 - v181
													if base.Ui32(v182) <= base.Ui32(int32(3)) {
														v185 = int32(3)
													} else {
														v185 = v182
													}
													if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
														v197 = base.B2i32(base.Ui32(v161) <= base.Ui32(v185))
													} else {
														v197 = base.B2i32(v161-v185 <= int32(0))
													}
													if v197 != 0 {
														v227 = int32(1)
													} else {
														v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
														v202 = *(*int32)(unsafe.Add(mBase, _consts[418]))
														if v14 < int32(0) {
															v205 = v202
														} else {
															v205 = v14
														}
														v206 = base.F64_convert_i32_s(v205)
														v209 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
														if base.F64_lt(v206, v209) != 0 {
															v211 = v206
														} else {
															v211 = v209
														}
														if base.F64_lt(base.F64_abs(v211), float64(2.147483648e+09)) != 0 {
															v215 = base.I32_trunc_f64_s(v211)
															v217 = v215
														} else {
															v217 = int32(-2147483648)
														}
														if v217 == v35 {
															v220 = int32(1)
														} else {
															v220 = v35 - v217
														}
														v227 = base.B2i32(v199-v220 <= int32(0))
													}
													return v227
												}
											}
										}
									}
								}
							}
						} else {
							v67 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								if v67 == int32(0) {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									if int32(base.Ui32(v84-v41)>>(uint(int32(31))%32)) == int32(0) {
										v111 = *(*int32)(unsafe.Add(mBase, _consts[415]))
										if v17 < int32(0) {
											v114 = v111
										} else {
											v114 = v17
										}
										v116 = *(*int32)(unsafe.Add(mBase, _consts[414]))
										v118 = base.I32_div_s(v116, int32(2))
										if v114 < v118 {
											v120 = v114
										} else {
											v120 = v118
										}
										v121 = v44 - v120
										if base.Ui32(v121) <= base.Ui32(int32(3)) {
											v124 = int32(3)
										} else {
											v124 = v121
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
										v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
											v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
										} else {
											v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
										}
										if v138 != 0 {
											v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
										} else {
										}
										v143 = *(*int32)(unsafe.Add(mBase, _consts[416]))
										if v16 < int32(0) {
											v146 = v143
										} else {
											v146 = v16
										}
										v148 = base.I32_div_s(v37, int32(2))
										if v146 < v148 {
											v150 = v146
										} else {
											v150 = v148
										}
										if v35 == v150 {
											v153 = int32(1)
										} else {
											v153 = v35 - v150
										}
										*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
										v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
										if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
											v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
										} else {
										}
										v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v164 = *(*int32)(unsafe.Add(mBase, _consts[417]))
										if v15 < int32(0) {
											v167 = v164
										} else {
											v167 = v15
										}
										v168 = base.F64_convert_i32_s(v167)
										v170 = *(*int32)(unsafe.Add(mBase, _consts[414]))
										v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
										if base.F64_lt(v168, v173) != 0 {
											v175 = v168
										} else {
											v175 = v173
										}
										if base.F64_lt(base.F64_abs(v175), float64(2.147483648e+09)) != 0 {
											v179 = base.I32_trunc_f64_s(v175)
											v181 = v179
										} else {
											v181 = int32(-2147483648)
										}
										v182 = v44 - v181
										if base.Ui32(v182) <= base.Ui32(int32(3)) {
											v185 = int32(3)
										} else {
											v185 = v182
										}
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
											v197 = base.B2i32(base.Ui32(v161) <= base.Ui32(v185))
										} else {
											v197 = base.B2i32(v161-v185 <= int32(0))
										}
										if v197 != 0 {
											v227 = int32(1)
										} else {
											v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
											v202 = *(*int32)(unsafe.Add(mBase, _consts[418]))
											if v14 < int32(0) {
												v205 = v202
											} else {
												v205 = v14
											}
											v206 = base.F64_convert_i32_s(v205)
											v209 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
											if base.F64_lt(v206, v209) != 0 {
												v211 = v206
											} else {
												v211 = v209
											}
											if base.F64_lt(base.F64_abs(v211), float64(2.147483648e+09)) != 0 {
												v215 = base.I32_trunc_f64_s(v211)
												v217 = v215
											} else {
												v217 = int32(-2147483648)
											}
											if v217 == v35 {
												v220 = int32(1)
											} else {
												v220 = v35 - v217
											}
											v227 = base.B2i32(v199-v220 <= int32(0))
										}
										return v227
									} else {
										v92 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											if v92 == int32(0) {
												v111 = *(*int32)(unsafe.Add(mBase, _consts[415]))
												if v17 < int32(0) {
													v114 = v111
												} else {
													v114 = v17
												}
												v116 = *(*int32)(unsafe.Add(mBase, _consts[414]))
												v118 = base.I32_div_s(v116, int32(2))
												if v114 < v118 {
													v120 = v114
												} else {
													v120 = v118
												}
												v121 = v44 - v120
												if base.Ui32(v121) <= base.Ui32(int32(3)) {
													v124 = int32(3)
												} else {
													v124 = v121
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
												v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
												if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
													v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
												} else {
													v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
												}
												if v138 != 0 {
													v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
												} else {
												}
												v143 = *(*int32)(unsafe.Add(mBase, _consts[416]))
												if v16 < int32(0) {
													v146 = v143
												} else {
													v146 = v16
												}
												v148 = base.I32_div_s(v37, int32(2))
												if v146 < v148 {
													v150 = v146
												} else {
													v150 = v148
												}
												if v35 == v150 {
													v153 = int32(1)
												} else {
													v153 = v35 - v150
												}
												*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
												v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
													v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
												} else {
												}
												v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												v164 = *(*int32)(unsafe.Add(mBase, _consts[417]))
												if v15 < int32(0) {
													v167 = v164
												} else {
													v167 = v15
												}
												v168 = base.F64_convert_i32_s(v167)
												v170 = *(*int32)(unsafe.Add(mBase, _consts[414]))
												v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
												if base.F64_lt(v168, v173) != 0 {
													v175 = v168
												} else {
													v175 = v173
												}
												if base.F64_lt(base.F64_abs(v175), float64(2.147483648e+09)) != 0 {
													v179 = base.I32_trunc_f64_s(v175)
													v181 = v179
												} else {
													v181 = int32(-2147483648)
												}
												v182 = v44 - v181
												if base.Ui32(v182) <= base.Ui32(int32(3)) {
													v185 = int32(3)
												} else {
													v185 = v182
												}
												if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
													v197 = base.B2i32(base.Ui32(v161) <= base.Ui32(v185))
												} else {
													v197 = base.B2i32(v161-v185 <= int32(0))
												}
												if v197 != 0 {
													v227 = int32(1)
												} else {
													v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
													v202 = *(*int32)(unsafe.Add(mBase, _consts[418]))
													if v14 < int32(0) {
														v205 = v202
													} else {
														v205 = v14
													}
													v206 = base.F64_convert_i32_s(v205)
													v209 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
													if base.F64_lt(v206, v209) != 0 {
														v211 = v206
													} else {
														v211 = v209
													}
													if base.F64_lt(base.F64_abs(v211), float64(2.147483648e+09)) != 0 {
														v215 = base.I32_trunc_f64_s(v211)
														v217 = v215
													} else {
														v217 = int32(-2147483648)
													}
													if v217 == v35 {
														v220 = int32(1)
													} else {
														v220 = v35 - v217
													}
													v227 = base.B2i32(v199-v220 <= int32(0))
												}
												return v227
											} else {
												F_errmsg(m, int32(77783), int32(0))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(568469), int32(0))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(496439), int32(1190), int32(156704))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return int32(0)
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, _consts[415]))
															if v17 < int32(0) {
																v114 = v111
															} else {
																v114 = v17
															}
															v116 = *(*int32)(unsafe.Add(mBase, _consts[414]))
															v118 = base.I32_div_s(v116, int32(2))
															if v114 < v118 {
																v120 = v114
															} else {
																v120 = v118
															}
															v121 = v44 - v120
															if base.Ui32(v121) <= base.Ui32(int32(3)) {
																v124 = int32(3)
															} else {
																v124 = v121
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
															v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
															if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
																v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
															} else {
																v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
															}
															if v138 != 0 {
																v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
																*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
															} else {
															}
															v143 = *(*int32)(unsafe.Add(mBase, _consts[416]))
															if v16 < int32(0) {
																v146 = v143
															} else {
																v146 = v16
															}
															v148 = base.I32_div_s(v37, int32(2))
															if v146 < v148 {
																v150 = v146
															} else {
																v150 = v148
															}
															if v35 == v150 {
																v153 = int32(1)
															} else {
																v153 = v35 - v150
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
															v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
															if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
																v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
															} else {
															}
															v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
															v164 = *(*int32)(unsafe.Add(mBase, _consts[417]))
															if v15 < int32(0) {
																v167 = v164
															} else {
																v167 = v15
															}
															v168 = base.F64_convert_i32_s(v167)
															v170 = *(*int32)(unsafe.Add(mBase, _consts[414]))
															v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
															if base.F64_lt(v168, v173) != 0 {
																v175 = v168
															} else {
																v175 = v173
															}
															if base.F64_lt(base.F64_abs(v175), float64(2.147483648e+09)) != 0 {
																v179 = base.I32_trunc_f64_s(v175)
																v181 = v179
															} else {
																v181 = int32(-2147483648)
															}
															v182 = v44 - v181
															if base.Ui32(v182) <= base.Ui32(int32(3)) {
																v185 = int32(3)
															} else {
																v185 = v182
															}
															if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
																v197 = base.B2i32(base.Ui32(v161) <= base.Ui32(v185))
															} else {
																v197 = base.B2i32(v161-v185 <= int32(0))
															}
															if v197 != 0 {
																v227 = int32(1)
															} else {
																v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
																v202 = *(*int32)(unsafe.Add(mBase, _consts[418]))
																if v14 < int32(0) {
																	v205 = v202
																} else {
																	v205 = v14
																}
																v206 = base.F64_convert_i32_s(v205)
																v209 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
																if base.F64_lt(v206, v209) != 0 {
																	v211 = v206
																} else {
																	v211 = v209
																}
																if base.F64_lt(base.F64_abs(v211), float64(2.147483648e+09)) != 0 {
																	v215 = base.I32_trunc_f64_s(v211)
																	v217 = v215
																} else {
																	v217 = int32(-2147483648)
																}
																if v217 == v35 {
																	v220 = int32(1)
																} else {
																	v220 = v35 - v217
																}
																v227 = base.B2i32(v199-v220 <= int32(0))
															}
															return v227
														}
													}
												}
											}
										}
									}
								} else {
									F_errmsg(m, int32(77833), int32(0))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										F_errhint(m, int32(568469), int32(0))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496439), int32(1185), int32(156704))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
												if int32(base.Ui32(v84-v41)>>(uint(int32(31))%32)) == int32(0) {
													v111 = *(*int32)(unsafe.Add(mBase, _consts[415]))
													if v17 < int32(0) {
														v114 = v111
													} else {
														v114 = v17
													}
													v116 = *(*int32)(unsafe.Add(mBase, _consts[414]))
													v118 = base.I32_div_s(v116, int32(2))
													if v114 < v118 {
														v120 = v114
													} else {
														v120 = v118
													}
													v121 = v44 - v120
													if base.Ui32(v121) <= base.Ui32(int32(3)) {
														v124 = int32(3)
													} else {
														v124 = v121
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
													v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
													if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
														v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
													} else {
														v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
													}
													if v138 != 0 {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
													} else {
													}
													v143 = *(*int32)(unsafe.Add(mBase, _consts[416]))
													if v16 < int32(0) {
														v146 = v143
													} else {
														v146 = v16
													}
													v148 = base.I32_div_s(v37, int32(2))
													if v146 < v148 {
														v150 = v146
													} else {
														v150 = v148
													}
													if v35 == v150 {
														v153 = int32(1)
													} else {
														v153 = v35 - v150
													}
													*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
													v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
													if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
														v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
													} else {
													}
													v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
													v164 = *(*int32)(unsafe.Add(mBase, _consts[417]))
													if v15 < int32(0) {
														v167 = v164
													} else {
														v167 = v15
													}
													v168 = base.F64_convert_i32_s(v167)
													v170 = *(*int32)(unsafe.Add(mBase, _consts[414]))
													v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
													if base.F64_lt(v168, v173) != 0 {
														v175 = v168
													} else {
														v175 = v173
													}
													if base.F64_lt(base.F64_abs(v175), float64(2.147483648e+09)) != 0 {
														v179 = base.I32_trunc_f64_s(v175)
														v181 = v179
													} else {
														v181 = int32(-2147483648)
													}
													v182 = v44 - v181
													if base.Ui32(v182) <= base.Ui32(int32(3)) {
														v185 = int32(3)
													} else {
														v185 = v182
													}
													if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
														v197 = base.B2i32(base.Ui32(v161) <= base.Ui32(v185))
													} else {
														v197 = base.B2i32(v161-v185 <= int32(0))
													}
													if v197 != 0 {
														v227 = int32(1)
													} else {
														v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
														v202 = *(*int32)(unsafe.Add(mBase, _consts[418]))
														if v14 < int32(0) {
															v205 = v202
														} else {
															v205 = v14
														}
														v206 = base.F64_convert_i32_s(v205)
														v209 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
														if base.F64_lt(v206, v209) != 0 {
															v211 = v206
														} else {
															v211 = v209
														}
														if base.F64_lt(base.F64_abs(v211), float64(2.147483648e+09)) != 0 {
															v215 = base.I32_trunc_f64_s(v211)
															v217 = v215
														} else {
															v217 = int32(-2147483648)
														}
														if v217 == v35 {
															v220 = int32(1)
														} else {
															v220 = v35 - v217
														}
														v227 = base.B2i32(v199-v220 <= int32(0))
													}
													return v227
												} else {
													v92 = F_errstart(m, int32(19), int32(0))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int32(0)
													} else {
														if v92 == int32(0) {
															v111 = *(*int32)(unsafe.Add(mBase, _consts[415]))
															if v17 < int32(0) {
																v114 = v111
															} else {
																v114 = v17
															}
															v116 = *(*int32)(unsafe.Add(mBase, _consts[414]))
															v118 = base.I32_div_s(v116, int32(2))
															if v114 < v118 {
																v120 = v114
															} else {
																v120 = v118
															}
															v121 = v44 - v120
															if base.Ui32(v121) <= base.Ui32(int32(3)) {
																v124 = int32(3)
															} else {
																v124 = v121
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
															v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
															if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
																v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
															} else {
																v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
															}
															if v138 != 0 {
																v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
																*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
															} else {
															}
															v143 = *(*int32)(unsafe.Add(mBase, _consts[416]))
															if v16 < int32(0) {
																v146 = v143
															} else {
																v146 = v16
															}
															v148 = base.I32_div_s(v37, int32(2))
															if v146 < v148 {
																v150 = v146
															} else {
																v150 = v148
															}
															if v35 == v150 {
																v153 = int32(1)
															} else {
																v153 = v35 - v150
															}
															*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
															v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
															if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
																v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
															} else {
															}
															v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
															v164 = *(*int32)(unsafe.Add(mBase, _consts[417]))
															if v15 < int32(0) {
																v167 = v164
															} else {
																v167 = v15
															}
															v168 = base.F64_convert_i32_s(v167)
															v170 = *(*int32)(unsafe.Add(mBase, _consts[414]))
															v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
															if base.F64_lt(v168, v173) != 0 {
																v175 = v168
															} else {
																v175 = v173
															}
															if base.F64_lt(base.F64_abs(v175), float64(2.147483648e+09)) != 0 {
																v179 = base.I32_trunc_f64_s(v175)
																v181 = v179
															} else {
																v181 = int32(-2147483648)
															}
															v182 = v44 - v181
															if base.Ui32(v182) <= base.Ui32(int32(3)) {
																v185 = int32(3)
															} else {
																v185 = v182
															}
															if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
																v197 = base.B2i32(base.Ui32(v161) <= base.Ui32(v185))
															} else {
																v197 = base.B2i32(v161-v185 <= int32(0))
															}
															if v197 != 0 {
																v227 = int32(1)
															} else {
																v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
																v202 = *(*int32)(unsafe.Add(mBase, _consts[418]))
																if v14 < int32(0) {
																	v205 = v202
																} else {
																	v205 = v14
																}
																v206 = base.F64_convert_i32_s(v205)
																v209 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
																if base.F64_lt(v206, v209) != 0 {
																	v211 = v206
																} else {
																	v211 = v209
																}
																if base.F64_lt(base.F64_abs(v211), float64(2.147483648e+09)) != 0 {
																	v215 = base.I32_trunc_f64_s(v211)
																	v217 = v215
																} else {
																	v217 = int32(-2147483648)
																}
																if v217 == v35 {
																	v220 = int32(1)
																} else {
																	v220 = v35 - v217
																}
																v227 = base.B2i32(v199-v220 <= int32(0))
															}
															return v227
														} else {
															F_errmsg(m, int32(77783), int32(0))
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																F_errhint(m, int32(568469), int32(0))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(496439), int32(1190), int32(156704))
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return int32(0)
																	} else {
																		v111 = *(*int32)(unsafe.Add(mBase, _consts[415]))
																		if v17 < int32(0) {
																			v114 = v111
																		} else {
																			v114 = v17
																		}
																		v116 = *(*int32)(unsafe.Add(mBase, _consts[414]))
																		v118 = base.I32_div_s(v116, int32(2))
																		if v114 < v118 {
																			v120 = v114
																		} else {
																			v120 = v118
																		}
																		v121 = v44 - v120
																		if base.Ui32(v121) <= base.Ui32(int32(3)) {
																			v124 = int32(3)
																		} else {
																			v124 = v121
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v124
																		v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
																		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v124))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v126)) == int32(0) {
																			v138 = base.B2i32(base.Ui32(v126) < base.Ui32(v124))
																		} else {
																			v138 = int32(base.Ui32(v126-v124) >> (uint(int32(31)) % 32))
																		}
																		if v138 != 0 {
																			v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v139
																		} else {
																		}
																		v143 = *(*int32)(unsafe.Add(mBase, _consts[416]))
																		if v16 < int32(0) {
																			v146 = v143
																		} else {
																			v146 = v16
																		}
																		v148 = base.I32_div_s(v37, int32(2))
																		if v146 < v148 {
																			v150 = v146
																		} else {
																			v150 = v148
																		}
																		if v35 == v150 {
																			v153 = int32(1)
																		} else {
																			v153 = v35 - v150
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v153
																		v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
																		if int32(base.Ui32(v155-v153)>>(uint(int32(31))%32)) != 0 {
																			v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
																			*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v159
																		} else {
																		}
																		v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
																		v164 = *(*int32)(unsafe.Add(mBase, _consts[417]))
																		if v15 < int32(0) {
																			v167 = v164
																		} else {
																			v167 = v15
																		}
																		v168 = base.F64_convert_i32_s(v167)
																		v170 = *(*int32)(unsafe.Add(mBase, _consts[414]))
																		v173 = base.F64_mul(base.F64_convert_i32_s(v170), float64(0.95))
																		if base.F64_lt(v168, v173) != 0 {
																			v175 = v168
																		} else {
																			v175 = v173
																		}
																		if base.F64_lt(base.F64_abs(v175), float64(2.147483648e+09)) != 0 {
																			v179 = base.I32_trunc_f64_s(v175)
																			v181 = v179
																		} else {
																			v181 = int32(-2147483648)
																		}
																		v182 = v44 - v181
																		if base.Ui32(v182) <= base.Ui32(int32(3)) {
																			v185 = int32(3)
																		} else {
																			v185 = v182
																		}
																		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v185))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v161)) == int32(0) {
																			v197 = base.B2i32(base.Ui32(v161) <= base.Ui32(v185))
																		} else {
																			v197 = base.B2i32(v161-v185 <= int32(0))
																		}
																		if v197 != 0 {
																			v227 = int32(1)
																		} else {
																			v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
																			v202 = *(*int32)(unsafe.Add(mBase, _consts[418]))
																			if v14 < int32(0) {
																				v205 = v202
																			} else {
																				v205 = v14
																			}
																			v206 = base.F64_convert_i32_s(v205)
																			v209 = base.F64_mul(base.F64_convert_i32_s(v37), float64(0.95))
																			if base.F64_lt(v206, v209) != 0 {
																				v211 = v206
																			} else {
																				v211 = v209
																			}
																			if base.F64_lt(base.F64_abs(v211), float64(2.147483648e+09)) != 0 {
																				v215 = base.I32_trunc_f64_s(v211)
																				v217 = v215
																			} else {
																				v217 = int32(-2147483648)
																			}
																			if v217 == v35 {
																				v220 = int32(1)
																			} else {
																				v220 = v35 - v217
																			}
																			v227 = base.B2i32(v199-v220 <= int32(0))
																		}
																		return v227
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_vacuum_rel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 float64
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v17
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v27
	F_StartTransactionCommand(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v33&int32(16) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v43 = F_LWLockAcquire(m, v39+int32(512), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v73 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+124)))
	v49 = v47 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+124)) = uint8(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v51 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v55 = v47 | int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+124)) = uint8(v55)
	v57 = v55
	goto L9
L8:
	;
	v57 = v49
	goto L9
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v61))) = uint8(v57)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v65+int32(512))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L5
L11:
	;
	F_PushActiveSnapshot(m, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v78 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v81&int32(16) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	m.G0 = v13 + int32(80)
	return v317
L18:
	;
	v91 = int32(8)
	goto L20
L19:
	;
	v91 = int32(4)
	goto L20
L20:
	;
	v92 = F_vacuum_open_relation(m, l0, l1, v81, int32(base.Ui32(v82^int32(-1))>>(uint(int32(31))%32)), v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v92 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	if v94 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L115
	}
L25:
	;
	F_relation_close(m, v92, v91)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L114
	}
L26:
	;
	v96 = v94
	goto L28
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+56))
	v96 = v95
	goto L28
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v101 = F_vacuum_is_permitted_for_relation(m, v96, v97, v98&int32(-3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v101 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+119)))
	v108 = v106 - int32(109)
	if int32(1)<<(uint(v108)%32)&int32(169) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = base.B2i32(base.Ui32(v108) <= base.Ui32(int32(7)))
	goto L33
L32:
	;
	v116 = int32(0)
	goto L33
L33:
	;
	if v116 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v119 = int32(0)
	v122 = F_errstart(m, int32(19), v119)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+118)))
	if v142 == int32(116) {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	if v122 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v124 + int32(4)
	F_errmsg(m, int32(165683), v13)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_relation_close(m, v92, v91)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	F_errfinish(m, int32(496439), int32(2144), int32(307202))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v317 = v119
	goto L17
L46:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+24)))
	if v145 == int32(0) {
		goto L25
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v106 == int32(112) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L48
L50:
	;
	v317 = int32(1)
	goto L17
L51:
	;
	F_relation_close(m, v92, v91)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v92)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v156
	F_LockRelationIdForSession(m, v13+int32(72), v91)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L50
L57:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v162 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v92)+180))
	if v165 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v92)+180))
	if v179 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L61:
	;
	v167 = int32(2)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+112))
	if v168 == v167 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v176 = int32(1)
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v176
	goto L60
L64:
	;
	v171 = int32(3)
	goto L66
L65:
	;
	v171 = v167
	goto L66
L66:
	;
	if v168 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v173 = v171
	goto L69
L68:
	;
	v173 = int32(1)
	goto L69
L69:
	;
	v176 = v173
	goto L63
L70:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v189 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v179)+120))
	if base.F64_ge(v182, float64(0)) == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l2)+40)) = v182
	goto L70
L73:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v92)+180))
	if v194 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v206 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v207&int32(128) == v206 {
		v218 = v206
		goto L85
	} else {
		goto L86
	}
L76:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+117)))
	if v198 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v201 = int32(4386672)
	goto L78
L78:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v202 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v199 = v194 + int32(116)
	goto L81
L80:
	;
	v199 = int32(4386672)
	goto L81
L81:
	;
	v201 = v199
	goto L78
L82:
	;
	v203 = int32(3)
	goto L84
L83:
	;
	v203 = int32(2)
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v203
	goto L75
L85:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(68)))) = v224
	v227 = *(*int32)(unsafe.Add(mBase, _consts[308]))
	*(*int32)(unsafe.Add(mBase, uint32(v13-int32(-64)))) = v227
	goto L88
L86:
	;
	v212 = int32(80)
	if v207&v212 == v212 {
		v218 = v206
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+112))
	v218 = v217
	goto L85
L88:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+80))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	*(*int32)(unsafe.Add(mBase, _consts[308])) = v231 | int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[31])) = v230
	goto L89
L89:
	;
	v239 = int32(4487480)
	v241 = *(*int32)(unsafe.Add(mBase, _consts[309]))
	v243 = v241 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[309])) = v243
	goto L90
L90:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v247&int32(64) != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	F_AtEOXact_GUC(m, int32(0), v243)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L101
	}
L93:
	;
	if v247&int32(16) != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	v267 = v92
	goto L92
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(base.Ui32(v247)>>(uint(int32(2))%32)) & int32(1)
	v257 = int32(0)
	F_cluster_rel(m, v92, v257, v13+int32(4))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v92)+188))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+128))
	m.T0[v264].(func(*base.Module, int32, int32, int32))(m, v92, l2, l3)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	v267 = v257
	goto L92
L100:
	;
	goto L95
L101:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	*(*int32)(unsafe.Add(mBase, _consts[308])) = v272
	*(*int32)(unsafe.Add(mBase, _consts[31])) = v271
	goto L102
L102:
	;
	if v267 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	F_relation_close(m, v267, int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L107
	}
L106:
	;
	goto L105
L107:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v218 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l0
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v285 | int32(64)
	v292 = F_vacuum_rel(m, v218, int32(0), v13+int32(8), l3)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_UnlockRelationIdForSession(m, v13+int32(72), v91)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	goto L50
L114:
	;
	goto L24
L115:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v317 = int32(0)
	goto L17
}
