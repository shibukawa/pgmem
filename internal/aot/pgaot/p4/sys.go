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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_SearchSysCacheCopyAttName[0]))
	v8 = F_SearchCatCache2(m, v5, base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v13)+91)))
			if v15 == int32(0) {
				v18 = F_heap_copytuple(m, v8)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = v18
					F_ReleaseCatCache(m, v8)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v24 = v20
						return v24
					}
				}
			} else {
				v20 = v3
				F_ReleaseCatCache(m, v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v24 = v20
					return v24
				}
			}
		} else {
			v24 = v3
			return v24
		}
	}
}
func F_SysCacheGetAttrNotNull(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_SysCacheGetAttr(m, l0, l1, l2, v9+int32(15))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
		if v17 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(5))%32))+uint32(_c_F_SysCacheGetAttrNotNull[0])))
				v29 = F_get_rel_name(m, v28)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SysCacheGetAttrNotNull[1])))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v36 + v37<<(uint(int32(3))%32) + l2*int32(100) - int32(68)
					F_errmsg_internal(m, int32(_a_F_SysCacheGetAttrNotNull_0), v9)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_SysCacheGetAttrNotNull_1), int32(639), int32(_a_F_SysCacheGetAttrNotNull_2))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(16)
			return v13
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
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v545 int64
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v554 int64
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int64
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int64
	_ = v580
	var v583 int64
	_ = v583
	var v585 int64
	_ = v585
	var v587 int64
	_ = v587
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v637 int64
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v766 int32
	_ = v766
	var v770 int64
	_ = v770
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int64
	_ = v779
	var v782 int64
	_ = v782
	var v784 int64
	_ = v784
	var v786 int64
	_ = v786
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v824 int64
	_ = v824
	var v826 int64
	_ = v826
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int64
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v848 int64
	_ = v848
	var v849 int32
	_ = v849
	var v851 int64
	_ = v851
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v864 int64
	_ = v864
	var v865 int32
	_ = v865
	var v867 int64
	_ = v867
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v881 int64
	_ = v881
	var v882 int32
	_ = v882
	var v884 int64
	_ = v884
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int64
	_ = v906
	var v907 int64
	_ = v907
	var v908 int64
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v936 int64
	_ = v936
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int64
	_ = v945
	var v948 int64
	_ = v948
	var v950 int64
	_ = v950
	var v952 int64
	_ = v952
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v968 int64
	_ = v968
	var v970 int64
	_ = v970
	var v971 int64
	_ = v971
	var v974 int64
	_ = v974
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1140 int32
	_ = v1140
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1320 int32
	_ = v1320
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1347 int32
	_ = v1347
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1401 int32
	_ = v1401
	var v1409 int32
	_ = v1409
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1458 int32
	_ = v1458
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1562 int32
	_ = v1562
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
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
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[2])) = uint8(v80)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[3])) = v78
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[4]))
	if v85 != 0 {
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
	F_MemoryContextDelete(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[5]))
	goto L28
L25:
	;
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[4])) = int32(0)
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
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[7]))
	v98 = F_GetBackendTypeDesc(m, v97)
	mBase = m.M
	goto L30
L30:
	;
	goto L27
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[8]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[8])) = int32(-1)
	v134 = m.G0
	v136 = v134 - int32(32)
	m.G0 = v136
	v139 = int32(967)
	switch v139 {
	case 0, 2:
		goto L38
	default:
		goto L39
	}
L37:
	;
	v174 = int32(0)
	v176 = m.G0
	v178 = v176 - int32(32)
	m.G0 = v178
	switch v174 {
	case 0, 2:
		goto L48
	default:
		goto L49
	}
L38:
	;
	F_sigemptyset(m, v136+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v136)+24)) = int32(268435456)
	switch v139 {
	case 0:
		goto L43
	default:
		goto L41
	case 2:
		goto L42
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[9])) = int32(965)
	goto L38
L40:
	;
	goto L45
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+12)) = int32(_a_F_SysLoggerMain_3)
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+12)) = int32(0)
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+12)) = int32(-2)
	goto L40
L45:
	;
	goto L46
L46:
	;
	v168 = F___sigaction(m, int32(1), v136+int32(12), int32(0))
	mBase = m.M
	m.G0 = v136 + int32(32)
	goto L37
L47:
	;
	v216 = int32(0)
	v218 = m.G0
	v220 = v218 - int32(32)
	m.G0 = v220
	switch v216 {
	case 0, 2:
		goto L58
	default:
		goto L59
	}
