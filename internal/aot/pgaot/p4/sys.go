package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReleaseSysCache(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_ReleaseCatCache(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_SearchSysCacheCopyAttName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1383]))
	v6 = F_SearchCatCache2(m, v5, l0, l1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v11)+91)))
			if v13 == int32(0) {
				v16 = F_heap_copytuple(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = v16
					F_ReleaseCatCache(m, v6)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v22 = v18
						return v22
					}
				}
			} else {
				v18 = v3
				F_ReleaseCatCache(m, v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v22 = v18
					return v22
				}
			}
		} else {
			v22 = v3
			return v22
		}
	}
}
func F_SysCacheGetAttrNotNull(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SysCacheGetAttr(m, l0, l1, l2, v8+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v16 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(5))%32))+uint32(_consts[1384])))
				v28 = F_get_rel_name(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1385])))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v35 + v36<<(uint(int32(4))%32) + l2*int32(100) - int32(76)
					F_errmsg_internal(m, int32(180690), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(487962), int32(644), int32(297155))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
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
			m.G0 = v8 + int32(16)
			return v12
		}
	}
}
func F_SysLoggerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int64
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v314 int32
	_ = v314
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v440 int32
	_ = v440
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v460 int64
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
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
	var v548 int64
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int64
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int64
	_ = v574
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int64
	_ = v583
	var v586 int64
	_ = v586
	var v588 int64
	_ = v588
	var v590 int64
	_ = v590
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v640 int64
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v763 int32
	_ = v763
	var v767 int64
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int64
	_ = v776
	var v779 int64
	_ = v779
	var v781 int64
	_ = v781
	var v783 int64
	_ = v783
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int64
	_ = v819
	var v821 int64
	_ = v821
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int64
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int64
	_ = v842
	var v843 int64
	_ = v843
	var v844 int64
	_ = v844
	var v846 int64
	_ = v846
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v862 int64
	_ = v862
	var v863 int64
	_ = v863
	var v864 int64
	_ = v864
	var v866 int64
	_ = v866
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v883 int64
	_ = v883
	var v884 int64
	_ = v884
	var v885 int64
	_ = v885
	var v887 int64
	_ = v887
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int64
	_ = v909
	var v910 int64
	_ = v910
	var v911 int64
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v939 int64
	_ = v939
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int64
	_ = v948
	var v951 int64
	_ = v951
	var v953 int64
	_ = v953
	var v955 int64
	_ = v955
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v969 int64
	_ = v969
	var v971 int64
	_ = v971
	var v972 int64
	_ = v972
	var v975 int64
	_ = v975
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1143 int32
	_ = v1143
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1280 int32
	_ = v1280
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1322 int32
	_ = v1322
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1349 int32
	_ = v1349
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1398 int32
	_ = v1398
	var v1402 int32
	_ = v1402
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1439 int32
	_ = v1439
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1454 int32
	_ = v1454
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1474 int32
	_ = v1474
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1545 int32
	_ = v1545
	var v1553 int32
	_ = v1553
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1698 int32
	_ = v1698
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1708 int32
	_ = v1708
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	v24 = m.G0
	v26 = v24 - int32(8240)
	m.G0 = v26
	v29 = *(*int32)(unsafe.Add(mBase, _consts[571]))
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_MemoryContextDelete(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[264])) = int32(17)
	v39 = *(*int64)(unsafe.Add(mBase, _consts[783]))
	goto L7
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[571])) = int32(0)
	goto L3
L6:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[784])))
	if v47 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[264]))
	v45 = F_GetBackendTypeDesc(m, v44)
	mBase = m.M
	goto L9
L9:
	;
	goto L6
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[785]))
	if int32(0) <= v70 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(0)
	v53 = int32(1)
	v56 = F_open(m, int32(295506), v53, v26+int32(16))
	mBase = m.M
	v58 = F_close(m, v53)
	mBase = m.M
	v60 = F_close(m, int32(2))
	mBase = m.M
	if v56 == int32(-1) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v64 = F_dup2(m, v56, int32(1))
	mBase = m.M
	v66 = F_dup2(m, v56, int32(2))
	mBase = m.M
	v67 = F_close(m, v56)
	mBase = m.M
	goto L10
L13:
	;
	v73 = F_close(m, v70)
	mBase = m.M
	goto L15
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[785])) = int32(-1)
	v78 = int32(913)
	v80 = m.G0
	v82 = v80 - int32(144)
	m.G0 = v82
	switch int32(915) {
	case 0, 2:
		v92 = v78
		goto L17
	default:
		goto L18
	}
L16:
	;
	v120 = int32(-2)
	v122 = m.G0
	v124 = v122 - int32(144)
	m.G0 = v124
	switch int32(0) {
	case 0, 2:
		v134 = v120
		goto L30
	default:
		goto L31
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v92
	F_sigemptyset(m, v82+int32(8))
	mBase = m.M
	goto L20
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[321])) = v78
	v92 = int32(4729)
	goto L17
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+136)) = int32(268435456)
	v104 = v82 + int32(4)
	goto L24
L22:
	;
	m.G0 = v82 + int32(144)
	goto L16
L24:
	;
	goto L25
L25:
	;
	if v104 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v115 = F___memcpy(m, int32(4629548), v104, int32(140))
	mBase = m.M
	goto L28
L27:
	;
	goto L28
L28:
	;
	goto L22
L29:
	;
	v162 = int32(-2)
	v164 = m.G0
	v166 = v164 - int32(144)
	m.G0 = v166
	switch int32(0) {
	case 0, 2:
		v176 = v162
		goto L43
	default:
		goto L44
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v134
	F_sigemptyset(m, v124+int32(8))
	mBase = m.M
	goto L33
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[322])) = v120
	v134 = int32(4729)
	goto L30
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+136)) = int32(268435456)
	v146 = v124 + int32(4)
	goto L37
