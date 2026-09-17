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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_BeginReportingGUCOptions[0]))
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
	*(*uint8)(unsafe.Add(mBase, _c_F_BeginReportingGUCOptions[1])) = uint8(v12)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BeginReportingGUCOptions[2])))
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
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_BeginReportingGUCOptions[3]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+316))
	v24 = base.B2i32(v22 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_BeginReportingGUCOptions[2])) = uint8(v24)
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
	v37 = F_set_config_with_handle(m, int32(_a_F_BeginReportingGUCOptions_0), v28, int32(_a_F_BeginReportingGUCOptions_1), v28, v31, v31, v28, int32(1), v28, v28)
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
	v40 = v5 + int32(12)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_BeginReportingGUCOptions[4]))
	F_hash_seq_init(m, v40, v42)
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
	v45 = F_hash_seq_search(m, v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if v45 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v49 = v45
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+20)))
	if v52&int32(64) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L1
L17:
	;
	F_ReportGUCOption(m, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v59 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	if v59 != 0 {
		v49 = v59
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
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v30 int64
	_ = v30
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = l0 + int32(16)
	v11 = base.I64_extend_i32_u(l3)
	v14 = v11 + int64(4354685564936845354)
	v15 = int64(30)
	v18 = int64(-4658895280553007687)
	v19 = (int64(base.Ui64(v14)>>(uint(v15)%64)) ^ v14) * v18
	v20 = int64(27)
	v23 = int64(-7723592293110705685)
	v24 = (int64(base.Ui64(v19)>>(uint(v20)%64)) ^ v19) * v23
	v25 = int64(31)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(base.Ui64(v24)>>(uint(v25)%64)) ^ v24
	v30 = v11 - int64(7046029254386353131)
	v35 = (int64(base.Ui64(v30)>>(uint(v15)%64)) ^ v30) * v18
	v40 = (int64(base.Ui64(v35)>>(uint(v20)%64)) ^ v35) * v23
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(base.Ui64(v40)>>(uint(v25)%64)) ^ v40
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v45) < base.Ui32(v46) {
		v48 = v45
	} else {
		v48 = v46
	}
	return v48
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
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
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
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1750 int32
	_ = v1750
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1825 int32
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1894 int32
	_ = v1894
	var v1898 int32
	_ = v1898
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
	v100 = v85&int32(63) | (v43<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v52<<(uint(int32(12))%32) | v68<<(uint(int32(6))%32))
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
	v100 = v43<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v52<<(uint(int32(6))%32) | v68
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
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v105)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v219 = v204&int32(63) | (v162<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v171<<(uint(int32(12))%32) | v187<<(uint(int32(6))%32))
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
	v219 = v162<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v171<<(uint(int32(6))%32) | v187
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
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v224)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v338 = v323&int32(63) | (v281<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v290<<(uint(int32(12))%32) | v306<<(uint(int32(6))%32))
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
	v338 = v281<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v290<<(uint(int32(6))%32) | v306
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
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v343)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v459 = v444&int32(63) | (v402<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v411<<(uint(int32(12))%32) | v427<<(uint(int32(6))%32))
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
	v459 = v402<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v411<<(uint(int32(6))%32) | v427
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
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v464)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v577 = v562&int32(63) | (v520<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v529<<(uint(int32(12))%32) | v545<<(uint(int32(6))%32))
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
	v577 = v520<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v529<<(uint(int32(6))%32) | v545
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
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v582)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v700 = v685&int32(63) | (v643<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v652<<(uint(int32(12))%32) | v668<<(uint(int32(6))%32))
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
	v700 = v643<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v652<<(uint(int32(6))%32) | v668
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
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v705)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v818 = v803&int32(63) | (v761<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v770<<(uint(int32(12))%32) | v786<<(uint(int32(6))%32))
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
	v818 = v761<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v770<<(uint(int32(6))%32) | v786
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
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v823)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v937 = v922&int32(63) | (v880<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v889<<(uint(int32(12))%32) | v905<<(uint(int32(6))%32))
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
	v937 = v880<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v889<<(uint(int32(6))%32) | v905
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
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v942)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v1058 = v1043&int32(63) | (v1001<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1010<<(uint(int32(12))%32) | v1026<<(uint(int32(6))%32))
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
	v1058 = v1001<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1010<<(uint(int32(6))%32) | v1026
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
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1063)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v1244 = v1229&int32(63) | (v1187<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1196<<(uint(int32(12))%32) | v1212<<(uint(int32(6))%32))
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
	v1244 = v1187<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1196<<(uint(int32(6))%32) | v1212
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
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1249)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v1366 = v1351&int32(63) | (v1309<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1318<<(uint(int32(12))%32) | v1334<<(uint(int32(6))%32))
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
	v1366 = v1309<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1318<<(uint(int32(6))%32) | v1334
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
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1371)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v1491 = v1476&int32(63) | (v1434<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1443<<(uint(int32(12))%32) | v1459<<(uint(int32(6))%32))
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
	v1491 = v1434<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1443<<(uint(int32(6))%32) | v1459
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
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1496)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	v1613 = v1598&int32(63) | (v1556<<(uint(int32(18))%32)&int32(_a_F_basque_UTF_8_stem_0) | v1565<<(uint(int32(12))%32) | v1581<<(uint(int32(6))%32))
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
	v1613 = v1556<<(uint(int32(12))%32)&int32(_a_F_basque_UTF_8_stem_1) | v1565<<(uint(int32(6))%32) | v1581
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
	v1624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1618)>>(uint(int32(3))%32)))+uint32(_c_F_basque_UTF_8_stem[0]))))
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
	return v1898
L364:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1736 = v1730 - v1731 + v1735
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1736
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1736
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1736 <= v1739 {
		v1842 = v1736
		v1843 = v1735
		goto L393
	} else {
		goto L394
	}
L365:
	;
	v1730 = v1652
	v1731 = v1652
	goto L364
L366:
	;
	goto L367
L367:
	;
	v1659 = v1652
	v1660 = v1652
	v1662 = v1656
	goto L368
L368:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1663+v1662))))
	if base.B2i32(v1665&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1665)%32)&int32(70566434) == int32(0)) != 0 {
		v1730 = v1659
		v1731 = v1660
		goto L364
	} else {
		goto L370
	}
L369:
	;
	v1730 = v1722
	v1731 = v1724
	goto L364
L370:
	;
	v1679 = F_find_among_b(m, l0, int32(_a_F_basque_UTF_8_stem_2), int32(109))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	return int32(0)
L372:
	;
	if v1679 == int32(0) {
		v1730 = v1659
		v1731 = v1660
		goto L364
	} else {
		goto L373
	}
L373:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1685
	switch v1679 - int32(1) {
	case 0:
		goto L379
	case 1:
		goto L378
	case 2:
		goto L377
	case 3:
		goto L376
	case 4:
		goto L375
	default:
		goto L374
	}
L374:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1722
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1726 = v1722 - int32(1)
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1727 < v1726 {
		v1659 = v1722
		v1660 = v1724
		v1662 = v1726
		goto L368
	} else {
		goto L392
	}
L375:
	;
	v1717 = F_slice_from_s(m, l0, int32(6), int32(_a_F_basque_UTF_8_stem_3))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L371
	} else {
		goto L390
	}
L376:
	;
	v1711 = F_slice_from_s(m, l0, int32(7), int32(_a_F_basque_UTF_8_stem_4))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L371
	} else {
		goto L388
	}
L377:
	;
	v1705 = F_slice_from_s(m, l0, int32(7), int32(_a_F_basque_UTF_8_stem_5))
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L371
	} else {
		goto L386
	}
L378:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1696)))
	if v1685 < v1697 {
		v1730 = v1659
		v1731 = v1660
		goto L364
	} else {
		goto L383
	}
L379:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+8))
	if v1685 < v1690 {
		v1730 = v1659
		v1731 = v1660
		goto L364
	} else {
		goto L380
	}
L380:
	;
	v1692 = F_slice_del(m, l0)
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L371
	} else {
		goto L381
	}
L381:
	;
	if int32(0) <= v1692 {
		goto L374
	} else {
		goto L382
	}
L382:
	;
	v1898 = v1692
	goto L363
L383:
	;
	v1699 = F_slice_del(m, l0)
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L371
	} else {
		goto L384
	}
L384:
	;
	if int32(0) <= v1699 {
		goto L374
	} else {
		goto L385
	}
L385:
	;
	v1898 = v1699
	goto L363
L386:
	;
	if int32(0) <= v1705 {
		goto L374
	} else {
		goto L387
	}
L387:
	;
	v1898 = v1705
	goto L363
L388:
	;
	if int32(0) <= v1711 {
		goto L374
	} else {
		goto L389
	}
L389:
	;
	v1898 = v1711
	goto L363
L390:
	;
	if v1717 < int32(0) {
		v1898 = v1717
		goto L363
	} else {
		goto L391
	}
L391:
	;
	goto L374
L392:
	;
	goto L369
L393:
	;
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1848 = v1846 + (v1842 - v1843)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1848
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1848
	v1852 = v1848 - int32(1)
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1852 <= v1853 {
		goto L435
	} else {
		goto L436
	}
L394:
	;
	v1742 = v1736
	v1743 = v1735
	goto L395
L395:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1748 = int32(1)
	v1750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1746+v1742-v1748))))
	if base.B2i32(v1750&int32(224) != int32(96))|base.B2i32(v1748<<(uint(v1750)%32)&int32(71162402) == int32(0)) != 0 {
		v1842 = v1742
		v1843 = v1743
		goto L393
	} else {
		goto L397
	}
L396:
	;
	v1842 = v1836
	v1843 = v1838
	goto L393
L397:
	;
	v1764 = F_find_among_b(m, l0, int32(_a_F_basque_UTF_8_stem_6), int32(295))
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L371
	} else {
		goto L398
	}
L398:
	;
	if v1764 == int32(0) {
		v1842 = v1742
		v1843 = v1743
		goto L393
	} else {
		goto L399
	}
L399:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1768
	switch v1764 - int32(1) {
	case 0:
		goto L410
	case 1:
		goto L409
	case 2:
		goto L408
	case 3:
		goto L407
	case 4:
		goto L406
	case 5:
		goto L405
	case 6:
		goto L404
	case 7:
		goto L403
	case 8:
		goto L402
	case 9:
		goto L401
	default:
		goto L400
	}
L400:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1836
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1839 < v1836 {
		v1742 = v1836
		v1743 = v1838
		goto L395
	} else {
		goto L434
	}
L401:
	;
	v1831 = F_slice_from_s(m, l0, int32(5), int32(_a_F_basque_UTF_8_stem_7))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L371
	} else {
		goto L432
	}
L402:
	;
	v1825 = F_slice_from_s(m, l0, int32(5), int32(_a_F_basque_UTF_8_stem_8))
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L371
	} else {
		goto L430
	}
L403:
	;
	v1819 = F_slice_from_s(m, l0, int32(5), int32(_a_F_basque_UTF_8_stem_9))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L371
	} else {
		goto L428
	}
L404:
	;
	v1813 = F_slice_from_s(m, l0, int32(5), int32(_a_F_basque_UTF_8_stem_10))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L371
	} else {
		goto L426
	}
L405:
	;
	v1807 = F_slice_from_s(m, l0, int32(6), int32(_a_F_basque_UTF_8_stem_11))
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L371
	} else {
		goto L424
	}
L406:
	;
	v1801 = F_slice_from_s(m, l0, int32(3), int32(_a_F_basque_UTF_8_stem_12))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L371
	} else {
		goto L422
	}
L407:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v1792)+4))
	if v1768 < v1793 {
		v1842 = v1742
		v1843 = v1743
		goto L393
	} else {
		goto L419
	}
L408:
	;
	v1788 = F_slice_from_s(m, l0, int32(3), int32(_a_F_basque_UTF_8_stem_13))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L371
	} else {
		goto L417
	}
L409:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v1779)))
	if v1768 < v1780 {
		v1842 = v1742
		v1843 = v1743
		goto L393
	} else {
		goto L414
	}
L410:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1772)+8))
	if v1768 < v1773 {
		v1842 = v1742
		v1843 = v1743
		goto L393
	} else {
		goto L411
	}
L411:
	;
	v1775 = F_slice_del(m, l0)
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L371
	} else {
		goto L412
	}
L412:
	;
	if int32(0) <= v1775 {
		goto L400
	} else {
		goto L413
	}
L413:
	;
	v1898 = v1775
	goto L363
L414:
	;
	v1782 = F_slice_del(m, l0)
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L371
	} else {
		goto L415
	}
L415:
	;
	if int32(0) <= v1782 {
		goto L400
	} else {
		goto L416
	}
L416:
	;
	v1898 = v1782
	goto L363
L417:
	;
	if int32(0) <= v1788 {
		goto L400
	} else {
		goto L418
	}
L418:
	;
	v1898 = v1788
	goto L363
L419:
	;
	v1795 = F_slice_del(m, l0)
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L371
	} else {
		goto L420
	}
L420:
	;
	if int32(0) <= v1795 {
		goto L400
	} else {
		goto L421
	}
L421:
	;
	v1898 = v1795
	goto L363
L422:
	;
	if int32(0) <= v1801 {
		goto L400
	} else {
		goto L423
	}
L423:
	;
	v1898 = v1801
	goto L363
L424:
	;
	if int32(0) <= v1807 {
		goto L400
	} else {
		goto L425
	}
L425:
	;
	v1898 = v1807
	goto L363
L426:
	;
	if int32(0) <= v1813 {
		goto L400
	} else {
		goto L427
	}
L427:
	;
	v1898 = v1813
	goto L363
L428:
	;
	if int32(0) <= v1819 {
		goto L400
	} else {
		goto L429
	}
L429:
	;
	v1898 = v1819
	goto L363
L430:
	;
	if int32(0) <= v1825 {
		goto L400
	} else {
		goto L431
	}
L431:
	;
	v1898 = v1825
	goto L363
L432:
	;
	if v1831 < int32(0) {
		v1898 = v1831
		goto L363
	} else {
		goto L433
	}