L48:
	;
	F_sigemptyset(m, v178+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v178)+24)) = int32(268435456)
	switch v174 {
	case 0:
		goto L53
	default:
		goto L51
	case 2:
		goto L52
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[10])) = int32(-2)
	goto L48
L50:
	;
	goto L55
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+12)) = int32(_a_F_SysLoggerMain_3)
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+12)) = int32(0)
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+12)) = int32(-2)
	goto L50
L55:
	;
	goto L56
L56:
	;
	v210 = F___sigaction(m, int32(2), v178+int32(12), int32(0))
	mBase = m.M
	m.G0 = v178 + int32(32)
	goto L47
L57:
	;
	v258 = int32(0)
	v260 = m.G0
	v262 = v260 - int32(32)
	m.G0 = v262
	switch v258 {
	case 0, 2:
		goto L68
	default:
		goto L69
	}
L58:
	;
	F_sigemptyset(m, v220+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = int32(268435456)
	switch v216 {
	case 0:
		goto L63
	default:
		goto L61
	case 2:
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[11])) = int32(-2)
	goto L58
L60:
	;
	goto L65
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = int32(_a_F_SysLoggerMain_3)
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = int32(0)
	goto L60
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+12)) = int32(-2)
	goto L60
L65:
	;
	goto L66
L66:
	;
	v252 = F___sigaction(m, int32(15), v220+int32(12), int32(0))
	mBase = m.M
	m.G0 = v220 + int32(32)
	goto L57
L67:
	;
	v300 = int32(0)
	v302 = m.G0
	v304 = v302 - int32(32)
	m.G0 = v304
	switch v300 {
	case 0, 2:
		goto L78
	default:
		goto L79
	}
L68:
	;
	F_sigemptyset(m, v262+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v262)+24)) = int32(268435456)
	switch v258 {
	case 0:
		goto L73
	default:
		goto L71
	case 2:
		goto L72
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[12])) = int32(-2)
	goto L68
L70:
	;
	goto L75
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v262)+12)) = int32(_a_F_SysLoggerMain_3)
	goto L70
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+12)) = int32(0)
	goto L70
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+12)) = int32(-2)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v294 = F___sigaction(m, int32(3), v262+int32(12), int32(0))
	mBase = m.M
	m.G0 = v262 + int32(32)
	goto L67
L77:
	;
	v342 = int32(0)
	v344 = m.G0
	v346 = v344 - int32(32)
	m.G0 = v346
	switch v342 {
	case 0, 2:
		goto L88
	default:
		goto L89
	}
L78:
	;
	F_sigemptyset(m, v304+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v304)+24)) = int32(268435456)
	switch v300 {
	case 0:
		goto L83
	default:
		goto L81
	case 2:
		goto L82
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[13])) = int32(-2)
	goto L78
L80:
	;
	goto L85
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v304)+12)) = int32(_a_F_SysLoggerMain_3)
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+12)) = int32(0)
	goto L80
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+12)) = int32(-2)
	goto L80
L85:
	;
	goto L86
L86:
	;
	v336 = F___sigaction(m, int32(14), v304+int32(12), int32(0))
	mBase = m.M
	m.G0 = v304 + int32(32)
	goto L77
L87:
	;
	v386 = m.G0
	v388 = v386 - int32(32)
	m.G0 = v388
	v391 = int32(1028)
	switch v391 {
	case 0, 2:
		goto L98
	default:
		goto L99
	}
L88:
	;
	F_sigemptyset(m, v346+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v346)+24)) = int32(268435456)
	switch v342 {
	case 0:
		goto L93
	default:
		goto L91
	case 2:
		goto L92
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[14])) = int32(-2)
	goto L88
L90:
	;
	goto L95
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v346)+12)) = int32(_a_F_SysLoggerMain_3)
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+12)) = int32(0)
	goto L90
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+12)) = int32(-2)
	goto L90
L95:
	;
	goto L96
L96:
	;
	v378 = F___sigaction(m, int32(13), v346+int32(12), int32(0))
	mBase = m.M
	m.G0 = v346 + int32(32)
	goto L87
L97:
	;
	v426 = int32(0)
	v428 = m.G0
	v430 = v428 - int32(32)
	m.G0 = v430
	switch v426 {
	case 0, 2:
		goto L108
	default:
		goto L109
	}