L35:
	;
	m.G0 = v124 + int32(144)
	goto L29
L37:
	;
	goto L38
L38:
	;
	if v146 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v157 = F___memcpy(m, int32(4629688), v146, int32(140))
	mBase = m.M
	goto L41
L40:
	;
	goto L41
L41:
	;
	goto L35
L42:
	;
	v204 = int32(-2)
	v206 = m.G0
	v208 = v206 - int32(144)
	m.G0 = v208
	switch int32(0) {
	case 0, 2:
		v218 = v204
		goto L56
	default:
		goto L57
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+4)) = v176
	F_sigemptyset(m, v166+int32(8))
	mBase = m.M
	goto L46
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[323])) = v162
	v176 = int32(4729)
	goto L43
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+136)) = int32(268435456)
	v188 = v166 + int32(4)
	goto L50
L48:
	;
	m.G0 = v166 + int32(144)
	goto L42
L50:
	;
	goto L51
L51:
	;
	if v188 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v199 = F___memcpy(m, int32(4631508), v188, int32(140))
	mBase = m.M
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L48
L55:
	;
	v246 = int32(-2)
	v248 = m.G0
	v250 = v248 - int32(144)
	m.G0 = v250
	switch int32(0) {
	case 0, 2:
		v260 = v246
		goto L69
	default:
		goto L70
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+4)) = v218
	F_sigemptyset(m, v208+int32(8))
	mBase = m.M
	goto L59
L57:
	;
	*(*int32)(unsafe.Add(mBase, _consts[324])) = v204
	v218 = int32(4729)
	goto L56
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208)+136)) = int32(268435456)
	v230 = v208 + int32(4)
	goto L63
L61:
	;
	m.G0 = v208 + int32(144)
	goto L55
L63:
	;
	goto L64
L64:
	;
	if v230 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v241 = F___memcpy(m, int32(4629828), v230, int32(140))
	mBase = m.M
	goto L67
L66:
	;
	goto L67
L67:
	;
	goto L61
L68:
	;
	v288 = int32(-2)
	v290 = m.G0
	v292 = v290 - int32(144)
	m.G0 = v292
	switch int32(0) {
	case 0, 2:
		v302 = v288
		goto L82
	default:
		goto L83
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = v260
	F_sigemptyset(m, v250+int32(8))
	mBase = m.M
	goto L72
L70:
	;
	*(*int32)(unsafe.Add(mBase, _consts[724])) = v246
	v260 = int32(4729)
	goto L69
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+136)) = int32(268435456)
	v272 = v250 + int32(4)
	goto L76
L74:
	;
	m.G0 = v250 + int32(144)
	goto L68
L76:
	;
	goto L77
L77:
	;
	if v272 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v283 = F___memcpy(m, int32(4631368), v272, int32(140))
	mBase = m.M
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L74
L81:
	;
	v330 = int32(963)
	v332 = m.G0
	v334 = v332 - int32(144)
	m.G0 = v334
	switch int32(965) {
	case 0, 2:
		v344 = v330
		goto L95
	default:
		goto L96
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+4)) = v302
	F_sigemptyset(m, v292+int32(8))
	mBase = m.M
	goto L85
L83:
	;
	*(*int32)(unsafe.Add(mBase, _consts[646])) = v288
	v302 = int32(4729)
	goto L82
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v292)+136)) = int32(268435456)
	v314 = v292 + int32(4)
	goto L89
L87:
	;
	m.G0 = v292 + int32(144)
	goto L81
L89:
	;
	goto L90
L90:
	;
	if v314 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v325 = F___memcpy(m, int32(4631228), v314, int32(140))
	mBase = m.M
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L87
L94:
	;
	v372 = int32(-2)
	v374 = m.G0
	v376 = v374 - int32(144)
	m.G0 = v376
	switch int32(0) {
	case 0, 2:
		v386 = v372
		goto L108
	default:
		goto L109
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+4)) = v344
	F_sigemptyset(m, v334+int32(8))
	mBase = m.M
	goto L98
L96:
	;
	*(*int32)(unsafe.Add(mBase, _consts[647])) = v330
	v344 = int32(4729)
	goto L95
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+136)) = int32(268435456)
	v356 = v334 + int32(4)
	goto L102
L100:
	;
	m.G0 = v334 + int32(144)
	goto L94
L102:
	;
	goto L103
L103:
	;
	if v356 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v367 = F___memcpy(m, int32(4630808), v356, int32(140))
	mBase = m.M
	goto L106
L105:
	;
	goto L106
L106:
	;
	goto L100
L107:
	;
	v414 = int32(0)
	v416 = m.G0
	v418 = v416 - int32(144)
	m.G0 = v418
	switch int32(2) {
	case 0, 2:
		v428 = v414
		goto L121
	default:
		goto L122
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376)+4)) = v386
	F_sigemptyset(m, v376+int32(8))
	mBase = m.M
	goto L111
L109:
	;
	*(*int32)(unsafe.Add(mBase, _consts[648])) = v372
	v386 = int32(4729)
	goto L108
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376)+136)) = int32(268435456)
	v398 = v376 + int32(4)
	goto L115
L113:
	;
	m.G0 = v376 + int32(144)
	goto L107
L115:
	;
	goto L116
L116:
	;
	if v398 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v409 = F___memcpy(m, int32(4631088), v398, int32(140))
	mBase = m.M
	goto L119
L118:
	;
	goto L119
L119:
	;
	goto L113
