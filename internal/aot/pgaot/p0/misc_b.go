package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BeginReportingGUCOptions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[958]))
	if v8 != int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v5 + int32(32)
	return
L2:
	;
	v12 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[959])) = uint8(v12)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
	if v16 == v12 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+316))
	v24 = base.B2i32(v22 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v24)
	v26 = v24
	goto L6
L5:
	;
	v26 = int32(0)
	goto L6
L6:
	;
	goto L3
L7:
	;
	v28 = int32(0)
	v31 = int32(10)
	v37 = F_set_config_with_handle(m, int32(23761), v28, int32(358953), v28, v31, v31, v28, int32(1), v28, v28)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[955]))
	F_hash_seq_init(m, v5+int32(12), v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	v47 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v47 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v52 = v47
	goto L15
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+20)))
	if v54&int32(64) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L1
L17:
	;
	F_ReportGUCOption(m, v53)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v61 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	if v61 != 0 {
		v52 = v61
		goto L15
	} else {
		goto L22
	}
L22:
	;
	goto L16
}
func F_BlockSampler_Init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v5 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = l0 + int32(16)
	v11 = base.I64_extend_i32_u(l3)
	v15 = v11 + int64(4354685564936845354)
	v16 = int64(30)
	v19 = int64(-4658895280553007687)
	v20 = (int64(base.Ui64(v15)>>(uint(v16)%64)) ^ v15) * v19
	v21 = int64(27)
	v24 = int64(-7723592293110705685)
	v25 = (int64(base.Ui64(v20)>>(uint(v21)%64)) ^ v20) * v24
	v26 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(base.Ui64(v25)>>(uint(v26)%64)) ^ v25
	v31 = v11 - int64(7046029254386353131)
	v36 = (int64(base.Ui64(v31)>>(uint(v16)%64)) ^ v31) * v19
	v41 = (int64(base.Ui64(v36)>>(uint(v21)%64)) ^ v36) * v24
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(base.Ui64(v41)>>(uint(v26)%64)) ^ v41
	if v31|v15 == v5 {
		*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(1442695040888963407)
		*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(6364136223846793005)
	} else {
	}
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v53) < base.Ui32(v54) {
		v56 = v53
	} else {
		v56 = v54
	}
	return v56
}
func F_basque_UTF_8_stem(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v323 int32
	_ = v323
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v444 int32
	_ = v444
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v562 int32
	_ = v562
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v607 int32
	_ = v607
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v685 int32
	_ = v685
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v722 int32
	_ = v722
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v803 int32
	_ = v803
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v922 int32
	_ = v922
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v966 int32
	_ = v966
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1043 int32
	_ = v1043
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1069 int32
	_ = v1069
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1178 int32
	_ = v1178
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1229 int32
	_ = v1229
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1262 int32
	_ = v1262
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1300 int32
	_ = v1300
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1351 int32
	_ = v1351
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1385 int32
	_ = v1385
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1425 int32
	_ = v1425
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1476 int32
	_ = v1476
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1502 int32
	_ = v1502
	var v1509 int32
	_ = v1509
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1547 int32
	_ = v1547
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1598 int32
	_ = v1598
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1618 int32
	_ = v1618
	var v1624 int32
	_ = v1624
	var v1632 int32
	_ = v1632
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1891 int32
	_ = v1891
	var v1895 int32
	_ = v1895
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L7
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1178 = v12
	goto L263
L2:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1152)+8)) = v1150
	goto L1
L3:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1150 = v1148 + v1146
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L136
L5:
	;
	if v130 != 0 {
		goto L4
	} else {
		goto L29
	}
L6:
	;
	v130 = v123
	goto L5
L7:
	;
	if v25 <= v12 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v123 = int32(0)
	goto L6
L9:
	;
	v130 = int32(-1)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v41 = int32(1)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v26))))
	if base.Ui32(v43) < base.Ui32(int32(192)) {
		v100 = v43
		v101 = v41
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if int32(117) < v100 {
		v123 = v101
		goto L6
	} else {
		goto L25
	}
L13:
	;
	v47 = v12 + int32(1)
	if v47 == v25 {
		v100 = v43
		v101 = v41
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v26))))
	v52 = v50 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v43) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v26))))
	v68 = v66 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v43) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v56 = v12 + int32(2)
	if v56 != v25 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v100 = v43<<(uint(int32(6))%32)&int32(1984) | v52
	v101 = int32(2)
	goto L12
L19:
	;
	goto L18
L20:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v72))))
	v100 = v85&int32(63) | (v43<<(uint(int32(18))%32)&int32(1835008) | v52<<(uint(int32(12))%32) | v68<<(uint(int32(6))%32))
	v101 = int32(4)
	goto L12
L21:
	;
	v72 = v12 + int32(3)
	if v72 != v25 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v100 = v43<<(uint(int32(12))%32)&int32(61440) | v52<<(uint(int32(6))%32) | v68
	v101 = int32(3)
	goto L12
L24:
	;
	goto L23
L25:
	;
	v105 = v100 - int32(97)
	if v105 < int32(0) {
		v123 = v101
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v105)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v111)>>(uint(v105&int32(7))%32))&int32(1) == int32(0) {
		v123 = v101
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v101 + v12
	goto L28
L28:
	;
	goto L8
L29:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L32
L30:
	;
	if v248 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L31:
	;
	v248 = v241
	goto L30
L32:
	;
	if v144 <= v131 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v241 = int32(0)
	goto L31
L34:
	;
	v248 = int32(-1)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v160 = int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v145))))
	if base.Ui32(v162) < base.Ui32(int32(192)) {
		v219 = v162
		v220 = v160
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if int32(117) < v219 {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v166 = v131 + int32(1)
	if v166 == v144 {
		v219 = v162
		v220 = v160
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+v145))))
	v171 = v169 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v162) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v145))))
	v187 = v185 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v162) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v175 = v131 + int32(2)
	if v175 != v144 {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v219 = v162<<(uint(int32(6))%32)&int32(1984) | v171
	v220 = int32(2)
	goto L37
L44:
	;
	goto L43
L45:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v191))))
	v219 = v204&int32(63) | (v162<<(uint(int32(18))%32)&int32(1835008) | v171<<(uint(int32(12))%32) | v187<<(uint(int32(6))%32))
	v220 = int32(4)
	goto L37
L46:
	;
	v191 = v131 + int32(3)
	if v191 != v144 {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v219 = v162<<(uint(int32(12))%32)&int32(61440) | v171<<(uint(int32(6))%32) | v187
	v220 = int32(3)
	goto L37
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v220 + v131
	goto L54
L51:
	;
	v224 = v219 - int32(97)
	if v224 < int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v224)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v230)>>(uint(v224&int32(7))%32))&int32(1) != 0 {
		v241 = v220
		goto L31
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	goto L33
L55:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v272 = v262
	goto L60
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v131
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L86
L58:
	;
	if int32(0) <= v367 {
		v1146 = v367
		goto L3
	} else {
		goto L83
	}
L59:
	;
	v367 = v339
	goto L58
L60:
	;
	if v263 <= v272 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v367 = int32(-1)
	goto L58
L63:
	;
	goto L64
L64:
	;
	v279 = int32(1)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272+v264))))
	if base.Ui32(v281) < base.Ui32(int32(192)) {
		v338 = v281
		v339 = v279
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if int32(117) < v338 {
		goto L78
	} else {
		goto L79
	}
L66:
	;
	v285 = v272 + int32(1)
	if v285 == v263 {
		v338 = v281
		v339 = v279
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285+v264))))
	v290 = v288 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v281) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v264))))
	v306 = v304 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v281) {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v294 = v272 + int32(2)
	if v294 != v263 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v338 = v281<<(uint(int32(6))%32)&int32(1984) | v290
	v339 = int32(2)
	goto L65
L72:
	;
	goto L71
L73:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264+v310))))
	v338 = v323&int32(63) | (v281<<(uint(int32(18))%32)&int32(1835008) | v290<<(uint(int32(12))%32) | v306<<(uint(int32(6))%32))
	v339 = int32(4)
	goto L65
L74:
	;
	v310 = v272 + int32(3)
	if v310 != v263 {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v338 = v281<<(uint(int32(12))%32)&int32(61440) | v290<<(uint(int32(6))%32) | v306
	v339 = int32(3)
	goto L65
L77:
	;
	goto L76
L78:
	;
	v356 = v339 + v272
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v356
	v272 = v356
	goto L60
L79:
	;
	v343 = v338 - int32(97)
	if v343 < int32(0) {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v343)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v349)>>(uint(v343&int32(7))%32))&int32(1) != 0 {
		goto L59
	} else {
		goto L81
	}
L81:
	;
	goto L78
L83:
	;
	goto L57
L84:
	;
	if v489 != 0 {
		goto L4
	} else {
		goto L108
	}
L85:
	;
	v489 = v482
	goto L84
L86:
	;
	if v384 <= v131 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v482 = int32(0)
	goto L85
L88:
	;
	v489 = int32(-1)
	goto L84
L89:
	;
	goto L90
L90:
	;
	v400 = int32(1)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v385))))
	if base.Ui32(v402) < base.Ui32(int32(192)) {
		v459 = v402
		v460 = v400
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if int32(117) < v459 {
		v482 = v460
		goto L85
	} else {
		goto L104
	}
L92:
	;
	v406 = v131 + int32(1)
	if v406 == v384 {
		v459 = v402
		v460 = v400
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406+v385))))
	v411 = v409 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v402) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v385))))
	v427 = v425 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v402) {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v415 = v131 + int32(2)
	if v415 != v384 {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v459 = v402<<(uint(int32(6))%32)&int32(1984) | v411
	v460 = int32(2)
	goto L91
L98:
	;
	goto L97
L99:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+v431))))
	v459 = v444&int32(63) | (v402<<(uint(int32(18))%32)&int32(1835008) | v411<<(uint(int32(12))%32) | v427<<(uint(int32(6))%32))
	v460 = int32(4)
	goto L91
L100:
	;
	v431 = v131 + int32(3)
	if v431 != v384 {
		goto L99
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v459 = v402<<(uint(int32(12))%32)&int32(61440) | v411<<(uint(int32(6))%32) | v427
	v460 = int32(3)
	goto L91
L103:
	;
	goto L102
L104:
	;
	v464 = v459 - int32(97)
	if v464 < int32(0) {
		v482 = v460
		goto L85
	} else {
		goto L105
	}
L105:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v464)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v470)>>(uint(v464&int32(7))%32))&int32(1) == int32(0) {
		v482 = v460
		goto L85
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 + v131
	goto L107
L107:
	;
	goto L87
L108:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v511 = v501
	goto L111
L109:
	;
	if int32(0) <= v607 {
		v1146 = v607
		goto L3
	} else {
		goto L133
	}
L110:
	;
	v607 = v578
	goto L109
L111:
	;
	if v502 <= v511 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v607 = int32(-1)
	goto L109
L114:
	;
	goto L115
L115:
	;
	v518 = int32(1)
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511+v503))))
	if base.Ui32(v520) < base.Ui32(int32(192)) {
		v577 = v520
		v578 = v518
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if int32(117) < v577 {
		goto L110
	} else {
		goto L129
	}
L117:
	;
	v524 = v511 + int32(1)
	if v524 == v502 {
		v577 = v520
		v578 = v518
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524+v503))))
	v529 = v527 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v520) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533+v503))))
	v545 = v543 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v520) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v533 = v511 + int32(2)
	if v533 != v502 {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v577 = v520<<(uint(int32(6))%32)&int32(1984) | v529
	v578 = int32(2)
	goto L116
L123:
	;
	goto L122
L124:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v549))))
	v577 = v562&int32(63) | (v520<<(uint(int32(18))%32)&int32(1835008) | v529<<(uint(int32(12))%32) | v545<<(uint(int32(6))%32))
	v578 = int32(4)
	goto L116
L125:
	;
	v549 = v511 + int32(3)
	if v549 != v502 {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v577 = v520<<(uint(int32(12))%32)&int32(61440) | v529<<(uint(int32(6))%32) | v545
	v578 = int32(3)
	goto L116
L128:
	;
	goto L127
L129:
	;
	v582 = v577 - int32(97)
	if v582 < int32(0) {
		goto L110
	} else {
		goto L130
	}
L130:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v582)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v588)>>(uint(v582&int32(7))%32))&int32(1) == int32(0) {
		goto L110
	} else {
		goto L131
	}
L131:
	;
	v596 = v578 + v511
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v596
	v511 = v596
	goto L111
L133:
	;
	goto L4
L134:
	;
	if v729 != 0 {
		goto L1
	} else {
		goto L159
	}
L135:
	;
	v729 = v722
	goto L134
L136:
	;
	if v625 <= v12 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v722 = int32(0)
	goto L135
L138:
	;
	v729 = int32(-1)
	goto L134
L139:
	;
	goto L140
L140:
	;
	v641 = int32(1)
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v626))))
	if base.Ui32(v643) < base.Ui32(int32(192)) {
		v700 = v643
		v701 = v641
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if int32(117) < v700 {
		goto L154
	} else {
		goto L155
	}
L142:
	;
	v647 = v12 + int32(1)
	if v647 == v625 {
		v700 = v643
		v701 = v641
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647+v626))))
	v652 = v650 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v643) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656+v626))))
	v668 = v666 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v643) {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	v656 = v12 + int32(2)
	if v656 != v625 {
		goto L144
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v700 = v643<<(uint(int32(6))%32)&int32(1984) | v652
	v701 = int32(2)
	goto L141
L148:
	;
	goto L147
L149:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626+v672))))
	v700 = v685&int32(63) | (v643<<(uint(int32(18))%32)&int32(1835008) | v652<<(uint(int32(12))%32) | v668<<(uint(int32(6))%32))
	v701 = int32(4)
	goto L141
L150:
	;
	v672 = v12 + int32(3)
	if v672 != v625 {
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v700 = v643<<(uint(int32(12))%32)&int32(61440) | v652<<(uint(int32(6))%32) | v668
	v701 = int32(3)
	goto L141
L153:
	;
	goto L152
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v701 + v12
	goto L158
L155:
	;
	v705 = v700 - int32(97)
	if v705 < int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v705)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v711)>>(uint(v705&int32(7))%32))&int32(1) != 0 {
		v722 = v701
		goto L135
	} else {
		goto L157
	}
L157:
	;
	goto L154
L158:
	;
	goto L137
L159:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L162
L160:
	;
	if v847 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L161:
	;
	v847 = v840
	goto L160
L162:
	;
	if v743 <= v730 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v840 = int32(0)
	goto L161
L164:
	;
	v847 = int32(-1)
	goto L160
L165:
	;
	goto L166
L166:
	;
	v759 = int32(1)
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730+v744))))
	if base.Ui32(v761) < base.Ui32(int32(192)) {
		v818 = v761
		v819 = v759
		goto L167
	} else {
		goto L168
	}
L167:
	;
	if int32(117) < v818 {
		goto L180
	} else {
		goto L181
	}
L168:
	;
	v765 = v730 + int32(1)
	if v765 == v743 {
		v818 = v761
		v819 = v759
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765+v744))))
	v770 = v768 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v761) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v774+v744))))
	v786 = v784 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v761) {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	v774 = v730 + int32(2)
	if v774 != v743 {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v818 = v761<<(uint(int32(6))%32)&int32(1984) | v770
	v819 = int32(2)
	goto L167
L174:
	;
	goto L173
L175:
	;
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744+v790))))
	v818 = v803&int32(63) | (v761<<(uint(int32(18))%32)&int32(1835008) | v770<<(uint(int32(12))%32) | v786<<(uint(int32(6))%32))
	v819 = int32(4)
	goto L167
L176:
	;
	v790 = v730 + int32(3)
	if v790 != v743 {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v818 = v761<<(uint(int32(12))%32)&int32(61440) | v770<<(uint(int32(6))%32) | v786
	v819 = int32(3)
	goto L167
L179:
	;
	goto L178
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v819 + v730
	goto L184
L181:
	;
	v823 = v818 - int32(97)
	if v823 < int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v823)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v829)>>(uint(v823&int32(7))%32))&int32(1) != 0 {
		v840 = v819
		goto L161
	} else {
		goto L183
	}
L183:
	;
	goto L180
L184:
	;
	goto L163
L185:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v871 = v861
	goto L190
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v730
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L216
L188:
	;
	if int32(0) <= v966 {
		v1146 = v966
		goto L3
	} else {
		goto L213
	}
L189:
	;
	v966 = v938
	goto L188
L190:
	;
	if v862 <= v871 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v966 = int32(-1)
	goto L188
L193:
	;
	goto L194
L194:
	;
	v878 = int32(1)
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871+v863))))
	if base.Ui32(v880) < base.Ui32(int32(192)) {
		v937 = v880
		v938 = v878
		goto L195
	} else {
		goto L196
	}
L195:
	;
	if int32(117) < v937 {
		goto L208
	} else {
		goto L209
	}
L196:
	;
	v884 = v871 + int32(1)
	if v884 == v862 {
		v937 = v880
		v938 = v878
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884+v863))))
	v889 = v887 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v880) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893+v863))))
	v905 = v903 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v880) {
		goto L204
	} else {
		goto L205
	}
L199:
	;
	v893 = v871 + int32(2)
	if v893 != v862 {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v937 = v880<<(uint(int32(6))%32)&int32(1984) | v889
	v938 = int32(2)
	goto L195
L202:
	;
	goto L201
L203:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863+v909))))
	v937 = v922&int32(63) | (v880<<(uint(int32(18))%32)&int32(1835008) | v889<<(uint(int32(12))%32) | v905<<(uint(int32(6))%32))
	v938 = int32(4)
	goto L195
L204:
	;
	v909 = v871 + int32(3)
	if v909 != v862 {
		goto L203
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v937 = v880<<(uint(int32(12))%32)&int32(61440) | v889<<(uint(int32(6))%32) | v905
	v938 = int32(3)
	goto L195
L207:
	;
	goto L206
L208:
	;
	v955 = v938 + v871
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v955
	v871 = v955
	goto L190
L209:
	;
	v942 = v937 - int32(97)
	if v942 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v942)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v948)>>(uint(v942&int32(7))%32))&int32(1) != 0 {
		goto L189
	} else {
		goto L211
	}
L211:
	;
	goto L208
L213:
	;
	goto L187
L214:
	;
	if v1088 != 0 {
		goto L1
	} else {
		goto L238
	}
L215:
	;
	v1088 = v1081
	goto L214
L216:
	;
	if v983 <= v730 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v1081 = int32(0)
	goto L215
L218:
	;
	v1088 = int32(-1)
	goto L214
L219:
	;
	goto L220
L220:
	;
	v999 = int32(1)
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730+v984))))
	if base.Ui32(v1001) < base.Ui32(int32(192)) {
		v1058 = v1001
		v1059 = v999
		goto L221
	} else {
		goto L222
	}
L221:
	;
	if int32(117) < v1058 {
		v1081 = v1059
		goto L215
	} else {
		goto L234
	}
L222:
	;
	v1005 = v730 + int32(1)
	if v1005 == v983 {
		v1058 = v1001
		v1059 = v999
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005+v984))))
	v1010 = v1008 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1001) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014+v984))))
	v1026 = v1024 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1001) {
		goto L230
	} else {
		goto L231
	}
L225:
	;
	v1014 = v730 + int32(2)
	if v1014 != v983 {
		goto L224
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v1058 = v1001<<(uint(int32(6))%32)&int32(1984) | v1010
	v1059 = int32(2)
	goto L221
L228:
	;
	goto L227
L229:
	;
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v984+v1030))))
	v1058 = v1043&int32(63) | (v1001<<(uint(int32(18))%32)&int32(1835008) | v1010<<(uint(int32(12))%32) | v1026<<(uint(int32(6))%32))
	v1059 = int32(4)
	goto L221
L230:
	;
	v1030 = v730 + int32(3)
	if v1030 != v983 {
		goto L229
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v1058 = v1001<<(uint(int32(12))%32)&int32(61440) | v1010<<(uint(int32(6))%32) | v1026
	v1059 = int32(3)
	goto L221
L233:
	;
	goto L232
L234:
	;
	v1063 = v1058 - int32(97)
	if v1063 < int32(0) {
		v1081 = v1059
		goto L215
	} else {
		goto L235
	}
L235:
	;
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1063)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v1069)>>(uint(v1063&int32(7))%32))&int32(1) == int32(0) {
		v1081 = v1059
		goto L215
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1059 + v730
	goto L237
L237:
	;
	goto L217
L238:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L241
L239:
	;
	if int32(0) <= v1143 {
		v1150 = v1143
		goto L2
	} else {
		goto L259
	}
L241:
	;
	goto L242
L242:
	;
	goto L243
L243:
	;
	v1098 = v1090
	v1100 = int32(1)
	goto L246
L245:
	;
	v1143 = v1128
	goto L239
L246:
	;
	if v1091 <= v1098 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	goto L245
L248:
	;
	v1143 = int32(-1)
	goto L239
L249:
	;
	goto L250
L250:
	;
	v1105 = v1098 + int32(1)
	v1107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089+v1098))))
	if base.Ui32(v1107) < base.Ui32(int32(192)) {
		v1128 = v1105
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1129 = int32(1)
	if v1129 < v1100 {
		v1098 = v1128
		v1100 = v1100 - v1129
		goto L246
	} else {
		goto L258
	}
L252:
	;
	if v1091 <= v1105 {
		v1128 = v1105
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v1114 = v1105
	goto L254
L254:
	;
	v1117 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1089+v1114))))
	if int32(-65) < v1117 {
		v1128 = v1114
		goto L251
	} else {
		goto L256
	}
L255:
	;
	v1128 = v1091
	goto L251