L98:
	;
	F_sigemptyset(m, v388+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v388)+24)) = int32(268435456)
	switch v391 {
	case 0:
		goto L103
	default:
		goto L101
	case 2:
		goto L102
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[15])) = int32(1026)
	goto L98
L100:
	;
	goto L105
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v388)+12)) = int32(_a_F_SysLoggerMain_3)
	goto L100
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388)+12)) = int32(0)
	goto L100
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388)+12)) = int32(-2)
	goto L100
L105:
	;
	goto L106
L106:
	;
	v420 = F___sigaction(m, int32(10), v388+int32(12), int32(0))
	mBase = m.M
	m.G0 = v388 + int32(32)
	goto L97
L107:
	;
	v470 = m.G0
	v472 = v470 - int32(32)
	m.G0 = v472
	v474 = int32(2)
	switch v474 {
	case 0, 2:
		goto L118
	default:
		goto L119
	}
L108:
	;
	F_sigemptyset(m, v430+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v430)+24)) = int32(268435456)
	switch v426 {
	case 0:
		goto L113
	default:
		goto L111
	case 2:
		goto L112
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[16])) = int32(-2)
	goto L108
L110:
	;
	goto L115
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v430)+12)) = int32(_a_F_SysLoggerMain_3)
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430)+12)) = int32(0)
	goto L110
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430)+12)) = int32(-2)
	goto L110
L115:
	;
	goto L116
L116:
	;
	v462 = F___sigaction(m, int32(12), v430+int32(12), int32(0))
	mBase = m.M
	m.G0 = v430 + int32(32)
	goto L107
L117:
	;
	F_pgmem_sigprocmask(m, int32(_a_F_SysLoggerMain_4), int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L25
	} else {
		goto L127
	}
L118:
	;
	F_sigemptyset(m, v472+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v472)+24)) = int32(268435456)
	switch v474 {
	case 0:
		goto L123
	default:
		goto L121
	case 2:
		goto L122
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[17])) = int32(0)
	goto L118
L120:
	;
	goto L124
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v472)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v472)+12)) = int32(_a_F_SysLoggerMain_3)
	v497 = int32(268435461)
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v472)+12)) = int32(0)
	v497 = int32(268435457)
	goto L120
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v472)+12)) = int32(-2)
	v497 = int32(268435457)
	goto L120
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v472)+24)) = v497
	goto L126
L126:
	;
	v504 = F___sigaction(m, int32(17), v472+int32(12), int32(0))
	mBase = m.M
	m.G0 = v472 + int32(32)
	goto L117
L127:
	;
	v513 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v513
	v516 = F_palloc(m, int32(1024))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L25
	} else {
		goto L128
	}
L128:
	;
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v519
	v523 = F_pg_snprintf(m, v516, int32(1024), int32(_a_F_SysLoggerMain_5), v26)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L25
	} else {
		goto L129
	}
L129:
	;
	v525 = F_strlen(m, v516)
	mBase = m.M
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[21]))
	v535 = F_pg_localtime(m, v26+int32(48), v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L25
	} else {
		goto L130
	}
L130:
	;
	v537 = F_pg_strftime(m, v516+v525, int32(1024)-v525, v530, v535)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L25
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[22])) = v516
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v542 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v545 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	v547 = F_logfile_getname(m, v545, int32(_a_F_SysLoggerMain_6))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L25
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[3]))
	if v551 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[23])) = v547
	goto L134
L136:
	;
	v554 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	v556 = F_logfile_getname(m, v554, int32(_a_F_SysLoggerMain_7))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L25
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v561 = F_pstrdup(m, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L25
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24])) = v556
	goto L138
L140:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v565 = F_pstrdup(m, v564)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L25
	} else {
		goto L141
	}
L141:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	if int32(0) < v568 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v571 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v571
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[21]))
	v577 = F_pg_localtime(m, v26+int32(48), v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L25
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L25
	} else {
		goto L146
	}
L145:
	;
	v580 = *(*int64)(unsafe.Add(mBase, uint32(v26)+48))
	v583 = base.I64_extend_i32_s(v568 * int32(60))
	v585 = int64(*(*int32)(unsafe.Add(mBase, uint32(v577)+36)))
	v587 = base.I64_rem_s(v580+v585, v583)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[26])) = v580 + v583 - v587
	goto L144
L146:
	;
	v596 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[27])) = v596
	v600 = F_CreateWaitEventSet(m, v596, int32(2))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L25
	} else {
		goto L147
	}