L120:
	;
	F_sigprocmask(m, int32(4371720), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L133
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v428
	F_sigemptyset(m, v418+int32(8))
	mBase = m.M
	goto L123
L122:
	;
	*(*int32)(unsafe.Add(mBase, _consts[650])) = v414
	v428 = int32(4729)
	goto L121
L123:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418)+136)) = int32(268435457)
	v440 = v418 + int32(4)
	goto L128
L126:
	;
	m.G0 = v418 + int32(144)
	goto L120
L128:
	;
	goto L129
L129:
	;
	if v440 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v451 = F___memcpy(m, int32(4631788), v440, int32(140))
	mBase = m.M
	goto L132
L131:
	;
	goto L132
L132:
	;
	goto L126
L133:
	;
	v460 = *(*int64)(unsafe.Add(mBase, _consts[786]))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v460
	v463 = F_palloc(m, int32(1024))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	v466 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v466
	v470 = F_pg_snprintf(m, v463, int32(1024), int32(542799), v26)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	if v463&int32(3) == int32(0) {
		v495 = v463
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v533 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	v537 = *(*int32)(unsafe.Add(mBase, _consts[479]))
	v538 = F_pg_localtime(m, v26+int32(48), v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L153
	}
L137:
	;
	v528 = v520 - v463
	goto L136
L138:
	;
	v499 = v495
	goto L147
L139:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463))))
	if v479 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v528 = int32(0)
	goto L136
L141:
	;
	goto L142
L142:
	;
	v484 = v463
	goto L143
L143:
	;
	v488 = v484 + int32(1)
	if v488&int32(3) == int32(0) {
		v495 = v488
		goto L138
	} else {
		goto L145
	}
L144:
	;
	v520 = v488
	goto L137
L145:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	if v493 != 0 {
		v484 = v488
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v508 = int32(-2139062144)
	if (int32(16843008)-v505|v505)&v508 == v508 {
		v499 = v499 + int32(4)
		goto L147
	} else {
		goto L149
	}
L148:
	;
	v514 = v499
	goto L150
L149:
	;
	goto L148
L150:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	if v518 != 0 {
		v514 = v514 + int32(1)
		goto L150
	} else {
		goto L152
	}
L151:
	;
	v520 = v514
	goto L137
L152:
	;
	goto L151
L153:
	;
	v540 = F_pg_strftime(m, v463+v528, int32(1024)-v528, v533, v538)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, _consts[789])) = v463
	v545 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	if v545 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v548 = *(*int64)(unsafe.Add(mBase, _consts[786]))
	v550 = F_logfile_getname(m, v548, int32(31670))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _consts[791]))
	if v554 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, _consts[792])) = v550
	goto L157
L159:
	;
	v557 = *(*int64)(unsafe.Add(mBase, _consts[786]))
	v559 = F_logfile_getname(m, v557, int32(239952))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v564 = F_pstrdup(m, v563)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L4
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, _consts[793])) = v559
	goto L161
L163:
	;
	v567 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	v568 = F_pstrdup(m, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v571 = *(*int32)(unsafe.Add(mBase, _consts[794]))
	if int32(0) < v571 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v574 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v574
	v579 = *(*int32)(unsafe.Add(mBase, _consts[479]))
	v580 = F_pg_localtime(m, v26+int32(48), v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L4
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L4
	} else {
		goto L169
	}
L168:
	;
	v583 = *(*int64)(unsafe.Add(mBase, uint32(v26)+48))
	v586 = base.I64_extend_i32_s(v571 * int32(60))
	v588 = int64(*(*int32)(unsafe.Add(mBase, uint32(v580)+36)))
	v590 = base.I64_rem_s(v583+v588, v586)
	*(*int64)(unsafe.Add(mBase, _consts[795])) = v583 + v586 - v590
	goto L167
L169:
	;
	v599 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[216])) = v599
	v603 = F_CreateWaitEventSet(m, v599, int32(2))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	v608 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	F_AddWaitEventToSet(m, v603, int32(1), int32(-1), v608)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _consts[796]))
	F_AddWaitEventToSet(m, v603, int32(2), v613, int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	v618 = int32(0)
	v628 = v571
	v630 = v564
	v631 = v568
	v640 = v39
	goto L173
L173:
	;
	v642 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	*(*int32)(unsafe.Add(mBase, uint32(v642))) = int32(0)
	goto L175
L174:
	;
	v1703 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1704 = m.ExcPending
	if v1704 != 0 {
		goto L4
	} else {
		goto L467
	}
L175:
	;
	v646 = *(*int32)(unsafe.Add(mBase, _consts[705]))
	if v646 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _consts[705])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	v805 = v628
	v806 = v630
	v807 = v631
	goto L178
L178:
	;
	v810 = int32(0)
	v812 = *(*int32)(unsafe.Add(mBase, _consts[794]))
	if v812 <= v810 {
		goto L225
	} else {
		goto L226
	}
L179:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630))))
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654))))
	if v658 == int32(0) {
		v677 = v657
		v678 = v658
		goto L181
	} else {
		goto L182
	}
L180:
	;
	if v678-v677 != 0 {
		goto L188
	} else {
		goto L189
	}
L181:
	;
	goto L180
L182:
	;
	if v657 != v658 {
		v677 = v657
		v678 = v658
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v662 = v654
	v663 = v630
	goto L184
L184:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663)+1)))
	v667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+1)))
	if v667 == int32(0) {
		v677 = v666
		v678 = v667
		goto L181
	} else {
		goto L186
	}
L185:
	;
	v677 = v666
	v678 = v667
	goto L181
