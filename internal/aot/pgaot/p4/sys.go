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
	v4 = Fn13839(m, l0, l1, int32(_a_F_SearchSysCacheCopyAttName_0))
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
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v685 int32
	_ = v685
	var v689 int64
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int64
	_ = v698
	var v701 int64
	_ = v701
	var v703 int64
	_ = v703
	var v705 int64
	_ = v705
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v743 int64
	_ = v743
	var v745 int64
	_ = v745
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int64
	_ = v752
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v767 int64
	_ = v767
	var v769 int64
	_ = v769
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v782 int64
	_ = v782
	var v784 int64
	_ = v784
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v798 int64
	_ = v798
	var v800 int64
	_ = v800
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int64
	_ = v822
	var v823 int64
	_ = v823
	var v824 int64
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int64
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int64
	_ = v861
	var v864 int64
	_ = v864
	var v866 int64
	_ = v866
	var v868 int64
	_ = v868
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v884 int64
	_ = v884
	var v886 int64
	_ = v886
	var v887 int64
	_ = v887
	var v890 int64
	_ = v890
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1191 int32
	_ = v1191
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1236 int32
	_ = v1236
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1263 int32
	_ = v1263
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1317 int32
	_ = v1317
	var v1325 int32
	_ = v1325
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1445 int32
	_ = v1445
	var v1468 int32
	_ = v1468
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
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
	v559 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v558))) = v559
	v564 = base.AtomicRmwOr32(m, v559, int32(_a_F_SysLoggerMain_8), v559)
	goto L116
L115:
	;
	v1473 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L25
	} else {
		goto L345
	}
L116:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[29]))
	if v566 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[29])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L25
	} else {
		goto L120
	}
L118:
	;
	v727 = v544
	v728 = v545
	v729 = v546
	goto L119
L119:
	;
	v732 = int32(0)
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24]))
	if v734 <= v732 {
		goto L164
	} else {
		goto L165
	}
L120:
	;
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	if base.B2i32(v577 == int32(0))|base.B2i32(v577 != v580) != 0 {
		v598 = v577
		v599 = v580
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v598-v599 != 0 {
		goto L128
	} else {
		goto L129
	}
L122:
	;
	goto L121
L123:
	;
	v583 = v574
	v584 = v545
	goto L124
L124:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584)+1)))
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+1)))
	if v588 == int32(0) {
		v598 = v588
		v599 = v587
		goto L122
	} else {
		goto L126
	}
L125:
	;
	v598 = v588
	v599 = v587
	goto L122
L126:
	;
	v591 = int32(1)
	if v588 == v587 {
		v583 = v583 + v591
		v584 = v584 + v591
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
	v602 = m.ExcPending
	if v602 != 0 {
		goto L25
	} else {
		goto L131
	}
L129:
	;
	v615 = v545
	goto L130
L130:
	;
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617))))
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546))))
	if base.B2i32(v620 == int32(0))|base.B2i32(v620 != v623) != 0 {
		v641 = v620
		v642 = v623
		goto L135
	} else {
		goto L136
	}
L131:
	;
	v604 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	v605 = F_pstrdup(m, v604)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L25
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v611 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[18]))
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[31]))
	v614 = F_mkdir(m, v611, v613)
	mBase = m.M
	goto L133
L133:
	;
	v615 = v605
	goto L130
L134:
	;
	if v641-v642 != 0 {
		goto L141
	} else {
		goto L142
	}
L135:
	;
	goto L134
L136:
	;
	v626 = v617
	v627 = v546
	goto L137
L137:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627)+1)))
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+1)))
	if v631 == int32(0) {
		v641 = v631
		v642 = v630
		goto L135
	} else {
		goto L139
	}
L138:
	;
	v641 = v631
	v642 = v630
	goto L135
L139:
	;
	v634 = int32(1)
	if v631 == v630 {
		v626 = v626 + v634
		v627 = v627 + v634
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
	v645 = m.ExcPending
	if v645 != 0 {
		goto L25
	} else {
		goto L144
	}
L142:
	;
	v653 = v546
	goto L143
L143:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[32]))
	v658 = int32(0)
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if base.B2i32(v655&int32(8) == v658)^base.B2i32(v661 != v658) == v658 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v647 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[19]))
	v648 = F_pstrdup(m, v647)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L25
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v653 = v648
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
	v672 = int32(0)
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	if base.B2i32(v655&int32(16) == v672)^base.B2i32(v675 != v672) == v672 {
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
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24]))
	if v685 != v544 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if int32(0) < v685 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v712 = v544
	goto L154