L147:
	;
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[28]))
	F_AddWaitEventToSet(m, v600, int32(1), int32(-1), v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L25
	} else {
		goto L148
	}
L148:
	;
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[29]))
	F_AddWaitEventToSet(m, v600, int32(2), v610, int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L25
	} else {
		goto L149
	}
L149:
	;
	v616 = int32(0)
	v625 = v568
	v626 = v561
	v627 = v565
	v637 = v92
	goto L150
L150:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[28]))
	v640 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v639))) = v640
	v645 = base.AtomicRmwOr32(m, v640, int32(_a_F_SysLoggerMain_8), v640)
	goto L152
L151:
	;
	v1557 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L25
	} else {
		goto L384
	}
L152:
	;
	v647 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30]))
	if v647 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L25
	} else {
		goto L156
	}
L154:
	;
	v808 = v625
	v809 = v626
	v810 = v627
	goto L155
L155:
	;
	v813 = int32(0)
	v815 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	if v815 <= v813 {
		goto L200
	} else {
		goto L201
	}
L156:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v655))))
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	if base.B2i32(v658 == int32(0))|base.B2i32(v658 != v661) != 0 {
		v679 = v658
		v680 = v661
		goto L158
	} else {
		goto L159
	}
L157:
	;
	if v679-v680 != 0 {
		goto L164
	} else {
		goto L165
	}
L158:
	;
	goto L157
L159:
	;
	v664 = v655
	v665 = v626
	goto L160
L160:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665)+1)))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664)+1)))
	if v669 == int32(0) {
		v679 = v669
		v680 = v668
		goto L158
	} else {
		goto L162
	}
L161:
	;
	v679 = v669
	v680 = v668
	goto L158
L162:
	;
	v672 = int32(1)
	if v669 == v668 {
		v664 = v664 + v672
		v665 = v665 + v672
		goto L160
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	F_pfree(m, v626)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L25
	} else {
		goto L167
	}
L165:
	;
	v696 = v626
	goto L166
L166:
	;
	v698 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698))))
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
	if base.B2i32(v701 == int32(0))|base.B2i32(v701 != v704) != 0 {
		v722 = v701
		v723 = v704
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v686 = F_pstrdup(m, v685)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L25
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = int32(1)
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[32]))
	v695 = F_mkdir(m, v692, v694)
	mBase = m.M
	goto L169
L169:
	;
	v696 = v686
	goto L166
L170:
	;
	if v722-v723 != 0 {
		goto L177
	} else {
		goto L178
	}
L171:
	;
	goto L170
L172:
	;
	v707 = v698
	v708 = v627
	goto L173
L173:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708)+1)))
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707)+1)))
	if v712 == int32(0) {
		v722 = v712
		v723 = v711
		goto L171
	} else {
		goto L175
	}
L174:
	;
	v722 = v712
	v723 = v711
	goto L171
L175:
	;
	v715 = int32(1)
	if v712 == v711 {
		v707 = v707 + v715
		v708 = v708 + v715
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	F_pfree(m, v627)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L25
	} else {
		goto L180
	}
L178:
	;
	v734 = v627
	goto L179
L179:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[33]))
	v739 = int32(0)
	v742 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if base.B2i32(v736&int32(8) == v739)^base.B2i32(v742 != v739) == v739 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v729 = F_pstrdup(m, v728)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L25
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = int32(1)
	v734 = v729
	goto L179
L182:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = int32(1)
	goto L184
L183:
	;
	goto L184
L184:
	;
	v753 = int32(0)
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[3]))
	if base.B2i32(v736&int32(16) == v753)^base.B2i32(v756 != v753) == v753 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = int32(1)
	goto L187
L186:
	;
	goto L187
L187:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	if v766 != v625 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	if int32(0) < v766 {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	v793 = v625
	goto L190
L190:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	if v797 != 0 {
		goto L195
	} else {
		goto L196
	}
L191:
	;
	v770 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v770
	v775 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[21]))
	v776 = F_pg_localtime(m, v26+int32(32), v775)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L25
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v793 = v766
	goto L190
L194:
	;
	v779 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v782 = base.I64_extend_i32_s(v766 * int32(60))
	v784 = int64(*(*int32)(unsafe.Add(mBase, uint32(v776)+36)))
	v786 = base.I64_rem_s(v779+v784, v782)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[26])) = v779 + v782 - v786
	goto L193
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = int32(1)
	v802 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])) = uint8(v802)
	goto L197
L196:
	;
	goto L197