L186:
	;
	v670 = int32(1)
	if v666 == v667 {
		v662 = v662 + v670
		v663 = v663 + v670
		goto L184
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	F_pfree(m, v630)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L4
	} else {
		goto L191
	}
L189:
	;
	v694 = v630
	goto L190
L190:
	;
	v696 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631))))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696))))
	if v700 == int32(0) {
		v719 = v699
		v720 = v700
		goto L195
	} else {
		goto L196
	}
L191:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v684 = F_pstrdup(m, v683)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, _consts[797])) = int32(1)
	v690 = *(*int32)(unsafe.Add(mBase, _consts[787]))
	v692 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	v693 = F_mkdir(m, v690, v692)
	mBase = m.M
	goto L193
L193:
	;
	v694 = v684
	goto L190
L194:
	;
	if v720-v719 != 0 {
		goto L202
	} else {
		goto L203
	}
L195:
	;
	goto L194
L196:
	;
	if v699 != v700 {
		v719 = v699
		v720 = v700
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v704 = v696
	v705 = v631
	goto L198
L198:
	;
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705)+1)))
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704)+1)))
	if v709 == int32(0) {
		v719 = v708
		v720 = v709
		goto L195
	} else {
		goto L200
	}
L199:
	;
	v719 = v708
	v720 = v709
	goto L195
L200:
	;
	v712 = int32(1)
	if v708 == v709 {
		v704 = v704 + v712
		v705 = v705 + v712
		goto L198
	} else {
		goto L201
	}
L201:
	;
	goto L199
L202:
	;
	F_pfree(m, v631)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L4
	} else {
		goto L205
	}
L203:
	;
	v731 = v631
	goto L204
L204:
	;
	v733 = *(*int32)(unsafe.Add(mBase, _consts[799]))
	v736 = int32(0)
	v739 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	if base.B2i32(v733&int32(8) == v736)^base.B2i32(v739 != v736) == v736 {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	v725 = *(*int32)(unsafe.Add(mBase, _consts[788]))
	v726 = F_pstrdup(m, v725)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, _consts[797])) = int32(1)
	v731 = v726
	goto L204
L207:
	;
	*(*int32)(unsafe.Add(mBase, _consts[797])) = int32(1)
	goto L209
L208:
	;
	goto L209
L209:
	;
	v750 = int32(0)
	v753 = *(*int32)(unsafe.Add(mBase, _consts[791]))
	if base.B2i32(v733&int32(16) == v750)^base.B2i32(v753 != v750) == v750 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, _consts[797])) = int32(1)
	goto L212
L211:
	;
	goto L212
L212:
	;
	v763 = *(*int32)(unsafe.Add(mBase, _consts[794]))
	if v763 != v628 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	if int32(0) < v763 {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	v790 = v628
	goto L215
L215:
	;
	v794 = int32(*(*uint8)(unsafe.Add(mBase, _consts[800])))
	if v794 != 0 {
		goto L220
	} else {
		goto L221
	}
L216:
	;
	v767 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v767
	v772 = *(*int32)(unsafe.Add(mBase, _consts[479]))
	v773 = F_pg_localtime(m, v26+int32(32), v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L4
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v790 = v763
	goto L215
L219:
	;
	v776 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v779 = base.I64_extend_i32_s(v763 * int32(60))
	v781 = int64(*(*int32)(unsafe.Add(mBase, uint32(v773)+36)))
	v783 = base.I64_rem_s(v776+v781, v779)
	*(*int64)(unsafe.Add(mBase, _consts[795])) = v776 + v779 - v783
	goto L218
L220:
	;
	*(*int32)(unsafe.Add(mBase, _consts[797])) = int32(1)
	v799 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[800])) = uint8(v799)
	goto L222
L221:
	;
	goto L222
L222:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L4
	} else {
		goto L223
	}
L223:
	;
	v805 = v790
	v806 = v694
	v807 = v731
	goto L178
L224:
	;
	v830 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	v832 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	if v832 != 0 {
		v897 = v810
		goto L230
	} else {
		goto L231
	}
L225:
	;
	v827 = int32(0)
	v828 = v640
	goto L224
L226:
	;
	goto L227
L227:
	;
	v816 = int32(0)
	v818 = int32(*(*uint8)(unsafe.Add(mBase, _consts[800])))
	if v818 != 0 {
		v827 = v816
		v828 = v640
		goto L224
	} else {
		goto L228
	}
L228:
	;
	v819 = F___time(m)
	mBase = m.M
	v821 = *(*int64)(unsafe.Add(mBase, _consts[795]))
	if v819 < v821 {
		v827 = v816
		v828 = v819
		goto L224
	} else {
		goto L229
	}
L229:
	;
	v823 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[797])) = v823
	v827 = v823
	v828 = v819
	goto L224
L230:
	;
	v899 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	if v899 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L231:
	;
	if v830 <= int32(0) {
		v897 = v810
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, _consts[800])))
	if v836 != 0 {
		v897 = v810
		goto L230
	} else {
		goto L233
	}
L233:
	;
	v838 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)+76))
	if v839 < int32(0) {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v846 = int64(*(*int32)(unsafe.Add(mBase, _consts[801])))
	if v846<<(uint(int64(10))%64) <= v844 {
		goto L238
	} else {
		goto L239
	}
L235:
	;
	v842 = F___ftello_unlocked(m, v838)
	mBase = m.M
	v844 = v842
	goto L234
L236:
	;
	goto L237
L237:
	;
	v843 = F___ftello_unlocked(m, v838)
	mBase = m.M
	v844 = v843
	goto L234
L238:
	;
	v851 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[797])) = v851
	v854 = v851
	goto L240
L239:
	;
	v854 = v810
	goto L240