L433:
	;
	goto L400
L434:
	;
	goto L396
L435:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1894
	v1898 = int32(1)
	goto L363
L436:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1855+v1852))))
	if base.B2i32(v1857&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1857)%32)&int32(_a_F_basque_UTF_8_stem_14) == int32(0)) != 0 {
		goto L435
	} else {
		goto L437
	}
L437:
	;
	v1871 = F_find_among_b(m, l0, int32(_a_F_basque_UTF_8_stem_15), int32(19))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L371
	} else {
		goto L438
	}
L438:
	;
	if v1871 == int32(0) {
		goto L435
	} else {
		goto L439
	}
L439:
	;
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1875
	switch v1871 - int32(1) {
	case 0:
		goto L441
	case 1:
		goto L440
	default:
		goto L435
	}
L440:
	;
	v1888 = F_slice_from_s(m, l0, int32(1), int32(_a_F_basque_UTF_8_stem_16))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L371
	} else {
		goto L445
	}
L441:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1879)+8))
	if v1875 < v1880 {
		goto L435
	} else {
		goto L442
	}
L442:
	;
	v1882 = F_slice_del(m, l0)
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L371
	} else {
		goto L443
	}
L443:
	;
	if int32(0) <= v1882 {
		goto L435
	} else {
		goto L444
	}
L444:
	;
	v1898 = v1882
	goto L363
L445:
	;
	if v1888 < int32(0) {
		v1898 = v1888
		goto L363
	} else {
		goto L446
	}
L446:
	;
	goto L435
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
	v23 = v21 & int32(_a_F_bernoulli_nextsampletuple_0)
	if base.Ui32(l2) < base.Ui32(v23) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v297)
	m.G0 = v8 + int32(16)
	return v297 & int32(_a_F_bernoulli_nextsampletuple_0)
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
	F_sigprocmask(m, int32(_a_F_bgworker_die_0), int32(0))
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
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_bgworker_die[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v18 + int32(96)
				F_errmsg(m, int32(_a_F_bgworker_die_1), v4)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_bgworker_die_2), int32(711), int32(_a_F_bgworker_die_3))
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
					F_errmsg(m, int32(_a_F_bitgetbit_0), v7)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bitgetbit_1), int32(1883), int32(_a_F_bitgetbit_2))
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
	var v90 int32
	_ = v90
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
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v32 < v35 {
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
	v90 = int32(0)
	goto L11
L11:
	;
	v109 = int32(0)
	goto L13
L12:
	;
	goto L4
L13:
	;
	v136 = int32(8) - v109
	v137 = v58 << (uint(v136) % 32)
	v138 = v30 + v61
	v139 = v90 + (v25 + v61)
	v141 = int32(base.Ui32(int32(255)) >> (uint(v109) % 32))
	v142 = int32(-256) >> (uint(v109) % 32)
	goto L15
L14:
	;
	if v90 != v47 {
		v90 = v90 + int32(1)
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
	return v109 + v90<<(uint(int32(3))%32) + int32(1)
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
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
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
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
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
		v19 = int32(0)
		if base.B2i32(v18 < v19)|base.B2i32(v17 <= v18) == v19 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if base.Ui32(int32(2)) <= base.Ui32(v25) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_bitsetbit_0), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bitsetbit_1), int32(1833), int32(_a_F_bitsetbit_2))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
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
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v31 = F_palloc(m, int32(base.Ui32(v28)>>(uint(int32(2))%32)))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v28 & int32(-4)
					v37 = int32(8)
					v38 = v31 + v37
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v43 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - v37
					if v43 != 0 {
						base.MemoryCopy(m, v38, v13+int32(8), v43)
					} else {
					}
					v49 = v38 + int32(base.Ui32(v18)>>(uint(int32(3))%32))
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
					v55 = (v18 ^ int32(-1)) & int32(7)
					if v25 != 0 {
						v61 = v50 | int32(1)<<(uint(v55)%32)
					} else {
						v61 = v50 & base.I32_rotl(int32(-2), v55)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v61)
					m.G0 = v10 + int32(16)
					return v31
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v17 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
					F_errmsg(m, int32(_a_F_bitsetbit_3), v10)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bitsetbit_1), int32(1825), int32(_a_F_bitsetbit_2))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
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
						F_errmsg(m, int32(_a_F_bitxor_0), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bitxor_1), int32(1342), int32(_a_F_bitxor_2))
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
	var v293 int32
	_ = v293
	var v309 int64
	_ = v309
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v330 int64
	_ = v330
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int64
	_ = v347
	var v357 int64
	_ = v357
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
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
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
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
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
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
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
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
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
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
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int64
	_ = v982
	var v984 int64
	_ = v984
	var v986 int64
	_ = v986
	var v988 int64
	_ = v988
	var v990 int64
	_ = v990
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1059 int64
	_ = v1059
	var v1060 int64
	_ = v1060
	var v1062 int64
	_ = v1062
	var v1063 int64
	_ = v1063
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
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
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1354 int32
	_ = v1354
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1370 int64
	_ = v1370
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1389 int32
	_ = v1389
	var v1399 int32
	_ = v1399
	var v1405 int32
	_ = v1405
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int64
	_ = v1420
	var v1422 int64
	_ = v1422
	var v1424 int64
	_ = v1424
	var v1426 int64
	_ = v1426
	var v1428 int64
	_ = v1428
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1452 int64
	_ = v1452
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1469 int32
	_ = v1469
	var v1484 int32
	_ = v1484
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int64
	_ = v1495
	var v1497 int64
	_ = v1497
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1550 int32
	_ = v1550
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
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
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v293 = base.B2i32(base.Ui32(v289) < base.Ui32(v290))
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
	if v293 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L60
	} else {
		goto L249
	}
L43:
	;
	goto L42
L44:
	;
	v1550 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1550
	v293 = v1550
	goto L41
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L60
	} else {
		goto L246
	}
L46:
	;
	v309 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v309 == int64(4294967296) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1037 = v1036 & (v280 ^ v272 - base.I32_rotl(v280, int32(24)))
	v1040 = v1035 + v1037*int32(40)
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040)+20)))
	if v1041 != 0 {
		goto L172
	} else {
		goto L173
	}
L49:
	;
	v312 = int32(0)
	v314 = int64(2)
	v316 = v309 << (uint(int64(1)) % 64)
	if base.Ui64(v316) <= base.Ui64(v314) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v293 = int32(1)
	goto L41
L51:
	;
	v319 = v314
	goto L53
L52:
	;
	v319 = v316
	goto L53
L53:
	;
	v320 = int64(1)
	if v319&(v319-v320) == int64(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v330 = v319
	goto L56
L55:
	;
	v330 = v320 << (uint(int64(64)-base.I64_clz(v319)) % 64)
	goto L56
L56:
	;
	if base.Ui64(v330*int64(40)) < base.Ui64(int64(2147483647)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v336 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v342 = F_MemoryContextAllocExtended(m, v337, base.I32_wrap_i64(v330)*int32(40), int32(5))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	goto L43
L60:
	;
	return int32(0)
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v342
	v347 = int64(1)
	if v330&(v330-v347) == int64(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v357 = v330
	goto L64
L63:
	;
	v357 = v347 << (uint(int64(64)-base.I64_clz(v330)) % 64)
	goto L64
L64:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v357*int64(40)) {
		goto L43
	} else {
		goto L65
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v357
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = base.I32_wrap_i64(v357) - int32(1)
	if v357 == int64(4294967296) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v374 = int32(-85899346)
	goto L68
L67:
	;
	v374 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v357), float64(0.9)))
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v374
	if v336 != int64(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v384 = v312
	goto L73
L70:
	;
	goto L71
L71:
	;
	F_pfree(m, v335)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L60
	} else {
		goto L170
	}
L72:
	;
	v678 = v312
	v679 = v672
	goto L118
L73:
	;
	v395 = v335 + v384*int32(40)
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+20)))
	if v396 != int32(1) {
		v672 = v384
		goto L72
	} else {
		goto L75
	}
L74:
	;
	v672 = int32(0)
	goto L72
L75:
	;
	v399 = int32(16)
	v405 = int32(-1636608416)
	if v395&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if (v659^v651-base.I32_rotl(v659, int32(24)))&v664 == v384 {
		v672 = v384
		goto L72
	} else {
		goto L116
	}
L77:
	;
	v637 = int32(14)
	v639 = v633 ^ v634 - base.I32_rotl(v633, v637)
	v643 = v639 ^ v632 - base.I32_rotl(v639, int32(11))
	v647 = v643 ^ v633 - base.I32_rotl(v643, int32(25))
	v651 = v647 ^ v639 - base.I32_rotl(v647, int32(16))
	v655 = v651 ^ v643 - base.I32_rotl(v651, int32(4))
	v659 = v655 ^ v647 - base.I32_rotl(v655, v637)
	goto L76
L78:
	;
	switch v559 - int32(1) {
	case 0:
		v625 = v550
		v626 = v551
		v627 = v555
		goto L105
	case 1:
		v618 = v550
		v619 = v551
		v620 = v555
		goto L106
	case 2:
		v611 = v550
		v612 = v551
		v613 = v555
		goto L107
	case 3:
		v605 = v551
		v606 = v555
		goto L108
	case 4:
		v601 = v551
		v602 = v555
		goto L109
	case 5:
		v595 = v551
		v596 = v555
		goto L110
	case 6:
		v589 = v551
		v590 = v555
		goto L111
	case 7:
		v584 = v555
		goto L112
	case 8:
		v579 = v555
		goto L113
	case 9:
		v574 = v555
		goto L114
	case 10:
		goto L115
	default:
		v632 = v550
		v633 = v551
		v634 = v555
		goto L77
	}
L79:
	;
	v514 = v395
	v515 = v399
	v516 = v405
	v517 = v405
	v518 = v405
	goto L102
L80:
	;
	goto L79
L81:
	;
	goto L82
L82:
	;
	goto L86
L84:
	;
	switch v457 - int32(1) {
	case 0:
		v511 = v448
		goto L91
	case 1:
		v506 = v448
		goto L92
	case 2:
		goto L93
	case 3:
		v499 = v449
		goto L94
	case 4:
		v496 = v449
		goto L95
	case 5:
		v491 = v449
		goto L96
	case 6:
		goto L97
	case 7:
		v482 = v453
		goto L98
	case 8:
		v477 = v453
		goto L99
	case 9:
		v472 = v453
		goto L100
	case 10:
		goto L101
	default:
		v632 = v448
		v633 = v449
		v634 = v453
		goto L77
	}
L86:
	;
	goto L87
L87:
	;
	v412 = v395
	v413 = v399
	v414 = v405
	v415 = v405
	v416 = v405
	goto L88
L88:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	v419 = v418 + v415
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v412)+8))
	v423 = v422 + v416
	v425 = int32(4)
	v427 = v420 + v414 - v423 ^ base.I32_rotl(v423, v425)
	v431 = v419 - v427 ^ base.I32_rotl(v427, int32(6))
	v432 = v423 + v419
	v433 = v427 + v432
	v434 = v431 + v433
	v438 = v432 - v431 ^ base.I32_rotl(v431, int32(8))
	v442 = v433 - v438 ^ base.I32_rotl(v438, int32(16))
	v446 = v434 - v442 ^ base.I32_rotl(v442, int32(19))
	v447 = v438 + v434
	v448 = v442 + v447
	v449 = v446 + v448
	v453 = v447 - v446 ^ base.I32_rotl(v446, v425)
	v454 = int32(12)
	v455 = v412 + v454
	v457 = v413 - v454
	if base.Ui32(int32(11)) < base.Ui32(v457) {
		v412 = v455
		v413 = v457
		v414 = v448
		v415 = v449
		v416 = v453
		goto L88
	} else {
		goto L90
	}
L89:
	;
	goto L84
L90:
	;
	goto L89
L91:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455))))
	v632 = v511 + v512
	v633 = v449
	v634 = v453
	goto L77
L92:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+1)))
	v511 = v507<<(uint(int32(8))%32) + v506
	goto L91
L93:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+2)))
	v506 = v502<<(uint(int32(16))%32) + v448
	goto L92
L94:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v632 = v500 + v448
	v633 = v499
	v634 = v453
	goto L77
L95:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+4)))
	v499 = v496 + v497
	goto L94
L96:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+5)))
	v496 = v492<<(uint(int32(8))%32) + v491
	goto L95
L97:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+6)))
	v491 = v487<<(uint(int32(16))%32) + v449
	goto L96
L98:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	v632 = v483 + v448
	v633 = v485 + v449
	v634 = v482
	goto L77
L99:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+8)))
	v482 = v478<<(uint(int32(8))%32) + v477
	goto L98
L100:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+9)))
	v477 = v473<<(uint(int32(16))%32) + v472
	goto L99
L101:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+10)))
	v472 = v468<<(uint(int32(24))%32) + v453
	goto L100
L102:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	v521 = v520 + v517
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v514)))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	v525 = v524 + v518
	v527 = int32(4)
	v529 = v522 + v516 - v525 ^ base.I32_rotl(v525, v527)
	v533 = v521 - v529 ^ base.I32_rotl(v529, int32(6))
	v534 = v525 + v521
	v535 = v529 + v534
	v536 = v533 + v535
	v540 = v534 - v533 ^ base.I32_rotl(v533, int32(8))
	v544 = v535 - v540 ^ base.I32_rotl(v540, int32(16))
	v548 = v536 - v544 ^ base.I32_rotl(v544, int32(19))
	v549 = v540 + v536
	v550 = v544 + v549
	v551 = v548 + v550
	v555 = v549 - v548 ^ base.I32_rotl(v548, v527)
	v556 = int32(12)
	v557 = v514 + v556
	v559 = v515 - v556
	if base.Ui32(int32(11)) < base.Ui32(v559) {
		v514 = v557
		v515 = v559
		v516 = v550
		v517 = v551
		v518 = v555
		goto L102
	} else {
		goto L104
	}
L103:
	;
	goto L78
L104:
	;
	goto L103
L105:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v632 = v625 + v628
	v633 = v626
	v634 = v627
	goto L77