L256:
	;
	v1121 = v1114 + int32(1)
	if v1121 != v1091 {
		v1114 = v1121
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	goto L247
L259:
	;
	goto L1
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v12
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1652
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1652
	v1656 = v1652 - int32(1)
	if v1656 <= v12 {
		goto L365
	} else {
		goto L366
	}
L261:
	;
	if v1273 < int32(0) {
		goto L260
	} else {
		goto L286
	}
L262:
	;
	v1273 = v1245
	goto L261
L263:
	;
	if v1169 <= v1178 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1273 = int32(-1)
	goto L261
L266:
	;
	goto L267
L267:
	;
	v1185 = int32(1)
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1178+v1170))))
	if base.Ui32(v1187) < base.Ui32(int32(192)) {
		v1244 = v1187
		v1245 = v1185
		goto L268
	} else {
		goto L269
	}
L268:
	;
	if int32(117) < v1244 {
		goto L281
	} else {
		goto L282
	}
L269:
	;
	v1191 = v1178 + int32(1)
	if v1191 == v1169 {
		v1244 = v1187
		v1245 = v1185
		goto L268
	} else {
		goto L270
	}
L270:
	;
	v1194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191+v1170))))
	v1196 = v1194 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1187) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200+v1170))))
	v1212 = v1210 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1187) {
		goto L277
	} else {
		goto L278
	}
L272:
	;
	v1200 = v1178 + int32(2)
	if v1200 != v1169 {
		goto L271
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1244 = v1187<<(uint(int32(6))%32)&int32(1984) | v1196
	v1245 = int32(2)
	goto L268
L275:
	;
	goto L274
L276:
	;
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1170+v1216))))
	v1244 = v1229&int32(63) | (v1187<<(uint(int32(18))%32)&int32(1835008) | v1196<<(uint(int32(12))%32) | v1212<<(uint(int32(6))%32))
	v1245 = int32(4)
	goto L268
L277:
	;
	v1216 = v1178 + int32(3)
	if v1216 != v1169 {
		goto L276
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1244 = v1187<<(uint(int32(12))%32)&int32(61440) | v1196<<(uint(int32(6))%32) | v1212
	v1245 = int32(3)
	goto L268
L280:
	;
	goto L279
L281:
	;
	v1262 = v1245 + v1178
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1262
	v1178 = v1262
	goto L263
L282:
	;
	v1249 = v1244 - int32(97)
	if v1249 < int32(0) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1249)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v1255)>>(uint(v1249&int32(7))%32))&int32(1) != 0 {
		goto L262
	} else {
		goto L284
	}
L284:
	;
	goto L281
L286:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1277 = v1276 + v1273
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1277
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1300 = v1277
	goto L289
L287:
	;
	if v1396 < int32(0) {
		goto L260
	} else {
		goto L311
	}
L288:
	;
	v1396 = v1367
	goto L287
L289:
	;
	if v1291 <= v1300 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1396 = int32(-1)
	goto L287
L292:
	;
	goto L293
L293:
	;
	v1307 = int32(1)
	v1309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300+v1292))))
	if base.Ui32(v1309) < base.Ui32(int32(192)) {
		v1366 = v1309
		v1367 = v1307
		goto L294
	} else {
		goto L295
	}
L294:
	;
	if int32(117) < v1366 {
		goto L288
	} else {
		goto L307
	}
L295:
	;
	v1313 = v1300 + int32(1)
	if v1313 == v1291 {
		v1366 = v1309
		v1367 = v1307
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1313+v1292))))
	v1318 = v1316 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1309) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322+v1292))))
	v1334 = v1332 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1309) {
		goto L303
	} else {
		goto L304
	}
L298:
	;
	v1322 = v1300 + int32(2)
	if v1322 != v1291 {
		goto L297
	} else {
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	v1366 = v1309<<(uint(int32(6))%32)&int32(1984) | v1318
	v1367 = int32(2)
	goto L294
L301:
	;
	goto L300
L302:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292+v1338))))
	v1366 = v1351&int32(63) | (v1309<<(uint(int32(18))%32)&int32(1835008) | v1318<<(uint(int32(12))%32) | v1334<<(uint(int32(6))%32))
	v1367 = int32(4)
	goto L294
L303:
	;
	v1338 = v1300 + int32(3)
	if v1338 != v1291 {
		goto L302
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1366 = v1309<<(uint(int32(12))%32)&int32(61440) | v1318<<(uint(int32(6))%32) | v1334
	v1367 = int32(3)
	goto L294
L306:
	;
	goto L305
L307:
	;
	v1371 = v1366 - int32(97)
	if v1371 < int32(0) {
		goto L288
	} else {
		goto L308
	}
L308:
	;
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1371)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v1377)>>(uint(v1371&int32(7))%32))&int32(1) == int32(0) {
		goto L288
	} else {
		goto L309
	}
L309:
	;
	v1385 = v1367 + v1300
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1385
	v1300 = v1385
	goto L289
L311:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1400 = v1399 + v1396
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1400
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1402)+4)) = v1400
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1425 = v1415
	goto L314
L312:
	;
	if v1520 < int32(0) {
		goto L260
	} else {
		goto L337
	}
L313:
	;
	v1520 = v1492
	goto L312
L314:
	;
	if v1416 <= v1425 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1520 = int32(-1)
	goto L312
L317:
	;
	goto L318
L318:
	;
	v1432 = int32(1)
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1425+v1417))))
	if base.Ui32(v1434) < base.Ui32(int32(192)) {
		v1491 = v1434
		v1492 = v1432
		goto L319
	} else {
		goto L320
	}
L319:
	;
	if int32(117) < v1491 {
		goto L332
	} else {
		goto L333
	}
L320:
	;
	v1438 = v1425 + int32(1)
	if v1438 == v1416 {
		v1491 = v1434
		v1492 = v1432
		goto L319
	} else {
		goto L321
	}
L321:
	;
	v1441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1438+v1417))))
	v1443 = v1441 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1434) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v1457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1447+v1417))))
	v1459 = v1457 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1434) {
		goto L328
	} else {
		goto L329
	}
L323:
	;
	v1447 = v1425 + int32(2)
	if v1447 != v1416 {
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1491 = v1434<<(uint(int32(6))%32)&int32(1984) | v1443
	v1492 = int32(2)
	goto L319
L326:
	;
	goto L325
L327:
	;
	v1476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1417+v1463))))
	v1491 = v1476&int32(63) | (v1434<<(uint(int32(18))%32)&int32(1835008) | v1443<<(uint(int32(12))%32) | v1459<<(uint(int32(6))%32))
	v1492 = int32(4)
	goto L319
L328:
	;
	v1463 = v1425 + int32(3)
	if v1463 != v1416 {
		goto L327
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v1491 = v1434<<(uint(int32(12))%32)&int32(61440) | v1443<<(uint(int32(6))%32) | v1459
	v1492 = int32(3)
	goto L319
L331:
	;
	goto L330
L332:
	;
	v1509 = v1492 + v1425
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1509
	v1425 = v1509
	goto L314
L333:
	;
	v1496 = v1491 - int32(97)
	if v1496 < int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1496)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v1502)>>(uint(v1496&int32(7))%32))&int32(1) != 0 {
		goto L313
	} else {
		goto L335
	}
L335:
	;
	goto L332
L337:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1524 = v1523 + v1520
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1524
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1547 = v1524
	goto L340
L338:
	;
	if v1643 < int32(0) {
		goto L260
	} else {
		goto L362
	}
L339:
	;
	v1643 = v1614
	goto L338
L340:
	;
	if v1538 <= v1547 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1643 = int32(-1)
	goto L338
L343:
	;
	goto L344
L344:
	;
	v1554 = int32(1)
	v1556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1547+v1539))))
	if base.Ui32(v1556) < base.Ui32(int32(192)) {
		v1613 = v1556
		v1614 = v1554
		goto L345
	} else {
		goto L346
	}
L345:
	;
	if int32(117) < v1613 {
		goto L339
	} else {
		goto L358
	}
L346:
	;
	v1560 = v1547 + int32(1)
	if v1560 == v1538 {
		v1613 = v1556
		v1614 = v1554
		goto L345
	} else {
		goto L347
	}
L347:
	;
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1560+v1539))))
	v1565 = v1563 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1556) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	v1579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1569+v1539))))
	v1581 = v1579 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1556) {
		goto L354
	} else {
		goto L355
	}
L349:
	;
	v1569 = v1547 + int32(2)
	if v1569 != v1538 {
		goto L348
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1613 = v1556<<(uint(int32(6))%32)&int32(1984) | v1565
	v1614 = int32(2)
	goto L345
L352:
	;
	goto L351
L353:
	;
	v1598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1539+v1585))))
	v1613 = v1598&int32(63) | (v1556<<(uint(int32(18))%32)&int32(1835008) | v1565<<(uint(int32(12))%32) | v1581<<(uint(int32(6))%32))
	v1614 = int32(4)
	goto L345
L354:
	;
	v1585 = v1547 + int32(3)
	if v1585 != v1538 {
		goto L353
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1613 = v1556<<(uint(int32(12))%32)&int32(61440) | v1565<<(uint(int32(6))%32) | v1581
	v1614 = int32(3)
	goto L345
L357:
	;
	goto L356
L358:
	;
	v1618 = v1613 - int32(97)
	if v1618 < int32(0) {
		goto L339
	} else {
		goto L359
	}
L359:
	;
	v1624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1618)>>(uint(int32(3))%32)))+uint32(_consts[1065]))))
	if int32(base.Ui32(v1624)>>(uint(v1618&int32(7))%32))&int32(1) == int32(0) {
		goto L339
	} else {
		goto L360
	}
L360:
	;
	v1632 = v1614 + v1547
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1632
	v1547 = v1632
	goto L340
L362:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1646))) = v1647 + v1643
	goto L260
L363:
	;
	return v1895
L364:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1735 = v1729 - v1731 + v1734
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1735
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1735
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1735 <= v1738 {
		v1840 = v1735
		v1842 = v1734
		goto L394
	} else {
		goto L395
	}
L365:
	;
	v1729 = v1652
	v1731 = v1652
	goto L364
L366:
	;
	goto L367
L367:
	;
	v1659 = v1652
	v1660 = v1656
	v1661 = v1652
	goto L368
L368:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1663+v1660))))
	if v1665&int32(224) != int32(96) {
		v1729 = v1659
		v1731 = v1661
		goto L364
	} else {
		goto L370
	}
L369:
	;
	v1729 = v1721
	v1731 = v1723
	goto L364
L370:
	;
	if int32(1)<<(uint(v1665)%32)&int32(70566434) == int32(0) {
		v1729 = v1659
		v1731 = v1661
		goto L364
	} else {
		goto L371
	}
L371:
	;
	v1678 = F_find_among_b(m, l0, int32(4271888), int32(109))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	return int32(0)
L373:
	;
	if v1678 == int32(0) {
		v1729 = v1659
		v1731 = v1661
		goto L364
	} else {
		goto L374
	}
L374:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1684
	switch v1678 - int32(1) {
	case 0:
		goto L380
	case 1:
		goto L379
	case 2:
		goto L378
	case 3:
		goto L377
	case 4:
		goto L376
	default:
		goto L375
	}
L375:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1721
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1725 = v1721 - int32(1)
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1726 < v1725 {
		v1659 = v1721
		v1660 = v1725
		v1661 = v1723
		goto L368
	} else {
		goto L393
	}
L376:
	;
	v1716 = F_slice_from_s(m, l0, int32(6), int32(2217984))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L372
	} else {
		goto L391
	}
L377:
	;
	v1710 = F_slice_from_s(m, l0, int32(7), int32(2217977))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L372
	} else {
		goto L389
	}
L378:
	;
	v1704 = F_slice_from_s(m, l0, int32(7), int32(2217970))
	mBase = m.M
	v1705 = m.ExcPending
	if v1705 != 0 {
		goto L372
	} else {
		goto L387
	}
L379:
	;
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1695)))
	if v1684 < v1696 {
		v1729 = v1659
		v1731 = v1661
		goto L364
	} else {
		goto L384
	}
L380:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+8))
	if v1684 < v1689 {
		v1729 = v1659
		v1731 = v1661
		goto L364
	} else {
		goto L381
	}
L381:
	;
	v1691 = F_slice_del(m, l0)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L372
	} else {
		goto L382
	}
L382:
	;
	if int32(0) <= v1691 {
		goto L375
	} else {
		goto L383
	}
L383:
	;
	v1895 = v1691
	goto L363
L384:
	;
	v1698 = F_slice_del(m, l0)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L372
	} else {
		goto L385
	}
L385:
	;
	if int32(0) <= v1698 {
		goto L375
	} else {
		goto L386
	}
L386:
	;
	v1895 = v1698
	goto L363
L387:
	;
	if int32(0) <= v1704 {
		goto L375
	} else {
		goto L388
	}
L388:
	;
	v1895 = v1704
	goto L363
L389:
	;
	if int32(0) <= v1710 {
		goto L375
	} else {
		goto L390
	}
L390:
	;
	v1895 = v1710
	goto L363
L391:
	;
	if v1716 < int32(0) {
		v1895 = v1716
		goto L363
	} else {
		goto L392
	}
L392:
	;
	goto L375
L393:
	;
	goto L369
L394:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1846 = v1844 + (v1840 - v1842)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1846
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1846
	v1850 = v1846 - int32(1)
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1850 <= v1851 {
		goto L437
	} else {
		goto L438
	}
L395:
	;
	v1741 = v1735
	v1743 = v1734
	goto L396
L396:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1745+v1741-int32(1)))))
	if v1749&int32(224) != int32(96) {
		v1840 = v1741
		v1842 = v1743
		goto L394
	} else {
		goto L398
	}
L397:
	;
	v1840 = v1834
	v1842 = v1836
	goto L394
L398:
	;
	if int32(1)<<(uint(v1749)%32)&int32(71162402) == int32(0) {
		v1840 = v1741
		v1842 = v1743
		goto L394
	} else {
		goto L399
	}
L399:
	;
	v1762 = F_find_among_b(m, l0, int32(4274080), int32(295))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L372
	} else {
		goto L400
	}
L400:
	;
	if v1762 == int32(0) {
		v1840 = v1741
		v1842 = v1743
		goto L394
	} else {
		goto L401
	}
L401:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1766
	switch v1762 - int32(1) {
	case 0:
		goto L412
	case 1:
		goto L411
	case 2:
		goto L410
	case 3:
		goto L409
	case 4:
		goto L408
	case 5:
		goto L407
	case 6:
		goto L406
	case 7:
		goto L405
	case 8:
		goto L404
	case 9:
		goto L403
	default:
		goto L402
	}
L402:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1834
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1837 < v1834 {
		v1741 = v1834
		v1743 = v1836
		goto L396
	} else {
		goto L436
	}
L403:
	;
	v1829 = F_slice_from_s(m, l0, int32(5), int32(2218484))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L372
	} else {
		goto L434
	}
L404:
	;
	v1823 = F_slice_from_s(m, l0, int32(5), int32(2218479))
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L372
	} else {
		goto L432
	}
L405:
	;
	v1817 = F_slice_from_s(m, l0, int32(5), int32(2218474))
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L372
	} else {
		goto L430
	}
L406:
	;
	v1811 = F_slice_from_s(m, l0, int32(5), int32(2218469))
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L372
	} else {
		goto L428
	}
L407:
	;
	v1805 = F_slice_from_s(m, l0, int32(6), int32(2218463))
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L372
	} else {
		goto L426
	}
L408:
	;
	v1799 = F_slice_from_s(m, l0, int32(3), int32(2218460))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L372
	} else {
		goto L424
	}
L409:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1790)+4))
	if v1766 < v1791 {
		v1840 = v1741
		v1842 = v1743
		goto L394
	} else {
		goto L421
	}
L410:
	;
	v1786 = F_slice_from_s(m, l0, int32(3), int32(2218457))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L372
	} else {
		goto L419
	}
L411:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1777)))
	if v1766 < v1778 {
		v1840 = v1741
		v1842 = v1743
		goto L394
	} else {
		goto L416
	}
L412:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1770)+8))
	if v1766 < v1771 {
		v1840 = v1741
		v1842 = v1743
		goto L394
	} else {
		goto L413
	}
L413:
	;
	v1773 = F_slice_del(m, l0)
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L372
	} else {
		goto L414
	}
L414:
	;
	if int32(0) <= v1773 {
		goto L402
	} else {
		goto L415
	}
L415:
	;
	v1895 = v1773
	goto L363
L416:
	;
	v1780 = F_slice_del(m, l0)
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L372
	} else {
		goto L417
	}
L417:
	;
	if int32(0) <= v1780 {
		goto L402
	} else {
		goto L418
	}
L418:
	;
	v1895 = v1780
	goto L363
L419:
	;
	if int32(0) <= v1786 {
		goto L402
	} else {
		goto L420
	}
L420:
	;
	v1895 = v1786
	goto L363
L421:
	;
	v1793 = F_slice_del(m, l0)
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L372
	} else {
		goto L422
	}
L422:
	;
	if int32(0) <= v1793 {
		goto L402
	} else {
		goto L423
	}
L423:
	;
	v1895 = v1793
	goto L363
L424:
	;
	if int32(0) <= v1799 {
		goto L402
	} else {
		goto L425
	}
L425:
	;
	v1895 = v1799
	goto L363
L426:
	;
	if int32(0) <= v1805 {
		goto L402
	} else {
		goto L427
	}
L427:
	;
	v1895 = v1805
	goto L363
L428:
	;
	if int32(0) <= v1811 {
		goto L402
	} else {
		goto L429
	}
L429:
	;
	v1895 = v1811
	goto L363
L430:
	;
	if int32(0) <= v1817 {
		goto L402
	} else {
		goto L431
	}
L431:
	;
	v1895 = v1817
	goto L363
L432:
	;
	if int32(0) <= v1823 {
		goto L402
	} else {
		goto L433
	}
L433:
	;
	v1895 = v1823
	goto L363
L434:
	;
	if v1829 < int32(0) {
		v1895 = v1829
		goto L363
	} else {
		goto L435
	}
L435:
	;
	goto L402
L436:
	;
	goto L397
L437:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1891
	v1895 = int32(1)
	goto L363
L438:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1853+v1850))))
	if v1855&int32(224) != int32(96) {
		goto L437
	} else {
		goto L439
	}
L439:
	;
	if int32(1)<<(uint(v1855)%32)&int32(35362) == int32(0) {
		goto L437
	} else {
		goto L440
	}
L440:
	;
	v1868 = F_find_among_b(m, l0, int32(4279984), int32(19))
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L372
	} else {
		goto L441
	}
L441:
	;
	if v1868 == int32(0) {
		goto L437
	} else {
		goto L442
	}
L442:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1872
	switch v1868 - int32(1) {
	case 0:
		goto L444
	case 1:
		goto L443
	default:
		goto L437
	}
L443:
	;
	v1885 = F_slice_from_s(m, l0, int32(1), int32(2219679))
	mBase = m.M
	v1886 = m.ExcPending
	if v1886 != 0 {
		goto L372
	} else {
		goto L448
	}
L444:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+8))
	if v1872 < v1877 {
		goto L437
	} else {
		goto L445
	}
L445:
	;
	v1879 = F_slice_del(m, l0)
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L372
	} else {
		goto L446
	}
L446:
	;
	if int32(0) <= v1879 {
		goto L437
	} else {
		goto L447
	}
L447:
	;
	v1895 = v1879
	goto L363
L448:
	;
	if v1885 < int32(0) {
		v1895 = v1885
		goto L363
	} else {
		goto L449
	}
L449:
	;
	goto L437
}
func F_bernoulli_initsamplescan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_palloc0(m, int32(16))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v4
		return
	}
}
func F_bernoulli_nextsampletuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int64
	_ = v294
	var v297 int32
	_ = v297
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v13
	v15 = v11
	goto L1
L1:
	;
	v21 = v15 + int32(1)
	v23 = v21 & int32(65535)
	if base.Ui32(l2) < base.Ui32(v23) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v297)
	m.G0 = v8 + int32(16)
	return v297 & int32(65535)
L3:
	;
	goto L2
L4:
	;
	v297 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v23
	v28 = v8 + int32(4)
	v29 = int32(12)
	v35 = int32(-1636608420)
	if v28&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui64(v294) <= base.Ui64(base.I64_extend_i32_u(v289^v281-base.I32_rotl(v289, int32(24)))) {
		v15 = v21
		goto L1
	} else {
		goto L47
	}
L8:
	;
	v267 = int32(14)
	v269 = v263 ^ v264 - base.I32_rotl(v263, v267)
	v273 = v269 ^ v262 - base.I32_rotl(v269, int32(11))
	v277 = v273 ^ v263 - base.I32_rotl(v273, int32(25))
	v281 = v277 ^ v269 - base.I32_rotl(v277, int32(16))
	v285 = v281 ^ v273 - base.I32_rotl(v281, int32(4))
	v289 = v285 ^ v277 - base.I32_rotl(v285, v267)
	goto L7
L9:
	;
	switch v189 - int32(1) {
	case 0:
		v255 = v180
		v256 = v181
		v257 = v185
		goto L36
	case 1:
		v248 = v180
		v249 = v181
		v250 = v185
		goto L37
	case 2:
		v241 = v180
		v242 = v181
		v243 = v185
		goto L38
	case 3:
		v235 = v181
		v236 = v185
		goto L39
	case 4:
		v231 = v181
		v232 = v185
		goto L40
	case 5:
		v225 = v181
		v226 = v185
		goto L41
	case 6:
		v219 = v181
		v220 = v185
		goto L42
	case 7:
		v214 = v185
		goto L43
	case 8:
		v209 = v185
		goto L44
	case 9:
		v204 = v185
		goto L45
	case 10:
		goto L46
	default:
		v262 = v180
		v263 = v181
		v264 = v185
		goto L8
	}