L240:
	;
	v856 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	if v856 == int32(0) {
		v875 = v854
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v877 = *(*int32)(unsafe.Add(mBase, _consts[791]))
	if v877 == int32(0) {
		v897 = v875
		goto L230
	} else {
		goto L248
	}
L242:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v856)+76))
	if v859 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v866 = int64(*(*int32)(unsafe.Add(mBase, _consts[801])))
	if v864 < v866<<(uint(int64(10))%64) {
		v875 = v854
		goto L241
	} else {
		goto L247
	}
L244:
	;
	v862 = F___ftello_unlocked(m, v856)
	mBase = m.M
	v864 = v862
	goto L243
L245:
	;
	goto L246
L246:
	;
	v863 = F___ftello_unlocked(m, v856)
	mBase = m.M
	v864 = v863
	goto L243
L247:
	;
	*(*int32)(unsafe.Add(mBase, _consts[797])) = int32(1)
	v875 = v854 | int32(8)
	goto L241
L248:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v877)+76))
	if v880 < int32(0) {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	v887 = int64(*(*int32)(unsafe.Add(mBase, _consts[801])))
	if v885 < v887<<(uint(int64(10))%64) {
		v897 = v875
		goto L230
	} else {
		goto L253
	}
L250:
	;
	v883 = F___ftello_unlocked(m, v877)
	mBase = m.M
	v885 = v883
	goto L249
L251:
	;
	goto L252
L252:
	;
	v884 = F___ftello_unlocked(m, v877)
	mBase = m.M
	v885 = v884
	goto L249
L253:
	;
	*(*int32)(unsafe.Add(mBase, _consts[797])) = int32(1)
	v897 = v875 | int32(16)
	goto L230
L254:
	;
	v962 = int32(-1)
	v964 = *(*int32)(unsafe.Add(mBase, _consts[794]))
	if v964 <= int32(0) {
		v983 = v962
		goto L275
	} else {
		goto L276
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, _consts[797])) = int32(0)
	if v897 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v906 = v897
	goto L258
L257:
	;
	v906 = int32(25)
	goto L258
L258:
	;
	if v827 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v907 = v897
	goto L261
L260:
	;
	v907 = v906
	goto L261
L261:
	;
	if v827 != 0 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v915 = F_logfile_rotate_dest(m, v827, v907, v911, int32(1), int32(4373296), int32(4373336))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L266
	}
L263:
	;
	v909 = *(*int64)(unsafe.Add(mBase, _consts[795]))
	v911 = v909
	goto L262
L264:
	;
	goto L265
L265:
	;
	v910 = F___time(m)
	mBase = m.M
	v911 = v910
	goto L262
L266:
	;
	if v915 == int32(0) {
		goto L254
	} else {
		goto L267
	}
L267:
	;
	v922 = F_logfile_rotate_dest(m, v827, v907, v911, int32(8), int32(4373304), int32(4373300))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	if v922 == int32(0) {
		goto L254
	} else {
		goto L269
	}
L269:
	;
	v929 = F_logfile_rotate_dest(m, v827, v907, v911, int32(16), int32(4373312), int32(4373308))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	if v929 == int32(0) {
		goto L254
	} else {
		goto L271
	}
L271:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	v936 = *(*int32)(unsafe.Add(mBase, _consts[794]))
	if v936 <= int32(0) {
		goto L254
	} else {
		goto L273
	}
L273:
	;
	v939 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v939
	v944 = *(*int32)(unsafe.Add(mBase, _consts[479]))
	v945 = F_pg_localtime(m, v26+int32(32), v944)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	v948 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v951 = base.I64_extend_i32_s(v936 * int32(60))
	v953 = int64(*(*int32)(unsafe.Add(mBase, uint32(v945)+36)))
	v955 = base.I64_rem_s(v948+v953, v951)
	*(*int64)(unsafe.Add(mBase, _consts[795])) = v948 + v951 - v955
	goto L254
L275:
	;
	v989 = F_WaitEventSetWait(m, v603, v983, v26+int32(32), int32(1), int32(83886093))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L4
	} else {
		goto L285
	}
L276:
	;
	v968 = int32(*(*uint8)(unsafe.Add(mBase, _consts[800])))
	if v968 != 0 {
		v983 = v962
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v969 = int64(2147483)
	v971 = *(*int64)(unsafe.Add(mBase, _consts[795]))
	v972 = v971 - v828
	if v969 <= v972 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v975 = v969
	goto L280
L279:
	;
	v975 = v972
	goto L280
L280:
	;
	if int64(0) < v972 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v982 = base.I32_wrap_i64(v975) * int32(1000)
	goto L283
L282:
	;
	v982 = int32(0)
	goto L283
L283:
	;
	v983 = v982
	goto L275
L284:
	;
	v1698 = int32(*(*uint8)(unsafe.Add(mBase, _consts[803])))
	if v1698 == int32(0) {
		v618 = v1674
		v628 = v805
		v630 = v806
		v631 = v807
		v640 = v828
		goto L173
	} else {
		goto L466
	}
L285:
	;
	if v989 != int32(1) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1674 = v618
	goto L284
L287:
	;
	goto L288
L288:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v993 != int32(2) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1674 = v618
	goto L284
L290:
	;
	goto L291
L291:
	;
	v997 = *(*int32)(unsafe.Add(mBase, _consts[796]))
	v1003 = F_read(m, v997, v26+int32(48)+v618, int32(8192)-v618)
	mBase = m.M
	if v1003 < int32(0) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v1007 == int32(27) {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	goto L294
L294:
	;
	if v1003 != 0 {
		goto L305
	} else {
		goto L306
	}
L295:
	;
	v1674 = v618
	goto L284
L296:
	;
	goto L297
L297:
	;
	v1012 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L4
	} else {
		goto L298
	}
L298:
	;
	if v1012 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1674 = v618
	goto L284
L300:
	;
	goto L301
L301:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L4
	} else {
		goto L302
	}