L106:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+1)))
	v625 = v621<<(uint(int32(8))%32) + v618
	v626 = v619
	v627 = v620
	goto L105
L107:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+2)))
	v618 = v614<<(uint(int32(16))%32) + v611
	v619 = v612
	v620 = v613
	goto L106
L108:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+3)))
	v611 = v607<<(uint(int32(24))%32) + v550
	v612 = v605
	v613 = v606
	goto L107
L109:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+4)))
	v605 = v601 + v603
	v606 = v602
	goto L108
L110:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+5)))
	v601 = v597<<(uint(int32(8))%32) + v595
	v602 = v596
	goto L109
L111:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+6)))
	v595 = v591<<(uint(int32(16))%32) + v589
	v596 = v590
	goto L110
L112:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+7)))
	v589 = v585<<(uint(int32(24))%32) + v551
	v590 = v584
	goto L111
L113:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+8)))
	v584 = v580<<(uint(int32(8))%32) + v579
	goto L112
L114:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+9)))
	v579 = v575<<(uint(int32(16))%32) + v574
	goto L113
L115:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+10)))
	v574 = v570<<(uint(int32(24))%32) + v555
	goto L114
L116:
	;
	v668 = v384 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v668)) < base.Ui64(v336) {
		v384 = v668
		goto L73
	} else {
		goto L117
	}
L117:
	;
	goto L74
L118:
	;
	v690 = v335 + v679*int32(40)
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690)+20)))
	if v691 == int32(1) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	goto L71
L120:
	;
	v694 = int32(16)
	v700 = int32(-1636608416)
	if v690&int32(3) != 0 {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	goto L122
L122:
	;
	v1008 = v679 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1008)) < base.Ui64(v336) {
		goto L166
	} else {
		goto L167
	}
L123:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v973 = v954 ^ v946 - base.I32_rotl(v954, int32(24))
	goto L163
L124:
	;
	v932 = int32(14)
	v934 = v928 ^ v929 - base.I32_rotl(v928, v932)
	v938 = v934 ^ v927 - base.I32_rotl(v934, int32(11))
	v942 = v938 ^ v928 - base.I32_rotl(v938, int32(25))
	v946 = v942 ^ v934 - base.I32_rotl(v942, int32(16))
	v950 = v946 ^ v938 - base.I32_rotl(v946, int32(4))
	v954 = v950 ^ v942 - base.I32_rotl(v950, v932)
	goto L123
L125:
	;
	switch v854 - int32(1) {
	case 0:
		v920 = v845
		v921 = v846
		v922 = v850
		goto L152
	case 1:
		v913 = v845
		v914 = v846
		v915 = v850
		goto L153
	case 2:
		v906 = v845
		v907 = v846
		v908 = v850
		goto L154
	case 3:
		v900 = v846
		v901 = v850
		goto L155
	case 4:
		v896 = v846
		v897 = v850
		goto L156
	case 5:
		v890 = v846
		v891 = v850
		goto L157
	case 6:
		v884 = v846
		v885 = v850
		goto L158
	case 7:
		v879 = v850
		goto L159
	case 8:
		v874 = v850
		goto L160
	case 9:
		v869 = v850
		goto L161
	case 10:
		goto L162
	default:
		v927 = v845
		v928 = v846
		v929 = v850
		goto L124
	}
L126:
	;
	v809 = v690
	v810 = v694
	v811 = v700
	v812 = v700
	v813 = v700
	goto L149
L127:
	;
	goto L126
L128:
	;
	goto L129
L129:
	;
	goto L133
L131:
	;
	switch v752 - int32(1) {
	case 0:
		v806 = v743
		goto L138
	case 1:
		v801 = v743
		goto L139
	case 2:
		goto L140
	case 3:
		v794 = v744
		goto L141
	case 4:
		v791 = v744
		goto L142
	case 5:
		v786 = v744
		goto L143
	case 6:
		goto L144
	case 7:
		v777 = v748
		goto L145
	case 8:
		v772 = v748
		goto L146
	case 9:
		v767 = v748
		goto L147
	case 10:
		goto L148
	default:
		v927 = v743
		v928 = v744
		v929 = v748
		goto L124
	}
L133:
	;
	goto L134
L134:
	;
	v707 = v690
	v708 = v694
	v709 = v700
	v710 = v700
	v711 = v700
	goto L135
L135:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	v714 = v713 + v710
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v707)+8))
	v718 = v717 + v711
	v720 = int32(4)
	v722 = v715 + v709 - v718 ^ base.I32_rotl(v718, v720)
	v726 = v714 - v722 ^ base.I32_rotl(v722, int32(6))
	v727 = v718 + v714
	v728 = v722 + v727
	v729 = v726 + v728
	v733 = v727 - v726 ^ base.I32_rotl(v726, int32(8))
	v737 = v728 - v733 ^ base.I32_rotl(v733, int32(16))
	v741 = v729 - v737 ^ base.I32_rotl(v737, int32(19))
	v742 = v733 + v729
	v743 = v737 + v742
	v744 = v741 + v743
	v748 = v742 - v741 ^ base.I32_rotl(v741, v720)
	v749 = int32(12)
	v750 = v707 + v749
	v752 = v708 - v749
	if base.Ui32(int32(11)) < base.Ui32(v752) {
		v707 = v750
		v708 = v752
		v709 = v743
		v710 = v744
		v711 = v748
		goto L135
	} else {
		goto L137
	}
L136:
	;
	goto L131
L137:
	;
	goto L136
L138:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	v927 = v806 + v807
	v928 = v744
	v929 = v748
	goto L124
L139:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+1)))
	v806 = v802<<(uint(int32(8))%32) + v801
	goto L138
L140:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+2)))
	v801 = v797<<(uint(int32(16))%32) + v743
	goto L139
L141:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v927 = v795 + v743
	v928 = v794
	v929 = v748
	goto L124
L142:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+4)))
	v794 = v791 + v792
	goto L141
L143:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+5)))
	v791 = v787<<(uint(int32(8))%32) + v786
	goto L142
L144:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+6)))
	v786 = v782<<(uint(int32(16))%32) + v744
	goto L143
L145:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v750)))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v750)+4))
	v927 = v778 + v743
	v928 = v780 + v744
	v929 = v777
	goto L124
L146:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+8)))
	v777 = v773<<(uint(int32(8))%32) + v772
	goto L145
L147:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+9)))
	v772 = v768<<(uint(int32(16))%32) + v767
	goto L146
L148:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+10)))
	v767 = v763<<(uint(int32(24))%32) + v748
	goto L147
L149:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v809)+4))
	v816 = v815 + v812
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v809)))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v809)+8))
	v820 = v819 + v813
	v822 = int32(4)
	v824 = v817 + v811 - v820 ^ base.I32_rotl(v820, v822)
	v828 = v816 - v824 ^ base.I32_rotl(v824, int32(6))
	v829 = v820 + v816
	v830 = v824 + v829
	v831 = v828 + v830
	v835 = v829 - v828 ^ base.I32_rotl(v828, int32(8))
	v839 = v830 - v835 ^ base.I32_rotl(v835, int32(16))
	v843 = v831 - v839 ^ base.I32_rotl(v839, int32(19))
	v844 = v835 + v831
	v845 = v839 + v844
	v846 = v843 + v845
	v850 = v844 - v843 ^ base.I32_rotl(v843, v822)
	v851 = int32(12)
	v852 = v809 + v851
	v854 = v810 - v851
	if base.Ui32(int32(11)) < base.Ui32(v854) {
		v809 = v852
		v810 = v854
		v811 = v845
		v812 = v846
		v813 = v850
		goto L149
	} else {
		goto L151
	}
L150:
	;
	goto L125
L151:
	;
	goto L150
L152:
	;
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852))))
	v927 = v920 + v923
	v928 = v921
	v929 = v922
	goto L124
L153:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+1)))
	v920 = v916<<(uint(int32(8))%32) + v913
	v921 = v914
	v922 = v915
	goto L152
L154:
	;
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+2)))
	v913 = v909<<(uint(int32(16))%32) + v906
	v914 = v907
	v915 = v908
	goto L153
L155:
	;
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+3)))
	v906 = v902<<(uint(int32(24))%32) + v845
	v907 = v900
	v908 = v901
	goto L154
L156:
	;
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+4)))
	v900 = v896 + v898
	v901 = v897
	goto L155
L157:
	;
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+5)))
	v896 = v892<<(uint(int32(8))%32) + v890
	v897 = v891
	goto L156
L158:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+6)))
	v890 = v886<<(uint(int32(16))%32) + v884
	v891 = v885
	goto L157
L159:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+7)))
	v884 = v880<<(uint(int32(24))%32) + v846
	v885 = v879
	goto L158
L160:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+8)))
	v879 = v875<<(uint(int32(8))%32) + v874
	goto L159
L161:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+9)))
	v874 = v870<<(uint(int32(16))%32) + v869
	goto L160
L162:
	;
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v852)+10)))
	v869 = v865<<(uint(int32(24))%32) + v850
	goto L161
L163:
	;
	v975 = v959 & v973
	v980 = v342 + v975*int32(40)
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980)+20)))
	if v981 != 0 {
		v973 = v975 + int32(1)
		goto L163
	} else {
		goto L165
	}
L164:
	;
	v982 = *(*int64)(unsafe.Add(mBase, uint32(v690)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+32)) = v982
	v984 = *(*int64)(unsafe.Add(mBase, uint32(v690)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+24)) = v984
	v986 = *(*int64)(unsafe.Add(mBase, uint32(v690)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+16)) = v986
	v988 = *(*int64)(unsafe.Add(mBase, uint32(v690)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v980)+8)) = v988
	v990 = *(*int64)(unsafe.Add(mBase, uint32(v690)))
	*(*int64)(unsafe.Add(mBase, uint32(v980))) = v990
	goto L122
L165:
	;
	goto L164
L166:
	;
	v1012 = v1008
	goto L168
L167:
	;
	v1012 = int32(0)
	goto L168
L168:
	;
	v1014 = v678 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1014)) < base.Ui64(v336) {
		v678 = v1014
		v679 = v1012
		goto L118
	} else {
		goto L169
	}
L169:
	;
	goto L119
L170:
	;
	goto L50
L171:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v1512)
	m.G0 = v18 + int32(16)
	return v1510
L172:
	;
	v1051 = int32(0)
	v1052 = v1040
	v1055 = v1037
	goto L176
L173:
	;
	v1484 = v1040
	goto L174
L174:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1492 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1491 + v1492
	v1495 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1484)+8)) = v1495
	v1497 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v1484))) = v1497
	*(*uint8)(unsafe.Add(mBase, uint32(v1484)+20)) = uint8(v1492)
	v1510 = v1484
	v1512 = int32(0)
	goto L171
L175:
	;
	v1484 = v1469
	goto L174
L176:
	;
	v1059 = *(*int64)(unsafe.Add(mBase, uint32(v1052)))
	v1060 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v1062 = *(*int64)(unsafe.Add(mBase, uint32(v1052)+8))
	v1063 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if v1059^v1060|(v1062^v1063) == int64(0) {
		v1510 = v1052
		v1512 = int32(1)
		goto L171
	} else {
		goto L178
	}
L177:
	;
	v1469 = v1459
	goto L175
L178:
	;
	v1068 = int32(16)
	v1074 = int32(-1636608416)
	if v1052&int32(3) != 0 {
		goto L183
	} else {
		goto L184
	}
L179:
	;
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1334 = (v1328 ^ v1320 - base.I32_rotl(v1328, int32(24))) & v1333
	if base.Ui32(v1055) < base.Ui32(v1334) {
		goto L219
	} else {
		goto L220
	}
L180:
	;
	v1306 = int32(14)
	v1308 = v1302 ^ v1303 - base.I32_rotl(v1302, v1306)
	v1312 = v1308 ^ v1301 - base.I32_rotl(v1308, int32(11))
	v1316 = v1312 ^ v1302 - base.I32_rotl(v1312, int32(25))
	v1320 = v1316 ^ v1308 - base.I32_rotl(v1316, int32(16))
	v1324 = v1320 ^ v1312 - base.I32_rotl(v1320, int32(4))
	v1328 = v1324 ^ v1316 - base.I32_rotl(v1324, v1306)
	goto L179
L181:
	;
	switch v1228 - int32(1) {
	case 0:
		v1294 = v1219
		v1295 = v1220
		v1296 = v1224
		goto L208
	case 1:
		v1287 = v1219
		v1288 = v1220
		v1289 = v1224
		goto L209
	case 2:
		v1280 = v1219
		v1281 = v1220
		v1282 = v1224
		goto L210
	case 3:
		v1274 = v1220
		v1275 = v1224
		goto L211
	case 4:
		v1270 = v1220
		v1271 = v1224
		goto L212
	case 5:
		v1264 = v1220
		v1265 = v1224
		goto L213
	case 6:
		v1258 = v1220
		v1259 = v1224
		goto L214
	case 7:
		v1253 = v1224
		goto L215
	case 8:
		v1248 = v1224
		goto L216
	case 9:
		v1243 = v1224
		goto L217
	case 10:
		goto L218
	default:
		v1301 = v1219
		v1302 = v1220
		v1303 = v1224
		goto L180
	}
L182:
	;
	v1183 = v1052
	v1184 = v1068
	v1185 = v1074
	v1186 = v1074
	v1187 = v1074
	goto L205
L183:
	;
	goto L182
L184:
	;
	goto L185
L185:
	;
	goto L189
L187:
	;
	switch v1126 - int32(1) {
	case 0:
		v1180 = v1117
		goto L194
	case 1:
		v1175 = v1117
		goto L195
	case 2:
		goto L196
	case 3:
		v1168 = v1118
		goto L197
	case 4:
		v1165 = v1118
		goto L198
	case 5:
		v1160 = v1118
		goto L199
	case 6:
		goto L200
	case 7:
		v1151 = v1122
		goto L201
	case 8:
		v1146 = v1122
		goto L202
	case 9:
		v1141 = v1122
		goto L203
	case 10:
		goto L204
	default:
		v1301 = v1117
		v1302 = v1118
		v1303 = v1122
		goto L180
	}