L10:
	;
	v144 = v28
	v145 = v29
	v146 = v35
	v147 = v35
	v148 = v35
	goto L33
L11:
	;
	goto L10
L12:
	;
	goto L13
L13:
	;
	goto L17
L15:
	;
	switch v87 - int32(1) {
	case 0:
		v141 = v78
		goto L22
	case 1:
		v136 = v78
		goto L23
	case 2:
		goto L24
	case 3:
		v129 = v79
		goto L25
	case 4:
		v126 = v79
		goto L26
	case 5:
		v121 = v79
		goto L27
	case 6:
		goto L28
	case 7:
		v112 = v83
		goto L29
	case 8:
		v107 = v83
		goto L30
	case 9:
		v102 = v83
		goto L31
	case 10:
		goto L32
	default:
		v262 = v78
		v263 = v79
		v264 = v83
		goto L8
	}
L17:
	;
	goto L18
L18:
	;
	v42 = v28
	v43 = v29
	v44 = v35
	v45 = v35
	v46 = v35
	goto L19
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v49 = v48 + v45
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v53 = v52 + v46
	v55 = int32(4)
	v57 = v50 + v44 - v53 ^ base.I32_rotl(v53, v55)
	v61 = v49 - v57 ^ base.I32_rotl(v57, int32(6))
	v62 = v53 + v49
	v63 = v57 + v62
	v64 = v61 + v63
	v68 = v62 - v61 ^ base.I32_rotl(v61, int32(8))
	v72 = v63 - v68 ^ base.I32_rotl(v68, int32(16))
	v76 = v64 - v72 ^ base.I32_rotl(v72, int32(19))
	v77 = v68 + v64
	v78 = v72 + v77
	v79 = v76 + v78
	v83 = v77 - v76 ^ base.I32_rotl(v76, v55)
	v84 = int32(12)
	v85 = v42 + v84
	v87 = v43 - v84
	if base.Ui32(int32(11)) < base.Ui32(v87) {
		v42 = v85
		v43 = v87
		v44 = v78
		v45 = v79
		v46 = v83
		goto L19
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	goto L20
L22:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v262 = v141 + v142
	v263 = v79
	v264 = v83
	goto L8
L23:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	v141 = v137<<(uint(int32(8))%32) + v136
	goto L22
L24:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+2)))
	v136 = v132<<(uint(int32(16))%32) + v78
	goto L23
L25:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v262 = v130 + v78
	v263 = v129
	v264 = v83
	goto L8
L26:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	v129 = v126 + v127
	goto L25
L27:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+5)))
	v126 = v122<<(uint(int32(8))%32) + v121
	goto L26
L28:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+6)))
	v121 = v117<<(uint(int32(16))%32) + v79
	goto L27
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v262 = v113 + v78
	v263 = v115 + v79
	v264 = v112
	goto L8
L30:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+8)))
	v112 = v108<<(uint(int32(8))%32) + v107
	goto L29
L31:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+9)))
	v107 = v103<<(uint(int32(16))%32) + v102
	goto L30
L32:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+10)))
	v102 = v98<<(uint(int32(24))%32) + v83
	goto L31
L33:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v151 = v150 + v147
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	v155 = v154 + v148
	v157 = int32(4)
	v159 = v152 + v146 - v155 ^ base.I32_rotl(v155, v157)
	v163 = v151 - v159 ^ base.I32_rotl(v159, int32(6))
	v164 = v155 + v151
	v165 = v159 + v164
	v166 = v163 + v165
	v170 = v164 - v163 ^ base.I32_rotl(v163, int32(8))
	v174 = v165 - v170 ^ base.I32_rotl(v170, int32(16))
	v178 = v166 - v174 ^ base.I32_rotl(v174, int32(19))
	v179 = v170 + v166
	v180 = v174 + v179
	v181 = v178 + v180
	v185 = v179 - v178 ^ base.I32_rotl(v178, v157)
	v186 = int32(12)
	v187 = v144 + v186
	v189 = v145 - v186
	if base.Ui32(int32(11)) < base.Ui32(v189) {
		v144 = v187
		v145 = v189
		v146 = v180
		v147 = v181
		v148 = v185
		goto L33
	} else {
		goto L35
	}
L34:
	;
	goto L9
L35:
	;
	goto L34
L36:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v262 = v255 + v258
	v263 = v256
	v264 = v257
	goto L8
L37:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	v255 = v251<<(uint(int32(8))%32) + v248
	v256 = v249
	v257 = v250
	goto L36
L38:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
	v248 = v244<<(uint(int32(16))%32) + v241
	v249 = v242
	v250 = v243
	goto L37
L39:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
	v241 = v237<<(uint(int32(24))%32) + v180
	v242 = v235
	v243 = v236
	goto L38
L40:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
	v235 = v231 + v233
	v236 = v232
	goto L39
L41:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
	v231 = v227<<(uint(int32(8))%32) + v225
	v232 = v226
	goto L40
L42:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
	v225 = v221<<(uint(int32(16))%32) + v219
	v226 = v220
	goto L41
L43:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+7)))
	v219 = v215<<(uint(int32(24))%32) + v181
	v220 = v214
	goto L42
L44:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
	v214 = v210<<(uint(int32(8))%32) + v209
	goto L43
L45:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
	v209 = v205<<(uint(int32(16))%32) + v204
	goto L44
L46:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+10)))
	v204 = v200<<(uint(int32(24))%32) + v185
	goto L45
L47:
	;
	v297 = v21
	goto L3
}
func F_bgworker_die(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_sigprocmask(m, int32(4456376), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errcode(m, int32(16908741))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[83]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v18 + int32(96)
				F_errmsg(m, int32(445506), v4)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_errfinish(m, int32(515389), int32(711), int32(414852))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
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
func F_bitgetbit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = int32(0)
		if base.B2i32(v16 <= v15)&base.B2i32(v15 < v14) == v16 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v14 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v15
					F_errmsg(m, int32(703972), v7)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512683), int32(1883), int32(109246))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+int32(base.Ui32(v15)>>(uint(int32(3))%32)))+8)))
			m.G0 = v7 + int32(16)
			return int32(base.Ui32(v44)>>(uint((v15^int32(-1))&int32(7))%32)) & int32(1)
		}
	}
}
func F_bitne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 == v16 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = int32(8)
	v19 = v7 + v18
	v21 = v12 + v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v23 = int32(2)
	v24 = int32(base.Ui32(v22) >> (uint(v23) % 32))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v27 = int32(base.Ui32(v25) >> (uint(v23) % 32))
	if base.Ui32(v24) < base.Ui32(v27) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v96 = int32(1)
	goto L6
L6:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v98 != v7 {
		goto L28
	} else {
		goto L29
	}
L7:
	;
	v29 = v24
	goto L9
L8:
	;
	v29 = v27
	goto L9
L9:
	;
	v31 = v29 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v31) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v96 = base.B2i32(v93 != int32(0))
	goto L6
L11:
	;
	v93 = int32(0)
	goto L10
L12:
	;
	v67 = v62
	v68 = v63
	v69 = v64
	goto L22
L13:
	;
	if (v19|v21)&int32(3) != 0 {
		v62 = v19
		v63 = v21
		v64 = v31
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v55 = v19
	v56 = v21
	v57 = v31
	goto L15
L15:
	;
	if v57 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L16:
	;
	v39 = v19
	v40 = v21
	v41 = v31
	goto L17
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v44 != v45 {
		v62 = v39
		v63 = v40
		v64 = v41
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v55 = v50
	v56 = v48
	v57 = v52
	goto L15
L19:
	;
	v47 = int32(4)
	v48 = v40 + v47
	v50 = v39 + v47
	v52 = v41 - v47
	if base.Ui32(int32(3)) < base.Ui32(v52) {
		v39 = v50
		v40 = v48
		v41 = v52
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v62 = v55
	v63 = v56
	v64 = v57
	goto L12
L22:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v72 == v73 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v93 = v72 - v73
	goto L10
L24:
	;
	v75 = int32(1)
	v80 = v69 - v75
	if v80 != 0 {
		v67 = v67 + v75
		v68 = v68 + v75
		v69 = v80
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L11
L28:
	;
	F_pfree(m, v7)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v102 != v12 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	F_pfree(m, v12)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	return v96
L35:
	;
	goto L34
}
func F_bitposition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v109 int32
	_ = v109
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	v2 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v32 == int32(0) {
		v220 = v2
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v220
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v32 < v35 {
		v220 = v2
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v35 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(1)
L8:
	;
	goto L9
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v42 = int32(2)
	v43 = int32(base.Ui32(v41) >> (uint(v42) % 32))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v46 = int32(base.Ui32(v44) >> (uint(v42) % 32))
	v47 = v43 - v46
	if v47 == int32(-1) {
		v220 = v2
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v50 = int32(255)
	v51 = int32(1)
	v53 = int32(-8)
	v56 = int32(-64)
	v58 = v50 << (uint(v44<<(uint(v51)%32)&v53-v35+v56) % 32)
	v61 = int32(8)
	v65 = v25 + v43
	v67 = v65 - v51
	v68 = v30 + v46
	v79 = v50 << (uint(v41<<(uint(v51)%32)&v53-v32+v56) % 32)
	v81 = v79 ^ int32(-1)
	v88 = v2
	goto L11
L11:
	;
	v109 = int32(0)
	goto L13
L12:
	;
	v220 = int32(0)
	goto L4
L13:
	;
	v136 = int32(8) - v109
	v137 = v58 << (uint(v136) % 32)
	v138 = v30 + v61
	v139 = v88 + (v25 + v61)
	v141 = int32(base.Ui32(int32(255)) >> (uint(v109) % 32))
	v142 = int32(-256) >> (uint(v109) % 32)
	goto L15
L14:
	;
	if v88 != v47 {
		v88 = v88 + int32(1)
		goto L11
	} else {
		goto L35
	}
L15:
	;
	if base.Ui32(v138) < base.Ui32(v68) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v210 = v109 + int32(1)
	if v210 != int32(8) {
		v109 = v210
		goto L13
	} else {
		goto L34
	}
L17:
	;
	goto L16
L18:
	;
	if v138 != v68-v51 {
		v194 = v142
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v163 = base.B2i32(v138 != v68-v51)
	if v138 != v68-v51 {
		v168 = v141
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	return v109 + v88<<(uint(int32(3))%32) + int32(1)
L22:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v168&(v169^int32(base.Ui32(v162)>>(uint(v109)%32))) != 0 {
		goto L17
	} else {
		goto L26
	}
L23:
	;
	v164 = v141 & int32(base.Ui32(v58&v50)>>(uint(v109)%32))
	if v139 != v67 {
		v168 = v164
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v164&v81 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v168 = v164 & v79
	goto L22
L26:
	;
	v174 = v139 + int32(1)
	if v174 != v65 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	if v137&int32(254) != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	goto L21
L29:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	if v194&(v197^v162<<(uint(v136)%32))&int32(255) == int32(0) {
		v138 = v138 + int32(1)
		v139 = v174
		v141 = v168
		v142 = v194
		goto L15
	} else {
		goto L33
	}
L30:
	;
	v188 = v142 & v137
	if v174 != v67 {
		v194 = v188
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if v188&v81&int32(255) != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v194 = v188 & v79
	goto L29
L33:
	;
	goto L17
L34:
	;
	goto L14
L35:
	;
	goto L12
}
func F_bitsetbit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v18 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v17 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
					F_errmsg(m, int32(703972), v10)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512683), int32(1825), int32(109236))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if v17 <= v18 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v17 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
						F_errmsg(m, int32(703972), v10)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512683), int32(1825), int32(109236))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if base.Ui32(int32(2)) <= base.Ui32(v22) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(584578), int32(0))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(512683), int32(1833), int32(109236))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v28 = F_palloc(m, int32(base.Ui32(v25)>>(uint(int32(2))%32)))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v25 & int32(-4)
						v34 = int32(8)
						v35 = v28 + v34
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v42 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - v34
						if v42 != 0 {
							v43 = F__emscripten_memcpy_bulkmem(m, v35, v13+v34, v42)
							mBase = m.M
							v44 = v43
						} else {
							v44 = v35
						}
						v47 = v44 + int32(base.Ui32(v18)>>(uint(int32(3))%32))
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
						v53 = (v18 ^ int32(-1)) & int32(7)
						if v22 != 0 {
							v59 = v48 | int32(1)<<(uint(v53)%32)
						} else {
							v59 = v48 & base.I32_rotl(int32(-2), v53)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v59)
						m.G0 = v10 + int32(16)
						return v28
					}
				}
			}
		}
	}
}
func F_bitsubstr_no_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_bitsubstring(m, v3, v7, int32(-1), int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_bitxor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v15 == v16 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v21 = F_palloc(m, int32(base.Ui32(v18)>>(uint(int32(2))%32)))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v15
					v24 = int32(-4)
					*(*int32)(unsafe.Add(mBase, uint32(v21))) = v18 & v24
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v27&v24 != int32(32) {
						v32 = int32(8)
						v38 = v8 + v32
						v39 = v13 + v32
						v41 = v21 + v32
						v43 = int32(0)
						for {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
							v46 = v44 ^ v45
							*(*uint8)(unsafe.Add(mBase, uint32(v41))) = uint8(v46)
							v48 = int32(1)
							v55 = v43 + v48
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							if base.Ui32(v55) < base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))-int32(8)) {
								v38 = v38 + v48
								v39 = v39 + v48
								v41 = v41 + v48
								v43 = v55
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return v21
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(101187714))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(165374), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512683), int32(1342), int32(216681))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
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
}
func F_blackhole_get_sink(m *base.Module, l0 int32, l1 int32) int32 {
	return l0
}
func F_blbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v4 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(1172))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			F_initBloomState(m, v9+int32(4), v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v9
				return v4
			}
		}
	}
}
func F_blendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 != 0 {
		F_pfree(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(0)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(0)
		return
	}
}
func F_blockreftable_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v307 int64
	_ = v307
	var v310 int32
	_ = v310
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v328 int64
	_ = v328
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int64
	_ = v345
	var v355 int64
	_ = v355
	var v370 float64
	_ = v370
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v387 int32
	_ = v387
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
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
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
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
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int64
	_ = v987
	var v989 int64
	_ = v989
	var v991 int64
	_ = v991
	var v993 int64
	_ = v993
	var v995 int64
	_ = v995
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1106 int32
	_ = v1106
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1459 int32
	_ = v1459
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1473 int64
	_ = v1473
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1524 int64
	_ = v1524
	var v1526 int64
	_ = v1526
	var v1528 int64
	_ = v1528
	var v1530 int64
	_ = v1530
	var v1532 int64
	_ = v1532
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1554 int64
	_ = v1554
	var v1556 int64
	_ = v1556
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1566 int64
	_ = v1566
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1588 int32
	_ = v1588
	var v1595 int32
	_ = v1595
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1608 int64
	_ = v1608
	var v1610 int64
	_ = v1610
	var v1620 int32
	_ = v1620
	var v1636 int32
	_ = v1636
	var v1645 int32
	_ = v1645
	var v1660 int32
	_ = v1660
	var v1666 int32
	_ = v1666
	v16 = m.G0
	v17 = int32(16)
	v18 = v16 - v17
	m.G0 = v18
	v26 = int32(-1636608416)
	if l1&int32(3) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v285 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v285
	v287 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v287
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v292 = v289
	v300 = v290
	goto L41
L2:
	;
	v258 = int32(14)
	v260 = v254 ^ v255 - base.I32_rotl(v254, v258)
	v264 = v260 ^ v253 - base.I32_rotl(v260, int32(11))
	v268 = v264 ^ v254 - base.I32_rotl(v264, int32(25))
	v272 = v268 ^ v260 - base.I32_rotl(v268, int32(16))
	v276 = v272 ^ v264 - base.I32_rotl(v272, int32(4))
	v280 = v276 ^ v268 - base.I32_rotl(v276, v258)
	goto L1
L3:
	;
	switch v180 - int32(1) {
	case 0:
		v246 = v171
		v247 = v172
		v248 = v176
		goto L30
	case 1:
		v239 = v171
		v240 = v172
		v241 = v176
		goto L31
	case 2:
		v232 = v171
		v233 = v172
		v234 = v176
		goto L32
	case 3:
		v226 = v172
		v227 = v176
		goto L33
	case 4:
		v222 = v172
		v223 = v176
		goto L34
	case 5:
		v216 = v172
		v217 = v176
		goto L35
	case 6:
		v210 = v172
		v211 = v176
		goto L36
	case 7:
		v205 = v176
		goto L37
	case 8:
		v200 = v176
		goto L38
	case 9:
		v195 = v176
		goto L39
	case 10:
		goto L40
	default:
		v253 = v171
		v254 = v172
		v255 = v176
		goto L2
	}
L4:
	;
	v135 = l1
	v136 = v17
	v137 = v26
	v138 = v26
	v139 = v26
	goto L27
L5:
	;
	goto L4
L6:
	;
	goto L7
L7:
	;
	goto L11
L9:
	;
	switch v78 - int32(1) {
	case 0:
		v132 = v69
		goto L16
	case 1:
		v127 = v69
		goto L17
	case 2:
		goto L18
	case 3:
		v120 = v70
		goto L19
	case 4:
		v117 = v70
		goto L20
	case 5:
		v112 = v70
		goto L21
	case 6:
		goto L22
	case 7:
		v103 = v74
		goto L23
	case 8:
		v98 = v74
		goto L24
	case 9:
		v93 = v74
		goto L25
	case 10:
		goto L26
	default:
		v253 = v69
		v254 = v70
		v255 = v74
		goto L2
	}
L11:
	;
	goto L12
L12:
	;
	v33 = l1
	v34 = v17
	v35 = v26
	v36 = v26
	v37 = v26
	goto L13
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v40 = v39 + v36
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v44 = v43 + v37
	v46 = int32(4)
	v48 = v41 + v35 - v44 ^ base.I32_rotl(v44, v46)
	v52 = v40 - v48 ^ base.I32_rotl(v48, int32(6))
	v53 = v44 + v40
	v54 = v48 + v53
	v55 = v52 + v54
	v59 = v53 - v52 ^ base.I32_rotl(v52, int32(8))
	v63 = v54 - v59 ^ base.I32_rotl(v59, int32(16))
	v67 = v55 - v63 ^ base.I32_rotl(v63, int32(19))
	v68 = v59 + v55
	v69 = v63 + v68
	v70 = v67 + v69
	v74 = v68 - v67 ^ base.I32_rotl(v67, v46)
	v75 = int32(12)
	v76 = v33 + v75
	v78 = v34 - v75
	if base.Ui32(int32(11)) < base.Ui32(v78) {
		v33 = v76
		v34 = v78
		v35 = v69
		v36 = v70
		v37 = v74
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L9
L15:
	;
	goto L14
L16:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v253 = v132 + v133
	v254 = v70
	v255 = v74
	goto L2
L17:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v132 = v128<<(uint(int32(8))%32) + v127
	goto L16
L18:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
	v127 = v123<<(uint(int32(16))%32) + v69
	goto L17
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v253 = v121 + v69
	v254 = v120
	v255 = v74
	goto L2
L20:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
	v120 = v117 + v118
	goto L19
L21:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+5)))
	v117 = v113<<(uint(int32(8))%32) + v112
	goto L20
L22:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+6)))
	v112 = v108<<(uint(int32(16))%32) + v70
	goto L21
L23:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v253 = v104 + v69
	v254 = v106 + v70
	v255 = v103
	goto L2
L24:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
	v103 = v99<<(uint(int32(8))%32) + v98
	goto L23
L25:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
	v98 = v94<<(uint(int32(16))%32) + v93
	goto L24
L26:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)))
	v93 = v89<<(uint(int32(24))%32) + v74
	goto L25
L27:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v142 = v141 + v138
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v146 = v145 + v139
	v148 = int32(4)
	v150 = v143 + v137 - v146 ^ base.I32_rotl(v146, v148)
	v154 = v142 - v150 ^ base.I32_rotl(v150, int32(6))
	v155 = v146 + v142
	v156 = v150 + v155
	v157 = v154 + v156
	v161 = v155 - v154 ^ base.I32_rotl(v154, int32(8))
	v165 = v156 - v161 ^ base.I32_rotl(v161, int32(16))
	v169 = v157 - v165 ^ base.I32_rotl(v165, int32(19))
	v170 = v161 + v157
	v171 = v165 + v170
	v172 = v169 + v171
	v176 = v170 - v169 ^ base.I32_rotl(v169, v148)
	v177 = int32(12)
	v178 = v135 + v177
	v180 = v136 - v177
	if base.Ui32(int32(11)) < base.Ui32(v180) {
		v135 = v178
		v136 = v180
		v137 = v171
		v138 = v172
		v139 = v176
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L3
L29:
	;
	goto L28
L30:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v253 = v246 + v249
	v254 = v247
	v255 = v248
	goto L2
L31:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	v246 = v242<<(uint(int32(8))%32) + v239
	v247 = v240
	v248 = v241
	goto L30
L32:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
	v239 = v235<<(uint(int32(16))%32) + v232
	v240 = v233
	v241 = v234
	goto L31
L33:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
	v232 = v228<<(uint(int32(24))%32) + v171
	v233 = v226
	v234 = v227
	goto L32
L34:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
	v226 = v222 + v224
	v227 = v223
	goto L33
L35:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
	v222 = v218<<(uint(int32(8))%32) + v216
	v223 = v217
	goto L34
L36:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
	v216 = v212<<(uint(int32(16))%32) + v210
	v217 = v211
	goto L35
L37:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
	v210 = v206<<(uint(int32(24))%32) + v172
	v211 = v205
	goto L36
L38:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
	v205 = v201<<(uint(int32(8))%32) + v200
	goto L37
L39:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
	v200 = v196<<(uint(int32(16))%32) + v195
	goto L38
L40:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+10)))
	v195 = v191<<(uint(int32(24))%32) + v176
	goto L39