L197:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L25
	} else {
		goto L198
	}
L198:
	;
	v808 = v793
	v809 = v696
	v810 = v734
	goto L155
L199:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31]))
	v841 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[35]))
	if v835&int32(1)|(v839|base.B2i32(v841 <= int32(0))) != 0 {
		v893 = v813
		goto L205
	} else {
		goto L206
	}
L200:
	;
	v832 = int32(0)
	v833 = v637
	goto L199
L201:
	;
	goto L202
L202:
	;
	v819 = int32(0)
	v821 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	if v821&int32(1) != 0 {
		v832 = v819
		v833 = v637
		goto L199
	} else {
		goto L203
	}
L203:
	;
	v824 = F_time(m)
	mBase = m.M
	v826 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[26]))
	if v824 < v826 {
		v832 = v819
		v833 = v824
		goto L199
	} else {
		goto L204
	}
L204:
	;
	v828 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = v828
	v832 = v828
	v833 = v824
	goto L199
L205:
	;
	v896 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31]))
	if v896 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L206:
	;
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v848 = F___ftello_unlocked(m, v847)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L25
	} else {
		goto L207
	}
L207:
	;
	v851 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[35])))
	if v851<<(uint(int64(10))%64) <= v848 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v856 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = v856
	v859 = v856
	goto L210
L209:
	;
	v859 = v813
	goto L210
L210:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v861 == int32(0) {
		v876 = v859
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[3]))
	if v878 == int32(0) {
		v893 = v876
		goto L205
	} else {
		goto L215
	}
L212:
	;
	v864 = F___ftello_unlocked(m, v861)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L25
	} else {
		goto L213
	}
L213:
	;
	v867 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[35])))
	if v864 < v867<<(uint(int64(10))%64) {
		v876 = v859
		goto L211
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = int32(1)
	v876 = v859 | int32(8)
	goto L211
L215:
	;
	v881 = F___ftello_unlocked(m, v878)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L25
	} else {
		goto L216
	}
L216:
	;
	v884 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[35])))
	if v881 < v884<<(uint(int64(10))%64) {
		v893 = v876
		goto L205
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = int32(1)
	v893 = v876 | int32(16)
	goto L205
L218:
	;
	v959 = int32(-1)
	v961 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	if v961 <= int32(0) {
		v982 = v959
		goto L239
	} else {
		goto L240
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31])) = int32(0)
	if v893 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v903 = v893
	goto L222
L221:
	;
	v903 = int32(25)
	goto L222
L222:
	;
	if v832 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v904 = v893
	goto L225
L224:
	;
	v904 = v903
	goto L225
L225:
	;
	if v832 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v912 = F_logfile_rotate_dest(m, v832, v904, v908, int32(1), int32(_a_F_SysLoggerMain_9), int32(_a_F_SysLoggerMain_10))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L25
	} else {
		goto L230
	}
L227:
	;
	v906 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[26]))
	v908 = v906
	goto L226
L228:
	;
	goto L229
L229:
	;
	v907 = F_time(m)
	mBase = m.M
	v908 = v907
	goto L226
L230:
	;
	if v912 == int32(0) {
		goto L218
	} else {
		goto L231
	}
L231:
	;
	v919 = F_logfile_rotate_dest(m, v832, v904, v908, int32(8), int32(_a_F_SysLoggerMain_11), int32(_a_F_SysLoggerMain_12))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L25
	} else {
		goto L232
	}
L232:
	;
	if v919 == int32(0) {
		goto L218
	} else {
		goto L233
	}
L233:
	;
	v926 = F_logfile_rotate_dest(m, v832, v904, v908, int32(16), int32(_a_F_SysLoggerMain_13), int32(_a_F_SysLoggerMain_14))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L25
	} else {
		goto L234
	}
L234:
	;
	if v926 == int32(0) {
		goto L218
	} else {
		goto L235
	}
L235:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L25
	} else {
		goto L236
	}
L236:
	;
	v933 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	if v933 <= int32(0) {
		goto L218
	} else {
		goto L237
	}
L237:
	;
	v936 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v936
	v941 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[21]))
	v942 = F_pg_localtime(m, v26+int32(32), v941)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L25
	} else {
		goto L238
	}
L238:
	;
	v945 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v948 = base.I64_extend_i32_s(v933 * int32(60))
	v950 = int64(*(*int32)(unsafe.Add(mBase, uint32(v942)+36)))
	v952 = base.I64_rem_s(v945+v950, v948)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[26])) = v945 + v948 - v952
	goto L218
