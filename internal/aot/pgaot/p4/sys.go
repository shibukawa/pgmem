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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13859(m, l0, l1, int32(_a_F_SearchSysCacheCopyAttName_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(5))%32))+uint32(_c_F_SysCacheGetAttrNotNull[0])))
				v28 = F_get_rel_name(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SysCacheGetAttrNotNull[1])))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v35 + v36<<(uint(int32(4))%32) + l2*int32(100) - int32(76)
					F_errmsg_internal(m, int32(_a_F_SysCacheGetAttrNotNull_0), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_SysCacheGetAttrNotNull_1), int32(644), int32(_a_F_SysCacheGetAttrNotNull_2))
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
	var v3 int32
	_ = v3
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v310 int32
	_ = v310
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v432 int64
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v464 int64
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int64
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int64
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int64
	_ = v499
	var v502 int64
	_ = v502
	var v504 int64
	_ = v504
	var v506 int64
	_ = v506
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v556 int64
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v681 int32
	_ = v681
	var v685 int64
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int64
	_ = v694
	var v697 int64
	_ = v697
	var v699 int64
	_ = v699
	var v701 int64
	_ = v701
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int64
	_ = v739
	var v741 int64
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int64
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v763 int64
	_ = v763
	var v765 int64
	_ = v765
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v778 int64
	_ = v778
	var v780 int64
	_ = v780
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v794 int64
	_ = v794
	var v796 int64
	_ = v796
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int64
	_ = v818
	var v819 int64
	_ = v819
	var v820 int64
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v848 int64
	_ = v848
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int64
	_ = v857
	var v860 int64
	_ = v860
	var v862 int64
	_ = v862
	var v864 int64
	_ = v864
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v880 int64
	_ = v880
	var v882 int64
	_ = v882
	var v883 int64
	_ = v883
	var v886 int64
	_ = v886
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v952 int32
	_ = v952
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1187 int32
	_ = v1187
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1232 int32
	_ = v1232
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1259 int32
	_ = v1259
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1313 int32
	_ = v1313
	var v1321 int32
	_ = v1321
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	v3 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(_a_F_SysLoggerMain_0)
	m.G0 = v26
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v29 != int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = F___fdopen(m, v29, int32(_a_F_SysLoggerMain_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = int32(-1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v43 = v3
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0])) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v46 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v43 = v33
	goto L3
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = int32(10)
	goto L7
L6:
	;
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v39 | int32(64)
	goto L4
L8:
	;
	v50 = F___fdopen(m, v46, int32(_a_F_SysLoggerMain_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v50)+80)) = int32(-1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	if v53 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v60 = v3
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1])) = v60
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v64 != int32(-1) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v60 = v50
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+80)) = int32(10)
	goto L14
L13:
	;
	goto L14
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v56 | int32(64)
	goto L11
L15:
	;
	v68 = F___fdopen(m, v64, int32(_a_F_SysLoggerMain_1))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v68)+80)) = int32(-1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+48))
	if v71 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v78 = int32(0)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2])) = v78
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[3]))
	if v82 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v78 = v68
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+80)) = int32(10)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v74 | int32(64)
	goto L18
L22:
	;
	F_MemoryContextDelete(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[4])) = int32(17)
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[5]))
	goto L28
L25:
	;
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[3])) = int32(0)
	goto L24
L27:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[6])))
	if v100 != int32(1) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[4]))
	v98 = F_GetBackendTypeDesc(m, v97)
	mBase = m.M
	goto L30
L30:
	;
	goto L27
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[7]))
	if int32(0) <= v123 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = int32(0)
	v106 = int32(1)
	v109 = F_open(m, int32(_a_F_SysLoggerMain_2), v106, v26+int32(16))
	mBase = m.M
	v111 = F_close(m, v106)
	mBase = m.M
	v113 = F_close(m, int32(2))
	mBase = m.M
	if v109 == int32(-1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v117 = F_dup2(m, v109, int32(1))
	mBase = m.M
	v119 = F_dup2(m, v109, int32(2))
	mBase = m.M
	v120 = F_close(m, v109)
	mBase = m.M
	goto L31
L34:
	;
	v126 = F_close(m, v123)
	mBase = m.M
	goto L36
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[7])) = int32(-1)
	v131 = int32(914)
	v133 = m.G0
	v135 = v133 - int32(32)
	m.G0 = v135
	switch int32(916) {
	case 0, 2:
		v145 = v131
		goto L38
	default:
		goto L39
	}