L41:
	;
	if base.Ui32(v292) <= base.Ui32(v300) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v1666 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1666
	v292 = v1666
	v300 = v1660
	goto L41
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v1645)
	m.G0 = v18 + int32(16)
	return v1636
L45:
	;
	v1636 = v1620
	v1645 = int32(0)
	goto L44
L46:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1605 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1604 + v1605
	v1608 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1595)+8)) = v1608
	v1610 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v1595))) = v1610
	*(*uint8)(unsafe.Add(mBase, uint32(v1595)+20)) = uint8(v1605)
	v1620 = v1595
	goto L45
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L63
	} else {
		goto L273
	}
L48:
	;
	v307 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v307 == int64(4294967296) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v1080 = int32(0)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1083 = v1082 & (v280 ^ v272 - base.I32_rotl(v280, int32(24)))
	v1086 = v1081 + v1083*int32(40)
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1086)+20)))
	if v1087 == v1080 {
		v1595 = v1086
		goto L46
	} else {
		goto L184
	}
L51:
	;
	v310 = int32(0)
	v312 = int64(2)
	v314 = v307 << (uint(int64(1)) % 64)
	if base.Ui64(v314) <= base.Ui64(v312) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L50
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L63
	} else {
		goto L181
	}
L54:
	;
	v317 = v312
	goto L56
L55:
	;
	v317 = v314
	goto L56
L56:
	;
	v318 = int64(1)
	if v317&(v317-v318) == int64(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v328 = v317
	goto L59
L58:
	;
	v328 = v318 << (uint(int64(64)-base.I64_clz(v317)) % 64)
	goto L59
L59:
	;
	if base.Ui64(v328*int64(40)) < base.Ui64(int64(2147483647)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v334 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v340 = F_MemoryContextAllocExtended(m, v335, base.I32_wrap_i64(v328)*int32(40), int32(5))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L63
	} else {
		goto L178
	}
L63:
	;
	return int32(0)
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v340
	v345 = int64(1)
	if v328&(v328-v345) == int64(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v355 = v328
	goto L67
L66:
	;
	v355 = v345 << (uint(int64(64)-base.I64_clz(v328)) % 64)
	goto L67
L67:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v355*int64(40)) {
		goto L53
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v355
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = base.I32_wrap_i64(v355) - int32(1)
	v370 = base.F64_mul(base.F64_convert_i64_u(v355), float64(0.9))
	if base.F64_lt(v370, float64(4.294967296e+09))&base.F64_ge(v370, float64(0)) != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v355 == int64(4294967296) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v376 = base.I32_trunc_f64_u(v370)
	v378 = v376
	goto L69
L71:
	;
	goto L72
L72:
	;
	v378 = int32(0)
	goto L69
L73:
	;
	v379 = int32(-85899346)
	goto L75
L74:
	;
	v379 = v378
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v379
	if v334 != int64(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v387 = v310
	goto L80
L77:
	;
	goto L78
L78:
	;
	F_pfree(m, v333)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L63
	} else {
		goto L177
	}
L79:
	;
	v682 = v677
	v684 = v310
	goto L125
L80:
	;
	v400 = v333 + v387*int32(40)
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+20)))
	if v401 != int32(1) {
		v677 = v387
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v677 = int32(0)
	goto L79
L82:
	;
	v404 = int32(16)
	v410 = int32(-1636608416)
	if v400&int32(3) != 0 {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if (v664^v656-base.I32_rotl(v664, int32(24)))&v669 == v387 {
		v677 = v387
		goto L79
	} else {
		goto L123
	}
L84:
	;
	v642 = int32(14)
	v644 = v638 ^ v639 - base.I32_rotl(v638, v642)
	v648 = v644 ^ v637 - base.I32_rotl(v644, int32(11))
	v652 = v648 ^ v638 - base.I32_rotl(v648, int32(25))
	v656 = v652 ^ v644 - base.I32_rotl(v652, int32(16))
	v660 = v656 ^ v648 - base.I32_rotl(v656, int32(4))
	v664 = v660 ^ v652 - base.I32_rotl(v660, v642)
	goto L83
L85:
	;
	switch v564 - int32(1) {
	case 0:
		v630 = v555
		v631 = v556
		v632 = v560
		goto L112
	case 1:
		v623 = v555
		v624 = v556
		v625 = v560
		goto L113
	case 2:
		v616 = v555
		v617 = v556
		v618 = v560
		goto L114
	case 3:
		v610 = v556
		v611 = v560
		goto L115
	case 4:
		v606 = v556
		v607 = v560
		goto L116
	case 5:
		v600 = v556
		v601 = v560
		goto L117
	case 6:
		v594 = v556
		v595 = v560
		goto L118
	case 7:
		v589 = v560
		goto L119
	case 8:
		v584 = v560
		goto L120
	case 9:
		v579 = v560
		goto L121
	case 10:
		goto L122
	default:
		v637 = v555
		v638 = v556
		v639 = v560
		goto L84
	}
L86:
	;
	v519 = v400
	v520 = v404
	v521 = v410
	v522 = v410
	v523 = v410
	goto L109
L87:
	;
	goto L86
L88:
	;
	goto L89
L89:
	;
	goto L93
L91:
	;
	switch v462 - int32(1) {
	case 0:
		v516 = v453
		goto L98
	case 1:
		v511 = v453
		goto L99
	case 2:
		goto L100
	case 3:
		v504 = v454
		goto L101
	case 4:
		v501 = v454
		goto L102
	case 5:
		v496 = v454
		goto L103
	case 6:
		goto L104
	case 7:
		v487 = v458
		goto L105
	case 8:
		v482 = v458
		goto L106
	case 9:
		v477 = v458
		goto L107
	case 10:
		goto L108
	default:
		v637 = v453
		v638 = v454
		v639 = v458
		goto L84
	}
L93:
	;
	goto L94
L94:
	;
	v417 = v400
	v418 = v404
	v419 = v410
	v420 = v410
	v421 = v410
	goto L95
L95:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	v424 = v423 + v420
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v417)+8))
	v428 = v427 + v421
	v430 = int32(4)
	v432 = v425 + v419 - v428 ^ base.I32_rotl(v428, v430)
	v436 = v424 - v432 ^ base.I32_rotl(v432, int32(6))
	v437 = v428 + v424
	v438 = v432 + v437
	v439 = v436 + v438
	v443 = v437 - v436 ^ base.I32_rotl(v436, int32(8))
	v447 = v438 - v443 ^ base.I32_rotl(v443, int32(16))
	v451 = v439 - v447 ^ base.I32_rotl(v447, int32(19))
	v452 = v443 + v439
	v453 = v447 + v452
	v454 = v451 + v453
	v458 = v452 - v451 ^ base.I32_rotl(v451, v430)
	v459 = int32(12)
	v460 = v417 + v459
	v462 = v418 - v459
	if base.Ui32(int32(11)) < base.Ui32(v462) {
		v417 = v460
		v418 = v462
		v419 = v453
		v420 = v454
		v421 = v458
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L91
L97:
	;
	goto L96
L98:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v637 = v516 + v517
	v638 = v454
	v639 = v458
	goto L84
L99:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+1)))
	v516 = v512<<(uint(int32(8))%32) + v511
	goto L98
L100:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+2)))
	v511 = v507<<(uint(int32(16))%32) + v453
	goto L99
L101:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v637 = v505 + v453
	v638 = v504
	v639 = v458
	goto L84
L102:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+4)))
	v504 = v501 + v502
	goto L101
L103:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+5)))
	v501 = v497<<(uint(int32(8))%32) + v496
	goto L102
L104:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+6)))
	v496 = v492<<(uint(int32(16))%32) + v454
	goto L103
L105:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	v637 = v488 + v453
	v638 = v490 + v454
	v639 = v487
	goto L84
L106:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+8)))
	v487 = v483<<(uint(int32(8))%32) + v482
	goto L105
L107:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+9)))
	v482 = v478<<(uint(int32(16))%32) + v477
	goto L106
L108:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+10)))
	v477 = v473<<(uint(int32(24))%32) + v458
	goto L107
L109:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
	v526 = v525 + v522
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	v530 = v529 + v523
	v532 = int32(4)
	v534 = v527 + v521 - v530 ^ base.I32_rotl(v530, v532)
	v538 = v526 - v534 ^ base.I32_rotl(v534, int32(6))
	v539 = v530 + v526
	v540 = v534 + v539
	v541 = v538 + v540
	v545 = v539 - v538 ^ base.I32_rotl(v538, int32(8))
	v549 = v540 - v545 ^ base.I32_rotl(v545, int32(16))
	v553 = v541 - v549 ^ base.I32_rotl(v549, int32(19))
	v554 = v545 + v541
	v555 = v549 + v554
	v556 = v553 + v555
	v560 = v554 - v553 ^ base.I32_rotl(v553, v532)
	v561 = int32(12)
	v562 = v519 + v561
	v564 = v520 - v561
	if base.Ui32(int32(11)) < base.Ui32(v564) {
		v519 = v562
		v520 = v564
		v521 = v555
		v522 = v556
		v523 = v560
		goto L109
	} else {
		goto L111
	}
L110:
	;
	goto L85
L111:
	;
	goto L110
L112:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	v637 = v630 + v633
	v638 = v631
	v639 = v632
	goto L84
L113:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+1)))
	v630 = v626<<(uint(int32(8))%32) + v623
	v631 = v624
	v632 = v625
	goto L112
L114:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+2)))
	v623 = v619<<(uint(int32(16))%32) + v616
	v624 = v617
	v625 = v618
	goto L113
L115:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+3)))
	v616 = v612<<(uint(int32(24))%32) + v555
	v617 = v610
	v618 = v611
	goto L114
L116:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+4)))
	v610 = v606 + v608
	v611 = v607
	goto L115
L117:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+5)))
	v606 = v602<<(uint(int32(8))%32) + v600
	v607 = v601
	goto L116
L118:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+6)))
	v600 = v596<<(uint(int32(16))%32) + v594
	v601 = v595
	goto L117
L119:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+7)))
	v594 = v590<<(uint(int32(24))%32) + v556
	v595 = v589
	goto L118
L120:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+8)))
	v589 = v585<<(uint(int32(8))%32) + v584
	goto L119
L121:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+9)))
	v584 = v580<<(uint(int32(16))%32) + v579
	goto L120
L122:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562)+10)))
	v579 = v575<<(uint(int32(24))%32) + v560
	goto L121
L123:
	;
	v673 = v387 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v673)) < base.Ui64(v334) {
		v387 = v673
		goto L80
	} else {
		goto L124
	}
L124:
	;
	goto L81
L125:
	;
	v695 = v333 + v682*int32(40)
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695)+20)))
	if v696 == int32(1) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L78
L127:
	;
	v699 = int32(16)
	v705 = int32(-1636608416)
	if v695&int32(3) != 0 {
		goto L134
	} else {
		goto L135
	}
L128:
	;
	goto L129
L129:
	;
	v1013 = v682 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1013)) < base.Ui64(v334) {
		goto L173
	} else {
		goto L174
	}
L130:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v966 = v959 ^ v951 - base.I32_rotl(v959, int32(24))
	goto L170
L131:
	;
	v937 = int32(14)
	v939 = v933 ^ v934 - base.I32_rotl(v933, v937)
	v943 = v939 ^ v932 - base.I32_rotl(v939, int32(11))
	v947 = v943 ^ v933 - base.I32_rotl(v943, int32(25))
	v951 = v947 ^ v939 - base.I32_rotl(v947, int32(16))
	v955 = v951 ^ v943 - base.I32_rotl(v951, int32(4))
	v959 = v955 ^ v947 - base.I32_rotl(v955, v937)
	goto L130
L132:
	;
	switch v859 - int32(1) {
	case 0:
		v925 = v850
		v926 = v851
		v927 = v855
		goto L159
	case 1:
		v918 = v850
		v919 = v851
		v920 = v855
		goto L160
	case 2:
		v911 = v850
		v912 = v851
		v913 = v855
		goto L161
	case 3:
		v905 = v851
		v906 = v855
		goto L162
	case 4:
		v901 = v851
		v902 = v855
		goto L163
	case 5:
		v895 = v851
		v896 = v855
		goto L164
	case 6:
		v889 = v851
		v890 = v855
		goto L165
	case 7:
		v884 = v855
		goto L166
	case 8:
		v879 = v855
		goto L167
	case 9:
		v874 = v855
		goto L168
	case 10:
		goto L169
	default:
		v932 = v850
		v933 = v851
		v934 = v855
		goto L131
	}
L133:
	;
	v814 = v695
	v815 = v699
	v816 = v705
	v817 = v705
	v818 = v705
	goto L156
L134:
	;
	goto L133
L135:
	;
	goto L136
L136:
	;
	goto L140
L138:
	;
	switch v757 - int32(1) {
	case 0:
		v811 = v748
		goto L145
	case 1:
		v806 = v748
		goto L146
	case 2:
		goto L147
	case 3:
		v799 = v749
		goto L148
	case 4:
		v796 = v749
		goto L149
	case 5:
		v791 = v749
		goto L150
	case 6:
		goto L151
	case 7:
		v782 = v753
		goto L152
	case 8:
		v777 = v753
		goto L153
	case 9:
		v772 = v753
		goto L154
	case 10:
		goto L155
	default:
		v932 = v748
		v933 = v749
		v934 = v753
		goto L131
	}
L140:
	;
	goto L141
L141:
	;
	v712 = v695
	v713 = v699
	v714 = v705
	v715 = v705
	v716 = v705
	goto L142
L142:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	v719 = v718 + v715
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v712)+8))
	v723 = v722 + v716
	v725 = int32(4)
	v727 = v720 + v714 - v723 ^ base.I32_rotl(v723, v725)
	v731 = v719 - v727 ^ base.I32_rotl(v727, int32(6))
	v732 = v723 + v719
	v733 = v727 + v732
	v734 = v731 + v733
	v738 = v732 - v731 ^ base.I32_rotl(v731, int32(8))
	v742 = v733 - v738 ^ base.I32_rotl(v738, int32(16))
	v746 = v734 - v742 ^ base.I32_rotl(v742, int32(19))
	v747 = v738 + v734
	v748 = v742 + v747
	v749 = v746 + v748
	v753 = v747 - v746 ^ base.I32_rotl(v746, v725)
	v754 = int32(12)
	v755 = v712 + v754
	v757 = v713 - v754
	if base.Ui32(int32(11)) < base.Ui32(v757) {
		v712 = v755
		v713 = v757
		v714 = v748
		v715 = v749
		v716 = v753
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L138
L144:
	;
	goto L143
L145:
	;
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	v932 = v811 + v812
	v933 = v749
	v934 = v753
	goto L131
L146:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+1)))
	v811 = v807<<(uint(int32(8))%32) + v806
	goto L145
L147:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+2)))
	v806 = v802<<(uint(int32(16))%32) + v748
	goto L146
L148:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	v932 = v800 + v748
	v933 = v799
	v934 = v753
	goto L131
L149:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+4)))
	v799 = v796 + v797
	goto L148
L150:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+5)))
	v796 = v792<<(uint(int32(8))%32) + v791
	goto L149
L151:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+6)))
	v791 = v787<<(uint(int32(16))%32) + v749
	goto L150
L152:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v755)+4))
	v932 = v783 + v748
	v933 = v785 + v749
	v934 = v782
	goto L131
L153:
	;
	v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+8)))
	v782 = v778<<(uint(int32(8))%32) + v777
	goto L152
L154:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+9)))
	v777 = v773<<(uint(int32(16))%32) + v772
	goto L153
L155:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755)+10)))
	v772 = v768<<(uint(int32(24))%32) + v753
	goto L154
L156:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v814)+4))
	v821 = v820 + v817
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v814)+8))
	v825 = v824 + v818
	v827 = int32(4)
	v829 = v822 + v816 - v825 ^ base.I32_rotl(v825, v827)
	v833 = v821 - v829 ^ base.I32_rotl(v829, int32(6))
	v834 = v825 + v821
	v835 = v829 + v834
	v836 = v833 + v835
	v840 = v834 - v833 ^ base.I32_rotl(v833, int32(8))
	v844 = v835 - v840 ^ base.I32_rotl(v840, int32(16))
	v848 = v836 - v844 ^ base.I32_rotl(v844, int32(19))
	v849 = v840 + v836
	v850 = v844 + v849
	v851 = v848 + v850
	v855 = v849 - v848 ^ base.I32_rotl(v848, v827)
	v856 = int32(12)
	v857 = v814 + v856
	v859 = v815 - v856
	if base.Ui32(int32(11)) < base.Ui32(v859) {
		v814 = v857
		v815 = v859
		v816 = v850
		v817 = v851
		v818 = v855
		goto L156
	} else {
		goto L158
	}
L157:
	;
	goto L132
L158:
	;
	goto L157
L159:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857))))
	v932 = v925 + v928
	v933 = v926
	v934 = v927
	goto L131
L160:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+1)))
	v925 = v921<<(uint(int32(8))%32) + v918
	v926 = v919
	v927 = v920
	goto L159
L161:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+2)))
	v918 = v914<<(uint(int32(16))%32) + v911
	v919 = v912
	v920 = v913
	goto L160
L162:
	;
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+3)))
	v911 = v907<<(uint(int32(24))%32) + v850
	v912 = v905
	v913 = v906
	goto L161
L163:
	;
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+4)))
	v905 = v901 + v903
	v906 = v902
	goto L162
L164:
	;
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+5)))
	v901 = v897<<(uint(int32(8))%32) + v895
	v902 = v896
	goto L163
L165:
	;
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+6)))
	v895 = v891<<(uint(int32(16))%32) + v889
	v896 = v890
	goto L164
L166:
	;
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+7)))
	v889 = v885<<(uint(int32(24))%32) + v851
	v890 = v884
	goto L165
L167:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+8)))
	v884 = v880<<(uint(int32(8))%32) + v879
	goto L166
L168:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+9)))
	v879 = v875<<(uint(int32(16))%32) + v874
	goto L167
L169:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857)+10)))
	v874 = v870<<(uint(int32(24))%32) + v855
	goto L168
L170:
	;
	v980 = v966 & v964
	v985 = v340 + v980*int32(40)
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v985)+20)))
	if v986 != 0 {
		v966 = v980 + int32(1)
		goto L170
	} else {
		goto L172
	}
L171:
	;
	v987 = *(*int64)(unsafe.Add(mBase, uint32(v695)))
	*(*int64)(unsafe.Add(mBase, uint32(v985))) = v987
	v989 = *(*int64)(unsafe.Add(mBase, uint32(v695)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v985)+32)) = v989
	v991 = *(*int64)(unsafe.Add(mBase, uint32(v695)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v985)+24)) = v991
	v993 = *(*int64)(unsafe.Add(mBase, uint32(v695)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v985)+16)) = v993
	v995 = *(*int64)(unsafe.Add(mBase, uint32(v695)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v985)+8)) = v995
	goto L129
L172:
	;
	goto L171
L173:
	;
	v1017 = v1013
	goto L175
L174:
	;
	v1017 = int32(0)
	goto L175
L175:
	;
	v1019 = v684 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1019)) < base.Ui64(v334) {
		v682 = v1017
		v684 = v1019
		goto L125
	} else {
		goto L176
	}
L176:
	;
	goto L126
L177:
	;
	goto L52
L178:
	;
	F_errmsg_internal(m, int32(416710), int32(0))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L63
	} else {
		goto L179
	}
L179:
	;
	F_errfinish(m, int32(340634), int32(327), int32(355704))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L63
	} else {
		goto L180
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	F_errmsg_internal(m, int32(416710), int32(0))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L63
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(340634), int32(327), int32(355704))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L63
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	v1093 = v1080
	v1094 = v1083
	v1096 = v1086
	goto L185
L185:
	;
	v1106 = int32(16)
	goto L190
L186:
	;
	v1595 = v1574
	goto L46
L187:
	;
	if v1168 == int32(0) {
		v1636 = v1096
		v1645 = int32(1)
		goto L44
	} else {
		goto L205
	}
L188:
	;
	v1168 = int32(0)
	goto L187
L189:
	;
	v1142 = v1137
	v1143 = v1138
	v1144 = v1139
	goto L199
L190:
	;
	if (v1096|v18)&int32(3) != 0 {
		v1137 = v1096
		v1138 = v18
		v1139 = v1106
		goto L189
	} else {
		goto L193
	}
L192:
	;
	if v1127 == int32(0) {
		goto L188
	} else {
		goto L198
	}
L193:
	;
	v1114 = v1096
	v1115 = v18
	v1116 = v1106
	goto L194
L194:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1114)))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1115)))
	if v1119 != v1120 {
		v1137 = v1114
		v1138 = v1115
		v1139 = v1116
		goto L189
	} else {
		goto L196
	}
L195:
	;
	goto L192
L196:
	;
	v1122 = int32(4)
	v1123 = v1115 + v1122
	v1125 = v1114 + v1122
	v1127 = v1116 - v1122
	if base.Ui32(int32(3)) < base.Ui32(v1127) {
		v1114 = v1125
		v1115 = v1123
		v1116 = v1127
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	v1137 = v1125
	v1138 = v1123
	v1139 = v1127
	goto L189
L199:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1142))))
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1143))))
	if v1147 == v1148 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v1168 = v1147 - v1148
	goto L187
