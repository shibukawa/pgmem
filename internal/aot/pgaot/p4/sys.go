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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1393]))
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
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(5))%32))+uint32(_consts[1394])))
				v28 = F_get_rel_name(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1395])))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v35 + v36<<(uint(int32(4))%32) + l2*int32(100) - int32(76)
					F_errmsg_internal(m, int32(184540), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(499338), int32(644), int32(303910))
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
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int64
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int64
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int64
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int64
	_ = v527
	var v530 int64
	_ = v530
	var v532 int64
	_ = v532
	var v534 int64
	_ = v534
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v584 int64
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v707 int32
	_ = v707
	var v711 int64
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int64
	_ = v720
	var v723 int64
	_ = v723
	var v725 int64
	_ = v725
	var v727 int64
	_ = v727
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int64
	_ = v763
	var v765 int64
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int64
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int64
	_ = v786
	var v787 int64
	_ = v787
	var v788 int64
	_ = v788
	var v790 int64
	_ = v790
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int64
	_ = v806
	var v807 int64
	_ = v807
	var v808 int64
	_ = v808
	var v810 int64
	_ = v810
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v827 int64
	_ = v827
	var v828 int64
	_ = v828
	var v829 int64
	_ = v829
	var v831 int64
	_ = v831
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int64
	_ = v853
	var v854 int64
	_ = v854
	var v855 int64
	_ = v855
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v883 int64
	_ = v883
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int64
	_ = v892
	var v895 int64
	_ = v895
	var v897 int64
	_ = v897
	var v899 int64
	_ = v899
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v913 int64
	_ = v913
	var v915 int64
	_ = v915
	var v916 int64
	_ = v916
	var v919 int64
	_ = v919
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1087 int32
	_ = v1087
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1224 int32
	_ = v1224
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1383 int32
	_ = v1383
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1398 int32
	_ = v1398
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1418 int32
	_ = v1418
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1489 int32
	_ = v1489
	var v1497 int32
	_ = v1497
	var v1518 int32
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	v24 = m.G0
	v26 = v24 - int32(8240)
	m.G0 = v26
	v29 = *(*int32)(unsafe.Add(mBase, _consts[581]))
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
	*(*int32)(unsafe.Add(mBase, _consts[273])) = int32(17)
	v39 = *(*int64)(unsafe.Add(mBase, _consts[793]))
	goto L7
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[581])) = int32(0)
	goto L3
L6:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[794])))
	if v47 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	v45 = F_GetBackendTypeDesc(m, v44)
	mBase = m.M
	goto L9
L9:
	;
	goto L6
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[795]))
	if int32(0) <= v70 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(0)
	v53 = int32(1)
	v56 = F_open(m, int32(302220), v53, v26+int32(16))
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
	*(*int32)(unsafe.Add(mBase, _consts[795])) = int32(-1)
	v78 = int32(914)
	v80 = m.G0
	v82 = v80 - int32(144)
	m.G0 = v82
	switch int32(916) {
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
	*(*int32)(unsafe.Add(mBase, _consts[330])) = v78
	v92 = int32(4730)
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
	v115 = F___memcpy(m, int32(4681100), v104, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[331])) = v120
	v134 = int32(4730)
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
	v157 = F___memcpy(m, int32(4681240), v146, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[332])) = v162
	v176 = int32(4730)
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
	v199 = F___memcpy(m, int32(4683060), v188, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[333])) = v204
	v218 = int32(4730)
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
	v241 = F___memcpy(m, int32(4681380), v230, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[734])) = v246
	v260 = int32(4730)
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
	v283 = F___memcpy(m, int32(4682920), v272, int32(140))
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
	v330 = int32(964)
	v332 = m.G0
	v334 = v332 - int32(144)
	m.G0 = v334
	switch int32(966) {
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
	*(*int32)(unsafe.Add(mBase, _consts[656])) = v288
	v302 = int32(4730)
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
	v325 = F___memcpy(m, int32(4682780), v314, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[657])) = v330
	v344 = int32(4730)
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
	v367 = F___memcpy(m, int32(4682360), v356, int32(140))
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
	*(*int32)(unsafe.Add(mBase, _consts[658])) = v372
	v386 = int32(4730)
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
	v409 = F___memcpy(m, int32(4682640), v398, int32(140))
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
	F_sigprocmask(m, int32(4422712), int32(0))
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
	*(*int32)(unsafe.Add(mBase, _consts[660])) = v414
	v428 = int32(4730)
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
	v451 = F___memcpy(m, int32(4683340), v440, int32(140))
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
	v460 = *(*int64)(unsafe.Add(mBase, _consts[796]))
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
	v466 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v466
	v470 = F_pg_snprintf(m, v463, int32(1024), int32(570764), v26)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v472 = F_strlen(m, v463)
	mBase = m.M
	v477 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	v481 = *(*int32)(unsafe.Add(mBase, _consts[489]))
	v482 = F_pg_localtime(m, v26+int32(48), v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v484 = F_pg_strftime(m, v463+v472, int32(1024)-v472, v477, v482)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, _consts[799])) = v463
	v489 = *(*int32)(unsafe.Add(mBase, _consts[800]))
	if v489 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v492 = *(*int64)(unsafe.Add(mBase, _consts[796]))
	v494 = F_logfile_getname(m, v492, int32(32616))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	if v498 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _consts[802])) = v494
	goto L140