L37:
	;
	v164 = int32(-2)
	v166 = m.G0
	v168 = v166 - int32(32)
	m.G0 = v168
	switch int32(0) {
	case 0, 2:
		v178 = v164
		goto L44
	default:
		goto L45
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v145
	F_sigemptyset(m, v135+int32(16))
	mBase = m.M
	goto L41
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[8])) = v131
	v145 = int32(_a_F_SysLoggerMain_3)
	goto L38
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+24)) = int32(268435456)
	v159 = F___sigaction(m, int32(1), v135+int32(12), int32(0))
	mBase = m.M
	m.G0 = v135 + int32(32)
	goto L37
L43:
	;
	v197 = int32(-2)
	v199 = m.G0
	v201 = v199 - int32(32)
	m.G0 = v201
	switch int32(0) {
	case 0, 2:
		v211 = v197
		goto L50
	default:
		goto L51
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v178
	F_sigemptyset(m, v168+int32(16))
	mBase = m.M
	goto L47
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[9])) = v164
	v178 = int32(_a_F_SysLoggerMain_3)
	goto L44
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+24)) = int32(268435456)
	v192 = F___sigaction(m, int32(2), v168+int32(12), int32(0))
	mBase = m.M
	m.G0 = v168 + int32(32)
	goto L43
L49:
	;
	v230 = int32(-2)
	v232 = m.G0
	v234 = v232 - int32(32)
	m.G0 = v234
	switch int32(0) {
	case 0, 2:
		v244 = v230
		goto L56
	default:
		goto L57
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+12)) = v211
	F_sigemptyset(m, v201+int32(16))
	mBase = m.M
	goto L53
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[10])) = v197
	v211 = int32(_a_F_SysLoggerMain_3)
	goto L50
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+24)) = int32(268435456)
	v225 = F___sigaction(m, int32(15), v201+int32(12), int32(0))
	mBase = m.M
	m.G0 = v201 + int32(32)
	goto L49
L55:
	;
	v263 = int32(-2)
	v265 = m.G0
	v267 = v265 - int32(32)
	m.G0 = v267
	switch int32(0) {
	case 0, 2:
		v277 = v263
		goto L62
	default:
		goto L63
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+12)) = v244
	F_sigemptyset(m, v234+int32(16))
	mBase = m.M
	goto L59
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[11])) = v230
	v244 = int32(_a_F_SysLoggerMain_3)
	goto L56
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+24)) = int32(268435456)
	v258 = F___sigaction(m, int32(3), v234+int32(12), int32(0))
	mBase = m.M
	m.G0 = v234 + int32(32)
	goto L55
L61:
	;
	v296 = int32(-2)
	v298 = m.G0
	v300 = v298 - int32(32)
	m.G0 = v300
	switch int32(0) {
	case 0, 2:
		v310 = v296
		goto L68
	default:
		goto L69
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267)+12)) = v277
	F_sigemptyset(m, v267+int32(16))
	mBase = m.M
	goto L65
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[12])) = v263
	v277 = int32(_a_F_SysLoggerMain_3)
	goto L62
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267)+24)) = int32(268435456)
	v291 = F___sigaction(m, int32(14), v267+int32(12), int32(0))
	mBase = m.M
	m.G0 = v267 + int32(32)
	goto L61
L67:
	;
	v329 = int32(964)
	v331 = m.G0
	v333 = v331 - int32(32)
	m.G0 = v333
	switch int32(966) {
	case 0, 2:
		v343 = v329
		goto L74
	default:
		goto L75
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v300)+12)) = v310
	F_sigemptyset(m, v300+int32(16))
	mBase = m.M
	goto L71
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[13])) = v296
	v310 = int32(_a_F_SysLoggerMain_3)
	goto L68
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v300)+24)) = int32(268435456)
	v324 = F___sigaction(m, int32(13), v300+int32(12), int32(0))
	mBase = m.M
	m.G0 = v300 + int32(32)
	goto L67