L201:
	;
	v1150 = int32(1)
	v1155 = v1144 - v1150
	if v1155 != 0 {
		v1142 = v1142 + v1150
		v1143 = v1143 + v1150
		v1144 = v1155
		goto L199
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	goto L200
L204:
	;
	goto L188
L205:
	;
	v1171 = int32(16)
	v1177 = int32(-1636608416)
	if v1096&int32(3) != 0 {
		goto L210
	} else {
		goto L211
	}
L206:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1437 = (v1431 ^ v1423 - base.I32_rotl(v1431, int32(24))) & v1436
	if base.Ui32(v1094) < base.Ui32(v1437) {
		goto L246
	} else {
		goto L247
	}
L207:
	;
	v1409 = int32(14)
	v1411 = v1405 ^ v1406 - base.I32_rotl(v1405, v1409)
	v1415 = v1411 ^ v1404 - base.I32_rotl(v1411, int32(11))
	v1419 = v1415 ^ v1405 - base.I32_rotl(v1415, int32(25))
	v1423 = v1419 ^ v1411 - base.I32_rotl(v1419, int32(16))
	v1427 = v1423 ^ v1415 - base.I32_rotl(v1423, int32(4))
	v1431 = v1427 ^ v1419 - base.I32_rotl(v1427, v1409)
	goto L206
L208:
	;
	switch v1331 - int32(1) {
	case 0:
		v1397 = v1322
		v1398 = v1323
		v1399 = v1327
		goto L235
	case 1:
		v1390 = v1322
		v1391 = v1323
		v1392 = v1327
		goto L236
	case 2:
		v1383 = v1322
		v1384 = v1323
		v1385 = v1327
		goto L237
	case 3:
		v1377 = v1323
		v1378 = v1327
		goto L238
	case 4:
		v1373 = v1323
		v1374 = v1327
		goto L239
	case 5:
		v1367 = v1323
		v1368 = v1327
		goto L240
	case 6:
		v1361 = v1323
		v1362 = v1327
		goto L241
	case 7:
		v1356 = v1327
		goto L242
	case 8:
		v1351 = v1327
		goto L243
	case 9:
		v1346 = v1327
		goto L244
	case 10:
		goto L245
	default:
		v1404 = v1322
		v1405 = v1323
		v1406 = v1327
		goto L207
	}
L209:
	;
	v1286 = v1096
	v1287 = v1171
	v1288 = v1177
	v1289 = v1177
	v1290 = v1177
	goto L232
L210:
	;
	goto L209
L211:
	;
	goto L212
L212:
	;
	goto L216
L214:
	;
	switch v1229 - int32(1) {
	case 0:
		v1283 = v1220
		goto L221
	case 1:
		v1278 = v1220
		goto L222
	case 2:
		goto L223
	case 3:
		v1271 = v1221
		goto L224
	case 4:
		v1268 = v1221
		goto L225
	case 5:
		v1263 = v1221
		goto L226
	case 6:
		goto L227
	case 7:
		v1254 = v1225
		goto L228
	case 8:
		v1249 = v1225
		goto L229
	case 9:
		v1244 = v1225
		goto L230
	case 10:
		goto L231
	default:
		v1404 = v1220
		v1405 = v1221
		v1406 = v1225
		goto L207
	}
L216:
	;
	goto L217
L217:
	;
	v1184 = v1096
	v1185 = v1171
	v1186 = v1177
	v1187 = v1177
	v1188 = v1177
	goto L218
L218:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1184)+4))
	v1191 = v1190 + v1187
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1184)))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1184)+8))
	v1195 = v1194 + v1188
	v1197 = int32(4)
	v1199 = v1192 + v1186 - v1195 ^ base.I32_rotl(v1195, v1197)
	v1203 = v1191 - v1199 ^ base.I32_rotl(v1199, int32(6))
	v1204 = v1195 + v1191
	v1205 = v1199 + v1204
	v1206 = v1203 + v1205
	v1210 = v1204 - v1203 ^ base.I32_rotl(v1203, int32(8))
	v1214 = v1205 - v1210 ^ base.I32_rotl(v1210, int32(16))
	v1218 = v1206 - v1214 ^ base.I32_rotl(v1214, int32(19))
	v1219 = v1210 + v1206
	v1220 = v1214 + v1219
	v1221 = v1218 + v1220
	v1225 = v1219 - v1218 ^ base.I32_rotl(v1218, v1197)
	v1226 = int32(12)
	v1227 = v1184 + v1226
	v1229 = v1185 - v1226
	if base.Ui32(int32(11)) < base.Ui32(v1229) {
		v1184 = v1227
		v1185 = v1229
		v1186 = v1220
		v1187 = v1221
		v1188 = v1225
		goto L218
	} else {
		goto L220
	}
L219:
	;
	goto L214
L220:
	;
	goto L219
L221:
	;
	v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227))))
	v1404 = v1283 + v1284
	v1405 = v1221
	v1406 = v1225
	goto L207
L222:
	;
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+1)))
	v1283 = v1279<<(uint(int32(8))%32) + v1278
	goto L221
L223:
	;
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+2)))
	v1278 = v1274<<(uint(int32(16))%32) + v1220
	goto L222
L224:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1227)))
	v1404 = v1272 + v1220
	v1405 = v1271
	v1406 = v1225
	goto L207
L225:
	;
	v1269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+4)))
	v1271 = v1268 + v1269
	goto L224
L226:
	;
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+5)))
	v1268 = v1264<<(uint(int32(8))%32) + v1263
	goto L225
L227:
	;
	v1259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+6)))
	v1263 = v1259<<(uint(int32(16))%32) + v1221
	goto L226
L228:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1227)))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+4))
	v1404 = v1255 + v1220
	v1405 = v1257 + v1221
	v1406 = v1254
	goto L207
L229:
	;
	v1250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+8)))
	v1254 = v1250<<(uint(int32(8))%32) + v1249
	goto L228
L230:
	;
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+9)))
	v1249 = v1245<<(uint(int32(16))%32) + v1244
	goto L229
L231:
	;
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227)+10)))
	v1244 = v1240<<(uint(int32(24))%32) + v1225
	goto L230
L232:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+4))
	v1293 = v1292 + v1289
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1286)))
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+8))
	v1297 = v1296 + v1290
	v1299 = int32(4)
	v1301 = v1294 + v1288 - v1297 ^ base.I32_rotl(v1297, v1299)
	v1305 = v1293 - v1301 ^ base.I32_rotl(v1301, int32(6))
	v1306 = v1297 + v1293
	v1307 = v1301 + v1306
	v1308 = v1305 + v1307
	v1312 = v1306 - v1305 ^ base.I32_rotl(v1305, int32(8))
	v1316 = v1307 - v1312 ^ base.I32_rotl(v1312, int32(16))
	v1320 = v1308 - v1316 ^ base.I32_rotl(v1316, int32(19))
	v1321 = v1312 + v1308
	v1322 = v1316 + v1321
	v1323 = v1320 + v1322
	v1327 = v1321 - v1320 ^ base.I32_rotl(v1320, v1299)
	v1328 = int32(12)
	v1329 = v1286 + v1328
	v1331 = v1287 - v1328
	if base.Ui32(int32(11)) < base.Ui32(v1331) {
		v1286 = v1329
		v1287 = v1331
		v1288 = v1322
		v1289 = v1323
		v1290 = v1327
		goto L232
	} else {
		goto L234
	}
L233:
	;
	goto L208
L234:
	;
	goto L233
L235:
	;
	v1400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329))))
	v1404 = v1397 + v1400
	v1405 = v1398
	v1406 = v1399
	goto L207
L236:
	;
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+1)))
	v1397 = v1393<<(uint(int32(8))%32) + v1390
	v1398 = v1391
	v1399 = v1392
	goto L235
L237:
	;
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+2)))
	v1390 = v1386<<(uint(int32(16))%32) + v1383
	v1391 = v1384
	v1392 = v1385
	goto L236
L238:
	;
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+3)))
	v1383 = v1379<<(uint(int32(24))%32) + v1322
	v1384 = v1377
	v1385 = v1378
	goto L237
L239:
	;
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+4)))
	v1377 = v1373 + v1375
	v1378 = v1374
	goto L238
L240:
	;
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+5)))
	v1373 = v1369<<(uint(int32(8))%32) + v1367
	v1374 = v1368
	goto L239
L241:
	;
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+6)))
	v1367 = v1363<<(uint(int32(16))%32) + v1361
	v1368 = v1362
	goto L240
L242:
	;
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+7)))
	v1361 = v1357<<(uint(int32(24))%32) + v1323
	v1362 = v1356
	goto L241
L243:
	;
	v1352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+8)))
	v1356 = v1352<<(uint(int32(8))%32) + v1351
	goto L242
L244:
	;
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+9)))
	v1351 = v1347<<(uint(int32(16))%32) + v1346
	goto L243
L245:
	;
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+10)))
	v1346 = v1342<<(uint(int32(24))%32) + v1327
	goto L244
L246:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1441 = v1094 + v1439
	goto L248
L247:
	;
	v1441 = v1094
	goto L248
L248:
	;
	v1444 = (v1094 + int32(1)) & v1436
	if base.Ui32(v1441-v1437) < base.Ui32(v1093) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1450 = v1081 + v1444*int32(40)
	v1451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1450)+20)))
	if v1451 != 0 {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	goto L251
L251:
	;
	v1561 = v1093 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1561) {
		goto L268
	} else {
		goto L269
	}
L252:
	;
	v1453 = v1444
	v1459 = int32(0)
	goto L255
L253:
	;
	v1487 = v1444
	v1489 = v1450
	goto L254
L254:
	;
	if v1487 != v1094 {
		goto L262
	} else {
		goto L263
	}
L255:
	;
	v1468 = v1459 + int32(1)
	if int32(151) <= v1468 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1487 = v1481
	v1489 = v1484
	goto L254
L257:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1473 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1471), base.F64_convert_i64_u(v1473)), float64(0.1)) != 0 {
		v1660 = v1471
		goto L43
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v1481 = (v1453 + int32(1)) & v1436
	v1484 = v1081 + v1481*int32(40)
	v1485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1484)+20)))
	if v1485 != 0 {
		v1453 = v1481
		v1459 = v1468
		goto L255
	} else {
		goto L261
	}
L260:
	;
	goto L259
L261:
	;
	goto L256
L262:
	;
	v1503 = v1487
	v1505 = v1489
	goto L265
L263:
	;
	goto L264
L264:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1551 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1550 + v1551
	v1554 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1096)+8)) = v1554
	v1556 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v1096))) = v1556
	*(*uint8)(unsafe.Add(mBase, uint32(v1096)+20)) = uint8(v1551)
	v1620 = v1096
	goto L45
L265:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1520 = v1517 & (v1503 - int32(1))
	v1523 = v1081 + v1520*int32(40)
	v1524 = *(*int64)(unsafe.Add(mBase, uint32(v1523)))
	*(*int64)(unsafe.Add(mBase, uint32(v1505))) = v1524
	v1526 = *(*int64)(unsafe.Add(mBase, uint32(v1523)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1505)+32)) = v1526
	v1528 = *(*int64)(unsafe.Add(mBase, uint32(v1523)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1505)+24)) = v1528
	v1530 = *(*int64)(unsafe.Add(mBase, uint32(v1523)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1505)+16)) = v1530
	v1532 = *(*int64)(unsafe.Add(mBase, uint32(v1523)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1505)+8)) = v1532
	if v1520 != v1094 {
		v1503 = v1520
		v1505 = v1523
		goto L265
	} else {
		goto L267
	}
L266:
	;
	goto L264
L267:
	;
	goto L266
L268:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1566 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1564), base.F64_convert_i64_u(v1566)), float64(0.1)) != 0 {
		v1660 = v1564
		goto L43
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1574 = v1081 + v1444*int32(40)
	v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1574)+20)))
	if v1575 != 0 {
		v1093 = v1561
		v1094 = v1444
		v1096 = v1574
		goto L185
	} else {
		goto L272
	}
L271:
	;
	goto L270
L272:
	;
	goto L186
L273:
	;
	F_errmsg_internal(m, int32(480207), int32(0))
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L63
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(340634), int32(630), int32(324489))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L63
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bloptions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1122]))
	v11 = F_build_reloptions(m, l0, l1, v6, int32(136), int32(4713072), int32(33))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v19 = base.I32_div_s(v15+int32(15), int32(16))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v19
		} else {
		}
		return v11
	}
}
func F_blvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v456 int32
	_ = v456
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	v17 = m.G0
	v19 = v17 - int32(128)
	m.G0 = v19
	v22 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	if int32(0) < v203 {
		goto L47
	} else {
		goto L48
	}
L2:
	;
	return int32(0)
L3:
	;
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
	v28 = v26 + v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+92))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+80))
	v32 = F_get_opfamily_name(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L44
	}
L7:
	;
	v36 = int32(0)
	v38 = F_SearchSysCacheList(m, int32(4), int32(1), v31, v36, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v40 = int32(1)
	v43 = int32(0)
	v45 = F_SearchSysCacheList(m, int32(5), v40, v31, v43, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+40))
	if v47 <= int32(0) {
		v190 = v40
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v50 = v30
	goto L13
L12:
	;
	v50 = v29
	goto L13
L13:
	;
	v56 = int32(0)
	v57 = v40
	goto L14
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v45+int32(48)+v56<<(uint(int32(2))%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+56))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
	v76 = v74 + v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	if v77 == v78 {
		v105 = v57
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v190 = v167
	goto L1
L16:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	if v106 != v29 {
		v167 = v105
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v80 = int32(0)
	v83 = F_errstart(m, int32(17), v80)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	if v83 == int32(0) {
		v105 = v80
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v91 = F_format_procedure(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+116)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v32
	F_errmsg(m, int32(270010), v19+int32(112))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(519039), int32(84), int32(371702))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v105 = v80
	goto L16
L24:
	;
	v171 = v56 + int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v45)+40))
	if v171 < v172 {
		v56 = v171
		v57 = v167
		goto L14
	} else {
		goto L43
	}
L25:
	;
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)))
	switch v108 - int32(1) {
	case 0:
		goto L28
	case 1:
		goto L30
	default:
		goto L29
	}
L26:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L39
	}
L27:
	;
	v136 = int32(0)
	v139 = F_errstart(m, int32(17), v136)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L37
	}
L28:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v50
	v129 = int32(1)
	v133 = F_check_amproc_signature(m, v125, int32(23), int32(0), v129, v129, v19+int32(96))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L35
	}
L29:
	;
	v116 = int32(0)
	v119 = F_errstart(m, int32(17), v116)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L33
	}
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v112 = F_check_amoptsproc_signature(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v112 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v167 = v105
	goto L24
L33:
	;
	if v119 == int32(0) {
		v167 = v116
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v145 = int32(489651)
	v147 = int32(111)
	goto L26
L35:
	;
	if v133 != 0 {
		v167 = v105
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L27
L37:
	;
	if v139 == int32(0) {
		v167 = v136
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v145 = int32(489464)
	v147 = int32(123)
	goto L26
L39:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v152 = F_format_procedure(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v32
	F_errmsg(m, v145, v19+int32(80))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(519039), v147, int32(371702))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v167 = int32(0)
	goto L24
L43:
	;
	goto L15
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	F_errmsg_internal(m, int32(45619), v19)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(519039), int32(50), int32(371702))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v211 = int32(0)
	v212 = v190
	goto L50
L48:
	;
	v335 = v190
	goto L49
L49:
	;
	v348 = F_identify_opfamily_groups(m, v38, v45)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L2
	} else {
		goto L83
	}
L50:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(48)+v211<<(uint(int32(2))%32))))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+56))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+22)))
	v231 = v229 + v230
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+16)))
	if v232 == int32(1) {
		v262 = v212
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v335 = v327
	goto L49
L52:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+18)))
	if v264 == int32(115) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v235 = int32(0)
	v238 = F_errstart(m, int32(17), v235)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	if v238 == int32(0) {
		v262 = v235
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	v246 = F_format_operator(m, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v231)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v32
	F_errmsg(m, int32(489269), v19-int32(-64))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(519039), int32(143), int32(371702))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v262 = v235
	goto L52
L60:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v300 = F_check_amop_signature(m, v296, int32(16), v298, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L2
	} else {
		goto L72
	}
L61:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	if v267 == int32(0) {
		v295 = v262
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v270 = int32(0)
	v273 = F_errstart(m, int32(17), v270)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if v273 == int32(0) {
		v295 = v270
		goto L60
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	v281 = F_format_operator(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v32
	F_errmsg(m, int32(189792), v19+int32(48))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(519039), int32(155), int32(371702))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v295 = v270
	goto L60
L71:
	;
	v329 = v211 + int32(1)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	if v329 < v330 {
		v211 = v329
		v212 = v327
		goto L50
	} else {
		goto L80
	}
L72:
	;
	if v300 != 0 {
		v327 = v295
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v302 = int32(0)
	v305 = F_errstart(m, int32(17), v302)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	if v305 == int32(0) {
		v327 = v302
		goto L71
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	v313 = F_format_operator(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v32
	F_errmsg(m, int32(378036), v19+int32(32))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(519039), int32(168), int32(371702))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v327 = v302
	goto L71
L80:
	;
	goto L51
L81:
	;
	F_ReleaseCatCacheList(m, v45)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L2
	} else {
		goto L121
	}
L82:
	;
	v475 = int32(0)
	v478 = F_errstart(m, int32(17), v475)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L2
	} else {
		goto L116
	}
L83:
	;
	if v348 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v352 = int32(0)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v353 <= v352 {
		v440 = v352
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if v440 == int32(0) {
		goto L82
	} else {
		goto L114
	}
L86:
	;
	v356 = int32(0)
	if v356 < v353 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v359 = v353
	goto L89
L88:
	;
	v359 = v356
	goto L89
L89:
	;
	v360 = int32(1)
	if v353 == v360 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v359&v360 == int32(0) {
		v440 = v412
		goto L85
	} else {
		goto L109
	}
L91:
	;
	v364 = int32(0)
	v410 = v364
	v412 = v364
	goto L90
L92:
	;
	goto L93
L93:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v348)+12))
	v369 = int32(0)
	v372 = v369
	v374 = v369
	v377 = v369
	goto L94
L94:
	;
	v390 = v368 + v372<<(uint(int32(2))%32)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	if v29 == v392 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v410 = v406
	v412 = v404
	goto L90
L96:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v394 == v29 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v397 = v374
	goto L98
L98:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	if v29 == v399 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v396 = v391
	goto L101
L100:
	;
	v396 = v374
	goto L101
L101:
	;
	v397 = v396
	goto L98
L102:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	if v401 == v29 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v404 = v397
	goto L104
L104:
	;
	v405 = int32(2)
	v406 = v372 + v405
	v408 = v377 + v405
	if v408 != v359&int32(2147483646) {
		v372 = v406
		v374 = v404
		v377 = v408
		goto L94
	} else {
		goto L108
	}
L105:
	;
	v403 = v398
	goto L107
L106:
	;
	v403 = v397
	goto L107
L107:
	;
	v404 = v403
	goto L104
L108:
	;
	goto L95
L109:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v348)+12))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v428+v410<<(uint(int32(2))%32))))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	if v433 != v29 {
		v440 = v412
		goto L85
	} else {
		goto L110
	}
L110:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v435 == v29 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v437 = v432
	goto L113
L112:
	;
	v437 = v412
	goto L113
L113:
	;
	v440 = v437
	goto L85
L114:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440)+16)))
	if v456&int32(2) != 0 {
		v503 = v335
		goto L81
	} else {
		goto L115
	}
L115:
	;
	goto L82
L116:
	;
	if v478 == int32(0) {
		v503 = v475
		goto L81
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v28 + int32(8)
	F_errmsg(m, int32(490800), v19+int32(16))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(519039), int32(206), int32(371702))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	v503 = v475
	goto L81