L142:
	;
	v501 = *(*int64)(unsafe.Add(mBase, _consts[796]))
	v503 = F_logfile_getname(m, v501, int32(245446))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	v508 = F_pstrdup(m, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L146
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _consts[803])) = v503
	goto L144
L146:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	v512 = F_pstrdup(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	if int32(0) < v515 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v518 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v518
	v523 = *(*int32)(unsafe.Add(mBase, _consts[489]))
	v524 = F_pg_localtime(m, v26+int32(48), v523)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L4
	} else {
		goto L152
	}
L151:
	;
	v527 = *(*int64)(unsafe.Add(mBase, uint32(v26)+48))
	v530 = base.I64_extend_i32_s(v515 * int32(60))
	v532 = int64(*(*int32)(unsafe.Add(mBase, uint32(v524)+36)))
	v534 = base.I64_rem_s(v527+v532, v530)
	*(*int64)(unsafe.Add(mBase, _consts[805])) = v527 + v530 - v534
	goto L150
L152:
	;
	v543 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[225])) = v543
	v547 = F_CreateWaitEventSet(m, v543, int32(2))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	v552 = *(*int32)(unsafe.Add(mBase, _consts[506]))
	F_AddWaitEventToSet(m, v547, int32(1), int32(-1), v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v557 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	F_AddWaitEventToSet(m, v547, int32(2), v557, int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v562 = int32(0)
	v572 = v515
	v574 = v508
	v575 = v512
	v584 = v39
	goto L156
L156:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _consts[506]))
	*(*int32)(unsafe.Add(mBase, uint32(v586))) = int32(0)
	goto L158
L157:
	;
	v1647 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1648 = m.ExcPending
	if v1648 != 0 {
		goto L4
	} else {
		goto L450
	}
L158:
	;
	v590 = *(*int32)(unsafe.Add(mBase, _consts[715]))
	if v590 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, _consts[715])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L162
	}
L160:
	;
	v749 = v572
	v750 = v574
	v751 = v575
	goto L161
L161:
	;
	v754 = int32(0)
	v756 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	if v756 <= v754 {
		goto L208
	} else {
		goto L209
	}
L162:
	;
	v598 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598))))
	if v602 == int32(0) {
		v621 = v601
		v622 = v602
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v622-v621 != 0 {
		goto L171
	} else {
		goto L172
	}
L164:
	;
	goto L163
L165:
	;
	if v601 != v602 {
		v621 = v601
		v622 = v602
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v606 = v598
	v607 = v574
	goto L167
L167:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607)+1)))
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606)+1)))
	if v611 == int32(0) {
		v621 = v610
		v622 = v611
		goto L164
	} else {
		goto L169
	}
L168:
	;
	v621 = v610
	v622 = v611
	goto L164
L169:
	;
	v614 = int32(1)
	if v610 == v611 {
		v606 = v606 + v614
		v607 = v607 + v614
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	F_pfree(m, v574)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L4
	} else {
		goto L174
	}
L172:
	;
	v638 = v574
	goto L173
L173:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575))))
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	if v644 == int32(0) {
		v663 = v643
		v664 = v644
		goto L178
	} else {
		goto L179
	}
L174:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	v628 = F_pstrdup(m, v627)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, _consts[807])) = int32(1)
	v634 = *(*int32)(unsafe.Add(mBase, _consts[797]))
	v636 = *(*int32)(unsafe.Add(mBase, _consts[808]))
	v637 = F_mkdir(m, v634, v636)
	mBase = m.M
	goto L176
L176:
	;
	v638 = v628
	goto L173
L177:
	;
	if v664-v663 != 0 {
		goto L185
	} else {
		goto L186
	}
L178:
	;
	goto L177