L302:
	;
	F_errmsg(m, int32(288227), int32(0))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L4
	} else {
		goto L303
	}
L303:
	;
	F_errfinish(m, int32(484362), int32(527), int32(272877))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L4
	} else {
		goto L304
	}
L304:
	;
	v1674 = v618
	goto L284
L305:
	;
	v1027 = v618 + v1003
	if v1027 < int32(10) {
		v618 = v1027
		v628 = v805
		v630 = v806
		v631 = v807
		v640 = v828
		goto L173
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	v1545 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[803])) = uint8(v1545)
	v1553 = int32(0)
	goto L444
L308:
	;
	v1033 = v1027
	v1036 = v26 + int32(48)
	v1040 = int32(1)
	goto L309
L309:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1036))))
	if v1056 != 0 {
		goto L314
	} else {
		goto L315
	}
L310:
	;
	if v1369 <= int32(0) {
		v618 = v1369
		v628 = v805
		v630 = v806
		v631 = v807
		v640 = v828
		goto L173
	} else {
		goto L396
	}
L311:
	;
	goto L310
L312:
	;
	v1365 = v1344 + v1036
	v1366 = v1033 - v1344
	if int32(9) < v1366 {
		v1033 = v1366
		v1036 = v1365
		v1040 = v1349
		goto L309
	} else {
		goto L395
	}
L313:
	;
	F_write_stderr(m, int32(719722), int32(0))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L4
	} else {
		goto L394
	}
L314:
	;
	v1280 = int32(1)
	goto L388
L315:
	;
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1036)+1)))
	if v1057 != 0 {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v1058 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1036)+2)))
	if base.Ui32(int32(4086)) < base.Ui32((v1058-int32(1))&int32(65535)) {
		goto L314
	} else {
		goto L317
	}
L317:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+4))
	if v1065 == int32(0) {
		goto L314
	} else {
		goto L318
	}
L318:
	;
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1036)+8)))
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068&int32(112))+uint32(_consts[804]))))
	if v1073 != int32(1) {
		goto L314
	} else {
		goto L319
	}
L319:
	;
	v1077 = v1058 + int32(9)
	if base.Ui32(v1033) < base.Ui32(v1077) {
		v1369 = v1033
		v1372 = v1036
		goto L311
	} else {
		goto L320
	}
L320:
	;
	if v1068&int32(16) != 0 {
		v1089 = int32(1)
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1090 = int32(0)
	v1092 = base.I32_rem_s(v1065, int32(256))
	v1094 = v1092 << (uint(int32(2)) % 32)
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+uint32(_consts[805])))
	if v1097 != 0 {
		goto L328
	} else {
		goto L329
	}
L322:
	;
	if v1068&int32(32) != 0 {
		v1089 = int32(8)
		goto L321
	} else {
		goto L323
	}
L323:
	;
	if v1068&int32(64) != 0 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1088 = int32(16)
	goto L326
L325:
	;
	v1088 = v1040
	goto L326
L326:
	;
	v1089 = v1088
	goto L321
L327:
	;
	if v1068&int32(1) == int32(0) {
		goto L339
	} else {
		goto L340
	}
L328:
	;
	v1098 = int32(0)
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	if v1099 <= v1098 {
		v1163 = v1098
		v1167 = v1090
		goto L327
	} else {
		goto L331
	}
L329:
	;
	v1143 = v1090
	goto L330
L330:
	;
	v1163 = int32(0)
	v1167 = v1143
	goto L327
L331:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+12))
	v1109 = v1090
	v1110 = int32(0)
	goto L332
L332:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1102+v1110<<(uint(int32(2))%32))))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1130)))
	if v1131 == v1065 {
		v1163 = v1130
		v1167 = v1109
		goto L327
	} else {
		goto L334
	}
L333:
	;
	v1143 = v1134
	goto L330
L334:
	;
	if v1109|v1131 != 0 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1134 = v1109
	goto L337
L336:
	;
	v1134 = v1130
	goto L337
L337:
	;
	v1136 = v1110 + int32(1)
	if v1099 != v1136 {
		v1109 = v1134
		v1110 = v1136
		goto L332
	} else {
		goto L338
	}
L338:
	;
	goto L333
L339:
	;
	if v1163 != 0 {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	goto L341
L341:
	;
	if v1163 != 0 {
		goto L353
	} else {
		goto L354
	}
L342:
	;
	F_appendBinaryStringInfo(m, v1163+int32(4), v1036+int32(9), v1058)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L4
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	if v1167 == int32(0) {
		goto L346
	} else {
		goto L347
	}
L345:
	;
	v1344 = v1077
	v1349 = v1089
	goto L312
L346:
	;
	v1198 = F_palloc(m, int32(20))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L4
	} else {
		goto L349
	}
L347:
	;
	v1203 = v1167
	goto L348
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1203))) = v1065
	v1206 = v1203 + int32(4)
	F_initStringInfo(m, v1206)
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L4
	} else {
		goto L351
	}
L349:
	;
	v1200 = F_lappend(m, v1097, v1198)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L4
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1094)+uint32(_consts[805]))) = v1200
	v1203 = v1198
	goto L348
L351:
	;
	F_appendBinaryStringInfo(m, v1206, v1036+int32(9), v1058)
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L4
	} else {
		goto L352
	}
L352:
	;
	v1344 = v1077
	v1349 = v1089
	goto L312