L121:
	;
	F_ReleaseCatCacheList(m, v38)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	F_ReleaseCatCache(m, v22)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	m.G0 = v19 + int32(128)
	return v503 & int32(1)
}
func F_boollt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v2 == v3) & base.B2i32(v5 != v3)
}
func F_boolrecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pq_getmsgbyte(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v3 != int32(0))
	}
}
func F_bounds_adjacent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v7 == int32(1) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
		if v6&int32(1) != 0 {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
			if v10 == v13 {
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
				return base.B2i32(v80 != v81)
			} else {
				if v10&int32(1) != 0 {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
					if v43 == int32(0) {
						return int32(0)
					} else {
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
						v49 = int32(1)
						v50 = v48 ^ v49
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v50)
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
						v54 = v52 ^ v49
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v54)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v49)
						v58 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v58)
						v62 = F_make_range(m, l0, l1, l2, v58, v58)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
							v68 = int32(1)
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(base.Ui32(v64)>>(uint(int32(2))%32))-v68))))
							return v70 & v68
						}
					}
				} else {
					return int32(0)
				}
			}
		} else {
			if v10&int32(1) != 0 {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
				if v43 == int32(0) {
					return int32(0)
				} else {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
					v49 = int32(1)
					v50 = v48 ^ v49
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v50)
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
					v54 = v52 ^ v49
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v54)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v49)
					v58 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v58)
					v62 = F_make_range(m, l0, l1, l2, v58, v58)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
						v68 = int32(1)
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(base.Ui32(v64)>>(uint(int32(2))%32))-v68))))
						return v70 & v68
					}
				}
			} else {
				return int32(0)
			}
		}
	} else {
		if v6&int32(1) != 0 {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
			if v25 == int32(0) {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
				if v43 == int32(0) {
					return int32(0)
				} else {
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
					v49 = int32(1)
					v50 = v48 ^ v49
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v50)
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
					v54 = v52 ^ v49
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v54)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v49)
					v58 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v58)
					v62 = F_make_range(m, l0, l1, l2, v58, v58)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
						v68 = int32(1)
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(base.Ui32(v64)>>(uint(int32(2))%32))-v68))))
						return v70 & v68
					}
				}
			} else {
				return int32(0)
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v35 = F_FunctionCall2Coll(m, l0+int32(212), v32, v33, v34)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if int32(0) <= v35 {
					if v35 == int32(0) {
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
						return base.B2i32(v80 != v81)
					} else {
						return int32(0)
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
					if v43 == int32(0) {
						return int32(0)
					} else {
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
						v49 = int32(1)
						v50 = v48 ^ v49
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v50)
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
						v54 = v52 ^ v49
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v54)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v49)
						v58 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v58)
						v62 = F_make_range(m, l0, l1, l2, v58, v58)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
							v68 = int32(1)
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+int32(base.Ui32(v64)>>(uint(int32(2))%32))-v68))))
							return v70 & v68
						}
					}
				}
			}
		}
	}
}
func F_bpcharcmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
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
	v17 = v12 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v22 = int32(1)
	v23 = v21 & v22
	if v21 == v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v23 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v23 != 0 {
		v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v37 = v26
	goto L10
L9:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v26
	goto L13
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v51 = v40
	goto L4
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v54 = v17
	goto L17
L16:
	;
	v54 = v12 + int32(4)
	goto L17
L17:
	;
	v59 = v51
	goto L18
L18:
	;
	if v59 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v78 = int32(1)
	v79 = v19 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v82 = v80 & v78
	if v80 == v78 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	v77 = v51 >> (uint(int32(31)) % 32) & v51
	goto L20
L22:
	;
	goto L23
L23:
	;
	v71 = v59 - int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v71))))
	if v73 == int32(32) {
		v59 = v71
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v77 = v59
	goto L20
L25:
	;
	if v82 != 0 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v85 = int32(4)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v87&int32(254) == int32(2) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v100 = int32(1)
	if v82 != 0 {
		v110 = int32(base.Ui32(v80)>>(uint(v100)%32)) - v100
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v96 = v85
	goto L31
L30:
	;
	v96 = base.B2i32(v87 == int32(18)) << (uint(v85) % 32)
	goto L31
L31:
	;
	if v87 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v85
	goto L34
L33:
	;
	v99 = v96
	goto L34
L34:
	;
	v110 = v99
	goto L25
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	v113 = v79
	goto L38
L37:
	;
	v113 = v19 + int32(4)
	goto L38
L38:
	;
	v118 = v110
	goto L39
L39:
	;
	if v118 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v137 = int32(1)
	if v21&v137 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v136 = v110 >> (uint(int32(31)) % 32) & v110
	goto L41
L43:
	;
	goto L44
L44:
	;
	v130 = v118 - int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v130))))
	if v132 == int32(32) {
		v118 = v130
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v136 = v118
	goto L41
L46:
	;
	v141 = v137
	goto L48
L47:
	;
	v141 = int32(4)
	goto L48
L48:
	;
	v143 = int32(1)
	if v80&v143 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v147 = v143
	goto L51
L50:
	;
	v147 = int32(4)
	goto L51
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = F_varstr_cmp(m, v12+v141, v77, v19+v147, v136, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v19 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v19)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return v150
L60:
	;
	goto L59
}
func F_bpcharlt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
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
	v17 = v12 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v22 = int32(1)
	v23 = v21 & v22
	if v21 == v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v23 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v23 != 0 {
		v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v37 = v26
	goto L10
L9:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v26
	goto L13
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v51 = v40
	goto L4
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v54 = v17
	goto L17
L16:
	;
	v54 = v12 + int32(4)
	goto L17
L17:
	;
	v59 = v51
	goto L18
L18:
	;
	if v59 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v78 = int32(1)
	v79 = v19 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v82 = v80 & v78
	if v80 == v78 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	v77 = v51 >> (uint(int32(31)) % 32) & v51
	goto L20
L22:
	;
	goto L23
L23:
	;
	v71 = v59 - int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v71))))
	if v73 == int32(32) {
		v59 = v71
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v77 = v59
	goto L20
L25:
	;
	if v82 != 0 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v85 = int32(4)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v87&int32(254) == int32(2) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v100 = int32(1)
	if v82 != 0 {
		v110 = int32(base.Ui32(v80)>>(uint(v100)%32)) - v100
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v96 = v85
	goto L31
L30:
	;
	v96 = base.B2i32(v87 == int32(18)) << (uint(v85) % 32)
	goto L31
L31:
	;
	if v87 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v85
	goto L34
L33:
	;
	v99 = v96
	goto L34
L34:
	;
	v110 = v99
	goto L25
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	v113 = v79
	goto L38
L37:
	;
	v113 = v19 + int32(4)
	goto L38
L38:
	;
	v118 = v110
	goto L39
L39:
	;
	if v118 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v137 = int32(1)
	if v21&v137 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v136 = v110 >> (uint(int32(31)) % 32) & v110
	goto L41
L43:
	;
	goto L44
L44:
	;
	v130 = v118 - int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v130))))
	if v132 == int32(32) {
		v118 = v130
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v136 = v118
	goto L41
L46:
	;
	v141 = v137
	goto L48
L47:
	;
	v141 = int32(4)
	goto L48
L48:
	;
	v143 = int32(1)
	if v80&v143 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v147 = v143
	goto L51
L50:
	;
	v147 = int32(4)
	goto L51
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = F_varstr_cmp(m, v12+v141, v77, v19+v147, v136, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v19 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v19)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return int32(base.Ui32(v150) >> (uint(int32(31)) % 32))
L60:
	;
	goto L59
}
func F_bpcharoctetlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_toast_raw_datum_size(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3 - int32(4)
	}
}
func F_brincostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 float64
	_ = v223
	var v224 float64
	_ = v224
	var v227 float64
	_ = v227
	var v228 int32
	_ = v228
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v238 float64
	_ = v238
	var v242 float64
	_ = v242
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v257 float64
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 float32
	_ = v392
	var v395 float64
	_ = v395
	var v396 float64
	_ = v396
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 float64
	_ = v441
	var v442 int32
	_ = v442
	var v443 float64
	_ = v443
	var v450 float64
	_ = v450
	var v452 float64
	_ = v452
	var v453 float64
	_ = v453
	var v454 float64
	_ = v454
	var v455 float64
	_ = v455
	var v463 float64
	_ = v463
	var v465 float64
	_ = v465
	var v466 int32
	_ = v466
	var v467 float64
	_ = v467
	var v468 int32
	_ = v468
	var v469 float64
	_ = v469
	var v472 float64
	_ = v472
	var v474 float64
	_ = v474
	var v479 float64
	_ = v479
	var v482 float64
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	v9 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(96)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v28 == v9 {
		v143 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v151 != 0 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 <= int32(0) {
		v143 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v44 = v31
	v45 = v9
	v50 = v9
	goto L4
L4:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v45<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 == int32(0) {
		v112 = v44
		v118 = v50
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v143 = v118
	goto L1
L6:
	;
	v125 = v45 + int32(1)
	if v125 < v112 {
		v44 = v112
		v45 = v125
		v50 = v118
		goto L4
	} else {
		goto L14
	}
L7:
	;
	v64 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v65 <= v64 {
		v112 = v44
		v118 = v50
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v77 = v64
	v84 = v50
	goto L9
L9:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v77<<(uint(int32(2))%32))))
	v95 = F_lappend(m, v84, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v112 = v101
	v118 = v95
	goto L6
L11:
	;
	return
L12:
	;
	v98 = v77 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v98 < v99 {
		v77 = v98
		v84 = v95
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L5
L15:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	F_get_tablespace_page_costs(m, v167, v25+int32(80), v25+int32(88))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L11
	} else {
		goto L19
	}
L16:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+68))
	v165 = v151 + v152<<(uint(int32(2))%32)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+52))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v149)+68))
	v165 = v158 + v159<<(uint(int32(2))%32) - int32(4)
	goto L15
L19:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+104)))
	if v174 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l6))) = int64(0)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v260 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L21:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v179 = F_index_open(m, v177, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L11
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v149)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = int32(128)
	v234 = base.F64_ceil(base.F64_mul(base.F64_convert_i32_u(v228), float64(0.0078125)))
	v235 = float64(1)
	if base.F64_gt(v234, v235) != 0 {
		goto L37
	} else {
		goto L38
	}
L24:
	;
	v182 = v25 + int32(72)
	v184 = F_ReadBuffer(m, v179, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	F_LockBuffer(m, v184, int32(1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	if v184 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v209 - int32(1)
	F_UnlockReleaseBuffer(m, v184)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L31
	}
L28:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192+(v184^int32(-1))<<(uint(int32(2))%32))))
	v206 = v198
	goto L27
L29:
	;
	goto L30
L30:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v206 = v200 + v184<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	F_relation_close(m, v179, int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v149)+116))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	v223 = base.F64_ceil(base.F64_div(base.F64_convert_i32_u(v218), base.F64_convert_i32_u(v220)))
	v224 = float64(1)
	if base.F64_gt(v223, v224) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v227 = v223
	goto L35
L34:
	;
	v227 = v224
	goto L35
L35:
	;
	v257 = v227
	goto L20
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+76)) = v250
	v257 = v238
	goto L20
L37:
	;
	v238 = v234
	goto L39
L38:
	;
	v238 = v235
	goto L39
L39:
	;
	v242 = base.F64_add(base.F64_div(v238, float64(1360)), float64(1))
	if base.F64_lt(v242, float64(4.294967296e+09))&base.F64_ge(v242, float64(0)) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v248 = base.I32_trunc_f64_u(v242)
	v250 = v248
	goto L36
L41:
	;
	goto L42
L42:
	;
	v250 = int32(0)
	goto L36
L43:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v149)+68))
	v439 = int32(0)
	v441 = F_clauselist_selectivity(m, l0, v143, v438, v439, v439)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L11
	} else {
		goto L89
	}
L44:
	;
	v263 = int32(0)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if v264 <= v263 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v280 = v263
	goto L46
L46:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v291 = int32(2)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290+v280<<(uint(v291)%32))))
	v295 = int32(*(*int16)(unsafe.Add(mBase, uint32(v294)+14)))
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289+v295<<(uint(v291)%32)))))
	if v299 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	goto L43
L48:
	;
	if v376 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = int32(1506)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v372
	v376 = v372
	goto L48
L50:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _consts[852]))
	if v301 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v336 = base.I32_extend16_s(v295 + int32(1))
	v338 = *(*int32)(unsafe.Add(mBase, _consts[853]))
	if v338 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v166)+16))
	v332 = F_SearchSysCache3(m, int32(65), v329, base.I32_extend16_s(v299), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L11
	} else {
		goto L62
	}
L54:
	;
	v307 = m.T0[v301].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v166, base.I32_extend16_s(v299), v25+int32(40))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	if v307 == int32(0) {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	if v311 == int32(0) {
		v376 = v311
		goto L48
	} else {
		goto L57
	}
L57:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	if v314 != 0 {
		v376 = v311
		goto L48
	} else {
		goto L58
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	F_errmsg_internal(m, int32(333616), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(514641), int32(8735), int32(369475))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v372 = v332
	goto L49
L63:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v368 = F_SearchSysCache3(m, int32(65), v366, v336, int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L11
	} else {
		goto L72
	}
L64:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v344 = m.T0[v338].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v341, v336, v25+int32(40))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	if v344 == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	if v348 == int32(0) {
		v376 = v348
		goto L48
	} else {
		goto L67
	}
L67:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	if v351 != 0 {
		v376 = v348
		goto L48
	} else {
		goto L68
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	F_errmsg_internal(m, int32(333616), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(514641), int32(8766), int32(369475))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v372 = v368
	goto L49
L73:
	;
	v413 = v280 + int32(1)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if v413 < v414 {
		v280 = v413
		goto L46
	} else {
		goto L88
	}
L74:
	;
	v385 = F_get_attstatsslot(m, v25+int32(4), v376, int32(3), int32(0), int32(2))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	if v385 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v387 <= int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	if v404 == int32(0) {
		goto L73
	} else {
		goto L86
	}
L79:
	;
	v395 = float64(0)
	goto L81
L80:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v392 = *(*float32)(unsafe.Add(mBase, uint32(v391)))
	v395 = base.F64_promote_f32(base.F32_abs(v392))
	goto L81
L81:
	;
	v396 = *(*float64)(unsafe.Add(mBase, uint32(l6)))
	if base.F64_gt(v395, v396) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l6))) = v395
	goto L84
L83:
	;
	goto L84
L84:
	;
	F_free_attstatsslot(m, v25+int32(4))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	goto L78
L86:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	m.T0[v407].(func(*base.Module, int32))(m, v404)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	goto L73
L88:
	;
	goto L47
L89:
	;
	v443 = *(*float64)(unsafe.Add(mBase, uint32(l6)))
	if base.F64_lt(v443, float64(1e-10)) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v463
	v465 = F_index_other_operands_eval_cost(m, l0, v143)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L11
	} else {
		goto L99
	}
L91:
	;
	v450 = base.F64_div(base.F64_ceil(base.F64_mul(v257, v441)), v443)
	if base.F64_lt(v450, v257) != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v453 = v257
	goto L93
L93:
	;
	v454 = float64(0)
	v455 = base.F64_div(v453, v257)
	if base.F64_lt(v455, v454) != 0 {
		v463 = v454
		goto L90
	} else {
		goto L97
	}
L94:
	;
	v452 = v450
	goto L96
L95:
	;
	v452 = v257
	goto L96
L96:
	;
	v453 = v452
	goto L93
L97:
	;
	if base.F64_gt(v455, float64(1)) == int32(0) {
		v463 = v455
		goto L90
	} else {
		goto L98
	}
L98:
	;
	v463 = float64(1)
	goto L90
L99:
	;
	v467 = *(*float64)(unsafe.Add(mBase, uint32(v25)+88))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v469 = base.F64_convert_i32_u(v468)
	v472 = base.F64_add(v465, base.F64_mul(l2, base.F64_mul(v467, v469)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v472
	v474 = *(*float64)(unsafe.Add(mBase, uint32(v25)+80))
	v479 = base.F64_add(base.F64_mul(base.F64_mul(v474, base.F64_sub(base.F64_convert_i32_u(v150), v469)), l2), v472)
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v479
	v482 = *(*float64)(unsafe.Add(mBase, _consts[382]))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(v453, base.F64_mul(v482, float64(0.1))), base.F64_convert_i32_u(v486)), v479)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l7))) = base.F64_convert_i32_u(v491)
	m.G0 = v25 + int32(96)
	return
}
func F_bringetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v278 int32
	_ = v278
	var v280 int64
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int64
	_ = v352
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v380 int64
	_ = v380
	var v383 int64
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v437 int32
	_ = v437
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v688 int64
	_ = v688
	var v689 int64
	_ = v689
	var v691 int64
	_ = v691
	var v722 int64
	_ = v722
	var v723 int64
	_ = v723
	var v729 int32
	_ = v729
	var v732 int64
	_ = v732
	var v733 int64
	_ = v733
	var v735 int64
	_ = v735
	var v736 int64
	_ = v736
	var v737 int64
	_ = v737
	var v739 int64
	_ = v739
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v771 int64
	_ = v771
	var v773 int64
	_ = v773
	var v774 int64
	_ = v774
	var v809 int64
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	v3 = int32(0)
	v26 = int64(0)
	v31 = m.G0
	v33 = v31 - int32(16)
	m.G0 = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v3
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+272))
	if v42 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v59 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+268)))
	if v45 != int32(1) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v53 = v42
	goto L4
L4:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v53)+16)) = v54 + int64(1)
	goto L1
L5:
	;
	F_pgstat_assoc_relation(m, v35)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int64(0)
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v35)+272))
	v53 = v52
	goto L4
L8:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
	*(*int64)(unsafe.Add(mBase, uint32(v59))) = v60 + int64(1)
	goto L10
L9:
	;
	goto L10
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v66 = F_IndexGetRelation(m, v64, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v69 = F_table_open(m, v66, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v72 = F_RelationGetNumberOfBlocksInFork(m, v69, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_sequence_close(m, v69, int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v81 = F_palloc0(m, v78*int32(28))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v102 = F_palloc(m, ((v84<<(uint(int32(3))%32)+int32(14))&int32(2147483632)+(v91<<(uint(int32(2))%32)+int32(7))&int32(2147483640)*v84)<<(uint(int32(1))%32))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v107 = v105 << (uint(int32(2)) % 32)
	v111 = (v107 + int32(7)) & int32(-8)
	v112 = v102 + v111
	v113 = v112 + v111
	v114 = v113 + v111
	v115 = int32(0)
	if v115 < v105 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v122 = v111 + v114
	v123 = int32(0)
	goto L20
L18:
	;
	v209 = v107
	goto L19
L19:
	;
	v211 = F__emscripten_memset_bulkmem(m, v113, base.I32_extend8_s(v115), v209)
	mBase = m.M
	goto L23
L20:
	;
	v150 = int32(2)
	v151 = v123 << (uint(v150) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v102+v151))) = v122
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v158 = int32(7)
	v160 = int32(-8)
	v162 = v122 + (v155<<(uint(v150)%32)+v158)&v160
	*(*int32)(unsafe.Add(mBase, uint32(v151+v112))) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v173 = v123 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v173 < v175 {
		v122 = v162 + (v164<<(uint(v150)%32)+v158)&v160
		v123 = v173
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v209 = v175 << (uint(int32(2)) % 32)
	goto L19
L22:
	;
	goto L21
L23:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v218 = F__emscripten_memset_bulkmem(m, v114, base.I32_extend8_s(int32(0)), v214<<(uint(int32(2))%32))
	mBase = m.M
	goto L24
L24:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v219 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v225 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v338 = F_brin_new_memtuple(m, v41)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L6
	} else {
		goto L42
	}
L28:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v256 = v253 + v225*int32(48)
	v257 = int32(*(*int16)(unsafe.Add(mBase, uint32(v256)+4)))
	v259 = v257 - int32(1)
	v262 = v81 + v259*int32(28)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v263 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	v267 = F_index_getprocinfo(m, v35, v257, int32(3))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v286 = v259 << (uint(int32(2)) % 32)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v289 = v287 & int32(1)
	if v289 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v273 = v262 + int32(16)
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v267)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v273))) = v274
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
	*(*int64)(unsafe.Add(mBase, uint32(v262))) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v267)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+24)) = v278
	v280 = *(*int64)(unsafe.Add(mBase, uint32(v267)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v262)+8)) = v280
	*(*int32)(unsafe.Add(mBase, uint32(v262)+20)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = int32(0)
	goto L34
L34:
	;
	goto L32
L35:
	;
	v290 = v112
	goto L37
L36:
	;
	v290 = v102
	goto L37
L37:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v286+v290)))
	if v289 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v293 = v218
	goto L40
L39:
	;
	v293 = v211
	goto L40
L40:
	;
	v294 = v293 + v286
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	*(*int32)(unsafe.Add(mBase, uint32(v292+v295<<(uint(int32(2))%32)))) = v256
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v301 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v300 + v301
	v305 = v225 + v301
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v305 < v306 {
		v225 = v305
		goto L28
	} else {
		goto L41
	}
L41:
	;
	goto L29
L42:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v346 = F_AllocSetContextCreateInternal(m, v341, int32(69532), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v348 = int32(4549024)
	v349 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v346
	if v72 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v352 = base.I64_extend_i32_u(v72)
	v367 = v338
	v371 = v3
	v380 = v26
	v383 = v26
	goto L47
L45:
	;
	v809 = int64(0)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v349
	F_MemoryContextDelete(m, v346)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L6
	} else {
		goto L113
	}
L47:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v386 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v809 = v771 * int64(10)
	goto L46
L49:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L6
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_MemoryContextReset(m, v346)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L6
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v397 = F_brinGetTupleForHeapBlock(m, v391, base.I32_wrap_i64(v380), v33+int32(12), v33+int32(6), v33)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L6
	} else {
		goto L56
	}
L54:
	;
	v773 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40))))
	v774 = v380 + v773
	if base.Ui64(v774) < base.Ui64(v352) {
		v367 = v755
		v371 = v759
		v380 = v774
		v383 = v771
		goto L47
	} else {
		goto L112
	}
L55:
	;
	v688 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40))))
	v689 = v380 + v688
	if base.Ui64(v689) < base.Ui64(v352) {
		goto L101
	} else {
		goto L102
	}
L56:
	;
	if v397 == int32(0) {
		v670 = v367
		v674 = v371
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v404 = F_brin_copy_tuple(m, v397, v401, v371, v33+int32(8))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	F_LockBuffer(m, v406, int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v410 = F_brin_deform_tuple(m, v41, v404, v367)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v412 != 0 {
		v670 = v410
		v674 = v404
		goto L55
	} else {
		goto L61
	}
L61:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	if v414 <= int32(0) {
		v670 = v410
		v674 = v404
		goto L55
	} else {
		goto L62
	}
L62:
	;
	v437 = int32(1)
	goto L63
L63:
	;
	v451 = v437 - int32(1)
	v453 = v451 << (uint(int32(2)) % 32)
	v454 = v211 + v453
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	if v455 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v670 = v410
	v674 = v404
	goto L55
L65:
	;
	v654 = v437 + int32(1)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	if v654 <= v656 {
		v437 = v654
		goto L63
	} else {
		goto L100
	}
L66:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v453+v218)))
	if v459 == int32(0) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+1)))
	if v462 != 0 {
		v755 = v410
		v759 = v404
		v771 = v383
		goto L54
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	v465 = v410 + int32(24) + v451*int32(20)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v453+(v41+int32(20)))))
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467)+2)))
	if v468 != int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v455 == int32(0) {
		goto L65
	} else {
		goto L86
	}