L179:
	;
	if v643 != v644 {
		v663 = v643
		v664 = v644
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v648 = v640
	v649 = v575
	goto L181
L181:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649)+1)))
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v648)+1)))
	if v653 == int32(0) {
		v663 = v652
		v664 = v653
		goto L178
	} else {
		goto L183
	}
L182:
	;
	v663 = v652
	v664 = v653
	goto L178
L183:
	;
	v656 = int32(1)
	if v652 == v653 {
		v648 = v648 + v656
		v649 = v649 + v656
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	F_pfree(m, v575)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L4
	} else {
		goto L188
	}
L186:
	;
	v675 = v575
	goto L187
L187:
	;
	v677 = *(*int32)(unsafe.Add(mBase, _consts[809]))
	v680 = int32(0)
	v683 = *(*int32)(unsafe.Add(mBase, _consts[800]))
	if base.B2i32(v677&int32(8) == v680)^base.B2i32(v683 != v680) == v680 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _consts[798]))
	v670 = F_pstrdup(m, v669)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, _consts[807])) = int32(1)
	v675 = v670
	goto L187
L190:
	;
	*(*int32)(unsafe.Add(mBase, _consts[807])) = int32(1)
	goto L192
L191:
	;
	goto L192
L192:
	;
	v694 = int32(0)
	v697 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	if base.B2i32(v677&int32(16) == v694)^base.B2i32(v697 != v694) == v694 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, _consts[807])) = int32(1)
	goto L195
L194:
	;
	goto L195
L195:
	;
	v707 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	if v707 != v572 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	if int32(0) < v707 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v734 = v572
	goto L198
L198:
	;
	v738 = int32(*(*uint8)(unsafe.Add(mBase, _consts[810])))
	if v738 != 0 {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	v711 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v711
	v716 = *(*int32)(unsafe.Add(mBase, _consts[489]))
	v717 = F_pg_localtime(m, v26+int32(32), v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L4
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v734 = v707
	goto L198
L202:
	;
	v720 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v723 = base.I64_extend_i32_s(v707 * int32(60))
	v725 = int64(*(*int32)(unsafe.Add(mBase, uint32(v717)+36)))
	v727 = base.I64_rem_s(v720+v725, v723)
	*(*int64)(unsafe.Add(mBase, _consts[805])) = v720 + v723 - v727
	goto L201
L203:
	;
	*(*int32)(unsafe.Add(mBase, _consts[807])) = int32(1)
	v743 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[810])) = uint8(v743)
	goto L205
L204:
	;
	goto L205
L205:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	v749 = v734
	v750 = v638
	v751 = v675
	goto L161
L207:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _consts[811]))
	v776 = *(*int32)(unsafe.Add(mBase, _consts[807]))
	if v776 != 0 {
		v841 = v754
		goto L213
	} else {
		goto L214
	}
L208:
	;
	v771 = int32(0)
	v772 = v584
	goto L207
L209:
	;
	goto L210
L210:
	;
	v760 = int32(0)
	v762 = int32(*(*uint8)(unsafe.Add(mBase, _consts[810])))
	if v762 != 0 {
		v771 = v760
		v772 = v584
		goto L207
	} else {
		goto L211
	}
L211:
	;
	v763 = F___time(m)
	mBase = m.M
	v765 = *(*int64)(unsafe.Add(mBase, _consts[805]))
	if v763 < v765 {
		v771 = v760
		v772 = v763
		goto L207
	} else {
		goto L212
	}
L212:
	;
	v767 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[807])) = v767
	v771 = v767
	v772 = v763
	goto L207
L213:
	;
	v843 = *(*int32)(unsafe.Add(mBase, _consts[807]))
	if v843 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L214:
	;
	if v774 <= int32(0) {
		v841 = v754
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, _consts[810])))
	if v780 != 0 {
		v841 = v754
		goto L213
	} else {
		goto L216
	}
L216:
	;
	v782 = *(*int32)(unsafe.Add(mBase, _consts[812]))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v782)+76))
	if v783 < int32(0) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v790 = int64(*(*int32)(unsafe.Add(mBase, _consts[811])))
	if v790<<(uint(int64(10))%64) <= v788 {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	v786 = F___ftello_unlocked(m, v782)
	mBase = m.M
	v788 = v786
	goto L217
L219:
	;
	goto L220
L220:
	;
	v787 = F___ftello_unlocked(m, v782)
	mBase = m.M
	v788 = v787
	goto L217
L221:
	;
	v795 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[807])) = v795
	v798 = v795
	goto L223