L154:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])))
	if v716 != 0 {
		goto L159
	} else {
		goto L160
	}
L155:
	;
	v689 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v689
	v694 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v695 = F_pg_localtime(m, v26+int32(32), v694)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L25
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v712 = v685
	goto L154
L158:
	;
	v698 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v701 = base.I64_extend_i32_s(v685 * int32(60))
	v703 = int64(*(*int32)(unsafe.Add(mBase, uint32(v695)+36)))
	v705 = base.I64_rem_s(v698+v703, v701)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25])) = v698 + v701 - v705
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v721 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])) = uint8(v721)
	goto L161
L160:
	;
	goto L161
L161:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L25
	} else {
		goto L162
	}
L162:
	;
	v727 = v712
	v728 = v615
	v729 = v653
	goto L119
L163:
	;
	v754 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])))
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30]))
	v760 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[34]))
	if v754&int32(1)|(v758|base.B2i32(v760 <= int32(0))) != 0 {
		v809 = v732
		goto L169
	} else {
		goto L170
	}
L164:
	;
	v751 = int32(0)
	v752 = v556
	goto L163
L165:
	;
	goto L166
L166:
	;
	v738 = int32(0)
	v740 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])))
	if v740&int32(1) != 0 {
		v751 = v738
		v752 = v556
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v743 = F_time(m)
	mBase = m.M
	v745 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	if v743 < v745 {
		v751 = v738
		v752 = v743
		goto L163
	} else {
		goto L168
	}
L168:
	;
	v747 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = v747
	v751 = v747
	v752 = v743
	goto L163
L169:
	;
	v812 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30]))
	if v812 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L170:
	;
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v767 = F___ftello_unlocked(m, v766)
	mBase = m.M
	v769 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	if v769<<(uint(int64(10))%64) <= v767 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v774 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = v774
	v777 = v774
	goto L173
L172:
	;
	v777 = v732
	goto L173
L173:
	;
	v779 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v779 == int32(0) {
		v793 = v777
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	if v795 == int32(0) {
		v809 = v793
		goto L169
	} else {
		goto L177
	}
L175:
	;
	v782 = F___ftello_unlocked(m, v779)
	mBase = m.M
	v784 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	if v782 < v784<<(uint(int64(10))%64) {
		v793 = v777
		goto L174
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v793 = v777 | int32(8)
	goto L174
L177:
	;
	v798 = F___ftello_unlocked(m, v795)
	mBase = m.M
	v800 = int64(*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[34])))
	if v798 < v800<<(uint(int64(10))%64) {
		v809 = v793
		goto L169
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(1)
	v809 = v793 | int32(16)
	goto L169
L179:
	;
	v875 = int32(-1)
	v877 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24]))
	if v877 <= int32(0) {
		v898 = v875
		goto L200
	} else {
		goto L201
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[30])) = int32(0)
	if v809 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v819 = v809
	goto L183
L182:
	;
	v819 = int32(25)
	goto L183
L183:
	;
	if v751 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v820 = v809
	goto L186
L185:
	;
	v820 = v819
	goto L186
L186:
	;
	if v751 != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v828 = F_logfile_rotate_dest(m, v751, v820, v824, int32(1), int32(_a_F_SysLoggerMain_9), int32(_a_F_SysLoggerMain_10))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L25
	} else {
		goto L191
	}
L188:
	;
	v822 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	v824 = v822
	goto L187
L189:
	;
	goto L190
L190:
	;
	v823 = F_time(m)
	mBase = m.M
	v824 = v823
	goto L187
L191:
	;
	if v828 == int32(0) {
		goto L179
	} else {
		goto L192
	}
L192:
	;
	v835 = F_logfile_rotate_dest(m, v751, v820, v824, int32(8), int32(_a_F_SysLoggerMain_11), int32(_a_F_SysLoggerMain_12))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L25
	} else {
		goto L193
	}
L193:
	;
	if v835 == int32(0) {
		goto L179
	} else {
		goto L194
	}
L194:
	;
	v842 = F_logfile_rotate_dest(m, v751, v820, v824, int32(16), int32(_a_F_SysLoggerMain_13), int32(_a_F_SysLoggerMain_14))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L25
	} else {
		goto L195
	}
L195:
	;
	if v842 == int32(0) {
		goto L179
	} else {
		goto L196
	}
L196:
	;
	F_update_metainfo_datafile(m)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L25
	} else {
		goto L197
	}
L197:
	;
	v849 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[24]))
	if v849 <= int32(0) {
		goto L179
	} else {
		goto L198
	}
L198:
	;
	v852 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = v852
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[20]))
	v858 = F_pg_localtime(m, v26+int32(32), v857)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L25
	} else {
		goto L199
	}