L72:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v453+v218)))
	if v472 <= int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v453+v112)))
	v481 = int32(0)
	goto L74
L74:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v476+v481<<(uint(int32(2))%32))))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)))
	if v512&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L71
L76:
	;
	v527 = v481 + int32(1)
	if v527 != v472 {
		v481 = v527
		goto L74
	} else {
		goto L85
	}
L77:
	;
	if v512&int32(64) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+3)))
	if v519 != 0 {
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v512&int32(128) == int32(0) {
		v755 = v410
		v759 = v404
		v771 = v383
		goto L54
	} else {
		goto L83
	}
L81:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+2)))
	if v520 != 0 {
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v755 = v410
	v759 = v404
	v771 = v383
	goto L54
L83:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+3)))
	if v525 != 0 {
		v755 = v410
		v759 = v404
		v771 = v383
		goto L54
	} else {
		goto L84
	}
L84:
	;
	goto L76
L85:
	;
	goto L75
L86:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465)+3)))
	if v561 != 0 {
		v755 = v410
		v759 = v404
		v771 = v383
		goto L54
	} else {
		goto L87
	}
L87:
	;
	v562 = v453 + v102
	v565 = v81 + v451*int32(28)
	v566 = int32(*(*int16)(unsafe.Add(mBase, uint32(v565)+8)))
	if v566 <= int32(3) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v582 = int32(0)
	goto L95
L89:
	;
	if v455 <= int32(0) {
		goto L65
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+12))
	v575 = F_FunctionCall4Coll(m, v565, v574, v41, v465, v572, v455)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L6
	} else {
		goto L93
	}
L92:
	;
	goto L88
L93:
	;
	if v575 == int32(0) {
		v755 = v410
		v759 = v404
		v771 = v383
		goto L54
	} else {
		goto L94
	}
L94:
	;
	goto L65
L95:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v562)))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v609+v582<<(uint(int32(2))%32))))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v613)+12))
	v615 = F_FunctionCall3Coll(m, v565, v614, v41, v465, v613)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L6
	} else {
		goto L97
	}
L96:
	;
	goto L65
L97:
	;
	if v615 == int32(0) {
		v755 = v410
		v759 = v404
		v771 = v383
		goto L54
	} else {
		goto L98
	}
L98:
	;
	v620 = v582 + int32(1)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	if v620 < v621 {
		v582 = v620
		goto L95
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	goto L64
L101:
	;
	v691 = v689
	goto L103
L102:
	;
	v691 = v352
	goto L103
L103:
	;
	if base.Ui64(v691-int64(1)) < base.Ui64(v380) {
		v755 = v670
		v759 = v674
		v771 = v383
		goto L54
	} else {
		goto L104
	}
L104:
	;
	v722 = v380
	v723 = v383
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v349
	F_tbm_add_page(m, l1, base.I32_wrap_i64(v722))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L6
	} else {
		goto L107
	}
L106:
	;
	v755 = v670
	v759 = v674
	v771 = v733
	goto L54
L107:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v346
	v732 = int64(1)
	v733 = v723 + v732
	v735 = v722 + v732
	v736 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40))))
	v737 = v380 + v736
	if base.Ui64(v737) < base.Ui64(v352) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v739 = v737
	goto L110
L109:
	;
	v739 = v352
	goto L110
L110:
	;
	if base.Ui64(v735) <= base.Ui64(v739-int64(1)) {
		v722 = v735
		v723 = v733
		goto L105
	} else {
		goto L111
	}
L111:
	;
	goto L106
L112:
	;
	goto L48
L113:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v814 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_ReleaseBuffer(m, v814)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L6
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	m.G0 = v33 + int32(16)
	return v809
L117:
	;
	goto L116
}
func F_brininsertcleanup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	if v3 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+136)) = int32(0)
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		F_brinRevmapTerminate(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_pfree(m, v3)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_btboolcmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v2 != v3) - base.B2i32(v5 != v3)
}
func F_btcharcmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	return v2 - v3
}
func F_btestimateparallelscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	v10 = l1<<(uint(int32(2))%32) + int32(40)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+10)))
	if v12 == int32(1) {
		v61 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v61
L2:
	;
	v15 = int32(0)
	v19 = F_datumEstimateSpace(m, v15, v15, int32(1), int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v12 < int32(2) {
		v61 = v10
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v27 = int32(1)
	v28 = v10
	goto L6
L6:
	;
	v33 = F_add_size(m, v28, int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v61 = v55
	goto L1
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v38 = v35 + v27<<(uint(int32(4))%32)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+10)))
	if v39 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v57 = v27 + int32(1)
	if v57 != v12 {
		v27 = v57
		v28 = v55
		goto L6
	} else {
		goto L17
	}
L10:
	;
	v42 = int32(0)
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+8)))
	v46 = F_datumEstimateSpace(m, v42, v42, int32(1), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v50 = F_add_size(m, v33, v19)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	v48 = F_add_size(m, v33, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v55 = v48
	goto L9
L15:
	;
	v53 = F_add_size(m, v50, int32(2704))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v55 = v53
	goto L9
L17:
	;
	goto L7
}
func F_btfloat4cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 float32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = int32(2147483647)
	v9 = base.I32_reinterpret_f32(v6) & v8
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = base.I32_reinterpret_f32(v10) & v8
	if base.Ui32(int32(2139095041)) <= base.Ui32(v13) {
		v22 = base.B2i32(base.Ui32(v9) < base.Ui32(int32(2139095041)))
		v30 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v13))|base.F32_lt(v6, v10))&v22
	} else {
		v18 = int32(1)
		if base.F32_gt(v6, v10) != 0 {
			v30 = v18
		} else {
			if base.Ui32(int32(2139095040)) < base.Ui32(v9) {
				v30 = v18
			} else {
				v22 = v18
				v30 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v13))|base.F32_lt(v6, v10))&v22
			}
		}
	}
	return v30
}
func F_btfloat8cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v7) & v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v15 = base.I64_reinterpret_f64(v12) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v15) {
		v24 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v32 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v24
	} else {
		v20 = int32(1)
		if base.F64_gt(v7, v12) != 0 {
			v32 = v20
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10) {
				v32 = v20
			} else {
				v24 = v20
				v32 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v24
			}
		}
	}
	return v32
}
func F_btfloat8fastcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 float64
	_ = v12
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v10 = int64(9223372036854775807)
	v11 = base.I64_reinterpret_f64(v8) & v10
	v12 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v15 = base.I64_reinterpret_f64(v12) & v10
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v15) {
		v24 = base.B2i32(base.Ui64(v11) < base.Ui64(int64(9218868437227405313)))
		v32 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v8, v12))&v24
	} else {
		v20 = int32(1)
		if base.F64_gt(v8, v12) != 0 {
			v32 = v20
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v11) {
				v32 = v20
			} else {
				v24 = v20
				v32 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v8, v12))&v24
			}
		}
	}
	return v32
}
func F_btint4sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(197)
	return int32(0)
}
func F_btint8sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(200)
	return int32(0)
}
func F_btrim(m *base.Module, l0 int32) int32 {
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
	var v86 int32
	_ = v86
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
			v86 = int32(1)
			v88 = F_dotrim(m, v21, v49, v57, v85, v86, v86)
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
func F_bttextsortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	v4 = int32(4549024)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v9
	F_varstr_sortsupport(m, v6, int32(25), v7)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v5
		return int32(0)
	}
}
func F_bttranslatecmptype(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	if base.Ui32(l0-int32(1)) < base.Ui32(int32(5)) {
		v8 = l0
	} else {
		v8 = int32(0)
	}
	return v8 & int32(65535)
}
func F_btvalidate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
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
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
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
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int64
	_ = v419
	var v422 int64
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int64
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	v16 = m.G0
	v18 = v16 - int32(240)
	m.G0 = v18
	v21 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	if int32(0) < v232 {
		goto L47
	} else {
		goto L48
	}
L2:
	;
	return int32(0)
L3:
	;
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
	v27 = v25 + v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	v31 = F_get_opfamily_name(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L44
	}
L7:
	;
	v35 = int32(0)
	v37 = F_SearchSysCacheList(m, int32(4), int32(1), v30, v35, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(1)
	v42 = int32(0)
	v44 = F_SearchSysCacheList(m, int32(5), v39, v30, v42, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v46 <= int32(0) {
		v221 = v39
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v51 = int32(0)
	v55 = v39
	goto L11
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(48)+v51<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+56))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+22)))
	v72 = v70 + v71
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+16)))
	switch v73 - int32(1) {
	case 0:
		goto L16
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	case 4:
		goto L19
	case 5:
		goto L18
	default:
		goto L17
	}
L12:
	;
	v221 = v197
	goto L1
L13:
	;
	v201 = v51 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v201 < v202 {
		v51 = v201
		v55 = v197
		goto L11
	} else {
		goto L43
	}
L14:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L39
	}
L15:
	;
	v161 = int32(0)
	v164 = F_errstart(m, int32(17), v161)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L37
	}
L16:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = v148
	v152 = int32(2)
	v156 = F_check_amproc_signature(m, v147, int32(23), int32(1), v152, v152, v18+int32(144))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L35
	}
L17:
	;
	v138 = int32(0)
	v141 = F_errstart(m, int32(17), v138)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L33
	}
L18:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+224)) = int32(2281)
	v129 = int32(1)
	v134 = F_check_amproc_signature(m, v125, int32(2278), v129, v129, v129, v18+int32(224))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L31
	}
L19:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v121 = F_check_amoptsproc_signature(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L2
	} else {
		goto L29
	}
L20:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = int32(26)
	v111 = int32(1)
	v116 = F_check_amproc_signature(m, v107, int32(16), v111, v111, v111, v18+int32(208))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L2
	} else {
		goto L27
	}
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+188)) = int64(68719476752)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v90
	v99 = int32(5)
	v103 = F_check_amproc_signature(m, v89, int32(16), int32(1), v99, v99, v18+int32(176))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L25
	}
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = int32(2281)
	v80 = int32(1)
	v85 = F_check_amproc_signature(m, v76, int32(2278), v80, v80, v80, v18+int32(160))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	if v85 == int32(0) {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v197 = v55
	goto L13
L25:
	;
	if v103 == int32(0) {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v197 = v55
	goto L13
L27:
	;
	if v116 == int32(0) {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v197 = v55
	goto L13
L29:
	;
	if v121 == int32(0) {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v197 = v55
	goto L13
L31:
	;
	if v134 == int32(0) {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v197 = v55
	goto L13
L33:
	;
	if v141 == int32(0) {
		v197 = v138
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v171 = int32(119)
	v174 = int32(489721)
	goto L14
L35:
	;
	if v156 != 0 {
		v197 = v55
		goto L13
	} else {
		goto L36
	}
L36:
	;
	goto L15
L37:
	;
	if v164 == int32(0) {
		v197 = v161
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v171 = int32(131)
	v174 = int32(489546)
	goto L14
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v179 = F_format_procedure(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = int32(426514)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v31
	F_errmsg(m, v174, v18+int32(128))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(518998), v171, int32(371624))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v197 = int32(0)
	goto L13
L43:
	;
	goto L12
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, int32(45619), v18)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(518998), int32(61), int32(371624))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v240 = int32(0)
	v242 = v221
	goto L50
L48:
	;
	v374 = v221
	goto L49
L49:
	;
	v385 = F_identify_opfamily_groups(m, v37, v44)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L2
	} else {
		goto L83
	}
L50:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(48)+v240<<(uint(int32(2))%32))))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+56))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+22)))
	v259 = v257 + v258
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+16)))
	if base.Ui32(int32(65530)) < base.Ui32((v260-int32(6))&int32(65535)) {
		v297 = v242
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v374 = v365
	goto L49
L52:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+18)))
	if v298 == int32(115) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v267 = int32(0)
	v270 = F_errstart(m, int32(17), v267)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	if v270 == int32(0) {
		v297 = v267
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	v278 = F_format_operator(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v280 = int32(*(*int16)(unsafe.Add(mBase, uint32(v259)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v280
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = int32(426514)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v31
	F_errmsg(m, int32(489340), v18+int32(112))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(518998), int32(151), int32(371624))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v297 = v267
	goto L52
L60:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v336 = F_check_amop_signature(m, v332, int32(16), v334, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L2
	} else {
		goto L72
	}
L61:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v259)+28))
	if v301 == int32(0) {
		v331 = v297
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v304 = int32(0)
	v307 = F_errstart(m, int32(17), v304)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L2
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if v307 == int32(0) {
		v331 = v304
		goto L60
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	v315 = F_format_operator(m, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = int32(426514)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v31
	F_errmsg(m, int32(189866), v18+int32(96))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(518998), int32(163), int32(371624))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v331 = v304
	goto L60
L71:
	;
	v367 = v240 + int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	if v367 < v368 {
		v240 = v367
		v242 = v365
		goto L50
	} else {
		goto L80
	}
L72:
	;
	if v336 != 0 {
		v365 = v331
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v338 = int32(0)
	v341 = F_errstart(m, int32(17), v338)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	if v341 == int32(0) {
		v365 = v338
		goto L71
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	v349 = F_format_operator(m, v348)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(426514)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v31
	F_errmsg(m, int32(378096), v18+int32(80))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(518998), int32(176), int32(371624))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v365 = v338
	goto L71
L80:
	;
	goto L51
L81:
	;
	if v578 != 0 {
		goto L133
	} else {
		goto L134
	}
L82:
	;
	v554 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L2
	} else {
		goto L125
	}
L83:
	;
	if v385 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v389 = int32(0)
	v540 = v389
	v542 = v389
	goto L82
L85:
	;
	goto L86
L86:
	;
	v391 = int32(0)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v392 <= v391 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v534 == int32(0) {
		v578 = v522
		v579 = v523
		v580 = v524
		goto L81
	} else {
		goto L124
	}
L88:
	;
	v522 = v391
	v523 = v374
	v524 = int32(0)
	v534 = int32(1)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v397 = int32(0)
	v401 = v397
	v402 = v391
	v403 = v374
	v404 = v397
	v410 = int32(0)
	goto L91
L91:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v385)+12))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414+v401<<(uint(int32(2))%32))))
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v418)+8))
	if v419 == int64(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v522 = v508
	v523 = v509
	v524 = v510
	v534 = base.B2i32(v512 == int32(0))
	goto L87
L93:
	;
	v514 = v401 + int32(1)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v514 < v515 {
		v401 = v514
		v402 = v508
		v403 = v509
		v404 = v510
		v410 = v512
		goto L91
	} else {
		goto L123
	}
L94:
	;
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v418)+16))
	if v422 == int64(8) {
		v508 = v402
		v509 = v403
		v510 = v404
		v512 = v410
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	if v28 == v425 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	if v427 == v28 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v430 = v410
	goto L100
L100:
	;
	v431 = F_list_append_unique_oid(m, v402, v425)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L2
	} else {
		goto L104
	}
L101:
	;
	v429 = v418
	goto L103
L102:
	;
	v429 = v410
	goto L103
L103:
	;
	v430 = v429
	goto L100
L104:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v434 = F_list_append_unique_oid(m, v431, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	v436 = *(*int64)(unsafe.Add(mBase, uint32(v418)+8))
	if v436 == int64(62) {
		v470 = v403
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v473 = v404 + int32(1)
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+16)))
	if v474&int32(2) != 0 {
		v508 = v434
		v509 = v470
		v510 = v473
		v512 = v430
		goto L93
	} else {
		goto L115
	}
L107:
	;
	v439 = int32(0)
	v442 = F_errstart(m, int32(17), v439)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	if v442 == int32(0) {
		v470 = v439
		goto L106
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L2
	} else {
		goto L110
	}
L110:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v450 = F_format_type_be(m, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v453 = F_format_type_be(m, v452)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = int32(426514)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v31
	F_errmsg(m, int32(205398), v18-int32(-64))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(518998), int32(235), int32(371624))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v470 = v439
	goto L106
L115:
	;
	v477 = int32(0)
	v480 = F_errstart(m, int32(17), v477)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	if v480 == int32(0) {
		v508 = v434
		v509 = v477
		v510 = v473
		v512 = v430
		goto L93
	} else {
		goto L117
	}
L117:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v488 = F_format_type_be(m, v487)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v491 = F_format_type_be(m, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = int32(426514)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v31
	F_errmsg(m, int32(205309), v18+int32(48))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(518998), int32(245), int32(371624))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v508 = v434
	v509 = v477
	v510 = v473
	v512 = v430
	goto L93
L123:
	;
	goto L92
L124:
	;
	v540 = v522
	v542 = v524
	goto L82
L125:
	;
	if v554 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L2
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v578 = v540
	v579 = int32(0)
	v580 = v542
	goto L81
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(426514)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v27 + int32(8)
	F_errmsg(m, int32(699937), v18+int32(32))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(518998), int32(257), int32(371624))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	F_ReleaseCatCacheList(m, v44)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L2
	} else {
		goto L142
	}
L133:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	v592 = v590
	goto L135
L134:
	;
	v592 = int32(0)
	goto L135
L135:
	;
	if v580 == v592*v592 {
		v618 = v579
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v595 = int32(0)
	v598 = F_errstart(m, int32(17), v595)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L2
	} else {
		goto L137
	}
L137:
	;
	if v598 == int32(0) {
		v618 = v595
		goto L132
	} else {
		goto L138
	}
L138:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(426514)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v31
	F_errmsg(m, int32(700000), v18+int32(16))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(518998), int32(273), int32(371624))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v618 = v595
	goto L132
L142:
	;
	F_ReleaseCatCacheList(m, v37)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	F_ReleaseCatCache(m, v21)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	m.G0 = v18 + int32(240)
	return v618 & int32(1)
}
func F_build_merged_partition_bounds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	if l1 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v14 = v12
	} else {
		v14 = int32(0)
	}
	v16 = F_palloc(m, int32(36))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
		v23 = v14 << (uint(int32(2)) % 32)
		v24 = F_palloc(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v24
			if l1 == int32(0) {
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v29 <= int32(0) {
				} else {
					v39 = int32(0)
					for {
						v44 = v39 << (uint(int32(2)) % 32)
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v47+v44)))
						*(*int32)(unsafe.Add(mBase, uint32(v44+v45))) = v49
						v52 = v39 + int32(1)
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if v52 < v53 {
							v39 = v52
							continue
						} else {
							break
						}
						break
					}
				}
			}
			if l0 == int32(114) {
				v68 = F_palloc(m, v23)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v68
					if l2 == int32(0) {
					} else {
						v73 = int32(0)
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						if v74 <= v73 {
						} else {
							v84 = v73
							for {
								v89 = v84 << (uint(int32(2)) % 32)
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
								v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v92+v89)))
								*(*int32)(unsafe.Add(mBase, uint32(v89+v90))) = v94
								v97 = v84 + int32(1)
								v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								if v97 < v98 {
									v84 = v97
									continue
								} else {
									break
								}
								break
							}
						}
					}
					v112 = v14 + int32(1)
					v116 = F_lappend_int(m, l3, int32(-1))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						v123 = v116
						v128 = v112
						v129 = v112 << (uint(int32(2)) % 32)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v128
						v132 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v132
						v135 = F_palloc(m, v129)
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v135
							if v123 == int32(0) {
							} else {
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
								if v140 <= int32(0) {
								} else {
									v150 = v132
									for {
										v155 = v150 << (uint(int32(2)) % 32)
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
										v158 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
										v160 = *(*int32)(unsafe.Add(mBase, uint32(v158+v155)))
										*(*int32)(unsafe.Add(mBase, uint32(v155+v156))) = v160
										v163 = v150 + int32(1)
										v164 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
										if v163 < v164 {
											v150 = v163
											continue
										} else {
											break
										}
										break
									}
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l5
							*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l4
							return v16
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
				v123 = l3
				v128 = v14
				v129 = v23
				*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v128
				v132 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v132
				v135 = F_palloc(m, v129)
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v135
					if v123 == int32(0) {
					} else {
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
						if v140 <= int32(0) {
						} else {
							v150 = v132
							for {
								v155 = v150 << (uint(int32(2)) % 32)
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
								v158 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v158+v155)))
								*(*int32)(unsafe.Add(mBase, uint32(v155+v156))) = v160
								v163 = v150 + int32(1)
								v164 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
								if v163 < v164 {
									v150 = v163
									continue
								} else {
									break
								}
								break
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l4
					return v16
				}
			}
		}
	}
}
func F_build_setop_child_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v340 int32
	_ = v340
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 float64
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v469 int32
	_ = v469
	var v470 float64
	_ = v470
	var v471 int32
	_ = v471
	var v473 float64
	_ = v473
	var v474 int32
	_ = v474
	var v491 float64
	_ = v491
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+180))
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l4 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_set_subquery_size_estimates(m, l0, l1)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L32
	}
L4:
	;
	goto L3
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v25 = v23
	goto L7
L6:
	;
	v25 = int32(0)
	goto L7
L7:
	;
	if l3 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L18
	} else {
		goto L29
	}
L9:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v106 != 0 {
		goto L25
	} else {
		goto L26
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v28 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v37 = v25
	v40 = v7
	goto L12
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v40<<(uint(int32(2))%32))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+26)))
	if v52 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	if v37 == int32(0) {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	v81 = v37
	goto L16
L16:
	;
	v85 = v40 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v85 < v86 {
		v37 = v81
		v40 = v85
		goto L12
	} else {
		goto L24
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v66 = F_exprType(m, v60)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_add_child_eq_member(m, l0, v58, int32(-1), v60, v61, v65, v64, v66, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v72 = v37 + int32(4)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.Ui32(v72) < base.Ui32(v74+v75<<(uint(int32(2))%32)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v80 = v72
	goto L23
L22:
	;
	v80 = int32(0)
	goto L23
L23:
	;
	v81 = v80
	goto L16
L24:
	;
	goto L13
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v111 = v107 - int32(1)
	goto L27
L26:
	;
	v111 = int32(-1)
	goto L27
L27:
	;
	v112 = F_bms_add_range(m, v104, int32(0), v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+136)) = v112
	goto L4
L29:
	;
	F_errmsg_internal(m, int32(270264), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(513170), int32(3100), int32(180400))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v149 = F_fetch_upper_rel(m, v146, int32(7), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v151)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149)+32))
	if v153 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _consts[393]))
	if v394 != 0 {
		goto L111
	} else {
		goto L112
	}