L239:
	;
	v988 = F_WaitEventSetWait(m, v600, v982, v26+int32(32), int32(1), int32(83886093))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L25
	} else {
		goto L249
	}
L240:
	;
	v965 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	if v965&int32(1) != 0 {
		v982 = v959
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v968 = int64(2147483)
	v970 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[26]))
	v971 = v970 - v833
	if v968 <= v971 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v974 = v968
	goto L244
L243:
	;
	v974 = v971
	goto L244
L244:
	;
	if int64(0) < v971 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v981 = base.I32_wrap_i64(v974) * int32(1000)
	goto L247
L246:
	;
	v981 = int32(0)
	goto L247
L247:
	;
	v982 = v981
	goto L239
L248:
	;
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[36])))
	if v1552 == int32(0) {
		v616 = v1529
		v625 = v808
		v626 = v809
		v627 = v810
		v637 = v833
		goto L150
	} else {
		goto L383
	}
L249:
	;
	if v988 != int32(1) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1529 = v616
	goto L248
L251:
	;
	goto L252
L252:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v992 != int32(2) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1529 = v616
	goto L248
L254:
	;
	goto L255
L255:
	;
	v996 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[29]))
	v1002 = F_read(m, v996, v26+int32(48)+v616, int32(_a_F_SysLoggerMain_15)-v616)
	mBase = m.M
	if v1002 < int32(0) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[37]))
	if v1006 == int32(27) {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L258
L258:
	;
	if v1002 != 0 {
		goto L269
	} else {
		goto L270
	}
L259:
	;
	v1529 = v616
	goto L248
L260:
	;
	goto L261
L261:
	;
	v1011 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L25
	} else {
		goto L262
	}
L262:
	;
	if v1011 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1529 = v616
	goto L248
L264:
	;
	goto L265
L265:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L25
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(_a_F_SysLoggerMain_16), int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L25
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(_a_F_SysLoggerMain_17), int32(546), int32(_a_F_SysLoggerMain_18))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L25
	} else {
		goto L268
	}
L268:
	;
	v1529 = v616
	goto L248
L269:
	;
	v1026 = v1002 + v616
	if v1026 < int32(10) {
		v616 = v1026
		v625 = v808
		v626 = v809
		v627 = v810
		v637 = v833
		goto L150
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1401 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[36])) = uint8(v1401)
	v1409 = int32(0)
	goto L361
L272:
	;
	v1033 = v1026
	v1034 = v26 + int32(48)
	v1040 = int32(1)
	goto L273
L273:
	;
	v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034))))
	if v1055 != 0 {
		goto L278
	} else {
		goto L279
	}
L274:
	;
	v1389 = int32(0)
	v1392 = v26 + int32(48)
	if base.B2i32(v1367 == v1389)|(base.B2i32(v1368 == v1392)|base.B2i32(v1367 <= v1389)) != 0 {
		v616 = v1367
		v625 = v808
		v626 = v809
		v627 = v810
		v637 = v833
		goto L150
	} else {
		goto L360
	}
L275:
	;
	goto L274
L276:
	;
	v1362 = v1339 + v1034
	v1363 = v1033 - v1339
	if int32(9) < v1363 {
		v1033 = v1363
		v1034 = v1362
		v1040 = v1347
		goto L273
	} else {
		goto L359
	}
L277:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_19), int32(0))
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L25
	} else {
		goto L358
	}
L278:
	;
	v1275 = int32(1)
	goto L352
L279:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+1)))
	if v1056 != 0 {
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v1057 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1034)+2)))
	if base.Ui32(int32(4086)) < base.Ui32((v1057-int32(1))&int32(_a_F_SysLoggerMain_20)) {
		goto L278
	} else {
		goto L281
	}
L281:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1034)+4))
	if v1064 == int32(0) {
		goto L278
	} else {
		goto L282
	}
L282:
	;
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+8)))
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067&int32(112))+uint32(_c_F_SysLoggerMain[38]))))
	if v1070 != int32(1) {
		goto L278
	} else {
		goto L283
	}
L283:
	;
	v1074 = v1057 + int32(9)
	if base.Ui32(v1033) < base.Ui32(v1074) {
		v1367 = v1033
		v1368 = v1034
		goto L275
	} else {
		goto L284
	}