L73:
	;
	v362 = int32(-2)
	v364 = m.G0
	v366 = v364 - int32(32)
	m.G0 = v366
	switch int32(0) {
	case 0, 2:
		v376 = v362
		goto L80
	default:
		goto L81
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+12)) = v343
	F_sigemptyset(m, v333+int32(16))
	mBase = m.M
	goto L77
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[14])) = v329
	v343 = int32(_a_F_SysLoggerMain_3)
	goto L74
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+24)) = int32(268435456)
	v357 = F___sigaction(m, int32(10), v333+int32(12), int32(0))
	mBase = m.M
	m.G0 = v333 + int32(32)
	goto L73
L79:
	;
	v395 = int32(0)
	v397 = m.G0
	v399 = v397 - int32(32)
	m.G0 = v399
	switch int32(2) {
	case 0, 2:
		v409 = v395
		goto L86
	default:
		goto L87
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+12)) = v376
	F_sigemptyset(m, v366+int32(16))
	mBase = m.M
	goto L83
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[15])) = v362
	v376 = int32(_a_F_SysLoggerMain_3)
	goto L80
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+24)) = int32(268435456)
	v390 = F___sigaction(m, int32(12), v366+int32(12), int32(0))
	mBase = m.M
	m.G0 = v366 + int32(32)
	goto L79
L85:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_SysLoggerMain_4), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L25
	} else {
		goto L91
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v399)+12)) = v409
	F_sigemptyset(m, v399+int32(16))
	mBase = m.M
	goto L88
L87:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[16])) = v395
	v409 = int32(_a_F_SysLoggerMain_3)
	goto L86
L88:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v399)+24)) = int32(268435457)
	v423 = F___sigaction(m, int32(17), v399+int32(12), int32(0))
	mBase = m.M
	m.G0 = v399 + int32(32)
	goto L85
L91:
	;
	v432 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[17]))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v432
	v435 = F_palloc(m, int32(1024))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L25
	} else {
		goto L92
	}
L92:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v438
	v442 = F_pg_snprintf(m, v435, int32(1024), int32(_a_F_SysLoggerMain_5), v26)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L25
	} else {
		goto L93
	}
L93:
	;
	v444 = F_strlen(m, v435)
	mBase = m.M
	v449 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v454 = F_pg_localtime(m, v26+int32(48), v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L25
	} else {
		goto L94
	}
L94:
	;
	v456 = F_pg_strftime(m, v435+v444, int32(1024)-v444, v449, v454)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L25
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[21])) = v435
	v461 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v461 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v464 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[17]))
	v466 = F_logfile_getname(m, v464, int32(_a_F_SysLoggerMain_6))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L25
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	if v470 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[22])) = v466
	goto L98
L100:
	;
	v473 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[17]))
	v475 = F_logfile_getname(m, v473, int32(_a_F_SysLoggerMain_7))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L25
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	v480 = F_pstrdup(m, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L25
	} else {
		goto L104
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[23])) = v475
	goto L102
L104:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v484 = F_pstrdup(m, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L25
	} else {
		goto L105
	}
L105:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24]))
	if int32(0) < v487 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v490 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v490
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v496 = F_pg_localtime(m, v26+int32(48), v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L25
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L25
	} else {
		goto L110
	}
L109:
	;
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v26)+48))
	v502 = base.I64_extend_i32_s(v487 * int32(60))
	v504 = int64(*(*int32)(unsafe.Add(mBase, uint32(v496)+36)))
	v506 = base.I64_rem_s(v499+v504, v502)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25])) = v499 + v502 - v506
	goto L108
L110:
	;
	v515 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[26])) = v515
	v519 = F_CreateWaitEventSet(m, v515, int32(2))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L25
	} else {
		goto L111
	}
L111:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[27]))
	F_AddWaitEventToSet(m, v519, int32(1), int32(-1), v524)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L25
	} else {
		goto L112
	}
L112:
	;
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[28]))
	F_AddWaitEventToSet(m, v519, int32(2), v529, int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L25
	} else {
		goto L113
	}