L199:
	;
	v861 = *(*int64)(unsafe.Add(mBase, uint32(v26)+32))
	v864 = base.I64_extend_i32_s(v849 * int32(60))
	v866 = int64(*(*int32)(unsafe.Add(mBase, uint32(v858)+36)))
	v868 = base.I64_rem_s(v861+v866, v864)
	*(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25])) = v861 + v864 - v868
	goto L179
L200:
	;
	v904 = F_WaitEventSetWait(m, v519, v898, v26+int32(32), int32(1), int32(83886093))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L25
	} else {
		goto L210
	}
L201:
	;
	v881 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[33])))
	if v881&int32(1) != 0 {
		v898 = v875
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v884 = int64(2147483)
	v886 = *(*int64)(unsafe.Add(mBase, _c_F_SysLoggerMain[25]))
	v887 = v886 - v752
	if v884 <= v887 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v890 = v884
	goto L205
L204:
	;
	v890 = v887
	goto L205
L205:
	;
	if int64(0) < v887 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v897 = base.I32_wrap_i64(v890) * int32(1000)
	goto L208
L207:
	;
	v897 = int32(0)
	goto L208
L208:
	;
	v898 = v897
	goto L200
L209:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[35])))
	if v1468 == int32(0) {
		v535 = v1445
		v544 = v727
		v545 = v728
		v546 = v729
		v556 = v752
		goto L114
	} else {
		goto L344
	}
L210:
	;
	if v904 != int32(1) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1445 = v535
	goto L209
L212:
	;
	goto L213
L213:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	if v908 != int32(2) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1445 = v535
	goto L209
L215:
	;
	goto L216
L216:
	;
	v912 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[28]))
	v918 = F_read(m, v912, v26+int32(48)+v535, int32(_a_F_SysLoggerMain_15)-v535)
	mBase = m.M
	if v918 < int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v922 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[36]))
	if v922 == int32(27) {
		goto L220
	} else {
		goto L221
	}
L218:
	;
	goto L219
L219:
	;
	if v918 != 0 {
		goto L230
	} else {
		goto L231
	}
L220:
	;
	v1445 = v535
	goto L209
L221:
	;
	goto L222
L222:
	;
	v927 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L25
	} else {
		goto L223
	}
L223:
	;
	if v927 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1445 = v535
	goto L209
L225:
	;
	goto L226
L226:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L25
	} else {
		goto L227
	}
L227:
	;
	F_errmsg(m, int32(_a_F_SysLoggerMain_16), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L25
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_SysLoggerMain_17), int32(527), int32(_a_F_SysLoggerMain_18))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L25
	} else {
		goto L229
	}
L229:
	;
	v1445 = v535
	goto L209
L230:
	;
	v942 = v918 + v535
	if v942 < int32(10) {
		v535 = v942
		v544 = v727
		v545 = v728
		v546 = v729
		v556 = v752
		goto L114
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v1317 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SysLoggerMain[35])) = uint8(v1317)
	v1325 = int32(0)
	goto L322
L233:
	;
	v949 = v942
	v950 = v26 + int32(48)
	v956 = int32(1)
	goto L234
L234:
	;
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950))))
	if v971 != 0 {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v1305 = int32(0)
	v1308 = v26 + int32(48)
	if base.B2i32(v1283 == v1305)|(base.B2i32(v1284 == v1308)|base.B2i32(v1283 <= v1305)) != 0 {
		v535 = v1283
		v544 = v727
		v545 = v728
		v546 = v729
		v556 = v752
		goto L114
	} else {
		goto L321
	}
L236:
	;
	goto L235
L237:
	;
	v1278 = v1255 + v950
	v1279 = v949 - v1255
	if int32(9) < v1279 {
		v949 = v1279
		v950 = v1278
		v956 = v1263
		goto L234
	} else {
		goto L320
	}
L238:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_19), int32(0))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L25
	} else {
		goto L319
	}
L239:
	;
	v1191 = int32(1)
	goto L313
L240:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+1)))
	if v972 != 0 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v973 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v950)+2)))
	if base.Ui32(int32(4086)) < base.Ui32((v973-int32(1))&int32(_a_F_SysLoggerMain_20)) {
		goto L239
	} else {
		goto L242
	}
L242:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v950)+4))
	if v980 == int32(0) {
		goto L239
	} else {
		goto L243
	}
L243:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+8)))
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v983&int32(112))+uint32(_c_F_SysLoggerMain[37]))))
	if v986 != int32(1) {
		goto L239
	} else {
		goto L244
	}