L284:
	;
	if v1067&int32(16) != 0 {
		v1086 = int32(1)
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1087 = int32(0)
	v1089 = base.I32_rem_s(v1064, int32(256))
	v1091 = v1089 << (uint(int32(2)) % 32)
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+uint32(_c_F_SysLoggerMain[39])))
	if v1094 != 0 {
		goto L292
	} else {
		goto L293
	}
L286:
	;
	if v1067&int32(32) != 0 {
		v1086 = int32(8)
		goto L285
	} else {
		goto L287
	}
L287:
	;
	if v1067&int32(64) != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1085 = int32(16)
	goto L290
L289:
	;
	v1085 = v1040
	goto L290
L290:
	;
	v1086 = v1085
	goto L285
L291:
	;
	if v1067&int32(1) == int32(0) {
		goto L303
	} else {
		goto L304
	}
L292:
	;
	v1095 = int32(0)
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+4))
	if v1096 <= v1095 {
		v1162 = v1095
		v1164 = v1087
		goto L291
	} else {
		goto L295
	}
L293:
	;
	v1140 = v1087
	goto L294
L294:
	;
	v1162 = int32(0)
	v1164 = v1140
	goto L291
L295:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+12))
	v1106 = v1087
	v1107 = int32(0)
	goto L296
L296:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1099+v1107<<(uint(int32(2))%32))))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1127)))
	if v1128 == v1064 {
		v1162 = v1127
		v1164 = v1106
		goto L291
	} else {
		goto L298
	}
L297:
	;
	v1140 = v1131
	goto L294
L298:
	;
	if v1106|v1128 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1131 = v1106
	goto L301
L300:
	;
	v1131 = v1127
	goto L301
L301:
	;
	v1133 = v1107 + int32(1)
	if v1096 != v1133 {
		v1106 = v1131
		v1107 = v1133
		goto L296
	} else {
		goto L302
	}
L302:
	;
	goto L297
L303:
	;
	if v1162 != 0 {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	goto L305
L305:
	;
	if v1162 != 0 {
		goto L317
	} else {
		goto L318
	}
L306:
	;
	F_appendBinaryStringInfo(m, v1162+int32(4), v1034+int32(9), v1057)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L25
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	if v1164 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	v1339 = v1074
	v1347 = v1086
	goto L276
L310:
	;
	v1195 = F_palloc(m, int32(20))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L25
	} else {
		goto L313
	}
L311:
	;
	v1200 = v1164
	goto L312
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1200))) = v1064
	v1203 = v1200 + int32(4)
	F_initStringInfo(m, v1203)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L25
	} else {
		goto L315
	}
L313:
	;
	v1197 = F_lappend(m, v1094, v1195)
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L25
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1091)+uint32(_c_F_SysLoggerMain[39]))) = v1197
	v1200 = v1195
	goto L312
L315:
	;
	F_appendBinaryStringInfo(m, v1203, v1034+int32(9), v1057)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L25
	} else {
		goto L316
	}
L316:
	;
	v1339 = v1074
	v1347 = v1086
	goto L276
L317:
	;
	F_appendBinaryStringInfo(m, v1162+int32(4), v1034+int32(9), v1057)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L25
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	if v1086&int32(8) != 0 {
		goto L339
	} else {
		goto L340
	}
L320:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+8))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+4))
	if v1086&int32(8) != 0 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v1236 = F_fwrite(m, v1217, int32(1), v1216, v1233)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L25
	} else {
		goto L332
	}
L322:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v1221 != 0 {
		v1233 = v1221
		goto L321
	} else {
		goto L325
	}
L323:
	;
	goto L324
L324:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[3]))
	v1226 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	if v1224 != 0 {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	goto L324
L326:
	;
	v1227 = v1224
	goto L328
L327:
	;
	v1227 = v1226
	goto L328
L328:
	;
	if int32(base.Ui32(v1086&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1232 = v1227
	goto L331
L330:
	;
	v1232 = v1226
	goto L331
L331:
	;
	v1233 = v1232
	goto L321
L332:
	;
	if v1236 != v1216 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_19), int32(0))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L25
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1162))) = int32(0)
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+4))
	F_pfree(m, v1245)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L25
	} else {
		goto L337
	}
L336:
	;
	goto L335
L337:
	;
	v1339 = v1074
	v1347 = v1086
	goto L276
L338:
	;
	v1268 = F_fwrite(m, v1034+int32(9), int32(1), v1057, v1264)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L25
	} else {
		goto L349
	}