L189:
	;
	goto L190
L190:
	;
	v1081 = v1052
	v1082 = v1068
	v1083 = v1074
	v1084 = v1074
	v1085 = v1074
	goto L191
L191:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	v1088 = v1087 + v1084
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1081)))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+8))
	v1092 = v1091 + v1085
	v1094 = int32(4)
	v1096 = v1089 + v1083 - v1092 ^ base.I32_rotl(v1092, v1094)
	v1100 = v1088 - v1096 ^ base.I32_rotl(v1096, int32(6))
	v1101 = v1092 + v1088
	v1102 = v1096 + v1101
	v1103 = v1100 + v1102
	v1107 = v1101 - v1100 ^ base.I32_rotl(v1100, int32(8))
	v1111 = v1102 - v1107 ^ base.I32_rotl(v1107, int32(16))
	v1115 = v1103 - v1111 ^ base.I32_rotl(v1111, int32(19))
	v1116 = v1107 + v1103
	v1117 = v1111 + v1116
	v1118 = v1115 + v1117
	v1122 = v1116 - v1115 ^ base.I32_rotl(v1115, v1094)
	v1123 = int32(12)
	v1124 = v1081 + v1123
	v1126 = v1082 - v1123
	if base.Ui32(int32(11)) < base.Ui32(v1126) {
		v1081 = v1124
		v1082 = v1126
		v1083 = v1117
		v1084 = v1118
		v1085 = v1122
		goto L191
	} else {
		goto L193
	}
L192:
	;
	goto L187
L193:
	;
	goto L192
L194:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	v1301 = v1180 + v1181
	v1302 = v1118
	v1303 = v1122
	goto L180
L195:
	;
	v1176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+1)))
	v1180 = v1176<<(uint(int32(8))%32) + v1175
	goto L194
L196:
	;
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+2)))
	v1175 = v1171<<(uint(int32(16))%32) + v1117
	goto L195
L197:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1124)))
	v1301 = v1169 + v1117
	v1302 = v1168
	v1303 = v1122
	goto L180
L198:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+4)))
	v1168 = v1165 + v1166
	goto L197
L199:
	;
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+5)))
	v1165 = v1161<<(uint(int32(8))%32) + v1160
	goto L198
L200:
	;
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+6)))
	v1160 = v1156<<(uint(int32(16))%32) + v1118
	goto L199
L201:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1124)))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+4))
	v1301 = v1152 + v1117
	v1302 = v1154 + v1118
	v1303 = v1151
	goto L180
L202:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+8)))
	v1151 = v1147<<(uint(int32(8))%32) + v1146
	goto L201
L203:
	;
	v1142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+9)))
	v1146 = v1142<<(uint(int32(16))%32) + v1141
	goto L202
L204:
	;
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124)+10)))
	v1141 = v1137<<(uint(int32(24))%32) + v1122
	goto L203
L205:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1183)+4))
	v1190 = v1189 + v1186
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1183)))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1183)+8))
	v1194 = v1193 + v1187
	v1196 = int32(4)
	v1198 = v1191 + v1185 - v1194 ^ base.I32_rotl(v1194, v1196)
	v1202 = v1190 - v1198 ^ base.I32_rotl(v1198, int32(6))
	v1203 = v1194 + v1190
	v1204 = v1198 + v1203
	v1205 = v1202 + v1204
	v1209 = v1203 - v1202 ^ base.I32_rotl(v1202, int32(8))
	v1213 = v1204 - v1209 ^ base.I32_rotl(v1209, int32(16))
	v1217 = v1205 - v1213 ^ base.I32_rotl(v1213, int32(19))
	v1218 = v1209 + v1205
	v1219 = v1213 + v1218
	v1220 = v1217 + v1219
	v1224 = v1218 - v1217 ^ base.I32_rotl(v1217, v1196)
	v1225 = int32(12)
	v1226 = v1183 + v1225
	v1228 = v1184 - v1225
	if base.Ui32(int32(11)) < base.Ui32(v1228) {
		v1183 = v1226
		v1184 = v1228
		v1185 = v1219
		v1186 = v1220
		v1187 = v1224
		goto L205
	} else {
		goto L207
	}
L206:
	;
	goto L181
L207:
	;
	goto L206
L208:
	;
	v1297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226))))
	v1301 = v1294 + v1297
	v1302 = v1295
	v1303 = v1296
	goto L180
L209:
	;
	v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+1)))
	v1294 = v1290<<(uint(int32(8))%32) + v1287
	v1295 = v1288
	v1296 = v1289
	goto L208
L210:
	;
	v1283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+2)))
	v1287 = v1283<<(uint(int32(16))%32) + v1280
	v1288 = v1281
	v1289 = v1282
	goto L209
L211:
	;
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+3)))
	v1280 = v1276<<(uint(int32(24))%32) + v1219
	v1281 = v1274
	v1282 = v1275
	goto L210
L212:
	;
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+4)))
	v1274 = v1270 + v1272
	v1275 = v1271
	goto L211
L213:
	;
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+5)))
	v1270 = v1266<<(uint(int32(8))%32) + v1264
	v1271 = v1265
	goto L212
L214:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+6)))
	v1264 = v1260<<(uint(int32(16))%32) + v1258
	v1265 = v1259
	goto L213
L215:
	;
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+7)))
	v1258 = v1254<<(uint(int32(24))%32) + v1220
	v1259 = v1253
	goto L214
L216:
	;
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+8)))
	v1253 = v1249<<(uint(int32(8))%32) + v1248
	goto L215
L217:
	;
	v1244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+9)))
	v1248 = v1244<<(uint(int32(16))%32) + v1243
	goto L216
L218:
	;
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226)+10)))
	v1243 = v1239<<(uint(int32(24))%32) + v1224
	goto L217
L219:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1338 = v1055 + v1336
	goto L221
L220:
	;
	v1338 = v1055
	goto L221
L221:
	;
	v1341 = (v1055 + int32(1)) & v1333
	if base.Ui32(v1338-v1334) < base.Ui32(v1051) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1346 = v1035 + v1341*int32(40)
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1346)+20)))
	if v1347 != 0 {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	v1447 = v1051 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1447) {
		goto L241
	} else {
		goto L242
	}
L225:
	;
	v1350 = v1341
	v1354 = int32(0)
	goto L228
L226:
	;
	v1383 = v1341
	v1389 = v1346
	goto L227
L227:
	;
	if v1383 != v1055 {
		goto L235
	} else {
		goto L236
	}
L228:
	;
	v1365 = v1354 + int32(1)
	if int32(151) <= v1365 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v1383 = v1377
	v1389 = v1380
	goto L227
L230:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1370 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1368), base.F64_convert_i64_u(v1370)), float64(0.1)) != 0 {
		goto L44
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v1377 = (v1350 + int32(1)) & v1333
	v1380 = v1035 + v1377*int32(40)
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1380)+20)))
	if v1381 != 0 {
		v1350 = v1377
		v1354 = v1365
		goto L228
	} else {
		goto L234
	}
L233:
	;
	goto L232
L234:
	;
	goto L229
L235:
	;
	v1399 = v1383
	v1405 = v1389
	goto L238
L236:
	;
	goto L237
L237:
	;
	v1469 = v1052
	goto L175
L238:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1416 = v1413 & (v1399 - int32(1))
	v1419 = v1035 + v1416*int32(40)
	v1420 = *(*int64)(unsafe.Add(mBase, uint32(v1419)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1405)+32)) = v1420
	v1422 = *(*int64)(unsafe.Add(mBase, uint32(v1419)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1405)+24)) = v1422
	v1424 = *(*int64)(unsafe.Add(mBase, uint32(v1419)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1405)+16)) = v1424
	v1426 = *(*int64)(unsafe.Add(mBase, uint32(v1419)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1405)+8)) = v1426
	v1428 = *(*int64)(unsafe.Add(mBase, uint32(v1419)))
	*(*int64)(unsafe.Add(mBase, uint32(v1405))) = v1428
	if v1416 != v1055 {
		v1399 = v1416
		v1405 = v1419
		goto L238
	} else {
		goto L240
	}
L239:
	;
	goto L237
L240:
	;
	goto L239
L241:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1452 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1450), base.F64_convert_i64_u(v1452)), float64(0.1)) != 0 {
		goto L44
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1459 = v1035 + v1341*int32(40)
	v1460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1459)+20)))
	if v1460 != 0 {
		v1051 = v1447
		v1052 = v1459
		v1055 = v1341
		goto L176
	} else {
		goto L245
	}
L244:
	;
	goto L243
L245:
	;
	goto L177
L246:
	;
	F_errmsg_internal(m, int32(_a_F_blockreftable_insert_0), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L60
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_blockreftable_insert_1), int32(630), int32(_a_F_blockreftable_insert_2))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L60
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_errmsg_internal(m, int32(_a_F_blockreftable_insert_3), int32(0))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L60
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_blockreftable_insert_1), int32(327), int32(_a_F_blockreftable_insert_4))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L60
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bloptions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_bloptions[0]))
	v8 = F_build_reloptions(m, l0, l1, v4, int32(136), int32(_a_F_bloptions_0), int32(33))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v16 = base.I32_div_s(v12+int32(15), int32(16))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v16
		} else {
		}
		return v8
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
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
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
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v453 int32
	_ = v453
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
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
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	if int32(0) < v201 {
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
	v175 = m.ExcPending
	if v175 != 0 {
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
		v188 = v40
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
	v188 = v165
	goto L1
L16:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	if v106 != v29 {
		v165 = v105
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
	F_errmsg(m, int32(_a_F_blvalidate_0), v19+int32(112))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(84), int32(_a_F_blvalidate_2))
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
	v169 = v56 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v45)+40))
	if v169 < v170 {
		v56 = v169
		v57 = v165
		goto L14
	} else {
		goto L43
	}
L25:
	;
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)))
	switch v108 - int32(1) {
	case 0:
		goto L30
	case 1:
		goto L28
	default:
		goto L29
	}
L26:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L39
	}
L27:
	;
	v135 = int32(0)
	v138 = F_errstart(m, int32(17), v135)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L37
	}
L28:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v133 = F_check_amoptsproc_signature(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L35
	}
L29:
	;
	v123 = int32(0)
	v126 = F_errstart(m, int32(17), v123)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L33
	}
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v50
	v115 = int32(1)
	v119 = F_check_amproc_signature(m, v111, int32(23), int32(0), v115, v115, v19+int32(96))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v119 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v165 = v105
	goto L24
L33:
	;
	if v126 == int32(0) {
		v165 = v123
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v144 = int32(_a_F_blvalidate_3)
	v145 = int32(111)
	goto L26
L35:
	;
	if v133 != 0 {
		v165 = v105
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L27
L37:
	;
	if v138 == int32(0) {
		v165 = v135
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v144 = int32(_a_F_blvalidate_4)
	v145 = int32(123)
	goto L26
L39:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v150 = F_format_procedure(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v76)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v32
	F_errmsg(m, v144, v19+int32(80))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), v145, int32(_a_F_blvalidate_2))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v165 = int32(0)
	goto L24
L43:
	;
	goto L15
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	F_errmsg_internal(m, int32(_a_F_blvalidate_5), v19)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(50), int32(_a_F_blvalidate_2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
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
	v209 = int32(0)
	v210 = v188
	goto L50
L48:
	;
	v333 = v188
	goto L49
L49:
	;
	v346 = F_identify_opfamily_groups(m, v38, v45)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L2
	} else {
		goto L83
	}
L50:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v38+int32(48)+v209<<(uint(int32(2))%32))))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+56))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+22)))
	v229 = v227 + v228
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229)+16)))
	if v230 == int32(1) {
		v260 = v210
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v333 = v325
	goto L49
L52:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+18)))
	if v262 == int32(115) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v233 = int32(0)
	v236 = F_errstart(m, int32(17), v233)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	if v236 == int32(0) {
		v260 = v233
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	v244 = F_format_operator(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v229)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v32
	F_errmsg(m, int32(_a_F_blvalidate_6), v19-int32(-64))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(143), int32(_a_F_blvalidate_2))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v260 = v233
	goto L52
L60:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v229)+8))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v298 = F_check_amop_signature(m, v294, int32(16), v296, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L72
	}
L61:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v229)+28))
	if v265 == int32(0) {
		v293 = v260
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v268 = int32(0)
	v271 = F_errstart(m, int32(17), v268)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if v271 == int32(0) {
		v293 = v268
		goto L60
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	v279 = F_format_operator(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v32
	F_errmsg(m, int32(_a_F_blvalidate_7), v19+int32(48))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(155), int32(_a_F_blvalidate_2))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v293 = v268
	goto L60
L71:
	;
	v327 = v209 + int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	if v327 < v328 {
		v209 = v327
		v210 = v325
		goto L50
	} else {
		goto L80
	}
L72:
	;
	if v298 != 0 {
		v325 = v293
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v300 = int32(0)
	v303 = F_errstart(m, int32(17), v300)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	if v303 == int32(0) {
		v325 = v300
		goto L71
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	v311 = F_format_operator(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v32
	F_errmsg(m, int32(_a_F_blvalidate_8), v19+int32(32))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(168), int32(_a_F_blvalidate_2))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v325 = v300
	goto L71
L80:
	;
	goto L51
L81:
	;
	F_ReleaseCatCacheList(m, v45)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L2
	} else {
		goto L122
	}
L82:
	;
	v472 = int32(0)
	v475 = F_errstart(m, int32(17), v472)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L2
	} else {
		goto L117
	}
L83:
	;
	if v346 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v350 <= int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v437 == int32(0) {
		goto L82
	} else {
		goto L115
	}
L86:
	;
	v437 = int32(0)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v354 = int32(0)
	if v350 != int32(1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v358 = int32(0)
	if v358 < v350 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v409 = v354
	v411 = v354
	goto L91
L91:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v425+v409<<(uint(int32(2))%32))))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if v430 != v29 {
		v437 = v411
		goto L85
	} else {
		goto L111
	}