L222:
	;
	v798 = v754
	goto L223
L223:
	;
	v800 = *(*int32)(unsafe.Add(mBase, _consts[800]))
	if v800 == int32(0) {
		v819 = v798
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v821 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	if v821 == int32(0) {
		v841 = v819
		goto L213
	} else {
		goto L231
	}
L225:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v800)+76))
	if v803 < int32(0) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v810 = int64(*(*int32)(unsafe.Add(mBase, _consts[811])))
	if v808 < v810<<(uint(int64(10))%64) {
		v819 = v798
		goto L224
	} else {
		goto L230
	}
L227:
	;
	v806 = F___ftello_unlocked(m, v800)
	mBase = m.M
	v808 = v806
	goto L226
L228:
	;
	goto L229
L229:
	;
	v807 = F___ftello_unlocked(m, v800)
	mBase = m.M
	v808 = v807
	goto L226
L230:
	;
	*(*int32)(unsafe.Add(mBase, _consts[807])) = int32(1)
	v819 = v798 | int32(8)
	goto L224
L231:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v821)+76))
	if v824 < int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	v831 = int64(*(*int32)(unsafe.Add(mBase, _consts[811])))
	if v829 < v831<<(uint(int64(10))%64) {
		v841 = v819
		goto L213
	} else {
		goto L236
	}
L233:
	;
	v827 = F___ftello_unlocked(m, v821)
	mBase = m.M
	v829 = v827
	goto L232
L234:
	;
	goto L235
L235:
	;
	v828 = F___ftello_unlocked(m, v821)
	mBase = m.M
	v829 = v828
	goto L232
L236:
	;
	*(*int32)(unsafe.Add(mBase, _consts[807])) = int32(1)
	v841 = v819 | int32(16)
	goto L213
L237:
	;
	v906 = int32(-1)
	v908 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	if v908 <= int32(0) {
		v927 = v906
		goto L258
	} else {
		goto L259
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, _consts[807])) = int32(0)
	if v841 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v850 = v841
	goto L241
L240:
	;
	v850 = int32(25)
	goto L241
L241:
	;
	if v771 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v851 = v841
	goto L244
L243:
	;
	v851 = v850
	goto L244
L244:
	;
	if v771 != 0 {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v859 = F_logfile_rotate_dest(m, v771, v851, v855, int32(1), int32(4424288), int32(4424328))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L4
	} else {
		goto L249
	}
L246:
	;
	v853 = *(*int64)(unsafe.Add(mBase, _consts[805]))
	v855 = v853
	goto L245
L247:
	;
	goto L248
L248:
	;
	v854 = F___time(m)
	mBase = m.M
	v855 = v854
	goto L245
L249:
	;
	if v859 == int32(0) {
		goto L237
	} else {
		goto L250
	}
L250:
	;
	v866 = F_logfile_rotate_dest(m, v771, v851, v855, int32(8), int32(4424296), int32(4424292))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L4
	} else {
		goto L251
	}
L251:
	;
	if v866 == int32(0) {
		goto L237
	} else {
		goto L252
	}
L252:
	;
	v873 = F_logfile_rotate_dest(m, v771, v851, v855, int32(16), int32(4424304), int32(4424300))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L4
	} else {
		goto L253
	}
L253:
	;
	if v873 == int32(0) {
		goto L237
	} else {
		goto L254
	}
L254:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	v880 = *(*int32)(unsafe.Add(mBase, _consts[804]))
	if v880 <= int32(0) {
		goto L237
	} else {
		goto L256
	}
L256:
	;
	v883 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v883
	v888 = *(*int32)(unsafe.Add(mBase, _consts[489]))
	v889 = F_pg_localtime(m, v26+int32(32), v888)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	v892 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v895 = base.I64_extend_i32_s(v880 * int32(60))
	v897 = int64(*(*int32)(unsafe.Add(mBase, uint32(v889)+36)))
	v899 = base.I64_rem_s(v892+v897, v895)
	*(*int64)(unsafe.Add(mBase, _consts[805])) = v892 + v895 - v899
	goto L237
L258:
	;
	v933 = F_WaitEventSetWait(m, v547, v927, v26+int32(32), int32(1), int32(83886093))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L4
	} else {
		goto L268
	}