L113:
	;
	v535 = int32(0)
	v544 = v487
	v545 = v480
	v546 = v484
	v556 = v92
	goto L114
L114:
	;
	v558 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[27]))
	*(*int32)(unsafe.Add(mBase, uint32(v558))) = int32(0)
	goto L116
L115:
	;
	v1469 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L25
	} else {
		goto L345
	}
L116:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[29]))
	if v562 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[29])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L25
	} else {
		goto L120
	}
L118:
	;
	v723 = v544
	v724 = v545
	v725 = v546
	goto L119
L119:
	;
	v728 = int32(0)
	v730 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24]))
	if v730 <= v728 {
		goto L164
	} else {
		goto L165
	}
L120:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570))))
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	if base.B2i32(v573 == int32(0))|base.B2i32(v573 != v576) != 0 {
		v594 = v573
		v595 = v576
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v594-v595 != 0 {
		goto L128
	} else {
		goto L129
	}
L122:
	;
	goto L121
L123:
	;
	v579 = v570
	v580 = v545
	goto L124
L124:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+1)))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+1)))
	if v584 == int32(0) {
		v594 = v584
		v595 = v583
		goto L122
	} else {
		goto L126
	}
L125:
	;
	v594 = v584
	v595 = v583
	goto L122
L126:
	;
	v587 = int32(1)
	if v584 == v583 {
		v579 = v579 + v587
		v580 = v580 + v587
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	F_pfree(m, v545)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L25
	} else {
		goto L131
	}
L129:
	;
	v611 = v545
	goto L130
L130:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546))))
	if base.B2i32(v616 == int32(0))|base.B2i32(v616 != v619) != 0 {
		v637 = v616
		v638 = v619
		goto L135
	} else {
		goto L136
	}
L131:
	;
	v600 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	v601 = F_pstrdup(m, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L25
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v607 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31]))
	v610 = F_mkdir(m, v607, v609)
	mBase = m.M
	goto L133
L133:
	;
	v611 = v601
	goto L130
L134:
	;
	if v637-v638 != 0 {
		goto L141
	} else {
		goto L142
	}
L135:
	;
	goto L134
L136:
	;
	v622 = v613
	v623 = v546
	goto L137
L137:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+1)))
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622)+1)))
	if v627 == int32(0) {
		v637 = v627
		v638 = v626
		goto L135
	} else {
		goto L139
	}
L138:
	;
	v637 = v627
	v638 = v626
	goto L135
L139:
	;
	v630 = int32(1)
	if v627 == v626 {
		v622 = v622 + v630
		v623 = v623 + v630
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	F_pfree(m, v546)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L25
	} else {
		goto L144
	}
L142:
	;
	v649 = v546
	goto L143
L143:
	;
	v651 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[32]))
	v654 = int32(0)
	v657 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if base.B2i32(v651&int32(8) == v654)^base.B2i32(v657 != v654) == v654 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v644 = F_pstrdup(m, v643)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L25
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v649 = v644
	goto L143
L146:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	goto L148
L147:
	;
	goto L148
L148:
	;
	v668 = int32(0)
	v671 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	if base.B2i32(v651&int32(16) == v668)^base.B2i32(v671 != v668) == v668 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	goto L151
L150:
	;
	goto L151
L151:
	;
	v681 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24]))
	if v681 != v544 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if int32(0) < v681 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v708 = v544
	goto L154
L154:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])))
	if v712 != 0 {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	v685 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v685
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v691 = F_pg_localtime(m, v26+int32(32), v690)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L25
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v708 = v681
	goto L154
L158:
	;
	v694 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v697 = base.I64_extend_i32_s(v681 * int32(60))
	v699 = int64(*(*int32)(unsafe.Add(mBase, uint32(v691)+36)))
	v701 = base.I64_rem_s(v694+v699, v697)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25])) = v694 + v697 - v701
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v717 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])) = uint8(v717)
	goto L161
L160:
	;
	goto L161
L161:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L25
	} else {
		goto L162
	}
L162:
	;
	v723 = v708
	v724 = v611
	v725 = v649
	goto L119
L163:
	;
	v750 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])))
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30]))
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[34]))
	if v750&int32(1)|(v754|base.B2i32(v756 <= int32(0))) != 0 {
		v805 = v728
		goto L169
	} else {
		goto L170
	}