L92:
	;
	v361 = v350
	goto L94
L93:
	;
	v361 = v358
	goto L94
L94:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	v367 = int32(0)
	v369 = v367
	v371 = v354
	v373 = v367
	goto L95
L95:
	;
	v387 = v366 + v369<<(uint(int32(2))%32)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	if v29 == v389 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v361&int32(1) == int32(0) {
		v437 = v401
		goto L85
	} else {
		goto L110
	}
L97:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	if v391 == v29 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v394 = v371
	goto L99
L99:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	if v29 == v396 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v393 = v388
	goto L102
L101:
	;
	v393 = v371
	goto L102
L102:
	;
	v394 = v393
	goto L99
L103:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v398 == v29 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v401 = v394
	goto L105
L105:
	;
	v402 = int32(2)
	v403 = v369 + v402
	v405 = v373 + v402
	if v405 != v361&int32(2147483646) {
		v369 = v403
		v371 = v401
		v373 = v405
		goto L95
	} else {
		goto L109
	}
L106:
	;
	v400 = v395
	goto L108
L107:
	;
	v400 = v394
	goto L108
L108:
	;
	v401 = v400
	goto L105
L109:
	;
	goto L96
L110:
	;
	v409 = v403
	v411 = v401
	goto L91
L111:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	if v432 == v29 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v434 = v429
	goto L114
L113:
	;
	v434 = v411
	goto L114
L114:
	;
	v437 = v434
	goto L85
L115:
	;
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+16)))
	if v453&int32(2) != 0 {
		v500 = v333
		goto L81
	} else {
		goto L116
	}
L116:
	;
	goto L82
L117:
	;
	if v475 == int32(0) {
		v500 = v472
		goto L81
	} else {
		goto L118
	}
L118:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v28 + int32(8)
	F_errmsg(m, int32(_a_F_blvalidate_9), v19+int32(16))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_blvalidate_1), int32(206), int32(_a_F_blvalidate_2))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v500 = v472
	goto L81
L122:
	;
	F_ReleaseCatCacheList(m, v38)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	F_ReleaseCatCache(m, v22)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	m.G0 = v19 + int32(128)
	return v500 & int32(1)
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
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v7 == int32(1) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
		if v6&int32(1) != 0 {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
			if v10 == v13 {
				v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
				return base.B2i32(v78 != v79)
			} else {
				if v10&int32(1) != 0 {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
					if v42 == int32(0) {
						return int32(0)
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
						v48 = int32(1)
						v49 = v47 ^ v48
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v49)
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
						v53 = v51 ^ v48
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v53)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v48)
						v57 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v57)
						v61 = F_make_range(m, l0, l1, l2, v57, v57)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
							v67 = int32(1)
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(base.Ui32(v63)>>(uint(int32(2))%32))-v67))))
							return v69 & v67
						}
					}
				} else {
					return int32(0)
				}
			}
		} else {
			if v10&int32(1) != 0 {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
				if v42 == int32(0) {
					return int32(0)
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
					v48 = int32(1)
					v49 = v47 ^ v48
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v49)
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
					v53 = v51 ^ v48
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v53)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v48)
					v57 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v57)
					v61 = F_make_range(m, l0, l1, l2, v57, v57)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
						v67 = int32(1)
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(base.Ui32(v63)>>(uint(int32(2))%32))-v67))))
						return v69 & v67
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
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
				if v42 == int32(0) {
					return int32(0)
				} else {
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
					v48 = int32(1)
					v49 = v47 ^ v48
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v49)
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
					v53 = v51 ^ v48
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v53)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v48)
					v57 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v57)
					v61 = F_make_range(m, l0, l1, l2, v57, v57)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
						v67 = int32(1)
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(base.Ui32(v63)>>(uint(int32(2))%32))-v67))))
						return v69 & v67
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
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
						return base.B2i32(v78 != v79)
					} else {
						return int32(0)
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
					if v42 == int32(0) {
						return int32(0)
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
						v48 = int32(1)
						v49 = v47 ^ v48
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v49)
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
						v53 = v51 ^ v48
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v53)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)) = uint8(v48)
						v57 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v57)
						v61 = F_make_range(m, l0, l1, l2, v57, v57)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
							v67 = int32(1)
							v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+int32(base.Ui32(v63)>>(uint(int32(2))%32))-v67))))
							return v69 & v67
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
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
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v25 = v23 & int32(1)
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v17
	goto L6
L5:
	;
	v26 = v12 + int32(4)
	goto L6
L6:
	;
	if v23 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v58 = v53
	goto L18
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v32 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v35 = int32(16)
	goto L13
L12:
	;
	v35 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = int32(4)
	goto L16
L15:
	;
	v42 = v35
	goto L16
L16:
	;
	v53 = v42
	goto L7
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v58 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v77 = int32(1)
	v78 = v19 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v83 = v81 & v77
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v76 = v53 & (v53 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v70 = v58 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v70))))
	if v72 == int32(32) {
		v58 = v70
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v76 = v58
	goto L20
L25:
	;
	v84 = v78
	goto L27
L26:
	;
	v84 = v19 + int32(4)
	goto L27
L27:
	;
	if v81 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v116 = v111
	goto L39
L29:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v90 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v101 = int32(1)
	if v83 != 0 {
		v111 = int32(base.Ui32(v81)>>(uint(v101)%32)) - v101
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v93 = int32(16)
	goto L34
L33:
	;
	v93 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v100 = int32(4)
	goto L37
L36:
	;
	v100 = v93
	goto L37
L37:
	;
	v111 = v100
	goto L28
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v116 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v135 = int32(1)
	if v23&v135 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v134 = v111 & (v111 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v128 = v116 - int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v128))))
	if v130 == int32(32) {
		v116 = v128
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v134 = v116
	goto L41
L46:
	;
	v139 = v135
	goto L48
L47:
	;
	v139 = int32(4)
	goto L48
L48:
	;
	v141 = int32(1)
	if v81&v141 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v145 = v141
	goto L51
L50:
	;
	v145 = int32(4)
	goto L51
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v148 = F_varstr_cmp(m, v12+v139, v76, v19+v145, v134, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v150 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v154 != v19 {
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
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return v148
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
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
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v25 = v23 & int32(1)
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v17
	goto L6
L5:
	;
	v26 = v12 + int32(4)
	goto L6
L6:
	;
	if v23 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v58 = v53
	goto L18
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v32 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v35 = int32(16)
	goto L13
L12:
	;
	v35 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = int32(4)
	goto L16
L15:
	;
	v42 = v35
	goto L16
L16:
	;
	v53 = v42
	goto L7
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v58 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v77 = int32(1)
	v78 = v19 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v83 = v81 & v77
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v76 = v53 & (v53 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v70 = v58 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v70))))
	if v72 == int32(32) {
		v58 = v70
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v76 = v58
	goto L20
L25:
	;
	v84 = v78
	goto L27
L26:
	;
	v84 = v19 + int32(4)
	goto L27
L27:
	;
	if v81 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v116 = v111
	goto L39
L29:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v90 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v101 = int32(1)
	if v83 != 0 {
		v111 = int32(base.Ui32(v81)>>(uint(v101)%32)) - v101
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v93 = int32(16)
	goto L34
L33:
	;
	v93 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v100 = int32(4)
	goto L37
L36:
	;
	v100 = v93
	goto L37
L37:
	;
	v111 = v100
	goto L28
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v116 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v135 = int32(1)
	if v23&v135 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v134 = v111 & (v111 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v128 = v116 - int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v128))))
	if v130 == int32(32) {
		v116 = v128
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v134 = v116
	goto L41
L46:
	;
	v139 = v135
	goto L48
L47:
	;
	v139 = int32(4)
	goto L48
L48:
	;
	v141 = int32(1)
	if v81&v141 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v145 = v141
	goto L51
L50:
	;
	v145 = int32(4)
	goto L51
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v148 = F_varstr_cmp(m, v12+v139, v76, v19+v145, v134, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v150 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v154 != v19 {
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
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return int32(base.Ui32(v148) >> (uint(int32(31)) % 32))
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v76 int32
	_ = v76
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
	var v116 int32
	_ = v116
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
	var v249 float64
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 float32
	_ = v384
	var v387 float64
	_ = v387
	var v388 float64
	_ = v388
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 float64
	_ = v433
	var v434 int32
	_ = v434
	var v435 float64
	_ = v435
	var v442 float64
	_ = v442
	var v444 float64
	_ = v444
	var v445 float64
	_ = v445
	var v446 float64
	_ = v446
	var v447 float64
	_ = v447
	var v455 float64
	_ = v455
	var v457 float64
	_ = v457
	var v458 int32
	_ = v458
	var v459 float64
	_ = v459
	var v460 int32
	_ = v460
	var v461 float64
	_ = v461
	var v464 float64
	_ = v464
	var v466 float64
	_ = v466
	var v471 float64
	_ = v471
	var v474 float64
	_ = v474
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
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
	v48 = v31
	v49 = v9
	v50 = v9
	goto L4
L4:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v49<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v61 == int32(0) {
		v116 = v48
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
	v125 = v49 + int32(1)
	if v125 < v116 {
		v48 = v116
		v49 = v125
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
		v116 = v48
		v118 = v50
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v76 = v64
	v84 = v50
	goto L9
L9:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v76<<(uint(int32(2))%32))))
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
	v116 = v101
	v118 = v95
	goto L6
L11:
	;
	return
L12:
	;
	v98 = v76 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v98 < v99 {
		v76 = v98
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
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v252 == int32(0) {
		goto L39
	} else {
		goto L40
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
		goto L36
	} else {
		goto L37
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
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_brincostestimate[0]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192+(v184^int32(-1))<<(uint(int32(2))%32))))
	v206 = v198
	goto L27
L29:
	;
	goto L30
L30:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_brincostestimate[1]))
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
	v249 = v227
	goto L20
L36:
	;
	v238 = v234
	goto L38
L37:
	;
	v238 = v235
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+76)) = base.I32_trunc_sat_f64_u(base.F64_add(base.F64_div(v238, float64(1360)), float64(1)))
	v249 = v238
	goto L20
L39:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v149)+68))
	v431 = int32(0)
	v433 = F_clauselist_selectivity(m, l0, v143, v430, v431, v431)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L11
	} else {
		goto L85
	}
L40:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v255 <= int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v269 = int32(0)
	goto L42
L42:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v283 = int32(2)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282+v269<<(uint(v283)%32))))
	v287 = int32(*(*int16)(unsafe.Add(mBase, uint32(v286)+14)))
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281+v287<<(uint(v283)%32)))))
	if v291 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L39
L44:
	;
	if v368 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = int32(1490)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v364
	v368 = v364
	goto L44
L46:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_brincostestimate[2]))
	if v293 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v328 = base.I32_extend16_s(v287 + int32(1))
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_brincostestimate[3]))
	if v330 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L49:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v166)+16))
	v324 = F_SearchSysCache3(m, int32(65), v321, base.I32_extend16_s(v291), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L11
	} else {
		goto L58
	}
L50:
	;
	v299 = m.T0[v293].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v166, base.I32_extend16_s(v291), v25+int32(40))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	if v299 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	if v303 == int32(0) {
		v368 = v303
		goto L44
	} else {
		goto L53
	}
L53:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	if v306 != 0 {
		v368 = v303
		goto L44
	} else {
		goto L54
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	F_errmsg_internal(m, int32(_a_F_brincostestimate_0), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_brincostestimate_1), int32(_a_F_brincostestimate_2), int32(_a_F_brincostestimate_3))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v364 = v324
	goto L45
L59:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v360 = F_SearchSysCache3(m, int32(65), v358, v328, int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L11
	} else {
		goto L68
	}
L60:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v336 = m.T0[v330].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v333, v328, v25+int32(40))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	if v336 == int32(0) {
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	if v340 == int32(0) {
		v368 = v340
		goto L44
	} else {
		goto L63
	}
L63:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	if v343 != 0 {
		v368 = v340
		goto L44
	} else {
		goto L64
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	F_errmsg_internal(m, int32(_a_F_brincostestimate_0), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_brincostestimate_1), int32(_a_F_brincostestimate_4), int32(_a_F_brincostestimate_3))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v364 = v360
	goto L45
L69:
	;
	v405 = v269 + int32(1)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v405 < v406 {
		v269 = v405
		goto L42
	} else {
		goto L84
	}
L70:
	;
	v377 = F_get_attstatsslot(m, v25+int32(4), v368, int32(3), int32(0), int32(2))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	if v377 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v379 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	if v396 == int32(0) {
		goto L69
	} else {
		goto L82
	}
L75:
	;
	v387 = float64(0)
	goto L77
L76:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v384 = *(*float32)(unsafe.Add(mBase, uint32(v383)))
	v387 = base.F64_promote_f32(base.F32_abs(v384))
	goto L77
L77:
	;
	v388 = *(*float64)(unsafe.Add(mBase, uint32(l6)))
	if base.F64_gt(v387, v388) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l6))) = v387
	goto L80
L79:
	;
	goto L80
L80:
	;
	F_free_attstatsslot(m, v25+int32(4))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	goto L74
L82:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	m.T0[v399].(func(*base.Module, int32))(m, v396)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	goto L69
L84:
	;
	goto L43
L85:
	;
	v435 = *(*float64)(unsafe.Add(mBase, uint32(l6)))
	if base.F64_lt(v435, float64(1e-10)) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v455
	v457 = F_index_other_operands_eval_cost(m, l0, v143)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L11
	} else {
		goto L95
	}
L87:
	;
	v442 = base.F64_div(base.F64_ceil(base.F64_mul(v249, v433)), v435)
	if base.F64_gt(v249, v442) != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v445 = v249
	goto L89
L89:
	;
	v446 = float64(0)
	v447 = base.F64_div(v445, v249)
	if base.F64_lt(v447, v446) != 0 {
		v455 = v446
		goto L86
	} else {
		goto L93
	}
L90:
	;
	v444 = v442
	goto L92
L91:
	;
	v444 = v249
	goto L92
L92:
	;
	v445 = v444
	goto L89