L259:
	;
	v912 = int32(*(*uint8)(unsafe.Add(mBase, _consts[810])))
	if v912 != 0 {
		v927 = v906
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v913 = int64(2147483)
	v915 = *(*int64)(unsafe.Add(mBase, _consts[805]))
	v916 = v915 - v772
	if v913 <= v916 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v919 = v913
	goto L263
L262:
	;
	v919 = v916
	goto L263
L263:
	;
	if int64(0) < v916 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v926 = base.I32_wrap_i64(v919) * int32(1000)
	goto L266
L265:
	;
	v926 = int32(0)
	goto L266
L266:
	;
	v927 = v926
	goto L258
L267:
	;
	v1642 = int32(*(*uint8)(unsafe.Add(mBase, _consts[813])))
	if v1642 == int32(0) {
		v562 = v1618
		v572 = v749
		v574 = v750
		v575 = v751
		v584 = v772
		goto L156
	} else {
		goto L449
	}
L268:
	;
	if v933 != int32(1) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1618 = v562
	goto L267
L270:
	;
	goto L271
L271:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v937 != int32(2) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1618 = v562
	goto L267
L273:
	;
	goto L274
L274:
	;
	v941 = *(*int32)(unsafe.Add(mBase, _consts[806]))
	v947 = F_read(m, v941, v26+int32(48)+v562, int32(8192)-v562)
	mBase = m.M
	if v947 < int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v951 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v951 == int32(27) {
		goto L278
	} else {
		goto L279
	}
L276:
	;
	goto L277
L277:
	;
	if v947 != 0 {
		goto L288
	} else {
		goto L289
	}
L278:
	;
	v1618 = v562
	goto L267
L279:
	;
	goto L280
L280:
	;
	v956 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L4
	} else {
		goto L281
	}
L281:
	;
	if v956 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1618 = v562
	goto L267
L283:
	;
	goto L284
L284:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L4
	} else {
		goto L285
	}
L285:
	;
	F_errmsg(m, int32(294806), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L4
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(495615), int32(527), int32(278921))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	v1618 = v562
	goto L267
L288:
	;
	v971 = v562 + v947
	if v971 < int32(10) {
		v562 = v971
		v572 = v749
		v574 = v750
		v575 = v751
		v584 = v772
		goto L156
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v1489 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[813])) = uint8(v1489)
	v1497 = int32(0)
	goto L427
L291:
	;
	v977 = v971
	v980 = v26 + int32(48)
	v984 = int32(1)
	goto L292
L292:
	;
	v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980))))
	if v1000 != 0 {
		goto L297
	} else {
		goto L298
	}
L293:
	;
	if v1313 <= int32(0) {
		v562 = v1313
		v572 = v749
		v574 = v750
		v575 = v751
		v584 = v772
		goto L156
	} else {
		goto L379
	}
L294:
	;
	goto L293
L295:
	;
	v1309 = v1288 + v980
	v1310 = v977 - v1288
	if int32(9) < v1310 {
		v977 = v1310
		v980 = v1309
		v984 = v1293
		goto L292
	} else {
		goto L378
	}
L296:
	;
	F_write_stderr(m, int32(748979), int32(0))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L4
	} else {
		goto L377
	}
L297:
	;
	v1224 = int32(1)
	goto L371
L298:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980)+1)))
	if v1001 != 0 {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1002 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v980)+2)))
	if base.Ui32(int32(4086)) < base.Ui32((v1002-int32(1))&int32(65535)) {
		goto L297
	} else {
		goto L300
	}
L300:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v980)+4))
	if v1009 == int32(0) {
		goto L297
	} else {
		goto L301
	}
L301:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980)+8)))
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012&int32(112))+uint32(_consts[814]))))
	if v1017 != int32(1) {
		goto L297
	} else {
		goto L302
	}
L302:
	;
	v1021 = v1002 + int32(9)
	if base.Ui32(v977) < base.Ui32(v1021) {
		v1313 = v977
		v1316 = v980
		goto L294
	} else {
		goto L303
	}
L303:
	;
	if v1012&int32(16) != 0 {
		v1033 = int32(1)
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1034 = int32(0)
	v1036 = base.I32_rem_s(v1009, int32(256))
	v1038 = v1036 << (uint(int32(2)) % 32)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+uint32(_consts[815])))
	if v1041 != 0 {
		goto L311
	} else {
		goto L312
	}
L305:
	;
	if v1012&int32(32) != 0 {
		v1033 = int32(8)
		goto L304
	} else {
		goto L306
	}
L306:
	;
	if v1012&int32(64) != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1032 = int32(16)
	goto L309
L308:
	;
	v1032 = v984
	goto L309
L309:
	;
	v1033 = v1032
	goto L304
L310:
	;
	if v1012&int32(1) == int32(0) {
		goto L322
	} else {
		goto L323
	}