L35:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v365 != 0 {
		goto L34
	} else {
		goto L107
	}
L36:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if int32(0) < v154 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v151&int32(1) == int32(0) {
		goto L34
	} else {
		goto L106
	}
L39:
	;
	v167 = v7
	goto L42
L40:
	;
	goto L41
L41:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v340&int32(1) == int32(0) {
		goto L34
	} else {
		goto L105
	}
L42:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173+v167<<(uint(int32(2))%32))))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v149)+48))
	if v177 == v178 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)+64))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v182 = F_make_tlist_from_pathtarget(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L18
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if l4 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v184 = F_convert_subquery_pathkeys(m, l0, l1, v180, v182)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v187 = F_create_subqueryscan_path(m, l0, l1, v177, l2, v184, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_add_path(m, l1, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v321 = v167 + int32(1)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v321 < v322 {
		v167 = v321
		goto L42
	} else {
		goto L104
	}
L52:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v177)+64))
	v195 = v19 + int32(12)
	if v22 == v193 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	if v298 == v178 {
		goto L51
	} else {
		goto L99
	}
L54:
	;
	if v273 != 0 {
		v298 = v177
		goto L53
	} else {
		goto L86
	}
L55:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v261
	v273 = int32(1)
	goto L54
L56:
	;
	if v22 != 0 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v22 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = int32(0)
	v273 = int32(1)
	goto L54
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = int32(0)
	v273 = int32(1)
	goto L54
L61:
	;
	goto L62
L62:
	;
	if v193 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v213 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v213
	v273 = v213
	goto L54
L64:
	;
	goto L65
L65:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v217 = int32(0)
	if v217 < v216 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v220 = v216
	goto L68
L67:
	;
	v220 = v217
	goto L68
L68:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v225 = int32(0)
	goto L69
L69:
	;
	if v225 < v221 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v237 = v233 + v225<<(uint(int32(2))%32)
	goto L73
L72:
	;
	v237 = int32(0)
	goto L73
L73:
	;
	if v225 == v220 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v220
	v273 = base.B2i32(v237 == int32(0))
	goto L54
L75:
	;
	goto L76
L76:
	;
	v243 = base.B2i32(v237 == int32(0))
	if v237 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v225
	v273 = v243
	goto L54
L78:
	;
	goto L79
L79:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v250 = v247 + v225<<(uint(int32(2))%32)
	if v250 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v225
	v273 = v243
	goto L54
L81:
	;
	goto L82
L82:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v254 != v255 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v225
	v273 = int32(0)
	goto L54
L84:
	;
	v225 = v225 + int32(1)
	goto L69
L86:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v274)+304))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v177 != v178 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v296 = F_create_incremental_sort_path(m, v274, v149, v177, v22, v276, v275)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L18
	} else {
		goto L98
	}
L88:
	;
	v293 = F_create_sort_path(m, v149, v177, v22, v275)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L18
	} else {
		goto L97
	}
L89:
	;
	if v276 == int32(0) {
		goto L51
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if v276 == int32(0) {
		goto L88
	} else {
		goto L95
	}
L92:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, _consts[392])))
	if v281 == int32(0) {
		goto L51
	} else {
		goto L93
	}
L93:
	;
	if v281 == int32(0) {
		goto L88
	} else {
		goto L94
	}
L94:
	;
	goto L87
L95:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, _consts[392])))
	if v289&int32(1) != 0 {
		goto L87
	} else {
		goto L96
	}
L96:
	;
	goto L88
L97:
	;
	v298 = v293
	goto L53
L98:
	;
	v298 = v296
	goto L53
L99:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v298)+64))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v306 = F_make_tlist_from_pathtarget(m, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L18
	} else {
		goto L100
	}
L100:
	;
	v308 = F_convert_subquery_pathkeys(m, l0, l1, v304, v306)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L18
	} else {
		goto L101
	}
L101:
	;
	v311 = F_create_subqueryscan_path(m, l0, l1, v298, l2, v308, int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L18
	} else {
		goto L102
	}
L102:
	;
	F_add_path(m, l1, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L18
	} else {
		goto L103
	}
L103:
	;
	goto L51
L104:
	;
	goto L43
L105:
	;
	goto L35
L106:
	;
	goto L35
L107:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v149)+40))
	if v366 == int32(0) {
		goto L34
	} else {
		goto L108
	}
L108:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v366)+12))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v371 = int32(0)
	v373 = F_create_subqueryscan_path(m, l0, l1, v370, l2, v371, v371)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L18
	} else {
		goto L109
	}
L109:
	;
	F_add_partial_path(m, l1, v373)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L18
	} else {
		goto L110
	}
L110:
	;
	goto L34
L111:
	;
	v395 = int32(0)
	m.T0[v394].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, v395, v395, l1, v395)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L18
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L18
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	if l5 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+100))
	if v404 != 0 {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	goto L118
L118:
	;
	m.G0 = v19 + int32(16)
	return
L119:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v491
	goto L118
L120:
	;
	v413 = int32(0)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v403)+76))
	if v415 != 0 {
		goto L129
	} else {
		goto L130
	}
L121:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v412 = *(*float64)(unsafe.Add(mBase, uint32(v411)+32))
	v491 = v412
	goto L119
L122:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v403)+108))
	if v405 != 0 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v403)+120))
	if v406 != 0 {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+318)))
	if v407 != 0 {
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+36)))
	if v408 != int32(1) {
		goto L120
	} else {
		goto L126
	}
L126:
	;
	goto L121
L127:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v470 = *(*float64)(unsafe.Add(mBase, uint32(v469)+32))
	v471 = int32(0)
	v473 = F_estimate_num_groups(m, v402, v453, v470, v471, v471)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L18
	} else {
		goto L140
	}
L128:
	;
	v419 = v413
	v421 = v413
	goto L133
L129:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	if int32(0) < v416 {
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v453 = v413
	goto L127
L132:
	;
	goto L131
L133:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v415)+12))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v435+v421<<(uint(int32(2))%32))))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439)+26)))
	if v440&int32(1) == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v453 = v448
	goto L127
L135:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v446 = F_lappend(m, v419, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L18
	} else {
		goto L138
	}
L136:
	;
	v448 = v419
	goto L137
L137:
	;
	v450 = v421 + int32(1)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	if v450 < v451 {
		v419 = v448
		v421 = v450
		goto L133
	} else {
		goto L139
	}
L138:
	;
	v448 = v446
	goto L137
L139:
	;
	goto L134
L140:
	;
	v491 = v473
	goto L119
}
func F_build_subplan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v30 int32
	_ = v30
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v133 int32
	_ = v133
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 float64
	_ = v318
	var v319 int32
	_ = v319
	var v330 float64
	_ = v330
	var v332 int32
	_ = v332
	var v336 float64
	_ = v336
	var v337 float64
	_ = v337
	var v340 float64
	_ = v340
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
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
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
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
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v542 int32
	_ = v542
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	v10 = l9
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v21 = F_palloc0(m, int32(72))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(23)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v30 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+37)) = uint8(v10)
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+36)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v52
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+37)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+38)) = uint8(v58)
	if l4 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = int64(-4294965018)
	v52 = int32(0)
	goto L3
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+26)))
	if v35 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v37 = F_exprType(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v37
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v41 = F_exprTypmod(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v45 = F_exprCollation(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v52 = v45
	goto L3
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if int32(0) < v64 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v152 = int32(1)
	goto L12
L12:
	;
	v154 = v21 + int32(40)
	if l5 != 0 {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v79 = v54
	goto L16
L14:
	;
	goto L15
L15:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v152 = base.B2i32(v133 == int32(0))
	goto L12
L16:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v83 = int32(2)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v79<<(uint(v83)%32))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if base.Ui32(v88-int32(9)) < base.Ui32(v83) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L15
L18:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v107 = F_lappend_int(m, v105, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L24
	}
L19:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+44)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l0
	v102 = F_process_sublinks_mutator(m, v87, v18+int32(40))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	if v88 == int32(319) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	if v88 != int32(61) {
		v104 = v87
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v104 = v102
	goto L18
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v107
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v111 = F_lappend(m, v110, v104)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v111
	v115 = v79 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v115 < v116 {
		v79 = v115
		goto L16
	} else {
		goto L26
	}
L26:
	;
	goto L17
L27:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+8))
	v573 = F_lappend(m, v572, v557)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L132
	}
L28:
	;
	v172 = int32(1)
	v174 = v152 ^ v172
	if v174|base.B2i32(l5 != int32(4)) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	if v152 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v160 = F_generate_new_exec_param(m, l0, int32(16), int32(-1), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v162
	v168 = F_list_make1_impl(m, int32(471), v18+int32(8))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v168
	v557 = l1
	v560 = v160
	v568 = int32(1)
	goto L27
L33:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v184 = F_exprType(m, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if base.B2i32(l5 != int32(6))|v174 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v187 = F_exprTypmod(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v190 = F_exprCollation(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v192 = F_generate_new_exec_param(m, l0, v184, v187, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v194
	v200 = F_list_make1_impl(m, int32(471), v18+int32(12))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v200
	v557 = l1
	v560 = v192
	v568 = v172
	goto L27
L41:
	;
	v557 = v542
	v560 = v21
	v568 = v553
	goto L27
L42:
	;
	v542 = l1
	v553 = int32(0)
	goto L41
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L127
	}
L44:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v212 = F_exprType(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v237 = v21 + int32(12)
	if v152^int32(1)|base.B2i32(l5 != int32(3)) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	v214 = F_get_promoted_array_type(m, v212)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v214 == int32(0) {
		goto L43
	} else {
		goto L49
	}
L49:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v219 = F_exprTypmod(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v222 = F_exprCollation(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v224 = F_generate_new_exec_param(m, l0, v214, v219, v222)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v226
	v232 = F_list_make1_impl(m, int32(471), v18+int32(24))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v232
	v557 = l1
	v560 = v224
	v568 = int32(1)
	goto L27
L54:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v246 = F_generate_subquery_params(m, l0, v245, v237)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if l5 == int32(5) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l0
	v252 = F_convert_testexpr_mutator(m, l7, v18+int32(40))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v255 = F_list_copy(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v255
	v557 = l1
	v560 = v252
	v568 = v172
	goto L27
L60:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v261 = F_generate_subquery_params(m, l0, v260, v154)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if l7 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L63:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v273 = v263
	goto L64
L64:
	;
	if v273 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v273)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v287+l6<<(uint(int32(2))%32)-int32(4)))) = v261
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v294 != 0 {
		goto L42
	} else {
		goto L73
	}
L66:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v281 = v279
	goto L68
L67:
	;
	v281 = int32(0)
	goto L68
L68:
	;
	if v281 < l6 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v284 = F_lappend(m, v273, int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	goto L65
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v284
	v273 = v284
	goto L64
L73:
	;
	v298 = F_makeNullConst(m, int32(2249), int32(-1), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v557 = l1
	v560 = v298
	v568 = v172
	goto L27
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v312
	if l5 != int32(2) {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = l8
	v312 = l7
	goto L75
L77:
	;
	if l8 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v303 = F_generate_subquery_params(m, l0, v302, v237)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l0
	v309 = F_convert_testexpr_mutator(m, l7, v18+int32(40))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v312 = v309
	goto L75
L81:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v489 != 0 {
		goto L42
	} else {
		goto L122
	}
L82:
	;
	v316 = int32(0)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v317 != 0 {
		v542 = l1
		v553 = v316
		goto L41
	} else {
		goto L83
	}
L83:
	;
	v318 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v330 = *(*float64)(unsafe.Add(mBase, _consts[343]))
	v332 = *(*int32)(unsafe.Add(mBase, _consts[324]))
	v336 = base.F64_mul(base.F64_mul(v330, base.F64_convert_i32_s(v332)), float64(1024))
	v337 = float64(4.294967295e+09)
	if base.F64_lt(v336, v337) != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if base.F64_gt(base.F64_mul(v318, base.F64_convert_i32_u((v319+int32(7))&int32(-8)+int32(24))), base.F64_convert_i32_u(v348)) != 0 {
		goto L81
	} else {
		goto L91
	}
L85:
	;
	v340 = v336
	goto L87
L86:
	;
	v340 = v337
	goto L87
L87:
	;
	if base.F64_lt(v340, float64(4.294967296e+09))&base.F64_ge(v340, float64(0)) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v346 = base.I32_trunc_f64_u(v340)
	v348 = v346
	goto L84
L89:
	;
	goto L90
L90:
	;
	v348 = int32(0)
	goto L84
L91:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v352 = int32(0)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v353 == v352 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if v469 == int32(0) {
		goto L81
	} else {
		goto L121
	}
L93:
	;
	v469 = int32(0)
	goto L92
L94:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v353)))
	switch v356 - int32(17) {
	case 0:
		goto L97
	default:
		goto L93
	case 4:
		goto L96
	}
L95:
	;
	v469 = v442
	goto L92
L96:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	if v381 != 0 {
		goto L93
	} else {
		goto L106
	}
L97:
	;
	v359 = F_hash_ok_operator(m, v353)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v359 == int32(0) {
		goto L93
	} else {
		goto L99
	}
L99:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v353)+28))
	if v363 == int32(0) {
		goto L93
	} else {
		goto L100
	}
L100:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	if v366 != int32(2) {
		goto L93
	} else {
		goto L101
	}
L101:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v363)+12))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v371 = F_contain_exec_param(m, v370, v351)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v371 != 0 {
		goto L93
	} else {
		goto L103
	}
L103:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v353)+28))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+12))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	v377 = F_contain_var_clause(m, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	if v377 == int32(0) {
		v442 = int32(1)
		goto L95
	} else {
		goto L105
	}
L105:
	;
	goto L93
L106:
	;
	v382 = int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v353)+8))
	if v383 == int32(0) {
		v442 = v382
		goto L95
	} else {
		goto L107
	}
L107:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v386 <= int32(0) {
		v442 = v382
		goto L95
	} else {
		goto L108
	}
L108:
	;
	v396 = v352
	goto L109
L109:
	;
	v404 = int32(0)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v405+v396<<(uint(int32(2))%32))))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	if v410 != int32(17) {
		v442 = v404
		goto L95
	} else {
		goto L111
	}
L110:
	;
	v442 = v432
	goto L95
L111:
	;
	v413 = F_hash_ok_operator(m, v409)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if v413 == int32(0) {
		v442 = v404
		goto L95
	} else {
		goto L113
	}
L113:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v409)+28))
	if v417 == int32(0) {
		v442 = v404
		goto L95
	} else {
		goto L114
	}
L114:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v417)+4))
	if v420 != int32(2) {
		v442 = v404
		goto L95
	} else {
		goto L115
	}
L115:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	v425 = F_contain_exec_param(m, v424, v351)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v425 != 0 {
		v442 = v404
		goto L95
	} else {
		goto L117
	}
L117:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v409)+28))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	v430 = F_contain_var_clause(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	if v430 != 0 {
		v442 = v404
		goto L95
	} else {
		goto L119
	}
L119:
	;
	v432 = int32(1)
	v434 = v396 + v432
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v434 < v435 {
		v396 = v434
		goto L109
	} else {
		goto L120
	}
L120:
	;
	goto L110
L121:
	;
	v472 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+36)) = uint8(v472)
	v542 = l1
	v553 = v316
	goto L41
L122:
	;
	v490 = int32(0)
	v492 = int32(*(*uint8)(unsafe.Add(mBase, _consts[388])))
	if v492 != int32(1) {
		v542 = l1
		v553 = v490
		goto L41
	} else {
		goto L123
	}
L123:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v497 = v495 - int32(348)
	goto L124
L124:
	;
	if base.B2i32(base.Ui32(v497) < base.Ui32(int32(15)))&int32(base.Ui32(int32(20541))>>(uint(v497)%32)) != 0 {
		v542 = l1
		v553 = v490
		goto L41
	} else {
		goto L125
	}
L125:
	;
	v503 = F_materialize_finished_plan(m, l1)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v542 = v503
	v553 = v490
	goto L41
L127:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v510 = F_exprType(m, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v512 = F_format_type_be(m, v510)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v512
	F_errmsg_internal(m, int32(197105), v18+int32(16))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(512850), int32(423), int32(294206))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v575)+8)) = v573
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v577)+12))
	v579 = F_lappend(m, v578, l2)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v581)+12)) = v579
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)+16))
	v585 = F_lappend(m, v584, l3)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v587)+16)) = v585
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v589)+8))
	if v590 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+4))
	v593 = v591
	goto L137
L136:
	;
	v593 = int32(0)
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v593
	if v568 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v610
	v615 = F_psprintf(m, int32(488525), v18)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L146
	}
L139:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v596 = F_lappend(m, v595, v21)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v600 = int32(294994)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v601 != 0 {
		v610 = v600
		goto L138
	} else {
		goto L143
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v596
	v610 = int32(294699)
	goto L138
L143:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+36)))
	if v602 != 0 {
		v610 = v600
		goto L138
	} else {
		goto L144
	}
L144:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v605 = F_bms_add_member(m, v604, v593)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v607)+20)) = v605
	v610 = v600
	goto L138
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v615
	F_cost_subplan(m, v21, v557)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	m.G0 = v18 + int32(48)
	return v560
}
func F_builtin_locale_encoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v8 != int32(67) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return v88
L2:
	;
	v13 = int32(6)
	v14 = int32(578035)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[808])))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 == int32(0) {
		v37 = v17
		v38 = v18
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v11 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v88 = int32(-1)
	goto L1
L5:
	;
	if v38-v37 == int32(0) {
		v88 = v13
		goto L1
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	if v17 != v18 {
		v37 = v17
		v38 = v18
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v22 = l0
	v23 = v14
	goto L9
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v26
		v38 = v27
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v37 = v26
	v38 = v27
	goto L6
L11:
	;
	v30 = int32(1)
	if v26 == v27 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v42 = int32(538446)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[809])))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v46 == int32(0) {
		v65 = v45
		v66 = v46
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v66-v65 == int32(0) {
		v88 = v13
		goto L1
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	if v45 != v46 {
		v65 = v45
		v66 = v46
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = l0
	v51 = v42
	goto L18
L18:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v55 == int32(0) {
		v65 = v54
		v66 = v55
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v65 = v54
	v66 = v55
	goto L15
L20:
	;
	v58 = int32(1)
	if v54 == v55 {
		v50 = v50 + v58
		v51 = v51 + v58
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(237287), v6)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(519586), int32(1499), int32(349928))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_byteale(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v47 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v19 = int32(4)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v21&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(1)
	if v16&v34 != 0 {
		v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v30 = v19
	goto L10
L9:
	;
	v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
	goto L10
L10:
	;
	if v21 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v19
	goto L13
L12:
	;
	v33 = v30
	goto L13
L13:
	;
	v46 = v33
	goto L4
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v78 = int32(1)
	if v16&v78 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v50 = int32(4)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52&int32(254) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v65 = int32(1)
	if v47&v65 != 0 {
		v77 = int32(base.Ui32(v47)>>(uint(v65)%32)) - v65
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v61 = v50
	goto L21
L20:
	;
	v61 = base.B2i32(v52 == int32(18)) << (uint(v50) % 32)
	goto L21
L21:
	;
	if v52 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v64 = v50
	goto L24
L23:
	;
	v64 = v61
	goto L24
L24:
	;
	v77 = v64
	goto L15
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v77 = int32(base.Ui32(v71)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v82 = v78
	goto L28
L27:
	;
	v82 = int32(4)
	goto L28
L28:
	;
	v83 = v9 + v82
	v84 = int32(1)
	if v47&v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v88 = v84
	goto L31
L30:
	;
	v88 = int32(4)
	goto L31
L31:
	;
	v89 = v14 + v88
	if v46 < v77 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v91 = v46
	goto L34
L33:
	;
	v91 = v77
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v91) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v154 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v153 = int32(0)
	goto L35
L37:
	;
	v127 = v122
	v128 = v123
	v129 = v124
	goto L47
L38:
	;
	if (v83|v89)&int32(3) != 0 {
		v122 = v83
		v123 = v89
		v124 = v91
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v115 = v83
	v116 = v89
	v117 = v91
	goto L40
L40:
	;
	if v117 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v99 = v83
	v100 = v89
	v101 = v91
	goto L42
L42:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v104 != v105 {
		v122 = v99
		v123 = v100
		v124 = v101
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v115 = v110
	v116 = v108
	v117 = v112
	goto L40
L44:
	;
	v107 = int32(4)
	v108 = v100 + v107
	v110 = v99 + v107
	v112 = v101 - v107
	if base.Ui32(int32(3)) < base.Ui32(v112) {
		v99 = v110
		v100 = v108
		v101 = v112
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v122 = v115
	v123 = v116
	v124 = v117
	goto L37
L47:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v132 == v133 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v153 = v132 - v133
	goto L35
L49:
	;
	v135 = int32(1)
	v140 = v129 - v135
	if v140 != 0 {
		v127 = v127 + v135
		v128 = v128 + v135
		v129 = v140
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v158 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v162 = int32(0)
	return base.B2i32(v153 == v162)&base.B2i32(v46 <= v77) | base.B2i32(v153 < v162)
L60:
	;
	goto L59
}