L93:
	;
	if base.F64_gt(v447, float64(1)) == int32(0) {
		v455 = v447
		goto L86
	} else {
		goto L94
	}
L94:
	;
	v455 = float64(1)
	goto L86
L95:
	;
	v459 = *(*float64)(unsafe.Add(mBase, uint32(v25)+88))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v461 = base.F64_convert_i32_u(v460)
	v464 = base.F64_add(v457, base.F64_mul(l2, base.F64_mul(v459, v461)))
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v464
	v466 = *(*float64)(unsafe.Add(mBase, uint32(v25)+80))
	v471 = base.F64_add(base.F64_mul(base.F64_mul(v466, base.F64_sub(base.F64_convert_i32_u(v150), v461)), l2), v464)
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = v471
	v474 = *(*float64)(unsafe.Add(mBase, _c_F_brincostestimate[4]))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	*(*float64)(unsafe.Add(mBase, uint32(l4))) = base.F64_add(base.F64_mul(base.F64_mul(v445, base.F64_mul(v474, float64(0.1))), base.F64_convert_i32_u(v478)), v471)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l7))) = base.F64_convert_i32_u(v483)
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int64
	_ = v346
	var v355 int32
	_ = v355
	var v367 int32
	_ = v367
	var v374 int64
	_ = v374
	var v378 int64
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
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
	var v423 int32
	_ = v423
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	var v670 int32
	_ = v670
	var v682 int64
	_ = v682
	var v683 int64
	_ = v683
	var v685 int64
	_ = v685
	var v716 int64
	_ = v716
	var v718 int64
	_ = v718
	var v723 int32
	_ = v723
	var v726 int64
	_ = v726
	var v727 int64
	_ = v727
	var v729 int64
	_ = v729
	var v730 int64
	_ = v730
	var v731 int64
	_ = v731
	var v733 int64
	_ = v733
	var v743 int32
	_ = v743
	var v755 int32
	_ = v755
	var v766 int64
	_ = v766
	var v767 int64
	_ = v767
	var v768 int64
	_ = v768
	var v803 int64
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
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
	F_relation_close(m, v69, int32(1))
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
	if int32(0) < v105 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v121 = int32(0)
	v122 = v111 + v114
	goto L20
L18:
	;
	v180 = v107
	goto L19
L19:
	;
	if v180 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v149 = int32(2)
	v150 = v121 << (uint(v149) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v102+v150))) = v122
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v157 = int32(7)
	v159 = int32(-8)
	v161 = v122 + (v154<<(uint(v149)%32)+v157)&v159
	*(*int32)(unsafe.Add(mBase, uint32(v150+v112))) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v172 = v121 + int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if v172 < v174 {
		v121 = v172
		v122 = v161 + (v163<<(uint(v149)%32)+v157)&v159
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v180 = v174 << (uint(int32(2)) % 32)
	goto L19
L22:
	;
	goto L21
L23:
	;
	base.MemoryFill(m, v113, int32(0), v180)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v213 = v211 << (uint(int32(2)) % 32)
	if v213 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	base.MemoryFill(m, v114, int32(0), v213)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v216 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v223 = int32(0)
	goto L32
L30:
	;
	goto L31
L31:
	;
	v332 = F_brin_new_memtuple(m, v41)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L6
	} else {
		goto L46
	}
L32:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v253 = v250 + v223*int32(48)
	v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v253)+4)))
	v256 = v254 - int32(1)
	v259 = v81 + v256*int32(28)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v260 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	v264 = F_index_getprocinfo(m, v35, v254, int32(3))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v280 = v256 << (uint(int32(2)) % 32)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v283 = v281 & int32(1)
	if v283 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0]))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v264)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v259)+16)) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v264)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v259)+24)) = v270
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v264)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v259)+8)) = v272
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v264)))
	*(*int64)(unsafe.Add(mBase, uint32(v259))) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v259)+20)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v259)+16)) = int32(0)
	goto L38
L38:
	;
	goto L36
L39:
	;
	v284 = v112
	goto L41
L40:
	;
	v284 = v102
	goto L41
L41:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v280+v284)))
	if v283 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v287 = v114
	goto L44
L43:
	;
	v287 = v113
	goto L44
L44:
	;
	v288 = v287 + v280
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	*(*int32)(unsafe.Add(mBase, uint32(v286+v289<<(uint(int32(2))%32)))) = v253
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v295 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v288))) = v294 + v295
	v299 = v223 + v295
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v299 < v300 {
		v223 = v299
		goto L32
	} else {
		goto L45
	}
L45:
	;
	goto L33
L46:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0]))
	v340 = F_AllocSetContextCreateInternal(m, v335, int32(_a_F_bringetbitmap_0), int32(0), int32(_a_F_bringetbitmap_1), int32(_a_F_bringetbitmap_2))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v342 = int32(_a_F_bringetbitmap_3)
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0])) = v340
	if v72 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v346 = base.I64_extend_i32_u(v72)
	v355 = v332
	v367 = v3
	v374 = v26
	v378 = v26
	goto L51
L49:
	;
	v803 = int64(0)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0])) = v343
	F_MemoryContextDelete(m, v340)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L6
	} else {
		goto L117
	}
L51:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[1]))
	if v380 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v803 = v766 * int64(10)
	goto L50
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_MemoryContextReset(m, v340)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L6
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v391 = F_brinGetTupleForHeapBlock(m, v385, base.I32_wrap_i64(v374), v33+int32(12), v33+int32(6), v33)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	v767 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40))))
	v768 = v374 + v767
	if base.Ui64(v768) < base.Ui64(v346) {
		v355 = v743
		v367 = v755
		v374 = v768
		v378 = v766
		goto L51
	} else {
		goto L116
	}
L59:
	;
	v682 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40))))
	v683 = v374 + v682
	if base.Ui64(v683) < base.Ui64(v346) {
		goto L105
	} else {
		goto L106
	}
L60:
	;
	if v391 == int32(0) {
		v658 = v355
		v670 = v367
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v398 = F_brin_copy_tuple(m, v391, v395, v367, v33+int32(8))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	F_LockBuffer(m, v400, int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v404 = F_brin_deform_tuple(m, v41, v398, v355)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404))))
	if v406 != 0 {
		v658 = v404
		v670 = v398
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	if v408 <= int32(0) {
		v658 = v404
		v670 = v398
		goto L59
	} else {
		goto L66
	}
L66:
	;
	v423 = int32(1)
	goto L67
L67:
	;
	v445 = v423 - int32(1)
	v447 = v445 << (uint(int32(2)) % 32)
	v448 = v113 + v447
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v449 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v658 = v404
	v670 = v398
	goto L59
L69:
	;
	v648 = v423 + int32(1)
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v649)))
	if v648 <= v650 {
		v423 = v648
		goto L67
	} else {
		goto L104
	}
L70:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v447+v114)))
	if v453 == int32(0) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+1)))
	if v456 != 0 {
		v743 = v404
		v755 = v398
		v766 = v378
		goto L58
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v459 = v404 + int32(24) + v445*int32(20)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v447+(v41+int32(20)))))
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+2)))
	if v462 != int32(1) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v449 == int32(0) {
		goto L69
	} else {
		goto L90
	}
L76:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v447+v114)))
	if v466 <= int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v447+v112)))
	v474 = int32(0)
	goto L78
L78:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v470+v474<<(uint(int32(2))%32))))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	if v506&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L75
L80:
	;
	v521 = v474 + int32(1)
	if v521 != v466 {
		v474 = v521
		goto L78
	} else {
		goto L89
	}
L81:
	;
	if v506&int32(64) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+3)))
	if v513 != 0 {
		goto L80
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v506&int32(128) == int32(0) {
		v743 = v404
		v755 = v398
		v766 = v378
		goto L58
	} else {
		goto L87
	}
L85:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+2)))
	if v514 != 0 {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v743 = v404
	v755 = v398
	v766 = v378
	goto L58
L87:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+3)))
	if v519 != 0 {
		v743 = v404
		v755 = v398
		v766 = v378
		goto L58
	} else {
		goto L88
	}
L88:
	;
	goto L80
L89:
	;
	goto L79
L90:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+3)))
	if v555 != 0 {
		v743 = v404
		v755 = v398
		v766 = v378
		goto L58
	} else {
		goto L91
	}
L91:
	;
	v556 = v447 + v102
	v559 = v81 + v445*int32(28)
	v560 = int32(*(*int16)(unsafe.Add(mBase, uint32(v559)+8)))
	if v560 <= int32(3) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v575 = int32(0)
	goto L99
L93:
	;
	if v449 <= int32(0) {
		goto L69
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v566)))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v569 = F_FunctionCall4Coll(m, v559, v568, v41, v459, v566, v449)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L6
	} else {
		goto L97
	}
L96:
	;
	goto L92
L97:
	;
	if v569 == int32(0) {
		v743 = v404
		v755 = v398
		v766 = v378
		goto L58
	} else {
		goto L98
	}
L98:
	;
	goto L69
L99:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v603+v575<<(uint(int32(2))%32))))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)+12))
	v609 = F_FunctionCall3Coll(m, v559, v608, v41, v459, v607)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L6
	} else {
		goto L101
	}
L100:
	;
	goto L69
L101:
	;
	if v609 == int32(0) {
		v743 = v404
		v755 = v398
		v766 = v378
		goto L58
	} else {
		goto L102
	}
L102:
	;
	v614 = v575 + int32(1)
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v614 < v615 {
		v575 = v614
		goto L99
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	goto L68
L105:
	;
	v685 = v683
	goto L107
L106:
	;
	v685 = v346
	goto L107
L107:
	;
	if base.Ui64(v685-int64(1)) < base.Ui64(v374) {
		v743 = v658
		v755 = v670
		v766 = v378
		goto L58
	} else {
		goto L108
	}
L108:
	;
	v716 = v374
	v718 = v378
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0])) = v343
	F_tbm_add_page(m, l1, base.I32_wrap_i64(v716))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L6
	} else {
		goto L111
	}
L110:
	;
	v743 = v658
	v755 = v670
	v766 = v727
	goto L58
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_bringetbitmap[0])) = v340
	v726 = int64(1)
	v727 = v718 + v726
	v729 = v716 + v726
	v730 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v40))))
	v731 = v374 + v730
	if base.Ui64(v731) < base.Ui64(v346) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v733 = v731
	goto L114
L113:
	;
	v733 = v346
	goto L114
L114:
	;
	if base.Ui64(v729) <= base.Ui64(v733-int64(1)) {
		v716 = v729
		v718 = v727
		goto L109
	} else {
		goto L115
	}
L115:
	;
	goto L110
L116:
	;
	goto L52
L117:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	if v808 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_ReleaseBuffer(m, v808)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L6
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	m.G0 = v33 + int32(16)
	return v803
L121:
	;
	goto L120
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
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = int32(2147483647)
	v9 = base.I32_reinterpret_f32(v6) & v8
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = base.I32_reinterpret_f32(v10) & v8
	if base.Ui32(int32(2139095041)) <= base.Ui32(v13) {
		v23 = base.B2i32(base.Ui32(v9) < base.Ui32(int32(2139095041)))
		v32 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v13))|base.F32_lt(v6, v10))&v23
	} else {
		v18 = int32(1)
		if base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v9))|base.F32_gt(v6, v10) != 0 {
			v32 = v18
		} else {
			v23 = v18
			v32 = int32(0) - (base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v13))|base.F32_lt(v6, v10))&v23
		}
	}
	return v32
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
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v7) & v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	v15 = base.I64_reinterpret_f64(v12) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v15) {
		v25 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v34 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v25
	} else {
		v20 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v10))|base.F64_gt(v7, v12) != 0 {
			v34 = v20
		} else {
			v25 = v20
			v34 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v7, v12))&v25
		}
	}
	return v34
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
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v10 = int64(9223372036854775807)
	v11 = base.I64_reinterpret_f64(v8) & v10
	v12 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v15 = base.I64_reinterpret_f64(v12) & v10
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v15) {
		v25 = base.B2i32(base.Ui64(v11) < base.Ui64(int64(9218868437227405313)))
		v34 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v8, v12))&v25
	} else {
		v20 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v11))|base.F64_gt(v8, v12) != 0 {
			v34 = v20
		} else {
			v25 = v20
			v34 = int32(0) - (base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v8, v12))&v25
		}
	}
	return v34
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
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = Fn13855(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_bttextsortsupport(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13854(m, l0, int32(25))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
	return v8 & int32(_a_F_bttranslatecmptype_0)
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
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
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
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
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
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int64
	_ = v417
	var v420 int64
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
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
	var v433 int32
	_ = v433
	var v434 int64
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
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
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
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
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
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
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	if int32(0) < v230 {
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
	v205 = m.ExcPending
	if v205 != 0 {
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
		v219 = v39
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
		goto L22
	case 1:
		goto L16
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
	v219 = v195
	goto L1
L13:
	;
	v199 = v51 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v199 < v200 {
		v51 = v199
		v55 = v195
		goto L11
	} else {
		goto L43
	}
L14:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L39
	}
L15:
	;
	v160 = int32(0)
	v163 = F_errstart(m, int32(17), v160)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L37
	}
L16:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = int32(2281)
	v151 = int32(1)
	v156 = F_check_amproc_signature(m, v147, int32(2278), v151, v151, v151, v18+int32(160))
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
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = v77
	v81 = int32(2)
	v85 = F_check_amproc_signature(m, v76, int32(23), int32(1), v81, v81, v18+int32(144))
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
	v195 = v55
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
	v195 = v55
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
	v195 = v55
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
	v195 = v55
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
	v195 = v55
	goto L13
L33:
	;
	if v141 == int32(0) {
		v195 = v138
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v170 = int32(119)
	v172 = int32(_a_F_btvalidate_0)
	goto L14
L35:
	;
	if v156 != 0 {
		v195 = v55
		goto L13
	} else {
		goto L36
	}
L36:
	;
	goto L15
L37:
	;
	if v163 == int32(0) {
		v195 = v160
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v170 = int32(131)
	v172 = int32(_a_F_btvalidate_1)
	goto L14
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v177 = F_format_procedure(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v31
	F_errmsg(m, v172, v18+int32(128))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), v170, int32(_a_F_btvalidate_4))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v195 = int32(0)
	goto L13
L43:
	;
	goto L12
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, int32(_a_F_btvalidate_5), v18)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(61), int32(_a_F_btvalidate_4))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
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
	v238 = int32(0)
	v240 = v219
	goto L50
L48:
	;
	v372 = v219
	goto L49
L49:
	;
	v383 = F_identify_opfamily_groups(m, v37, v44)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L2
	} else {
		goto L83
	}