L311:
	;
	v1042 = int32(0)
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+4))
	if v1043 <= v1042 {
		v1107 = v1042
		v1111 = v1034
		goto L310
	} else {
		goto L314
	}
L312:
	;
	v1087 = v1034
	goto L313
L313:
	;
	v1107 = int32(0)
	v1111 = v1087
	goto L310
L314:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+12))
	v1053 = v1034
	v1054 = int32(0)
	goto L315
L315:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1046+v1054<<(uint(int32(2))%32))))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1074)))
	if v1075 == v1009 {
		v1107 = v1074
		v1111 = v1053
		goto L310
	} else {
		goto L317
	}
L316:
	;
	v1087 = v1078
	goto L313
L317:
	;
	if v1053|v1075 != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1078 = v1053
	goto L320
L319:
	;
	v1078 = v1074
	goto L320
L320:
	;
	v1080 = v1054 + int32(1)
	if v1043 != v1080 {
		v1053 = v1078
		v1054 = v1080
		goto L315
	} else {
		goto L321
	}
L321:
	;
	goto L316
L322:
	;
	if v1107 != 0 {
		goto L325
	} else {
		goto L326
	}
L323:
	;
	goto L324
L324:
	;
	if v1107 != 0 {
		goto L336
	} else {
		goto L337
	}
L325:
	;
	F_appendBinaryStringInfo(m, v1107+int32(4), v980+int32(9), v1002)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L4
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	if v1111 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v1288 = v1021
	v1293 = v1033
	goto L295
L329:
	;
	v1142 = F_palloc(m, int32(20))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L4
	} else {
		goto L332
	}
L330:
	;
	v1147 = v1111
	goto L331
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1147))) = v1009
	v1150 = v1147 + int32(4)
	F_initStringInfo(m, v1150)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L4
	} else {
		goto L334
	}
L332:
	;
	v1144 = F_lappend(m, v1041, v1142)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L4
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1038)+uint32(_consts[815]))) = v1144
	v1147 = v1142
	goto L331
L334:
	;
	F_appendBinaryStringInfo(m, v1150, v980+int32(9), v1002)
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L4
	} else {
		goto L335
	}
L335:
	;
	v1288 = v1021
	v1293 = v1033
	goto L295
L336:
	;
	F_appendBinaryStringInfo(m, v1107+int32(4), v980+int32(9), v1002)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L4
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	if v1033&int32(8) != 0 {
		goto L358
	} else {
		goto L359
	}
L339:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+8))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+4))
	if v1033&int32(8) != 0 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	v1183 = F_fwrite(m, v1164, int32(1), v1163, v1180)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L4
	} else {
		goto L351
	}
L341:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, _consts[800]))
	if v1168 != 0 {
		v1180 = v1168
		goto L340
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	v1173 = *(*int32)(unsafe.Add(mBase, _consts[812]))
	if v1171 != 0 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	goto L343
L345:
	;
	v1174 = v1171
	goto L347
L346:
	;
	v1174 = v1173
	goto L347
L347:
	;
	if int32(base.Ui32(v1033&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1179 = v1174
	goto L350
L349:
	;
	v1179 = v1173
	goto L350
L350:
	;
	v1180 = v1179
	goto L340
L351:
	;
	if v1183 != v1163 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	F_write_stderr(m, int32(748979), int32(0))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L4
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1107))) = int32(0)
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1107)+4))
	F_pfree(m, v1192)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L4
	} else {
		goto L356
	}
L355:
	;
	goto L354
L356:
	;
	v1288 = v1021
	v1293 = v1033
	goto L295
L357:
	;
	v1215 = F_fwrite(m, v980+int32(9), int32(1), v1002, v1210)
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L4
	} else {
		goto L368
	}
L358:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, _consts[800]))
	if v1198 != 0 {
		v1210 = v1198
		goto L357
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	v1203 = *(*int32)(unsafe.Add(mBase, _consts[812]))
	if v1201 != 0 {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	goto L360
L362:
	;
	v1204 = v1201
	goto L364
L363:
	;
	v1204 = v1203
	goto L364
L364:
	;
	if int32(base.Ui32(v1033&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1209 = v1204
	goto L367
L366:
	;
	v1209 = v1203
	goto L367
L367:
	;
	v1210 = v1209
	goto L357
L368:
	;
	if v1215 == v1002 {
		v1288 = v1021
		v1293 = v1033
		goto L295
	} else {
		goto L369
	}
L369:
	;
	v1261 = v1021
	v1266 = v1033
	goto L296
L370:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, _consts[812]))
	v1256 = F_fwrite(m, v980, int32(1), v1252, v1255)
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L4
	} else {
		goto L375
	}
L371:
	;
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1224+v980))))
	if v1246 == int32(0) {
		v1252 = v1224
		goto L370
	} else {
		goto L373
	}