L353:
	;
	F_appendBinaryStringInfo(m, v1163+int32(4), v1036+int32(9), v1058)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L4
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	if v1089&int32(8) != 0 {
		goto L375
	} else {
		goto L376
	}
L356:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+8))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+4))
	if v1089&int32(8) != 0 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v1239 = F_fwrite(m, v1220, int32(1), v1219, v1236)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L4
	} else {
		goto L368
	}
L358:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	if v1224 != 0 {
		v1236 = v1224
		goto L357
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, _consts[791]))
	v1229 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	if v1227 != 0 {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	goto L360
L362:
	;
	v1230 = v1227
	goto L364
L363:
	;
	v1230 = v1229
	goto L364
L364:
	;
	if int32(base.Ui32(v1089&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1235 = v1230
	goto L367
L366:
	;
	v1235 = v1229
	goto L367
L367:
	;
	v1236 = v1235
	goto L357
L368:
	;
	if v1239 != v1219 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	F_write_stderr(m, int32(719722), int32(0))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L4
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1163))) = int32(0)
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+4))
	F_pfree(m, v1248)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L4
	} else {
		goto L373
	}
L372:
	;
	goto L371
L373:
	;
	v1344 = v1077
	v1349 = v1089
	goto L312
L374:
	;
	v1271 = F_fwrite(m, v1036+int32(9), int32(1), v1058, v1266)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L4
	} else {
		goto L385
	}
L375:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, _consts[790]))
	if v1254 != 0 {
		v1266 = v1254
		goto L374
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, _consts[791]))
	v1259 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	if v1257 != 0 {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	goto L377
L379:
	;
	v1260 = v1257
	goto L381
L380:
	;
	v1260 = v1259
	goto L381
L381:
	;
	if int32(base.Ui32(v1089&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1265 = v1260
	goto L384
L383:
	;
	v1265 = v1259
	goto L384
L384:
	;
	v1266 = v1265
	goto L374
L385:
	;
	if v1271 == v1058 {
		v1344 = v1077
		v1349 = v1089
		goto L312
	} else {
		goto L386
	}
L386:
	;
	v1317 = v1077
	v1322 = v1089
	goto L313
L387:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v1312 = F_fwrite(m, v1036, int32(1), v1308, v1311)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L4
	} else {
		goto L392
	}
L388:
	;
	v1302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280+v1036))))
	if v1302 == int32(0) {
		v1308 = v1280
		goto L387
	} else {
		goto L390
	}
L389:
	;
	v1308 = v1033
	goto L387
L390:
	;
	v1306 = v1280 + int32(1)
	if v1306 != v1033 {
		v1280 = v1306
		goto L388
	} else {
		goto L391
	}
L391:
	;
	goto L389
L392:
	;
	if v1312 == v1308 {
		v1344 = v1308
		v1349 = v1040
		goto L312
	} else {
		goto L393
	}
L393:
	;
	v1317 = v1308
	v1322 = v1040
	goto L313
L394:
	;
	v1344 = v1317
	v1349 = v1322
	goto L312
L395:
	;
	v1369 = v1366
	v1372 = v1365
	goto L311
L396:
	;
	if v1372 == v26+int32(48) {
		v618 = v1369
		v628 = v805
		v630 = v806
		v631 = v807
		v640 = v828
		goto L173
	} else {
		goto L397
	}
L397:
	;
	v1398 = v26 + int32(48)
	if v1398 == v1372 {
		goto L399
	} else {
		goto L400
	}
L398:
	;
	v618 = v1369
	v628 = v805
	v630 = v806
	v631 = v807
	v640 = v828
	goto L173
L399:
	;
	goto L398
L400:
	;
	v1402 = v1398 + v1369
	if base.Ui32(v1372-v1402) <= base.Ui32(int32(0)-v1369<<(uint(int32(1))%32)) {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1409 = F___memcpy(m, v1398, v1372, v1369)
	mBase = m.M
	goto L398
L402:
	;
	goto L403
L403:
	;
	v1412 = (v1398 ^ v1372) & int32(3)
	if base.Ui32(v1398) < base.Ui32(v1372) {
		goto L406
	} else {
		goto L407
	}
L404:
	;
	if v1514 == int32(0) {
		goto L399
	} else {
		goto L440
	}
L405:
	;
	if base.Ui32(v1492) <= base.Ui32(int32(3)) {
		v1513 = v1491
		v1514 = v1492
		v1515 = v1493
		goto L404
	} else {
		goto L436
	}
L406:
	;
	if v1412 != 0 {
		goto L409
	} else {
		goto L410
	}
L407:
	;
	goto L408
L408:
	;
	if v1412 != 0 {
		v1474 = v1369
		goto L419
	} else {
		goto L420
	}
L409:
	;
	v1513 = v1372
	v1514 = v1369
	v1515 = v1398
	goto L404
L410:
	;
	goto L411
L411:
	;
	if v1398&int32(3) == int32(0) {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1491 = v1372
	v1492 = v1369
	v1493 = v1398
	goto L405
L413:
	;
	goto L414
L414:
	;
	v1419 = v1372
	v1420 = v1369
	v1421 = v1398
	goto L415
L415:
	;
	if v1420 == int32(0) {
		goto L399
	} else {
		goto L417
	}
L416:
	;
	v1491 = v1428
	v1492 = v1430
	v1493 = v1432
	goto L405
L417:
	;
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1419))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1421))) = uint8(v1425)
	v1427 = int32(1)
	v1428 = v1419 + v1427
	v1430 = v1420 - v1427
	v1432 = v1421 + v1427
	if v1432&int32(3) != 0 {
		v1419 = v1428
		v1420 = v1430
		v1421 = v1432
		goto L415
	} else {
		goto L418
	}