L164:
	;
	v747 = int32(0)
	v748 = v556
	goto L163
L165:
	;
	goto L166
L166:
	;
	v734 = int32(0)
	v736 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])))
	if v736&int32(1) != 0 {
		v747 = v734
		v748 = v556
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v739 = F_time(m)
	mBase = m.M
	v741 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	if v739 < v741 {
		v747 = v734
		v748 = v739
		goto L163
	} else {
		goto L168
	}
L168:
	;
	v743 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = v743
	v747 = v743
	v748 = v739
	goto L163
L169:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30]))
	if v808 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L170:
	;
	v762 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v763 = F___ftello_unlocked(m, v762)
	mBase = m.M
	v765 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	if v765<<(uint(int64(10))%64) <= v763 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v770 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = v770
	v773 = v770
	goto L173
L172:
	;
	v773 = v728
	goto L173
L173:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v775 == int32(0) {
		v789 = v773
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	if v791 == int32(0) {
		v805 = v789
		goto L169
	} else {
		goto L177
	}
L175:
	;
	v778 = F___ftello_unlocked(m, v775)
	mBase = m.M
	v780 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	if v778 < v780<<(uint(int64(10))%64) {
		v789 = v773
		goto L174
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v789 = v773 | int32(8)
	goto L174
L177:
	;
	v794 = F___ftello_unlocked(m, v791)
	mBase = m.M
	v796 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	if v794 < v796<<(uint(int64(10))%64) {
		v805 = v789
		goto L169
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v805 = v789 | int32(16)
	goto L169
L179:
	;
	v871 = int32(-1)
	v873 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24]))
	if v873 <= int32(0) {
		v894 = v871
		goto L200
	} else {
		goto L201
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(0)
	if v805 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v815 = v805
	goto L183
L182:
	;
	v815 = int32(25)
	goto L183
L183:
	;
	if v747 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v816 = v805
	goto L186
L185:
	;
	v816 = v815
	goto L186
L186:
	;
	if v747 != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v824 = F_logfile_rotate_dest(m, v747, v816, v820, int32(1), int32(_a_F_SysLoggerMain_8), int32(_a_F_SysLoggerMain_9))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L25
	} else {
		goto L191
	}
L188:
	;
	v818 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	v820 = v818
	goto L187
L189:
	;
	goto L190
L190:
	;
	v819 = F_time(m)
	mBase = m.M
	v820 = v819
	goto L187
L191:
	;
	if v824 == int32(0) {
		goto L179
	} else {
		goto L192
	}
L192:
	;
	v831 = F_logfile_rotate_dest(m, v747, v816, v820, int32(8), int32(_a_F_SysLoggerMain_10), int32(_a_F_SysLoggerMain_11))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L25
	} else {
		goto L193
	}
L193:
	;
	if v831 == int32(0) {
		goto L179
	} else {
		goto L194
	}
L194:
	;
	v838 = F_logfile_rotate_dest(m, v747, v816, v820, int32(16), int32(_a_F_SysLoggerMain_12), int32(_a_F_SysLoggerMain_13))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L25
	} else {
		goto L195
	}
L195:
	;
	if v838 == int32(0) {
		goto L179
	} else {
		goto L196
	}
L196:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L25
	} else {
		goto L197
	}
L197:
	;
	v845 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24]))
	if v845 <= int32(0) {
		goto L179
	} else {
		goto L198
	}
L198:
	;
	v848 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v848
	v853 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v854 = F_pg_localtime(m, v26+int32(32), v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L25
	} else {
		goto L199
	}
L199:
	;
	v857 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v860 = base.I64_extend_i32_s(v845 * int32(60))
	v862 = int64(*(*int32)(unsafe.Add(mBase, uint32(v854)+36)))
	v864 = base.I64_rem_s(v857+v862, v860)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25])) = v857 + v860 - v864
	goto L179
L200:
	;
	v900 = F_WaitEventSetWait(m, v519, v894, v26+int32(32), int32(1), int32(83886093))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L25
	} else {
		goto L210
	}