L372:
	;
	v1252 = v977
	goto L370
L373:
	;
	v1250 = v1224 + int32(1)
	if v1250 != v977 {
		v1224 = v1250
		goto L371
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	if v1256 == v1252 {
		v1288 = v1252
		v1293 = v984
		goto L295
	} else {
		goto L376
	}
L376:
	;
	v1261 = v1252
	v1266 = v984
	goto L296
L377:
	;
	v1288 = v1261
	v1293 = v1266
	goto L295
L378:
	;
	v1313 = v1310
	v1316 = v1309
	goto L294
L379:
	;
	if v1316 == v26+int32(48) {
		v562 = v1313
		v572 = v749
		v574 = v750
		v575 = v751
		v584 = v772
		goto L156
	} else {
		goto L380
	}
L380:
	;
	v1342 = v26 + int32(48)
	if v1342 == v1316 {
		goto L382
	} else {
		goto L383
	}
L381:
	;
	v562 = v1313
	v572 = v749
	v574 = v750
	v575 = v751
	v584 = v772
	goto L156
L382:
	;
	goto L381
L383:
	;
	v1346 = v1342 + v1313
	if base.Ui32(v1316-v1346) <= base.Ui32(int32(0)-v1313<<(uint(int32(1))%32)) {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1353 = F___memcpy(m, v1342, v1316, v1313)
	mBase = m.M
	goto L381
L385:
	;
	goto L386
L386:
	;
	v1356 = (v1342 ^ v1316) & int32(3)
	if base.Ui32(v1342) < base.Ui32(v1316) {
		goto L389
	} else {
		goto L390
	}
L387:
	;
	if v1458 == int32(0) {
		goto L382
	} else {
		goto L423
	}
L388:
	;
	if base.Ui32(v1436) <= base.Ui32(int32(3)) {
		v1457 = v1435
		v1458 = v1436
		v1459 = v1437
		goto L387
	} else {
		goto L419
	}
L389:
	;
	if v1356 != 0 {
		goto L392
	} else {
		goto L393
	}
L390:
	;
	goto L391
L391:
	;
	if v1356 != 0 {
		v1418 = v1313
		goto L402
	} else {
		goto L403
	}
L392:
	;
	v1457 = v1316
	v1458 = v1313
	v1459 = v1342
	goto L387
L393:
	;
	goto L394
L394:
	;
	if v1342&int32(3) == int32(0) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1435 = v1316
	v1436 = v1313
	v1437 = v1342
	goto L388
L396:
	;
	goto L397
L397:
	;
	v1363 = v1316
	v1364 = v1313
	v1365 = v1342
	goto L398
L398:
	;
	if v1364 == int32(0) {
		goto L382
	} else {
		goto L400
	}
L399:
	;
	v1435 = v1372
	v1436 = v1374
	v1437 = v1376
	goto L388
L400:
	;
	v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1363))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1365))) = uint8(v1369)
	v1371 = int32(1)
	v1372 = v1363 + v1371
	v1374 = v1364 - v1371
	v1376 = v1365 + v1371
	if v1376&int32(3) != 0 {
		v1363 = v1372
		v1364 = v1374
		v1365 = v1376
		goto L398
	} else {
		goto L401
	}
L401:
	;
	goto L399
L402:
	;
	if v1418 == int32(0) {
		goto L382
	} else {
		goto L415
	}
L403:
	;
	if v1346&int32(3) != 0 {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1383 = v1313
	goto L407
L405:
	;
	v1398 = v1313
	goto L406
L406:
	;
	if base.Ui32(v1398) <= base.Ui32(int32(3)) {
		v1418 = v1398
		goto L402
	} else {
		goto L411
	}
L407:
	;
	if v1383 == int32(0) {
		goto L382
	} else {
		goto L409
	}
L408:
	;
	v1398 = v1389
	goto L406
L409:
	;
	v1389 = v1383 - int32(1)
	v1390 = v1342 + v1389
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1316+v1389))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1390))) = uint8(v1392)
	if v1390&int32(3) != 0 {
		v1383 = v1389
		goto L407
	} else {
		goto L410
	}
L410:
	;
	goto L408
L411:
	;
	v1405 = v1398
	goto L412