L244:
	;
	v990 = v973 + int32(9)
	if base.Ui32(v949) < base.Ui32(v990) {
		v1283 = v949
		v1284 = v950
		goto L236
	} else {
		goto L245
	}
L245:
	;
	if v983&int32(16) != 0 {
		v1002 = int32(1)
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1003 = int32(0)
	v1005 = base.I32_rem_s(v980, int32(256))
	v1007 = v1005 << (uint(int32(2)) % 32)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+uint32(_c_F_SysLoggerMain[38])))
	if v1010 != 0 {
		goto L253
	} else {
		goto L254
	}
L247:
	;
	if v983&int32(32) != 0 {
		v1002 = int32(8)
		goto L246
	} else {
		goto L248
	}
L248:
	;
	if v983&int32(64) != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1001 = int32(16)
	goto L251
L250:
	;
	v1001 = v956
	goto L251
L251:
	;
	v1002 = v1001
	goto L246
L252:
	;
	if v983&int32(1) == int32(0) {
		goto L264
	} else {
		goto L265
	}
L253:
	;
	v1011 = int32(0)
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+4))
	if v1012 <= v1011 {
		v1078 = v1011
		v1080 = v1003
		goto L252
	} else {
		goto L256
	}
L254:
	;
	v1056 = v1003
	goto L255
L255:
	;
	v1078 = int32(0)
	v1080 = v1056
	goto L252
L256:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1010)+12))
	v1022 = v1003
	v1023 = int32(0)
	goto L257
L257:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1015+v1023<<(uint(int32(2))%32))))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1043)))
	if v1044 == v980 {
		v1078 = v1043
		v1080 = v1022
		goto L252
	} else {
		goto L259
	}
L258:
	;
	v1056 = v1047
	goto L255
L259:
	;
	if v1022|v1044 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1047 = v1022
	goto L262
L261:
	;
	v1047 = v1043
	goto L262
L262:
	;
	v1049 = v1023 + int32(1)
	if v1012 != v1049 {
		v1022 = v1047
		v1023 = v1049
		goto L257
	} else {
		goto L263
	}
L263:
	;
	goto L258
L264:
	;
	if v1078 != 0 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	if v1078 != 0 {
		goto L278
	} else {
		goto L279
	}
L267:
	;
	F_appendBinaryStringInfo(m, v1078+int32(4), v950+int32(9), v973)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L25
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	if v1080 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1255 = v990
	v1263 = v1002
	goto L237
L271:
	;
	v1111 = F_palloc(m, int32(20))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L25
	} else {
		goto L274
	}
L272:
	;
	v1116 = v1080
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1116))) = v980
	v1119 = v1116 + int32(4)
	F_initStringInfo(m, v1119)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L25
	} else {
		goto L276
	}
L274:
	;
	v1113 = F_lappend(m, v1010, v1111)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L25
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1007)+uint32(_c_F_SysLoggerMain[38]))) = v1113
	v1116 = v1111
	goto L273
L276:
	;
	F_appendBinaryStringInfo(m, v1119, v950+int32(9), v973)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L25
	} else {
		goto L277
	}
L277:
	;
	v1255 = v990
	v1263 = v1002
	goto L237
L278:
	;
	F_appendBinaryStringInfo(m, v1078+int32(4), v950+int32(9), v973)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L25
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	if v1002&int32(8) != 0 {
		goto L300
	} else {
		goto L301
	}
L281:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+8))
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+4))
	if v1002&int32(8) != 0 {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	v1152 = F_fwrite(m, v1133, int32(1), v1132, v1149)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L25
	} else {
		goto L293
	}
L283:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v1137 != 0 {
		v1149 = v1137
		goto L282
	} else {
		goto L286
	}
L284:
	;
	goto L285
L285:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	v1142 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	if v1140 != 0 {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	goto L285
L287:
	;
	v1143 = v1140
	goto L289
L288:
	;
	v1143 = v1142
	goto L289
L289:
	;
	if int32(base.Ui32(v1002&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1148 = v1143
	goto L292
L291:
	;
	v1148 = v1142
	goto L292
L292:
	;
	v1149 = v1148
	goto L282
L293:
	;
	if v1152 != v1132 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_19), int32(0))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L25
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1078))) = int32(0)
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+4))
	F_pfree(m, v1161)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L25
	} else {
		goto L298
	}
L297:
	;
	goto L296
L298:
	;
	v1255 = v990
	v1263 = v1002
	goto L237
L299:
	;
	v1184 = F_fwrite(m, v950+int32(9), int32(1), v973, v1180)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L25
	} else {
		goto L310
	}