L418:
	;
	goto L416
L419:
	;
	if v1474 == int32(0) {
		goto L399
	} else {
		goto L432
	}
L420:
	;
	if v1402&int32(3) != 0 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1439 = v1369
	goto L424
L422:
	;
	v1454 = v1369
	goto L423
L423:
	;
	if base.Ui32(v1454) <= base.Ui32(int32(3)) {
		v1474 = v1454
		goto L419
	} else {
		goto L428
	}
L424:
	;
	if v1439 == int32(0) {
		goto L399
	} else {
		goto L426
	}
L425:
	;
	v1454 = v1445
	goto L423
L426:
	;
	v1445 = v1439 - int32(1)
	v1446 = v1398 + v1445
	v1448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1372+v1445))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1446))) = uint8(v1448)
	if v1446&int32(3) != 0 {
		v1439 = v1445
		goto L424
	} else {
		goto L427
	}
L427:
	;
	goto L425
L428:
	;
	v1461 = v1454
	goto L429
L429:
	;
	v1465 = v1461 - int32(4)
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1372+v1465)))
	*(*int32)(unsafe.Add(mBase, uint32(v1398+v1465))) = v1468
	if base.Ui32(int32(3)) < base.Ui32(v1465) {
		v1461 = v1465
		goto L429
	} else {
		goto L431
	}
L430:
	;
	v1474 = v1465
	goto L419
L431:
	;
	goto L430
L432:
	;
	v1481 = v1474
	goto L433
L433:
	;
	v1485 = v1481 - int32(1)
	v1488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1372+v1485))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1398+v1485))) = uint8(v1488)
	if v1485 != 0 {
		v1481 = v1485
		goto L433
	} else {
		goto L435
	}
L434:
	;
	goto L399
L435:
	;
	goto L434
L436:
	;
	v1498 = v1491
	v1499 = v1492
	v1500 = v1493
	goto L437
L437:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1498)))
	*(*int32)(unsafe.Add(mBase, uint32(v1500))) = v1502
	v1504 = int32(4)
	v1505 = v1498 + v1504
	v1507 = v1500 + v1504
	v1509 = v1499 - v1504
	if base.Ui32(int32(3)) < base.Ui32(v1509) {
		v1498 = v1505
		v1499 = v1509
		v1500 = v1507
		goto L437
	} else {
		goto L439
	}
L438:
	;
	v1513 = v1505
	v1514 = v1509
	v1515 = v1507
	goto L404
L439:
	;
	goto L438
L440:
	;
	v1520 = v1513
	v1521 = v1514
	v1522 = v1515
	goto L441
L441:
	;
	v1524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1520))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1522))) = uint8(v1524)
	v1526 = int32(1)
	v1531 = v1521 - v1526
	if v1531 != 0 {
		v1520 = v1520 + v1526
		v1521 = v1531
		v1522 = v1522 + v1526
		goto L441
	} else {
		goto L443
	}
L442:
	;
	goto L399
L443:
	;
	goto L442
L444:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1553<<(uint(int32(2))%32))+uint32(_consts[805])))
	if v1574 == int32(0) {
		goto L446
	} else {
		goto L447
	}
L445:
	;
	v1659 = int32(0)
	if v618 <= v1659 {
		v1674 = v1659
		goto L284
	} else {
		goto L462
	}
L446:
	;
	v1656 = v1553 + int32(1)
	if v1656 != int32(256) {
		v1553 = v1656
		goto L444
	} else {
		goto L461
	}
L447:
	;
	v1577 = int32(0)
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+4))
	if v1578 <= v1577 {
		goto L446
	} else {
		goto L448
	}
L448:
	;
	v1581 = v1578
	v1584 = v1577
	goto L449
L449:
	;
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+12))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1604+v1584<<(uint(int32(2))%32))))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1608)))
	if v1609 != 0 {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	goto L446
L451:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+4))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+8))
	v1614 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v1615 = F_fwrite(m, v1610, int32(1), v1612, v1614)
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L4
	} else {
		goto L454
	}
L452:
	;
	v1628 = v1581
	goto L453
L453:
	;
	v1630 = v1584 + int32(1)
	if v1630 < v1628 {
		v1581 = v1628
		v1584 = v1630
		goto L449
	} else {
		goto L460
	}
L454:
	;
	if v1615 != v1612 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	F_write_stderr(m, int32(719722), int32(0))
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L4
	} else {
		goto L458
	}
L456:
	;
	goto L457
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1608))) = int32(0)
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+4))
	F_pfree(m, v1624)
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L4
	} else {
		goto L459
	}
L458:
	;
	goto L457
L459:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+4))
	v1628 = v1627
	goto L453
L460:
	;
	goto L450
L461:
	;
	goto L445
L462:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v1667 = F_fwrite(m, v26+int32(48), int32(1), v618, v1666)
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L4
	} else {
		goto L463
	}
L463:
	;
	if v1667 == v618 {
		v1674 = v1659
		goto L284
	} else {
		goto L464
	}
L464:
	;
	F_write_stderr(m, int32(719722), int32(0))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L4
	} else {
		goto L465
	}
L465:
	;
	v1674 = v1659
	goto L284
L466:
	;
	goto L174
L467:
	;
	if v1703 != 0 {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	F_errmsg_internal(m, int32(239008), int32(0))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L4
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L4
	} else {
		goto L473
	}
L471:
	;
	F_errfinish(m, int32(484362), int32(575), int32(272877))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L4
	} else {
		goto L472
	}
L472:
	;
	goto L470
L473:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