L201:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])))
	if v877&int32(1) != 0 {
		v894 = v871
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v880 = int64(2147483)
	v882 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	v883 = v882 - v748
	if v880 <= v883 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v886 = v880
	goto L205
L204:
	;
	v886 = v883
	goto L205
L205:
	;
	if int64(0) < v883 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v893 = base.I32_wrap_i64(v886) * int32(1000)
	goto L208
L207:
	;
	v893 = int32(0)
	goto L208
L208:
	;
	v894 = v893
	goto L200
L209:
	;
	v1464 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[35])))
	if v1464 == int32(0) {
		v535 = v1441
		v544 = v723
		v545 = v724
		v546 = v725
		v556 = v748
		goto L114
	} else {
		goto L344
	}
L210:
	;
	if v900 != int32(1) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1441 = v535
	goto L209
L212:
	;
	goto L213
L213:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v904 != int32(2) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1441 = v535
	goto L209
L215:
	;
	goto L216
L216:
	;
	v908 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[28]))
	v914 = F_read(m, v908, v26+int32(48)+v535, int32(_a_F_SysLoggerMain_14)-v535)
	mBase = m.M
	if v914 < int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v918 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[36]))
	if v918 == int32(27) {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	goto L219
L219:
	;
	if v914 != 0 {
		goto L230
	} else {
		goto L231
	}
L220:
	;
	v1441 = v535
	goto L209
L221:
	;
	goto L222
L222:
	;
	v923 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L25
	} else {
		goto L223
	}
L223:
	;
	if v923 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1441 = v535
	goto L209
L225:
	;
	goto L226
L226:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L25
	} else {
		goto L227
	}
L227:
	;
	F_errmsg(m, int32(_a_F_SysLoggerMain_15), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L25
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_SysLoggerMain_16), int32(527), int32(_a_F_SysLoggerMain_17))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L25
	} else {
		goto L229
	}
L229:
	;
	v1441 = v535
	goto L209
L230:
	;
	v938 = v914 + v535
	if v938 < int32(10) {
		v535 = v938
		v544 = v723
		v545 = v724
		v546 = v725
		v556 = v748
		goto L114
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v1313 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[35])) = uint8(v1313)
	v1321 = int32(0)
	goto L322
L233:
	;
	v945 = v938
	v946 = v26 + int32(48)
	v952 = int32(1)
	goto L234
L234:
	;
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946))))
	if v967 != 0 {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v1301 = int32(0)
	v1304 = v26 + int32(48)
	if base.B2i32(v1279 == v1301)|(base.B2i32(v1280 == v1304)|base.B2i32(v1279 <= v1301)) != 0 {
		v535 = v1279
		v544 = v723
		v545 = v724
		v546 = v725
		v556 = v748
		goto L114
	} else {
		goto L321
	}
L236:
	;
	goto L235
L237:
	;
	v1274 = v1251 + v946
	v1275 = v945 - v1251
	if int32(9) < v1275 {
		v945 = v1275
		v946 = v1274
		v952 = v1259
		goto L234
	} else {
		goto L320
	}
L238:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_18), int32(0))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L25
	} else {
		goto L319
	}
L239:
	;
	v1187 = int32(1)
	goto L313
L240:
	;
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946)+1)))
	if v968 != 0 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v969 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v946)+2)))
	if base.Ui32(int32(4086)) < base.Ui32((v969-int32(1))&int32(_a_F_SysLoggerMain_19)) {
		goto L239
	} else {
		goto L242
	}
L242:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v946)+4))
	if v976 == int32(0) {
		goto L239
	} else {
		goto L243
	}
L243:
	;
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v946)+8)))
	v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979&int32(112))+uint32(_c_F_SysLoggerMain[37]))))
	if v982 != int32(1) {
		goto L239
	} else {
		goto L244
	}
L244:
	;
	v986 = v969 + int32(9)
	if base.Ui32(v945) < base.Ui32(v986) {
		v1279 = v945
		v1280 = v946
		goto L236
	} else {
		goto L245
	}
L245:
	;
	if v979&int32(16) != 0 {
		v998 = int32(1)
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v999 = int32(0)
	v1001 = base.I32_rem_s(v976, int32(256))
	v1003 = v1001 << (uint(int32(2)) % 32)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+uint32(_c_F_SysLoggerMain[38])))
	if v1006 != 0 {
		goto L253
	} else {
		goto L254
	}