L50:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(48)+v238<<(uint(int32(2))%32))))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+56))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+22)))
	v257 = v255 + v256
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+16)))
	if base.Ui32(int32(_a_F_btvalidate_6)) < base.Ui32((v258-int32(6))&int32(_a_F_btvalidate_7)) {
		v294 = v240
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v372 = v363
	goto L49
L52:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+18)))
	if v296 == int32(115) {
		goto L61
	} else {
		goto L62
	}
L53:
	;
	v265 = int32(0)
	v268 = F_errstart(m, int32(17), v265)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	if v268 == int32(0) {
		v294 = v265
		goto L52
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v257)+20))
	v276 = F_format_operator(m, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v278 = int32(*(*int16)(unsafe.Add(mBase, uint32(v257)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v31
	F_errmsg(m, int32(_a_F_btvalidate_8), v18+int32(112))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(151), int32(_a_F_btvalidate_4))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v294 = v265
	goto L52
L60:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v257)+20))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v334 = F_check_amop_signature(m, v330, int32(16), v332, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L72
	}
L61:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v257)+28))
	if v299 == int32(0) {
		v329 = v294
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v302 = int32(0)
	v305 = F_errstart(m, int32(17), v302)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L2
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	if v305 == int32(0) {
		v329 = v302
		goto L60
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v257)+20))
	v313 = F_format_operator(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v31
	F_errmsg(m, int32(_a_F_btvalidate_9), v18+int32(96))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(163), int32(_a_F_btvalidate_4))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v329 = v302
	goto L60
L71:
	;
	v365 = v238 + int32(1)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v37)+40))
	if v365 < v366 {
		v238 = v365
		v240 = v363
		goto L50
	} else {
		goto L80
	}
L72:
	;
	if v334 != 0 {
		v363 = v329
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v336 = int32(0)
	v339 = F_errstart(m, int32(17), v336)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	if v339 == int32(0) {
		v363 = v336
		goto L71
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v257)+20))
	v347 = F_format_operator(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v31
	F_errmsg(m, int32(_a_F_btvalidate_10), v18+int32(80))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(176), int32(_a_F_btvalidate_4))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v363 = v336
	goto L71
L80:
	;
	goto L51
L81:
	;
	if v575 != 0 {
		goto L133
	} else {
		goto L134
	}
L82:
	;
	v551 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L2
	} else {
		goto L125
	}
L83:
	;
	if v383 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v387 = int32(0)
	v537 = v387
	v539 = v387
	goto L82
L85:
	;
	goto L86
L86:
	;
	v390 = int32(0)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v391 <= v390 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v517 == int32(0) {
		v575 = v520
		v576 = v521
		v577 = v522
		goto L81
	} else {
		goto L124
	}
L88:
	;
	v517 = int32(1)
	v520 = int32(0)
	v521 = v372
	v522 = v390
	goto L87
L89:
	;
	goto L90
L90:
	;
	v395 = int32(0)
	v399 = v395
	v400 = v395
	v401 = v372
	v402 = v390
	v408 = int32(0)
	goto L91
L91:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v412+v399<<(uint(int32(2))%32))))
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v416)+8))
	if v417 == int64(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v517 = base.B2i32(v510 == int32(0))
	v520 = v506
	v521 = v507
	v522 = v508
	goto L87
L93:
	;
	v512 = v399 + int32(1)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v512 < v513 {
		v399 = v512
		v400 = v506
		v401 = v507
		v402 = v508
		v408 = v510
		goto L91
	} else {
		goto L123
	}
L94:
	;
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v416)+16))
	if v420 == int64(8) {
		v506 = v400
		v507 = v401
		v508 = v402
		v510 = v408
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	if v28 == v423 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	if v425 == v28 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v428 = v408
	goto L100
L100:
	;
	v429 = F_list_append_unique_oid(m, v400, v423)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L2
	} else {
		goto L104
	}
L101:
	;
	v427 = v416
	goto L103
L102:
	;
	v427 = v408
	goto L103
L103:
	;
	v428 = v427
	goto L100
L104:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v432 = F_list_append_unique_oid(m, v429, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	v434 = *(*int64)(unsafe.Add(mBase, uint32(v416)+8))
	if v434 == int64(62) {
		v468 = v401
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v471 = v402 + int32(1)
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+16)))
	if v472&int32(2) != 0 {
		v506 = v432
		v507 = v468
		v508 = v471
		v510 = v428
		goto L93
	} else {
		goto L115
	}
L107:
	;
	v437 = int32(0)
	v440 = F_errstart(m, int32(17), v437)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	if v440 == int32(0) {
		v468 = v437
		goto L106
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L2
	} else {
		goto L110
	}
L110:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v448 = F_format_type_be(m, v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v451 = F_format_type_be(m, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v31
	F_errmsg(m, int32(_a_F_btvalidate_11), v18-int32(-64))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(235), int32(_a_F_btvalidate_4))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v468 = v437
	goto L106
L115:
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
L116:
	;
	if v478 == int32(0) {
		v506 = v432
		v507 = v475
		v508 = v471
		v510 = v428
		goto L93
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
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v486 = F_format_type_be(m, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	v489 = F_format_type_be(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v31
	F_errmsg(m, int32(_a_F_btvalidate_12), v18+int32(48))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(245), int32(_a_F_btvalidate_4))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v506 = v432
	v507 = v475
	v508 = v471
	v510 = v428
	goto L93
L123:
	;
	goto L92
L124:
	;
	v537 = v520
	v539 = v522
	goto L82
L125:
	;
	if v551 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L2
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v575 = v537
	v576 = int32(0)
	v577 = v539
	goto L81
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v27 + int32(8)
	F_errmsg(m, int32(_a_F_btvalidate_13), v18+int32(32))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(257), int32(_a_F_btvalidate_4))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
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
	v618 = m.ExcPending
	if v618 != 0 {
		goto L2
	} else {
		goto L142
	}
L133:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	v591 = v587 * v587
	goto L135
L134:
	;
	v591 = int32(0)
	goto L135
L135:
	;
	if v591 == v577 {
		v616 = v576
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v593 = int32(0)
	v596 = F_errstart(m, int32(17), v593)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L2
	} else {
		goto L137
	}
L137:
	;
	if v596 == int32(0) {
		v616 = v593
		goto L132
	} else {
		goto L138
	}
L138:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(_a_F_btvalidate_2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v31
	F_errmsg(m, int32(_a_F_btvalidate_14), v18+int32(16))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_btvalidate_3), int32(273), int32(_a_F_btvalidate_4))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	v616 = v593
	goto L132
L142:
	;
	F_ReleaseCatCacheList(m, v37)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	F_ReleaseCatCache(m, v21)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	m.G0 = v18 + int32(240)
	return v616 & int32(1)
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
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
						if v73 <= int32(0) {
						} else {
							v84 = int32(0)
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v348 int32
	_ = v348
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
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
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 float64
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 float64
	_ = v501
	var v502 int32
	_ = v502
	var v504 float64
	_ = v504
	var v505 int32
	_ = v505
	var v523 float64
	_ = v523
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+180))
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
	v149 = m.ExcPending
	if v149 != 0 {
		goto L18
	} else {
		goto L32
	}
L4:
	;
	goto L3
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v26 = v24
	goto L7
L6:
	;
	v26 = int32(0)
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
	v121 = m.ExcPending
	if v121 != 0 {
		goto L18
	} else {
		goto L29
	}
L9:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v109 != 0 {
		goto L25
	} else {
		goto L26
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v29 <= int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v38 = v26
	v41 = v7
	goto L12
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v41<<(uint(int32(2))%32))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+26)))
	if v54 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	if v38 == int32(0) {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	v83 = v38
	goto L16
L16:
	;
	v87 = v41 + int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v87 < v88 {
		v38 = v83
		v41 = v87
		goto L12
	} else {
		goto L24
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	v68 = F_exprType(m, v62)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	F_add_child_eq_member(m, l0, v60, int32(-1), v62, v63, v67, v66, v68, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v74 = v38 + int32(4)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.Ui32(v74) < base.Ui32(v76+v77<<(uint(int32(2))%32)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v82 = v74
	goto L23
L22:
	;
	v82 = int32(0)
	goto L23
L23:
	;
	v83 = v82
	goto L16
L24:
	;
	goto L13
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v114 = v110 - int32(1)
	goto L27
L26:
	;
	v114 = int32(-1)
	goto L27
L27:
	;
	v115 = F_bms_add_range(m, v107, int32(0), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+136)) = v115
	goto L4
L29:
	;
	F_errmsg_internal(m, int32(_a_F_build_setop_child_paths_0), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_build_setop_child_paths_1), int32(3100), int32(_a_F_build_setop_child_paths_2))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
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
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v153 = F_fetch_upper_rel(m, v150, int32(7), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)) = uint8(v155)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	if v157 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_build_setop_child_paths[0]))
	if v402 != 0 {
		goto L113
	} else {
		goto L114
	}
L35:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v372 != 0 {
		goto L34
	} else {
		goto L109
	}
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if int32(0) < v158 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v155&int32(1) == int32(0) {
		goto L34
	} else {
		goto L108
	}
L39:
	;
	v174 = v7
	goto L42
L40:
	;
	goto L41
L41:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	if v348 == int32(0) {
		goto L34
	} else {
		goto L107
	}
L42:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v174<<(uint(int32(2))%32))))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v153)+48))
	if v182 == v183 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182)+64))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v187 = F_make_tlist_from_pathtarget(m, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
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
	v189 = F_convert_subquery_pathkeys(m, l0, l1, v185, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v192 = F_create_subqueryscan_path(m, l0, l1, v182, l2, v189, int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_add_path(m, l1, v192)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v328 = v174 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v328 < v329 {
		v174 = v328
		goto L42
	} else {
		goto L106
	}
L52:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v182)+64))
	v200 = v20 + int32(12)
	if v23 == v198 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	if v303 == v183 {
		goto L51
	} else {
		goto L101
	}
L54:
	;
	if v278 != 0 {
		v303 = v182
		goto L53
	} else {
		goto L86
	}
L55:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v266
	v278 = int32(1)
	goto L54
L56:
	;
	if v23 != 0 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v23 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = int32(0)
	v278 = int32(1)
	goto L54
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = int32(0)
	v278 = int32(1)
	goto L54
L61:
	;
	goto L62
L62:
	;
	if v198 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v218 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v218
	v278 = v218
	goto L54
L64:
	;
	goto L65
L65:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v222 = int32(0)
	if v222 < v221 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v225 = v221
	goto L68
L67:
	;
	v225 = v222
	goto L68
L68:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v231 = int32(0)
	goto L69
L69:
	;
	if v231 < v226 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v242 = v238 + v231<<(uint(int32(2))%32)
	goto L73
L72:
	;
	v242 = int32(0)
	goto L73
L73:
	;
	if v231 == v225 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v225
	v278 = base.B2i32(v242 == int32(0))
	goto L54
L75:
	;
	goto L76
L76:
	;
	v248 = base.B2i32(v242 == int32(0))
	if v242 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v231
	v278 = v248
	goto L54
L78:
	;
	goto L79
L79:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	if v252 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v231
	v278 = v248
	goto L54
L81:
	;
	goto L82
L82:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v252+v231<<(uint(int32(2))%32))))
	if v256 != v260 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v231
	v278 = int32(0)
	goto L54
L84:
	;
	v231 = v231 + int32(1)
	goto L69
L86:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_setop_child_paths[1])))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v281)+304))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v182 == v183 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v292&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v292 = v280
	goto L87
L89:
	;
	goto L90
L90:
	;
	if v283 == int32(0) {
		goto L51
	} else {
		goto L91
	}
L91:
	;
	v287 = int32(1)
	if v280&v287 == int32(0) {
		goto L51
	} else {
		goto L92
	}
L92:
	;
	v292 = v287
	goto L87
L93:
	;
	v296 = v283
	goto L95
L94:
	;
	v296 = int32(0)
	goto L95
L95:
	;
	if v296 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v299 = F_create_sort_path(m, v153, v182, v23, v282)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L18
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v301 = F_create_incremental_sort_path(m, v281, v153, v182, v23, v283, v282)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L18
	} else {
		goto L100
	}
L99:
	;
	v303 = v299
	goto L53
L100:
	;
	v303 = v301
	goto L53
L101:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v303)+64))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v303)+12))
	v312 = F_make_tlist_from_pathtarget(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L18
	} else {
		goto L102
	}
L102:
	;
	v314 = F_convert_subquery_pathkeys(m, l0, l1, v310, v312)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L18
	} else {
		goto L103
	}
L103:
	;
	v317 = F_create_subqueryscan_path(m, l0, l1, v303, l2, v314, int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L18
	} else {
		goto L104
	}
L104:
	;
	F_add_path(m, l1, v317)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L18
	} else {
		goto L105
	}
L105:
	;
	goto L51
L106:
	;
	goto L43
L107:
	;
	goto L35
L108:
	;
	goto L35
L109:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v153)+40))
	if v373 == int32(0) {
		goto L34
	} else {
		goto L110
	}
L110:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	v378 = int32(0)
	v380 = F_create_subqueryscan_path(m, l0, l1, v377, l2, v378, v378)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L18
	} else {
		goto L111
	}
L111:
	;
	F_add_partial_path(m, l1, v380)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L18
	} else {
		goto L112
	}
L112:
	;
	goto L34