L412:
	;
	v1409 = v1405 - int32(4)
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1316+v1409)))
	*(*int32)(unsafe.Add(mBase, uint32(v1342+v1409))) = v1412
	if base.Ui32(int32(3)) < base.Ui32(v1409) {
		v1405 = v1409
		goto L412
	} else {
		goto L414
	}
L413:
	;
	v1418 = v1409
	goto L402
L414:
	;
	goto L413
L415:
	;
	v1425 = v1418
	goto L416
L416:
	;
	v1429 = v1425 - int32(1)
	v1432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1316+v1429))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1342+v1429))) = uint8(v1432)
	if v1429 != 0 {
		v1425 = v1429
		goto L416
	} else {
		goto L418
	}
L417:
	;
	goto L382
L418:
	;
	goto L417
L419:
	;
	v1442 = v1435
	v1443 = v1436
	v1444 = v1437
	goto L420
L420:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1442)))
	*(*int32)(unsafe.Add(mBase, uint32(v1444))) = v1446
	v1448 = int32(4)
	v1449 = v1442 + v1448
	v1451 = v1444 + v1448
	v1453 = v1443 - v1448
	if base.Ui32(int32(3)) < base.Ui32(v1453) {
		v1442 = v1449
		v1443 = v1453
		v1444 = v1451
		goto L420
	} else {
		goto L422
	}
L421:
	;
	v1457 = v1449
	v1458 = v1453
	v1459 = v1451
	goto L387
L422:
	;
	goto L421
L423:
	;
	v1464 = v1457
	v1465 = v1458
	v1466 = v1459
	goto L424
L424:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1466))) = uint8(v1468)
	v1470 = int32(1)
	v1475 = v1465 - v1470
	if v1475 != 0 {
		v1464 = v1464 + v1470
		v1465 = v1475
		v1466 = v1466 + v1470
		goto L424
	} else {
		goto L426
	}
L425:
	;
	goto L382
L426:
	;
	goto L425
L427:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1497<<(uint(int32(2))%32))+uint32(_consts[815])))
	if v1518 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L428:
	;
	v1603 = int32(0)
	if v562 <= v1603 {
		v1618 = v1603
		goto L267
	} else {
		goto L445
	}
L429:
	;
	v1600 = v1497 + int32(1)
	if v1600 != int32(256) {
		v1497 = v1600
		goto L427
	} else {
		goto L444
	}
L430:
	;
	v1521 = int32(0)
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1518)+4))
	if v1522 <= v1521 {
		goto L429
	} else {
		goto L431
	}
L431:
	;
	v1525 = v1522
	v1528 = v1521
	goto L432
L432:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1518)+12))
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1548+v1528<<(uint(int32(2))%32))))
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1552)))
	if v1553 != 0 {
		goto L434
	} else {
		goto L435
	}
L433:
	;
	goto L429
L434:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+4))
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+8))
	v1558 = *(*int32)(unsafe.Add(mBase, _consts[812]))
	v1559 = F_fwrite(m, v1554, int32(1), v1556, v1558)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L4
	} else {
		goto L437
	}
L435:
	;
	v1572 = v1525
	goto L436
L436:
	;
	v1574 = v1528 + int32(1)
	if v1574 < v1572 {
		v1525 = v1572
		v1528 = v1574
		goto L432
	} else {
		goto L443
	}
L437:
	;
	if v1559 != v1556 {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	F_write_stderr(m, int32(748979), int32(0))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L4
	} else {
		goto L441
	}
L439:
	;
	goto L440
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1552))) = int32(0)
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+4))
	F_pfree(m, v1568)
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L4
	} else {
		goto L442
	}
L441:
	;
	goto L440
L442:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1518)+4))
	v1572 = v1571
	goto L436
L443:
	;
	goto L433
L444:
	;
	goto L428
L445:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, _consts[812]))
	v1611 = F_fwrite(m, v26+int32(48), int32(1), v562, v1610)
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L4
	} else {
		goto L446
	}
L446:
	;
	if v1611 == v562 {
		v1618 = v1603
		goto L267
	} else {
		goto L447
	}
L447:
	;
	F_write_stderr(m, int32(748979), int32(0))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L4
	} else {
		goto L448
	}
L448:
	;
	v1618 = v1603
	goto L267
L449:
	;
	goto L157
L450:
	;
	if v1647 != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	F_errmsg_internal(m, int32(244463), int32(0))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L4
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L4
	} else {
		goto L456
	}
L454:
	;
	F_errfinish(m, int32(495615), int32(575), int32(278921))
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L4
	} else {
		goto L455
	}
L455:
	;
	goto L453
L456:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