L247:
	;
	if v979&int32(32) != 0 {
		v998 = int32(8)
		goto L246
	} else {
		goto L248
	}
L248:
	;
	if v979&int32(64) != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v997 = int32(16)
	goto L251
L250:
	;
	v997 = v952
	goto L251
L251:
	;
	v998 = v997
	goto L246
L252:
	;
	if v979&int32(1) == int32(0) {
		goto L264
	} else {
		goto L265
	}
L253:
	;
	v1007 = int32(0)
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+4))
	if v1008 <= v1007 {
		v1074 = v1007
		v1076 = v999
		goto L252
	} else {
		goto L256
	}
L254:
	;
	v1052 = v999
	goto L255
L255:
	;
	v1074 = int32(0)
	v1076 = v1052
	goto L252
L256:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+12))
	v1018 = v999
	v1019 = int32(0)
	goto L257
L257:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1011+v1019<<(uint(int32(2))%32))))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1039)))
	if v1040 == v976 {
		v1074 = v1039
		v1076 = v1018
		goto L252
	} else {
		goto L259
	}
L258:
	;
	v1052 = v1043
	goto L255
L259:
	;
	if v1018|v1040 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1043 = v1018
	goto L262
L261:
	;
	v1043 = v1039
	goto L262
L262:
	;
	v1045 = v1019 + int32(1)
	if v1008 != v1045 {
		v1018 = v1043
		v1019 = v1045
		goto L257
	} else {
		goto L263
	}
L263:
	;
	goto L258
L264:
	;
	if v1074 != 0 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	if v1074 != 0 {
		goto L278
	} else {
		goto L279
	}
L267:
	;
	F_appendBinaryStringInfo(m, v1074+int32(4), v946+int32(9), v969)
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L25
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	if v1076 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1251 = v986
	v1259 = v998
	goto L237
L271:
	;
	v1107 = F_palloc(m, int32(20))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L25
	} else {
		goto L274
	}
L272:
	;
	v1112 = v1076
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1112))) = v976
	v1115 = v1112 + int32(4)
	F_initStringInfo(m, v1115)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L25
	} else {
		goto L276
	}
L274:
	;
	v1109 = F_lappend(m, v1006, v1107)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L25
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1003)+uint32(_c_F_SysLoggerMain[38]))) = v1109
	v1112 = v1107
	goto L273
L276:
	;
	F_appendBinaryStringInfo(m, v1115, v946+int32(9), v969)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L25
	} else {
		goto L277
	}
L277:
	;
	v1251 = v986
	v1259 = v998
	goto L237
L278:
	;
	F_appendBinaryStringInfo(m, v1074+int32(4), v946+int32(9), v969)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L25
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	if v998&int32(8) != 0 {
		goto L300
	} else {
		goto L301
	}
L281:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+8))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+4))
	if v998&int32(8) != 0 {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	v1148 = F_fwrite(m, v1129, int32(1), v1128, v1145)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L25
	} else {
		goto L293
	}
L283:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v1133 != 0 {
		v1145 = v1133
		goto L282
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	v1138 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	if v1136 != 0 {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	goto L285
L287:
	;
	v1139 = v1136
	goto L289
L288:
	;
	v1139 = v1138
	goto L289
L289:
	;
	if int32(base.Ui32(v998&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1144 = v1139
	goto L292
L291:
	;
	v1144 = v1138
	goto L292
L292:
	;
	v1145 = v1144
	goto L282
L293:
	;
	if v1148 != v1128 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_18), int32(0))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L25
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1074))) = int32(0)
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1074)+4))
	F_pfree(m, v1157)
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L25
	} else {
		goto L298
	}
L297:
	;
	goto L296
L298:
	;
	v1251 = v986
	v1259 = v998
	goto L237
L299:
	;
	v1180 = F_fwrite(m, v946+int32(9), int32(1), v969, v1176)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L25
	} else {
		goto L310
	}