L300:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[1]))
	if v1167 != 0 {
		v1180 = v1167
		goto L299
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[2]))
	v1172 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	if v1170 != 0 {
		goto L304
	} else {
		goto L305
	}
L303:
	;
	goto L302
L304:
	;
	v1173 = v1170
	goto L306
L305:
	;
	v1173 = v1172
	goto L306
L306:
	;
	if int32(base.Ui32(v1002&int32(16))>>(uint(int32(4))%32)) != 0 {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1178 = v1173
	goto L309
L308:
	;
	v1178 = v1172
	goto L309
L309:
	;
	v1180 = v1178
	goto L299
L310:
	;
	if v1184 == v973 {
		v1255 = v990
		v1263 = v1002
		goto L237
	} else {
		goto L311
	}
L311:
	;
	v1228 = v990
	v1236 = v1002
	goto L238
L312:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v1225 = F_fwrite(m, v950, int32(1), v1221, v1224)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L25
	} else {
		goto L317
	}
L313:
	;
	v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191+v950))))
	if v1215 == int32(0) {
		v1221 = v1191
		goto L312
	} else {
		goto L315
	}
L314:
	;
	v1221 = v949
	goto L312
L315:
	;
	v1219 = v1191 + int32(1)
	if v1219 != v949 {
		v1191 = v1219
		goto L313
	} else {
		goto L316
	}
L316:
	;
	goto L314
L317:
	;
	if v1225 == v1221 {
		v1255 = v1221
		v1263 = v956
		goto L237
	} else {
		goto L318
	}
L318:
	;
	v1228 = v1221
	v1236 = v956
	goto L238
L319:
	;
	v1255 = v1228
	v1263 = v1236
	goto L237
L320:
	;
	v1283 = v1279
	v1284 = v1278
	goto L236
L321:
	;
	base.MemoryCopy(m, v1308, v1284, v1283)
	v535 = v1283
	v544 = v727
	v545 = v728
	v546 = v729
	v556 = v752
	goto L114
L322:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1325<<(uint(int32(2))%32))+uint32(_c_F_SysLoggerMain[38])))
	if v1344 == int32(0) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v1429 = int32(0)
	if v535 <= v1429 {
		v1445 = v1429
		goto L209
	} else {
		goto L340
	}
L324:
	;
	v1426 = v1325 + int32(1)
	if v1426 != int32(256) {
		v1325 = v1426
		goto L322
	} else {
		goto L339
	}
L325:
	;
	v1347 = int32(0)
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+4))
	if v1348 <= v1347 {
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v1352 = v1348
	v1353 = v1347
	goto L327
L327:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+12))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1374+v1353<<(uint(int32(2))%32))))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1378)))
	if v1379 != 0 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	goto L324
L329:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+4))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+8))
	v1384 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v1385 = F_fwrite(m, v1380, int32(1), v1382, v1384)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L25
	} else {
		goto L332
	}
L330:
	;
	v1398 = v1352
	goto L331
L331:
	;
	v1400 = v1353 + int32(1)
	if v1400 < v1398 {
		v1352 = v1398
		v1353 = v1400
		goto L327
	} else {
		goto L338
	}
L332:
	;
	if v1385 != v1382 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_19), int32(0))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L25
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1378))) = int32(0)
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1378)+4))
	F_pfree(m, v1394)
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L25
	} else {
		goto L337
	}
L336:
	;
	goto L335
L337:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+4))
	v1398 = v1397
	goto L331
L338:
	;
	goto L328
L339:
	;
	goto L323
L340:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, _c_F_SysLoggerMain[0]))
	v1437 = F_fwrite(m, v26+int32(48), int32(1), v535, v1436)
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L25
	} else {
		goto L341
	}
L341:
	;
	if v1437 == v535 {
		v1445 = v1429
		goto L209
	} else {
		goto L342
	}
L342:
	;
	F_write_stderr(m, int32(_a_F_SysLoggerMain_19), int32(0))
	mBase = m.M
	v1443 = m.ExcPending
	if v1443 != 0 {
		goto L25
	} else {
		goto L343
	}
L343:
	;
	v1445 = v1429
	goto L209
L344:
	;
	goto L115
L345:
	;
	if v1473 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	F_errmsg_internal(m, int32(_a_F_SysLoggerMain_21), int32(0))
	mBase = m.M
	v1478 = m.ExcPending
	if v1478 != 0 {
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
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L25
	} else {
		goto L351
	}
L349:
	;
	F_errfinish(m, int32(_a_F_SysLoggerMain_17), int32(575), int32(_a_F_SysLoggerMain_18))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
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