L113:
	;
	v403 = int32(0)
	m.T0[v402].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, v403, v403, l1, v403)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L18
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	F_set_cheapest(m, l1)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L18
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	if l5 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+100))
	if v412 != 0 {
		goto L123
	} else {
		goto L124
	}
L119:
	;
	goto L120
L120:
	;
	m.G0 = v20 + int32(16)
	return
L121:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l5))) = v523
	goto L120
L122:
	;
	v421 = int32(0)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v411)+76))
	if v424 == v421 {
		v499 = v421
		goto L129
	} else {
		goto L130
	}
L123:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v420 = *(*float64)(unsafe.Add(mBase, uint32(v419)+32))
	v523 = v420
	goto L121
L124:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v411)+108))
	if v413 != 0 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v411)+120))
	if v414 != 0 {
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410)+318)))
	if v415 != 0 {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+36)))
	if v416 != int32(1) {
		goto L122
	} else {
		goto L128
	}
L128:
	;
	goto L123
L129:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v501 = *(*float64)(unsafe.Add(mBase, uint32(v500)+32))
	v502 = int32(0)
	v504 = F_estimate_num_groups(m, v410, v499, v501, v502, v502)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L18
	} else {
		goto L141
	}
L130:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	if int32(0) < v427 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v430 = v421
	v432 = v421
	goto L134
L132:
	;
	v465 = v421
	goto L133
L133:
	;
	v499 = v465
	goto L129
L134:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v424)+12))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v447+v432<<(uint(int32(2))%32))))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451)+26)))
	if v452&int32(1) == int32(0) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v465 = v460
	goto L133
L136:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v458 = F_lappend(m, v430, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L18
	} else {
		goto L139
	}
L137:
	;
	v460 = v430
	goto L138
L138:
	;
	v462 = v432 + int32(1)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	if v462 < v463 {
		v430 = v460
		v432 = v462
		goto L134
	} else {
		goto L140
	}
L139:
	;
	v460 = v458
	goto L138
L140:
	;
	goto L135
L141:
	;
	v523 = v504
	goto L121
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
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v139 int32
	_ = v139
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
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
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
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
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 float64
	_ = v330
	var v331 int32
	_ = v331
	var v342 float64
	_ = v342
	var v344 int32
	_ = v344
	var v348 float64
	_ = v348
	var v349 float64
	_ = v349
	var v352 float64
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
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
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
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
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
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
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v599 int32
	_ = v599
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
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
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
	v158 = int32(1)
	goto L12
L12:
	;
	v160 = v21 + int32(40)
	v161 = int32(0)
	if l5|base.B2i32(v158 == v161) == v161 {
		goto L27
	} else {
		goto L28
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
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v158 = base.B2i32(v139 == int32(0))
	goto L12
L16:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v83 = int32(2)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v79<<(uint(v83)%32))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v96 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32(v88-int32(9)) < base.Ui32(v83))|base.B2i32(v88 == int32(61)) == v96)&base.B2i32(v88 != int32(319)) == v96 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+44)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l0
	v108 = F_process_sublinks_mutator(m, v87, v18+int32(40))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v110 = v87
	goto L20
L20:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v113 = F_lappend_int(m, v111, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v110 = v108
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v113
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v117 = F_lappend(m, v116, v110)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v117
	v121 = v79 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v121 < v122 {
		v79 = v121
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L17
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L137
	}
L26:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+8))
	v545 = F_lappend(m, v544, v529)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L121
	}
L27:
	;
	v169 = F_generate_new_exec_param(m, l0, int32(16), int32(-1), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v181 = int32(1)
	v183 = v158 ^ v181
	if v183|base.B2i32(l5 != int32(4)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v171
	v177 = F_list_make1_impl(m, int32(471), v18+int32(8))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v177
	v529 = l1
	v532 = v169
	v540 = int32(1)
	goto L26
L32:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v193 = F_exprType(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.B2i32(l5 != int32(6))|v183 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v196 = F_exprTypmod(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v199 = F_exprCollation(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v201 = F_generate_new_exec_param(m, l0, v193, v196, v199)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v203
	v209 = F_list_make1_impl(m, int32(471), v18+int32(12))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v209
	v529 = l1
	v532 = v201
	v540 = v181
	goto L26
L40:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v221 = F_exprType(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v245 = v21 + int32(12)
	if v158^int32(1)|base.B2i32(l5 != int32(3)) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	v223 = F_get_promoted_array_type(m, v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v223 == int32(0) {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v228 = F_exprTypmod(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v231 = F_exprCollation(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v233 = F_generate_new_exec_param(m, l0, v223, v228, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v235
	v241 = F_list_make1_impl(m, int32(471), v18+int32(24))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v241
	v529 = l1
	v532 = v233
	v540 = v181
	goto L26
L50:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v254 = F_generate_subquery_params(m, l0, v253, v245)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if l5 == int32(5) {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l0
	v260 = F_convert_testexpr_mutator(m, l7, v18+int32(40))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v263 = F_list_copy(m, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v263
	v529 = l1
	v532 = v260
	v540 = v181
	goto L26
L56:
	;
	v529 = v514
	v532 = v21
	v540 = v525
	goto L26
L57:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v269 = F_generate_subquery_params(m, l0, v268, v160)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v310 = int32(0)
	if base.B2i32(l7 == v310)|l8 == v310 {
		goto L73
	} else {
		goto L74
	}
L60:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v281 = v271
	goto L61
L61:
	;
	if v281 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v296+l6<<(uint(int32(2))%32)-int32(4)))) = v269
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v303 != 0 {
		v514 = l1
		v525 = int32(0)
		goto L56
	} else {
		goto L70
	}
L63:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v290 = v288
	goto L65
L64:
	;
	v290 = int32(0)
	goto L65
L65:
	;
	if v290 < l6 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v293 = F_lappend(m, v281, int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L62
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v293
	v281 = v293
	goto L61
L70:
	;
	v308 = F_makeNullConst(m, int32(2249), int32(-1), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v529 = l1
	v532 = v308
	v540 = int32(1)
	goto L26
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v325
	if l5 != int32(2) {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v316 = F_generate_subquery_params(m, l0, v315, v245)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = l8
	v325 = l7
	goto L72
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l0
	v322 = F_convert_testexpr_mutator(m, l7, v18+int32(40))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v325 = v322
	goto L72
L78:
	;
	v495 = int32(0)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v496 != 0 {
		v514 = l1
		v525 = v495
		goto L56
	} else {
		goto L116
	}
L79:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v329 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v330 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v342 = *(*float64)(unsafe.Add(mBase, _c_F_build_subplan[0]))
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_build_subplan[1]))
	v348 = base.F64_mul(base.F64_mul(v342, base.F64_convert_i32_s(v344)), float64(1024))
	v349 = float64(4.294967295e+09)
	if base.F64_lt(v348, v349) != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if base.F64_gt(base.F64_mul(v330, base.F64_convert_i32_u((v331+int32(7))&int32(-8)+int32(24))), base.F64_convert_i32_u(base.I32_trunc_sat_f64_u(v352))) != 0 {
		goto L78
	} else {
		goto L85
	}
L82:
	;
	v352 = v348
	goto L84
L83:
	;
	v352 = v349
	goto L84
L84:
	;
	goto L81
L85:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v357 = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v358 == v357 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v474 == int32(0) {
		goto L78
	} else {
		goto L115
	}
L87:
	;
	v474 = int32(0)
	goto L86
L88:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	switch v361 - int32(17) {
	case 0:
		goto L91
	default:
		goto L87
	case 4:
		goto L90
	}
L89:
	;
	v474 = v447
	goto L86
L90:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v386 != 0 {
		goto L87
	} else {
		goto L100
	}
L91:
	;
	v364 = F_hash_ok_operator(m, v358)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v364 == int32(0) {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v358)+28))
	if v368 == int32(0) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	if v371 != int32(2) {
		goto L87
	} else {
		goto L95
	}
L95:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v368)+12))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v376 = F_contain_exec_param(m, v375, v356)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if v376 != 0 {
		goto L87
	} else {
		goto L97
	}
L97:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v358)+28))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	v382 = F_contain_var_clause(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v382 == int32(0) {
		v447 = int32(1)
		goto L89
	} else {
		goto L99
	}
L99:
	;
	goto L87
L100:
	;
	v387 = int32(1)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v358)+8))
	if v388 == int32(0) {
		v447 = v387
		goto L89
	} else {
		goto L101
	}
L101:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	if v391 <= int32(0) {
		v447 = v387
		goto L89
	} else {
		goto L102
	}
L102:
	;
	v401 = v357
	goto L103
L103:
	;
	v409 = int32(0)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v388)+12))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v410+v401<<(uint(int32(2))%32))))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	if v415 != int32(17) {
		v447 = v409
		goto L89
	} else {
		goto L105
	}
L104:
	;
	v447 = v437
	goto L89
L105:
	;
	v418 = F_hash_ok_operator(m, v414)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	if v418 == int32(0) {
		v447 = v409
		goto L89
	} else {
		goto L107
	}
L107:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v414)+28))
	if v422 == int32(0) {
		v447 = v409
		goto L89
	} else {
		goto L108
	}
L108:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422)+4))
	if v425 != int32(2) {
		v447 = v409
		goto L89
	} else {
		goto L109
	}
L109:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v422)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	v430 = F_contain_exec_param(m, v429, v356)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v430 != 0 {
		v447 = v409
		goto L89
	} else {
		goto L111
	}
L111:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v414)+28))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v435 = F_contain_var_clause(m, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if v435 != 0 {
		v447 = v409
		goto L89
	} else {
		goto L113
	}
L113:
	;
	v437 = int32(1)
	v439 = v401 + v437
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	if v439 < v440 {
		v401 = v439
		goto L103
	} else {
		goto L114
	}
L114:
	;
	goto L104
L115:
	;
	v477 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+36)) = uint8(v477)
	v514 = l1
	v525 = int32(0)
	goto L56
L116:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_subplan[2])))
	if v498&int32(1) == int32(0) {
		v514 = l1
		v525 = v495
		goto L56
	} else {
		goto L117
	}
L117:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v505 = v503 - int32(348)
	goto L118
L118:
	;
	if base.B2i32(base.Ui32(v505) < base.Ui32(int32(15)))&int32(base.Ui32(int32(_a_F_build_subplan_0))>>(uint(v505)%32)) != 0 {
		v514 = l1
		v525 = v495
		goto L56
	} else {
		goto L119
	}
L119:
	;
	v511 = F_materialize_finished_plan(m, l1)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v514 = v511
	v525 = v495
	goto L56
L121:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+8)) = v545
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+12))
	v551 = F_lappend(m, v550, l2)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v553)+12)) = v551
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+16))
	v557 = F_lappend(m, v556, l3)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v559)+16)) = v557
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
	if v562 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+4))
	v565 = v563
	goto L126
L125:
	;
	v565 = int32(0)
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v565
	if v540 != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v582
	v587 = F_psprintf(m, int32(_a_F_build_subplan_1), v18)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L135
	}
L128:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v568 = F_lappend(m, v567, v21)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v572 = int32(_a_F_build_subplan_2)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v573 != 0 {
		v582 = v572
		goto L127
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v568
	v582 = int32(_a_F_build_subplan_3)
	goto L127
L132:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+36)))
	if v574 != 0 {
		v582 = v572
		goto L127
	} else {
		goto L133
	}
L133:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v575)+20))
	v577 = F_bms_add_member(m, v576, v565)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v579)+20)) = v577
	v582 = v572
	goto L127
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v587
	F_cost_subplan(m, v21, v529)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	m.G0 = v18 + int32(48)
	return v532
L137:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v601 = F_exprType(m, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v603 = F_format_type_be(m, v601)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v603
	F_errmsg_internal(m, int32(_a_F_build_subplan_4), v18+int32(16))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_build_subplan_5), int32(423), int32(_a_F_build_subplan_6))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	return v90
L2:
	;
	v13 = int32(6)
	v14 = int32(_a_F_builtin_locale_encoding_0)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_builtin_locale_encoding[0])))
	if base.B2i32(v17 == int32(0))|base.B2i32(v17 != v20) != 0 {
		v38 = v17
		v39 = v20
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
	v90 = int32(-1)
	goto L1
L5:
	;
	if v38-v39 == int32(0) {
		v90 = v13
		goto L1
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v23 = l0
	v24 = v14
	goto L8
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v28
		v39 = v27
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v38 = v28
	v39 = v27
	goto L6
L10:
	;
	v31 = int32(1)
	if v28 == v27 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v43 = int32(_a_F_builtin_locale_encoding_1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_builtin_locale_encoding[1])))
	if base.B2i32(v46 == int32(0))|base.B2i32(v46 != v49) != 0 {
		v67 = v46
		v68 = v49
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v67-v68 == int32(0) {
		v90 = v13
		goto L1
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v52 = l0
	v53 = v43
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v57
		v68 = v56
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v67 = v57
	v68 = v56
	goto L14
L18:
	;
	v60 = int32(1)
	if v57 == v56 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(_a_F_builtin_locale_encoding_2), v6)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_builtin_locale_encoding_3), int32(1499), int32(_a_F_builtin_locale_encoding_4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
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
	var v160 int32
	_ = v160
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
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v76 = int32(1)
	if v16&v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = int32(1)
	if v46&v63 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = int32(16)
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(4)
	goto L24
L23:
	;
	v62 = v55
	goto L24
L24:
	;
	v75 = v62
	goto L15
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v80 = v76
	goto L28
L27:
	;
	v80 = int32(4)
	goto L28
L28:
	;
	v81 = v9 + v80
	v82 = int32(1)
	if v46&v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = int32(4)
	goto L31
L31:
	;
	v87 = v14 + v86
	if v45 < v75 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v45
	goto L34
L33:
	;
	v89 = v75
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v81|v87)&int32(3) != 0 {
		v120 = v81
		v121 = v87
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v81
	v114 = v87
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v81
	v98 = v87
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
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
	if v156 != v14 {
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
	v160 = int32(0)
	return base.B2i32(v151 == v160)&base.B2i32(v45 <= v75) | base.B2i32(v151 < v160)
L60:
	;
	goto L59
}