L300:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v1163 != 0 {
		v1176 = v1163
		goto L299
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	v1168 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	if v1166 != 0 {
		goto L304
	} else {
		goto L305
	}
L303:
	;
	goto L302
L304:
	;
	v1169 = v1166
	goto L306
L305:
	;
	v1169 = v1168
	goto L306
L306:
	;
	if int32(base.Ui32(v998&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1174 = v1169
	goto L309
L308:
	;
	v1174 = v1168
	goto L309
L309:
	;
	v1176 = v1174
	goto L299
L310:
	;
	if v1180 == v969 {
		v1251 = v986
		v1259 = v998
		goto L237
	} else {
		goto L311
	}
L311:
	;
	v1224 = v986
	v1232 = v998
	goto L238
L312:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v1221 = F_fwrite(m, v946, int32(1), v1217, v1220)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L25
	} else {
		goto L317
	}
L313:
	;
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187+v946))))
	if v1211 == int32(0) {
		v1217 = v1187
		goto L312
	} else {
		goto L315
	}
L314:
	;
	v1217 = v945
	goto L312
L315:
	;
	v1215 = v1187 + int32(1)
	if v1215 != v945 {
		v1187 = v1215
		goto L313
	} else {
		goto L316
	}
L316:
	;
	goto L314
L317:
	;
	if v1221 == v1217 {
		v1251 = v1217
		v1259 = v952
		goto L237
	} else {
		goto L318
	}
L318:
	;
	v1224 = v1217
	v1232 = v952
	goto L238
L319:
	;
	v1251 = v1224
	v1259 = v1232
	goto L237
L320:
	;
	v1279 = v1275
	v1280 = v1274
	goto L236
L321:
	;
	base.MemoryCopy(m, v1304, v1280, v1279)
	v535 = v1279
	v544 = v723
	v545 = v724
	v546 = v725
	v556 = v748
	goto L114
L322:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1321<<(uint(int32(2))%32))+uint32(_c_F_SysLoggerMain[38])))
	if v1340 == int32(0) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1425 = int32(0)
	if v535 <= v1425 {
		v1441 = v1425
		goto L209
	} else {
		goto L340
	}
L324:
	;
	v1422 = v1321 + int32(1)
	if v1422 != int32(256) {
		v1321 = v1422
		goto L322
	} else {
		goto L339
	}
L325:
	;
	v1343 = int32(0)
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1340)+4))
	if v1344 <= v1343 {
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v1348 = v1344
	v1349 = v1343
	goto L327
L327:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1340)+12))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1370+v1349<<(uint(int32(2))%32))))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1374)))
	if v1375 != 0 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	goto L324
L329:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+4))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+8))
	v1380 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v1381 = F_fwrite(m, v1376, int32(1), v1378, v1380)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L25
	} else {
		goto L332
	}
L330:
	;
	v1394 = v1348
	goto L331
L331:
	;
	v1396 = v1349 + int32(1)
	if v1396 < v1394 {
		v1348 = v1394
		v1349 = v1396
		goto L327
	} else {
		goto L338
	}
L332:
	;
	if v1381 != v1378 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_18), int32(0))
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L25
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1374))) = int32(0)
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+4))
	F_pfree(m, v1390)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L25
	} else {
		goto L337
	}
L336:
	;
	goto L335
L337:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1340)+4))
	v1394 = v1393
	goto L331
L338:
	;
	goto L328
L339:
	;
	goto L323
L340:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v1433 = F_fwrite(m, v26+int32(48), int32(1), v535, v1432)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L25
	} else {
		goto L341
	}
L341:
	;
	if v1433 == v535 {
		v1441 = v1425
		goto L209
	} else {
		goto L342
	}
L342:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_18), int32(0))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L25
	} else {
		goto L343
	}
L343:
	;
	v1441 = v1425
	goto L209
L344:
	;
	goto L115
L345:
	;
	if v1469 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	F_errmsg_internal(m, int32(_a_F_SysLoggerMain_20), int32(0))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L25
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L25
	} else {
		goto L351
	}
L349:
	;
	F_errfinish(m, int32(_a_F_SysLoggerMain_16), int32(575), int32(_a_F_SysLoggerMain_17))
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L25
	} else {
		goto L350
	}
L350:
	;
	goto L348
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