L339:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v1251 != 0 {
		v1264 = v1251
		goto L338
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[3]))
	v1256 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	if v1254 != 0 {
		goto L343
	} else {
		goto L344
	}
L342:
	;
	goto L341
L343:
	;
	v1257 = v1254
	goto L345
L344:
	;
	v1257 = v1256
	goto L345
L345:
	;
	if int32(base.Ui32(v1086&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1262 = v1257
	goto L348
L347:
	;
	v1262 = v1256
	goto L348
L348:
	;
	v1264 = v1262
	goto L338
L349:
	;
	if v1268 == v1057 {
		v1339 = v1074
		v1347 = v1086
		goto L276
	} else {
		goto L350
	}
L350:
	;
	v1312 = v1074
	v1320 = v1086
	goto L277
L351:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v1309 = F_fwrite(m, v1034, int32(1), v1305, v1308)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L25
	} else {
		goto L356
	}
L352:
	;
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1275+v1034))))
	if v1299 == int32(0) {
		v1305 = v1275
		goto L351
	} else {
		goto L354
	}
L353:
	;
	v1305 = v1033
	goto L351
L354:
	;
	v1303 = v1275 + int32(1)
	if v1303 != v1033 {
		v1275 = v1303
		goto L352
	} else {
		goto L355
	}
L355:
	;
	goto L353
L356:
	;
	if v1309 == v1305 {
		v1339 = v1305
		v1347 = v1040
		goto L276
	} else {
		goto L357
	}
L357:
	;
	v1312 = v1305
	v1320 = v1040
	goto L277
L358:
	;
	v1339 = v1312
	v1347 = v1320
	goto L276
L359:
	;
	v1367 = v1363
	v1368 = v1362
	goto L275
L360:
	;
	base.MemoryCopy(m, v1392, v1368, v1367)
	v616 = v1367
	v625 = v808
	v626 = v809
	v627 = v810
	v637 = v833
	goto L150
L361:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v1409<<(uint(int32(2))%32))+uint32(_c_F_SysLoggerMain[39])))
	if v1428 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	v1513 = int32(0)
	if v616 <= v1513 {
		v1529 = v1513
		goto L248
	} else {
		goto L379
	}
L363:
	;
	v1510 = v1409 + int32(1)
	if v1510 != int32(256) {
		v1409 = v1510
		goto L361
	} else {
		goto L378
	}
L364:
	;
	v1431 = int32(0)
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+4))
	if v1432 <= v1431 {
		goto L363
	} else {
		goto L365
	}
L365:
	;
	v1436 = v1432
	v1437 = v1431
	goto L366
L366:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+12))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1458+v1437<<(uint(int32(2))%32))))
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1462)))
	if v1463 != 0 {
		goto L368
	} else {
		goto L369
	}
L367:
	;
	goto L363
L368:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1462)+4))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1462)+8))
	v1468 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v1469 = F_fwrite(m, v1464, int32(1), v1466, v1468)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L25
	} else {
		goto L371
	}
L369:
	;
	v1482 = v1436
	goto L370
L370:
	;
	v1484 = v1437 + int32(1)
	if v1484 < v1482 {
		v1436 = v1482
		v1437 = v1484
		goto L366
	} else {
		goto L377
	}
L371:
	;
	if v1469 != v1466 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_19), int32(0))
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L25
	} else {
		goto L375
	}
L373:
	;
	goto L374
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1462))) = int32(0)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1462)+4))
	F_pfree(m, v1478)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L25
	} else {
		goto L376
	}
L375:
	;
	goto L374
L376:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1428)+4))
	v1482 = v1481
	goto L370
L377:
	;
	goto L367
L378:
	;
	goto L362
L379:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v1521 = F_fwrite(m, v26+int32(48), int32(1), v616, v1520)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L25
	} else {
		goto L380
	}
L380:
	;
	if v1521 == v616 {
		v1529 = v1513
		goto L248
	} else {
		goto L381
	}
L381:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_19), int32(0))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L25
	} else {
		goto L382
	}
L382:
	;
	v1529 = v1513
	goto L248
L383:
	;
	goto L151
L384:
	;
	if v1557 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	F_errmsg_internal(m, int32(_a_F_SysLoggerMain_21), int32(0))
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L25
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L25
	} else {
		goto L390
	}
L388:
	;
	F_errfinish(m, int32(_a_F_SysLoggerMain_17), int32(594), int32(_a_F_SysLoggerMain_18))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L25
	} else {
		goto L389
	}
L389:
	;
	goto L387
L390:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
